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

// ChatV2Credential struct for ChatV2Credential
type ChatV2Credential struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// [APN only] Whether to send the credential to sandbox APNs
	Sandbox *string `json:"Sandbox,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// The type of push-notification service the credential is for
	Type *string `json:"Type,omitempty"`
	// The absolute URL of the Credential resource
	Url *string `json:"Url,omitempty"`
}
