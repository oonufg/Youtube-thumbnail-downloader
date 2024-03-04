package cache

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	INIT_TABLE = `
		CREATE TABLE  IF NOT EXISTS thumbnails(
			video_id TEXT,
			thumbnail BLOB
		);
	`
	CACHE_THUMBNAIL = `
		INSERT INTO thumbnails (video_id, thumbnail)
		VALUES (?, ?);
	`
	GET_THUMBNAIL_FROM_CACHE = `
		SELECT thumbnail FROM  thumbnails
		WHERE video_id = ?;
	`
	IS_THUMBNAIL_CACHED = `
		SELECT EXISTS(SELECT video_id FROM thumbnails WHERE video_id = ?);`
)

func initTableIfNotExists(db *sql.DB) {
	_, err := db.Exec(INIT_TABLE)
	if err != nil {
		log.Fatalf("Error while init sqlite cache %w\n", err)
	}
}

func cacheThumbnail(ctx context.Context, db *sql.DB, videoId string, thumbnailBytes []byte) error {
	transaction, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("Error while caching thumbnail | %s\n", videoId)
		return err
	}

	defer func() {
		if panicValue := recover(); panicValue != nil {
			transaction.Rollback()
		}
	}()

	_, err = transaction.Exec(CACHE_THUMBNAIL, videoId, thumbnailBytes)

	if err != nil {
		transaction.Rollback()
		return err
	}

	err = transaction.Commit()
	if err != nil {
		return err
	}
	return nil
}

func getThumbnailFromCache(ctx context.Context, db *sql.DB, videoId string) ([]byte, error) {
	transaction, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("Error while getting thumbnail from cache | %s\n", videoId)
		return nil, err
	}

	var thumbnailBytes []byte

	err = transaction.QueryRow(GET_THUMBNAIL_FROM_CACHE, videoId).Scan(&thumbnailBytes)
	if err != nil {
		log.Printf("Error reading getting thumbnail from cache | %s\n", videoId)
		return nil, err
	}
	transaction.Commit()
	return thumbnailBytes, nil
}

func isThumbnailCached(ctx context.Context, db *sql.DB, videoId string) bool {
	transaction, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("Error while checking thumbnail in cache | %s\n", videoId)
		return false
	}

	var isThumbnailCached bool

	err = transaction.QueryRow(IS_THUMBNAIL_CACHED, videoId).Scan(&isThumbnailCached)
	if err != nil {
		log.Printf("Error while checking thumbnail from cache | %s\n", videoId)
		return false
	}
	transaction.Commit()
	return isThumbnailCached
}
