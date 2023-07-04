package initializers

import "github.com/cosmos-sajal/go_boilerplate/crons"

func InitialiseCron() error {
	return crons.Initialise()
}
