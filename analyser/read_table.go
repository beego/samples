package analyser
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"samples/WebIM/models"
	"os"
	"path/filepath"
	"os/exec"
)

func main() {
	db, err := gorm.Open("mysql", "root/:@/mb?charset=utf8&parseTime=True&loc=Local")
	fmt.Printf("%v\n", err)
	fmt.Printf("%v\n", db)
	fmt.Printf("%v\n", db.HasTable("events"))

	var events []models.Event
	db.Find(&events)
	fmt.Printf("%v\n", events)
	// TODO: Push to settings/config
	inputFile := "/tmp/batch-output.txt"
	writeToFile(events, inputFile)
	// analysePartsOfSpeech(inputFile)
	AnalyseDependencies(inputFile)
}

func AnalyseDependencies(inputFile string) string {
	// TODO: Push all constants to settings/config
	var StanfordLibPath = "/Users/adi/Downloads/stanford-parser-full-2015-12-09/"
	outputFile := "/tmp/batch-dep-output.txt"
	parserLibs := []string{ "stanford-parser.jar", "stanford-parser-3.6.0-models.jar", "slf4j-api.jar"}
	classPath := classPath(StanfordLibPath, parserLibs)
	cmd := fmt.Sprintf("java -mx200m -cp %v edu.stanford.nlp.parser.lexparser.LexicalizedParser -retainTMPSubcategories -outputFormat 'wordsAndTags,penn,typedDependencies' edu/stanford/nlp/models/lexparser/englishPCFG.ser.gz %v | tee %v", classPath, inputFile, outputFile)
	out := PrintAndExec(cmd)
	return string(out)
}

func PrintAndExec(cmd string) string {
	fmt.Printf("%v\n", cmd)
	out, err := exec.Command("sh","-c",cmd).Output()
	fmt.Printf("StdOut:%s\n", out)
	fmt.Printf("StdErr:%s\n", err)
	return string(out)
}

func classPath(root string, libs []string)(path string) {
	for _, lib := range libs {
		path += filepath.Join(root, lib) + ":"
	}
	return path
}

func analysePartsOfSpeech(inputFile string) {
	// TODO: Push all constants to settings/config
	var StanfordLibPath = "/Users/adi/Downloads/stanford-postagger-2015-12-09/"
	mainJar  := filepath.Join(StanfordLibPath, "stanford-postagger.jar")
	libs := filepath.Join(StanfordLibPath, "lib/*")
	model := filepath.Join(StanfordLibPath, "models/english-bidirectional-distsim.tagger")
	outputFile := "/tmp/batch-tagged-output.txt"
	cmd := fmt.Sprintf("java -mx300m -classpath %v:%v edu.stanford.nlp.tagger.maxent.MaxentTagger -model %v -textFile %v | tee %v", mainJar, libs, model, inputFile, outputFile)
	PrintAndExec(cmd)
}

func writeToFile(events []models.Event, inputFile string) {

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
