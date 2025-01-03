package main

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

// WeatherInfo represents meteorological data for a location
type WeatherInfo struct {
    Location      string
    AvgTemp       float64 // in Celsius
    Precipitation float64 // in millimeters
}

// WeatherAnalyzer handles analysis operations on weather data
type WeatherAnalyzer struct {
    locations []WeatherInfo
    input     *bufio.Scanner
}

// NewWeatherAnalyzer initializes a new instance with mock data
func NewWeatherAnalyzer() *WeatherAnalyzer {
    // Mock data for various locations
    mockData := []WeatherInfo{
        {"Springfield", 22.5, 1340.0},
        {"Rivertown", 24.3, 980.0},
        {"Hillview", 19.1, 780.0},
        {"Lakeland", 23.7, 1560.0},
        {"Greenfield", 21.4, 1120.0},
        {"Bayport", 25.8, 890.0},
        {"Woodside", 20.2, 730.0},
        {"Meadowbrook", 23.0, 920.0},
    }

    return &WeatherAnalyzer{
        locations: mockData,
        input:     bufio.NewScanner(os.Stdin),
    }
}

// AddLocation adds a new location to the dataset
func (wa *WeatherAnalyzer) AddLocation(name string, temp, precipitation float64) error {
    // Validate input
    if name == "" {
        return errors.New("location name cannot be empty")
    }
    if temp < -50 || temp > 50 {
        return fmt.Errorf("invalid temperature: %.2f°C (must be between -50°C and 50°C)", temp)
    }
    if precipitation < 0 {
        return fmt.Errorf("invalid precipitation: %.2f mm (must be non-negative)", precipitation)
    }

    // Check for duplicate location
    for _, loc := range wa.locations {
        if strings.EqualFold(loc.Location, name) {
            return fmt.Errorf("location '%s' already exists in the dataset", name)
        }
    }

    newLocation := WeatherInfo{
        Location:      name,
        AvgTemp:       temp,
        Precipitation: precipitation,
    }
    wa.locations = append(wa.locations, newLocation)
    return nil
}

// FindHighestTemperature returns the location with the highest average temperature
func (wa *WeatherAnalyzer) FindHighestTemperature() (*WeatherInfo, error) {
    if len(wa.locations) == 0 {
        return nil, errors.New("no locations in dataset")
    }

    highest := &wa.locations[0]
    for i := range wa.locations {
        if wa.locations[i].AvgTemp > highest.AvgTemp {
            highest = &wa.locations[i]
        }
    }
    return highest, nil
}

// FindLowestTemperature returns the location with the lowest average temperature
func (wa *WeatherAnalyzer) FindLowestTemperature() (*WeatherInfo, error) {
    if len(wa.locations) == 0 {
        return nil, errors.New("no locations in dataset")
    }

    lowest := &wa.locations[0]
    for i := range wa.locations {
        if wa.locations[i].AvgTemp < lowest.AvgTemp {
            lowest = &wa.locations[i]
        }
    }
    return lowest, nil
}

// CalculateAveragePrecipitation computes the mean precipitation across all locations
func (wa *WeatherAnalyzer) CalculateAveragePrecipitation() (float64, error) {
    if len(wa.locations) == 0 {
        return 0, errors.New("no locations in dataset")
    }

    var total float64
    for _, loc := range wa.locations {
        total += loc.Precipitation
    }
    return total / float64(len(wa.locations)), nil
}

// FilterLocationsByPrecipitation returns locations with precipitation above a threshold
func (wa *WeatherAnalyzer) FilterLocationsByPrecipitation(threshold float64) []WeatherInfo {
    var filtered []WeatherInfo
    for _, loc := range wa.locations {
        if loc.Precipitation > threshold {
            filtered = append(filtered, loc)
        }
    }

    // Sort by precipitation in descending order
    sort.Slice(filtered, func(i, j int) bool {
        return filtered[i].Precipitation > filtered[j].Precipitation
    })

    return filtered
}

// SearchLocation finds a location by name (case-insensitive)
func (wa *WeatherAnalyzer) SearchLocation(name string) (*WeatherInfo, error) {
    if name == "" {
        return nil, errors.New("location name cannot be empty")
    }

    for i := range wa.locations {
        if strings.EqualFold(wa.locations[i].Location, name) {
            return &wa.locations[i], nil
        }
    }
    return nil, fmt.Errorf("location '%s' not found", name)
}

// DisplayAllLocations prints all locations in a formatted table
func (wa *WeatherAnalyzer) DisplayAllLocations() {
    fmt.Printf("\n%-15s | %12s | %12s\n", "Location", "Temperature", "Precipitation")
    fmt.Println(strings.Repeat("-", 45))

    for _, loc := range wa.locations {
        fmt.Printf("%-15s | %9.1f°C | %8.1f mm\n",
            loc.Location, loc.AvgTemp, loc.Precipitation)
    }
    fmt.Println()
}

func main() {
    analyzer := NewWeatherAnalyzer()
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Println("\nWeather Data Analysis Tool")
        fmt.Println("1. Display All Locations")
        fmt.Println("2. Add New Location")
        fmt.Println("3. Find Temperature Extremes")
        fmt.Println("4. Calculate Average Precipitation")
        fmt.Println("5. Filter Locations by Precipitation")
        fmt.Println("6. Search Location")
        fmt.Println("7. Exit")
        fmt.Print("Enter your choice (1-7): ")

        scanner.Scan()
        choice := strings.TrimSpace(scanner.Text())

        switch choice {
        case "1":
            analyzer.DisplayAllLocations()

        case "2":
            fmt.Print("Enter location name: ")
            scanner.Scan()
            name := strings.TrimSpace(scanner.Text())

            fmt.Print("Enter average temperature (°C): ")
            scanner.Scan()
            temp, err := strconv.ParseFloat(scanner.Text(), 64)
            if err != nil {
                fmt.Println("Invalid temperature value")
                continue
            }

            fmt.Print("Enter precipitation (mm): ")
            scanner.Scan()
            precipitation, err := strconv.ParseFloat(scanner.Text(), 64)
            if err != nil {
                fmt.Println("Invalid precipitation value")
                continue
            }

            if err := analyzer.AddLocation(name, temp, precipitation); err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                fmt.Println("Location added successfully!")
            }

        case "3":
            highest, err := analyzer.FindHighestTemperature()
            if err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                fmt.Printf("\nHighest Temperature: %s (%.1f°C)\n",
                    highest.Location, highest.AvgTemp)
            }

            lowest, err := analyzer.FindLowestTemperature()
            if err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                fmt.Printf("Lowest Temperature: %s (%.1f°C)\n",
                    lowest.Location, lowest.AvgTemp)
            }

        case "4":
            avg, err := analyzer.CalculateAveragePrecipitation()
            if err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                fmt.Printf("\nAverage Precipitation across all locations: %.1f mm\n", avg)
            }

        case "5":
            fmt.Print("Enter precipitation threshold (mm): ")
            scanner.Scan()
            threshold, err := strconv.ParseFloat(scanner.Text(), 64)
            if err != nil {
                fmt.Println("Invalid threshold value")
                continue
            }

            locations := analyzer.FilterLocationsByPrecipitation(threshold)
            if len(locations) == 0 {
                fmt.Printf("No locations found with precipitation above %.1f mm\n", threshold)
            } else {
                fmt.Printf("\nLocations with precipitation above %.1f mm:\n", threshold)
                for _, loc := range locations {
                    fmt.Printf("%-15s: %.1f mm\n", loc.Location, loc.Precipitation)
                }
            }

        case "6":
            fmt.Print("Enter location name to search: ")
            scanner.Scan()
            name := strings.TrimSpace(scanner.Text())

            location, err := analyzer.SearchLocation(name)
            if err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                fmt.Printf("\nLocation: %s\n", location.Location)
                fmt.Printf("Average Temperature: %.1f°C\n", location.AvgTemp)
                fmt.Printf("Precipitation: %.1f mm\n", location.Precipitation)
            }

        case "7":
            fmt.Println("Thank you for using the Weather Data Analysis Tool!")
            return

        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}
