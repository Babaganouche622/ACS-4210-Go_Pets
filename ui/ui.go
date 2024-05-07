package ui

import (
	"ACS-4210-Go_Pets/colour"
	"ACS-4210-Go_Pets/monster"
	"ACS-4210-Go_Pets/pet"
	"ACS-4210-Go_Pets/storage"
	"ACS-4210-Go_Pets/weather"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// StartGame starts the game UI
func StartGame() {
	var userPet pet.Tamagotchi
	var err error
	var choice string
	var petName string
	var eggName string

	// Display the weather
	fmt.Println(weather.GetWeather())

	reader := bufio.NewReader(os.Stdin)

	for choice == "" {
		fmt.Println(colour.Yellow + "\nWould you like to load a saved pet?" + colour.Reset)
		fmt.Printf("%sYes:%s 1 %sNo:%s 2 %sExit:%s 3\n\n",
			colour.Blue, colour.Reset,
			colour.Blue, colour.Reset,
			colour.Blue, colour.Reset,
		)
		choice, _ = reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		switch choice {
		case "1":
			for petName == "" {
				fmt.Println(colour.Yellow + "\nWhich pet would you like to load?" + colour.Reset)
				allPets, err := storage.GetAllPetData("data/pets.json")
				if err != nil {
					fmt.Println(colour.Red+"Failed to load pet data:"+colour.Reset, err)
					petName = ""
				}
				for name := range allPets {
					fmt.Println(colour.Green + name + colour.Reset)
				}
				petName, _ = reader.ReadString('\n')
				petName = strings.TrimSpace(petName)
				petData, err := storage.LoadPetData("data/pets.json", petName)
				if err != nil {
					fmt.Println(colour.Red+"Failed to load pet data:"+colour.Reset, err)
					petName = ""
				}
				if petName != "" {
					userPet = *petData
				}
			}
			// Continue with the loaded pet data
		case "2":
			// Fall through to create a new pet
		case "3":
			fmt.Println(colour.Green + "Goodbye!" + colour.Reset)
			os.Exit(0)
		default:
			fmt.Println(colour.Red + "Invalid choice. Please enter 1 or 2." + colour.Reset)
			// choice = ""
		}
		fmt.Println(choice)
	}

	if userPet.Name == "" {
		fmt.Println(colour.Yellow + "Let's hatch a new pet!" + colour.Reset)
		for eggName == "" {
			fmt.Println(colour.Yellow + "What would you like to name your egg?" + colour.Reset)
			eggName, _ = reader.ReadString('\n')
			eggName = strings.TrimSpace(eggName)
			_, err = storage.LoadPetData("data/pets.json", eggName)
			if err == nil {
				fmt.Println(colour.Orange + "A pet with this name already exists. Please choose another name." + colour.Reset)
				eggName = ""
			}
		}
		userPet = pet.Tamagotchi{
			Name:          eggName,
			Happiness:     100,
			CurrentHealth: 100,
			MaxHealth:     100,
			Attack:        20,
			Dirty:         0,
			PoopState:     0,
			Hunger:        0,
			Age:           0,
		}
		// Incubate the egg
		fmt.Println(userPet.DisplayStats("Egg"))
		go func() {
			for i := 0; i < 99; i++ {
				fmt.Print("|")
				time.Sleep(100 * time.Millisecond)
			}
		}()
		time.Sleep(10 * time.Second)
	}

	for !userPet.IsDead() {
		// Clear the screen and display the pet stats
		fmt.Println(userPet.DisplayStats("Idle"))

		// We should always print the options to interact with the egg
		// Display options
		fmt.Println(colour.Yellow + "What would you like to do?" + colour.Reset)
		fmt.Printf("%sFeed:%s 1 %sClean:%s 2 %sPlay:%s 3 %sPoop:%s 4 %sBattle:%s 5 %sSave:%s 6 %sExit:%s 7\n",
			colour.Blue, colour.Reset,
			colour.Blue, colour.Reset,
			colour.Blue, colour.Reset,
			colour.Blue, colour.Reset,
			colour.Blue, colour.Reset,
			colour.Blue, colour.Reset,
			colour.Blue, colour.Reset)

		// We should always get user input
		// Get user cli input
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		// We should always handle user input
		// Handle user input
		switch choice {
		case "1":
			if userPet.IsFull() {
				fmt.Println(userPet.DisplayStats("Angry"))
			} else {
				fmt.Println(userPet.DisplayStats("Happy"))
			}
			fmt.Println(userPet.Feed() + "\n")
		case "2":
			if userPet.IsDirty() {
				fmt.Println(userPet.DisplayStats("Happy"))
			} else {
				fmt.Println(userPet.DisplayStats("Angry"))
			}
			fmt.Println(userPet.Clean() + "\n")
		case "3":
			fmt.Println(userPet.DisplayStats("Happy"))
			fmt.Println(userPet.Play() + "\n")
		case "4":
			if userPet.NeedsToPoop() {
				fmt.Println(userPet.DisplayStats("Happy"))
			} else {
				fmt.Println(userPet.DisplayStats("Angry"))
			}
			fmt.Println(colour.Yellow + userPet.Poop() + "\n" + colour.Reset)
		case "5":
			fmt.Println(userPet.Battle(monster.Monster{Name: "Wolf", CurrentHealth: 100, MaxHealth: 100, Attack: 5}) + "\n")
			fmt.Println(colour.Yellow + "Press any key to continue." + colour.Reset)
			reader.ReadString('\n')
		case "6":
			// Save the pet data to the pets.json file
			storage.SavePetData("data/pets.json", userPet)
			fmt.Println(colour.Green + "Pet saved." + colour.Reset)
		case "7":
			fmt.Println(colour.Green + "Goodbye!" + colour.Reset)
			os.Exit(0)
		default:
			fmt.Println("Invalid choice.")
		}
		time.Sleep(2 * time.Second)
		// We should always check if the pet has died
		if userPet.IsDead() {
			fmt.Println(userPet.DisplayStats("Dead"))
			fmt.Printf("%s has died.", userPet.Name)
			os.Exit(0)
		}
	}
}
