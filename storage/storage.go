package storage

import (
	"ACS-4210-Go_Pets/pet"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// LoadPetData loads pet data from JSON file
func LoadPetData(petName string) (*pet.Tamagotchi, error) {
	dataPath, err := getFilePath()
	if err != nil {
		return nil, err
	}

	file, err := os.ReadFile(dataPath)
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

func GetAllPetData() (map[string]pet.Tamagotchi, error) {
	dataPath, err := getFilePath()
	if err != nil {
		return nil, err
	}

	file, err := os.ReadFile(dataPath)
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
func SavePetData(petData pet.Tamagotchi) error {
	var loadedPetData map[string]pet.Tamagotchi

	dataPath, err := getFilePath()
	if err != nil {
		return err
	}

	file, err := os.ReadFile(dataPath)
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
	err = os.WriteFile(dataPath, writeData, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func getFilePath() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	exeDir := filepath.Dir(exePath)
	return filepath.Join(exeDir, "data/pets.json"), nil
}
