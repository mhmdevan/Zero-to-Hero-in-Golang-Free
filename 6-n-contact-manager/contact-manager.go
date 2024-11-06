package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

// Contact represents a contact in the contact manager
type Contact struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

// ContactManager manages a list of contacts
type ContactManager struct {
	contacts []Contact
	mu       sync.Mutex
}

// AddContact adds a new contact
func (cm *ContactManager) AddContact(contact Contact) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.contacts = append(cm.contacts, contact)
}

// ListContacts lists all contacts
func (cm *ContactManager) ListContacts() {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	for _, contact := range cm.contacts {
		fmt.Printf("Name: %s, Phone: %s, Email: %s, Address: %s\n", contact.Name, contact.Phone, contact.Email, contact.Address)
	}
}

// SearchContact searches for a contact by name
func (cm *ContactManager) SearchContact(name string) *Contact {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	for _, contact := range cm.contacts {
		if contact.Name == name {
			return &contact
		}
	}
	return nil
}

// SaveToFile saves contacts to a JSON file
func (cm *ContactManager) SaveToFile(filename string) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	data, err := json.Marshal(cm.contacts)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// LoadFromFile loads contacts from a JSON file
func (cm *ContactManager) LoadFromFile(filename string) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &cm.contacts)
}

func main() {
	cm := &ContactManager{}

	// Load contacts from file if it exists
	err := cm.LoadFromFile("contacts.json")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Error loading contacts:", err)
		return
	}

	for {
		var choice int
		fmt.Println("\nContact Manager")
		fmt.Println("1. Add Contact")
		fmt.Println("2. List Contacts")
		fmt.Println("3. Search Contact")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var contact Contact
			fmt.Print("Enter Name: ")
			fmt.Scan(&contact.Name)
			fmt.Print("Enter Phone: ")
			fmt.Scan(&contact.Phone)
			fmt.Print("Enter Email: ")
			fmt.Scan(&contact.Email)
			fmt.Print("Enter Address: ")
			fmt.Scan(&contact.Address)
			cm.AddContact(contact)
			cm.SaveToFile("contacts.json")
			fmt.Println("Contact added successfully!")

		case 2:
			fmt.Println("Listing Contacts:")
			cm.ListContacts()

		case 3:
			var name string
			fmt.Print("Enter Name to Search: ")
			fmt.Scan(&name)
			contact := cm.SearchContact(name)
			if contact != nil {
				fmt.Printf("Found Contact - Name: %s, Phone: %s, Email: %s, Address: %s\n", contact.Name, contact.Phone, contact.Email, contact.Address)
			} else {
				fmt.Println("Contact not found.")
			}

		case 4:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
