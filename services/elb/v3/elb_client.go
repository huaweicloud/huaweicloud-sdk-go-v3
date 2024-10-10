package v3

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v3/model"
)

type ElbClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewElbClient(hcClient *httpclient.HcHttpClient) *ElbClient {
	return &ElbClient{HcClient: hcClient}
}

func ElbClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// BatchAddAvailableZones 新增负载均衡器可用区
//
// 给负载均衡器新增可用区。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) BatchAddAvailableZones(request *model.BatchAddAvailableZonesRequest) (*model.BatchAddAvailableZonesResponse, error) {
	requestDef := GenReqDefForBatchAddAvailableZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddAvailableZonesResponse), nil
	}
}

// BatchAddAvailableZonesInvoker 新增负载均衡器可用区
func (c *ElbClient) BatchAddAvailableZonesInvoker(request *model.BatchAddAvailableZonesRequest) *BatchAddAvailableZonesInvoker {
	requestDef := GenReqDefForBatchAddAvailableZones()
	return &BatchAddAvailableZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateMembers 批量创建后端服务器
//
// 在指定pool下批量创建后端服务器。一次最多创建200个。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) BatchCreateMembers(request *model.BatchCreateMembersRequest) (*model.BatchCreateMembersResponse, error) {
	requestDef := GenReqDefForBatchCreateMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateMembersResponse), nil
	}
}

// BatchCreateMembersInvoker 批量创建后端服务器
func (c *ElbClient) BatchCreateMembersInvoker(request *model.BatchCreateMembersRequest) *BatchCreateMembersInvoker {
	requestDef := GenReqDefForBatchCreateMembers()
	return &BatchCreateMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteMembers 批量删除后端服务器
//
// 在指定pool下批量删除后端服务器。一次最多添加200个。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) BatchDeleteMembers(request *model.BatchDeleteMembersRequest) (*model.BatchDeleteMembersResponse, error) {
	requestDef := GenReqDefForBatchDeleteMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteMembersResponse), nil
	}
}

// BatchDeleteMembersInvoker 批量删除后端服务器
func (c *ElbClient) BatchDeleteMembersInvoker(request *model.BatchDeleteMembersRequest) *BatchDeleteMembersInvoker {
	requestDef := GenReqDefForBatchDeleteMembers()
	return &BatchDeleteMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchRemoveAvailableZones 移除负载均衡器可用区
//
// 移除负载均衡器的可用区。
// &gt; 移除可用区可能导致已有链接断开，请谨慎操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) BatchRemoveAvailableZones(request *model.BatchRemoveAvailableZonesRequest) (*model.BatchRemoveAvailableZonesResponse, error) {
	requestDef := GenReqDefForBatchRemoveAvailableZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRemoveAvailableZonesResponse), nil
	}
}

// BatchRemoveAvailableZonesInvoker 移除负载均衡器可用区
func (c *ElbClient) BatchRemoveAvailableZonesInvoker(request *model.BatchRemoveAvailableZonesRequest) *BatchRemoveAvailableZonesInvoker {
	requestDef := GenReqDefForBatchRemoveAvailableZones()
	return &BatchRemoveAvailableZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateMembers 批量更新后端服务器
//
// 在指定pool下批量更新后端服务器。一次最多添加200个。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) BatchUpdateMembers(request *model.BatchUpdateMembersRequest) (*model.BatchUpdateMembersResponse, error) {
	requestDef := GenReqDefForBatchUpdateMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateMembersResponse), nil
	}
}

// BatchUpdateMembersInvoker 批量更新后端服务器
func (c *ElbClient) BatchUpdateMembersInvoker(request *model.BatchUpdateMembersRequest) *BatchUpdateMembersInvoker {
	requestDef := GenReqDefForBatchUpdateMembers()
	return &BatchUpdateMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdatePoliciesPriority 批量更新转发策略优先级
//
// 批量更新转发策略的优先级。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) BatchUpdatePoliciesPriority(request *model.BatchUpdatePoliciesPriorityRequest) (*model.BatchUpdatePoliciesPriorityResponse, error) {
	requestDef := GenReqDefForBatchUpdatePoliciesPriority()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdatePoliciesPriorityResponse), nil
	}
}

// BatchUpdatePoliciesPriorityInvoker 批量更新转发策略优先级
func (c *ElbClient) BatchUpdatePoliciesPriorityInvoker(request *model.BatchUpdatePoliciesPriorityRequest) *BatchUpdatePoliciesPriorityInvoker {
	requestDef := GenReqDefForBatchUpdatePoliciesPriority()
	return &BatchUpdatePoliciesPriorityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeLoadbalancerChargeMode 变更负载均衡器计费模式
//
// 负载均衡器计费模式变更，当前支持的计费模式变更为：
// 1. 按需计费转包周期计费；
// 2. 按需按规格计费转按需按使用量计费；
// 3. 按需按使用量计费转按需按规格计费；
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ChangeLoadbalancerChargeMode(request *model.ChangeLoadbalancerChargeModeRequest) (*model.ChangeLoadbalancerChargeModeResponse, error) {
	requestDef := GenReqDefForChangeLoadbalancerChargeMode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeLoadbalancerChargeModeResponse), nil
	}
}

// ChangeLoadbalancerChargeModeInvoker 变更负载均衡器计费模式
func (c *ElbClient) ChangeLoadbalancerChargeModeInvoker(request *model.ChangeLoadbalancerChargeModeRequest) *ChangeLoadbalancerChargeModeInvoker {
	requestDef := GenReqDefForChangeLoadbalancerChargeMode()
	return &ChangeLoadbalancerChargeModeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCertificate 创建证书
//
// 创建证书。用于HTTPS协议监听器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) CreateCertificate(request *model.CreateCertificateRequest) (*model.CreateCertificateResponse, error) {
	requestDef := GenReqDefForCreateCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCertificateResponse), nil
	}
}

// CreateCertificateInvoker 创建证书
func (c *ElbClient) CreateCertificateInvoker(request *model.CreateCertificateRequest) *CreateCertificateInvoker {
	requestDef := GenReqDefForCreateCertificate()
	return &CreateCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCertificatePrivateKeyEcho 修改证书私钥字段回显开关
//
// 开启或关闭证书私钥字段回显开关。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) CreateCertificatePrivateKeyEcho(request *model.CreateCertificatePrivateKeyEchoRequest) (*model.CreateCertificatePrivateKeyEchoResponse, error) {
	requestDef := GenReqDefForCreateCertificatePrivateKeyEcho()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCertificatePrivateKeyEchoResponse), nil
	}
}

// CreateCertificatePrivateKeyEchoInvoker 修改证书私钥字段回显开关
func (c *ElbClient) CreateCertificatePrivateKeyEchoInvoker(request *model.CreateCertificatePrivateKeyEchoRequest) *CreateCertificatePrivateKeyEchoInvoker {
	requestDef := GenReqDefForCreateCertificatePrivateKeyEcho()
	return &CreateCertificatePrivateKeyEchoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHealthMonitor 创建健康检查
//
// 创建健康检查。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) CreateHealthMonitor(request *model.CreateHealthMonitorRequest) (*model.CreateHealthMonitorResponse, error) {
	requestDef := GenReqDefForCreateHealthMonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHealthMonitorResponse), nil
	}
}

// CreateHealthMonitorInvoker 创建健康检查
func (c *ElbClient) CreateHealthMonitorInvoker(request *model.CreateHealthMonitorRequest) *CreateHealthMonitorInvoker {
	requestDef := GenReqDefForCreateHealthMonitor()
	return &CreateHealthMonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateL7Policy 创建转发策略
//
// 创建七层转发策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) CreateL7Policy(request *model.CreateL7PolicyRequest) (*model.CreateL7PolicyResponse, error) {
	requestDef := GenReqDefForCreateL7Policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateL7PolicyResponse), nil
	}
}

// CreateL7PolicyInvoker 创建转发策略
func (c *ElbClient) CreateL7PolicyInvoker(request *model.CreateL7PolicyRequest) *CreateL7PolicyInvoker {
	requestDef := GenReqDefForCreateL7Policy()
	return &CreateL7PolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateL7Rule 创建转发规则
//
// 创建七层转发规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) CreateL7Rule(request *model.CreateL7RuleRequest) (*model.CreateL7RuleResponse, error) {
	requestDef := GenReqDefForCreateL7Rule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateL7RuleResponse), nil
	}
}

// CreateL7RuleInvoker 创建转发规则
func (c *ElbClient) CreateL7RuleInvoker(request *model.CreateL7RuleRequest) *CreateL7RuleInvoker {
	requestDef := GenReqDefForCreateL7Rule()
	return &CreateL7RuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateListener 创建监听器
//
// 创建监听器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) CreateListener(request *model.CreateListenerRequest) (*model.CreateListenerResponse, error) {
	requestDef := GenReqDefForCreateListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateListenerResponse), nil
	}
}

// CreateListenerInvoker 创建监听器
func (c *ElbClient) CreateListenerInvoker(request *model.CreateListenerRequest) *CreateListenerInvoker {
	requestDef := GenReqDefForCreateListener()
	return &CreateListenerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLoadBalancer 创建负载均衡器
//
// 创建独享型负载均衡器，包括按需及包周期计费负载均衡器。
// 1. 若要创建内网IPv4负载均衡器，则需要传入vip_subnet_cidr_id。
// 2. 若要创建公网IPv4负载均衡器，则需要传入publicip，以及传入vpc_id和vip_subnet_cidr_id这两个参数中的一个。
// 3. 若要绑定有已有公网IPv4地址，则需要传入publicip_ids，以及传入vpc_id和vip_subnet_cidr_id这两个参数中的一个。
// 4. 若要创建内网双栈负载均衡器，则需要传入ipv6_vip_virsubnet_id。
// 5. 若要创建公网双栈负载均衡器，则需要传入ipv6_vip_virsubnet_id和ipv6_bandwidth。
// 6. 若要创建网络型负载均衡器，则需要传入l4_flavor_id（网络型规格ID）；若要创建应用型负载均衡器，则需要传入l7_flavor_id（应用型规格ID）；若要创建网络型+应用型负载均衡器，则需要传入l4_flavor_id和l7_flavor_id。
// 7. 若要创建包周期负载均衡器，则需要传入prepaid_options，否则创建按需计费负载均衡器。
// 8. 按需计费分为固定规格计费和弹性规格计费，根据创建时所选规格的类型决定计费方式。具体规格说明见创建LB请求参数l4_flavor_id和l7_flavor_id。
// [9.若要创建gateway类型的负载均衡器，则需要：
//   - 指定loadbalancer_type&#x3D;\&quot;gateway\&quot;，且不支持指定vip_address，ipv6_vip_address。
//   - vip_subnet_cidr_id和ipv6_subnet_cidr_id两者不能都为空，如果两者都传入，则必须属于同一子网。
//   - 不支持创建公网gateway类型LB。
//   - 如果要指定规格，则从请求参数gw_flavor_id传入。](tag:hws_eu)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) CreateLoadBalancer(request *model.CreateLoadBalancerRequest) (*model.CreateLoadBalancerResponse, error) {
	requestDef := GenReqDefForCreateLoadBalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLoadBalancerResponse), nil
	}
}

// CreateLoadBalancerInvoker 创建负载均衡器
func (c *ElbClient) CreateLoadBalancerInvoker(request *model.CreateLoadBalancerRequest) *CreateLoadBalancerInvoker {
	requestDef := GenReqDefForCreateLoadBalancer()
	return &CreateLoadBalancerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLogtank 创建云日志
//
// 创建云日志。[荷兰region不支持云日志功能，请勿使用。](tag:dt,dt_test)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) CreateLogtank(request *model.CreateLogtankRequest) (*model.CreateLogtankResponse, error) {
	requestDef := GenReqDefForCreateLogtank()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLogtankResponse), nil
	}
}

// CreateLogtankInvoker 创建云日志
func (c *ElbClient) CreateLogtankInvoker(request *model.CreateLogtankRequest) *CreateLogtankInvoker {
	requestDef := GenReqDefForCreateLogtank()
	return &CreateLogtankInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMasterSlavePool 创建主备后端服务器组
//
// 创建主备后端服务器组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) CreateMasterSlavePool(request *model.CreateMasterSlavePoolRequest) (*model.CreateMasterSlavePoolResponse, error) {
	requestDef := GenReqDefForCreateMasterSlavePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMasterSlavePoolResponse), nil
	}
}

// CreateMasterSlavePoolInvoker 创建主备后端服务器组
func (c *ElbClient) CreateMasterSlavePoolInvoker(request *model.CreateMasterSlavePoolRequest) *CreateMasterSlavePoolInvoker {
	requestDef := GenReqDefForCreateMasterSlavePool()
	return &CreateMasterSlavePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMember 创建后端服务器
//
// 创建后端服务器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) CreateMember(request *model.CreateMemberRequest) (*model.CreateMemberResponse, error) {
	requestDef := GenReqDefForCreateMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMemberResponse), nil
	}
}

// CreateMemberInvoker 创建后端服务器
func (c *ElbClient) CreateMemberInvoker(request *model.CreateMemberRequest) *CreateMemberInvoker {
	requestDef := GenReqDefForCreateMember()
	return &CreateMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePool 创建后端服务器组
//
// 创建后端服务器组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) CreatePool(request *model.CreatePoolRequest) (*model.CreatePoolResponse, error) {
	requestDef := GenReqDefForCreatePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePoolResponse), nil
	}
}

// CreatePoolInvoker 创建后端服务器组
func (c *ElbClient) CreatePoolInvoker(request *model.CreatePoolRequest) *CreatePoolInvoker {
	requestDef := GenReqDefForCreatePool()
	return &CreatePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSecurityPolicy 创建自定义安全策略
//
// 创建自定义安全策略。用于在创建HTTPS监听器时，请求参数中指定security_policy_id来设置监听器的自定义安全策略。
//
// [荷兰region不支持自定义安全策略功能，请勿使用。](tag:dt,dt_test)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) CreateSecurityPolicy(request *model.CreateSecurityPolicyRequest) (*model.CreateSecurityPolicyResponse, error) {
	requestDef := GenReqDefForCreateSecurityPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecurityPolicyResponse), nil
	}
}

// CreateSecurityPolicyInvoker 创建自定义安全策略
func (c *ElbClient) CreateSecurityPolicyInvoker(request *model.CreateSecurityPolicyRequest) *CreateSecurityPolicyInvoker {
	requestDef := GenReqDefForCreateSecurityPolicy()
	return &CreateSecurityPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCertificate 删除证书
//
// 删除证书。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) DeleteCertificate(request *model.DeleteCertificateRequest) (*model.DeleteCertificateResponse, error) {
	requestDef := GenReqDefForDeleteCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCertificateResponse), nil
	}
}

// DeleteCertificateInvoker 删除证书
func (c *ElbClient) DeleteCertificateInvoker(request *model.DeleteCertificateRequest) *DeleteCertificateInvoker {
	requestDef := GenReqDefForDeleteCertificate()
	return &DeleteCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHealthMonitor 删除健康检查
//
// 删除健康检查。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) DeleteHealthMonitor(request *model.DeleteHealthMonitorRequest) (*model.DeleteHealthMonitorResponse, error) {
	requestDef := GenReqDefForDeleteHealthMonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHealthMonitorResponse), nil
	}
}

// DeleteHealthMonitorInvoker 删除健康检查
func (c *ElbClient) DeleteHealthMonitorInvoker(request *model.DeleteHealthMonitorRequest) *DeleteHealthMonitorInvoker {
	requestDef := GenReqDefForDeleteHealthMonitor()
	return &DeleteHealthMonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteL7Policy 删除转发策略
//
// 删除七层转发策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) DeleteL7Policy(request *model.DeleteL7PolicyRequest) (*model.DeleteL7PolicyResponse, error) {
	requestDef := GenReqDefForDeleteL7Policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteL7PolicyResponse), nil
	}
}

// DeleteL7PolicyInvoker 删除转发策略
func (c *ElbClient) DeleteL7PolicyInvoker(request *model.DeleteL7PolicyRequest) *DeleteL7PolicyInvoker {
	requestDef := GenReqDefForDeleteL7Policy()
	return &DeleteL7PolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteL7Rule 删除转发规则
//
// 删除七层转发规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) DeleteL7Rule(request *model.DeleteL7RuleRequest) (*model.DeleteL7RuleResponse, error) {
	requestDef := GenReqDefForDeleteL7Rule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteL7RuleResponse), nil
	}
}

// DeleteL7RuleInvoker 删除转发规则
func (c *ElbClient) DeleteL7RuleInvoker(request *model.DeleteL7RuleRequest) *DeleteL7RuleInvoker {
	requestDef := GenReqDefForDeleteL7Rule()
	return &DeleteL7RuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteListener 删除监听器
//
// 删除监听器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) DeleteListener(request *model.DeleteListenerRequest) (*model.DeleteListenerResponse, error) {
	requestDef := GenReqDefForDeleteListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteListenerResponse), nil
	}
}

// DeleteListenerInvoker 删除监听器
func (c *ElbClient) DeleteListenerInvoker(request *model.DeleteListenerRequest) *DeleteListenerInvoker {
	requestDef := GenReqDefForDeleteListener()
	return &DeleteListenerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteListenerForce 级联删除监听器
//
// 删除监听器且级联删除其下子资源（删除监听器、转发策略等，解绑后端服务器组）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) DeleteListenerForce(request *model.DeleteListenerForceRequest) (*model.DeleteListenerForceResponse, error) {
	requestDef := GenReqDefForDeleteListenerForce()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteListenerForceResponse), nil
	}
}

// DeleteListenerForceInvoker 级联删除监听器
func (c *ElbClient) DeleteListenerForceInvoker(request *model.DeleteListenerForceRequest) *DeleteListenerForceInvoker {
	requestDef := GenReqDefForDeleteListenerForce()
	return &DeleteListenerForceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLoadBalancer 删除负载均衡器
//
// 删除负载均衡器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) DeleteLoadBalancer(request *model.DeleteLoadBalancerRequest) (*model.DeleteLoadBalancerResponse, error) {
	requestDef := GenReqDefForDeleteLoadBalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLoadBalancerResponse), nil
	}
}

// DeleteLoadBalancerInvoker 删除负载均衡器
func (c *ElbClient) DeleteLoadBalancerInvoker(request *model.DeleteLoadBalancerRequest) *DeleteLoadBalancerInvoker {
	requestDef := GenReqDefForDeleteLoadBalancer()
	return &DeleteLoadBalancerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLoadBalancerForce 级联删除负载均衡器
//
// 删除负载均衡器且级联删除其下子资源（删除负载均衡器及其绑定的监听器、后端服务器组、后端服务器等一系列资源）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) DeleteLoadBalancerForce(request *model.DeleteLoadBalancerForceRequest) (*model.DeleteLoadBalancerForceResponse, error) {
	requestDef := GenReqDefForDeleteLoadBalancerForce()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLoadBalancerForceResponse), nil
	}
}

// DeleteLoadBalancerForceInvoker 级联删除负载均衡器
func (c *ElbClient) DeleteLoadBalancerForceInvoker(request *model.DeleteLoadBalancerForceRequest) *DeleteLoadBalancerForceInvoker {
	requestDef := GenReqDefForDeleteLoadBalancerForce()
	return &DeleteLoadBalancerForceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLogtank 删除云日志
//
// 删除云日志。[荷兰region不支持云日志功能，请勿使用。](tag:dt,dt_test)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) DeleteLogtank(request *model.DeleteLogtankRequest) (*model.DeleteLogtankResponse, error) {
	requestDef := GenReqDefForDeleteLogtank()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLogtankResponse), nil
	}
}

// DeleteLogtankInvoker 删除云日志
func (c *ElbClient) DeleteLogtankInvoker(request *model.DeleteLogtankRequest) *DeleteLogtankInvoker {
	requestDef := GenReqDefForDeleteLogtank()
	return &DeleteLogtankInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMasterSlavePool 删除主备后端服务器组
//
// 删除主备后端服务器组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) DeleteMasterSlavePool(request *model.DeleteMasterSlavePoolRequest) (*model.DeleteMasterSlavePoolResponse, error) {
	requestDef := GenReqDefForDeleteMasterSlavePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMasterSlavePoolResponse), nil
	}
}

// DeleteMasterSlavePoolInvoker 删除主备后端服务器组
func (c *ElbClient) DeleteMasterSlavePoolInvoker(request *model.DeleteMasterSlavePoolRequest) *DeleteMasterSlavePoolInvoker {
	requestDef := GenReqDefForDeleteMasterSlavePool()
	return &DeleteMasterSlavePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMember 删除后端服务器
//
// 删除后端服务器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) DeleteMember(request *model.DeleteMemberRequest) (*model.DeleteMemberResponse, error) {
	requestDef := GenReqDefForDeleteMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMemberResponse), nil
	}
}

// DeleteMemberInvoker 删除后端服务器
func (c *ElbClient) DeleteMemberInvoker(request *model.DeleteMemberRequest) *DeleteMemberInvoker {
	requestDef := GenReqDefForDeleteMember()
	return &DeleteMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePool 删除后端服务器组
//
// 删除后端服务器组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) DeletePool(request *model.DeletePoolRequest) (*model.DeletePoolResponse, error) {
	requestDef := GenReqDefForDeletePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePoolResponse), nil
	}
}

// DeletePoolInvoker 删除后端服务器组
func (c *ElbClient) DeletePoolInvoker(request *model.DeletePoolRequest) *DeletePoolInvoker {
	requestDef := GenReqDefForDeletePool()
	return &DeletePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSecurityPolicy 删除自定义安全策略
//
// 删除自定义安全策略。[荷兰region不支持自定义安全策略功能，请勿使用。](tag:dt,dt_test)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) DeleteSecurityPolicy(request *model.DeleteSecurityPolicyRequest) (*model.DeleteSecurityPolicyResponse, error) {
	requestDef := GenReqDefForDeleteSecurityPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecurityPolicyResponse), nil
	}
}

// DeleteSecurityPolicyInvoker 删除自定义安全策略
func (c *ElbClient) DeleteSecurityPolicyInvoker(request *model.DeleteSecurityPolicyRequest) *DeleteSecurityPolicyInvoker {
	requestDef := GenReqDefForDeleteSecurityPolicy()
	return &DeleteSecurityPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllMembers 后端服务器全局列表
//
// 查询当前项目下的后端服务器列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ListAllMembers(request *model.ListAllMembersRequest) (*model.ListAllMembersResponse, error) {
	requestDef := GenReqDefForListAllMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllMembersResponse), nil
	}
}

// ListAllMembersInvoker 后端服务器全局列表
func (c *ElbClient) ListAllMembersInvoker(request *model.ListAllMembersRequest) *ListAllMembersInvoker {
	requestDef := GenReqDefForListAllMembers()
	return &ListAllMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAvailabilityZones 查询可用区列表
//
// 返回租户创建LB时可使用的可用区集合列表情况。
//
// - 默认情况下，会返回一个可用区集合。
// 在（如创建LB）设置可用区时，填写的可用区必须包含在可用区集合中、为这个可用区集合的子集。
//
// - 特殊场景下，部分客户要求负载均衡只能创建在指定可用区集合中，此时会返回客户定制的可用区集合。
// 返回可用区集合可能为一个也可能为多个，比如列表有两个可用区集合\\[az1,az2\\],\\[az2,az3\\]。
// 在创建负载均衡器时，可以选择创建在多个可用区，但所选的多个可用区必须同属于其中一个可用区集合，
// 如可以选az2和az3，但不能选择az1和az3。你可以选择多个可用区，只要这些可用区在一个子集中
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ListAvailabilityZones(request *model.ListAvailabilityZonesRequest) (*model.ListAvailabilityZonesResponse, error) {
	requestDef := GenReqDefForListAvailabilityZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailabilityZonesResponse), nil
	}
}

// ListAvailabilityZonesInvoker 查询可用区列表
func (c *ElbClient) ListAvailabilityZonesInvoker(request *model.ListAvailabilityZonesRequest) *ListAvailabilityZonesInvoker {
	requestDef := GenReqDefForListAvailabilityZones()
	return &ListAvailabilityZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCertificates 查询证书列表
//
// 查询证书列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ListCertificates(request *model.ListCertificatesRequest) (*model.ListCertificatesResponse, error) {
	requestDef := GenReqDefForListCertificates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCertificatesResponse), nil
	}
}

// ListCertificatesInvoker 查询证书列表
func (c *ElbClient) ListCertificatesInvoker(request *model.ListCertificatesRequest) *ListCertificatesInvoker {
	requestDef := GenReqDefForListCertificates()
	return &ListCertificatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlavors 查询规格列表
//
// 查询当前region下可用的负载均衡规格列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsResponse), nil
	}
}

// ListFlavorsInvoker 查询规格列表
func (c *ElbClient) ListFlavorsInvoker(request *model.ListFlavorsRequest) *ListFlavorsInvoker {
	requestDef := GenReqDefForListFlavors()
	return &ListFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHealthMonitors 查询健康检查列表
//
// 健康检查列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ListHealthMonitors(request *model.ListHealthMonitorsRequest) (*model.ListHealthMonitorsResponse, error) {
	requestDef := GenReqDefForListHealthMonitors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHealthMonitorsResponse), nil
	}
}

// ListHealthMonitorsInvoker 查询健康检查列表
func (c *ElbClient) ListHealthMonitorsInvoker(request *model.ListHealthMonitorsRequest) *ListHealthMonitorsInvoker {
	requestDef := GenReqDefForListHealthMonitors()
	return &ListHealthMonitorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListL7Policies 查询转发策略列表
//
// 查询七层转发策略列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ListL7Policies(request *model.ListL7PoliciesRequest) (*model.ListL7PoliciesResponse, error) {
	requestDef := GenReqDefForListL7Policies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListL7PoliciesResponse), nil
	}
}

// ListL7PoliciesInvoker 查询转发策略列表
func (c *ElbClient) ListL7PoliciesInvoker(request *model.ListL7PoliciesRequest) *ListL7PoliciesInvoker {
	requestDef := GenReqDefForListL7Policies()
	return &ListL7PoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListL7Rules 查询转发规则列表
//
// 查询转发规则列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ListL7Rules(request *model.ListL7RulesRequest) (*model.ListL7RulesResponse, error) {
	requestDef := GenReqDefForListL7Rules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListL7RulesResponse), nil
	}
}

// ListL7RulesInvoker 查询转发规则列表
func (c *ElbClient) ListL7RulesInvoker(request *model.ListL7RulesRequest) *ListL7RulesInvoker {
	requestDef := GenReqDefForListL7Rules()
	return &ListL7RulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListListeners 查询监听器列表
//
// 查询监听器列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ListListeners(request *model.ListListenersRequest) (*model.ListListenersResponse, error) {
	requestDef := GenReqDefForListListeners()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListListenersResponse), nil
	}
}

// ListListenersInvoker 查询监听器列表
func (c *ElbClient) ListListenersInvoker(request *model.ListListenersRequest) *ListListenersInvoker {
	requestDef := GenReqDefForListListeners()
	return &ListListenersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLoadBalancers 查询负载均衡器列表
//
// 查询负载均衡器列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ListLoadBalancers(request *model.ListLoadBalancersRequest) (*model.ListLoadBalancersResponse, error) {
	requestDef := GenReqDefForListLoadBalancers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLoadBalancersResponse), nil
	}
}

// ListLoadBalancersInvoker 查询负载均衡器列表
func (c *ElbClient) ListLoadBalancersInvoker(request *model.ListLoadBalancersRequest) *ListLoadBalancersInvoker {
	requestDef := GenReqDefForListLoadBalancers()
	return &ListLoadBalancersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLogtanks 查询云日志列表
//
// 查询云日志列表。[荷兰region不支持云日志功能，请勿使用。](tag:dt,dt_test)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ListLogtanks(request *model.ListLogtanksRequest) (*model.ListLogtanksResponse, error) {
	requestDef := GenReqDefForListLogtanks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLogtanksResponse), nil
	}
}

// ListLogtanksInvoker 查询云日志列表
func (c *ElbClient) ListLogtanksInvoker(request *model.ListLogtanksRequest) *ListLogtanksInvoker {
	requestDef := GenReqDefForListLogtanks()
	return &ListLogtanksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMasterSlavePools 查询主备后端服务器组列表
//
// 主备后端服务器组列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ListMasterSlavePools(request *model.ListMasterSlavePoolsRequest) (*model.ListMasterSlavePoolsResponse, error) {
	requestDef := GenReqDefForListMasterSlavePools()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMasterSlavePoolsResponse), nil
	}
}

// ListMasterSlavePoolsInvoker 查询主备后端服务器组列表
func (c *ElbClient) ListMasterSlavePoolsInvoker(request *model.ListMasterSlavePoolsRequest) *ListMasterSlavePoolsInvoker {
	requestDef := GenReqDefForListMasterSlavePools()
	return &ListMasterSlavePoolsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMembers 查询后端服务器列表
//
// Pool下的后端服务器列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ListMembers(request *model.ListMembersRequest) (*model.ListMembersResponse, error) {
	requestDef := GenReqDefForListMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMembersResponse), nil
	}
}

// ListMembersInvoker 查询后端服务器列表
func (c *ElbClient) ListMembersInvoker(request *model.ListMembersRequest) *ListMembersInvoker {
	requestDef := GenReqDefForListMembers()
	return &ListMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPools 查询后端服务器组列表
//
// 后端服务器组列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ListPools(request *model.ListPoolsRequest) (*model.ListPoolsResponse, error) {
	requestDef := GenReqDefForListPools()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPoolsResponse), nil
	}
}

// ListPoolsInvoker 查询后端服务器组列表
func (c *ElbClient) ListPoolsInvoker(request *model.ListPoolsRequest) *ListPoolsInvoker {
	requestDef := GenReqDefForListPools()
	return &ListPoolsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuotaDetails 查询配额使用详情
//
// 查询指定项目中负载均衡相关的各类资源的当前配额和已使用配额信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ListQuotaDetails(request *model.ListQuotaDetailsRequest) (*model.ListQuotaDetailsResponse, error) {
	requestDef := GenReqDefForListQuotaDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotaDetailsResponse), nil
	}
}

// ListQuotaDetailsInvoker 查询配额使用详情
func (c *ElbClient) ListQuotaDetailsInvoker(request *model.ListQuotaDetailsRequest) *ListQuotaDetailsInvoker {
	requestDef := GenReqDefForListQuotaDetails()
	return &ListQuotaDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecurityPolicies 查询自定义安全策略列表
//
// 查询自定义安全策略列表。[荷兰region不支持自定义安全策略功能，请勿使用。](tag:dt,dt_test)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ListSecurityPolicies(request *model.ListSecurityPoliciesRequest) (*model.ListSecurityPoliciesResponse, error) {
	requestDef := GenReqDefForListSecurityPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityPoliciesResponse), nil
	}
}

// ListSecurityPoliciesInvoker 查询自定义安全策略列表
func (c *ElbClient) ListSecurityPoliciesInvoker(request *model.ListSecurityPoliciesRequest) *ListSecurityPoliciesInvoker {
	requestDef := GenReqDefForListSecurityPolicies()
	return &ListSecurityPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSystemSecurityPolicies 查询系统安全策略列表
//
// 查询系统安全策略列表。
//
// 系统安全策略为预置的所有租户通用的安全策略，租户不可新增或修改。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ListSystemSecurityPolicies(request *model.ListSystemSecurityPoliciesRequest) (*model.ListSystemSecurityPoliciesResponse, error) {
	requestDef := GenReqDefForListSystemSecurityPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSystemSecurityPoliciesResponse), nil
	}
}

// ListSystemSecurityPoliciesInvoker 查询系统安全策略列表
func (c *ElbClient) ListSystemSecurityPoliciesInvoker(request *model.ListSystemSecurityPoliciesRequest) *ListSystemSecurityPoliciesInvoker {
	requestDef := GenReqDefForListSystemSecurityPolicies()
	return &ListSystemSecurityPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCertificate 查询证书详情
//
// 查询证书详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ShowCertificate(request *model.ShowCertificateRequest) (*model.ShowCertificateResponse, error) {
	requestDef := GenReqDefForShowCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCertificateResponse), nil
	}
}

// ShowCertificateInvoker 查询证书详情
func (c *ElbClient) ShowCertificateInvoker(request *model.ShowCertificateRequest) *ShowCertificateInvoker {
	requestDef := GenReqDefForShowCertificate()
	return &ShowCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCertificatePrivateKeyEcho 查询证书私钥字段回显开关
//
// 查询证书私钥回显开关当前的状态，开启或关闭。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ShowCertificatePrivateKeyEcho(request *model.ShowCertificatePrivateKeyEchoRequest) (*model.ShowCertificatePrivateKeyEchoResponse, error) {
	requestDef := GenReqDefForShowCertificatePrivateKeyEcho()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCertificatePrivateKeyEchoResponse), nil
	}
}

// ShowCertificatePrivateKeyEchoInvoker 查询证书私钥字段回显开关
func (c *ElbClient) ShowCertificatePrivateKeyEchoInvoker(request *model.ShowCertificatePrivateKeyEchoRequest) *ShowCertificatePrivateKeyEchoInvoker {
	requestDef := GenReqDefForShowCertificatePrivateKeyEcho()
	return &ShowCertificatePrivateKeyEchoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFlavor 查询规格详情
//
// 查询规格的详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ShowFlavor(request *model.ShowFlavorRequest) (*model.ShowFlavorResponse, error) {
	requestDef := GenReqDefForShowFlavor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFlavorResponse), nil
	}
}

// ShowFlavorInvoker 查询规格详情
func (c *ElbClient) ShowFlavorInvoker(request *model.ShowFlavorRequest) *ShowFlavorInvoker {
	requestDef := GenReqDefForShowFlavor()
	return &ShowFlavorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHealthMonitor 查询健康检查详情
//
// 查询健康检查详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ShowHealthMonitor(request *model.ShowHealthMonitorRequest) (*model.ShowHealthMonitorResponse, error) {
	requestDef := GenReqDefForShowHealthMonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHealthMonitorResponse), nil
	}
}

// ShowHealthMonitorInvoker 查询健康检查详情
func (c *ElbClient) ShowHealthMonitorInvoker(request *model.ShowHealthMonitorRequest) *ShowHealthMonitorInvoker {
	requestDef := GenReqDefForShowHealthMonitor()
	return &ShowHealthMonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowL7Policy 查询转发策略详情
//
// 查询七层转发策略详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ShowL7Policy(request *model.ShowL7PolicyRequest) (*model.ShowL7PolicyResponse, error) {
	requestDef := GenReqDefForShowL7Policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowL7PolicyResponse), nil
	}
}

// ShowL7PolicyInvoker 查询转发策略详情
func (c *ElbClient) ShowL7PolicyInvoker(request *model.ShowL7PolicyRequest) *ShowL7PolicyInvoker {
	requestDef := GenReqDefForShowL7Policy()
	return &ShowL7PolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowL7Rule 查询转发规则详情
//
// 查询七层转发规则详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ShowL7Rule(request *model.ShowL7RuleRequest) (*model.ShowL7RuleResponse, error) {
	requestDef := GenReqDefForShowL7Rule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowL7RuleResponse), nil
	}
}

// ShowL7RuleInvoker 查询转发规则详情
func (c *ElbClient) ShowL7RuleInvoker(request *model.ShowL7RuleRequest) *ShowL7RuleInvoker {
	requestDef := GenReqDefForShowL7Rule()
	return &ShowL7RuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowListener 查询监听器详情
//
// 监听器详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ShowListener(request *model.ShowListenerRequest) (*model.ShowListenerResponse, error) {
	requestDef := GenReqDefForShowListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowListenerResponse), nil
	}
}

// ShowListenerInvoker 查询监听器详情
func (c *ElbClient) ShowListenerInvoker(request *model.ShowListenerRequest) *ShowListenerInvoker {
	requestDef := GenReqDefForShowListener()
	return &ShowListenerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLoadBalancer 查询负载均衡器详情
//
// 查询负载均衡器详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ShowLoadBalancer(request *model.ShowLoadBalancerRequest) (*model.ShowLoadBalancerResponse, error) {
	requestDef := GenReqDefForShowLoadBalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLoadBalancerResponse), nil
	}
}

// ShowLoadBalancerInvoker 查询负载均衡器详情
func (c *ElbClient) ShowLoadBalancerInvoker(request *model.ShowLoadBalancerRequest) *ShowLoadBalancerInvoker {
	requestDef := GenReqDefForShowLoadBalancer()
	return &ShowLoadBalancerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLoadBalancerStatus 查询负载均衡器状态树
//
// 查询负载均衡器状态树，包括负载均衡器及其关联的子资源的状态信息。
// 注意：该接口中的operating_status不一定与对应资源的operating_status相同。
// 如：当Member的admin_state_up&#x3D;false且operating_status&#x3D;OFFLINE时，
// 该接口返回member的operating_status&#x3D;DISABLE。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ShowLoadBalancerStatus(request *model.ShowLoadBalancerStatusRequest) (*model.ShowLoadBalancerStatusResponse, error) {
	requestDef := GenReqDefForShowLoadBalancerStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLoadBalancerStatusResponse), nil
	}
}

// ShowLoadBalancerStatusInvoker 查询负载均衡器状态树
func (c *ElbClient) ShowLoadBalancerStatusInvoker(request *model.ShowLoadBalancerStatusRequest) *ShowLoadBalancerStatusInvoker {
	requestDef := GenReqDefForShowLoadBalancerStatus()
	return &ShowLoadBalancerStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLogtank 查询云日志详情
//
// 云日志详情。[荷兰region不支持云日志功能，请勿使用。](tag:dt,dt_test)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ShowLogtank(request *model.ShowLogtankRequest) (*model.ShowLogtankResponse, error) {
	requestDef := GenReqDefForShowLogtank()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLogtankResponse), nil
	}
}

// ShowLogtankInvoker 查询云日志详情
func (c *ElbClient) ShowLogtankInvoker(request *model.ShowLogtankRequest) *ShowLogtankInvoker {
	requestDef := GenReqDefForShowLogtank()
	return &ShowLogtankInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMasterSlavePool 查询主备后端服务器组详情
//
// 主备后端服务器组详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ShowMasterSlavePool(request *model.ShowMasterSlavePoolRequest) (*model.ShowMasterSlavePoolResponse, error) {
	requestDef := GenReqDefForShowMasterSlavePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMasterSlavePoolResponse), nil
	}
}

// ShowMasterSlavePoolInvoker 查询主备后端服务器组详情
func (c *ElbClient) ShowMasterSlavePoolInvoker(request *model.ShowMasterSlavePoolRequest) *ShowMasterSlavePoolInvoker {
	requestDef := GenReqDefForShowMasterSlavePool()
	return &ShowMasterSlavePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMember 查询后端服务器详情
//
// 后端服务器详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ShowMember(request *model.ShowMemberRequest) (*model.ShowMemberResponse, error) {
	requestDef := GenReqDefForShowMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMemberResponse), nil
	}
}

// ShowMemberInvoker 查询后端服务器详情
func (c *ElbClient) ShowMemberInvoker(request *model.ShowMemberRequest) *ShowMemberInvoker {
	requestDef := GenReqDefForShowMember()
	return &ShowMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPool 查询后端服务器组详情
//
// 后端服务器组详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ShowPool(request *model.ShowPoolRequest) (*model.ShowPoolResponse, error) {
	requestDef := GenReqDefForShowPool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPoolResponse), nil
	}
}

// ShowPoolInvoker 查询后端服务器组详情
func (c *ElbClient) ShowPoolInvoker(request *model.ShowPoolRequest) *ShowPoolInvoker {
	requestDef := GenReqDefForShowPool()
	return &ShowPoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQuota 查询配额详情
//
// 查询指定项目中负载均衡相关的各类资源的当前配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ShowQuota(request *model.ShowQuotaRequest) (*model.ShowQuotaResponse, error) {
	requestDef := GenReqDefForShowQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotaResponse), nil
	}
}

// ShowQuotaInvoker 查询配额详情
func (c *ElbClient) ShowQuotaInvoker(request *model.ShowQuotaRequest) *ShowQuotaInvoker {
	requestDef := GenReqDefForShowQuota()
	return &ShowQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSecurityPolicy 查询自定义安全策略详情
//
// 查询自定义安全策略详情。[荷兰region不支持自定义安全策略功能，请勿使用。](tag:dt,dt_test)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ShowSecurityPolicy(request *model.ShowSecurityPolicyRequest) (*model.ShowSecurityPolicyResponse, error) {
	requestDef := GenReqDefForShowSecurityPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecurityPolicyResponse), nil
	}
}

// ShowSecurityPolicyInvoker 查询自定义安全策略详情
func (c *ElbClient) ShowSecurityPolicyInvoker(request *model.ShowSecurityPolicyRequest) *ShowSecurityPolicyInvoker {
	requestDef := GenReqDefForShowSecurityPolicy()
	return &ShowSecurityPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCertificate 更新证书
//
// 更新证书。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) UpdateCertificate(request *model.UpdateCertificateRequest) (*model.UpdateCertificateResponse, error) {
	requestDef := GenReqDefForUpdateCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCertificateResponse), nil
	}
}

// UpdateCertificateInvoker 更新证书
func (c *ElbClient) UpdateCertificateInvoker(request *model.UpdateCertificateRequest) *UpdateCertificateInvoker {
	requestDef := GenReqDefForUpdateCertificate()
	return &UpdateCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHealthMonitor 更新健康检查
//
// 更新健康检查。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) UpdateHealthMonitor(request *model.UpdateHealthMonitorRequest) (*model.UpdateHealthMonitorResponse, error) {
	requestDef := GenReqDefForUpdateHealthMonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHealthMonitorResponse), nil
	}
}

// UpdateHealthMonitorInvoker 更新健康检查
func (c *ElbClient) UpdateHealthMonitorInvoker(request *model.UpdateHealthMonitorRequest) *UpdateHealthMonitorInvoker {
	requestDef := GenReqDefForUpdateHealthMonitor()
	return &UpdateHealthMonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateL7Policy 更新转发策略
//
// 更新七层转发策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) UpdateL7Policy(request *model.UpdateL7PolicyRequest) (*model.UpdateL7PolicyResponse, error) {
	requestDef := GenReqDefForUpdateL7Policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateL7PolicyResponse), nil
	}
}

// UpdateL7PolicyInvoker 更新转发策略
func (c *ElbClient) UpdateL7PolicyInvoker(request *model.UpdateL7PolicyRequest) *UpdateL7PolicyInvoker {
	requestDef := GenReqDefForUpdateL7Policy()
	return &UpdateL7PolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateL7Rule 更新转发规则
//
// 更新七层转发规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) UpdateL7Rule(request *model.UpdateL7RuleRequest) (*model.UpdateL7RuleResponse, error) {
	requestDef := GenReqDefForUpdateL7Rule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateL7RuleResponse), nil
	}
}

// UpdateL7RuleInvoker 更新转发规则
func (c *ElbClient) UpdateL7RuleInvoker(request *model.UpdateL7RuleRequest) *UpdateL7RuleInvoker {
	requestDef := GenReqDefForUpdateL7Rule()
	return &UpdateL7RuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateListener 更新监听器
//
// 更新监听器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) UpdateListener(request *model.UpdateListenerRequest) (*model.UpdateListenerResponse, error) {
	requestDef := GenReqDefForUpdateListener()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateListenerResponse), nil
	}
}

// UpdateListenerInvoker 更新监听器
func (c *ElbClient) UpdateListenerInvoker(request *model.UpdateListenerRequest) *UpdateListenerInvoker {
	requestDef := GenReqDefForUpdateListener()
	return &UpdateListenerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLoadBalancer 更新负载均衡器
//
// 更新负载均衡器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) UpdateLoadBalancer(request *model.UpdateLoadBalancerRequest) (*model.UpdateLoadBalancerResponse, error) {
	requestDef := GenReqDefForUpdateLoadBalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLoadBalancerResponse), nil
	}
}

// UpdateLoadBalancerInvoker 更新负载均衡器
func (c *ElbClient) UpdateLoadBalancerInvoker(request *model.UpdateLoadBalancerRequest) *UpdateLoadBalancerInvoker {
	requestDef := GenReqDefForUpdateLoadBalancer()
	return &UpdateLoadBalancerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLogtank 更新云日志
//
// 更新云日志。[荷兰region不支持云日志功能，请勿使用。](tag:dt,dt_test)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) UpdateLogtank(request *model.UpdateLogtankRequest) (*model.UpdateLogtankResponse, error) {
	requestDef := GenReqDefForUpdateLogtank()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLogtankResponse), nil
	}
}

// UpdateLogtankInvoker 更新云日志
func (c *ElbClient) UpdateLogtankInvoker(request *model.UpdateLogtankRequest) *UpdateLogtankInvoker {
	requestDef := GenReqDefForUpdateLogtank()
	return &UpdateLogtankInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMember 更新后端服务器
//
// 更新后端服务器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) UpdateMember(request *model.UpdateMemberRequest) (*model.UpdateMemberResponse, error) {
	requestDef := GenReqDefForUpdateMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMemberResponse), nil
	}
}

// UpdateMemberInvoker 更新后端服务器
func (c *ElbClient) UpdateMemberInvoker(request *model.UpdateMemberRequest) *UpdateMemberInvoker {
	requestDef := GenReqDefForUpdateMember()
	return &UpdateMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePool 更新后端服务器组
//
// 更新后端服务器组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) UpdatePool(request *model.UpdatePoolRequest) (*model.UpdatePoolResponse, error) {
	requestDef := GenReqDefForUpdatePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePoolResponse), nil
	}
}

// UpdatePoolInvoker 更新后端服务器组
func (c *ElbClient) UpdatePoolInvoker(request *model.UpdatePoolRequest) *UpdatePoolInvoker {
	requestDef := GenReqDefForUpdatePool()
	return &UpdatePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSecurityPolicy 更新自定义安全策略
//
// 更新自定义安全策略。[荷兰region不支持自定义安全策略功能，请勿使用。](tag:dt,dt_test)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) UpdateSecurityPolicy(request *model.UpdateSecurityPolicyRequest) (*model.UpdateSecurityPolicyResponse, error) {
	requestDef := GenReqDefForUpdateSecurityPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSecurityPolicyResponse), nil
	}
}

// UpdateSecurityPolicyInvoker 更新自定义安全策略
func (c *ElbClient) UpdateSecurityPolicyInvoker(request *model.UpdateSecurityPolicyRequest) *UpdateSecurityPolicyInvoker {
	requestDef := GenReqDefForUpdateSecurityPolicy()
	return &UpdateSecurityPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiVersions 查询API版本列表信息
//
// 返回ELB当前所有可用的API版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ListApiVersions(request *model.ListApiVersionsRequest) (*model.ListApiVersionsResponse, error) {
	requestDef := GenReqDefForListApiVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionsResponse), nil
	}
}

// ListApiVersionsInvoker 查询API版本列表信息
func (c *ElbClient) ListApiVersionsInvoker(request *model.ListApiVersionsRequest) *ListApiVersionsInvoker {
	requestDef := GenReqDefForListApiVersions()
	return &ListApiVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteIpList 删除IP地址组的IP列表项
//
// 批量删除IP地址组的IP列表信息。[荷兰region不支持该API](tag:dt,dt_test)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) BatchDeleteIpList(request *model.BatchDeleteIpListRequest) (*model.BatchDeleteIpListResponse, error) {
	requestDef := GenReqDefForBatchDeleteIpList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteIpListResponse), nil
	}
}

// BatchDeleteIpListInvoker 删除IP地址组的IP列表项
func (c *ElbClient) BatchDeleteIpListInvoker(request *model.BatchDeleteIpListRequest) *BatchDeleteIpListInvoker {
	requestDef := GenReqDefForBatchDeleteIpList()
	return &BatchDeleteIpListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountPreoccupyIpNum 计算预占IP数
//
// 计算以下几种场景的预占用IP数量：
//
// - 计算创建LB的第一个七层监听器后总占用IP数量：
// 传入loadbalancer_id、l7_flavor_id为空、ip_target_enable不传或为false。
//
// - 计算LB规格变更或开启跨VPC后总占用IP数量：
// 传入参数loadbalancer_id，及l7_flavor_id不为空或ip_target_enable为true。
//
// - 计算创建LB所需IP数量：传入参数availability_zone_id，
// 及可选参数l7_flavor_id、ip_target_enable、ip_version，不能传loadbalancer_id。
//
// 说明：
// - 计算出来的预占IP数大于等于最终实际占用的IP数。
// - 总占用IP数量，即整个LB所占用的IP数量。
//
// [不支持传入l7_flavor_id](tag:hcso,hk_vdf,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) CountPreoccupyIpNum(request *model.CountPreoccupyIpNumRequest) (*model.CountPreoccupyIpNumResponse, error) {
	requestDef := GenReqDefForCountPreoccupyIpNum()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountPreoccupyIpNumResponse), nil
	}
}

// CountPreoccupyIpNumInvoker 计算预占IP数
func (c *ElbClient) CountPreoccupyIpNumInvoker(request *model.CountPreoccupyIpNumRequest) *CountPreoccupyIpNumInvoker {
	requestDef := GenReqDefForCountPreoccupyIpNum()
	return &CountPreoccupyIpNumInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateIpGroup 创建IP地址组
//
// 创建IP地址组。输入的ip可为ip地址或者CIDR子网，支持IPV4和IPV6。
//
// 需要注意0.0.0.0与0.0.0.0/32视为重复，0:0:0:0:0:0:0:1与::1与::1/128视为重复，只会保存其中一个。
//
// [荷兰region不支持IP地址组功能，请勿使用。](tag:dt,dt_test)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) CreateIpGroup(request *model.CreateIpGroupRequest) (*model.CreateIpGroupResponse, error) {
	requestDef := GenReqDefForCreateIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIpGroupResponse), nil
	}
}

// CreateIpGroupInvoker 创建IP地址组
func (c *ElbClient) CreateIpGroupInvoker(request *model.CreateIpGroupRequest) *CreateIpGroupInvoker {
	requestDef := GenReqDefForCreateIpGroup()
	return &CreateIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteIpGroup 删除IP地址组
//
// 删除ip地址组。[荷兰region不支持IP地址组功能，请勿使用。](tag:dt,dt_test)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) DeleteIpGroup(request *model.DeleteIpGroupRequest) (*model.DeleteIpGroupResponse, error) {
	requestDef := GenReqDefForDeleteIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteIpGroupResponse), nil
	}
}

// DeleteIpGroupInvoker 删除IP地址组
func (c *ElbClient) DeleteIpGroupInvoker(request *model.DeleteIpGroupRequest) *DeleteIpGroupInvoker {
	requestDef := GenReqDefForDeleteIpGroup()
	return &DeleteIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIpGroups 查询IP地址组列表
//
// 查询IP地址组列表。[荷兰region不支持IP地址组功能，请勿使用。](tag:dt,dt_test)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ListIpGroups(request *model.ListIpGroupsRequest) (*model.ListIpGroupsResponse, error) {
	requestDef := GenReqDefForListIpGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIpGroupsResponse), nil
	}
}

// ListIpGroupsInvoker 查询IP地址组列表
func (c *ElbClient) ListIpGroupsInvoker(request *model.ListIpGroupsRequest) *ListIpGroupsInvoker {
	requestDef := GenReqDefForListIpGroups()
	return &ListIpGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIpGroup 查询IP地址组详情
//
// 获取IP地址组详情。[荷兰region不支持IP地址组功能，请勿使用。](tag:dt,dt_test)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) ShowIpGroup(request *model.ShowIpGroupRequest) (*model.ShowIpGroupResponse, error) {
	requestDef := GenReqDefForShowIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIpGroupResponse), nil
	}
}

// ShowIpGroupInvoker 查询IP地址组详情
func (c *ElbClient) ShowIpGroupInvoker(request *model.ShowIpGroupRequest) *ShowIpGroupInvoker {
	requestDef := GenReqDefForShowIpGroup()
	return &ShowIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateIpGroup 更新IP地址组
//
// 更新IP地址组，只支持全量更新IP。即IP地址组中的ip_list将被全量覆盖，不在请求参数中的IP地址将被移除。
// 输入的ip可为ip地址或者CIDR子网，支持IPV4和IPV6。
//
// 需要注意0.0.0.0与0.0.0.0/32视为重复，0:0:0:0:0:0:0:1与::1与::1/128视为重复，只会保存其中一个。
//
// [荷兰region不支持IP地址组功能，请勿使用。](tag:dt,dt_test)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) UpdateIpGroup(request *model.UpdateIpGroupRequest) (*model.UpdateIpGroupResponse, error) {
	requestDef := GenReqDefForUpdateIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIpGroupResponse), nil
	}
}

// UpdateIpGroupInvoker 更新IP地址组
func (c *ElbClient) UpdateIpGroupInvoker(request *model.UpdateIpGroupRequest) *UpdateIpGroupInvoker {
	requestDef := GenReqDefForUpdateIpGroup()
	return &UpdateIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateIpList 更新IP地址组的IP列表项
//
// 添加新的IP地址到IP地址组的IP列表信息，或更新已有IP地址的描述。[荷兰region不支持该API](tag:dt,dt_test)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ElbClient) UpdateIpList(request *model.UpdateIpListRequest) (*model.UpdateIpListResponse, error) {
	requestDef := GenReqDefForUpdateIpList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIpListResponse), nil
	}
}

// UpdateIpListInvoker 更新IP地址组的IP列表项
func (c *ElbClient) UpdateIpListInvoker(request *model.UpdateIpListRequest) *UpdateIpListInvoker {
	requestDef := GenReqDefForUpdateIpList()
	return &UpdateIpListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
