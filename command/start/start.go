package start

import (
	"flag"
	"fmt"
	"net/http"
)

type Start struct {
	Port    int
	Version string
	HelpF   bool
}

func (s *Start) Name() string {
	return "start"
}

func (s *Start) Example() string {
	return `
  isaacli start --port 3030 --version  1.0.0
  isaacli start --version 1.0.0
  `
}

func (s *Start) Help() string {
	return "Responsável por inicializar a aplicação"
}

func (s *Start) LongHelp() string {
	return `
  Responsável por inicializar a aplicação

  isaacli start --port [serverPort] --version [serverVersion]
  - serverPort: int
  - serverVersion: string
  `
}

func (s *Start) Register(fs *flag.FlagSet) {
	fs.IntVar(&s.Port, "port", 8080, "porta do servidor")
	fs.StringVar(&s.Version, "version", "", "versão do servidor")
	fs.BoolVar(&s.HelpF, "help", false, "comando de ajuda")
}

func (s *Start) Run() {
	if s.HelpF {
		fmt.Println(s.LongHelp())
		return
	}

	if s.Version == "" {
		fmt.Println("[--version] é obrigatório")
		return
	}

	fmt.Printf("Servidor v%s rodando na porta %v", s.Version, s.Port)
	http.ListenAndServe(fmt.Sprintf(":%v", s.Port), nil)
}
