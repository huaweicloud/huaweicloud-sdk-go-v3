package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dataartsfabric/v1/model"
)

type CreateAgencyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAgencyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAgencyInvoker) Invoke() (*model.CreateAgencyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAgencyResponse), nil
	}
}

type DeleteAgencyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAgencyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAgencyInvoker) Invoke() (*model.DeleteAgencyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAgencyResponse), nil
	}
}

type ListAgencyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAgencyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAgencyInvoker) Invoke() (*model.ListAgencyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAgencyResponse), nil
	}
}

type CreateAgreementInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAgreementInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAgreementInvoker) Invoke() (*model.CreateAgreementResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAgreementResponse), nil
	}
}

type DeleteAgreementInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAgreementInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAgreementInvoker) Invoke() (*model.DeleteAgreementResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAgreementResponse), nil
	}
}

type ShowAgreementInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAgreementInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAgreementInvoker) Invoke() (*model.ShowAgreementResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAgreementResponse), nil
	}
}

type ShowAgreementRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAgreementRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAgreementRuleInvoker) Invoke() (*model.ShowAgreementRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAgreementRuleResponse), nil
	}
}

type ListFeaturesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFeaturesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFeaturesInvoker) Invoke() (*model.ListFeaturesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFeaturesResponse), nil
	}
}

type CreateEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEndpointInvoker) Invoke() (*model.CreateEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEndpointResponse), nil
	}
}

type DeleteEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEndpointInvoker) Invoke() (*model.DeleteEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEndpointResponse), nil
	}
}

type ListEndpointsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEndpointsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEndpointsInvoker) Invoke() (*model.ListEndpointsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEndpointsResponse), nil
	}
}

type ShowEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEndpointInvoker) Invoke() (*model.ShowEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEndpointResponse), nil
	}
}

type SubscribeEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *SubscribeEndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SubscribeEndpointInvoker) Invoke() (*model.SubscribeEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubscribeEndpointResponse), nil
	}
}

type UpdateEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEndpointInvoker) Invoke() (*model.UpdateEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEndpointResponse), nil
	}
}

type ShowAdminHealthCheckInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAdminHealthCheckInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAdminHealthCheckInvoker) Invoke() (*model.ShowAdminHealthCheckResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAdminHealthCheckResponse), nil
	}
}

type CreateMessageNotificationPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMessageNotificationPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMessageNotificationPolicyInvoker) Invoke() (*model.CreateMessageNotificationPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMessageNotificationPolicyResponse), nil
	}
}

type DeleteMessageNotificationPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMessageNotificationPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMessageNotificationPolicyInvoker) Invoke() (*model.DeleteMessageNotificationPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMessageNotificationPolicyResponse), nil
	}
}

type ListMessageNotificationPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMessageNotificationPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMessageNotificationPolicyInvoker) Invoke() (*model.ListMessageNotificationPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMessageNotificationPolicyResponse), nil
	}
}

type UpdateMetricsConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMetricsConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMetricsConfigInvoker) Invoke() (*model.UpdateMetricsConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMetricsConfigResponse), nil
	}
}

type CleanupModelInvoker struct {
	*invoker.BaseInvoker
}

func (i *CleanupModelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CleanupModelInvoker) Invoke() (*model.CleanupModelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CleanupModelResponse), nil
	}
}

type CreateModelDefinitionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateModelDefinitionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateModelDefinitionInvoker) Invoke() (*model.CreateModelDefinitionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateModelDefinitionResponse), nil
	}
}

type DeleteModelVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteModelVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteModelVersionInvoker) Invoke() (*model.DeleteModelVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteModelVersionResponse), nil
	}
}

type ListBaseModelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBaseModelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBaseModelsInvoker) Invoke() (*model.ListBaseModelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBaseModelsResponse), nil
	}
}

type ListModelVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListModelVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListModelVersionsInvoker) Invoke() (*model.ListModelVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListModelVersionsResponse), nil
	}
}

type ListModelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListModelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListModelsInvoker) Invoke() (*model.ListModelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListModelsResponse), nil
	}
}

type UpdateModelDefinitionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateModelDefinitionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateModelDefinitionInvoker) Invoke() (*model.UpdateModelDefinitionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateModelDefinitionResponse), nil
	}
}

type ListSpecsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSpecsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSpecsInvoker) Invoke() (*model.ListSpecsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSpecsResponse), nil
	}
}

type BatchCreateFabricWorkspaceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateFabricWorkspaceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateFabricWorkspaceTagsInvoker) Invoke() (*model.BatchCreateFabricWorkspaceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateFabricWorkspaceTagsResponse), nil
	}
}

type BatchDeleteFabricWorkspaceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteFabricWorkspaceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteFabricWorkspaceTagsInvoker) Invoke() (*model.BatchDeleteFabricWorkspaceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteFabricWorkspaceTagsResponse), nil
	}
}

type CountTagFabricWorkspacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountTagFabricWorkspacesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountTagFabricWorkspacesInvoker) Invoke() (*model.CountTagFabricWorkspacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountTagFabricWorkspacesResponse), nil
	}
}

type ListFabricProjectTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFabricProjectTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFabricProjectTagsInvoker) Invoke() (*model.ListFabricProjectTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFabricProjectTagsResponse), nil
	}
}

type ListFabricWorkspaceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFabricWorkspaceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFabricWorkspaceTagsInvoker) Invoke() (*model.ListFabricWorkspaceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFabricWorkspaceTagsResponse), nil
	}
}

type ListTagFabricWorkspacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagFabricWorkspacesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTagFabricWorkspacesInvoker) Invoke() (*model.ListTagFabricWorkspacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagFabricWorkspacesResponse), nil
	}
}

type CreateWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWorkspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWorkspaceInvoker) Invoke() (*model.CreateWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWorkspaceResponse), nil
	}
}

type DeleteWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteWorkspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteWorkspaceInvoker) Invoke() (*model.DeleteWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteWorkspaceResponse), nil
	}
}

type ListWorkspacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkspacesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkspacesInvoker) Invoke() (*model.ListWorkspacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkspacesResponse), nil
	}
}

type UpdateWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWorkspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateWorkspaceInvoker) Invoke() (*model.UpdateWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWorkspaceResponse), nil
	}
}
