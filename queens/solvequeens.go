package queens

import "github.com/wwi23ama-prog/schachprobleme/board"

// SolveQueensFromRow erwartet ein Spielfeld und eine Zeile.
// Die Funktion versucht rekursiv, das Damenproblem für
// dieses Spielfeld ab dieser Zeile zu lösen.
// Sie liefert true zurück, falls das Problem gelöst werden konnte.
func SolveQueensFromRow(b board.Board, row int) bool {
	// Rekursionsanker: Wenn die Zeile außerhalb des Spielfelds liegt,
	// ist das Problem gelöst.
	// Rekursionsschritt:
	// Für jede Spalte, falls erlaubt:
	// Setze eine Dame und versuche, das Problem ab der nächsten Zeile zu lösen.
	// Gelingt dies, ist das Problem gelöst.
	// Gelingt dies nicht, entferne die Dame wieder und versuche die nächste Spalte.
	if row >= b.Rows() {
		return true
	}

	for col := 0; col < b.Cols(); col++ {
		if QueenAllowed(b, row, col) {
			b[row][col] = "*"
			if SolveQueensFromRow(b, row+1) {
				return true
			}
			b[row][col] = " "
		}
	}

	// Wenn das Problem an dieser Stelle immer noch nicht gelöst wurde,
	// ist es in der aktuellen Konfiguration gar nicht lösbar.
	return false
}

// SolveQueens erwartet ein Spielfeld und löst das Damenproblem für dieses Spielfeld.
// Die Funktion liefert true zurück, falls das Problem gelöst werden konnte.
func SolveQueens(b board.Board) bool {
	// Die eigentliche Arbeit macht die Funktion SolveQueensFromRow.
	// Sie muss mit der ersten Zeile aufgerufen werden.
	//
	// Anmerkung: Dieses Muster kommt bei Rekursion häufig vor:
	// Man hat eine Funktion, die die eigentliche Arbeit macht und
	// eine zweite Funktion, die die erste Funktion mit den richtigen
	// Parametern aufruft, so dass diese Parameter bei der Benutzung
	// der Funktion nicht mehr bekannt sein müssen.
	return SolveQueensFromRow(b, 0)
}
