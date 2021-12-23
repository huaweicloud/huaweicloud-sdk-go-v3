package v3

import (
	http_client "github.com/RandolphCYG/huaweicloud-sdk-go-v3/core"

	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/services/elb/v3/model"
)

type ElbClient struct {
	HcClient *http_client.HcHttpClient
}

func NewElbClient(hcClient *http_client.HcHttpClient) *ElbClient {
	return &ElbClient{HcClient: hcClient}
}

func ElbClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//批量更新转发策略的优先级。
func (c *ElbClient) BatchUpdatePoliciesPriority(request *model.BatchUpdatePoliciesPriorityRequest) (*model.BatchUpdatePoliciesPriorityResponse, error) {
	requestDef := GenReqDefForBatchUpdatePoliciesPriority()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdatePoliciesPriorityResponse), nil
	}
}

//负载均衡器计费模式变更，当前只支持按需计费转包周期计费。
func (c *ElbClient) ChangeLoadbalancerChargeMode(request *model.ChangeLoadbalancerChargeModeRequest) (*model.ChangeLoadbalancerChargeModeResponse, error) {
	requestDef := GenReqDefForChangeLoadbalancerChargeMode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeLoadbalancerChargeModeResponse), nil
	}
}

//创建证书。用于HTTPS协议监听器。
func (c *ElbClient) CreateCertificate(request *model.CreateCertificateRequest) (*model.CreateCertificateResponse, error) {
	requestDef := GenReqDefForCreateCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCertificateResponse), nil
	}
}

//创建健康检查。
func (c *ElbClient) CreateHealthMonitor(request *model.CreateHealthMonitorRequest) (*model.CreateHealthMonitorResponse, error) {
	requestDef := GenReqDefForCreateHealthMonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHealthMonitorResponse), nil
	}
}

//创建七层转发策略。
func (c *ElbClient) CreateL7Policy(request *model.CreateL7PolicyRequest) (*model.CreateL7PolicyResponse, error) {
	requestDef := GenReqDefForCreateL7Policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateL7PolicyResponse), nil
	}
}

//创建七层转发规则。
func (c *ElbClient) CreateL7Rule(request *model.CreateL7RuleRequest) (*model.CreateL7RuleResponse, error) {
	requestDef := GenReqDefForCreateL7Rule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateL7RuleResponse), nil
	}
}

//创建监听器。
func (c *ElbClient) CreateListener(request *model.CreateListenerRequest) (*model.CreateListenerResponse, error) {
	requestDef := GenReqDefForCreateListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateListenerResponse), nil
	}
}

//创建负载均衡器。  1.若要创建内网IPv4负载均衡器，则需要设置vip_subnet_cidr_id和vip_address。  2.若要创建公网IPv4负载均衡器，则需要设置publicip，以及设置vpc_id和vip_subnet_cidr_id这两个参数中的一个。  3.若要绑定有已有公网IPv4地址，需要设置publicip_ids，以及设置vpc_id和vip_subnet_cidr_id这两个参数中的一个。  4.若要创建内网双栈负载均衡器，则需要设置ipv6_vip_virsubnet_id。  5.若要创建公网双栈负载均衡器，则需要设置ipv6_vip_virsubnet_id和ipv6_bandwidth。  6.不支持绑定已有未使用的内网IPv4、内网IPv6或公网IPv6地址。  [>不支持创建IPv6地址负载均衡器](tag:otc,otc_test,dt,dt_test)
func (c *ElbClient) CreateLoadBalancer(request *model.CreateLoadBalancerRequest) (*model.CreateLoadBalancerResponse, error) {
	requestDef := GenReqDefForCreateLoadBalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLoadBalancerResponse), nil
	}
}

//创建后端服务器。
func (c *ElbClient) CreateMember(request *model.CreateMemberRequest) (*model.CreateMemberResponse, error) {
	requestDef := GenReqDefForCreateMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMemberResponse), nil
	}
}

//创建后端服务器组。
func (c *ElbClient) CreatePool(request *model.CreatePoolRequest) (*model.CreatePoolResponse, error) {
	requestDef := GenReqDefForCreatePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePoolResponse), nil
	}
}

//创建自定义安全策略。用于在创建HTTPS监听器时，请求参数中指定security_policy_id来设置监听器的自定义安全策略。
func (c *ElbClient) CreateSecurityPolicy(request *model.CreateSecurityPolicyRequest) (*model.CreateSecurityPolicyResponse, error) {
	requestDef := GenReqDefForCreateSecurityPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecurityPolicyResponse), nil
	}
}

//删除证书。
func (c *ElbClient) DeleteCertificate(request *model.DeleteCertificateRequest) (*model.DeleteCertificateResponse, error) {
	requestDef := GenReqDefForDeleteCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCertificateResponse), nil
	}
}

//删除健康检查。
func (c *ElbClient) DeleteHealthMonitor(request *model.DeleteHealthMonitorRequest) (*model.DeleteHealthMonitorResponse, error) {
	requestDef := GenReqDefForDeleteHealthMonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHealthMonitorResponse), nil
	}
}

//删除七层转发策略。
func (c *ElbClient) DeleteL7Policy(request *model.DeleteL7PolicyRequest) (*model.DeleteL7PolicyResponse, error) {
	requestDef := GenReqDefForDeleteL7Policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteL7PolicyResponse), nil
	}
}

//删除七层转发规则。
func (c *ElbClient) DeleteL7Rule(request *model.DeleteL7RuleRequest) (*model.DeleteL7RuleResponse, error) {
	requestDef := GenReqDefForDeleteL7Rule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteL7RuleResponse), nil
	}
}

//删除监听器。
func (c *ElbClient) DeleteListener(request *model.DeleteListenerRequest) (*model.DeleteListenerResponse, error) {
	requestDef := GenReqDefForDeleteListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteListenerResponse), nil
	}
}

//删除负载均衡器。
func (c *ElbClient) DeleteLoadBalancer(request *model.DeleteLoadBalancerRequest) (*model.DeleteLoadBalancerResponse, error) {
	requestDef := GenReqDefForDeleteLoadBalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLoadBalancerResponse), nil
	}
}

//删除后端服务器。
func (c *ElbClient) DeleteMember(request *model.DeleteMemberRequest) (*model.DeleteMemberResponse, error) {
	requestDef := GenReqDefForDeleteMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMemberResponse), nil
	}
}

//删除后端服务器组。
func (c *ElbClient) DeletePool(request *model.DeletePoolRequest) (*model.DeletePoolResponse, error) {
	requestDef := GenReqDefForDeletePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePoolResponse), nil
	}
}

//删除自定义安全策略。
func (c *ElbClient) DeleteSecurityPolicy(request *model.DeleteSecurityPolicyRequest) (*model.DeleteSecurityPolicyResponse, error) {
	requestDef := GenReqDefForDeleteSecurityPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecurityPolicyResponse), nil
	}
}

//查询当前租户下的后端服务器列表。
func (c *ElbClient) ListAllMembers(request *model.ListAllMembersRequest) (*model.ListAllMembersResponse, error) {
	requestDef := GenReqDefForListAllMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllMembersResponse), nil
	}
}

//返回租户创建LB时可使用的可用区集合列表情况。  默认情况下，会返回一个可用区集合。在（如创建LB）设置可用区时，填写的可用区必须包含在可用区集合中、为这个可用区集合的子集。  [特殊场景下，部分客户要求负载均衡只能创建在指定可用区集合中，此时会返回客户定制的可用区集合。返回可用区集合可能为一个也可能为多个，比如列表有两个可用区集合[az1,az2], [az2,az3]。在创建负载均衡器时，可以选择创建在多个可用区，但所选的多个可用区必须同属于其中一个可用区集合，如可以选az2和az3，但不能选择az1和az3。你可以选择多个可用区，只要这些可用区在一个子集中](tag:hc,hws,hk,ocb,tlf,ctc,hcso,sbc,g42,tm,cmcc,hk-g42)
func (c *ElbClient) ListAvailabilityZones(request *model.ListAvailabilityZonesRequest) (*model.ListAvailabilityZonesResponse, error) {
	requestDef := GenReqDefForListAvailabilityZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailabilityZonesResponse), nil
	}
}

//查询证书列表。
func (c *ElbClient) ListCertificates(request *model.ListCertificatesRequest) (*model.ListCertificatesResponse, error) {
	requestDef := GenReqDefForListCertificates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCertificatesResponse), nil
	}
}

//查询租户在当前region下可用的负载均衡规格列表。
func (c *ElbClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsResponse), nil
	}
}

//健康检查列表。
func (c *ElbClient) ListHealthMonitors(request *model.ListHealthMonitorsRequest) (*model.ListHealthMonitorsResponse, error) {
	requestDef := GenReqDefForListHealthMonitors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHealthMonitorsResponse), nil
	}
}

//查询七层转发策略列表。
func (c *ElbClient) ListL7Policies(request *model.ListL7PoliciesRequest) (*model.ListL7PoliciesResponse, error) {
	requestDef := GenReqDefForListL7Policies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListL7PoliciesResponse), nil
	}
}

//查询转发规则列表。
func (c *ElbClient) ListL7Rules(request *model.ListL7RulesRequest) (*model.ListL7RulesResponse, error) {
	requestDef := GenReqDefForListL7Rules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListL7RulesResponse), nil
	}
}

//查询监听器列表。
func (c *ElbClient) ListListeners(request *model.ListListenersRequest) (*model.ListListenersResponse, error) {
	requestDef := GenReqDefForListListeners()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListListenersResponse), nil
	}
}

//查询负载均衡器列表。
func (c *ElbClient) ListLoadBalancers(request *model.ListLoadBalancersRequest) (*model.ListLoadBalancersResponse, error) {
	requestDef := GenReqDefForListLoadBalancers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLoadBalancersResponse), nil
	}
}

//Pool下的后端服务器列表。
func (c *ElbClient) ListMembers(request *model.ListMembersRequest) (*model.ListMembersResponse, error) {
	requestDef := GenReqDefForListMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMembersResponse), nil
	}
}

//后端服务器组列表。
func (c *ElbClient) ListPools(request *model.ListPoolsRequest) (*model.ListPoolsResponse, error) {
	requestDef := GenReqDefForListPools()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPoolsResponse), nil
	}
}

//查询指定项目中负载均衡相关的各类资源的当前配额和已使用配额信息。
func (c *ElbClient) ListQuotaDetails(request *model.ListQuotaDetailsRequest) (*model.ListQuotaDetailsResponse, error) {
	requestDef := GenReqDefForListQuotaDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotaDetailsResponse), nil
	}
}

//查询自定义安全策略列表。
func (c *ElbClient) ListSecurityPolicies(request *model.ListSecurityPoliciesRequest) (*model.ListSecurityPoliciesResponse, error) {
	requestDef := GenReqDefForListSecurityPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityPoliciesResponse), nil
	}
}

//查询系统安全策略列表。  系统安全策略为预置的所有租户通用的安全策略，租户不可新增或修改。
func (c *ElbClient) ListSystemSecurityPolicies(request *model.ListSystemSecurityPoliciesRequest) (*model.ListSystemSecurityPoliciesResponse, error) {
	requestDef := GenReqDefForListSystemSecurityPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSystemSecurityPoliciesResponse), nil
	}
}

//查询证书详情。
func (c *ElbClient) ShowCertificate(request *model.ShowCertificateRequest) (*model.ShowCertificateResponse, error) {
	requestDef := GenReqDefForShowCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCertificateResponse), nil
	}
}

//查询规格的详情。
func (c *ElbClient) ShowFlavor(request *model.ShowFlavorRequest) (*model.ShowFlavorResponse, error) {
	requestDef := GenReqDefForShowFlavor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFlavorResponse), nil
	}
}

//查询健康检查详情。
func (c *ElbClient) ShowHealthMonitor(request *model.ShowHealthMonitorRequest) (*model.ShowHealthMonitorResponse, error) {
	requestDef := GenReqDefForShowHealthMonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHealthMonitorResponse), nil
	}
}

//查询七层转发策略详情。
func (c *ElbClient) ShowL7Policy(request *model.ShowL7PolicyRequest) (*model.ShowL7PolicyResponse, error) {
	requestDef := GenReqDefForShowL7Policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowL7PolicyResponse), nil
	}
}

//查询七层转发规则详情。
func (c *ElbClient) ShowL7Rule(request *model.ShowL7RuleRequest) (*model.ShowL7RuleResponse, error) {
	requestDef := GenReqDefForShowL7Rule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowL7RuleResponse), nil
	}
}

//监听器详情。
func (c *ElbClient) ShowListener(request *model.ShowListenerRequest) (*model.ShowListenerResponse, error) {
	requestDef := GenReqDefForShowListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowListenerResponse), nil
	}
}

//查询负载均衡器详情。
func (c *ElbClient) ShowLoadBalancer(request *model.ShowLoadBalancerRequest) (*model.ShowLoadBalancerResponse, error) {
	requestDef := GenReqDefForShowLoadBalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLoadBalancerResponse), nil
	}
}

//查询负载均衡器状态树，包括负载均衡器及其关联的子资源的状态信息。 注意：该接口中的operating_status不一定与对应资源的operating_status相同。如：当Member的admin_state_up=false且operating_status=OFFLINE时，该接口返回member的operating_status=DISABLE。
func (c *ElbClient) ShowLoadBalancerStatus(request *model.ShowLoadBalancerStatusRequest) (*model.ShowLoadBalancerStatusResponse, error) {
	requestDef := GenReqDefForShowLoadBalancerStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLoadBalancerStatusResponse), nil
	}
}

//后端服务器详情。
func (c *ElbClient) ShowMember(request *model.ShowMemberRequest) (*model.ShowMemberResponse, error) {
	requestDef := GenReqDefForShowMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMemberResponse), nil
	}
}

//后端服务器组详情。
func (c *ElbClient) ShowPool(request *model.ShowPoolRequest) (*model.ShowPoolResponse, error) {
	requestDef := GenReqDefForShowPool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPoolResponse), nil
	}
}

//查询指定项目中负载均衡相关的各类资源的当前配额。
func (c *ElbClient) ShowQuota(request *model.ShowQuotaRequest) (*model.ShowQuotaResponse, error) {
	requestDef := GenReqDefForShowQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotaResponse), nil
	}
}

//查询自定义安全策略详情。
func (c *ElbClient) ShowSecurityPolicy(request *model.ShowSecurityPolicyRequest) (*model.ShowSecurityPolicyResponse, error) {
	requestDef := GenReqDefForShowSecurityPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecurityPolicyResponse), nil
	}
}

//更新证书。
func (c *ElbClient) UpdateCertificate(request *model.UpdateCertificateRequest) (*model.UpdateCertificateResponse, error) {
	requestDef := GenReqDefForUpdateCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCertificateResponse), nil
	}
}

//更新健康检查。
func (c *ElbClient) UpdateHealthMonitor(request *model.UpdateHealthMonitorRequest) (*model.UpdateHealthMonitorResponse, error) {
	requestDef := GenReqDefForUpdateHealthMonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHealthMonitorResponse), nil
	}
}

//更新七层转发策略。
func (c *ElbClient) UpdateL7Policy(request *model.UpdateL7PolicyRequest) (*model.UpdateL7PolicyResponse, error) {
	requestDef := GenReqDefForUpdateL7Policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateL7PolicyResponse), nil
	}
}

//更新七层转发规则。
func (c *ElbClient) UpdateL7Rule(request *model.UpdateL7RuleRequest) (*model.UpdateL7RuleResponse, error) {
	requestDef := GenReqDefForUpdateL7Rule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateL7RuleResponse), nil
	}
}

//更新监听器。
func (c *ElbClient) UpdateListener(request *model.UpdateListenerRequest) (*model.UpdateListenerResponse, error) {
	requestDef := GenReqDefForUpdateListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateListenerResponse), nil
	}
}

//更新负载均衡器。
func (c *ElbClient) UpdateLoadBalancer(request *model.UpdateLoadBalancerRequest) (*model.UpdateLoadBalancerResponse, error) {
	requestDef := GenReqDefForUpdateLoadBalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLoadBalancerResponse), nil
	}
}

//更新后端服务器。
func (c *ElbClient) UpdateMember(request *model.UpdateMemberRequest) (*model.UpdateMemberResponse, error) {
	requestDef := GenReqDefForUpdateMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMemberResponse), nil
	}
}

//更新后端服务器组。
func (c *ElbClient) UpdatePool(request *model.UpdatePoolRequest) (*model.UpdatePoolResponse, error) {
	requestDef := GenReqDefForUpdatePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePoolResponse), nil
	}
}

//更新自定义安全策略。
func (c *ElbClient) UpdateSecurityPolicy(request *model.UpdateSecurityPolicyRequest) (*model.UpdateSecurityPolicyResponse, error) {
	requestDef := GenReqDefForUpdateSecurityPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSecurityPolicyResponse), nil
	}
}

//返回ELB当前所有可用的API版本。
func (c *ElbClient) ListApiVersions(request *model.ListApiVersionsRequest) (*model.ListApiVersionsResponse, error) {
	requestDef := GenReqDefForListApiVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionsResponse), nil
	}
}

//批量删除IP地址组的IP列表项。
func (c *ElbClient) BatchDeleteIpList(request *model.BatchDeleteIpListRequest) (*model.BatchDeleteIpListResponse, error) {
	requestDef := GenReqDefForBatchDeleteIpList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteIpListResponse), nil
	}
}

//计算以下几种场景的预占用IP数量： - 计算创建LB的第一个七层监听器后总占用IP数量：传入loadbalancer_id、l7_flavor_id为空、ip_target_enable不传或为false。 - 计算LB规格变更或开启跨VPC后总占用IP数量：传入参数loadbalancer_id，及l7_flavor_id不为空或ip_target_enable为true。 - 计算创建LB所需IP数量：传入参数availability_zone_id，及可选参数l7_flavor_id、ip_target_enable、ip_version，不能传loadbalancer_id。 > 查询出的预占IP数大于等于最终实际占用的IP数。
func (c *ElbClient) CountPreoccupyIpNum(request *model.CountPreoccupyIpNumRequest) (*model.CountPreoccupyIpNumResponse, error) {
	requestDef := GenReqDefForCountPreoccupyIpNum()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountPreoccupyIpNumResponse), nil
	}
}

//创建IP地址组。输入的ip可为ip地址或者CIDR子网，支持IPV4和IPV6。需要注意，0.0.0.0与0.0.0.0/32视为重复，0&#58;0&#58;0&#58;0&#58;0&#58;0&#58;0&#58;1与&#58;&#58;1与&#58;&#58;1/128视为重复，会只保留其中一个写入。
func (c *ElbClient) CreateIpGroup(request *model.CreateIpGroupRequest) (*model.CreateIpGroupResponse, error) {
	requestDef := GenReqDefForCreateIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIpGroupResponse), nil
	}
}

//删除ip地址组。
func (c *ElbClient) DeleteIpGroup(request *model.DeleteIpGroupRequest) (*model.DeleteIpGroupResponse, error) {
	requestDef := GenReqDefForDeleteIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteIpGroupResponse), nil
	}
}

//查询IP地址组列表。
func (c *ElbClient) ListIpGroups(request *model.ListIpGroupsRequest) (*model.ListIpGroupsResponse, error) {
	requestDef := GenReqDefForListIpGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIpGroupsResponse), nil
	}
}

//获取IP地址组详情。
func (c *ElbClient) ShowIpGroup(request *model.ShowIpGroupRequest) (*model.ShowIpGroupResponse, error) {
	requestDef := GenReqDefForShowIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIpGroupResponse), nil
	}
}

//更新IP地址组，只支持全量更新IP。即IP地址组中的ip_list将被全量覆盖，不在请求参数中的IP地址将被移除。输入的ip可为ip地址或者CIDR子网，支持IPV4和IPV6。需要注意，0.0.0.0与0.0.0.0/32视为重复，0&#58;0&#58;0&#58;0&#58;0&#58;0&#58;0&#58;1与&#58;&#58;1与&#58;&#58;1/128视为重复，会只保留其中一个写入。
func (c *ElbClient) UpdateIpGroup(request *model.UpdateIpGroupRequest) (*model.UpdateIpGroupResponse, error) {
	requestDef := GenReqDefForUpdateIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIpGroupResponse), nil
	}
}

//更新IP地址组的IP列表信息。
func (c *ElbClient) UpdateIpList(request *model.UpdateIpListRequest) (*model.UpdateIpListResponse, error) {
	requestDef := GenReqDefForUpdateIpList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIpListResponse), nil
	}
}
