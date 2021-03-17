// Copyright 2021 Harran Ali. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package models

import "github.com/gincoat/gincoat/core/database"

//MigrateDB the database
func MigrateDB() {
	db := database.Resolve()
	//add your models to be auto migrated here
	db.AutoMigrate(&Product{})
}
