package repository

import "context"

type KeyValuerepository interface {
	Set(ctx context.Context, key, value string)
	Get(ctx context.Context, key string) (string, bool)
	Del(ctx context.Context, key string) int
	Expire(ctx context.Context, key string, seconds int) bool
	TTl(ctx context.Context, key string) int64
	Persist(ctx context.Context, key string) bool
	Keys(ctx context.Context, pattern string) bool
	Exists(ctx context.Context, key string) bool
	Size(ctx context.Context) int
	StartCleanup(intervalMs int64)
	StopClenup()
}
