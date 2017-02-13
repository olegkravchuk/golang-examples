package hello_world

const testVersion = 3

func HelloWorld(name string) string {

	if name != ""{
		return "Hello, " + name + "!"
	} else {
		return "Hello, World!"
	}

}
