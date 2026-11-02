

package client

import (
	"MonEpicerie/boutique"
)



type Panier struct {

	AchatsClient map[string]int


}







func CalculerTotalPanier(AchatsClient Panier, catalogues map[string]boutique.Produit ) (float64,[]string) {










	var ListProduitManquants[]string
	var TotArticle float64
	var TotPanier float64
	i := 0

	for nomProduit, QteArticles := range AchatsClient.AchatsClient {

		i = i + 1
		produit, existe := catalogues[nomProduit]

		if existe {

			TotArticle = float64(produit.Prix) *  float64(QteArticles)


		}else{

			ListProduitManquants[i] = produit.Nom

			
		}

		TotPanier = TotPanier + TotArticle

	}

	return TotPanier, ListProduitManquants


}