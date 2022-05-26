package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cloudide/v2/model"
)

type CloudIDEClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCloudIDEClient(hcClient *http_client.HcHttpClient) *CloudIDEClient {
	return &CloudIDEClient{HcClient: hcClient}
}

func CloudIDEClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CreateExtensionAuthorization 设置ide实例对插件的授权
//
// 设置ide实例对插件的授权。同意、不同意、未知（下次重新询问）
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudIDEClient) CreateExtensionAuthorization(request *model.CreateExtensionAuthorizationRequest) (*model.CreateExtensionAuthorizationResponse, error) {
	requestDef := GenReqDefForCreateExtensionAuthorization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateExtensionAuthorizationResponse), nil
	}
}

// CreateExtensionAuthorizationInvoker 设置ide实例对插件的授权
func (c *CloudIDEClient) CreateExtensionAuthorizationInvoker(request *model.CreateExtensionAuthorizationRequest) *CreateExtensionAuthorizationInvoker {
	requestDef := GenReqDefForCreateExtensionAuthorization()
	return &CreateExtensionAuthorizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectTemplates 查询技术栈模板工程
//
// 查询技术栈模板工程
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudIDEClient) ListProjectTemplates(request *model.ListProjectTemplatesRequest) (*model.ListProjectTemplatesResponse, error) {
	requestDef := GenReqDefForListProjectTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectTemplatesResponse), nil
	}
}

// ListProjectTemplatesInvoker 查询技术栈模板工程
func (c *CloudIDEClient) ListProjectTemplatesInvoker(request *model.ListProjectTemplatesRequest) *ListProjectTemplatesInvoker {
	requestDef := GenReqDefForListProjectTemplates()
	return &ListProjectTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStacks 按region获取标签所有技术栈
//
// 按region获取标签所有技术栈
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudIDEClient) ListStacks(request *model.ListStacksRequest) (*model.ListStacksResponse, error) {
	requestDef := GenReqDefForListStacks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStacksResponse), nil
	}
}

// ListStacksInvoker 按region获取标签所有技术栈
func (c *CloudIDEClient) ListStacksInvoker(request *model.ListStacksRequest) *ListStacksInvoker {
	requestDef := GenReqDefForListStacks()
	return &ListStacksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAccountStatus 查询当前帐号访问权限
//
// 查询当前帐号访问权限
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudIDEClient) ShowAccountStatus(request *model.ShowAccountStatusRequest) (*model.ShowAccountStatusResponse, error) {
	requestDef := GenReqDefForShowAccountStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAccountStatusResponse), nil
	}
}

// ShowAccountStatusInvoker 查询当前帐号访问权限
func (c *CloudIDEClient) ShowAccountStatusInvoker(request *model.ShowAccountStatusRequest) *ShowAccountStatusInvoker {
	requestDef := GenReqDefForShowAccountStatus()
	return &ShowAccountStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowExtensionAuthorization 查询ide实例对插件的授权情况
//
// 查询ide实例对插件的授权情况，同意授权的插件能在ide实例内携带登陆用户的token调用第三方服务
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudIDEClient) ShowExtensionAuthorization(request *model.ShowExtensionAuthorizationRequest) (*model.ShowExtensionAuthorizationResponse, error) {
	requestDef := GenReqDefForShowExtensionAuthorization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowExtensionAuthorizationResponse), nil
	}
}

// ShowExtensionAuthorizationInvoker 查询ide实例对插件的授权情况
func (c *CloudIDEClient) ShowExtensionAuthorizationInvoker(request *model.ShowExtensionAuthorizationRequest) *ShowExtensionAuthorizationInvoker {
	requestDef := GenReqDefForShowExtensionAuthorization()
	return &ShowExtensionAuthorizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPrice 获取技术栈计费信息
//
// 获取技术栈计费信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudIDEClient) ShowPrice(request *model.ShowPriceRequest) (*model.ShowPriceResponse, error) {
	requestDef := GenReqDefForShowPrice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPriceResponse), nil
	}
}

// ShowPriceInvoker 获取技术栈计费信息
func (c *CloudIDEClient) ShowPriceInvoker(request *model.ShowPriceRequest) *ShowPriceInvoker {
	requestDef := GenReqDefForShowPrice()
	return &ShowPriceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckInstanceAccess 查询用户是否有权限访问某个IDE实例
//
// 查询用户是否有权限访问某个IDE实例
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudIDEClient) CheckInstanceAccess(request *model.CheckInstanceAccessRequest) (*model.CheckInstanceAccessResponse, error) {
	requestDef := GenReqDefForCheckInstanceAccess()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckInstanceAccessResponse), nil
	}
}

// CheckInstanceAccessInvoker 查询用户是否有权限访问某个IDE实例
func (c *CloudIDEClient) CheckInstanceAccessInvoker(request *model.CheckInstanceAccessRequest) *CheckInstanceAccessInvoker {
	requestDef := GenReqDefForCheckInstanceAccess()
	return &CheckInstanceAccessInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckName 查询IDE实例名是否重复
//
// 查询IDE实例名是否重复
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudIDEClient) CheckName(request *model.CheckNameRequest) (*model.CheckNameResponse, error) {
	requestDef := GenReqDefForCheckName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckNameResponse), nil
	}
}

// CheckNameInvoker 查询IDE实例名是否重复
func (c *CloudIDEClient) CheckNameInvoker(request *model.CheckNameRequest) *CheckNameInvoker {
	requestDef := GenReqDefForCheckName()
	return &CheckNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstance 创建IDE实例
//
// 创建IDE实例
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudIDEClient) CreateInstance(request *model.CreateInstanceRequest) (*model.CreateInstanceResponse, error) {
	requestDef := GenReqDefForCreateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceResponse), nil
	}
}

// CreateInstanceInvoker 创建IDE实例
func (c *CloudIDEClient) CreateInstanceInvoker(request *model.CreateInstanceRequest) *CreateInstanceInvoker {
	requestDef := GenReqDefForCreateInstance()
	return &CreateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstanceBy3rd 外部第三方集成商创建IDE实例
//
// 创建IDE实例
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudIDEClient) CreateInstanceBy3rd(request *model.CreateInstanceBy3rdRequest) (*model.CreateInstanceBy3rdResponse, error) {
	requestDef := GenReqDefForCreateInstanceBy3rd()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceBy3rdResponse), nil
	}
}

// CreateInstanceBy3rdInvoker 外部第三方集成商创建IDE实例
func (c *CloudIDEClient) CreateInstanceBy3rdInvoker(request *model.CreateInstanceBy3rdRequest) *CreateInstanceBy3rdInvoker {
	requestDef := GenReqDefForCreateInstanceBy3rd()
	return &CreateInstanceBy3rdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstance 删除IDE实例
//
// 删除IDE实例（同时删除磁盘数据）
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudIDEClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResponse), nil
	}
}

// DeleteInstanceInvoker 删除IDE实例
func (c *CloudIDEClient) DeleteInstanceInvoker(request *model.DeleteInstanceRequest) *DeleteInstanceInvoker {
	requestDef := GenReqDefForDeleteInstance()
	return &DeleteInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstances 查询IDE实例列表
//
// 查询IDE实例列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudIDEClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

// ListInstancesInvoker 查询IDE实例列表
func (c *CloudIDEClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOrgInstances 查询某个租户下的IDE实例列表
//
// 查询某个租户下的IDE实例列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudIDEClient) ListOrgInstances(request *model.ListOrgInstancesRequest) (*model.ListOrgInstancesResponse, error) {
	requestDef := GenReqDefForListOrgInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOrgInstancesResponse), nil
	}
}

// ListOrgInstancesInvoker 查询某个租户下的IDE实例列表
func (c *CloudIDEClient) ListOrgInstancesInvoker(request *model.ListOrgInstancesRequest) *ListOrgInstancesInvoker {
	requestDef := GenReqDefForListOrgInstances()
	return &ListOrgInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstance 查询某个IDE实例
//
// 查询某个IDE实例
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudIDEClient) ShowInstance(request *model.ShowInstanceRequest) (*model.ShowInstanceResponse, error) {
	requestDef := GenReqDefForShowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceResponse), nil
	}
}

// ShowInstanceInvoker 查询某个IDE实例
func (c *CloudIDEClient) ShowInstanceInvoker(request *model.ShowInstanceRequest) *ShowInstanceInvoker {
	requestDef := GenReqDefForShowInstance()
	return &ShowInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartInstance 启动IDE实例
//
// 启动IDE实例
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudIDEClient) StartInstance(request *model.StartInstanceRequest) (*model.StartInstanceResponse, error) {
	requestDef := GenReqDefForStartInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartInstanceResponse), nil
	}
}

// StartInstanceInvoker 启动IDE实例
func (c *CloudIDEClient) StartInstanceInvoker(request *model.StartInstanceRequest) *StartInstanceInvoker {
	requestDef := GenReqDefForStartInstance()
	return &StartInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopInstance 停止IDE实例
//
// 停止IDE实例（不删除磁盘数据）
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudIDEClient) StopInstance(request *model.StopInstanceRequest) (*model.StopInstanceResponse, error) {
	requestDef := GenReqDefForStopInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopInstanceResponse), nil
	}
}

// StopInstanceInvoker 停止IDE实例
func (c *CloudIDEClient) StopInstanceInvoker(request *model.StopInstanceRequest) *StopInstanceInvoker {
	requestDef := GenReqDefForStopInstance()
	return &StopInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstance 修改IDE实例
//
// 修改IDE实例
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudIDEClient) UpdateInstance(request *model.UpdateInstanceRequest) (*model.UpdateInstanceResponse, error) {
	requestDef := GenReqDefForUpdateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceResponse), nil
	}
}

// UpdateInstanceInvoker 修改IDE实例
func (c *CloudIDEClient) UpdateInstanceInvoker(request *model.UpdateInstanceRequest) *UpdateInstanceInvoker {
	requestDef := GenReqDefForUpdateInstance()
	return &UpdateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceActivity 刷新IDE实例活跃状态
//
// 刷新IDE实例活跃状态，超过该实例设置的过期时间后实例自动关闭。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudIDEClient) UpdateInstanceActivity(request *model.UpdateInstanceActivityRequest) (*model.UpdateInstanceActivityResponse, error) {
	requestDef := GenReqDefForUpdateInstanceActivity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceActivityResponse), nil
	}
}

// UpdateInstanceActivityInvoker 刷新IDE实例活跃状态
func (c *CloudIDEClient) UpdateInstanceActivityInvoker(request *model.UpdateInstanceActivityRequest) *UpdateInstanceActivityInvoker {
	requestDef := GenReqDefForUpdateInstanceActivity()
	return &UpdateInstanceActivityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
