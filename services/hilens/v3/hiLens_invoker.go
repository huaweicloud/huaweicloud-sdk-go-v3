package v3

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/hilens/v3/model"
)

type AddDeploymentNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddDeploymentNodesInvoker) Invoke() (*model.AddDeploymentNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddDeploymentNodesResponse), nil
	}
}

type BatchCreateNodeTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateNodeTagsInvoker) Invoke() (*model.BatchCreateNodeTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateNodeTagsResponse), nil
	}
}

type CreateConfigMapInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateConfigMapInvoker) Invoke() (*model.CreateConfigMapResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConfigMapResponse), nil
	}
}

type CreateDeploymentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDeploymentInvoker) Invoke() (*model.CreateDeploymentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDeploymentResponse), nil
	}
}

type CreateNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNodeInvoker) Invoke() (*model.CreateNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNodeResponse), nil
	}
}

type CreateOrderFormInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOrderFormInvoker) Invoke() (*model.CreateOrderFormResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOrderFormResponse), nil
	}
}

type CreateResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResourceTagsInvoker) Invoke() (*model.CreateResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResourceTagsResponse), nil
	}
}

type CreateSecretInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSecretInvoker) Invoke() (*model.CreateSecretResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSecretResponse), nil
	}
}

type CreateTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTaskInvoker) Invoke() (*model.CreateTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTaskResponse), nil
	}
}

type CreateWorkSpaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWorkSpaceInvoker) Invoke() (*model.CreateWorkSpaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWorkSpaceResponse), nil
	}
}

type DeleteConfigMapInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteConfigMapInvoker) Invoke() (*model.DeleteConfigMapResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteConfigMapResponse), nil
	}
}

type DeleteDeploymentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDeploymentInvoker) Invoke() (*model.DeleteDeploymentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDeploymentResponse), nil
	}
}

type DeleteNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNodeInvoker) Invoke() (*model.DeleteNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNodeResponse), nil
	}
}

type DeletePodInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePodInvoker) Invoke() (*model.DeletePodResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePodResponse), nil
	}
}

type DeleteResourceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteResourceTagInvoker) Invoke() (*model.DeleteResourceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteResourceTagResponse), nil
	}
}

type DeleteSecretInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSecretInvoker) Invoke() (*model.DeleteSecretResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSecretResponse), nil
	}
}

type DeleteTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTaskInvoker) Invoke() (*model.DeleteTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTaskResponse), nil
	}
}

type DeleteWorkSpaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteWorkSpaceInvoker) Invoke() (*model.DeleteWorkSpaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteWorkSpaceResponse), nil
	}
}

type FreezeNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *FreezeNodeInvoker) Invoke() (*model.FreezeNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.FreezeNodeResponse), nil
	}
}

type ListConfigMapsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConfigMapsInvoker) Invoke() (*model.ListConfigMapsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConfigMapsResponse), nil
	}
}

type ListFirmwaresInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFirmwaresInvoker) Invoke() (*model.ListFirmwaresResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFirmwaresResponse), nil
	}
}

type ListPlatformManagerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPlatformManagerInvoker) Invoke() (*model.ListPlatformManagerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPlatformManagerResponse), nil
	}
}

type ListResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceTagsInvoker) Invoke() (*model.ListResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceTagsResponse), nil
	}
}

type ListSecretsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecretsInvoker) Invoke() (*model.ListSecretsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecretsResponse), nil
	}
}

type ListTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTasksInvoker) Invoke() (*model.ListTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTasksResponse), nil
	}
}

type ListWorkSpacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkSpacesInvoker) Invoke() (*model.ListWorkSpacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkSpacesResponse), nil
	}
}

type SetDefaultOrderFormInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetDefaultOrderFormInvoker) Invoke() (*model.SetDefaultOrderFormResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetDefaultOrderFormResponse), nil
	}
}

type ShowConfigMapInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConfigMapInvoker) Invoke() (*model.ShowConfigMapResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConfigMapResponse), nil
	}
}

type ShowDeploymentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeploymentInvoker) Invoke() (*model.ShowDeploymentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeploymentResponse), nil
	}
}

type ShowDeploymentPodsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeploymentPodsInvoker) Invoke() (*model.ShowDeploymentPodsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeploymentPodsResponse), nil
	}
}

type ShowDeploymentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeploymentsInvoker) Invoke() (*model.ShowDeploymentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeploymentsResponse), nil
	}
}

type ShowNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNodeInvoker) Invoke() (*model.ShowNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNodeResponse), nil
	}
}

type ShowNodeActivationRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNodeActivationRecordsInvoker) Invoke() (*model.ShowNodeActivationRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNodeActivationRecordsResponse), nil
	}
}

type ShowNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNodesInvoker) Invoke() (*model.ShowNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNodesResponse), nil
	}
}

type ShowResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceTagsInvoker) Invoke() (*model.ShowResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceTagsResponse), nil
	}
}

type ShowSecretInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSecretInvoker) Invoke() (*model.ShowSecretResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSecretResponse), nil
	}
}

type ShowSkillInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSkillInfoInvoker) Invoke() (*model.ShowSkillInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSkillInfoResponse), nil
	}
}

type ShowSkillListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSkillListInvoker) Invoke() (*model.ShowSkillListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSkillListResponse), nil
	}
}

type ShowSkillOrderInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSkillOrderInfoInvoker) Invoke() (*model.ShowSkillOrderInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSkillOrderInfoResponse), nil
	}
}

type ShowSkillOrderListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSkillOrderListInvoker) Invoke() (*model.ShowSkillOrderListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSkillOrderListResponse), nil
	}
}

type ShowTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskInvoker) Invoke() (*model.ShowTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskResponse), nil
	}
}

type ShowUpgradeProgressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUpgradeProgressInvoker) Invoke() (*model.ShowUpgradeProgressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUpgradeProgressResponse), nil
	}
}

type ShowWorkSpaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkSpaceInvoker) Invoke() (*model.ShowWorkSpaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkSpaceResponse), nil
	}
}

type StartAndStopDeploymentInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartAndStopDeploymentInvoker) Invoke() (*model.StartAndStopDeploymentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartAndStopDeploymentResponse), nil
	}
}

type StartAndStopDeploymentPodInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartAndStopDeploymentPodInvoker) Invoke() (*model.StartAndStopDeploymentPodResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartAndStopDeploymentPodResponse), nil
	}
}

type SwitchNodeConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchNodeConnectionInvoker) Invoke() (*model.SwitchNodeConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchNodeConnectionResponse), nil
	}
}

type UnfreezeNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *UnfreezeNodeInvoker) Invoke() (*model.UnfreezeNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnfreezeNodeResponse), nil
	}
}

type UpdateConfigMapInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateConfigMapInvoker) Invoke() (*model.UpdateConfigMapResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateConfigMapResponse), nil
	}
}

type UpdateDeploymentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDeploymentInvoker) Invoke() (*model.UpdateDeploymentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDeploymentResponse), nil
	}
}

type UpdateDeploymentUsingPatchInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDeploymentUsingPatchInvoker) Invoke() (*model.UpdateDeploymentUsingPatchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDeploymentUsingPatchResponse), nil
	}
}

type UpdateNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNodeInvoker) Invoke() (*model.UpdateNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNodeResponse), nil
	}
}

type UpdateNodeCertInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNodeCertInvoker) Invoke() (*model.UpdateNodeCertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNodeCertResponse), nil
	}
}

type UpdateNodeFirmwareInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNodeFirmwareInvoker) Invoke() (*model.UpdateNodeFirmwareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNodeFirmwareResponse), nil
	}
}

type UpdateSecretInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSecretInvoker) Invoke() (*model.UpdateSecretResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSecretResponse), nil
	}
}

type UpdateTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTaskInvoker) Invoke() (*model.UpdateTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTaskResponse), nil
	}
}

type UpdateWorkSpaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWorkSpaceInvoker) Invoke() (*model.UpdateWorkSpaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWorkSpaceResponse), nil
	}
}
