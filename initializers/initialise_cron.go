package initializers

import "github.com/cosmos-sajal/go_boilerplate/crons"

func InitialiseCron() {
	err := crons.Initialise()
	if err != nil {
		panic(err)
	}
}
