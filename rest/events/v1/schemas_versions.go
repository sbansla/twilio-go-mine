/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"

	"github.com/twilio/twilio-go/client"
)

// Fetch a specific schema and version.
func (c *ApiService) FetchSchemaVersion(Id string, SchemaVersion int) (*EventsV1SchemaSchemaVersion, error) {
	path := "/v1/Schemas/{Id}/Versions/{SchemaVersion}"
	path = strings.Replace(path, "{"+"Id"+"}", Id, -1)
	path = strings.Replace(path, "{"+"SchemaVersion"+"}", fmt.Sprint(SchemaVersion), -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1SchemaSchemaVersion{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSchemaVersion'
type ListSchemaVersionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSchemaVersionParams) SetPageSize(PageSize int) *ListSchemaVersionParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSchemaVersionParams) SetLimit(Limit int) *ListSchemaVersionParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of SchemaVersion records from the API. Request is executed immediately.
func (c *ApiService) PageSchemaVersion(Id string, params *ListSchemaVersionParams, pageToken string, pageNumber string) (*ListSchemaVersionResponse, error) {
	path := "/v1/Schemas/{Id}/Versions"

	path = strings.Replace(path, "{"+"Id"+"}", Id, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageToken != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSchemaVersionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SchemaVersion records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSchemaVersion(Id string, params *ListSchemaVersionParams) ([]EventsV1SchemaSchemaVersion, error) {
	if params == nil {
		params = &ListSchemaVersionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSchemaVersion(Id, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []EventsV1SchemaSchemaVersion

	for response != nil {
		records = append(records, response.SchemaVersions...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListSchemaVersionResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListSchemaVersionResponse)
	}

	return records, err
}

// Streams SchemaVersion records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSchemaVersion(Id string, params *ListSchemaVersionParams) (chan EventsV1SchemaSchemaVersion, error) {
	if params == nil {
		params = &ListSchemaVersionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSchemaVersion(Id, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan EventsV1SchemaSchemaVersion, 1)

	go func() {
		for response != nil {
			for item := range response.SchemaVersions {
				channel <- response.SchemaVersions[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListSchemaVersionResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListSchemaVersionResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListSchemaVersionResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSchemaVersionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
