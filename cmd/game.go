package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/sethjback/jagg/pkg/guess"
	"github.com/spf13/cobra"
)

type gameFlags struct {
	rounds int
}

func NewGame() *cobra.Command {
	f := &gameFlags{}

	cmd := &cobra.Command{
		Use:   "new",
		Short: "start a new game",
		RunE:  runGame(f),
	}

	cmd.Flags().IntVarP(&f.rounds, "rounds", "r", 3, "number of rounds")

	return cmd
}

func runGame(f *gameFlags) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if f.rounds < 1 {
			return errors.New("must have at least one round")
		}

		ng := guess.NewNumber(f.rounds, 2)

		won := false

		for ng.RoundsLeft() > 0 {
			fmt.Printf("guess a number 0 and %d\n", ng.Max)

			guess := ""

			survey.AskOne(&survey.Input{Message: "What is your guess?"}, &guess)

			gn, err := strconv.Atoi(guess)
			if err != nil {
				fmt.Printf("I didn't understand that....\n")
			} else if ng.Guess(gn) {
				fmt.Println("you win!!!")
				won = true
				break
			} else {
				fmt.Println("nope...")
			}
		}

		if !won {
			fmt.Printf("sorry, you lost! The number was: %d\n", ng.Secret)
		}

		return nil
	}
}
