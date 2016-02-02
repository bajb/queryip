package queryip

import (
	"encoding/json"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, OPTIONS")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	enc := json.NewEncoder(w)
	d := map[string]string{"ip": r.RemoteAddr, "country": r.Header.Get("X-Appengine-Country"), "region": r.Header.Get("X-Appengine-Region"), "city": r.Header.Get("X-Appengine-City"), "latlong": r.Header.Get("X-Appengine-Citylatlong")}
	enc.Encode(d)
}
