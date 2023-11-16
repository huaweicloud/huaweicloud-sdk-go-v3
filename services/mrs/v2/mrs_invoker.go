package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/mrs/v2/model"
)

type BatchDeleteJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteJobsInvoker) Invoke() (*model.BatchDeleteJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteJobsResponse), nil
	}
}

type CreateClusterInvoker struct {
	*invoker.BaseInvoker
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

func (i *CreateExecuteJobInvoker) Invoke() (*model.CreateExecuteJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateExecuteJobResponse), nil
	}
}

type RunJobFlowInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdateAgencyMappingInvoker) Invoke() (*model.UpdateAgencyMappingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAgencyMappingResponse), nil
	}
}

type UpdateClusterNameInvoker struct {
	*invoker.BaseInvoker
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

func (i *ExpandClusterInvoker) Invoke() (*model.ExpandClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExpandClusterResponse), nil
	}
}

type ShrinkClusterInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowHdfsFileListInvoker) Invoke() (*model.ShowHdfsFileListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHdfsFileListResponse), nil
	}
}

type CancelSqlInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowSqlResultInvoker) Invoke() (*model.ShowSqlResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlResultResponse), nil
	}
}

type ShowMrsVersionListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMrsVersionListInvoker) Invoke() (*model.ShowMrsVersionListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMrsVersionListResponse), nil
	}
}
