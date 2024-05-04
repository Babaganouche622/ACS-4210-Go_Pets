package storage

import (
	"ACS-4210-Go_Pets/pet"
	"encoding/json"
	"fmt"
	"os"
)

// LoadPetData loads pet data from JSON file
func LoadPetData(filename string, petName string) (*pet.Tamagotchi, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var allPetData map[string]*pet.Tamagotchi
	err = json.Unmarshal(file, &allPetData)
	if err != nil {
		return nil, err
	}

	petData, ok := allPetData[petName]
	if !ok {
		return nil, fmt.Errorf("pet %s not found", petName)
	}

	return petData, nil
}

func GetAllPetData(filename string) (map[string]pet.Tamagotchi, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var allPetData map[string]pet.Tamagotchi
	err = json.Unmarshal(file, &allPetData)
	if err != nil {
		return nil, err
	}
	return allPetData, nil
}

// SavePetData saves pet data to JSON file
func SavePetData(filename string, petData pet.Tamagotchi) error {
	var loadedPetData map[string]pet.Tamagotchi

	file, err := os.ReadFile(filename)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		// If the file does not exist, initialize loadedPetData as an empty map
		loadedPetData = make(map[string]pet.Tamagotchi)
	} else {
		// If the file exists, load the existing pet data
		err = json.Unmarshal(file, &loadedPetData)
		if err != nil {
			return err
		}
	}
	// Update the pet data
	loadedPetData[petData.Name] = petData

	// Marshal the updated pet data
	writeData, err := json.MarshalIndent(loadedPetData, "", "  ")
	if err != nil {
		return err
	}

	// Write the updated pet data to the file
	err = os.WriteFile(filename, writeData, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
