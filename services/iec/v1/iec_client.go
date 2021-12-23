package v1

import (
	http_client "github.com/RandolphCYG/huaweicloud-sdk-go-v3/core"

	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/services/iec/v1/model"
)

type IecClient struct {
	HcClient *http_client.HcHttpClient
}

func NewIecClient(hcClient *http_client.HcHttpClient) *IecClient {
	return &IecClient{HcClient: hcClient}
}

func IecClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

//添加网卡。
func (c *IecClient) AddNics(request *model.AddNicsRequest) (*model.AddNicsResponse, error) {
	requestDef := GenReqDefForAddNics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddNicsResponse), nil
	}
}

//批量重启边缘实例。
func (c *IecClient) BatchRebootInstance(request *model.BatchRebootInstanceRequest) (*model.BatchRebootInstanceResponse, error) {
	requestDef := GenReqDefForBatchRebootInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRebootInstanceResponse), nil
	}
}

//批量操作启动边缘实例。
func (c *IecClient) BatchStartInstance(request *model.BatchStartInstanceRequest) (*model.BatchStartInstanceResponse, error) {
	requestDef := GenReqDefForBatchStartInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStartInstanceResponse), nil
	}
}

//批量关闭边缘实例。
func (c *IecClient) BatchStopInstance(request *model.BatchStopInstanceRequest) (*model.BatchStopInstanceResponse, error) {
	requestDef := GenReqDefForBatchStopInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStopInstanceResponse), nil
	}
}

//切换边缘实例操作系统，支持边缘实例创建成功后，保持ip、数据盘不变的情况下重装操作系统。  调用该接口后，系统将卸载系统盘，然后使用新镜像重新创建系统盘，并挂载至实例，实现切换操作系统功能。
func (c *IecClient) ChangeOs(request *model.ChangeOsRequest) (*model.ChangeOsResponse, error) {
	requestDef := GenReqDefForChangeOs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeOsResponse), nil
	}
}

//为方便您的统一管理，以及跨边缘站点管理资源，IEC基于业务场景角度，定义了边缘业务。 边缘业务即为逻辑层面的一套资源管理集合。这里的资源主要是指计算实例，包含实例规格、镜像、硬盘、网络等方面。通过指定计算实例的数量、调度策略以及区域分布等形成一套管理集合。[了解更多](https://support.huaweicloud.com/usermanual-iec/iec_02_0301.html)  创建一个部署计划并执行，即可创建一个边缘业务。  - 边缘业务下实例分布取决于部署计划的实例分布与调度策略。 - 边缘业务下实例名称、规格、镜像等参数取决于部署计划配置计算实例字段。
func (c *IecClient) CreateDeployment(request *model.CreateDeploymentRequest) (*model.CreateDeploymentResponse, error) {
	requestDef := GenReqDefForCreateDeployment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDeploymentResponse), nil
	}
}

//创建SSH密钥，或把公钥导入系统，生成密钥对。  创建SSH密钥成功后，请把响应数据中的私钥内容保存到本地文件，用户使用该私钥登录边缘实例。为保证边缘实例安全，私钥数据只能读取一次，请妥善保管。
func (c *IecClient) CreateKeypair(request *model.CreateKeypairRequest) (*model.CreateKeypairResponse, error) {
	requestDef := GenReqDefForCreateKeypair()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateKeypairResponse), nil
	}
}

//创建端口。
func (c *IecClient) CreatePort(request *model.CreatePortRequest) (*model.CreatePortResponse, error) {
	requestDef := GenReqDefForCreatePort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePortResponse), nil
	}
}

//根据用户的请求内容，创建对应的安全组。
func (c *IecClient) CreateSecurityGroup(request *model.CreateSecurityGroupRequest) (*model.CreateSecurityGroupResponse, error) {
	requestDef := GenReqDefForCreateSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecurityGroupResponse), nil
	}
}

//根据用户的请求内容，创建安全组规则。
func (c *IecClient) CreateSecurityGroupRule(request *model.CreateSecurityGroupRuleRequest) (*model.CreateSecurityGroupRuleResponse, error) {
	requestDef := GenReqDefForCreateSecurityGroupRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecurityGroupRuleResponse), nil
	}
}

//根据用户的请求内容，创建虚拟私有云。
func (c *IecClient) CreateVpc(request *model.CreateVpcRequest) (*model.CreateVpcResponse, error) {
	requestDef := GenReqDefForCreateVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVpcResponse), nil
	}
}

//删除部署计划。
func (c *IecClient) DeleteDeployment(request *model.DeleteDeploymentRequest) (*model.DeleteDeploymentResponse, error) {
	requestDef := GenReqDefForDeleteDeployment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeploymentResponse), nil
	}
}

//删除边缘业务以及其下边缘实例。
func (c *IecClient) DeleteEdgeCloud(request *model.DeleteEdgeCloudRequest) (*model.DeleteEdgeCloudResponse, error) {
	requestDef := GenReqDefForDeleteEdgeCloud()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEdgeCloudResponse), nil
	}
}

//批量删除边缘实例。
func (c *IecClient) DeleteInstances(request *model.DeleteInstancesRequest) (*model.DeleteInstancesResponse, error) {
	requestDef := GenReqDefForDeleteInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstancesResponse), nil
	}
}

//删除密钥。
func (c *IecClient) DeleteKeypair(request *model.DeleteKeypairRequest) (*model.DeleteKeypairResponse, error) {
	requestDef := GenReqDefForDeleteKeypair()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteKeypairResponse), nil
	}
}

//删除网卡。
func (c *IecClient) DeleteNics(request *model.DeleteNicsRequest) (*model.DeleteNicsResponse, error) {
	requestDef := GenReqDefForDeleteNics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNicsResponse), nil
	}
}

//删除端口。
func (c *IecClient) DeletePort(request *model.DeletePortRequest) (*model.DeletePortResponse, error) {
	requestDef := GenReqDefForDeletePort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePortResponse), nil
	}
}

//根据安全组的ID，删除对应的安全组。
func (c *IecClient) DeleteSecurityGroup(request *model.DeleteSecurityGroupRequest) (*model.DeleteSecurityGroupResponse, error) {
	requestDef := GenReqDefForDeleteSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecurityGroupResponse), nil
	}
}

//根据安全组的ID，删除对应的安全组。
func (c *IecClient) DeleteSecurityGroupRule(request *model.DeleteSecurityGroupRuleRequest) (*model.DeleteSecurityGroupRuleResponse, error) {
	requestDef := GenReqDefForDeleteSecurityGroupRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecurityGroupRuleResponse), nil
	}
}

//根据子网的ID，删除子网。
func (c *IecClient) DeleteSubnet(request *model.DeleteSubnetRequest) (*model.DeleteSubnetResponse, error) {
	requestDef := GenReqDefForDeleteSubnet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSubnetResponse), nil
	}
}

//根据虚拟机私有云的ID，删除对应的虚拟私有云。
func (c *IecClient) DeleteVpc(request *model.DeleteVpcRequest) (*model.DeleteVpcResponse, error) {
	requestDef := GenReqDefForDeleteVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVpcResponse), nil
	}
}

//执行部署计划，创建一个边缘业务。单租户默认可创建10个边缘业务。
func (c *IecClient) ExecuteDeployment(request *model.ExecuteDeploymentRequest) (*model.ExecuteDeploymentResponse, error) {
	requestDef := GenReqDefForExecuteDeployment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteDeploymentResponse), nil
	}
}

//执行部署计划，对边缘业务进行扩容操作。
func (c *IecClient) ExpandEdgecloud(request *model.ExpandEdgecloudRequest) (*model.ExpandEdgecloudResponse, error) {
	requestDef := GenReqDefForExpandEdgecloud()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandEdgecloudResponse), nil
	}
}

//查询带宽列表。
func (c *IecClient) ListBandwidths(request *model.ListBandwidthsRequest) (*model.ListBandwidthsResponse, error) {
	requestDef := GenReqDefForListBandwidths()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBandwidthsResponse), nil
	}
}

//查询部署计划列表。
func (c *IecClient) ListDeployments(request *model.ListDeploymentsRequest) (*model.ListDeploymentsResponse, error) {
	requestDef := GenReqDefForListDeployments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDeploymentsResponse), nil
	}
}

//查询边缘业务列表。
func (c *IecClient) ListEdgeCloud(request *model.ListEdgeCloudRequest) (*model.ListEdgeCloudResponse, error) {
	requestDef := GenReqDefForListEdgeCloud()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEdgeCloudResponse), nil
	}
}

//查询边缘规格列表。
func (c *IecClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsResponse), nil
	}
}

//根据不同条件查询镜像列表，例:  -  查询已注册的私有镜像列表: visibility=private - 公共镜像: visibility=public
func (c *IecClient) ListImages(request *model.ListImagesRequest) (*model.ListImagesResponse, error) {
	requestDef := GenReqDefForListImages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImagesResponse), nil
	}
}

//查询边缘实例列表。
func (c *IecClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

//查询密钥信息列表。
func (c *IecClient) ListKeypairs(request *model.ListKeypairsRequest) (*model.ListKeypairsResponse, error) {
	requestDef := GenReqDefForListKeypairs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKeypairsResponse), nil
	}
}

//查询端口的列表信息
func (c *IecClient) ListPorts(request *model.ListPortsRequest) (*model.ListPortsResponse, error) {
	requestDef := GenReqDefForListPorts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPortsResponse), nil
	}
}

//查询租户资源配额。
func (c *IecClient) ListQuota(request *model.ListQuotaRequest) (*model.ListQuotaResponse, error) {
	requestDef := GenReqDefForListQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotaResponse), nil
	}
}

//根据用户的查询条件，获取安全组规则的列表信息。
func (c *IecClient) ListSecurityGroupRules(request *model.ListSecurityGroupRulesRequest) (*model.ListSecurityGroupRulesResponse, error) {
	requestDef := GenReqDefForListSecurityGroupRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityGroupRulesResponse), nil
	}
}

//根据特定查询条件，获取安全组的列表信息。
func (c *IecClient) ListSecurityGroups(request *model.ListSecurityGroupsRequest) (*model.ListSecurityGroupsResponse, error) {
	requestDef := GenReqDefForListSecurityGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityGroupsResponse), nil
	}
}

//查询边缘站点列表。  - 边缘站点：靠近终端应用的位置，基于一个或多个运营商建立的一个城市级站点。边缘站点提供物理隔离的资源池，提供多元算力、存储和网络的能力。用户可以将业务灵活就近部署在边缘站点上，以降低网络时延和成本。 - 边缘区域：为依据边缘站点的物理位置划分的区域，一个边缘区域包含多个相靠近的边缘站点的集合。IEC当前提供城市级、省级和大区级三个分布层级的边缘区域。
func (c *IecClient) ListSites(request *model.ListSitesRequest) (*model.ListSitesResponse, error) {
	requestDef := GenReqDefForListSites()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSitesResponse), nil
	}
}

//根据查询条件获取子网的列表信息。
func (c *IecClient) ListSubnets(request *model.ListSubnetsRequest) (*model.ListSubnetsResponse, error) {
	requestDef := GenReqDefForListSubnets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubnetsResponse), nil
	}
}

//获取虚拟私有云的列表。
func (c *IecClient) ListVpcs(request *model.ListVpcsRequest) (*model.ListVpcsResponse, error) {
	requestDef := GenReqDefForListVpcs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpcsResponse), nil
	}
}

//查询带宽详情。
func (c *IecClient) ShowBandwidth(request *model.ShowBandwidthRequest) (*model.ShowBandwidthResponse, error) {
	requestDef := GenReqDefForShowBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBandwidthResponse), nil
	}
}

//查询边缘业务详情。
func (c *IecClient) ShowEdgeCloud(request *model.ShowEdgeCloudRequest) (*model.ShowEdgeCloudResponse, error) {
	requestDef := GenReqDefForShowEdgeCloud()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEdgeCloudResponse), nil
	}
}

//查询镜像详情。
func (c *IecClient) ShowImage(request *model.ShowImageRequest) (*model.ShowImageResponse, error) {
	requestDef := GenReqDefForShowImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowImageResponse), nil
	}
}

//查询边缘实例详情。
func (c *IecClient) ShowInstance(request *model.ShowInstanceRequest) (*model.ShowInstanceResponse, error) {
	requestDef := GenReqDefForShowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceResponse), nil
	}
}

//查询密钥信息列表。
func (c *IecClient) ShowKeypair(request *model.ShowKeypairRequest) (*model.ShowKeypairResponse, error) {
	requestDef := GenReqDefForShowKeypair()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKeypairResponse), nil
	}
}

//根据端口的ID，获取端口的详细信息。
func (c *IecClient) ShowPort(request *model.ShowPortRequest) (*model.ShowPortResponse, error) {
	requestDef := GenReqDefForShowPort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPortResponse), nil
	}
}

//根据安全组的ID，获取特定安全组的详细信息。
func (c *IecClient) ShowSecurityGroup(request *model.ShowSecurityGroupRequest) (*model.ShowSecurityGroupResponse, error) {
	requestDef := GenReqDefForShowSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecurityGroupResponse), nil
	}
}

//根据安全组规则的ID，获取安全组规则的详细信息。
func (c *IecClient) ShowSecurityGroupRule(request *model.ShowSecurityGroupRuleRequest) (*model.ShowSecurityGroupRuleResponse, error) {
	requestDef := GenReqDefForShowSecurityGroupRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecurityGroupRuleResponse), nil
	}
}

//根据子网的ID，获取子网的详细信息。
func (c *IecClient) ShowSubnet(request *model.ShowSubnetRequest) (*model.ShowSubnetResponse, error) {
	requestDef := GenReqDefForShowSubnet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSubnetResponse), nil
	}
}

//查询硬盘详情。
func (c *IecClient) ShowVolume(request *model.ShowVolumeRequest) (*model.ShowVolumeResponse, error) {
	requestDef := GenReqDefForShowVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVolumeResponse), nil
	}
}

//根据虚拟私有云ID，获取虚拟私有云的详情。
func (c *IecClient) ShowVpc(request *model.ShowVpcRequest) (*model.ShowVpcResponse, error) {
	requestDef := GenReqDefForShowVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVpcResponse), nil
	}
}

//修改边缘实例。
func (c *IecClient) UpdateInstance(request *model.UpdateInstanceRequest) (*model.UpdateInstanceResponse, error) {
	requestDef := GenReqDefForUpdateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceResponse), nil
	}
}

//更新端口。
func (c *IecClient) UpdatePort(request *model.UpdatePortRequest) (*model.UpdatePortResponse, error) {
	requestDef := GenReqDefForUpdatePort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePortResponse), nil
	}
}

//更新子网的基本信息。
func (c *IecClient) UpdateSubnet(request *model.UpdateSubnetRequest) (*model.UpdateSubnetResponse, error) {
	requestDef := GenReqDefForUpdateSubnet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSubnetResponse), nil
	}
}

//更新虚拟私有云的信息
func (c *IecClient) UpdateVpc(request *model.UpdateVpcRequest) (*model.UpdateVpcResponse, error) {
	requestDef := GenReqDefForUpdateVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVpcResponse), nil
	}
}

//创建网络ACL。
func (c *IecClient) CreateFirewall(request *model.CreateFirewallRequest) (*model.CreateFirewallResponse, error) {
	requestDef := GenReqDefForCreateFirewall()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFirewallResponse), nil
	}
}

//删除网络ACL。
func (c *IecClient) DeleteFirewall(request *model.DeleteFirewallRequest) (*model.DeleteFirewallResponse, error) {
	requestDef := GenReqDefForDeleteFirewall()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFirewallResponse), nil
	}
}

//查询网络ACL列表。
func (c *IecClient) ListFirewalls(request *model.ListFirewallsRequest) (*model.ListFirewallsResponse, error) {
	requestDef := GenReqDefForListFirewalls()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFirewallsResponse), nil
	}
}

//查询网络ACL详情。
func (c *IecClient) ShowFirewall(request *model.ShowFirewallRequest) (*model.ShowFirewallResponse, error) {
	requestDef := GenReqDefForShowFirewall()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFirewallResponse), nil
	}
}

//更新网络ACL。
func (c *IecClient) UpdateFirewall(request *model.UpdateFirewallRequest) (*model.UpdateFirewallResponse, error) {
	requestDef := GenReqDefForUpdateFirewall()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFirewallResponse), nil
	}
}

//更新网络ACL规则。
func (c *IecClient) UpdateFirewallRule(request *model.UpdateFirewallRuleRequest) (*model.UpdateFirewallRuleResponse, error) {
	requestDef := GenReqDefForUpdateFirewallRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFirewallRuleResponse), nil
	}
}

//根据用户的请求内容，创建弹性公网IP
func (c *IecClient) CreatePublicIp(request *model.CreatePublicIpRequest) (*model.CreatePublicIpResponse, error) {
	requestDef := GenReqDefForCreatePublicIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePublicIpResponse), nil
	}
}

//根据弹性公网IP的ID，删除对应的弹性公网IP。
func (c *IecClient) DeletePublicIp(request *model.DeletePublicIpRequest) (*model.DeletePublicIpResponse, error) {
	requestDef := GenReqDefForDeletePublicIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePublicIpResponse), nil
	}
}

//获取弹性公网IP列表信息。
func (c *IecClient) ListPublicIps(request *model.ListPublicIpsRequest) (*model.ListPublicIpsResponse, error) {
	requestDef := GenReqDefForListPublicIps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPublicIpsResponse), nil
	}
}

//获取弹性公网IP的详情信息。
func (c *IecClient) ShowPublicIp(request *model.ShowPublicIpRequest) (*model.ShowPublicIpResponse, error) {
	requestDef := GenReqDefForShowPublicIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPublicIpResponse), nil
	}
}

//更新弹性公网IP的信息，主要用于解绑和绑定EIP和VIP之间的关系。
func (c *IecClient) UpdatePublicIp(request *model.UpdatePublicIpRequest) (*model.UpdatePublicIpResponse, error) {
	requestDef := GenReqDefForUpdatePublicIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePublicIpResponse), nil
	}
}
