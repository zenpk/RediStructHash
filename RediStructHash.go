// Copyright 2022 zenpk
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"reflect"
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
		tag := ref.Type().Field(i).Tag.Get("redistructhash")
		if tag == "no" {
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
