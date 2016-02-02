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
	enc := json.NewEncoder(w)
	d := map[string]string{"ip": r.RemoteAddr, "country": r.Header.Get("X-Appengine-Country"), "region": r.Header.Get("X-Appengine-Region"), "city": r.Header.Get("X-Appengine-City"), "latlong": r.Header.Get("X-Appengine-Citylatlong")}
	enc.Encode(d)
}
