package main

func initializations() {
	// INIT START OMIT
	var i int
	var j int = 42
	k := 42
	var foo, bar int64 = 42, 1093
	// INIT END OMIT
}

// FUNC START OMIT
func foo1(param1 int, param2 string) {}

func foo2(param1, param2 int) {}

func foo3() string {
	return "success"
}

func foo4() (string, error) {
	return "success", nil
}

// FUNC END OMIT
