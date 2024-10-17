package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/aad/v1/model"
)

type ExecuteUnblockIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteUnblockIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListBlockIpsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListUnblockQuotaStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowBlockStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowUnblockRecordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowUnblockRecordInvoker) Invoke() (*model.ShowUnblockRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUnblockRecordResponse), nil
	}
}

type AddBlackWhiteIpListInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddBlackWhiteIpListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddBlackWhiteIpListInvoker) Invoke() (*model.AddBlackWhiteIpListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddBlackWhiteIpListResponse), nil
	}
}

type AddPolicyBlackAndWhiteIpListInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddPolicyBlackAndWhiteIpListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *AssociateIpToPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateIpToPolicyInvoker) Invoke() (*model.AssociateIpToPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateIpToPolicyResponse), nil
	}
}

type AssociateIpToPolicyAndPackageInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateIpToPolicyAndPackageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateIpToPolicyAndPackageInvoker) Invoke() (*model.AssociateIpToPolicyAndPackageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateIpToPolicyAndPackageResponse), nil
	}
}

type BatchCreateInstanceIpRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateInstanceIpRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *BatchDeleteInstanceIpRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteInstanceIpRuleInvoker) Invoke() (*model.BatchDeleteInstanceIpRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteInstanceIpRuleResponse), nil
	}
}

type CreateAadDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAadDomainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAadDomainInvoker) Invoke() (*model.CreateAadDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAadDomainResponse), nil
	}
}

type CreatePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteAlarmConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAlarmConfigInvoker) Invoke() (*model.DeleteAlarmConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAlarmConfigResponse), nil
	}
}

type DeleteBlackWhiteIpListInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBlackWhiteIpListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteBlackWhiteIpListInvoker) Invoke() (*model.DeleteBlackWhiteIpListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBlackWhiteIpListResponse), nil
	}
}

type DeletePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeletePolicyBlackAndWhiteIpListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DisassociateIpFromPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisassociateIpFromPolicyInvoker) Invoke() (*model.DisassociateIpFromPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateIpFromPolicyResponse), nil
	}
}

type DisassociateIpFromPolicyAndPackageInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateIpFromPolicyAndPackageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisassociateIpFromPolicyAndPackageInvoker) Invoke() (*model.DisassociateIpFromPolicyAndPackageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateIpFromPolicyAndPackageResponse), nil
	}
}

type ListDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDomainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListInstanceIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListInstanceIpRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListPackageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListPeakInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListProtectedIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProtectedIpInvoker) Invoke() (*model.ListProtectedIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProtectedIpResponse), nil
	}
}

type ListSourceIpsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSourceIpsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSourceIpsInvoker) Invoke() (*model.ListSourceIpsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSourceIpsResponse), nil
	}
}

type ListUnboundProtectedIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUnboundProtectedIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUnboundProtectedIpInvoker) Invoke() (*model.ListUnboundProtectedIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUnboundProtectedIpResponse), nil
	}
}

type ModifyDomainWebSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyDomainWebSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyDomainWebSwitchInvoker) Invoke() (*model.ModifyDomainWebSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyDomainWebSwitchResponse), nil
	}
}

type SetCertForDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetCertForDomainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetCertForDomainInvoker) Invoke() (*model.SetCertForDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetCertForDomainResponse), nil
	}
}

type ShowAlarmConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAlarmConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateAlarmConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateDomainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateInstanceIpRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdatePackageIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdatePackageNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdatePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateTagForProtectedIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTagForProtectedIpInvoker) Invoke() (*model.UpdateTagForProtectedIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTagForProtectedIpResponse), nil
	}
}
