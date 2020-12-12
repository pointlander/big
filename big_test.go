// Copyright 2020 The Big Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package big

import (
	"math/big"
	"testing"
)

func TestAbs(t *testing.T) {
	a := NewFloat(big.NewFloat(5), big.NewFloat(5))
	a.Abs(a)
	t.Log(a.String())
	if a.String() != "7.071067812 + 0i" {
		t.Fatal("invalid result")
	}
}

func TestImaginary_Div(t *testing.T) {
	a := NewFloat(big.NewFloat(4), big.NewFloat(5))
	b := NewFloat(big.NewFloat(2), big.NewFloat(6))
	a.Div(a, b)
	t.Log(a.String())
	if a.String() != "0.95 + -0.35i" {
		t.Fatal("invalid result")
	}
}

func TestImaginary_Sqrt(t *testing.T) {
	a := NewFloat(big.NewFloat(5), big.NewFloat(12))
	a.Sqrt(a)
	t.Log(a.String())
	if a.String() != "3 + 2i" {
		t.Fatal("invalid result")
	}
}

func TestImaginary_Pow(t *testing.T) {
	a := NewFloat(big.NewFloat(2).SetPrec(64), big.NewFloat(1).SetPrec(64))
	b := NewFloat(big.NewFloat(2).SetPrec(64), big.NewFloat(1).SetPrec(64))
	c := NewFloat(big.NewFloat(0).SetPrec(64), big.NewFloat(0).SetPrec(64))
	c.Pow(a, b)
	t.Log(c.String())
	if c.String() != "-0.504824689 + 3.104144077i" {
		t.Fatal("invalid result")
	}
}

func TestRational_Div(t *testing.T) {
	a := NewRational(big.NewRat(4, 1), big.NewRat(5, 1))
	b := NewRational(big.NewRat(2, 1), big.NewRat(6, 1))
	a.Div(a, b)
	t.Log(a.String())
	if a.String() != "19/20 + -7/20i" {
		t.Fatal("invalid result")
	}
}
