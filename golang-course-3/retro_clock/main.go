package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func clear() { // clear command for linux
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	colon := placeholder{
		"   ",
		" █ ",
		"   ",
		" █ ",
		"   ",
	}

	colon2 := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	_ = colon2

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	for {
		clear()
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		//fmt.Printf("hour: %d, min: %d, sec: %d\n", hour, min, sec)

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10], colon, digits[min/10], digits[min%10], colon, digits[sec/10], digits[sec%10],
		}

		if sec%2 == 0 {
			clock[2] = colon2
			clock[5] = colon2
		}

		for line := range clock[0] {
			for index, digit := range clock {
				next := clock[index][line]
				if digit == colon && sec%3 == 0 {
					next = "   "
				}
				fmt.Print(next)
				fmt.Print("  ")
			}
			fmt.Println()
		}
		fmt.Println()
		time.Sleep(time.Second)
	}
}
