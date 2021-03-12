/*
 * Twilio - Taskrouter
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

// TaskrouterV1WorkspaceWorkerWorkerChannel struct for TaskrouterV1WorkspaceWorkerWorkerChannel
type TaskrouterV1WorkspaceWorkerWorkerChannel struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The total number of Tasks assigned to Worker for the TaskChannel type
	AssignedTasks *int32 `json:"AssignedTasks,omitempty"`
	// Whether the Worker should receive Tasks of the TaskChannel type
	Available *bool `json:"Available,omitempty"`
	// The current available capacity between 0 to 100 for the TaskChannel
	AvailableCapacityPercentage *int32 `json:"AvailableCapacityPercentage,omitempty"`
	// The current configured capacity for the WorkerChannel
	ConfiguredCapacity *int32 `json:"ConfiguredCapacity,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// The SID of the TaskChannel
	TaskChannelSid *string `json:"TaskChannelSid,omitempty"`
	// The unique name of the TaskChannel, such as 'voice' or 'sms'
	TaskChannelUniqueName *string `json:"TaskChannelUniqueName,omitempty"`
	// The absolute URL of the WorkerChannel resource
	Url *string `json:"Url,omitempty"`
	// The SID of the Worker that contains the WorkerChannel
	WorkerSid *string `json:"WorkerSid,omitempty"`
	// The SID of the Workspace that contains the WorkerChannel
	WorkspaceSid *string `json:"WorkspaceSid,omitempty"`
}
