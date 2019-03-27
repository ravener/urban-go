package urban

import (
  "testing"
)

func TestMarkdown(t *testing.T) {
  if res := Markdown("[some term]"); res != "[some term](https://www.urbandictionary.com/define.php?term=some+term)" {
    t.Errorf("Markdown([some term]) gave unexpected results: %s", res)
  }
}

func TestHTML(t *testing.T) {
  if res := HTML("[some term]"); res != "<a href=\"https://www.urbandictionary.com/define.php?term=some+term\">some term</a>" {
    t.Errorf("HTML([some term]) gave unexpected results: %s", res)
  }
}
