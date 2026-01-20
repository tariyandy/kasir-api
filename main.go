package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Produk struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
	Stok  int    `json:"stok"`
}

var produk = []Produk{
	{ID: 1, Nama: "Indomie Godok", Harga: 3500, Stok: 200},
	{ID: 2, Nama: "Vitamin C", Harga: 2000, Stok: 20},
	{ID: 3, Nama: "nasi Goreng", Harga: 25000, Stok: 2},
}

// GET localhost:8080/api/produk{id}
func getProdukByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid produk id", http.StatusBadRequest)
		return
	}
	for _, p := range produk {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	http.Error(w, "produk belum ada", http.StatusNotFound)
}

// PUT localhost:8080/api/produk{id}
func updateProduk(w http.ResponseWriter, r *http.Request) {
	// get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	// ganti int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid produk id", http.StatusBadRequest)
		return
	}
	// get data dari request
	var updateProduk Produk
	err = json.NewDecoder(r.Body).Decode(&updateProduk)
	if err != nil {
		http.Error(w, "invalid produk id", http.StatusBadRequest)
		return
	}
	// loop produk, cari id, ganti sesuai data dari request
	for i := range produk {
		if produk[i].ID == id {
			updateProduk.ID = id
			produk[i] = updateProduk
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateProduk)
			return
		}
	}
}

func deleteProduk(w http.ResponseWriter, r *http.Request) {
	//getid
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	//gantiidmenjadiint
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid produk id", http.StatusBadRequest)
		return
	}

	//loopprodukcariiddapetindeksyangmaudihapus
	for i, p := range produk {
		if p.ID == id {
			//bikinslicebaruindeksdengandatasebelum dan sesudah indeks
			produk = append(produk[:i], produk[i+1:]...)
			//return produk
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "sukses delete",
			})
			return
		}

		http.Error(w, "produk belum ada", http.StatusNotFound)
	}
}

func main() {
	//GET localhost:8080/api/produk{id}
	//PUT localhost:8080/api/produk{id}
	//DELETE localhost:8080/api/produk{id}
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProdukByID(w, r)
		} else if r.Method == "PUT" {
			updateProduk(w, r)
		} else if r.Method == "DELETE" {
			deleteProduk(w, r)
		}

	})
	//GET localhost:8080/api/produk
	//POST localhost:8080/api/produk
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk)
		} else if r.Method == "POST" {
			//bacadarirequest
			var produkBaru Produk
			err := json.NewDecoder(r.Body).Decode(&produkBaru)
			if err != nil {
				http.Error(w, "invalid request", http.StatusBadRequest)
				return
			}
			// masukin data dari dalam variabel produk
			produkBaru.ID = len(produk) + 1
			produk = append(produk, produkBaru)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated) //201
			json.NewEncoder(w).Encode(produkBaru)
		}
	})

	//localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API RUNNING",
		})
		w.Write([]byte("OK"))
	})
	//masukkelocalhost:8080
	fmt.Println("Server running di localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("gagal running server")
	}

}
