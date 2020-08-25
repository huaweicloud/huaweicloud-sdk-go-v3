package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v3/model"
)

type ElbClient struct {
	hcClient *http_client.HcHttpClient
}

func NewElbClient(hcClient *http_client.HcHttpClient) *ElbClient {
	return &ElbClient{hcClient: hcClient}
}

func ElbClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//创建证书。
func (c *ElbClient) CreateCertificate(request *model.CreateCertificateRequest) (*model.CreateCertificateResponse, error) {
	requestDef := GenReqDefForCreateCertificate(request)
	resp, responseDef := GenRespForCreateCertificate()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//创建健康检查。
func (c *ElbClient) CreateHealthMonitor(request *model.CreateHealthMonitorRequest) (*model.CreateHealthMonitorResponse, error) {
	requestDef := GenReqDefForCreateHealthMonitor(request)
	resp, responseDef := GenRespForCreateHealthMonitor()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//创建转发策略.
func (c *ElbClient) CreateL7Policy(request *model.CreateL7PolicyRequest) (*model.CreateL7PolicyResponse, error) {
	requestDef := GenReqDefForCreateL7Policy(request)
	resp, responseDef := GenRespForCreateL7Policy()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//创建转发规则。
func (c *ElbClient) CreateL7Rule(request *model.CreateL7RuleRequest) (*model.CreateL7RuleResponse, error) {
	requestDef := GenReqDefForCreateL7Rule(request)
	resp, responseDef := GenRespForCreateL7Rule()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//ElbV3 创建监听器。
func (c *ElbClient) CreateListener(request *model.CreateListenerRequest) (*model.CreateListenerResponse, error) {
	requestDef := GenReqDefForCreateListener(request)
	resp, responseDef := GenRespForCreateListener()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//创建负载均衡器。 1.创建公网负载均衡器的场合，需要传入vpc_id。 2.创建内网负载均衡器的场合，需要传入vip_subnet_cidr_id。 3.创建内网双栈负载均衡器的场合，需要传入ipv6_vip_virsubnet_id。  关联有已有公网ip地址，需要传入publicip_ids 新建公网ip地址，需要传入publicip 包括IPV4私网类型，IPV4公网类型，IPV6私网，IPV6公网
func (c *ElbClient) CreateLoadBalancer(request *model.CreateLoadBalancerRequest) (*model.CreateLoadBalancerResponse, error) {
	requestDef := GenReqDefForCreateLoadBalancer(request)
	resp, responseDef := GenRespForCreateLoadBalancer()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//创建后端服务器。
func (c *ElbClient) CreateMember(request *model.CreateMemberRequest) (*model.CreateMemberResponse, error) {
	requestDef := GenReqDefForCreateMember(request)
	resp, responseDef := GenRespForCreateMember()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//创建后端服务器组。
func (c *ElbClient) CreatePool(request *model.CreatePoolRequest) (*model.CreatePoolResponse, error) {
	requestDef := GenReqDefForCreatePool(request)
	resp, responseDef := GenRespForCreatePool()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除SSL证书。
func (c *ElbClient) DeleteCertificate(request *model.DeleteCertificateRequest) (*model.DeleteCertificateResponse, error) {
	requestDef := GenReqDefForDeleteCertificate(request)
	resp, responseDef := GenRespForDeleteCertificate()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除健康检查。
func (c *ElbClient) DeleteHealthMonitor(request *model.DeleteHealthMonitorRequest) (*model.DeleteHealthMonitorResponse, error) {
	requestDef := GenReqDefForDeleteHealthMonitor(request)
	resp, responseDef := GenRespForDeleteHealthMonitor()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除转发策略。
func (c *ElbClient) DeleteL7Policy(request *model.DeleteL7PolicyRequest) (*model.DeleteL7PolicyResponse, error) {
	requestDef := GenReqDefForDeleteL7Policy(request)
	resp, responseDef := GenRespForDeleteL7Policy()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除转发规则。
func (c *ElbClient) DeleteL7Rule(request *model.DeleteL7RuleRequest) (*model.DeleteL7RuleResponse, error) {
	requestDef := GenReqDefForDeleteL7Rule(request)
	resp, responseDef := GenRespForDeleteL7Rule()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除监听器。
func (c *ElbClient) DeleteListener(request *model.DeleteListenerRequest) (*model.DeleteListenerResponse, error) {
	requestDef := GenReqDefForDeleteListener(request)
	resp, responseDef := GenRespForDeleteListener()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除负载均衡器。
func (c *ElbClient) DeleteLoadBalancer(request *model.DeleteLoadBalancerRequest) (*model.DeleteLoadBalancerResponse, error) {
	requestDef := GenReqDefForDeleteLoadBalancer(request)
	resp, responseDef := GenRespForDeleteLoadBalancer()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除后端服务器。
func (c *ElbClient) DeleteMember(request *model.DeleteMemberRequest) (*model.DeleteMemberResponse, error) {
	requestDef := GenReqDefForDeleteMember(request)
	resp, responseDef := GenRespForDeleteMember()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除后端服务器组。
func (c *ElbClient) DeletePool(request *model.DeletePoolRequest) (*model.DeletePoolResponse, error) {
	requestDef := GenReqDefForDeletePool(request)
	resp, responseDef := GenRespForDeletePool()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//返回用户可使用的可用区的列表情况
func (c *ElbClient) ListAvailabilityZones(request *model.ListAvailabilityZonesRequest) (*model.ListAvailabilityZonesResponse, error) {
	requestDef := GenReqDefForListAvailabilityZones(request)
	resp, responseDef := GenRespForListAvailabilityZones()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询SSL证书列表。
func (c *ElbClient) ListCertificates(request *model.ListCertificatesRequest) (*model.ListCertificatesResponse, error) {
	requestDef := GenReqDefForListCertificates(request)
	resp, responseDef := GenRespForListCertificates()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询所有的规格。
func (c *ElbClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors(request)
	resp, responseDef := GenRespForListFlavors()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//健康检查列表。
func (c *ElbClient) ListHealthMonitors(request *model.ListHealthMonitorsRequest) (*model.ListHealthMonitorsResponse, error) {
	requestDef := GenReqDefForListHealthMonitors(request)
	resp, responseDef := GenRespForListHealthMonitors()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询转发策略列表。
func (c *ElbClient) ListL7Policies(request *model.ListL7PoliciesRequest) (*model.ListL7PoliciesResponse, error) {
	requestDef := GenReqDefForListL7Policies(request)
	resp, responseDef := GenRespForListL7Policies()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询转发规则列表。
func (c *ElbClient) ListL7Rules(request *model.ListL7RulesRequest) (*model.ListL7RulesResponse, error) {
	requestDef := GenReqDefForListL7Rules(request)
	resp, responseDef := GenRespForListL7Rules()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询监听器列表。
func (c *ElbClient) ListListeners(request *model.ListListenersRequest) (*model.ListListenersResponse, error) {
	requestDef := GenReqDefForListListeners(request)
	resp, responseDef := GenRespForListListeners()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询负载均衡器列表，支持过滤查询和分页查询
func (c *ElbClient) ListLoadBalancers(request *model.ListLoadBalancersRequest) (*model.ListLoadBalancersResponse, error) {
	requestDef := GenReqDefForListLoadBalancers(request)
	resp, responseDef := GenRespForListLoadBalancers()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//Pool下的后端服务器列表。
func (c *ElbClient) ListMembers(request *model.ListMembersRequest) (*model.ListMembersResponse, error) {
	requestDef := GenReqDefForListMembers(request)
	resp, responseDef := GenRespForListMembers()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//后端服务器组列表。
func (c *ElbClient) ListPools(request *model.ListPoolsRequest) (*model.ListPoolsResponse, error) {
	requestDef := GenReqDefForListPools(request)
	resp, responseDef := GenRespForListPools()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询SSL证书详情。
func (c *ElbClient) ShowCertificate(request *model.ShowCertificateRequest) (*model.ShowCertificateResponse, error) {
	requestDef := GenReqDefForShowCertificate(request)
	resp, responseDef := GenRespForShowCertificate()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询规格的详情。
func (c *ElbClient) ShowFlavor(request *model.ShowFlavorRequest) (*model.ShowFlavorResponse, error) {
	requestDef := GenReqDefForShowFlavor(request)
	resp, responseDef := GenRespForShowFlavor()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询健康检查详情。
func (c *ElbClient) ShowHealthMonitor(request *model.ShowHealthMonitorRequest) (*model.ShowHealthMonitorResponse, error) {
	requestDef := GenReqDefForShowHealthMonitor(request)
	resp, responseDef := GenRespForShowHealthMonitor()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询转发策略详情。
func (c *ElbClient) ShowL7Policy(request *model.ShowL7PolicyRequest) (*model.ShowL7PolicyResponse, error) {
	requestDef := GenReqDefForShowL7Policy(request)
	resp, responseDef := GenRespForShowL7Policy()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询转发规则详情
func (c *ElbClient) ShowL7Rule(request *model.ShowL7RuleRequest) (*model.ShowL7RuleResponse, error) {
	requestDef := GenReqDefForShowL7Rule(request)
	resp, responseDef := GenRespForShowL7Rule()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//监听器详情。
func (c *ElbClient) ShowListener(request *model.ShowListenerRequest) (*model.ShowListenerResponse, error) {
	requestDef := GenReqDefForShowListener(request)
	resp, responseDef := GenRespForShowListener()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询负载均衡器详情
func (c *ElbClient) ShowLoadBalancer(request *model.ShowLoadBalancerRequest) (*model.ShowLoadBalancerResponse, error) {
	requestDef := GenReqDefForShowLoadBalancer(request)
	resp, responseDef := GenRespForShowLoadBalancer()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询负载均衡器状态树，列出负载均衡器关联的子资源的信息
func (c *ElbClient) ShowLoadBalancerStatus(request *model.ShowLoadBalancerStatusRequest) (*model.ShowLoadBalancerStatusResponse, error) {
	requestDef := GenReqDefForShowLoadBalancerStatus(request)
	resp, responseDef := GenRespForShowLoadBalancerStatus()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//后端服务器详情
func (c *ElbClient) ShowMember(request *model.ShowMemberRequest) (*model.ShowMemberResponse, error) {
	requestDef := GenReqDefForShowMember(request)
	resp, responseDef := GenRespForShowMember()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//后端服务器组详情。
func (c *ElbClient) ShowPool(request *model.ShowPoolRequest) (*model.ShowPoolResponse, error) {
	requestDef := GenReqDefForShowPool(request)
	resp, responseDef := GenRespForShowPool()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//【不开放】查询特定项目的配额数。
func (c *ElbClient) ShowQuota(request *model.ShowQuotaRequest) (*model.ShowQuotaResponse, error) {
	requestDef := GenReqDefForShowQuota(request)
	resp, responseDef := GenRespForShowQuota()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//【不开放】查询默认配额数。
func (c *ElbClient) ShowQuotaDefaults(request *model.ShowQuotaDefaultsRequest) (*model.ShowQuotaDefaultsResponse, error) {
	requestDef := GenReqDefForShowQuotaDefaults(request)
	resp, responseDef := GenRespForShowQuotaDefaults()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//更新SSL证书。
func (c *ElbClient) UpdateCertificate(request *model.UpdateCertificateRequest) (*model.UpdateCertificateResponse, error) {
	requestDef := GenReqDefForUpdateCertificate(request)
	resp, responseDef := GenRespForUpdateCertificate()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//更新健康检查。
func (c *ElbClient) UpdateHealthMonitor(request *model.UpdateHealthMonitorRequest) (*model.UpdateHealthMonitorResponse, error) {
	requestDef := GenReqDefForUpdateHealthMonitor(request)
	resp, responseDef := GenRespForUpdateHealthMonitor()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//更新转发策略。
func (c *ElbClient) UpdateL7Policy(request *model.UpdateL7PolicyRequest) (*model.UpdateL7PolicyResponse, error) {
	requestDef := GenReqDefForUpdateL7Policy(request)
	resp, responseDef := GenRespForUpdateL7Policy()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//更新转发规则。
func (c *ElbClient) UpdateL7Rule(request *model.UpdateL7RuleRequest) (*model.UpdateL7RuleResponse, error) {
	requestDef := GenReqDefForUpdateL7Rule(request)
	resp, responseDef := GenRespForUpdateL7Rule()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//更新监听器。
func (c *ElbClient) UpdateListener(request *model.UpdateListenerRequest) (*model.UpdateListenerResponse, error) {
	requestDef := GenReqDefForUpdateListener(request)
	resp, responseDef := GenRespForUpdateListener()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//更新负载均衡器。
func (c *ElbClient) UpdateLoadBalancer(request *model.UpdateLoadBalancerRequest) (*model.UpdateLoadBalancerResponse, error) {
	requestDef := GenReqDefForUpdateLoadBalancer(request)
	resp, responseDef := GenRespForUpdateLoadBalancer()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//如果member绑定的负载均衡器的provisioning status不是ACTIVE，则不能更新该member。
func (c *ElbClient) UpdateMember(request *model.UpdateMemberRequest) (*model.UpdateMemberResponse, error) {
	requestDef := GenReqDefForUpdateMember(request)
	resp, responseDef := GenRespForUpdateMember()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//更新后端服务器组。
func (c *ElbClient) UpdatePool(request *model.UpdatePoolRequest) (*model.UpdatePoolResponse, error) {
	requestDef := GenReqDefForUpdatePool(request)
	resp, responseDef := GenRespForUpdatePool()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//计算创建一个负载均衡实例和第一个七层监听器需预占用的IP数
func (c *ElbClient) CountPreoccupyIpNum(request *model.CountPreoccupyIpNumRequest) (*model.CountPreoccupyIpNumResponse, error) {
	requestDef := GenReqDefForCountPreoccupyIpNum(request)
	resp, responseDef := GenRespForCountPreoccupyIpNum()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//创建ip地址组
func (c *ElbClient) CreateIpGroup(request *model.CreateIpGroupRequest) (*model.CreateIpGroupResponse, error) {
	requestDef := GenReqDefForCreateIpGroup(request)
	resp, responseDef := GenRespForCreateIpGroup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除ip地址组。
func (c *ElbClient) DeleteIpGroup(request *model.DeleteIpGroupRequest) (*model.DeleteIpGroupResponse, error) {
	requestDef := GenReqDefForDeleteIpGroup(request)
	resp, responseDef := GenRespForDeleteIpGroup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询IP地址组列表
func (c *ElbClient) ListIpGroups(request *model.ListIpGroupsRequest) (*model.ListIpGroupsResponse, error) {
	requestDef := GenReqDefForListIpGroups(request)
	resp, responseDef := GenRespForListIpGroups()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//获取ip地址组详情
func (c *ElbClient) ShowIpGroup(request *model.ShowIpGroupRequest) (*model.ShowIpGroupResponse, error) {
	requestDef := GenReqDefForShowIpGroup(request)
	resp, responseDef := GenRespForShowIpGroup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//更新ip地址组，只支持全量更新ip。
func (c *ElbClient) UpdateIpGroup(request *model.UpdateIpGroupRequest) (*model.UpdateIpGroupResponse, error) {
	requestDef := GenReqDefForUpdateIpGroup(request)
	resp, responseDef := GenRespForUpdateIpGroup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}
