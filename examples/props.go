// +build ignore

package main

import (
	"fmt"
	//"github.com/fari-99/fsm"
	".."
)

func main() {
	fsmTest := fsm.NewFSM(
		"closed",
		fsm.Events{
			{
				Name: "pull",
				Src: []string{"closed"},
				Dst: "open",
				Props:fsm.Properties{
					"editable": true,
					"deletable": false,},},
			{
				Name: "push",
				Src: []string{"closed"},
				Dst: "open",
				Props:fsm.Properties{
					"editable": false,
					"deletable": false,},},
			{
				Name: "close",
				Src: []string{"open"},
				Dst: "closed"},
		},
		fsm.Callbacks{},
	)

	fmt.Println(fsmTest.Current())
	fmt.Println(fsmTest.AvailableTransitions())

	fmt.Println(fsmTest.GetPropertiesTransitions("close"))
	fmt.Println(fsmTest.GetPropertiesTransitions("pull"))
	fmt.Println(fsmTest.GetPropertiesTransitions("push"))

	fsmTest.Event("pull")
	fmt.Println(fsmTest.Current())
}
