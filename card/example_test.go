package card_test

import (
	"fmt"

	"github.com/Souvik-Cyclic/card-validator/card"
)

func ExampleIsValid() {
	fmt.Println(card.IsValid("4111 1111 1111 1111"))
	fmt.Println(card.IsValid("4111111111111112"))
	// Output:
	// true
	// false
}

func ExampleIssuer() {
	fmt.Println(card.Issuer("4111111111111111"))
	fmt.Println(card.Issuer("36000000000008"))
	fmt.Println(card.Issuer("6521000000000000"))
	// Output:
	// Visa
	// Diners Club
	// RuPay
}
