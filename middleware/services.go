package middleware

import (
	"fmt"
	"net/http"
	"regexp"
)

func MACAddressMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		regex := regexp.MustCompile(`(.*)mac=([A-Za-z\d]*)(.*)`)
		if regex.MatchString(r.RequestURI) {
			//fmt.Println("MATCH ! ")
			submatches := regex.FindStringSubmatch(r.RequestURI)
			mac := submatches[2]
			fmt.Printf("Request from MAC '%s'\n", mac)
		} else {
			//fmt.Println("NO MATCH")
		}

		// #####
		// finally, serve the other handlers
		next.ServeHTTP(w, r)
	})
}

func XMLEncodingLineAddingWrapper(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\" ?>"))

		// #####
		// finally, serve the other handlers
		next.ServeHTTP(w, r)
	})
}
