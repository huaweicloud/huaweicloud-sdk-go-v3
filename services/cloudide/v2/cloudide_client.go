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
	requestDef := GenReqDefForListProjectTemplates(request)
	resp, responseDef := GenRespForListProjectTemplates()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//获取标签所有技术栈
func (c *CloudideClient) ListStacksByTag(request *model.ListStacksByTagRequest) (*model.ListStacksByTagResponse, error) {
	requestDef := GenReqDefForListStacksByTag(request)
	resp, responseDef := GenRespForListStacksByTag()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//获取技术栈计费信息
func (c *CloudideClient) ShowPrice(request *model.ShowPriceRequest) (*model.ShowPriceResponse, error) {
	requestDef := GenReqDefForShowPrice(request)
	resp, responseDef := GenRespForShowPrice()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询IDE实例名是否重复
func (c *CloudideClient) CheckName(request *model.CheckNameRequest) (*model.CheckNameResponse, error) {
	requestDef := GenReqDefForCheckName(request)
	resp, responseDef := GenRespForCheckName()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//创建IDE实例
func (c *CloudideClient) CreateInstance(request *model.CreateInstanceRequest) (*model.CreateInstanceResponse, error) {
	requestDef := GenReqDefForCreateInstance(request)
	resp, responseDef := GenRespForCreateInstance()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//创建IDE实例
func (c *CloudideClient) CreateInstanceBy3rd(request *model.CreateInstanceBy3rdRequest) (*model.CreateInstanceBy3rdResponse, error) {
	requestDef := GenReqDefForCreateInstanceBy3rd(request)
	resp, responseDef := GenRespForCreateInstanceBy3rd()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除IDE实例（同时删除磁盘数据）
func (c *CloudideClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance(request)
	resp, responseDef := GenRespForDeleteInstance()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询IDE实例列表
func (c *CloudideClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances(request)
	resp, responseDef := GenRespForListInstances()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询某个组织下的IDE实例列表
func (c *CloudideClient) ListOrgInstances(request *model.ListOrgInstancesRequest) (*model.ListOrgInstancesResponse, error) {
	requestDef := GenReqDefForListOrgInstances(request)
	resp, responseDef := GenRespForListOrgInstances()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询某个IDE实例
func (c *CloudideClient) ShowInstance(request *model.ShowInstanceRequest) (*model.ShowInstanceResponse, error) {
	requestDef := GenReqDefForShowInstance(request)
	resp, responseDef := GenRespForShowInstance()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//启动IDE实例
func (c *CloudideClient) StartInstance(request *model.StartInstanceRequest) (*model.StartInstanceResponse, error) {
	requestDef := GenReqDefForStartInstance(request)
	resp, responseDef := GenRespForStartInstance()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//停止IDE实例（不删除磁盘数据）
func (c *CloudideClient) StopInstance(request *model.StopInstanceRequest) (*model.StopInstanceResponse, error) {
	requestDef := GenReqDefForStopInstance(request)
	resp, responseDef := GenRespForStopInstance()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//修改IDE实例
func (c *CloudideClient) UpdateInstance(request *model.UpdateInstanceRequest) (*model.UpdateInstanceResponse, error) {
	requestDef := GenReqDefForUpdateInstance(request)
	resp, responseDef := GenRespForUpdateInstance()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}
