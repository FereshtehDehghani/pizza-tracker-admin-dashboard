package models


type DBModel struct {
	order OrderModel
}

func InitDB(dataSourceName string)(*DBModel,error){
	db,err := gorm.Open(sqlit.open(dataSourceName,&gorm.Config{}))

	if err != nil{
		return nil, fmt.Error("Failed to migrate database: %v",err)
	}

	err =db.AutoMigrate(&Order{},&OrderItem{})

	if err != nil{
		return nil, fmt.Error("Failed to migrate database: %v",err)
	}

	dbModel := &DBModel{
		order:OrderModel{DB:db}
	}

	return dbModel,nil
}