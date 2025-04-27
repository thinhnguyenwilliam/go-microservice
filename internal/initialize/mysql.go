package initialize

import (
	"fmt"
	"log"
	"time"

	"github.com/thinhcompany/go-ecommerce-backend-api/global"
	persistentobject "github.com/thinhcompany/go-ecommerce-backend-api/internal/persistentObject"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// checkErr is a helper to panic and log errors
func checkErr(err error, msg string) {
	if err != nil {
		global.Logger.Error(msg, zap.Error(err))
		panic(fmt.Sprintf("%s: %v", msg, err))
	}
}

// InitializeMySQL connects to the MySQL database and sets up GORM
func InitializeMySQL() {
	mysqlConfig := global.Config.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.User,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.DBName,
	)

	var err error
	global.Mdb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: false,
	})

	checkErr(err, "Failed to connect to MySQL database")

	global.Logger.Info("Connected to MySQL database successfully")
	SetPool()
	migrateTables()
}

// SetPool sets up connection pool configuration
func SetPool() {
	sqlDB, err := global.Mdb.DB()
	if err != nil {
		log.Fatalf("failed to get database connection pool: %v", err)
	}

	mysqlConfig := global.Config.MySQL

	sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(mysqlConfig.ConnexMaxLifetime) * time.Second)
}

func migrateTables() {
	// Run AutoMigrate for your models to create or update the tables
	err := global.Mdb.AutoMigrate(&persistentobject.User{},
		&persistentobject.Role{},
	) // Add more models as needed

	if err != nil {
		// Handle the error if migration fails
		global.Logger.Error("AutoMigrate failed", zap.Error(err))
		log.Fatalf("AutoMigrate failed: %v", err)
		return
	}

	// Log success message
	global.Logger.Info("Database migration completed successfully")
}
