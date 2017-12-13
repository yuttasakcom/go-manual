package main

import "fmt"

var ligthSwitchIsOn bool = false

func main() {
	printLightSwitchState()

	toggleLightSwitch()
	printLightSwitchState()

	toggleLightSwitch()
	printLightSwitchState()
}

func printLightSwitchState() {
	fmt.Println("The light switch is off:", ligthSwitchIsOn)
}

func toggleLightSwitch() {
	ligthSwitchIsOn = !ligthSwitchIsOn
}
