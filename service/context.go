package service

import (
	"context"
	"errors"

	"github.com/slapec93/bitrise-attendance/configs"
	"github.com/slapec93/bitrise-attendance/sheets"
)

type tRequestContextKey string

const (
	// ContextKeyConfig ...
	ContextKeyConfig tRequestContextKey = "rck-config"
	// ContextKeySheetsClient ...
	ContextKeySheetsClient tRequestContextKey = "rck-sheets-client"
)

// ContextWithConfig ...
func ContextWithConfig(ctx context.Context, conf configs.Model) context.Context {
	return context.WithValue(ctx, ContextKeyConfig, conf)
}

// GetConfigFromContext ...
func GetConfigFromContext(ctx context.Context) (configs.Model, error) {
	config, ok := ctx.Value(ContextKeyConfig).(configs.Model)
	if !ok {
		return config, errors.New("Config not found in Context")
	}
	return config, nil
}

// ContextWithSheetsClient ...
func ContextWithSheetsClient(ctx context.Context, client sheets.Client) context.Context {
	return context.WithValue(ctx, ContextKeySheetsClient, client)
}

// GetSheetsClientFromContext ...
func GetSheetsClientFromContext(ctx context.Context) (sheets.Client, error) {
	client, ok := ctx.Value(ContextKeySheetsClient).(sheets.Client)
	if !ok {
		return client, errors.New("Sheets Client not found in Context")
	}
	return client, nil
}
