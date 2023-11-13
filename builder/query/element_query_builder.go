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

package query

import (
	"github.com/chenmingyong0423/go-mongox/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type elementQueryBuilder struct {
	parent *Builder
}

func (b *elementQueryBuilder) Exists(key string, exists bool) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{bson.E{Key: types.Exists, Value: exists}}})
	return b.parent
}

func (b *elementQueryBuilder) Type(key string, t bsontype.Type) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{bson.E{Key: types.Type, Value: t}}})
	return b.parent
}

func (b *elementQueryBuilder) TypeAlias(key string, alias string) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{bson.E{Key: types.Type, Value: alias}}})
	return b.parent
}

func (b *elementQueryBuilder) TypeArray(key string, ts ...bsontype.Type) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{bson.E{Key: types.Type, Value: ts}}})
	return b.parent
}

func (b *elementQueryBuilder) TypeArrayAlias(key string, aliases ...string) *Builder {
	b.parent.data = append(b.parent.data, bson.E{Key: key, Value: bson.D{bson.E{Key: types.Type, Value: aliases}}})
	return b.parent
}
