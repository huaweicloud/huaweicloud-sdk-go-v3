package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eps/v1/model"
)

type EpsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewEpsClient(hcClient *http_client.HcHttpClient) *EpsClient {
	return &EpsClient{HcClient: hcClient}
}

func EpsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// CreateEnterpriseProject 创建企业项目
//
// 创建企业项目。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EpsClient) CreateEnterpriseProject(request *model.CreateEnterpriseProjectRequest) (*model.CreateEnterpriseProjectResponse, error) {
	requestDef := GenReqDefForCreateEnterpriseProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEnterpriseProjectResponse), nil
	}
}

// CreateEnterpriseProjectInvoker 创建企业项目
func (c *EpsClient) CreateEnterpriseProjectInvoker(request *model.CreateEnterpriseProjectRequest) *CreateEnterpriseProjectInvoker {
	requestDef := GenReqDefForCreateEnterpriseProject()
	return &CreateEnterpriseProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableEnterpriseProject 停用企业项目
//
// 停用企业项目。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EpsClient) DisableEnterpriseProject(request *model.DisableEnterpriseProjectRequest) (*model.DisableEnterpriseProjectResponse, error) {
	requestDef := GenReqDefForDisableEnterpriseProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableEnterpriseProjectResponse), nil
	}
}

// DisableEnterpriseProjectInvoker 停用企业项目
func (c *EpsClient) DisableEnterpriseProjectInvoker(request *model.DisableEnterpriseProjectRequest) *DisableEnterpriseProjectInvoker {
	requestDef := GenReqDefForDisableEnterpriseProject()
	return &DisableEnterpriseProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableEnterpriseProject 启用企业项目
//
// 启用企业项目。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EpsClient) EnableEnterpriseProject(request *model.EnableEnterpriseProjectRequest) (*model.EnableEnterpriseProjectResponse, error) {
	requestDef := GenReqDefForEnableEnterpriseProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableEnterpriseProjectResponse), nil
	}
}

// EnableEnterpriseProjectInvoker 启用企业项目
func (c *EpsClient) EnableEnterpriseProjectInvoker(request *model.EnableEnterpriseProjectRequest) *EnableEnterpriseProjectInvoker {
	requestDef := GenReqDefForEnableEnterpriseProject()
	return &EnableEnterpriseProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiVersions 查询API版本列表
//
// 查询企业项目的API版本列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EpsClient) ListApiVersions(request *model.ListApiVersionsRequest) (*model.ListApiVersionsResponse, error) {
	requestDef := GenReqDefForListApiVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionsResponse), nil
	}
}

// ListApiVersionsInvoker 查询API版本列表
func (c *EpsClient) ListApiVersionsInvoker(request *model.ListApiVersionsRequest) *ListApiVersionsInvoker {
	requestDef := GenReqDefForListApiVersions()
	return &ListApiVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnterpriseProject 查询企业项目列表
//
// 查询当前用户已授权的企业项目列表，用户可以使用企业项目绑定资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EpsClient) ListEnterpriseProject(request *model.ListEnterpriseProjectRequest) (*model.ListEnterpriseProjectResponse, error) {
	requestDef := GenReqDefForListEnterpriseProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnterpriseProjectResponse), nil
	}
}

// ListEnterpriseProjectInvoker 查询企业项目列表
func (c *EpsClient) ListEnterpriseProjectInvoker(request *model.ListEnterpriseProjectRequest) *ListEnterpriseProjectInvoker {
	requestDef := GenReqDefForListEnterpriseProject()
	return &ListEnterpriseProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProviders 查询企业项目支持的服务
//
// 查询企业项目支持的服务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EpsClient) ListProviders(request *model.ListProvidersRequest) (*model.ListProvidersResponse, error) {
	requestDef := GenReqDefForListProviders()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProvidersResponse), nil
	}
}

// ListProvidersInvoker 查询企业项目支持的服务
func (c *EpsClient) ListProvidersInvoker(request *model.ListProvidersRequest) *ListProvidersInvoker {
	requestDef := GenReqDefForListProviders()
	return &ListProvidersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// MigrateResource 迁移资源
//
// 迁移资源到目标企业项目。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EpsClient) MigrateResource(request *model.MigrateResourceRequest) (*model.MigrateResourceResponse, error) {
	requestDef := GenReqDefForMigrateResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MigrateResourceResponse), nil
	}
}

// MigrateResourceInvoker 迁移资源
func (c *EpsClient) MigrateResourceInvoker(request *model.MigrateResourceRequest) *MigrateResourceInvoker {
	requestDef := GenReqDefForMigrateResource()
	return &MigrateResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApiVersion 查询API版本号详情
//
// 查询指定的企业项目API版本号详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EpsClient) ShowApiVersion(request *model.ShowApiVersionRequest) (*model.ShowApiVersionResponse, error) {
	requestDef := GenReqDefForShowApiVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApiVersionResponse), nil
	}
}

// ShowApiVersionInvoker 查询API版本号详情
func (c *EpsClient) ShowApiVersionInvoker(request *model.ShowApiVersionRequest) *ShowApiVersionInvoker {
	requestDef := GenReqDefForShowApiVersion()
	return &ShowApiVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEnterpriseProject 查询企业项目详情
//
// 查询企业项目详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EpsClient) ShowEnterpriseProject(request *model.ShowEnterpriseProjectRequest) (*model.ShowEnterpriseProjectResponse, error) {
	requestDef := GenReqDefForShowEnterpriseProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEnterpriseProjectResponse), nil
	}
}

// ShowEnterpriseProjectInvoker 查询企业项目详情
func (c *EpsClient) ShowEnterpriseProjectInvoker(request *model.ShowEnterpriseProjectRequest) *ShowEnterpriseProjectInvoker {
	requestDef := GenReqDefForShowEnterpriseProject()
	return &ShowEnterpriseProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEnterpriseProjectQuota 查询企业项目配额
//
// 查询企业项目的配额信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EpsClient) ShowEnterpriseProjectQuota(request *model.ShowEnterpriseProjectQuotaRequest) (*model.ShowEnterpriseProjectQuotaResponse, error) {
	requestDef := GenReqDefForShowEnterpriseProjectQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEnterpriseProjectQuotaResponse), nil
	}
}

// ShowEnterpriseProjectQuotaInvoker 查询企业项目配额
func (c *EpsClient) ShowEnterpriseProjectQuotaInvoker(request *model.ShowEnterpriseProjectQuotaRequest) *ShowEnterpriseProjectQuotaInvoker {
	requestDef := GenReqDefForShowEnterpriseProjectQuota()
	return &ShowEnterpriseProjectQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceBindEnterpriseProject 查询企业项目绑定的资源列表
//
// 查询企业项目下绑定的资源详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EpsClient) ShowResourceBindEnterpriseProject(request *model.ShowResourceBindEnterpriseProjectRequest) (*model.ShowResourceBindEnterpriseProjectResponse, error) {
	requestDef := GenReqDefForShowResourceBindEnterpriseProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceBindEnterpriseProjectResponse), nil
	}
}

// ShowResourceBindEnterpriseProjectInvoker 查询企业项目绑定的资源列表
func (c *EpsClient) ShowResourceBindEnterpriseProjectInvoker(request *model.ShowResourceBindEnterpriseProjectRequest) *ShowResourceBindEnterpriseProjectInvoker {
	requestDef := GenReqDefForShowResourceBindEnterpriseProject()
	return &ShowResourceBindEnterpriseProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEnterpriseProject 修改企业项目
//
// 修改企业项目。当前仅支持修改名称和描述。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EpsClient) UpdateEnterpriseProject(request *model.UpdateEnterpriseProjectRequest) (*model.UpdateEnterpriseProjectResponse, error) {
	requestDef := GenReqDefForUpdateEnterpriseProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEnterpriseProjectResponse), nil
	}
}

// UpdateEnterpriseProjectInvoker 修改企业项目
func (c *EpsClient) UpdateEnterpriseProjectInvoker(request *model.UpdateEnterpriseProjectRequest) *UpdateEnterpriseProjectInvoker {
	requestDef := GenReqDefForUpdateEnterpriseProject()
	return &UpdateEnterpriseProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
