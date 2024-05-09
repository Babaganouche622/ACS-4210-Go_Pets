package pet_test

import (
	"ACS-4210-Go_Pets/monster"
	"ACS-4210-Go_Pets/pet"
	"testing"
)

func TestFeed(t *testing.T) {
	tests := []struct {
		name           string
		initialHunger  int
		expectedHunger int
	}{
		{"Not Full", 50, 0},
		{"Already Full", 0, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tamagotchi := &pet.Tamagotchi{Name: "Test", Hunger: test.initialHunger}
			tamagotchi.Feed()
			if tamagotchi.Hunger != test.expectedHunger {
				t.Errorf("Feed() = ( %d hunger), expected ( %d hunger)", tamagotchi.Hunger, test.expectedHunger)
			}
		})
	}
}

func TestClean(t *testing.T) {
	tests := []struct {
		name              string
		initialDirty      int
		expectedDirty     int
		expectedHappiness int
	}{
		{"Dirty", 50, 0, 100},
		{"Clean", 0, 0, 90},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tamagotchi := &pet.Tamagotchi{Name: "Test", Dirty: test.initialDirty, Happiness: 100}
			tamagotchi.Clean()
			if tamagotchi.Dirty != test.expectedDirty || tamagotchi.Happiness != test.expectedHappiness {
				t.Errorf("Clean() = (dirty: %d, happiness: %d), expected (dirty: %d, happiness: %d)", tamagotchi.Dirty, tamagotchi.Happiness, test.expectedDirty, test.expectedHappiness)
			}
		})
	}
}

func TestPlay(t *testing.T) {
	tests := []struct {
		name              string
		initialHappiness  int
		expectedHappiness int
	}{
		{"Normal", 50, 60},
		{"Max", 90, 100},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tamagotchi := &pet.Tamagotchi{Name: "Test", Happiness: test.initialHappiness}
			tamagotchi.Play()
			if tamagotchi.Happiness != test.expectedHappiness {
				t.Errorf("Play() = %d happiness, expected %d happiness", tamagotchi.Happiness, test.expectedHappiness)
			}
		})
	}
}

func TestPoop(t *testing.T) {
	tests := []struct {
		name          string
		initialPoop   int
		expectedPoop  int
		expectedDirty int
	}{
		{"Needs To Poop", 3, 0, 10},
		{"No Need To Poop", 0, 0, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tamagotchi := &pet.Tamagotchi{Name: "Test", PoopState: test.initialPoop, Dirty: 0}
			tamagotchi.Poop()
			if tamagotchi.PoopState != test.expectedPoop || tamagotchi.Dirty != test.expectedDirty {
				t.Errorf("Poop() = (poop: %d, dirty: %d), expected (poop: %d, dirty: %d)", tamagotchi.PoopState, tamagotchi.Dirty, test.expectedPoop, test.expectedDirty)
			}
		})
	}
}

func TestTakeDamage(t *testing.T) {
	tamagotchi := &pet.Tamagotchi{Name: "Test", CurrentHealth: 100}
	tamagotchi.TakeDamage(20)
	if tamagotchi.CurrentHealth != 80 {
		t.Errorf("TakeDamage() = %d health, expected 80 health", tamagotchi.CurrentHealth)
	}
}

func TestBattle(t *testing.T) {
	// Mock Monster for testing
	mockMonster := monster.Monster{Name: "Mock Monster", CurrentHealth: 100, Attack: 20}

	tests := []struct {
		name             string
		initialHunger    int
		initialHealth    int
		initialDirty     int
		expectedHunger   int
		expectedHealth   int
		expectedDirty    int
		expectedResponse string
	}{
		{"Hungry", 100, 100, 0, 100, 100, 0, "Test is too hungry to battle."},
		{"Normal", 50, 100, 0, 80, 20, 10, "Mock Monster has died."},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tamagotchi := &pet.Tamagotchi{Name: "Test", Hunger: test.initialHunger, CurrentHealth: test.initialHealth, Dirty: test.initialDirty, MaxHealth: 100, Attack: 30}
			response := tamagotchi.Battle(mockMonster)
			if tamagotchi.Hunger != test.expectedHunger || tamagotchi.CurrentHealth != test.expectedHealth || tamagotchi.Dirty != test.expectedDirty || response != test.expectedResponse {
				t.Errorf("Battle() = (%d hunger, %d health, %d dirty, response: %s), expected (%d hunger, %d health, %d dirty, response: %s)", tamagotchi.Hunger, tamagotchi.CurrentHealth, tamagotchi.Dirty, response, test.expectedHunger, test.expectedHealth, test.expectedDirty, test.expectedResponse)
			}
		})
	}
}

func TestIncreaseHunger(t *testing.T) {
	tests := []struct {
		name           string
		initialHunger  int
		increaseAmount int
		expectedHunger int
		expectedHealth int
	}{
		{"Normal", 50, 30, 80, 100},
		{"Max", 90, 30, 100, 90},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tamagotchi := &pet.Tamagotchi{Name: "Test", Hunger: test.initialHunger, CurrentHealth: 100}
			tamagotchi.IncreaseHunger(test.increaseAmount)
			if tamagotchi.Hunger != test.expectedHunger || tamagotchi.CurrentHealth != test.expectedHealth {
				t.Errorf("IncreaseHunger() = (%d hunger, %d health), expected (%d hunger, %d health)", tamagotchi.Hunger, tamagotchi.CurrentHealth, test.expectedHunger, test.expectedHealth)
			}
		})
	}
}

// This test doesn't work right now.
// func TestDisplayStats(t *testing.T) {
// 	tests := []struct {
// 		name             string
// 		initialHunger    int
// 		initialPoopState int
// 		initialDirty     int
// 		initialHappiness int
// 		initialHealth    int
// 		expectedOutput   string
// 	}{
// 		{"Normal", 50, 2, 25, 75, 80, "\033[H\033[2J" + "Current Weather" + "\n" + "\x1b[34mName: \x1b[0m\x1b[35mTest\x1b[0m\n\x1b[34mHunger:\x1b[0m \x1b[32mSatisfied\x1b[0m \x1b[34mPoop level:\x1b[0m \x1b[33mCode Brown\x1b[0m \x1b[34mDirty level:\x1b[0m \x1b[33mDusty\x1b[0m \x1b[34mHappiness:\x1b[0m \x1b[32mContent\x1b[0m \x1b[34mHealth:\x1b[0m 80\n"},
// 		{"Max Values", 100, 3, 100, 100, 100, "\033[H\033[2J" + "Current Weather" + "\n" + "\x1b[34mName: \x1b[0m\x1b[35mTest\x1b[0m\n\x1b[34mHunger:\x1b[0m \x1b[32mFull\x1b[0m \x1b[34mPoop level:\x1b[0m \x1b[31mDefcon Poop\x1b[0m \x1b[34mDirty level:\x1b[0m \x1b[31mFilthy\x1b[0m \x1b[34mHappiness:\x1b[0m \x1b[32mEcstatic\x1b[0m \x1b[34mHealth:\x1b[0m 100\n"},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			tamagotchi := &pet.Tamagotchi{Name: "Test", Hunger: test.initialHunger, PoopState: test.initialPoopState, Dirty: test.initialDirty, Happiness: test.initialHappiness, CurrentHealth: test.initialHealth}
// 			output := tamagotchi.DisplayStats("")
// 			if output != test.expectedOutput {
// 				t.Errorf("DisplayStats() = %s, expected %s", output, test.expectedOutput)
// 			}
// 		})
// 	}
// }

func TestIsFull(t *testing.T) {
	tests := []struct {
		name          string
		initialHunger int
		expected      bool
	}{
		{"Not Full", 50, false},
		{"Full", 0, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tamagotchi := &pet.Tamagotchi{Name: "Test", Hunger: test.initialHunger}
			result := tamagotchi.IsFull()
			if result != test.expected {
				t.Errorf("IsFull() = %v, expected %v", result, test.expected)
			}
		})
	}
}

func TestIsDirty(t *testing.T) {
	tests := []struct {
		name         string
		initialDirty int
		expected     bool
	}{
		{"Not Dirty", 25, false},
		{"Dirty", 50, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tamagotchi := &pet.Tamagotchi{Name: "Test", Dirty: test.initialDirty}
			result := tamagotchi.IsDirty()
			if result != test.expected {
				t.Errorf("IsDirty() = %v, expected %v", result, test.expected)
			}
		})
	}
}

func TestNeedsToPoop(t *testing.T) {
	tests := []struct {
		name        string
		initialPoop int
		expected    bool
	}{
		{"Needs To Poop", 1, true},
		{"No Need To Poop", 0, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tamagotchi := &pet.Tamagotchi{Name: "Test", PoopState: test.initialPoop}
			result := tamagotchi.NeedsToPoop()
			if result != test.expected {
				t.Errorf("NeedsToPoop() = %v, expected %v", result, test.expected)
			}
		})
	}
}
