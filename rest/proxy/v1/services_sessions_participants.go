/*
 * Twilio - Proxy
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

// Optional parameters for the method 'CreateParticipant'
type CreateParticipantParams struct {
	// [Experimental] For accounts with the ProxyAllowParticipantConflict account flag, setting to true enables per-request opt-in to allowing Proxy to reject a Participant create request that could cause the same Identifier/ProxyIdentifier pair to be active in multiple Sessions. Depending on the context, this could be a 409 error (Twilio error code 80623) or a 400 error (Twilio error code 80604). If not provided, requests will be allowed to succeed and a Debugger notification (80802) will be emitted. Having multiple, active Participants with the same Identifier/ProxyIdentifier pair causes calls and messages from affected Participants to be routed incorrectly. Please note, the default behavior for accounts without the ProxyAllowParticipantConflict flag is to reject the request as described.  This will eventually be the default for all accounts.
	FailOnParticipantConflict *bool `json:"FailOnParticipantConflict,omitempty"`
	// The string that you assigned to describe the participant. This value must be 255 characters or fewer. **This value should not have PII.**
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The phone number of the Participant.
	Identifier *string `json:"Identifier,omitempty"`
	// The proxy phone number to use for the Participant. If not specified, Proxy will select a number from the pool.
	ProxyIdentifier *string `json:"ProxyIdentifier,omitempty"`
	// The SID of the Proxy Identifier to assign to the Participant.
	ProxyIdentifierSid *string `json:"ProxyIdentifierSid,omitempty"`
}

func (params *CreateParticipantParams) SetFailOnParticipantConflict(FailOnParticipantConflict bool) *CreateParticipantParams {
	params.FailOnParticipantConflict = &FailOnParticipantConflict
	return params
}
func (params *CreateParticipantParams) SetFriendlyName(FriendlyName string) *CreateParticipantParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateParticipantParams) SetIdentifier(Identifier string) *CreateParticipantParams {
	params.Identifier = &Identifier
	return params
}
func (params *CreateParticipantParams) SetProxyIdentifier(ProxyIdentifier string) *CreateParticipantParams {
	params.ProxyIdentifier = &ProxyIdentifier
	return params
}
func (params *CreateParticipantParams) SetProxyIdentifierSid(ProxyIdentifierSid string) *CreateParticipantParams {
	params.ProxyIdentifierSid = &ProxyIdentifierSid
	return params
}

// Add a new Participant to the Session
func (c *ApiService) CreateParticipant(ServiceSid string, SessionSid string, params *CreateParticipantParams) (*ProxyV1ServiceSessionParticipant, error) {
	path := "/v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"SessionSid"+"}", SessionSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FailOnParticipantConflict != nil {
		data.Set("FailOnParticipantConflict", fmt.Sprint(*params.FailOnParticipantConflict))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Identifier != nil {
		data.Set("Identifier", *params.Identifier)
	}
	if params != nil && params.ProxyIdentifier != nil {
		data.Set("ProxyIdentifier", *params.ProxyIdentifier)
	}
	if params != nil && params.ProxyIdentifierSid != nil {
		data.Set("ProxyIdentifierSid", *params.ProxyIdentifierSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyV1ServiceSessionParticipant{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Participant. This is a soft-delete. The participant remains associated with the session and cannot be re-added. Participants are only permanently deleted when the [Session](https://www.twilio.com/docs/proxy/api/session) is deleted.
func (c *ApiService) DeleteParticipant(ServiceSid string, SessionSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"SessionSid"+"}", SessionSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch a specific Participant.
func (c *ApiService) FetchParticipant(ServiceSid string, SessionSid string, Sid string) (*ProxyV1ServiceSessionParticipant, error) {
	path := "/v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"SessionSid"+"}", SessionSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyV1ServiceSessionParticipant{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListParticipant'
type ListParticipantParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListParticipantParams) SetPageSize(PageSize int) *ListParticipantParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListParticipantParams) SetLimit(Limit int) *ListParticipantParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Participant records from the API. Request is executed immediately.
func (c *ApiService) PageParticipant(ServiceSid string, SessionSid string, params *ListParticipantParams, pageToken string, pageNumber string) (*ListParticipantResponse, error) {
	path := "/v1/Services/{ServiceSid}/Sessions/{SessionSid}/Participants"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"SessionSid"+"}", SessionSid, -1)

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

	ps := &ListParticipantResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Participant records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListParticipant(ServiceSid string, SessionSid string, params *ListParticipantParams) ([]ProxyV1ServiceSessionParticipant, error) {
	if params == nil {
		params = &ListParticipantParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageParticipant(ServiceSid, SessionSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ProxyV1ServiceSessionParticipant

	for response != nil {
		records = append(records, response.Participants...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListParticipantResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListParticipantResponse)
	}

	return records, err
}

// Streams Participant records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamParticipant(ServiceSid string, SessionSid string, params *ListParticipantParams) (chan ProxyV1ServiceSessionParticipant, error) {
	if params == nil {
		params = &ListParticipantParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageParticipant(ServiceSid, SessionSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ProxyV1ServiceSessionParticipant, 1)

	go func() {
		for response != nil {
			for item := range response.Participants {
				channel <- response.Participants[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListParticipantResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListParticipantResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListParticipantResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListParticipantResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
