package model

import "time"

// log-level
var (
	LogLevel      = "log-level"
	LogLevelInfo  = "info"
	LogLevelError = "error"
	LogLevelDebug = "debug"
	LogLevelFatal = "fatal"
	LogLevelWarn  = "warn"
)

// package name

var (
	ApiPackage         = "api"
	StorePackage       = "store"
	ModelPackage       = "model"
	ControllersPackage = "controllers"
	UtilsPackage       = "utils"
	MainPackage        = "main"
)

// function name

var (
	Controllers = "controllers"
	Store       = "store"
	Cmd         = "cmd"
	Utils       = "utils"
	Api         = "api"
)

var (
	NewServer       = "new-server"
	NewStore        = "new-store"
	CreateUser      = "create-user"
	GetUsers        = "get-users"
	GetUserByFilter = "get-user-by-filter"
	GetUserByID     = "get-user-by-id"
	UpdateUser      = "update-user"
	DeleteUser      = "delete-user"
	SignUP          = "signup"
	SignIn          = "signin"
	AuthMiddleware  = "authmiddleware"

	CreateCollege      = "create-college"
	GetCollege         = "get-college"
	GetCollegeByFilter = "get-college-by-filter"
	GetCollegeByID     = "get-college-by-id"
	UpdateCollege      = "update-college"
	DeleteCollege      = "delete-college"

	CreateClass      = "create-class"
	GetClass         = "get-class"
	GetClassByFilter = "get-class-by-filter"
	GetClassByID     = "get-class-by-id"
	UpdateClass      = "update-class"
	DeleteClass      = "delete-class"

	CreateTeacher      = "create-Teacher"
	GetTeacher         = "get-Teacher"
	GetTeacherByFilter = "get-Teacher-by-filter"
	GetTeacherByID     = "get-Teacher-by-id"
	UpdateTeacher      = "update-Teacher"
	DeleteTeacher      = "delete-Teacher"

	SetLimitAndPage = "setlimitandpage"
	SetDateRangeFilter = "setdaterangefilter"
)

var (
	Value    = "value"
	Email    = "email"
	Password = "password"
	UserID   = "userID"
	Expire   = "exp"

	Authorization = "X-Token"

	DNS = "host=localhost user=yash1 password=password1 dbname=project1 port=5432 sslmode=disable"

	DataPerPage = "limit"
	PageNumber  = "page"
	StartDate   = "start_date"
	EndDate     = "end_date"
	TimeLayout  = "2006-01-02 15:04:05.000 -0700"
)

var (
	TokenExpiration = time.Hour * 24
)

var SecretKey = []byte("college-management-key")

var (
	SuperAdminUser = "superadmin"
	AdminUser      = "adminuser"
	NormalUser     = "normaluser"
)
