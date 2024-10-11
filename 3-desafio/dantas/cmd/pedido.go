package cmd

// Pedido representa um pedido no restaurante.
type Pedido struct {
	ID       int
	Items    []Item
	Pronto   bool
	IsVIP    bool
	Expirado bool
}
