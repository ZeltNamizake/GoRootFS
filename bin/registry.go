package bin

var Commands = make(map[string]func([]string))

func Register(name string, fn func([]string)) {
	Commands[name] = fn
}