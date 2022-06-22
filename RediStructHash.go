package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"reflect"
)

const (
	TagName = "redistructhash"
	NoHash  = "no"
)

// convertCase - from CamelCase to camel_case
func convertCase(in string) string {
	var out []rune
	for i, c := range in {
		if c >= 'A' && c <= 'Z' {
			c += 32
			if i > 0 {
				out = append(out, '_')
			}
		}
		out = append(out, c)
	}
	return string(out)
}

// RedisStructHash - Automatically create hash from struct
func RedisStructHash(rdb *redis.Client, ctx context.Context, t interface{}, key string) error {
	ref := reflect.ValueOf(t)
	for i := 0; i < ref.NumField(); i++ {
		tag := ref.Type().Field(i).Tag.Get(TagName)
		if tag == NoHash {
			continue
		}
		fieldName := ref.Type().Field(i).Name
		dbName := convertCase(fieldName)
		if err := rdb.HSet(ctx, key, dbName, ref.Field(i).Interface()).Err(); err != nil {
			return err
		}
	}
	return nil
}
