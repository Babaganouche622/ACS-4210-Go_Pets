// Package monster defines the Monster struct and its associated methods.
package monster

// Monster represents a monster in a game. It has a name, current health, maximum health, and an attack power.
type Monster struct {
	Name          string // Name of the monster
	CurrentHealth int    // Current health of the monster
	MaxHealth     int    // Maximum health of the monster
	Attack        int    // Attack power of the monster, range 0-100
}

// TakeDamage reduces the monster's current health by the given damage amount.
func (m *Monster) TakeDamage(damage int) {
	m.CurrentHealth -= damage
}

// IsDead checks if the monster's current health is 0 or less, indicating that the monster is dead.
func (m *Monster) IsDead() bool {
	return m.CurrentHealth <= 0
}
