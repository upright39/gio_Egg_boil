module mygioTutorial

go 1.18

require gioui.org v0.0.0-20220808141521-a55065af9c1e

require (
	gioui.org/cpu v0.0.0-20210817075930-8d6a761490d2 // indirect
	gioui.org/shader v1.0.6 // indirect
	github.com/benoitkugler/textlayout v0.1.1 // indirect
	github.com/gioui/uax v0.2.1-0.20220325163150-e3d987515a12 // indirect
	github.com/go-text/typesetting v0.0.0-20220411150340-35994bc27a7b // indirect
	golang.org/x/exp v0.0.0-20210722180016-6781d3edade3 // indirect
	golang.org/x/image v0.0.0-20210628002857-a66eb6448b8d // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	golang.org/x/text v0.3.7 // indirect
)

/////////////////////////////////
//package main

// import (
// 	"log"
// 	"os"
// 	"time"

// 	"gioui.org/app"
// 	"gioui.org/font/gofont"
// 	"gioui.org/io/system"
// 	"gioui.org/layout"
// 	"gioui.org/op"
// 	"gioui.org/unit"
// 	"gioui.org/widget"
// 	"gioui.org/widget/material"
// )

// type C = layout.Context
// type D = layout.Dimensions

// var progress float32

// // is the egg boiling?
// var boiling bool

// // Define  a progressIncrementer channel
// var progressIncrementer chan float32

// func main() {

// 	// Setup a separate channel to provide ticks to increment progress
// 	progressIncrementer = make(chan float32)

// 	go func() {
// 		for {
// 			time.Sleep(time.Second / 25)
// 			progressIncrementer <- 0.004
// 		}

// 	}()

// 	go func() {
// 		//creat a new window
// 		w := app.NewWindow(
// 			app.Title("Egg Timer"),
// 			app.Size(unit.Dp(400), unit.Dp(600)),
// 		)

// 		if err := draw(w); err != nil {

// 			log.Fatal(err)
// 		}

// 		os.Exit(0)
// 	}()

// 	app.Main()
// }

// func draw(w *app.Window) error {
// 	//operatioss from the UI
// 	var ops op.Ops
// 	//startButton is a clickable widget
// 	var startButton widget.Clickable
// 	//th defines the material design style
// 	th := material.NewTheme(gofont.Collection())

// 	//listen for events in the window
// 	for e := range w.Events() {

// 		//detect what type of event
// 		switch e := e.(type) {

// 		//this is sent when the application should re-render.
// 		case system.FrameEvent:
// 			gtx := layout.NewContext(&ops, e)

// 			if startButton.Clicked() {
// 				boiling = !boiling
// 			}

// 			//flexbox layout concept
// 			layout.Flex{
// 				//vertical alignment from top to buttom
// 				Axis: layout.Vertical,

// 				//Empty space is left at the start i.e at the top
// 				Spacing: layout.SpaceStart,
// 			}.Layout(gtx,
// 				// progress bar Rigid
// 				layout.Rigid(
// 					//progress bar widget
// 					func(gtx C) D {
// 						bar := material.ProgressBar(th, progress)
// 						return bar.Layout(gtx)
// 					},
// 				),
// 				//we insert two rigid elements
// 				//first a button
// 				layout.Rigid(

// 					func(gtx C) D {
// 						// ONE: First define margins around the button using layout.Inset ...
// 						margins := layout.Inset{
// 							Top:    unit.Dp(25),
// 							Bottom: unit.Dp(25),
// 							Right:  unit.Dp(40),
// 							Left:   unit.Dp(40),
// 						}
// 						// TWO: ... then we lay out those margins ...
// 						return margins.Layout(gtx,
// 							// THREE: ... and finally within the margins, we define and lay out the button
// 							func(gtx C) D {

// 								var text string
// 								if !boiling {
// 									text = "Start"
// 								} else {
// 									text = "Stop"
// 								}
// 								btn := material.Button(th, &startButton, text)

// 								return btn.Layout(gtx)
// 							},
// 						)
// 					},
// 				),
// 				//then an empty spacer
// 				layout.Rigid(
// 					//the height of the spacer is 25 Device independent pixels
// 					layout.Spacer{Height: unit.Dp(25)}.Layout,
// 				),
// 			)

// 			e.Frame(gtx.Ops)
// 			// this is sent when the application is closed.
// 		case system.DestroyEvent:
// 			return e.Err
// 		}

// 	}
// 	return nil
// }

//image viewer rigid
// layout.Rigid(
// 	func(gtx C) D {
// 		circle := clip.Ellipse{
// 			// Hard coding the x coordinate. Try resizing the window
// 			// Min: image.Pt(80, 0),
// 			// Max: image.Pt(320, 240),
// 			// Soft coding the x coordinate. Try resizing the window
// 			Min: image.Pt(gtx.Constraints.Max.X/2-120, 60),
// 			Max: image.Pt(gtx.Constraints.Max.X/2+100, 300),
// 		}.Op(gtx.Ops)

// 		color := color.NRGBA{R: 200, A: 255}

// 		paint.FillShape(gtx.Ops, color, circle)
// 		d := image.Point{Y: 350}
// 		return layout.Dimensions{Size: d}
// 	},
// ),
