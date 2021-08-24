package color

type Rgb struct {
	R, G, B float32
}

func InterpolateColor(color1 Rgb, color2 Rgb, fraction float32) Rgb {
	var result Rgb
	result.R = (color2.R-color1.R)*fraction + color1.R
	result.G = (color2.G-color1.G)*fraction + color1.G
	result.B = (color2.B-color1.B)*fraction + color1.B
	return result
}

func InterpolateColors(colors1 []Rgb, colors2 []Rgb, step int, steps int) []Rgb {
	fraction := float32(step / steps)
	var colors []Rgb

	for i, color1 := range colors1 {
		colors = append(colors, InterpolateColor(color1, colors2[i], fraction))
	}

	return colors
}
