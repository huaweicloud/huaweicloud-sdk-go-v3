package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/model"
)

type VpcClient struct {
	hcClient *http_client.HcHttpClient
}

func NewVpcClient(hcClient *http_client.HcHttpClient) *VpcClient {
	return &VpcClient{hcClient: hcClient}
}

func VpcClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//租户A名下的VPC申请和租户B的VPC建立对等连接，需要等待租户B接受该请求。此接口用于租户接受其他租户发起的对等连接请求。
func (c *VpcClient) AcceptVpcPeering(request *model.AcceptVpcPeeringRequest) (*model.AcceptVpcPeeringResponse, error) {
	requestDef := GenReqDefForAcceptVpcPeering(request)
	resp, responseDef := GenRespForAcceptVpcPeering()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//创建端口。
func (c *VpcClient) CreatePort(request *model.CreatePortRequest) (*model.CreatePortResponse, error) {
	requestDef := GenReqDefForCreatePort(request)
	resp, responseDef := GenRespForCreatePort()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//创建安全组。
func (c *VpcClient) CreateSecurityGroup(request *model.CreateSecurityGroupRequest) (*model.CreateSecurityGroupResponse, error) {
	requestDef := GenReqDefForCreateSecurityGroup(request)
	resp, responseDef := GenRespForCreateSecurityGroup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//创建安全组规则。
func (c *VpcClient) CreateSecurityGroupRule(request *model.CreateSecurityGroupRuleRequest) (*model.CreateSecurityGroupRuleResponse, error) {
	requestDef := GenReqDefForCreateSecurityGroupRule(request)
	resp, responseDef := GenRespForCreateSecurityGroupRule()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//创建子网。
func (c *VpcClient) CreateSubnet(request *model.CreateSubnetRequest) (*model.CreateSubnetResponse, error) {
	requestDef := GenReqDefForCreateSubnet(request)
	resp, responseDef := GenRespForCreateSubnet()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//创建对等连接。
func (c *VpcClient) CreateVpcPeering(request *model.CreateVpcPeeringRequest) (*model.CreateVpcPeeringResponse, error) {
	requestDef := GenReqDefForCreateVpcPeering(request)
	resp, responseDef := GenRespForCreateVpcPeering()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除端口。
func (c *VpcClient) DeletePort(request *model.DeletePortRequest) (*model.DeletePortResponse, error) {
	requestDef := GenReqDefForDeletePort(request)
	resp, responseDef := GenRespForDeletePort()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除安全组。
func (c *VpcClient) DeleteSecurityGroup(request *model.DeleteSecurityGroupRequest) (*model.DeleteSecurityGroupResponse, error) {
	requestDef := GenReqDefForDeleteSecurityGroup(request)
	resp, responseDef := GenRespForDeleteSecurityGroup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除安全组规则。
func (c *VpcClient) DeleteSecurityGroupRule(request *model.DeleteSecurityGroupRuleRequest) (*model.DeleteSecurityGroupRuleResponse, error) {
	requestDef := GenReqDefForDeleteSecurityGroupRule(request)
	resp, responseDef := GenRespForDeleteSecurityGroupRule()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除子网
func (c *VpcClient) DeleteSubnet(request *model.DeleteSubnetRequest) (*model.DeleteSubnetResponse, error) {
	requestDef := GenReqDefForDeleteSubnet(request)
	resp, responseDef := GenRespForDeleteSubnet()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除对等连接。 可以在在本端或对端任何一端删除对等连接。
func (c *VpcClient) DeleteVpcPeering(request *model.DeleteVpcPeeringRequest) (*model.DeleteVpcPeeringResponse, error) {
	requestDef := GenReqDefForDeleteVpcPeering(request)
	resp, responseDef := GenRespForDeleteVpcPeering()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询提交请求的租户的所有端口，单次查询最多返回2000条数据。
func (c *VpcClient) ListPorts(request *model.ListPortsRequest) (*model.ListPortsResponse, error) {
	requestDef := GenReqDefForListPorts(request)
	resp, responseDef := GenRespForListPorts()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询安全组规则列表。
func (c *VpcClient) ListSecurityGroupRules(request *model.ListSecurityGroupRulesRequest) (*model.ListSecurityGroupRulesResponse, error) {
	requestDef := GenReqDefForListSecurityGroupRules(request)
	resp, responseDef := GenRespForListSecurityGroupRules()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询安全组列表
func (c *VpcClient) ListSecurityGroups(request *model.ListSecurityGroupsRequest) (*model.ListSecurityGroupsResponse, error) {
	requestDef := GenReqDefForListSecurityGroups(request)
	resp, responseDef := GenRespForListSecurityGroups()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询子网列表
func (c *VpcClient) ListSubnets(request *model.ListSubnetsRequest) (*model.ListSubnetsResponse, error) {
	requestDef := GenReqDefForListSubnets(request)
	resp, responseDef := GenRespForListSubnets()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询提交请求的租户的所有对等连接。根据过滤条件进行过滤。
func (c *VpcClient) ListVpcPeerings(request *model.ListVpcPeeringsRequest) (*model.ListVpcPeeringsResponse, error) {
	requestDef := GenReqDefForListVpcPeerings(request)
	resp, responseDef := GenRespForListVpcPeerings()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//租户A名下的VPC申请和租户B的VPC建立对等连接，需要等待租户B接受该请求。此接口用于租户拒绝其他租户发起的对等连接请求。
func (c *VpcClient) RejectVpcPeering(request *model.RejectVpcPeeringRequest) (*model.RejectVpcPeeringResponse, error) {
	requestDef := GenReqDefForRejectVpcPeering(request)
	resp, responseDef := GenRespForRejectVpcPeering()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询单个端口详情。
func (c *VpcClient) ShowPort(request *model.ShowPortRequest) (*model.ShowPortResponse, error) {
	requestDef := GenReqDefForShowPort(request)
	resp, responseDef := GenRespForShowPort()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询单租户在VPC服务下的网络资源配额，包括vpc配额、子网配额、安全组配额、安全组规则配额、弹性公网IP配额，vpn配额等。
func (c *VpcClient) ShowQuota(request *model.ShowQuotaRequest) (*model.ShowQuotaResponse, error) {
	requestDef := GenReqDefForShowQuota(request)
	resp, responseDef := GenRespForShowQuota()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询单个安全组详情。
func (c *VpcClient) ShowSecurityGroup(request *model.ShowSecurityGroupRequest) (*model.ShowSecurityGroupResponse, error) {
	requestDef := GenReqDefForShowSecurityGroup(request)
	resp, responseDef := GenRespForShowSecurityGroup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询单个安全组规则详情
func (c *VpcClient) ShowSecurityGroupRule(request *model.ShowSecurityGroupRuleRequest) (*model.ShowSecurityGroupRuleResponse, error) {
	requestDef := GenReqDefForShowSecurityGroupRule(request)
	resp, responseDef := GenRespForShowSecurityGroupRule()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询子网详情。
func (c *VpcClient) ShowSubnet(request *model.ShowSubnetRequest) (*model.ShowSubnetResponse, error) {
	requestDef := GenReqDefForShowSubnet(request)
	resp, responseDef := GenRespForShowSubnet()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询对等连接详情。
func (c *VpcClient) ShowVpcPeering(request *model.ShowVpcPeeringRequest) (*model.ShowVpcPeeringResponse, error) {
	requestDef := GenReqDefForShowVpcPeering(request)
	resp, responseDef := GenRespForShowVpcPeering()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//更新端口。
func (c *VpcClient) UpdatePort(request *model.UpdatePortRequest) (*model.UpdatePortResponse, error) {
	requestDef := GenReqDefForUpdatePort(request)
	resp, responseDef := GenRespForUpdatePort()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//更新子网。
func (c *VpcClient) UpdateSubnet(request *model.UpdateSubnetRequest) (*model.UpdateSubnetResponse, error) {
	requestDef := GenReqDefForUpdateSubnet(request)
	resp, responseDef := GenRespForUpdateSubnet()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//更新对等连接。
func (c *VpcClient) UpdateVpcPeering(request *model.UpdateVpcPeeringRequest) (*model.UpdateVpcPeeringResponse, error) {
	requestDef := GenReqDefForUpdateVpcPeering(request)
	resp, responseDef := GenRespForUpdateVpcPeering()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//申请私有IP。
func (c *VpcClient) CreatePrivateip(request *model.CreatePrivateipRequest) (*model.CreatePrivateipResponse, error) {
	requestDef := GenReqDefForCreatePrivateip(request)
	resp, responseDef := GenRespForCreatePrivateip()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除私有IP。
func (c *VpcClient) DeletePrivateip(request *model.DeletePrivateipRequest) (*model.DeletePrivateipResponse, error) {
	requestDef := GenReqDefForDeletePrivateip(request)
	resp, responseDef := GenRespForDeletePrivateip()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询指定子网下的私有IP列表。
func (c *VpcClient) ListPrivateips(request *model.ListPrivateipsRequest) (*model.ListPrivateipsResponse, error) {
	requestDef := GenReqDefForListPrivateips(request)
	resp, responseDef := GenRespForListPrivateips()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//显示一个指定网络中的IPv4地址使用情况。 包括此网络中的IP总数以及已用IP总数，以及网络下每一个子网的IP地址总数和可用IP地址总数。  > 须知  - 系统预留地址指的是子网的第1个以及最后4个地址，一般用于网关、DHCP等服务。 - 这里以及下文描述的IP地址总数、已用IP地址总数不包含系统预留地址。 - 在分配IP时，用户可以指定系统预留的IP地址。但是不论IP是如何分配的，只要是处于系统预留IP地址段的IP均不会被统计到已用IP地址数目和IP地址总数中。
func (c *VpcClient) ShowNetworkIpAvailabilities(request *model.ShowNetworkIpAvailabilitiesRequest) (*model.ShowNetworkIpAvailabilitiesResponse, error) {
	requestDef := GenReqDefForShowNetworkIpAvailabilities(request)
	resp, responseDef := GenRespForShowNetworkIpAvailabilities()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//指定ID查询私有IP。
func (c *VpcClient) ShowPrivateip(request *model.ShowPrivateipRequest) (*model.ShowPrivateipResponse, error) {
	requestDef := GenReqDefForShowPrivateip(request)
	resp, responseDef := GenRespForShowPrivateip()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//创建虚拟私有云。
func (c *VpcClient) CreateVpc(request *model.CreateVpcRequest) (*model.CreateVpcResponse, error) {
	requestDef := GenReqDefForCreateVpc(request)
	resp, responseDef := GenRespForCreateVpc()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//创建路由
func (c *VpcClient) CreateVpcRoute(request *model.CreateVpcRouteRequest) (*model.CreateVpcRouteResponse, error) {
	requestDef := GenReqDefForCreateVpcRoute(request)
	resp, responseDef := GenRespForCreateVpcRoute()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除虚拟私有云。
func (c *VpcClient) DeleteVpc(request *model.DeleteVpcRequest) (*model.DeleteVpcResponse, error) {
	requestDef := GenReqDefForDeleteVpc(request)
	resp, responseDef := GenRespForDeleteVpc()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除路由
func (c *VpcClient) DeleteVpcRoute(request *model.DeleteVpcRouteRequest) (*model.DeleteVpcRouteResponse, error) {
	requestDef := GenReqDefForDeleteVpcRoute(request)
	resp, responseDef := GenRespForDeleteVpcRoute()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询提交请求的租户的所有路由列表，并根据过滤条件进行过滤。
func (c *VpcClient) ListVpcRoutes(request *model.ListVpcRoutesRequest) (*model.ListVpcRoutesResponse, error) {
	requestDef := GenReqDefForListVpcRoutes(request)
	resp, responseDef := GenRespForListVpcRoutes()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询虚拟私有云列表。
func (c *VpcClient) ListVpcs(request *model.ListVpcsRequest) (*model.ListVpcsResponse, error) {
	requestDef := GenReqDefForListVpcs(request)
	resp, responseDef := GenRespForListVpcs()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询虚拟私有云。
func (c *VpcClient) ShowVpc(request *model.ShowVpcRequest) (*model.ShowVpcResponse, error) {
	requestDef := GenReqDefForShowVpc(request)
	resp, responseDef := GenRespForShowVpc()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询路由详情
func (c *VpcClient) ShowVpcRoute(request *model.ShowVpcRouteRequest) (*model.ShowVpcRouteResponse, error) {
	requestDef := GenReqDefForShowVpcRoute(request)
	resp, responseDef := GenRespForShowVpcRoute()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//更新虚拟私有云。
func (c *VpcClient) UpdateVpc(request *model.UpdateVpcRequest) (*model.UpdateVpcResponse, error) {
	requestDef := GenReqDefForUpdateVpc(request)
	resp, responseDef := GenRespForUpdateVpc()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}
