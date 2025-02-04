/*
 * Twilio - Numbers
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

// Fetch a specific Supporting Document Type Instance.
func (c *ApiService) FetchSupportingDocumentType(Sid string) (*NumbersV2SupportingDocumentType, error) {
	path := "/v2/RegulatoryCompliance/SupportingDocumentTypes/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2SupportingDocumentType{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSupportingDocumentType'
type ListSupportingDocumentTypeParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSupportingDocumentTypeParams) SetPageSize(PageSize int) *ListSupportingDocumentTypeParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSupportingDocumentTypeParams) SetLimit(Limit int) *ListSupportingDocumentTypeParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of SupportingDocumentType records from the API. Request is executed immediately.
func (c *ApiService) PageSupportingDocumentType(params *ListSupportingDocumentTypeParams, pageToken, pageNumber string) (*ListSupportingDocumentTypeResponse, error) {
	path := "/v2/RegulatoryCompliance/SupportingDocumentTypes"

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

	ps := &ListSupportingDocumentTypeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SupportingDocumentType records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSupportingDocumentType(params *ListSupportingDocumentTypeParams) ([]NumbersV2SupportingDocumentType, error) {
	response, errors := c.StreamSupportingDocumentType(params)

	records := make([]NumbersV2SupportingDocumentType, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams SupportingDocumentType records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSupportingDocumentType(params *ListSupportingDocumentTypeParams) (chan NumbersV2SupportingDocumentType, chan error) {
	if params == nil {
		params = &ListSupportingDocumentTypeParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan NumbersV2SupportingDocumentType, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageSupportingDocumentType(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamSupportingDocumentType(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamSupportingDocumentType(response *ListSupportingDocumentTypeResponse, params *ListSupportingDocumentTypeParams, recordChannel chan NumbersV2SupportingDocumentType, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.SupportingDocumentTypes
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListSupportingDocumentTypeResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListSupportingDocumentTypeResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListSupportingDocumentTypeResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSupportingDocumentTypeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
