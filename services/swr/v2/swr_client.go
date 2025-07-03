package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/swr/v2/model"
)

type SwrClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewSwrClient(hcClient *httpclient.HcHttpClient) *SwrClient {
	return &SwrClient{HcClient: hcClient}
}

func SwrClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateAuthorizationToken 生成增强型登录指令(新)
//
// 调用该接口，通过获取响应消息头的X-Swr-Dockerlogin的值及响应消息体的host值，可生成增强型登录指令,注：此接口只支持IAM新平面的调用方式。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) CreateAuthorizationToken(request *model.CreateAuthorizationTokenRequest) (*model.CreateAuthorizationTokenResponse, error) {
	requestDef := GenReqDefForCreateAuthorizationToken()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAuthorizationTokenResponse), nil
	}
}

// CreateAuthorizationTokenInvoker 生成增强型登录指令(新)
func (c *SwrClient) CreateAuthorizationTokenInvoker(request *model.CreateAuthorizationTokenRequest) *CreateAuthorizationTokenInvoker {
	requestDef := GenReqDefForCreateAuthorizationToken()
	return &CreateAuthorizationTokenInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CreateRepoTag 创建镜像tag
//
// 创建镜像tag
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) CreateRepoTag(request *model.CreateRepoTagRequest) (*model.CreateRepoTagResponse, error) {
	requestDef := GenReqDefForCreateRepoTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRepoTagResponse), nil
	}
}

// CreateRepoTagInvoker 创建镜像tag
func (c *SwrClient) CreateRepoTagInvoker(request *model.CreateRepoTagRequest) *CreateRepoTagInvoker {
	requestDef := GenReqDefForCreateRepoTag()
	return &CreateRepoTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListRepoDetails 查询镜像仓库列表详情
//
// 查询镜像仓库列表详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListRepoDetails(request *model.ListRepoDetailsRequest) (*model.ListRepoDetailsResponse, error) {
	requestDef := GenReqDefForListRepoDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepoDetailsResponse), nil
	}
}

// ListRepoDetailsInvoker 查询镜像仓库列表详情
func (c *SwrClient) ListRepoDetailsInvoker(request *model.ListRepoDetailsRequest) *ListRepoDetailsInvoker {
	requestDef := GenReqDefForListRepoDetails()
	return &ListRepoDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListRepositoryTag 查询镜像tag列表详情
//
// 查询镜像tag列表详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListRepositoryTag(request *model.ListRepositoryTagRequest) (*model.ListRepositoryTagResponse, error) {
	requestDef := GenReqDefForListRepositoryTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRepositoryTagResponse), nil
	}
}

// ListRepositoryTagInvoker 查询镜像tag列表详情
func (c *SwrClient) ListRepositoryTagInvoker(request *model.ListRepositoryTagRequest) *ListRepositoryTagInvoker {
	requestDef := GenReqDefForListRepositoryTag()
	return &ListRepositoryTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListSharedRepoDetails 查询共享镜像列表详情
//
// 查询共享镜像列表详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListSharedRepoDetails(request *model.ListSharedRepoDetailsRequest) (*model.ListSharedRepoDetailsResponse, error) {
	requestDef := GenReqDefForListSharedRepoDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSharedRepoDetailsResponse), nil
	}
}

// ListSharedRepoDetailsInvoker 查询共享镜像列表详情
func (c *SwrClient) ListSharedRepoDetailsInvoker(request *model.ListSharedRepoDetailsRequest) *ListSharedRepoDetailsInvoker {
	requestDef := GenReqDefForListSharedRepoDetails()
	return &ListSharedRepoDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowDomainOverview 获取租户总览信息
//
// 获取租户总览信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowDomainOverview(request *model.ShowDomainOverviewRequest) (*model.ShowDomainOverviewResponse, error) {
	requestDef := GenReqDefForShowDomainOverview()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainOverviewResponse), nil
	}
}

// ShowDomainOverviewInvoker 获取租户总览信息
func (c *SwrClient) ShowDomainOverviewInvoker(request *model.ShowDomainOverviewRequest) *ShowDomainOverviewInvoker {
	requestDef := GenReqDefForShowDomainOverview()
	return &ShowDomainOverviewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDomainResourceReports 获取租户资源统计信息
//
// 获取租户资源统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowDomainResourceReports(request *model.ShowDomainResourceReportsRequest) (*model.ShowDomainResourceReportsResponse, error) {
	requestDef := GenReqDefForShowDomainResourceReports()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainResourceReportsResponse), nil
	}
}

// ShowDomainResourceReportsInvoker 获取租户资源统计信息
func (c *SwrClient) ShowDomainResourceReportsInvoker(request *model.ShowDomainResourceReportsRequest) *ShowDomainResourceReportsInvoker {
	requestDef := GenReqDefForShowDomainResourceReports()
	return &ShowDomainResourceReportsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowRepoTag 查询指定tag的镜像详情
//
// 查询镜像仓库中指定tag的镜像
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowRepoTag(request *model.ShowRepoTagRequest) (*model.ShowRepoTagResponse, error) {
	requestDef := GenReqDefForShowRepoTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRepoTagResponse), nil
	}
}

// ShowRepoTagInvoker 查询指定tag的镜像详情
func (c *SwrClient) ShowRepoTagInvoker(request *model.ShowRepoTagRequest) *ShowRepoTagInvoker {
	requestDef := GenReqDefForShowRepoTag()
	return &ShowRepoTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowShareFeatureGates 查询服务特性开关信息
//
// 查询服务特性开关信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowShareFeatureGates(request *model.ShowShareFeatureGatesRequest) (*model.ShowShareFeatureGatesResponse, error) {
	requestDef := GenReqDefForShowShareFeatureGates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowShareFeatureGatesResponse), nil
	}
}

// ShowShareFeatureGatesInvoker 查询服务特性开关信息
func (c *SwrClient) ShowShareFeatureGatesInvoker(request *model.ShowShareFeatureGatesRequest) *ShowShareFeatureGatesInvoker {
	requestDef := GenReqDefForShowShareFeatureGates()
	return &ShowShareFeatureGatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// AddDomainName 增加域名
//
// 增加域名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) AddDomainName(request *model.AddDomainNameRequest) (*model.AddDomainNameResponse, error) {
	requestDef := GenReqDefForAddDomainName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddDomainNameResponse), nil
	}
}

// AddDomainNameInvoker 增加域名
func (c *SwrClient) AddDomainNameInvoker(request *model.AddDomainNameRequest) *AddDomainNameInvoker {
	requestDef := GenReqDefForAddDomainName()
	return &AddDomainNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateImmutableRule 创建不可变Tag策略
//
// 创建不可变Tag策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) CreateImmutableRule(request *model.CreateImmutableRuleRequest) (*model.CreateImmutableRuleResponse, error) {
	requestDef := GenReqDefForCreateImmutableRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateImmutableRuleResponse), nil
	}
}

// CreateImmutableRuleInvoker 创建不可变Tag策略
func (c *SwrClient) CreateImmutableRuleInvoker(request *model.CreateImmutableRuleRequest) *CreateImmutableRuleInvoker {
	requestDef := GenReqDefForCreateImmutableRule()
	return &CreateImmutableRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstance 创建实例
//
// 创建企业仓库实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) CreateInstance(request *model.CreateInstanceRequest) (*model.CreateInstanceResponse, error) {
	requestDef := GenReqDefForCreateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceResponse), nil
	}
}

// CreateInstanceInvoker 创建实例
func (c *SwrClient) CreateInstanceInvoker(request *model.CreateInstanceRequest) *CreateInstanceInvoker {
	requestDef := GenReqDefForCreateInstance()
	return &CreateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstanceEndpointPolicy 开启或关闭公网访问
//
// 开启或关闭公网访问
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) CreateInstanceEndpointPolicy(request *model.CreateInstanceEndpointPolicyRequest) (*model.CreateInstanceEndpointPolicyResponse, error) {
	requestDef := GenReqDefForCreateInstanceEndpointPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceEndpointPolicyResponse), nil
	}
}

// CreateInstanceEndpointPolicyInvoker 开启或关闭公网访问
func (c *SwrClient) CreateInstanceEndpointPolicyInvoker(request *model.CreateInstanceEndpointPolicyRequest) *CreateInstanceEndpointPolicyInvoker {
	requestDef := GenReqDefForCreateInstanceEndpointPolicy()
	return &CreateInstanceEndpointPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstanceInternalEndpoint 新增内网访问
//
// 新增内网访问
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) CreateInstanceInternalEndpoint(request *model.CreateInstanceInternalEndpointRequest) (*model.CreateInstanceInternalEndpointResponse, error) {
	requestDef := GenReqDefForCreateInstanceInternalEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceInternalEndpointResponse), nil
	}
}

// CreateInstanceInternalEndpointInvoker 新增内网访问
func (c *SwrClient) CreateInstanceInternalEndpointInvoker(request *model.CreateInstanceInternalEndpointRequest) *CreateInstanceInternalEndpointInvoker {
	requestDef := GenReqDefForCreateInstanceInternalEndpoint()
	return &CreateInstanceInternalEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstanceLtCredential 创建长期访问凭证
//
// 创建长期访问凭证
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) CreateInstanceLtCredential(request *model.CreateInstanceLtCredentialRequest) (*model.CreateInstanceLtCredentialResponse, error) {
	requestDef := GenReqDefForCreateInstanceLtCredential()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceLtCredentialResponse), nil
	}
}

// CreateInstanceLtCredentialInvoker 创建长期访问凭证
func (c *SwrClient) CreateInstanceLtCredentialInvoker(request *model.CreateInstanceLtCredentialRequest) *CreateInstanceLtCredentialInvoker {
	requestDef := GenReqDefForCreateInstanceLtCredential()
	return &CreateInstanceLtCredentialInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstanceNamespace 创建命名空间
//
// 创建命名空间
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) CreateInstanceNamespace(request *model.CreateInstanceNamespaceRequest) (*model.CreateInstanceNamespaceResponse, error) {
	requestDef := GenReqDefForCreateInstanceNamespace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceNamespaceResponse), nil
	}
}

// CreateInstanceNamespaceInvoker 创建命名空间
func (c *SwrClient) CreateInstanceNamespaceInvoker(request *model.CreateInstanceNamespaceRequest) *CreateInstanceNamespaceInvoker {
	requestDef := GenReqDefForCreateInstanceNamespace()
	return &CreateInstanceNamespaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstanceRegistry 创建镜像同步的目标仓库
//
// 创建镜像同步的目标仓库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) CreateInstanceRegistry(request *model.CreateInstanceRegistryRequest) (*model.CreateInstanceRegistryResponse, error) {
	requestDef := GenReqDefForCreateInstanceRegistry()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceRegistryResponse), nil
	}
}

// CreateInstanceRegistryInvoker 创建镜像同步的目标仓库
func (c *SwrClient) CreateInstanceRegistryInvoker(request *model.CreateInstanceRegistryRequest) *CreateInstanceRegistryInvoker {
	requestDef := GenReqDefForCreateInstanceRegistry()
	return &CreateInstanceRegistryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstanceReplicationPolicy 创建镜像同步策略
//
// 创建镜像同步策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) CreateInstanceReplicationPolicy(request *model.CreateInstanceReplicationPolicyRequest) (*model.CreateInstanceReplicationPolicyResponse, error) {
	requestDef := GenReqDefForCreateInstanceReplicationPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceReplicationPolicyResponse), nil
	}
}

// CreateInstanceReplicationPolicyInvoker 创建镜像同步策略
func (c *SwrClient) CreateInstanceReplicationPolicyInvoker(request *model.CreateInstanceReplicationPolicyRequest) *CreateInstanceReplicationPolicyInvoker {
	requestDef := GenReqDefForCreateInstanceReplicationPolicy()
	return &CreateInstanceReplicationPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstanceResourceTags 批量添加资源标签
//
// 批量添加资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) CreateInstanceResourceTags(request *model.CreateInstanceResourceTagsRequest) (*model.CreateInstanceResourceTagsResponse, error) {
	requestDef := GenReqDefForCreateInstanceResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceResourceTagsResponse), nil
	}
}

// CreateInstanceResourceTagsInvoker 批量添加资源标签
func (c *SwrClient) CreateInstanceResourceTagsInvoker(request *model.CreateInstanceResourceTagsRequest) *CreateInstanceResourceTagsInvoker {
	requestDef := GenReqDefForCreateInstanceResourceTags()
	return &CreateInstanceResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstanceRetentionPolicy 创建老化策略
//
// 创建老化策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) CreateInstanceRetentionPolicy(request *model.CreateInstanceRetentionPolicyRequest) (*model.CreateInstanceRetentionPolicyResponse, error) {
	requestDef := GenReqDefForCreateInstanceRetentionPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceRetentionPolicyResponse), nil
	}
}

// CreateInstanceRetentionPolicyInvoker 创建老化策略
func (c *SwrClient) CreateInstanceRetentionPolicyInvoker(request *model.CreateInstanceRetentionPolicyRequest) *CreateInstanceRetentionPolicyInvoker {
	requestDef := GenReqDefForCreateInstanceRetentionPolicy()
	return &CreateInstanceRetentionPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstanceSignPolicy 创建签名策略
//
// 创建签名策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) CreateInstanceSignPolicy(request *model.CreateInstanceSignPolicyRequest) (*model.CreateInstanceSignPolicyResponse, error) {
	requestDef := GenReqDefForCreateInstanceSignPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceSignPolicyResponse), nil
	}
}

// CreateInstanceSignPolicyInvoker 创建签名策略
func (c *SwrClient) CreateInstanceSignPolicyInvoker(request *model.CreateInstanceSignPolicyRequest) *CreateInstanceSignPolicyInvoker {
	requestDef := GenReqDefForCreateInstanceSignPolicy()
	return &CreateInstanceSignPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstanceTempCredential 获取临时访问凭证
//
// 获取临时访问凭证
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) CreateInstanceTempCredential(request *model.CreateInstanceTempCredentialRequest) (*model.CreateInstanceTempCredentialResponse, error) {
	requestDef := GenReqDefForCreateInstanceTempCredential()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceTempCredentialResponse), nil
	}
}

// CreateInstanceTempCredentialInvoker 获取临时访问凭证
func (c *SwrClient) CreateInstanceTempCredentialInvoker(request *model.CreateInstanceTempCredentialRequest) *CreateInstanceTempCredentialInvoker {
	requestDef := GenReqDefForCreateInstanceTempCredential()
	return &CreateInstanceTempCredentialInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstanceWebhook 创建触发器策略
//
// 创建触发器策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) CreateInstanceWebhook(request *model.CreateInstanceWebhookRequest) (*model.CreateInstanceWebhookResponse, error) {
	requestDef := GenReqDefForCreateInstanceWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceWebhookResponse), nil
	}
}

// CreateInstanceWebhookInvoker 创建触发器策略
func (c *SwrClient) CreateInstanceWebhookInvoker(request *model.CreateInstanceWebhookRequest) *CreateInstanceWebhookInvoker {
	requestDef := GenReqDefForCreateInstanceWebhook()
	return &CreateInstanceWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSubResourceTags 批量添加子资源标签
//
// 批量添加子资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) CreateSubResourceTags(request *model.CreateSubResourceTagsRequest) (*model.CreateSubResourceTagsResponse, error) {
	requestDef := GenReqDefForCreateSubResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSubResourceTagsResponse), nil
	}
}

// CreateSubResourceTagsInvoker 批量添加子资源标签
func (c *SwrClient) CreateSubResourceTagsInvoker(request *model.CreateSubResourceTagsRequest) *CreateSubResourceTagsInvoker {
	requestDef := GenReqDefForCreateSubResourceTags()
	return &CreateSubResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDomainName 删除域名
//
// 删除域名，SWR企业仓库的默认域名无法删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) DeleteDomainName(request *model.DeleteDomainNameRequest) (*model.DeleteDomainNameResponse, error) {
	requestDef := GenReqDefForDeleteDomainName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDomainNameResponse), nil
	}
}

// DeleteDomainNameInvoker 删除域名
func (c *SwrClient) DeleteDomainNameInvoker(request *model.DeleteDomainNameRequest) *DeleteDomainNameInvoker {
	requestDef := GenReqDefForDeleteDomainName()
	return &DeleteDomainNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteImmutableRule 删除不可变Tag策略
//
// 删除不可变Tag策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) DeleteImmutableRule(request *model.DeleteImmutableRuleRequest) (*model.DeleteImmutableRuleResponse, error) {
	requestDef := GenReqDefForDeleteImmutableRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteImmutableRuleResponse), nil
	}
}

// DeleteImmutableRuleInvoker 删除不可变Tag策略
func (c *SwrClient) DeleteImmutableRuleInvoker(request *model.DeleteImmutableRuleRequest) *DeleteImmutableRuleInvoker {
	requestDef := GenReqDefForDeleteImmutableRule()
	return &DeleteImmutableRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstance 删除实例
//
// 删除实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResponse), nil
	}
}

// DeleteInstanceInvoker 删除实例
func (c *SwrClient) DeleteInstanceInvoker(request *model.DeleteInstanceRequest) *DeleteInstanceInvoker {
	requestDef := GenReqDefForDeleteInstance()
	return &DeleteInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstanceArtifact 删除制品
//
// 删除制品
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) DeleteInstanceArtifact(request *model.DeleteInstanceArtifactRequest) (*model.DeleteInstanceArtifactResponse, error) {
	requestDef := GenReqDefForDeleteInstanceArtifact()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceArtifactResponse), nil
	}
}

// DeleteInstanceArtifactInvoker 删除制品
func (c *SwrClient) DeleteInstanceArtifactInvoker(request *model.DeleteInstanceArtifactRequest) *DeleteInstanceArtifactInvoker {
	requestDef := GenReqDefForDeleteInstanceArtifact()
	return &DeleteInstanceArtifactInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstanceInternalEndpoint 删除内网访问
//
// 删除内网访问
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) DeleteInstanceInternalEndpoint(request *model.DeleteInstanceInternalEndpointRequest) (*model.DeleteInstanceInternalEndpointResponse, error) {
	requestDef := GenReqDefForDeleteInstanceInternalEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceInternalEndpointResponse), nil
	}
}

// DeleteInstanceInternalEndpointInvoker 删除内网访问
func (c *SwrClient) DeleteInstanceInternalEndpointInvoker(request *model.DeleteInstanceInternalEndpointRequest) *DeleteInstanceInternalEndpointInvoker {
	requestDef := GenReqDefForDeleteInstanceInternalEndpoint()
	return &DeleteInstanceInternalEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstanceJob 删除任务
//
// 删除任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) DeleteInstanceJob(request *model.DeleteInstanceJobRequest) (*model.DeleteInstanceJobResponse, error) {
	requestDef := GenReqDefForDeleteInstanceJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceJobResponse), nil
	}
}

// DeleteInstanceJobInvoker 删除任务
func (c *SwrClient) DeleteInstanceJobInvoker(request *model.DeleteInstanceJobRequest) *DeleteInstanceJobInvoker {
	requestDef := GenReqDefForDeleteInstanceJob()
	return &DeleteInstanceJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstanceLtCredential 删除长期访问凭证
//
// 删除长期访问凭证
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) DeleteInstanceLtCredential(request *model.DeleteInstanceLtCredentialRequest) (*model.DeleteInstanceLtCredentialResponse, error) {
	requestDef := GenReqDefForDeleteInstanceLtCredential()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceLtCredentialResponse), nil
	}
}

// DeleteInstanceLtCredentialInvoker 删除长期访问凭证
func (c *SwrClient) DeleteInstanceLtCredentialInvoker(request *model.DeleteInstanceLtCredentialRequest) *DeleteInstanceLtCredentialInvoker {
	requestDef := GenReqDefForDeleteInstanceLtCredential()
	return &DeleteInstanceLtCredentialInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstanceNamespace 删除命名空间
//
// 删除命名空间
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) DeleteInstanceNamespace(request *model.DeleteInstanceNamespaceRequest) (*model.DeleteInstanceNamespaceResponse, error) {
	requestDef := GenReqDefForDeleteInstanceNamespace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceNamespaceResponse), nil
	}
}

// DeleteInstanceNamespaceInvoker 删除命名空间
func (c *SwrClient) DeleteInstanceNamespaceInvoker(request *model.DeleteInstanceNamespaceRequest) *DeleteInstanceNamespaceInvoker {
	requestDef := GenReqDefForDeleteInstanceNamespace()
	return &DeleteInstanceNamespaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstanceRegistry 删除镜像同步的目标仓库
//
// 删除同步仓库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) DeleteInstanceRegistry(request *model.DeleteInstanceRegistryRequest) (*model.DeleteInstanceRegistryResponse, error) {
	requestDef := GenReqDefForDeleteInstanceRegistry()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceRegistryResponse), nil
	}
}

// DeleteInstanceRegistryInvoker 删除镜像同步的目标仓库
func (c *SwrClient) DeleteInstanceRegistryInvoker(request *model.DeleteInstanceRegistryRequest) *DeleteInstanceRegistryInvoker {
	requestDef := GenReqDefForDeleteInstanceRegistry()
	return &DeleteInstanceRegistryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstanceReplicationPolicy 删除镜像同步策略
//
// 删除镜像同步策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) DeleteInstanceReplicationPolicy(request *model.DeleteInstanceReplicationPolicyRequest) (*model.DeleteInstanceReplicationPolicyResponse, error) {
	requestDef := GenReqDefForDeleteInstanceReplicationPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceReplicationPolicyResponse), nil
	}
}

// DeleteInstanceReplicationPolicyInvoker 删除镜像同步策略
func (c *SwrClient) DeleteInstanceReplicationPolicyInvoker(request *model.DeleteInstanceReplicationPolicyRequest) *DeleteInstanceReplicationPolicyInvoker {
	requestDef := GenReqDefForDeleteInstanceReplicationPolicy()
	return &DeleteInstanceReplicationPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstanceRepository 删除制品仓库
//
// 删除仓库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) DeleteInstanceRepository(request *model.DeleteInstanceRepositoryRequest) (*model.DeleteInstanceRepositoryResponse, error) {
	requestDef := GenReqDefForDeleteInstanceRepository()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceRepositoryResponse), nil
	}
}

// DeleteInstanceRepositoryInvoker 删除制品仓库
func (c *SwrClient) DeleteInstanceRepositoryInvoker(request *model.DeleteInstanceRepositoryRequest) *DeleteInstanceRepositoryInvoker {
	requestDef := GenReqDefForDeleteInstanceRepository()
	return &DeleteInstanceRepositoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstanceResourceTags 批量删除资源标签
//
// 批量删除资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) DeleteInstanceResourceTags(request *model.DeleteInstanceResourceTagsRequest) (*model.DeleteInstanceResourceTagsResponse, error) {
	requestDef := GenReqDefForDeleteInstanceResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResourceTagsResponse), nil
	}
}

// DeleteInstanceResourceTagsInvoker 批量删除资源标签
func (c *SwrClient) DeleteInstanceResourceTagsInvoker(request *model.DeleteInstanceResourceTagsRequest) *DeleteInstanceResourceTagsInvoker {
	requestDef := GenReqDefForDeleteInstanceResourceTags()
	return &DeleteInstanceResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstanceRetentionPolicy 删除老化策略
//
// 删除老化策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) DeleteInstanceRetentionPolicy(request *model.DeleteInstanceRetentionPolicyRequest) (*model.DeleteInstanceRetentionPolicyResponse, error) {
	requestDef := GenReqDefForDeleteInstanceRetentionPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceRetentionPolicyResponse), nil
	}
}

// DeleteInstanceRetentionPolicyInvoker 删除老化策略
func (c *SwrClient) DeleteInstanceRetentionPolicyInvoker(request *model.DeleteInstanceRetentionPolicyRequest) *DeleteInstanceRetentionPolicyInvoker {
	requestDef := GenReqDefForDeleteInstanceRetentionPolicy()
	return &DeleteInstanceRetentionPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstanceSignPolicy 删除签名策略
//
// 删除签名策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) DeleteInstanceSignPolicy(request *model.DeleteInstanceSignPolicyRequest) (*model.DeleteInstanceSignPolicyResponse, error) {
	requestDef := GenReqDefForDeleteInstanceSignPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceSignPolicyResponse), nil
	}
}

// DeleteInstanceSignPolicyInvoker 删除签名策略
func (c *SwrClient) DeleteInstanceSignPolicyInvoker(request *model.DeleteInstanceSignPolicyRequest) *DeleteInstanceSignPolicyInvoker {
	requestDef := GenReqDefForDeleteInstanceSignPolicy()
	return &DeleteInstanceSignPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstanceWebhook 删除触发器策略
//
// 删除触发器策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) DeleteInstanceWebhook(request *model.DeleteInstanceWebhookRequest) (*model.DeleteInstanceWebhookResponse, error) {
	requestDef := GenReqDefForDeleteInstanceWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceWebhookResponse), nil
	}
}

// DeleteInstanceWebhookInvoker 删除触发器策略
func (c *SwrClient) DeleteInstanceWebhookInvoker(request *model.DeleteInstanceWebhookRequest) *DeleteInstanceWebhookInvoker {
	requestDef := GenReqDefForDeleteInstanceWebhook()
	return &DeleteInstanceWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSubResourceTags 批量删除子资源标签
//
// 批量删除子资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) DeleteSubResourceTags(request *model.DeleteSubResourceTagsRequest) (*model.DeleteSubResourceTagsResponse, error) {
	requestDef := GenReqDefForDeleteSubResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSubResourceTagsResponse), nil
	}
}

// DeleteSubResourceTagsInvoker 批量删除子资源标签
func (c *SwrClient) DeleteSubResourceTagsInvoker(request *model.DeleteSubResourceTagsRequest) *DeleteSubResourceTagsInvoker {
	requestDef := GenReqDefForDeleteSubResourceTags()
	return &DeleteSubResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteInstanceReplicationPolicy 手动执行镜像同步策略
//
// 手动执行同步策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ExecuteInstanceReplicationPolicy(request *model.ExecuteInstanceReplicationPolicyRequest) (*model.ExecuteInstanceReplicationPolicyResponse, error) {
	requestDef := GenReqDefForExecuteInstanceReplicationPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteInstanceReplicationPolicyResponse), nil
	}
}

// ExecuteInstanceReplicationPolicyInvoker 手动执行镜像同步策略
func (c *SwrClient) ExecuteInstanceReplicationPolicyInvoker(request *model.ExecuteInstanceReplicationPolicyRequest) *ExecuteInstanceReplicationPolicyInvoker {
	requestDef := GenReqDefForExecuteInstanceReplicationPolicy()
	return &ExecuteInstanceReplicationPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteInstanceRetentionPolicy 执行老化策略
//
// 执行老化策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ExecuteInstanceRetentionPolicy(request *model.ExecuteInstanceRetentionPolicyRequest) (*model.ExecuteInstanceRetentionPolicyResponse, error) {
	requestDef := GenReqDefForExecuteInstanceRetentionPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteInstanceRetentionPolicyResponse), nil
	}
}

// ExecuteInstanceRetentionPolicyInvoker 执行老化策略
func (c *SwrClient) ExecuteInstanceRetentionPolicyInvoker(request *model.ExecuteInstanceRetentionPolicyRequest) *ExecuteInstanceRetentionPolicyInvoker {
	requestDef := GenReqDefForExecuteInstanceRetentionPolicy()
	return &ExecuteInstanceRetentionPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteInstanceSignPolicy 手动执行签名策略
//
// 手动执行签名策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ExecuteInstanceSignPolicy(request *model.ExecuteInstanceSignPolicyRequest) (*model.ExecuteInstanceSignPolicyResponse, error) {
	requestDef := GenReqDefForExecuteInstanceSignPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteInstanceSignPolicyResponse), nil
	}
}

// ExecuteInstanceSignPolicyInvoker 手动执行签名策略
func (c *SwrClient) ExecuteInstanceSignPolicyInvoker(request *model.ExecuteInstanceSignPolicyRequest) *ExecuteInstanceSignPolicyInvoker {
	requestDef := GenReqDefForExecuteInstanceSignPolicy()
	return &ExecuteInstanceSignPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditLogs 获取上传下载的相关审计日志列表
//
// 获取上传下载的相关审计日志列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListAuditLogs(request *model.ListAuditLogsRequest) (*model.ListAuditLogsResponse, error) {
	requestDef := GenReqDefForListAuditLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditLogsResponse), nil
	}
}

// ListAuditLogsInvoker 获取上传下载的相关审计日志列表
func (c *SwrClient) ListAuditLogsInvoker(request *model.ListAuditLogsRequest) *ListAuditLogsInvoker {
	requestDef := GenReqDefForListAuditLogs()
	return &ListAuditLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDomainNames 查询所有域名列表
//
// 查询当前实例的所有域名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListDomainNames(request *model.ListDomainNamesRequest) (*model.ListDomainNamesResponse, error) {
	requestDef := GenReqDefForListDomainNames()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainNamesResponse), nil
	}
}

// ListDomainNamesInvoker 查询所有域名列表
func (c *SwrClient) ListDomainNamesInvoker(request *model.ListDomainNamesRequest) *ListDomainNamesInvoker {
	requestDef := GenReqDefForListDomainNames()
	return &ListDomainNamesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFeatureGates 查询特性开关信息
//
// 查询特性开关信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListFeatureGates(request *model.ListFeatureGatesRequest) (*model.ListFeatureGatesResponse, error) {
	requestDef := GenReqDefForListFeatureGates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFeatureGatesResponse), nil
	}
}

// ListFeatureGatesInvoker 查询特性开关信息
func (c *SwrClient) ListFeatureGatesInvoker(request *model.ListFeatureGatesRequest) *ListFeatureGatesInvoker {
	requestDef := GenReqDefForListFeatureGates()
	return &ListFeatureGatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGlobalFeatureGates 查询全局特性开关信息
//
// 查询全局特性开关信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListGlobalFeatureGates(request *model.ListGlobalFeatureGatesRequest) (*model.ListGlobalFeatureGatesResponse, error) {
	requestDef := GenReqDefForListGlobalFeatureGates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGlobalFeatureGatesResponse), nil
	}
}

// ListGlobalFeatureGatesInvoker 查询全局特性开关信息
func (c *SwrClient) ListGlobalFeatureGatesInvoker(request *model.ListGlobalFeatureGatesRequest) *ListGlobalFeatureGatesInvoker {
	requestDef := GenReqDefForListGlobalFeatureGates()
	return &ListGlobalFeatureGatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImmutableRules 获取不可变Tag策略列表
//
// 获取不可变Tag策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListImmutableRules(request *model.ListImmutableRulesRequest) (*model.ListImmutableRulesResponse, error) {
	requestDef := GenReqDefForListImmutableRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImmutableRulesResponse), nil
	}
}

// ListImmutableRulesInvoker 获取不可变Tag策略列表
func (c *SwrClient) ListImmutableRulesInvoker(request *model.ListImmutableRulesRequest) *ListImmutableRulesInvoker {
	requestDef := GenReqDefForListImmutableRules()
	return &ListImmutableRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstance 查询实例列表
//
// 查询实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstance(request *model.ListInstanceRequest) (*model.ListInstanceResponse, error) {
	requestDef := GenReqDefForListInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceResponse), nil
	}
}

// ListInstanceInvoker 查询实例列表
func (c *SwrClient) ListInstanceInvoker(request *model.ListInstanceRequest) *ListInstanceInvoker {
	requestDef := GenReqDefForListInstance()
	return &ListInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceAccessories 获取制品附件列表
//
// 获取制品附件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceAccessories(request *model.ListInstanceAccessoriesRequest) (*model.ListInstanceAccessoriesResponse, error) {
	requestDef := GenReqDefForListInstanceAccessories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceAccessoriesResponse), nil
	}
}

// ListInstanceAccessoriesInvoker 获取制品附件列表
func (c *SwrClient) ListInstanceAccessoriesInvoker(request *model.ListInstanceAccessoriesRequest) *ListInstanceAccessoriesInvoker {
	requestDef := GenReqDefForListInstanceAccessories()
	return &ListInstanceAccessoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceAllArtifacts 获取仓库实例的所有制品版本列表
//
// 获取仓库实例的所有制品版本列表（此接口只在企业仓库实例版本大于25.6.0以上的版本才支持）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceAllArtifacts(request *model.ListInstanceAllArtifactsRequest) (*model.ListInstanceAllArtifactsResponse, error) {
	requestDef := GenReqDefForListInstanceAllArtifacts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceAllArtifactsResponse), nil
	}
}

// ListInstanceAllArtifactsInvoker 获取仓库实例的所有制品版本列表
func (c *SwrClient) ListInstanceAllArtifactsInvoker(request *model.ListInstanceAllArtifactsRequest) *ListInstanceAllArtifactsInvoker {
	requestDef := GenReqDefForListInstanceAllArtifacts()
	return &ListInstanceAllArtifactsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceArtifacts 获取制品版本列表
//
// 获取制品版本列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceArtifacts(request *model.ListInstanceArtifactsRequest) (*model.ListInstanceArtifactsResponse, error) {
	requestDef := GenReqDefForListInstanceArtifacts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceArtifactsResponse), nil
	}
}

// ListInstanceArtifactsInvoker 获取制品版本列表
func (c *SwrClient) ListInstanceArtifactsInvoker(request *model.ListInstanceArtifactsRequest) *ListInstanceArtifactsInvoker {
	requestDef := GenReqDefForListInstanceArtifacts()
	return &ListInstanceArtifactsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceInternalEndpoints 获取内网访问列表
//
// 获取内网访问列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceInternalEndpoints(request *model.ListInstanceInternalEndpointsRequest) (*model.ListInstanceInternalEndpointsResponse, error) {
	requestDef := GenReqDefForListInstanceInternalEndpoints()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceInternalEndpointsResponse), nil
	}
}

// ListInstanceInternalEndpointsInvoker 获取内网访问列表
func (c *SwrClient) ListInstanceInternalEndpointsInvoker(request *model.ListInstanceInternalEndpointsRequest) *ListInstanceInternalEndpointsInvoker {
	requestDef := GenReqDefForListInstanceInternalEndpoints()
	return &ListInstanceInternalEndpointsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceJobs 获取任务列表
//
// 获取任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceJobs(request *model.ListInstanceJobsRequest) (*model.ListInstanceJobsResponse, error) {
	requestDef := GenReqDefForListInstanceJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceJobsResponse), nil
	}
}

// ListInstanceJobsInvoker 获取任务列表
func (c *SwrClient) ListInstanceJobsInvoker(request *model.ListInstanceJobsRequest) *ListInstanceJobsInvoker {
	requestDef := GenReqDefForListInstanceJobs()
	return &ListInstanceJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceLtCredentials 查询长期访问凭证列表
//
// 查询长期访问凭证列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceLtCredentials(request *model.ListInstanceLtCredentialsRequest) (*model.ListInstanceLtCredentialsResponse, error) {
	requestDef := GenReqDefForListInstanceLtCredentials()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceLtCredentialsResponse), nil
	}
}

// ListInstanceLtCredentialsInvoker 查询长期访问凭证列表
func (c *SwrClient) ListInstanceLtCredentialsInvoker(request *model.ListInstanceLtCredentialsRequest) *ListInstanceLtCredentialsInvoker {
	requestDef := GenReqDefForListInstanceLtCredentials()
	return &ListInstanceLtCredentialsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceNamespaces 获取命名空间列表
//
// 获取命名空间列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceNamespaces(request *model.ListInstanceNamespacesRequest) (*model.ListInstanceNamespacesResponse, error) {
	requestDef := GenReqDefForListInstanceNamespaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceNamespacesResponse), nil
	}
}

// ListInstanceNamespacesInvoker 获取命名空间列表
func (c *SwrClient) ListInstanceNamespacesInvoker(request *model.ListInstanceNamespacesRequest) *ListInstanceNamespacesInvoker {
	requestDef := GenReqDefForListInstanceNamespaces()
	return &ListInstanceNamespacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceProjectTags 查询项目下所有实例标签
//
// 查询项目下所有实例标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceProjectTags(request *model.ListInstanceProjectTagsRequest) (*model.ListInstanceProjectTagsResponse, error) {
	requestDef := GenReqDefForListInstanceProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceProjectTagsResponse), nil
	}
}

// ListInstanceProjectTagsInvoker 查询项目下所有实例标签
func (c *SwrClient) ListInstanceProjectTagsInvoker(request *model.ListInstanceProjectTagsRequest) *ListInstanceProjectTagsInvoker {
	requestDef := GenReqDefForListInstanceProjectTags()
	return &ListInstanceProjectTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceRegistries 获取镜像同步的目标仓库列表
//
// 获取镜像同步的目标仓库列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceRegistries(request *model.ListInstanceRegistriesRequest) (*model.ListInstanceRegistriesResponse, error) {
	requestDef := GenReqDefForListInstanceRegistries()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceRegistriesResponse), nil
	}
}

// ListInstanceRegistriesInvoker 获取镜像同步的目标仓库列表
func (c *SwrClient) ListInstanceRegistriesInvoker(request *model.ListInstanceRegistriesRequest) *ListInstanceRegistriesInvoker {
	requestDef := GenReqDefForListInstanceRegistries()
	return &ListInstanceRegistriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceReplicationPolicies 获取镜像同步策略列表
//
// 获取镜像同步策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceReplicationPolicies(request *model.ListInstanceReplicationPoliciesRequest) (*model.ListInstanceReplicationPoliciesResponse, error) {
	requestDef := GenReqDefForListInstanceReplicationPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceReplicationPoliciesResponse), nil
	}
}

// ListInstanceReplicationPoliciesInvoker 获取镜像同步策略列表
func (c *SwrClient) ListInstanceReplicationPoliciesInvoker(request *model.ListInstanceReplicationPoliciesRequest) *ListInstanceReplicationPoliciesInvoker {
	requestDef := GenReqDefForListInstanceReplicationPolicies()
	return &ListInstanceReplicationPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceReplicationPolicyExecSubTasks 获取镜像同步策略执行任务的镜像版本列表
//
// 获取镜像同步策略执行任务的镜像版本列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceReplicationPolicyExecSubTasks(request *model.ListInstanceReplicationPolicyExecSubTasksRequest) (*model.ListInstanceReplicationPolicyExecSubTasksResponse, error) {
	requestDef := GenReqDefForListInstanceReplicationPolicyExecSubTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceReplicationPolicyExecSubTasksResponse), nil
	}
}

// ListInstanceReplicationPolicyExecSubTasksInvoker 获取镜像同步策略执行任务的镜像版本列表
func (c *SwrClient) ListInstanceReplicationPolicyExecSubTasksInvoker(request *model.ListInstanceReplicationPolicyExecSubTasksRequest) *ListInstanceReplicationPolicyExecSubTasksInvoker {
	requestDef := GenReqDefForListInstanceReplicationPolicyExecSubTasks()
	return &ListInstanceReplicationPolicyExecSubTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceReplicationPolicyExecTasks 获取镜像同步策略执行任务的镜像列表
//
// 获取镜像同步策略执行任务的镜像列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceReplicationPolicyExecTasks(request *model.ListInstanceReplicationPolicyExecTasksRequest) (*model.ListInstanceReplicationPolicyExecTasksResponse, error) {
	requestDef := GenReqDefForListInstanceReplicationPolicyExecTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceReplicationPolicyExecTasksResponse), nil
	}
}

// ListInstanceReplicationPolicyExecTasksInvoker 获取镜像同步策略执行任务的镜像列表
func (c *SwrClient) ListInstanceReplicationPolicyExecTasksInvoker(request *model.ListInstanceReplicationPolicyExecTasksRequest) *ListInstanceReplicationPolicyExecTasksInvoker {
	requestDef := GenReqDefForListInstanceReplicationPolicyExecTasks()
	return &ListInstanceReplicationPolicyExecTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceReplicationPolicyExecutions 获取镜像同步策略执行记录列表
//
// 获取镜像同步策略执行记录列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceReplicationPolicyExecutions(request *model.ListInstanceReplicationPolicyExecutionsRequest) (*model.ListInstanceReplicationPolicyExecutionsResponse, error) {
	requestDef := GenReqDefForListInstanceReplicationPolicyExecutions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceReplicationPolicyExecutionsResponse), nil
	}
}

// ListInstanceReplicationPolicyExecutionsInvoker 获取镜像同步策略执行记录列表
func (c *SwrClient) ListInstanceReplicationPolicyExecutionsInvoker(request *model.ListInstanceReplicationPolicyExecutionsRequest) *ListInstanceReplicationPolicyExecutionsInvoker {
	requestDef := GenReqDefForListInstanceReplicationPolicyExecutions()
	return &ListInstanceReplicationPolicyExecutionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceRepositories 获取制品仓库列表
//
// 获取制品仓库列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceRepositories(request *model.ListInstanceRepositoriesRequest) (*model.ListInstanceRepositoriesResponse, error) {
	requestDef := GenReqDefForListInstanceRepositories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceRepositoriesResponse), nil
	}
}

// ListInstanceRepositoriesInvoker 获取制品仓库列表
func (c *SwrClient) ListInstanceRepositoriesInvoker(request *model.ListInstanceRepositoriesRequest) *ListInstanceRepositoriesInvoker {
	requestDef := GenReqDefForListInstanceRepositories()
	return &ListInstanceRepositoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceResourceInstances 按照标签检索资源列表
//
// 按照标签检索资源列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceResourceInstances(request *model.ListInstanceResourceInstancesRequest) (*model.ListInstanceResourceInstancesResponse, error) {
	requestDef := GenReqDefForListInstanceResourceInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceResourceInstancesResponse), nil
	}
}

// ListInstanceResourceInstancesInvoker 按照标签检索资源列表
func (c *SwrClient) ListInstanceResourceInstancesInvoker(request *model.ListInstanceResourceInstancesRequest) *ListInstanceResourceInstancesInvoker {
	requestDef := GenReqDefForListInstanceResourceInstances()
	return &ListInstanceResourceInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceResourceTags 查询资源标签
//
// 查询资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceResourceTags(request *model.ListInstanceResourceTagsRequest) (*model.ListInstanceResourceTagsResponse, error) {
	requestDef := GenReqDefForListInstanceResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceResourceTagsResponse), nil
	}
}

// ListInstanceResourceTagsInvoker 查询资源标签
func (c *SwrClient) ListInstanceResourceTagsInvoker(request *model.ListInstanceResourceTagsRequest) *ListInstanceResourceTagsInvoker {
	requestDef := GenReqDefForListInstanceResourceTags()
	return &ListInstanceResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceRetentionPolicies 获取老化策略列表
//
// 获取老化策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceRetentionPolicies(request *model.ListInstanceRetentionPoliciesRequest) (*model.ListInstanceRetentionPoliciesResponse, error) {
	requestDef := GenReqDefForListInstanceRetentionPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceRetentionPoliciesResponse), nil
	}
}

// ListInstanceRetentionPoliciesInvoker 获取老化策略列表
func (c *SwrClient) ListInstanceRetentionPoliciesInvoker(request *model.ListInstanceRetentionPoliciesRequest) *ListInstanceRetentionPoliciesInvoker {
	requestDef := GenReqDefForListInstanceRetentionPolicies()
	return &ListInstanceRetentionPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceRetentionPolicyExecSubTasks 获取老化策略执行任务的镜像版本列表
//
// 获取老化策略执行任务的镜像版本列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceRetentionPolicyExecSubTasks(request *model.ListInstanceRetentionPolicyExecSubTasksRequest) (*model.ListInstanceRetentionPolicyExecSubTasksResponse, error) {
	requestDef := GenReqDefForListInstanceRetentionPolicyExecSubTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceRetentionPolicyExecSubTasksResponse), nil
	}
}

// ListInstanceRetentionPolicyExecSubTasksInvoker 获取老化策略执行任务的镜像版本列表
func (c *SwrClient) ListInstanceRetentionPolicyExecSubTasksInvoker(request *model.ListInstanceRetentionPolicyExecSubTasksRequest) *ListInstanceRetentionPolicyExecSubTasksInvoker {
	requestDef := GenReqDefForListInstanceRetentionPolicyExecSubTasks()
	return &ListInstanceRetentionPolicyExecSubTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceRetentionPolicyExecTasks 获取老化策略执行任务的镜像列表
//
// 获取老化策略执行任务的镜像列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceRetentionPolicyExecTasks(request *model.ListInstanceRetentionPolicyExecTasksRequest) (*model.ListInstanceRetentionPolicyExecTasksResponse, error) {
	requestDef := GenReqDefForListInstanceRetentionPolicyExecTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceRetentionPolicyExecTasksResponse), nil
	}
}

// ListInstanceRetentionPolicyExecTasksInvoker 获取老化策略执行任务的镜像列表
func (c *SwrClient) ListInstanceRetentionPolicyExecTasksInvoker(request *model.ListInstanceRetentionPolicyExecTasksRequest) *ListInstanceRetentionPolicyExecTasksInvoker {
	requestDef := GenReqDefForListInstanceRetentionPolicyExecTasks()
	return &ListInstanceRetentionPolicyExecTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceRetentionPolicyExecutions 获取老化策略执行记录列表
//
// 获取老化策略执行记录列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceRetentionPolicyExecutions(request *model.ListInstanceRetentionPolicyExecutionsRequest) (*model.ListInstanceRetentionPolicyExecutionsResponse, error) {
	requestDef := GenReqDefForListInstanceRetentionPolicyExecutions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceRetentionPolicyExecutionsResponse), nil
	}
}

// ListInstanceRetentionPolicyExecutionsInvoker 获取老化策略执行记录列表
func (c *SwrClient) ListInstanceRetentionPolicyExecutionsInvoker(request *model.ListInstanceRetentionPolicyExecutionsRequest) *ListInstanceRetentionPolicyExecutionsInvoker {
	requestDef := GenReqDefForListInstanceRetentionPolicyExecutions()
	return &ListInstanceRetentionPolicyExecutionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceSignPolicies 获取签名策略列表
//
// 获取签名策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceSignPolicies(request *model.ListInstanceSignPoliciesRequest) (*model.ListInstanceSignPoliciesResponse, error) {
	requestDef := GenReqDefForListInstanceSignPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceSignPoliciesResponse), nil
	}
}

// ListInstanceSignPoliciesInvoker 获取签名策略列表
func (c *SwrClient) ListInstanceSignPoliciesInvoker(request *model.ListInstanceSignPoliciesRequest) *ListInstanceSignPoliciesInvoker {
	requestDef := GenReqDefForListInstanceSignPolicies()
	return &ListInstanceSignPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceSignPolicyExecTasks 获取签名执行记录任务的镜像列表
//
// 获取签名执行记录任务的镜像列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceSignPolicyExecTasks(request *model.ListInstanceSignPolicyExecTasksRequest) (*model.ListInstanceSignPolicyExecTasksResponse, error) {
	requestDef := GenReqDefForListInstanceSignPolicyExecTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceSignPolicyExecTasksResponse), nil
	}
}

// ListInstanceSignPolicyExecTasksInvoker 获取签名执行记录任务的镜像列表
func (c *SwrClient) ListInstanceSignPolicyExecTasksInvoker(request *model.ListInstanceSignPolicyExecTasksRequest) *ListInstanceSignPolicyExecTasksInvoker {
	requestDef := GenReqDefForListInstanceSignPolicyExecTasks()
	return &ListInstanceSignPolicyExecTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceSignPolicyExecutions 获取签名执行记录列表
//
// 获取签名执行记录列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceSignPolicyExecutions(request *model.ListInstanceSignPolicyExecutionsRequest) (*model.ListInstanceSignPolicyExecutionsResponse, error) {
	requestDef := GenReqDefForListInstanceSignPolicyExecutions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceSignPolicyExecutionsResponse), nil
	}
}

// ListInstanceSignPolicyExecutionsInvoker 获取签名执行记录列表
func (c *SwrClient) ListInstanceSignPolicyExecutionsInvoker(request *model.ListInstanceSignPolicyExecutionsRequest) *ListInstanceSignPolicyExecutionsInvoker {
	requestDef := GenReqDefForListInstanceSignPolicyExecutions()
	return &ListInstanceSignPolicyExecutionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceSignatureExecutionSubtasks 获取签名策略执行任务的镜像版本列表
//
// 获取签名策略执行任务的镜像版本列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceSignatureExecutionSubtasks(request *model.ListInstanceSignatureExecutionSubtasksRequest) (*model.ListInstanceSignatureExecutionSubtasksResponse, error) {
	requestDef := GenReqDefForListInstanceSignatureExecutionSubtasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceSignatureExecutionSubtasksResponse), nil
	}
}

// ListInstanceSignatureExecutionSubtasksInvoker 获取签名策略执行任务的镜像版本列表
func (c *SwrClient) ListInstanceSignatureExecutionSubtasksInvoker(request *model.ListInstanceSignatureExecutionSubtasksRequest) *ListInstanceSignatureExecutionSubtasksInvoker {
	requestDef := GenReqDefForListInstanceSignatureExecutionSubtasks()
	return &ListInstanceSignatureExecutionSubtasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceStatistics 获取实例统计数据
//
// 获取实例统计数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceStatistics(request *model.ListInstanceStatisticsRequest) (*model.ListInstanceStatisticsResponse, error) {
	requestDef := GenReqDefForListInstanceStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceStatisticsResponse), nil
	}
}

// ListInstanceStatisticsInvoker 获取实例统计数据
func (c *SwrClient) ListInstanceStatisticsInvoker(request *model.ListInstanceStatisticsRequest) *ListInstanceStatisticsInvoker {
	requestDef := GenReqDefForListInstanceStatistics()
	return &ListInstanceStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceTags 获取制品仓库的Tag列表
//
// 获取制品仓库的Tag列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceTags(request *model.ListInstanceTagsRequest) (*model.ListInstanceTagsResponse, error) {
	requestDef := GenReqDefForListInstanceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceTagsResponse), nil
	}
}

// ListInstanceTagsInvoker 获取制品仓库的Tag列表
func (c *SwrClient) ListInstanceTagsInvoker(request *model.ListInstanceTagsRequest) *ListInstanceTagsInvoker {
	requestDef := GenReqDefForListInstanceTags()
	return &ListInstanceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceWebhookJobs 获取触发器执行任务列表
//
// 获取触发器执行任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceWebhookJobs(request *model.ListInstanceWebhookJobsRequest) (*model.ListInstanceWebhookJobsResponse, error) {
	requestDef := GenReqDefForListInstanceWebhookJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceWebhookJobsResponse), nil
	}
}

// ListInstanceWebhookJobsInvoker 获取触发器执行任务列表
func (c *SwrClient) ListInstanceWebhookJobsInvoker(request *model.ListInstanceWebhookJobsRequest) *ListInstanceWebhookJobsInvoker {
	requestDef := GenReqDefForListInstanceWebhookJobs()
	return &ListInstanceWebhookJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceWebhooks 获取触发器策略列表
//
// 获取触发器策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListInstanceWebhooks(request *model.ListInstanceWebhooksRequest) (*model.ListInstanceWebhooksResponse, error) {
	requestDef := GenReqDefForListInstanceWebhooks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceWebhooksResponse), nil
	}
}

// ListInstanceWebhooksInvoker 获取触发器策略列表
func (c *SwrClient) ListInstanceWebhooksInvoker(request *model.ListInstanceWebhooksRequest) *ListInstanceWebhooksInvoker {
	requestDef := GenReqDefForListInstanceWebhooks()
	return &ListInstanceWebhooksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNamespaceRepositories 获取命名空间下所有制品仓库列表
//
// 获取命名空间下所有制品仓库列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListNamespaceRepositories(request *model.ListNamespaceRepositoriesRequest) (*model.ListNamespaceRepositoriesResponse, error) {
	requestDef := GenReqDefForListNamespaceRepositories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNamespaceRepositoriesResponse), nil
	}
}

// ListNamespaceRepositoriesInvoker 获取命名空间下所有制品仓库列表
func (c *SwrClient) ListNamespaceRepositoriesInvoker(request *model.ListNamespaceRepositoriesRequest) *ListNamespaceRepositoriesInvoker {
	requestDef := GenReqDefForListNamespaceRepositories()
	return &ListNamespaceRepositoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNamespaceTags 查询实例下所有命名空间标签
//
// 查询实例下所有命名空间标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListNamespaceTags(request *model.ListNamespaceTagsRequest) (*model.ListNamespaceTagsResponse, error) {
	requestDef := GenReqDefForListNamespaceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNamespaceTagsResponse), nil
	}
}

// ListNamespaceTagsInvoker 查询实例下所有命名空间标签
func (c *SwrClient) ListNamespaceTagsInvoker(request *model.ListNamespaceTagsRequest) *ListNamespaceTagsInvoker {
	requestDef := GenReqDefForListNamespaceTags()
	return &ListNamespaceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubResourceInstances 按照标签检索子资源列表
//
// 按照标签检索子资源列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListSubResourceInstances(request *model.ListSubResourceInstancesRequest) (*model.ListSubResourceInstancesResponse, error) {
	requestDef := GenReqDefForListSubResourceInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubResourceInstancesResponse), nil
	}
}

// ListSubResourceInstancesInvoker 按照标签检索子资源列表
func (c *SwrClient) ListSubResourceInstancesInvoker(request *model.ListSubResourceInstancesRequest) *ListSubResourceInstancesInvoker {
	requestDef := GenReqDefForListSubResourceInstances()
	return &ListSubResourceInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubResourceTags 查询子资源标签
//
// 查询子资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ListSubResourceTags(request *model.ListSubResourceTagsRequest) (*model.ListSubResourceTagsResponse, error) {
	requestDef := GenReqDefForListSubResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubResourceTagsResponse), nil
	}
}

// ListSubResourceTagsInvoker 查询子资源标签
func (c *SwrClient) ListSubResourceTagsInvoker(request *model.ListSubResourceTagsRequest) *ListSubResourceTagsInvoker {
	requestDef := GenReqDefForListSubResourceTags()
	return &ListSubResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstance 获取实例详情
//
// 获取实例详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowInstance(request *model.ShowInstanceRequest) (*model.ShowInstanceResponse, error) {
	requestDef := GenReqDefForShowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceResponse), nil
	}
}

// ShowInstanceInvoker 获取实例详情
func (c *SwrClient) ShowInstanceInvoker(request *model.ShowInstanceRequest) *ShowInstanceInvoker {
	requestDef := GenReqDefForShowInstance()
	return &ShowInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceArtifact 获取制品详情
//
// 获取制品详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowInstanceArtifact(request *model.ShowInstanceArtifactRequest) (*model.ShowInstanceArtifactResponse, error) {
	requestDef := GenReqDefForShowInstanceArtifact()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceArtifactResponse), nil
	}
}

// ShowInstanceArtifactInvoker 获取制品详情
func (c *SwrClient) ShowInstanceArtifactInvoker(request *model.ShowInstanceArtifactRequest) *ShowInstanceArtifactInvoker {
	requestDef := GenReqDefForShowInstanceArtifact()
	return &ShowInstanceArtifactInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceArtifactAddition 获取制品附加信息
//
// 获取制品附加信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowInstanceArtifactAddition(request *model.ShowInstanceArtifactAdditionRequest) (*model.ShowInstanceArtifactAdditionResponse, error) {
	requestDef := GenReqDefForShowInstanceArtifactAddition()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceArtifactAdditionResponse), nil
	}
}

// ShowInstanceArtifactAdditionInvoker 获取制品附加信息
func (c *SwrClient) ShowInstanceArtifactAdditionInvoker(request *model.ShowInstanceArtifactAdditionRequest) *ShowInstanceArtifactAdditionInvoker {
	requestDef := GenReqDefForShowInstanceArtifactAddition()
	return &ShowInstanceArtifactAdditionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceConfiguration 查看实例配置
//
// 查看实例配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowInstanceConfiguration(request *model.ShowInstanceConfigurationRequest) (*model.ShowInstanceConfigurationResponse, error) {
	requestDef := GenReqDefForShowInstanceConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceConfigurationResponse), nil
	}
}

// ShowInstanceConfigurationInvoker 查看实例配置
func (c *SwrClient) ShowInstanceConfigurationInvoker(request *model.ShowInstanceConfigurationRequest) *ShowInstanceConfigurationInvoker {
	requestDef := GenReqDefForShowInstanceConfiguration()
	return &ShowInstanceConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceEndpointPolicy 获取公网访问信息
//
// 获取公网访问信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowInstanceEndpointPolicy(request *model.ShowInstanceEndpointPolicyRequest) (*model.ShowInstanceEndpointPolicyResponse, error) {
	requestDef := GenReqDefForShowInstanceEndpointPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceEndpointPolicyResponse), nil
	}
}

// ShowInstanceEndpointPolicyInvoker 获取公网访问信息
func (c *SwrClient) ShowInstanceEndpointPolicyInvoker(request *model.ShowInstanceEndpointPolicyRequest) *ShowInstanceEndpointPolicyInvoker {
	requestDef := GenReqDefForShowInstanceEndpointPolicy()
	return &ShowInstanceEndpointPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceInternalEndpoint 查询内网访问详情
//
// 查询内网访问详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowInstanceInternalEndpoint(request *model.ShowInstanceInternalEndpointRequest) (*model.ShowInstanceInternalEndpointResponse, error) {
	requestDef := GenReqDefForShowInstanceInternalEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceInternalEndpointResponse), nil
	}
}

// ShowInstanceInternalEndpointInvoker 查询内网访问详情
func (c *SwrClient) ShowInstanceInternalEndpointInvoker(request *model.ShowInstanceInternalEndpointRequest) *ShowInstanceInternalEndpointInvoker {
	requestDef := GenReqDefForShowInstanceInternalEndpoint()
	return &ShowInstanceInternalEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceJob 获取任务详情
//
// 获取任务详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowInstanceJob(request *model.ShowInstanceJobRequest) (*model.ShowInstanceJobResponse, error) {
	requestDef := GenReqDefForShowInstanceJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceJobResponse), nil
	}
}

// ShowInstanceJobInvoker 获取任务详情
func (c *SwrClient) ShowInstanceJobInvoker(request *model.ShowInstanceJobRequest) *ShowInstanceJobInvoker {
	requestDef := GenReqDefForShowInstanceJob()
	return &ShowInstanceJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceNamespace 获取命名空间详情
//
// 获取命名空间详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowInstanceNamespace(request *model.ShowInstanceNamespaceRequest) (*model.ShowInstanceNamespaceResponse, error) {
	requestDef := GenReqDefForShowInstanceNamespace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceNamespaceResponse), nil
	}
}

// ShowInstanceNamespaceInvoker 获取命名空间详情
func (c *SwrClient) ShowInstanceNamespaceInvoker(request *model.ShowInstanceNamespaceRequest) *ShowInstanceNamespaceInvoker {
	requestDef := GenReqDefForShowInstanceNamespace()
	return &ShowInstanceNamespaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceRegistry 获取镜像同步的目标仓库详情
//
// 获取镜像同步的目标仓库详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowInstanceRegistry(request *model.ShowInstanceRegistryRequest) (*model.ShowInstanceRegistryResponse, error) {
	requestDef := GenReqDefForShowInstanceRegistry()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceRegistryResponse), nil
	}
}

// ShowInstanceRegistryInvoker 获取镜像同步的目标仓库详情
func (c *SwrClient) ShowInstanceRegistryInvoker(request *model.ShowInstanceRegistryRequest) *ShowInstanceRegistryInvoker {
	requestDef := GenReqDefForShowInstanceRegistry()
	return &ShowInstanceRegistryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceReplicationPolicy 获取镜像同步策略详情
//
// 获取镜像同步策略详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowInstanceReplicationPolicy(request *model.ShowInstanceReplicationPolicyRequest) (*model.ShowInstanceReplicationPolicyResponse, error) {
	requestDef := GenReqDefForShowInstanceReplicationPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceReplicationPolicyResponse), nil
	}
}

// ShowInstanceReplicationPolicyInvoker 获取镜像同步策略详情
func (c *SwrClient) ShowInstanceReplicationPolicyInvoker(request *model.ShowInstanceReplicationPolicyRequest) *ShowInstanceReplicationPolicyInvoker {
	requestDef := GenReqDefForShowInstanceReplicationPolicy()
	return &ShowInstanceReplicationPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceRepository 获取制品仓库详情
//
// 获取制品仓库详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowInstanceRepository(request *model.ShowInstanceRepositoryRequest) (*model.ShowInstanceRepositoryResponse, error) {
	requestDef := GenReqDefForShowInstanceRepository()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceRepositoryResponse), nil
	}
}

// ShowInstanceRepositoryInvoker 获取制品仓库详情
func (c *SwrClient) ShowInstanceRepositoryInvoker(request *model.ShowInstanceRepositoryRequest) *ShowInstanceRepositoryInvoker {
	requestDef := GenReqDefForShowInstanceRepository()
	return &ShowInstanceRepositoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceResourceInstancesCount 按照标签检索资源数量
//
// 按照标签检索资源数量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowInstanceResourceInstancesCount(request *model.ShowInstanceResourceInstancesCountRequest) (*model.ShowInstanceResourceInstancesCountResponse, error) {
	requestDef := GenReqDefForShowInstanceResourceInstancesCount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceResourceInstancesCountResponse), nil
	}
}

// ShowInstanceResourceInstancesCountInvoker 按照标签检索资源数量
func (c *SwrClient) ShowInstanceResourceInstancesCountInvoker(request *model.ShowInstanceResourceInstancesCountRequest) *ShowInstanceResourceInstancesCountInvoker {
	requestDef := GenReqDefForShowInstanceResourceInstancesCount()
	return &ShowInstanceResourceInstancesCountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceRetentionPolicy 获取老化策略详情
//
// 获取老化策略详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowInstanceRetentionPolicy(request *model.ShowInstanceRetentionPolicyRequest) (*model.ShowInstanceRetentionPolicyResponse, error) {
	requestDef := GenReqDefForShowInstanceRetentionPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceRetentionPolicyResponse), nil
	}
}

// ShowInstanceRetentionPolicyInvoker 获取老化策略详情
func (c *SwrClient) ShowInstanceRetentionPolicyInvoker(request *model.ShowInstanceRetentionPolicyRequest) *ShowInstanceRetentionPolicyInvoker {
	requestDef := GenReqDefForShowInstanceRetentionPolicy()
	return &ShowInstanceRetentionPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceSignPolicy 获取签名策略详情
//
// 获取签名策略详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowInstanceSignPolicy(request *model.ShowInstanceSignPolicyRequest) (*model.ShowInstanceSignPolicyResponse, error) {
	requestDef := GenReqDefForShowInstanceSignPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceSignPolicyResponse), nil
	}
}

// ShowInstanceSignPolicyInvoker 获取签名策略详情
func (c *SwrClient) ShowInstanceSignPolicyInvoker(request *model.ShowInstanceSignPolicyRequest) *ShowInstanceSignPolicyInvoker {
	requestDef := GenReqDefForShowInstanceSignPolicy()
	return &ShowInstanceSignPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceWebhook 获取触发器策略详情
//
// 获取触发器策略详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowInstanceWebhook(request *model.ShowInstanceWebhookRequest) (*model.ShowInstanceWebhookResponse, error) {
	requestDef := GenReqDefForShowInstanceWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceWebhookResponse), nil
	}
}

// ShowInstanceWebhookInvoker 获取触发器策略详情
func (c *SwrClient) ShowInstanceWebhookInvoker(request *model.ShowInstanceWebhookRequest) *ShowInstanceWebhookInvoker {
	requestDef := GenReqDefForShowInstanceWebhook()
	return &ShowInstanceWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSubResourceInstancesCount 按照标签检索子资源数量
//
// 按照标签检索子资源数量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) ShowSubResourceInstancesCount(request *model.ShowSubResourceInstancesCountRequest) (*model.ShowSubResourceInstancesCountResponse, error) {
	requestDef := GenReqDefForShowSubResourceInstancesCount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSubResourceInstancesCountResponse), nil
	}
}

// ShowSubResourceInstancesCountInvoker 按照标签检索子资源数量
func (c *SwrClient) ShowSubResourceInstancesCountInvoker(request *model.ShowSubResourceInstancesCountRequest) *ShowSubResourceInstancesCountInvoker {
	requestDef := GenReqDefForShowSubResourceInstancesCount()
	return &ShowSubResourceInstancesCountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopInstanceReplicationPolicyExecution 停止镜像同步任务
//
// 停止镜像同步任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) StopInstanceReplicationPolicyExecution(request *model.StopInstanceReplicationPolicyExecutionRequest) (*model.StopInstanceReplicationPolicyExecutionResponse, error) {
	requestDef := GenReqDefForStopInstanceReplicationPolicyExecution()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopInstanceReplicationPolicyExecutionResponse), nil
	}
}

// StopInstanceReplicationPolicyExecutionInvoker 停止镜像同步任务
func (c *SwrClient) StopInstanceReplicationPolicyExecutionInvoker(request *model.StopInstanceReplicationPolicyExecutionRequest) *StopInstanceReplicationPolicyExecutionInvoker {
	requestDef := GenReqDefForStopInstanceReplicationPolicyExecution()
	return &StopInstanceReplicationPolicyExecutionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDomainName 更新域名
//
// 更新域名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) UpdateDomainName(request *model.UpdateDomainNameRequest) (*model.UpdateDomainNameResponse, error) {
	requestDef := GenReqDefForUpdateDomainName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDomainNameResponse), nil
	}
}

// UpdateDomainNameInvoker 更新域名
func (c *SwrClient) UpdateDomainNameInvoker(request *model.UpdateDomainNameRequest) *UpdateDomainNameInvoker {
	requestDef := GenReqDefForUpdateDomainName()
	return &UpdateDomainNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateImmutableRule 修改不可变Tag策略
//
// 修改不可变Tag策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) UpdateImmutableRule(request *model.UpdateImmutableRuleRequest) (*model.UpdateImmutableRuleResponse, error) {
	requestDef := GenReqDefForUpdateImmutableRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateImmutableRuleResponse), nil
	}
}

// UpdateImmutableRuleInvoker 修改不可变Tag策略
func (c *SwrClient) UpdateImmutableRuleInvoker(request *model.UpdateImmutableRuleRequest) *UpdateImmutableRuleInvoker {
	requestDef := GenReqDefForUpdateImmutableRule()
	return &UpdateImmutableRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceConfiguration 修改实例配置
//
// 修改实例配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) UpdateInstanceConfiguration(request *model.UpdateInstanceConfigurationRequest) (*model.UpdateInstanceConfigurationResponse, error) {
	requestDef := GenReqDefForUpdateInstanceConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceConfigurationResponse), nil
	}
}

// UpdateInstanceConfigurationInvoker 修改实例配置
func (c *SwrClient) UpdateInstanceConfigurationInvoker(request *model.UpdateInstanceConfigurationRequest) *UpdateInstanceConfigurationInvoker {
	requestDef := GenReqDefForUpdateInstanceConfiguration()
	return &UpdateInstanceConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceEndpointPolicy 更新公网访问白名单
//
// 更新公网访问白名单，更新为全量更新方式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) UpdateInstanceEndpointPolicy(request *model.UpdateInstanceEndpointPolicyRequest) (*model.UpdateInstanceEndpointPolicyResponse, error) {
	requestDef := GenReqDefForUpdateInstanceEndpointPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceEndpointPolicyResponse), nil
	}
}

// UpdateInstanceEndpointPolicyInvoker 更新公网访问白名单
func (c *SwrClient) UpdateInstanceEndpointPolicyInvoker(request *model.UpdateInstanceEndpointPolicyRequest) *UpdateInstanceEndpointPolicyInvoker {
	requestDef := GenReqDefForUpdateInstanceEndpointPolicy()
	return &UpdateInstanceEndpointPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceLtCredential 启用/停用长期访问凭证
//
// 启用/停用长期访问凭证
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) UpdateInstanceLtCredential(request *model.UpdateInstanceLtCredentialRequest) (*model.UpdateInstanceLtCredentialResponse, error) {
	requestDef := GenReqDefForUpdateInstanceLtCredential()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceLtCredentialResponse), nil
	}
}

// UpdateInstanceLtCredentialInvoker 启用/停用长期访问凭证
func (c *SwrClient) UpdateInstanceLtCredentialInvoker(request *model.UpdateInstanceLtCredentialRequest) *UpdateInstanceLtCredentialInvoker {
	requestDef := GenReqDefForUpdateInstanceLtCredential()
	return &UpdateInstanceLtCredentialInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceNamespace 修改命名空间
//
// 修改命名空间
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) UpdateInstanceNamespace(request *model.UpdateInstanceNamespaceRequest) (*model.UpdateInstanceNamespaceResponse, error) {
	requestDef := GenReqDefForUpdateInstanceNamespace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceNamespaceResponse), nil
	}
}

// UpdateInstanceNamespaceInvoker 修改命名空间
func (c *SwrClient) UpdateInstanceNamespaceInvoker(request *model.UpdateInstanceNamespaceRequest) *UpdateInstanceNamespaceInvoker {
	requestDef := GenReqDefForUpdateInstanceNamespace()
	return &UpdateInstanceNamespaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceRegistry 修改镜像同步的目标仓库
//
// 修改镜像同步的目标仓库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) UpdateInstanceRegistry(request *model.UpdateInstanceRegistryRequest) (*model.UpdateInstanceRegistryResponse, error) {
	requestDef := GenReqDefForUpdateInstanceRegistry()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceRegistryResponse), nil
	}
}

// UpdateInstanceRegistryInvoker 修改镜像同步的目标仓库
func (c *SwrClient) UpdateInstanceRegistryInvoker(request *model.UpdateInstanceRegistryRequest) *UpdateInstanceRegistryInvoker {
	requestDef := GenReqDefForUpdateInstanceRegistry()
	return &UpdateInstanceRegistryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceReplicationPolicy 修改镜像同步策略
//
// 修改镜像同步策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) UpdateInstanceReplicationPolicy(request *model.UpdateInstanceReplicationPolicyRequest) (*model.UpdateInstanceReplicationPolicyResponse, error) {
	requestDef := GenReqDefForUpdateInstanceReplicationPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceReplicationPolicyResponse), nil
	}
}

// UpdateInstanceReplicationPolicyInvoker 修改镜像同步策略
func (c *SwrClient) UpdateInstanceReplicationPolicyInvoker(request *model.UpdateInstanceReplicationPolicyRequest) *UpdateInstanceReplicationPolicyInvoker {
	requestDef := GenReqDefForUpdateInstanceReplicationPolicy()
	return &UpdateInstanceReplicationPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceRepository 修改制品仓库
//
// 修改制品仓库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) UpdateInstanceRepository(request *model.UpdateInstanceRepositoryRequest) (*model.UpdateInstanceRepositoryResponse, error) {
	requestDef := GenReqDefForUpdateInstanceRepository()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceRepositoryResponse), nil
	}
}

// UpdateInstanceRepositoryInvoker 修改制品仓库
func (c *SwrClient) UpdateInstanceRepositoryInvoker(request *model.UpdateInstanceRepositoryRequest) *UpdateInstanceRepositoryInvoker {
	requestDef := GenReqDefForUpdateInstanceRepository()
	return &UpdateInstanceRepositoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceRetentionPolicy 修改老化策略
//
// 修改老化策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) UpdateInstanceRetentionPolicy(request *model.UpdateInstanceRetentionPolicyRequest) (*model.UpdateInstanceRetentionPolicyResponse, error) {
	requestDef := GenReqDefForUpdateInstanceRetentionPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceRetentionPolicyResponse), nil
	}
}

// UpdateInstanceRetentionPolicyInvoker 修改老化策略
func (c *SwrClient) UpdateInstanceRetentionPolicyInvoker(request *model.UpdateInstanceRetentionPolicyRequest) *UpdateInstanceRetentionPolicyInvoker {
	requestDef := GenReqDefForUpdateInstanceRetentionPolicy()
	return &UpdateInstanceRetentionPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceSignPolicy 修改签名策略
//
// 修改签名策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) UpdateInstanceSignPolicy(request *model.UpdateInstanceSignPolicyRequest) (*model.UpdateInstanceSignPolicyResponse, error) {
	requestDef := GenReqDefForUpdateInstanceSignPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceSignPolicyResponse), nil
	}
}

// UpdateInstanceSignPolicyInvoker 修改签名策略
func (c *SwrClient) UpdateInstanceSignPolicyInvoker(request *model.UpdateInstanceSignPolicyRequest) *UpdateInstanceSignPolicyInvoker {
	requestDef := GenReqDefForUpdateInstanceSignPolicy()
	return &UpdateInstanceSignPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceWebhook 修改触发器策略
//
// 修改触发器策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SwrClient) UpdateInstanceWebhook(request *model.UpdateInstanceWebhookRequest) (*model.UpdateInstanceWebhookResponse, error) {
	requestDef := GenReqDefForUpdateInstanceWebhook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceWebhookResponse), nil
	}
}

// UpdateInstanceWebhookInvoker 修改触发器策略
func (c *SwrClient) UpdateInstanceWebhookInvoker(request *model.UpdateInstanceWebhookRequest) *UpdateInstanceWebhookInvoker {
	requestDef := GenReqDefForUpdateInstanceWebhook()
	return &UpdateInstanceWebhookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
