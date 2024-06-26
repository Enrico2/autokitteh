package authloginhttpsvc

import (
	"github.com/dghubble/gologin/v2"

	"go.autokitteh.dev/autokitteh/internal/backend/configset"
)

type oauth2Config struct {
	Enabled      bool                  `koanf:"enabled"`
	ClientID     string                `koanf:"client_id"`
	ClientSecret string                `koanf:"client_secret"`
	RedirectURL  string                `koanf:"redirect_url"`
	Cookie       *gologin.CookieConfig `koanf:"cookie"`
}

func (c oauth2Config) cookieConfig() gologin.CookieConfig {
	if c.Cookie == nil {
		return gologin.DefaultCookieConfig
	}

	return *c.Cookie
}

type descopeConfig struct {
	Enabled       bool   `koanf:"enabled"`
	ProjectID     string `koanf:"project_id"`
	ManagementKey string `koanf:"management_key"`
}

type Config struct {
	GoogleOAuth oauth2Config  `koanf:"google_oauth"`
	GithubOAuth oauth2Config  `konf:"github_oauth"`
	Descope     descopeConfig `koanf:"descope"`

	// Allowed login patterns, separated by commas.
	// Pattern format is either of:
	// - "*"       - matches any login
	// - "*@host"  - matches any login from host
	// - otherwise - matches exact login
	AllowedLogins string `koanf:"allowed_logins"`
}

var Configs = configset.Set[Config]{
	Default: &Config{},
	Dev: &Config{
		GoogleOAuth: oauth2Config{
			RedirectURL: "http://localhost:9980/auth/google/callback",
			Cookie:      &gologin.DebugOnlyCookieConfig,
		},
		GithubOAuth: oauth2Config{
			RedirectURL: "http://localhost:9980/auth/github/callback",
			Cookie:      &gologin.DebugOnlyCookieConfig,
		},
	},
}
