package main
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"samples/WebIM/models"
	"os"
	"path/filepath"
)

func main() {
	db, err := gorm.Open("mysql", "newuser:password@/mb?charset=utf8&parseTime=True&loc=Local")
	fmt.Printf("%v\n", err)
	fmt.Printf("%v\n", db)
	fmt.Printf("%v\n", db.HasTable("events"))

	var events []models.Event
	db.Find(&events)
	fmt.Printf("%v\n", events)
	writeToFile(events)
}

func writeToFile(events []models.Event) {

	var StanfordLibPath = "/Users/adi/Downloads/stanford-postagger-2015-12-09/"
	mainJar  := filepath.Join(StanfordLibPath, "stanford-postagger.jar")
	libs := filepath.Join(StanfordLibPath, "lib/*")
	model := filepath.Join(StanfordLibPath, "models/english-bidirectional-distsim.tagger")
	inputFile := "/tmp/batch-output.txt"
	outputFile := "/tmp/batch-tagged-output.txt"
	cmd := fmt.Sprintf("java -mx300m -classpath %v:%v edu.stanford.nlp.tagger.maxent.MaxentTagger -model %v -textFile %v > %v", mainJar, libs, model, inputFile, outputFile)
	fmt.Printf("%v\n", cmd)

	fo, err := os.Create(inputFile)
	if err != nil {
			panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
			if err := fo.Close(); err != nil {
					panic(err)
			}
	}()

	for _, event := range events {
				fmt.Printf("%v\n", event)

		if _, err := fo.WriteString(event.User + "_USER " + event.Content + "\n"); err != nil {
				panic(err)
		}
	}

}
