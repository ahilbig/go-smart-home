package main

import (
	"fmt"
	"go-smart-home/pkg/client"
	"go-smart-home/pkg/client/switch_operations"
	"go-smart-home/pkg/models"
	"log"
)

func main() {
	c := client.NewHTTPClientWithConfig(nil, client.DefaultTransportConfig().WithHost("127.0.0.1:57861"))

	// make the request to get all switches
	resp, err := c.SwitchOperations.ListSwitches(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Existing switches: %#v\n", resp.Payload)

	// Generate some sample switches

	desc1 := "Switch1"
	desc2 := "Switch2"
	desc3 := "Switch3"

	_, err = c.SwitchOperations.AddSwitch(
		switch_operations.NewAddSwitchParams().WithBody(&models.Switch{ID: 1, Description: &desc1}))
	if err != nil {
		log.Println(err)
	}
	_, err = c.SwitchOperations.AddSwitch(
		switch_operations.NewAddSwitchParams().WithBody(&models.Switch{ID: 2, Description: &desc2}))
	if err != nil {
		log.Println(err)
	}
	_, err = c.SwitchOperations.AddSwitch(
		switch_operations.NewAddSwitchParams().WithBody(&models.Switch{ID: 3, Description: &desc3}))
	if err != nil {
		log.Println(err)
	}

	// make the request to get all switches
	resp, err = c.SwitchOperations.ListSwitches(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Generated switches: %s\n", resp.Payload)
	fmt.Printf("Switch1=%s", resp.Payload[0])
}
