package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbr/v1/model"
)

type CbrClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCbrClient(hcClient *http_client.HcHttpClient) *CbrClient {
	return &CbrClient{HcClient: hcClient}
}

func CbrClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// AddMember 添加备份成员
//
// 添加备份可共享的成员，只有云服务器备份可以添加备份共享成员，且仅支持在同一区域的不同用户间共享。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) AddMember(request *model.AddMemberRequest) (*model.AddMemberResponse, error) {
	requestDef := GenReqDefForAddMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddMemberResponse), nil
	}
}

// AddMemberInvoker 添加备份成员
func (c *CbrClient) AddMemberInvoker(request *model.AddMemberRequest) *AddMemberInvoker {
	requestDef := GenReqDefForAddMember()
	return &AddMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddVaultResource 添加资源
//
// 存储库添加资源
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) AddVaultResource(request *model.AddVaultResourceRequest) (*model.AddVaultResourceResponse, error) {
	requestDef := GenReqDefForAddVaultResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddVaultResourceResponse), nil
	}
}

// AddVaultResourceInvoker 添加资源
func (c *CbrClient) AddVaultResourceInvoker(request *model.AddVaultResourceRequest) *AddVaultResourceInvoker {
	requestDef := GenReqDefForAddVaultResource()
	return &AddVaultResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateVaultPolicy 设置存储库策略
//
// 存储库设置策略
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) AssociateVaultPolicy(request *model.AssociateVaultPolicyRequest) (*model.AssociateVaultPolicyResponse, error) {
	requestDef := GenReqDefForAssociateVaultPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateVaultPolicyResponse), nil
	}
}

// AssociateVaultPolicyInvoker 设置存储库策略
func (c *CbrClient) AssociateVaultPolicyInvoker(request *model.AssociateVaultPolicyRequest) *AssociateVaultPolicyInvoker {
	requestDef := GenReqDefForAssociateVaultPolicy()
	return &AssociateVaultPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateAndDeleteVaultTags 批量添加删除存储库资源标签
//
// 为指定实例批量添加或删除标签
// 标签管理服务需要使用该接口批量管理实例的标签。
// 一个资源上最多有10个标签。
// 此接口为幂等接口：
//     创建时如果请求体中存在重复key则报错。
//     创建时，不允许重复key，如果数据库存在就覆盖。
//     删除时，允许重复key。
//     删除时，如果删除的标签不存在，默认处理成功,删除时不对标签字符集范围做校验。key长度127个字符，value为255个字符。删除时tags结构体不能缺失，key不能为空，或者空字符串。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) BatchCreateAndDeleteVaultTags(request *model.BatchCreateAndDeleteVaultTagsRequest) (*model.BatchCreateAndDeleteVaultTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateAndDeleteVaultTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateAndDeleteVaultTagsResponse), nil
	}
}

// BatchCreateAndDeleteVaultTagsInvoker 批量添加删除存储库资源标签
func (c *CbrClient) BatchCreateAndDeleteVaultTagsInvoker(request *model.BatchCreateAndDeleteVaultTagsRequest) *BatchCreateAndDeleteVaultTagsInvoker {
	requestDef := GenReqDefForBatchCreateAndDeleteVaultTags()
	return &BatchCreateAndDeleteVaultTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CopyBackup 复制备份
//
// 跨区域复制备份。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) CopyBackup(request *model.CopyBackupRequest) (*model.CopyBackupResponse, error) {
	requestDef := GenReqDefForCopyBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyBackupResponse), nil
	}
}

// CopyBackupInvoker 复制备份
func (c *CbrClient) CopyBackupInvoker(request *model.CopyBackupRequest) *CopyBackupInvoker {
	requestDef := GenReqDefForCopyBackup()
	return &CopyBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CopyCheckpoint 复制备份还原点
//
// 执行复制
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) CopyCheckpoint(request *model.CopyCheckpointRequest) (*model.CopyCheckpointResponse, error) {
	requestDef := GenReqDefForCopyCheckpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyCheckpointResponse), nil
	}
}

// CopyCheckpointInvoker 复制备份还原点
func (c *CbrClient) CopyCheckpointInvoker(request *model.CopyCheckpointRequest) *CopyCheckpointInvoker {
	requestDef := GenReqDefForCopyCheckpoint()
	return &CopyCheckpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCheckpoint 创建备份还原点
//
// 对存储库执行备份，生成备份还原点
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) CreateCheckpoint(request *model.CreateCheckpointRequest) (*model.CreateCheckpointResponse, error) {
	requestDef := GenReqDefForCreateCheckpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCheckpointResponse), nil
	}
}

// CreateCheckpointInvoker 创建备份还原点
func (c *CbrClient) CreateCheckpointInvoker(request *model.CreateCheckpointRequest) *CreateCheckpointInvoker {
	requestDef := GenReqDefForCreateCheckpoint()
	return &CreateCheckpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePolicy 创建策略
//
// 创建策略，策略分为备份策略和复制策略。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) CreatePolicy(request *model.CreatePolicyRequest) (*model.CreatePolicyResponse, error) {
	requestDef := GenReqDefForCreatePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePolicyResponse), nil
	}
}

// CreatePolicyInvoker 创建策略
func (c *CbrClient) CreatePolicyInvoker(request *model.CreatePolicyRequest) *CreatePolicyInvoker {
	requestDef := GenReqDefForCreatePolicy()
	return &CreatePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVault 创建存储库
//
// 创建存储库
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) CreateVault(request *model.CreateVaultRequest) (*model.CreateVaultResponse, error) {
	requestDef := GenReqDefForCreateVault()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVaultResponse), nil
	}
}

// CreateVaultInvoker 创建存储库
func (c *CbrClient) CreateVaultInvoker(request *model.CreateVaultRequest) *CreateVaultInvoker {
	requestDef := GenReqDefForCreateVault()
	return &CreateVaultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVaultTags 添加存储库资源标签
//
// 一个资源上最多有10个标签。
// 此接口为幂等接口：创建时，如果创建的标签已经存在（key相同），则覆盖。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) CreateVaultTags(request *model.CreateVaultTagsRequest) (*model.CreateVaultTagsResponse, error) {
	requestDef := GenReqDefForCreateVaultTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVaultTagsResponse), nil
	}
}

// CreateVaultTagsInvoker 添加存储库资源标签
func (c *CbrClient) CreateVaultTagsInvoker(request *model.CreateVaultTagsRequest) *CreateVaultTagsInvoker {
	requestDef := GenReqDefForCreateVaultTags()
	return &CreateVaultTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBackup 删除备份
//
// 删除单个备份。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) DeleteBackup(request *model.DeleteBackupRequest) (*model.DeleteBackupResponse, error) {
	requestDef := GenReqDefForDeleteBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBackupResponse), nil
	}
}

// DeleteBackupInvoker 删除备份
func (c *CbrClient) DeleteBackupInvoker(request *model.DeleteBackupRequest) *DeleteBackupInvoker {
	requestDef := GenReqDefForDeleteBackup()
	return &DeleteBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMember 删除指定备份成员
//
// 删除指定的备份共享成员
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) DeleteMember(request *model.DeleteMemberRequest) (*model.DeleteMemberResponse, error) {
	requestDef := GenReqDefForDeleteMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMemberResponse), nil
	}
}

// DeleteMemberInvoker 删除指定备份成员
func (c *CbrClient) DeleteMemberInvoker(request *model.DeleteMemberRequest) *DeleteMemberInvoker {
	requestDef := GenReqDefForDeleteMember()
	return &DeleteMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePolicy 删除策略
//
// 删除策略
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) DeletePolicy(request *model.DeletePolicyRequest) (*model.DeletePolicyResponse, error) {
	requestDef := GenReqDefForDeletePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePolicyResponse), nil
	}
}

// DeletePolicyInvoker 删除策略
func (c *CbrClient) DeletePolicyInvoker(request *model.DeletePolicyRequest) *DeletePolicyInvoker {
	requestDef := GenReqDefForDeletePolicy()
	return &DeletePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVault 删除存储库
//
// 删除存储库。若删除储存库，将一并删除存储库中的所有备份。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) DeleteVault(request *model.DeleteVaultRequest) (*model.DeleteVaultResponse, error) {
	requestDef := GenReqDefForDeleteVault()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVaultResponse), nil
	}
}

// DeleteVaultInvoker 删除存储库
func (c *CbrClient) DeleteVaultInvoker(request *model.DeleteVaultRequest) *DeleteVaultInvoker {
	requestDef := GenReqDefForDeleteVault()
	return &DeleteVaultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVaultTag 删除存储库资源标签
//
// 幂等接口：删除时，如果删除的标签不存在，返回404。Key不能为空或者空字符串。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) DeleteVaultTag(request *model.DeleteVaultTagRequest) (*model.DeleteVaultTagResponse, error) {
	requestDef := GenReqDefForDeleteVaultTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVaultTagResponse), nil
	}
}

// DeleteVaultTagInvoker 删除存储库资源标签
func (c *CbrClient) DeleteVaultTagInvoker(request *model.DeleteVaultTagRequest) *DeleteVaultTagInvoker {
	requestDef := GenReqDefForDeleteVaultTag()
	return &DeleteVaultTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateVaultPolicy 解除存储库策略
//
// 存储库解除策略
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) DisassociateVaultPolicy(request *model.DisassociateVaultPolicyRequest) (*model.DisassociateVaultPolicyResponse, error) {
	requestDef := GenReqDefForDisassociateVaultPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateVaultPolicyResponse), nil
	}
}

// DisassociateVaultPolicyInvoker 解除存储库策略
func (c *CbrClient) DisassociateVaultPolicyInvoker(request *model.DisassociateVaultPolicyRequest) *DisassociateVaultPolicyInvoker {
	requestDef := GenReqDefForDisassociateVaultPolicy()
	return &DisassociateVaultPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportBackup 同步备份
//
// 同步线下混合云VMware备份副本
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) ImportBackup(request *model.ImportBackupRequest) (*model.ImportBackupResponse, error) {
	requestDef := GenReqDefForImportBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportBackupResponse), nil
	}
}

// ImportBackupInvoker 同步备份
func (c *CbrClient) ImportBackupInvoker(request *model.ImportBackupRequest) *ImportBackupInvoker {
	requestDef := GenReqDefForImportBackup()
	return &ImportBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBackups 查询所有备份
//
// 查询所有副本
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) ListBackups(request *model.ListBackupsRequest) (*model.ListBackupsResponse, error) {
	requestDef := GenReqDefForListBackups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackupsResponse), nil
	}
}

// ListBackupsInvoker 查询所有备份
func (c *CbrClient) ListBackupsInvoker(request *model.ListBackupsRequest) *ListBackupsInvoker {
	requestDef := GenReqDefForListBackups()
	return &ListBackupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOpLogs 查询任务列表
//
// 查询任务列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) ListOpLogs(request *model.ListOpLogsRequest) (*model.ListOpLogsResponse, error) {
	requestDef := GenReqDefForListOpLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOpLogsResponse), nil
	}
}

// ListOpLogsInvoker 查询任务列表
func (c *CbrClient) ListOpLogsInvoker(request *model.ListOpLogsRequest) *ListOpLogsInvoker {
	requestDef := GenReqDefForListOpLogs()
	return &ListOpLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicies 查询策略列表
//
// 查询策略列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) ListPolicies(request *model.ListPoliciesRequest) (*model.ListPoliciesResponse, error) {
	requestDef := GenReqDefForListPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPoliciesResponse), nil
	}
}

// ListPoliciesInvoker 查询策略列表
func (c *CbrClient) ListPoliciesInvoker(request *model.ListPoliciesRequest) *ListPoliciesInvoker {
	requestDef := GenReqDefForListPolicies()
	return &ListPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProtectable 查询可保护资源
//
// 查询可保护性资源列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) ListProtectable(request *model.ListProtectableRequest) (*model.ListProtectableResponse, error) {
	requestDef := GenReqDefForListProtectable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProtectableResponse), nil
	}
}

// ListProtectableInvoker 查询可保护资源
func (c *CbrClient) ListProtectableInvoker(request *model.ListProtectableRequest) *ListProtectableInvoker {
	requestDef := GenReqDefForListProtectable()
	return &ListProtectableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVault 查询存储库列表
//
// 查询存储库列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) ListVault(request *model.ListVaultRequest) (*model.ListVaultResponse, error) {
	requestDef := GenReqDefForListVault()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVaultResponse), nil
	}
}

// ListVaultInvoker 查询存储库列表
func (c *CbrClient) ListVaultInvoker(request *model.ListVaultRequest) *ListVaultInvoker {
	requestDef := GenReqDefForListVault()
	return &ListVaultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// MigrateVaultResource 迁移资源
//
// 支持资源迁移到另一个存储库，不删除备份。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) MigrateVaultResource(request *model.MigrateVaultResourceRequest) (*model.MigrateVaultResourceResponse, error) {
	requestDef := GenReqDefForMigrateVaultResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MigrateVaultResourceResponse), nil
	}
}

// MigrateVaultResourceInvoker 迁移资源
func (c *CbrClient) MigrateVaultResourceInvoker(request *model.MigrateVaultResourceRequest) *MigrateVaultResourceInvoker {
	requestDef := GenReqDefForMigrateVaultResource()
	return &MigrateVaultResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveVaultResource 移除资源
//
// 移除存储库中的资源，若移除资源，将一并删除该资源在保管库中的备份
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) RemoveVaultResource(request *model.RemoveVaultResourceRequest) (*model.RemoveVaultResourceResponse, error) {
	requestDef := GenReqDefForRemoveVaultResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveVaultResourceResponse), nil
	}
}

// RemoveVaultResourceInvoker 移除资源
func (c *CbrClient) RemoveVaultResourceInvoker(request *model.RemoveVaultResourceRequest) *RemoveVaultResourceInvoker {
	requestDef := GenReqDefForRemoveVaultResource()
	return &RemoveVaultResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestoreBackup 备份恢复
//
// 恢复备份数据
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) RestoreBackup(request *model.RestoreBackupRequest) (*model.RestoreBackupResponse, error) {
	requestDef := GenReqDefForRestoreBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreBackupResponse), nil
	}
}

// RestoreBackupInvoker 备份恢复
func (c *CbrClient) RestoreBackupInvoker(request *model.RestoreBackupRequest) *RestoreBackupInvoker {
	requestDef := GenReqDefForRestoreBackup()
	return &RestoreBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBackup 查询指定备份
//
// 根据指定id查询单个副本。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) ShowBackup(request *model.ShowBackupRequest) (*model.ShowBackupResponse, error) {
	requestDef := GenReqDefForShowBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackupResponse), nil
	}
}

// ShowBackupInvoker 查询指定备份
func (c *CbrClient) ShowBackupInvoker(request *model.ShowBackupRequest) *ShowBackupInvoker {
	requestDef := GenReqDefForShowBackup()
	return &ShowBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCheckpoint 查询备份还原点
//
// 根据还原点ID查询指定还原点
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) ShowCheckpoint(request *model.ShowCheckpointRequest) (*model.ShowCheckpointResponse, error) {
	requestDef := GenReqDefForShowCheckpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCheckpointResponse), nil
	}
}

// ShowCheckpointInvoker 查询备份还原点
func (c *CbrClient) ShowCheckpointInvoker(request *model.ShowCheckpointRequest) *ShowCheckpointInvoker {
	requestDef := GenReqDefForShowCheckpoint()
	return &ShowCheckpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMemberDetail 获取备份成员详情
//
// 获取备份成员的详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) ShowMemberDetail(request *model.ShowMemberDetailRequest) (*model.ShowMemberDetailResponse, error) {
	requestDef := GenReqDefForShowMemberDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMemberDetailResponse), nil
	}
}

// ShowMemberDetailInvoker 获取备份成员详情
func (c *CbrClient) ShowMemberDetailInvoker(request *model.ShowMemberDetailRequest) *ShowMemberDetailInvoker {
	requestDef := GenReqDefForShowMemberDetail()
	return &ShowMemberDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMembersDetail 获取备份成员列表
//
// 获取备份共享成员的列表信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) ShowMembersDetail(request *model.ShowMembersDetailRequest) (*model.ShowMembersDetailResponse, error) {
	requestDef := GenReqDefForShowMembersDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMembersDetailResponse), nil
	}
}

// ShowMembersDetailInvoker 获取备份成员列表
func (c *CbrClient) ShowMembersDetailInvoker(request *model.ShowMembersDetailRequest) *ShowMembersDetailInvoker {
	requestDef := GenReqDefForShowMembersDetail()
	return &ShowMembersDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOpLog 查询单个任务
//
// 根据指定任务ID查询任务
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) ShowOpLog(request *model.ShowOpLogRequest) (*model.ShowOpLogResponse, error) {
	requestDef := GenReqDefForShowOpLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOpLogResponse), nil
	}
}

// ShowOpLogInvoker 查询单个任务
func (c *CbrClient) ShowOpLogInvoker(request *model.ShowOpLogRequest) *ShowOpLogInvoker {
	requestDef := GenReqDefForShowOpLog()
	return &ShowOpLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPolicy 查询单个策略
//
// 查询单个策略
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) ShowPolicy(request *model.ShowPolicyRequest) (*model.ShowPolicyResponse, error) {
	requestDef := GenReqDefForShowPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPolicyResponse), nil
	}
}

// ShowPolicyInvoker 查询单个策略
func (c *CbrClient) ShowPolicyInvoker(request *model.ShowPolicyRequest) *ShowPolicyInvoker {
	requestDef := GenReqDefForShowPolicy()
	return &ShowPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProtectable 查询指定可保护资源
//
// 根据ID查询可保护性资源
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) ShowProtectable(request *model.ShowProtectableRequest) (*model.ShowProtectableResponse, error) {
	requestDef := GenReqDefForShowProtectable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProtectableResponse), nil
	}
}

// ShowProtectableInvoker 查询指定可保护资源
func (c *CbrClient) ShowProtectableInvoker(request *model.ShowProtectableRequest) *ShowProtectableInvoker {
	requestDef := GenReqDefForShowProtectable()
	return &ShowProtectableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowReplicationCapabilities 查询复制能力
//
// 查询本区域的复制能力
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) ShowReplicationCapabilities(request *model.ShowReplicationCapabilitiesRequest) (*model.ShowReplicationCapabilitiesResponse, error) {
	requestDef := GenReqDefForShowReplicationCapabilities()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReplicationCapabilitiesResponse), nil
	}
}

// ShowReplicationCapabilitiesInvoker 查询复制能力
func (c *CbrClient) ShowReplicationCapabilitiesInvoker(request *model.ShowReplicationCapabilitiesRequest) *ShowReplicationCapabilitiesInvoker {
	requestDef := GenReqDefForShowReplicationCapabilities()
	return &ShowReplicationCapabilitiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVault 查询指定存储库
//
// 根据ID查询指定存储库
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) ShowVault(request *model.ShowVaultRequest) (*model.ShowVaultResponse, error) {
	requestDef := GenReqDefForShowVault()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVaultResponse), nil
	}
}

// ShowVaultInvoker 查询指定存储库
func (c *CbrClient) ShowVaultInvoker(request *model.ShowVaultRequest) *ShowVaultInvoker {
	requestDef := GenReqDefForShowVault()
	return &ShowVaultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVaultProjectTag 查询存储库项目标签
//
// 查询租户在指定Region和实例类型的所有标签集合
// 标签管理服务需要能够列出当前租户全部已使用的标签集合，为各服务Console打标签和过滤实例时提供标签联想功能
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) ShowVaultProjectTag(request *model.ShowVaultProjectTagRequest) (*model.ShowVaultProjectTagResponse, error) {
	requestDef := GenReqDefForShowVaultProjectTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVaultProjectTagResponse), nil
	}
}

// ShowVaultProjectTagInvoker 查询存储库项目标签
func (c *CbrClient) ShowVaultProjectTagInvoker(request *model.ShowVaultProjectTagRequest) *ShowVaultProjectTagInvoker {
	requestDef := GenReqDefForShowVaultProjectTag()
	return &ShowVaultProjectTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVaultResourceInstances 查询存储库资源实例
//
// 使用标签过滤实例
// 标签管理服务需要提供按标签过滤各服务实例并汇总显示在列表中，需要各服务提供查询能力
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) ShowVaultResourceInstances(request *model.ShowVaultResourceInstancesRequest) (*model.ShowVaultResourceInstancesResponse, error) {
	requestDef := GenReqDefForShowVaultResourceInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVaultResourceInstancesResponse), nil
	}
}

// ShowVaultResourceInstancesInvoker 查询存储库资源实例
func (c *CbrClient) ShowVaultResourceInstancesInvoker(request *model.ShowVaultResourceInstancesRequest) *ShowVaultResourceInstancesInvoker {
	requestDef := GenReqDefForShowVaultResourceInstances()
	return &ShowVaultResourceInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVaultTag 查询存储库资源标签
//
// 查询指定实例的标签信息
// 标签管理服务需要使用该接口查询指定实例的全部标签数据
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) ShowVaultTag(request *model.ShowVaultTagRequest) (*model.ShowVaultTagResponse, error) {
	requestDef := GenReqDefForShowVaultTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVaultTagResponse), nil
	}
}

// ShowVaultTagInvoker 查询存储库资源标签
func (c *CbrClient) ShowVaultTagInvoker(request *model.ShowVaultTagRequest) *ShowVaultTagInvoker {
	requestDef := GenReqDefForShowVaultTag()
	return &ShowVaultTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMemberStatus 更新备份成员状态
//
// 更新备份共享成员的状态，需要接收方执行此API。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) UpdateMemberStatus(request *model.UpdateMemberStatusRequest) (*model.UpdateMemberStatusResponse, error) {
	requestDef := GenReqDefForUpdateMemberStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMemberStatusResponse), nil
	}
}

// UpdateMemberStatusInvoker 更新备份成员状态
func (c *CbrClient) UpdateMemberStatusInvoker(request *model.UpdateMemberStatusRequest) *UpdateMemberStatusInvoker {
	requestDef := GenReqDefForUpdateMemberStatus()
	return &UpdateMemberStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePolicy 修改策略
//
// 修改策略
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) UpdatePolicy(request *model.UpdatePolicyRequest) (*model.UpdatePolicyResponse, error) {
	requestDef := GenReqDefForUpdatePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePolicyResponse), nil
	}
}

// UpdatePolicyInvoker 修改策略
func (c *CbrClient) UpdatePolicyInvoker(request *model.UpdatePolicyRequest) *UpdatePolicyInvoker {
	requestDef := GenReqDefForUpdatePolicy()
	return &UpdatePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVault 修改存储库
//
// 根据存储库ID修改存储库
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbrClient) UpdateVault(request *model.UpdateVaultRequest) (*model.UpdateVaultResponse, error) {
	requestDef := GenReqDefForUpdateVault()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVaultResponse), nil
	}
}

// UpdateVaultInvoker 修改存储库
func (c *CbrClient) UpdateVaultInvoker(request *model.UpdateVaultRequest) *UpdateVaultInvoker {
	requestDef := GenReqDefForUpdateVault()
	return &UpdateVaultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
