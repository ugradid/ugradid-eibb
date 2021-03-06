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
	"encoding/binary"
	"encoding/hex"
	"errors"
	"math"
)

const boltDBFileMode = 0600
const KeyDelimiter = 0x10

// Reference equals a document hash. In an index, the values are references to docs.
type Reference []byte

// EncodeToString encodes the reference as hex encoded string
func (r Reference) EncodeToString() string {
	return hex.EncodeToString(r)
}

// ByteSize returns the size of the reference, eg: 32 bytes for a sha256
func (r Reference) ByteSize() int {
	return len(r)
}

func toBytes(data interface{}) ([]byte, error) {
	switch data.(type) {
	case []uint8:
		return data.([]byte), nil
	case string:
		return []byte(data.(string)), nil
	case float64:
		var buf [8]byte
		binary.BigEndian.PutUint64(buf[:], math.Float64bits(data.(float64)))
		return buf[:], nil
	}

	return nil, errors.New("couldn't convert data to []byte")
}
