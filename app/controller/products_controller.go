package controller

import (
	"belajar-golang/app/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := models.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newProduct, err := models.CreateProduct(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Produk berhasil ditambahkan",
		"product": newProduct,
	})

	w.WriteHeader(http.StatusCreated)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	// Mengambil ID produk dari query parameter
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing product ID", http.StatusBadRequest)
		return
	}

	// Mengubah string ID menjadi uint
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Menghapus produk berdasarkan ID
	err = models.DeleteProduct(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mengembalikan respons sukses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Produk berhasil dihapus"}`))
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    if idStr == "" {
        http.Error(w, "Missing product ID", http.StatusBadRequest)
        return
    }
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }

    var p models.Product
    if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Jika ID di badan permintaan kosong, gunakan ID dari URL
    if p.ID == 0 {
        p.ID = uint(id)
    }

    // Pastikan ID di badan permintaan cocok dengan ID di URL
    if uint(id) != p.ID {
        http.Error(w, "Product ID in URL and body must match", http.StatusBadRequest)
        return
    }
	log.Printf("ID from URL: %d, ID from body: %d\n", id, p.ID)

    err = models.UpdateProduct(p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Produk berhasil diperbarui",
    })
}

func GetProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Query().Get("id")
	if idstr == "" {
		http.Error(w, "Missing product ID", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
	}

	product, err := models.GetProductById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if product == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(product)
}
