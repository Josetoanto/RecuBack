package domain

type Producto struct {
	Nombre    string `json:"nombre"`
	Precio    int    `json:"precio"`
	Codigo    string `json:"codigo"`
	Descuento bool   `json:"descuento"`
}
