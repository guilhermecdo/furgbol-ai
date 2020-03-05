package control

import (
	"fmt"
	"testing"
	"time"
)

func TestJoystick(t *testing.T) {
	joystick := NewJoystick(1)
	err := joystick.Init()
	if err != nil {
		t.Errorf("Error on initializing joystick: %v", err)
	} else {
		running := true
		ticker := time.NewTicker(time.Minute)
		done := make(chan bool)

		go func() {
			for {
				select {
				case <-done:
					return
				case <-ticker.C:
					running = false
					ticker.Stop()
				}
			}
		}()

		for running {
			cmdRepo, err := joystick.GetCommands()
			if err != nil {
				t.Errorf("Error on reading data: %v", err)
				break
			} else {
				fmt.Println(cmdRepo)
				time.Sleep(10 * time.Millisecond)
			}
		}
	}
}
