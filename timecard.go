package main

import (
	"path/filepath"
	"io/ioutil"
	"encoding/json"
	"os"
)

const dbdir = "./db"

type TimeCard struct {
	User, Password string
	Time           [52][7]int
}

func (t *TimeCard) punch(week, day, time int) {
	t.Time[week][day] = time
}

func writeTC(t *TimeCard) {
	dest := filepath.Join(dbdir, t.User + ".json")
	data, err := json.Marshal(t)
	check(err)

	err = ioutil.WriteFile(dest, data, 0644)
	check(err)
}

func readFile(file string) *TimeCard {
	var out TimeCard
	f, err := os.Open(file)
	check(err)
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	check(err)

	err = json.Unmarshal(data, &out)
	check(err)
	return &out
}

func readDB(user string) *TimeCard {
	return readFile(filepath.Join(dbdir, user + ".json"))
}

func readAll() []*TimeCard{
	var out []*TimeCard
	files, err := ioutil.ReadDir(dbdir)
	check(err)

	for _, f := range files {
		out = append(out, readFile(filepath.Join(dbdir, f.Name())))
	}

	return out
}