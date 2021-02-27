English | [简体中文](./README_CN.md)

<p align="center">
<a href="https://www.huaweicloud.com/"><img width="450px" height="102px" src="https://console-static.huaweicloud.com/static/authui/20210202115135/public/custom/images/logo-en.svg"></a>
</p>

<h1 align="center">Huawei Cloud Go Software Development Kit (Go SDK)</h1>

The Huawei Cloud Go SDK allows you to easily work with Huawei Cloud services such as Elastic Compute Service (ECS) and
Virtual Private Cloud (VPC) without the need to handle API related tasks.

This document introduces how to obtain and use Huawei Cloud Go SDK.

## Requirements

- To use Huawei Cloud Go SDK, you must have Huawei Cloud account as well as the Access Key and Secret key of the Huawei
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
- Substitute the values for `{your ak string}`, `{your sk string}`, `{your endpoint string}` and `{your project id}`.

``` go
package main

import (
    "fmt"
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/httphandler"
    vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2"
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/model"
    "net/http"
)

func RequestHandler(request http.Request) {
    fmt.Println(request)
}

func ResponseHandler(response http.Response) {
    fmt.Println(response)
}

func main() {
    client := vpc.NewVpcClient(
        vpc.VpcClientBuilder().
            WithEndpoint("{your endpoint}").
            WithCredential(
                basic.NewCredentialsBuilder().
                    WithAk("{your ak string}").
                    WithSk("{your sk string}").
                    WithProjectId("{your project id}").
                    Build()).
            WithHttpConfig(config.DefaultHttpConfig().
                WithIgnoreSSLVerification(true).
                WithHttpHandler(httphandler.
                    NewHttpHandler().
                        AddRequestHandler(RequestHandler).
                        AddResponseHandler(ResponseHandler))).
            Build())

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
}
```

## Changelog

Detailed changes for each released version are documented in
the [CHANGELOG.md](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/blob/master/CHANGELOG.md).

## User Manual [:top:](#huawei-cloud-go-software-development-kit-go-sdk)

* [1. Client Configuration](#1-client-configuration-top)
    * [1.1  Default Configuration](#11-default-configuration-top)
    * [1.2  Network Proxy](#12-network-proxy-top)
    * [1.3  Connection](#13-connection-top)
    * [1.4  SSL Certification](#14-ssl-certification-top)
* [2. Credentials Configuration](#2-credentials-configuration-top)
    * [2.1  Use Permanent AK&SK](#21-use-permanent-aksk-top)
    * [2.2  Use Temporary AK&SK](#22-use-temporary-aksk-top)
* [3. Client Initialization](#3-client-initialization-top)
    * [3.1  Initialize client with specified Endpoint](#31-initialize-the-serviceclient-with-specified-endpoint-top)
    * [3.2  Initialize client with specified Region](#32-initialize-the-serviceclient-with-specified-region-recommended-top)
* [4. Send Request and Handle response](#4-send-requests-and-handle-responses-top)
    * [4.1  Exceptions](#41-exceptions-top)
* [5. Troubleshooting](#5-troubleshooting-top)
    * [5.1  Original HTTP Listener](#51-original-http-listener-top)

### 1. Client Configuration [:top:](#user-manual-top)

#### 1.1 Default Configuration [:top:](#user-manual-top)

``` go
// Use default configuration
httpConfig := config.DefaultHttpConfig()
```

#### 1.2 Network Proxy [:top:](#user-manual-top)

``` go
// Use proxy if needed
httpConfig.WithProxy(config.NewProxy().
    WithSchema("http").
    WithHost("proxy.huaweicloud.com").
    WithPort(80).
    WithUsername("testuser").
    WithPassword("password"))))
```

#### 1.3 Connection [:top:](#user-manual-top)

``` go
// Seconds to wait for the server to send data before giving up
httpConfig.WithTimeout(120);
```

#### 1.4 SSL Certification [:top:](#user-manual-top)

``` go
// Skip ssl certification checking while using https protocol if needed
httpConfig.WithIgnoreSSLVerification(true);
```

### 2. Credentials Configuration [:top:](#user-manual-top)

There are two types of Huawei Cloud services, `regional` services and `global` services.

Global services contain BSS, DevStar, EPS, IAM, RMS.

For `Regional` services' authentication, projectId is required. For `global` services' authentication, domainId is
required.

`Parameter description`:

- `ak` is the access key ID for your account.
- `sk` is the secret access key for your account.
- `projectId` is the ID of your project depending on your region which you want to operate.
- `domainId` is the account ID of HUAWEI CLOUD.
- `securityToken` is the security token when using temporary AK/SK.

You could use permanent AK plus SK **or** use temporary AK plus SK plus SecurityToken to complete credentials'
configuration.

#### 2.1 Use Permanent AK&SK [:top:](#user-manual-top)

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

**Notice**: projectId/domainId supports **automatic acquisition** in version `0.0.26-beta` or later, if you want to use
this feature, it's required to build your client instance with method `WithRegion()`, detailed example could refer
to [3.2  Initialize client with specified Region](#32-initialize-the-serviceclient-with-specified-region-recommended-top)
.

#### 2.2 Use Temporary AK&SK [:top:](#user-manual-top)

It's required to obtain temporary access key, security key and security token first, which could be obtained through
permanent access key and security key or through an agency.

Obtaining a temporary access key token through permanent access key and security key, you could refer to
document: https://support.huaweicloud.com/en-us/api-iam/iam_04_0002.html . The API mentioned in the document above
corresponds to the method of `CreateTemporaryAccessKeyByToken` in IAM SDK.

Obtaining a temporary access key and security token through an agency, you could refer to
document: https://support.huaweicloud.com/en-us/api-iam/iam_04_0101.html . The API mentioned in the document above
corresponds to the method of `CreateTemporaryAccessKeyByAgency` in IAM SDK.

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

### 3. Client Initialization [:top:](#user-manual-top)

There are two ways to initialize the {Service}Client, you could choose one you preferred.

#### 3.1 Initialize the {Service}Client with specified Endpoint [:top:](#user-manual-top)

``` go
// Initialize specified New{Service}Client, take NewVpcClient for example
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

#### 3.2 Initialize the {Service}Client with specified Region **(Recommended)** [:top:](#user-manual-top)

``` go
import (
    // dependency for region module
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/region"
)

globalCredentials := global.NewCredentialsBuilder().
    WithAk(ak).
    WithSk(sk).
    // domainId could be unassigned in this situation
    Build()

// Initialize specified New{Service}Client, take NewIamClient for example
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
- Multiple ProjectId situation is not supported.

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
func RequestHandler(request http.Request) {
    fmt.Println(request)
}

func ResponseHandler(response http.Response) {
    fmt.Println(response)
}

client := vpc.NewVpcClient(
    vpc.VpcClientBuilder().
        WithEndpoint("{your endpoint}").
        WithCredential(
            basic.NewCredentialsBuilder().
                WithAk("{your ak string}").
                WithSk("{your sk string}").
                WithProjectId("{your project id}").
                   Build()).
        WithHttpConfig(config.DefaultHttpConfig().
            WithIgnoreSSLVerification(true).
            WithHttpHandler(httphandler.
                NewHttpHandler().
                    AddRequestHandler(RequestHandler).
                    AddResponseHandler(ResponseHandler))).
        Build())
```
