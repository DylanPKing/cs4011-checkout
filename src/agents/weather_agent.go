package agents

import (
	"math/rand"
)

// Weather struct that models the weather agent in our application
type Weather struct {
	TimesChangedToday          int
	CustomerPatienceMultiplier float32
	CustomerEntryRate          float32
	Conditions                 *Condition
	CurrentCondition           string
	Seed                       *rand.Source
}

// Condition struct models a map of possible weather outcomes and the assocciated multipliers
type Condition struct {
	possibleConditions map[string]float32
}

// NewWeather creates a new Weather struct NOTE: this returns a pointer
func NewWeather(seed *rand.Source) *Weather {
	// Roll for the set of conditions that the simulation will use
	conditions := NewCondition()
	conditions.setConditions(seed)

	// Create the struct
	weather := Weather{
		TimesChangedToday:          0,
		CustomerPatienceMultiplier: 1,
		CustomerEntryRate:          1,
		Seed:                       seed,
		Conditions:                 conditions,
	}
	return &weather
}

// NewCondition creates a new condition struct to hold the possible weather conditions
func NewCondition() *Condition {
	// Create an empty struct here and populate when creating the weather agent
	conditions := Condition{}
	return &conditions
}

// ToggleWeather tries to change the current weather to a random weather from the chosen set of possible conditions
func (weather *Weather) ToggleWeather() {
	var conditionsArray []string
	randomCondition := (rand.New(*weather.Seed)).Intn(4)

	// Since I can't really index a map, I'm using an array to hold a copy of each key for wasy access
	for i := range weather.Conditions.possibleConditions {
		conditionsArray = append(conditionsArray, i)
	}

	// Set the weather datafields with the new weather information
	weather.CurrentCondition = conditionsArray[randomCondition]
	weather.CustomerPatienceMultiplier = weather.Conditions.possibleConditions[weather.CurrentCondition]
	weather.CustomerEntryRate = weather.Conditions.possibleConditions[weather.CurrentCondition]
	// Incerement the number is weather changes in a given simulation
	weather.TimesChangedToday++
}

// setConditions selects a one set from the available sets of weathers to chose the weather later
func (conditions *Condition) setConditions(seed *rand.Source) {
	// The normal day package
	set1 := make(map[string]float32)
	set1["Sunny"] = 1.0 // base stats
	set1["Overcast"] = 0.8
	set1["Shower"] = 0.7
	set1["Strong Rain"] = 1.25

	// The rainy day package
	set2 := make(map[string]float32)
	set2["Lashing"] = 1.25
	set2["Thunderstorm"] = 1.5
	set2["Shower"] = 0.7
	set2["Cloudy"] = 1.0

	// The winter day package
	set3 := make(map[string]float32)
	set3["Foggy"] = 1.1
	set3["Snowing"] = 1.3
	set3["Blizzard"] = 1.5
	set3["Snow and Rain"] = 1.4

	// Select a random set
	setsArray := [3]map[string]float32{
		set1,
		set2,
		set3,
	}
	conditions.possibleConditions = setsArray[(rand.New(*seed)).Intn(3)]
}
