// Copyright 2020 The Big Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package big

import (
	"math/big"
	"testing"
)

func TestAbs(t *testing.T) {
	a := NewImaginary(big.NewFloat(5), big.NewFloat(5))
	a.Abs(a)
	t.Log(a.String())
	if a.String() != "7.071067812 + 0i" {
		t.Fatal("invalid result")
	}
}

func TestDiv(t *testing.T) {
	a := NewImaginary(big.NewFloat(4), big.NewFloat(5))
	b := NewImaginary(big.NewFloat(2), big.NewFloat(6))
	a.Div(a, b)
	t.Log(a.String())
	if a.String() != "0.95 + -0.35i" {
		t.Fatal("invalid result")
	}
}
