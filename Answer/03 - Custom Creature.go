package main

import "fmt"

type Creature interface {
	SelfIntroduction()
	NormalAttack()
	SpecialAttack()
}

type Human struct {
	name string
	atk int
	spcAtk int
}

func (this *Human) SelfIntroduction() {
	fmt.Println("I'm a Human, My name is", this.name)
}

func (this *Human)NormalAttack() {
	fmt.Println("Normal Attack: Sword\nDamage", this.atk)
}

func (this *Human)SpecialAttack() {
	fmt.Println("Special Attack: Excalibur\nDamage", this.spcAtk)
}

type Elf struct {
	name string
	atk int
	spcAtk int
}

func (this *Elf) SelfIntroduction() {
	fmt.Println("Nice to meet you, My name is", this.name, "and I'm a Elf.")
}

func (this *Elf)NormalAttack() {
	fmt.Println("Normal Attack: Bow\nDamage", this.atk)
}

func (this *Elf)SpecialAttack() {
	fmt.Println("Special Attack: FireArrow\nDamage", this.spcAtk)
}

type Dwarf struct {
	name string
	atk int
	spcAtk int
}

func (this *Dwarf) SelfIntroduction() {
	fmt.Println("Hello! My name is", this.name, ", a Dwarf")
}

func (this *Dwarf)NormalAttack() {
	fmt.Println("Normal Attack: Hammer\nDamage", this.atk)
}

func (this *Dwarf)SpecialAttack() {
	fmt.Println("Special Attack: Volcanic Eruption\nDamage", this.spcAtk)
}

func CreatureInfo(c Creature) {
	c.SelfIntroduction()
	c.NormalAttack()
	c.SpecialAttack()
}

func main () {
	human := Human{"King Arthur", 100, 1000}
	elf := Elf{"Serenial", 50, 300}
	dwarf := Dwarf{"Excelhand", 150, 750}

	fmt.Println("<< Creature Information Service >>")
	CreatureInfo(&human)
	fmt.Println("===================================")
	CreatureInfo(&elf)
	fmt.Println("===================================")
	CreatureInfo(&dwarf)
}
