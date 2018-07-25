package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jontorrado/ine-callejero-dataset/processors"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter filename (relative path): ")
	filename, _ := reader.ReadString('\n')
	filename = strings.TrimSpace(filename)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if strings.Contains(filename, "PSEU") {
		processors.PseudoProcess(file)
	} else if strings.Contains(filename, "TRAMOS") {
		processors.TramosProcess(file)
	} else if strings.Contains(filename, "UP") {
		processors.UnidadProcess(file)
	} else if strings.Contains(filename, "VIAS") {
		processors.ViaProcess(file)
	} else {
		fmt.Println("File name not recognised")
		os.Exit(3)
	}
	os.Exit(3)
}
