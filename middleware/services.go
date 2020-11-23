package middleware

import (
	"context"
	"fmt"
	"hash/crc64"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func GenerateStationID() string {
	rand.Seed(time.Now().UnixNano())
	var id uint64 = rand.Uint64()
	res := fmt.Sprintf("%02x", id)
	return res
}

func CalculateStationID(val string) string {
	table := crc64.MakeTable(crc64.ISO)
	crc := crc64.Checksum([]byte(val), table)
	//fmt.Printf("CRC64 of '%s': \n\t0x%08x\n", val, crc)
	return fmt.Sprintf("%08x", crc)
}

func MACAddressMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		regex := regexp.MustCompile(`(.*)mac=([A-Za-z\d]*)(.*)`)
		if regex.MatchString(r.RequestURI) {
			//fmt.Println("MATCH ! ")
			submatches := regex.FindStringSubmatch(r.RequestURI)
			mac := submatches[2]
			fmt.Printf("Request from MAC '%s'\n", mac)

			ctx = context.WithValue(r.Context(), "MAC", mac)

		} else {
			//fmt.Println("NO MATCH")
		}

		fmt.Printf("  Host: %s\n", r.Host)
		fmt.Printf("  Header: %s\n", r.Header)

		// #####
		// finally, serve the other handlers
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func StationIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		regex := regexp.MustCompile(`(.*)id=([A-Za-z\d-_.]*)(.*)`) // format of station ID is somehow limited
		if regex.MatchString(r.RequestURI) {
			submatches := regex.FindStringSubmatch(r.RequestURI)
			stationId := submatches[2]
			fmt.Printf("Request for Station ID '%s'\n", stationId)

			ctx = context.WithValue(r.Context(), "StationId", stationId)

		} else {
			//fmt.Println("NO MATCH")
		}

		// #####
		// finally, serve the other handlers
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func XMLEncodingLineAddingWrapper(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.RequestURI, "loginXML.asp?token=0") { // FIXME: this is actually nasty
			// quite hacky way to prevent
			// ATTENTION: RX-V receiver will REJECT the response if it contains '<?xml version="1.0" encoding="UTF-8" standalone="yes" ?>' reamble in response!
		} else {
			w.Write([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\" ?>"))
		}

		// #####
		// finally, serve the other handlers
		next.ServeHTTP(w, r)
	})
}
