English | [简体中文](./README_CN.md)

<p align="center">
<a href="https://www.huaweicloud.com/"><img width="450px" height="102px" src="https://console-static.huaweicloud.com/static/authui/20210202115135/public/custom/images/logo-en.svg"></a>
</p>

<h1 align="center">Huawei Cloud Go Software Development Kit (Go SDK)</h1>

The Huawei Cloud Go SDK allows you to easily work with Huawei Cloud services such as Elastic Compute Service (ECS) and
Virtual Private Cloud (VPC) without the need to handle API related tasks.

This document introduces how to obtain and use Huawei Cloud Go SDK.

## Requirements

- To use Huawei Cloud Go SDK, you must have Huawei Cloud account as well as the Access Key (AK) and Secret key (SK) of the Huawei
  Cloud account. You can create an AccessKey in the Huawei Cloud console. For more information,
  see [My Credentials](https://support.huaweicloud.com/en-us/usermanual-ca/en-us_topic_0046606340.html).

- To use Huawei Cloud Go SDK to access the APIs of specific service, please make sure you do have activated the service
  in [Huawei Cloud console](https://console.huaweicloud.com/?locale=en-us) if needed.

- Huawei Cloud Go SDK requires go 1.14 or later, run command `go version` to check the version of Go.

## Install Go SDK

Run the following command to install Huawei Cloud Go SDK:

``` bash
# Install the library of Huawei Cloud Go SDK
go get github.com/huaweicloud/huaweicloud-sdk-go-v3
```

## Code Example

- The following example shows how to query a list of VPCs in a specific region, you need to substitute your
  real `{service} "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/{service}/{version}"`
  for `vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2"` in actual use, and initialize the client
  as `{service}.New{Service}Client`.
- Substitute the values for `{your ak string}` and `{your sk string}`。

**Simplified Demo**

``` go
package main

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2"
	vpcModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/model"
	vpcRegion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/region"
)

func main() {
	// Configure authentication
	auth := basic.NewCredentialsBuilder().
		WithAk("{your ak string}").
		WithSk("{your ak string}").
		Build()

	// Create a service client
	client := vpc.NewVpcClient(
		vpc.VpcClientBuilder().
			WithRegion(vpcRegion.ValueOf("cn-north-4")).
			WithCredential(auth).
			Build())

	// Send the request and get the response
	request := &vpcModel.ListVpcsRequest{}
	response, err := client.ListVpcs(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
```

**Detailed Demo**

```go
package main

import (
	"context"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/httphandler"
	vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2"
	vpcModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/model"
	vpcRegion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/region"
	"net"
	"net/http"
)

func main() {
	// Configure authentication
	auth := basic.NewCredentialsBuilder().
		// Access Key and Secret Access Key used for authentication
		WithAk("{your ak string}").
		WithSk("{your ak string}").
		// If ProjectId is not filled in, the SDK will automatically call the IAM service to query the project id corresponding to the region.
		WithProjectId("{your projectId string}").
		// Configure the SDK built-in IAM service endpoint, default is https://iam.myhuaweicloud.com
		WithIamEndpointOverride("https://iam.cn-north-4.myhuaweicloud.com").
		Build()

	// Use default configuration
	httpConfig := config.DefaultHttpConfig()
	// Configure whether to ignore the SSL certificate verification, default is false
	httpConfig.WithIgnoreSSLVerification(true)
	// Configure timeout as needed, default timeout is 120 seconds
	httpConfig.WithTimeout(120)
	// Configure proxy as needed
	proxy := config.NewProxy().
		WithSchema("http").
		WithHost("proxy.huaweicloud.com").
		WithPort(80).
		WithUsername("testuser").
		WithPassword("password")
	httpConfig.WithProxy(proxy)
	// Configure how to create network connections as needed
	dialContext := func(ctx context.Context, network string, addr string) (net.Conn, error) {
		return net.Dial(network, addr)
	}
	httpConfig.WithDialContext(dialContext)
	// Configure HTTP handler for debugging, do not use in production environment
	requestHandler := func(request http.Request) {
		fmt.Println(request)
	}
	responseHandler := func(response http.Response) {
		fmt.Println(response)
	}
	httpHandler := httphandler.NewHttpHandler().AddRequestHandler(requestHandler).AddResponseHandler(responseHandler)
	httpConfig.WithHttpHandler(httpHandler)

	// Create a service client
	client := vpc.NewVpcClient(
		vpc.VpcClientBuilder().
			// Configure region, it will cause a panic if the region does not exist
			WithRegion(vpcRegion.ValueOf("cn-north-4")).
			// Configure authentication
			WithCredential(auth).
			// Configure HTTP
			WithHttpConfig(httpConfig).
			Build())

	// Create a request
	request := &vpcModel.ListVpcsRequest{}
	// Configure the number of records on each page
	limit := int32(1)
	request.Limit = &limit

	// Send the request and get the response
	response, err := client.ListVpcs(request)
	// Handle error and print response
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
```

## Online Debugging

[API Explorer](https://apiexplorer.developer.intl.huaweicloud.com/apiexplorer/overview) provides api retrieval and online debugging, supports full fast retrieval, visual debugging, help document viewing, and online consultation.

## Changelog

Detailed changes for each released version are documented in
the [CHANGELOG.md](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/blob/master/CHANGELOG.md).

## User Manual [:top:](#huawei-cloud-go-software-development-kit-go-sdk)

* [1. Client Configuration](#1-client-configuration-top)
    * [1.1 Default Configuration](#11-default-configuration-top)
    * [1.2 Network Proxy](#12-network-proxy-top)
    * [1.3 Timeout Configuration](#13-timeout-configuration-top)
    * [1.4 SSL Certification](#14-ssl-certification-top)
    * [1.5 Custom Network Connection](#15-custom-network-connection-top)
* [2. Credentials Configuration](#2-credentials-configuration-top)
  * [2.1 Use Permanent AK&SK](#21-use-permanent-aksk-top)
  * [2.2 Use Temporary AK&SK](#22-use-temporary-aksk-top)
  * [2.3 Use IdpId&IdTokenFile](#23-use-idpididtokenfile-top)
  * [2.4 Authentication Management](#24-authentication-management-top)
    * [2.4.1 Environment Variables](#241-environment-variables-top)
	* [2.4.2 Profile](#242-profile-top)
	* [2.4.3 Metadata](#243-metadata-top)
	* [2.4.4 Provider Chain](#244-provider-chain-top)
* [3. Client Initialization](#3-client-initialization-top)
  * [3.1 Initialize client with specified Endpoint](#31-initialize-the-serviceclient-with-specified-endpoint-top)
  * [3.2 Initialize client with specified Region](#32-initialize-the-serviceclient-with-specified-region-recommended-top)
  * [3.3 Custom Configuration](#33-custom-configuration-top)
    * [3.3.1 IAM endpoint configuration](#331-iam-endpoint-configuration-top)
    * [3.3.2 Region configuration](#332-region-configuration-top)
* [4. Send Request and Handle response](#4-send-requests-and-handle-responses-top)
    * [4.1 Exceptions](#41-exceptions-top)
* [5.Troubleshooting](#5-troubleshooting-top)
    * [5.1 Original HTTP Listener](#51-original-http-listener-top)
* [6. Upload and download files](#6-upload-and-download-files-top)
* [7. Retry For Request](#7-retry-for-request-top)

### 1. Client Configuration [:top:](#user-manual-top)

#### 1.1 Default Configuration [:top:](#user-manual-top)

``` go
import "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"

// Use default configuration
httpConfig := config.DefaultHttpConfig()

client := vpc.NewVpcClient(
    vpc.VpcClientBuilder().
    WithHttpConfig(httpConfig).
    Build())
```

#### 1.2 Network Proxy [:top:](#user-manual-top)

``` go
// Use proxy if needed
proxy := config.NewProxy().
    WithSchema("http").
    WithHost("proxy.huaweicloud.com").
    WithPort(80).
    WithUsername("testuser").
    WithPassword("password")
httpConfig := config.DefaultHttpConfig().WithProxy(proxy)

client := vpc.NewVpcClient(
    vpc.VpcClientBuilder().
    WithHttpConfig(httpConfig).
    Build())
```

#### 1.3 Timeout Configuration [:top:](#user-manual-top)

``` go
// The default timeout is 120 seconds, which can be adjusted as needed
httpConfig := config.DefaultHttpConfig().WithTimeout(120)

client := vpc.NewVpcClient(
    vpc.VpcClientBuilder().
    WithHttpConfig(httpConfig).
    Build())
```

#### 1.4 SSL Certification [:top:](#user-manual-top)

``` go
// Skip SSL certification checking while using https protocol if needed
httpConfig := config.DefaultHttpConfig().WithIgnoreSSLVerification(true)

client := vpc.NewVpcClient(
    vpc.VpcClientBuilder().
    WithHttpConfig(httpConfig).
    Build())
```

#### 1.5 Custom Network Connection [:top:](#user-manual-top)

``` go
// Config network connection dial function if needed
func DialContext(ctx context.Context, network string, addr string) (net.Conn, error) {
	return net.Dial(network, addr)
}
httpConfig := config.DefaultHttpConfig().WithDialContext(DialContext)

client := vpc.NewVpcClient(
    vpc.VpcClientBuilder().
    WithHttpConfig(httpConfig).
    Build())
```

### 2. Credentials Configuration [:top:](#user-manual-top)

There are two types of Huawei Cloud services, `regional` services and `global` services.

Global services contain BSS, DevStar, EPS, IAM, RMS.

For `regional` services' authentication, projectId is required to initialize basic.NewCredentialsBuilder(). 

For `global` services' authentication, domainId is required to initialize global.NewCredentialsBuilder().

The following authentications are supported:

- permanent AK&SK
- temporary AK&SK + SecurityToken
- IdpId&IdTokenFile

#### 2.1 Use Permanent AK&SK [:top:](#user-manual-top)

**Parameter description**:

- `ak` is the access key ID for your account.
- `sk` is the secret access key for your account.
- `projectId` is the ID of your project depending on your region which you want to operate.
- `domainId` is the account ID of Huawei Cloud.

``` go
// Regional Services
basicCredentials := basic.NewCredentialsBuilder().
    WithAk(ak).
    WithSk(sk).
    WithProjectId(projectId).
    Build()

// Global Services
globalCredentials := global.NewCredentialsBuilder().
    WithAk(ak).
    WithSk(sk).
    WithDomainId(domainId).
    Build()
```

**Notice**:

- projectId/domainId supports **automatic acquisition** in version `0.0.26-beta` or later, if you want to use this
  feature, you need to provide the ak and sk of your account and the id of the region, and then build your client
  instance with method `WithRegion()`, detailed example could refer
  to [3.2  Initialize client with specified Region](#32-initialize-the-serviceclient-with-specified-region-recommended-top)
  .

#### 2.2 Use Temporary AK&SK [:top:](#user-manual-top)

It's required to obtain temporary AK&SK and security token first, which could be obtained through
permanent AK&SK or through an agency.A temporary access key and securityToken are issued by the system to IAM users, and can be valid for 15 minutes to 24 hours.

- Obtaining a temporary access key and security token through token, you could refer to
document: https://support.huaweicloud.com/en-us/api-iam/iam_04_0002.html . The API mentioned in the document above
corresponds to the method of `CreateTemporaryAccessKeyByToken` in IAM SDK.

- Obtaining a temporary access key and security token through an agency, you could refer to
document: https://support.huaweicloud.com/en-us/api-iam/iam_04_0101.html . The API mentioned in the document above
corresponds to the method of `CreateTemporaryAccessKeyByAgency` in IAM SDK.

**Parameter description**:

- `ak` is the access key ID for your account.
- `sk` is the secret access key for your account.
- `securityToken` is the security token when using temporary AK/SK.
- `projectId` is the ID of your project depending on your region which you want to operate.
- `domainId` is the account ID of Huawei Cloud.

``` go
// Regional Services
basicCredentials := basic.NewCredentialsBuilder().
            WithAk(ak).
            WithSk(sk).
            WithProjectId(projectId).
            WithSecurityToken(securityToken).
            Build()

// Global Services
globalCredentials := global.NewCredentialsBuilder().
            WithAk(ak).
            WithSk(sk).
            WithDomainId(domainId).
            WithSecurityToken(securityToken).
            Build()
```

In the following two cases, the temporary AK/SK and securitytoken will be obtained from the **metadata of the instance**:

1. basic.Credentials or global.Credentials were not explicitly specified when creating the client.
2. AK/SK was not explicitly specified when creating basic.Credentials or global.Credentials.

Refer to the [Obtaining Metadata](https://support.huaweicloud.com/intl/en-us/usermanual-ecs/ecs_03_0166.html) for more information.

```go
// Regional Services
basicAuth := basic.NewCredentialsBuilder().WithProjectId(projectId).Build()

// Global Services
globalAuth := global.NewCredentialsBuilder().WithDomainId(domainId).Build()
```

#### 2.3 Use IdpId&IdTokenFile [:top:](#user-manual-top)

Obtain a federated identity authentication token using an OpenID Connect ID token, refer to the [Obtaining a Token with an OpenID Connect ID Token](https://support.huaweicloud.com/intl/en-us/api-iam/iam_13_0605.html)

**Parameter description**:

- `idpId` Identity provider ID.
- `idTokenFile` Id token file path. Id token is constructed by the enterprise IdP to carry the identity information of federated users.
- `projectId` is the ID of your project depending on your region which you want to operate.
- `domainId` is the account ID of Huawei Cloud.

```go
import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/global"
)

// Regional service
basicAuth := basic.NewCredentialsBuilder().
	WithIdpId(idpId).
	WithIdTokenFile(idTokenFile).
	WithProjectId(projectId).
	Build()

// Global service
globalAuth := global.NewCredentialsBuilder().
	WithIdpId(idpId).
	WithIdTokenFile(idTokenFile).
	WithDomainId(domainId).
	Build()
```

#### 2.4 Authentication Management [:top:](#user-manual-top)

Getting Authentication from providers is supported since `v0.0.96`

**Regional services** use `BasicCredentialXxxProvider`, **Global services** use `GlobalCredentialXxxProvider`

##### 2.4.1 Environment Variables [:top:](#user-manual-top)

**AK/SK Auth**

| Environment Variables  |  Notice |
| ------------ | ------------ |
| HUAWEICLOUD_SDK_AK  | Required，AccessKey  |
| HUAWEICLOUD_SDK_SK  |  Required，SecretKey |
| HUAWEICLOUD_SDK_SECURITY_TOKEN  | Optional, this parameter needs to be specified when using temporary ak/sk  |
| HUAWEICLOUD_SDK_PROJECT_ID  | Optional, used for regional services, required in multi-ProjectId scenarios  |
| HUAWEICLOUD_SDK_DOMAIN_ID  | Optional, used for global services  |

Configure environment variables:

```
// Linux
export HUAWEICLOUD_SDK_AK=YOUR_AK
export HUAWEICLOUD_SDK_SK=YOUR_SK

// Windows
set HUAWEICLOUD_SDK_AK=YOUR_AK
set HUAWEICLOUD_SDK_SK=YOUR_SK
```

Get the credentials from configured environment variables:

```go
import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/provider"
)

// basic
basicProvider := provider.BasicCredentialEnvProvider()
basicCred, err := basicProvider.GetCredentials()

// global
globalProvider := provider.GlobalCredentialEnvProvider()
globalCred, err := globalProvider.GetCredentials()
```

**IdpId/IdTokenFile Auth**

| Environment Variables  |  Notice |
| ------------ | ------------ |
| HUAWEICLOUD_SDK_IDP_ID  | Required, identity provider Id |
| HUAWEICLOUD_SDK_ID_TOKEN_FILE  |  Required, id token file path |
| HUAWEICLOUD_SDK_PROJECT_ID  | For basic credentials, this parameter is required  |
| HUAWEICLOUD_SDK_DOMAIN_ID  | For global credentials, this parameter is required  |

Configure environment variables:

```
// Linux
export HUAWEICLOUD_SDK_IDP_ID=YOUR_IDP_ID
export HUAWEICLOUD_SDK_ID_TOKEN_FILE=/some_path/your_token_file
export HUAWEICLOUD_SDK_PROJECT_ID=YOUR_PROJECT_ID // For basic credentials, this parameter is required
export HUAWEICLOUD_SDK_DOMAIN_ID=YOUR_DOMAIN_ID // For global credentials, this parameter is required

// Windows
set HUAWEICLOUD_SDK_IDP_ID=YOUR_IDP_ID
set HUAWEICLOUD_SDK_ID_TOKEN_FILE=/some_path/your_token_file
set HUAWEICLOUD_SDK_PROJECT_ID=YOUR_PROJECT_ID // For basic credentials, this parameter is required
set HUAWEICLOUD_SDK_DOMAIN_ID=YOUR_DOMAIN_ID // For global credentials, this parameter is required
```

Get the credentials from configured environment variables:

```go
import (
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/provider"
)

// basic
basicProvider := provider.BasicCredentialEnvProvider()
basicCred, err := basicProvider.GetCredentials()

// global
globalProvider := provider.GlobalCredentialEnvProvider()
globalCred, err := globalProvider.GetCredentials()
```

##### 2.4.2 Profile [:top:](#user-manual-top)

The profile will be read from the user's home directory by default, linux`~/.huaweicloud/credentials`,windows`C:\Users\USER_NAME\.huaweicloud\credentials`, the path to the profile can be modified by configuring the environment variable `HUAWEICLOUD_SDK_CREDENTIALS_FILE`

**AK/SK Auth**

| Configuration Parameters  |  Notice |
| ------------ | ------------ |
| ak  | Required，AccessKey  |
| sk  |  Required，SecretKey |
| security_token  | Optional, this parameter needs to be specified when using temporary ak/sk  |
| project_id  | Optional, used for regional services, required in multi-ProjectId scenarios  |
| domain_id  | Optional, used for global services  |
| iam_endpoint  | optional, endpoint for authentication, default is `https://iam.myhuaweicloud.com` |

The content of the profile is as follows:

```ini
[basic]
ak = your_ak
sk = your_sk

[global]
ak = your_ak
sk = your_sk
```

Get the credentials from profile:

```go
import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/provider"
)

// basic
basicProvider := provider.BasicCredentialProfileProvider()
basicCred, err := basicProvider.GetCredentials()

// global
globalProvider := provider.GlobalCredentialProfileProvider()
globalCred, err := globalProvider.GetCredentials()
```

**IdpId/IdTokenFile Auth**

| Configuration Parameters  |  Notice |
| ------------ | ------------ |
| idp_id  | Required, identity provider Id |
| id_token_file  |  Required, id token file path |
| project_id  | For basic credentials, this parameter is required  |
| domain_id  | For global credentials, this parameter is required  |
| iam_endpoint  | optional, endpoint for authentication, default is `https://iam.myhuaweicloud.com` |

The content of the profile is as follows:

```ini
[basic]
idp_id = your_idp_id
id_token_file = /some_path/your_token_file
project_id = your_project_id

[global]
idp_id = your_idp_id
id_token_file = /some_path/your_token_file
domainId = your_domain_id
```

Get the credentials from profile:

```go
import (
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/provider"
)

// basic
basicProvider := provider.BasicCredentialProfileProvider()
basicCred, err := basicProvider.GetCredentials()

// global
globalProvider := provider.GlobalCredentialProfileProvider()
globalCred, err := globalProvider.GetCredentials()
```

##### 2.4.3 Metadata [:top:](#user-manual-top)

Get temporary AK/SK and securitytoken from instance's metadata. Refer to the [Obtaining Metadata](https://support.huaweicloud.com/intl/en-us/usermanual-ecs/ecs_03_0166.html) for more information.

Manually obtain authentication from instance metadata:

```go
import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/provider"
)

// basic
basicProvider := provider.BasicCredentialMetadataProvider()
basicCred, err := basicProvider.GetCredentials()

// global
globalProvider := provider.GlobalCredentialMetadataProvider()
globalCred, err := globalProvider.GetCredentials()
```

##### 2.4.4 Provider Chain [:top:](#user-manual-top)

When creating a service client without credentials, try to load authentication in the order **Environment Variables -> Profile -> Metadata**

Get authentication from provider chain:

```go
import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/provider"
)

// basic
basicChain := provider.BasicCredentialProviderChain()
basicCred, err := basicChain.GetCredentials()

// global
globalChain := provider.GlobalCredentialProviderChain()
globalCred, err := globalChain.GetCredentials()
```

Custom credentials provider chain is supported:

```go
import (
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/provider"
)

providers := []provider.ICredentialProvider{
    provider.BasicCredentialMetadataProvider(),
    provider.BasicCredentialProfileProvider(),
}
chain := provider.NewCredentialProviderChain(providers)
cred, err := chain.GetCredentials()
```

### 3. Client Initialization [:top:](#user-manual-top)

There are two ways to initialize the {Service}Client, you could choose one you preferred.

#### 3.1 Initialize the {Service}Client with specified Endpoint [:top:](#user-manual-top)

``` go
// Specify the endpoint, take the endpoint of VPC service in region of cn-north-4 for example
endpoint := "https://vpc.cn-north-4.myhuaweicloud.com"

// Initialize the credentials, you should provide projectId or domainId in this way, take initializing BasicCredentials for example
basicAuth := basic.NewCredentialsBuilder().
    WithAk(ak).
    WithSk(sk).
    WithProjectId(projectId).
    Build()

// Initialize specified New{Service}Client, take initializing the regional service VPC's VpcClient for example
client := vpc.NewVpcClient(
    vpc.VpcClientBuilder().
        WithEndpoint(endpoint).
        WithCredential(basicCredentials).
        WithHttpConfig(config.DefaultHttpConfig()).  
        Build())
```

**where:**

- `endpoint` varies with services and regions,
  see [Regions and Endpoints](https://developer.huaweicloud.com/intl/en-us/endpoint) to obtain correct endpoint.

- When you meet some trouble in getting projectId using the specified region way, you could use this way instead.

#### 3.2 Initialize the {Service}Client with specified Region **(Recommended)** [:top:](#user-manual-top)

``` go
import (
    // dependency for region module
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/region"
)

// Initialize the credentials, projectId or domainId could be unassigned in this situation, take initializing GlobalCredentials for example
globalCredentials := global.NewCredentialsBuilder().
    WithAk(ak).
    WithSk(sk).
    // domainId could be unassigned in this situation
    Build()

// Initialize specified New{Service}Client, take initializing the global service IAM's NewIamClient for example
client := iam.NewIamClient(
    iam.IamClientBuilder().
        WithRegion(region.CN_NORTH_4).
        WithCredential(globalCredentials).
        WithHttpConfig(config.DefaultHttpConfig()).  
        Build())
```

**Notice:**

- If you use `region` to initialize {Service}Client, projectId/domainId supports automatic acquisition, you don't need
  to configure it when initializing Credentials.

- Multiple ProjectId situation is **not supported**.

- Supported region list: af-south-1, ap-southeast-1, ap-southeast-2, ap-southeast-3, cn-east-2, cn-east-3, cn-north-1,
  cn-north-4, cn-south-1, cn-southwest-2, ru-northwest-2. You may get exception such as `Unsupported regionId` if your
  region don't in the list above.

**Comparison of the two ways:**

| Initialization | Advantages | Disadvantage |
| :---- | :---- | :---- |
| Specified Endpoint | The API can be invoked successfully once it has been published in the environment. | You need to prepare projectId and endpoint yourself.
| Specified Region | No need for projectId and endpoint, it supports automatic acquisition if you configure it in the right way. | The supported services and regions are limited.

#### 3.3 Custom Configuration

**Notice:** Supported since v0.0.92

##### 3.3.1 IAM endpoint configuration

Automatically acquiring projectId/domainId will invoke the [KeystoneListProjects](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=IAM&api=KeystoneListProjects) /[KeystoneListAuthDomains](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=IAM&api=KeystoneListAuthDomains) interface of IAM service. The default iam endpoint is `https://iam.myhuaweicloud.com`, you can modify the endpoint in the following two ways:

###### 3.3.1.1 Global scope

This configuration takes effect globally, specified by environment variable `HUAWEICLOUD_SDK_IAM_ENDPOINT`

```
//linux
export HUAWEICLOUD_SDK_IAM_ENDPOINT=https://iam.cn-north-4.myhuaweicloud.com

//windows
set HUAWEICLOUD_SDK_IAM_ENDPOINT=https://iam.cn-north-4.myhuaweicloud.com
```

###### 3.3.1.2 Credentials scope

This configuration is only valid for a credential, and it will override the global configuration

```go
import "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"

iamEndpoint := "https://iam.cn-north-4.myhuaweicloud.com"
cred := basic.NewCredentialsBuilder().
			WithAk(ak).
			WithSk(sk).
			WithIamEndpointOverride(iamEndpoint).
			Build()
```

##### 3.3.2 Region configuration

###### 3.3.2.1 Environment variable

Specified by environment variable, the format is `HUAWEICLOUD_SDK_REGION_{SERIVCE_NAME}_{REGION_ID}={endpoint}`

Notice: the name of environment variable is UPPER-CASE, replacing hyphens with underscores.

```
// Take ECS and IoTDA services as examples

// linux
export HUAWEICLOUD_SDK_REGION_ECS_CN_NORTH_99=https://ecs.cn-north-99.myhuaweicloud.com
export HUAWEICLOUD_SDK_REGION_IOTDA_AP_SOUTHEAST_10=https://iotda.ap-southwest-10.myhuaweicloud.com

// windows
set HUAWEICLOUD_SDK_REGION_ECS_CN_NORTH_99=https://ecs.cn-north-99.myhuaweicloud.com
set HUAWEICLOUD_SDK_REGION_IOTDA_AP_SOUTHEAST_10=https://iotda.ap-southwest-10.myhuaweicloud.com
```

###### 3.3.2.2 Profile

The profile will be read from the user's home directory by default, linux`~/.huaweicloud/regions.yaml`, windows`C:\Users\USER_NAME\.huaweicloud\regions.yaml`, the default file may not exist, but if the file exists and the content format is incorrect, an exception will be thrown for parsing errors.

The path to the profile can be modified by configuring the environment variable `HUAWEICLOUD_SDK_REGIONS_FILE`, like `HUAWEICLOUD_SDK_REGIONS_FILE=/tmp/my_regions.yml`

The file content format is as follows:

```yaml
# Serivce name is case-insensitive
ECS:
  - id: 'cn-north-10'
    endpoint: 'https://ecs.cn-north-10.myhuaweicloud.com'
  - id: 'cn-north-11'
    endpoint: 'https://ecs.cn-north-11.myhuaweicloud.com'
IoTDA:
  - id: 'ap-southwest-9'
    endpoint: 'https://iotda.ap-southwest-9.myhuaweicloud.com'
```

###### 3.3.2.3 Region supply chain

The default order is **environment variables -> profile -> region defined in SDK**, if the region is not found in the above ways, an exception will be thrown.

```go
import "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"

var (
	region1 = region.ValueOf("cn-north-10")
	region2 = region.ValueOf("cn-north-11")
)
```

### 4. Send Requests and Handle Responses [:top:](#user-manual-top)

``` go
// send a request and print response, take interface of ListVpcs for example
limit := int32(1)
request := &model.ListVpcsRequest{
    Limit: &limit,
}

response, err := client.ListVpcs(request)
if err == nil {
    fmt.Printf("%+v\n\n", response.Vpcs)
} else {
    fmt.Println(err)
}
```

#### 4.1 Exceptions [:top:](#user-manual-top)

| Level 1 | Notice |
| :---- | :---- |
| ServiceResponseError | service response error |
| url.Error | connect endpoint error |

``` go
response, err := client.ListVpcs(request)
if err == nil {
    fmt.Printf("%+v\n\n", response.Vpcs)
} else {
    fmt.Println(err)
}
```

### 5. Troubleshooting [:top:](#user-manual-top)

#### 5.1 Original HTTP Listener [:top:](#user-manual-top)

In some situation, you may need to debug your http requests, original http request and response information will be
needed. The SDK provides a listener function to obtain the original encrypted http request and response information.

> :warning:  Warning: The original http log information is used in debugging stage only, please do not print the original http header or body in the production environment. This log information is not encrypted and contains sensitive data such as the password of your ECS virtual machine, or the password of your IAM user account, etc. When the response body is binary content, the body will be printed as "***" without detailed information.

``` go
import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/httphandler"
	vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2"
	"net/http"
)

func RequestHandler(request http.Request) {
	fmt.Println(request)
}

func ResponseHandler(response http.Response) {
	fmt.Println(response)
}

handler := httphandler.NewHttpHandler().
	AddRequestHandler(RequestHandler).
	AddResponseHandler(ResponseHandler)
httpConfig := config.DefaultHttpConfig().WithHttpHandler(handler)

client := vpc.NewVpcClient(
	vpc.VpcClientBuilder().
		WithHttpConfig(httpConfig).
		Build())
```

### 6. Upload and download files [:top:](#user-manual-top)

Take the interface `CreateImageWatermark` of the service `Data Security Center` as an example, this interface needs to upload an image file and return the watermarked image file stream:

```go
package main

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	dsc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dsc/v1"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dsc/v1/model"
	"os"
)

func createImageWatermark(client *dsc.DscClient) {
	
	// Open the file.
	file, err := os.Open("demo.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	body := &model.CreateImageWatermarkRequestBody{
		File:           def.NewFilePart(file),
		BlindWatermark: def.NewMultiPart("test_watermark"),
	}

	request := &model.CreateImageWatermarkRequest{Body: body}
	response, err := client.CreateImageWatermark(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
		return
	}

	// Download the file.
	result, err := os.Create("result.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	response.Consume(result)

}

func main() {
	ak := "{your ak string}"
	sk := "{your sk string}"
	endpoint := "{your endpoint string}"
	projectId := "{your project id}"

	credentials := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		WithProjectId(projectId).
		Build()

	client := dsc.NewDscClient(
		dsc.DscClientBuilder().
			WithEndpoint(endpoint).
			WithCredential(credentials).
			Build())

	createImageWatermark(client)
}
```

### 7. Retry For Request [:top:](#user-manual-top)

When a request encounters a network exception or flow control on the interface, the request needs to be retried. The
Go SDK provides the retry method for our users which could be used to the requests of `GET` HTTP method. 
If you want to use the retry method, the following parameters are required:

- _maxRetryTimes_: the max retry times
- _retryCondition_: a function, which determine the condition of when to retry
- _backoffStrategy_: calculate the wait duration before next retry

Take the interface `ListVpcs` of VPC service for example, assume the request would retry at most 3 times, 
retry when service responses an error, the code would be like the following:

``` go
// initialize the client
client := vpc.NewVpcClient(
	vpc.VpcClientBuilder().
		WithEndpoint("<input your endpoint>").
		WithCredential(
			basic.NewCredentialsBuilder().
				WithAk("<input your ak>").
				WithSk("<input your sk>").
				WithProjectId("<input your project id>").
				Build()).
		Build())

// initialize the request
request := &model.ListVpcsRequest{}

// send the requet and retry when service responses an error
response, err := client.ListVpcsInvoker(request).WithRetry(3, func(i interface{}, err error) bool {
	return err != nil
}, new(retry.None)).Invoke()

if err == nil {
	fmt.Printf("%+v\n", response)
} else {
	fmt.Printf("%+v\n", err)
}
```
