package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/aad/v1/model"
)

type ExecuteUnblockIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteUnblockIpInvoker) Invoke() (*model.ExecuteUnblockIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteUnblockIpResponse), nil
	}
}

type ListBlockIpsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBlockIpsInvoker) Invoke() (*model.ListBlockIpsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBlockIpsResponse), nil
	}
}

type ListUnblockQuotaStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUnblockQuotaStatisticsInvoker) Invoke() (*model.ListUnblockQuotaStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUnblockQuotaStatisticsResponse), nil
	}
}

type ShowBlockStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBlockStatisticsInvoker) Invoke() (*model.ShowBlockStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBlockStatisticsResponse), nil
	}
}

type ShowUnblockRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUnblockRecordInvoker) Invoke() (*model.ShowUnblockRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUnblockRecordResponse), nil
	}
}

type AddPolicyBlackAndWhiteIpListInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddPolicyBlackAndWhiteIpListInvoker) Invoke() (*model.AddPolicyBlackAndWhiteIpListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddPolicyBlackAndWhiteIpListResponse), nil
	}
}

type AssociateIpToPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateIpToPolicyInvoker) Invoke() (*model.AssociateIpToPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateIpToPolicyResponse), nil
	}
}

type BatchCreateInstanceIpRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateInstanceIpRuleInvoker) Invoke() (*model.BatchCreateInstanceIpRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateInstanceIpRuleResponse), nil
	}
}

type BatchDeleteInstanceIpRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteInstanceIpRuleInvoker) Invoke() (*model.BatchDeleteInstanceIpRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteInstanceIpRuleResponse), nil
	}
}

type CreatePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePolicyInvoker) Invoke() (*model.CreatePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePolicyResponse), nil
	}
}

type DeleteAlarmConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAlarmConfigInvoker) Invoke() (*model.DeleteAlarmConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAlarmConfigResponse), nil
	}
}

type DeletePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePolicyInvoker) Invoke() (*model.DeletePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePolicyResponse), nil
	}
}

type DeletePolicyBlackAndWhiteIpListInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePolicyBlackAndWhiteIpListInvoker) Invoke() (*model.DeletePolicyBlackAndWhiteIpListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePolicyBlackAndWhiteIpListResponse), nil
	}
}

type DisassociateIpFromPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateIpFromPolicyInvoker) Invoke() (*model.DisassociateIpFromPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateIpFromPolicyResponse), nil
	}
}

type ListDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDomainInvoker) Invoke() (*model.ListDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDomainResponse), nil
	}
}

type ListInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceInvoker) Invoke() (*model.ListInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceResponse), nil
	}
}

type ListInstanceIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceIdInvoker) Invoke() (*model.ListInstanceIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceIdResponse), nil
	}
}

type ListInstanceIpRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceIpRuleInvoker) Invoke() (*model.ListInstanceIpRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceIpRuleResponse), nil
	}
}

type ListPackageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPackageInvoker) Invoke() (*model.ListPackageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPackageResponse), nil
	}
}

type ListPeakInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPeakInvoker) Invoke() (*model.ListPeakResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPeakResponse), nil
	}
}

type ListPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPolicyInvoker) Invoke() (*model.ListPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPolicyResponse), nil
	}
}

type ListProtectedIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProtectedIpInvoker) Invoke() (*model.ListProtectedIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProtectedIpResponse), nil
	}
}

type ListUnboundProtectedIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUnboundProtectedIpInvoker) Invoke() (*model.ListUnboundProtectedIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUnboundProtectedIpResponse), nil
	}
}

type ShowAlarmConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAlarmConfigInvoker) Invoke() (*model.ShowAlarmConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAlarmConfigResponse), nil
	}
}

type ShowPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPolicyInvoker) Invoke() (*model.ShowPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPolicyResponse), nil
	}
}

type UpdateAlarmConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAlarmConfigInvoker) Invoke() (*model.UpdateAlarmConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAlarmConfigResponse), nil
	}
}

type UpdateDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDomainInvoker) Invoke() (*model.UpdateDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDomainResponse), nil
	}
}

type UpdateInstanceIpRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceIpRuleInvoker) Invoke() (*model.UpdateInstanceIpRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceIpRuleResponse), nil
	}
}

type UpdatePackageIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePackageIpInvoker) Invoke() (*model.UpdatePackageIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePackageIpResponse), nil
	}
}

type UpdatePackageNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePackageNameInvoker) Invoke() (*model.UpdatePackageNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePackageNameResponse), nil
	}
}

type UpdatePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePolicyInvoker) Invoke() (*model.UpdatePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePolicyResponse), nil
	}
}

type UpdateTagForProtectedIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTagForProtectedIpInvoker) Invoke() (*model.UpdateTagForProtectedIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTagForProtectedIpResponse), nil
	}
}
