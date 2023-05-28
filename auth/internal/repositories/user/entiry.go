package user

type user struct {
	Id       int64 // Id is automatically detected as primary key
	Requests int64
	Username string `pg:",unique"`
}
