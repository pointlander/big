// Copyright 2020 The C0mpl3x Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package big

import (
	"math/big"
	"testing"
)

func TestMatrix_Add(t *testing.T) {
	a1 := NewRational(big.NewRat(1, 1), big.NewRat(0, 1))
	a2 := NewRational(big.NewRat(2, 1), big.NewRat(0, 1))
	a3 := NewRational(big.NewRat(3, 1), big.NewRat(0, 1))
	a4 := NewRational(big.NewRat(4, 1), big.NewRat(0, 1))
	m := Matrix{}
	m.Values = append(m.Values, []Rational{*a1, *a2})
	m.Values = append(m.Values, []Rational{*a3, *a4})
	m.Add(&m, &m)
	t.Log(m.String())
	if m.String() != "[2/1 + 0/1i 4/1 + 0/1i ;6/1 + 0/1i 8/1 + 0/1i ;]" {
		t.Fatal("invalid result")
	}

	m.Values = [][]Rational{}
	m.Values = append(m.Values, []Rational{*a1, *a2})
	m.Values = append(m.Values, []Rational{*a3, *a4})

	n := Matrix{}
	n.Values = make([][]Rational, 1)
	n.Values[0] = append(n.Values[0], *a2)
	m.Add(&m, &n)
	t.Log(m.String())
	if m.String() != "[3/1 + 0/1i 4/1 + 0/1i ;5/1 + 0/1i 6/1 + 0/1i ;]" {
		t.Fatal("invalid result")
	}

	m.Values = [][]Rational{}
	m.Values = append(m.Values, []Rational{*a1, *a2})
	m.Values = append(m.Values, []Rational{*a3, *a4})

	n.Values = make([][]Rational, 1)
	n.Values[0] = append(n.Values[0], *a2)
	m.Add(&n, &m)
	t.Log(m.String())
	if m.String() != "[3/1 + 0/1i 4/1 + 0/1i ;5/1 + 0/1i 6/1 + 0/1i ;]" {
		t.Fatal("invalid result")
	}
}

func TestMatrix_Sub(t *testing.T) {
	a1 := NewRational(big.NewRat(1, 1), big.NewRat(0, 1))
	a2 := NewRational(big.NewRat(2, 1), big.NewRat(0, 1))
	a3 := NewRational(big.NewRat(3, 1), big.NewRat(0, 1))
	a4 := NewRational(big.NewRat(4, 1), big.NewRat(0, 1))
	m := Matrix{}
	m.Values = append(m.Values, []Rational{*a1, *a2})
	m.Values = append(m.Values, []Rational{*a3, *a4})
	m.Sub(&m, &m)
	t.Log(m.String())
	if m.String() != "[0/1 + 0/1i 0/1 + 0/1i ;0/1 + 0/1i 0/1 + 0/1i ;]" {
		t.Fatal("invalid result")
	}

	m.Values = [][]Rational{}
	m.Values = append(m.Values, []Rational{*a1, *a2})
	m.Values = append(m.Values, []Rational{*a3, *a4})

	n := Matrix{}
	n.Values = make([][]Rational, 1)
	n.Values[0] = append(n.Values[0], *a2)
	m.Sub(&m, &n)
	t.Log(m.String())
	if m.String() != "[-1/1 + 0/1i 0/1 + 0/1i ;1/1 + 0/1i 2/1 + 0/1i ;]" {
		t.Fatal("invalid result")
	}

	m.Values = [][]Rational{}
	m.Values = append(m.Values, []Rational{*a1, *a2})
	m.Values = append(m.Values, []Rational{*a3, *a4})

	n.Values = make([][]Rational, 1)
	n.Values[0] = append(n.Values[0], *a2)
	m.Sub(&n, &m)
	t.Log(m.String())
	if m.String() != "[1/1 + 0/1i 0/1 + 0/1i ;-1/1 + 0/1i -2/1 + 0/1i ;]" {
		t.Fatal("invalid result")
	}
}

func TestMatrix_Mul(t *testing.T) {
	a1 := NewRational(big.NewRat(1, 1), big.NewRat(0, 1))
	a2 := NewRational(big.NewRat(2, 1), big.NewRat(0, 1))
	a3 := NewRational(big.NewRat(3, 1), big.NewRat(0, 1))
	a4 := NewRational(big.NewRat(4, 1), big.NewRat(0, 1))
	m := Matrix{}
	m.Values = append(m.Values, []Rational{*a1, *a2})
	m.Values = append(m.Values, []Rational{*a3, *a4})
	m.Mul(&m, &m)
	t.Log(m.String())
	if m.String() != "[7/1 + 0/1i 10/1 + 0/1i ;15/1 + 0/1i 22/1 + 0/1i ;]" {
		t.Fatal("invalid result")
	}

	m.Values = [][]Rational{}
	m.Values = append(m.Values, []Rational{*a1, *a2})
	m.Values = append(m.Values, []Rational{*a3, *a4})

	n := Matrix{}
	n.Values = make([][]Rational, 1)
	n.Values[0] = append(n.Values[0], *a2)
	m.Mul(&m, &n)
	t.Log(m.String())
	if m.String() != "[2/1 + 0/1i 4/1 + 0/1i ;6/1 + 0/1i 8/1 + 0/1i ;]" {
		t.Fatal("invalid result")
	}

	m.Values = [][]Rational{}
	m.Values = append(m.Values, []Rational{*a1, *a2})
	m.Values = append(m.Values, []Rational{*a3, *a4})

	n.Values = make([][]Rational, 1)
	n.Values[0] = append(n.Values[0], *a2)
	m.Mul(&n, &m)
	t.Log(m.String())
	if m.String() != "[2/1 + 0/1i 4/1 + 0/1i ;6/1 + 0/1i 8/1 + 0/1i ;]" {
		t.Fatal("invalid result")
	}
}

func TestMatrix_Div(t *testing.T) {
	a1 := NewRational(big.NewRat(1, 1), big.NewRat(0, 1))
	a2 := NewRational(big.NewRat(2, 1), big.NewRat(0, 1))
	m := Matrix{}
	m.Values = append(m.Values, []Rational{*a1})
	n := Matrix{}
	n.Values = append(n.Values, []Rational{*a2})
	m.Div(&m, &n)
	t.Log(m.String())
	if m.String() != "[1/2 + 0/1i ;]" {
		t.Fatal("invalid result")
	}
}

func TestFloat_Abs(t *testing.T) {
	a := NewFloat(big.NewFloat(5), big.NewFloat(5))
	a.Abs(a)
	t.Log(a.String())
	if a.String() != "7.071067812 + 0i" {
		t.Fatal("invalid result")
	}
}

func TestFloat_Div(t *testing.T) {
	a := NewFloat(big.NewFloat(4), big.NewFloat(5))
	b := NewFloat(big.NewFloat(2), big.NewFloat(6))
	a.Div(a, b)
	t.Log(a.String())
	if a.String() != "0.95 + -0.35i" {
		t.Fatal("invalid result")
	}
}

func TestFloat_Sqrt(t *testing.T) {
	a := NewFloat(big.NewFloat(5), big.NewFloat(12))
	a.Sqrt(a)
	t.Log(a.String())
	if a.String() != "3 + 2i" {
		t.Fatal("invalid result")
	}
}

func TestFloat_Exp(t *testing.T) {
	a := NewFloat(big.NewFloat(1), big.NewFloat(1))
	a.Exp(a)
	t.Log(a.String())
	if a.String() != "1.46869394 + 2.287355287i" {
		t.Fatal("invalid result")
	}

	a = NewFloat(big.NewFloat(0), big.NewFloat(1))
	a.Exp(a)
	t.Log(a.String())
	if a.String() != "0.5403023059 + 0.8414709848i" {
		t.Fatal("invalid result")
	}

	a = NewFloat(big.NewFloat(1), big.NewFloat(0))
	a.Exp(a)
	t.Log(a.String())
	if a.String() != "2.718281828 + 0i" {
		t.Fatal("invalid result")
	}

	a = NewFloat(big.NewFloat(1), big.NewFloat(2))
	a.Exp(a)
	t.Log(a.String())
	if a.String() != "-1.131204384 + 2.471726672i" {
		t.Fatal("invalid result")
	}

	a = NewFloat(big.NewFloat(2), big.NewFloat(1))
	a.Exp(a)
	t.Log(a.String())
	if a.String() != "3.992324048 + 6.217676312i" {
		t.Fatal("invalid result")
	}
}

func TestFloat_Cos(t *testing.T) {
	a := NewFloat(big.NewFloat(1), big.NewFloat(1))
	a.Cos(a)
	t.Log(a.String())
	if a.String() != "0.8337300251 + -0.9888977058i" {
		t.Fatal("invalid result")
	}

	a = NewFloat(big.NewFloat(0), big.NewFloat(1))
	a.Cos(a)
	t.Log(a.String())
	if a.String() != "1.543080635 + 0i" {
		t.Fatal("invalid result")
	}

	a = NewFloat(big.NewFloat(1), big.NewFloat(0))
	a.Cos(a)
	t.Log(a.String())
	if a.String() != "0.5403023059 + 0i" {
		t.Fatal("invalid result")
	}

	a = NewFloat(big.NewFloat(1), big.NewFloat(2))
	a.Cos(a)
	t.Log(a.String())
	if a.String() != "2.032723007 + -3.051897799i" {
		t.Fatal("invalid result")
	}

	a = NewFloat(big.NewFloat(2), big.NewFloat(1))
	a.Cos(a)
	t.Log(a.String())
	if a.String() != "-0.6421481247 + -1.068607421i" {
		t.Fatal("invalid result")
	}
}

func TestFloat_Sin(t *testing.T) {
	a := NewFloat(big.NewFloat(1), big.NewFloat(1))
	a.Sin(a)
	t.Log(a.String())
	if a.String() != "1.298457581 + 0.6349639148i" {
		t.Fatal("invalid result")
	}

	a = NewFloat(big.NewFloat(0), big.NewFloat(1))
	a.Sin(a)
	t.Log(a.String())
	if a.String() != "0 + 1.175201194i" {
		t.Fatal("invalid result")
	}

	a = NewFloat(big.NewFloat(1), big.NewFloat(0))
	a.Sin(a)
	t.Log(a.String())
	if a.String() != "0.8414709848 + 0i" {
		t.Fatal("invalid result")
	}

	a = NewFloat(big.NewFloat(1), big.NewFloat(2))
	a.Sin(a)
	t.Log(a.String())
	if a.String() != "3.165778513 + 1.959601041i" {
		t.Fatal("invalid result")
	}

	a = NewFloat(big.NewFloat(2), big.NewFloat(1))
	a.Sin(a)
	t.Log(a.String())
	if a.String() != "1.403119251 + -0.489056259i" {
		t.Fatal("invalid result")
	}
}

func TestFloat_Tan(t *testing.T) {
	a := NewFloat(big.NewFloat(1), big.NewFloat(1))
	a.Tan(a)
	t.Log(a.String())
	if a.String() != "0.2717525853 + 1.083923327i" {
		t.Fatal("invalid result")
	}

	a = NewFloat(big.NewFloat(0), big.NewFloat(1))
	a.Tan(a)
	t.Log(a.String())
	if a.String() != "0 + 0.761594156i" {
		t.Fatal("invalid result")
	}

	a = NewFloat(big.NewFloat(1), big.NewFloat(0))
	a.Tan(a)
	t.Log(a.String())
	if a.String() != "1.557407725 + 0i" {
		t.Fatal("invalid result")
	}

	a = NewFloat(big.NewFloat(1), big.NewFloat(2))
	a.Tan(a)
	t.Log(a.String())
	if a.String() != "0.03381282608 + 1.014793616i" {
		t.Fatal("invalid result")
	}

	a = NewFloat(big.NewFloat(2), big.NewFloat(1))
	a.Tan(a)
	t.Log(a.String())
	if a.String() != "-0.2434582012 + 1.166736257i" {
		t.Fatal("invalid result")
	}
}

func TestFloat_Log(t *testing.T) {
	a := NewFloat(big.NewFloat(1), big.NewFloat(1))
	a.Log(a)
	t.Log(a.String())
	if a.String() != "0.3465735903 + 0.7853981634i" {
		t.Fatal("invalid result")
	}

	a = NewFloat(big.NewFloat(1), big.NewFloat(0))
	a.Log(a)
	t.Log(a.String())
	if a.String() != "0 + 0i" {
		t.Fatal("invalid result")
	}

	a = NewFloat(big.NewFloat(0), big.NewFloat(1))
	a.Log(a)
	t.Log(a.String())
	if a.String() != "0 + 1.570796327i" {
		t.Fatal("invalid result")
	}

	a = NewFloat(big.NewFloat(1), big.NewFloat(2))
	a.Log(a)
	t.Log(a.String())
	if a.String() != "0.8047189562 + 1.107148718i" {
		t.Fatal("invalid result")
	}

	a = NewFloat(big.NewFloat(2), big.NewFloat(1))
	a.Log(a)
	t.Log(a.String())
	if a.String() != "0.8047189562 + 0.463647609i" {
		t.Fatal("invalid result")
	}
}

func TestFloat_Pow(t *testing.T) {
	a := NewFloat(big.NewFloat(2).SetPrec(64), big.NewFloat(1).SetPrec(64))
	b := NewFloat(big.NewFloat(2).SetPrec(64), big.NewFloat(1).SetPrec(64))
	c := NewFloat(big.NewFloat(0).SetPrec(64), big.NewFloat(0).SetPrec(64))
	c.Pow(a, b)
	t.Log(c.String())
	if c.String() != "-0.504824689 + 3.104144077i" {
		t.Fatal("invalid result")
	}

	a = NewFloat(big.NewFloat(1).SetPrec(64), big.NewFloat(2).SetPrec(64))
	b = NewFloat(big.NewFloat(1).SetPrec(64), big.NewFloat(2).SetPrec(64))
	c = NewFloat(big.NewFloat(0).SetPrec(64), big.NewFloat(0).SetPrec(64))
	c.Pow(a, b)
	t.Log(c.String())
	if c.String() != "-0.2225171568 + 0.1007091311i" {
		t.Fatal("invalid result")
	}

	a = NewFloat(big.NewFloat(0).SetPrec(64), big.NewFloat(1).SetPrec(64))
	b = NewFloat(big.NewFloat(0).SetPrec(64), big.NewFloat(1).SetPrec(64))
	c = NewFloat(big.NewFloat(0).SetPrec(64), big.NewFloat(0).SetPrec(64))
	c.Pow(a, b)
	t.Log(c.String())
	if c.String() != "0.2078795764 + 0i" {
		t.Fatal("invalid result")
	}

	a = NewFloat(big.NewFloat(1).SetPrec(64), big.NewFloat(0).SetPrec(64))
	b = NewFloat(big.NewFloat(1).SetPrec(64), big.NewFloat(0).SetPrec(64))
	c = NewFloat(big.NewFloat(0).SetPrec(64), big.NewFloat(0).SetPrec(64))
	c.Pow(a, b)
	t.Log(c.String())
	if c.String() != "1 + 0i" {
		t.Fatal("invalid result")
	}

	a = NewFloat(big.NewFloat(0).SetPrec(64), big.NewFloat(0).SetPrec(64))
	b = NewFloat(big.NewFloat(0).SetPrec(64), big.NewFloat(0).SetPrec(64))
	c = NewFloat(big.NewFloat(0).SetPrec(64), big.NewFloat(0).SetPrec(64))
	c.Pow(a, b)
	t.Log(c.String())
	if c.String() != "+Inf + 0i" {
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
