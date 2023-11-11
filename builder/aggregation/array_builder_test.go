// Copyright 2023 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aggregation

import (
	"testing"

	"github.com/chenmingyong0423/go-mongox/types"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_arrayBuilder_ArrayElemAt(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		index      int64
		expected   bson.D
	}{
		{
			name:       "valid expression",
			expression: "$favorites",
			index:      0,
			expected:   bson.D{{Key: "$arrayElemAt", Value: []any{"$favorites", int64(0)}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().ArrayElemAt(tc.expression, tc.index).Build())
		})
	}
}

func Test_arrayBuilder_ConcatArrays(t *testing.T) {
	testCases := []struct {
		name     string
		arrays   []any
		expected bson.D
	}{
		{
			name:     "valid arrays",
			arrays:   []any{"$instock", "$ordered"},
			expected: bson.D{{Key: "$concatArrays", Value: []any{"$instock", "$ordered"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().ConcatArrays(tc.arrays...).Build())
		})
	}
}

func Test_arrayBuilder_ArrayToObject(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		expected   bson.D
	}{
		{
			name:       "string expression",
			expression: "$dimensions",
			expected:   bson.D{{Key: "$arrayToObject", Value: "$dimensions"}},
		},
		{
			name:       "array expression",
			expression: []any{BsonBuilder().AddKeyValues(types.KV("k", "item"), types.KV("v", "abc123")).Build()},
			expected:   bson.D{{Key: "$arrayToObject", Value: []any{bson.D{{Key: "k", Value: "item"}, {Key: "v", Value: "abc123"}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().ArrayToObject(tc.expression).Build())
		})
	}
}

func Test_arrayBuilder_Size(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		expected   bson.D
	}{
		{
			name:       "valid expression",
			expression: "$items",
			expected:   bson.D{{Key: "$size", Value: "$items"}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().Size(tc.expression).Build())
		})
	}
}

func Test_arrayBuilder_Slice(t *testing.T) {
	testCases := []struct {
		name      string
		array     any
		nElements int64
		expected  bson.D
	}{
		{
			name:      "valid expression",
			array:     "$items",
			nElements: 5,
			expected:  bson.D{{Key: "$slice", Value: []any{"$items", int64(5)}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().Slice(tc.array, tc.nElements).Build())
		})
	}
}

func Test_arrayBuilder_SliceWithPosition(t *testing.T) {
	testCases := []struct {
		name      string
		array     any
		position  int64
		nElements int64
		expected  bson.D
	}{
		{
			name:      "valid expression",
			array:     "$items",
			position:  20,
			nElements: 5,
			expected:  bson.D{{Key: "$slice", Value: []any{"$items", int64(20), int64(5)}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().SliceWithPosition(tc.array, tc.position, tc.nElements).Build())
		})
	}
}
