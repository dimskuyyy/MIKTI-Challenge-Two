package soal

import "fmt"

// calculating the average
func calculatingTheAverage(score []int) float32{
	var average float32
	var count int = 0

	// Looping to calculate the values in the array
	for _, value := range score{
		count = count+value
	}

	// get the average
	average = float32(count) / float32(len(score))
	return average
}

// Question for Data 1
func QuestionOne(){

	// Initialisation
	var averageLumbaLumba, averageKoala float32
	var lumbaLumbaScore = []int{96,108,89}
	var koalaScore = []int{88,91,110}

	// calculating the average
	averageLumbaLumba = calculatingTheAverage(lumbaLumbaScore);
	averageKoala = calculatingTheAverage(koalaScore);
	
	// Output
	fmt.Println("\n/////////// Question Data 1 ///////////")
	fmt.Println("Average skor tim lumba-lumba	: ",averageLumbaLumba)
	fmt.Println("Average skor tim koala		: ",averageKoala)
	
	// Determining the winning team
	if averageKoala < averageLumbaLumba {
		fmt.Println("Tim Lumba-Lumba Win!!!")
	}else if averageKoala > averageLumbaLumba {
		fmt.Println("Tim Koala Win!!!")
	}else{
		fmt.Println("SKOR SERI OLEH KEDUA TIM!!!")
	}
}

// Question for Data Bonus 1

func QuestionBonusOne(){
	// Initialisation
	var averageLumbaLumba, averageKoala float32
	var lumbaLumbaScore = []int{97,112,101}
	var koalaScore = []int{109,95,123}
	
	// calculating the average
	averageLumbaLumba = calculatingTheAverage(lumbaLumbaScore)
	averageKoala = calculatingTheAverage(koalaScore)

	// Output
	fmt.Println("\n/////////// Question Bonus 1 ///////////")
	fmt.Println("Average skor tim lumba-lumba	: ",averageLumbaLumba)
	fmt.Println("Average skor tim koala		: ",averageKoala)

	/*
		Determine the winning team with the following conditions: 
		Include a requirement for a minimum score of 100. With this rule, 
		a team only wins if it has a higher score than the other teams, 
		and at the same time a minimum score of 100 points.
	*/

	// If same average
	if averageLumbaLumba == averageKoala{
		fmt.Println("SKOR SERI OLEH KEDUA TIM!!!")
	}
	// If the average is eligible and there is a winner
	if (averageLumbaLumba > averageKoala) && averageLumbaLumba >= 100{
		fmt.Println("Tim Lumba-Lumba Win!!!")
	}else if (averageLumbaLumba < averageKoala) && averageKoala >= 100{
		fmt.Println("Tim Koala Win!!!")
	}else{
		fmt.Println("Tim Tidak Ada Yang Menang, Syarat Minimum Adalah 100!!!")
	}
}


// Question for Bonus Two
func QuestionBonusTwo(){
	// Initialisation
	var averageLumbaLumba, averageKoala float32
	var lumbaLumbaScore = []int{97,112,101}
	var koalaScore = []int{109,95,106}
	
	// calculating the average
	averageLumbaLumba = calculatingTheAverage(lumbaLumbaScore)
	averageKoala = calculatingTheAverage(koalaScore)

	// Output
	fmt.Println("\n/////////// Question Bonus 1 ///////////")
	fmt.Println("Average skor tim lumba-lumba	: ",averageLumbaLumba)
	fmt.Println("Average skor tim koala		: ",averageKoala)

	/* 
		Determine the winning team with the following conditions: 
		Minimum score also applies to ties! So a tie only happens when both 
		teams have the same score and both have a score greater than or equal 
		to 100 points. Otherwise, neither team wins the trophy.
	*/

	// Checking if the average is eligible and there is a winner
	if (averageLumbaLumba > averageKoala) && averageLumbaLumba >= 100{
		fmt.Println("Tim Lumba-Lumba Win!!!")
	}else if (averageLumbaLumba < averageKoala) && averageKoala >= 100{
		fmt.Println("Tim Koala Win!!!")
	}else if (averageLumbaLumba == averageKoala) && (averageLumbaLumba >=100 && averageKoala >= 100){
		fmt.Println("SKOR SERI OLEH KEDUA TIM!!!")
	}else{
		fmt.Println("Tim Tidak Ada Yang Menang, Syarat Minimum Adalah 100!!!")
	}
}