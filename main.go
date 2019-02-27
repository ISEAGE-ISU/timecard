package main

import (
	"github.com/julienschmidt/httprouter"
	"strconv"
    "net/http"
    "log"
)

func Index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	parseAll(w)
}

func MakeUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	check(r.ParseForm())
	var t TimeCard

	t.User = r.PostFormValue("user")
	t.Password = r.PostFormValue("password")
	writeTC(&t)
	http.Redirect(w, r, "/", 301)
}

func Timecard(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	parseDB(w, ps.ByName("user"))
}

func Punch(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	check(r.ParseForm())
	week, err := strconv.Atoi(r.PostFormValue("week"))
	check(err)

	day, err := strconv.Atoi(r.PostFormValue("day"))
	check(err)

	time, err := strconv.Atoi(r.PostFormValue("time"))
	check(err)

	tc := readDB(ps.ByName("user"))
	if tc.Password == r.PostFormValue("password") {
		tc.punch(week, day, time)
		writeTC(tc)
	}

	http.Redirect(w, r, "/tc/" + ps.ByName("user"), 301)
}

func main() {
	router := httprouter.New()
    router.GET("/", Index)
	router.POST("/", MakeUser)

    router.GET("/tc/:user", Timecard)
	router.POST("/tc/:user", Punch)

	router.Handler("GET", "/raw/*file", http.StripPrefix("/raw/", http.FileServer(http.Dir(dbdir))))

    log.Fatal(http.ListenAndServe(":80", router))

}
