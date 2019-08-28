-package main

import (
    // "fmt"
    // "io/ioutil"
    // "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)

type Action struct{
    Name string
    Amount int
}

type Player struct{
    Id int
    Title string
    Score int
}

type Game struct{
    Id int
    MaxBet int
    Players *[]Player
    Starter *Player
    Winner *Player
}


func initDB(filepath string) *sql.DB {
    db, err := sq.Open("sqlite3", filepath)

    if err != nil {
      panic(err)
    }

    if db == nil {
      panic("db nil")
    }

    return db
}


func migrate(db *sql.DB) {
    sql :=
    ` CREATE TABLE IF NOT EXISTS Players(
      id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
      name VARCHAR NOT NULL,
      score INTEGER NOT NULL
      UNIQUE(name)
    );

    `
}

// todo
func main(){
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Define the HTTP routes

    // Players
    e.GET("/players",func(c echo.Context) error{
      return c.JSON(200, "GET Players")
    })

    e.PUT("/players", func(c echo.Context) error{
      return c.JSON(200, "PUT Players")
    })

    e.PUT("/players/:id", func(c echo.Context) error{
      return c.JSON(200, "UPDATE Player " + c.Param("id"))
    })

    e.POST("players/updateScore/:id")




    // Start Server
    e.Logger.Fatal(e.Start(":9000"))
}
