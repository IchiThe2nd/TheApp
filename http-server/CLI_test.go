package poker_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	poker "github.com/IchiThe2nd/TheApp/http-server"
)

var dummySpyAlerter = &SpyBlindAlerter{}

type scheduledAlert struct {
	at     time.Duration
	amount int
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{at, amount})
}

func TestCLI(t *testing.T) {
	t.Run("recortds chris from input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})
	t.Run("records cleo from input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})
	t.Run("it schedules printing in  blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerstore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(playerstore, in, blindAlerter)
		cli.PlayPoker()

		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, want := range cases {

			t.Run(fmt.Sprint(want), func(t *testing.T) {

				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
				}
				got := blindAlerter.alerts[i]

				assertScheduledAlert(t, got, want)
			})
		}
	})
}

func assertScheduledAlert(t testing.TB, gotAlert scheduledAlert, wantAlert scheduledAlert) {
	t.Helper()
	if gotAlert.amount != wantAlert.amount {
		t.Errorf("got amount %d,want %d", gotAlert.amount, wantAlert.amount)
	}
	if gotAlert.at != wantAlert.at {
		t.Errorf(" got %d time and wanted %d ", gotAlert.at, wantAlert.at)
	}
	return
}
