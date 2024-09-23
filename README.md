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

You can get the SDK version information through [SDK center](https://console-intl.huaweicloud.com/apiexplorer/#/sdkcenter?language=Go) or [Github Releases](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/releases?page=1).

## Code Example

- The following example shows how to query a list of VPCs in a specific region, you need to substitute your
  real `{service} "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/{service}/{version}"`
  for `vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2"` in actual use, and initialize the client
  as `{service}.New{Service}Client`.
- Hard-coding ak and sk for authentication into the code has a great security risk. It is recommended to store the ciphertext in the profile or environment variables and decrypt it when used to ensure security.
- In this example, ak and sk are stored in environment variables. Please configure the environment variables `HUAWEICLOUD_SDK_AK` and `HUAWEICLOUD_SDK_SK` before running this example.

**Simplified Demo**

``` go
package main

import (
    "os"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2"
	vpcModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/model"
	vpcRegion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/region"
)

func main() {
    // Configure authentication
    // Authentication can be configured through environment variables and other methods. Please refer to Chapter 2.4 Authentication Management
    auth, err := basic.NewCredentialsBuilder().
        WithAk(os.Getenv("HUAWEICLOUD_SDK_AK")).
        WithSk(os.Getenv("HUAWEICLOUD_SDK_SK")).
        SafeBuild()
    if err != nil {
        fmt.Println(err)
        return
    }

    // Get available region
    region, err := vpcRegion.SafeValueOf("cn-north-4")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Create a service client
    hcClient, err := vpc.VpcClientBuilder().
        WithRegion(region).
        WithCredential(auth).
        SafeBuild()
    if err != nil {
        fmt.Println(err)
        return
    }
    client := vpc.NewVpcClient(hcClient)

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
	"os"
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
    auth, err := basic.NewCredentialsBuilder().
        // Authentication can be configured through environment variables and other methods. Please refer to Chapter 2.4 Authentication Management
        WithAk(os.Getenv("HUAWEICLOUD_SDK_AK")).
        WithSk(os.Getenv("HUAWEICLOUD_SDK_SK")).
        // If ProjectId is not filled in, the SDK will automatically call the IAM service to query the project id corresponding to the region.
        WithProjectId("{your projectId string}").
        // Configure the SDK built-in IAM service endpoint, default is https://iam.myhuaweicloud.com
        WithIamEndpointOverride("https://iam.cn-north-4.myhuaweicloud.com").
        SafeBuild()
    if err != nil {
        fmt.Println(err)
        return
    }

    // Use default configuration
    httpConfig := config.DefaultHttpConfig()
    // Configure whether to ignore the SSL certificate verification, default is false
    httpConfig.WithIgnoreSSLVerification(true)
    // Configure timeout as needed, default timeout is 120 seconds
    httpConfig.WithTimeout(120)
    // Configure proxy as needed
    proxy := config.NewProxy().
        // Replace the proxy schema, host and port in the example according to the actual situation
        WithSchema("http").
        WithHost("proxy.huaweicloud.com").
        WithPort(80).
        // Configure the username and password if the proxy requires authentication
        WithUsername(os.Getenv("PROXY_USERNAME")).
        WithPassword(os.Getenv("PROXY_PASSWORD"))
    httpConfig.WithProxy(proxy)
    // Configure how to create network connections as needed
    dialContext := func(ctx context.Context, network string, addr string) (net.Conn, error) {
        // You need to implement this function
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

    // Get available region
    region, err := vpcRegion.SafeValueOf("cn-north-4")
    if err != nil {
        fmt.Println(err)
        return
    }
    // Create a service client
    hcClient, err := vpc.VpcClientBuilder().
        // Configure region, it will cause a panic if the region does not exist
        WithRegion(region).
        // Configure authentication
        WithCredential(auth).
        // Configure HTTP
        WithHttpConfig(httpConfig).
        SafeBuild()
    if err != nil {
        fmt.Println(err)
        return
    }
    client := vpc.NewVpcClient(hcClient)

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

[API Explorer](https://apiexplorer.developer.intl.huaweicloud.com/apiexplorer/overview) provides api retrieval, SDK samples and online debugging, supports full fast retrieval, visual debugging, help document viewing, and online consultation.

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
    * [1.6 Custom HTTP Transport](#16-custom-http-transport-top)
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
* [5. Troubleshooting](#5-troubleshooting-top)
    * [5.1 Original HTTP Listener](#51-original-http-listener-top)
* [6. Upload and download files](#6-upload-and-download-files-top)
* [7. API Invoker](#7-api-invoker-top)
    * [7.1 Custom request headers](#71-custom-request-headers-top)
    * [7.2 Retry for request](#72-retry-for-request-top)

### 1. Client Configuration [:top:](#user-manual-top)

#### 1.1 Default Configuration [:top:](#user-manual-top)

``` go
import (
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
    vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2"
)

// Use default configuration
httpConfig := config.DefaultHttpConfig()

hcClient, err := vpc.VpcClientBuilder().
    WithHttpConfig(httpConfig).
    SafeBuild()
if err != nil {
    // handle error
}
client := vpc.NewVpcClient(hcClient)
```

#### 1.2 Network Proxy [:top:](#user-manual-top)

``` go
// Use proxy if needed
proxy := config.NewProxy().
    // Replace the proxy schema, host and port in the example according to the actual situation
    WithSchema("http").
    WithHost("proxy.huaweicloud.com").
    WithPort(80).
    // Configure the username and password if the proxy requires authentication
    // In this example, username and password are stored in environment variables. Please configure the environment variables PROXY_USERNAME and PROXY_PASSWORD before running this example.
    WithUsername(os.Getenv("PROXY_USERNAME")).
    WithPassword(os.Getenv("PROXY_PASSWORD"))
httpConfig := config.DefaultHttpConfig().WithProxy(proxy)

hcClient, err := vpc.VpcClientBuilder().
    WithHttpConfig(httpConfig).
    SafeBuild()
if err != nil {
    // handle error
}
client := vpc.NewVpcClient(hcClient)
```

#### 1.3 Timeout Configuration [:top:](#user-manual-top)

``` go
// The default timeout is 120 seconds, which can be adjusted as needed
httpConfig := config.DefaultHttpConfig().WithTimeout(120)

hcClient, err := vpc.VpcClientBuilder().
    WithHttpConfig(httpConfig).
    SafeBuild()
if err != nil {
    // handle error
}
client := vpc.NewVpcClient(hcClient)
```

#### 1.4 SSL Certification [:top:](#user-manual-top)

``` go
// Skip SSL certification checking while using https protocol if needed
httpConfig := config.DefaultHttpConfig().WithIgnoreSSLVerification(true)

hcClient, err := vpc.VpcClientBuilder().
    WithHttpConfig(httpConfig).
    SafeBuild()
if err != nil {
    // handle error
}
client := vpc.NewVpcClient(hcClient)
```

#### 1.5 Custom Network Connection [:top:](#user-manual-top)

``` go
// Config network connection dial function if needed
func DialContext(ctx context.Context, network string, addr string) (net.Conn, error) {
    return net.Dial(network, addr)
}
httpConfig := config.DefaultHttpConfig().WithDialContext(DialContext)

hcClient, err := vpc.VpcClientBuilder().
    WithHttpConfig(httpConfig).
    SafeBuild()
if err != nil {
    // handle error
}
client := vpc.NewVpcClient(hcClient)
```

#### 1.6 Custom HTTP Transport [:top:](#user-manual-top)

Supports configuring **HttpTransport** or **HttpRoundTripper** (v0.1.114 or above). The former is an interface implementation of the latter. Just choose one to configure.

**NOTE:** HttpTransport has the highest priority.

Specifying the custom HTTP transport or roundTripper  **will invalidate the configurations [1.2 Network Proxy](#12-network-proxy-top), [1.4 SSL Certification](#14-ssl-certification-top), [1.5 Custom Network Connection](#15-custom-network-connection-top).**

``` go
transport := &http.Transport{}
httpConfig := config.DefaultHttpConfig().WithHttpTransport(transport)
// httpConfig.WithHttpRoundTripper(&YourRoundTripper{})

hcClient, err := vpc.VpcClientBuilder().
    WithHttpConfig(httpConfig).
    SafeBuild()
if err != nil {
    // handle error
}
client := vpc.NewVpcClient(hcClient)
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
ak := os.Getenv("HUAWEICLOUD_SDK_AK")
sk := os.Getenv("HUAWEICLOUD_SDK_SK")
projectId := "{your projectId string}"

basicAuth, err := basic.NewCredentialsBuilder().
    WithAk(ak).
    WithSk(sk).
    WithProjectId(projectId).
    SafeBuild()

// Global Services
ak := os.Getenv("HUAWEICLOUD_SDK_AK")
sk := os.Getenv("HUAWEICLOUD_SDK_SK")
domainId := "{your domainId string}"

globalAuth, err := global.NewCredentialsBuilder().
    WithAk(ak).
    WithSk(sk).
    WithDomainId(domainId).
    SafeBuild()
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
ak := os.Getenv("HUAWEICLOUD_SDK_AK")
sk := os.Getenv("HUAWEICLOUD_SDK_SK")
securityToken := os.Getenv("HUAWEICLOUD_SDK_SECURITY_TOKEN")
projectId := "{your projectId string}"

basicAuth, err := basic.NewCredentialsBuilder().
    WithAk(ak).
    WithSk(sk).
    WithSecurityToken(securityToken).
    WithProjectId(projectId).
    SafeBuild()

// Global Services
ak := os.Getenv("HUAWEICLOUD_SDK_AK")
sk := os.Getenv("HUAWEICLOUD_SDK_SK")
securityToken := os.Getenv("HUAWEICLOUD_SDK_SECURITY_TOKEN")
domainId := "{your domainId string}"

globalAuth, err := global.NewCredentialsBuilder().
    WithAk(ak).
    WithSk(sk).
    WithSecurityToken(securityToken).
    WithDomainId(domainId).
    SafeBuild()
```

In the following two cases, the temporary AK/SK and securitytoken will be obtained from the **metadata of the instance**:

1. basic.Credentials or global.Credentials were not explicitly specified when creating the client.
2. AK/SK was not explicitly specified when creating basic.Credentials or global.Credentials.

Refer to the [Obtaining Metadata](https://support.huaweicloud.com/intl/en-us/usermanual-ecs/ecs_03_0166.html) for more information.

```go
// Regional Services
basicAuth, err := basic.NewCredentialsBuilder().WithProjectId(projectId).SafeBuild()

// Global Services
globalAuth, err := global.NewCredentialsBuilder().WithDomainId(domainId).SafeBuild()
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
basicAuth, err := basic.NewCredentialsBuilder().
    WithIdpId(idpId).
    WithIdTokenFile(idTokenFile).
    WithProjectId(projectId).
    SafeBuild()

// Global service
globalAuth, err := global.NewCredentialsBuilder().
    WithIdpId(idpId).
    WithIdTokenFile(idTokenFile).
    WithDomainId(domainId).
    SafeBuild()
```

#### 2.4 Authentication Management [:top:](#user-manual-top)

Getting Authentication from providers is supported since `v0.0.96`

**Regional services** use `BasicCredentialXxxProvider`, **Global services** use `GlobalCredentialXxxProvider`

##### 2.4.1 Environment Variables [:top:](#user-manual-top)

**AK/SK Auth**

| Environment Variables  |  Notice |
| ------------ | ------------ |
| HUAWEICLOUD_SDK_AK  | Required, AccessKey  |
| HUAWEICLOUD_SDK_SK  |  Required, SecretKey |
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
import "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/provider"

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
import "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/provider"

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
| ak  | Required, AccessKey  |
| sk  |  Required, SecretKey |
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
import "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/provider"

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
import "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/provider"

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
import "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/provider"

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
import "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/provider"

// basic
basicChain := provider.BasicCredentialProviderChain()
basicCred, err := basicChain.GetCredentials()

// global
globalChain := provider.GlobalCredentialProviderChain()
globalCred, err := globalChain.GetCredentials()
```

Custom credentials provider chain is supported:

```go
import "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/provider"

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
package main

import (
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
    vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2"
    "os"
)

func main() {
    // Specify the endpoint, take the endpoint of VPC service in region of cn-north-4 for example
    endpoint := "https://vpc.cn-north-4.myhuaweicloud.com"
    // Initialize the credentials, you should provide projectId or domainId in this way, take initializing BasicCredentials for example
    basicAuth, err := basic.NewCredentialsBuilder().
        WithAk(os.Getenv("HUAWEICLOUD_SDK_AK")).
        WithSk(os.Getenv("HUAWEICLOUD_SDK_SK")).
        WithProjectId("{your projectId string}").
        SafeBuild()
    if err != nil {
        fmt.Println(err)
        return
    }

    // Initialize specified New{Service}Client, take initializing the regional service VPC's VpcClient for example
    hcClient, err := vpc.VpcClientBuilder().
        WithEndpoint(endpoint).
        WithCredential(basicAuth).
        WithHttpConfig(config.DefaultHttpConfig()).
        SafeBuild()
    if err != nil {
        fmt.Println(err)
        return
    }
    client := vpc.NewVpcClient(hcClient)
}
```

**where:**

- `endpoint` varies with services and regions,
  see [Regions and Endpoints](https://developer.huaweicloud.com/intl/en-us/endpoint) to obtain correct endpoint.

- When you meet some trouble in getting projectId using the specified region way, you could use this way instead.

#### 3.2 Initialize the {Service}Client with specified Region **(Recommended)** [:top:](#user-manual-top)

``` go
package main

import (
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/global"
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
    iam "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3"
    iamRegion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/region"
    "os"
)

func main() {
    // Initialize the credentials, projectId or domainId could be unassigned in this situation, take initializing GlobalCredentials for example
    globalAuth, err := global.NewCredentialsBuilder().
        WithAk(os.Getenv("HUAWEICLOUD_SDK_AK")).
        WithSk(os.Getenv("HUAWEICLOUD_SDK_SK")).
        // domainId could be unassigned in this situation
        WithDomainId(domainId).
        SafeBuild()
    if err != nil {
        fmt.Println(err)
        return
    }

    // Initialize specified New{Service}Client, take initializing the global service IAM's NewIamClient for example
    hcClient, err := iam.IamClientBuilder().
        WithRegion(iamRegion.CN_NORTH_4).
        WithCredential(globalAuth).
        WithHttpConfig(config.DefaultHttpConfig()).
        SafeBuild()
    if err != nil {
        fmt.Println(err)
        return
    }
    client := iam.NewIamClient(hcClient)
}
```

**Notice:**

- If you use `region` to initialize {Service}Client, projectId/domainId supports automatic acquisition, you don't need
  to configure it when initializing Credentials.

- Multiple ProjectId situation is **not supported**.

- You can query the supported regions through [Regions and Endpoints](https://console-intl.huaweicloud.com/apiexplorer/#/endpoint
). You may get exception such as `Unsupported regionId` if you specify an unsupported region.

**Comparison of the two ways:**

| Initialization | Advantages | Disadvantage |
| :---- | :---- | :---- |
| Specified Endpoint | The API can be invoked successfully once it has been published in the environment. | You need to prepare projectId and endpoint yourself.
| Specified Region | No need for projectId and endpoint, it supports automatic acquisition if you configure it in the right way. | The supported services and regions are limited.

#### 3.3 Custom Configuration [:top:](#user-manual-top)

**Notice:** Supported since v0.0.92

##### 3.3.1 IAM endpoint configuration [:top:](#user-manual-top)

Automatically acquiring projectId/domainId will invoke the [KeystoneListProjects](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=IAM&api=KeystoneListProjects) /[KeystoneListAuthDomains](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=IAM&api=KeystoneListAuthDomains) interface of IAM service. The default iam endpoint is `https://iam.myhuaweicloud.com`, **European station users need to specify the endpoint as https://iam.eu-west-101.myhuaweicloud.eu**, you can modify the endpoint in the following two ways:

###### 3.3.1.1 Global scope [:top:](#user-manual-top)

This configuration takes effect globally, specified by environment variable `HUAWEICLOUD_SDK_IAM_ENDPOINT`

```
//linux
export HUAWEICLOUD_SDK_IAM_ENDPOINT=https://iam.cn-north-4.myhuaweicloud.com

//windows
set HUAWEICLOUD_SDK_IAM_ENDPOINT=https://iam.cn-north-4.myhuaweicloud.com
```

###### 3.3.1.2 Credentials scope [:top:](#user-manual-top)

This configuration is only valid for a credential, and it will override the global configuration

```go
import "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"

iamEndpoint := "https://iam.cn-north-4.myhuaweicloud.com"
cred, err := basic.NewCredentialsBuilder().
            WithAk(os.Getenv("HUAWEICLOUD_SDK_AK")).
            WithSk(os.Getenv("HUAWEICLOUD_SDK_SK")).
            WithIamEndpointOverride(iamEndpoint).
            SafeBuild()
```

##### 3.3.2 Region configuration [:top:](#user-manual-top)

###### 3.3.2.1 Code [:top:](#user-manual-top)

```go
import (
    vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

// Create a region with custom region id and endpoint
reg := region.NewRegion("cn-north-9", "https://ecs.cn-north-9.myhuaweicloud.com")

hcClient, err := ecs.EcsClientBuilder().
    WithRegion(reg).
    WithCredential(auth).
    SafeBuild()
if err != nil {
    // handle error
}
client := ecs.NewEcsClient(hcClient)
```

###### 3.3.2.2 Environment variable [:top:](#user-manual-top)

Specified by environment variable, the format is `HUAWEICLOUD_SDK_REGION_{SERVICE_NAME}_{REGION_ID}={endpoint}`

Notice: the name of environment variable is UPPER-CASE, replacing hyphens with underscores.

```
// Take ECS and IoTDA services as examples

// linux
export HUAWEICLOUD_SDK_REGION_ECS_CN_NORTH_9=https://ecs.cn-north-9.myhuaweicloud.com
export HUAWEICLOUD_SDK_REGION_IOTDA_AP_SOUTHEAST_1=https://iotda.ap-southwest-1.myhuaweicloud.com

// windows
set HUAWEICLOUD_SDK_REGION_ECS_CN_NORTH_9=https://ecs.cn-north-9.myhuaweicloud.com
set HUAWEICLOUD_SDK_REGION_IOTDA_AP_SOUTHEAST_1=https://iotda.ap-southwest-1.myhuaweicloud.com
```

A region corresponding to multiple endpoints is supported since **v0.1.60**, if the main endpoint cannot be connected, it will automatically switch to the backup endpoint.

The format is `HUAWEICLOUD_SDK_REGION_{SERVICE_NAME}_{REGION_ID}={endpoint1},{endpoint2}`, separate multiple endpoints with commas, such as `HUAWEICLOUD_SDK_REGION_ECS_CN_NORTH_9=https://ecs.cn-north-9.myhuaweicloud.com,https://ecs.cn-north-9.myhuaweicloud.cn`

###### 3.3.2.3 Profile [:top:](#user-manual-top)

The profile will be read from the user's home directory by default, linux`~/.huaweicloud/regions.yaml`, windows`C:\Users\USER_NAME\.huaweicloud\regions.yaml`, the default file may not exist, but if the file exists and the content format is incorrect, an exception will be thrown for parsing errors.

The path to the profile can be modified by configuring the environment variable `HUAWEICLOUD_SDK_REGIONS_FILE`, like `HUAWEICLOUD_SDK_REGIONS_FILE=/tmp/my_regions.yml`

The file content format is as follows:

```yaml
# Service name is case-insensitive
ECS:
  - id: 'cn-north-1'
    endpoint: 'https://ecs.cn-north-1.myhuaweicloud.com'
  - id: 'cn-north-9'
    endpoint: 'https://ecs.cn-north-9.myhuaweicloud.com'
IoTDA:
  - id: 'ap-southwest-1'
    endpoint: 'https://iotda.ap-southwest-1.myhuaweicloud.com'
```

A region corresponding to multiple endpoints is supported since v0.1.62, if the main endpoint cannot be connected, it will automatically switch to the backup endpoint.

```yaml
ECS:
  - id: 'cn-north-1'
    endpoints:
      - 'https://ecs.cn-north-1.myhuaweicloud.com'
      - 'https://ecs.cn-north-1.myhuaweicloud.cn'
```

###### 3.3.2.4 Region supply chain [:top:](#user-manual-top)

The default lookup order is **environment variables -> profile -> region defined in SDK** of method **region.ValueOf(regionId)**, if the region is not found in the above ways, an exception will be thrown.

```go
import "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"

region1, err := region.SafeValueOf("cn-north-1")
region2, err := region.SafeValueOf("cn-north-9")
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
package main

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

func main() {
    handler := httphandler.NewHttpHandler().
        AddRequestHandler(RequestHandler).
        AddResponseHandler(ResponseHandler)
    httpConfig := config.DefaultHttpConfig().WithHttpHandler(handler)

    hcClient, err := vpc.VpcClientBuilder().
        WithHttpConfig(httpConfig).
        SafeBuild()
    if err != nil {
        fmt.Println(err)
        return
    }
    client := vpc.NewVpcClient(hcClient)
}
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

func createImageWatermark(client *dsc.DscClient) error {
    // Open the file.
    file, err := os.Open("demo.jpg")
    if err != nil {
        return err
    }
    defer file.Close()

    body := &model.CreateImageWatermarkRequestBody{
        File:           def.NewFilePart(file),
        BlindWatermark: def.NewMultiPart("test123"),
    }

    request := &model.CreateImageWatermarkRequest{Body: body}
    response, err := client.CreateImageWatermark(request)
    if err != nil {
        return err
    }

    fmt.Printf("status code: %d\n", response.HttpStatusCode)

    // Download the file.
    result, err := os.Create("result.jpg")
    if err != nil {
        return err
    }
    _, err = response.Consume(result)
    return err
}

func main() {
    ak := os.Getenv("HUAWEICLOUD_SDK_AK")
    sk := os.Getenv("HUAWEICLOUD_SDK_SK")
    endpoint := "{your endpoint string}"
    projectId := "{your project id}"

    credentials, err := basic.NewCredentialsBuilder().
        WithAk(ak).
        WithSk(sk).
        WithProjectId(projectId).
        SafeBuild()
    if err != nil {
        fmt.Println(err)
        return
    }

    hcClient, err := dsc.DscClientBuilder().
        WithEndpoint(endpoint).
        WithCredential(credentials).
        SafeBuild()
    if err != nil {
        fmt.Println(err)
        return
    }
    client := dsc.NewDscClient(hcClient)
    err := createImageWatermark(client)
}
```

### 7. API Invoker [:top:](#user-manual-top)

#### 7.1 Custom request headers [:top:](#user-manual-top)

You can flexibly configure request headers as needed. **Do not** specify common request headers such as `Host`, `Authorization`, `User-Agent`, `Content-Type` unless necessary, as this may cause the errors.

```go
package main

import (
    "fmt"
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
    vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2"
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/model"
    "os"
)

func main() {
    auth, err := basic.NewCredentialsBuilder().
        WithAk(os.Getenv("HUAWEICLOUD_SDK_AK")).
        WithSk(os.Getenv("HUAWEICLOUD_SDK_SK")).
        WithProjectId("<input your project id>").
        SafeBuild()
    if err != nil {
        fmt.Println(err)
        return
    }

    hcClient, err := vpc.VpcClientBuilder().
        WithEndpoint("<input your endpoint>").
        WithCredential(auth).
        SafeBuild()
    if err != nil {
        fmt.Println(err)
        return
    }
    client := vpc.NewVpcClient(hcClient)

    request := &model.ListVpcsRequest{}
    response, err := client.ListVpcsInvoker(request).
        // custom request headers
        AddHeaders(map[string]string{"key1": "value1", "key2": "value2"}).
        Invoke()

    if err == nil {
        fmt.Printf("%+v\n", response)
    } else {
        fmt.Printf("%+v\n", err)
    }
}
```

#### 7.2 Retry for request [:top:](#user-manual-top)

When a request encounters a network exception or flow control on the interface, the request needs to be retried. The
Go SDK provides the retry method for our users which could be used to the requests of `GET` HTTP method. 
If you want to use the retry method, the following parameters are required:

- _maxRetryTimes_: the max retry times
- _retryCondition_: a function, which determine the condition of when to retry
- _backoffStrategy_: calculate the wait duration before next retry

Take the interface `ListVpcs` of VPC service for example, assume the request would retry at most 3 times, 
retry when service responses an error, the code would be like the following:

``` go
package main

import (
    "fmt"
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker/retry"
    vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2"
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/model"
    "os"
)

func main() {
    auth, err := basic.NewCredentialsBuilder().
        WithAk(os.Getenv("HUAWEICLOUD_SDK_AK")).
        WithSk(os.Getenv("HUAWEICLOUD_SDK_SK")).
        WithProjectId("<input your project id>").
        SafeBuild()
    if err != nil {
        fmt.Println(err)
        return
    }
    // initialize the client
    hcClient, err := vpc.VpcClientBuilder().
        WithEndpoint("<input your endpoint>").
        WithCredential(auth).
        SafeBuild()
    if err != nil {
        fmt.Println(err)
        return
    }
    client := vpc.NewVpcClient(hcClient)

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
}
```
