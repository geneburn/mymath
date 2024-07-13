package mymath

import "math"

// Sqrt вычисляет квадратный корень из числа x и возвращает результат
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

// Abs вычисляет абсолютное значение числа x и возвращает результат
func Abs(x float64) float64 {
	return math.Abs(x)
}

// Pow вычисляет степень числа x и возвращает результат
func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

// Log вычисляет логарифм числа x и возвращает результат
func Log(x float64) float64 {
	return math.Log(x)
}

// Exp вычисляет экспоненту числа x и возвращает результат
func Exp(x float64) float64 {
	return math.Exp(x)
}

// Sin вычисляет синус числа x и возвращает результат
func Sin(x float64) float64 {
	return math.Sin(x)
}

// Cos вычисляет косинус числа x и возвращает результат
func Cos(x float64) float64 {
	return math.Cos(x)
}

// Tan вычисляет тангенс числа x и возвращает результат
func Tan(x float64) float64 {
	return math.Tan(x)
}

// Asin вычисляет арксинус числа x и возвращает результат
func Asin(x float64) float64 {
	return math.Asin(x)
}

// Acos вычисляет арккосинус числа x и возвращает результат
func Acos(x float64) float64 {
	return math.Acos(x)
}

// Atan вычисляет арктангенс числа x и возвращает результат
func Atan(x float64) float64 {
	return math.Atan(x)
}

// Atan2 вычисляет арктангенс числа x и возвращает результат
func Atan2(y, x float64) float64 {
	return math.Atan2(y, x)
}

// Cbrt вычисляет кубический корень из числа x и возвращает результат
func Cbrt(x float64) float64 {
	return math.Cbrt(x)
}

// Ceil вычисляет ближайшее большее целое число из числа x и возвращает результат
func Ceil(x float64) float64 {
	return math.Ceil(x)
}

// Floor вычисляет ближайшее меньшее целое число из числа x и возвращает результат
func Floor(x float64) float64 {
	return math.Floor(x)
}

// Round вычисляет ближайшее целое число из числа x и возвращает результат
func Round(x float64) float64 {
	return math.Round(x)
}

// Trunc вычисляет целое число из числа x и возвращает результат
func Trunc(x float64) float64 {
	return math.Trunc(x)
}

// Mod вычисляет остаток от деления числа x на число y и возвращает результат
func Mod(x, y float64) float64 {
	return math.Mod(x, y)
}

// Hypot вычисляет гипотенузу треугольника с катетами x и y и возвращает результат
func Hypot(x, y float64) float64 {
	return math.Hypot(x, y)
}

// Pow10 вычисляет 10-й степень числа x и возвращает результат
func Pow10(x int) float64 {
	return math.Pow10(x)
}

// Min возвращает наименьшее число из x и y
func Min(x, y float64) float64 {
	return math.Min(x, y)
}

// Max возвращает наибольшее число из x и y
func Max(x, y float64) float64 {
	return math.Max(x, y)
}
