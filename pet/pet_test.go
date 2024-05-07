package pet_test

import (
	"testing"

	"ACS-4210-Go_Pets/colour"
	"ACS-4210-Go_Pets/monster"
	"ACS-4210-Go_Pets/pet"
)

// func TestMain(m *testing.M) {
// 	// Suppress console output
// 	old := os.Stdout                   // keep backup of the real stdout
// 	os.Stdout = nil                    // this discards all output
// 	defer func() { os.Stdout = old }() // restore real stdout after tests

// 	// Run the tests
// 	code := m.Run()

// 	// Exit with the code returned from the tests
// 	os.Exit(code)
// }

func TestTamagotchi(t *testing.T) {
	t.Run("Feed", func(t *testing.T) {
		// Test feeding
		tama := &pet.Tamagotchi{Name: "TestPet"}
		tama.Feed()
		if tama.CurrentHealth != tama.MaxHealth || tama.Hunger != 0 {
			t.Errorf("Expected health to be restored and hunger to be reset after feeding.")
		}
	})

	t.Run("Clean", func(t *testing.T) {
		// Test cleaning
		tama := &pet.Tamagotchi{Name: "TestPet"}
		tama.Dirty = 50
		tama.Clean()
		if tama.Dirty != 0 {
			t.Errorf("Expected pet to be clean after cleaning.")
		}
	})

	t.Run("Play", func(t *testing.T) {
		// Test playing
		tama := &pet.Tamagotchi{Name: "TestPet"}
		initialHappiness := tama.Happiness
		tama.Play()
		if tama.Happiness != initialHappiness+10 {
			t.Errorf("Expected happiness to increase after playing.")
		}
	})

	t.Run("Poop", func(t *testing.T) {
		// Test pooping
		tama := &pet.Tamagotchi{Name: "TestPet"}
		tama.PoopState = 3
		tama.Poop()
		if tama.PoopState != 0 {
			t.Errorf("Expected poop state to be reset after pooping.")
		}
	})

	t.Run("TakeDamage", func(t *testing.T) {
		// Test taking damage
		tama := &pet.Tamagotchi{Name: "TestPet", CurrentHealth: 100}
		initialHealth := tama.CurrentHealth
		tama.TakeDamage(20)
		if tama.CurrentHealth != initialHealth-20 {
			t.Errorf("Expected health to decrease after taking damage.")
		}
	})

	t.Run("Battle", func(t *testing.T) {
		// Test battling
		tama := &pet.Tamagotchi{Name: "TestPet", Hunger: 50, CurrentHealth: 100}
		monster := &monster.Monster{Name: "TestMonster", Attack: 10, CurrentHealth: 50}
		result := tama.Battle(*monster)
		if result == "" {
			t.Errorf("Expected battle result to be returned.")
		}
	})

	t.Run("IsDead", func(t *testing.T) {
		// Test IsDead method
		tama := &pet.Tamagotchi{Name: "TestPet", CurrentHealth: 0}
		if !tama.IsDead() {
			t.Errorf("Expected pet to be dead.")
		}
	})

	t.Run("IncreaseHunger", func(t *testing.T) {
		// Test IncreaseHunger method
		tama := &pet.Tamagotchi{Name: "TestPet"}
		initialHunger := tama.Hunger
		tama.IncreaseHunger(20)
		if tama.Hunger != initialHunger+20 {
			t.Errorf("Expected hunger to increase after feeding.")
		}
	})

	t.Run("DisplayStats", func(t *testing.T) {
		// Test DisplayStats method
		tama := &pet.Tamagotchi{Name: "TestPet", Hunger: 75, PoopState: 2, Dirty: 60, Happiness: 80, CurrentHealth: 90}
		stats := tama.DisplayStats("Idle")
		expectedStats := "\033[H\033[2J" +
			colour.Blue + "Name: " + colour.Reset + colour.Purple + "TestPet" + colour.Reset + "\n" +
			colour.Blue + "Hunger:" + colour.Reset + " " + colour.Red + "Starving" + colour.Reset + " " +
			colour.Blue + "Poop level:" + colour.Reset + " " + colour.Orange + "Code Brown" + colour.Reset + " " +
			colour.Blue + "Dirty level:" + colour.Reset + " " + colour.Orange + "Dirty" + colour.Reset + " " +
			colour.Blue + "Happiness:" + colour.Reset + " " + colour.Yellow + "Happy" + colour.Reset + " " +
			colour.Blue + "Health:" + colour.Reset + " 90\n" +
			tama.Display("Idle")
		if stats != expectedStats {
			t.Errorf("Expected display stats to match.")
		}
	})

	t.Run("Display", func(t *testing.T) {
		// Test Display method
		tama := &pet.Tamagotchi{Name: "TestPet"}
		display := tama.Display("Happy")
		expectedDisplay := colour.Green + pet.Happy + colour.Reset + colour.Yellow + "TestPet had a good time!" + colour.Reset
		if display != expectedDisplay {
			t.Errorf("Expected display to match.")
		}
	})

	t.Run("IsFull", func(t *testing.T) {
		// Test IsFull method
		tama := &pet.Tamagotchi{Name: "TestPet", Hunger: 0}
		if !tama.IsFull() {
			t.Errorf("Expected pet to be full.")
		}
	})

	t.Run("IsDirty", func(t *testing.T) {
		// Test IsDirty method
		tama := &pet.Tamagotchi{Name: "TestPet", Dirty: 60}
		if !tama.IsDirty() {
			t.Errorf("Expected pet to be dirty.")
		}
	})

	t.Run("NeedsToPoop", func(t *testing.T) {
		// Test NeedsToPoop method
		tama := &pet.Tamagotchi{Name: "TestPet", PoopState: 1}
		if !tama.NeedsToPoop() {
			t.Errorf("Expected pet to need to poop.")
		}
	})
}
