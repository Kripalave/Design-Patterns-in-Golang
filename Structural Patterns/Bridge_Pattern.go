package main

import "fmt"

// Database abstraction interface
type Database interface {
	Connect() error
	Query(query string) (string, error)
	Close() error
}

// SQL implementation
type SQLDatabase struct {
	// Additional SQL-specific fields if needed
}

func (s *SQLDatabase) Connect() error {
	fmt.Println("Connecting to SQL database")
	// Actual SQL connection logic
	return nil
}

func (s *SQLDatabase) Query(query string) (string, error) {
	fmt.Println("Executing SQL query:", query)
	// Actual SQL query execution logic
	return "SQL result", nil
}

func (s *SQLDatabase) Close() error {
	fmt.Println("Closing SQL database connection")
	// Actual SQL connection closing logic
	return nil
}

// NoSQL implementation
type NoSQLDatabase struct {
	// Additional NoSQL-specific fields if needed
}

func (n *NoSQLDatabase) Connect() error {
	fmt.Println("Connecting to NoSQL database")
	// Actual NoSQL connection logic
	return nil
}

func (n *NoSQLDatabase) Query(query string) (string, error) {
	fmt.Println("Executing NoSQL query:", query)
	// Actual NoSQL query execution logic
	return "NoSQL result", nil
}

func (n *NoSQLDatabase) Close() error {
	fmt.Println("Closing NoSQL database connection")
	// Actual NoSQL connection closing logic
	return nil
}

func main() {
	// Client code using the Database abstraction

	// Create instances of concrete implementations
	sqlDB := &SQLDatabase{}
	noSQLDB := &NoSQLDatabase{}

	// Connect, query, and close for SQL
	useDatabase(sqlDB)

	// Connect, query, and close for NoSQL
	useDatabase(noSQLDB)
}

func useDatabase(db Database) {
	db.Connect()
	result, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println("Error querying database:", err)
		return
	}
	fmt.Println("Query result:", result)
	db.Close()
}
