English | [简体中文](./README_CN.md)

# HuaweiCloud Go Software Development Kit (Go SDK)

The HuaweiCloud Go SDK allows you to easily work with Huawei Cloud services such as Elastic Compute Service (ECS) and Virtual Private Cloud(VPC) without the need to handle API related tasks.

This document introduces how to obtain and use HuaweiCloud Go SDK.

## Getting Started

- To use HuaweiCloud Go SDK, you must have Huawei Cloud account as well as the Access Key and Secret key of the HuaweiCloud account.

    The accessKey is required when initializing `{Service}Client`. You can create an AccessKey in the Huawei Cloud console. For more information, see [My Credentials](https://support.huaweicloud.com/en-us/usermanual-ca/en-us_topic_0046606340.html).

- HuaweiCloud Go SDK requires go 1.14 or later.


## Install Go SDK

HuaweiCloud Go SDK supports go 1.14 or later. Run ``go version`` to check the version of Go.

- Use go get

    Run the following command to install the individual libraries of HuaweiCloud services:

    ```bash
    # Install the core library
    go get github.com/huaweicloud/huaweicloud-sdk-go-v3
 
    # Install the dependent library
    go get github.com/json-iterator/go
    ```

## Use Go SDK

1. Import the required modules as follows:

    ```go
    import (
        "fmt"
        "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
        "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
        "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/httphandler"
        vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2"
        "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/model"
        "net/http"
    )
    ```

2. Config `{Service}Client` Configurations

    2.1 Default Configuration

    ```go
    // Use default configuration
    httpConfig := config.DefaultHttpConfig()
    ```

    2.2  Proxy(Optional)

    ```go
    // Use proxy if needed
    httpConfig.WithProxy(config.NewProxy().
        WithSchema("http").
        WithHost("proxy.huaweicloud.com").
        WithPort(80).
        WithUsername("testuser").
        WithPassword("password"))))
    ```

    2.3 Connection(Optional)

    ``` go
    // seconds to wait for the server to send data before giving up
    httpConfig.WithTimeout(30);
    ```

    2.4 SSL Certification(Optional)

    ``` go
    // Skip ssl certification checking while using https protocol if needed
    httpConfig.WithIgnoreSSLVerification(true);

3. Initialize the `{Service}Client` instance:

    ```go
    client := vpc.NewVpcClient(
        vpc.VpcClientBuilder().
            WithEndpoint("{your endpoint}").
            WithCredential(
                basic.NewCredentialsBuilder().
                    WithAk("{your ak string}").
                    WithSk("{your sk string}").
                    WithProjectId("{your project id}").
                    Build()).
            WithHttpConfig(config.DefaultHttpConfig()).  
            Build())
    ```

    **where:**

    - `ak` is the access key id for your account.
    - `sk` is the secret access key for your account.
    - `project_id` is the id of the project.
    - `endpoint` is the service specific endpoints, see [Regions and Endpoints](https://developer.huaweicloud.com/intl/en-us/endpoint)

4. Send a request and print response.

    ```go
    // 初始化请求
    request := &model.CreateVpcRequest{
        Body: &model.CreateVpcRequestBody{
            Vpc: &model.CreateVpcOption{
                Cidr: "192.168.1.0/24",
                Name: "TestVpc",
            },
        },
    }
    response, err := client.CreateVpc(request)
    if err == nil {
        fmt.Printf("%+v\n",response.Vpc)
    } else {
        fmt.Println(err)
    }
    ```

5. Exceptions

    | Level 1 | Notice | 
    | :---- | :---- | 
    | ServiceResponseError | service response error |
    | url.Error | connect endpoint error |
    
    ```go
    response, err := client.CreateVpc(request)
    if err == nil {
        fmt.Println(response)
    } else {
        fmt.Println(err)
    }
    ```

6. Original HTTP Listener

    In some situation, you may need to debug your http requests, original http request and response information will be needed. The SDK provides a listener function to obtain the original encrypted http request and response information.

    > :warning:  Warning: The original http log information are used in debugging stage only, please do not print the original http header or body in the production environment. These log information are not encrypted and contain sensitive data such as the password of your ECS virtual machine or the password of your IAM user account, etc.

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


## Code example

The following example shows how to query a list of VPC in a specific region. Substitute the values for `{your ak string}`, `{your sk string}`, `{your endpoint}` and `{your project id}`.

```go
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

    request := &model.CreateVpcRequest{
        Body: &model.CreateVpcRequestBody{
            Vpc: &model.CreateVpcOption{
                Cidr: "192.168.1.0/24",
                Name: "TestVpc",
            },
        },
    }
    response, err := client.CreateVpc(request)
    if err == nil {
        fmt.Println("%+v\n",response.Vpc)
    } else {
        fmt.Println(err)
    }
}
```
