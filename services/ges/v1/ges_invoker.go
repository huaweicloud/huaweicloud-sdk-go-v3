package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ges/v1/model"
)

type AttachEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachEipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AttachEipInvoker) Invoke() (*model.AttachEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachEipResponse), nil
	}
}

type ClearGraphInvoker struct {
	*invoker.BaseInvoker
}

func (i *ClearGraphInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ClearGraphInvoker) Invoke() (*model.ClearGraphResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ClearGraphResponse), nil
	}
}

type CreateBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBackupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateBackupInvoker) Invoke() (*model.CreateBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBackupResponse), nil
	}
}

type CreateGraphInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGraphInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateGraphInvoker) Invoke() (*model.CreateGraphResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGraphResponse), nil
	}
}

type CreateMetadataInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMetadataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMetadataInvoker) Invoke() (*model.CreateMetadataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMetadataResponse), nil
	}
}

type DeleteBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBackupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteBackupInvoker) Invoke() (*model.DeleteBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBackupResponse), nil
	}
}

type DeleteGraphInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGraphInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteGraphInvoker) Invoke() (*model.DeleteGraphResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGraphResponse), nil
	}
}

type DeleteMetadataInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMetadataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMetadataInvoker) Invoke() (*model.DeleteMetadataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMetadataResponse), nil
	}
}

type DetachEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetachEipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetachEipInvoker) Invoke() (*model.DetachEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetachEipResponse), nil
	}
}

type ExpandGraphInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExpandGraphInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExpandGraphInvoker) Invoke() (*model.ExpandGraphResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExpandGraphResponse), nil
	}
}

type ExportGraphInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportGraphInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportGraphInvoker) Invoke() (*model.ExportGraphResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportGraphResponse), nil
	}
}

type ImportGraphInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportGraphInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportGraphInvoker) Invoke() (*model.ImportGraphResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportGraphResponse), nil
	}
}

type ListBackupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBackupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBackupsInvoker) Invoke() (*model.ListBackupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBackupsResponse), nil
	}
}

type ListGraphBackupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGraphBackupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGraphBackupsInvoker) Invoke() (*model.ListGraphBackupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGraphBackupsResponse), nil
	}
}

type ListGraphMetadatasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGraphMetadatasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGraphMetadatasInvoker) Invoke() (*model.ListGraphMetadatasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGraphMetadatasResponse), nil
	}
}

type ListGraphsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGraphsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGraphsInvoker) Invoke() (*model.ListGraphsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGraphsResponse), nil
	}
}

type ListJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListJobsInvoker) Invoke() (*model.ListJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJobsResponse), nil
	}
}

type ListMetadatasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMetadatasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMetadatasInvoker) Invoke() (*model.ListMetadatasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMetadatasResponse), nil
	}
}

type ListQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQuotasInvoker) Invoke() (*model.ListQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotasResponse), nil
	}
}

type ResizeGraphInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeGraphInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResizeGraphInvoker) Invoke() (*model.ResizeGraphResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeGraphResponse), nil
	}
}

type RestartGraphInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestartGraphInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RestartGraphInvoker) Invoke() (*model.RestartGraphResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartGraphResponse), nil
	}
}

type ShowGraphInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGraphInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGraphInvoker) Invoke() (*model.ShowGraphResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGraphResponse), nil
	}
}

type ShowJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobInvoker) Invoke() (*model.ShowJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobResponse), nil
	}
}

type StartGraphInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartGraphInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartGraphInvoker) Invoke() (*model.StartGraphResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartGraphResponse), nil
	}
}

type StopGraphInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopGraphInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopGraphInvoker) Invoke() (*model.StopGraphResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopGraphResponse), nil
	}
}

type UpgradeGraphInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpgradeGraphInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpgradeGraphInvoker) Invoke() (*model.UpgradeGraphResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpgradeGraphResponse), nil
	}
}

type UploadFromObsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadFromObsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UploadFromObsInvoker) Invoke() (*model.UploadFromObsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadFromObsResponse), nil
	}
}
