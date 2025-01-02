package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game *Game
}

func NewCLI(store PlayerStore, in io.Reader, out io.Writer, alerter BlindAlerter) *CLI {
	return &CLI{
		in:  bufio.NewScanner(in),
		out: out,
		game: &Game{
			alerter: alerter,
			store:   store,
		},
	}
}

const PlayerPrompt = "Please enter number of players: "

// PlayPoker starts the game.
func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)

	numberOfPlayersInput := cli.readLine()
	numberOfPlayers, _ := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))

	cli.game.Start(numberOfPlayers)
	winnerInput := cli.readLine()
	winner := extractWinner(winnerInput)

	cli.game.Finish(winner)
}

//	func (cli *CLI) scheduleBlindAlerts(numberOfPlayers int) {
//		blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute
//		blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
//		blindTime := 0 * time.Second
//		for _, blind := range blinds {
//			cli.alerter.ScheduleAlertAt(blindTime, blind)
//			blindTime = blindTime + blindIncrement
//		}
//
// }
func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins\n", "", 1)
}
