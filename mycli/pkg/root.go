package mycli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd representa o comando raiz da aplicação
var rootCmd = &cobra.Command{
	Use:   "mycli",
	Short: "MyCLI é uma ferramenta de linha de comando para demonstração",
	Long: `MyCLI é uma ferramenta de linha de comando escrita em Go.
Esta CLI demonstra como usar o Cobra para construir aplicativos CLI.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bem-vindo ao Desafio Go Bases!")
		fmt.Println("A aplicação lê o arquivo CSV de acordo com a sua necessidade de informaçao para cada país!")
	},
}
var Destination string // Define a variável para o destino padrão

var name string

func init() {
	greetCmd.Flags().StringVarP(&name, "name", "n", "Mundo", "Nome para personalizar a saudação")
	rootCmd.AddCommand(greetCmd)
	// Adicionando a flag para definir o destino
	rootCmd.Flags().StringVarP(&Destination, "destination", "d", "Brazil", "Destino para o processamento")
}

var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Exibe uma mensagem de saudação",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Olá, %s! Seja bem-vindo ao MyCLI!\n", name)
	},
}

// Execute adiciona todos os comandos filhos
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing Zero '%s'\n", err)
		os.Exit(1)
	}
}

/*
var destinationCmd = &cobra.Command{
	Use:   "dest",
	Short: "Roda o programa para o destino idealizado",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Olá, %s! Seja bem-vindo ao MyCLI!\n", name)
	},
}
*/
