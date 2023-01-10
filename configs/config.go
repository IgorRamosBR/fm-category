package configs

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type AppConfig struct {
	Environment         string
	DynamoCategoryTable string
	DynamoEndpoint      string
	DynamoRegion        string
}

func GetAppConfigs() AppConfig {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		log.Fatal("Failed to read ENVIRONMENT")
	}

	dynamoCategoryTable := os.Getenv("DYNAMODB_CATEGORY_TABLE")
	if dynamoCategoryTable == "" {
		log.Fatal("Failed to read DYNAMODB_CATEGORY_TABLE")
	}

	dynamoEndpoint := os.Getenv("DYNAMODB_CATEGORY_ENDPOINT")
	if dynamoEndpoint == "" {
		log.Fatal("Failed to read DYNAMODB_CATEGORY_ENDPOINT")
	}

	dynamoRegion := os.Getenv("DYNAMODB_CATEGORY_REGION")
	if dynamoRegion == "" {
		log.Fatal("Failed to read DYNAMODB_CATEGORY_REGION")
	}

	return AppConfig{
		Environment:         environment,
		DynamoCategoryTable: dynamoCategoryTable,
		DynamoEndpoint:      dynamoEndpoint,
		DynamoRegion:        dynamoRegion,
	}
}
