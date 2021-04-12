// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	CreateVillager(ctx context.Context, arg CreateVillagerParams) (Villager, error)
	GetVillager(ctx context.Context, name string) (Villager, error)
	GetVillagers(ctx context.Context, limit int32) ([]Villager, error)
}

var _ Querier = (*Queries)(nil)
