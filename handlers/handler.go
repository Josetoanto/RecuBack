package handlers

import (
	"encoding/json"
	"net/http"
	"recuAPI/domain"
	"recuAPI/repository"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductoHandler struct {
    repo *repository.ProductoRepository
}

func NewProductoHandler(repo *repository.ProductoRepository) *ProductoHandler {
    return &ProductoHandler{repo: repo}
}

func (h *ProductoHandler) AddProduct(c *gin.Context) {
    var producto domain.Producto
    if err := c.ShouldBindJSON(&producto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }

    h.repo.AddProduct(producto)
    c.JSON(http.StatusCreated, gin.H{"message": "Producto agregado"})
}

func (h *ProductoHandler) GetTemporaryProducts(c *gin.Context) {
    temporaryProductos := h.repo.GetTemporaryProductos()

    if len(temporaryProductos) == 0 {
        c.JSON(http.StatusOK, []domain.Producto{}) 
        return
    }

    c.JSON(http.StatusOK, temporaryProductos)
}


func (h *ProductoHandler) CountProductInDiscount(c *gin.Context) {
    flusher, ok := c.Writer.(http.Flusher)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Streaming no soportado"})
        return
    }

    for {
        cantidad := h.repo.CountProductInDiscount()
        json.NewEncoder(c.Writer).Encode(gin.H{"productos_con_descuento": cantidad})
        flusher.Flush()
        time.Sleep(5 * time.Second)
    }
}
