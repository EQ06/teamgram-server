/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dataobject

type UserSettingsDO struct {
	Id      int64  `db:"id"`
	UserId  int64  `db:"user_id"`
	Key2    string `db:"key2"`
	Value   string `db:"value"`
	Deleted bool   `db:"deleted"`
}
