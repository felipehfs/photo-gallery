package middleware

import (
	"context"
	"errors"
	"net/http"

	mgo "gopkg.in/mgo.v2"
)

// DbContext is a wrapper of string
type DbContext string

// WithDB setup the database connection
func WithDB(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := mgo.Dial("localhost")
		if err != nil {
			http.Error(w, "Connection failed", 500)
			return
		}
		defer session.Close()
		ctx := context.WithValue(r.Context(), DbContext("session"), session)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

// ExtractSession returns the session of the database
func ExtractSession(r *http.Request) (*mgo.Session, error) {
	session, ok := r.Context().Value(DbContext("session")).(*mgo.Session)
	if !ok {
		return nil, errors.New("connection not found")
	}
	return session, nil
}
