package main

import "fmt"

type Player struct {
	name      string
	invertory []Item
}

type Item struct {
	itemName string
	itemType string
}

func CreatePlayer(playerName string) Player {
	return Player{name: playerName}
}

func CreateItem(itemName string, itemType string) Item {
	return Item{itemName: itemName, itemType: itemType}
}

func (p *Player) PickUpItem(item Item) {
	p.invertory = append(p.invertory, item)
	fmt.Println(p.name + " picked item: " + item.itemName)
}

func (p *Player) removeItem(itemName string) (int, Item) {
	for index, item := range p.invertory {
		if item.itemName == itemName {
			p.invertory = append(p.invertory[:index], p.invertory[index+1:]...)
			return index, item
		}
	}
	return -1, Item{itemName: "", itemType: ""}
}
func (p *Player) DropItem(itemName string) {
	resultIndex, resultItem := p.removeItem(itemName)
	if resultIndex != -1 {
		fmt.Println(p.name + " dropped Item: " + resultItem.itemName)
	} else {
		fmt.Printf("Cannot drop %s as item not found!\n", itemName)
	}
}

func (p *Player) UseItem(itemName string) {
	resultIndex, resultItem := p.removeItem(itemName)
	if resultIndex != -1 {
		switch resultItem.itemName {
		case "potion":
			switch resultItem.itemType {
			case "health":
				fmt.Println(p.name + " used " + resultItem.itemType + " " + resultItem.itemName + ", health increased by 20.")
			case "attack":
				fmt.Println(p.name + " used " + resultItem.itemType + " " + resultItem.itemName + ", attack damage increased by 20.")
			}
		}
	} else {
		fmt.Printf("Cannot use %s as item not found!\n", itemName)
	}
}

func main() {
	p1 := CreatePlayer("Akash")
	p1.PickUpItem(CreateItem("potion", "attack"))
	p1.PickUpItem(CreateItem("potion", "health"))
	p1.PickUpItem(CreateItem("sword", "silver"))
	p1.UseItem("potion")
	p1.DropItem("sword")
	p1.UseItem("potion")
	p1.UseItem("potion")
	p1.DropItem("sword")
}
