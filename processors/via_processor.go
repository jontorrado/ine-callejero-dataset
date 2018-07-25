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

func ViaProcess(file io.Reader) {
	scanner := bufio.NewScanner(file)
	var elements []models.Via
	for scanner.Scan() {
		line := scanner.Text()
		elements = append(elements, models.Via{
			strings.TrimSpace(line[0:2]),
			strings.TrimSpace(line[2:5]),
			strings.TrimSpace(line[5:10]),
			strings.TrimSpace(line[10:11]),
			strings.TrimSpace(line[11:13]),
			strings.TrimSpace(line[13:21]),
			strings.TrimSpace(line[21:22]),
			strings.TrimSpace(line[22:27]),
			strings.TrimSpace(line[27:32]),
			strings.TrimSpace(line[32:33]),
			strings.TrimSpace(line[33:83]),
			strings.TrimSpace(line[83:108]),
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("TOTAL ELEMENTS: ", len(elements))
	viasJson, _ := json.Marshal(elements)
	err := ioutil.WriteFile("vias_012018.json", viasJson, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
