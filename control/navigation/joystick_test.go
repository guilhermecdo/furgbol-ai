package control

import (
	"fmt"
	"testing"
)

func TestJoystick(t *testing.T) {
	joystick := NewJoystick(1)
	err := joystick.Init()
	if err != nil {
		t.Errorf("Error on initializing joystick: ", err)
	} else {
		for {
			cmdRepo, err := joystick.GetCommands()
			if err != nil {
				t.Errorf("Error on reading data: ", err)
				break
			} else {
				fmt.Println(cmdRepo)
			}
		}
	}
}
