package text

import (
	"flag"
	"fmt"
)

type Text struct {
	Message string
	HelpF   bool
}

func (s *Text) Name() string {
	return "text"
}

func (s *Text) Example() string {
	return `
  isaacli text --msg teste
  `
}

func (s *Text) Help() string {
	return "Responsável por exibir um texto"
}

func (s *Text) LongHelp() string {
	return `
	Responsável por exibir um texto

  isaacli text --msg [text]
  - text: string
  `
}

func (s *Text) Register(fs *flag.FlagSet) {
	fs.StringVar(&s.Message, "msg", "", "texto a ser exibido")
	fs.BoolVar(&s.HelpF, "help", false, "comando de ajuda")
}

func (s *Text) Run() {
	if s.HelpF {
		fmt.Println(s.LongHelp())
		return
	}

	fmt.Printf("Mensagem: %s", s.Message)
}
