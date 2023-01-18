package user

//repository interface that repository struct implements
type UserRepository interface {
	GetUserFromDb() string
}

//repository model that constructor returns
type repository struct {
	Database dbInstance
}

//temp struct it can be any database config
type dbInstance struct {
	uri      string
	database string
}

func NewRepository(uri, database string) UserRepository {

	return &repository{
		Database: dbInstance{
			uri:      uri,
			database: database,
		},
	}
}

func (receiver *repository) GetUserFromDb() string {
	return "User Data comes from" + receiver.Database.database
}
