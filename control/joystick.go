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
	return &CommandsRepository{joystickID: id}
}

// Init initiates the joystick working
func (joystick Joystick) Init() error {
	joystick.js, err := joy.Open(joystick.joystickID)
	return err
}

// GetCommands builds the command repository to be sent
func (joystick Joystick) GetCommands() (CommandsRepository, error) {
	State, err := joystick.js.Read()
	if err != nil {
		return nil, err
	}

	linearVelocity, angularVelocity float64
	if State.AxisData[0] > -8192 && State.AxisData[0] < 8192 {
		angularVelocity = 0
	} else {
		angularVelocity = float64(State.AxisData[0] / 256)
	}

	if State.AxisData[1] > -8192 && State.AxisData[1] < 8192 {
		linearVelocity = 0
	} else {
		linearVelocity = float64(State.AxisData[0] / 256)
	}

	cmdRepo := NewCommandsRepository(3)
	if State.Buttons == 1 { // If the A button is pressed
		cmdRepo[0] = StandardCommand{RobotID: 1, LinearVelocity: linearVelocity, AngularVelocity: angularVelocity}
		cmdRepo[1] = StandardCommand{RobotID: 2, LinearVelocity: 0, AngularVelocity: 0}
		cmdRepo[2] = StandardCommand{RobotID: 3, LinearVelocity: 0, AngularVelocity: 0}
	} else if State.Buttons == 2 { // If the B button is pressed
		cmdRepo[0] = StandardCommand{RobotID: 2, LinearVelocity: linearVelocity, AngularVelocity: angularVelocity}
		cmdRepo[1] = StandardCommand{RobotID: 1, LinearVelocity: 0, AngularVelocity: 0}
		cmdRepo[2] = StandardCommand{RobotID: 3, LinearVelocity: 0, AngularVelocity: 0}
	} else if State.Buttons == 4 { // If the X button is pressed
		cmdRepo[0] = StandardCommand{RobotID: 3, LinearVelocity: linearVelocity, AngularVelocity: angularVelocity}
		cmdRepo[1] = StandardCommand{RobotID: 1, LinearVelocity: 0, AngularVelocity: 0}
		cmdRepo[2] = StandardCommand{RobotID: 2, LinearVelocity: 0, AngularVelocity: 0}
	} else {
		cmdRepo[0] = StandardCommand{RobotID: 1, LinearVelocity: 0, AngularVelocity: 0}
		cmdRepo[1] = StandardCommand{RobotID: 2, LinearVelocity: 0, AngularVelocity: 0}
		cmdRepo[2] = StandardCommand{RobotID: 3, LinearVelocity: 0, AngularVelocity: 0}
	}

	return cmdRepo, nil
}
