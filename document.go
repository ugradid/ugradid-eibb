/*
 * Copyright (c) 2021 ugradid community
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <https://www.gnu.org/licenses/>.
 */

package eibb

import (
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
)

// Document represents a JSON document in []byte format
type Document struct {
	raw []byte
}

// Bytes returns the raw bytes
func (d Document) Bytes() []byte {
	return d.raw
}

// DocumentFromString creates a Document from a JSON string
func DocumentFromString(json string) Document {
	return Document{raw: []byte(json)}
}

// DocumentFromBytes creates a Document from a JSON string
func DocumentFromBytes(json []byte) Document {
	return Document{raw: json}
}

// ErrInvalidJSON is returned when invalid JSON is parsed
var ErrInvalidJSON = errors.New("invalid json")

// KeysAtPath returns the values found at the JSON path query as Key
func (d Document) KeysAtPath(jsonPath string) ([]Key, error) {
	rawKeys, err := d.ValuesAtPath(jsonPath)
	if err != nil {
		return nil, err
	}

	keys := make([]Key, len(rawKeys))
	for i, rk := range rawKeys {
		// valuesFromResult has already filtered values that are not supported by toBytes
		key, _ := toBytes(rk)
		keys[i] = key
	}
	return keys, nil
}

// ValuesAtPath returns a slice with the values found at the given JSON path query
func (d Document) ValuesAtPath(jsonPath string) ([]interface{}, error) {
	if !gjson.ValidBytes(d.raw) {
		return nil, ErrInvalidJSON
	}
	result := gjson.GetBytes(d.raw, jsonPath)

	return valuesFromResult(result)
}

func valuesFromResult(result gjson.Result) ([]interface{}, error) {
	switch result.Type {
	case gjson.String:
		return []interface{}{result.Str}, nil
	case gjson.Number:
		return []interface{}{result.Num}, nil
	case gjson.Null:
		return []interface{}{}, nil
	default:
		if result.IsArray() {
			keys := make([]interface{}, 0)
			for _, subResult := range result.Array() {
				subKeys, err := valuesFromResult(subResult)
				if err != nil {
					return nil, err
				}
				keys = append(keys, subKeys...)
			}
			return keys, nil
		}
	}
	return nil, fmt.Errorf("type at path not supported for indexing: %s", result.String())
}
