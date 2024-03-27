package soal

import "fmt"

// Question 2 for Data 1
func QuestionTwoDataOne() {
	// initialisation
	var mark = []float32 {78,1.69}
	var john = []float32 {92,1.95}
	
	// Output
	fmt.Println("\n/////////// QUESTION 2 DATA 1 ///////////")
	findOutHigherScore(mark,john)
}

// Question 2 for Data 2
func QuestionTwoDataTwo() {
	// initialisation
	var mark = []float32 {95,1.88}
	var john = []float32 {85,1.76}
	
	// Output
	fmt.Println("\n/////////// QUESTION 2 DATA 1 ///////////")
	findOutHigherScore(mark,john)

}

// A function that serves to find the result
func findOutHigherScore(mark,john []float32 ){
	// create a boolean to contain the condition result value
	var markHigherBMI bool

	// catch mark data
	markWeight := mark[0]
	markHeight := mark[1]
	
	// catch john data
	johnWeight := john[0]
	johnHeight := john[1]

	// Calculating BMI mark and John
	BMIMark  := calculateBMI(markWeight, markHeight)
	BMIJohn := calculateBMI(johnWeight, johnHeight)

	// Output BMI
	fmt.Printf("BMI Mark adalah %.2f \n",  BMIMark)
	fmt.Printf("BMI John adalah %.2f \n",  BMIJohn)


	if BMIMark > BMIJohn {
		markHigherBMI = true
	} else {
		markHigherBMI = false
	}
	// Display the result from condition above
   fmt.Println("Apakah benar BMI milik Mark lebih tinggi dari John? ", markHigherBMI)
}

// Calculating to get the BMI
func calculateBMI(weight, height float32) float32 {
	BMI  := float32(weight) / (height * height)
	return BMI
}