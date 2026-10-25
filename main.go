

package main

import (
	"fmt"
	"MonEpicerie/boutique"
)

func main() {

	for _, produit := range boutique.ObtenirCatalogue() {


		statut_stock := boutique.EstEnRupture(produit.Qte)

		if statut_stock == false || statut_stock == true {

			fmt.Printf("\n %s - %v - %d ", produit.Nom, produit.Prix, produit.Qte)
		}


	}



}