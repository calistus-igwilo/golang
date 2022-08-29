/*
Let us assume the following formula for displacement s as a function
of time t, acceleration a, initial velocity vo, and initial displacement so.

s = ½ a t2 + vot + so

Write a program which first prompts the user to enter values for
acceleration, initial velocity, and initial displacement. Then the
program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn()
which takes three float64 arguments, acceleration a, initial
velocity vo, and initial displacement so. GenDisplaceFn() should
return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement. The function returned by GenDisplaceFn() should take
one float64 argument t, representing time, and return one float64
argument which is the displacement travelled after time t.
*/
package main

import (
	"fmt"
)

var(
	accel float64
	vel float64
	disp float64
	t float64
)

func ReadInput(){
	fmt.Println("Enter the acceleration, velocity and distance, seperated by spaces")
	fmt.Scanf("%f %f %f", &accel, &vel, &disp)
	fmt.Println("Enter the value of time: ")
	fmt.Scanf("%f", &t)
}

func GenDisplaceFn(accel float64, vel float64, disp float64) func (float64) float64 {
	fn := func (t float64) float64{
		s := 0.5 * accel * t * t + vel * t + disp
		return s
	}
	return fn
}

func main() {
	ReadInput()

	fn := GenDisplaceFn(accel, vel, disp)
	fmt.Println(fn(t))
}
