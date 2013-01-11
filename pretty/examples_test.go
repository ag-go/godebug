package pretty_test

import (
	"fmt"
	"github.com/kylelemons/godebug/pretty"
)

func ExampleConfig_Sprint() {
	type Pair [2]int
	type Map struct {
		Name      string
		Players   map[string]Pair
		Obstacles map[Pair]string
	}

	m := Map{
		Name: "Rock Creek",
		Players: map[string]Pair{
			"player1": {1, 3},
			"player2": {0, -1},
		},
		Obstacles: map[Pair]string{
			Pair{0, 0}: "rock",
			Pair{2, 1}: "pond",
			Pair{1, 1}: "stream",
			Pair{0, 1}: "stream",
		},
	}

	// Specific output formats
	compact := &pretty.Config{
		Compact: true,
	}
	diffable := &pretty.Config{
		Diffable: true,
	}

	// Print out a summary
	fmt.Printf("Players: %s\n", compact.Sprint(m.Players))

	// Print diffable output
	fmt.Printf("Map State:\n%s", diffable.Sprint(m))

	// Output:
	// Players: {player1:[1,3],player2:[0,-1]}
	// Map State:
	// {
	//  Name:      "Rock Creek",
	//  Players:   {
	//              player1: [
	//                        1,
	//                        3,
	//                       ],
	//              player2: [
	//                        0,
	//                        -1,
	//                       ],
	//             },
	//  Obstacles: {
	//              [0,0]: "rock",
	//              [0,1]: "stream",
	//              [1,1]: "stream",
	//              [2,1]: "pond",
	//             },
	// }
}

func ExampleConfig_Print() {
	type ShipManifest struct {
		Name     string
		Crew     map[string]string
		Androids int
		Stolen   bool
	}

	manifest := &ShipManifest{
		Name: "Spaceship Heart of Gold",
		Crew: map[string]string{
			"Zaphod Beeblebrox": "Galactic President",
			"Trillian":          "Human",
			"Ford Prefect":      "A Hoopy Frood",
			"Arthur Dent":       "Along for the Ride",
		},
		Androids: 1,
		Stolen:   true,
	}

	pretty.Print(manifest)

	// Output:
	// {Name:     "Spaceship Heart of Gold",
	//  Crew:     {Arthur Dent:       "Along for the Ride",
	//             Ford Prefect:      "A Hoopy Frood",
	//             Trillian:          "Human",
	//             Zaphod Beeblebrox: "Galactic President"},
	//  Androids: 1,
	//  Stolen:   true}
}

func ExampleCompare() {
	type ShipManifest struct {
		Name     string
		Crew     map[string]string
		Androids int
		Stolen   bool
	}

	reported := &ShipManifest{
		Name: "Spaceship Heart of Gold",
		Crew: map[string]string{
			"Zaphod Beeblebrox": "Galactic President",
			"Trillian":          "Human",
			"Ford Prefect":      "A Hoopy Frood",
			"Arthur Dent":       "Along for the Ride",
		},
		Androids: 1,
		Stolen:   true,
	}

	expected := &ShipManifest{
		Name: "Spaceship Heart of Gold",
		Crew: map[string]string{
			"Rowan Artosok": "Captain",
		},
		Androids: 1,
		Stolen:   false,
	}

	fmt.Println(pretty.Compare(reported, expected))
	// Output:
	//  {
	//   Name:     "Spaceship Heart of Gold",
	//   Crew:     {
	// -            Arthur Dent:       "Along for the Ride",
	// -            Ford Prefect:      "A Hoopy Frood",
	// -            Trillian:          "Human",
	// -            Zaphod Beeblebrox: "Galactic President",
	// +            Rowan Artosok: "Captain",
	//             },
	//   Androids: 1,
	// - Stolen:   true,
	// + Stolen:   false,
	//  }
}
