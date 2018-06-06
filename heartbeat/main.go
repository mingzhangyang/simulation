package heartbeat

import (
	"fmt"
	"log"
	"strconv"
	"os"
)

// Exec start the simulation
func Exec() {
	fmt.Print("Please set the interval (if omitted, default value 1000 will be set):  ")
	var interval string
	_, err := fmt.Scan(&interval)
	if err != nil {
		fmt.Println("A single numeric argument expected")
		for i := 0; i < 5; i++ {
			_, err = fmt.Scan(&interval)
			if err == nil {
				break
			}
			log.Println("you are exiting the programme")
		}
	}
	fmt.Print("Please set the strength (if omitted, default value 10 will be set):  ")
	var strength string
	_, err = fmt.Scan(&strength)
	if err != nil {
		fmt.Println("A single numeric argument expected")
		for i := 0; i < 5; i++ {
			_, err = fmt.Scan(&strength)
			if err == nil {
				break
			}
			log.Println("you are exiting the programme")
		}
	}
	intev, err := strconv.Atoi(interval)
	if err != nil {
		intev = 1000
	}
	stren, err := strconv.Atoi(strength)
	if err != nil {
		stren = 10
	}
	hb := Initialize(State{intev, stren}, Deviation{500, 2}, Probablities{0.6, 0.2})
	log.Println("starting")
	go hb.Begin()
	var ins string
	for {
		fmt.Scan(&ins)
		i, err := strconv.Atoi(ins)
		if err != nil {
			log.Println("Please input 0, 1 or 2 to control the system")
		}
		switch i {
		case 0:
			if !hb.Paused {
				fmt.Println("been running")
			} else {
				hb.Paused = false
				fmt.Println("resuming")
				hb.Instruction <- 0
			}
		case 1:
			if hb.Paused {
				fmt.Println("paused. Input 0 to resume.")
			} else {
				hb.Paused = true
				fmt.Println("paused. Input 0 to resume")
			}
		case 2:
			close(hb.Instruction)
			log.Println("Exiting...")
			log.Println("Bye")
			os.Exit(0)
		default:
			fmt.Println("invalid instruction omitted. 0 for continue/resume, 1 for pause, 2 for stop")
		}
	}
}

