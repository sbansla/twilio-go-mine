/*
 * Twilio - Api
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

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'ListDependentPhoneNumber'
type ListDependentPhoneNumberParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the DependentPhoneNumber resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListDependentPhoneNumberParams) SetPathAccountSid(PathAccountSid string) *ListDependentPhoneNumberParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListDependentPhoneNumberParams) SetPageSize(PageSize int) *ListDependentPhoneNumberParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListDependentPhoneNumberParams) SetLimit(Limit int) *ListDependentPhoneNumberParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of DependentPhoneNumber records from the API. Request is executed immediately.
func (c *ApiService) PageDependentPhoneNumber(AddressSid string, params *ListDependentPhoneNumberParams, pageToken, pageNumber string) (*ListDependentPhoneNumberResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Addresses/{AddressSid}/DependentPhoneNumbers.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"AddressSid"+"}", AddressSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

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

	ps := &ListDependentPhoneNumberResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists DependentPhoneNumber records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListDependentPhoneNumber(AddressSid string, params *ListDependentPhoneNumberParams) ([]ApiV2010DependentPhoneNumber, error) {
	response, errors := c.StreamDependentPhoneNumber(AddressSid, params)

	records := make([]ApiV2010DependentPhoneNumber, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams DependentPhoneNumber records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamDependentPhoneNumber(AddressSid string, params *ListDependentPhoneNumberParams) (chan ApiV2010DependentPhoneNumber, chan error) {
	if params == nil {
		params = &ListDependentPhoneNumberParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ApiV2010DependentPhoneNumber, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageDependentPhoneNumber(AddressSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamDependentPhoneNumber(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamDependentPhoneNumber(response *ListDependentPhoneNumberResponse, params *ListDependentPhoneNumberParams, recordChannel chan ApiV2010DependentPhoneNumber, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.DependentPhoneNumbers
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListDependentPhoneNumberResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListDependentPhoneNumberResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListDependentPhoneNumberResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListDependentPhoneNumberResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
