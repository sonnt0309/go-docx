package docx

import (
	"errors"
	"fmt"
	"testing"
)

func BenchmarkDocument_ReplaceAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		replaceMap := PlaceholderMap{
			"key":                         "REPLACE some more",
			"key-with-dash":               "REPLACE",
			"key-with-dashes":             "REPLACE",
			"key with space":              "REPLACE",
			"key_with_underscore":         "REPLACE",
			"multiline":                   "REPLACE",
			"key.with.dots":               "REPLACE",
			"mixed-key.separator_styles#": "REPLACE",
			"yet-another_placeholder":     "REPLACE",
			"foo":                         "bar",
		}

		doc, err := Open("./examples/simple/template.docx")
		if err != nil {
			panic(err)
		}

		result, err := doc.GetVariableInFile()
		if err != nil {
			panic(err)
		}
		if len(result) != 1 {
			panic(errors.New("result != 1"))
		}

		err = doc.ReplaceAll(replaceMap)
		if err != nil {
			panic(err)
		}

		err = doc.WriteToFile("/tmp/replaced.docx")
		if err != nil {
			panic(err)
		}
	}
}
