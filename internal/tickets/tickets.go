package tickets

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strings"
	"time"
)

type Ticket struct {
	ID          string
	Name        string
	Email       string
	Destination string
	Time        string
	Flight      string
}

// TicketLoader processa tickets linha por linha a partir de um arquivo CSV.
type TicketLoader struct {
	FilePath string
}

// DestinationCounter encapsula a lógica de contagem de tickets para um destino específico.
type DestinationCounter struct {
	Destination      string
	CountDestination int
	Period           string
	CountPeriod      int
	CountTotal       int
}

// TicketProcessor define o contrato para processadores de tickets.
type TicketProcessor interface {
	GetTotalTickets(ticket Ticket) error
	GetCountByPeriod(ticket Ticket) error
}

// LoadAndProcess carrega tickets do arquivo e os processa usando um TicketProcessor.
func (tl *TicketLoader) LoadAndProcess(processor TicketProcessor) error {
	file, err := os.Open(tl.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Ler linha por linha
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// Validar formato
		if len(record) != 6 {
			return errors.New("formato inválido no arquivo CSV")
		}

		// Criar o ticket
		ticket := Ticket{
			ID:          record[0],
			Name:        record[1],
			Email:       record[2],
			Destination: strings.TrimSpace(record[3]),
			Time:        record[4],
			Flight:      record[5],
		}

		// Processar o ticket
		if err := processor.GetTotalTickets(ticket); err != nil {
			return err
		}

		if err := processor.GetCountByPeriod(ticket); err != nil {
			return err
		}
	}

	return nil
}

// GetTotalTickets implementa TicketProcessor para contar tickets.
func (dc *DestinationCounter) GetTotalTickets(ticket Ticket) error {
	dc.CountTotal++
	if strings.EqualFold(ticket.Destination, dc.Destination) {
		dc.CountDestination++
	}
	return nil
}

// GetCountByPeriod implementa TicketProcessor para contar tickets p/ determinados periodos.
func (dc *DestinationCounter) GetCountByPeriod(ticket Ticket) error {
	tempo, err := time.Parse("15:04", ticket.Time)
	if err != nil {
		return err
	}

	tempoMinutos := tempo.Hour()*60 + tempo.Minute()

	switch dc.Period {
	case "Dawn":
		if tempoMinutos >= 0 && tempoMinutos <= 6*60 {
			dc.CountPeriod++
		}
	case "Morning":
		if tempoMinutos > 6*60 && tempoMinutos <= 12*60 {
			dc.CountPeriod++
		}
	case "Afternoon":
		if tempoMinutos > 12*60 && tempoMinutos <= 19*60 {
			dc.CountPeriod++
		}
	case "Night":
		if tempoMinutos > 19*60 && tempoMinutos < 24*60 {
			dc.CountPeriod++
		}
	}
	return nil
}

// AverageDestination retorna o percentual de viagens realizadas para o destino
func AverageDestination(CountDestination, total int) float64 {
	return (float64(CountDestination) / float64(total)) * 100
}
