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

