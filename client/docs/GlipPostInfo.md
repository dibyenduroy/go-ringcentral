# GlipPostInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of a post | [optional] [default to null]
**GroupId** | **string** | Internal identifier of a group a post belongs to | [optional] [default to null]
**Type_** | **string** | Type of a post | [optional] [default to null]
**Text** | **string** | For &#39;TextMessage&#39; post type only. Message text | [optional] [default to null]
**CreatorId** | **string** | Internal identifier of a user - author of a post | [optional] [default to null]
**AddedPersonIds** | **[]string** | For PersonsAdded post type only. Identifiers of persons added to a group | [optional] [default to null]
**CreationTime** | [**time.Time**](time.Time.md) | Post creation datetime in ISO 8601 format | [optional] [default to null]
**LastModifiedTime** | [**time.Time**](time.Time.md) | Post last modification datetime in ISO 8601 format | [optional] [default to null]
**Attachments** | [**[]GlipMessageAttachmentInfo**](GlipMessageAttachmentInfo.md) | List of posted attachments | [optional] [default to null]
**Mentions** | [**[]GlipMentionsInfo**](GlipMentionsInfo.md) | List of posted attachments | [optional] [default to null]
**Activity** | **string** | activity type | [optional] [default to null]
**Title** | **string** | Title of the message. Can be set for bot messages only | [optional] [default to null]
**IconUri** | **string** | URI to an image to use as the icon for this message. | [optional] [default to null]
**IconEmoji** | **string** | Emoji to use as the icon for a message | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


