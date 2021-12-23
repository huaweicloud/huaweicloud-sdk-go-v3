package v1

import (
	http_client "github.com/RandolphCYG/huaweicloud-sdk-go-v3/core"

	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/services/devstar/v1/model"
)

type DevStarClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDevStarClient(hcClient *http_client.HcHttpClient) *DevStarClient {
	return &DevStarClient{HcClient: hcClient}
}

func DevStarClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

//下载模板产物。
func (c *DevStarClient) DownloadApplicationCode(request *model.DownloadApplicationCodeRequest) (*model.DownloadApplicationCodeResponse, error) {
	requestDef := GenReqDefForDownloadApplicationCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadApplicationCodeResponse), nil
	}
}

//使用CodeHub模板创建应用代码。  通过 Codehub 模板创建生成应用代码的任务，并将应用代码存储于指定的 CodeHub 仓库中或者生成代码压缩包，可以通过返回的任务 ID 查询相关任务状态。  - 接口鉴权方式 通过华为云服务获取的用户token。  - 代码生成位置 应用代码生成后的地址，目前支持codehub地址和压缩包下载地址。
func (c *DevStarClient) RunCodehubTemplateJob(request *model.RunCodehubTemplateJobRequest) (*model.RunCodehubTemplateJobResponse, error) {
	requestDef := GenReqDefForRunCodehubTemplateJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunCodehubTemplateJobResponse), nil
	}
}

//使用DevStar的模板创建应用代码。  通过 DevStar 模板创建生成应用代码的任务，并将应用代码存储于指定的 CodeHub 仓库中，可以通过返回的任务 ID 查询相关任务状态。  - 接口鉴权方式 通过华为云服务获取的用户token。  - 代码生成位置 应用代码生成后的地址，目前支持codehub地址和压缩包下载地址。
func (c *DevStarClient) RunDevstarTemplateJob(request *model.RunDevstarTemplateJobRequest) (*model.RunDevstarTemplateJobResponse, error) {
	requestDef := GenReqDefForRunDevstarTemplateJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunDevstarTemplateJobResponse), nil
	}
}

//查询任务的详情。  通过任务ID可以查看任务的状态 当任务结束时返回应用代码存放的位置。  - 接口鉴权方式 通过华为云服务获取的用户token。  - 代码生成位置 应用代码生成后的地址，目前支持codehub地址和压缩包下载地址。
func (c *DevStarClient) ShowJobDetail(request *model.ShowJobDetailRequest) (*model.ShowJobDetailResponse, error) {
	requestDef := GenReqDefForShowJobDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobDetailResponse), nil
	}
}

//该接口可以用于模板作者或模板维护人读取模板文件内容。
func (c *DevStarClient) ShowTemplateFile(request *model.ShowTemplateFileRequest) (*model.ShowTemplateFileResponse, error) {
	requestDef := GenReqDefForShowTemplateFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTemplateFileResponse), nil
	}
}

//未登录状态下，将用户浏览过的模板缓存在浏览器中，登录时，调用该接口同步模板浏览记录。
func (c *DevStarClient) CreateTemplateViewHistories(request *model.CreateTemplateViewHistoriesRequest) (*model.CreateTemplateViewHistoriesResponse, error) {
	requestDef := GenReqDefForCreateTemplateViewHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTemplateViewHistoriesResponse), nil
	}
}

//查询模板列表，推荐使用V2版本接口。
func (c *DevStarClient) ListPublishedTemplates(request *model.ListPublishedTemplatesRequest) (*model.ListPublishedTemplatesResponse, error) {
	requestDef := GenReqDefForListPublishedTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPublishedTemplatesResponse), nil
	}
}

//查询DevStar或者CodeLabs登录用户浏览过的模板（只返回最近浏览的5个模板）。
func (c *DevStarClient) ListTemplateViewHistories(request *model.ListTemplateViewHistoriesRequest) (*model.ListTemplateViewHistoriesResponse, error) {
	requestDef := GenReqDefForListTemplateViewHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplateViewHistoriesResponse), nil
	}
}

//查询模板列表。
func (c *DevStarClient) ListTemplatesV2(request *model.ListTemplatesV2Request) (*model.ListTemplatesV2Response, error) {
	requestDef := GenReqDefForListTemplatesV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplatesV2Response), nil
	}
}

//获取指定模板详情，包括模板id、名称、描述、作者、标签、上架时间等信息。
func (c *DevStarClient) ShowTemplateV3(request *model.ShowTemplateV3Request) (*model.ShowTemplateV3Response, error) {
	requestDef := GenReqDefForShowTemplateV3()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTemplateV3Response), nil
	}
}

//查询模板详情，推荐使用V3版本接口。
func (c *DevStarClient) ShowTemplateDetail(request *model.ShowTemplateDetailRequest) (*model.ShowTemplateDetailResponse, error) {
	requestDef := GenReqDefForShowTemplateDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTemplateDetailResponse), nil
	}
}
