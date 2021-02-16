package main

import (
	"context"
	"flag"
	"os"
	"runtime"
	"time"

	"github.com/flexzuu/linear-tracker/linear"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
	"github.com/rs/zerolog/log"
)

func main() {

	authToken := flag.String("token", os.Getenv("LINEAR_TOKEN"), "linear auth token")
	flag.Parse()

	client := linear.NewClient(*authToken)

	runtime.LockOSThread()

	app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
		obj := cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength)
		obj.Retain()
		title := "Issue Tracker"
		btn := obj.Button()
		btn.SetTitle(title)

		items := map[string]cocoa.NSMenuItem{}

		menu := cocoa.NSMenu_New()

		render := func() {
			var iss linear.AssignedIssues
			err := client.Query(context.Background(), &iss, nil)
			if err != nil {
				log.Error().AnErr("error", err).Msg("linear query failed")
				return
			}
			core.Dispatch(func() {
				renderMenu(iss, title, menu, items, btn)
			})
		}

		cocoa.DefaultDelegateClass.AddMethod("menuWillOpen:", func(_ objc.Object) {
			go render()
		})

		ticker := time.NewTicker(30 * time.Second)
		done := make(chan bool)
		go func() {
			for {
				select {
				case <-done:
					return
				case <-ticker.C:
					go render()
				}
			}
		}()

		menu.Send("setDelegate:", cocoa.DefaultDelegate)
		go render()
		obj.SetMenu(menu)
	})
	app.Run()
}
