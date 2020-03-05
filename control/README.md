# Control Package

The control package contains modules that provide control of the environment, such as: commands, strategies, controllers, etc.

## Contents

- [Standard Command Type](#stdcmd)
- [Commands Repository Type](#cmdrepo)
- [Simulator Command Type](#slrcmd)
- [Joystick Type](#joystick)

<a name="stdcmd"></a>

### Standard Command Type

This type models the standard system command, containing the following fields:

- *RobotID()*
- *LinearVelocity()*
- *AngularVelocity()*

<a name="cmdrepo"></a>

### Commands Repository Type

This type is an alias to a slice: *[]StandardCommand*

To create a new Commands Repository, the **NewCommandsRepository**(*numberOfRobots* int) function is provided.

<a name="slrcmd"></a>

### Simulator Command Type

This type models the simulator command, having the following fields:

- *RobotID()*
- *LeftWheelVelocity()*
- *RightWheelVelocity()*
- *FinalPosition()*

### Joystick Type

This type implements the command generation interface, providing the method: (*joystick Joystick*) **GetCommands()** (*CommandsRepository, error*).
To create and initiate a joystick use the methods:

- **NewJoystick**(*joystickID int*)
- (*joystick *Joystick*) **Init** (*err error*)