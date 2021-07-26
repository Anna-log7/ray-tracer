package matrices

import (
	"testing"
)

func Test4By4Construction(t *testing.T) {
	matrix := Matrix{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	}

	if !(matrix[0][0] == 1 && matrix[0][3] == 4 && matrix[1][0] == 5.5 &&
		matrix[1][2] == 7.5 && matrix[2][2] == 11 && matrix[3][0] == 13.5 &&
		matrix[3][2] == 15.5) {
		t.Error("4x4 Matrix init error")
	}
}

func Test2By2Construction(t *testing.T) {
	matrix := Matrix{
		{-3, 5},
		{1, -2},
	}

	if !(matrix[0][0] == -3 && matrix[0][1] == 5 && matrix[1][0] == 1 &&
		matrix[1][1] == -2) {
		t.Error("2x2 Matrix init error")
	}
}
