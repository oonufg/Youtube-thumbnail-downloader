package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type ThumbnailCache interface {
	CacheThumbnail(ctx context.Context, videoId string, thumbnailBytes []byte)
	GetThumbnail(ctx context.Context, videoId string) []byte
	IsThumbnailCached(videoId string) bool
	Close()
}

type SqliteCache struct {
	ThumbnailCache
	db *sql.DB
}

func New(path string) (*SqliteCache, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("Can't open database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Can't connect to database: %w", err)
	}
	sqliteCache := &SqliteCache{db: db}
	sqliteCache.initTable()

	return sqliteCache, nil
}

func (cache *SqliteCache) Close() {
	cache.db.Close()
}

func (cache *SqliteCache) CacheThumbnail(ctx context.Context, videoId string, thumbnailBytes []byte) {
	if !isThumbnailCached(ctx, cache.db, videoId) {
		err := cacheThumbnail(ctx, cache.db, videoId, thumbnailBytes)
		if err != nil {
			log.Println(err)
		}
	}
}

func (cache *SqliteCache) GetThumbnail(ctx context.Context, videoId string) []byte {
	thumbnailBytes, _ := getThumbnailFromCache(ctx, cache.db, videoId)
	return thumbnailBytes
}

func (cache *SqliteCache) IsThumbnailCached(videoId string) bool {
	return isThumbnailCached(context.TODO(), cache.db, videoId)
}

func (cache *SqliteCache) initTable() {
	initTable(cache.db)
}
