package poker

import (
	"bufio"
	"io"
	"strings"
	"time"
)

type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
	alerter     BlindAlerter
}

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

func NewCLI(store PlayerStore, in io.Reader, alerter BlindAlerter) *CLI {
	return &CLI{
		playerStore: store,
		in:          bufio.NewScanner(in),
		alerter:     alerter,
	}

}

func (cli *CLI) scheduleBlindAlerts() {

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		cli.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + 10*time.Minute
	}

}
func (cli *CLI) PlayPoker() {

	cli.scheduleBlindAlerts()

	userInput := cli.readLine()			
		amountGot := alert.amount
	if amountGot != c.expectedAmount {
		t.Errorf("got amount %d,want %d", amountGot, c.expectedAmount)
	}

	gotScheduledTime := alert.at
	if gotScheduledTime != c.expectedScheduleTime {
		t.Errorf("got scheduled time of %v, wanted %v", gotScheduledTime, c.expectedScheduleTime)
	}1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
