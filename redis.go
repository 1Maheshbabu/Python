package main

import (
	"fmt"
	"github.com/mediocregopher/radix.v2/redis"
)

func dataStore(connection *redis.Client) error {

	resp := connection.Cmd("HMSET", "Employee_Details", "Name", "Mahesh", "Role", "DevOps Engineer", "Salary", 5, "Age", 25)
	// Check the Err field of the *Resp object for any errors.
	if resp.Err != nil {
		return resp.Err
	}

	fmt.Println("Employee record has been created!")

	return nil
}

func dataRetrieve(connection *redis.Client, key string) error {
	// To Fetch all the key-value pairs stored in the table
	employeeData, err := connection.Cmd("HGETALL", "Employee_Details").Map()
	if err != nil {
		return err
	}

	// To print the employee record as key and values

	for key1, value := range employeeData {
		fmt.Println(key1, ":", value)
	}

	// To Print values of all the key value pairs  as a map object
	fmt.Printf("Employee record is  %+v\n", employeeData)

	// To print value of specific key
	fmt.Printf("Employee %s : %v\n", key, employeeData[key])

	// To Fetch value of particular key from database
	employeeName, err := connection.Cmd("HGET", "Employee_Details", "Name").Str()
	if err != nil {
		return err
	}

	fmt.Printf("Name of the employee is %s\n", employeeName)
	return nil
}

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Printf("Attempt to database connection has failed with error : %v", err)
	}

	// Importantly, use defer to ensure the connection is always properly
	// closed before exiting the main() function.
	defer conn.Close()

	errorMessageStore := dataStore(conn)
	if errorMessageStore != nil {
		fmt.Printf("Attempt to entry into the database is failed with error : %v", errorMessageStore)
	}

	errorMessageRetrieval := dataRetrieve(conn, "Salary")
	if errorMessageStore != nil {
		fmt.Printf("Attempt to data retrieve from the database is failed with error : %v", errorMessageRetrieval)
	}

}

