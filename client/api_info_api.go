/*
 * RingCentral Connect Platform API Explorer
 *
 * <p>This is a beta interactive API explorer for the RingCentral Connect Platform. To use this service, you will need to have an account with the proper credentials to generate an OAuth2 access token.</p><p><h2>Quick Start</h2></p><ol><li>1) Go to <b>Authentication > /oauth/token</b></li><li>2) Enter <b>app_key, app_secret, username, password</b> fields and then click \"Try it out!\"</li><li>3) Upon success, your access_token is loaded and you can access any form requiring authorization.</li></ol><h2>Links</h2><ul><li><a href=\"https://github.com/ringcentral\" target=\"_blank\">RingCentral SDKs on Github</a></li><li><a href=\"mailto:devsupport@ringcentral.com\">RingCentral Developer Support Email</a></li></ul>
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ringcentral

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type APIInfoApiService service


/* APIInfoApiService Get Version Info
 &lt;p style&#x3D;&#39;font-style:italic;&#39;&gt;Since 1.0.0&lt;/p&gt;&lt;p&gt;Returns current API version info by apiVersion.&lt;/p&gt;&lt;h4&gt;Usage Plan Group&lt;/h4&gt;&lt;p&gt;Light&lt;/p&gt;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param apiVersion API version to be requested, for example &#39;v1.0&#39;
 @return GetVersionResponse*/
func (a *APIInfoApiService) GetVersionInfo(ctx context.Context, apiVersion string) (GetVersionResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  GetVersionResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/restapi/{apiVersion}"
	localVarPath = strings.Replace(localVarPath, "{"+"apiVersion"+"}", fmt.Sprintf("%v", apiVersion), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* APIInfoApiService Get Status
 &lt;p style&#x3D;&#39;font-style:italic;&#39;&gt;Since 1.0.27 (Release 8.3)&lt;/p&gt;&lt;p&gt;Returns the API status; status &#39;200&#39; means the API is working fine, and &#39;503&#39; means it is temporary unavailable.&lt;/p&gt;&lt;h4&gt;API Group&lt;/h4&gt;&lt;p&gt;Light&lt;/p&gt;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @return */
func (a *APIInfoApiService) LoadAPIStatus(ctx context.Context) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/restapi/v1.0/status"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* APIInfoApiService Get API Versions
 &lt;p style&#x3D;&#39;font-style:italic;&#39;&gt;Since 1.0.0&lt;/p&gt;&lt;p&gt;Returns current API version(s) and server info.&lt;/p&gt;&lt;h4&gt;Usage Plan Group&lt;/h4&gt;&lt;p&gt;Light&lt;/p&gt;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @return GetVersionsResponse*/
func (a *APIInfoApiService) RestapiGet(ctx context.Context) (GetVersionsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  GetVersionsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/restapi"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

