package controllers

import (
	"net/http"
	"strings"
	"time"

	cache "github.com/patrickmn/go-cache"
)

type Controller struct {
	cache *cache.Cache
}

func New() *Controller {
	cache := cache.New(5*time.Minute, 10*time.Minute)

	return &Controller{
		cache: cache,
	}
}

func (c *Controller) HandleRequests(w http.ResponseWriter, r *http.Request) {
	p := strings.Split(r.URL.Path, "/")[1:]
	n := len(p)

	switch {
	case n == 3 && p[2] == "register" && r.Method == "POST":
		c.postRegister(w, r)
	case n == 3 && p[2] == "login" && r.Method == "POST":
		c.postLogin(w, r)
	default:
		http.NotFound(w, r)
	}
}