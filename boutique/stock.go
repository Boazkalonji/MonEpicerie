


package boutique 


/*

	definisions de la structure


*/


type Produit struct {

	Nom string
	Prix float64
	Qte uint8
}

func EstEnRupture(Qte uint8) bool{

	if Qte == 0 {

		return true

	}else{
		return false
	}

}


func ObtenirCatalogue() map[string]Produit {

	stocks := map[string]Produit{

		"pomme":{Nom: "pomme", Prix: 15.5, Qte:100},
		"pain":{Nom: "pain", Prix: 200, Qte:5},
		"lait":{Nom: "lait", Prix: 5.5, Qte:0},
	}

	return stocks
	 

}