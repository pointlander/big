// Copyright 2020 The Big Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package big

import (
	"math/big"

	"github.com/ALTree/bigfloat"
)

// Float is an imaginary number
type Float struct {
	a, b *big.Float
}

// NewFloat creates a new imaginary number
func NewFloat(a, b *big.Float) *Float {
	return &Float{
		a: a,
		b: b,
	}
}

// Abs computes the absolute value of a
func (f *Float) Abs(a *Float) *Float {
	f.a.Mul(a.a, a.a)
	f.b.Mul(a.b, a.b)
	f.a = bigfloat.Sqrt(f.a.Add(f.a, f.b))
	f.b = big.NewFloat(0)
	return f
}

// Add add two imaginary numbers
func (f *Float) Add(a, b *Float) *Float {
	f.a.Add(a.a, b.a)
	f.b.Add(a.b, b.b)
	return f
}

// Sub subtracts two imaginary numbers
func (f *Float) Sub(a, b *Float) *Float {
	f.a.Sub(a.a, b.a)
	f.b.Sub(a.b, b.b)
	return f
}

// Mul multiples two imaginary numbers
func (f *Float) Mul(a, b *Float) *Float {
	x1, x2, x3, x4 :=
		big.NewFloat(0).SetPrec(f.a.Prec()), big.NewFloat(0).SetPrec(f.b.Prec()),
		big.NewFloat(0).SetPrec(f.b.Prec()), big.NewFloat(0).SetPrec(f.a.Prec())
	x1.Mul(a.a, b.a) // a*a
	x2.Mul(a.a, b.b) // a*ib
	x3.Mul(a.b, b.a) // ib*a
	x4.Mul(a.b, b.b) // i^2 * b = -b
	f.a.Add(x1, x4.Neg(x4))
	f.b.Add(x2, x3)
	return f
}

// Conj computes the complex conjugate of a
func (f *Float) Conj(a *Float) *Float {
	f.a.Set(a.a)
	f.b.Neg(a.b)
	return f
}

// Div divides two imaginary numbers
func (f *Float) Div(a, b *Float) *Float {
	c := NewFloat(big.NewFloat(0).SetPrec(f.a.Prec()), big.NewFloat(0).SetPrec(f.b.Prec()))
	c.Conj(b)
	x := NewFloat(a.a.Copy(a.a), a.b.Copy(a.b))
	y := NewFloat(b.a.Copy(b.a), b.b.Copy(b.b))
	x.Mul(x, c)
	y.Mul(y, c)
	f.a.Quo(x.a, y.a)
	f.b.Quo(x.b, y.a)
	return f
}

// String returns a string of the imaginary number
func (f *Float) String() string {
	return f.a.String() + " + " + f.b.String() + "i"
}

// Rational is an imaginary number
type Rational struct {
	a, b *big.Rat
}

// NewRational creates a new imaginary number
func NewRational(a, b *big.Rat) *Rational {
	return &Rational{
		a: a,
		b: b,
	}
}

// Add add two imaginary numbers
func (r *Rational) Add(a, b *Rational) *Rational {
	r.a.Add(a.a, b.a)
	r.b.Add(a.b, b.b)
	return r
}

// Sub subtracts two imaginary numbers
func (r *Rational) Sub(a, b *Rational) *Rational {
	r.a.Sub(a.a, b.a)
	r.b.Sub(a.b, b.b)
	return r
}

// Mul multiples two imaginary numbers
func (r *Rational) Mul(a, b *Rational) *Rational {
	x1, x2, x3, x4 :=
		big.NewRat(0, 1), big.NewRat(0, 1),
		big.NewRat(0, 1), big.NewRat(0, 1)
	x1.Mul(a.a, b.a) // a*a
	x2.Mul(a.a, b.b) // a*ib
	x3.Mul(a.b, b.a) // ib*a
	x4.Mul(a.b, b.b) // i^2 * b = -b
	r.a.Add(x1, x4.Neg(x4))
	r.b.Add(x2, x3)
	return r
}

// Conj computes the complex conjugate of a
func (r *Rational) Conj(a *Rational) *Rational {
	r.a.Set(a.a)
	r.b.Neg(a.b)
	return r
}

// Div divides two imaginary numbers
func (r *Rational) Div(a, b *Rational) *Rational {
	c := NewRational(big.NewRat(0, 1), big.NewRat(0, 1))
	c.Conj(b)
	x := NewRational(a.a.Set(a.a), a.b.Set(a.b))
	y := NewRational(b.a.Set(b.a), b.b.Set(b.b))
	x.Mul(x, c)
	y.Mul(y, c)
	r.a.Quo(x.a, y.a)
	r.b.Quo(x.b, y.a)
	return r
}

// String returns a string of the imaginary number
func (r *Rational) String() string {
	return r.a.String() + " + " + r.b.String() + "i"
}
