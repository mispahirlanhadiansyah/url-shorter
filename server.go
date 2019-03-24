package main

import (
	"encoding/json"
	"net/http"
	"time"
	"url-shorter/util"

	"github.com/labstack/echo"
)

var shorten_data map[string]shortenStorage = map[string]shortenStorage{}
var shorten_check map[string]interface{} = map[string]interface{}{}

type shortenStorage struct {
	Url           string
	Shortcode     string
	Stardate      time.Time
	Lastseendate  time.Time
	Redirectcount int
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Say,Hello World!")
	})

	e.POST("/shorten", shortencreate)
	e.GET("/:shortcode", shortenget)
	e.GET("/:shortcode/stats", shortstatus)

	e.Logger.Fatal(e.Start(":1323"))

}

func shortstatus(c echo.Context) error {
	shortencode := c.Param("shortcode")
	shortobject := shorten_data[shortencode]
	map_response := make(map[string]interface{})
	map_response["stardate"] = shortobject.Stardate
	map_response["lastseendate"] = shortobject.Lastseendate
	map_response["redirecount"] = shortobject.Redirectcount
	return c.JSON(http.StatusFound, map_response)

}

func shortenget(c echo.Context) error {
	shortencode := c.Param("shortcode")
	shortobject := shorten_data[shortencode]
	shortobject.Redirectcount++
	shortobject.Lastseendate = time.Now()
	shorten_data[shortencode] = shortobject
	c.Response().Header().Set("Location", shortobject.Url)
	c.Response().WriteHeader(http.StatusFound)
	return json.NewEncoder(c.Response()).Encode(shortobject.Url)
}

func shortencreate(c echo.Context) error {
	map_request := make(map[string]interface{})

	c.Bind(&map_request)
	map_response := make(map[string]interface{})
	shorten_url := map_request["url"]
	shorten_code := map_request["shortcode"]
	if shorten_url == nil || shorten_code == nil {
		map_response["error"] = "url or shortcode is not present"
		return c.JSON(http.StatusBadRequest, map_response)
	}

	if shorten_check[shorten_code.(string)] != nil {
		map_response["error"] = "is already in used"
		return c.JSON(http.StatusConflict, map_response)
	}

	check_regex := util.Shortener{}.ValidateUrl(shorten_code.(string))
	if check_regex == false {
		map_response["error"] = "fail meet regex ^[0-9a-zA-Z_]{4,}$"
		return c.JSON(http.StatusUnprocessableEntity, map_response)
	}

	shorten_object := shortenStorage{
		shorten_url.(string),
		shorten_code.(string),
		time.Now(),
		time.Now(),
		0}
	random_string := util.Shortener{}.GetRandomString()
	shorten_data[random_string] = shorten_object
	map_response["shortcode"] = random_string
	shorten_check[shorten_code.(string)] = shorten_url
	return c.JSON(http.StatusCreated, map_response)
}
