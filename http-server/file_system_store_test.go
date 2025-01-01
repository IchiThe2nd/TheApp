package main

import (
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
		{"Name": "Cleo", "Wins":10},
		{"Name": "Chris", "Wins":33}]`)

		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		got := store.GetLeague()

		want := League{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)

		//read again
		got = store.GetLeague()
		assertLeague(t, got, want)

	})
	t.Run("get player score", func(t *testing.T) {

		database, cleanDatabase := createTempFile(t, `[
	{"Name": "Cleo", "Wins" : 10},
	{"Name": "Chris", "Wins" : 33}	
	]`)

		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		got := store.GetPlayerScore("Chris")
		want := 33
		assertScoreEquals(t, got, want)

	})
	t.Run("records player wins", func(t *testing.T) {

		database, cleanDatabase := createTempFile(t, `[
	{"Name": "Cleo", "Wins" : 10},
	{"Name": "Chris", "Wins" : 33}	
	]`)

		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)
		store.RecordWin("Chris")
		got := store.GetPlayerScore("Chris")
		want := 34
		assertScoreEquals(t, got, want)

	})

	t.Run("records player wins", func(t *testing.T) {

		database, cleanDatabase := createTempFile(t, `[
	{"Name": "Cleo", "Wins" : 10},
	{"Name": "Chris", "Wins" : 33}	
	]`)

		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)
		store.RecordWin("Pepper")
		got := store.GetPlayerScore("Pepper")
		want := 1
		assertScoreEquals(t, got, want)
	})
}

// cant even write an assert wel
func assertScoreEquals(t testing.TB, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d wanted %d", got, want)
	}
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {

	t.Helper()
	tmpfile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("coul not create remp file %v", err)
	}
	tmpfile.Write([]byte(initialData))
	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}
	return tmpfile, removeFile
}
