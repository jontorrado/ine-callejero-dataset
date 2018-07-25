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

func TramosProcess(file io.Reader) {
	scanner := bufio.NewScanner(file)
	var elements []models.Tramo
	for scanner.Scan() {
		line := scanner.Text()
		elements = append(elements, models.Tramo{
			strings.TrimSpace(line[0:2]),
			strings.TrimSpace(line[2:5]),
			strings.TrimSpace(line[5:7]),
			strings.TrimSpace(line[7:10]),
			strings.TrimSpace(line[10:11]),
			strings.TrimSpace(line[11:13]),
			strings.TrimSpace(line[13:20]),
			strings.TrimSpace(line[20:25]),
			strings.TrimSpace(line[25:30]),
			strings.TrimSpace(line[30:42]),
			strings.TrimSpace(line[42:47]),
			strings.TrimSpace(line[47:48]),
			strings.TrimSpace(line[48:52]),
			strings.TrimSpace(line[52:53]),
			strings.TrimSpace(line[53:57]),
			strings.TrimSpace(line[57:58]),
			strings.TrimSpace(line[58:59]),
			strings.TrimSpace(line[59:61]),
			strings.TrimSpace(line[61:69]),
			strings.TrimSpace(line[69:70]),
			strings.TrimSpace(line[70:72]),
			strings.TrimSpace(line[72:75]),
			strings.TrimSpace(line[75:76]),
			strings.TrimSpace(line[76:78]),
			strings.TrimSpace(line[78:85]),
			strings.TrimSpace(line[85:110]),
			strings.TrimSpace(line[110:135]),
			strings.TrimSpace(line[135:160]),
			strings.TrimSpace(line[160:165]),
			strings.TrimSpace(line[165:190]),
			strings.TrimSpace(line[190:195]),
			strings.TrimSpace(line[195:245]),
			strings.TrimSpace(line[245:257]),
			strings.TrimSpace(line[257:262]),
			strings.TrimSpace(line[262:263]),
			strings.TrimSpace(line[263:267]),
			strings.TrimSpace(line[267:268]),
			strings.TrimSpace(line[268:272]),
			strings.TrimSpace(line[272:273]),
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("TOTAL ELEMENTS: ", len(elements))
	tramosJson, _ := json.Marshal(elements)
	err := ioutil.WriteFile("tramos_012018.json", tramosJson, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
