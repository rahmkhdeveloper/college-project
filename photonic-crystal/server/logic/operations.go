package logic
 
import "math"

func ArrayTimesFloat(arr [][]float64, t float64) ([][]float64) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			arr[i][j] = arr[i][j] * t
		}
	}
	return arr
}

func ComplexTimesFloat(eye complex128, r float64) (complex128) {
	real := real(eye)*r
	imaginary := imag(eye)*r	

	return complex(real, imaginary)
}

func ComplexPowerFloat(c complex128, m float64) (complex128) {
	real := math.Pow(real(c), m)
	imaginary := math.Pow(imag(c), m)
	
	return complex(real, imaginary)
}

func ComplexTimesComplex(arr1 [][]complex128, arr2 [][]complex128) ([][]complex128) {

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[0]); j++ {
			arr1[i][j] = arr1[i][j] * arr2[i][j]
		}
	}

	return arr1
}