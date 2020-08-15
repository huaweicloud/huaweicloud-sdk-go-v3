package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbr/v1/model"
)

type CbrClient struct {
	hcClient *http_client.HcHttpClient
}

func NewCbrClient(hcClient *http_client.HcHttpClient) *CbrClient {
	return &CbrClient{hcClient: hcClient}
}

func CbrClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//添加备份可共享的成员，只有云服务器备份可以添加备份共享成员，且仅支持在同一区域的不同用户间共享。
func (c *CbrClient) AddMember(request *model.AddMemberRequest) (*model.AddMemberResponse, error) {
	requestDef := GenReqDefForAddMember(request)
	resp, responseDef := GenRespForAddMember()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//存储库添加资源
func (c *CbrClient) AddVaultResource(request *model.AddVaultResourceRequest) (*model.AddVaultResourceResponse, error) {
	requestDef := GenReqDefForAddVaultResource(request)
	resp, responseDef := GenRespForAddVaultResource()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//存储库设置策略
func (c *CbrClient) AssociateVaultPolicy(request *model.AssociateVaultPolicyRequest) (*model.AssociateVaultPolicyResponse, error) {
	requestDef := GenReqDefForAssociateVaultPolicy(request)
	resp, responseDef := GenRespForAssociateVaultPolicy()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//跨区域复制备份。
func (c *CbrClient) CopyBackup(request *model.CopyBackupRequest) (*model.CopyBackupResponse, error) {
	requestDef := GenReqDefForCopyBackup(request)
	resp, responseDef := GenRespForCopyBackup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//执行复制
func (c *CbrClient) CopyCheckpoint(request *model.CopyCheckpointRequest) (*model.CopyCheckpointResponse, error) {
	requestDef := GenReqDefForCopyCheckpoint(request)
	resp, responseDef := GenRespForCopyCheckpoint()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//对存储库执行备份，生成备份还原点
func (c *CbrClient) CreateCheckpoint(request *model.CreateCheckpointRequest) (*model.CreateCheckpointResponse, error) {
	requestDef := GenReqDefForCreateCheckpoint(request)
	resp, responseDef := GenRespForCreateCheckpoint()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//[创建策略，策略分为备份策略和复制策略。](tag:hws,hws_hk) [创建备份策略。](tag:dt,ocb,tlf,sbc,fcs_vm,ctc)
func (c *CbrClient) CreatePolicy(request *model.CreatePolicyRequest) (*model.CreatePolicyResponse, error) {
	requestDef := GenReqDefForCreatePolicy(request)
	resp, responseDef := GenRespForCreatePolicy()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//创建存储库
func (c *CbrClient) CreateVault(request *model.CreateVaultRequest) (*model.CreateVaultResponse, error) {
	requestDef := GenReqDefForCreateVault(request)
	resp, responseDef := GenRespForCreateVault()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除单个备份。
func (c *CbrClient) DeleteBackup(request *model.DeleteBackupRequest) (*model.DeleteBackupResponse, error) {
	requestDef := GenReqDefForDeleteBackup(request)
	resp, responseDef := GenRespForDeleteBackup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除指定的备份共享成员
func (c *CbrClient) DeleteMember(request *model.DeleteMemberRequest) (*model.DeleteMemberResponse, error) {
	requestDef := GenReqDefForDeleteMember(request)
	resp, responseDef := GenRespForDeleteMember()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除策略
func (c *CbrClient) DeletePolicy(request *model.DeletePolicyRequest) (*model.DeletePolicyResponse, error) {
	requestDef := GenReqDefForDeletePolicy(request)
	resp, responseDef := GenRespForDeletePolicy()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除存储库。若删除储存库，将一并删除存储库中的所有备份。
func (c *CbrClient) DeleteVault(request *model.DeleteVaultRequest) (*model.DeleteVaultResponse, error) {
	requestDef := GenReqDefForDeleteVault(request)
	resp, responseDef := GenRespForDeleteVault()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//存储库解除策略
func (c *CbrClient) DisassociateVaultPolicy(request *model.DisassociateVaultPolicyRequest) (*model.DisassociateVaultPolicyResponse, error) {
	requestDef := GenReqDefForDisassociateVaultPolicy(request)
	resp, responseDef := GenRespForDisassociateVaultPolicy()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//同步备份副本
func (c *CbrClient) ImportBackup(request *model.ImportBackupRequest) (*model.ImportBackupResponse, error) {
	requestDef := GenReqDefForImportBackup(request)
	resp, responseDef := GenRespForImportBackup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询所有副本
func (c *CbrClient) ListBackups(request *model.ListBackupsRequest) (*model.ListBackupsResponse, error) {
	requestDef := GenReqDefForListBackups(request)
	resp, responseDef := GenRespForListBackups()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询任务列表
func (c *CbrClient) ListOpLogs(request *model.ListOpLogsRequest) (*model.ListOpLogsResponse, error) {
	requestDef := GenReqDefForListOpLogs(request)
	resp, responseDef := GenRespForListOpLogs()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询策略列表
func (c *CbrClient) ListPolicies(request *model.ListPoliciesRequest) (*model.ListPoliciesResponse, error) {
	requestDef := GenReqDefForListPolicies(request)
	resp, responseDef := GenRespForListPolicies()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询可保护性资源列表
func (c *CbrClient) ListProtectable(request *model.ListProtectableRequest) (*model.ListProtectableResponse, error) {
	requestDef := GenReqDefForListProtectable(request)
	resp, responseDef := GenRespForListProtectable()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询存储库列表
func (c *CbrClient) ListVault(request *model.ListVaultRequest) (*model.ListVaultResponse, error) {
	requestDef := GenReqDefForListVault(request)
	resp, responseDef := GenRespForListVault()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//移除存储库中的资源，若移除资源，将一并删除该资源在保管库中的备份
func (c *CbrClient) RemoveVaultResource(request *model.RemoveVaultResourceRequest) (*model.RemoveVaultResourceResponse, error) {
	requestDef := GenReqDefForRemoveVaultResource(request)
	resp, responseDef := GenRespForRemoveVaultResource()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//恢复备份数据
func (c *CbrClient) RestoreBackup(request *model.RestoreBackupRequest) (*model.RestoreBackupResponse, error) {
	requestDef := GenReqDefForRestoreBackup(request)
	resp, responseDef := GenRespForRestoreBackup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//根据指定id查询单个副本。
func (c *CbrClient) ShowBackup(request *model.ShowBackupRequest) (*model.ShowBackupResponse, error) {
	requestDef := GenReqDefForShowBackup(request)
	resp, responseDef := GenRespForShowBackup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//根据还原点ID查询指定还原点
func (c *CbrClient) ShowCheckpoint(request *model.ShowCheckpointRequest) (*model.ShowCheckpointResponse, error) {
	requestDef := GenReqDefForShowCheckpoint(request)
	resp, responseDef := GenRespForShowCheckpoint()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//获取备份成员的详情
func (c *CbrClient) ShowMemberDetail(request *model.ShowMemberDetailRequest) (*model.ShowMemberDetailResponse, error) {
	requestDef := GenReqDefForShowMemberDetail(request)
	resp, responseDef := GenRespForShowMemberDetail()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//获取备份共享成员的列表信息
func (c *CbrClient) ShowMembersDetail(request *model.ShowMembersDetailRequest) (*model.ShowMembersDetailResponse, error) {
	requestDef := GenReqDefForShowMembersDetail(request)
	resp, responseDef := GenRespForShowMembersDetail()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//根据指定任务ID查询任务
func (c *CbrClient) ShowOpLog(request *model.ShowOpLogRequest) (*model.ShowOpLogResponse, error) {
	requestDef := GenReqDefForShowOpLog(request)
	resp, responseDef := GenRespForShowOpLog()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询单个策略
func (c *CbrClient) ShowPolicy(request *model.ShowPolicyRequest) (*model.ShowPolicyResponse, error) {
	requestDef := GenReqDefForShowPolicy(request)
	resp, responseDef := GenRespForShowPolicy()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//根据ID查询可保护性资源
func (c *CbrClient) ShowProtectable(request *model.ShowProtectableRequest) (*model.ShowProtectableResponse, error) {
	requestDef := GenReqDefForShowProtectable(request)
	resp, responseDef := GenRespForShowProtectable()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询本区域的复制能力
func (c *CbrClient) ShowReplicationCapabilities(request *model.ShowReplicationCapabilitiesRequest) (*model.ShowReplicationCapabilitiesResponse, error) {
	requestDef := GenReqDefForShowReplicationCapabilities(request)
	resp, responseDef := GenRespForShowReplicationCapabilities()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//根据ID查询指定存储库
func (c *CbrClient) ShowVault(request *model.ShowVaultRequest) (*model.ShowVaultResponse, error) {
	requestDef := GenReqDefForShowVault(request)
	resp, responseDef := GenRespForShowVault()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//更新备份共享成员的状态，需要接收方执行此API。
func (c *CbrClient) UpdateMemberStatus(request *model.UpdateMemberStatusRequest) (*model.UpdateMemberStatusResponse, error) {
	requestDef := GenReqDefForUpdateMemberStatus(request)
	resp, responseDef := GenRespForUpdateMemberStatus()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//修改策略
func (c *CbrClient) UpdatePolicy(request *model.UpdatePolicyRequest) (*model.UpdatePolicyResponse, error) {
	requestDef := GenReqDefForUpdatePolicy(request)
	resp, responseDef := GenRespForUpdatePolicy()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//根据存储库ID修改存储库
func (c *CbrClient) UpdateVault(request *model.UpdateVaultRequest) (*model.UpdateVaultResponse, error) {
	requestDef := GenReqDefForUpdateVault(request)
	resp, responseDef := GenRespForUpdateVault()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}
