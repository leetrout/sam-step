// © 2019-2020 nextmv.io inc. All rights reserved.
// nextmv.io, inc. CONFIDENTIAL
//
// This file includes unpublished proprietary source code of nextmv.io, inc.
// The copyright notice above does not evidence any actual or intended
// publication of such source code. Disclosure of this source code or any
// related proprietary information is strictly prohibited without the express
// written permission of nextmv.io, inc.

package samstep

import (
	"encoding/json"

	"github.com/nextmv-io/hop/model"
)

type board struct {
	rows model.IntDomains
}

// Feasible is true if the board has a single queen in each row, no two rows
// have a queen in the same column and there are no diagonally-arranged queens
func (b board) Feasible() bool {
	// iterate over rows until penultimate (second-to-last) row
	for i, row1 := range b.rows[:len(b.rows)-1] {
		// iterate over rows from one after row1 to last row
		for j, row2 := range b.rows[i+1:] {
			// check rows for threats between queens
			if threat(row1, row2, j+1) {
				return false
			}
		}
	}

	return true
}

func threat(r1, r2 model.IntDomain, offset int) bool {
	// get each row's column value and whether it is a singleton:
	col1, ok1 := r1.Value()
	col2, ok2 := r2.Value()
	// if either row is not a singleton, the solution is infeasible. We then
	// skip the checks for a threat and return true:
	if !ok1 || !ok2 {
		return true
	}

	// if column values are equal, these queens are in the same column, posing
	// a threat
	if col1 == col2 {
		return true
	}

	// if the distance between columns is equal to that between rows,
	// the queens are positioned diagonally; there is a threat:
	return (col1-col2 == offset || col1-col2 == -offset)
}

// Next places a queen in each available column of the most constrained row.
func (b board) Next() []model.State {
	next := []model.State{}

	if row, ok := b.rows.Smallest(); ok {
		for it := b.rows[row].Iterator(); it.Next(); {
			next = append(next, b.place(row, it.Value()))
		}
	}
	return next
}

// MarshalJSON provides custom marshaling for the board.
func (b board) MarshalJSON() ([]byte, error) {
	if values, ok := b.rows.Values(); ok {
		return json.Marshal(values)
	}
	return json.Marshal(nil)
}

// Place a queen on the board. Remove columns that threaten it from other rows.
func (b board) place(row, col int) board {
	rows := b.rows.Assign(row, col)
	for i := range rows {
		if i != row {
			offset := row - i
			rows[i] = rows[i].Remove(col-offset, col, col+offset)
		}
	}
	return board{rows: rows}
}
