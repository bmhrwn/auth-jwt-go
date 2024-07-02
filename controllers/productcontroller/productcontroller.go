package productcontroller

import (
	"net/http"

	"github.com/jeypc/go-jwt-mux/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := []map[string]interface{}{
		{
			"id":           1,
			"nama_product": "Kemeja",
			"stok":         1000,
		},
		{
			"id":           2,
			"nama_product": "Kemeja 2",
			"stok":         2000,
		},
		{
			"id":           3,
			"nama_product": "Kemeja 3",
			"stok":         3000,
		},
	}

	helper.ResponseJson(w, http.StatusOK, data)
}
