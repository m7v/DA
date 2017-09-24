package main

import (
	"encoding/json"
	"github.com/kataras/iris"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"
	"os"
	"log"
)

type SergejAnswers struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

type Answers struct {
	Answers []SergejAnswers
}

type HomePage struct {
	Bg     string
	Answer string
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	app := iris.New()

	app.Favicon("./static/favicon.ico")
	app.StaticWeb("/static", "./static")

	app.RegisterView(iris.HTML("./views", ".html"))

	app.Get("/", home)

	app.Run(iris.Addr(":" + port), iris.WithCharset("UTF-8"))
}

func home(ctx iris.Context) {
	dat, err := ioutil.ReadFile("./answers.json")
	check(err)

	res := Answers{}
	err = json.Unmarshal(dat, &res)
	check(err)

	max := len(res.Answers)
	min := res.Answers[0].Id - 1

	a := random(min, max)
	bg := random(1, 11)

	ctx.ViewData("", HomePage{Bg: strconv.Itoa(bg), Answer: string(res.Answers[a].Text)})
	ctx.View("home.html")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func random(min int, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
