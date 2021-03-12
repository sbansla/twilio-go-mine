/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// ChatV2ServiceChannelMessage struct for ChatV2ServiceChannelMessage
type ChatV2ServiceChannelMessage struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The JSON string that stores application-specific data
	Attributes *string `json:"Attributes,omitempty"`
	// The content of the message
	Body *string `json:"Body,omitempty"`
	// The SID of the Channel the Message resource belongs to
	ChannelSid *string `json:"ChannelSid,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The Identity of the message's author
	From *string `json:"From,omitempty"`
	// The index of the message within the Channel
	Index *int32 `json:"Index,omitempty"`
	// The Identity of the User who last updated the Message
	LastUpdatedBy *string `json:"LastUpdatedBy,omitempty"`
	// A Media object that describes the Message's media if attached; otherwise, null
	Media *map[string]interface{} `json:"Media,omitempty"`
	// The SID of the Service that the resource is associated with
	ServiceSid *string `json:"ServiceSid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// The SID of the Channel that the message was sent to
	To *string `json:"To,omitempty"`
	// The Message type
	Type *string `json:"Type,omitempty"`
	// The absolute URL of the Message resource
	Url *string `json:"Url,omitempty"`
	// Whether the message has been edited since  it was created
	WasEdited *bool `json:"WasEdited,omitempty"`
}
