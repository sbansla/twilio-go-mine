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

// TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics struct for TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics
type TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The average time in seconds between Task creation and acceptance
	AvgTaskAcceptanceTime *int32 `json:"AvgTaskAcceptanceTime,omitempty"`
	// The end of the interval during which these statistics were calculated
	EndTime *time.Time `json:"EndTime,omitempty"`
	// The total number of Reservations accepted for Tasks in the TaskQueue
	ReservationsAccepted *int32 `json:"ReservationsAccepted,omitempty"`
	// The total number of Reservations canceled for Tasks in the TaskQueue
	ReservationsCanceled *int32 `json:"ReservationsCanceled,omitempty"`
	// The total number of Reservations created for Tasks in the TaskQueue
	ReservationsCreated *int32 `json:"ReservationsCreated,omitempty"`
	// The total number of Reservations rejected for Tasks in the TaskQueue
	ReservationsRejected *int32 `json:"ReservationsRejected,omitempty"`
	// The total number of Reservations rescinded
	ReservationsRescinded *int32 `json:"ReservationsRescinded,omitempty"`
	// The total number of Reservations that timed out for Tasks in the TaskQueue
	ReservationsTimedOut *int32 `json:"ReservationsTimedOut,omitempty"`
	// A list of objects that describe the Tasks canceled and reservations accepted above and below the specified thresholds
	SplitByWaitTime *map[string]interface{} `json:"SplitByWaitTime,omitempty"`
	// The beginning of the interval during which these statistics were calculated
	StartTime *time.Time `json:"StartTime,omitempty"`
	// The SID of the TaskQueue from which these statistics were calculated
	TaskQueueSid *string `json:"TaskQueueSid,omitempty"`
	// The total number of Tasks canceled in the TaskQueue
	TasksCanceled *int32 `json:"TasksCanceled,omitempty"`
	// The total number of Tasks completed in the TaskQueue
	TasksCompleted *int32 `json:"TasksCompleted,omitempty"`
	// The total number of Tasks deleted in the TaskQueue
	TasksDeleted *int32 `json:"TasksDeleted,omitempty"`
	// The total number of Tasks entered into the TaskQueue
	TasksEntered *int32 `json:"TasksEntered,omitempty"`
	// The total number of Tasks that were moved from one queue to another
	TasksMoved *int32 `json:"TasksMoved,omitempty"`
	// The absolute URL of the TaskQueue statistics resource
	Url *string `json:"Url,omitempty"`
	// The relative wait duration statistics for Tasks accepted while in the TaskQueue
	WaitDurationInQueueUntilAccepted *map[string]interface{} `json:"WaitDurationInQueueUntilAccepted,omitempty"`
	// The wait duration statistics for Tasks accepted while in the TaskQueue
	WaitDurationUntilAccepted *map[string]interface{} `json:"WaitDurationUntilAccepted,omitempty"`
	// The wait duration statistics for Tasks canceled while in the TaskQueue
	WaitDurationUntilCanceled *map[string]interface{} `json:"WaitDurationUntilCanceled,omitempty"`
	// The SID of the Workspace that contains the TaskQueue
	WorkspaceSid *string `json:"WorkspaceSid,omitempty"`
}
