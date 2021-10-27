package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/flexzuu/linear-tracker/linear"
	"github.com/flexzuu/linear-tracker/pasteboard"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

const (
	SFSymbolsTicket  = "􀪄" // ticket.fill 🎟
	SFSymbolsSpinner = "􀖇" // hourglass ⌛️
)

func renderMenu(app cocoa.NSApplication, q linear.AssignedIssues, title string, menu cocoa.NSMenu, items map[string]cocoa.NSMenuItem, btn cocoa.NSStatusBarButton) {
	itemsToRemove := map[string]struct{}{}
	// we start off by marking all prev items as to delete
	for k := range items {
		itemsToRemove[k] = struct{}{}
	}
	for _, issue := range q.Viewer.AssignedIssues.Nodes {
		// ignore unneeded issues
		if issue.State.ID != linear.StateInProgressID {
			continue
		}
		// check if the item existed before
		existingMenu, ok := items[string(issue.ID)]
		if ok {
			// item should remain - unmark it from the delete map
			delete(itemsToRemove, string(issue.ID))
			// update title etc to make sure we stay up-to-date with potential updates
			existingMenu.SetTitle(string(issue.BranchName))
			// update tooltip
			existingMenu.SetToolTip(string(issue.Title))
			// update on click action
			b := string(issue.BranchName)
			cb, cbSel := core.Callback(func(o objc.Object) {
				err := pasteboard.Copy(strings.NewReader(b))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(issue.URL)
				// app.Send("openURL:", core.URL(string(issue.URL)))
			})
			existingMenu.SetAction(cbSel)
			existingMenu.SetTarget(cb)

		} else {
			// if its not found its a new menu entry
			item := cocoa.NSMenuItem_New()
			item.SetTitle(string(issue.BranchName))
			item.SetToolTip(string(issue.Title))
			b := string(issue.BranchName)
			cb, cbSel := core.Callback(func(o objc.Object) {
				err := pasteboard.Copy(strings.NewReader(b))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(issue.URL)
				// app.Send("openURL:", core.URL(string(issue.URL)))
			})
			item.SetAction(cbSel)
			item.SetTarget(cb)
			// add item to the list
			menu.AddItem(item)
			// save item in the map for future updates
			items[string(issue.ID)] = item
		}
	}
	// now only the no longer relevant items remain in our remove list
	for k := range itemsToRemove {
		// run over each and remove the item
		item, ok := items[k]
		if ok {
			// remove it from the menue
			menu.RemoveItem(item)
			// get the target and free memory
			t := item.Target()
			t.Release()
			// free item memory
			item.Release()
			// delete reference from our bookkeeping map
			delete(items, k)
		}
	}

	// update title based on how many issues are in our menu
	if len(items) == 0 {
		title = SFSymbolsTicket
	} else if len(items) == 1 {
		for k := range items {
			for _, issue := range q.Viewer.AssignedIssues.Nodes {
				if string(issue.ID) == k {
					title = fmt.Sprintf("%s - %s", SFSymbolsTicket, string(issue.Identifier))
				}
			}
		}

	} else {
		title = fmt.Sprintf("%s (%d)", SFSymbolsTicket, len(items))
	}
	// update button title
	btn.SetTitle(title)
}
