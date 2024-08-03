package infrastructure

import (
	"context"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var i18Bundle *i18n.Bundle

func NewLocalizer() {
	i18Bundle = i18n.NewBundle(language.English)
	i18Bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	if _, err := i18Bundle.LoadMessageFile("./locale/en.toml"); err != nil {
		fmt.Printf("localize: %v", err.Error())
		os.Exit(1)
	}
	if _, err := i18Bundle.LoadMessageFile("./locale/id.toml"); err != nil {
		fmt.Printf("localize: %v", err.Error())
		os.Exit(1)
	}
	fmt.Println("localize: load successfully")
}

var i18Localizer *i18n.Localizer

func LocalizerMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		header := c.GetHeader("Accept-Language")
		query := c.Query("lang")

		i18Localizer = i18n.NewLocalizer(i18Bundle, string(header), query)
		c.Next(ctx)
	}
}

func Localize(params any) string {
	switch p := params.(type) {
	case string:
		return i18Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: p,
		})
	default:
		return i18Localizer.MustLocalize(p.(*i18n.LocalizeConfig))
	}
}