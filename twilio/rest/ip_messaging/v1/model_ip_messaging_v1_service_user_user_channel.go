/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// IpMessagingV1ServiceUserUserChannel struct for IpMessagingV1ServiceUserUserChannel
type IpMessagingV1ServiceUserUserChannel struct {
	AccountSid               *string                 `json:"AccountSid,omitempty"`
	ChannelSid               *string                 `json:"ChannelSid,omitempty"`
	LastConsumedMessageIndex *int32                  `json:"LastConsumedMessageIndex,omitempty"`
	Links                    *map[string]interface{} `json:"Links,omitempty"`
	MemberSid                *string                 `json:"MemberSid,omitempty"`
	ServiceSid               *string                 `json:"ServiceSid,omitempty"`
	Status                   *string                 `json:"Status,omitempty"`
	UnreadMessagesCount      *int32                  `json:"UnreadMessagesCount,omitempty"`
}
