package config

import (
	"charlie-parker/pkg/types"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/gommon/log"
)

// Configuration contains relevant app environment variables
type Configuration struct {
	Mode           string `default:"local"`
	AppName        string `default:"charlie-parker"`
	Region         string `default:"localhost"`
	RatesTable     string `default:"parker-rates-local"`
	RatesTableConn dynamo.Table
	WebServerPort  string `default:"8554"`
	APIURL         string `default:"http://localhost:8554"`
	DyDBEndpoint   string `default:"http://localhost:8000"`
}

// Config is the app-wide Configuration
var Config Configuration

func init() {
	err := envconfig.Process("settings", &Config)
	if err != nil {
		panic(err)
	}
	Config.RatesTableConn = connectDynamoDB(Config.RatesTable, types.Rate{})
}

func connectDynamoDB(tableName string, tableDataType interface{}) dynamo.Table {
	// Setup a session to DynamoDB
	dy := dynamo.New(session.New(), &aws.Config{Endpoint: aws.String(Config.DyDBEndpoint), Region: aws.String(Config.Region)})
	// Get all existing tables from DynamoDB.  Panic and exit if
	// unable to communicate with Dynamo and check for tables
	dynamoTables, err := dy.ListTables().All()
	if err != nil {
		log.Error("Unable to check dynamo for existing tables.  Failed with error", err.Error())
		os.Exit(1)
	}
	// Before attempting to perform a create table operation, check
	// and see if our table already exists by doing a string slice
	// comparison
	tableCheck := isStringInSlice(dynamoTables, tableName)
	if tableCheck {
		log.Infof("%v exists in dynamo, create table operation will not be performed", tableName)
		return dy.Table(tableName)
	}
	// If table does not exist, create the table.  Panic and exit the
	// application if it is unable to create a table.
	err = dy.CreateTable(tableName, tableDataType).OnDemand(true).Run()
	if err != nil {
		log.Errorf("Error creating %v table: %v", tableName, err)
		os.Exit(1)
	}
	log.Infof("%v table has been successfully created", tableName)
	return dy.Table(tableName)
}

// If a string (b) is found in our slice of strings (a)
// return true.  Otherwise, return false.
func isStringInSlice(a []string, b string) bool {
	for _, x := range a {
		if x == b {
			return true
		}
	}
	return false
}
