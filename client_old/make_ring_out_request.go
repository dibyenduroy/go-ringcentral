/*
 * RingCentral Connect Platform API Explorer
 *
 * <p>This is a beta interactive API explorer for the RingCentral Connect Platform. To use this service, you will need to have an account with the proper credentials to generate an OAuth2 access token.</p><p><h2>Quick Start</h2></p><ol><li>1) Go to <b>Authentication > /oauth/token</b></li><li>2) Enter <b>app_key, app_secret, username, password</b> fields and then click \"Try it out!\"</li><li>3) Upon success, your access_token is loaded and you can access any form requiring authorization.</li></ol><h2>Links</h2><ul><li><a href=\"https://github.com/ringcentral\" target=\"_blank\">RingCentral SDKs on Github</a></li><li><a href=\"mailto:devsupport@ringcentral.com\">RingCentral Developer Support Email</a></li></ul>
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ringcentral

type MakeRingOutRequest struct {

	// Phone number of the caller. This number corresponds to the 1st leg of the RingOut call. This number can be one of user's configured forwarding numbers or arbitrary number
	From *MakeRingOutCallerInfoRequestFrom `json:"from"`

	// Phone number of the called party. This number corresponds to the 2nd leg of the RingOut call
	To *MakeRingOutCallerInfoRequestTo `json:"to"`

	// The number which will be displayed to the called party
	CallerId *MakeRingOutCallerInfoRequestTo `json:"callerId,omitempty"`

	// The audio prompt that the calling party hears when the call is connected
	PlayPrompt bool `json:"playPrompt,omitempty"`

	// Optional. Dialing plan country data. If not specified, then extension home country is applied by default
	Country *MakeRingOutCoutryInfo `json:"country,omitempty"`
}