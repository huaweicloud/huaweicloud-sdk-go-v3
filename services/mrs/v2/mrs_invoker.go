package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/mrs/v2/model"
)

type BatchDeleteJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteJobsInvoker) Invoke() (*model.BatchDeleteJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteJobsResponse), nil
	}
}

type CreateAutoScalingPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAutoScalingPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAutoScalingPolicyInvoker) Invoke() (*model.CreateAutoScalingPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAutoScalingPolicyResponse), nil
	}
}

type CreateClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateClusterInvoker) Invoke() (*model.CreateClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClusterResponse), nil
	}
}

type CreateExecuteJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateExecuteJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateExecuteJobInvoker) Invoke() (*model.CreateExecuteJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateExecuteJobResponse), nil
	}
}

type DeleteAutoScalingPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAutoScalingPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAutoScalingPolicyInvoker) Invoke() (*model.DeleteAutoScalingPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAutoScalingPolicyResponse), nil
	}
}

type RunJobFlowInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunJobFlowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunJobFlowInvoker) Invoke() (*model.RunJobFlowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunJobFlowResponse), nil
	}
}

type ShowAgencyMappingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAgencyMappingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAgencyMappingInvoker) Invoke() (*model.ShowAgencyMappingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAgencyMappingResponse), nil
	}
}

type ShowAutoScalingPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAutoScalingPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAutoScalingPolicyInvoker) Invoke() (*model.ShowAutoScalingPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAutoScalingPolicyResponse), nil
	}
}

type ShowJobExeListNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobExeListNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobExeListNewInvoker) Invoke() (*model.ShowJobExeListNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobExeListNewResponse), nil
	}
}

type ShowSingleJobExeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSingleJobExeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSingleJobExeInvoker) Invoke() (*model.ShowSingleJobExeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSingleJobExeResponse), nil
	}
}

type ShowSqlResultWithJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSqlResultWithJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSqlResultWithJobInvoker) Invoke() (*model.ShowSqlResultWithJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlResultWithJobResponse), nil
	}
}

type StopJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopJobInvoker) Invoke() (*model.StopJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopJobResponse), nil
	}
}

type UpdateAgencyMappingInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAgencyMappingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAgencyMappingInvoker) Invoke() (*model.UpdateAgencyMappingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAgencyMappingResponse), nil
	}
}

type UpdateAutoScalingPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAutoScalingPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAutoScalingPolicyInvoker) Invoke() (*model.UpdateAutoScalingPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAutoScalingPolicyResponse), nil
	}
}

type UpdateClusterNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateClusterNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateClusterNameInvoker) Invoke() (*model.UpdateClusterNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateClusterNameResponse), nil
	}
}

type AddComponentInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddComponentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddComponentInvoker) Invoke() (*model.AddComponentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddComponentResponse), nil
	}
}

type ExpandClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExpandClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExpandClusterInvoker) Invoke() (*model.ExpandClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExpandClusterResponse), nil
	}
}

type ListNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNodesInvoker) Invoke() (*model.ListNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNodesResponse), nil
	}
}

type ShrinkClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShrinkClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShrinkClusterInvoker) Invoke() (*model.ShrinkClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShrinkClusterResponse), nil
	}
}

type CreateDataConnectorInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDataConnectorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDataConnectorInvoker) Invoke() (*model.CreateDataConnectorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDataConnectorResponse), nil
	}
}

type DeleteDataConnectorInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDataConnectorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDataConnectorInvoker) Invoke() (*model.DeleteDataConnectorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDataConnectorResponse), nil
	}
}

type ListDataConnectorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDataConnectorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDataConnectorInvoker) Invoke() (*model.ListDataConnectorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDataConnectorResponse), nil
	}
}

type UpdateDataConnectorInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDataConnectorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDataConnectorInvoker) Invoke() (*model.UpdateDataConnectorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDataConnectorResponse), nil
	}
}

type ShowHdfsFileListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHdfsFileListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHdfsFileListInvoker) Invoke() (*model.ShowHdfsFileListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHdfsFileListResponse), nil
	}
}

type CancelSyncIamUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelSyncIamUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelSyncIamUserInvoker) Invoke() (*model.CancelSyncIamUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelSyncIamUserResponse), nil
	}
}

type ShowSyncIamUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSyncIamUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSyncIamUserInvoker) Invoke() (*model.ShowSyncIamUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSyncIamUserResponse), nil
	}
}

type UpdateSyncIamUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSyncIamUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSyncIamUserInvoker) Invoke() (*model.UpdateSyncIamUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSyncIamUserResponse), nil
	}
}

type CancelSqlInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelSqlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelSqlInvoker) Invoke() (*model.CancelSqlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelSqlResponse), nil
	}
}

type ExecuteSqlInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteSqlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteSqlInvoker) Invoke() (*model.ExecuteSqlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteSqlResponse), nil
	}
}

type ShowSqlResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSqlResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSqlResultInvoker) Invoke() (*model.ShowSqlResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlResultResponse), nil
	}
}

type ShowTagQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTagQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTagQuotaInvoker) Invoke() (*model.ShowTagQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTagQuotaResponse), nil
	}
}

type ShowTagStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTagStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTagStatusInvoker) Invoke() (*model.ShowTagStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTagStatusResponse), nil
	}
}

type SwitchClusterTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchClusterTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchClusterTagsInvoker) Invoke() (*model.SwitchClusterTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchClusterTagsResponse), nil
	}
}

type ShowMrsFlavorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMrsFlavorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMrsFlavorsInvoker) Invoke() (*model.ShowMrsFlavorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMrsFlavorsResponse), nil
	}
}

type ShowMrsVersionListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMrsVersionListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMrsVersionListInvoker) Invoke() (*model.ShowMrsVersionListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMrsVersionListResponse), nil
	}
}
