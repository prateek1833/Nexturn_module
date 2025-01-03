package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Action constants
const (
	ADD_FUNDS     = 1
	REMOVE_FUNDS  = 2
	CHECK_FUNDS   = 3
	VIEW_LOGS     = 4
	QUIT_SYSTEM   = 5
)

// Record types
const (
	ADD_FUNDS_TYPE   = "ADD_FUNDS"
	REMOVE_FUNDS_TYPE = "REMOVE_FUNDS"
)

// UserAccount represents a user's bank account
type UserAccount struct {
	ProfileID      int
	FullName       string
	CurrentFunds   float64
	ActivityLog    []string
}

// FinancialManager handles bank operations
type FinancialManager struct {
	users          []*UserAccount
	inputReader    *bufio.Scanner
}

// InitializeManager creates a new instance of FinancialManager
func InitializeManager() *FinancialManager {
	return &FinancialManager{
		users:       make([]*UserAccount, 0),
		inputReader: bufio.NewScanner(os.Stdin),
	}
}

// RegisterUser creates a new user account
func (fm *FinancialManager) RegisterUser(profileID int, fullName string) (*UserAccount, error) {
	for _, user := range fm.users {
		if user.ProfileID == profileID {
			return nil, fmt.Errorf("profile with ID %d already exists", profileID)
		}
	}

	newUser := &UserAccount{
		ProfileID:    profileID,
		FullName:     fullName,
		CurrentFunds: 0,
		ActivityLog:  make([]string, 0),
	}

	fm.users = append(fm.users, newUser)
	return newUser, nil
}

// LocateUser retrieves a user account by ProfileID
func (fm *FinancialManager) LocateUser(profileID int) (*UserAccount, error) {
	for _, user := range fm.users {
		if user.ProfileID == profileID {
			return user, nil
		}
	}
	return nil, fmt.Errorf("profile with ID %d not found", profileID)
}

// AddFunds adds money to a user account
func (fm *FinancialManager) AddFunds(profileID int, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	user, err := fm.LocateUser(profileID)
	if err != nil {
		return err
	}

	user.CurrentFunds += amount
	logEntry := fmt.Sprintf("%s: +Rs.%.2f (Funds: Rs.%.2f) - %s",
		ADD_FUNDS_TYPE, amount, user.CurrentFunds, time.Now().Format("2006-01-02 15:04:05"))
	user.ActivityLog = append(user.ActivityLog, logEntry)

	return nil
}

// RemoveFunds withdraws money from a user account
func (fm *FinancialManager) RemoveFunds(profileID int, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	user, err := fm.LocateUser(profileID)
	if err != nil {
		return err
	}

	if user.CurrentFunds < amount {
		return fmt.Errorf("insufficient funds. Available funds: Rs.%.2f", user.CurrentFunds)
	}

	user.CurrentFunds -= amount
	logEntry := fmt.Sprintf("%s: -Rs.%.2f (Funds: Rs.%.2f) - %s",
		REMOVE_FUNDS_TYPE, amount, user.CurrentFunds, time.Now().Format("2006-01-02 15:04:05"))
	user.ActivityLog = append(user.ActivityLog, logEntry)

	return nil
}

// ShowActivityLog displays a user's transaction history
func (fm *FinancialManager) ShowActivityLog(profileID int) error {
	user, err := fm.LocateUser(profileID)
	if err != nil {
		return err
	}

	if len(user.ActivityLog) == 0 {
		fmt.Println("No activity logs found.")
		return nil
	}

	fmt.Printf("\nActivity Logs for Profile %d (%s):\n", user.ProfileID, user.FullName)
	fmt.Println("----------------------------------------")
	for _, log := range user.ActivityLog {
		fmt.Println(log)
	}
	return nil
}

// readInputLine reads input from the user
func (fm *FinancialManager) readInputLine() string {
	fm.inputReader.Scan()
	return strings.TrimSpace(fm.inputReader.Text())
}

// LaunchMenu starts the interactive menu system
func (fm *FinancialManager) LaunchMenu() {
	fmt.Println("Welcome to the Financial Management System!")
	user, err := fm.RegisterUser(101, "Rahul Sharma")
	if err != nil {
		fmt.Printf("Error registering user: %v\n", err)
		return
	}
	fmt.Printf("Registered user: %s (Profile ID: %d)\n\n", user.FullName, user.ProfileID)

	for {
		fmt.Println("\nSelect an option:")
		fmt.Printf("%d. Add Funds\n", ADD_FUNDS)
		fmt.Printf("%d. Withdraw Funds\n", REMOVE_FUNDS)
		fmt.Printf("%d. Check Funds\n", CHECK_FUNDS)
		fmt.Printf("%d. View Logs\n", VIEW_LOGS)
		fmt.Printf("%d. Exit\n", QUIT_SYSTEM)

		choice, err := strconv.Atoi(fm.readInputLine())
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case ADD_FUNDS:
			fmt.Print("Enter amount to add: Rs.")
			amount, err := strconv.ParseFloat(fm.readInputLine(), 64)
			if err != nil {
				fmt.Println("Invalid amount.")
				continue
			}

			if err := fm.AddFunds(101, amount); err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Printf("Successfully added Rs.%.2f\n", amount)
			}

		case REMOVE_FUNDS:
			fmt.Print("Enter amount to withdraw: Rs.")
			amount, err := strconv.ParseFloat(fm.readInputLine(), 64)
			if err != nil {
				fmt.Println("Invalid amount.")
				continue
			}

			if err := fm.RemoveFunds(101, amount); err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Printf("Successfully withdrew Rs.%.2f\n", amount)
			}

		case CHECK_FUNDS:
			user, err := fm.LocateUser(101)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Printf("Available funds: Rs.%.2f\n", user.CurrentFunds)
			}

		case VIEW_LOGS:
			if err := fm.ShowActivityLog(101); err != nil {
				fmt.Printf("Error: %v\n", err)
			}

		case QUIT_SYSTEM:
			fmt.Println("Thank you for using the Financial Management System!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func main() {
	manager := InitializeManager()
	manager.LaunchMenu()
}
