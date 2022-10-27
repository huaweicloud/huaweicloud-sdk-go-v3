package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v2/model"
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

// BatchCreateListenerTags 批量添加监听器标签
//
// 批量添加监听器标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) BatchCreateListenerTags(request *model.BatchCreateListenerTagsRequest) (*model.BatchCreateListenerTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateListenerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateListenerTagsResponse), nil
	}
}

// BatchCreateListenerTagsInvoker 批量添加监听器标签
func (c *ElbClient) BatchCreateListenerTagsInvoker(request *model.BatchCreateListenerTagsRequest) *BatchCreateListenerTagsInvoker {
	requestDef := GenReqDefForBatchCreateListenerTags()
	return &BatchCreateListenerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateLoadbalancerTags 批量添加负载均衡器标签
//
// 批量添加负载均衡器标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) BatchCreateLoadbalancerTags(request *model.BatchCreateLoadbalancerTagsRequest) (*model.BatchCreateLoadbalancerTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateLoadbalancerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateLoadbalancerTagsResponse), nil
	}
}

// BatchCreateLoadbalancerTagsInvoker 批量添加负载均衡器标签
func (c *ElbClient) BatchCreateLoadbalancerTagsInvoker(request *model.BatchCreateLoadbalancerTagsRequest) *BatchCreateLoadbalancerTagsInvoker {
	requestDef := GenReqDefForBatchCreateLoadbalancerTags()
	return &BatchCreateLoadbalancerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteListenerTags 批量删除监听器标签
//
// 批量删除监听器标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) BatchDeleteListenerTags(request *model.BatchDeleteListenerTagsRequest) (*model.BatchDeleteListenerTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteListenerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteListenerTagsResponse), nil
	}
}

// BatchDeleteListenerTagsInvoker 批量删除监听器标签
func (c *ElbClient) BatchDeleteListenerTagsInvoker(request *model.BatchDeleteListenerTagsRequest) *BatchDeleteListenerTagsInvoker {
	requestDef := GenReqDefForBatchDeleteListenerTags()
	return &BatchDeleteListenerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteLoadbalancerTags 批量删除负载均衡器标签
//
// 批量删除负载均衡器标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) BatchDeleteLoadbalancerTags(request *model.BatchDeleteLoadbalancerTagsRequest) (*model.BatchDeleteLoadbalancerTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteLoadbalancerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteLoadbalancerTagsResponse), nil
	}
}

// BatchDeleteLoadbalancerTagsInvoker 批量删除负载均衡器标签
func (c *ElbClient) BatchDeleteLoadbalancerTagsInvoker(request *model.BatchDeleteLoadbalancerTagsRequest) *BatchDeleteLoadbalancerTagsInvoker {
	requestDef := GenReqDefForBatchDeleteLoadbalancerTags()
	return &BatchDeleteLoadbalancerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHealthmonitor 创建健康检查
//
// 给后端云服务器组添加健康检查
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) CreateHealthmonitor(request *model.CreateHealthmonitorRequest) (*model.CreateHealthmonitorResponse, error) {
	requestDef := GenReqDefForCreateHealthmonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHealthmonitorResponse), nil
	}
}

// CreateHealthmonitorInvoker 创建健康检查
func (c *ElbClient) CreateHealthmonitorInvoker(request *model.CreateHealthmonitorRequest) *CreateHealthmonitorInvoker {
	requestDef := GenReqDefForCreateHealthmonitor()
	return &CreateHealthmonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateL7policy 创建转发策略
//
// 创建listener关联的转发策略
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) CreateL7policy(request *model.CreateL7policyRequest) (*model.CreateL7policyResponse, error) {
	requestDef := GenReqDefForCreateL7policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateL7policyResponse), nil
	}
}

// CreateL7policyInvoker 创建转发策略
func (c *ElbClient) CreateL7policyInvoker(request *model.CreateL7policyRequest) *CreateL7policyInvoker {
	requestDef := GenReqDefForCreateL7policy()
	return &CreateL7policyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateL7rule 创建转发规则
//
// 创建转发规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) CreateL7rule(request *model.CreateL7ruleRequest) (*model.CreateL7ruleResponse, error) {
	requestDef := GenReqDefForCreateL7rule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateL7ruleResponse), nil
	}
}

// CreateL7ruleInvoker 创建转发规则
func (c *ElbClient) CreateL7ruleInvoker(request *model.CreateL7ruleRequest) *CreateL7ruleInvoker {
	requestDef := GenReqDefForCreateL7rule()
	return &CreateL7ruleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateListener 创建监听器
//
// 创建与负载均衡器绑定的监听器。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// CreateListenerTags 添加监听器标签
//
// 给指定监听器添加标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) CreateListenerTags(request *model.CreateListenerTagsRequest) (*model.CreateListenerTagsResponse, error) {
	requestDef := GenReqDefForCreateListenerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateListenerTagsResponse), nil
	}
}

// CreateListenerTagsInvoker 添加监听器标签
func (c *ElbClient) CreateListenerTagsInvoker(request *model.CreateListenerTagsRequest) *CreateListenerTagsInvoker {
	requestDef := GenReqDefForCreateListenerTags()
	return &CreateListenerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLoadbalancer 创建负载均衡器
//
// 创建私网类型的增强型负载均衡器。创建成功后，该接口会返回创建的增强型负载均衡器的ID、所属子网ID、负载均衡器IP等详细信息。若要创建公网类型的增强型负载均衡器，还需调用创建浮动IP的接口，将浮动IP与私网负载均衡器的vip_port_id绑定。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) CreateLoadbalancer(request *model.CreateLoadbalancerRequest) (*model.CreateLoadbalancerResponse, error) {
	requestDef := GenReqDefForCreateLoadbalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLoadbalancerResponse), nil
	}
}

// CreateLoadbalancerInvoker 创建负载均衡器
func (c *ElbClient) CreateLoadbalancerInvoker(request *model.CreateLoadbalancerRequest) *CreateLoadbalancerInvoker {
	requestDef := GenReqDefForCreateLoadbalancer()
	return &CreateLoadbalancerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLoadbalancerTags 添加负载均衡器标签
//
// 给指定负载均衡器添加标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) CreateLoadbalancerTags(request *model.CreateLoadbalancerTagsRequest) (*model.CreateLoadbalancerTagsResponse, error) {
	requestDef := GenReqDefForCreateLoadbalancerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLoadbalancerTagsResponse), nil
	}
}

// CreateLoadbalancerTagsInvoker 添加负载均衡器标签
func (c *ElbClient) CreateLoadbalancerTagsInvoker(request *model.CreateLoadbalancerTagsRequest) *CreateLoadbalancerTagsInvoker {
	requestDef := GenReqDefForCreateLoadbalancerTags()
	return &CreateLoadbalancerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMember 创建后端云服务器
//
// 添加属于某个后端云服务器组的后端云服务器。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) CreateMember(request *model.CreateMemberRequest) (*model.CreateMemberResponse, error) {
	requestDef := GenReqDefForCreateMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMemberResponse), nil
	}
}

// CreateMemberInvoker 创建后端云服务器
func (c *ElbClient) CreateMemberInvoker(request *model.CreateMemberRequest) *CreateMemberInvoker {
	requestDef := GenReqDefForCreateMember()
	return &CreateMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePool 创建后端云服务器组
//
// 创建后端云服务器组。将多个后端云服务器添加到后端云服务器组中后，请求会在后端云服务器间按后端云服务器组的负载均衡算法和后端云服务器的权重来做请求分发。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) CreatePool(request *model.CreatePoolRequest) (*model.CreatePoolResponse, error) {
	requestDef := GenReqDefForCreatePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePoolResponse), nil
	}
}

// CreatePoolInvoker 创建后端云服务器组
func (c *ElbClient) CreatePoolInvoker(request *model.CreatePoolRequest) *CreatePoolInvoker {
	requestDef := GenReqDefForCreatePool()
	return &CreatePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWhitelist 创建白名单
//
// 创建白名单，控制监听器的访问权限。若开启了白名单功能，只有白名单中放通的IP可以访问该监听器的后端服务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) CreateWhitelist(request *model.CreateWhitelistRequest) (*model.CreateWhitelistResponse, error) {
	requestDef := GenReqDefForCreateWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWhitelistResponse), nil
	}
}

// CreateWhitelistInvoker 创建白名单
func (c *ElbClient) CreateWhitelistInvoker(request *model.CreateWhitelistRequest) *CreateWhitelistInvoker {
	requestDef := GenReqDefForCreateWhitelist()
	return &CreateWhitelistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHealthmonitor 删除健康检查
//
// 删除健康检查
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) DeleteHealthmonitor(request *model.DeleteHealthmonitorRequest) (*model.DeleteHealthmonitorResponse, error) {
	requestDef := GenReqDefForDeleteHealthmonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHealthmonitorResponse), nil
	}
}

// DeleteHealthmonitorInvoker 删除健康检查
func (c *ElbClient) DeleteHealthmonitorInvoker(request *model.DeleteHealthmonitorRequest) *DeleteHealthmonitorInvoker {
	requestDef := GenReqDefForDeleteHealthmonitor()
	return &DeleteHealthmonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteL7policy 删除转发策略
//
// 删除转发策略
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) DeleteL7policy(request *model.DeleteL7policyRequest) (*model.DeleteL7policyResponse, error) {
	requestDef := GenReqDefForDeleteL7policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteL7policyResponse), nil
	}
}

// DeleteL7policyInvoker 删除转发策略
func (c *ElbClient) DeleteL7policyInvoker(request *model.DeleteL7policyRequest) *DeleteL7policyInvoker {
	requestDef := GenReqDefForDeleteL7policy()
	return &DeleteL7policyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteL7rule 删除转发规则
//
// 删除转发规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) DeleteL7rule(request *model.DeleteL7ruleRequest) (*model.DeleteL7ruleResponse, error) {
	requestDef := GenReqDefForDeleteL7rule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteL7ruleResponse), nil
	}
}

// DeleteL7ruleInvoker 删除转发规则
func (c *ElbClient) DeleteL7ruleInvoker(request *model.DeleteL7ruleRequest) *DeleteL7ruleInvoker {
	requestDef := GenReqDefForDeleteL7rule()
	return &DeleteL7ruleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteListener 删除监听器
//
// 根据指定ID删除监听器。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// DeleteListenerTags 删除监听器标签
//
// 删除监听器的某个key对应的标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) DeleteListenerTags(request *model.DeleteListenerTagsRequest) (*model.DeleteListenerTagsResponse, error) {
	requestDef := GenReqDefForDeleteListenerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteListenerTagsResponse), nil
	}
}

// DeleteListenerTagsInvoker 删除监听器标签
func (c *ElbClient) DeleteListenerTagsInvoker(request *model.DeleteListenerTagsRequest) *DeleteListenerTagsInvoker {
	requestDef := GenReqDefForDeleteListenerTags()
	return &DeleteListenerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLoadbalancer 删除负载均衡
//
// 根据指定ID删除负载均衡器。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) DeleteLoadbalancer(request *model.DeleteLoadbalancerRequest) (*model.DeleteLoadbalancerResponse, error) {
	requestDef := GenReqDefForDeleteLoadbalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLoadbalancerResponse), nil
	}
}

// DeleteLoadbalancerInvoker 删除负载均衡
func (c *ElbClient) DeleteLoadbalancerInvoker(request *model.DeleteLoadbalancerRequest) *DeleteLoadbalancerInvoker {
	requestDef := GenReqDefForDeleteLoadbalancer()
	return &DeleteLoadbalancerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLoadbalancerTags 删除负载均衡标签
//
// 删除负载均衡器的某个key对应的标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) DeleteLoadbalancerTags(request *model.DeleteLoadbalancerTagsRequest) (*model.DeleteLoadbalancerTagsResponse, error) {
	requestDef := GenReqDefForDeleteLoadbalancerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLoadbalancerTagsResponse), nil
	}
}

// DeleteLoadbalancerTagsInvoker 删除负载均衡标签
func (c *ElbClient) DeleteLoadbalancerTagsInvoker(request *model.DeleteLoadbalancerTagsRequest) *DeleteLoadbalancerTagsInvoker {
	requestDef := GenReqDefForDeleteLoadbalancerTags()
	return &DeleteLoadbalancerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMember 删除后端云服务器
//
// 删除后端云服务器
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) DeleteMember(request *model.DeleteMemberRequest) (*model.DeleteMemberResponse, error) {
	requestDef := GenReqDefForDeleteMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMemberResponse), nil
	}
}

// DeleteMemberInvoker 删除后端云服务器
func (c *ElbClient) DeleteMemberInvoker(request *model.DeleteMemberRequest) *DeleteMemberInvoker {
	requestDef := GenReqDefForDeleteMember()
	return &DeleteMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePool 删除后端云服务器组
//
// 删除后端云服务器组。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) DeletePool(request *model.DeletePoolRequest) (*model.DeletePoolResponse, error) {
	requestDef := GenReqDefForDeletePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePoolResponse), nil
	}
}

// DeletePoolInvoker 删除后端云服务器组
func (c *ElbClient) DeletePoolInvoker(request *model.DeletePoolRequest) *DeletePoolInvoker {
	requestDef := GenReqDefForDeletePool()
	return &DeletePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWhitelist 删除白名单
//
// 删除白名单
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) DeleteWhitelist(request *model.DeleteWhitelistRequest) (*model.DeleteWhitelistResponse, error) {
	requestDef := GenReqDefForDeleteWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWhitelistResponse), nil
	}
}

// DeleteWhitelistInvoker 删除白名单
func (c *ElbClient) DeleteWhitelistInvoker(request *model.DeleteWhitelistRequest) *DeleteWhitelistInvoker {
	requestDef := GenReqDefForDeleteWhitelist()
	return &DeleteWhitelistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHealthmonitors 查询健康检查列表
//
// 查询健康检查列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListHealthmonitors(request *model.ListHealthmonitorsRequest) (*model.ListHealthmonitorsResponse, error) {
	requestDef := GenReqDefForListHealthmonitors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHealthmonitorsResponse), nil
	}
}

// ListHealthmonitorsInvoker 查询健康检查列表
func (c *ElbClient) ListHealthmonitorsInvoker(request *model.ListHealthmonitorsRequest) *ListHealthmonitorsInvoker {
	requestDef := GenReqDefForListHealthmonitors()
	return &ListHealthmonitorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListL7policies 查询转发策略列表
//
// 查询转发策略。支持过滤查询和分页查询。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListL7policies(request *model.ListL7policiesRequest) (*model.ListL7policiesResponse, error) {
	requestDef := GenReqDefForListL7policies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListL7policiesResponse), nil
	}
}

// ListL7policiesInvoker 查询转发策略列表
func (c *ElbClient) ListL7policiesInvoker(request *model.ListL7policiesRequest) *ListL7policiesInvoker {
	requestDef := GenReqDefForListL7policies()
	return &ListL7policiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListL7rules 查询转发规则列表
//
// 查询指定转发策略下关联的转发规则列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListL7rules(request *model.ListL7rulesRequest) (*model.ListL7rulesResponse, error) {
	requestDef := GenReqDefForListL7rules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListL7rulesResponse), nil
	}
}

// ListL7rulesInvoker 查询转发规则列表
func (c *ElbClient) ListL7rulesInvoker(request *model.ListL7rulesRequest) *ListL7rulesInvoker {
	requestDef := GenReqDefForListL7rules()
	return &ListL7rulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListListenerTags 查询所有监听器的标签列表
//
// 查询指定项目下所有监听器的标签列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListListenerTags(request *model.ListListenerTagsRequest) (*model.ListListenerTagsResponse, error) {
	requestDef := GenReqDefForListListenerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListListenerTagsResponse), nil
	}
}

// ListListenerTagsInvoker 查询所有监听器的标签列表
func (c *ElbClient) ListListenerTagsInvoker(request *model.ListListenerTagsRequest) *ListListenerTagsInvoker {
	requestDef := GenReqDefForListListenerTags()
	return &ListListenerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListListeners 查询监听器列表
//
// 查询监听器列表。支持过滤查询和分页查询。可以通过监听器ID、协议类型、监听端口号、关联的后端云服务器的IP等查询监听器。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ListListenersByTags 根据标签查询监听器
//
// 根据标签过滤查询监听器实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListListenersByTags(request *model.ListListenersByTagsRequest) (*model.ListListenersByTagsResponse, error) {
	requestDef := GenReqDefForListListenersByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListListenersByTagsResponse), nil
	}
}

// ListListenersByTagsInvoker 根据标签查询监听器
func (c *ElbClient) ListListenersByTagsInvoker(request *model.ListListenersByTagsRequest) *ListListenersByTagsInvoker {
	requestDef := GenReqDefForListListenersByTags()
	return &ListListenersByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLoadbalancerTags 查询所有负载均衡器的标签列表
//
// 查询指定项目下所有负载均衡器的标签列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListLoadbalancerTags(request *model.ListLoadbalancerTagsRequest) (*model.ListLoadbalancerTagsResponse, error) {
	requestDef := GenReqDefForListLoadbalancerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLoadbalancerTagsResponse), nil
	}
}

// ListLoadbalancerTagsInvoker 查询所有负载均衡器的标签列表
func (c *ElbClient) ListLoadbalancerTagsInvoker(request *model.ListLoadbalancerTagsRequest) *ListLoadbalancerTagsInvoker {
	requestDef := GenReqDefForListLoadbalancerTags()
	return &ListLoadbalancerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLoadbalancers 查询负载均衡列表
//
// 查询负载均衡器列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListLoadbalancers(request *model.ListLoadbalancersRequest) (*model.ListLoadbalancersResponse, error) {
	requestDef := GenReqDefForListLoadbalancers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLoadbalancersResponse), nil
	}
}

// ListLoadbalancersInvoker 查询负载均衡列表
func (c *ElbClient) ListLoadbalancersInvoker(request *model.ListLoadbalancersRequest) *ListLoadbalancersInvoker {
	requestDef := GenReqDefForListLoadbalancers()
	return &ListLoadbalancersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLoadbalancersByTags 根据标签查询负载均衡器
//
// 根据标签过滤查询负载均衡实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListLoadbalancersByTags(request *model.ListLoadbalancersByTagsRequest) (*model.ListLoadbalancersByTagsResponse, error) {
	requestDef := GenReqDefForListLoadbalancersByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLoadbalancersByTagsResponse), nil
	}
}

// ListLoadbalancersByTagsInvoker 根据标签查询负载均衡器
func (c *ElbClient) ListLoadbalancersByTagsInvoker(request *model.ListLoadbalancersByTagsRequest) *ListLoadbalancersByTagsInvoker {
	requestDef := GenReqDefForListLoadbalancersByTags()
	return &ListLoadbalancersByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMembers 查询后端云服务器列表
//
// 查询属于某个后端云服务器组的后端云服务器。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListMembers(request *model.ListMembersRequest) (*model.ListMembersResponse, error) {
	requestDef := GenReqDefForListMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMembersResponse), nil
	}
}

// ListMembersInvoker 查询后端云服务器列表
func (c *ElbClient) ListMembersInvoker(request *model.ListMembersRequest) *ListMembersInvoker {
	requestDef := GenReqDefForListMembers()
	return &ListMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPools 查询后端云服务器组列表
//
// 查询后端云服务器组列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListPools(request *model.ListPoolsRequest) (*model.ListPoolsResponse, error) {
	requestDef := GenReqDefForListPools()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPoolsResponse), nil
	}
}

// ListPoolsInvoker 查询后端云服务器组列表
func (c *ElbClient) ListPoolsInvoker(request *model.ListPoolsRequest) *ListPoolsInvoker {
	requestDef := GenReqDefForListPools()
	return &ListPoolsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWhitelists 查询白名单列表
//
// 查询白名单，支持过滤查询和分页查询。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListWhitelists(request *model.ListWhitelistsRequest) (*model.ListWhitelistsResponse, error) {
	requestDef := GenReqDefForListWhitelists()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWhitelistsResponse), nil
	}
}

// ListWhitelistsInvoker 查询白名单列表
func (c *ElbClient) ListWhitelistsInvoker(request *model.ListWhitelistsRequest) *ListWhitelistsInvoker {
	requestDef := GenReqDefForListWhitelists()
	return &ListWhitelistsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHealthmonitors 查询健康检查详情
//
// 根据指定ID查询健康检查详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ShowHealthmonitors(request *model.ShowHealthmonitorsRequest) (*model.ShowHealthmonitorsResponse, error) {
	requestDef := GenReqDefForShowHealthmonitors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHealthmonitorsResponse), nil
	}
}

// ShowHealthmonitorsInvoker 查询健康检查详情
func (c *ElbClient) ShowHealthmonitorsInvoker(request *model.ShowHealthmonitorsRequest) *ShowHealthmonitorsInvoker {
	requestDef := GenReqDefForShowHealthmonitors()
	return &ShowHealthmonitorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowL7policy 查询转发策略详情
//
// 根据指定ID查询转发策略详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ShowL7policy(request *model.ShowL7policyRequest) (*model.ShowL7policyResponse, error) {
	requestDef := GenReqDefForShowL7policy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowL7policyResponse), nil
	}
}

// ShowL7policyInvoker 查询转发策略详情
func (c *ElbClient) ShowL7policyInvoker(request *model.ShowL7policyRequest) *ShowL7policyInvoker {
	requestDef := GenReqDefForShowL7policy()
	return &ShowL7policyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowL7rule 查询转发规则详情
//
// 根据指定ID查询某转发策略下关联的转发规则详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ShowL7rule(request *model.ShowL7ruleRequest) (*model.ShowL7ruleResponse, error) {
	requestDef := GenReqDefForShowL7rule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowL7ruleResponse), nil
	}
}

// ShowL7ruleInvoker 查询转发规则详情
func (c *ElbClient) ShowL7ruleInvoker(request *model.ShowL7ruleRequest) *ShowL7ruleInvoker {
	requestDef := GenReqDefForShowL7rule()
	return &ShowL7ruleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowListener 查询监听器详情
//
// 根据指定ID查询监听器详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ShowListenerTags 查询监听器的标签详情
//
// 查询指定监听器的所有标签信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ShowListenerTags(request *model.ShowListenerTagsRequest) (*model.ShowListenerTagsResponse, error) {
	requestDef := GenReqDefForShowListenerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowListenerTagsResponse), nil
	}
}

// ShowListenerTagsInvoker 查询监听器的标签详情
func (c *ElbClient) ShowListenerTagsInvoker(request *model.ShowListenerTagsRequest) *ShowListenerTagsInvoker {
	requestDef := GenReqDefForShowListenerTags()
	return &ShowListenerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLoadbalancer 查询负载均衡详情
//
// 根据指定负载均衡器ID查询负载均衡器详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ShowLoadbalancer(request *model.ShowLoadbalancerRequest) (*model.ShowLoadbalancerResponse, error) {
	requestDef := GenReqDefForShowLoadbalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLoadbalancerResponse), nil
	}
}

// ShowLoadbalancerInvoker 查询负载均衡详情
func (c *ElbClient) ShowLoadbalancerInvoker(request *model.ShowLoadbalancerRequest) *ShowLoadbalancerInvoker {
	requestDef := GenReqDefForShowLoadbalancer()
	return &ShowLoadbalancerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLoadbalancerTags 查询负载均衡器的标签详情
//
// 查询指定负载均衡器的所有标签信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ShowLoadbalancerTags(request *model.ShowLoadbalancerTagsRequest) (*model.ShowLoadbalancerTagsResponse, error) {
	requestDef := GenReqDefForShowLoadbalancerTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLoadbalancerTagsResponse), nil
	}
}

// ShowLoadbalancerTagsInvoker 查询负载均衡器的标签详情
func (c *ElbClient) ShowLoadbalancerTagsInvoker(request *model.ShowLoadbalancerTagsRequest) *ShowLoadbalancerTagsInvoker {
	requestDef := GenReqDefForShowLoadbalancerTags()
	return &ShowLoadbalancerTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLoadbalancersStatus 查询负载均衡状态树
//
// 查询负载均衡器状态树。可通过该接口查询负载均衡器关联的监听器、后端云服务器组、后端云服务器、健康检查、转发策略、转发规则的主要信息，了解负载均衡器下资源的拓扑情况。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ShowLoadbalancersStatus(request *model.ShowLoadbalancersStatusRequest) (*model.ShowLoadbalancersStatusResponse, error) {
	requestDef := GenReqDefForShowLoadbalancersStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLoadbalancersStatusResponse), nil
	}
}

// ShowLoadbalancersStatusInvoker 查询负载均衡状态树
func (c *ElbClient) ShowLoadbalancersStatusInvoker(request *model.ShowLoadbalancersStatusRequest) *ShowLoadbalancersStatusInvoker {
	requestDef := GenReqDefForShowLoadbalancersStatus()
	return &ShowLoadbalancersStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMember 查询后端云服务器详情
//
// 根据指定ID查询后端云服务器详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ShowMember(request *model.ShowMemberRequest) (*model.ShowMemberResponse, error) {
	requestDef := GenReqDefForShowMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMemberResponse), nil
	}
}

// ShowMemberInvoker 查询后端云服务器详情
func (c *ElbClient) ShowMemberInvoker(request *model.ShowMemberRequest) *ShowMemberInvoker {
	requestDef := GenReqDefForShowMember()
	return &ShowMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPool 查询后端云服务器组详情
//
// 根据指定ID查询后端云服务器组详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ShowPool(request *model.ShowPoolRequest) (*model.ShowPoolResponse, error) {
	requestDef := GenReqDefForShowPool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPoolResponse), nil
	}
}

// ShowPoolInvoker 查询后端云服务器组详情
func (c *ElbClient) ShowPoolInvoker(request *model.ShowPoolRequest) *ShowPoolInvoker {
	requestDef := GenReqDefForShowPool()
	return &ShowPoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWhitelist 查询白名单详情
//
// 根据指定ID查询白名单详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ShowWhitelist(request *model.ShowWhitelistRequest) (*model.ShowWhitelistResponse, error) {
	requestDef := GenReqDefForShowWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWhitelistResponse), nil
	}
}

// ShowWhitelistInvoker 查询白名单详情
func (c *ElbClient) ShowWhitelistInvoker(request *model.ShowWhitelistRequest) *ShowWhitelistInvoker {
	requestDef := GenReqDefForShowWhitelist()
	return &ShowWhitelistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHealthmonitor 更新健康检查
//
// 更新健康检查
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) UpdateHealthmonitor(request *model.UpdateHealthmonitorRequest) (*model.UpdateHealthmonitorResponse, error) {
	requestDef := GenReqDefForUpdateHealthmonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHealthmonitorResponse), nil
	}
}

// UpdateHealthmonitorInvoker 更新健康检查
func (c *ElbClient) UpdateHealthmonitorInvoker(request *model.UpdateHealthmonitorRequest) *UpdateHealthmonitorInvoker {
	requestDef := GenReqDefForUpdateHealthmonitor()
	return &UpdateHealthmonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateL7policies 更新转发策略
//
// 更新转发策略
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) UpdateL7policies(request *model.UpdateL7policiesRequest) (*model.UpdateL7policiesResponse, error) {
	requestDef := GenReqDefForUpdateL7policies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateL7policiesResponse), nil
	}
}

// UpdateL7policiesInvoker 更新转发策略
func (c *ElbClient) UpdateL7policiesInvoker(request *model.UpdateL7policiesRequest) *UpdateL7policiesInvoker {
	requestDef := GenReqDefForUpdateL7policies()
	return &UpdateL7policiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateL7rule 更新转发规则
//
// 更新指定的转发规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) UpdateL7rule(request *model.UpdateL7ruleRequest) (*model.UpdateL7ruleResponse, error) {
	requestDef := GenReqDefForUpdateL7rule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateL7ruleResponse), nil
	}
}

// UpdateL7ruleInvoker 更新转发规则
func (c *ElbClient) UpdateL7ruleInvoker(request *model.UpdateL7ruleRequest) *UpdateL7ruleInvoker {
	requestDef := GenReqDefForUpdateL7rule()
	return &UpdateL7ruleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateListener 更新监听器
//
// 更新监听器。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// UpdateLoadbalancer 更新负载均衡器
//
// 更新负载均衡器。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) UpdateLoadbalancer(request *model.UpdateLoadbalancerRequest) (*model.UpdateLoadbalancerResponse, error) {
	requestDef := GenReqDefForUpdateLoadbalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLoadbalancerResponse), nil
	}
}

// UpdateLoadbalancerInvoker 更新负载均衡器
func (c *ElbClient) UpdateLoadbalancerInvoker(request *model.UpdateLoadbalancerRequest) *UpdateLoadbalancerInvoker {
	requestDef := GenReqDefForUpdateLoadbalancer()
	return &UpdateLoadbalancerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMember 更新后端云服务器
//
// 更新后端云服务器
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) UpdateMember(request *model.UpdateMemberRequest) (*model.UpdateMemberResponse, error) {
	requestDef := GenReqDefForUpdateMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMemberResponse), nil
	}
}

// UpdateMemberInvoker 更新后端云服务器
func (c *ElbClient) UpdateMemberInvoker(request *model.UpdateMemberRequest) *UpdateMemberInvoker {
	requestDef := GenReqDefForUpdateMember()
	return &UpdateMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePool 更新后端云服务器组
//
// 更新后端云服务器组。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) UpdatePool(request *model.UpdatePoolRequest) (*model.UpdatePoolResponse, error) {
	requestDef := GenReqDefForUpdatePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePoolResponse), nil
	}
}

// UpdatePoolInvoker 更新后端云服务器组
func (c *ElbClient) UpdatePoolInvoker(request *model.UpdatePoolRequest) *UpdatePoolInvoker {
	requestDef := GenReqDefForUpdatePool()
	return &UpdatePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWhitelist 更新白名单
//
// 更新白名单。可以打开或关闭白名单，或更新访问控制的IP。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) UpdateWhitelist(request *model.UpdateWhitelistRequest) (*model.UpdateWhitelistResponse, error) {
	requestDef := GenReqDefForUpdateWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWhitelistResponse), nil
	}
}

// UpdateWhitelistInvoker 更新白名单
func (c *ElbClient) UpdateWhitelistInvoker(request *model.UpdateWhitelistRequest) *UpdateWhitelistInvoker {
	requestDef := GenReqDefForUpdateWhitelist()
	return &UpdateWhitelistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCertificate 创建SSL证书
//
// 创建SSL证书。将监听器和SSL证书绑定后，可以通过负载均衡器实现服务端认证，后端服务器只要提供HTTP服务就能实现安全可靠的连接。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) CreateCertificate(request *model.CreateCertificateRequest) (*model.CreateCertificateResponse, error) {
	requestDef := GenReqDefForCreateCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCertificateResponse), nil
	}
}

// CreateCertificateInvoker 创建SSL证书
func (c *ElbClient) CreateCertificateInvoker(request *model.CreateCertificateRequest) *CreateCertificateInvoker {
	requestDef := GenReqDefForCreateCertificate()
	return &CreateCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCertificate 删除SSL证书
//
// 删除指定的SSL证书
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) DeleteCertificate(request *model.DeleteCertificateRequest) (*model.DeleteCertificateResponse, error) {
	requestDef := GenReqDefForDeleteCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCertificateResponse), nil
	}
}

// DeleteCertificateInvoker 删除SSL证书
func (c *ElbClient) DeleteCertificateInvoker(request *model.DeleteCertificateRequest) *DeleteCertificateInvoker {
	requestDef := GenReqDefForDeleteCertificate()
	return &DeleteCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCertificates 查询SSL证书列表
//
// 查询SSL证书。支持过滤查询和分页查询。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ListCertificates(request *model.ListCertificatesRequest) (*model.ListCertificatesResponse, error) {
	requestDef := GenReqDefForListCertificates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCertificatesResponse), nil
	}
}

// ListCertificatesInvoker 查询SSL证书列表
func (c *ElbClient) ListCertificatesInvoker(request *model.ListCertificatesRequest) *ListCertificatesInvoker {
	requestDef := GenReqDefForListCertificates()
	return &ListCertificatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCertificate 查询SSL证书详情
//
// 查询指定SSL证书的详情信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) ShowCertificate(request *model.ShowCertificateRequest) (*model.ShowCertificateResponse, error) {
	requestDef := GenReqDefForShowCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCertificateResponse), nil
	}
}

// ShowCertificateInvoker 查询SSL证书详情
func (c *ElbClient) ShowCertificateInvoker(request *model.ShowCertificateRequest) *ShowCertificateInvoker {
	requestDef := GenReqDefForShowCertificate()
	return &ShowCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCertificate 更新SSL证书
//
// 更新指定的SSL证书
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ElbClient) UpdateCertificate(request *model.UpdateCertificateRequest) (*model.UpdateCertificateResponse, error) {
	requestDef := GenReqDefForUpdateCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCertificateResponse), nil
	}
}

// UpdateCertificateInvoker 更新SSL证书
func (c *ElbClient) UpdateCertificateInvoker(request *model.UpdateCertificateRequest) *UpdateCertificateInvoker {
	requestDef := GenReqDefForUpdateCertificate()
	return &UpdateCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
