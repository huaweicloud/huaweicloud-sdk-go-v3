[English](./README.md) | 简体中文

<p align="center">
  <a href="https://www.huaweicloud.com/"><img width="270px" height="90px" src="https://console-static.huaweicloud.com/static/authui/20210202115135/public/custom/images/logo.svg"></a>
</p>

<h1 align="center">华为云开发者 Go 软件开发工具包（Go SDK）</h1>

欢迎使用华为云 Go SDK。

华为云 Go SDK 让您无需关心请求细节即可快速使用弹性云服务器（ECS）、虚拟私有云（VPC）等多个华为云服务。

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

您可以通过 [SDK中心](https://console.huaweicloud.com/apiexplorer/#/sdkcenter?language=Go) 或 [Github Releases](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/releases?page=1) 查询SDK版本信息。

## 代码示例

- 使用如下代码在指定 Region 下查询 VPC 列表，实际使用中请将 `vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2"`
  替换为您使用的产品/服务相应的 `{service} "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/{service}/{version}"`
  ，且初始化为 `{service}.New{Service}Client` 。
- 认证用的ak和sk直接写到代码中有很大的安全风险，建议在配置文件或者环境变量中密文存放，使用时解密，确保安全。
- 本示例中的ak和sk保存在环境变量中，运行本示例前请先在本地环境中配置环境变量`HUAWEICLOUD_SDK_AK`和`HUAWEICLOUD_SDK_SK`。

**精简示例**

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
    // 配置认证信息
    // 可通过环境变量等方式配置认证信息，参考2.4认证信息管理章节
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

    // 创建服务客户端
    hcClient, err := vpc.VpcClientBuilder().
        WithRegion(region).
        WithCredential(auth).
        SafeBuild()
    if err != nil {
        fmt.Println(err)
        return
    }
    client := vpc.NewVpcClient(hcClient)

    // 发送请求并获取响应
    request := &vpcModel.ListVpcsRequest{}
    response, err := client.ListVpcs(request)
    if err == nil {
        fmt.Printf("%+v\n", response)
    } else {
        fmt.Println(err)
    }
}
```

**详细示例**

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
    // 配置认证信息
    auth, err := basic.NewCredentialsBuilder().
        // 可通过环境变量等方式配置认证信息，参考2.4认证信息管理章节
        WithAk(os.Getenv("HUAWEICLOUD_SDK_AK")).
        WithSk(os.Getenv("HUAWEICLOUD_SDK_SK")).
        // 如果未填写ProjectId，SDK会自动调用IAM服务查询所在region对应的项目id
        WithProjectId("{your projectId string}").
        // 配置SDK内置的IAM服务地址，默认为https://iam.myhuaweicloud.com
        WithIamEndpointOverride("https://iam.cn-north-4.myhuaweicloud.com").
        SafeBuild()
    if err != nil {
        fmt.Println(err)
        return
    }

    // 使用默认配置
    httpConfig := config.DefaultHttpConfig()
    // 配置是否忽略SSL证书校验， 默认不忽略
    httpConfig.WithIgnoreSSLVerification(true)
    // 默认超时时间为120秒，可根据需要配置
    httpConfig.WithTimeout(120)
    // 根据需要配置网络代理
    proxy := config.NewProxy().
        // 请根据实际情况替换示例中的代理协议、地址和端口号
        WithSchema("http").
        WithHost("proxy.huaweicloud.com").
        WithPort(80).
        // 如果代理需要认证，请配置用户名和密码
        WithUsername(os.Getenv("PROXY_USERNAME")).
        WithPassword(os.Getenv("PROXY_PASSWORD"))
    httpConfig.WithProxy(proxy)
    // 根据需要配置如何创建网络连接
    dialContext := func(ctx context.Context, network string, addr string) (net.Conn, error) {
        // 此处需自行实现
    }
    httpConfig.WithDialContext(dialContext)
    // 配置HTTP监听器进行调试，请勿用于生产环境
    requestHandler := func(request http.Request) {
        fmt.Println(request)
    }
    responseHandler := func(response http.Response) {
        fmt.Println(response)
    }
    httpHandler := httphandler.NewHttpHandler().AddRequestHandler(requestHandler).AddResponseHandler(responseHandler)
    httpConfig.WithHttpHandler(httpHandler)

    // 获取可用地区
    region, err := vpcRegion.SafeValueOf("cn-north-4")
    if err != nil {
        fmt.Println(err)
        return
    }
    // 创建服务客户端
    hcClient, err := vpc.VpcClientBuilder().
        // 配置地区, 如果地区不存在会导致panic
        WithRegion(region).
        // 配置认证信息
        WithCredential(auth).
        // HTTP配置
        WithHttpConfig(httpConfig).
        SafeBuild()
    if err != nil {
        fmt.Println(err)
        return
    }
    client := vpc.NewVpcClient(hcClient)

    // 创建请求
    request := &vpcModel.ListVpcsRequest{}
    // 配置每页返回的个数
    limit := int32(1)
    request.Limit = &limit

    // 发送请求并获取响应
    response, err := client.ListVpcs(request)
    // 处理异常，打印响应信息
    if err == nil {
        fmt.Printf("%+v\n", response)
    } else {
        fmt.Println(err)
    }
}
```

## 在线调试

[API Explorer](https://apiexplorer.developer.huaweicloud.com/apiexplorer/overview)
提供API检索、SDK示例及平台调试，支持全量快速检索、可视化调试、帮助文档查看、在线咨询。

## 变更日志

每个版本的详细更改记录可在 [变更日志](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/blob/master/CHANGELOG_CN.md) 中查看。

## 用户手册 [:top:](#华为云开发者-go-软件开发工具包go-sdk)

* [1. 客户端连接参数](#1-客户端连接参数-top)
    * [1.1 默认配置](#11-默认配置-top)
    * [1.2 网络代理](#12-网络代理-top)
    * [1.3 超时配置](#13-超时配置-top)
    * [1.4 SSL配置](#14-ssl-配置-top)
    * [1.5 自定义网络连接创建](#15-自定义网络连接创建-top)
    * [1.6 自定义HTTP传输](#16-自定义http传输-top)
* [2. 认证信息配置](#2-认证信息配置-top)
  * [2.1 使用永久 AK 和 SK](#21-使用永久-ak-和-sk-top)
  * [2.2 使用临时 AK 和 SK](#22-使用临时-ak-和-sk-top)
  * [2.3 使用 IdpId 和 IdTokenFile](#23-使用-idpid-和-idtokenfile-top)
  * [2.4 认证信息管理](#24-认证信息管理-top)
    * [2.4.1 环境变量](#241-环境变量-top)
    * [2.4.2 配置文件](#242-配置文件-top)
    * [2.4.3 实例元数据](#243-实例元数据-top)
    * [2.4.4 认证信息提供链](#244-认证信息提供链-top)
* [3. 客户端初始化](#3-客户端初始化-top)
  * [3.1 指定云服务 Endpoint 方式](#31-指定云服务-endpoint-方式-top)
  * [3.2 指定 Region 方式（推荐）](#32-指定-region-方式-推荐-top)
  * [3.3 自定义配置](#33-自定义配置-top)
    * [3.3.1 IAM endpoint配置](#331-iam-endpoint配置-top)
    * [3.3.2 Region配置](#332-region配置-top)
* [4. 发送请求并查看响应](#4-发送请求并查看响应-top)
    * [4.1 异常处理](#41-异常处理-top)
* [5. 故障处理](#5-故障处理-top)
    * [5.1 HTTP 监听器](#51-http监听器-top)
* [6. 文件上传与下载](#6-文件上传与下载-top)
* [7. 接口调用器](#7-接口调用器-top)
    * [7.1 自定义请求头](#71-自定义请求头-top)
    * [7.2 请求重试](#72-请求重试-top)

### 1. 客户端连接参数 [:top:](#用户手册-top)

#### 1.1 默认配置 [:top:](#用户手册-top)

``` go
import (
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
    vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2"
)

// 使用默认配置
httpConfig := config.DefaultHttpConfig()

hcClient, err := vpc.VpcClientBuilder().
    WithHttpConfig(httpConfig).
    SafeBuild()
if err != nil {
    // 处理错误
}
client := vpc.NewVpcClient(hcClient)
```

#### 1.2 网络代理 [:top:](#用户手册-top)

``` go
// 根据需要配置网络代理
proxy := config.NewProxy().
    // 请根据实际情况替换示例中的代理协议、地址和端口号
    WithSchema("http").
    WithHost("proxy.huaweicloud.com").
    WithPort(80).
    // 如果代理需要认证，请配置用户名和密码
    // 本示例中的账号和密码保存在环境变量中，运行本示例前请先在本地环境中配置环境变量PROXY_USERNAME和PROXY_PASSWORD
    WithUsername(os.Getenv("PROXY_USERNAME")).
    WithPassword(os.Getenv("PROXY_PASSWORD"))
httpConfig := config.DefaultHttpConfig().WithProxy(proxy)

hcClient, err := vpc.VpcClientBuilder().
    WithHttpConfig(httpConfig).
    SafeBuild()
if err != nil {
    // 处理错误
}
client := vpc.NewVpcClient(hcClient)
```

#### 1.3 超时配置 [:top:](#用户手册-top)

``` go
// 默认超时时间为120秒，可根据需要配置
httpConfig := config.DefaultHttpConfig().WithTimeout(120)

hcClient, err := vpc.VpcClientBuilder().
    WithHttpConfig(httpConfig).
    SafeBuild()
if err != nil {
    // 处理错误
}
client := vpc.NewVpcClient(hcClient)
```

#### 1.4 SSL 配置 [:top:](#用户手册-top)

``` go
// 根据需要配置是否跳过SSL证书校验, 默认不忽略
httpConfig := config.DefaultHttpConfig().WithIgnoreSSLVerification(true)

hcClient, err := vpc.VpcClientBuilder().
    WithHttpConfig(httpConfig).
    SafeBuild()
if err != nil {
    // 处理错误
}
client := vpc.NewVpcClient(hcClient)
```

#### 1.5 自定义网络连接创建 [:top:](#用户手册-top)

``` go
// 根据需要配置如何创建网络连接
func DialContext(ctx context.Context, network string, addr string) (net.Conn, error) {
    return net.Dial(network, addr)
}
httpConfig := config.DefaultHttpConfig().WithDialContext(DialContext)

hcClient, err := vpc.VpcClientBuilder().
    WithHttpConfig(httpConfig).
    SafeBuild()
if err != nil {
    // 处理错误
}
client := vpc.NewVpcClient(hcClient)
```

#### 1.6 自定义HTTP传输 [:top:](#用户手册-top)

支持配置**HttpTransport**或**HttpRoundTripper**(v0.1.114版本以上)，前者是对后者的接口实现，选择其一配置即可。

**注意：** HttpTransport配置项拥有最高优先级。

如果配置了**HttpTransport**或**HttpRoundTripper**，**会导致 [1.2 网络代理](#12-网络代理-top)、[1.4 SSL配置](#14-ssl-配置-top)、[1.5 自定义网络连接创建](#15-自定义网络连接创建-top) 配置失效。**

``` go
transport := &http.Transport{}
httpConfig := config.DefaultHttpConfig().WithHttpTransport(transport)
// httpConfig.WithHttpRoundTripper(&YourRoundTripper{})

hcClient, err := vpc.VpcClientBuilder().
    WithHttpConfig(httpConfig).
    SafeBuild()
if err != nil {
    // 处理错误
}
client := vpc.NewVpcClient(hcClient)
```

### 2. 认证信息配置 [:top:](#用户手册-top)

华为云服务存在两种部署方式，Region 级服务和 Global 级服务。

Global 级服务有 BSS、DevStar、EPS、IAM、RMS。

Region 级服务使用 basic.NewCredentialsBuilder() 初始化，需要提供 projectId 。

Global 级服务使用 global.NewCredentialsBuilder() 初始化，需要提供 domainId 。

客户端认证方式支持以下几种：

- 永久 AK&SK 认证
- 临时 AK&SK&SecurityToken 认证
- IdpId&IdTokenFile 认证

#### 2.1 使用永久 AK 和 SK [:top:](#用户手册-top)

**认证参数说明**：

- `ak` 华为云账号 Access Key
- `sk` 华为云账号 Secret Access Key
- `projectId` 云服务所在项目 ID ，根据你想操作的项目所属区域选择对应的项目 ID
- `domainId` 华为云账号 ID

``` go
import (
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/global"
    "os"
)

// Region 级服务
ak := os.Getenv("HUAWEICLOUD_SDK_AK")
sk := os.Getenv("HUAWEICLOUD_SDK_SK")
projectId := "{your projectId string}"

basicAuth, err := basic.NewCredentialsBuilder().
    WithAk(ak).
    WithSk(sk).
    WithProjectId(projectId).
    SafeBuild()

// Global 级服务
ak := os.Getenv("HUAWEICLOUD_SDK_AK")
sk := os.Getenv("HUAWEICLOUD_SDK_SK")
domainId := "{your domainId string}"

globalAuth, err := global.NewCredentialsBuilder().
    WithAk(ak).
    WithSk(sk).
    WithDomainId(domainId).
    SafeBuild()
```

**说明**：

- `0.0.26-beta` 及以上版本支持自动获取 projectId/domainId ，用户需要指定当前华为云账号的永久 AK&SK 和 对应的 region_id，同时在初始化客户端时配合 `WithRegion()`
  方法使用。 代码示例详见 [3.2 指定Region方式（推荐）](#32-指定-region-方式-推荐-top) 。

#### 2.2 使用临时 AK 和 SK [:top:](#用户手册-top)

临时AK/SK和securitytoken是系统颁发给IAM用户的临时访问令牌，有效期可在15分钟至24小时范围内设置，过期后需要重新获取。首先需要获得临时 AK、SK 和 SecurityToken ，可以从永久 AK&SK 获得，或者通过委托授权获得。

- 通过永久 AK&SK 获得可以参考文档：https://support.huaweicloud.com/api-iam/iam_04_0002.html ，对应 IAM SDK
  中的 `CreateTemporaryAccessKeyByToken` 方法。

- 通过委托授权获得可以参考文档：https://support.huaweicloud.com/api-iam/iam_04_0101.html ，对应 IAM SDK
  中的 `CreateTemporaryAccessKeyByAgency` 方法。

**认证参数说明**：

- `ak` 华为云账号 Access Key
- `sk` 华为云账号 Secret Access Key
- `securityToken` 采用临时 AK&SK 认证场景下的安全票据
- `projectId` 云服务所在项目 ID ，根据你想操作的项目所属区域选择对应的项目 ID
- `domainId` 华为云账号 ID

临时 AK&SK&SecurityToken 获取成功后，可使用如下方式初始化认证信息：

``` go
import (
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/global"
    "os"
)

// Region级服务
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

// Global级服务
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

以下两种情况，会尝试从**实例元数据**中读取获取临时AK/SK和securitytoken：

1. 创建客户端时未显式指定 basic.Credentials 或 global.Credentials
2. 创建 basic.Credentials 或 global.Credentials 时未显式指定 AK/SK

关于元数据获取请参阅：[元数据获取](https://support.huaweicloud.com/usermanual-ecs/ecs_03_0166.html)

```go
import (
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/global"
)

// Region级服务
basicAuth, err := basic.NewCredentialsBuilder().WithProjectId(projectId).SafeBuild()

// Global级服务
globalAuth, err := global.NewCredentialsBuilder().WithDomainId(domainId).SafeBuild()
```

#### 2.3 使用 IdpId 和 IdTokenFile [:top:](#用户手册-top)

通过OpenID Connect ID token方式获取联邦认证token, 可参考文档：[获取联邦认证token(OpenID Connect ID token方式)](https://support.huaweicloud.com/api-iam/iam_13_0605.html)

**认证参数说明**：

- `idpId` 身份提供商ID
- `idTokenFile` 存放id_token的文件路径，id_token由企业IdP构建，携带联邦用户身份信息
- `projectId` 云服务所在项目 ID ，根据你想操作的项目所属区域选择对应的项目 ID
- `domainId` 华为云账号 ID

```go
import (
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/global"
)

// Region级服务
basicAuth, err := basic.NewCredentialsBuilder().
    WithIdpId(idpId).
    WithIdTokenFile(idTokenFile).
    WithProjectId(projectId).
    SafeBuild()

// Global级服务
globalAuth, err := global.NewCredentialsBuilder().
    WithIdpId(idpId).
    WithIdTokenFile(idTokenFile).
    WithDomainId(domainId).
    SafeBuild()
```

#### 2.4 认证信息管理 [:top:](#用户手册-top)

从**0.0.96**版本起，支持从各类提供器中获取认证信息

**Region级服务** 请使用 `BasicCredentialXxxProvider`, **Global级服务** 请使用 `GlobalCredentialXxxProvider`

##### 2.4.1 环境变量 [:top:](#用户手册-top)

**AK/SK认证**

| 环境变量  |  说明 |
| ------------ | ------------ |
| HUAWEICLOUD_SDK_AK  | 必填，AccessKey  |
| HUAWEICLOUD_SDK_SK  |  必填，SecretKey |
| HUAWEICLOUD_SDK_SECURITY_TOKEN  | 可选, 使用临时ak/sk认证时需要指定该参数  |
| HUAWEICLOUD_SDK_PROJECT_ID  | 可选，用于Region级服务，多ProjectId场景下必填  |
| HUAWEICLOUD_SDK_DOMAIN_ID  | 可选，用于Global级服务  |

配置环境变量：

```
// Linux
export HUAWEICLOUD_SDK_AK=YOUR_AK
export HUAWEICLOUD_SDK_SK=YOUR_SK

// Windows
set HUAWEICLOUD_SDK_AK=YOUR_AK
set HUAWEICLOUD_SDK_SK=YOUR_SK
```

从配置的环境变量中获取认证信息：

```go
import "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/provider"

// basic
basicProvider := provider.BasicCredentialEnvProvider()
basicCred, err := basicProvider.GetCredentials()

// global
globalProvider := provider.GlobalCredentialEnvProvider()
globalCred, err := globalProvider.GetCredentials()
```

**IdpId/IdTokenFile认证**

| 环境变量  |  说明 |
| ------------ | ------------ |
| HUAWEICLOUD_SDK_IDP_ID  | 必填，身份提供商ID |
| HUAWEICLOUD_SDK_ID_TOKEN_FILE  |  必填，存放id_token的文件路径 |
| HUAWEICLOUD_SDK_PROJECT_ID  | basic类型认证时，该参数必填  |
| HUAWEICLOUD_SDK_DOMAIN_ID  | global类型认证时，该参数必填  |

配置环境变量：

```
// Linux
export HUAWEICLOUD_SDK_IDP_ID=YOUR_IDP_ID
export HUAWEICLOUD_SDK_ID_TOKEN_FILE=/some_path/your_token_file
export HUAWEICLOUD_SDK_PROJECT_ID=YOUR_PROJECT_ID // basic认证时必填
export HUAWEICLOUD_SDK_DOMAIN_ID=YOUR_DOMAIN_ID // global认证时必填

// Windows
set HUAWEICLOUD_SDK_IDP_ID=YOUR_IDP_ID
set HUAWEICLOUD_SDK_ID_TOKEN_FILE=/some_path/your_token_file
set HUAWEICLOUD_SDK_PROJECT_ID=YOUR_PROJECT_ID // basic认证时必填
set HUAWEICLOUD_SDK_DOMAIN_ID=YOUR_DOMAIN_ID // global认证时必填
```

从配置的环境变量中获取认证信息：

```go
import "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/provider"

// basic
basicProvider := provider.BasicCredentialEnvProvider()
basicCred, err := basicProvider.GetCredentials()

// global
globalProvider := provider.GlobalCredentialEnvProvider()
globalCred, err := globalProvider.GetCredentials()
```

##### 2.4.2 配置文件 [:top:](#用户手册-top)

默认会从用户主目录下读取认证信息配置文件，linux为`~/.huaweicloud/credentials`，windows为`C:\Users\USER_NAME\.huaweicloud\credentials`，可以通过配置环境变量`HUAWEICLOUD_SDK_CREDENTIALS_FILE`来修改默认文件的路径

**AK/SK认证**

| 配置参数  |  说明 |
| ------------ | ------------ |
| ak  | 必填，AccessKey  |
| sk  |  必填，SecretKey |
| security_token  | 可选, 使用临时ak/sk认证时需要指定该参数  |
| project_id  | 可选，用于Region级服务，多ProjectId场景下必填  |
| domain_id  | 可选，用于Global级服务  |
| iam_endpoint  | 可选，用于身份认证的endpoint，默认为`https://iam.myhuaweicloud.com` |

配置文件内容如下：

```ini
[basic]
ak = your_ak
sk = your_sk

[global]
ak = your_ak
sk = your_sk
```

从配置文件中读取认证信息：

```go
import "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/provider"

// basic
basicProvider := provider.BasicCredentialProfileProvider()
basicCred, err := basicProvider.GetCredentials()

// global
globalProvider := provider.GlobalCredentialProfileProvider()
globalCred, err := globalProvider.GetCredentials()
```

**IdpId/IdTokenFile认证**

| 配置参数  |  说明 |
| ------------ | ------------ |
| idp_id  | 必填，身份提供商ID |
| id_token_file  |  必填，存放id_token的文件路径 |
| project_id  | basic类型认证时，该参数必填  |
| domain_id  | global类型认证时，该参数必填  |
| iam_endpoint  | 可选，用于身份认证的endpoint，默认为`https://iam.myhuaweicloud.com` |

配置文件内容如下：

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

从配置文件中读取认证信息：

```go
import "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/provider"

// basic
basicProvider := provider.BasicCredentialProfileProvider()
basicCred, err := basicProvider.GetCredentials()

// global
globalProvider := provider.GlobalCredentialProfileProvider()
globalCred, err := globalProvider.GetCredentials()
```

##### 2.4.3 实例元数据 [:top:](#用户手册-top)

从实例元数据获取临时AK/SK和securitytoken，关于元数据获取请参阅：[元数据获取](https://support.huaweicloud.com/usermanual-ecs/ecs_03_0166.html)

手动获取实例元数据认证信息：

```go
import "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/provider"

// basic
basicProvider := provider.BasicCredentialMetadataProvider()
basicCred, err := basicProvider.GetCredentials()

// global
globalProvider := provider.GlobalCredentialMetadataProvider()
globalCred, err := globalProvider.GetCredentials()
```

##### 2.4.4 认证信息提供链 [:top:](#用户手册-top)

在创建服务客户端，未显式指定认证信息时，按照顺序 **环境变量 -> 配置文件 -> 实例元数据** 尝试加载认证信息

通过提供链获取认证信息：

```go
import "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/provider"

// basic
basicChain := provider.BasicCredentialProviderChain()
basicCred, err := basicChain.GetCredentials()

// global
globalChain := provider.GlobalCredentialProviderChain()
globalCred, err := globalChain.GetCredentials()
```

支持自定义认证信息提供链：

```go
import "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/provider"

providers := []provider.ICredentialProvider{
    provider.BasicCredentialMetadataProvider(),
    provider.BasicCredentialProfileProvider(),
}
chain := provider.NewCredentialProviderChain(providers)
cred, err := chain.GetCredentials()
```

### 3. 客户端初始化 [:top:](#用户手册-top)

客户端初始化有两种方式，可根据需要选择下列两种方式中的一种：

#### 3.1 指定云服务 Endpoint 方式 [:top:](#用户手册-top)

``` go
package main

import (
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
    vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2"
    "os"
)

func main() {
    // 指定终端节点，以 VPC 服务北京四的 endpoint 为例
    endpoint := "https://vpc.cn-north-4.myhuaweicloud.com"
    // 初始化客户端认证信息，需要填写相应 projectId/domainId，以初始化 basic.NewCredentialsBuilder() 为例
    basicAuth, err := basic.NewCredentialsBuilder().
        WithAk(os.Getenv("HUAWEICLOUD_SDK_AK")).
        WithSk(os.Getenv("HUAWEICLOUD_SDK_SK")).
        WithProjectId("{your projectId string}").
        SafeBuild()
    if err != nil {
        fmt.Println(err)
        return
    }

    // 初始化指定云服务的客户端 New{Service}Client ，以初始化 Region 级服务 VPC 的 NewVpcClient 为例
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

**说明:**

- `endpoint` 是华为云各服务应用区域和各服务的终端节点，详情请查看 [地区和终端节点](https://developer.huaweicloud.com/endpoint) 。
- 当用户使用指定 Region 方式无法自动获取 projectId 时，可以使用当前方式调用接口。

#### 3.2 指定 Region 方式 **（推荐）** [:top:](#用户手册-top)

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
    // 初始化客户端认证信息，使用当前客户端初始化方式可不填 projectId/domainId，以初始化 global.NewCredentialsBuilder() 为例
    globalAuth, err := global.NewCredentialsBuilder().
        WithAk(os.Getenv("HUAWEICLOUD_SDK_AK")).
        WithSk(os.Getenv("HUAWEICLOUD_SDK_SK")).
        // 可不填 domainId
        WithDomainId(domainId).
        SafeBuild()
    if err != nil {
        fmt.Println(err)
        return
    }

    // 初始化指定云服务的客户端 New{Service}Client ，以初始化 Global 级服务 IAM 的 IamClient 为例
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

**说明：**

- 指定 Region 方式创建客户端的场景，支持自动获取用户的 projectId 或者 domainId，初始化认证信息时可无需指定相应参数。

- **不适用**于 `多ProjectId` 的场景。

- 支持指定的 Region 可通过[地区和终端节点](https://console.huaweicloud.com/apiexplorer/#/endpoint)查询。调用不支持的 region 可能会抛出 `Unsupported regionId` 的异常信息。

**两种方式对比：**

| 初始化方式 | 优势 | 劣势 |
| :---- | :---- | :---- | 
| 指定云服务 Endpoint 方式 | 只要接口已在当前环境发布就可以成功调用 | 需要用户自行查找并填写 projectId 和 endpoint
| 指定 Region 方式 | 无需指定 projectId 和 endpoint，按照要求配置即可自动获取该值并回填 | 支持的服务和 region 有限制

#### 3.3 自定义配置 [:top:](#用户手册-top)

**注：**0.0.92版本起支持

##### 3.3.1 IAM endpoint配置 [:top:](#用户手册-top)

自动获取用户的 projectId 和 domainId 会分别调用统一身份认证服务的 [KeystoneListProjects](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=IAM&api=KeystoneListProjects) 和 [KeystoneListAuthDomains](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=IAM&api=KeystoneListAuthDomains) 接口，默认访问的endpoint为 https://iam.myhuaweicloud.com， **欧洲站用户需要指定 endpoint 为 https://iam.eu-west-101.myhuaweicloud.eu**

用户可以通过以下两种方式来修改endpoint

###### 3.3.1.1 全局级 [:top:](#用户手册-top)

全局范围生效，通过环境变量`HUAWEICLOUD_SDK_IAM_ENDPOINT`指定

```
//linux
export HUAWEICLOUD_SDK_IAM_ENDPOINT=https://iam.cn-north-4.myhuaweicloud.com

//windows
set HUAWEICLOUD_SDK_IAM_ENDPOINT=https://iam.cn-north-4.myhuaweicloud.com
```

###### 3.3.1.2 凭证级 [:top:](#用户手册-top)

只对单个凭证生效，会覆盖全局配置

```go
import "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"

iamEndpoint := "https://iam.cn-north-4.myhuaweicloud.com"
cred, err := basic.NewCredentialsBuilder().
            WithAk(os.Getenv("HUAWEICLOUD_SDK_AK")).
            WithSk(os.Getenv("HUAWEICLOUD_SDK_SK")).
            WithIamEndpointOverride(iamEndpoint).
            SafeBuild()
```

##### 3.3.2 Region配置 [:top:](#用户手册-top)

###### 3.3.2.1 代码配置  [:top:](#用户手册-top)

```go
import (
    ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

// 使用自定义的regionId和endpoint创建一个region
reg := region.NewRegion("cn-north-9", "https://ecs.cn-north-9.myhuaweicloud.com")

hcClient, err := ecs.EcsClientBuilder().
    WithRegion(reg).
    WithCredential(auth).
    SafeBuild()
if err != nil {
    // 处理错误
}
client := ecs.NewEcsClient(hcClient)
```

###### 3.3.2.2 环境变量 [:top:](#用户手册-top)

通过环境变量配置，格式为`HUAWEICLOUD_SDK_REGION_{SERVICE_NAME}_{REGION_ID}={endpoint}`

注：环境变量名全大写，中划线替换为下划线

```
// 以ECS和IoTDA服务为例

// linux
export HUAWEICLOUD_SDK_REGION_ECS_CN_NORTH_9=https://ecs.cn-north-9.myhuaweicloud.com
export HUAWEICLOUD_SDK_REGION_IOTDA_AP_SOUTHEAST_1=https://iotda.ap-southwest-1.myhuaweicloud.com

// windows
set HUAWEICLOUD_SDK_REGION_ECS_CN_NORTH_9=https://ecs.cn-north-9.myhuaweicloud.com
set HUAWEICLOUD_SDK_REGION_IOTDA_AP_SOUTHEAST_1=https://iotda.ap-southwest-1.myhuaweicloud.com
```

从**0.1.62**版本起，支持配置一个region对应多个endpoint，主要endpoint无法连接会自动切换到备用endpoint

格式为`HUAWEICLOUD_SDK_REGION_{SERVICE_NAME}_{REGION_ID}={endpoint1},{endpoint2}`, 多个endpoint之间用英文逗号隔开, 比如`HUAWEICLOUD_SDK_REGION_ECS_CN_NORTH_9=https://ecs.cn-north-9.myhuaweicloud.com,https://ecs.cn-north-9.myhuaweicloud.cn`

###### 3.3.2.3 文件配置 [:top:](#用户手册-top)

通过yaml文件配置，默认会从用户主目录下读取region配置文件，linux为`~/.huaweicloud/regions.yaml`，windows为`C:\Users\USER_NAME\.huaweicloud\regions.yaml`，默认配置文件可以不存在，但是如果配置文件存在且内容格式不对会解析错误抛出异常。

可以通过配置环境变量`HUAWEICLOUD_SDK_REGIONS_FILE`来修改默认文件的路径，如`HUAWEICLOUD_SDK_REGIONS_FILE=/tmp/my_regions.yml`

文件内容格式如下：

```yaml
# 服务名不区分大小写
ECS:
  - id: 'cn-north-1'
    endpoint: 'https://ecs.cn-north-1.myhuaweicloud.com'
  - id: 'cn-north-9'
    endpoint: 'https://ecs.cn-north-9.myhuaweicloud.com'
IoTDA:
  - id: 'ap-southwest-1'
    endpoint: 'https://iotda.ap-southwest-1.myhuaweicloud.com'
```

从**0.1.62**版本起，支持配置一个region对应多个endpoint，主要endpoint无法连接会自动切换到备用endpoint，格式如下：

```yaml
ECS:
  - id: 'cn-north-1'
    endpoints:
      - 'https://ecs.cn-north-1.myhuaweicloud.com'
      - 'https://ecs.cn-north-1.myhuaweicloud.cn'
```

###### 3.3.2.4 Region提供链 [:top:](#用户手册-top)

**region.ValueOf(regionId)** 默认查找顺序为 **环境变量 -> 配置文件 -> SDK中已定义Region**，以上方式都找不到region会抛出异常，获取region示例：

```go
import "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"

region1, err := region.SafeValueOf("cn-north-1")
region2, err := region.SafeValueOf("cn-north-9")
```

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

func createImageWatermark(client *dsc.DscClient) error {
    // 打开文件
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

    // 下载文件
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

### 7. 接口调用器 [:top:](#用户手册-top)

#### 7.1 自定义请求头 [:top:](#用户手册-top)

可以根据需要灵活地配置请求头域参数，非必要**请勿**指定诸如`Host`、`Authorization`、`User-Agent`、`Content-Type`等通用请求头，可能会导致接口调用错误。

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
        // 自定义请求头
        AddHeaders(map[string]string{"key1": "value1", "key2": "value2"}).
        Invoke()

    if err == nil {
        fmt.Printf("%+v\n", response)
    } else {
        fmt.Printf("%+v\n", err)
    }
}
```

#### 7.2 请求重试 [:top:](#用户手册-top)

当请求遇到网络异常或者流控场景的时候，通常需要对请求进行重试。Go SDK 提供了请求重试的入口，可用于请求方式为 `GET`
的请求。如需使用重试，需要配置最大重试次数、重试条件和重试策略。其中，

- _最大重试次数_：最大重试次数
- _重试条件_：函数，根据上一次请求的返回情况给出是否需要重试
- _重试策略_：每次重试前的等待时间，支持多种重试策略

以 VPC 服务的 `ListVpcs` 接口为例，最多重试3次，服务端返回异常时进行重试，代码如下：

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
    // 初始化客户端
    hcClient, err := vpc.VpcClientBuilder().
        WithEndpoint("<input your endpoint>").
        WithCredential(auth).
        SafeBuild()
    if err != nil {
        fmt.Println(err)
        return
    }
    client := vpc.NewVpcClient(hcClient)

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
}
```
