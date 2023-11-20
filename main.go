package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serve static files
	r.Static("/static", "./static")

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// Define routes
	setupRoutes(r)

	// Start the server
	r.Run()
}

func setupRoutes(r *gin.Engine) {
	// Root route
	r.GET("/", homeHandler)
	r.GET("/overview", OverviewHandler)
	r.GET("/toggle-content", ToggleContentHandler)
}

func homeHandler(c *gin.Context) {
	name := struct {
		Title       string
		Description string
	}{
		Title:       "Varmint Cong",
		Description: "Simple Go Web Starter Kit",
	}

	// Render the HTML template
	c.HTML(http.StatusOK, "index.html", name)

}

func OverviewHandler(c *gin.Context) {
	name := struct {
		Title string
	}{
		Title: "Project Overview",
	}

	// Render the HTML template for the project overview
	c.HTML(http.StatusOK, "overview.html", name)
}

func ToggleContentHandler(c *gin.Context) {
	// You can add logic here to determine the content to be toggled
	name := struct {
		Title string
	}{
		Title: "Toggled Content",
	}

	// Render the HTML template for the toggled content
	c.HTML(http.StatusOK, "toggle-content.html", name)
}
