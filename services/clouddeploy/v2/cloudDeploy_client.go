package v2

import (
	http_client "github.com/RandolphCYG/huaweicloud-sdk-go-v3/core"

	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/services/clouddeploy/v2/model"
)

type CloudDeployClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCloudDeployClient(hcClient *http_client.HcHttpClient) *CloudDeployClient {
	return &CloudDeployClient{HcClient: hcClient}
}

func CloudDeployClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//通过模板新建部署任务cloudpipeline流水线调用
func (c *CloudDeployClient) CreateDeployTaskByTemplate(request *model.CreateDeployTaskByTemplateRequest) (*model.CreateDeployTaskByTemplateResponse, error) {
	requestDef := GenReqDefForCreateDeployTaskByTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDeployTaskByTemplateResponse), nil
	}
}

//Delete Task by id
func (c *CloudDeployClient) DeleteDeployTask(request *model.DeleteDeployTaskRequest) (*model.DeleteDeployTaskResponse, error) {
	requestDef := GenReqDefForDeleteDeployTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeployTaskResponse), nil
	}
}

//Get Task Detail By Id
func (c *CloudDeployClient) ShowDeployTaskDetail(request *model.ShowDeployTaskDetailRequest) (*model.ShowDeployTaskDetailResponse, error) {
	requestDef := GenReqDefForShowDeployTaskDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeployTaskDetailResponse), nil
	}
}

//Deploy Task Start
func (c *CloudDeployClient) StartDeployTask(request *model.StartDeployTaskRequest) (*model.StartDeployTaskResponse, error) {
	requestDef := GenReqDefForStartDeployTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartDeployTaskResponse), nil
	}
}

//在指定主机组下新建主机
func (c *CloudDeployClient) CreateDeploymentHost(request *model.CreateDeploymentHostRequest) (*model.CreateDeploymentHostResponse, error) {
	requestDef := GenReqDefForCreateDeploymentHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDeploymentHostResponse), nil
	}
}

//根据主机id删除主机
func (c *CloudDeployClient) DeleteDeploymentHost(request *model.DeleteDeploymentHostRequest) (*model.DeleteDeploymentHostResponse, error) {
	requestDef := GenReqDefForDeleteDeploymentHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeploymentHostResponse), nil
	}
}

//查询指定主机组下的主机列表
func (c *CloudDeployClient) ListHosts(request *model.ListHostsRequest) (*model.ListHostsResponse, error) {
	requestDef := GenReqDefForListHosts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostsResponse), nil
	}
}

//根据主机id查询主机详情
func (c *CloudDeployClient) ShowDeploymentHostDetail(request *model.ShowDeploymentHostDetailRequest) (*model.ShowDeploymentHostDetailResponse, error) {
	requestDef := GenReqDefForShowDeploymentHostDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeploymentHostDetailResponse), nil
	}
}

//修改主机信息
func (c *CloudDeployClient) UpdateDeploymentHost(request *model.UpdateDeploymentHostRequest) (*model.UpdateDeploymentHostResponse, error) {
	requestDef := GenReqDefForUpdateDeploymentHost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeploymentHostResponse), nil
	}
}

//新建主机组
func (c *CloudDeployClient) CreateDeploymentGroup(request *model.CreateDeploymentGroupRequest) (*model.CreateDeploymentGroupResponse, error) {
	requestDef := GenReqDefForCreateDeploymentGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDeploymentGroupResponse), nil
	}
}

//根据主机组id删除主机组
func (c *CloudDeployClient) DeleteDeploymentGroup(request *model.DeleteDeploymentGroupRequest) (*model.DeleteDeploymentGroupResponse, error) {
	requestDef := GenReqDefForDeleteDeploymentGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeploymentGroupResponse), nil
	}
}

//查询主机组列表
func (c *CloudDeployClient) ListHostGroups(request *model.ListHostGroupsRequest) (*model.ListHostGroupsResponse, error) {
	requestDef := GenReqDefForListHostGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHostGroupsResponse), nil
	}
}

//根据主机组id查询主机组详情
func (c *CloudDeployClient) ShowDeploymentGroupDetail(request *model.ShowDeploymentGroupDetailRequest) (*model.ShowDeploymentGroupDetailResponse, error) {
	requestDef := GenReqDefForShowDeploymentGroupDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeploymentGroupDetailResponse), nil
	}
}

//修改主机组信息
func (c *CloudDeployClient) UpdateDeploymentGroup(request *model.UpdateDeploymentGroupRequest) (*model.UpdateDeploymentGroupResponse, error) {
	requestDef := GenReqDefForUpdateDeploymentGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeploymentGroupResponse), nil
	}
}
