package queens

import "github.com/wwi23ama-prog/schachprobleme/board"

// RowAllowed erwartet ein Spielfeld und eine Zeilennummer.
// Die Funktion liefert true, falls in dieser Zeile noch eine Dame gesetzt werden darf.
func RowAllowed(b board.Board, row int) bool {
	// begin:hint
	// Benutzen Sie die Funktion RowContainsOnly aus dem Paket board.
	// Prüfen Sie, ob in der Zeile nur Leerzeichen vorkommen.
	// end:hint
	// begin:solution
	return b.RowContainsOnly(row, " ")
	// end:solution
	// iftask: return false
}

// ColAllowed erwartet ein Spielfeld und eine Spaltennummer.
// Die Funktion liefert true, falls in dieser Spalte noch eine Dame gesetzt werden darf.
func ColAllowed(b board.Board, col int) bool {
	// begin:hint
	// Benutzen Sie die Funktion ColContainsOnly aus dem Paket board.
	// end:hint
	// begin:solution
	return b.ColContainsOnly(col, " ")
	// end:solution
	// iftask: return false
}

// DiagRightAllowed erwartet ein Spielfeld und Koordinaten.
// Die Funktion liefert true, falls in der Diagonalen, die durch dieses Feld
// von links oben nach rechts unten verläuft, noch eine Dame gesetzt werden darf.
func DiagRightAllowed(b board.Board, row, col int) bool {
	// begin:hint
	// Benutzen Sie die Funktion DiagRightContainsOnly aus dem Paket board.
	// Diese Funktion erwartet die Spalte, in der diese Diagonale ganz oben
	// beginnt. Diese Spalte müssen Sie hier berechnen.
	// end:hint
	// begin:solution
	return b.DiagRightContainsOnly(col-row, " ")
	// end:solution
	// iftask: return false
}

// DiagLeftAllowed erwartet ein Spielfeld und Koordinaten.
// Die Funktion liefert true, falls in der Diagonalen, die durch dieses Feld
// von rechts oben nach links unten verläuft, noch eine Dame gesetzt werden darf.
func DiagLeftAllowed(b board.Board, row, col int) bool {
	// begin:hint
	// Benutzen Sie die Funktion DiagLeftContainsOnly aus dem Paket board.
	// end:hint
	// begin:solution
	return b.DiagLeftContainsOnly(col+row, " ")
	// end:solution
	// iftask: return false
}

// QueenAllowed erwartet ein Spielfeld und Koordinaten.
// Die Funktion liefert true, falls an diese Stelle noch eine Dame gesetzt werden darf.
func QueenAllowed(b board.Board, row, col int) bool {
	// begin:hint
	// Eine Dame darf gesetzt werden, wenn die Zeile,
	// die Spalte und beide Diagonalen frei sind.
	return RowAllowed(b, row) &&
		ColAllowed(b, col) &&
		DiagRightAllowed(b, row, col) &&
		DiagLeftAllowed(b, row, col)
}
