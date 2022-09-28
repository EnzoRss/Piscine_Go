package main

func UltimateDivMod(a *int, b *int) {
	var stockdiv int
	var stockmod int

	stockdiv = *a / *b
	stockmod = *a % *b
	*a = stockdiv
	*b = stockmod
}
