package logic;

import (
	"fmt"
	"math"
	"math/cmplx"
);

type Input struct {
	N1      float64 `json:"n1"`
	N2      float64 `json:"n2"`
	D1      float64 `json:"d1"`
	D2      float64 `json:"d2"`
	Theta   int     `json:"th"`
	N       int     `json:"n"`
    TT11    [1000]int
}

func (inputVal *Input) Calculate() {

	PI := math.Pi

	k := 1.75 * math.Pow10(-6)
    j := 1.3 * math.Pow10(-6)
    p := (k - j) / 1000;

    var lambda[1000]float64

	for i := 0; i < 1000; i++ {
                
		lambda[i] = j;
        j = j + p;
		
		theta1 := float64(inputVal.N) * PI / 180
		theta2 := math.Asin(((inputVal.N2 * math.Sin(theta1)) / inputVal.N1));
                
		d31 := inputVal.D1 * (1 * math.Pow10(-6));
		d41 := inputVal.D2 * (1 * math.Pow10(-6));

        a1 := math.Cos(theta1) 
		a2 := math.Cos(theta2)

        d1 := d31 / a1 
		d2 := d41 / a2

		r12 := ((inputVal.N1 * a1) - (inputVal.N2* a2))/((inputVal.N1 * a1) + (inputVal.N2 *a2));
        r21 := -r12;

        t12 := math.Sqrt(math.Abs(1 - r12 * r12));
        t21 := math.Sqrt(math.Abs(1 - r21 * r21));

        A := [][]float64{{1, r12}, {r12, 1}}
        B := [][]float64{{1, r21}, {r21, 1}}
        
        M12 := ArrayTimesFloat(A, (1 / t12))
        M21 := ArrayTimesFloat(B, (1 / t21))

		fmt.Printf("%f, %f", M12, M21)

        k1 := (2 * PI * inputVal.N1 / lambda[i]);
        k2 := (2 * PI * inputVal.N2 / lambda[i]);

        t := -(1 / (t12 * t21));

        eye1 := complex(math.Cos(k1 * d1), math.Sin(k1 * d1))
        eye2 := complex(math.Cos(k1 * d1), -math.Sin(k1 * d1))

        eye11 := ComplexTimesFloat(eye1, r21)
        eye22 := ComplexTimesFloat(eye2, r21)

        eye3 := complex(math.Cos(k2 * d2), math.Sin(k2 * d2))
        eye4 := complex(math.Cos(k2 * d2), -math.Sin(k2 * d2))

        eye33 := ComplexTimesFloat(eye3, r12)
        eye44 := ComplexTimesFloat(eye4, r12)

        e1 := ComplexTimesFloat(((eye1 * eye3) + (eye22 * eye33)), t)
        e2 := ComplexTimesFloat(((eye1 * eye44) + (eye22 * eye4)), t)
        e3 := ComplexTimesFloat(((eye11 * eye3) + (eye2 * eye33)), t)
        e4 := ComplexTimesFloat(((eye11 * eye44) + (eye2 * eye4)), t)

        ML1 := [][]complex128{{e1,e2},{e3,e4}}

        e5 := ComplexTimesFloat(((eye3 * eye1) + (eye44 * eye11)), t)
		e6 := ComplexTimesFloat(((eye3 * eye22) + (eye44 * eye2)), t)

        e7 := ComplexTimesFloat(((eye33 * eye1) + (eye4 * eye11)), t)
        e8 := ComplexTimesFloat(((eye33 * eye22) + (eye4 * eye2)), t)

        ML2 := [][]complex128{{e5,e6},{e7,e8}}
		fmt.Println(ML2)

        for z := 0; z < inputVal.Theta; z++ {
            ML1 = ComplexTimesComplex(ML1, ML1)
		}

        tt1 := 1/ML1[0][0]
        T1 := ComplexPowerFloat(tt1, 2.0)

        inputVal.TT11[i] = int(math.Log(cmplx.Abs(T1)));
    }
}