package main

import (
	"fmt"
	"net/http"
)

// fungsi log untuk middleware

func log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ini dari middleware Log... \n")
		fmt.Println(r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// fungsi untuk getmsiswa untuk mengambil text pada string browser

func Getsiswa(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ini dari function siswa()"))
}

func main() {
	//konfigurasi server
	server := &http.Server{
		Addr: ":9000",
	}
	// routing
	http.Handle("/", log(http.HandlerFunc(Getsiswa)))

	// jalankan server
	server.ListenAndServe()
}
