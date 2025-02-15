// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"context"
	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
)

// Dao dao.
type Dao struct {
	*Mysql
	sqlc.CachedConn
}

// New new a dao and return.
func New(db *sqlx.DB, c sqlc.CachedConn) (dao *Dao) {
	dao = &Dao{
		Mysql:      newMysqlDao(db),
		CachedConn: c,
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.Mysql.Close()
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	return d.Mysql.Ping(ctx)
}
