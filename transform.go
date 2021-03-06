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
	"regexp"
	"strings"
)

// Transform is a function definition for transforming values and search terms.
type Transform func(interface{}) interface{}

// ToLower transforms all Unicode letters mapped to their lower case.
// It only transforms objects that conform to the Stringer interface.
func ToLower(terms interface{}) interface{} {
	switch terms.(type) {
	case string:
		return strings.ToLower(terms.(string))
	case Key:
		return strings.ToLower(terms.(Key).String())
	case []byte:
		return strings.ToLower(string(terms.([]byte)))
	default:
		return terms
	}
}

// Tokenizer is a function definition that transforms a text into tokens
type Tokenizer func(string) []string

const nonWhitespaceRegex = `\S+`

// WhiteSpaceTokenizer tokenizes the string based on the /\S/g regex
func WhiteSpaceTokenizer(text string) []string {
	exp, _ := regexp.Compile(nonWhitespaceRegex)
	return exp.FindAllString(text, -1)
}
