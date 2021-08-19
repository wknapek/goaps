package queenattack

import "errors"

func CanQueenAttack(w string, b string) (bool, error) {
	if w == b {
		return false, errors.New("invalid input. Same square")
	}
	if len(w) != 2 || len(b) != 2 {
		return false, errors.New("invalid input. Not a square")
	}

	wRank, bRank := w[0]-"a"[0], b[0]-"a"[0]
	wFile, bFile := w[1]-"1"[0], b[1]-"1"[0]

	for _, v := range []byte{wRank, bRank, wFile, bFile} {
		if v < 0 || v > 7 {
			return false, errors.New("Invalid input. Off the board.")
		}
	}

	switch {
	case wRank == bRank || wFile == bFile:
		return true, nil // Axial Attack
	case wRank-bRank == wFile-bFile || wRank-bRank == -(wFile-bFile):
		return true, nil // Diagonal Attack
	default:
		return false, nil
	}
}
