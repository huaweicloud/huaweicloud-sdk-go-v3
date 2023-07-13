package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cms/v1/model"
)

type CmsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCmsClient(hcClient *http_client.HcHttpClient) *CmsClient {
	return &CmsClient{HcClient: hcClient}
}

func CmsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// CreateAutoLaunchGroup 创建智能购买组
//
// 创建智能购买组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CmsClient) CreateAutoLaunchGroup(request *model.CreateAutoLaunchGroupRequest) (*model.CreateAutoLaunchGroupResponse, error) {
	requestDef := GenReqDefForCreateAutoLaunchGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAutoLaunchGroupResponse), nil
	}
}

// CreateAutoLaunchGroupInvoker 创建智能购买组
func (c *CmsClient) CreateAutoLaunchGroupInvoker(request *model.CreateAutoLaunchGroupRequest) *CreateAutoLaunchGroupInvoker {
	requestDef := GenReqDefForCreateAutoLaunchGroup()
	return &CreateAutoLaunchGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAutoLaunchGroup 删除智能购买组
//
// 删除指定的智能购买组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CmsClient) DeleteAutoLaunchGroup(request *model.DeleteAutoLaunchGroupRequest) (*model.DeleteAutoLaunchGroupResponse, error) {
	requestDef := GenReqDefForDeleteAutoLaunchGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAutoLaunchGroupResponse), nil
	}
}

// DeleteAutoLaunchGroupInvoker 删除智能购买组
func (c *CmsClient) DeleteAutoLaunchGroupInvoker(request *model.DeleteAutoLaunchGroupRequest) *DeleteAutoLaunchGroupInvoker {
	requestDef := GenReqDefForDeleteAutoLaunchGroup()
	return &DeleteAutoLaunchGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAutoLaunchGroups 查询智能购买组列表
//
// 获取租户创建的所有的智能购买组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CmsClient) ListAutoLaunchGroups(request *model.ListAutoLaunchGroupsRequest) (*model.ListAutoLaunchGroupsResponse, error) {
	requestDef := GenReqDefForListAutoLaunchGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAutoLaunchGroupsResponse), nil
	}
}

// ListAutoLaunchGroupsInvoker 查询智能购买组列表
func (c *CmsClient) ListAutoLaunchGroupsInvoker(request *model.ListAutoLaunchGroupsRequest) *ListAutoLaunchGroupsInvoker {
	requestDef := GenReqDefForListAutoLaunchGroups()
	return &ListAutoLaunchGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstances 查询智能购买组实例列表
//
// 获取智能购买组创建的实例信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CmsClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

// ListInstancesInvoker 查询智能购买组实例列表
func (c *CmsClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSupplyRecommendation 地域推荐
//
// 对ECS的资源供给的地域和规格进行推荐，推荐结果以打分的形式呈现，分数越高推荐程度越高
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CmsClient) ListSupplyRecommendation(request *model.ListSupplyRecommendationRequest) (*model.ListSupplyRecommendationResponse, error) {
	requestDef := GenReqDefForListSupplyRecommendation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSupplyRecommendationResponse), nil
	}
}

// ListSupplyRecommendationInvoker 地域推荐
func (c *CmsClient) ListSupplyRecommendationInvoker(request *model.ListSupplyRecommendationRequest) *ListSupplyRecommendationInvoker {
	requestDef := GenReqDefForListSupplyRecommendation()
	return &ListSupplyRecommendationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAutoLaunchGroup 查询智能购买组详情
//
// 查询指定智能购买组的详细信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CmsClient) ShowAutoLaunchGroup(request *model.ShowAutoLaunchGroupRequest) (*model.ShowAutoLaunchGroupResponse, error) {
	requestDef := GenReqDefForShowAutoLaunchGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAutoLaunchGroupResponse), nil
	}
}

// ShowAutoLaunchGroupInvoker 查询智能购买组详情
func (c *CmsClient) ShowAutoLaunchGroupInvoker(request *model.ShowAutoLaunchGroupRequest) *ShowAutoLaunchGroupInvoker {
	requestDef := GenReqDefForShowAutoLaunchGroup()
	return &ShowAutoLaunchGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAutoLaunchGroup 修改智能购买组
//
// 更新指定智能购买组信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CmsClient) UpdateAutoLaunchGroup(request *model.UpdateAutoLaunchGroupRequest) (*model.UpdateAutoLaunchGroupResponse, error) {
	requestDef := GenReqDefForUpdateAutoLaunchGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAutoLaunchGroupResponse), nil
	}
}

// UpdateAutoLaunchGroupInvoker 修改智能购买组
func (c *CmsClient) UpdateAutoLaunchGroupInvoker(request *model.UpdateAutoLaunchGroupRequest) *UpdateAutoLaunchGroupInvoker {
	requestDef := GenReqDefForUpdateAutoLaunchGroup()
	return &UpdateAutoLaunchGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
