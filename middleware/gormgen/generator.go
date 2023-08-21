package gormgen

import (
	"main/dao"
	"main/model"

	"gorm.io/gen"
)

// References:
// https://gorm.io/gen/index.html
// https://juejin.cn/post/7143968556609175583

func DaoGenerator() {
	// Connect database
	db := dao.InitDB()

	// Automatically migrate models into mysql database
	db.AutoMigrate(&model.Users{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Users{})

	// Automatically generate queries
	g := gen.NewGenerator(gen.Config{
		OutPath: "dao",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(db)
	g.ApplyBasic(&model.Users{})
	g.Execute()
}
