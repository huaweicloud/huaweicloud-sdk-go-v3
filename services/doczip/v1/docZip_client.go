package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/doczip/v1/model"
)

type DocZipClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewDocZipClient(hcClient *httpclient.HcHttpClient) *DocZipClient {
	return &DocZipClient{HcClient: hcClient}
}

func DocZipClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateApplication 创建应用
//
// 租户可通过该接口创建应用，每个租户允许创建应用个数可以通过“获取允许创建的资源数量”接口获取。其中租户ID就是用户的账号ID。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DocZipClient) CreateApplication(request *model.CreateApplicationRequest) (*model.CreateApplicationResponse, error) {
	requestDef := GenReqDefForCreateApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApplicationResponse), nil
	}
}

// CreateApplicationInvoker 创建应用
func (c *DocZipClient) CreateApplicationInvoker(request *model.CreateApplicationRequest) *CreateApplicationInvoker {
	requestDef := GenReqDefForCreateApplication()
	return &CreateApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApplication 删除应用
//
// 租户可通过该接口删除应用，应用被删除后，将不能被恢复。其中租户ID就是用户的账号ID。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DocZipClient) DeleteApplication(request *model.DeleteApplicationRequest) (*model.DeleteApplicationResponse, error) {
	requestDef := GenReqDefForDeleteApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApplicationResponse), nil
	}
}

// DeleteApplicationInvoker 删除应用
func (c *DocZipClient) DeleteApplicationInvoker(request *model.DeleteApplicationRequest) *DeleteApplicationInvoker {
	requestDef := GenReqDefForDeleteApplication()
	return &DeleteApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApplicationDetails 获取应用详情
//
// 租户通过该接口可以查询某个租户的详细信息，配合query_type字段查询应用的基本信息，系统配置信息，文档块信息，也可以查询全部信息。其中租户ID就是用户的账号ID。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DocZipClient) ListApplicationDetails(request *model.ListApplicationDetailsRequest) (*model.ListApplicationDetailsResponse, error) {
	requestDef := GenReqDefForListApplicationDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationDetailsResponse), nil
	}
}

// ListApplicationDetailsInvoker 获取应用详情
func (c *DocZipClient) ListApplicationDetailsInvoker(request *model.ListApplicationDetailsRequest) *ListApplicationDetailsInvoker {
	requestDef := GenReqDefForListApplicationDetails()
	return &ListApplicationDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApplications 获取应用列表
//
// 租户可通过该接口获取租户下所有应用的信息列表。其中租户ID就是用户的账号ID。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DocZipClient) ListApplications(request *model.ListApplicationsRequest) (*model.ListApplicationsResponse, error) {
	requestDef := GenReqDefForListApplications()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationsResponse), nil
	}
}

// ListApplicationsInvoker 获取应用列表
func (c *DocZipClient) ListApplicationsInvoker(request *model.ListApplicationsRequest) *ListApplicationsInvoker {
	requestDef := GenReqDefForListApplications()
	return &ListApplicationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PublishApplication 发布应用
//
// 租户可通过该接口发布应用，应用发布后，才可以被使用。其中租户ID就是用户的账号ID。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DocZipClient) PublishApplication(request *model.PublishApplicationRequest) (*model.PublishApplicationResponse, error) {
	requestDef := GenReqDefForPublishApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PublishApplicationResponse), nil
	}
}

// PublishApplicationInvoker 发布应用
func (c *DocZipClient) PublishApplicationInvoker(request *model.PublishApplicationRequest) *PublishApplicationInvoker {
	requestDef := GenReqDefForPublishApplication()
	return &PublishApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAppResourceLimit 获取允许创建的应用上限
//
// 租户可通过该接口获取允许创建的应用数量上限，每个租户允许创建的租户数量不能超过配置的被允许创建的应用数量。其中租户ID就是用户的账号ID。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DocZipClient) ShowAppResourceLimit(request *model.ShowAppResourceLimitRequest) (*model.ShowAppResourceLimitResponse, error) {
	requestDef := GenReqDefForShowAppResourceLimit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppResourceLimitResponse), nil
	}
}

// ShowAppResourceLimitInvoker 获取允许创建的应用上限
func (c *DocZipClient) ShowAppResourceLimitInvoker(request *model.ShowAppResourceLimitRequest) *ShowAppResourceLimitInvoker {
	requestDef := GenReqDefForShowAppResourceLimit()
	return &ShowAppResourceLimitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAppSecretKey 获取应用密钥
//
// 租户可通过该接口获取应用密钥。租户可通过application_id和应用密钥完成租户认证。其中租户ID就是用户的账号ID。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DocZipClient) ShowAppSecretKey(request *model.ShowAppSecretKeyRequest) (*model.ShowAppSecretKeyResponse, error) {
	requestDef := GenReqDefForShowAppSecretKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppSecretKeyResponse), nil
	}
}

// ShowAppSecretKeyInvoker 获取应用密钥
func (c *DocZipClient) ShowAppSecretKeyInvoker(request *model.ShowAppSecretKeyRequest) *ShowAppSecretKeyInvoker {
	requestDef := GenReqDefForShowAppSecretKey()
	return &ShowAppSecretKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBlocks 获取应用文档块信息
//
// 租户可通过该接口获取当前租户下支持的所有文档块信息，在“创建应用”接口中，可根据业务需求为不同的应用选择不同的文档块数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DocZipClient) ShowBlocks(request *model.ShowBlocksRequest) (*model.ShowBlocksResponse, error) {
	requestDef := GenReqDefForShowBlocks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBlocksResponse), nil
	}
}

// ShowBlocksInvoker 获取应用文档块信息
func (c *DocZipClient) ShowBlocksInvoker(request *model.ShowBlocksRequest) *ShowBlocksInvoker {
	requestDef := GenReqDefForShowBlocks()
	return &ShowBlocksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateApplication 修改应用
//
// 租户可通过该接口修改某个应用的名称，系统设置等信息。其中租户ID就是用户的账号ID。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DocZipClient) UpdateApplication(request *model.UpdateApplicationRequest) (*model.UpdateApplicationResponse, error) {
	requestDef := GenReqDefForUpdateApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateApplicationResponse), nil
	}
}

// UpdateApplicationInvoker 修改应用
func (c *DocZipClient) UpdateApplicationInvoker(request *model.UpdateApplicationRequest) *UpdateApplicationInvoker {
	requestDef := GenReqDefForUpdateApplication()
	return &UpdateApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
