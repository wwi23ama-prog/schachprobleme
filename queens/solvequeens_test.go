package queens

import (
	"fmt"
	"strings"

	"github.com/wwi23ama-prog/schachprobleme/board"
)

func ExampleSolveQueens_size3() {
	b := MakeEmptyBoard(3)
	success := SolveQueens(b)
	if success {
		fmt.Println(b)
	} else {
		fmt.Println("Keine Lösung")
	}

	// Output:
	// Keine Lösung
}

func ExampleSolveQueens_size4() {
	b := MakeEmptyBoard(4)
	success := SolveQueens(b)
	if success {
		fmt.Println(b)
	} else {
		fmt.Println("Keine Lösung")
	}

	// Output:
	// +---+---+---+---+
	// |   | * |   |   |
	// +---+---+---+---+
	// |   |   |   | * |
	// +---+---+---+---+
	// | * |   |   |   |
	// +---+---+---+---+
	// |   |   | * |   |
	// +---+---+---+---+
}

func ExampleSolveQueens_size5() {
	b := MakeEmptyBoard(5)
	success := SolveQueens(b)
	if success {
		fmt.Println(b)
	} else {
		fmt.Println("Keine Lösung")
	}

	// Output:
	// +---+---+---+---+---+
	// | * |   |   |   |   |
	// +---+---+---+---+---+
	// |   |   | * |   |   |
	// +---+---+---+---+---+
	// |   |   |   |   | * |
	// +---+---+---+---+---+
	// |   | * |   |   |   |
	// +---+---+---+---+---+
	// |   |   |   | * |   |
	// +---+---+---+---+---+
}

func ExampleSolveQueens_size6() {
	b := MakeEmptyBoard(6)
	success := SolveQueens(b)
	if success {
		fmt.Println(b)
	} else {
		fmt.Println("Keine Lösung")
	}

	// Output:
	// +---+---+---+---+---+---+
	// |   | * |   |   |   |   |
	// +---+---+---+---+---+---+
	// |   |   |   | * |   |   |
	// +---+---+---+---+---+---+
	// |   |   |   |   |   | * |
	// +---+---+---+---+---+---+
	// | * |   |   |   |   |   |
	// +---+---+---+---+---+---+
	// |   |   | * |   |   |   |
	// +---+---+---+---+---+---+
	// |   |   |   |   | * |   |
	// +---+---+---+---+---+---+
}

func ExampleSolveQueens_size7() {
	b := MakeEmptyBoard(7)
	success := SolveQueens(b)
	if success {
		fmt.Println(b)
	} else {
		fmt.Println("Keine Lösung")
	}

	// Output:
	// +---+---+---+---+---+---+---+
	// | * |   |   |   |   |   |   |
	// +---+---+---+---+---+---+---+
	// |   |   | * |   |   |   |   |
	// +---+---+---+---+---+---+---+
	// |   |   |   |   | * |   |   |
	// +---+---+---+---+---+---+---+
	// |   |   |   |   |   |   | * |
	// +---+---+---+---+---+---+---+
	// |   | * |   |   |   |   |   |
	// +---+---+---+---+---+---+---+
	// |   |   |   | * |   |   |   |
	// +---+---+---+---+---+---+---+
	// |   |   |   |   |   | * |   |
	// +---+---+---+---+---+---+---+
}

func ExampleSolveQueens_size8() {
	b := MakeEmptyBoard(8)
	success := SolveQueens(b)
	if success {
		fmt.Println(b)
	} else {
		fmt.Println("Keine Lösung")
	}

	// Output:
	// +---+---+---+---+---+---+---+---+
	// | * |   |   |   |   |   |   |   |
	// +---+---+---+---+---+---+---+---+
	// |   |   |   |   | * |   |   |   |
	// +---+---+---+---+---+---+---+---+
	// |   |   |   |   |   |   |   | * |
	// +---+---+---+---+---+---+---+---+
	// |   |   |   |   |   | * |   |   |
	// +---+---+---+---+---+---+---+---+
	// |   |   | * |   |   |   |   |   |
	// +---+---+---+---+---+---+---+---+
	// |   |   |   |   |   |   | * |   |
	// +---+---+---+---+---+---+---+---+
	// |   | * |   |   |   |   |   |   |
	// +---+---+---+---+---+---+---+---+
	// |   |   |   | * |   |   |   |   |
	// +---+---+---+---+---+---+---+---+
}

// Hilfsfunktion: Liefert ein Board der Größe nxn.
func MakeEmptyBoard(n int) board.Board {
	row := strings.Repeat(" ", n)
	rows := make([]string, n)
	for i := range rows {
		rows[i] = row
	}

	return board.MakeBoardFromStrings(rows...)
}
