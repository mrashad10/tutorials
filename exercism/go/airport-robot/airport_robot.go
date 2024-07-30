package airportrobot

type Greeter interface {
	LanguageName() string
	Greet() string
}

func SayHello(name string, greeter Greeter) string {
	return greeter.Greet() + " " + name + "!"
}

type Italian struct{}
type Portuguese struct{}

func (language Italian) LanguageName() string {
	return "Italian"
}

func (language Italian) Greet() string {
	return "I can speak " + language.LanguageName() + ": Ciao"
}

func (language Portuguese) LanguageName() string {
	return "Portuguese"
}

func (language Portuguese) Greet() string {
	return "I can speak " + language.LanguageName() + ": Olá"
}
