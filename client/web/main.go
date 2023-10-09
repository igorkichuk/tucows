package main

import (
	"fmt"
	"github.com/caarlos0/env/v9"
	"github.com/igorkichuk/tucows/api/img"
	"github.com/igorkichuk/tucows/api/quote"
	"github.com/igorkichuk/tucows/client/web/controller"
	"github.com/igorkichuk/tucows/client/web/handler"
	"github.com/igorkichuk/tucows/common"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"net/http"
)

func main() {
	cfg := common.Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Println(err)
	}

	c := controller.NewPostController(
		cfg,
		quote.NewQuoteProvider(common.DefaultHTTPClient, quote.HTMLFormat),
		img.NewPicsumImageProvider(common.DefaultHTTPClient))
	h := handler.NewPostHandler(c)
	http.HandleFunc("/post", h.ShowRandomPost)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
