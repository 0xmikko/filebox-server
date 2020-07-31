/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package config

type Config struct {
	AuthAppleTeamID     string `env:"AUTH_APPLE_TEAM_ID" validate:"required"`
	AuthAppleClientID   string `env:"AUTH_APPLE_CLIENT_ID" validate:"required"`
	AuthAppleKeyID      string `env:"AUTH_APPLE_KEY_ID" validate:"required"`
	AuthAppleSigningKey string `env:"AUTH_APPLE_SIGNING_KEY" validate:"required"`

	AuthJWTSecretKey string `env:"AUTH_JWT_SECRET" validate:"required"`

	DatabaseUrl  string `env:"DATABASE_URL" validate:"required"`
	DatabaseName string `env:"DATABASE_NAME" validate:"required"`

	Env string `env:"ENV" default:"development" validate:"required"`

	IpfsEndpoint string `env:"IPFS_ENDPOINT" validate:"required"`
	TemporaryDir string `env:"TEMP_DIR" default:"tmp/" validate:"required"`

	SentryDSN   string `env:"SENTRY_DSN" validate:"required"`
	SSLRedirect string `env:"SSL_REDIRECT" validate:"required"`

	Port string `env:"PORT" default:"8080" validate:"required"`
}
