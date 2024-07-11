package models

import "gorm.io/gorm"

type Model struct{}

func New() *Model {
	return &Model{}
}

func (m *Model) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Database{},
		&Job{},
		&Permission{},
		&RolePermission{},
		&Role{},
		&User{},
	)
}

func (m *Model) Seeds(db *gorm.DB) error {
	permission := &Permission{}
	permission.Seed(db)

	role := &Role{}
	role.Seed(db)

	rolePermission := &RolePermission{}
	rolePermission.Seed(db)

	user := &User{}
	user.Seed(db)

	return nil
}
