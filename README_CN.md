[English](./README.md) | 简体中文

<p align="center">
  <a href="https://www.huaweicloud.com/"><img width="270px" height="90px" src="https://console-static.huaweicloud.com/static/authui/20210202115135/public/custom/images/logo.svg"></a>
</p>

<h1 align="center">华为云开发者 Go 软件开发工具包（Go SDK）</h1>

欢迎使用华为云 Go SDK。

华为云 Go SDK 让您无需关心请求细节即可快速使用弹性云服务器、虚拟私有云等多个华为云服务。

这里将向您介绍如何获取并使用华为云 Go SDK 。

## 使用前提

- 要使用华为云 Go SDK ，您需要拥有云账号以及该账号对应的 Access Key（AK）和 Secret Access Key（SK）。请在华为云控制台“我的凭证-访问密钥”页面上创建和查看您的 AK&SK
  。更多信息请查看 [访问密钥](https://support.huaweicloud.com/usermanual-ca/zh-cn_topic_0046606340.html) 。

- 要使用华为云 Go SDK 访问指定服务的 API
  ，您需要确认已在 [华为云控制台](https://console.huaweicloud.com/console/?locale=zh-cn&region=cn-north-4#/home) 开通当前服务。

- 华为云 Go SDK 支持 go 1.14 及以上版本，可执行 `go version` 检查当前 Go 的版本信息。

## SDK 获取和安装

使用 go get 安装华为云 Go SDK ，执行如下命令安装华为云 Go SDK 库：

``` bash
# 安装华为云 Go SDK 库
go get github.com/huaweicloud/huaweicloud-sdk-go-v3
```

## 代码示例

- 使用如下代码在指定 Region 下查询 VPC 列表，实际使用中请将 `vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2"`
  替换为您使用的产品/服务相应的 `{service} "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/{service}/{version}"`
  ，且初始化为 `{service}.New{Service}Client` 。
- 调用前请根据实际情况替换如下变量： `{your ak string}`、`{your sk string}`、`{your endpoint string}` 以及 `{your project id}` 。

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

## 在线调试

[API Explorer](https://apiexplorer.developer.huaweicloud.com/apiexplorer/overview)
提供API检索及平台调试，支持全量快速检索、可视化调试、帮助文档查看、在线咨询。

## 变更日志

每个版本的详细更改记录可在 [变更日志](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/blob/master/CHANGELOG_CN.md) 中查看。

## 用户手册 [:top:](#华为云开发者-go-软件开发工具包go-sdk)

* [1. 客户端连接参数](#1-客户端连接参数-top)
    * [1.1 默认配置](#11-默认配置-top)
    * [1.2 网络代理](#12-网络代理-top)
    * [1.3 超时配置](#13-超时配置-top)
    * [1.4 SSL配置](#14-ssl-配置-top)
    * [1.5 自定义网络连接创建](#15-自定义网络连接创建-top)
* [2. 客户端认证信息](#2-客户端认证信息-top)
  * [2.1 使用永久 AK 和 SK](#21-使用永久-ak-和-sk-top)
    * [2.1.1 手动指定](#211-手动指定)
    * [2.1.2 环境变量](#212-环境变量)
  * [2.2 使用临时 AK 和 SK](#22-使用临时-ak-和-sk-top)
    * [2.2.1 手动指定](#221-手动指定)
    * [2.2.2 元数据获取](#222-元数据获取)
  * [2.3 认证信息提供链](#23-认证信息提供链)
* [3. 客户端初始化](#3-客户端初始化-top)
    * [3.1 指定云服务 Endpoint 方式](#31-指定云服务-endpoint-方式-top)
    * [3.2 指定 Region 方式（推荐）](#32-指定-region-方式-推荐-top)
* [4. 发送请求并查看响应](#4-发送请求并查看响应-top)
    * [4.1 异常处理](#41-异常处理-top)
* [5. 故障处理](#5-故障处理-top)
    * [5.1 HTTP 监听器](#51-http监听器-top)
* [6. 文件上传与下载](#6-文件上传与下载-top)
* [7. 请求重试](#7-请求重试-top)

### 1. 客户端连接参数 [:top:](#用户手册-top)

#### 1.1 默认配置 [:top:](#用户手册-top)

``` go
// 使用默认配置
httpConfig := config.DefaultHttpConfig()
```

#### 1.2 网络代理 [:top:](#用户手册-top)

``` go
// 根据需要配置网络代理
httpConfig.WithProxy(config.NewProxy().
    WithSchema("http").
    WithHost("proxy.huaweicloud.com").
    WithPort(80).
    WithUsername("testuser").
    WithPassword("password"))))
```

#### 1.3 超时配置 [:top:](#用户手册-top)

``` go
// 默认超时时间为120秒，可根据需要配置
httpConfig.WithTimeout(120);
```

#### 1.4 SSL 配置 [:top:](#用户手册-top)

``` go
// 根据需要配置是否跳过SSL证书校验
httpConfig.WithIgnoreSSLVerification(true);
```

#### 1.5 自定义网络连接创建 [:top:](#用户手册-top)

``` go
// 根据需要配置如何创建网络连接
func DialContext(ctx context.Context, network string, addr string) (net.Conn, error) {
	return net.Dial(network, addr)
}
httpConfig.WithDialContext(DialContext)
```

### 2. 客户端认证信息 [:top:](#用户手册-top)

华为云服务存在两种部署方式，Region 级服务和 Global 级服务。

Global 级服务有 BSS、DevStar、EPS、IAM、RMS。

Region 级服务使用 basic.NewCredentialsBuilder() 初始化，需要提供 projectId 。Global 级服务使用 global.NewCredentialsBuilder() 初始化，需要提供
domainId 。

客户端认证可以使用永久 AK&SK 认证，也可以使用临时 AK&SK&SecurityToken 认证。

**认证参数说明**：

- `ak` 华为云账号 Access Key
- `sk` 华为云账号 Secret Access Key
- `projectId` 云服务所在项目 ID ，根据你想操作的项目所属区域选择对应的项目 ID
- `domainId` 华为云账号 ID
- `securityToken` 采用临时 AK&SK 认证场景下的安全票据

#### 2.1 使用永久 AK 和 SK [:top:](#用户手册-top)

##### 2.1.1 手动指定

``` go
// Region 级服务
basicAuth := basic.NewCredentialsBuilder().
    WithAk(ak).
    WithSk(sk).
    WithProjectId(projectId).
    Build()

// Global 级服务
globalAuth := global.NewCredentialsBuilder().
    WithAk(ak).
    WithSk(sk).
    WithDomainId(domainId).
    Build()
```

**说明**：

- `0.0.26-beta` 及以上版本支持自动获取 projectId/domainId ，用户需要指定当前华为云账号的永久 AK&SK 和 对应的 region_id，同时在初始化客户端时配合 `WithRegion()`
  方法使用。 代码示例详见 [3.2 指定Region方式（推荐）](#32-指定-region-方式-推荐-top) 。

##### 2.1.2 环境变量

默认从环境变量`HUAWEICLOUD_SDK_AK`、`HUAWEICLOUD_SDK_SK`、`HUAWEICLOUD_SDK_PROJECT_ID`和`HUAWEICLOUD_SDK_DOMAIN_ID`中读取 ak、sk、projectId 和 domainId。

#### 2.2 使用临时 AK 和 SK [:top:](#用户手册-top)

##### 2.2.1 手动指定

临时AK/SK和securitytoken是系统颁发给IAM用户的临时访问令牌，有效期可在15分钟至24小时范围内设置，过期后需要重新获取。首先需要获得临时 AK、SK 和 SecurityToken ，可以从永久 AK&SK 获得，或者通过委托授权获得。

- 通过永久 AK&SK 获得可以参考文档：https://support.huaweicloud.com/api-iam/iam_04_0002.html ，对应 IAM SDK
  中的 `CreateTemporaryAccessKeyByToken` 方法。

- 通过委托授权获得可以参考文档：https://support.huaweicloud.com/api-iam/iam_04_0101.html ，对应 IAM SDK
  中的 `CreateTemporaryAccessKeyByAgency` 方法。

临时 AK&SK&SecurityToken 获取成功后，可使用如下方式初始化认证信息：

``` go
// Region级服务
basicAuth := basic.NewCredentialsBuilder().
    WithAk(ak).
    WithSk(sk).
    WithProjectId(projectId).
    WithSecurityToken(securityToken).
    Build()

// Global级服务
globalAuth := global.NewCredentialsBuilder().
    WithAk(ak).
    WithSk(sk).
    WithDomainId(domainId).
    WithSecurityToken(securityToken).
    Build()
```

##### 2.2.2 元数据获取

从实例元数据获取临时AK/SK和securitytoken，关于元数据获取请参阅：[元数据获取](https://support.huaweicloud.com/usermanual-ecs/ecs_03_0166.html)

以下两种情况，会尝试从实例元数据中读取认证信息：

1. 创建客户端时未手动指定 basic.Credentials 或 global.Credentials
2. 创建 basic.Credentials 或 global.Credentials 时未指定 AK/SK

```go
// Region级服务
basicAuth := basic.NewCredentialsBuilder().WithProjectId(projectId).Build()

// Global级服务
globalAuth := global.NewCredentialsBuilder().WithDomainId(domainId).Build()
```

#### 2.3 认证信息提供链

在创建客户端时按照以下顺序加载认证信息：

1. [手动指定](#211-手动指定) basic.Credentials 或 global.Credentials
2. 未手动指定，尝试从 [环境变量](#212-环境变量) 加载
3. 环境变量中获取不到，尝试从 [实例元数据](#222-元数据获取) 读取临时认证信息

### 3. 客户端初始化 [:top:](#用户手册-top)

客户端初始化有两种方式，可根据需要选择下列两种方式中的一种：

#### 3.1 指定云服务 Endpoint 方式 [:top:](#用户手册-top)

``` go
// 指定终端节点，以 VPC 服务北京四的 endpoint 为例
endpoint := "https://vpc.cn-north-4.myhuaweicloud.com"
// 初始化客户端认证信息，需要填写相应 projectId/domainId，以初始化 basic.NewCredentialsBuilder() 为例
basicAuth := basic.NewCredentialsBuilder().
    WithAk(ak).
    WithSk(sk).
    WithProjectId(projectId).
    Build()
 
// 初始化指定云服务的客户端 New{Service}Client ，以初始化 Region 级服务 VPC 的 NewVpcClient 为例
client := vpc.NewVpcClient(
    vpc.VpcClientBuilder().
        WithEndpoint(endpoint).
        WithCredential(basicAuth).
        WithHttpConfig(config.DefaultHttpConfig()).  
        Build())
```

**说明:**

- `endpoint` 是华为云各服务应用区域和各服务的终端节点，详情请查看 [地区和终端节点](https://developer.huaweicloud.com/endpoint) 。
- 当用户使用指定 Region 方式无法自动获取 projectId 时，可以使用当前方式调用接口。

#### 3.2 指定 Region 方式 **（推荐）** [:top:](#用户手册-top)

``` go
import (
    // 增加region依赖
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/region"
)

// 初始化客户端认证信息，使用当前客户端初始化方式可不填 projectId/domainId，以初始化 global.NewCredentialsBuilder() 为例
globalAuth := global.NewCredentialsBuilder().
            WithAk(ak).
            WithSk(sk).
            // 可不填 domainId
            Build()

// 初始化指定云服务的客户端 New{Service}Client ，以初始化 Global 级服务 IAM 的 IamClient 为例
client := iam.NewIamClient(
    iam.IamClientBuilder().
        WithRegion(region.CN_NORTH_4).
        WithCredential(globalAuth).
        WithHttpConfig(config.DefaultHttpConfig()).  
        Build())
```

**说明：**

- 指定 Region 方式创建客户端的场景，支持自动获取用户的 projectId 或者 domainId，初始化认证信息时可无需指定相应参数。

- **不适用**于 `多ProjectId` 的场景。

- 当前支持指定 Region 方式初始化客户端的 region_id : af-south-1, ap-southeast-1, ap-southeast-2, ap-southeast-3, cn-east-2, cn-east-3,
  cn-north-1, cn-north-4, cn-south-1, cn-southwest-2, ru-northwest-2。调用其他 region 可能会抛出 `Unsupported regionId` 的异常信息。

**两种方式对比：**

| 初始化方式 | 优势 | 劣势 |
| :---- | :---- | :---- | 
| 指定云服务 Endpoint 方式 | 只要接口已在当前环境发布就可以成功调用 | 需要用户自行查找并填写 projectId 和 endpoint
| 指定 Region 方式 | 无需指定 projectId 和 endpoint，按照要求配置即可自动获取该值并回填 | 支持的服务和 region 有限制

### 4. 发送请求并查看响应 [:top:](#用户手册-top)

``` go
// 初始化请求,，以调用接口 ListVpcs 为例
limit := int32(1)
request := &model.ListVpcsRequest{
    Limit: &limit,
}

response, err := client.ListVpcs(request)
if err == nil {
    fmt.Printf("%+v\n", response.Vpcs)
} else {
    fmt.Println(err)
}
```

#### 4.1 异常处理 [:top:](#用户手册-top)

| 一级分类 | 一级分类说明 |
| :---- | :---- | 
| ServiceResponseError | service response error |
| url.Error | connect endpoint error |

``` go
response, err := client.ListVpcs(request)
if err == nil {
    fmt.Printf("%+v\n", response.Vpcs)
} else {
    fmt.Println(err)
}
```

### 5. 故障处理 [:top:](#用户手册-top)

#### 5.1 HTTP监听器 [:top:](#用户手册-top)

在某些场景下可能对业务发出的Http请求进行Debug，需要看到原始的Http请求和返回信息，SDK提供监听器功能来获取原始的为加密的Http请求和返回信息。

> :warning:  Warning: 原始信息打印仅在 Debug 阶段使用，请不要在生产系统中将原始的 HTTP 头和 Body 信息打印到日志，这些信息并未加密且其中包含敏感数据，例如所创建虚拟机的密码，IAM 用户的密码等；当 Body 体为二进制内容，即 Content-Type 标识为二进制时，Body 为"***"，详细内容不输出。

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

### 6. 文件上传与下载 [:top:](#用户手册-top)

以数据安全中心服务的嵌入图片水印接口为例，该接口需要上传一个图片文件，并返回加过水印的图片文件流：

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
	
	// 打开文件
	file, err := os.Open("demo.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	body := &model.CreateImageWatermarkRequestBody{
		File:           def.NewFilePart(file),
		BlindWatermark: def.NewMultiPart("test123"),
	}

	request := &model.CreateImageWatermarkRequest{Body: body}
	response, err := client.CreateImageWatermark(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
		return
	}

	// 下载文件
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

### 7. 请求重试 [:top:](#用户手册-top)

当请求遇到网络异常或者流控场景的时候，通常需要对请求进行重试。Go SDK 提供了请求重试的入口，可用于请求方式为 `GET`
的请求。如需使用重试，需要配置最大重试次数、重试条件和重试策略。其中，

- _最大重试次数_：最大重试次数
- _重试条件_：函数，根据上一次请求的返回情况给出是否需要重试
- _重试策略_：每次重试前的等待时间，支持多种重试策略

以 VPC 服务的 `ListVpcs` 接口为例，最多重试3次，服务端返回异常时进行重试，代码如下：

``` go
// 初始化同步客户端
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

// 初始化请求
request := &model.ListVpcsRequest{}

// 发送请求，当请求异常时进行重试
response, err := client.ListVpcsInvoker(request).WithRetry(3, func(i interface{}, err error) bool {
	return err != nil
}, new(retry.None)).Invoke()

if err == nil {
	fmt.Printf("%+v\n", response)
} else {
	fmt.Printf("%+v\n", err)
}
```
