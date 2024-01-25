package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iec/v1/model"
)

type IecClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewIecClient(hcClient *httpclient.HcHttpClient) *IecClient {
	return &IecClient{HcClient: hcClient}
}

func IecClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// AddNics 添加网卡
//
// 添加网卡。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) AddNics(request *model.AddNicsRequest) (*model.AddNicsResponse, error) {
	requestDef := GenReqDefForAddNics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddNicsResponse), nil
	}
}

// AddNicsInvoker 添加网卡
func (c *IecClient) AddNicsInvoker(request *model.AddNicsRequest) *AddNicsInvoker {
	requestDef := GenReqDefForAddNics()
	return &AddNicsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateSubnet 路由表关联子网
//
// 路由表关联子网
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) AssociateSubnet(request *model.AssociateSubnetRequest) (*model.AssociateSubnetResponse, error) {
	requestDef := GenReqDefForAssociateSubnet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateSubnetResponse), nil
	}
}

// AssociateSubnetInvoker 路由表关联子网
func (c *IecClient) AssociateSubnetInvoker(request *model.AssociateSubnetRequest) *AssociateSubnetInvoker {
	requestDef := GenReqDefForAssociateSubnet()
	return &AssociateSubnetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchRebootInstance 批量重启边缘实例
//
// 批量重启边缘实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) BatchRebootInstance(request *model.BatchRebootInstanceRequest) (*model.BatchRebootInstanceResponse, error) {
	requestDef := GenReqDefForBatchRebootInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRebootInstanceResponse), nil
	}
}

// BatchRebootInstanceInvoker 批量重启边缘实例
func (c *IecClient) BatchRebootInstanceInvoker(request *model.BatchRebootInstanceRequest) *BatchRebootInstanceInvoker {
	requestDef := GenReqDefForBatchRebootInstance()
	return &BatchRebootInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchStartInstance 批量启动边缘实例
//
// 批量操作启动边缘实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) BatchStartInstance(request *model.BatchStartInstanceRequest) (*model.BatchStartInstanceResponse, error) {
	requestDef := GenReqDefForBatchStartInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStartInstanceResponse), nil
	}
}

// BatchStartInstanceInvoker 批量启动边缘实例
func (c *IecClient) BatchStartInstanceInvoker(request *model.BatchStartInstanceRequest) *BatchStartInstanceInvoker {
	requestDef := GenReqDefForBatchStartInstance()
	return &BatchStartInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchStopInstance 批量关机边缘实例
//
// 批量关闭边缘实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) BatchStopInstance(request *model.BatchStopInstanceRequest) (*model.BatchStopInstanceResponse, error) {
	requestDef := GenReqDefForBatchStopInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStopInstanceResponse), nil
	}
}

// BatchStopInstanceInvoker 批量关机边缘实例
func (c *IecClient) BatchStopInstanceInvoker(request *model.BatchStopInstanceRequest) *BatchStopInstanceInvoker {
	requestDef := GenReqDefForBatchStopInstance()
	return &BatchStopInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeOs 切换操作系统
//
// 切换边缘实例操作系统，支持边缘实例创建成功后，保持ip、数据盘不变的情况下重装操作系统。
//
// 调用该接口后，系统将卸载系统盘，然后使用新镜像重新创建系统盘，并挂载至实例，实现切换操作系统功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ChangeOs(request *model.ChangeOsRequest) (*model.ChangeOsResponse, error) {
	requestDef := GenReqDefForChangeOs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeOsResponse), nil
	}
}

// ChangeOsInvoker 切换操作系统
func (c *IecClient) ChangeOsInvoker(request *model.ChangeOsRequest) *ChangeOsInvoker {
	requestDef := GenReqDefForChangeOs()
	return &ChangeOsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDeployment 创建部署计划
//
// 为方便您的统一管理，以及跨边缘站点管理资源，IEC基于业务场景角度，定义了边缘业务。
// 边缘业务即为逻辑层面的一套资源管理集合。这里的资源主要是指计算实例，包含实例规格、镜像、硬盘、网络等方面。通过指定计算实例的数量、调度策略以及区域分布等形成一套管理集合。[了解更多](https://support.huaweicloud.com/usermanual-iec/iec_02_0301.html)
//
// 创建一个部署计划并执行，即可创建一个边缘业务。
//
// - 边缘业务下实例分布取决于部署计划的实例分布与调度策略。
// - 边缘业务下实例名称、规格、镜像等参数取决于部署计划配置计算实例字段。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) CreateDeployment(request *model.CreateDeploymentRequest) (*model.CreateDeploymentResponse, error) {
	requestDef := GenReqDefForCreateDeployment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDeploymentResponse), nil
	}
}

// CreateDeploymentInvoker 创建部署计划
func (c *IecClient) CreateDeploymentInvoker(request *model.CreateDeploymentRequest) *CreateDeploymentInvoker {
	requestDef := GenReqDefForCreateDeployment()
	return &CreateDeploymentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateImage 从边缘实例创建边缘私有镜像
//
// 使用指定边缘实例的系统盘创建边缘私有镜像。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) CreateImage(request *model.CreateImageRequest) (*model.CreateImageResponse, error) {
	requestDef := GenReqDefForCreateImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateImageResponse), nil
	}
}

// CreateImageInvoker 从边缘实例创建边缘私有镜像
func (c *IecClient) CreateImageInvoker(request *model.CreateImageRequest) *CreateImageInvoker {
	requestDef := GenReqDefForCreateImage()
	return &CreateImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstance 创建边缘实例
//
// 创建边缘实例。单租户默认可创建50个边缘实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) CreateInstance(request *model.CreateInstanceRequest) (*model.CreateInstanceResponse, error) {
	requestDef := GenReqDefForCreateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceResponse), nil
	}
}

// CreateInstanceInvoker 创建边缘实例
func (c *IecClient) CreateInstanceInvoker(request *model.CreateInstanceRequest) *CreateInstanceInvoker {
	requestDef := GenReqDefForCreateInstance()
	return &CreateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateKeypair 创建和导入密钥
//
// 创建SSH密钥，或把公钥导入系统，生成密钥对。
//
// 创建SSH密钥成功后，请把响应数据中的私钥内容保存到本地文件，用户使用该私钥登录边缘实例。为保证边缘实例安全，私钥数据只能读取一次，请妥善保管。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) CreateKeypair(request *model.CreateKeypairRequest) (*model.CreateKeypairResponse, error) {
	requestDef := GenReqDefForCreateKeypair()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateKeypairResponse), nil
	}
}

// CreateKeypairInvoker 创建和导入密钥
func (c *IecClient) CreateKeypairInvoker(request *model.CreateKeypairRequest) *CreateKeypairInvoker {
	requestDef := GenReqDefForCreateKeypair()
	return &CreateKeypairInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePort 创建端口
//
// 创建端口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) CreatePort(request *model.CreatePortRequest) (*model.CreatePortResponse, error) {
	requestDef := GenReqDefForCreatePort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePortResponse), nil
	}
}

// CreatePortInvoker 创建端口
func (c *IecClient) CreatePortInvoker(request *model.CreatePortRequest) *CreatePortInvoker {
	requestDef := GenReqDefForCreatePort()
	return &CreatePortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRoutes 创建路由
//
// 创建路由
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) CreateRoutes(request *model.CreateRoutesRequest) (*model.CreateRoutesResponse, error) {
	requestDef := GenReqDefForCreateRoutes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRoutesResponse), nil
	}
}

// CreateRoutesInvoker 创建路由
func (c *IecClient) CreateRoutesInvoker(request *model.CreateRoutesRequest) *CreateRoutesInvoker {
	requestDef := GenReqDefForCreateRoutes()
	return &CreateRoutesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRoutetable 创建路由表
//
// 创建路由表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) CreateRoutetable(request *model.CreateRoutetableRequest) (*model.CreateRoutetableResponse, error) {
	requestDef := GenReqDefForCreateRoutetable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRoutetableResponse), nil
	}
}

// CreateRoutetableInvoker 创建路由表
func (c *IecClient) CreateRoutetableInvoker(request *model.CreateRoutetableRequest) *CreateRoutetableInvoker {
	requestDef := GenReqDefForCreateRoutetable()
	return &CreateRoutetableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSecurityGroup 创建边缘安全组
//
// 根据用户的请求内容，创建对应的安全组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) CreateSecurityGroup(request *model.CreateSecurityGroupRequest) (*model.CreateSecurityGroupResponse, error) {
	requestDef := GenReqDefForCreateSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecurityGroupResponse), nil
	}
}

// CreateSecurityGroupInvoker 创建边缘安全组
func (c *IecClient) CreateSecurityGroupInvoker(request *model.CreateSecurityGroupRequest) *CreateSecurityGroupInvoker {
	requestDef := GenReqDefForCreateSecurityGroup()
	return &CreateSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSecurityGroupRule 创建安全组规则
//
// 根据用户的请求内容，创建安全组规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) CreateSecurityGroupRule(request *model.CreateSecurityGroupRuleRequest) (*model.CreateSecurityGroupRuleResponse, error) {
	requestDef := GenReqDefForCreateSecurityGroupRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecurityGroupRuleResponse), nil
	}
}

// CreateSecurityGroupRuleInvoker 创建安全组规则
func (c *IecClient) CreateSecurityGroupRuleInvoker(request *model.CreateSecurityGroupRuleRequest) *CreateSecurityGroupRuleInvoker {
	requestDef := GenReqDefForCreateSecurityGroupRule()
	return &CreateSecurityGroupRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVpc 创建虚拟私有云
//
// 根据用户的请求内容，创建虚拟私有云。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) CreateVpc(request *model.CreateVpcRequest) (*model.CreateVpcResponse, error) {
	requestDef := GenReqDefForCreateVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVpcResponse), nil
	}
}

// CreateVpcInvoker 创建虚拟私有云
func (c *IecClient) CreateVpcInvoker(request *model.CreateVpcRequest) *CreateVpcInvoker {
	requestDef := GenReqDefForCreateVpc()
	return &CreateVpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBandwidth 删除带宽
//
// 删除带宽。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) DeleteBandwidth(request *model.DeleteBandwidthRequest) (*model.DeleteBandwidthResponse, error) {
	requestDef := GenReqDefForDeleteBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBandwidthResponse), nil
	}
}

// DeleteBandwidthInvoker 删除带宽
func (c *IecClient) DeleteBandwidthInvoker(request *model.DeleteBandwidthRequest) *DeleteBandwidthInvoker {
	requestDef := GenReqDefForDeleteBandwidth()
	return &DeleteBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDeployment 删除部署计划
//
// 删除部署计划。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) DeleteDeployment(request *model.DeleteDeploymentRequest) (*model.DeleteDeploymentResponse, error) {
	requestDef := GenReqDefForDeleteDeployment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeploymentResponse), nil
	}
}

// DeleteDeploymentInvoker 删除部署计划
func (c *IecClient) DeleteDeploymentInvoker(request *model.DeleteDeploymentRequest) *DeleteDeploymentInvoker {
	requestDef := GenReqDefForDeleteDeployment()
	return &DeleteDeploymentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEdgeCloud 删除边缘业务
//
// 删除边缘业务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) DeleteEdgeCloud(request *model.DeleteEdgeCloudRequest) (*model.DeleteEdgeCloudResponse, error) {
	requestDef := GenReqDefForDeleteEdgeCloud()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEdgeCloudResponse), nil
	}
}

// DeleteEdgeCloudInvoker 删除边缘业务
func (c *IecClient) DeleteEdgeCloudInvoker(request *model.DeleteEdgeCloudRequest) *DeleteEdgeCloudInvoker {
	requestDef := GenReqDefForDeleteEdgeCloud()
	return &DeleteEdgeCloudInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteImage 删除边缘私有镜像
//
// 将指定ID的边缘私有镜像删除
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) DeleteImage(request *model.DeleteImageRequest) (*model.DeleteImageResponse, error) {
	requestDef := GenReqDefForDeleteImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteImageResponse), nil
	}
}

// DeleteImageInvoker 删除边缘私有镜像
func (c *IecClient) DeleteImageInvoker(request *model.DeleteImageRequest) *DeleteImageInvoker {
	requestDef := GenReqDefForDeleteImage()
	return &DeleteImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstances 批量删除边缘实例
//
// 批量删除边缘实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) DeleteInstances(request *model.DeleteInstancesRequest) (*model.DeleteInstancesResponse, error) {
	requestDef := GenReqDefForDeleteInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstancesResponse), nil
	}
}

// DeleteInstancesInvoker 批量删除边缘实例
func (c *IecClient) DeleteInstancesInvoker(request *model.DeleteInstancesRequest) *DeleteInstancesInvoker {
	requestDef := GenReqDefForDeleteInstances()
	return &DeleteInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteKeypair 删除密钥
//
// 删除密钥。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) DeleteKeypair(request *model.DeleteKeypairRequest) (*model.DeleteKeypairResponse, error) {
	requestDef := GenReqDefForDeleteKeypair()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteKeypairResponse), nil
	}
}

// DeleteKeypairInvoker 删除密钥
func (c *IecClient) DeleteKeypairInvoker(request *model.DeleteKeypairRequest) *DeleteKeypairInvoker {
	requestDef := GenReqDefForDeleteKeypair()
	return &DeleteKeypairInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteNics 删除网卡
//
// 删除网卡。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) DeleteNics(request *model.DeleteNicsRequest) (*model.DeleteNicsResponse, error) {
	requestDef := GenReqDefForDeleteNics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNicsResponse), nil
	}
}

// DeleteNicsInvoker 删除网卡
func (c *IecClient) DeleteNicsInvoker(request *model.DeleteNicsRequest) *DeleteNicsInvoker {
	requestDef := GenReqDefForDeleteNics()
	return &DeleteNicsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePort 删除端口
//
// 删除端口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) DeletePort(request *model.DeletePortRequest) (*model.DeletePortResponse, error) {
	requestDef := GenReqDefForDeletePort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePortResponse), nil
	}
}

// DeletePortInvoker 删除端口
func (c *IecClient) DeletePortInvoker(request *model.DeletePortRequest) *DeletePortInvoker {
	requestDef := GenReqDefForDeletePort()
	return &DeletePortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRoutes 删除路由
//
// 删除路由
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) DeleteRoutes(request *model.DeleteRoutesRequest) (*model.DeleteRoutesResponse, error) {
	requestDef := GenReqDefForDeleteRoutes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRoutesResponse), nil
	}
}

// DeleteRoutesInvoker 删除路由
func (c *IecClient) DeleteRoutesInvoker(request *model.DeleteRoutesRequest) *DeleteRoutesInvoker {
	requestDef := GenReqDefForDeleteRoutes()
	return &DeleteRoutesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRoutetable 删除路由表
//
// 删除路由表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) DeleteRoutetable(request *model.DeleteRoutetableRequest) (*model.DeleteRoutetableResponse, error) {
	requestDef := GenReqDefForDeleteRoutetable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRoutetableResponse), nil
	}
}

// DeleteRoutetableInvoker 删除路由表
func (c *IecClient) DeleteRoutetableInvoker(request *model.DeleteRoutetableRequest) *DeleteRoutetableInvoker {
	requestDef := GenReqDefForDeleteRoutetable()
	return &DeleteRoutetableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSecurityGroup 删除安全组
//
// 根据安全组的ID，删除对应的安全组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) DeleteSecurityGroup(request *model.DeleteSecurityGroupRequest) (*model.DeleteSecurityGroupResponse, error) {
	requestDef := GenReqDefForDeleteSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecurityGroupResponse), nil
	}
}

// DeleteSecurityGroupInvoker 删除安全组
func (c *IecClient) DeleteSecurityGroupInvoker(request *model.DeleteSecurityGroupRequest) *DeleteSecurityGroupInvoker {
	requestDef := GenReqDefForDeleteSecurityGroup()
	return &DeleteSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSecurityGroupRule 删除安全组规则
//
// 根据安全组的ID，删除对应的安全组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) DeleteSecurityGroupRule(request *model.DeleteSecurityGroupRuleRequest) (*model.DeleteSecurityGroupRuleResponse, error) {
	requestDef := GenReqDefForDeleteSecurityGroupRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecurityGroupRuleResponse), nil
	}
}

// DeleteSecurityGroupRuleInvoker 删除安全组规则
func (c *IecClient) DeleteSecurityGroupRuleInvoker(request *model.DeleteSecurityGroupRuleRequest) *DeleteSecurityGroupRuleInvoker {
	requestDef := GenReqDefForDeleteSecurityGroupRule()
	return &DeleteSecurityGroupRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSubnet 删除子网
//
// 根据子网的ID，删除子网。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) DeleteSubnet(request *model.DeleteSubnetRequest) (*model.DeleteSubnetResponse, error) {
	requestDef := GenReqDefForDeleteSubnet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSubnetResponse), nil
	}
}

// DeleteSubnetInvoker 删除子网
func (c *IecClient) DeleteSubnetInvoker(request *model.DeleteSubnetRequest) *DeleteSubnetInvoker {
	requestDef := GenReqDefForDeleteSubnet()
	return &DeleteSubnetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVpc 删除虚拟私有云
//
// 根据虚拟机私有云的ID，删除对应的虚拟私有云。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) DeleteVpc(request *model.DeleteVpcRequest) (*model.DeleteVpcResponse, error) {
	requestDef := GenReqDefForDeleteVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVpcResponse), nil
	}
}

// DeleteVpcInvoker 删除虚拟私有云
func (c *IecClient) DeleteVpcInvoker(request *model.DeleteVpcRequest) *DeleteVpcInvoker {
	requestDef := GenReqDefForDeleteVpc()
	return &DeleteVpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateSubnet 路由表解关联子网
//
// 路由表解关联子网
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) DisassociateSubnet(request *model.DisassociateSubnetRequest) (*model.DisassociateSubnetResponse, error) {
	requestDef := GenReqDefForDisassociateSubnet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateSubnetResponse), nil
	}
}

// DisassociateSubnetInvoker 路由表解关联子网
func (c *IecClient) DisassociateSubnetInvoker(request *model.DisassociateSubnetRequest) *DisassociateSubnetInvoker {
	requestDef := GenReqDefForDisassociateSubnet()
	return &DisassociateSubnetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteDeployment 执行部署计划
//
// 执行部署计划，创建一个边缘业务。单租户默认可创建10个边缘业务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ExecuteDeployment(request *model.ExecuteDeploymentRequest) (*model.ExecuteDeploymentResponse, error) {
	requestDef := GenReqDefForExecuteDeployment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteDeploymentResponse), nil
	}
}

// ExecuteDeploymentInvoker 执行部署计划
func (c *IecClient) ExecuteDeploymentInvoker(request *model.ExecuteDeploymentRequest) *ExecuteDeploymentInvoker {
	requestDef := GenReqDefForExecuteDeployment()
	return &ExecuteDeploymentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExpandEdgecloud 扩容边缘业务
//
// 执行部署计划，对边缘业务进行扩容操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ExpandEdgecloud(request *model.ExpandEdgecloudRequest) (*model.ExpandEdgecloudResponse, error) {
	requestDef := GenReqDefForExpandEdgecloud()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandEdgecloudResponse), nil
	}
}

// ExpandEdgecloudInvoker 扩容边缘业务
func (c *IecClient) ExpandEdgecloudInvoker(request *model.ExpandEdgecloudRequest) *ExpandEdgecloudInvoker {
	requestDef := GenReqDefForExpandEdgecloud()
	return &ExpandEdgecloudInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBandwidthTypes 查询共享带宽类型列表
//
// 查询共享带宽类型列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ListBandwidthTypes(request *model.ListBandwidthTypesRequest) (*model.ListBandwidthTypesResponse, error) {
	requestDef := GenReqDefForListBandwidthTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBandwidthTypesResponse), nil
	}
}

// ListBandwidthTypesInvoker 查询共享带宽类型列表
func (c *IecClient) ListBandwidthTypesInvoker(request *model.ListBandwidthTypesRequest) *ListBandwidthTypesInvoker {
	requestDef := GenReqDefForListBandwidthTypes()
	return &ListBandwidthTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBandwidths 查询带宽列表
//
// 查询带宽列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ListBandwidths(request *model.ListBandwidthsRequest) (*model.ListBandwidthsResponse, error) {
	requestDef := GenReqDefForListBandwidths()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBandwidthsResponse), nil
	}
}

// ListBandwidthsInvoker 查询带宽列表
func (c *IecClient) ListBandwidthsInvoker(request *model.ListBandwidthsRequest) *ListBandwidthsInvoker {
	requestDef := GenReqDefForListBandwidths()
	return &ListBandwidthsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCloudImages 查询中心镜像列表
//
// 查询租户在某个云Region的可见镜像列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ListCloudImages(request *model.ListCloudImagesRequest) (*model.ListCloudImagesResponse, error) {
	requestDef := GenReqDefForListCloudImages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCloudImagesResponse), nil
	}
}

// ListCloudImagesInvoker 查询中心镜像列表
func (c *IecClient) ListCloudImagesInvoker(request *model.ListCloudImagesRequest) *ListCloudImagesInvoker {
	requestDef := GenReqDefForListCloudImages()
	return &ListCloudImagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDeployments 查询部署计划列表
//
// 查询部署计划列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ListDeployments(request *model.ListDeploymentsRequest) (*model.ListDeploymentsResponse, error) {
	requestDef := GenReqDefForListDeployments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDeploymentsResponse), nil
	}
}

// ListDeploymentsInvoker 查询部署计划列表
func (c *IecClient) ListDeploymentsInvoker(request *model.ListDeploymentsRequest) *ListDeploymentsInvoker {
	requestDef := GenReqDefForListDeployments()
	return &ListDeploymentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEdgeCloud 查询边缘业务列表
//
// 查询边缘业务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ListEdgeCloud(request *model.ListEdgeCloudRequest) (*model.ListEdgeCloudResponse, error) {
	requestDef := GenReqDefForListEdgeCloud()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEdgeCloudResponse), nil
	}
}

// ListEdgeCloudInvoker 查询边缘业务列表
func (c *IecClient) ListEdgeCloudInvoker(request *model.ListEdgeCloudRequest) *ListEdgeCloudInvoker {
	requestDef := GenReqDefForListEdgeCloud()
	return &ListEdgeCloudInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlavors 查询边缘规格列表
//
// 查询边缘规格列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsResponse), nil
	}
}

// ListFlavorsInvoker 查询边缘规格列表
func (c *IecClient) ListFlavorsInvoker(request *model.ListFlavorsRequest) *ListFlavorsInvoker {
	requestDef := GenReqDefForListFlavors()
	return &ListFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImages 查询镜像列表
//
// 根据不同条件查询镜像列表，例:
//
// -  查询已注册的私有镜像列表: visibility&#x3D;private
// - 公共镜像: visibility&#x3D;public
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ListImages(request *model.ListImagesRequest) (*model.ListImagesResponse, error) {
	requestDef := GenReqDefForListImages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImagesResponse), nil
	}
}

// ListImagesInvoker 查询镜像列表
func (c *IecClient) ListImagesInvoker(request *model.ListImagesRequest) *ListImagesInvoker {
	requestDef := GenReqDefForListImages()
	return &ListImagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstances 查询边缘实例列表
//
// 查询边缘实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

// ListInstancesInvoker 查询边缘实例列表
func (c *IecClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListKeypairs 查询密钥列表
//
// 查询密钥信息列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ListKeypairs(request *model.ListKeypairsRequest) (*model.ListKeypairsResponse, error) {
	requestDef := GenReqDefForListKeypairs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKeypairsResponse), nil
	}
}

// ListKeypairsInvoker 查询密钥列表
func (c *IecClient) ListKeypairsInvoker(request *model.ListKeypairsRequest) *ListKeypairsInvoker {
	requestDef := GenReqDefForListKeypairs()
	return &ListKeypairsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPorts 查询端口列表
//
// 查询端口的列表信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ListPorts(request *model.ListPortsRequest) (*model.ListPortsResponse, error) {
	requestDef := GenReqDefForListPorts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPortsResponse), nil
	}
}

// ListPortsInvoker 查询端口列表
func (c *IecClient) ListPortsInvoker(request *model.ListPortsRequest) *ListPortsInvoker {
	requestDef := GenReqDefForListPorts()
	return &ListPortsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuota 查询配额
//
// 查询租户资源配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ListQuota(request *model.ListQuotaRequest) (*model.ListQuotaResponse, error) {
	requestDef := GenReqDefForListQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotaResponse), nil
	}
}

// ListQuotaInvoker 查询配额
func (c *IecClient) ListQuotaInvoker(request *model.ListQuotaRequest) *ListQuotaInvoker {
	requestDef := GenReqDefForListQuota()
	return &ListQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRelatedRoutetables 查询子网关联的路由表
//
// 查询子网关联的路由表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ListRelatedRoutetables(request *model.ListRelatedRoutetablesRequest) (*model.ListRelatedRoutetablesResponse, error) {
	requestDef := GenReqDefForListRelatedRoutetables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRelatedRoutetablesResponse), nil
	}
}

// ListRelatedRoutetablesInvoker 查询子网关联的路由表
func (c *IecClient) ListRelatedRoutetablesInvoker(request *model.ListRelatedRoutetablesRequest) *ListRelatedRoutetablesInvoker {
	requestDef := GenReqDefForListRelatedRoutetables()
	return &ListRelatedRoutetablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRoutes 查询路由列表
//
// 查询路由列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ListRoutes(request *model.ListRoutesRequest) (*model.ListRoutesResponse, error) {
	requestDef := GenReqDefForListRoutes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRoutesResponse), nil
	}
}

// ListRoutesInvoker 查询路由列表
func (c *IecClient) ListRoutesInvoker(request *model.ListRoutesRequest) *ListRoutesInvoker {
	requestDef := GenReqDefForListRoutes()
	return &ListRoutesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRoutetables 查询路由表列表
//
// 查询路由列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ListRoutetables(request *model.ListRoutetablesRequest) (*model.ListRoutetablesResponse, error) {
	requestDef := GenReqDefForListRoutetables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRoutetablesResponse), nil
	}
}

// ListRoutetablesInvoker 查询路由表列表
func (c *IecClient) ListRoutetablesInvoker(request *model.ListRoutetablesRequest) *ListRoutetablesInvoker {
	requestDef := GenReqDefForListRoutetables()
	return &ListRoutetablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecurityGroupRules 查询安全组规则列表
//
// 根据用户的查询条件，获取安全组规则的列表信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ListSecurityGroupRules(request *model.ListSecurityGroupRulesRequest) (*model.ListSecurityGroupRulesResponse, error) {
	requestDef := GenReqDefForListSecurityGroupRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityGroupRulesResponse), nil
	}
}

// ListSecurityGroupRulesInvoker 查询安全组规则列表
func (c *IecClient) ListSecurityGroupRulesInvoker(request *model.ListSecurityGroupRulesRequest) *ListSecurityGroupRulesInvoker {
	requestDef := GenReqDefForListSecurityGroupRules()
	return &ListSecurityGroupRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecurityGroups 查询安全组列表
//
// 根据特定查询条件，获取安全组的列表信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ListSecurityGroups(request *model.ListSecurityGroupsRequest) (*model.ListSecurityGroupsResponse, error) {
	requestDef := GenReqDefForListSecurityGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityGroupsResponse), nil
	}
}

// ListSecurityGroupsInvoker 查询安全组列表
func (c *IecClient) ListSecurityGroupsInvoker(request *model.ListSecurityGroupsRequest) *ListSecurityGroupsInvoker {
	requestDef := GenReqDefForListSecurityGroups()
	return &ListSecurityGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSites 查询边缘站点列表
//
// 查询边缘站点列表。
//
// - 边缘站点：靠近终端应用的位置，基于一个或多个运营商建立的一个城市级站点。边缘站点提供物理隔离的资源池，提供多元算力、存储和网络的能力。用户可以将业务灵活就近部署在边缘站点上，以降低网络时延和成本。
// - 边缘区域：为依据边缘站点的物理位置划分的区域，一个边缘区域包含多个相靠近的边缘站点的集合。IEC当前提供城市级、省级和大区级三个分布层级的边缘区域。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ListSites(request *model.ListSitesRequest) (*model.ListSitesResponse, error) {
	requestDef := GenReqDefForListSites()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSitesResponse), nil
	}
}

// ListSitesInvoker 查询边缘站点列表
func (c *IecClient) ListSitesInvoker(request *model.ListSitesRequest) *ListSitesInvoker {
	requestDef := GenReqDefForListSites()
	return &ListSitesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubnets 查询子网列表
//
// 根据查询条件获取子网的列表信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ListSubnets(request *model.ListSubnetsRequest) (*model.ListSubnetsResponse, error) {
	requestDef := GenReqDefForListSubnets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubnetsResponse), nil
	}
}

// ListSubnetsInvoker 查询子网列表
func (c *IecClient) ListSubnetsInvoker(request *model.ListSubnetsRequest) *ListSubnetsInvoker {
	requestDef := GenReqDefForListSubnets()
	return &ListSubnetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVolume 查询硬盘列表
//
// 查询硬盘列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ListVolume(request *model.ListVolumeRequest) (*model.ListVolumeResponse, error) {
	requestDef := GenReqDefForListVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVolumeResponse), nil
	}
}

// ListVolumeInvoker 查询硬盘列表
func (c *IecClient) ListVolumeInvoker(request *model.ListVolumeRequest) *ListVolumeInvoker {
	requestDef := GenReqDefForListVolume()
	return &ListVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVpcs 查询虚拟私有云列表
//
// 获取虚拟私有云的列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ListVpcs(request *model.ListVpcsRequest) (*model.ListVpcsResponse, error) {
	requestDef := GenReqDefForListVpcs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpcsResponse), nil
	}
}

// ListVpcsInvoker 查询虚拟私有云列表
func (c *IecClient) ListVpcsInvoker(request *model.ListVpcsRequest) *ListVpcsInvoker {
	requestDef := GenReqDefForListVpcs()
	return &ListVpcsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RebuildImage 重试边缘镜像任务
//
// 重试边缘镜像任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) RebuildImage(request *model.RebuildImageRequest) (*model.RebuildImageResponse, error) {
	requestDef := GenReqDefForRebuildImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RebuildImageResponse), nil
	}
}

// RebuildImageInvoker 重试边缘镜像任务
func (c *IecClient) RebuildImageInvoker(request *model.RebuildImageRequest) *RebuildImageInvoker {
	requestDef := GenReqDefForRebuildImage()
	return &RebuildImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RegisterImage 注册边缘私有镜像
//
// 将指定Region和ID的IMS镜像注册到边缘IEC-IMS;
// 注意指定的Region必须在当前IEC-IMS支持的Region列表中。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) RegisterImage(request *model.RegisterImageRequest) (*model.RegisterImageResponse, error) {
	requestDef := GenReqDefForRegisterImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterImageResponse), nil
	}
}

// RegisterImageInvoker 注册边缘私有镜像
func (c *IecClient) RegisterImageInvoker(request *model.RegisterImageRequest) *RegisterImageInvoker {
	requestDef := GenReqDefForRegisterImage()
	return &RegisterImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBandwidth 查询带宽详情
//
// 查询带宽详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ShowBandwidth(request *model.ShowBandwidthRequest) (*model.ShowBandwidthResponse, error) {
	requestDef := GenReqDefForShowBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBandwidthResponse), nil
	}
}

// ShowBandwidthInvoker 查询带宽详情
func (c *IecClient) ShowBandwidthInvoker(request *model.ShowBandwidthRequest) *ShowBandwidthInvoker {
	requestDef := GenReqDefForShowBandwidth()
	return &ShowBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEdgeCloud 查询边缘业务详情
//
// 查询边缘业务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ShowEdgeCloud(request *model.ShowEdgeCloudRequest) (*model.ShowEdgeCloudResponse, error) {
	requestDef := GenReqDefForShowEdgeCloud()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEdgeCloudResponse), nil
	}
}

// ShowEdgeCloudInvoker 查询边缘业务详情
func (c *IecClient) ShowEdgeCloudInvoker(request *model.ShowEdgeCloudRequest) *ShowEdgeCloudInvoker {
	requestDef := GenReqDefForShowEdgeCloud()
	return &ShowEdgeCloudInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowImage 查询镜像详情
//
// 查询镜像详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ShowImage(request *model.ShowImageRequest) (*model.ShowImageResponse, error) {
	requestDef := GenReqDefForShowImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowImageResponse), nil
	}
}

// ShowImageInvoker 查询镜像详情
func (c *IecClient) ShowImageInvoker(request *model.ShowImageRequest) *ShowImageInvoker {
	requestDef := GenReqDefForShowImage()
	return &ShowImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstance 查询边缘实例详情
//
// 查询边缘实例详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ShowInstance(request *model.ShowInstanceRequest) (*model.ShowInstanceResponse, error) {
	requestDef := GenReqDefForShowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceResponse), nil
	}
}

// ShowInstanceInvoker 查询边缘实例详情
func (c *IecClient) ShowInstanceInvoker(request *model.ShowInstanceRequest) *ShowInstanceInvoker {
	requestDef := GenReqDefForShowInstance()
	return &ShowInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowKeypair 查询密钥详情
//
// 查询密钥信息列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ShowKeypair(request *model.ShowKeypairRequest) (*model.ShowKeypairResponse, error) {
	requestDef := GenReqDefForShowKeypair()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKeypairResponse), nil
	}
}

// ShowKeypairInvoker 查询密钥详情
func (c *IecClient) ShowKeypairInvoker(request *model.ShowKeypairRequest) *ShowKeypairInvoker {
	requestDef := GenReqDefForShowKeypair()
	return &ShowKeypairInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPort 查询端口详情
//
// 根据端口的ID，获取端口的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ShowPort(request *model.ShowPortRequest) (*model.ShowPortResponse, error) {
	requestDef := GenReqDefForShowPort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPortResponse), nil
	}
}

// ShowPortInvoker 查询端口详情
func (c *IecClient) ShowPortInvoker(request *model.ShowPortRequest) *ShowPortInvoker {
	requestDef := GenReqDefForShowPort()
	return &ShowPortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRoutetable 查询路由表详情
//
// 查询路由表详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ShowRoutetable(request *model.ShowRoutetableRequest) (*model.ShowRoutetableResponse, error) {
	requestDef := GenReqDefForShowRoutetable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRoutetableResponse), nil
	}
}

// ShowRoutetableInvoker 查询路由表详情
func (c *IecClient) ShowRoutetableInvoker(request *model.ShowRoutetableRequest) *ShowRoutetableInvoker {
	requestDef := GenReqDefForShowRoutetable()
	return &ShowRoutetableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSecurityGroup 查询安全组详情
//
// 根据安全组的ID，获取特定安全组的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ShowSecurityGroup(request *model.ShowSecurityGroupRequest) (*model.ShowSecurityGroupResponse, error) {
	requestDef := GenReqDefForShowSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecurityGroupResponse), nil
	}
}

// ShowSecurityGroupInvoker 查询安全组详情
func (c *IecClient) ShowSecurityGroupInvoker(request *model.ShowSecurityGroupRequest) *ShowSecurityGroupInvoker {
	requestDef := GenReqDefForShowSecurityGroup()
	return &ShowSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSecurityGroupRule 查询安全组规则详情
//
// 根据安全组规则的ID，获取安全组规则的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ShowSecurityGroupRule(request *model.ShowSecurityGroupRuleRequest) (*model.ShowSecurityGroupRuleResponse, error) {
	requestDef := GenReqDefForShowSecurityGroupRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecurityGroupRuleResponse), nil
	}
}

// ShowSecurityGroupRuleInvoker 查询安全组规则详情
func (c *IecClient) ShowSecurityGroupRuleInvoker(request *model.ShowSecurityGroupRuleRequest) *ShowSecurityGroupRuleInvoker {
	requestDef := GenReqDefForShowSecurityGroupRule()
	return &ShowSecurityGroupRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSubnet 查询子网详情
//
// 根据子网的ID，获取子网的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ShowSubnet(request *model.ShowSubnetRequest) (*model.ShowSubnetResponse, error) {
	requestDef := GenReqDefForShowSubnet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSubnetResponse), nil
	}
}

// ShowSubnetInvoker 查询子网详情
func (c *IecClient) ShowSubnetInvoker(request *model.ShowSubnetRequest) *ShowSubnetInvoker {
	requestDef := GenReqDefForShowSubnet()
	return &ShowSubnetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVolume 查询硬盘详情
//
// 查询硬盘详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ShowVolume(request *model.ShowVolumeRequest) (*model.ShowVolumeResponse, error) {
	requestDef := GenReqDefForShowVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVolumeResponse), nil
	}
}

// ShowVolumeInvoker 查询硬盘详情
func (c *IecClient) ShowVolumeInvoker(request *model.ShowVolumeRequest) *ShowVolumeInvoker {
	requestDef := GenReqDefForShowVolume()
	return &ShowVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVolumeTypes 查询硬盘类型列表
//
// 查询硬盘类型列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ShowVolumeTypes(request *model.ShowVolumeTypesRequest) (*model.ShowVolumeTypesResponse, error) {
	requestDef := GenReqDefForShowVolumeTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVolumeTypesResponse), nil
	}
}

// ShowVolumeTypesInvoker 查询硬盘类型列表
func (c *IecClient) ShowVolumeTypesInvoker(request *model.ShowVolumeTypesRequest) *ShowVolumeTypesInvoker {
	requestDef := GenReqDefForShowVolumeTypes()
	return &ShowVolumeTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVpc 查询虚拟私有云详情
//
// 根据虚拟私有云ID，获取虚拟私有云的详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ShowVpc(request *model.ShowVpcRequest) (*model.ShowVpcResponse, error) {
	requestDef := GenReqDefForShowVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVpcResponse), nil
	}
}

// ShowVpcInvoker 查询虚拟私有云详情
func (c *IecClient) ShowVpcInvoker(request *model.ShowVpcRequest) *ShowVpcInvoker {
	requestDef := GenReqDefForShowVpc()
	return &ShowVpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateBandwidth 更新带宽
//
// 更新带宽。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) UpdateBandwidth(request *model.UpdateBandwidthRequest) (*model.UpdateBandwidthResponse, error) {
	requestDef := GenReqDefForUpdateBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBandwidthResponse), nil
	}
}

// UpdateBandwidthInvoker 更新带宽
func (c *IecClient) UpdateBandwidthInvoker(request *model.UpdateBandwidthRequest) *UpdateBandwidthInvoker {
	requestDef := GenReqDefForUpdateBandwidth()
	return &UpdateBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstance 修改边缘实例
//
// 修改边缘实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) UpdateInstance(request *model.UpdateInstanceRequest) (*model.UpdateInstanceResponse, error) {
	requestDef := GenReqDefForUpdateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceResponse), nil
	}
}

// UpdateInstanceInvoker 修改边缘实例
func (c *IecClient) UpdateInstanceInvoker(request *model.UpdateInstanceRequest) *UpdateInstanceInvoker {
	requestDef := GenReqDefForUpdateInstance()
	return &UpdateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePort 更新端口
//
// 更新端口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) UpdatePort(request *model.UpdatePortRequest) (*model.UpdatePortResponse, error) {
	requestDef := GenReqDefForUpdatePort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePortResponse), nil
	}
}

// UpdatePortInvoker 更新端口
func (c *IecClient) UpdatePortInvoker(request *model.UpdatePortRequest) *UpdatePortInvoker {
	requestDef := GenReqDefForUpdatePort()
	return &UpdatePortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRoutes 更新路由
//
// 更新路由信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) UpdateRoutes(request *model.UpdateRoutesRequest) (*model.UpdateRoutesResponse, error) {
	requestDef := GenReqDefForUpdateRoutes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRoutesResponse), nil
	}
}

// UpdateRoutesInvoker 更新路由
func (c *IecClient) UpdateRoutesInvoker(request *model.UpdateRoutesRequest) *UpdateRoutesInvoker {
	requestDef := GenReqDefForUpdateRoutes()
	return &UpdateRoutesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRoutetable 更新路由表
//
// 更新路由表基本信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) UpdateRoutetable(request *model.UpdateRoutetableRequest) (*model.UpdateRoutetableResponse, error) {
	requestDef := GenReqDefForUpdateRoutetable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRoutetableResponse), nil
	}
}

// UpdateRoutetableInvoker 更新路由表
func (c *IecClient) UpdateRoutetableInvoker(request *model.UpdateRoutetableRequest) *UpdateRoutetableInvoker {
	requestDef := GenReqDefForUpdateRoutetable()
	return &UpdateRoutetableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSubnet 更新子网
//
// 更新子网的基本信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) UpdateSubnet(request *model.UpdateSubnetRequest) (*model.UpdateSubnetResponse, error) {
	requestDef := GenReqDefForUpdateSubnet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSubnetResponse), nil
	}
}

// UpdateSubnetInvoker 更新子网
func (c *IecClient) UpdateSubnetInvoker(request *model.UpdateSubnetRequest) *UpdateSubnetInvoker {
	requestDef := GenReqDefForUpdateSubnet()
	return &UpdateSubnetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVpc 更新虚拟私有云
//
// 更新虚拟私有云的信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) UpdateVpc(request *model.UpdateVpcRequest) (*model.UpdateVpcResponse, error) {
	requestDef := GenReqDefForUpdateVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVpcResponse), nil
	}
}

// UpdateVpcInvoker 更新虚拟私有云
func (c *IecClient) UpdateVpcInvoker(request *model.UpdateVpcRequest) *UpdateVpcInvoker {
	requestDef := GenReqDefForUpdateVpc()
	return &UpdateVpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFirewall 创建网络ACL
//
// 创建网络ACL。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) CreateFirewall(request *model.CreateFirewallRequest) (*model.CreateFirewallResponse, error) {
	requestDef := GenReqDefForCreateFirewall()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFirewallResponse), nil
	}
}

// CreateFirewallInvoker 创建网络ACL
func (c *IecClient) CreateFirewallInvoker(request *model.CreateFirewallRequest) *CreateFirewallInvoker {
	requestDef := GenReqDefForCreateFirewall()
	return &CreateFirewallInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFirewall 删除网络ACL
//
// 删除网络ACL。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) DeleteFirewall(request *model.DeleteFirewallRequest) (*model.DeleteFirewallResponse, error) {
	requestDef := GenReqDefForDeleteFirewall()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFirewallResponse), nil
	}
}

// DeleteFirewallInvoker 删除网络ACL
func (c *IecClient) DeleteFirewallInvoker(request *model.DeleteFirewallRequest) *DeleteFirewallInvoker {
	requestDef := GenReqDefForDeleteFirewall()
	return &DeleteFirewallInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFirewalls 查询网络ACL列表
//
// 查询网络ACL列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ListFirewalls(request *model.ListFirewallsRequest) (*model.ListFirewallsResponse, error) {
	requestDef := GenReqDefForListFirewalls()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFirewallsResponse), nil
	}
}

// ListFirewallsInvoker 查询网络ACL列表
func (c *IecClient) ListFirewallsInvoker(request *model.ListFirewallsRequest) *ListFirewallsInvoker {
	requestDef := GenReqDefForListFirewalls()
	return &ListFirewallsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFirewall 查询网络ACL详情
//
// 查询网络ACL详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ShowFirewall(request *model.ShowFirewallRequest) (*model.ShowFirewallResponse, error) {
	requestDef := GenReqDefForShowFirewall()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFirewallResponse), nil
	}
}

// ShowFirewallInvoker 查询网络ACL详情
func (c *IecClient) ShowFirewallInvoker(request *model.ShowFirewallRequest) *ShowFirewallInvoker {
	requestDef := GenReqDefForShowFirewall()
	return &ShowFirewallInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFirewall 更新网络ACL
//
// 更新网络ACL。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) UpdateFirewall(request *model.UpdateFirewallRequest) (*model.UpdateFirewallResponse, error) {
	requestDef := GenReqDefForUpdateFirewall()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFirewallResponse), nil
	}
}

// UpdateFirewallInvoker 更新网络ACL
func (c *IecClient) UpdateFirewallInvoker(request *model.UpdateFirewallRequest) *UpdateFirewallInvoker {
	requestDef := GenReqDefForUpdateFirewall()
	return &UpdateFirewallInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFirewallRule 更新网络ACL规则
//
// 更新网络ACL规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) UpdateFirewallRule(request *model.UpdateFirewallRuleRequest) (*model.UpdateFirewallRuleResponse, error) {
	requestDef := GenReqDefForUpdateFirewallRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFirewallRuleResponse), nil
	}
}

// UpdateFirewallRuleInvoker 更新网络ACL规则
func (c *IecClient) UpdateFirewallRuleInvoker(request *model.UpdateFirewallRuleRequest) *UpdateFirewallRuleInvoker {
	requestDef := GenReqDefForUpdateFirewallRule()
	return &UpdateFirewallRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePublicIp 创建弹性公网IP
//
// 根据用户的请求内容，创建弹性公网IP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) CreatePublicIp(request *model.CreatePublicIpRequest) (*model.CreatePublicIpResponse, error) {
	requestDef := GenReqDefForCreatePublicIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePublicIpResponse), nil
	}
}

// CreatePublicIpInvoker 创建弹性公网IP
func (c *IecClient) CreatePublicIpInvoker(request *model.CreatePublicIpRequest) *CreatePublicIpInvoker {
	requestDef := GenReqDefForCreatePublicIp()
	return &CreatePublicIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePublicIp 删除弹性公网IP
//
// 根据弹性公网IP的ID，删除对应的弹性公网IP。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) DeletePublicIp(request *model.DeletePublicIpRequest) (*model.DeletePublicIpResponse, error) {
	requestDef := GenReqDefForDeletePublicIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePublicIpResponse), nil
	}
}

// DeletePublicIpInvoker 删除弹性公网IP
func (c *IecClient) DeletePublicIpInvoker(request *model.DeletePublicIpRequest) *DeletePublicIpInvoker {
	requestDef := GenReqDefForDeletePublicIp()
	return &DeletePublicIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPublicIps 查询弹性公网IP列表
//
// 获取弹性公网IP列表信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ListPublicIps(request *model.ListPublicIpsRequest) (*model.ListPublicIpsResponse, error) {
	requestDef := GenReqDefForListPublicIps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPublicIpsResponse), nil
	}
}

// ListPublicIpsInvoker 查询弹性公网IP列表
func (c *IecClient) ListPublicIpsInvoker(request *model.ListPublicIpsRequest) *ListPublicIpsInvoker {
	requestDef := GenReqDefForListPublicIps()
	return &ListPublicIpsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPublicIp 查询弹性公网IP
//
// 获取弹性公网IP的详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) ShowPublicIp(request *model.ShowPublicIpRequest) (*model.ShowPublicIpResponse, error) {
	requestDef := GenReqDefForShowPublicIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPublicIpResponse), nil
	}
}

// ShowPublicIpInvoker 查询弹性公网IP
func (c *IecClient) ShowPublicIpInvoker(request *model.ShowPublicIpRequest) *ShowPublicIpInvoker {
	requestDef := GenReqDefForShowPublicIp()
	return &ShowPublicIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePublicIp 更新弹性公网IP
//
// 更新弹性公网IP的信息，主要用于解绑和绑定EIP和VIP之间的关系。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) UpdatePublicIp(request *model.UpdatePublicIpRequest) (*model.UpdatePublicIpResponse, error) {
	requestDef := GenReqDefForUpdatePublicIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePublicIpResponse), nil
	}
}

// UpdatePublicIpInvoker 更新弹性公网IP
func (c *IecClient) UpdatePublicIpInvoker(request *model.UpdatePublicIpRequest) *UpdatePublicIpInvoker {
	requestDef := GenReqDefForUpdatePublicIp()
	return &UpdatePublicIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AttachVipBandwidth 端口绑定带宽
//
// IPv6虚拟IP或者IPv6私网IP绑定带宽，支持公网访问。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) AttachVipBandwidth(request *model.AttachVipBandwidthRequest) (*model.AttachVipBandwidthResponse, error) {
	requestDef := GenReqDefForAttachVipBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachVipBandwidthResponse), nil
	}
}

// AttachVipBandwidthInvoker 端口绑定带宽
func (c *IecClient) AttachVipBandwidthInvoker(request *model.AttachVipBandwidthRequest) *AttachVipBandwidthInvoker {
	requestDef := GenReqDefForAttachVipBandwidth()
	return &AttachVipBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetachVipBandwidth 端口解绑带宽
//
// IPv6虚拟IP或者IPv6私网IP解绑带宽。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) DetachVipBandwidth(request *model.DetachVipBandwidthRequest) (*model.DetachVipBandwidthResponse, error) {
	requestDef := GenReqDefForDetachVipBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetachVipBandwidthResponse), nil
	}
}

// DetachVipBandwidthInvoker 端口解绑带宽
func (c *IecClient) DetachVipBandwidthInvoker(request *model.DetachVipBandwidthRequest) *DetachVipBandwidthInvoker {
	requestDef := GenReqDefForDetachVipBandwidth()
	return &DetachVipBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSubnet 创建子网
//
// 根据用户的请求内容，创建子网。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IecClient) CreateSubnet(request *model.CreateSubnetRequest) (*model.CreateSubnetResponse, error) {
	requestDef := GenReqDefForCreateSubnet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSubnetResponse), nil
	}
}

// CreateSubnetInvoker 创建子网
func (c *IecClient) CreateSubnetInvoker(request *model.CreateSubnetRequest) *CreateSubnetInvoker {
	requestDef := GenReqDefForCreateSubnet()
	return &CreateSubnetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
