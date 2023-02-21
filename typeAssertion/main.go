package main

import "fmt"

type ninjaStar struct {
	owner string
}
type ninjaSword struct {
	owner string
}
type ninjaWeapon interface {
	attack()
}

func (n ninjaStar) pickup() {
	fmt.Println("Picking up the ninja star..")
}
func (n ninjaStar) attack() {
	fmt.Println("Throwing ninja star..")
	

	fmt.Printf("answer of attack cloneis : %v")
}
func (n ninjaStar)attack( )string{
	fmt.Println("clone of attack")
	str:=n.owner
	fmt.Printf("answer string of clone attack is : %v",str)
}
func (n ninjaSword) attack() {
	fmt.Println("Throwing ninja sword")
}

func main() {

	// doubt hai
	
	var nin ninjaWeapon
	nin = ninjaStar{"Aman"}
	// fmt.Println(nin.(ninjaStar).owner)
	


	weapons :=
		[]ninjaWeapon{
			ninjaStar{"aman"},
			ninjaSword{"tarar"},
		}

	for _, weapon := range weapons {
		weapon.attack()
		fmt.Println(weapon.(ninjaStar).owner)

	}
}