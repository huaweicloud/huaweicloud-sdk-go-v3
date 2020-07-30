[English](./README.md) | 简体中文

# 华为云开发者 Go 软件开发工具包 （Go SDK）

欢迎使用华为云 Go SDK。

华为云 Go SDK让您无需关心请求细节即可快速使用云服务器、虚拟私有云等多个华为云服务。

这里将向您介绍如何获取并使用华为云 Go SDK。

## 在线示例

[API Explorer](https://apiexplorer.developer.huaweicloud.com/apiexplorer/overview) 提供API检索及平台调试，支持全量快速检索、可视化调试、帮助文档查看、在线咨询。



## 现在开始

- 要使用华为云 Go SDK，您需要拥有云账号以及该账号对应的Access Key（AK）和Secret Access Key（SK）。 请在华为云控制台“我的凭证-访问密钥”页面上创建和查看您的AKSK。更多信息请查看[访问密钥](https://support.huaweicloud.com/usermanual-ca/zh-cn_topic_0046606340.html).

- 华为云 Go SDK 支持 go 1.14 及以上版本。


## SDK 获取和安装

华为云 Go SDK 支持go 1.14及以上版本。执行``go version``检查当前Go的版本信息.

- 使用 go get 安装华为云Go SDK

    执行如下命令安装华为云 Go SDK库以及相关依赖库:
    
    ```bash
    # 安装华为云Go库
    go get github.com/huaweicloud/huaweicloud-sdk-go-v3
    
    # 安装依赖
    go get github.com/json-iterator/go
    ```

## 开始使用

1. 导入依赖模块:

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

2. 配置客户端属性

    2.1 默认配置

    ```go
    # Use default configuration
    httpConfig := config.DefaultHttpConfig()
    ```

    2.2 网络代理（可选）

    ```go
    // 根据需要配置网络代理
    httpConfig.WithProxy(config.NewProxy().
        WithSchema("http").
        WithHost("proxy.huaweicloud.com").
        WithPort(80).
        WithUsername("testuser").
        WithPassword("password"))))
    ```

    2.3 超时配置（可选）

    ``` go
    httpConfig.WithTimeout(30);
    ```

    2.4 SSL配置（可选）

    ``` go
    // 根据需要配置是否跳过SSL证书校验
    httpConfig.WithIgnoreSSLVerification(true);

3. 初始化客户端:

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

    **说明:**
    - `ak` 华为云账号Access Key。
    - `sk` 华为云账号Secret Access Key。
    - `project_id` 云服务所在项目 ID。
    - `endpoint` 华为云各服务应用区域和各服务的终端节点，详情请查看[地区和终端节点](https://developer.huaweicloud.com/endpoint)。
    
4. 发送请求并查看响应.

    ```go
    # 初始化请求
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

5. 异常处理

    | 一级分类 | 一级分类说明 |
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

6. 原始Http侦听器

    在某些场景下可能对业务发出的Http请求进行Debug，需要看到原始的Http请求和返回信息，SDK提供侦听器功能来获取原始的为加密的Http请求和返回信息。

    > :warning:  Warning: 原始信息打印仅在debug阶段使用，请不要在生产系统中将原始的Http头和Body信息打印到日志，这些信息并未加密且其中包含敏感数据，例如所创建虚拟机的密码，IAM用户的密码等;

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

## 代码实例

使用如下代码同步查询特定 Region 下的 VPC 清单，调用前请根据实际情况替换如下变量：`{your ak string}`、 `{your sk string}`、 `{your endpoint}` 以及 `{your project id}`。

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
