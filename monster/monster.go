package monster

type Monster struct {
	Name          string
	CurrentHealth int
	MaxHealth     int
	Attack        int // 0-100
}

func (m *Monster) TakeDamage(damage int) {
	m.CurrentHealth -= damage
}

func (m *Monster) IsDead() bool {
	return m.CurrentHealth <= 0
}
