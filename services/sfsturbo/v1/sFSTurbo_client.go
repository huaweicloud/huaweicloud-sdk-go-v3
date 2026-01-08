package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/sfsturbo/v1/model"
)

type SFSTurboClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewSFSTurboClient(hcClient *httpclient.HcHttpClient) *SFSTurboClient {
	return &SFSTurboClient{HcClient: hcClient}
}

func SFSTurboClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AddActiveDirectoryDomain 加入AD域
//
// 加入AD域。Active Directory域（简称：AD域）提供统一的身份认证和授权管理。通过将SFS Turbo文件系统的挂载点接入AD域内，您可以在AD域中实现文件系统用户身份的认证管理和文件级别的访问权限控制。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) AddActiveDirectoryDomain(request *model.AddActiveDirectoryDomainRequest) (*model.AddActiveDirectoryDomainResponse, error) {
	requestDef := GenReqDefForAddActiveDirectoryDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddActiveDirectoryDomainResponse), nil
	}
}

// AddActiveDirectoryDomainInvoker 加入AD域
func (c *SFSTurboClient) AddActiveDirectoryDomainInvoker(request *model.AddActiveDirectoryDomainRequest) *AddActiveDirectoryDomainInvoker {
	requestDef := GenReqDefForAddActiveDirectoryDomain()
	return &AddActiveDirectoryDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAddSharedTags 批量添加共享标签
//
// 指定共享批量添加标签。
//
// 一个共享上最多有20个标签。
// 一个共享上的多个标签的key不允许重复。
// 此接口为幂等接口：如果要添加的key在共享上已存在，则覆盖更新标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) BatchAddSharedTags(request *model.BatchAddSharedTagsRequest) (*model.BatchAddSharedTagsResponse, error) {
	requestDef := GenReqDefForBatchAddSharedTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddSharedTagsResponse), nil
	}
}

// BatchAddSharedTagsInvoker 批量添加共享标签
func (c *SFSTurboClient) BatchAddSharedTagsInvoker(request *model.BatchAddSharedTagsRequest) *BatchAddSharedTagsInvoker {
	requestDef := GenReqDefForBatchAddSharedTags()
	return &BatchAddSharedTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeSecurityGroup 修改文件系统绑定的安全组
//
// 修改SFS Turbo文件系统绑定的安全组。修改安全组为异步任务，可以通过“查询单个文件系统”返回的子状态字段“sub_status”来判断是否修改安全组状态，子状态为“232”即为修改安全组成功。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) ChangeSecurityGroup(request *model.ChangeSecurityGroupRequest) (*model.ChangeSecurityGroupResponse, error) {
	requestDef := GenReqDefForChangeSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeSecurityGroupResponse), nil
	}
}

// ChangeSecurityGroupInvoker 修改文件系统绑定的安全组
func (c *SFSTurboClient) ChangeSecurityGroupInvoker(request *model.ChangeSecurityGroupRequest) *ChangeSecurityGroupInvoker {
	requestDef := GenReqDefForChangeSecurityGroup()
	return &ChangeSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeShareChargeModeV2 修改文件系统计费模式由按需转为包周期
//
// 修改文件系统计费模式由按需转为包周期。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) ChangeShareChargeModeV2(request *model.ChangeShareChargeModeV2Request) (*model.ChangeShareChargeModeV2Response, error) {
	requestDef := GenReqDefForChangeShareChargeModeV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeShareChargeModeV2Response), nil
	}
}

// ChangeShareChargeModeV2Invoker 修改文件系统计费模式由按需转为包周期
func (c *SFSTurboClient) ChangeShareChargeModeV2Invoker(request *model.ChangeShareChargeModeV2Request) *ChangeShareChargeModeV2Invoker {
	requestDef := GenReqDefForChangeShareChargeModeV2()
	return &ChangeShareChargeModeV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeShareName 修改文件系统名称
//
// 修改文件系统名称
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) ChangeShareName(request *model.ChangeShareNameRequest) (*model.ChangeShareNameResponse, error) {
	requestDef := GenReqDefForChangeShareName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeShareNameResponse), nil
	}
}

// ChangeShareNameInvoker 修改文件系统名称
func (c *SFSTurboClient) ChangeShareNameInvoker(request *model.ChangeShareNameRequest) *ChangeShareNameInvoker {
	requestDef := GenReqDefForChangeShareName()
	return &ChangeShareNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateBackendTarget 绑定后端存储
//
// 为SFS Turbo 文件系统绑定后端存储
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) CreateBackendTarget(request *model.CreateBackendTargetRequest) (*model.CreateBackendTargetResponse, error) {
	requestDef := GenReqDefForCreateBackendTarget()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBackendTargetResponse), nil
	}
}

// CreateBackendTargetInvoker 绑定后端存储
func (c *SFSTurboClient) CreateBackendTargetInvoker(request *model.CreateBackendTargetRequest) *CreateBackendTargetInvoker {
	requestDef := GenReqDefForCreateBackendTarget()
	return &CreateBackendTargetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFsDir 创建目录
//
// 创建目录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) CreateFsDir(request *model.CreateFsDirRequest) (*model.CreateFsDirResponse, error) {
	requestDef := GenReqDefForCreateFsDir()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFsDirResponse), nil
	}
}

// CreateFsDirInvoker 创建目录
func (c *SFSTurboClient) CreateFsDirInvoker(request *model.CreateFsDirRequest) *CreateFsDirInvoker {
	requestDef := GenReqDefForCreateFsDir()
	return &CreateFsDirInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFsDirQuota 创建目标文件夹配额
//
// 创建目标文件夹配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) CreateFsDirQuota(request *model.CreateFsDirQuotaRequest) (*model.CreateFsDirQuotaResponse, error) {
	requestDef := GenReqDefForCreateFsDirQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFsDirQuotaResponse), nil
	}
}

// CreateFsDirQuotaInvoker 创建目标文件夹配额
func (c *SFSTurboClient) CreateFsDirQuotaInvoker(request *model.CreateFsDirQuotaRequest) *CreateFsDirQuotaInvoker {
	requestDef := GenReqDefForCreateFsDirQuota()
	return &CreateFsDirQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFsTask 创建文件系统异步任务
//
// 创建文件系统异步任务，仅支持异步查询目录资源使用情况，API请求路径的feature取值为dir-usage，以下简称为DU任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) CreateFsTask(request *model.CreateFsTaskRequest) (*model.CreateFsTaskResponse, error) {
	requestDef := GenReqDefForCreateFsTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFsTaskResponse), nil
	}
}

// CreateFsTaskInvoker 创建文件系统异步任务
func (c *SFSTurboClient) CreateFsTaskInvoker(request *model.CreateFsTaskRequest) *CreateFsTaskInvoker {
	requestDef := GenReqDefForCreateFsTask()
	return &CreateFsTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHpcCacheTask 创建数据导入导出任务
//
// 创建数据导入导出任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) CreateHpcCacheTask(request *model.CreateHpcCacheTaskRequest) (*model.CreateHpcCacheTaskResponse, error) {
	requestDef := GenReqDefForCreateHpcCacheTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHpcCacheTaskResponse), nil
	}
}

// CreateHpcCacheTaskInvoker 创建数据导入导出任务
func (c *SFSTurboClient) CreateHpcCacheTaskInvoker(request *model.CreateHpcCacheTaskRequest) *CreateHpcCacheTaskInvoker {
	requestDef := GenReqDefForCreateHpcCacheTask()
	return &CreateHpcCacheTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLdapConfig 创建并绑定LDAP配置
//
// 创建并绑定LDAP配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) CreateLdapConfig(request *model.CreateLdapConfigRequest) (*model.CreateLdapConfigResponse, error) {
	requestDef := GenReqDefForCreateLdapConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLdapConfigResponse), nil
	}
}

// CreateLdapConfigInvoker 创建并绑定LDAP配置
func (c *SFSTurboClient) CreateLdapConfigInvoker(request *model.CreateLdapConfigRequest) *CreateLdapConfigInvoker {
	requestDef := GenReqDefForCreateLdapConfig()
	return &CreateLdapConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePermRule 创建权限规则
//
// 创建权限规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) CreatePermRule(request *model.CreatePermRuleRequest) (*model.CreatePermRuleResponse, error) {
	requestDef := GenReqDefForCreatePermRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePermRuleResponse), nil
	}
}

// CreatePermRuleInvoker 创建权限规则
func (c *SFSTurboClient) CreatePermRuleInvoker(request *model.CreatePermRuleRequest) *CreatePermRuleInvoker {
	requestDef := GenReqDefForCreatePermRule()
	return &CreatePermRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateShare 创建文件系统
//
// 创建文件系统。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) CreateShare(request *model.CreateShareRequest) (*model.CreateShareResponse, error) {
	requestDef := GenReqDefForCreateShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateShareResponse), nil
	}
}

// CreateShareInvoker 创建文件系统
func (c *SFSTurboClient) CreateShareInvoker(request *model.CreateShareRequest) *CreateShareInvoker {
	requestDef := GenReqDefForCreateShare()
	return &CreateShareInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSharedTag 创建共享标签
//
// 指定共享添加一个标签。
// 一个共享上最多有20个标签。
// 一个共享上的多个标签的key不允许重复。
// 此接口为幂等接口：如果要添加的key在共享上已存在，则覆盖更新标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) CreateSharedTag(request *model.CreateSharedTagRequest) (*model.CreateSharedTagResponse, error) {
	requestDef := GenReqDefForCreateSharedTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSharedTagResponse), nil
	}
}

// CreateSharedTagInvoker 创建共享标签
func (c *SFSTurboClient) CreateSharedTagInvoker(request *model.CreateSharedTagRequest) *CreateSharedTagInvoker {
	requestDef := GenReqDefForCreateSharedTag()
	return &CreateSharedTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteActiveDirectoryDomain 退出AD域
//
// 退出AD域。Active Directory域（简称：AD域）提供统一的身份认证和授权管理。通过将SFS Turbo文件系统的挂载点接入AD域内，您可以在AD域中实现文件系统用户身份的认证管理和文件级别的访问权限控制。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) DeleteActiveDirectoryDomain(request *model.DeleteActiveDirectoryDomainRequest) (*model.DeleteActiveDirectoryDomainResponse, error) {
	requestDef := GenReqDefForDeleteActiveDirectoryDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteActiveDirectoryDomainResponse), nil
	}
}

// DeleteActiveDirectoryDomainInvoker 退出AD域
func (c *SFSTurboClient) DeleteActiveDirectoryDomainInvoker(request *model.DeleteActiveDirectoryDomainRequest) *DeleteActiveDirectoryDomainInvoker {
	requestDef := GenReqDefForDeleteActiveDirectoryDomain()
	return &DeleteActiveDirectoryDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBackendTarget 删除后端存储
//
// 删除后端存储
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) DeleteBackendTarget(request *model.DeleteBackendTargetRequest) (*model.DeleteBackendTargetResponse, error) {
	requestDef := GenReqDefForDeleteBackendTarget()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBackendTargetResponse), nil
	}
}

// DeleteBackendTargetInvoker 删除后端存储
func (c *SFSTurboClient) DeleteBackendTargetInvoker(request *model.DeleteBackendTargetRequest) *DeleteBackendTargetInvoker {
	requestDef := GenReqDefForDeleteBackendTarget()
	return &DeleteBackendTargetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFsDir 删除文件系统目录
//
// 删除文件系统目录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) DeleteFsDir(request *model.DeleteFsDirRequest) (*model.DeleteFsDirResponse, error) {
	requestDef := GenReqDefForDeleteFsDir()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFsDirResponse), nil
	}
}

// DeleteFsDirInvoker 删除文件系统目录
func (c *SFSTurboClient) DeleteFsDirInvoker(request *model.DeleteFsDirRequest) *DeleteFsDirInvoker {
	requestDef := GenReqDefForDeleteFsDir()
	return &DeleteFsDirInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFsDirQuota 删除目标文件夹配额
//
// 删除目标文件夹配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) DeleteFsDirQuota(request *model.DeleteFsDirQuotaRequest) (*model.DeleteFsDirQuotaResponse, error) {
	requestDef := GenReqDefForDeleteFsDirQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFsDirQuotaResponse), nil
	}
}

// DeleteFsDirQuotaInvoker 删除目标文件夹配额
func (c *SFSTurboClient) DeleteFsDirQuotaInvoker(request *model.DeleteFsDirQuotaRequest) *DeleteFsDirQuotaInvoker {
	requestDef := GenReqDefForDeleteFsDirQuota()
	return &DeleteFsDirQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFsTask 取消/删除文件系统异步任务
//
// 如果异步任务正在执行，则取消并删除任务；否则，删除任务。仅支持删除目录资源使用情况的任务，API请求路径的feature取值为dir-usage，以下简称为DU任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) DeleteFsTask(request *model.DeleteFsTaskRequest) (*model.DeleteFsTaskResponse, error) {
	requestDef := GenReqDefForDeleteFsTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFsTaskResponse), nil
	}
}

// DeleteFsTaskInvoker 取消/删除文件系统异步任务
func (c *SFSTurboClient) DeleteFsTaskInvoker(request *model.DeleteFsTaskRequest) *DeleteFsTaskInvoker {
	requestDef := GenReqDefForDeleteFsTask()
	return &DeleteFsTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHpcCacheTask 删除数据导入导出任务
//
// 删除数据导入导出任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) DeleteHpcCacheTask(request *model.DeleteHpcCacheTaskRequest) (*model.DeleteHpcCacheTaskResponse, error) {
	requestDef := GenReqDefForDeleteHpcCacheTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHpcCacheTaskResponse), nil
	}
}

// DeleteHpcCacheTaskInvoker 删除数据导入导出任务
func (c *SFSTurboClient) DeleteHpcCacheTaskInvoker(request *model.DeleteHpcCacheTaskRequest) *DeleteHpcCacheTaskInvoker {
	requestDef := GenReqDefForDeleteHpcCacheTask()
	return &DeleteHpcCacheTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLdapConfig 删除LDAP配置
//
// 删除LDAP配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) DeleteLdapConfig(request *model.DeleteLdapConfigRequest) (*model.DeleteLdapConfigResponse, error) {
	requestDef := GenReqDefForDeleteLdapConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLdapConfigResponse), nil
	}
}

// DeleteLdapConfigInvoker 删除LDAP配置
func (c *SFSTurboClient) DeleteLdapConfigInvoker(request *model.DeleteLdapConfigRequest) *DeleteLdapConfigInvoker {
	requestDef := GenReqDefForDeleteLdapConfig()
	return &DeleteLdapConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePermRule 删除权限规则
//
// 删除权限规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) DeletePermRule(request *model.DeletePermRuleRequest) (*model.DeletePermRuleResponse, error) {
	requestDef := GenReqDefForDeletePermRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePermRuleResponse), nil
	}
}

// DeletePermRuleInvoker 删除权限规则
func (c *SFSTurboClient) DeletePermRuleInvoker(request *model.DeletePermRuleRequest) *DeletePermRuleInvoker {
	requestDef := GenReqDefForDeletePermRule()
	return &DeletePermRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteShare 删除文件系统
//
// 删除文件系统。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) DeleteShare(request *model.DeleteShareRequest) (*model.DeleteShareResponse, error) {
	requestDef := GenReqDefForDeleteShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteShareResponse), nil
	}
}

// DeleteShareInvoker 删除文件系统
func (c *SFSTurboClient) DeleteShareInvoker(request *model.DeleteShareRequest) *DeleteShareInvoker {
	requestDef := GenReqDefForDeleteShare()
	return &DeleteShareInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSharedTag 删除共享标签
//
// 指定共享删除一个标签。当共享中不存在指定要删除的key时，接口调用将会返回404错误。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) DeleteSharedTag(request *model.DeleteSharedTagRequest) (*model.DeleteSharedTagResponse, error) {
	requestDef := GenReqDefForDeleteSharedTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSharedTagResponse), nil
	}
}

// DeleteSharedTagInvoker 删除共享标签
func (c *SFSTurboClient) DeleteSharedTagInvoker(request *model.DeleteSharedTagRequest) *DeleteSharedTagInvoker {
	requestDef := GenReqDefForDeleteSharedTag()
	return &DeleteSharedTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExpandShare 扩容文件系统
//
// 扩容文件系统。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) ExpandShare(request *model.ExpandShareRequest) (*model.ExpandShareResponse, error) {
	requestDef := GenReqDefForExpandShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandShareResponse), nil
	}
}

// ExpandShareInvoker 扩容文件系统
func (c *SFSTurboClient) ExpandShareInvoker(request *model.ExpandShareRequest) *ExpandShareInvoker {
	requestDef := GenReqDefForExpandShare()
	return &ExpandShareInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBackendTargets 查询后端存储列表
//
// 查询后端存储列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) ListBackendTargets(request *model.ListBackendTargetsRequest) (*model.ListBackendTargetsResponse, error) {
	requestDef := GenReqDefForListBackendTargets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackendTargetsResponse), nil
	}
}

// ListBackendTargetsInvoker 查询后端存储列表
func (c *SFSTurboClient) ListBackendTargetsInvoker(request *model.ListBackendTargetsRequest) *ListBackendTargetsInvoker {
	requestDef := GenReqDefForListBackendTargets()
	return &ListBackendTargetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFsTasks 获取文件系统异步任务列表
//
// 获取文件系统异步任务列表。仅支持查询目录资源使用情况的任务，API请求路径的feature取值为dir-usage，以下简称为DU任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) ListFsTasks(request *model.ListFsTasksRequest) (*model.ListFsTasksResponse, error) {
	requestDef := GenReqDefForListFsTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFsTasksResponse), nil
	}
}

// ListFsTasksInvoker 获取文件系统异步任务列表
func (c *SFSTurboClient) ListFsTasksInvoker(request *model.ListFsTasksRequest) *ListFsTasksInvoker {
	requestDef := GenReqDefForListFsTasks()
	return &ListFsTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHpcCacheTasks 查询数据导入导出任务列表
//
// 查询数据导入导出任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) ListHpcCacheTasks(request *model.ListHpcCacheTasksRequest) (*model.ListHpcCacheTasksResponse, error) {
	requestDef := GenReqDefForListHpcCacheTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHpcCacheTasksResponse), nil
	}
}

// ListHpcCacheTasksInvoker 查询数据导入导出任务列表
func (c *SFSTurboClient) ListHpcCacheTasksInvoker(request *model.ListHpcCacheTasksRequest) *ListHpcCacheTasksInvoker {
	requestDef := GenReqDefForListHpcCacheTasks()
	return &ListHpcCacheTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPermRules 查询文件系统的权限规则列表
//
// 查询文件系统的权限规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) ListPermRules(request *model.ListPermRulesRequest) (*model.ListPermRulesResponse, error) {
	requestDef := GenReqDefForListPermRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPermRulesResponse), nil
	}
}

// ListPermRulesInvoker 查询文件系统的权限规则列表
func (c *SFSTurboClient) ListPermRulesInvoker(request *model.ListPermRulesRequest) *ListPermRulesInvoker {
	requestDef := GenReqDefForListPermRules()
	return &ListPermRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListShareTypes 查询文件系统类型和配额
//
// 查询文件系统类型和配额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) ListShareTypes(request *model.ListShareTypesRequest) (*model.ListShareTypesResponse, error) {
	requestDef := GenReqDefForListShareTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListShareTypesResponse), nil
	}
}

// ListShareTypesInvoker 查询文件系统类型和配额
func (c *SFSTurboClient) ListShareTypesInvoker(request *model.ListShareTypesRequest) *ListShareTypesInvoker {
	requestDef := GenReqDefForListShareTypes()
	return &ListShareTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSharedTags 查询租户所有共享的标签
//
// 查询租户所有共享的标签集合。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) ListSharedTags(request *model.ListSharedTagsRequest) (*model.ListSharedTagsResponse, error) {
	requestDef := GenReqDefForListSharedTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSharedTagsResponse), nil
	}
}

// ListSharedTagsInvoker 查询租户所有共享的标签
func (c *SFSTurboClient) ListSharedTagsInvoker(request *model.ListSharedTagsRequest) *ListSharedTagsInvoker {
	requestDef := GenReqDefForListSharedTags()
	return &ListSharedTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListShares 获取文件系统列表
//
// 获取文件系统列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) ListShares(request *model.ListSharesRequest) (*model.ListSharesResponse, error) {
	requestDef := GenReqDefForListShares()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSharesResponse), nil
	}
}

// ListSharesInvoker 获取文件系统列表
func (c *SFSTurboClient) ListSharesInvoker(request *model.ListSharesRequest) *ListSharesInvoker {
	requestDef := GenReqDefForListShares()
	return &ListSharesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSharesByTag 通过标签查询文件系统列表
//
// 通过标签查询文件系统列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) ListSharesByTag(request *model.ListSharesByTagRequest) (*model.ListSharesByTagResponse, error) {
	requestDef := GenReqDefForListSharesByTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSharesByTagResponse), nil
	}
}

// ListSharesByTagInvoker 通过标签查询文件系统列表
func (c *SFSTurboClient) ListSharesByTagInvoker(request *model.ListSharesByTagRequest) *ListSharesByTagInvoker {
	requestDef := GenReqDefForListSharesByTag()
	return &ListSharesByTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetHpcCacheBackend 配置hpc缓存型后端信息
//
// 配置hpc缓存型后端信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) SetHpcCacheBackend(request *model.SetHpcCacheBackendRequest) (*model.SetHpcCacheBackendResponse, error) {
	requestDef := GenReqDefForSetHpcCacheBackend()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetHpcCacheBackendResponse), nil
	}
}

// SetHpcCacheBackendInvoker 配置hpc缓存型后端信息
func (c *SFSTurboClient) SetHpcCacheBackendInvoker(request *model.SetHpcCacheBackendRequest) *SetHpcCacheBackendInvoker {
	requestDef := GenReqDefForSetHpcCacheBackend()
	return &SetHpcCacheBackendInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowActiveDirectoryDomain 查询AD域配置
//
// 查询AD域配置。Active Directory域（简称：AD域）提供统一的身份认证和授权管理。通过将SFS Turbo文件系统的挂载点接入AD域内，您可以在AD域中实现文件系统用户身份的认证管理和文件级别的访问权限控制。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) ShowActiveDirectoryDomain(request *model.ShowActiveDirectoryDomainRequest) (*model.ShowActiveDirectoryDomainResponse, error) {
	requestDef := GenReqDefForShowActiveDirectoryDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowActiveDirectoryDomainResponse), nil
	}
}

// ShowActiveDirectoryDomainInvoker 查询AD域配置
func (c *SFSTurboClient) ShowActiveDirectoryDomainInvoker(request *model.ShowActiveDirectoryDomainRequest) *ShowActiveDirectoryDomainInvoker {
	requestDef := GenReqDefForShowActiveDirectoryDomain()
	return &ShowActiveDirectoryDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBackendTargetInfo 获取后端存储详细信息
//
// 获取后端存储详细信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) ShowBackendTargetInfo(request *model.ShowBackendTargetInfoRequest) (*model.ShowBackendTargetInfoResponse, error) {
	requestDef := GenReqDefForShowBackendTargetInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackendTargetInfoResponse), nil
	}
}

// ShowBackendTargetInfoInvoker 获取后端存储详细信息
func (c *SFSTurboClient) ShowBackendTargetInfoInvoker(request *model.ShowBackendTargetInfoRequest) *ShowBackendTargetInfoInvoker {
	requestDef := GenReqDefForShowBackendTargetInfo()
	return &ShowBackendTargetInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowClientIpInfo 获取已挂载的客户端ip信息
//
// 获取已挂载的客户端ip信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) ShowClientIpInfo(request *model.ShowClientIpInfoRequest) (*model.ShowClientIpInfoResponse, error) {
	requestDef := GenReqDefForShowClientIpInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClientIpInfoResponse), nil
	}
}

// ShowClientIpInfoInvoker 获取已挂载的客户端ip信息
func (c *SFSTurboClient) ShowClientIpInfoInvoker(request *model.ShowClientIpInfoRequest) *ShowClientIpInfoInvoker {
	requestDef := GenReqDefForShowClientIpInfo()
	return &ShowClientIpInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFsDir 查询目录是否存在
//
// 查询目录是否存在
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) ShowFsDir(request *model.ShowFsDirRequest) (*model.ShowFsDirResponse, error) {
	requestDef := GenReqDefForShowFsDir()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFsDirResponse), nil
	}
}

// ShowFsDirInvoker 查询目录是否存在
func (c *SFSTurboClient) ShowFsDirInvoker(request *model.ShowFsDirRequest) *ShowFsDirInvoker {
	requestDef := GenReqDefForShowFsDir()
	return &ShowFsDirInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFsDirQuota 查询目标文件夹配额
//
// 查询目标文件夹配额。查询的used_capacity、used_inode数据可能有延迟。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) ShowFsDirQuota(request *model.ShowFsDirQuotaRequest) (*model.ShowFsDirQuotaResponse, error) {
	requestDef := GenReqDefForShowFsDirQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFsDirQuotaResponse), nil
	}
}

// ShowFsDirQuotaInvoker 查询目标文件夹配额
func (c *SFSTurboClient) ShowFsDirQuotaInvoker(request *model.ShowFsDirQuotaRequest) *ShowFsDirQuotaInvoker {
	requestDef := GenReqDefForShowFsDirQuota()
	return &ShowFsDirQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFsDirUsage 查询目录资源使用情况
//
// 查询目录资源使用情况(包括子目录的资源)。后端有5min的缓存时间，查询的数据可能有延迟。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) ShowFsDirUsage(request *model.ShowFsDirUsageRequest) (*model.ShowFsDirUsageResponse, error) {
	requestDef := GenReqDefForShowFsDirUsage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFsDirUsageResponse), nil
	}
}

// ShowFsDirUsageInvoker 查询目录资源使用情况
func (c *SFSTurboClient) ShowFsDirUsageInvoker(request *model.ShowFsDirUsageRequest) *ShowFsDirUsageInvoker {
	requestDef := GenReqDefForShowFsDirUsage()
	return &ShowFsDirUsageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFsTask 获取文件系统异步任务详情
//
// 获取文件系统异步任务详情。仅支持查询目录资源使用情况的任务，API请求路径的feature取值为dir-usage，以下简称为DU任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) ShowFsTask(request *model.ShowFsTaskRequest) (*model.ShowFsTaskResponse, error) {
	requestDef := GenReqDefForShowFsTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFsTaskResponse), nil
	}
}

// ShowFsTaskInvoker 获取文件系统异步任务详情
func (c *SFSTurboClient) ShowFsTaskInvoker(request *model.ShowFsTaskRequest) *ShowFsTaskInvoker {
	requestDef := GenReqDefForShowFsTask()
	return &ShowFsTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHpcCacheTask 查询数据导入导出任务详情
//
// 查询数据导入导出任务详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) ShowHpcCacheTask(request *model.ShowHpcCacheTaskRequest) (*model.ShowHpcCacheTaskResponse, error) {
	requestDef := GenReqDefForShowHpcCacheTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHpcCacheTaskResponse), nil
	}
}

// ShowHpcCacheTaskInvoker 查询数据导入导出任务详情
func (c *SFSTurboClient) ShowHpcCacheTaskInvoker(request *model.ShowHpcCacheTaskRequest) *ShowHpcCacheTaskInvoker {
	requestDef := GenReqDefForShowHpcCacheTask()
	return &ShowHpcCacheTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobDetail 查询job的状态详情
//
// 用于查询SFS Turbo异步API的执行状态。例如：可使用调用创建并绑定LDAP配置接口时返回的jobId，通过该接口查询job的执行状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) ShowJobDetail(request *model.ShowJobDetailRequest) (*model.ShowJobDetailResponse, error) {
	requestDef := GenReqDefForShowJobDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobDetailResponse), nil
	}
}

// ShowJobDetailInvoker 查询job的状态详情
func (c *SFSTurboClient) ShowJobDetailInvoker(request *model.ShowJobDetailRequest) *ShowJobDetailInvoker {
	requestDef := GenReqDefForShowJobDetail()
	return &ShowJobDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLdapConfig 查询LDAP的配置
//
// 查询LDAP的配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) ShowLdapConfig(request *model.ShowLdapConfigRequest) (*model.ShowLdapConfigResponse, error) {
	requestDef := GenReqDefForShowLdapConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLdapConfigResponse), nil
	}
}

// ShowLdapConfigInvoker 查询LDAP的配置
func (c *SFSTurboClient) ShowLdapConfigInvoker(request *model.ShowLdapConfigRequest) *ShowLdapConfigInvoker {
	requestDef := GenReqDefForShowLdapConfig()
	return &ShowLdapConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPermRule 查询文件系统的某一个权限规则
//
// 查询文件系统的某一个权限规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) ShowPermRule(request *model.ShowPermRuleRequest) (*model.ShowPermRuleResponse, error) {
	requestDef := GenReqDefForShowPermRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPermRuleResponse), nil
	}
}

// ShowPermRuleInvoker 查询文件系统的某一个权限规则
func (c *SFSTurboClient) ShowPermRuleInvoker(request *model.ShowPermRuleRequest) *ShowPermRuleInvoker {
	requestDef := GenReqDefForShowPermRule()
	return &ShowPermRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQuota 查询租户配额
//
// 查询租户配额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) ShowQuota(request *model.ShowQuotaRequest) (*model.ShowQuotaResponse, error) {
	requestDef := GenReqDefForShowQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotaResponse), nil
	}
}

// ShowQuotaInvoker 查询租户配额
func (c *SFSTurboClient) ShowQuotaInvoker(request *model.ShowQuotaRequest) *ShowQuotaInvoker {
	requestDef := GenReqDefForShowQuota()
	return &ShowQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowShare 查询文件系统详细信息
//
// 查询SFS Turbo文件系统详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) ShowShare(request *model.ShowShareRequest) (*model.ShowShareResponse, error) {
	requestDef := GenReqDefForShowShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowShareResponse), nil
	}
}

// ShowShareInvoker 查询文件系统详细信息
func (c *SFSTurboClient) ShowShareInvoker(request *model.ShowShareRequest) *ShowShareInvoker {
	requestDef := GenReqDefForShowShare()
	return &ShowShareInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSharedTags 查询共享标签
//
// 查询指定共享的所有标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) ShowSharedTags(request *model.ShowSharedTagsRequest) (*model.ShowSharedTagsResponse, error) {
	requestDef := GenReqDefForShowSharedTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSharedTagsResponse), nil
	}
}

// ShowSharedTagsInvoker 查询共享标签
func (c *SFSTurboClient) ShowSharedTagsInvoker(request *model.ShowSharedTagsRequest) *ShowSharedTagsInvoker {
	requestDef := GenReqDefForShowSharedTags()
	return &ShowSharedTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateActiveDirectoryDomain 修改AD域配置
//
// 修改AD域配置。Active Directory域（简称：AD域）提供统一的身份认证和授权管理。通过将SFS Turbo文件系统的挂载点接入AD域内，您可以在AD域中实现文件系统用户身份的认证管理和文件级别的访问权限控制。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) UpdateActiveDirectoryDomain(request *model.UpdateActiveDirectoryDomainRequest) (*model.UpdateActiveDirectoryDomainResponse, error) {
	requestDef := GenReqDefForUpdateActiveDirectoryDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateActiveDirectoryDomainResponse), nil
	}
}

// UpdateActiveDirectoryDomainInvoker 修改AD域配置
func (c *SFSTurboClient) UpdateActiveDirectoryDomainInvoker(request *model.UpdateActiveDirectoryDomainRequest) *UpdateActiveDirectoryDomainInvoker {
	requestDef := GenReqDefForUpdateActiveDirectoryDomain()
	return &UpdateActiveDirectoryDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFsDirQuota 更新目标文件夹配额
//
// 更新目标文件夹配额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) UpdateFsDirQuota(request *model.UpdateFsDirQuotaRequest) (*model.UpdateFsDirQuotaResponse, error) {
	requestDef := GenReqDefForUpdateFsDirQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFsDirQuotaResponse), nil
	}
}

// UpdateFsDirQuotaInvoker 更新目标文件夹配额
func (c *SFSTurboClient) UpdateFsDirQuotaInvoker(request *model.UpdateFsDirQuotaRequest) *UpdateFsDirQuotaInvoker {
	requestDef := GenReqDefForUpdateFsDirQuota()
	return &UpdateFsDirQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHpcShare 更新文件系统
//
// 更新文件系统冷数据淘汰时间
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) UpdateHpcShare(request *model.UpdateHpcShareRequest) (*model.UpdateHpcShareResponse, error) {
	requestDef := GenReqDefForUpdateHpcShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHpcShareResponse), nil
	}
}

// UpdateHpcShareInvoker 更新文件系统
func (c *SFSTurboClient) UpdateHpcShareInvoker(request *model.UpdateHpcShareRequest) *UpdateHpcShareInvoker {
	requestDef := GenReqDefForUpdateHpcShare()
	return &UpdateHpcShareInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLdapConfig 修改LDAP配置
//
// 修改LDAP配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) UpdateLdapConfig(request *model.UpdateLdapConfigRequest) (*model.UpdateLdapConfigResponse, error) {
	requestDef := GenReqDefForUpdateLdapConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLdapConfigResponse), nil
	}
}

// UpdateLdapConfigInvoker 修改LDAP配置
func (c *SFSTurboClient) UpdateLdapConfigInvoker(request *model.UpdateLdapConfigRequest) *UpdateLdapConfigInvoker {
	requestDef := GenReqDefForUpdateLdapConfig()
	return &UpdateLdapConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateObsTargetAttributes 更新后端存储属性
//
// 更新后端存储属性
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) UpdateObsTargetAttributes(request *model.UpdateObsTargetAttributesRequest) (*model.UpdateObsTargetAttributesResponse, error) {
	requestDef := GenReqDefForUpdateObsTargetAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateObsTargetAttributesResponse), nil
	}
}

// UpdateObsTargetAttributesInvoker 更新后端存储属性
func (c *SFSTurboClient) UpdateObsTargetAttributesInvoker(request *model.UpdateObsTargetAttributesRequest) *UpdateObsTargetAttributesInvoker {
	requestDef := GenReqDefForUpdateObsTargetAttributes()
	return &UpdateObsTargetAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateObsTargetPolicy 更新后端存储自动同步策略
//
// 更新后端存储自动同步策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) UpdateObsTargetPolicy(request *model.UpdateObsTargetPolicyRequest) (*model.UpdateObsTargetPolicyResponse, error) {
	requestDef := GenReqDefForUpdateObsTargetPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateObsTargetPolicyResponse), nil
	}
}

// UpdateObsTargetPolicyInvoker 更新后端存储自动同步策略
func (c *SFSTurboClient) UpdateObsTargetPolicyInvoker(request *model.UpdateObsTargetPolicyRequest) *UpdateObsTargetPolicyInvoker {
	requestDef := GenReqDefForUpdateObsTargetPolicy()
	return &UpdateObsTargetPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePermRule 修改权限规则
//
// 修改权限规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SFSTurboClient) UpdatePermRule(request *model.UpdatePermRuleRequest) (*model.UpdatePermRuleResponse, error) {
	requestDef := GenReqDefForUpdatePermRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePermRuleResponse), nil
	}
}

// UpdatePermRuleInvoker 修改权限规则
func (c *SFSTurboClient) UpdatePermRuleInvoker(request *model.UpdatePermRuleRequest) *UpdatePermRuleInvoker {
	requestDef := GenReqDefForUpdatePermRule()
	return &UpdatePermRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
