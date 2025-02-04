# SupersimV1EsimProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account to which the eSIM Profile resource belongs |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**Eid** | Pointer to **string** | Identifier of the eUICC that can claim the eSIM Profile |
**ErrorCode** | Pointer to **string** | Code indicating the failure if the download of the SIM Profile failed and the eSIM Profile is in `failed` state |
**ErrorMessage** | Pointer to **string** | Error message describing the failure if the download of the SIM Profile failed and the eSIM Profile is in `failed` state |
**Iccid** | Pointer to **string** | The ICCID associated with the Sim resource |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**SimSid** | Pointer to **string** | The SID of the Sim resource that this eSIM Profile controls |
**SmdpPlusAddress** | Pointer to **string** | Address of the SM-DP+ server from which the Profile will be downloaded |
**Status** | Pointer to **string** | The status of the eSIM Profile |
**Url** | Pointer to **string** | The absolute URL of the eSIM Profile resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


