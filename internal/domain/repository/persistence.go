package repository

import "context"

type PersistenceRepository interface {
	Append(ctx context.Context, command string, args []string) error
	Replay(ctx context.Context, store KeyValuerepository) error
	Close() error
}
