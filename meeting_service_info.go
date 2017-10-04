/* 
 * Ringcentral API
 *
 * RingCentral Connect Platform API
 *
 * OpenAPI spec version: v1.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package ringcentral

type MeetingServiceInfo struct {

	// Canonical URI of a meeting service info resource
	Uri string `json:"uri,omitempty"`

	// URI to retrieve support information for meetings functionality
	SupportUri string `json:"supportUri,omitempty"`

	// URI to retrieve international dial in numbers
	IntlDialInNumbersUri string `json:"intlDialInNumbersUri,omitempty"`

	// External user data
	ExternalUserInfo ExternalUserInfo `json:"externalUserInfo,omitempty"`

	// Dial-in numbers data
	DialInNumbers DialInNumbers `json:"dialInNumbers,omitempty"`
}