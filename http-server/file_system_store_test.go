package main

import (
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("get player score", func(t *testing.T) {

		database, cleanDatabase := createTempFile(t, `[
	{"Name": "Cleo", "Wins" : 10},
	{"Name": "Chris", "Wins" : 33}	
	]`)

		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

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

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)
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

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)
		store.RecordWin("Pepper")
		got := store.GetPlayerScore("Pepper")
		want := 1
		assertScoreEquals(t, got, want)
	})

	t.Run("works withj an emptyfile", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()
		_, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)
	})
	t.Run("league sorted ", func(t *testing.T) {

		database, cleanDatabase := createTempFile(t, `[
		{"Name":"Cleo", "Wins": 10},
		{"Name":"Chris", "Wins": 33}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		assertNoError(t, err)

		got := store.GetLeague()
		want := League{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assertLeague(t, got, want)
		//read again
		got = store.GetLeague()
		assertLeague(t, got, want)
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
func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didnt expect an error but gor %v", err)
	}
}
