package tutorials

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/binding"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func bindingScreen(_ fyne.Window) fyne.CanvasObject {
	data := binding.NewFloat()
	label := widget.NewLabel(floatLabel(data))
	data.AddListener(binding.NewNotifyFunction(func(binding.DataItem) {
		label.SetText(floatLabel(data))
	}))

	slide := widget.NewSliderWithData(0, 1, data)
	slide.Step = 0.1
	bar := widget.NewProgressBarWithData(data)

	buttons := container.NewGridWithColumns(4,
		widget.NewButton("0%", func() {
			data.Set(0)
		}),
		widget.NewButton("30%", func() {
			data.Set(0.3)
		}),
		widget.NewButton("70%", func() {
			data.Set(0.7)
		}),
		widget.NewButton("100%", func() {
			data.Set(1)
		}))

	return container.NewVBox(label, slide, bar, buttons)
}

func floatLabel(data binding.Float) string { // TODO add type conversion
	return fmt.Sprintf("Float current value: %0.1f", data.Get())
}
