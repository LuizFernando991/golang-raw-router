package router

import (
	"context"
	"net/http"
)

type contextKey string

func AddURLParam(ctx context.Context, key, value string) context.Context {
	return context.WithValue(ctx, contextKey(key), value)
}

func GetURLParam(r *http.Request, key string) string {
	val := r.Context().Value(contextKey(key))
	if val == nil {
		return ""
	}
	return val.(string)
}

func GetUrlQuery(r *http.Request, key string) string {
	val := r.URL.Query().Get(key)

	return val
}
