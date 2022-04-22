package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/swr/v2/model"
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

// 创建镜像自动同步任务
//
// 创建镜像自动同步任务
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) CreateImageSyncRepo(request *model.CreateImageSyncRepoRequest) (*model.CreateImageSyncRepoResponse, error) {
	requestDef := GenReqDefForCreateImageSyncRepo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateImageSyncRepoResponse), nil
	}
}

// 手动同步镜像
//
// 手动同步镜像
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) CreateManualImageSyncRepo(request *model.CreateManualImageSyncRepoRequest) (*model.CreateManualImageSyncRepoResponse, error) {
	requestDef := GenReqDefForCreateManualImageSyncRepo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateManualImageSyncRepoResponse), nil
	}
}

// 创建组织
//
// 创建组织
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) CreateNamespace(request *model.CreateNamespaceRequest) (*model.CreateNamespaceResponse, error) {
	requestDef := GenReqDefForCreateNamespace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNamespaceResponse), nil
	}
}

// 创建组织权限
//
// 创建组织权限
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) CreateNamespaceAuth(request *model.CreateNamespaceAuthRequest) (*model.CreateNamespaceAuthResponse, error) {
	requestDef := GenReqDefForCreateNamespaceAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNamespaceAuthResponse), nil
	}
}

// 在组织下创建镜像仓库
//
// 在组织下创建镜像仓库。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) CreateRepo(request *model.CreateRepoRequest) (*model.CreateRepoResponse, error) {
	requestDef := GenReqDefForCreateRepo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRepoResponse), nil
	}
}

// 创建共享帐号
//
// 创建共享帐号。镜像上传后，您可以共享私有镜像给其他帐号，并授予下载该镜像的权限。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) CreateRepoDomains(request *model.CreateRepoDomainsRequest) (*model.CreateRepoDomainsResponse, error) {
	requestDef := GenReqDefForCreateRepoDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRepoDomainsResponse), nil
	}
}

// 创建镜像老化规则
//
// 创建镜像老化规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) CreateRetention(request *model.CreateRetentionRequest) (*model.CreateRetentionResponse, error) {
	requestDef := GenReqDefForCreateRetention()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRetentionResponse), nil
	}
}

// 生成临时登录指令
//
// 调用该接口，通过获取响应消息头的X-Swr-Dockerlogin的值及响应消息体的host值，可生成临时登录指令。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) CreateSecret(request *model.CreateSecretRequest) (*model.CreateSecretResponse, error) {
	requestDef := GenReqDefForCreateSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecretResponse), nil
	}
}

// 创建触发器
//
// 创建触发器
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) CreateTrigger(request *model.CreateTriggerRequest) (*model.CreateTriggerResponse, error) {
	requestDef := GenReqDefForCreateTrigger()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTriggerResponse), nil
	}
}

// 创建镜像权限
//
// 创建镜像权限
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) CreateUserRepositoryAuth(request *model.CreateUserRepositoryAuthRequest) (*model.CreateUserRepositoryAuthResponse, error) {
	requestDef := GenReqDefForCreateUserRepositoryAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateUserRepositoryAuthResponse), nil
	}
}

// 删除镜像自动同步任务
//
// 删除镜像自动同步任务
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) DeleteImageSyncRepo(request *model.DeleteImageSyncRepoRequest) (*model.DeleteImageSyncRepoResponse, error) {
	requestDef := GenReqDefForDeleteImageSyncRepo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteImageSyncRepoResponse), nil
	}
}

// 删除组织权限
//
// 删除组织权限
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) DeleteNamespaceAuth(request *model.DeleteNamespaceAuthRequest) (*model.DeleteNamespaceAuthResponse, error) {
	requestDef := GenReqDefForDeleteNamespaceAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNamespaceAuthResponse), nil
	}
}

// 删除组织
//
// 删除组织
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) DeleteNamespaces(request *model.DeleteNamespacesRequest) (*model.DeleteNamespacesResponse, error) {
	requestDef := GenReqDefForDeleteNamespaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNamespacesResponse), nil
	}
}

// 删除组织下的镜像仓库
//
// 删除组织下的镜像仓库。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) DeleteRepo(request *model.DeleteRepoRequest) (*model.DeleteRepoResponse, error) {
	requestDef := GenReqDefForDeleteRepo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRepoResponse), nil
	}
}

// 删除共享帐号
//
// 删除共享帐号
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) DeleteRepoDomains(request *model.DeleteRepoDomainsRequest) (*model.DeleteRepoDomainsResponse, error) {
	requestDef := GenReqDefForDeleteRepoDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRepoDomainsResponse), nil
	}
}

// 删除指定tag的镜像
//
// 删除镜像仓库中指定tag的镜像
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) DeleteRepoTag(request *model.DeleteRepoTagRequest) (*model.DeleteRepoTagResponse, error) {
	requestDef := GenReqDefForDeleteRepoTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRepoTagResponse), nil
	}
}

// 删除镜像老化规则
//
// 删除镜像老化规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) DeleteRetention(request *model.DeleteRetentionRequest) (*model.DeleteRetentionResponse, error) {
	requestDef := GenReqDefForDeleteRetention()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRetentionResponse), nil
	}
}

// 删除触发器
//
// 删除触发器
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) DeleteTrigger(request *model.DeleteTriggerRequest) (*model.DeleteTriggerResponse, error) {
	requestDef := GenReqDefForDeleteTrigger()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTriggerResponse), nil
	}
}

// 删除镜像权限
//
// 删除镜像权限
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) DeleteUserRepositoryAuth(request *model.DeleteUserRepositoryAuthRequest) (*model.DeleteUserRepositoryAuthResponse, error) {
	requestDef := GenReqDefForDeleteUserRepositoryAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteUserRepositoryAuthResponse), nil
	}
}

// 获取镜像自动同步任务信息
//
// 获取镜像自动同步任务列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) ListImageAutoSyncReposDetails(request *model.ListImageAutoSyncReposDetailsRequest) (*model.ListImageAutoSyncReposDetailsResponse, error) {
	requestDef := GenReqDefForListImageAutoSyncReposDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImageAutoSyncReposDetailsResponse), nil
	}
}

// 查询组织列表
//
// 查询组织列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) ListNamespaces(request *model.ListNamespacesRequest) (*model.ListNamespacesResponse, error) {
	requestDef := GenReqDefForListNamespaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNamespacesResponse), nil
	}
}

// 获取配额信息
//
// 获取配额信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) ListQuotas(request *model.ListQuotasRequest) (*model.ListQuotasResponse, error) {
	requestDef := GenReqDefForListQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotasResponse), nil
	}
}

// 获取共享帐号列表
//
// 获取共享帐号列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) ListRepoDomains(request *model.ListRepoDomainsRequest) (*model.ListRepoDomainsResponse, error) {
	requestDef := GenReqDefForListRepoDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepoDomainsResponse), nil
	}
}

// 查询镜像仓库列表
//
// 查询镜像仓库列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) ListReposDetails(request *model.ListReposDetailsRequest) (*model.ListReposDetailsResponse, error) {
	requestDef := GenReqDefForListReposDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListReposDetailsResponse), nil
	}
}

// 查询镜像tag列表
//
// 查询镜像tag列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) ListRepositoryTags(request *model.ListRepositoryTagsRequest) (*model.ListRepositoryTagsResponse, error) {
	requestDef := GenReqDefForListRepositoryTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryTagsResponse), nil
	}
}

// 获取镜像老化记录
//
// 获取镜像老化记录
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) ListRetentionHistories(request *model.ListRetentionHistoriesRequest) (*model.ListRetentionHistoriesResponse, error) {
	requestDef := GenReqDefForListRetentionHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRetentionHistoriesResponse), nil
	}
}

// 获取镜像老化规则列表
//
// 获取镜像老化规则列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) ListRetentions(request *model.ListRetentionsRequest) (*model.ListRetentionsResponse, error) {
	requestDef := GenReqDefForListRetentions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRetentionsResponse), nil
	}
}

// 查询共享镜像列表
//
// 查询共享镜像列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) ListSharedReposDetails(request *model.ListSharedReposDetailsRequest) (*model.ListSharedReposDetailsResponse, error) {
	requestDef := GenReqDefForListSharedReposDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSharedReposDetailsResponse), nil
	}
}

// 获取镜像仓库下的触发器列表
//
// 获取镜像仓库下的触发器列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) ListTriggersDetails(request *model.ListTriggersDetailsRequest) (*model.ListTriggersDetailsResponse, error) {
	requestDef := GenReqDefForListTriggersDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTriggersDetailsResponse), nil
	}
}

// 判断共享帐号是否存在
//
// 判断共享租户是否存在
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) ShowAccessDomain(request *model.ShowAccessDomainRequest) (*model.ShowAccessDomainResponse, error) {
	requestDef := GenReqDefForShowAccessDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAccessDomainResponse), nil
	}
}

// 获取组织详情
//
// 获取组织详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) ShowNamespace(request *model.ShowNamespaceRequest) (*model.ShowNamespaceResponse, error) {
	requestDef := GenReqDefForShowNamespace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNamespaceResponse), nil
	}
}

// 查询组织权限
//
// 查询组织权限
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) ShowNamespaceAuth(request *model.ShowNamespaceAuthRequest) (*model.ShowNamespaceAuthResponse, error) {
	requestDef := GenReqDefForShowNamespaceAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNamespaceAuthResponse), nil
	}
}

// 查询镜像仓库概要信息
//
// 查询镜像仓库概要信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) ShowRepository(request *model.ShowRepositoryRequest) (*model.ShowRepositoryResponse, error) {
	requestDef := GenReqDefForShowRepository()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepositoryResponse), nil
	}
}

// 获取镜像老化规则记录
//
// 获取镜像老化规则记录
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) ShowRetention(request *model.ShowRetentionRequest) (*model.ShowRetentionResponse, error) {
	requestDef := GenReqDefForShowRetention()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRetentionResponse), nil
	}
}

// 获取镜像同步任务信息
//
// 获取镜像同步任务信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) ShowSyncJob(request *model.ShowSyncJobRequest) (*model.ShowSyncJobResponse, error) {
	requestDef := GenReqDefForShowSyncJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSyncJobResponse), nil
	}
}

// 获取触发器详情
//
// 获取触发器详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) ShowTrigger(request *model.ShowTriggerRequest) (*model.ShowTriggerResponse, error) {
	requestDef := GenReqDefForShowTrigger()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTriggerResponse), nil
	}
}

// 查询镜像权限
//
// 查询镜像权限
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) ShowUserRepositoryAuth(request *model.ShowUserRepositoryAuthRequest) (*model.ShowUserRepositoryAuthResponse, error) {
	requestDef := GenReqDefForShowUserRepositoryAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserRepositoryAuthResponse), nil
	}
}

// 更新组织权限
//
// 更新组织权限
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) UpdateNamespaceAuth(request *model.UpdateNamespaceAuthRequest) (*model.UpdateNamespaceAuthResponse, error) {
	requestDef := GenReqDefForUpdateNamespaceAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNamespaceAuthResponse), nil
	}
}

// 更新镜像仓库的概要信息
//
// 更新租户命名空间下的镜像概要信息，包括镜像类型、是否公有、描述信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) UpdateRepo(request *model.UpdateRepoRequest) (*model.UpdateRepoResponse, error) {
	requestDef := GenReqDefForUpdateRepo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRepoResponse), nil
	}
}

// 更新共享帐号
//
// 更新共享帐号
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) UpdateRepoDomains(request *model.UpdateRepoDomainsRequest) (*model.UpdateRepoDomainsResponse, error) {
	requestDef := GenReqDefForUpdateRepoDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRepoDomainsResponse), nil
	}
}

// 修改镜像老化规则
//
// 修改镜像老化规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) UpdateRetention(request *model.UpdateRetentionRequest) (*model.UpdateRetentionResponse, error) {
	requestDef := GenReqDefForUpdateRetention()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRetentionResponse), nil
	}
}

// 更新触发器配置
//
// 更新触发器配置
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) UpdateTrigger(request *model.UpdateTriggerRequest) (*model.UpdateTriggerResponse, error) {
	requestDef := GenReqDefForUpdateTrigger()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTriggerResponse), nil
	}
}

// 更新镜像权限
//
// 更新镜像权限
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) UpdateUserRepositoryAuth(request *model.UpdateUserRepositoryAuthRequest) (*model.UpdateUserRepositoryAuthResponse, error) {
	requestDef := GenReqDefForUpdateUserRepositoryAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserRepositoryAuthResponse), nil
	}
}

// 查询所有API版本信息
//
// 查询所有API版本信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) ListApiVersions(request *model.ListApiVersionsRequest) (*model.ListApiVersionsResponse, error) {
	requestDef := GenReqDefForListApiVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionsResponse), nil
	}
}

// 查询指定API版本信息
//
// 查询指定API版本信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SwrClient) ShowApiVersion(request *model.ShowApiVersionRequest) (*model.ShowApiVersionResponse, error) {
	requestDef := GenReqDefForShowApiVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApiVersionResponse), nil
	}
}
