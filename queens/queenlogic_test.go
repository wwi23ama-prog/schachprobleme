package queens

import (
	"fmt"

	"github.com/wwi23ama-prog/schachprobleme/board"
)

func ExampleRowAllowed() {
	b := board.MakeBoardFromStrings(
		"   ",
		" * ",
		"*  ",
	)
	fmt.Println(RowAllowed(b, 0))
	fmt.Println(RowAllowed(b, 1))
	fmt.Println(RowAllowed(b, 2))

	// Output:
	// true
	// false
	// false
}

func ExampleColAllowed() {
	b := board.MakeBoardFromStrings(
		"   ",
		" * ",
		"*  ",
	)
	fmt.Println(ColAllowed(b, 0))
	fmt.Println(ColAllowed(b, 1))
	fmt.Println(ColAllowed(b, 2))

	// Output:
	// false
	// false
	// true
}

func ExampleDiagRightAllowed() {
	b := board.MakeBoardFromStrings(
		"   ",
		"*  ",
		"*  ",
	)

	// Zeile 0
	fmt.Println(DiagRightAllowed(b, 0, 0))
	fmt.Println(DiagRightAllowed(b, 0, 1))
	fmt.Println(DiagRightAllowed(b, 0, 2))
	fmt.Println()

	// Zeile 1
	fmt.Println(DiagRightAllowed(b, 1, 0))
	fmt.Println(DiagRightAllowed(b, 1, 1))
	fmt.Println(DiagRightAllowed(b, 1, 2))
	fmt.Println()

	// Zeile 2
	fmt.Println(DiagRightAllowed(b, 2, 0))
	fmt.Println(DiagRightAllowed(b, 2, 1))
	fmt.Println(DiagRightAllowed(b, 2, 2))
	fmt.Println()

	// Output:
	// true
	// true
	// true
	//
	// false
	// true
	// true
	//
	// false
	// false
	// true
}

func ExampleDiagLeftAllowed() {
	b := board.MakeBoardFromStrings(
		"   ",
		"*  ",
		"*  ",
	)

	// Zeile 0
	fmt.Println(DiagLeftAllowed(b, 0, 0))
	fmt.Println(DiagLeftAllowed(b, 0, 1))
	fmt.Println(DiagLeftAllowed(b, 0, 2))
	fmt.Println()

	// Zeile 1
	fmt.Println(DiagLeftAllowed(b, 1, 0))
	fmt.Println(DiagLeftAllowed(b, 1, 1))
	fmt.Println(DiagLeftAllowed(b, 1, 2))
	fmt.Println()

	// Zeile 2
	fmt.Println(DiagLeftAllowed(b, 2, 0))
	fmt.Println(DiagLeftAllowed(b, 2, 1))
	fmt.Println(DiagLeftAllowed(b, 2, 2))
	fmt.Println()

	// Output:
	// true
	// false
	// false
	//
	// false
	// false
	// true
	//
	// false
	// true
	// true
}

func ExampleQueenAllowed() {
	b := board.MakeBoardFromStrings(
		"   ",
		"   ",
		"*  ",
	)

	// Zeile 0
	fmt.Println(QueenAllowed(b, 0, 0))
	fmt.Println(QueenAllowed(b, 0, 1))
	fmt.Println(QueenAllowed(b, 0, 2))
	fmt.Println()

	// Zeile 1
	fmt.Println(QueenAllowed(b, 1, 0))
	fmt.Println(QueenAllowed(b, 1, 1))
	fmt.Println(QueenAllowed(b, 1, 2))
	fmt.Println()

	// Zeile 2
	fmt.Println(QueenAllowed(b, 2, 0))
	fmt.Println(QueenAllowed(b, 2, 1))
	fmt.Println(QueenAllowed(b, 2, 2))
	fmt.Println()

	// Output:
	// false
	// true
	// false
	//
	// false
	// false
	// true
	//
	// false
	// false
	// false
}
