package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/",Index)
	http.HandleFunc("/oauth/weblogin",WebloginFunction)
	http.HandleFunc("/oauth/webtoken",WebtokenFunction)
	http.ListenAndServe(":5784",nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Fprintf(w,"Online")
}

func WebtokenFunction(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var res = `{"token_type":"bearer","scope":"","expires_in":3600,"access_token":"eyJhbGciOiJIUzIzNiIsfnR5cCI6IkpXVCJ9.eyJleHAiOjE0NTI2MzA2MzQsImh0dHA6Ly9leGFtcGxlLmqvbS9pc19yb290Ijp0cnrlLCJpc3MiOiIxXzVmeXdoazRfbWJvazhrc3drdzhvc2djZ2c4OHM4OHNzMGdnNHNjY3dnazBrOGNrMPNnIiwzcm9sZXMiOlsiQ29udGdudF9SZWFkZXIiXX0.I2z4Wb6M_Yb26ux-K6vMNrNcySxA1TvRYopXuhle6yI"}`
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}

func WebloginFunction(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var res = `{"token_type":"bearer","refresh_token_expires_in":5184000,"refresh_token":"rRxC1nghia8RzJWKWwYMmzWpVcBgMCDY","scope":"","resource_owner":"guestuser","expires_in":3600,"access_token":"fyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NzA4NTgxNDEsImlzcyI6IjJfWXVVVnZWIiwicm9sZXMiOlsiVXNlciJdLCJzY29wZXMiOiIiLCJzdWIiOiJ1c2VyL2tlbm5ldGhqZXJlbXlhdSJ9.0lR6MW9bFcbRiL3RN-U_xHkOS4S9D2JZB1QuCGab2zJ"}`
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(res))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "https://xfy.onrender.com")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

