package tickets_test

import (
	"testing"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

// Implementação de um processador de tickets para testes.
type mockProcessor struct {
	totalCount  int
	periodCount int
}

func (m *mockProcessor) GetTotalTickets(ticket tickets.Ticket) error {
	m.totalCount++
	return nil
}

func (m *mockProcessor) GetCountByPeriod(ticket tickets.Ticket) error {
	m.periodCount++
	return nil
}

func TestTicketLoader_LoadAndProcess(t *testing.T) {
	tests := []struct {
		name      string
		filePath  string
		processor tickets.TicketProcessor
		wantErr   bool
	}{
		{
			name:      "Sucesso ao carregar tickets",
			filePath:  "../../tickets.csv", // Este deve ser um arquivo CSV válido com 6 colunas
			processor: &mockProcessor{},
			wantErr:   false,
		},
		{
			name:      "Arquivo não encontrado",
			filePath:  "non_existent_file.csv",
			processor: &mockProcessor{},
			wantErr:   true,
		},
		{
			name:      "Formato inválido no CSV",
			filePath:  "invalid_tickets.csv", // Deve conter linhas com formato inválido
			processor: &mockProcessor{},
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tl := &tickets.TicketLoader{
				FilePath: tt.filePath,
			}
			err := tl.LoadAndProcess(tt.processor)
			if (err != nil) != tt.wantErr {
				t.Errorf("TicketLoader.LoadAndProcess() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDestinationCounter_GetTotalTickets(t *testing.T) {
	tests := []struct {
		name                 string
		destination          string
		ticket               tickets.Ticket
		wantCountDestination int
		wantCountTotal       int
	}{
		{
			name:                 "Cliente com destino correspondente",
			destination:          "United States",
			ticket:               tickets.Ticket{Destination: "United States"},
			wantCountDestination: 1,
			wantCountTotal:       1,
		},
		{
			name:                 "Cliente com destino não correspondente",
			destination:          "Canada",
			ticket:               tickets.Ticket{Destination: "United States"},
			wantCountDestination: 0,
			wantCountTotal:       1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dc := &tickets.DestinationCounter{Destination: tt.destination}
			dc.GetTotalTickets(tt.ticket)

			if dc.CountDestination != tt.wantCountDestination {
				t.Errorf("CountDestination = %v, want %v", dc.CountDestination, tt.wantCountDestination)
			}
			if dc.CountTotal != tt.wantCountTotal {
				t.Errorf("CountTotal = %v, want %v", dc.CountTotal, tt.wantCountTotal)
			}
		})
	}
}

func TestDestinationCounter_GetCountByPeriod(t *testing.T) {
	tests := []struct {
		name            string
		period          string
		ticket          tickets.Ticket
		wantCountPeriod int
		wantErr         bool
	}{
		{
			name:            "Cliente no período da noite",
			period:          "Night",
			ticket:          tickets.Ticket{Time: "22:00"},
			wantCountPeriod: 1,
			wantErr:         false,
		},
		{
			name:            "Cliente no período da madrugada",
			period:          "Dawn",
			ticket:          tickets.Ticket{Time: "02:00"},
			wantCountPeriod: 1,
			wantErr:         false,
		},
		{
			name:            "Cliente no período da tarde",
			period:          "Afternoon",
			ticket:          tickets.Ticket{Time: "12:30"},
			wantCountPeriod: 1,
			wantErr:         false,
		},
		{
			name:            "Cliente fora do período",
			period:          "Night",
			ticket:          tickets.Ticket{Time: "10:00"},
			wantCountPeriod: 0,
			wantErr:         false,
		},
		{
			name:            "Cliente no período da manhã",
			period:          "Morning",
			ticket:          tickets.Ticket{Time: "09:00"},
			wantCountPeriod: 1,
			wantErr:         false,
		},
		{
			name:            "Cliente com formato de hora inválido",
			period:          "Night",
			ticket:          tickets.Ticket{Time: "invalid"},
			wantCountPeriod: 0,
			wantErr:         true, // Agora esperamos que um erro aconteça
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dc := &tickets.DestinationCounter{Period: tt.period}

			// Chame o método GetCountByPeriod
			err := dc.GetCountByPeriod(tt.ticket)

			// Verifique se o retorno do erro está de acordo com o esperado
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCountByPeriod() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Verifique a contagem de período
			if dc.CountPeriod != tt.wantCountPeriod {
				t.Errorf("CountPeriod = %v, want %v", dc.CountPeriod, tt.wantCountPeriod)
			}
		})
	}
}
func TestAverageDestination(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			CountDestination int
			total            int
		}
		want float64
	}{
		{
			name: "Teste percentual com valores válidos",
			args: struct {
				CountDestination int
				total            int
			}{CountDestination: 2, total: 10},
			want: 20.0,
		},
		{
			name: "Teste percentual com zero total",
			args: struct {
				CountDestination int
				total            int
			}{CountDestination: 0, total: 10},
			want: 0.0, // Definir que 0/0 deve ser tratado como 0%
		},
		{
			name: "Teste percentual padrão",
			args: struct {
				CountDestination int
				total            int
			}{CountDestination: 1, total: 1},
			want: 100.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tickets.AverageDestination(tt.args.CountDestination, tt.args.total)
			if got != tt.want {
				t.Errorf("AverageDestination() = %v, want %v", got, tt.want)
			}
		})
	}
}
