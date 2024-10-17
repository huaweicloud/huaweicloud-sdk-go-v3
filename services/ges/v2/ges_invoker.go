package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ges/v2/model"
)

type AttachEip2Invoker struct {
	*invoker.BaseInvoker
}

func (i *AttachEip2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AttachEip2Invoker) Invoke() (*model.AttachEip2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachEip2Response), nil
	}
}

type ChangeSecurityGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeSecurityGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeSecurityGroupInvoker) Invoke() (*model.ChangeSecurityGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeSecurityGroupResponse), nil
	}
}

type ClearGraph2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ClearGraph2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ClearGraph2Invoker) Invoke() (*model.ClearGraph2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ClearGraph2Response), nil
	}
}

type CreateBackup2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBackup2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateBackup2Invoker) Invoke() (*model.CreateBackup2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBackup2Response), nil
	}
}

type CreateGraph2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGraph2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateGraph2Invoker) Invoke() (*model.CreateGraph2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGraph2Response), nil
	}
}

type CreateMetadata2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMetadata2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMetadata2Invoker) Invoke() (*model.CreateMetadata2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMetadata2Response), nil
	}
}

type DeleteBackup2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBackup2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteBackup2Invoker) Invoke() (*model.DeleteBackup2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBackup2Response), nil
	}
}

type DeleteGraph2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGraph2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteGraph2Invoker) Invoke() (*model.DeleteGraph2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGraph2Response), nil
	}
}

type DeleteMetadata2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMetadata2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMetadata2Invoker) Invoke() (*model.DeleteMetadata2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMetadata2Response), nil
	}
}

type DetachEip2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DetachEip2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetachEip2Invoker) Invoke() (*model.DetachEip2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetachEip2Response), nil
	}
}

type ExpandGraph2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ExpandGraph2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExpandGraph2Invoker) Invoke() (*model.ExpandGraph2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExpandGraph2Response), nil
	}
}

type ExportBackup2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ExportBackup2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportBackup2Invoker) Invoke() (*model.ExportBackup2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportBackup2Response), nil
	}
}

type ExportGraph2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ExportGraph2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportGraph2Invoker) Invoke() (*model.ExportGraph2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportGraph2Response), nil
	}
}

type ImportBackup2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ImportBackup2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportBackup2Invoker) Invoke() (*model.ImportBackup2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportBackup2Response), nil
	}
}

type ImportGraph2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ImportGraph2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportGraph2Invoker) Invoke() (*model.ImportGraph2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportGraph2Response), nil
	}
}

type ListBackups2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListBackups2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBackups2Invoker) Invoke() (*model.ListBackups2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBackups2Response), nil
	}
}

type ListGraphBackups2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListGraphBackups2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGraphBackups2Invoker) Invoke() (*model.ListGraphBackups2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGraphBackups2Response), nil
	}
}

type ListGraphs2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListGraphs2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGraphs2Invoker) Invoke() (*model.ListGraphs2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGraphs2Response), nil
	}
}

type ListJobs2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListJobs2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListJobs2Invoker) Invoke() (*model.ListJobs2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJobs2Response), nil
	}
}

type ListMetadatas2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListMetadatas2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMetadatas2Invoker) Invoke() (*model.ListMetadatas2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMetadatas2Response), nil
	}
}

type ListQuotas2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuotas2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQuotas2Invoker) Invoke() (*model.ListQuotas2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotas2Response), nil
	}
}

type ResizeGraph2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeGraph2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResizeGraph2Invoker) Invoke() (*model.ResizeGraph2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeGraph2Response), nil
	}
}

type RestartGraph2Invoker struct {
	*invoker.BaseInvoker
}

func (i *RestartGraph2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RestartGraph2Invoker) Invoke() (*model.RestartGraph2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartGraph2Response), nil
	}
}

type ShowBackupDownloadLinkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBackupDownloadLinkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBackupDownloadLinkInvoker) Invoke() (*model.ShowBackupDownloadLinkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBackupDownloadLinkResponse), nil
	}
}

type ShowGraph2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGraph2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGraph2Invoker) Invoke() (*model.ShowGraph2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGraph2Response), nil
	}
}

type ShowJob2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJob2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJob2Invoker) Invoke() (*model.ShowJob2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJob2Response), nil
	}
}

type ShowMetadata2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMetadata2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMetadata2Invoker) Invoke() (*model.ShowMetadata2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMetadata2Response), nil
	}
}

type StartGraph2Invoker struct {
	*invoker.BaseInvoker
}

func (i *StartGraph2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartGraph2Invoker) Invoke() (*model.StartGraph2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartGraph2Response), nil
	}
}

type StopGraph2Invoker struct {
	*invoker.BaseInvoker
}

func (i *StopGraph2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopGraph2Invoker) Invoke() (*model.StopGraph2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopGraph2Response), nil
	}
}

type UpgradeGraph2Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpgradeGraph2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpgradeGraph2Invoker) Invoke() (*model.UpgradeGraph2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpgradeGraph2Response), nil
	}
}

type UploadFromObs2Invoker struct {
	*invoker.BaseInvoker
}

func (i *UploadFromObs2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UploadFromObs2Invoker) Invoke() (*model.UploadFromObs2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadFromObs2Response), nil
	}
}

type DeregisterScenes2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeregisterScenes2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeregisterScenes2Invoker) Invoke() (*model.DeregisterScenes2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeregisterScenes2Response), nil
	}
}

type ListScenes2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListScenes2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScenes2Invoker) Invoke() (*model.ListScenes2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScenes2Response), nil
	}
}

type RegisterScenes2Invoker struct {
	*invoker.BaseInvoker
}

func (i *RegisterScenes2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RegisterScenes2Invoker) Invoke() (*model.RegisterScenes2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterScenes2Response), nil
	}
}
