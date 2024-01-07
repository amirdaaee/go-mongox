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
	t.Run("test arrayElemAt", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "first", Value: bson.D{{Key: "$arrayElemAt", Value: []any{"$favorites", int64(0)}}}}},
			BsonBuilder().ArrayElemAt("first", "$favorites", 0).Build())
	})
}

func Test_arrayBuilder_ConcatArrays(t *testing.T) {
	t.Run("test concatArrays", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "favorites", Value: bson.D{{Key: "$concatArrays", Value: []any{"$favorites", "$hobbies"}}}}},
			BsonBuilder().ConcatArrays("favorites", "$favorites", "$hobbies").Build(),
		)
	})
}

func Test_arrayBuilder_ArrayToObject(t *testing.T) {
	t.Run("test arrayToObject", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "favorites", Value: bson.D{{Key: "$arrayToObject", Value: "$favorites"}}}},
			BsonBuilder().ArrayToObject("favorites", "$favorites").Build(),
		)
	})
}

func Test_arrayBuilder_Size(t *testing.T) {
	t.Run("test size", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "favorites", Value: bson.D{{Key: "$size", Value: "$favorites"}}}},
			BsonBuilder().Size("favorites", "$favorites").Build(),
		)
	})
}

func Test_arrayBuilder_Slice(t *testing.T) {
	t.Run("test slice", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "favorites", Value: bson.D{{Key: "$slice", Value: []any{"$favorites", int64(5)}}}}},
			BsonBuilder().Slice("favorites", "$favorites", 5).Build(),
		)
	})
}

func Test_arrayBuilder_SliceWithPosition(t *testing.T) {
	t.Run("test slice with position", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "favorites", Value: bson.D{{Key: "$slice", Value: []any{"$favorites", int64(2), int64(5)}}}}},
			BsonBuilder().SliceWithPosition("favorites", "$favorites", 2, 5).Build(),
		)
	})
}

func Test_arrayBuilder_Map(t *testing.T) {
	t.Run("test map", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "favorites", Value: bson.D{{Key: "$map", Value: bson.D{{Key: "input", Value: "$items"}, {Key: "as", Value: "item"}, {Key: "in", Value: "$$item.price * 1.25"}}}}}},
			BsonBuilder().Map("favorites", "$items", "item", "$$item.price * 1.25").Build(),
		)
	})
}

func Test_arrayBuilder_Filter(t *testing.T) {
	testCases := []struct {
		name       string
		key        string
		inputArray any
		cond       any
		opt        *types.FilterOptions
		expected   bson.D
	}{
		{
			name:       "nil options",
			key:        "items",
			inputArray: "$items",
			cond:       "$$item.price > 100",
			opt:        nil,
			expected:   bson.D{bson.E{Key: "items", Value: bson.D{{Key: "$filter", Value: bson.D{{Key: "input", Value: "$items"}, {Key: "cond", Value: "$$item.price > 100"}}}}}},
		},
		{
			name:       "with options",
			key:        "items",
			inputArray: "$items",
			cond:       "$$item.price > 100",
			opt:        &types.FilterOptions{As: "item", Limit: 5},
			expected:   bson.D{bson.E{Key: "items", Value: bson.D{{Key: "$filter", Value: bson.D{{Key: "input", Value: "$items"}, {Key: "cond", Value: "$$item.price > 100"}, {Key: "as", Value: "item"}, {Key: "limit", Value: int64(5)}}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().Filter(tc.key, tc.inputArray, tc.cond, tc.opt).Build())
		})
	}
}
