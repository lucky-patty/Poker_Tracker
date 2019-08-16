package main

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

// todo
func main(){
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Define the HTTP routes
    e.GET("/players",func(c echo.Context) error{
      return c.JSON(200, "GET Players")
    })

    e.PUT("/players", func(c echo.Context) error{
      return c.JSON(200, "PUT Players")
    })

    e.PUT("/players/:id", func(c echo.Context) error{
      return c.JSON(200, "UPDATE Player " + c.Param("id"))
    })

    // Start Server
    e.Logger.Fatal(e.Start(":9000"))
}
