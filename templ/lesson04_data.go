package templ

// Lesson4Items returns the sample todo list used by lesson 4. Keeping the data
// in Go avoids embedding a braced composite literal inside a {@ } statement.
func Lesson4Items() []string {
	return []string{"Learn Go", "Learn templ", "Build something"}
}
