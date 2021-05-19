package models

// Database Connection Type
type DBConnection struct {
	// host
	Host string

	// database
	Database string

	// username
	User string

	// password
	Password string

	// port
	Port int
}
