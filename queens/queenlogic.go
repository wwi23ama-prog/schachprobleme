package queens

import "github.com/wwi23ama-prog/schachprobleme/board"

// RowAllowed erwartet ein Spielfeld und eine Zeilennummer.
// Die Funktion liefert true, falls in dieser Zeile noch eine Dame gesetzt werden darf.
func RowAllowed(b board.Board, row int) bool {
	// TODO
	return false
}

// ColAllowed erwartet ein Spielfeld und eine Spaltennummer.
// Die Funktion liefert true, falls in dieser Spalte noch eine Dame gesetzt werden darf.
func ColAllowed(b board.Board, col int) bool {
	// TODO
	return false
}

// DiagRightAllowed erwartet ein Spielfeld und Koordinaten.
// Die Funktion liefert true, falls in der Diagonalen, die durch dieses Feld
// von links oben nach rechts unten verläuft, noch eine Dame gesetzt werden darf.
func DiagRightAllowed(b board.Board, row, col int) bool {
	// TODO
	return false
}

// DiagLeftAllowed erwartet ein Spielfeld und Koordinaten.
// Die Funktion liefert true, falls in der Diagonalen, die durch dieses Feld
// von rechts oben nach links unten verläuft, noch eine Dame gesetzt werden darf.
func DiagLeftAllowed(b board.Board, row, col int) bool {
	// TODO
	return false
}

// QueenAllowed erwartet ein Spielfeld und Koordinaten.
// Die Funktion liefert true, falls an diese Stelle noch eine Dame gesetzt werden darf.
func QueenAllowed(b board.Board, row, col int) bool {
	// TODO
	return false
}
