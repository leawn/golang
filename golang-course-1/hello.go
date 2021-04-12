package main

func main() {
	// fmt.Print("Wie alt bist du? ")
	// reader := bufio.NewReader(os.Stdin)
	// input, err := reader.ReadString('\n')

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// input = strings.TrimSpace(input)

	// old, err := strconv.ParseFloat(input, 64)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if old >= 18 {
	// 	fmt.Println("Wow ur rly old")
	// } else {
	// 	fmt.Println("Go home")
	// }

	// seconds := time.Now().Unix()
	// rand.Seed(seconds)

	// target := rand.Intn(100) + 1
	// fmt.Println("I choose the number from 1 to 100")
	// fmt.Println("Number is chosen")
	// fmt.Println(target)

	// readernew := bufio.NewReader(os.Stdin)

	// success := false

	// for guesses := 0; guesses < 10; guesses++ {

	// 	fmt.Println("You got ", 10-guesses, " chances to guess")
	// 	fmt.Print("Choose your number: ")
	// 	inputnew, err := readernew.ReadString('\n')
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	inputnew = strings.TrimSpace(inputnew)

	// 	guess, err := strconv.Atoi(input)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	if guess > target {
	// 		fmt.Println("Your number is greater than the secret number")
	// 	} else if guess < target {
	// 		fmt.Println("Your number is less than the secret number")
	// 	} else {
	// 		success = true
	// 		fmt.Println("Congrats! You won!")
	// 		break
	// 	}
	// }

	// if !success {
	// 	fmt.Println("Sorry, you don't have any more chances! The secret number was: ", target)
	// }

	// notes := [7]string{
	// 	"do",
	// 	"re",
	// 	"mi",
	// }

	// for _, value := range notes {
	// 	fmt.Println(value)
	// }

	// lines, err := datafile.GetStrings("data.txt")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// counts := make(map[string]int)

	// for _, line := range lines {
	// 	counts[line]++
	// }

	// fmt.Print(counts)

	// langs := map[string]float64{
	// 	"java":   100,
	// 	"golang": 10,
	// 	"python": 150,
	// }

	// var names []string

	// for name := range langs {
	// 	names = append(names, name)
	// }

	// sort.Strings(names)

	// for _, value := range names {
	// 	fmt.Println(value, langs[value])
	// }

	// arrayLiters := [5]string{"a", "b", "c", "d", "e"}

	// slice := []string{
	// 	"q",
	// 	"w",
	// }
	// slice2 := arrayLiters[2:]

	// slice = append(slice, "e")
	// fmt.Println(slice, slice2)

	var auto struct {
		name string
		rate float64
		new  bool
	}

	auto.name = "Porsche 911"
	auto.rate = 9.11
	auto.new = true
}
