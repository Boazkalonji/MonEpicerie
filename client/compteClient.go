

package client

import "fmt"



type Payeur interface {

	Payer(montant *float64) float64

}


type CarteBancaire struct{

	InfoClient Client

}


type PayPal struct{

	InfoClient Client
	
}




type Client struct{


	Nom string
	Solde float64

}


func (cb CarteBancaire ) Payer(montant *float64) float64{


	*montant = cb.InfoClient.Solde 

	fmt.Printf("paiement faite sous CB")

	return cb.InfoClient.Solde 

}



func (pp PayPal ) Payer(montant *float64) float64{

	*montant += pp.InfoClient.Solde 

	fmt.Printf("paiement faite sous paypal")

	return pp.InfoClient.Solde  
	
}




func ValiderPaiement(p Payeur, montant *float64){

	soldeFinal := p.Payer(montant)

	fmt.Printf("Le solde est de : %.2f", soldeFinal)


}



func (c Client) Crediter(montant float64) {

	AvantSolde := c.Solde

	c.Solde += montant

	if AvantSolde != c.Solde {

		fmt.Printf("LE SOLDE DU CLIENT %s a ete modifier avec succes nouveau solde  : %.2f", c.Nom, c.Solde)

	}






}