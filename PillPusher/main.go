/*
DESCRIPTION:
The target dose for the oral medication Katamol has been calculated: d (in mg).
The only options to prescribe the medication are pills of strength
a (in mg) or b (in mg), so it may not be possible to prescribe the exact target dose.
For a given target dose d, return the highest possible dose that can be made
with any amount of pills a and/or b that does not exceed the target dose.
Input: 3 integers, d,a,b representing the calculated target dose, the dose of pill a,
and the dose of pill b (all in mg)
Output: Integer value of the closest dose to the target that can be made
with any amount of pills a and/or b without going above the target dose.

Example:

Constraints:

a <= d
2 <=d <=10000
1 <= a < b <5000
*/

package main

import "fmt"

func Prescribe(d, a, b int) int {

	// d - граница передоза
	// a - к-во витаминов в смолтабле
	// b - к-во витаминов в бигтабле
	// срезов и переменных можно было бы и поменьше наляпатб, но так нагляднее в дебаге будет

	dose := 0 // сюда потом результат, ну и для промежуточных вычислений пойдёт

	if a > d && b > d { // на всякий случай если оба табла больше дозы - досвидоний
		return 0
	}

	maxbp := 0 // для начала считаю, сколько больших таблов влезает в дозу без передоза

	if b <= d {
		for dose < d {
			dose += b
			maxbp++
		}
		if dose > d {
			dose -= b
			maxbp--
		}
	}

	var bigdoses []int            // считаю дозы для всех возможных количеств бигтаблов в зиплоке
	for i := 0; i <= maxbp; i++ { // i - это к-во бигтаблов в текущей итерации
		if d >= b*i {
			bigdoses = append(bigdoses, b*i) // если ещё не передоз бигтаблами
		}
	}

	var smalldoses []int // для каждого к-ва бигтаблов надо посчитать добивку смолтаблами
	var sumdoses []int   // суммарные дозы бигтаблов и смолтаблов

	for i := 0; i < len(bigdoses); i++ {
		max_dobivka := d - bigdoses[i]
		dose = 0
		if a <= max_dobivka {
			for dose < max_dobivka {
				dose += a
			}
			if dose > max_dobivka {
				dose -= a
			}
			smalldoses = append(smalldoses, dose)
		}
		sumdoses = append(sumdoses, dose+bigdoses[i]) // в этом срезе все возможные дозы, из них надо взять максималбную
	}

	dose = MaxPossibleDose(sumdoses)
	return dose
}

func MaxPossibleDose(values []int) int { // в 1.21.3 есть либа slices с готовым поиском максимума,
	max := values[0]                // но на момент написания этого высера тут 1.20
	for _, number := range values { // и её ещё не прикрутили
		if number > max { // поэтому ищем максимум по старинке
			max = number
		}
	}
	return max
}

func main() {

	fmt.Println(Prescribe(1000, 15, 63))
}
