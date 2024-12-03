package main

import (
	"fmt"
	"log"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {

	// Caminho do arquivo CSV
	filePath := "tickets.csv"

	// Criar o TicketLoader
	loader := &tickets.TicketLoader{FilePath: filePath}

	// Criar o DestinationCounter para o destino
	counter := &tickets.DestinationCounter{Destination: "United States", Period: "Night"}

	// Processar os tickets usando DestinationCounter
	err := loader.LoadAndProcess(counter)
	if err != nil {
		log.Fatalf("Erro ao processar tickets: %v", err)
	}

	fmt.Printf("Total de tickets para %s: %d\n", counter.Destination, counter.CountDestination)
	fmt.Printf("Total de tickets no período %s: %d\n", counter.Period, counter.CountPeriod)
	fmt.Printf("Percentual de tickets para %s: %.2f%%\n", counter.Destination, tickets.AverageDestination(counter.CountDestination, counter.CountTotal))

}
