package pet

import (
	"ACS-4210-Go_Pets/colour"
	"ACS-4210-Go_Pets/monster"
	"fmt"
)

type Tamagotchi struct {
	Name          string
	Happiness     int // 0-100
	CurrentHealth int
	MaxHealth     int
	Attack        int // 0-100
	Dirty         int // 0-100
	PoopState     int // 0-4
	Hunger        int // 0-100
	Age           int // 0-100
}

const (
	Egg = `
     _______
    /       \
   /         \
  /           \
 |             |
  \           /
   \_________/
`
	Idle = `
     .-"''''"-.
   .'           '.
  /  o       o   \
 |                 |
 |    -------     |
(_____/     \_____)
     |       |
    (|_____|)
`
	Angry = `
     .-"''''"-.
   .' \     /   '.
  /   o     o    \
 |     \_U_/      |
 |       |        |
 (_____/   \_____)
    / \__ __/\ 
   (/       \)
`
	Dead = `
     .-"''''"-.
   .'           '.
  /  x       x   \
 |                 |
 |    -------     |
(_____/     \_____)
     |  ---  |
    (|_____|)
`
	Happy = `
     .-"''''"-.
   .'           '.
  /   ^     ^    \
 |     \___/      |
 |       -        |
(_____/     \_____)
    \  ' '  /
     \_____/
 `
)

func (t *Tamagotchi) Feed() string {
	if !t.IsFull() {
		t.CurrentHealth = t.MaxHealth
	}
	t.Hunger = 0
	t.PoopState++
	return fmt.Sprintf(colour.Yellow + "You fed " + t.Name + "!" + colour.Reset)
}

func (t *Tamagotchi) Clean() string {
	t.Dirty = 0
	return fmt.Sprintf(colour.Yellow+"%s is now clean."+colour.Reset, t.Name)
}

func (t *Tamagotchi) Play() string {
	t.Happiness += 10
	if t.Happiness > 100 {
		t.Happiness = 100
	}
	return fmt.Sprintf(colour.Green+"%s played!"+colour.Reset, t.Name)
}

func (t *Tamagotchi) Poop() string {
	if t.NeedsToPoop() {
		t.PoopState = 0
		return fmt.Sprintf("%s pooped.", t.Name)
	} else {
		return fmt.Sprintf("%s doesn't need to poop.", t.Name)
	}
}

func (t *Tamagotchi) TakeDamage(damage int) {
	t.CurrentHealth -= damage
}

func (t *Tamagotchi) Battle(monster monster.Monster) string {
	if t.Hunger >= 100 {
		return fmt.Sprintf("%s is too hungry to battle.", t.Name)
	}

	for !t.IsDead() && !monster.IsDead() {
		t.TakeDamage(monster.Attack)
		monster.TakeDamage(t.Attack)
		fmt.Printf("%s%s%s took %s%d%s damage!\n", colour.Green, t.Name, colour.Reset, colour.Red, monster.Attack, colour.Reset)
		fmt.Printf("%s%s%s took %s%d%s damage!\n", colour.Green, monster.Name, colour.Reset, colour.Green, t.Attack, colour.Reset)
	}

	if t.IsDead() {
		return fmt.Sprintf("%s has died.", t.Name)
	}

	// End of combat status changes.
	t.Dirty += 10
	t.IncreaseHunger(30)

	return fmt.Sprintf("%s has died.", monster.Name)
}

func (t *Tamagotchi) IsDead() bool {
	return t.CurrentHealth <= 0
}

func (t *Tamagotchi) IncreaseHunger(amount int) {
	t.Hunger += amount
	if t.Hunger >= 100 {
		t.Hunger = 100
		t.TakeDamage(10)
	}
}

func (t *Tamagotchi) DisplayStats(state string) string {
	// List of Hunger states
	hungerStates := []string{colour.Green + "Full" + colour.Reset, colour.Yellow + "Satisfied" + colour.Reset, colour.Orange + "Hungry" + colour.Reset, colour.Red + "Starving" + colour.Reset}
	hungerState := hungerStates[t.Hunger/25]

	// List of Poop states:
	poopStates := []string{colour.Green + "Poopless Bliss" + colour.Reset, colour.Yellow + "Pootential Building" + colour.Reset, colour.Orange + "Code Brown" + colour.Reset, colour.Red + "Defcon Poop" + colour.Reset}
	poopState := poopStates[t.PoopState]

	// List of Dirty states
	dirtyStates := []string{colour.Green + "Clean" + colour.Reset, colour.Yellow + "Dusty" + colour.Reset, colour.Orange + "Dirty" + colour.Reset, colour.Red + "Filthy" + colour.Reset}
	dirtyState := dirtyStates[t.Dirty/25]

	// List of Happiness states
	happinessStates := []string{colour.Red + "Sad" + colour.Reset, colour.Orange + "Content" + colour.Reset, colour.Yellow + "Happy" + colour.Reset, colour.Green + "Ecstatic" + colour.Reset}
	happinessState := happinessStates[t.Happiness/25-1]

	return fmt.Sprintf(
		"\033[H\033[2J"+
			colour.Blue+"Name: "+colour.Reset+colour.Purple+t.Name+colour.Reset+"\n"+
			"%sHunger:%s %s %sPoop level:%s %s %sDirty level:%s %s %sHappiness:%s %s %sHealth:%s %d\n"+t.Display(state),
		colour.Blue, colour.Reset, hungerState,
		colour.Blue, colour.Reset, poopState,
		colour.Blue, colour.Reset, dirtyState,
		colour.Blue, colour.Reset, happinessState,
		colour.Blue, colour.Reset, t.CurrentHealth,
	)
}

// Based on state change text line to display mood of pet
// Idle: <name of pet> is just chilling.
// Angry: <name of pet> did not like that!
// Dead: <name of pet> has died.
// Happy: <name of pet> had a good time!
func (t *Tamagotchi) Display(state string) string {
	switch state {
	case "Egg":
		return colour.Orange + Egg + colour.Reset + colour.Yellow + "Your egg is incubating!" + colour.Reset
	case "Idle":
		return colour.Purple + Idle + colour.Reset + colour.Yellow + t.Name + " is just chilling." + colour.Reset
	case "Angry":
		return colour.Red + Angry + colour.Reset + colour.Yellow + t.Name + " did not like that!" + colour.Reset
	case "Dead":
		return colour.Red + Dead + colour.Reset + colour.Yellow + t.Name + " has died." + colour.Reset
	case "Happy":
		return colour.Green + Happy + colour.Reset + colour.Yellow + t.Name + " had a good time!" + colour.Reset
	default:
		return colour.Purple + Idle + colour.Reset + colour.Yellow + t.Name + " is just chilling." + colour.Reset
	}
}

func (t *Tamagotchi) IsFull() bool {
	return t.Hunger == 0
}

func (t *Tamagotchi) IsDirty() bool {
	return t.Dirty >= 50
}

func (t *Tamagotchi) NeedsToPoop() bool {
	return t.PoopState >= 1
}
