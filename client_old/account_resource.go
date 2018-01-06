/*
 * RingCentral Connect Platform API Explorer
 *
 * <p>This is a beta interactive API explorer for the RingCentral Connect Platform. To use this service, you will need to have an account with the proper credentials to generate an OAuth2 access token.</p><p><h2>Quick Start</h2></p><ol><li>1) Go to <b>Authentication > /oauth/token</b></li><li>2) Enter <b>app_key, app_secret, username, password</b> fields and then click \"Try it out!\"</li><li>3) Upon success, your access_token is loaded and you can access any form requiring authorization.</li></ol><h2>Links</h2><ul><li><a href=\"https://github.com/ringcentral\" target=\"_blank\">RingCentral SDKs on Github</a></li><li><a href=\"mailto:devsupport@ringcentral.com\">RingCentral Developer Support Email</a></li></ul>
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ringcentral

type AccountResource struct {
	Uri string `json:"uri,omitempty"`

	Id string `json:"id,omitempty"`

	ServiceInfo *AccountServiceInfoResource `json:"serviceInfo,omitempty"`

	PartnerId string `json:"partnerId,omitempty"`

	Operator *ExtensionResource `json:"operator"`

	MainNumber string `json:"mainNumber,omitempty"`

	ReservationId string `json:"reservationId,omitempty"`

	SessionId string `json:"sessionId,omitempty"`

	Status string `json:"status,omitempty"`

	StatusInfo *StatusInfo `json:"statusInfo,omitempty"`

	SignupInfo *SignupInfoResource `json:"signupInfo,omitempty"`

	SetupWizardState string `json:"setupWizardState,omitempty"`

	TesterFlags []string `json:"testerFlags,omitempty"`

	PromotionCode string `json:"promotionCode,omitempty"`
}
