package main

import (
    "errors"
    "fmt"
    "sort"
    "strconv"
    "strings"
)

// Item represents an element in the catalog
type Item struct {
    Code  int
    Title string
    Cost  float64
    Quantity int
}

// CatalogManager handles all catalog operations
type CatalogManager struct {
    items []Item
}

// NewCatalogManager creates a new instance of CatalogManager
func NewCatalogManager() *CatalogManager {
    return &CatalogManager{
        items: make([]Item, 0),
    }
}

// AddItem adds a new item to the catalog
func (cm *CatalogManager) AddItem(code int, title string, costStr string, quantity int) error {
    // Check for duplicate Code
    for _, item := range cm.items {
        if item.Code == code {
            return fmt.Errorf("item with Code %d already exists", code)
        }
    }

    // Convert string cost to float64
    cost, err := strconv.ParseFloat(costStr, 64)
    if err != nil {
        return errors.New("invalid cost format")
    }

    // Validate inputs
    if cost <= 0 {
        return errors.New("cost must be greater than zero")
    }
    if quantity < 0 {
        return errors.New("quantity cannot be negative")
    }
    if title == "" {
        return errors.New("item title cannot be empty")
    }

    // Create and add new item
    newItem := Item{
        Code:     code,
        Title:    title,
        Cost:     cost,
        Quantity: quantity,
    }

    cm.items = append(cm.items, newItem)
    return nil
}

// UpdateQuantity modifies the quantity of an item
func (cm *CatalogManager) UpdateQuantity(code int, newQuantity int) error {
    if newQuantity < 0 {
        return errors.New("quantity cannot be negative")
    }

    for i := range cm.items {
        if cm.items[i].Code == code {
            cm.items[i].Quantity = newQuantity
            return nil
        }
    }

    return fmt.Errorf("item with Code %d not found", code)
}

// FindByCode searches for an item by Code
func (cm *CatalogManager) FindByCode(code int) (*Item, error) {
    for i := range cm.items {
        if cm.items[i].Code == code {
            return &cm.items[i], nil
        }
    }
    return nil, fmt.Errorf("item with Code %d not found", code)
}

// FindByTitle searches for items by title (case-insensitive partial match)
func (cm *CatalogManager) FindByTitle(title string) []Item {
    var matches []Item
    searchPhrase := strings.ToLower(title)

    for _, item := range cm.items {
        if strings.Contains(strings.ToLower(item.Title), searchPhrase) {
            matches = append(matches, item)
        }
    }

    return matches
}

// ArrangeByCost sorts items by cost
func (cm *CatalogManager) ArrangeByCost() {
    sort.Slice(cm.items, func(i, j int) bool {
        return cm.items[i].Cost < cm.items[j].Cost
    })
}

// ArrangeByQuantity sorts items by quantity
func (cm *CatalogManager) ArrangeByQuantity() {
    sort.Slice(cm.items, func(i, j int) bool {
        return cm.items[i].Quantity < cm.items[j].Quantity
    })
}

// ShowCatalog displays all items in a formatted table
func (cm *CatalogManager) ShowCatalog() {
    if len(cm.items) == 0 {
        fmt.Println("Catalog is empty")
        return
    }

    // Print table header
    fmt.Printf("\n%-6s | %-20s | %-10s | %-9s\n", "Code", "Title", "Cost ($)", "Quantity")
    fmt.Println(strings.Repeat("-", 50))

    // Print each item
    for _, item := range cm.items {
        fmt.Printf("%-6d | %-20s | %10.2f | %-9d\n",
            item.Code, item.Title, item.Cost, item.Quantity)
    }
    fmt.Println()
}

func main() {
    // Create new catalog manager
    catalog := NewCatalogManager()

    // Add sample items
    examples := []struct {
        code     int
        title    string
        cost     string
        quantity int
    }{
        {1001, "Smartphone", "799.99", 25},
        {1002, "Tablet", "399.99", 40},
        {1003, "Charger", "19.99", 200},
        {1004, "Headphones", "49.99", 75},
        {1005, "Backpack", "59.99", 50},
    }

    // Add sample items and handle errors
    for _, example := range examples {
        err := catalog.AddItem(example.code, example.title, example.cost, example.quantity)
        if err != nil {
            fmt.Printf("Error adding %s: %v\n", example.title, err)
        }
    }

    // Display original catalog
    fmt.Println("Original Catalog:")
    catalog.ShowCatalog()

    // Update quantity
    err := catalog.UpdateQuantity(1001, 30)
    if err != nil {
        fmt.Printf("Error updating quantity: %v\n", err)
    }

    // Search by Code
    item, err := catalog.FindByCode(1001)
    if err != nil {
        fmt.Printf("Search error: %v\n", err)
    } else {
        fmt.Printf("\nFound item by Code: %+v\n", *item)
    }

    // Search by title
    searchResults := catalog.FindByTitle("Tab")
    fmt.Printf("\nItems containing 'Tab':\n")
    for _, item := range searchResults {
        fmt.Printf("%+v\n", item)
    }

    // Sort by cost and display
    fmt.Println("\nCatalog sorted by cost:")
    catalog.ArrangeByCost()
    catalog.ShowCatalog()

    // Sort by quantity and display
    fmt.Println("\nCatalog sorted by quantity:")
    catalog.ArrangeByQuantity()
    catalog.ShowCatalog()
}
