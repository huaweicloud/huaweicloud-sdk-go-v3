package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cloudide/v2/model"
)

type CloudideClient struct {
	hcClient *http_client.HcHttpClient
}

func NewCloudideClient(hcClient *http_client.HcHttpClient) *CloudideClient {
	return &CloudideClient{hcClient: hcClient}
}

func CloudideClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//查询技术栈模板工程
func (c *CloudideClient) ListProjectTemplates(request *model.ListProjectTemplatesRequest) (*model.ListProjectTemplatesResponse, error) {
	requestDef := GenReqDefForListProjectTemplates()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectTemplatesResponse), nil
	}
}

//获取标签所有技术栈
func (c *CloudideClient) ListStacksByTag(request *model.ListStacksByTagRequest) (*model.ListStacksByTagResponse, error) {
	requestDef := GenReqDefForListStacksByTag()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStacksByTagResponse), nil
	}
}

//获取技术栈计费信息
func (c *CloudideClient) ShowPrice(request *model.ShowPriceRequest) (*model.ShowPriceResponse, error) {
	requestDef := GenReqDefForShowPrice()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPriceResponse), nil
	}
}

//查询IDE实例名是否重复
func (c *CloudideClient) CheckName(request *model.CheckNameRequest) (*model.CheckNameResponse, error) {
	requestDef := GenReqDefForCheckName()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckNameResponse), nil
	}
}

//创建IDE实例
func (c *CloudideClient) CreateInstance(request *model.CreateInstanceRequest) (*model.CreateInstanceResponse, error) {
	requestDef := GenReqDefForCreateInstance()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceResponse), nil
	}
}

//创建IDE实例
func (c *CloudideClient) CreateInstanceBy3rd(request *model.CreateInstanceBy3rdRequest) (*model.CreateInstanceBy3rdResponse, error) {
	requestDef := GenReqDefForCreateInstanceBy3rd()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceBy3rdResponse), nil
	}
}

//删除IDE实例（同时删除磁盘数据）
func (c *CloudideClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResponse), nil
	}
}

//查询IDE实例列表
func (c *CloudideClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

//查询某个组织下的IDE实例列表
func (c *CloudideClient) ListOrgInstances(request *model.ListOrgInstancesRequest) (*model.ListOrgInstancesResponse, error) {
	requestDef := GenReqDefForListOrgInstances()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOrgInstancesResponse), nil
	}
}

//查询某个IDE实例
func (c *CloudideClient) ShowInstance(request *model.ShowInstanceRequest) (*model.ShowInstanceResponse, error) {
	requestDef := GenReqDefForShowInstance()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceResponse), nil
	}
}

//启动IDE实例
func (c *CloudideClient) StartInstance(request *model.StartInstanceRequest) (*model.StartInstanceResponse, error) {
	requestDef := GenReqDefForStartInstance()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartInstanceResponse), nil
	}
}

//停止IDE实例（不删除磁盘数据）
func (c *CloudideClient) StopInstance(request *model.StopInstanceRequest) (*model.StopInstanceResponse, error) {
	requestDef := GenReqDefForStopInstance()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopInstanceResponse), nil
	}
}

//修改IDE实例
func (c *CloudideClient) UpdateInstance(request *model.UpdateInstanceRequest) (*model.UpdateInstanceResponse, error) {
	requestDef := GenReqDefForUpdateInstance()

	if resp, err := c.hcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceResponse), nil
	}
}
