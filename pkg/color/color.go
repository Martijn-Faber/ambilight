package color

type Rgb struct {
	R, G, B float32
}

func InterpolateColor(color1 Rgb, color2 Rgb, fraction float32) (interpolated Rgb) {
	interpolated.R = (color2.R-color1.R)*fraction + color1.R
	interpolated.G = (color2.G-color1.G)*fraction + color1.G
	interpolated.B = (color2.B-color1.B)*fraction + color1.B
	return
}

func CorrectColor(rgb Rgb, corrections Rgb) (corrected Rgb) {
	corrected = rgb
	corrected.R *= corrections.R
	corrected.G *= corrections.G
	corrected.B *= corrections.B
	return
}

func InterpolateColors(colors1 []Rgb, colors2 []Rgb, step int, steps int) (colors []Rgb) {
	fraction := float32(step / steps)
	for i, color1 := range colors1 {
		colors = append(colors, InterpolateColor(color1, colors2[i], fraction))
	}
	return
}
