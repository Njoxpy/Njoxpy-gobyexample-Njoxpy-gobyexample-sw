// Go has various value types including strings,
// integers, floats, booleans, etc. Here are a few
// basic examples.
// Go ina aina mbalimbali za value ambazo ni sentensi,
// namba kamili, decimali, buliani na nyinginezo n.k. Hapa ni baadhi ya chache
// mifano ya kawaida

package main

import "fmt"

func main() {

	// Strings, which can be added together with `+`.
	// Sentensi, zinaweza kuongezwa pamoja kwa kutumia `+`.
	fmt.Println("go" + "lang")

	// Integers and floats.
	// Namba kamili na desimali
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Booleans, with boolean operators as you'd expect.
	// Buliani, pamoja na operesheni za buliani kama ambazo unategemea.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
