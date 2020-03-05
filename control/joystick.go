package control

import (
	joy "github.com/simulatedsimian/joystick"
)

// Joystick - Data type to store the joystick data
type Joystick struct {
	joystickID int
	js         joy.Joystick
}

// NewJoystick creates a joystick instance
func NewJoystick(id int) *Joystick {
	return &Joystick{joystickID: id}
}

// Init initiates the joystick working
func (joystick *Joystick) Init() (err error) {
	joystick.js, err = joy.Open(joystick.joystickID)
	return err
}

// GetCommands builds the command repository to be sent
func (joystick Joystick) GetCommands() (CommandsRepository, error) {
	State, err := joystick.js.Read()
	if err != nil {
		return nil, err
	}

	angularVelocity := float64(State.AxisData[0] / 256)
	if angularVelocity > -32 && angularVelocity < 32 {
		angularVelocity = 0
	}

	linearVelocity := float64(State.AxisData[0] / 256)
	if linearVelocity > -32 && linearVelocity < 32 {
		linearVelocity = 0
	}

	cmdRepo := NewCommandsRepository(3)
	if State.Buttons == 1 { // If the A button is pressed, the green robot have to be controlled
		cmdRepo[0] = StandardCommand{RobotID: 1, LinearVelocity: 0, AngularVelocity: 0}
		cmdRepo[1] = StandardCommand{RobotID: 2, LinearVelocity: linearVelocity, AngularVelocity: angularVelocity}
		cmdRepo[2] = StandardCommand{RobotID: 3, LinearVelocity: 0, AngularVelocity: 0}
	} else if State.Buttons == 2 { // If the B button is pressed, the red robot have to be controlled
		cmdRepo[0] = StandardCommand{RobotID: 1, LinearVelocity: linearVelocity, AngularVelocity: angularVelocity}
		cmdRepo[1] = StandardCommand{RobotID: 2, LinearVelocity: 0, AngularVelocity: 0}
		cmdRepo[2] = StandardCommand{RobotID: 3, LinearVelocity: 0, AngularVelocity: 0}
	} else if State.Buttons == 4 { // If the X button is pressed, the magenta robot have to be controlled
		cmdRepo[0] = StandardCommand{RobotID: 1, LinearVelocity: 0, AngularVelocity: 0}
		cmdRepo[1] = StandardCommand{RobotID: 2, LinearVelocity: 0, AngularVelocity: 0}
		cmdRepo[2] = StandardCommand{RobotID: 3, LinearVelocity: linearVelocity, AngularVelocity: angularVelocity}
	} else { // If none activation button is pressed, none robot have to be controlled
		cmdRepo[0] = StandardCommand{RobotID: 1, LinearVelocity: 0, AngularVelocity: 0}
		cmdRepo[1] = StandardCommand{RobotID: 2, LinearVelocity: 0, AngularVelocity: 0}
		cmdRepo[2] = StandardCommand{RobotID: 3, LinearVelocity: 0, AngularVelocity: 0}
	}

	return cmdRepo, nil
}
