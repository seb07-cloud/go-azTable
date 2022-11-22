package manipulateAzTable

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

func GetTableData(client *aztables.Client, partitionKey string, rowKey string, tableName string) *string {

	filter := fmt.Sprintf("PartitionKey eq '%s'", partitionKey)
	options := &aztables.ListEntitiesOptions{
		Filter: &filter,
		Top:    to.Ptr(int32(500)),
	}

	pager := client.NewListEntitiesPager(options)
	pageCount := 0

	var export string

	for pager.More() {
		response, err := pager.NextPage(context.TODO())
		if err != nil {
			panic(err)
		}
		pageCount += 1

		for _, entity := range response.Entities {
			var myEntity aztables.EDMEntity
			err = json.Unmarshal(entity, &myEntity)
			if err != nil {
				panic(err)
			}

			if myEntity.RowKey == rowKey {

				jsonStr, err := json.Marshal(myEntity.Properties)
				if err != nil {
					fmt.Printf("Error: %s", err.Error())
				} else {
					fmt.Println(string(jsonStr))
				}

				err = ioutil.WriteFile("data.json", jsonStr, 0644)
				if err != nil {
					log.Fatal(err)
				}
				export = fmt.Sprintln(string(jsonStr))
			}
		}
	}
	return &export
}

func GetSingleTableValue(client *aztables.Client, partitionKey string, rowKey string, tableName string, tableProperty string) *string {

	// type ExportStruct struct {
	// 	Name  string `json:"Key"`
	// 	Value string `json:"Value"`
	// }

	filter := fmt.Sprintf("PartitionKey eq '%s'", partitionKey)
	options := &aztables.ListEntitiesOptions{
		Filter: &filter,
		Top:    to.Ptr(int32(500)),
	}

	pager := client.NewListEntitiesPager(options)
	pageCount := 0

	var export string

	for pager.More() {
		response, err := pager.NextPage(context.TODO())
		if err != nil {
			panic(err)
		}
		pageCount += 1

		for _, entity := range response.Entities {
			var myEntity aztables.EDMEntity
			err = json.Unmarshal(entity, &myEntity)
			if err != nil {
				panic(err)
			}

			if myEntity.RowKey == rowKey {

				for k, v := range myEntity.Properties {
					if k == tableProperty {

						//jsonStr, err := json.Marshal(ExportStruct{k, v.(string)})
						jsonStr, err := json.Marshal(v.(string))

						if err != nil {
							fmt.Printf("Error: %s", err.Error())
						} else {
							fmt.Println(string(jsonStr))
						}

						err = ioutil.WriteFile("data.json", jsonStr, 0644)
						if err != nil {
							log.Fatal(err)
						}
						export = fmt.Sprintln(string(jsonStr))
					}
				}
			}
		}
	}
	return &export
}

func UpdateTableProperties(client *aztables.Client, partitionKey string, rowKey string, tableName string, propertyName string, propertyValue string) (string, error) {

	// Inserting an entity with int64s, binary, datetime, or guid types
	myAddEntity := aztables.EDMEntity{
		Entity: aztables.Entity{
			PartitionKey: partitionKey,
			RowKey:       rowKey,
		},
		Properties: map[string]interface{}{
			propertyName: propertyValue,
		},
	}

	upsertEntityOptions := aztables.UpsertEntityOptions{
		UpdateMode: "merge",
	}

	marshalled, err := json.Marshal(myAddEntity)
	if err != nil {
		return "", errors.New("couldnt convert to json")
	}

	_, err = client.UpsertEntity(context.TODO(), marshalled, &upsertEntityOptions)
	if err != nil {
		return "", errors.New("couldnt update or create value")
	}

	r := fmt.Sprintf("%q : %q", propertyName, propertyValue)
	return r, nil
}

func DeleteTableProperties(client *aztables.Client, partitionKey string, rowKey string, tableName string, propertyName string) (string, error) {

	var s string
	e := azcore.ETag(propertyName)
	etag := azcore.ETagAny.Equals(e)

	if etag {
		_, err := client.DeleteEntity(context.TODO(), partitionKey, rowKey, &aztables.DeleteEntityOptions{IfMatch: &e,})
		fmt.Println(etag)
		if err != nil {
			return "", err
			//"", errors.New("couldnt delete property")
		}
		s = fmt.Sprintf("deleted: %q", propertyName)
	} else {
		fmt.Printf("Value %v not found in Etags", propertyName)
	}
	return s, nil
}


// &aztables.DeleteEntityOptions{IfMatch: &e}