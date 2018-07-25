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

func UnidadProcess(file io.Reader) {
	scanner := bufio.NewScanner(file)
	var elements []models.Unidad
	for scanner.Scan() {
		line := scanner.Text()
		elements = append(elements, models.Unidad{
			strings.TrimSpace(line[0:2]),
			strings.TrimSpace(line[2:5]),
			strings.TrimSpace(line[5:12]),
			strings.TrimSpace(line[12:13]),
			strings.TrimSpace(line[13:15]),
			strings.TrimSpace(line[15:23]),
			strings.TrimSpace(line[23:24]),
			strings.TrimSpace(line[24:94]),
			strings.TrimSpace(line[94:144]),
			strings.TrimSpace(line[144:169]),
			strings.TrimSpace(line[169:239]),
			strings.TrimSpace(line[239:289]),
			strings.TrimSpace(line[289:324]),
			strings.TrimSpace(line[324:384]),
			strings.TrimSpace(line[384:434]),
			strings.TrimSpace(line[434:459]),
			strings.TrimSpace(line[459:529]),
			strings.TrimSpace(line[529:579]),
			strings.TrimSpace(line[579:604]),
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("TOTAL ELEMENTS: ", len(elements))
	unidadesJson, _ := json.Marshal(elements)
	err := ioutil.WriteFile("unidades_poblacionales_012018.json", unidadesJson, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
