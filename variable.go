package config

import (
	"github.com/xybor/x/logging"
	gormlogger "gorm.io/gorm/logger"
)

type Variable struct {
	Server         ServerVariable         `envconfig:"server"`
	Postgres       PostgresVariable       `envconfig:"postgres"`
	Redis          RedisVariable          `envconfig:"redis"`
	Authentication AuthenticationVariable `envconfig:"authentication"`
	OAuth2         OAuth2Variable         `envconfig:"oauth2"`
	Session        SessionVariable        `envconfig:"session"`
}

func DefaultVariable() Variable {
	return Variable{
		Server:         DefaultServerVariable(),
		Postgres:       DefaultPostgresVariable(),
		Redis:          DefaultRedisVariable(),
		Authentication: DefaultAuthenticationVariable(),
		OAuth2:         DefaultOAuth2Variable(),
		Session:        DefaultSessionVariable(),
	}
}

type ServerVariable struct {
	Host           string `envconfig:"host"`
	Port           int    `envconfig:"port"`
	NodeID         int    `envconfig:"nodeid"`
	LogLevel       int    `envconfig:"loglevel"`
	RequestTimeout int    `envconfig:"timeout"` // The timeout of each request (in millisecond).
}

func DefaultServerVariable() ServerVariable {
	return ServerVariable{
		Host:           "0.0.0.0",
		Port:           8080,
		NodeID:         0,
		LogLevel:       int(logging.LevelDebug),
		RequestTimeout: 3000, // 3s
	}
}

type PostgresVariable struct {
	LogLevel      int    `envconfig:"loglevel"`
	Host          string `envconfig:"host"`
	Port          int    `envconfig:"port"`
	SSLMode       string `envconfig:"sslmode"`
	Timezone      string `envconfig:"timezone"`
	RetryAttempts int    `envconfig:"retry_attempts"`
	RetryInterval int    `envconfig:"retry_interval"` // in second
}

func DefaultPostgresVariable() PostgresVariable {
	return PostgresVariable{
		LogLevel:      int(gormlogger.Warn),
		Host:          "localhost",
		Port:          5432,
		SSLMode:       "disable",
		RetryAttempts: 3,
		RetryInterval: 1,
	}
}

type RedisVariable struct {
	Addr string `envconfig:"addr"`
	DB   int    `envconfig:"db"`
}

func DefaultRedisVariable() RedisVariable {
	return RedisVariable{
		Addr: "localhost:6379",
		DB:   0,
	}
}

type AuthenticationVariable struct {
	AccessTokenExpiration  int    `envconfig:"access_token_expiration"`  // in second
	RefreshTokenExpiration int    `envconfig:"refresh_token_expiration"` // in second
	IDTokenExpiration      int    `envconfig:"id_token_expiration"`      // in second
	TokenIssuer            string `envconfig:"token_issuer"`
}

func DefaultAuthenticationVariable() AuthenticationVariable {
	return AuthenticationVariable{
		AccessTokenExpiration:  60,           // 60s
		RefreshTokenExpiration: 60 * 60,      // 1h
		IDTokenExpiration:      24 * 60 * 60, // 1d
	}
}

type OAuth2Variable struct {
	IdPLoginURL                      string `envconfig:"idp_login_url"`
	ClientSecretLength               int    `envconfig:"client_secret_length"`
	AuthorizationCodeFlowExpiration  int    `envconfig:"authorization_code_flow_expiration"` // in second
	AuthenticationCallbackExpiration int    `envconfig:"authentication_callback_expiration"` // in second
	SessionUpdateExpiration          int    `envconfig:"session_update_expiration"`          // in second
}

func DefaultOAuth2Variable() OAuth2Variable {
	return OAuth2Variable{
		IdPLoginURL:                      "http://localhost:7063/login",
		ClientSecretLength:               64,
		AuthorizationCodeFlowExpiration:  10 * 60, // 10m
		AuthenticationCallbackExpiration: 5 * 60,  // 5m
		SessionUpdateExpiration:          15,      // 15s
	}
}

type SessionVariable struct {
	Expiration int `envconfig:"expiration"`
}

func DefaultSessionVariable() SessionVariable {
	return SessionVariable{
		Expiration: 24 * 60 * 60, // 24h
	}
}
