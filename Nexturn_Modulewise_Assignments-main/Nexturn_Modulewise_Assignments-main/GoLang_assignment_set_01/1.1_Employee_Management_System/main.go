package main

import (
	"errors"
	"fmt"
	"strings"
)

// Constants for teams
const (
	ADMIN_TEAM   = "ADMIN"
	DEVELOP_TEAM = "DEVELOPMENT"
	FINANCE_TEAM = "FINANCE"
)

// Member struct to hold team member information
type Member struct {
	ID        int
	FullName  string
	Years     int
	Team      string
}

// TeamManager handles all team member operations
type TeamManager struct {
	members []Member
}

// NewTeamManager creates a new instance of TeamManager
func NewTeamManager() *TeamManager {
	return &TeamManager{
		members: make([]Member, 0),
	}
}

// AddMember adds a new member after validation
func (tm *TeamManager) AddMember(id int, fullName string, years int, team string) error {
	// Validate years
	if years < 18 {
		return errors.New("member must be at least 18 years old")
	}

	// Validate team
	team = strings.ToUpper(team)
	if team != ADMIN_TEAM && team != DEVELOP_TEAM && team != FINANCE_TEAM {
		return fmt.Errorf("invalid team: %s", team)
	}

	// Check for duplicate ID
	for _, m := range tm.members {
		if m.ID == id {
			return fmt.Errorf("member with ID %d already exists", id)
		}
	}

	// Create and add new member
	newMember := Member{
		ID:        id,
		FullName:  fullName,
		Years:     years,
		Team:      team,
	}

	tm.members = append(tm.members, newMember)
	return nil
}

// SearchByID searches for a member by their ID
func (tm *TeamManager) SearchByID(id int) (*Member, error) {
	for i := range tm.members {
		if tm.members[i].ID == id {
			return &tm.members[i], nil
		}
	}
	return nil, fmt.Errorf("member with ID %d not found", id)
}

// SearchByName searches for a member by their name
func (tm *TeamManager) SearchByName(name string) ([]*Member, error) {
	var found []*Member
	name = strings.ToLower(name)

	for i := range tm.members {
		if strings.Contains(strings.ToLower(tm.members[i].FullName), name) {
			found = append(found, &tm.members[i])
		}
	}

	if len(found) == 0 {
		return nil, fmt.Errorf("no members found with name containing '%s'", name)
	}
	return found, nil
}

// ListByTeam returns all members in a given team
func (tm *TeamManager) ListByTeam(team string) ([]*Member, error) {
	team = strings.ToUpper(team)
	var teamMembers []*Member

	for i := range tm.members {
		if tm.members[i].Team == team {
			teamMembers = append(teamMembers, &tm.members[i])
		}
	}

	if len(teamMembers) == 0 {
		return nil, fmt.Errorf("no members found in team %s", team)
	}
	return teamMembers, nil
}

// CountByTeam returns the number of members in a team
func (tm *TeamManager) CountByTeam(team string) int {
	team = strings.ToUpper(team)
	count := 0

	for _, m := range tm.members {
		if m.Team == team {
			count++
		}
	}
	return count
}

func main() {
	// Create new team manager
	manager := NewTeamManager()

	// Example
	fmt.Println("Adding team members...")

	// Add some members
	errors := []error{
		manager.AddMember(101, "Aman Singh", 29, "DEVELOPMENT"),
		manager.AddMember(102, "Priya Sharma", 24, "ADMIN"),
		manager.AddMember(103, "Rohan Gupta", 34, "DEVELOPMENT"),
		manager.AddMember(104, "Sonal Jain", 27, "FINANCE"),
	}

	// Check for errors during addition
	for _, err := range errors {
		if err != nil {
			fmt.Printf("Error adding member: %v\n", err)
		}
	}

	// Try to add a member with duplicate ID
	err := manager.AddMember(101, "Duplicate User", 22, "ADMIN")
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}

	// Search by ID
	member, err := manager.SearchByID(102)
	if err != nil {
		fmt.Printf("Search error: %v\n", err)
	} else {
		fmt.Printf("Found member: %+v\n", *member)
	}

	// Search by name
	members, err := manager.SearchByName("rohan")
	if err != nil {
		fmt.Printf("Search error: %v\n", err)
	} else {
		fmt.Println("Members found by name:")
		for _, m := range members {
			fmt.Printf("%+v\n", *m)
		}
	}

	// List DEVELOPMENT team members
	devTeam, err := manager.ListByTeam("DEVELOPMENT")
	if err != nil {
		fmt.Printf("List error: %v\n", err)
	} else {
		fmt.Println("\nDEVELOPMENT Team members:")
		for _, m := range devTeam {
			fmt.Printf("%+v\n", *m)
		}
	}

	// Count members by team
	fmt.Printf("\nMember counts by team:\n")
	fmt.Printf("DEVELOPMENT: %d\n", manager.CountByTeam(DEVELOP_TEAM))
	fmt.Printf("ADMIN: %d\n", manager.CountByTeam(ADMIN_TEAM))
	fmt.Printf("Finance: %d\n", manager.CountByTeam(FINANCE_TEAM))
}
