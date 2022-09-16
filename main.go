package main

import (
	"log"
	my_draws "mygioTutorial/my_draw"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/unit"
)

// // Define  a progressIncrementer channel
// var progressIncrementer chan float32

func main() {

	// Setup a separate channel to provide ticks to increment progress
	my_draws.ProgressIncrementer = make(chan float32)

	go func() {
		for {
			time.Sleep(time.Second / 25)
			my_draws.ProgressIncrementer <- 0.004
		}

	}()

	go func() {
		//creat a new window
		w := app.NewWindow(
			app.Title("Egg Timer"),
			app.Size(unit.Dp(400), unit.Dp(600)),
		)

		if err := my_draws.Draw(w); err != nil {

			log.Fatal(err)
		}

		os.Exit(0)
	}()

	app.Main()
}

// func draw(w *app.Window) error {
// 	// is the egg boiling?
// 	var boiling bool
// 	var boilDuration float32
// 	// boilDurationInput is a textfield to input boil duration
// 	var boilDurationInput widget.Editor
// 	//operatioss from the UI
// 	var ops op.Ops
// 	//startButton is a clickable widget
// 	var startButton widget.Clickable
// 	//th defines the material design style
// 	th := material.NewTheme(gofont.Collection())

// 	for {
// 		select {
// 		// listen for events in the window.
// 		case e := <-w.Events():
// 			//detect what type of event
// 			switch e := e.(type) {

// 			//this is sent when the application should re-render.
// 			case system.FrameEvent:
// 				gtx := layout.NewContext(&ops, e)

// 				if startButton.Clicked() {
// 					boiling = !boiling

// 					// Read from the input box
// 					inputString := boilDurationInput.Text()
// 					inputString = strings.TrimSpace(inputString)
// 					inputFloat, _ := strconv.ParseFloat(inputString, 32)
// 					boilDuration = float32(inputFloat)
// 					boilDuration = boilDuration / (1 - progress)

// 				}
// 				//flexbox layout concept
// 				layout.Flex{
// 					//vertical alignment from top to buttom
// 					Axis: layout.Vertical,

// 					//Empty space is left at the start i.e at the top
// 					Spacing: layout.SpaceStart,
// 				}.Layout(gtx,
// 					layout.Rigid(
// 						func(gtx C) D {
// 							// Draw a custom path, shaped like an egg
// 							var eggPath clip.Path
// 							op.Offset(image.Pt(gtx.Dp(200), gtx.Dp(150))).Add(gtx.Ops)
// 							eggPath.Begin(gtx.Ops)
// 							// Rotate from 0 to 360 degrees
// 							for deg := 0.0; deg <= 360; deg++ {

// 								// Egg math (really) at this brilliant site. Thanks!
// 								// https://observablehq.com/@toja/egg-curve
// 								// Convert degrees to radians
// 								rad := deg / 360 * 2 * math.Pi
// 								// Trig gives the distance in X and Y direction
// 								cosT := math.Cos(rad)
// 								sinT := math.Sin(rad)
// 								// Constants to define the eggshape
// 								a := 110.0
// 								b := 150.0
// 								d := 20.0
// 								// The x/y coordinates
// 								x := a * cosT
// 								y := -(math.Sqrt(b*b-d*d*cosT*cosT) + d*sinT) * sinT
// 								// Finally the point on the outline
// 								p := f32.Pt(float32(x), float32(y))
// 								// Draw the line to this point
// 								eggPath.LineTo(p)
// 							}
// 							// Close the path
// 							eggPath.Close()

// 							// Get hold of the actual clip
// 							eggArea := clip.Outline{Path: eggPath.End()}.Op()

// 							// Fill the shape
// 							//color := color.NRGBA{R: 255, G: 239, B: 174, A: 255}
// 							color := color.NRGBA{R: 255, G: uint8(239 * (1 - progress)), B: uint8(174 * (1 - progress)), A: 255}
// 							paint.FillShape(gtx.Ops, color, eggArea)

// 							d := image.Point{Y: 375}
// 							return layout.Dimensions{Size: d}
// 						},
// 					),

// 					layout.Rigid(
// 						func(gtx C) D {
// 							//Wrap the Editor in material design
// 							ed := material.Editor(th, &boilDurationInput, "sec")

// 							//define characteristics of the input box
// 							boilDurationInput.SingleLine = true
// 							boilDurationInput.Alignment = text.Middle

// 							if boiling && progress < 1 {
// 								boilRemain := (1 - progress) * boilDuration
// 								// Format to 1 decimal.
// 								inputStr := fmt.Sprintf("%.1f", math.Round(float64(boilRemain)*10)/10)
// 								// Update the text in the inputbox
// 								boilDurationInput.SetText(inputStr)
// 							}

// 							// Define insets ...
// 							margins := layout.Inset{
// 								Top:    unit.Dp(0),
// 								Right:  unit.Dp(170),
// 								Bottom: unit.Dp(40),
// 								Left:   unit.Dp(170),
// 							}
// 							// ... and borders ...
// 							border := widget.Border{
// 								Color:        color.NRGBA{R: 204, G: 204, B: 204, A: 255},
// 								CornerRadius: unit.Dp(3),
// 								Width:        unit.Dp(2),
// 							}
// 							// ... before laying it out, one inside the other
// 							return margins.Layout(gtx,
// 								func(gtx C) D {
// 									return border.Layout(gtx, ed.Layout)
// 								},
// 							)
// 						},
// 					),
// 					// 	},
// 					// ),

// 					layout.Rigid(
// 						//progress bar widget
// 						func(gtx C) D {
// 							bar := material.ProgressBar(th, progress)
// 							return bar.Layout(gtx)
// 						},
// 					),
// 					//we insert two rigid elements
// 					//first a button
// 					layout.Rigid(

// 						func(gtx C) D {
// 							// ONE: First define margins around the button using layout.Inset ...
// 							margins := layout.Inset{
// 								Top:    unit.Dp(25),
// 								Bottom: unit.Dp(25),
// 								Right:  unit.Dp(40),
// 								Left:   unit.Dp(40),
// 							}
// 							// TWO: ... then we lay out those margins ...
// 							return margins.Layout(gtx,
// 								// THREE: ... and finally within the margins, we define and lay out the button
// 								func(gtx C) D {

// 									var text string
// 									if !boiling {
// 										text = "Start"
// 									}
// 									if boiling && progress < 1 {
// 										text = "Stop"
// 									}

// 									if boiling && progress >= 1 {
// 										text = "Finished"
// 									}

// 									btn := material.Button(th, &startButton, text)

// 									return btn.Layout(gtx)
// 								},
// 							)
// 						},
// 					),
// 					//then an empty spacer
// 					layout.Rigid(
// 						//the height of the spacer is 25 Device independent pixels
// 						layout.Spacer{Height: unit.Dp(25)}.Layout,
// 					),
// 				)

// 				e.Frame(gtx.Ops)
// 				// this is sent when the application is closed.
// 			case system.DestroyEvent:
// 				return e.Err
// 			}

// 		// listen for events in the incrementor channel
// 		case p := <-progressIncrementer:
// 			if boiling && progress < 1 {
// 				progress += p
// 				w.Invalidate()
// 			}
// 		}

// 	}

// 	//return nil
// }
