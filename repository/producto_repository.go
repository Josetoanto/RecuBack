package repository

import (
    "recuAPI/domain"
    "sync"
    "time"
)

type ProductoRepository struct {
    productos         []domain.Producto
    temporaryProductos []domain.Producto
    productosConDescuento int
    mutex             sync.Mutex
}

func NewProductoRepository() *ProductoRepository {
    return &ProductoRepository{
        productos:         []domain.Producto{},
        temporaryProductos: []domain.Producto{},
        productosConDescuento: 0,
    }
}

func (r *ProductoRepository) AddProduct(producto domain.Producto) {
    r.mutex.Lock()
    defer r.mutex.Unlock()

    r.productos = append(r.productos, producto)
    r.temporaryProductos = append(r.temporaryProductos, producto)

    if producto.Descuento {
        r.productosConDescuento++
    }

    go func() {
        time.Sleep(5 * time.Second)
        r.mutex.Lock()
        defer r.mutex.Unlock()
        if len(r.temporaryProductos) > 0 {
            r.temporaryProductos = r.temporaryProductos[1:]
        }
    }()
}

func (r *ProductoRepository) GetTemporaryProductos() []domain.Producto {
    r.mutex.Lock()
    defer r.mutex.Unlock()
    return r.temporaryProductos
}

func (r *ProductoRepository) CountProductInDiscount() int {
    r.mutex.Lock()
    defer r.mutex.Unlock()
    return r.productosConDescuento
}
