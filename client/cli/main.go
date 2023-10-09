package main

import (
	"flag"
	"fmt"

	"github.com/igorkichuk/tucows/api/img"
	"github.com/igorkichuk/tucows/api/quote"
	"github.com/igorkichuk/tucows/client/cli/controller"
	"github.com/igorkichuk/tucows/client/cli/handler"
	"github.com/igorkichuk/tucows/common"

	"github.com/caarlos0/env/v9"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/term"
)

func main() {
	cfg := common.Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Println(err)
	}

	key := flag.Int("key", common.DefaultKey, "Numeric key, which influences the choice of quotation, the maximum length is 6 characters")
	grayscale := flag.Bool("grayscale", common.DefaultGrayscale, "Draws grayscale image if set true.")
	flag.Parse()

	termWidth := getTermWidth()

	c := controller.NewPostController(
		cfg,
		quote.NewQuoteProvider(common.DefaultHTTPClient, quote.TextFormat),
		img.NewPicsumImageProvider(common.DefaultHTTPClient),
	)
	h := handler.NewPostHandler(c)
	h.ShowRandomPost(*grayscale, termWidth, *key)
}

func getTermWidth() int {
	if !term.IsTerminal(0) {
		return common.DefaultTermWidth
	}

	width, _, err := term.GetSize(0)
	if err != nil || width == 0 {
		return common.DefaultTermWidth
	}

	return width
}
