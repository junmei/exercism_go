package greeting

const testVersion = 3

// HelloWorld ...
// Greeting name
func HelloWorld(name string) string {
	if name == "" {
		return "Hello, World!"
	}
	return "Hello, " + name + "!"

}
