/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/twilio/twilio-go/client"
)

// Retrieve a specific log.
func (c *ApiService) FetchLog(ServiceSid string, EnvironmentSid string, Sid string) (*ServerlessV1Log, error) {
	path := "/v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Logs/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"EnvironmentSid"+"}", EnvironmentSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ServerlessV1Log{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListLog'
type ListLogParams struct {
	// The SID of the function whose invocation produced the Log resources to read.
	FunctionSid *string `json:"FunctionSid,omitempty"`
	// The date/time (in GMT, ISO 8601) after which the Log resources must have been created. Defaults to 1 day prior to current date/time.
	StartDate *time.Time `json:"StartDate,omitempty"`
	// The date/time (in GMT, ISO 8601) before which the Log resources must have been created. Defaults to current date/time.
	EndDate *time.Time `json:"EndDate,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListLogParams) SetFunctionSid(FunctionSid string) *ListLogParams {
	params.FunctionSid = &FunctionSid
	return params
}
func (params *ListLogParams) SetStartDate(StartDate time.Time) *ListLogParams {
	params.StartDate = &StartDate
	return params
}
func (params *ListLogParams) SetEndDate(EndDate time.Time) *ListLogParams {
	params.EndDate = &EndDate
	return params
}
func (params *ListLogParams) SetPageSize(PageSize int) *ListLogParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListLogParams) SetLimit(Limit int) *ListLogParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Log records from the API. Request is executed immediately.
func (c *ApiService) PageLog(ServiceSid string, EnvironmentSid string, params *ListLogParams, pageToken, pageNumber string) (*ListLogResponse, error) {
	path := "/v1/Services/{ServiceSid}/Environments/{EnvironmentSid}/Logs"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"EnvironmentSid"+"}", EnvironmentSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FunctionSid != nil {
		data.Set("FunctionSid", *params.FunctionSid)
	}
	if params != nil && params.StartDate != nil {
		data.Set("StartDate", fmt.Sprint((*params.StartDate).Format(time.RFC3339)))
	}
	if params != nil && params.EndDate != nil {
		data.Set("EndDate", fmt.Sprint((*params.EndDate).Format(time.RFC3339)))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListLogResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Log records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListLog(ServiceSid string, EnvironmentSid string, params *ListLogParams) ([]ServerlessV1Log, error) {
	response, errors := c.StreamLog(ServiceSid, EnvironmentSid, params)

	records := make([]ServerlessV1Log, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Log records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamLog(ServiceSid string, EnvironmentSid string, params *ListLogParams) (chan ServerlessV1Log, chan error) {
	if params == nil {
		params = &ListLogParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ServerlessV1Log, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageLog(ServiceSid, EnvironmentSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamLog(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamLog(response *ListLogResponse, params *ListLogParams, recordChannel chan ServerlessV1Log, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Logs
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListLogResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListLogResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListLogResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListLogResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
