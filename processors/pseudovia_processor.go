package processors

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"github.com/jontorrado/ine-callejero-dataset/models"
)

func PseudoProcess(file io.Reader) {
	scanner := bufio.NewScanner(file)
	var elements []models.Pseudovia
	for scanner.Scan() {
		line := scanner.Text()
		elements = append(elements, models.Pseudovia{
			strings.TrimSpace(line[0:2]),
			strings.TrimSpace(line[2:5]),
			strings.TrimSpace(line[5:10]),
			strings.TrimSpace(line[10:11]),
			strings.TrimSpace(line[11:13]),
			strings.TrimSpace(line[13:21]),
			strings.TrimSpace(line[21:22]),
			strings.TrimSpace(line[22:27]),
			strings.TrimSpace(line[27:50]),
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("TOTAL ELEMENTS: ", len(elements))
	pseudoviasJson, _ := json.Marshal(elements)
	err := ioutil.WriteFile("pseudovias_012018.json", pseudoviasJson, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
