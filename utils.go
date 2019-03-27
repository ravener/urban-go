package urban

import (
  "html"
  "regexp"
  "fmt"
)

// The regexp used to match [] boxes
var Regexp = regexp.MustCompile("\\[([^\\][]+)\\]")
// Regexp to replace whitespaces with +
var WS = regexp.MustCompile("\\s+")

// Markdown renders a given string to Markdown style.
// [some term] => [some term](https://www.urbandictionary.com/define.php?term=some+term)
func Markdown(str string) string {
  return Regexp.ReplaceAllStringFunc(str, func(m string) string {
    sub := Regexp.FindStringSubmatch(m)
    return fmt.Sprintf("%s(https://www.urbandictionary.com/define.php?term=%s)",
      m,
      WS.ReplaceAllString(sub[1], "+"),
    )
  })
}

// HTML renders a given string to HTML style.
// [some term] => <a href="https://www.urbandictionary.com/define.php?term=some+term">some term</a>
// str is also HTML-escaped
func HTML(str string) string {
  return Regexp.ReplaceAllStringFunc(str, func(m string) string {
    sub := Regexp.FindStringSubmatch(m)
    return fmt.Sprintf("<a href=\"https://www.urbandictionary.com/define.php?term=%s\">%s</a>",
      WS.ReplaceAllString(sub[1], "+"),
      html.EscapeString(m[1:len(m) - 1]),
    )
  })
}
