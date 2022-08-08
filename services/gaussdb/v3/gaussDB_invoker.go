package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/gaussdb/v3/model"
)

type BatchTagActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchTagActionInvoker) Invoke() (*model.BatchTagActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchTagActionResponse), nil
	}
}

type ChangeGaussMySqlInstanceSpecificationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeGaussMySqlInstanceSpecificationInvoker) Invoke() (*model.ChangeGaussMySqlInstanceSpecificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeGaussMySqlInstanceSpecificationResponse), nil
	}
}

type ChangeGaussMySqlProxySpecificationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeGaussMySqlProxySpecificationInvoker) Invoke() (*model.ChangeGaussMySqlProxySpecificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeGaussMySqlProxySpecificationResponse), nil
	}
}

type CreateGaussMySqlBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGaussMySqlBackupInvoker) Invoke() (*model.CreateGaussMySqlBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGaussMySqlBackupResponse), nil
	}
}

type CreateGaussMySqlInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGaussMySqlInstanceInvoker) Invoke() (*model.CreateGaussMySqlInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGaussMySqlInstanceResponse), nil
	}
}

type CreateGaussMySqlProxyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGaussMySqlProxyInvoker) Invoke() (*model.CreateGaussMySqlProxyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGaussMySqlProxyResponse), nil
	}
}

type CreateGaussMySqlReadonlyNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGaussMySqlReadonlyNodeInvoker) Invoke() (*model.CreateGaussMySqlReadonlyNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGaussMySqlReadonlyNodeResponse), nil
	}
}

type DeleteGaussMySqlInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGaussMySqlInstanceInvoker) Invoke() (*model.DeleteGaussMySqlInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGaussMySqlInstanceResponse), nil
	}
}

type DeleteGaussMySqlProxyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGaussMySqlProxyInvoker) Invoke() (*model.DeleteGaussMySqlProxyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGaussMySqlProxyResponse), nil
	}
}

type DeleteGaussMySqlReadonlyNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGaussMySqlReadonlyNodeInvoker) Invoke() (*model.DeleteGaussMySqlReadonlyNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGaussMySqlReadonlyNodeResponse), nil
	}
}

type ExpandGaussMySqlInstanceVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExpandGaussMySqlInstanceVolumeInvoker) Invoke() (*model.ExpandGaussMySqlInstanceVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExpandGaussMySqlInstanceVolumeResponse), nil
	}
}

type ExpandGaussMySqlProxyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExpandGaussMySqlProxyInvoker) Invoke() (*model.ExpandGaussMySqlProxyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExpandGaussMySqlProxyResponse), nil
	}
}

type ListGaussMySqlConfigurationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGaussMySqlConfigurationsInvoker) Invoke() (*model.ListGaussMySqlConfigurationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGaussMySqlConfigurationsResponse), nil
	}
}

type ListGaussMySqlDedicatedResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGaussMySqlDedicatedResourcesInvoker) Invoke() (*model.ListGaussMySqlDedicatedResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGaussMySqlDedicatedResourcesResponse), nil
	}
}

type ListGaussMySqlErrorLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGaussMySqlErrorLogInvoker) Invoke() (*model.ListGaussMySqlErrorLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGaussMySqlErrorLogResponse), nil
	}
}

type ListGaussMySqlInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGaussMySqlInstancesInvoker) Invoke() (*model.ListGaussMySqlInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGaussMySqlInstancesResponse), nil
	}
}

type ListGaussMySqlSlowLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGaussMySqlSlowLogInvoker) Invoke() (*model.ListGaussMySqlSlowLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGaussMySqlSlowLogResponse), nil
	}
}

type ListInstanceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceTagsInvoker) Invoke() (*model.ListInstanceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceTagsResponse), nil
	}
}

type ListProjectTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectTagsInvoker) Invoke() (*model.ListProjectTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectTagsResponse), nil
	}
}

type ResetGaussMySqlPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetGaussMySqlPasswordInvoker) Invoke() (*model.ResetGaussMySqlPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetGaussMySqlPasswordResponse), nil
	}
}

type SetGaussMySqlProxyWeightInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetGaussMySqlProxyWeightInvoker) Invoke() (*model.SetGaussMySqlProxyWeightResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetGaussMySqlProxyWeightResponse), nil
	}
}

type SetGaussMySqlQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetGaussMySqlQuotasInvoker) Invoke() (*model.SetGaussMySqlQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetGaussMySqlQuotasResponse), nil
	}
}

type ShowAuditLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuditLogInvoker) Invoke() (*model.ShowAuditLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuditLogResponse), nil
	}
}

type ShowDedicatedResourceInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDedicatedResourceInfoInvoker) Invoke() (*model.ShowDedicatedResourceInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDedicatedResourceInfoResponse), nil
	}
}

type ShowGaussMySqlBackupListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGaussMySqlBackupListInvoker) Invoke() (*model.ShowGaussMySqlBackupListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGaussMySqlBackupListResponse), nil
	}
}

type ShowGaussMySqlBackupPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGaussMySqlBackupPolicyInvoker) Invoke() (*model.ShowGaussMySqlBackupPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGaussMySqlBackupPolicyResponse), nil
	}
}

type ShowGaussMySqlEngineVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGaussMySqlEngineVersionInvoker) Invoke() (*model.ShowGaussMySqlEngineVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGaussMySqlEngineVersionResponse), nil
	}
}

type ShowGaussMySqlFlavorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGaussMySqlFlavorsInvoker) Invoke() (*model.ShowGaussMySqlFlavorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGaussMySqlFlavorsResponse), nil
	}
}

type ShowGaussMySqlInstanceInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGaussMySqlInstanceInfoInvoker) Invoke() (*model.ShowGaussMySqlInstanceInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGaussMySqlInstanceInfoResponse), nil
	}
}

type ShowGaussMySqlJobInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGaussMySqlJobInfoInvoker) Invoke() (*model.ShowGaussMySqlJobInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGaussMySqlJobInfoResponse), nil
	}
}

type ShowGaussMySqlProjectQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGaussMySqlProjectQuotasInvoker) Invoke() (*model.ShowGaussMySqlProjectQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGaussMySqlProjectQuotasResponse), nil
	}
}

type ShowGaussMySqlProxyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGaussMySqlProxyInvoker) Invoke() (*model.ShowGaussMySqlProxyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGaussMySqlProxyResponse), nil
	}
}

type ShowGaussMySqlProxyFlavorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGaussMySqlProxyFlavorsInvoker) Invoke() (*model.ShowGaussMySqlProxyFlavorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGaussMySqlProxyFlavorsResponse), nil
	}
}

type ShowGaussMySqlProxyListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGaussMySqlProxyListInvoker) Invoke() (*model.ShowGaussMySqlProxyListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGaussMySqlProxyListResponse), nil
	}
}

type ShowGaussMySqlQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGaussMySqlQuotasInvoker) Invoke() (*model.ShowGaussMySqlQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGaussMySqlQuotasResponse), nil
	}
}

type ShowInstanceMonitorExtendInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceMonitorExtendInvoker) Invoke() (*model.ShowInstanceMonitorExtendResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceMonitorExtendResponse), nil
	}
}

type UpdateAuditLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAuditLogInvoker) Invoke() (*model.UpdateAuditLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAuditLogResponse), nil
	}
}

type UpdateGaussMySqlBackupPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGaussMySqlBackupPolicyInvoker) Invoke() (*model.UpdateGaussMySqlBackupPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGaussMySqlBackupPolicyResponse), nil
	}
}

type UpdateGaussMySqlInstanceNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGaussMySqlInstanceNameInvoker) Invoke() (*model.UpdateGaussMySqlInstanceNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGaussMySqlInstanceNameResponse), nil
	}
}

type UpdateGaussMySqlQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGaussMySqlQuotasInvoker) Invoke() (*model.UpdateGaussMySqlQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGaussMySqlQuotasResponse), nil
	}
}

type UpdateInstanceMonitorInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceMonitorInvoker) Invoke() (*model.UpdateInstanceMonitorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceMonitorResponse), nil
	}
}

type DeleteSqlFilterRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSqlFilterRuleInvoker) Invoke() (*model.DeleteSqlFilterRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSqlFilterRuleResponse), nil
	}
}

type SetSqlFilterRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetSqlFilterRuleInvoker) Invoke() (*model.SetSqlFilterRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetSqlFilterRuleResponse), nil
	}
}

type ShowSqlFilterControlInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSqlFilterControlInvoker) Invoke() (*model.ShowSqlFilterControlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlFilterControlResponse), nil
	}
}

type ShowSqlFilterRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSqlFilterRuleInvoker) Invoke() (*model.ShowSqlFilterRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlFilterRuleResponse), nil
	}
}

type UpdateSqlFilterControlInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSqlFilterControlInvoker) Invoke() (*model.UpdateSqlFilterControlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSqlFilterControlResponse), nil
	}
}
