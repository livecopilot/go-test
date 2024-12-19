package main

func main() {
	label := Label{Widget{10, 10}, "State:"}
	label.X = 11
	label.Y = 12

	println(label.X, label.Y, label.Text)

	button1 := Button{Label{Widget{20, 20}, "OK"}}
	button2 := NewButton(100, 100, "Cancel")

	listBox := ListBox{Widget{30, 30}, []string{"AL", "AK", "AZ", "AR"}, 0}

	for _, painter := range []Painter{label, button1, button2, listBox} {
		painter.Paint()
	}

	for _, widget := range []interface{}{label, button1, button2, listBox} {
		widget.(Painter).Paint()
		if Clicker, ok := widget.(Clicker); ok {
			Clicker.Click()
		}

		println()
	}
}