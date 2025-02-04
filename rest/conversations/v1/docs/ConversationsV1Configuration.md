# ConversationsV1Configuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account responsible for this configuration. |
**DefaultChatServiceSid** | Pointer to **string** | The SID of the default Conversation Service that every new conversation is associated with. |
**DefaultClosedTimer** | Pointer to **string** | Default ISO8601 duration when conversation will be switched to `closed` state. |
**DefaultInactiveTimer** | Pointer to **string** | Default ISO8601 duration when conversation will be switched to `inactive` state. |
**DefaultMessagingServiceSid** | Pointer to **string** | The SID of the default Messaging Service that every new conversation is associated with. |
**Links** | Pointer to **map[string]interface{}** | Absolute URLs to access the webhook and default service configurations. |
**Url** | Pointer to **string** | An absolute URL for this global configuration. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


