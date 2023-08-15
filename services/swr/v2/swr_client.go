package v2

import (
	http_client "github.com/dysodeng/huaweicloud-sdk-go-v3/core"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/swr/v2/model"
)

type SwrClient struct {
	HcClient *http_client.HcHttpClient
}

func NewSwrClient(hcClient *http_client.HcHttpClient) *SwrClient {
	return &SwrClient{HcClient: hcClient}
}

func SwrClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CreateImageSyncRepo 创建镜像自动同步任务
//
// 创建镜像自动同步任务，帮助您把最新推送的镜像自动同步到其他区域镜像仓库内。 镜像自动同步帮助您把最新推送的镜像自动同步到其他区域镜像仓库内，后期镜像有更新时，目标仓库的镜像也会自动更新，但已有的镜像不会自动同步。已有镜像的同步需要手动操作，详情请参见手动同步镜像。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) CreateImageSyncRepo(request *model.CreateImageSyncRepoRequest) (*model.CreateImageSyncRepoResponse, error) {
	requestDef := GenReqDefForCreateImageSyncRepo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateImageSyncRepoResponse), nil
	}
}

// CreateImageSyncRepoInvoker 创建镜像自动同步任务
func (c *SwrClient) CreateImageSyncRepoInvoker(request *model.CreateImageSyncRepoRequest) *CreateImageSyncRepoInvoker {
	requestDef := GenReqDefForCreateImageSyncRepo()
	return &CreateImageSyncRepoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateManualImageSyncRepo 手动同步镜像
//
// 对于镜像仓库已有的镜像，如果想在其他区域使用，需要手动触发镜像同步。 判断是否同步成功的方法如下：响应状态码为200，无报错信息，表示同步成功。通过SWR管理控制台或调用查询镜像仓库概要信息接口，在目标区域的目标组织下，若存在所同步的镜像版本表示同步成功。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) CreateManualImageSyncRepo(request *model.CreateManualImageSyncRepoRequest) (*model.CreateManualImageSyncRepoResponse, error) {
	requestDef := GenReqDefForCreateManualImageSyncRepo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateManualImageSyncRepoResponse), nil
	}
}

// CreateManualImageSyncRepoInvoker 手动同步镜像
func (c *SwrClient) CreateManualImageSyncRepoInvoker(request *model.CreateManualImageSyncRepoRequest) *CreateManualImageSyncRepoInvoker {
	requestDef := GenReqDefForCreateManualImageSyncRepo()
	return &CreateManualImageSyncRepoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateNamespace 创建组织
//
// 创建组织
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) CreateNamespace(request *model.CreateNamespaceRequest) (*model.CreateNamespaceResponse, error) {
	requestDef := GenReqDefForCreateNamespace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNamespaceResponse), nil
	}
}

// CreateNamespaceInvoker 创建组织
func (c *SwrClient) CreateNamespaceInvoker(request *model.CreateNamespaceRequest) *CreateNamespaceInvoker {
	requestDef := GenReqDefForCreateNamespace()
	return &CreateNamespaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateNamespaceAuth 创建组织权限
//
// 创建组织权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) CreateNamespaceAuth(request *model.CreateNamespaceAuthRequest) (*model.CreateNamespaceAuthResponse, error) {
	requestDef := GenReqDefForCreateNamespaceAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNamespaceAuthResponse), nil
	}
}

// CreateNamespaceAuthInvoker 创建组织权限
func (c *SwrClient) CreateNamespaceAuthInvoker(request *model.CreateNamespaceAuthRequest) *CreateNamespaceAuthInvoker {
	requestDef := GenReqDefForCreateNamespaceAuth()
	return &CreateNamespaceAuthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRepo 在组织下创建镜像仓库
//
// 在组织下创建镜像仓库。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) CreateRepo(request *model.CreateRepoRequest) (*model.CreateRepoResponse, error) {
	requestDef := GenReqDefForCreateRepo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRepoResponse), nil
	}
}

// CreateRepoInvoker 在组织下创建镜像仓库
func (c *SwrClient) CreateRepoInvoker(request *model.CreateRepoRequest) *CreateRepoInvoker {
	requestDef := GenReqDefForCreateRepo()
	return &CreateRepoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRepoDomains 创建共享帐号
//
// 创建共享帐号。镜像上传后，您可以共享私有镜像给其他帐号，并授予下载该镜像的权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) CreateRepoDomains(request *model.CreateRepoDomainsRequest) (*model.CreateRepoDomainsResponse, error) {
	requestDef := GenReqDefForCreateRepoDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRepoDomainsResponse), nil
	}
}

// CreateRepoDomainsInvoker 创建共享帐号
func (c *SwrClient) CreateRepoDomainsInvoker(request *model.CreateRepoDomainsRequest) *CreateRepoDomainsInvoker {
	requestDef := GenReqDefForCreateRepoDomains()
	return &CreateRepoDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRetention 创建镜像老化规则
//
// 创建镜像老化规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) CreateRetention(request *model.CreateRetentionRequest) (*model.CreateRetentionResponse, error) {
	requestDef := GenReqDefForCreateRetention()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRetentionResponse), nil
	}
}

// CreateRetentionInvoker 创建镜像老化规则
func (c *SwrClient) CreateRetentionInvoker(request *model.CreateRetentionRequest) *CreateRetentionInvoker {
	requestDef := GenReqDefForCreateRetention()
	return &CreateRetentionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSecret 生成临时登录指令
//
// 调用该接口，通过获取响应消息头的X-Swr-Dockerlogin的值及响应消息体的host值，可生成临时登录指令。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) CreateSecret(request *model.CreateSecretRequest) (*model.CreateSecretResponse, error) {
	requestDef := GenReqDefForCreateSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecretResponse), nil
	}
}

// CreateSecretInvoker 生成临时登录指令
func (c *SwrClient) CreateSecretInvoker(request *model.CreateSecretRequest) *CreateSecretInvoker {
	requestDef := GenReqDefForCreateSecret()
	return &CreateSecretInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTrigger 创建触发器
//
// 创建触发器
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) CreateTrigger(request *model.CreateTriggerRequest) (*model.CreateTriggerResponse, error) {
	requestDef := GenReqDefForCreateTrigger()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTriggerResponse), nil
	}
}

// CreateTriggerInvoker 创建触发器
func (c *SwrClient) CreateTriggerInvoker(request *model.CreateTriggerRequest) *CreateTriggerInvoker {
	requestDef := GenReqDefForCreateTrigger()
	return &CreateTriggerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateUserRepositoryAuth 创建镜像权限
//
// 创建镜像权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) CreateUserRepositoryAuth(request *model.CreateUserRepositoryAuthRequest) (*model.CreateUserRepositoryAuthResponse, error) {
	requestDef := GenReqDefForCreateUserRepositoryAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateUserRepositoryAuthResponse), nil
	}
}

// CreateUserRepositoryAuthInvoker 创建镜像权限
func (c *SwrClient) CreateUserRepositoryAuthInvoker(request *model.CreateUserRepositoryAuthRequest) *CreateUserRepositoryAuthInvoker {
	requestDef := GenReqDefForCreateUserRepositoryAuth()
	return &CreateUserRepositoryAuthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteImageSyncRepo 删除镜像自动同步任务
//
// 根据目标区域、目标组织删除指定的镜像自动同步任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) DeleteImageSyncRepo(request *model.DeleteImageSyncRepoRequest) (*model.DeleteImageSyncRepoResponse, error) {
	requestDef := GenReqDefForDeleteImageSyncRepo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteImageSyncRepoResponse), nil
	}
}

// DeleteImageSyncRepoInvoker 删除镜像自动同步任务
func (c *SwrClient) DeleteImageSyncRepoInvoker(request *model.DeleteImageSyncRepoRequest) *DeleteImageSyncRepoInvoker {
	requestDef := GenReqDefForDeleteImageSyncRepo()
	return &DeleteImageSyncRepoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteNamespaceAuth 删除组织权限
//
// 删除组织权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) DeleteNamespaceAuth(request *model.DeleteNamespaceAuthRequest) (*model.DeleteNamespaceAuthResponse, error) {
	requestDef := GenReqDefForDeleteNamespaceAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNamespaceAuthResponse), nil
	}
}

// DeleteNamespaceAuthInvoker 删除组织权限
func (c *SwrClient) DeleteNamespaceAuthInvoker(request *model.DeleteNamespaceAuthRequest) *DeleteNamespaceAuthInvoker {
	requestDef := GenReqDefForDeleteNamespaceAuth()
	return &DeleteNamespaceAuthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteNamespaces 删除组织
//
// 删除组织
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) DeleteNamespaces(request *model.DeleteNamespacesRequest) (*model.DeleteNamespacesResponse, error) {
	requestDef := GenReqDefForDeleteNamespaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNamespacesResponse), nil
	}
}

// DeleteNamespacesInvoker 删除组织
func (c *SwrClient) DeleteNamespacesInvoker(request *model.DeleteNamespacesRequest) *DeleteNamespacesInvoker {
	requestDef := GenReqDefForDeleteNamespaces()
	return &DeleteNamespacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRepo 删除组织下的镜像仓库
//
// 删除组织下的镜像仓库。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) DeleteRepo(request *model.DeleteRepoRequest) (*model.DeleteRepoResponse, error) {
	requestDef := GenReqDefForDeleteRepo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRepoResponse), nil
	}
}

// DeleteRepoInvoker 删除组织下的镜像仓库
func (c *SwrClient) DeleteRepoInvoker(request *model.DeleteRepoRequest) *DeleteRepoInvoker {
	requestDef := GenReqDefForDeleteRepo()
	return &DeleteRepoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRepoDomains 删除共享帐号
//
// 删除共享帐号
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) DeleteRepoDomains(request *model.DeleteRepoDomainsRequest) (*model.DeleteRepoDomainsResponse, error) {
	requestDef := GenReqDefForDeleteRepoDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRepoDomainsResponse), nil
	}
}

// DeleteRepoDomainsInvoker 删除共享帐号
func (c *SwrClient) DeleteRepoDomainsInvoker(request *model.DeleteRepoDomainsRequest) *DeleteRepoDomainsInvoker {
	requestDef := GenReqDefForDeleteRepoDomains()
	return &DeleteRepoDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRepoTag 删除指定tag的镜像
//
// 删除镜像仓库中指定tag的镜像
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) DeleteRepoTag(request *model.DeleteRepoTagRequest) (*model.DeleteRepoTagResponse, error) {
	requestDef := GenReqDefForDeleteRepoTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRepoTagResponse), nil
	}
}

// DeleteRepoTagInvoker 删除指定tag的镜像
func (c *SwrClient) DeleteRepoTagInvoker(request *model.DeleteRepoTagRequest) *DeleteRepoTagInvoker {
	requestDef := GenReqDefForDeleteRepoTag()
	return &DeleteRepoTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRetention 删除镜像老化规则
//
// 删除镜像老化规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) DeleteRetention(request *model.DeleteRetentionRequest) (*model.DeleteRetentionResponse, error) {
	requestDef := GenReqDefForDeleteRetention()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRetentionResponse), nil
	}
}

// DeleteRetentionInvoker 删除镜像老化规则
func (c *SwrClient) DeleteRetentionInvoker(request *model.DeleteRetentionRequest) *DeleteRetentionInvoker {
	requestDef := GenReqDefForDeleteRetention()
	return &DeleteRetentionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTrigger 删除触发器
//
// 删除触发器
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) DeleteTrigger(request *model.DeleteTriggerRequest) (*model.DeleteTriggerResponse, error) {
	requestDef := GenReqDefForDeleteTrigger()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTriggerResponse), nil
	}
}

// DeleteTriggerInvoker 删除触发器
func (c *SwrClient) DeleteTriggerInvoker(request *model.DeleteTriggerRequest) *DeleteTriggerInvoker {
	requestDef := GenReqDefForDeleteTrigger()
	return &DeleteTriggerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteUserRepositoryAuth 删除镜像权限
//
// 删除镜像权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) DeleteUserRepositoryAuth(request *model.DeleteUserRepositoryAuthRequest) (*model.DeleteUserRepositoryAuthResponse, error) {
	requestDef := GenReqDefForDeleteUserRepositoryAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteUserRepositoryAuthResponse), nil
	}
}

// DeleteUserRepositoryAuthInvoker 删除镜像权限
func (c *SwrClient) DeleteUserRepositoryAuthInvoker(request *model.DeleteUserRepositoryAuthRequest) *DeleteUserRepositoryAuthInvoker {
	requestDef := GenReqDefForDeleteUserRepositoryAuth()
	return &DeleteUserRepositoryAuthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImageAutoSyncReposDetails 获取镜像自动同步任务列表
//
// 获取为当前镜像仓库所创建的所有自动同步任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListImageAutoSyncReposDetails(request *model.ListImageAutoSyncReposDetailsRequest) (*model.ListImageAutoSyncReposDetailsResponse, error) {
	requestDef := GenReqDefForListImageAutoSyncReposDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImageAutoSyncReposDetailsResponse), nil
	}
}

// ListImageAutoSyncReposDetailsInvoker 获取镜像自动同步任务列表
func (c *SwrClient) ListImageAutoSyncReposDetailsInvoker(request *model.ListImageAutoSyncReposDetailsRequest) *ListImageAutoSyncReposDetailsInvoker {
	requestDef := GenReqDefForListImageAutoSyncReposDetails()
	return &ListImageAutoSyncReposDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNamespaces 查询组织列表
//
// 查询组织列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListNamespaces(request *model.ListNamespacesRequest) (*model.ListNamespacesResponse, error) {
	requestDef := GenReqDefForListNamespaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNamespacesResponse), nil
	}
}

// ListNamespacesInvoker 查询组织列表
func (c *SwrClient) ListNamespacesInvoker(request *model.ListNamespacesRequest) *ListNamespacesInvoker {
	requestDef := GenReqDefForListNamespaces()
	return &ListNamespacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuotas 获取配额信息
//
// 获取配额信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListQuotas(request *model.ListQuotasRequest) (*model.ListQuotasResponse, error) {
	requestDef := GenReqDefForListQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotasResponse), nil
	}
}

// ListQuotasInvoker 获取配额信息
func (c *SwrClient) ListQuotasInvoker(request *model.ListQuotasRequest) *ListQuotasInvoker {
	requestDef := GenReqDefForListQuotas()
	return &ListQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepoDomains 获取共享帐号列表
//
// 获取共享帐号列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListRepoDomains(request *model.ListRepoDomainsRequest) (*model.ListRepoDomainsResponse, error) {
	requestDef := GenReqDefForListRepoDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepoDomainsResponse), nil
	}
}

// ListRepoDomainsInvoker 获取共享帐号列表
func (c *SwrClient) ListRepoDomainsInvoker(request *model.ListRepoDomainsRequest) *ListRepoDomainsInvoker {
	requestDef := GenReqDefForListRepoDomains()
	return &ListRepoDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListReposDetails 查询镜像仓库列表
//
// 查询镜像仓库列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListReposDetails(request *model.ListReposDetailsRequest) (*model.ListReposDetailsResponse, error) {
	requestDef := GenReqDefForListReposDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListReposDetailsResponse), nil
	}
}

// ListReposDetailsInvoker 查询镜像仓库列表
func (c *SwrClient) ListReposDetailsInvoker(request *model.ListReposDetailsRequest) *ListReposDetailsInvoker {
	requestDef := GenReqDefForListReposDetails()
	return &ListReposDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRepositoryTags 查询镜像tag列表
//
// 查询镜像tag列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListRepositoryTags(request *model.ListRepositoryTagsRequest) (*model.ListRepositoryTagsResponse, error) {
	requestDef := GenReqDefForListRepositoryTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryTagsResponse), nil
	}
}

// ListRepositoryTagsInvoker 查询镜像tag列表
func (c *SwrClient) ListRepositoryTagsInvoker(request *model.ListRepositoryTagsRequest) *ListRepositoryTagsInvoker {
	requestDef := GenReqDefForListRepositoryTags()
	return &ListRepositoryTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRetentionHistories 获取镜像老化记录
//
// 获取镜像老化记录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListRetentionHistories(request *model.ListRetentionHistoriesRequest) (*model.ListRetentionHistoriesResponse, error) {
	requestDef := GenReqDefForListRetentionHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRetentionHistoriesResponse), nil
	}
}

// ListRetentionHistoriesInvoker 获取镜像老化记录
func (c *SwrClient) ListRetentionHistoriesInvoker(request *model.ListRetentionHistoriesRequest) *ListRetentionHistoriesInvoker {
	requestDef := GenReqDefForListRetentionHistories()
	return &ListRetentionHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRetentions 获取镜像老化规则列表
//
// 获取镜像老化规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListRetentions(request *model.ListRetentionsRequest) (*model.ListRetentionsResponse, error) {
	requestDef := GenReqDefForListRetentions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRetentionsResponse), nil
	}
}

// ListRetentionsInvoker 获取镜像老化规则列表
func (c *SwrClient) ListRetentionsInvoker(request *model.ListRetentionsRequest) *ListRetentionsInvoker {
	requestDef := GenReqDefForListRetentions()
	return &ListRetentionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSharedReposDetails 查询共享镜像列表
//
// 查询共享镜像列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListSharedReposDetails(request *model.ListSharedReposDetailsRequest) (*model.ListSharedReposDetailsResponse, error) {
	requestDef := GenReqDefForListSharedReposDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSharedReposDetailsResponse), nil
	}
}

// ListSharedReposDetailsInvoker 查询共享镜像列表
func (c *SwrClient) ListSharedReposDetailsInvoker(request *model.ListSharedReposDetailsRequest) *ListSharedReposDetailsInvoker {
	requestDef := GenReqDefForListSharedReposDetails()
	return &ListSharedReposDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTriggersDetails 获取镜像仓库下的触发器列表
//
// 获取镜像仓库下的触发器列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListTriggersDetails(request *model.ListTriggersDetailsRequest) (*model.ListTriggersDetailsResponse, error) {
	requestDef := GenReqDefForListTriggersDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTriggersDetailsResponse), nil
	}
}

// ListTriggersDetailsInvoker 获取镜像仓库下的触发器列表
func (c *SwrClient) ListTriggersDetailsInvoker(request *model.ListTriggersDetailsRequest) *ListTriggersDetailsInvoker {
	requestDef := GenReqDefForListTriggersDetails()
	return &ListTriggersDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAccessDomain 判断共享帐号是否存在
//
// 判断共享租户是否存在
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowAccessDomain(request *model.ShowAccessDomainRequest) (*model.ShowAccessDomainResponse, error) {
	requestDef := GenReqDefForShowAccessDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAccessDomainResponse), nil
	}
}

// ShowAccessDomainInvoker 判断共享帐号是否存在
func (c *SwrClient) ShowAccessDomainInvoker(request *model.ShowAccessDomainRequest) *ShowAccessDomainInvoker {
	requestDef := GenReqDefForShowAccessDomain()
	return &ShowAccessDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNamespace 获取组织详情
//
// 获取组织详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowNamespace(request *model.ShowNamespaceRequest) (*model.ShowNamespaceResponse, error) {
	requestDef := GenReqDefForShowNamespace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNamespaceResponse), nil
	}
}

// ShowNamespaceInvoker 获取组织详情
func (c *SwrClient) ShowNamespaceInvoker(request *model.ShowNamespaceRequest) *ShowNamespaceInvoker {
	requestDef := GenReqDefForShowNamespace()
	return &ShowNamespaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNamespaceAuth 查询组织权限
//
// 查询组织权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowNamespaceAuth(request *model.ShowNamespaceAuthRequest) (*model.ShowNamespaceAuthResponse, error) {
	requestDef := GenReqDefForShowNamespaceAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNamespaceAuthResponse), nil
	}
}

// ShowNamespaceAuthInvoker 查询组织权限
func (c *SwrClient) ShowNamespaceAuthInvoker(request *model.ShowNamespaceAuthRequest) *ShowNamespaceAuthInvoker {
	requestDef := GenReqDefForShowNamespaceAuth()
	return &ShowNamespaceAuthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRepository 查询镜像仓库概要信息
//
// 查询镜像仓库概要信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowRepository(request *model.ShowRepositoryRequest) (*model.ShowRepositoryResponse, error) {
	requestDef := GenReqDefForShowRepository()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryResponse), nil
	}
}

// ShowRepositoryInvoker 查询镜像仓库概要信息
func (c *SwrClient) ShowRepositoryInvoker(request *model.ShowRepositoryRequest) *ShowRepositoryInvoker {
	requestDef := GenReqDefForShowRepository()
	return &ShowRepositoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRetention 获取镜像老化规则记录
//
// 获取镜像老化规则记录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowRetention(request *model.ShowRetentionRequest) (*model.ShowRetentionResponse, error) {
	requestDef := GenReqDefForShowRetention()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRetentionResponse), nil
	}
}

// ShowRetentionInvoker 获取镜像老化规则记录
func (c *SwrClient) ShowRetentionInvoker(request *model.ShowRetentionRequest) *ShowRetentionInvoker {
	requestDef := GenReqDefForShowRetention()
	return &ShowRetentionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSyncJob 获取镜像自动同步任务信息
//
// 创建镜像自动同步任务后，可以通过此接口查询该任务的状态，以判断是否同步成功。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowSyncJob(request *model.ShowSyncJobRequest) (*model.ShowSyncJobResponse, error) {
	requestDef := GenReqDefForShowSyncJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSyncJobResponse), nil
	}
}

// ShowSyncJobInvoker 获取镜像自动同步任务信息
func (c *SwrClient) ShowSyncJobInvoker(request *model.ShowSyncJobRequest) *ShowSyncJobInvoker {
	requestDef := GenReqDefForShowSyncJob()
	return &ShowSyncJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTrigger 获取触发器详情
//
// 获取触发器详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowTrigger(request *model.ShowTriggerRequest) (*model.ShowTriggerResponse, error) {
	requestDef := GenReqDefForShowTrigger()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTriggerResponse), nil
	}
}

// ShowTriggerInvoker 获取触发器详情
func (c *SwrClient) ShowTriggerInvoker(request *model.ShowTriggerRequest) *ShowTriggerInvoker {
	requestDef := GenReqDefForShowTrigger()
	return &ShowTriggerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUserRepositoryAuth 查询镜像权限
//
// 查询镜像权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowUserRepositoryAuth(request *model.ShowUserRepositoryAuthRequest) (*model.ShowUserRepositoryAuthResponse, error) {
	requestDef := GenReqDefForShowUserRepositoryAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserRepositoryAuthResponse), nil
	}
}

// ShowUserRepositoryAuthInvoker 查询镜像权限
func (c *SwrClient) ShowUserRepositoryAuthInvoker(request *model.ShowUserRepositoryAuthRequest) *ShowUserRepositoryAuthInvoker {
	requestDef := GenReqDefForShowUserRepositoryAuth()
	return &ShowUserRepositoryAuthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNamespaceAuth 更新组织权限
//
// 更新组织权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) UpdateNamespaceAuth(request *model.UpdateNamespaceAuthRequest) (*model.UpdateNamespaceAuthResponse, error) {
	requestDef := GenReqDefForUpdateNamespaceAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNamespaceAuthResponse), nil
	}
}

// UpdateNamespaceAuthInvoker 更新组织权限
func (c *SwrClient) UpdateNamespaceAuthInvoker(request *model.UpdateNamespaceAuthRequest) *UpdateNamespaceAuthInvoker {
	requestDef := GenReqDefForUpdateNamespaceAuth()
	return &UpdateNamespaceAuthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRepo 更新镜像仓库的概要信息
//
// 更新租户组织下的镜像概要信息，包括镜像类型、是否公有、描述信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) UpdateRepo(request *model.UpdateRepoRequest) (*model.UpdateRepoResponse, error) {
	requestDef := GenReqDefForUpdateRepo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRepoResponse), nil
	}
}

// UpdateRepoInvoker 更新镜像仓库的概要信息
func (c *SwrClient) UpdateRepoInvoker(request *model.UpdateRepoRequest) *UpdateRepoInvoker {
	requestDef := GenReqDefForUpdateRepo()
	return &UpdateRepoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRepoDomains 更新共享帐号
//
// 更新共享帐号
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) UpdateRepoDomains(request *model.UpdateRepoDomainsRequest) (*model.UpdateRepoDomainsResponse, error) {
	requestDef := GenReqDefForUpdateRepoDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRepoDomainsResponse), nil
	}
}

// UpdateRepoDomainsInvoker 更新共享帐号
func (c *SwrClient) UpdateRepoDomainsInvoker(request *model.UpdateRepoDomainsRequest) *UpdateRepoDomainsInvoker {
	requestDef := GenReqDefForUpdateRepoDomains()
	return &UpdateRepoDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRetention 修改镜像老化规则
//
// 修改镜像老化规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) UpdateRetention(request *model.UpdateRetentionRequest) (*model.UpdateRetentionResponse, error) {
	requestDef := GenReqDefForUpdateRetention()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRetentionResponse), nil
	}
}

// UpdateRetentionInvoker 修改镜像老化规则
func (c *SwrClient) UpdateRetentionInvoker(request *model.UpdateRetentionRequest) *UpdateRetentionInvoker {
	requestDef := GenReqDefForUpdateRetention()
	return &UpdateRetentionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTrigger 更新触发器配置
//
// 更新触发器配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) UpdateTrigger(request *model.UpdateTriggerRequest) (*model.UpdateTriggerResponse, error) {
	requestDef := GenReqDefForUpdateTrigger()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTriggerResponse), nil
	}
}

// UpdateTriggerInvoker 更新触发器配置
func (c *SwrClient) UpdateTriggerInvoker(request *model.UpdateTriggerRequest) *UpdateTriggerInvoker {
	requestDef := GenReqDefForUpdateTrigger()
	return &UpdateTriggerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUserRepositoryAuth 更新镜像权限
//
// 更新镜像权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) UpdateUserRepositoryAuth(request *model.UpdateUserRepositoryAuthRequest) (*model.UpdateUserRepositoryAuthResponse, error) {
	requestDef := GenReqDefForUpdateUserRepositoryAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserRepositoryAuthResponse), nil
	}
}

// UpdateUserRepositoryAuthInvoker 更新镜像权限
func (c *SwrClient) UpdateUserRepositoryAuthInvoker(request *model.UpdateUserRepositoryAuthRequest) *UpdateUserRepositoryAuthInvoker {
	requestDef := GenReqDefForUpdateUserRepositoryAuth()
	return &UpdateUserRepositoryAuthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiVersions 查询所有API版本信息
//
// 查询所有API版本信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListApiVersions(request *model.ListApiVersionsRequest) (*model.ListApiVersionsResponse, error) {
	requestDef := GenReqDefForListApiVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionsResponse), nil
	}
}

// ListApiVersionsInvoker 查询所有API版本信息
func (c *SwrClient) ListApiVersionsInvoker(request *model.ListApiVersionsRequest) *ListApiVersionsInvoker {
	requestDef := GenReqDefForListApiVersions()
	return &ListApiVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApiVersion 查询指定API版本信息
//
// 查询指定API版本信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowApiVersion(request *model.ShowApiVersionRequest) (*model.ShowApiVersionResponse, error) {
	requestDef := GenReqDefForShowApiVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApiVersionResponse), nil
	}
}

// ShowApiVersionInvoker 查询指定API版本信息
func (c *SwrClient) ShowApiVersionInvoker(request *model.ShowApiVersionRequest) *ShowApiVersionInvoker {
	requestDef := GenReqDefForShowApiVersion()
	return &ShowApiVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
