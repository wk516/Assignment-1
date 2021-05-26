package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var category = []string{"Household", "Food", "Drinks"}

type itemInformation struct {
	category int
	quantity int
	unitCost float64
}

var entry int            //global variable for menu()
var index int            //global variable for contains()
var categoryCheck string //global variable for contains()

func main() {

	Cups := itemInformation{category: 0, quantity: 5, unitCost: 3}
	Bread := itemInformation{category: 1, quantity: 2, unitCost: 2}
	Fork := itemInformation{category: 0, quantity: 4, unitCost: 3}
	Plates := itemInformation{category: 0, quantity: 4, unitCost: 3}
	Cake := itemInformation{category: 1, quantity: 3, unitCost: 1}
	Coke := itemInformation{category: 2, quantity: 5, unitCost: 2}
	Sprite := itemInformation{category: 2, quantity: 5, unitCost: 2}

	shoppingList := map[string]itemInformation{
		"Cups":   Cups,
		"Bread":  Bread,
		"Fork":   Fork,
		"Plates": Plates,
		"Coke":   Coke,
		"Cake":   Cake,
		"Sprite": Sprite,
	}

	for {
		menu()

		switch entry {
		case 1:
			shoppingListMenu(shoppingList)
		case 2:
			report(shoppingList)
		case 3:
			addItems(shoppingList)
		case 4:
			modify(shoppingList)
		case 5:
			deleteItems(shoppingList)
		case 6:
			data(shoppingList)
		case 7:
			newCategory()
		default:
			fmt.Println("Invalid entry! Please enter a number from 1 to 7.")
		}
	}

}

//1 Shopping list menu
func menu() {
	title := "Shopping List Application"
	fmt.Println(title)
	fmt.Println(strings.Repeat("=", len(title)))
	fmt.Println("1. View Entire Shopping List")
	fmt.Println("2. Generate Shopping List Report")
	fmt.Println("3. Add Item Information.")
	fmt.Println("4. Modify Existing Items.")
	fmt.Println("5. Delete Item from Shopping List.")
	fmt.Println("6. Print Current Data Fields.")
	fmt.Println("7. Add New Category.")
	fmt.Println("Select your choice:")
	fmt.Scanln(&entry)
}

//2 View entire shopping list
func shoppingListMenu(shoppingList map[string]itemInformation) {

	fmt.Println("Shopping list contents:")

	for k, v := range shoppingList {
		fmt.Printf("For %s: Category is %s - Quantity is %d - Unit Cost is %.2f\n", k, category[v.category], v.quantity, v.unitCost)
	}

}

//3 Generate shopping list report
func report(shoppingList map[string]itemInformation) {
	//menu to choose report type
	fmt.Println("Generate report")
	fmt.Println("1. Total cost of each category")
	fmt.Println("2. List of items by category")
	fmt.Println("3. Main menu")
	fmt.Println("Choose your report")
	var reportNumber int
	fmt.Scanln(&reportNumber)

	switch reportNumber {
	case 1:
		//report by category total cost
		fmt.Println("Total cost by each category:")
		var itemCost float64
		var categoryCost float64

		for i := 0; i < len(category); i++ {
			itemCost = 0
			categoryCost = 0
			for _, v := range shoppingList {
				if v.category == i {
					itemCost += v.unitCost
					categoryCost += itemCost
				}
			}
			fmt.Printf("%s cost is %.2f\n", category[i], categoryCost)
		}

	case 2:
		//report by category list
		fmt.Println("List of items by category.")
		for i := 0; i < (len(category) + 1); i++ {
			for k, v := range shoppingList {
				if v.category == i {
					fmt.Printf("Category: %s - Item: %s - Quantiy: %d - Unit Cost: %.2f, at index %d\n", category[v.category], k, v.quantity, v.unitCost, i)
				}
			}
		}

	case 3:
		menu()
	default:
		fmt.Println("Error! Please choose 1, 2, or 3 only")
	}

}

//4 Add items information
func addItems(shoppingList map[string]itemInformation) {
	var nameNew, categoryNew string
	var quantityNew int
	var unitCostNew float64

	fmt.Println("What is the name of your item?")
	fmt.Scanln(&nameNew)

	for k, _ := range shoppingList {
		//check item name doesnt already exists and is not empty field
		if strings.EqualFold(nameNew, k) {
			fmt.Println("Item already exists!")
			break
		} else if nameNew == "" {
			fmt.Println("Nothing is entered.")
			break
		} else {
			//check category. will only proceed with adding item if category already exists
			fmt.Println("What category does it belong to?")
			fmt.Scanln(&categoryNew)

			if contains(category, categoryNew) == true {
				fmt.Println("What is the quantity?")
				fmt.Scanln(&quantityNew)

				fmt.Println("How much does it cost per unit?")
				fmt.Scanln(&unitCostNew)
			} else if contains(category, categoryNew) == false {
				fmt.Printf("Category %s doesn't exist yet?\n", categoryNew)
				fmt.Println("Will you like to add it as new category? Choose Yes or No")
				fmt.Scanln(&categoryCheck)

				if strings.EqualFold(categoryCheck, "Yes") {
					fmt.Println("Please add as new category before adding the new item.")
					newCategory() //redirect to add new category
					break
				} else {
					fmt.Println("Please choose existing category to proceed")
					break
				}
			}

			c := itemInformation{
				category: index,
				quantity: quantityNew,
				unitCost: unitCostNew,
			}

			shoppingList[strings.Title(nameNew)] = c
			fmt.Println("Added new item:", strings.Title(nameNew), shoppingList[strings.Title(nameNew)])
			fmt.Println(shoppingList)
			break
		}
	}

}

func itemExists(item string, shoppingList map[string]itemInformation) bool {
	for existingItem := range shoppingList {
		if strings.EqualFold(item, existingItem) {
			return false
		}
	}
	return true
}

//read the console input from client
func readInput() (input string) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	input = (strings.TrimSpace(strings.Title(strings.ToLower(input))))
	return input
}

//5 Modify Existing Items
func modify(shoppingList map[string]itemInformation) {
	var nameModify, nameModified, categoryModified string
	var quantityModified int
	var unitCostModified float64

	found := false

	fmt.Println("Modify items:")
	fmt.Println("Which item would you like to modify?")
	fmt.Scanln(&nameModify)

	for k, v := range shoppingList {
		// check item name already exisis in the list
		if strings.Title(nameModify) == k {
			found = true
			v = shoppingList[strings.Title(nameModify)]
			fmt.Printf("Current item name is %s - Category is %s - Quantity is %d - Unit Cost is %.2f\n", strings.Title(nameModify), category[v.category], v.quantity, v.unitCost)
			fmt.Println("Enter new Name. Enter for no change.")
			fmt.Scanln(&nameModified)

			if nameModified == "" { //enter for no change
				nameModified = nameModify
				fmt.Printf("No changes to name for %s made. Name still remains as %s\n", strings.Title(nameModify), strings.Title(nameModify))
			} else {
				fmt.Printf("New name for %s will be %s\n", strings.Title(nameModify), nameModified)
				delete(shoppingList, strings.Title(nameModify))
			}

			fmt.Println("Enter new Category. Enter for no change.")
			fmt.Println("Please only enter the following as Category:", category) //let user knows current existing categories
			fmt.Scanln(&categoryModified)
			if categoryModified == "" { //enter for no change
				categoryModified = category[v.category]
				fmt.Printf("No changes to category for %s made. Category still remains as %s. Index is %d.\n", nameModify, category[v.category], v.category)
			} else if contains(category, categoryModified) == false {
				fmt.Printf("Category %s doesn't exist yet?\n", categoryModified)
				fmt.Println("Will you like to add it as new category? Choose Yes or No")
				fmt.Scanln(&categoryCheck)

				if strings.EqualFold(categoryCheck, "Yes") {
					fmt.Println("Please add as new category before adding the new item.")
					newCategory() //redirect to add category
					continue
				} else {
					fmt.Println("Please choose existing category to proceed")
					continue
				}
			} else { //will only proceed with item modification if category entered already exists
				contains(category, categoryModified)
				v.category = index
				fmt.Printf("New category for %s will be %s. Index is %d.\n", nameModify, categoryModified, index)

			}

			fmt.Println("Enter new Quantity. Enter for no change.")
			fmt.Scanln(&quantityModified)

			if quantityModified == 0 {
				quantityModified = v.quantity
				fmt.Printf("No changes to quantity for %s made. Quantity still remains as %d\n", nameModify, quantityModified)
			} else {
				fmt.Printf("New quantity for %s will be %d\n", nameModify, quantityModified)
			}

			fmt.Println("Enter new Unit Cost. Enter for no change.")
			fmt.Scanln(&unitCostModified)

			if unitCostModified == 0 {
				unitCostModified = v.unitCost
				fmt.Printf("No changes to unit cost for %s made. Unit cost still remains as %.2f\n", nameModify, unitCostModified)
			} else {
				fmt.Printf("New unit cost for %s will be %.2f\n", nameModify, unitCostModified)
			}

			shoppingList[strings.Title(nameModified)] = itemInformation{v.category, quantityModified, unitCostModified}
			fmt.Println("Item modified:", strings.Title(nameModified), shoppingList[strings.Title(nameModified)])
			fmt.Println(shoppingList)

		}
	}
	if !found {
		fmt.Println("No such item in the list!")
	}

}

//6 Delete Item from shopping list
func deleteItems(shoppingList map[string]itemInformation) {
	var nameDelete string
	fmt.Println("Delete item:")

	//notify user the existing items
	keys := []string{}
	for key, _ := range shoppingList {
		keys = append(keys, key)
	}
	fmt.Println("")
	fmt.Println("Current list of items are", keys)
	fmt.Println("Choose which item to delete:")
	fmt.Scanln(&nameDelete)

	_, ok := shoppingList[strings.Title(nameDelete)]
	if ok {
		delete(shoppingList, strings.Title(nameDelete))
		fmt.Printf("%s deleted\n", strings.Title(nameDelete))
	} else {
		fmt.Println("No such item in the list. Nothing deleted.")
	}

}

//7 Print current data fields
func data(shoppingList map[string]itemInformation) {
	fmt.Println("Print Current Data.")

	if len(shoppingList) == 0 {
		fmt.Println("No data found!")
	}

	for k, v := range shoppingList {
		fmt.Printf("%s - {%d, %d, %.2f}\n", k, v.category, v.quantity, v.unitCost)
	}
}

//To check if a category already exists and to return index
func contains(category []string, inputCategory string) bool {
	for i, v := range category {
		if strings.EqualFold(v, inputCategory) {
			index = i
			return true
		}
	}
	return false
}

//8 Add New Category Name
func newCategory() {
	fmt.Println("Add new category name,")
	fmt.Println("Current categories:", category)
	fmt.Println("What is the new category name to add?")
	var newCategory string
	fmt.Scanln(&newCategory)

	for i, _ := range category {
		if contains(category, newCategory) == true {
			fmt.Printf("Category: %s already exists at index %d!\n", newCategory, index)
			break
		}
		if contains(category, newCategory) == false {
			if newCategory == "" {
				fmt.Println("No input")
				break
			}
			i = len(category)
			category = append(category, strings.Title(newCategory))
			fmt.Printf("New category: %s added at index %d\n", strings.Title(newCategory), i)
			break
		}
	}
}
