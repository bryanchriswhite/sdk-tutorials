package types

import (
	"fmt"
)

// GenesisState - all namespace state that must be provided at genesis
type GenesisState struct {
	WhoisRecords []Whois `json:"whois_records"`
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState(whoIsRecords []Whois) GenesisState {
	return GenesisState{WhoisRecords: nil}
}

// ValidateGenesis validates the namespace genesis parameters
func ValidateGenesis(data GenesisState) error {
	for _, record := range data.WhoisRecords {
		if record.Owner == nil {
			return fmt.Errorf("invalid WhoisRecord: Owner: %s. Error: Missing Owner", record.Owner)
		}
		if record.Value == "" {
			return fmt.Errorf("invalid WhoisRecord: Value: %s. Error: Missing Value", record.Value)
		}
		if record.Price == nil {
			return fmt.Errorf("invalid WhoisRecord: Price: %s. Error: Missing Price", record.Price)
		}
	}
	return nil
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	return GenesisState{
		WhoisRecords: []Whois{},
	}
}

