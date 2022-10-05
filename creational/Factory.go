package creational

type Attack interface {
	Hit(damage int)
}
type AttackType int

const (
	Physic AttackType = iota
	Spell
)

type spellAttack struct {
}
type physicAttack struct {
}

func (s *spellAttack) Hit(damage int) {
}
func (p *physicAttack) Hit(damage int) {

}

// same as switch(case)
var attackMap = map[AttackType]Attack{
	Physic: new(spellAttack),
	Spell:  new(physicAttack),
}

func ChooseAttack(t AttackType) Attack {
	return attackMap[t]
}
