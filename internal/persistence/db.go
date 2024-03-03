package persistence

import (
	"context"
	"database/sql"
	"log"
)

const (
	INIT_TABLE = `
		CREATE TABLE thumbnails IF NOT EXISTS(
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

func initTable(ctx context.Context, db *sql.DB) {
	_, err := db.ExecContext(ctx, INIT_TABLE)
	if err != nil {
		log.Fatalln("Error while init sqlite cache")
	}
}

func cacheThumbnail(ctx context.Context, db *sql.DB, video_id string, thumbnailBytes []byte) error {
	transaction, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("Error while caching thumbnail | %s\n", video_id)
		return err
	}

	defer func() {
		if panicValue := recover(); panicValue != nil {
			transaction.Rollback()
		}
	}()

	_, err = transaction.Exec(CACHE_THUMBNAIL, video_id, thumbnailBytes)

	if err != nil {
		transaction.Rollback()
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
	return isThumbnailCached
}
