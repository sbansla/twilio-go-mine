/*
 * Twilio - Video
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

// VideoV1Composition struct for VideoV1Composition
type VideoV1Composition struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The array of track names to include in the composition
	AudioSources *[]string `json:"AudioSources,omitempty"`
	// The array of track names to exclude from the composition
	AudioSourcesExcluded *[]string `json:"AudioSourcesExcluded,omitempty"`
	// The average bit rate of the composition's media
	Bitrate *int32 `json:"Bitrate,omitempty"`
	// Date when the media processing task finished
	DateCompleted *time.Time `json:"DateCompleted,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The ISO 8601 date and time in GMT when the composition generated media was deleted
	DateDeleted *time.Time `json:"DateDeleted,omitempty"`
	// The duration of the composition's media file in seconds
	Duration *int32 `json:"Duration,omitempty"`
	// The container format of the composition's media files as specified in the POST request that created the Composition resource
	Format *string `json:"Format,omitempty"`
	// The URL of the media file associated with the composition
	Links *map[string]interface{} `json:"Links,omitempty"`
	// The dimensions of the video image in pixels expressed as columns (width) and rows (height)
	Resolution *string `json:"Resolution,omitempty"`
	// The SID of the Group Room that generated the audio and video tracks used in the composition
	RoomSid *string `json:"RoomSid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// The size of the composed media file in bytes
	Size *int32 `json:"Size,omitempty"`
	// The status of the composition
	Status *string `json:"Status,omitempty"`
	// Whether to remove intervals with no media
	Trim *bool `json:"Trim,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"Url,omitempty"`
	// An object that describes the video layout of the composition
	VideoLayout *map[string]interface{} `json:"VideoLayout,omitempty"`
}
