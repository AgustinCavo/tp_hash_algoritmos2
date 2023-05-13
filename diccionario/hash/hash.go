package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

const TAM = 100
const FACTORCARGA = 2

//FactorCarga--->achicar 0,5

type parClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K comparable, V any] struct {
	tabla    []TDALista.Lista[parClaveValor[K, V]]
	tam      int
	cantidad int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := new(hashAbierto[K, V])
	hash.tam = TAM
	return hash
}

func (d *hashAbierto[K, V]) Cantidad() int {
	return d.cantidad
}

func (d *hashAbierto[K, V]) Guardar(clave K, dato V) {

}

func (d *hashAbierto[K, V]) Pertenece(clave K) bool { //Reutilizo en otras primitivas
	//aplico funcion hashing a la clave y me devuelve un numero que es la posicion en donde esta en el vector
	for _, lista := range d.tabla { //itero por indice el array, en cada indice hay una lista
		iter := lista.Iterador()
		for iter.HaySiguiente() {
			if iter.VerActual().clave == clave {
				return true
			}
			iter.Siguiente()
		}
	}
	return false
}

func (d *hashAbierto[K, V]) Obtener(clave K) V {
	d.ValidarPertenece(clave)

}

func (d *hashAbierto[K, V]) Borrar(clave K) V {

}

func (d *hashAbierto[K, V]) redimensionar(tam int) { //Redimensiono cuando el F=n/M >(2 ó 3)??? Siendo n=cant de elementos realmente en la tabla
	//y M tamaño total del vector,  F es el factor de carga
	tabla_nueva := make([]TDALista.Lista[parClaveValor[K, V]], tam) //Me creo una tabla nueva
	//Los elementos de la tabla que ya tenia los tengo que "rehashear" y mandarlos a la tabla_nueva
	//Deberia iterar la tabla que tenia..
}

func (d *hashAbierto[K, V]) ValidarPertenece(clave K) {
	if !d.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}
