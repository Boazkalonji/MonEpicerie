

package commande


import (

	"fmt"
)


type Commande struct {


	Id uint `json:id`
	Produit string `json:produit`
	Prix float64 `josn:prix`



}





func  ChanelCommandes(commande chan map[string]any, state chan bool){


	for v := range commande {

		fmt.Printf("TTT dans le channel ..... %v", v)
	}

	state <- true 
}