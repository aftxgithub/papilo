package papilo

import "log"

// SumComponent continuously pushes the sum of all numbers passing through it.
// The input and output data type is float64
func SumComponent(p *Pipe) {
	var sum float64
	for !p.IsClosed {
		d, err := p.Next()
		if err != nil {
			continue
		}
		num, err := ToFloat64(d)
		if err != nil {
			log.Println(err)
			continue
		}
		sum += num
		p.Write(sum)
	}
}
