package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/aidatalake/v2/model"
)

type AddV2TenantImageVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddV2TenantImageVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddV2TenantImageVersionInvoker) Invoke() (*model.AddV2TenantImageVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddV2TenantImageVersionResponse), nil
	}
}

type AttachV2AgencyPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachV2AgencyPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AttachV2AgencyPolicyInvoker) Invoke() (*model.AttachV2AgencyPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachV2AgencyPolicyResponse), nil
	}
}

type CreateV2AgencyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateV2AgencyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateV2AgencyInvoker) Invoke() (*model.CreateV2AgencyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateV2AgencyResponse), nil
	}
}

type CreateV2ComputeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateV2ComputeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateV2ComputeInvoker) Invoke() (*model.CreateV2ComputeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateV2ComputeResponse), nil
	}
}

type CreateV2EndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateV2EndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateV2EndpointInvoker) Invoke() (*model.CreateV2EndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateV2EndpointResponse), nil
	}
}

type CreateV2NetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateV2NetworkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateV2NetworkInvoker) Invoke() (*model.CreateV2NetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateV2NetworkResponse), nil
	}
}

type CreateV2TenantImageSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateV2TenantImageSnapshotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateV2TenantImageSnapshotInvoker) Invoke() (*model.CreateV2TenantImageSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateV2TenantImageSnapshotResponse), nil
	}
}

type CreateV2WorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateV2WorkspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateV2WorkspaceInvoker) Invoke() (*model.CreateV2WorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateV2WorkspaceResponse), nil
	}
}

type DeleteV2AgencyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteV2AgencyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteV2AgencyInvoker) Invoke() (*model.DeleteV2AgencyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteV2AgencyResponse), nil
	}
}

type DeleteV2AgreementInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteV2AgreementInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteV2AgreementInvoker) Invoke() (*model.DeleteV2AgreementResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteV2AgreementResponse), nil
	}
}

type DeleteV2ComputeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteV2ComputeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteV2ComputeInvoker) Invoke() (*model.DeleteV2ComputeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteV2ComputeResponse), nil
	}
}

type DeleteV2EndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteV2EndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteV2EndpointInvoker) Invoke() (*model.DeleteV2EndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteV2EndpointResponse), nil
	}
}

type DeleteV2TenantImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteV2TenantImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteV2TenantImageInvoker) Invoke() (*model.DeleteV2TenantImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteV2TenantImageResponse), nil
	}
}

type DeleteV2TenantImageSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteV2TenantImageSnapshotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteV2TenantImageSnapshotInvoker) Invoke() (*model.DeleteV2TenantImageSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteV2TenantImageSnapshotResponse), nil
	}
}

type DeleteV2TenantImageVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteV2TenantImageVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteV2TenantImageVersionInvoker) Invoke() (*model.DeleteV2TenantImageVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteV2TenantImageVersionResponse), nil
	}
}

type DeleteV2WorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteV2WorkspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteV2WorkspaceInvoker) Invoke() (*model.DeleteV2WorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteV2WorkspaceResponse), nil
	}
}

type DetachV2AgencyPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetachV2AgencyPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetachV2AgencyPolicyInvoker) Invoke() (*model.DetachV2AgencyPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetachV2AgencyPolicyResponse), nil
	}
}

type ListV2AgreementRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListV2AgreementRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListV2AgreementRuleInvoker) Invoke() (*model.ListV2AgreementRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListV2AgreementRuleResponse), nil
	}
}

type ListV2AllEndpointsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListV2AllEndpointsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListV2AllEndpointsInvoker) Invoke() (*model.ListV2AllEndpointsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListV2AllEndpointsResponse), nil
	}
}

type ListV2ComputeVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListV2ComputeVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListV2ComputeVersionsInvoker) Invoke() (*model.ListV2ComputeVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListV2ComputeVersionsResponse), nil
	}
}

type ListV2ComputesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListV2ComputesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListV2ComputesInvoker) Invoke() (*model.ListV2ComputesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListV2ComputesResponse), nil
	}
}

type ListV2EndpointsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListV2EndpointsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListV2EndpointsInvoker) Invoke() (*model.ListV2EndpointsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListV2EndpointsResponse), nil
	}
}

type ListV2NetworksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListV2NetworksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListV2NetworksInvoker) Invoke() (*model.ListV2NetworksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListV2NetworksResponse), nil
	}
}

type ListV2TenantImageSnapshotsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListV2TenantImageSnapshotsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListV2TenantImageSnapshotsInvoker) Invoke() (*model.ListV2TenantImageSnapshotsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListV2TenantImageSnapshotsResponse), nil
	}
}

type ListV2TenantImageVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListV2TenantImageVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListV2TenantImageVersionsInvoker) Invoke() (*model.ListV2TenantImageVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListV2TenantImageVersionsResponse), nil
	}
}

type ListV2TenantImagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListV2TenantImagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListV2TenantImagesInvoker) Invoke() (*model.ListV2TenantImagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListV2TenantImagesResponse), nil
	}
}

type ListV2WorkspacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListV2WorkspacesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListV2WorkspacesInvoker) Invoke() (*model.ListV2WorkspacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListV2WorkspacesResponse), nil
	}
}

type RegisterV2TenantImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *RegisterV2TenantImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RegisterV2TenantImageInvoker) Invoke() (*model.RegisterV2TenantImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterV2TenantImageResponse), nil
	}
}

type ShowV2AgencyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowV2AgencyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowV2AgencyInvoker) Invoke() (*model.ShowV2AgencyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowV2AgencyResponse), nil
	}
}

type ShowV2AgreementInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowV2AgreementInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowV2AgreementInvoker) Invoke() (*model.ShowV2AgreementResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowV2AgreementResponse), nil
	}
}

type ShowV2ComputeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowV2ComputeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowV2ComputeInvoker) Invoke() (*model.ShowV2ComputeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowV2ComputeResponse), nil
	}
}

type ShowV2ComputesMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowV2ComputesMetricsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowV2ComputesMetricsInvoker) Invoke() (*model.ShowV2ComputesMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowV2ComputesMetricsResponse), nil
	}
}

type ShowV2EndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowV2EndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowV2EndpointInvoker) Invoke() (*model.ShowV2EndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowV2EndpointResponse), nil
	}
}

type ShowV2EndpointResourceLimitInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowV2EndpointResourceLimitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowV2EndpointResourceLimitInvoker) Invoke() (*model.ShowV2EndpointResourceLimitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowV2EndpointResourceLimitResponse), nil
	}
}

type ShowV2EndpointsMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowV2EndpointsMetricsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowV2EndpointsMetricsInvoker) Invoke() (*model.ShowV2EndpointsMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowV2EndpointsMetricsResponse), nil
	}
}

type ShowV2TenantImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowV2TenantImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowV2TenantImageInvoker) Invoke() (*model.ShowV2TenantImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowV2TenantImageResponse), nil
	}
}

type ShowV2TenantImageVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowV2TenantImageVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowV2TenantImageVersionInvoker) Invoke() (*model.ShowV2TenantImageVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowV2TenantImageVersionResponse), nil
	}
}

type ShowV2WorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowV2WorkspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowV2WorkspaceInvoker) Invoke() (*model.ShowV2WorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowV2WorkspaceResponse), nil
	}
}

type SignV2AgreementInvoker struct {
	*invoker.BaseInvoker
}

func (i *SignV2AgreementInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SignV2AgreementInvoker) Invoke() (*model.SignV2AgreementResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SignV2AgreementResponse), nil
	}
}

type UpdateV2ComputeInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateV2ComputeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateV2ComputeInvoker) Invoke() (*model.UpdateV2ComputeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateV2ComputeResponse), nil
	}
}

type UpdateV2EndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateV2EndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateV2EndpointInvoker) Invoke() (*model.UpdateV2EndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateV2EndpointResponse), nil
	}
}

type UpdateV2TenantImageVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateV2TenantImageVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateV2TenantImageVersionInvoker) Invoke() (*model.UpdateV2TenantImageVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateV2TenantImageVersionResponse), nil
	}
}

type UpdateV2WorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateV2WorkspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateV2WorkspaceInvoker) Invoke() (*model.UpdateV2WorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateV2WorkspaceResponse), nil
	}
}

type CancelAuraDataframeJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelAuraDataframeJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelAuraDataframeJobInvoker) Invoke() (*model.CancelAuraDataframeJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelAuraDataframeJobResponse), nil
	}
}

type CreateAuraDataframeJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAuraDataframeJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAuraDataframeJobInvoker) Invoke() (*model.CreateAuraDataframeJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAuraDataframeJobResponse), nil
	}
}

type ListAuraDataframeJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuraDataframeJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuraDataframeJobInvoker) Invoke() (*model.ListAuraDataframeJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuraDataframeJobResponse), nil
	}
}

type ShowAuraDataframeJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuraDataframeJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuraDataframeJobInvoker) Invoke() (*model.ShowAuraDataframeJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuraDataframeJobResponse), nil
	}
}

type CancelAuraSqlStatementInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelAuraSqlStatementInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelAuraSqlStatementInvoker) Invoke() (*model.CancelAuraSqlStatementResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelAuraSqlStatementResponse), nil
	}
}

type CloseAuraSqlSessionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CloseAuraSqlSessionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CloseAuraSqlSessionInvoker) Invoke() (*model.CloseAuraSqlSessionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CloseAuraSqlSessionResponse), nil
	}
}

type CreateAuraSqlSessionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAuraSqlSessionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAuraSqlSessionInvoker) Invoke() (*model.CreateAuraSqlSessionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAuraSqlSessionResponse), nil
	}
}

type ExecuteAuraSqlStatementInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteAuraSqlStatementInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteAuraSqlStatementInvoker) Invoke() (*model.ExecuteAuraSqlStatementResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteAuraSqlStatementResponse), nil
	}
}

type ListAuraSessionStatementRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuraSessionStatementRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuraSessionStatementRecordsInvoker) Invoke() (*model.ListAuraSessionStatementRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuraSessionStatementRecordsResponse), nil
	}
}

type ListAuraSqlSessionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuraSqlSessionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuraSqlSessionsInvoker) Invoke() (*model.ListAuraSqlSessionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuraSqlSessionsResponse), nil
	}
}

type ListAuraStatementRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuraStatementRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuraStatementRecordsInvoker) Invoke() (*model.ListAuraStatementRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuraStatementRecordsResponse), nil
	}
}

type ShowAuraSqlSessionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuraSqlSessionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuraSqlSessionInvoker) Invoke() (*model.ShowAuraSqlSessionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuraSqlSessionResponse), nil
	}
}

type ShowAuraSqlStatementResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuraSqlStatementResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuraSqlStatementResultInvoker) Invoke() (*model.ShowAuraSqlStatementResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuraSqlStatementResultResponse), nil
	}
}

type ShowAuraStatementOperatorMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuraStatementOperatorMetricsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuraStatementOperatorMetricsInvoker) Invoke() (*model.ShowAuraStatementOperatorMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuraStatementOperatorMetricsResponse), nil
	}
}

type ShowAuraStatementQueryMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuraStatementQueryMetricsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuraStatementQueryMetricsInvoker) Invoke() (*model.ShowAuraStatementQueryMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuraStatementQueryMetricsResponse), nil
	}
}

type ListFlinkJobAutoscalingHistoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlinkJobAutoscalingHistoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFlinkJobAutoscalingHistoriesInvoker) Invoke() (*model.ListFlinkJobAutoscalingHistoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlinkJobAutoscalingHistoriesResponse), nil
	}
}

type ShowFlinkJobAutoscalingPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFlinkJobAutoscalingPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFlinkJobAutoscalingPolicyInvoker) Invoke() (*model.ShowFlinkJobAutoscalingPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFlinkJobAutoscalingPolicyResponse), nil
	}
}

type UpdateFlinkJobAutoscalingPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFlinkJobAutoscalingPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateFlinkJobAutoscalingPolicyInvoker) Invoke() (*model.UpdateFlinkJobAutoscalingPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFlinkJobAutoscalingPolicyResponse), nil
	}
}

type CancelFlinkJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelFlinkJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelFlinkJobInvoker) Invoke() (*model.CancelFlinkJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelFlinkJobResponse), nil
	}
}

type ListFlinkJobEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlinkJobEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFlinkJobEventsInvoker) Invoke() (*model.ListFlinkJobEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlinkJobEventsResponse), nil
	}
}

type ListFlinkJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlinkJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFlinkJobsInvoker) Invoke() (*model.ListFlinkJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlinkJobsResponse), nil
	}
}

type RunFlinkJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunFlinkJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunFlinkJobInvoker) Invoke() (*model.RunFlinkJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunFlinkJobResponse), nil
	}
}

type ShowFlinkJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFlinkJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFlinkJobInvoker) Invoke() (*model.ShowFlinkJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFlinkJobResponse), nil
	}
}

type ShowFlinkJobStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFlinkJobStateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFlinkJobStateInvoker) Invoke() (*model.ShowFlinkJobStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFlinkJobStateResponse), nil
	}
}

type CreateFlinkSavepointInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFlinkSavepointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFlinkSavepointInvoker) Invoke() (*model.CreateFlinkSavepointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFlinkSavepointResponse), nil
	}
}

type DeleteFlinkSavepointInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFlinkSavepointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteFlinkSavepointInvoker) Invoke() (*model.DeleteFlinkSavepointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFlinkSavepointResponse), nil
	}
}

type ListFlinkSavepointInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlinkSavepointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFlinkSavepointInvoker) Invoke() (*model.ListFlinkSavepointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlinkSavepointResponse), nil
	}
}

type ShowFlinkSavepointInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFlinkSavepointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFlinkSavepointInvoker) Invoke() (*model.ShowFlinkSavepointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFlinkSavepointResponse), nil
	}
}

type CreateV2RayFederationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateV2RayFederationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateV2RayFederationInvoker) Invoke() (*model.CreateV2RayFederationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateV2RayFederationResponse), nil
	}
}

type DeleteV2RayFederationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteV2RayFederationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteV2RayFederationInvoker) Invoke() (*model.DeleteV2RayFederationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteV2RayFederationResponse), nil
	}
}

type ListV2RayFederationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListV2RayFederationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListV2RayFederationsInvoker) Invoke() (*model.ListV2RayFederationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListV2RayFederationsResponse), nil
	}
}

type ShowV2RayFederationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowV2RayFederationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowV2RayFederationInvoker) Invoke() (*model.ShowV2RayFederationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowV2RayFederationResponse), nil
	}
}

type UpdateV2RayFederationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateV2RayFederationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateV2RayFederationInvoker) Invoke() (*model.UpdateV2RayFederationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateV2RayFederationResponse), nil
	}
}

type CancelRayJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelRayJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelRayJobInvoker) Invoke() (*model.CancelRayJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelRayJobResponse), nil
	}
}

type ListRayJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRayJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRayJobsInvoker) Invoke() (*model.ListRayJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRayJobsResponse), nil
	}
}

type RunRayJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunRayJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunRayJobInvoker) Invoke() (*model.RunRayJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunRayJobResponse), nil
	}
}

type ShowRayJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRayJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRayJobInvoker) Invoke() (*model.ShowRayJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRayJobResponse), nil
	}
}

type CloseRaySessionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CloseRaySessionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CloseRaySessionInvoker) Invoke() (*model.CloseRaySessionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CloseRaySessionResponse), nil
	}
}

type CreateRaySessionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRaySessionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRaySessionInvoker) Invoke() (*model.CreateRaySessionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRaySessionResponse), nil
	}
}

type CancelSparkJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelSparkJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelSparkJobInvoker) Invoke() (*model.CancelSparkJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelSparkJobResponse), nil
	}
}

type ListSparkJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSparkJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSparkJobsInvoker) Invoke() (*model.ListSparkJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSparkJobsResponse), nil
	}
}

type RunSparkJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunSparkJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunSparkJobInvoker) Invoke() (*model.RunSparkJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunSparkJobResponse), nil
	}
}

type ShowSparkJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSparkJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSparkJobInvoker) Invoke() (*model.ShowSparkJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSparkJobResponse), nil
	}
}

type ShowSparkJobStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSparkJobStateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSparkJobStateInvoker) Invoke() (*model.ShowSparkJobStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSparkJobStateResponse), nil
	}
}

type ShowSparkOperatorStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSparkOperatorStateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSparkOperatorStateInvoker) Invoke() (*model.ShowSparkOperatorStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSparkOperatorStateResponse), nil
	}
}

type CancelSparkSqlInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelSparkSqlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelSparkSqlInvoker) Invoke() (*model.CancelSparkSqlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelSparkSqlResponse), nil
	}
}

type ListSparkSqlsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSparkSqlsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSparkSqlsInvoker) Invoke() (*model.ListSparkSqlsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSparkSqlsResponse), nil
	}
}

type PreviewSparkSqlResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *PreviewSparkSqlResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PreviewSparkSqlResultInvoker) Invoke() (*model.PreviewSparkSqlResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PreviewSparkSqlResultResponse), nil
	}
}

type RestartSparkSqlClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestartSparkSqlClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RestartSparkSqlClusterInvoker) Invoke() (*model.RestartSparkSqlClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartSparkSqlClusterResponse), nil
	}
}

type RunSparkSqlInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunSparkSqlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunSparkSqlInvoker) Invoke() (*model.RunSparkSqlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunSparkSqlResponse), nil
	}
}

type ShowSparkSqlInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSparkSqlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSparkSqlInvoker) Invoke() (*model.ShowSparkSqlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSparkSqlResponse), nil
	}
}

type ShowSparkSqlStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSparkSqlStateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSparkSqlStateInvoker) Invoke() (*model.ShowSparkSqlStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSparkSqlStateResponse), nil
	}
}

type BatchCreateV2ResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateV2ResourceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateV2ResourceTagsInvoker) Invoke() (*model.BatchCreateV2ResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateV2ResourceTagsResponse), nil
	}
}

type BatchDeleteV2ResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteV2ResourceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteV2ResourceTagsInvoker) Invoke() (*model.BatchDeleteV2ResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteV2ResourceTagsResponse), nil
	}
}

type CountV2TagResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountV2TagResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountV2TagResourcesInvoker) Invoke() (*model.CountV2TagResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountV2TagResourcesResponse), nil
	}
}

type ListV2ProjectTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListV2ProjectTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListV2ProjectTagsInvoker) Invoke() (*model.ListV2ProjectTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListV2ProjectTagsResponse), nil
	}
}

type ListV2ResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListV2ResourceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListV2ResourceTagsInvoker) Invoke() (*model.ListV2ResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListV2ResourceTagsResponse), nil
	}
}

type ListV2TagResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListV2TagResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListV2TagResourcesInvoker) Invoke() (*model.ListV2TagResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListV2TagResourcesResponse), nil
	}
}
