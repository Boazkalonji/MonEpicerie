

package main

import (
	"fmt"
	"encoding/json"
	"MonEpicerie/commande"
	
	/*"MonEpicerie/boutique"
	"MonEpicerie/statistiques"
	"MonEpicerie/client"
	*/
)

/*
func VendrePommes(stock *int){

	*stock = *stock - 10

}






func updateB(item map[string]int){

	item["bonbon"] = 10
}

func updateA(item []int){

	item[0]= 900
	item[1]=800

}



func updateD(AdresseoFpointeur *int , value int){

	*AdresseoFpointeur = value

}




*/






func main() {





	commandeOne := commande.Commande {

		Id : 1,
		Produit : "pain",
		Prix : 500,

	}




	SortieJSON, err := json.Marshal(commandeOne)


	if err != nil{
		fmt.Printf("La map a ete bien transforrme le JSON : %v \n ", SortieJSON)
	}

	RecuJson := []byte(`{

		"id" : 2, "produit" : "lait", "prix" : 800,

	}`)


	var resultat map[string]any


	json.Unmarshal(RecuJson, &resultat)

	fmt.Printf("LES JSON RECU EST : \n  %v", resultat)





	var commandes = make(chan map[string]any)
	var state = make(chan bool)


	

	go commande.ChanelCommandes(commandes, state)

	commandes <- map[string]any{
		"id": "bb",
		"hfhfh":52,
	}

	close(commandes)

	<- state
	fmt.Printf("fin du travail du goroutine")
	
	












 
/*




	original := 500

	adresseOForiginal := &original

	updateD(adresseOForiginal,200)

	fmt.Printf("la'adresse de la vairable original est %v et sa valeur est ", adresseOForiginal, original )


	fmt.Printf("---------------------------------------------------------------------------------------")










	MonSlice := []int{1,2,3,6,9,8,5,1,2,5,6,}

	updateA(MonSlice)

	fmt.Println(MonSlice)


	fmt.Printf("--------------------------------------------------------------\n")


	groceryList := map[string]int{

		"pain": 6,
		"lait": 7,

	}

	updateB(groceryList)

	fmt.Printf("%v", groceryList)

















	/*

	stockPommes := 50


	VendrePommes(&stockPommes)

	fmt.Printf("LA VALEUR DE LA POMME EST %d",stockPommes )


	



	for _, produit := range boutique.ObtenirCatalogue() {


		statut_stock := boutique.EstEnRupture(produit.Qte)

		if statut_stock == false || statut_stock == true {

			fmt.Printf("\n %s - %v - %d ", produit.Nom, produit.Prix, produit.Qte)
		}


	}

	fmt.Println("----------------------------------------")

	catalogue := boutique.ObtenirCatalogue()
	totalCatalogues , statut := statistiques.CalculerValeurTotale(catalogue)



	if totalCatalogues != 0 || statut == nil {

		fmt.Printf("\b LES RESULTAT EST : TOTAL CATALOGUES %v - STATUT", totalCatalogues, statut)

	}else{

		fmt.Println("\b LES RESULATAT EST NEGATIVE CAR LE MAP EST VIDE ", totalCatalogues, statut)

	}

	fmt.Println("-----------------------------------------------------------------")



	panier_client := client.Panier{

		AchatsClient: map[string]int{
			"pomme": 300,
			"lait":6,
		},
	}



	

	facture, ListProduitManquants := client.CalculerTotalPanier(panier_client, catalogue)




	fmt.Printf("-----LE MONTANT TOTAL FACTURE EST DE : -------- %v$ Americain", facture)

	if len(ListProduitManquants) == 0 {
		fmt.Printf("\n-----LE(S) PRODUITS MANQUATS SONT  : \n", 0)
	}else{

		fmt.Printf("-----LE(S) PRODUITS MANQUATS SONT  : \n")
		for i, ProduitManquant := range ListProduitManquants{

			fmt.Println(ProduitManquant[i])

		}

	}
	
	*/
	
	

	




}