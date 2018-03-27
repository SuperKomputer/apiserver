package dao

func FactoryDao(e string) UserDao {
	return UserImpl{}
}
