package benchmark

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);not null"`
}

// insertRecord inserts a new user record into the database
func insertRecord(b *testing.B, db *gorm.DB) {
	user := User{Name: "William"}
	if err := db.Create(&user).Error; err != nil {
		b.Fatalf("Error inserting record: %v", err)
	}
}

// BenchmarkMaxOpenConns1 benchmarks inserting a user with a specified max open connections (1 connection)
func BenchmarkMaxOpenConns1(b *testing.B) {
	// Define DSN and establish DB connection
	dsn := "root:1234@tcp(localhost:3306)/go_eco?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		b.Fatalf("Error opening DB connection: %v", err)
	}

	// Migrate the User table (create it if it doesn't exist)
	if db.Migrator().HasTable(&User{}) {
		if err := db.Migrator().DropTable(&User{}); err != nil {
			b.Fatalf("Error dropping User table: %v", err)
		}
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		b.Fatalf("Error migrating User table: %v", err)
	}

	// Set database connection pool configuration
	sqlDB, err := db.DB()
	if err != nil {
		b.Fatalf("Error getting SQL DB instance: %v", err)
	}

	sqlDB.SetMaxOpenConns(1) // Set max open connections to 1 for this benchmark
	defer sqlDB.Close()      // Close the DB connection once done

	// Benchmark with parallel inserts
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db) // Call insertRecord for each parallel execution
		}
	})
}

// BenchmarkMaxOpenConns10 benchmarks inserting a user with a specified max open connections (10 connections)
func BenchmarkMaxOpenConns10(b *testing.B) {
	// Define DSN and establish DB connection
	dsn := "root:1234@tcp(localhost:3306)/go_eco?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		b.Fatalf("Error opening DB connection: %v", err)
	}

	// Migrate the User table (create it if it doesn't exist)
	if db.Migrator().HasTable(&User{}) {
		if err := db.Migrator().DropTable(&User{}); err != nil {
			b.Fatalf("Error dropping User table: %v", err)
		}
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		b.Fatalf("Error migrating User table: %v", err)
	}

	// Set database connection pool configuration
	sqlDB, err := db.DB()
	if err != nil {
		b.Fatalf("Error getting SQL DB instance: %v", err)
	}

	sqlDB.SetMaxOpenConns(10) // Set max open connections to 10 for this benchmark
	defer sqlDB.Close()       // Close the DB connection once done

	// Benchmark with parallel inserts
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db) // Call insertRecord for each parallel execution
		}
	})
}
