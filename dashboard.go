package main

import (
    "fmt"             // Import the fmt package for formatted I/O
    "log"             // Import the log package for logging errors and information
    "time"            // Import the time package for handling time-related operations
    "github.com/jroimartin/gocui" // Import the gocui package for creating the GUI
)

// layout function defines the layout of the GUI
func layout(g *gocui.Gui) error {
    maxX, maxY := g.Size() // Get the size of the terminal window
    if v, err := g.SetView("stock", maxX/4, maxY/4, 3*maxX/4, 3*maxY/4); err != nil {
        if err != gocui.ErrUnknownView {
            return err // Return error if it's not an unknown view error
        }
        fmt.Fprintln(v, "Stock data will appear here") // Placeholder text for stock data
    }
    return nil // Return nil if no errors
}

// updateView function updates the stock data in the view
func updateView(g *gocui.Gui) error {
    v, err := g.View("stock") // Get the view named "stock"
    if err != nil {
        return err // Return error if view not found
    }
    v.Clear() // Clear the view
    stockData, err := getStockPrice("AAPL") // Fetch stock data for Apple (AAPL)
    if err != nil {
        fmt.Fprintln(v, "Error fetching stock data:", err) // Print error message in the view
        return nil // Return nil since we handle the error by displaying it
    }
    fmt.Fprintln(v, stockData) // Print fetched stock data in the view
    return nil // Return nil if no errors
}

// main function is the entry point of the application
func main() {
    g, err := gocui.NewGui(gocui.OutputNormal) // Create a new GUI instance
    if err != nil {
        log.Panicln(err) // Log and panic if there's an error creating the GUI
    }
    defer g.Close() // Ensure the GUI is closed when main function exits

    g.SetManagerFunc(layout) // Set the layout function to manage the GUI

    // Set a keybinding for Ctrl+C to quit the application
    if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
        return gocui.ErrQuit // Return the quit error to exit the main loop
    }); err != nil {
        log.Panicln(err) // Log and panic if there's an error setting the keybinding
    }

    // Start a goroutine to periodically update the view
    go func() {
        for {
            time.Sleep(1 * time.Second) // Wait for 1 second
            g.Update(updateView)        // Update the view with new stock data
        }
    }()

    // Start the main loop of the GUI
