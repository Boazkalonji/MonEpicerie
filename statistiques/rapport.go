

package statistiques


import (
	
	"errors"
	"MonEpicerie/boutique"
	

)



func CalculerValeurTotale(catalogues map[string]boutique.Produit) (float64, error) {

	if len(catalogues) == 0{

		return 0, errors.New("le catalogue est vide")

	}

	var totalCatalogues float64


	for _, produit := range catalogues{
		totalCatalogues += produit.Prix * float64(produit.Qte)

		
	}

	return totalCatalogues, nil



}