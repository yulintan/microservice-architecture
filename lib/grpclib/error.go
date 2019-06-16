package grpclib

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type errorBody struct {
	StatusCode int32    `json:"status_code"`
	Errors     []string `json:"errors"`
}

// Error converts any error to gRPC status.Error which later will be converted
// to http status by grpc-gateway.
func Error(err error) error {
	switch err.(type) {
	default:
		switch err {
		case sql.ErrNoRows:
			return status.Error(codes.NotFound, "not found")
		}
	}

	return status.Error(codes.Unknown, "internal error")
}

// errorHandler is a custom error handler. It used by grpc-gateway mux.
func ErrorHandler(ctx context.Context, _ *runtime.ServeMux, _ runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	const fallback = `{"error": "failed to marshal error message"}`

	w.Header().Del("Trailer")
	w.Header().Set("Content-Type", "application/json")

	s, ok := status.FromError(err)
	if !ok {
		s = status.New(codes.Unknown, err.Error())
	}

	st := runtime.HTTPStatusFromCode(s.Code())

	buf, merr := json.Marshal(&errorBody{
		Errors:     []string{s.Message()},
		StatusCode: int32(st),
	})
	if merr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := io.WriteString(w, fallback); err != nil {
			// clog.FromContext(ctx).Error("error", err)
		}
		return
	}

	w.WriteHeader(st)
	if _, err := w.Write(buf); err != nil {
		// clog.FromContext(ctx).Error("error", err)
	}
}
