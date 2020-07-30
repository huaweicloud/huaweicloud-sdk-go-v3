package v1

import (
    http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/devstar/v1/model"
)

type DevstarClient struct {
    hcClient *http_client.HcHttpClient
}

func NewDevstarClient(hcClient *http_client.HcHttpClient) *DevstarClient {
    return &DevstarClient{hcClient: hcClient}
}

func DevstarClientBuilder() *http_client.HcHttpClientBuilder {
    builder := http_client.NewHcHttpClientBuilder()
    return builder
}


//通过DevStar的模板进行应用代码创建  新建任务时会返回任务ID，通过任务ID可以查看任务的状态以及最终代码生成的地址  - 接口鉴权方式 通过华为云服务获取的用户token  - 代码生成位置 应用代码生成后的地址，目前支持codehub地址和压缩包下载地址。
func (c *DevstarClient) RunTemplateJobV2(request *model.RunTemplateJobV2Request) (*model.RunTemplateJobV2Response, error) {
    requestDef := GenReqDefForRunTemplateJobV2(request)
    resp, responseDef := GenRespForRunTemplateJobV2()

    if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
        return nil, err
    } else {
        return resp, nil
    }
}

//查询任务的详情  通过任务ID可以查看任务的状态 当任务结束时返回应用代码存放的位置  - 接口鉴权方式 通过华为云服务获取的用户token  - 代码生成位置 应用代码生成后的地址，目前支持codehub地址和压缩包下载地址
func (c *DevstarClient) ShowJobDetail(request *model.ShowJobDetailRequest) (*model.ShowJobDetailResponse, error) {
    requestDef := GenReqDefForShowJobDetail(request)
    resp, responseDef := GenRespForShowJobDetail()

    if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
        return nil, err
    } else {
        return resp, nil
    }
}

//查询模板列表
func (c *DevstarClient) ListPublishedTemplates(request *model.ListPublishedTemplatesRequest) (*model.ListPublishedTemplatesResponse, error) {
    requestDef := GenReqDefForListPublishedTemplates(request)
    resp, responseDef := GenRespForListPublishedTemplates()

    if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
        return nil, err
    } else {
        return resp, nil
    }
}

//查询模板详情
func (c *DevstarClient) ShowTemplateDetail(request *model.ShowTemplateDetailRequest) (*model.ShowTemplateDetailResponse, error) {
    requestDef := GenReqDefForShowTemplateDetail(request)
    resp, responseDef := GenRespForShowTemplateDetail()

    if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
        return nil, err
    } else {
        return resp, nil
    }
}

