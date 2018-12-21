package server

import (
	"context"
	"github.com/go-chi/chi"
	"net/http"
	"strings"
)

func GopherID(r *http.Request) string {
	return ContextMap(r)["GopherID"]
}
func ThingID(r *http.Request) string {
	return ContextMap(r)["ThingID"]
}
func GopherName(r *http.Request) string {
	return ContextMap(r)["GopherName"]
}
func GopherDescription(r *http.Request) string {
	return ContextMap(r)["GopherDescription"]
}

func IsAuthenticated(r *http.Request) (authed bool) {
	ctxMap := ContextMap(r)
	if ctxMap["AccessKey"] != "" && ctxMap["SecretKey"] != "" {
		authed = true
	}
	return authed
}

func ContextMap(r *http.Request) map[string]string {
	return (r.Context().Value("ctxMap")).(map[string]string)
}

// InitialCtx allows Non-authenicated read-only requests and sets
//  	GopherID
//	  GopherName
//	  GopherDescription
func InitCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ctxMap := make(map[string]string)

		xKeys := strings.Split(r.Header.Get("X-ApiKeys"), ";")
		for x := range xKeys {
			keys := strings.Split(xKeys[x], "=")
			switch {
			case strings.ToLower(keys[0]) == "accesskey":
				ctxMap["AccessKey"] = keys[1]

			case strings.ToLower(keys[0]) == "secretkey":
				ctxMap["SecretKey"] = keys[1]
			}
		}
		ctx := context.WithValue(r.Context(), "ctxMap", ctxMap)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GopherCtx allows non-authenicated read-only (GET/HEAD) requests and sets:
//  	GopherID
//	  GopherName
//	  GopherDescription
func GopherCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !IsAuthenticated(r) {
			switch strings.ToUpper(r.Method) {
			case "GET", "HEAD":
				// ALLOW!
			default:
				// DENY ALL OTHERS!
				http.Error(w, http.StatusText(403), 403)
				return
			}
		}

		ctxMap := r.Context().Value("ctxMap").(map[string]string)
		ctxMap["GopherID"] = chi.URLParam(r, "GopherID")
		ctxMap["GopherName"] = r.FormValue("GopherName")
		ctxMap["GopherDescription"] = r.FormValue("GopherDescription")
		ctx := context.WithValue(r.Context(), "ctxMap", ctxMap)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ThingCtx requires IsAuthenticated() for ALL HTTP methods:
//  	ThingID
//	  ThingName
//	  ThingDescription
func ThingCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !IsAuthenticated(r) {
			// PS: Never /actually/ do this.. create a proper SecurityCtx and evaluates the uri etc. :-)
			http.Error(w, http.StatusText(403), 403)
			return
		}

		ctxMap := r.Context().Value("ctxMap").(map[string]string)
		ctxMap["ThingID"] = chi.URLParam(r, "ThingID")
		ctxMap["ThingName"] = r.FormValue("ThingName")
		ctxMap["ThingDescription"] = r.FormValue("ThingDescription")
		ctx := context.WithValue(r.Context(), "ctxMap", ctxMap)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
