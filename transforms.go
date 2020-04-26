package gopherplot

import (
	"errors"
	"fmt"
)

// Shape is the description of a matrix's dimensions
type Shape []int

// Equals determines if two shapes are the same
func (s Shape) Equals(s2 Shape) bool {
	if len(s) != len(s2) {
		return false
	}
	for i := range s {
		if s[i] != s2[i] {
			return false
		}
	}
	return true
}

// Valuable will return a value
type Valuable interface {
	Value() interface{}
}

// Number is a number structure
type Number struct {
	Val float64
}

// Value of the Number
func (n *Number) Value() interface{} {
	return n.Val
}

func (n *Number) String() string {
	return fmt.Sprintf("%f", n.Value().(float64))
}

//Getable will retrieve the Valuable at the provided coordinates
type Getable interface {
	Get(...int) (Valuable, error)
}

// Dimensioned provides access to the shape of the structure
type Dimensioned interface {
	Dimension() Shape
}

// Vector is a list of float64 values
type Vector struct {
	Data []*Number
	dims Shape
	Getable
	Valuable
	Dimensioned
}

// Max return the maximum int value from a []int
func Max(vals ...int) int {
	var max int
	for i := 0; i < len(vals); i++ {
		if vals[i] > vals[max] {
			max = i
		}
	}
	return vals[max]
}

/*
NewVector creates a new vector.
  for a horizontal vector, the shape is [n,1]
  for a vertical vector, the shape is [1,n]
  One of shape's values must be 1.
  Shape cannot have negative values
  If the len(values) < n, the remaining values are set to 0
*/
func NewVector(shape []int, values ...float64) (*Vector, error) {
	if (shape[0] != 1 && shape[1] != 1) || shape[0] < 1 || shape[1] < 1 || len(shape) != 2 {
		return nil, errors.New("Invalid Shape")
	}
	newData := []*Number{}
	for i := 0; i < Max(shape...); i++ {
		v := &Number{Val: 0}
		if i < len(values) {
			v = &Number{Val: values[i]}
		}
		newData = append(newData, v)
	}
	v := Vector{
		dims: shape,
		Data: newData,
	}
	return &v, nil
}

// Dimension describes the shape of the Vector
func (v *Vector) Dimension() Shape {
	return v.dims
}

//Get the float64 value at the given address
func (v *Vector) Get(addr ...int) (Valuable, error) {
	err := errors.New("Invalid Address")
	var val Valuable
	if len(addr) > len(v.dims) {
		return val, err
	}
	for i, a := range addr {
		if a < 0 {
			return val, err
		}
		if a > v.dims[i] {
			return val, err
		}
	}
	return val, nil
}

func (v *Vector) String() string {
	trailer := ""
	prefix := ""
	if v.dims[0] > 1 {
		trailer = " "
	} else {
		trailer = "\n"
		prefix = " "
	}
	s := "[" + trailer
	for _, v := range v.Data {
		s += fmt.Sprintf("%s%f%s", prefix, v.Value().(float64), trailer)
	}
	s += "]"

	return s
}

//Times multiplies a number to each value of the Vector
func (v *Vector) Times(n *Number) {
	for i, d := range v.Data {
		v.Data[i] = &Number{Val: d.Value().(float64) * n.Value().(float64)}
	}
}

//Plus adds n to each element of Vector
func (v *Vector) Plus(n *Number) {
	for i, d := range v.Data {
		v.Data[i] = &Number{Val: d.Value().(float64) + n.Value().(float64)}
	}
}

// DotProduct produces the inner product of two Vectors
func (v *Vector) DotProduct(v2 *Vector) (*Number, error) {
	if v.dims[0] != v2.dims[1] && v.dims[1] != v2.dims[0] {
		return nil, errors.New("Vectors lack compatible shapes")
	}
	if v.dims[1] != 1 {
		return nil, errors.New("Vector must be a row")
	}
	var p float64
	for i := range v.Data {
		p += v.Data[i].Value().(float64) * v2.Data[i].Value().(float64)
	}

	return &Number{p}, nil
}

/*
T Generates a transpose the vector using the same data.
*/
func (v *Vector) T() *Vector {
	return &Vector{
		Data: v.Data,
		dims: []int{v.dims[1], v.dims[0]},
	}
}

// Matrix is an n x m matrix
type Matrix struct {
	dims []int
	Data []*Number
	Getable
	Valuable
	Dimensioned
}

/*
NewMatrix creates a new n-dimensional matrix
All dimensions must be >= 1
Data populates dimensions in the order presented
Unfilled areas default to 0
*/
func NewMatrix(dims []int, data ...float64) (*Matrix, error) {
	if len(dims) < 2 {
		return nil, errors.New("Matrix must have 2 or more dimension")
	}

	dLen := 1
	for i, d := range dims {
		dLen = dLen * d
		if dims[i] < 1 {
			return nil, fmt.Errorf("Dimension %d of %d must be >=1", i, d)
		}
	}
	if len(data) > dLen {
		return nil, fmt.Errorf("Invalid dimensions - %d data would be lost", len(data)-dLen)
	}

	newData := []*Number{}

	for i := 0; i < dLen; i++ {
		n := &Number{0}
		if i < len(data) {
			n.Val = data[i]
		}
		newData = append(newData, n)
	}

	return &Matrix{
		Data: newData,
		dims: dims,
	}, nil
}

//Get allows Matrix to be Gettable
func (m *Matrix) Get(dims ...int) (Valuable, error) {
	if len(dims) != len(m.dims) {
		return nil, fmt.Errorf("Shape has wrong dimensions.  Matrix has %d dimensions", len(dims))
	}
	for i, d := range dims {
		if m.dims[i] < dims[i] {
			return nil, fmt.Errorf("Requested %d Dimension %d exceeds Matrix shape %v", i, d, m.dims)
		}
	}

	i := dims[0]
	for idx, addr := range dims[1:] {
		i += (addr * m.dims[idx])
	}

	return m.Data[i], nil
}

// Dimension gives the dimension of the matrix
func (m *Matrix) Dimension() Shape {
	return m.dims
}

// AddNumber adds a number to each element of the Matrix
// To subtract, multiply n by -1
func (m *Matrix) AddNumber(n *Number) {
	for _, c := range m.Data {
		c.Val += n.Val
	}
}

//AddMatrix is for adding same-dimensioned matrixes
func (m *Matrix) AddMatrix(m2 *Matrix) error {
	if !m.Dimension().Equals(m2.Dimension()) {
		return errors.New("Matrixes are not the same dimensions")
	}
	for i := range m.Data {
		m.Data[i].Val += m2.Data[i].Val
	}
	return nil
}

func (m *Matrix) String() string {
	var s string
	for j := 0; j < m.Dimension()[1]; j++ {
		prefix := "\n"
		for i := 0; i < m.Dimension()[0]; i++ {
			v, err := m.Get(i, j)
			if err != nil {
				fmt.Println(err)
				return ""
			}
			s += fmt.Sprintf("%s[%0.2f]", prefix, v.Value().(float64))
			prefix = " "
		}
	}
	return s
}
