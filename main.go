// declare my main package
package main

// import fmt function
import "fmt"

// declare the CONST variable for the boiling point of water in F
const boilingF = 212

// main function
func main() {

    tempFF := 32
    tempFC := 0
    tempF := boilingF                      // variable for temperature value in degrees F
    tempC := (tempF - 32) * 5 / 9         // Conversion from F to C
    // display the result on the screen
    fmt.Printf("The boiling point of water in °F = %v (%T), boiling point of water in °C = %v (%T).", tempF, tempF, tempC, tempC)
    fmt.Printf("The melting point of water in °F = %v (%T), and the melting point of water in °C = %v (%T). ", tempFF, tempFF, tempFC, tempFC)
    // We expect to see F = 212 and C = 100
}
