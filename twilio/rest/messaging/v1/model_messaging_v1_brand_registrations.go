/*
 * Twilio - Messaging
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

// MessagingV1BrandRegistrations struct for MessagingV1BrandRegistrations
type MessagingV1BrandRegistrations struct {
	// A2P Messaging Profile Bundle BundleSid
	A2pProfileBundleSid *string `json:"A2pProfileBundleSid,omitempty"`
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// A2P Messaging Profile Bundle BundleSid
	CustomerProfileBundleSid *string `json:"CustomerProfileBundleSid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// A reason why brand registration has failed
	FailureReason *string `json:"FailureReason,omitempty"`
	// A2P BrandRegistration Sid
	Sid *string `json:"Sid,omitempty"`
	// Brand Registration status
	Status *string `json:"Status,omitempty"`
	// Campaign Registry (TCR) Brand ID
	TcrId *string `json:"TcrId,omitempty"`
	// The absolute URL of the Brand Registration
	Url *string `json:"Url,omitempty"`
}
