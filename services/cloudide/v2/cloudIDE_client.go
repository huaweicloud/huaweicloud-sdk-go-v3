package v2

import (
	http_client "code.byted.org/ti/huaweicloud-sdk-go-v3/core"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/services/cloudide/v2/model"
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

//设置ide实例对插件的授权。同意、不同意、未知（下次重新询问）
func (c *CloudIDEClient) CreateExtensionAuthorization(request *model.CreateExtensionAuthorizationRequest) (*model.CreateExtensionAuthorizationResponse, error) {
	requestDef := GenReqDefForCreateExtensionAuthorization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateExtensionAuthorizationResponse), nil
	}
}

//查询技术栈模板工程
func (c *CloudIDEClient) ListProjectTemplates(request *model.ListProjectTemplatesRequest) (*model.ListProjectTemplatesResponse, error) {
	requestDef := GenReqDefForListProjectTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectTemplatesResponse), nil
	}
}

//按region获取标签所有技术栈
func (c *CloudIDEClient) ListStacks(request *model.ListStacksRequest) (*model.ListStacksResponse, error) {
	requestDef := GenReqDefForListStacks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStacksResponse), nil
	}
}

//查询当前帐号访问权限
func (c *CloudIDEClient) ShowAccountStatus(request *model.ShowAccountStatusRequest) (*model.ShowAccountStatusResponse, error) {
	requestDef := GenReqDefForShowAccountStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAccountStatusResponse), nil
	}
}

//查询ide实例对插件的授权情况，同意授权的插件能在ide实例内携带登陆用户的token调用第三方服务
func (c *CloudIDEClient) ShowExtensionAuthorization(request *model.ShowExtensionAuthorizationRequest) (*model.ShowExtensionAuthorizationResponse, error) {
	requestDef := GenReqDefForShowExtensionAuthorization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowExtensionAuthorizationResponse), nil
	}
}

//获取技术栈计费信息
func (c *CloudIDEClient) ShowPrice(request *model.ShowPriceRequest) (*model.ShowPriceResponse, error) {
	requestDef := GenReqDefForShowPrice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPriceResponse), nil
	}
}

//查询用户是否有权限访问某个IDE实例
func (c *CloudIDEClient) CheckInstanceAccess(request *model.CheckInstanceAccessRequest) (*model.CheckInstanceAccessResponse, error) {
	requestDef := GenReqDefForCheckInstanceAccess()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckInstanceAccessResponse), nil
	}
}

//查询IDE实例名是否重复
func (c *CloudIDEClient) CheckName(request *model.CheckNameRequest) (*model.CheckNameResponse, error) {
	requestDef := GenReqDefForCheckName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckNameResponse), nil
	}
}

//创建IDE实例
func (c *CloudIDEClient) CreateInstance(request *model.CreateInstanceRequest) (*model.CreateInstanceResponse, error) {
	requestDef := GenReqDefForCreateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceResponse), nil
	}
}

//创建IDE实例
func (c *CloudIDEClient) CreateInstanceBy3rd(request *model.CreateInstanceBy3rdRequest) (*model.CreateInstanceBy3rdResponse, error) {
	requestDef := GenReqDefForCreateInstanceBy3rd()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceBy3rdResponse), nil
	}
}

//删除IDE实例（同时删除磁盘数据）
func (c *CloudIDEClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResponse), nil
	}
}

//查询IDE实例列表
func (c *CloudIDEClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

//查询某个租户下的IDE实例列表
func (c *CloudIDEClient) ListOrgInstances(request *model.ListOrgInstancesRequest) (*model.ListOrgInstancesResponse, error) {
	requestDef := GenReqDefForListOrgInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOrgInstancesResponse), nil
	}
}

//查询某个IDE实例
func (c *CloudIDEClient) ShowInstance(request *model.ShowInstanceRequest) (*model.ShowInstanceResponse, error) {
	requestDef := GenReqDefForShowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceResponse), nil
	}
}

//启动IDE实例
func (c *CloudIDEClient) StartInstance(request *model.StartInstanceRequest) (*model.StartInstanceResponse, error) {
	requestDef := GenReqDefForStartInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartInstanceResponse), nil
	}
}

//停止IDE实例（不删除磁盘数据）
func (c *CloudIDEClient) StopInstance(request *model.StopInstanceRequest) (*model.StopInstanceResponse, error) {
	requestDef := GenReqDefForStopInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopInstanceResponse), nil
	}
}

//修改IDE实例
func (c *CloudIDEClient) UpdateInstance(request *model.UpdateInstanceRequest) (*model.UpdateInstanceResponse, error) {
	requestDef := GenReqDefForUpdateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceResponse), nil
	}
}

//刷新IDE实例活跃状态，超过该实例设置的过期时间后实例自动关闭。
func (c *CloudIDEClient) UpdateInstanceActivity(request *model.UpdateInstanceActivityRequest) (*model.UpdateInstanceActivityResponse, error) {
	requestDef := GenReqDefForUpdateInstanceActivity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceActivityResponse), nil
	}
}
