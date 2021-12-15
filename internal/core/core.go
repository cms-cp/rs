package core

import (
	"bytes"
	stdlog "log"
	"os"

	"github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/dustin/go-humanize"
)

type (
	Core interface {
		Run()
	}

	core struct {
		app *fiber.App

		log *stdlog.Logger
	}
)

var (
	buf bytes.Buffer
)

func NewCore() Core {
	f, err := os.OpenFile("/var/log/rs.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		os.Exit(254)
	}

	os.Stdout.Close()
	os.Stdout = f
	os.Stderr.Close()
	os.Stderr = f

	c := &core{}

	c.log = stdlog.New(&buf, "rs: ", stdlog.LstdFlags|stdlog.Lmsgprefix|stdlog.Lshortfile)
	c.log.SetOutput(f)

	c.log.Print(os.Getwd())
	c.log.Print(humanize.Bytes(82854982))

	c.app = fiber.New()

	return c
}

func (c *core) Run() {
	c.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("service a not implementation")
	})

	c.app.Get("/favicon.ico", func(c *fiber.Ctx) error {
		return c.SendString("")
	})

	c.app.Get("/swagger/*", swagger.Handler) // default swagger

	c.app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL: "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:39080/swagger/oauth2-redirect.html",
	}))

	c.app.Listen(":39080")
}
