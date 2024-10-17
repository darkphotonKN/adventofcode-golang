package designpatterns

import "fmt"

// -- Complex Subsystems --
// Weapons subsystem
type Weapons struct {
	inventory []string
}

func (w *Weapons) AddWeapon(weapon string) {
	w.inventory = append(w.inventory, weapon)
	fmt.Printf("Added weapon: %s\n", weapon)
}

// Potions subsystem
type Potions struct {
	inventory []string
}

func (p *Potions) AddPotion(potion string) {
	p.inventory = append(p.inventory, potion)
	fmt.Printf("Added potion: %s\n", potion)
}

// Armor subsystem
type Armor struct {
	inventory []string
}

func (a *Armor) AddArmor(armor string) {
	a.inventory = append(a.inventory, armor)
	fmt.Printf("Added armor: %s\n", armor)
}

// -- Inventory Facade --
type InventoryFacade struct {
	weapons Weapons
	potions Potions
	armor   Armor
}

func NewInventoryFacade() *InventoryFacade {
	return &InventoryFacade{
		weapons: Weapons{},
		potions: Potions{},
		armor:   Armor{},
	}
}

// Facade to reduce perceived complexity
func (i *InventoryFacade) AddItem(itemType string, item string) {
	switch itemType {
	case "weapon":
		i.weapons.AddWeapon(item)
	case "potion":
		i.potions.AddPotion(item)
	case "armor":
		i.armor.AddArmor(item)
	}
}
