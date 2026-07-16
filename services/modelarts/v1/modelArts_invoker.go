package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/modelarts/v1/model"
)

type AcceptScheduledEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *AcceptScheduledEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AcceptScheduledEventInvoker) Invoke() (*model.AcceptScheduledEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AcceptScheduledEventResponse), nil
	}
}

type AttachDynamicStorageInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachDynamicStorageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AttachDynamicStorageInvoker) Invoke() (*model.AttachDynamicStorageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachDynamicStorageResponse), nil
	}
}

type BatchBindInferApiKeysInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchBindInferApiKeysInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchBindInferApiKeysInvoker) Invoke() (*model.BatchBindInferApiKeysResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchBindInferApiKeysResponse), nil
	}
}

type BatchBindPoolNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchBindPoolNodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchBindPoolNodesInvoker) Invoke() (*model.BatchBindPoolNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchBindPoolNodesResponse), nil
	}
}

type BatchCreatePoolTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreatePoolTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreatePoolTagsInvoker) Invoke() (*model.BatchCreatePoolTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreatePoolTagsResponse), nil
	}
}

type BatchDeleteInferIntranetConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteInferIntranetConnectionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteInferIntranetConnectionsInvoker) Invoke() (*model.BatchDeleteInferIntranetConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteInferIntranetConnectionsResponse), nil
	}
}

type BatchDeleteInferServicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteInferServicesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteInferServicesInvoker) Invoke() (*model.BatchDeleteInferServicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteInferServicesResponse), nil
	}
}

type BatchDeletePoolNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeletePoolNodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeletePoolNodesInvoker) Invoke() (*model.BatchDeletePoolNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeletePoolNodesResponse), nil
	}
}

type BatchDeletePoolTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeletePoolTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeletePoolTagsInvoker) Invoke() (*model.BatchDeletePoolTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeletePoolTagsResponse), nil
	}
}

type BatchLockPoolNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchLockPoolNodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchLockPoolNodesInvoker) Invoke() (*model.BatchLockPoolNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchLockPoolNodesResponse), nil
	}
}

type BatchMigratePoolNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchMigratePoolNodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchMigratePoolNodesInvoker) Invoke() (*model.BatchMigratePoolNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchMigratePoolNodesResponse), nil
	}
}

type BatchRebootPoolNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchRebootPoolNodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchRebootPoolNodesInvoker) Invoke() (*model.BatchRebootPoolNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchRebootPoolNodesResponse), nil
	}
}

type BatchResetPoolNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchResetPoolNodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchResetPoolNodesInvoker) Invoke() (*model.BatchResetPoolNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchResetPoolNodesResponse), nil
	}
}

type BatchResizePoolNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchResizePoolNodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchResizePoolNodesInvoker) Invoke() (*model.BatchResizePoolNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchResizePoolNodesResponse), nil
	}
}

type BatchUnbindInferApiKeysInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUnbindInferApiKeysInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUnbindInferApiKeysInvoker) Invoke() (*model.BatchUnbindInferApiKeysResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUnbindInferApiKeysResponse), nil
	}
}

type BatchUnlockPoolNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUnlockPoolNodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUnlockPoolNodesInvoker) Invoke() (*model.BatchUnlockPoolNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUnlockPoolNodesResponse), nil
	}
}

type BatchUpdatePoolNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdatePoolNodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdatePoolNodesInvoker) Invoke() (*model.BatchUpdatePoolNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdatePoolNodesResponse), nil
	}
}

type BindInferApiKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *BindInferApiKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BindInferApiKeyInvoker) Invoke() (*model.BindInferApiKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BindInferApiKeyResponse), nil
	}
}

type CancelInferDeploymentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelInferDeploymentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelInferDeploymentInvoker) Invoke() (*model.CancelInferDeploymentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelInferDeploymentResponse), nil
	}
}

type ChangeAlgorithmInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeAlgorithmInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeAlgorithmInvoker) Invoke() (*model.ChangeAlgorithmResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeAlgorithmResponse), nil
	}
}

type ChangeTrainingExperimentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeTrainingExperimentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeTrainingExperimentInvoker) Invoke() (*model.ChangeTrainingExperimentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeTrainingExperimentResponse), nil
	}
}

type ChangeTrainingJobDescriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeTrainingJobDescriptionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeTrainingJobDescriptionInvoker) Invoke() (*model.ChangeTrainingJobDescriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeTrainingJobDescriptionResponse), nil
	}
}

type CheckTrainingExperimentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckTrainingExperimentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckTrainingExperimentInvoker) Invoke() (*model.CheckTrainingExperimentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckTrainingExperimentResponse), nil
	}
}

type CountInferServicesByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountInferServicesByTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountInferServicesByTagsInvoker) Invoke() (*model.CountInferServicesByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountInferServicesByTagsResponse), nil
	}
}

type CreateAlgorithmInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAlgorithmInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAlgorithmInvoker) Invoke() (*model.CreateAlgorithmResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAlgorithmResponse), nil
	}
}

type CreateAlgorithmVersionToGalleryInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAlgorithmVersionToGalleryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAlgorithmVersionToGalleryInvoker) Invoke() (*model.CreateAlgorithmVersionToGalleryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAlgorithmVersionToGalleryResponse), nil
	}
}

type CreateAuthorizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAuthorizationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAuthorizationInvoker) Invoke() (*model.CreateAuthorizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAuthorizationResponse), nil
	}
}

type CreateInferApiKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInferApiKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInferApiKeyInvoker) Invoke() (*model.CreateInferApiKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInferApiKeyResponse), nil
	}
}

type CreateInferDeploymentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInferDeploymentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInferDeploymentInvoker) Invoke() (*model.CreateInferDeploymentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInferDeploymentResponse), nil
	}
}

type CreateInferIntranetConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInferIntranetConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInferIntranetConnectionInvoker) Invoke() (*model.CreateInferIntranetConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInferIntranetConnectionResponse), nil
	}
}

type CreateInferServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInferServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInferServiceInvoker) Invoke() (*model.CreateInferServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInferServiceResponse), nil
	}
}

type CreateInferServiceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInferServiceTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInferServiceTagInvoker) Invoke() (*model.CreateInferServiceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInferServiceTagResponse), nil
	}
}

type CreateModelArtsAgencyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateModelArtsAgencyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateModelArtsAgencyInvoker) Invoke() (*model.CreateModelArtsAgencyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateModelArtsAgencyResponse), nil
	}
}

type CreateNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNetworkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateNetworkInvoker) Invoke() (*model.CreateNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNetworkResponse), nil
	}
}

type CreateNodePoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNodePoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateNodePoolInvoker) Invoke() (*model.CreateNodePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNodePoolResponse), nil
	}
}

type CreateOrderIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOrderIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateOrderIdInvoker) Invoke() (*model.CreateOrderIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOrderIdResponse), nil
	}
}

type CreatePoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePoolInvoker) Invoke() (*model.CreatePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePoolResponse), nil
	}
}

type CreatePoolPluginInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePoolPluginInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePoolPluginInvoker) Invoke() (*model.CreatePoolPluginResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePoolPluginResponse), nil
	}
}

type CreateSaveImageJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSaveImageJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSaveImageJobInvoker) Invoke() (*model.CreateSaveImageJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSaveImageJobResponse), nil
	}
}

type CreateTrainJobTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTrainJobTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTrainJobTagsInvoker) Invoke() (*model.CreateTrainJobTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTrainJobTagsResponse), nil
	}
}

type CreateTrainingExperimentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTrainingExperimentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTrainingExperimentInvoker) Invoke() (*model.CreateTrainingExperimentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTrainingExperimentResponse), nil
	}
}

type CreateTrainingJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTrainingJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTrainingJobInvoker) Invoke() (*model.CreateTrainingJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTrainingJobResponse), nil
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

type DeleteAlgorithmInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAlgorithmInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAlgorithmInvoker) Invoke() (*model.DeleteAlgorithmResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAlgorithmResponse), nil
	}
}

type DeleteAuthorizationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAuthorizationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAuthorizationsInvoker) Invoke() (*model.DeleteAuthorizationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAuthorizationsResponse), nil
	}
}

type DeleteImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteImageInvoker) Invoke() (*model.DeleteImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteImageResponse), nil
	}
}

type DeleteImageGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteImageGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteImageGroupInvoker) Invoke() (*model.DeleteImageGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteImageGroupResponse), nil
	}
}

type DeleteInferApiKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInferApiKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInferApiKeyInvoker) Invoke() (*model.DeleteInferApiKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInferApiKeyResponse), nil
	}
}

type DeleteInferDeploymentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInferDeploymentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInferDeploymentInvoker) Invoke() (*model.DeleteInferDeploymentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInferDeploymentResponse), nil
	}
}

type DeleteInferDeploymentInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInferDeploymentInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInferDeploymentInstanceInvoker) Invoke() (*model.DeleteInferDeploymentInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInferDeploymentInstanceResponse), nil
	}
}

type DeleteInferDeploymentPodInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInferDeploymentPodInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInferDeploymentPodInvoker) Invoke() (*model.DeleteInferDeploymentPodResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInferDeploymentPodResponse), nil
	}
}

type DeleteInferDeploymentVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInferDeploymentVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInferDeploymentVersionInvoker) Invoke() (*model.DeleteInferDeploymentVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInferDeploymentVersionResponse), nil
	}
}

type DeleteInferServiceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInferServiceTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInferServiceTagInvoker) Invoke() (*model.DeleteInferServiceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInferServiceTagResponse), nil
	}
}

type DeleteNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNetworkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteNetworkInvoker) Invoke() (*model.DeleteNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNetworkResponse), nil
	}
}

type DeleteNodePoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNodePoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteNodePoolInvoker) Invoke() (*model.DeleteNodePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNodePoolResponse), nil
	}
}

type DeletePoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePoolInvoker) Invoke() (*model.DeletePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePoolResponse), nil
	}
}

type DeleteTrainJobTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTrainJobTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTrainJobTagsInvoker) Invoke() (*model.DeleteTrainJobTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTrainJobTagsResponse), nil
	}
}

type DeleteTrainingExperimentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTrainingExperimentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTrainingExperimentInvoker) Invoke() (*model.DeleteTrainingExperimentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTrainingExperimentResponse), nil
	}
}

type DeleteTrainingJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTrainingJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTrainingJobInvoker) Invoke() (*model.DeleteTrainingJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTrainingJobResponse), nil
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

type DetachDynamicStorageInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetachDynamicStorageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetachDynamicStorageInvoker) Invoke() (*model.DetachDynamicStorageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetachDynamicStorageResponse), nil
	}
}

type GetAuthorizationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetAuthorizationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetAuthorizationsInvoker) Invoke() (*model.GetAuthorizationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetAuthorizationsResponse), nil
	}
}

type GetHyperinstanceOperationInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetHyperinstanceOperationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetHyperinstanceOperationInvoker) Invoke() (*model.GetHyperinstanceOperationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetHyperinstanceOperationResponse), nil
	}
}

type ListAlgorithmsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlgorithmsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAlgorithmsInvoker) Invoke() (*model.ListAlgorithmsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlgorithmsResponse), nil
	}
}

type ListDynamicStoragesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDynamicStoragesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDynamicStoragesInvoker) Invoke() (*model.ListDynamicStoragesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDynamicStoragesResponse), nil
	}
}

type ListEventCategoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEventCategoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEventCategoriesInvoker) Invoke() (*model.ListEventCategoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEventCategoriesResponse), nil
	}
}

type ListEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEventsInvoker) Invoke() (*model.ListEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEventsResponse), nil
	}
}

type ListImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageInvoker) Invoke() (*model.ListImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageResponse), nil
	}
}

type ListImageGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageGroupInvoker) Invoke() (*model.ListImageGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageGroupResponse), nil
	}
}

type ListInferApiKeysInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInferApiKeysInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInferApiKeysInvoker) Invoke() (*model.ListInferApiKeysResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInferApiKeysResponse), nil
	}
}

type ListInferClusterFlavorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInferClusterFlavorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInferClusterFlavorsInvoker) Invoke() (*model.ListInferClusterFlavorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInferClusterFlavorsResponse), nil
	}
}

type ListInferDeploymentInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInferDeploymentInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInferDeploymentInstancesInvoker) Invoke() (*model.ListInferDeploymentInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInferDeploymentInstancesResponse), nil
	}
}

type ListInferDeploymentPodEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInferDeploymentPodEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInferDeploymentPodEventsInvoker) Invoke() (*model.ListInferDeploymentPodEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInferDeploymentPodEventsResponse), nil
	}
}

type ListInferDeploymentPodsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInferDeploymentPodsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInferDeploymentPodsInvoker) Invoke() (*model.ListInferDeploymentPodsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInferDeploymentPodsResponse), nil
	}
}

type ListInferDeploymentVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInferDeploymentVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInferDeploymentVersionsInvoker) Invoke() (*model.ListInferDeploymentVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInferDeploymentVersionsResponse), nil
	}
}

type ListInferDeploymentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInferDeploymentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInferDeploymentsInvoker) Invoke() (*model.ListInferDeploymentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInferDeploymentsResponse), nil
	}
}

type ListInferIntranetConnectionApplicationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInferIntranetConnectionApplicationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInferIntranetConnectionApplicationsInvoker) Invoke() (*model.ListInferIntranetConnectionApplicationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInferIntranetConnectionApplicationsResponse), nil
	}
}

type ListInferIntranetConnectionReviewsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInferIntranetConnectionReviewsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInferIntranetConnectionReviewsInvoker) Invoke() (*model.ListInferIntranetConnectionReviewsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInferIntranetConnectionReviewsResponse), nil
	}
}

type ListInferServiceEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInferServiceEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInferServiceEventsInvoker) Invoke() (*model.ListInferServiceEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInferServiceEventsResponse), nil
	}
}

type ListInferServiceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInferServiceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInferServiceTagsInvoker) Invoke() (*model.ListInferServiceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInferServiceTagsResponse), nil
	}
}

type ListInferServicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInferServicesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInferServicesInvoker) Invoke() (*model.ListInferServicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInferServicesResponse), nil
	}
}

type ListInferServicesByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInferServicesByTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInferServicesByTagsInvoker) Invoke() (*model.ListInferServicesByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInferServicesByTagsResponse), nil
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

type ListNetworksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNetworksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNetworksInvoker) Invoke() (*model.ListNetworksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNetworksResponse), nil
	}
}

type ListNodePoolNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNodePoolNodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNodePoolNodesInvoker) Invoke() (*model.ListNodePoolNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNodePoolNodesResponse), nil
	}
}

type ListNodePoolsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNodePoolsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNodePoolsInvoker) Invoke() (*model.ListNodePoolsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNodePoolsResponse), nil
	}
}

type ListPluginTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPluginTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPluginTemplatesInvoker) Invoke() (*model.ListPluginTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPluginTemplatesResponse), nil
	}
}

type ListPoolNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPoolNodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPoolNodesInvoker) Invoke() (*model.ListPoolNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPoolNodesResponse), nil
	}
}

type ListPoolPluginsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPoolPluginsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPoolPluginsInvoker) Invoke() (*model.ListPoolPluginsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPoolPluginsResponse), nil
	}
}

type ListPoolTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPoolTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPoolTagsInvoker) Invoke() (*model.ListPoolTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPoolTagsResponse), nil
	}
}

type ListPoolsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPoolsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPoolsInvoker) Invoke() (*model.ListPoolsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPoolsResponse), nil
	}
}

type ListResourceFlavorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceFlavorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceFlavorsInvoker) Invoke() (*model.ListResourceFlavorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceFlavorsResponse), nil
	}
}

type ListScheduledEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScheduledEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScheduledEventsInvoker) Invoke() (*model.ListScheduledEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScheduledEventsResponse), nil
	}
}

type ListTrainingExperimentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTrainingExperimentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTrainingExperimentsInvoker) Invoke() (*model.ListTrainingExperimentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTrainingExperimentsResponse), nil
	}
}

type ListTrainingJobEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTrainingJobEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTrainingJobEventsInvoker) Invoke() (*model.ListTrainingJobEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTrainingJobEventsResponse), nil
	}
}

type ListTrainingJobStagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTrainingJobStagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTrainingJobStagesInvoker) Invoke() (*model.ListTrainingJobStagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTrainingJobStagesResponse), nil
	}
}

type ListTrainingJobTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTrainingJobTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTrainingJobTasksInvoker) Invoke() (*model.ListTrainingJobTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTrainingJobTasksResponse), nil
	}
}

type ListTrainingJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTrainingJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTrainingJobsInvoker) Invoke() (*model.ListTrainingJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTrainingJobsResponse), nil
	}
}

type ListWorkloadsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkloadsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkloadsInvoker) Invoke() (*model.ListWorkloadsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkloadsResponse), nil
	}
}

type ListWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkspaceInvoker) Invoke() (*model.ListWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkspaceResponse), nil
	}
}

type ModifyInferIntranetConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyInferIntranetConnectionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyInferIntranetConnectionsInvoker) Invoke() (*model.ModifyInferIntranetConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyInferIntranetConnectionsResponse), nil
	}
}

type NotifyTrainingJobInformationInvoker struct {
	*invoker.BaseInvoker
}

func (i *NotifyTrainingJobInformationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *NotifyTrainingJobInformationInvoker) Invoke() (*model.NotifyTrainingJobInformationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NotifyTrainingJobInformationResponse), nil
	}
}

type PatchNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *PatchNetworkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PatchNetworkInvoker) Invoke() (*model.PatchNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PatchNetworkResponse), nil
	}
}

type PatchNodePoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *PatchNodePoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PatchNodePoolInvoker) Invoke() (*model.PatchNodePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PatchNodePoolResponse), nil
	}
}

type PatchPoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *PatchPoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PatchPoolInvoker) Invoke() (*model.PatchPoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PatchPoolResponse), nil
	}
}

type RegisterImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *RegisterImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RegisterImageInvoker) Invoke() (*model.RegisterImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterImageResponse), nil
	}
}

type ShowAlgorithmByUuidInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAlgorithmByUuidInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAlgorithmByUuidInvoker) Invoke() (*model.ShowAlgorithmByUuidResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAlgorithmByUuidResponse), nil
	}
}

type ShowAuthmodeDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuthmodeDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuthmodeDetailInvoker) Invoke() (*model.ShowAuthmodeDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuthmodeDetailResponse), nil
	}
}

type ShowAutoSearchParamAnalysisResultPathInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAutoSearchParamAnalysisResultPathInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAutoSearchParamAnalysisResultPathInvoker) Invoke() (*model.ShowAutoSearchParamAnalysisResultPathResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAutoSearchParamAnalysisResultPathResponse), nil
	}
}

type ShowAutoSearchParamsAnalysisInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAutoSearchParamsAnalysisInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAutoSearchParamsAnalysisInvoker) Invoke() (*model.ShowAutoSearchParamsAnalysisResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAutoSearchParamsAnalysisResponse), nil
	}
}

type ShowAutoSearchPerTrialInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAutoSearchPerTrialInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAutoSearchPerTrialInvoker) Invoke() (*model.ShowAutoSearchPerTrialResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAutoSearchPerTrialResponse), nil
	}
}

type ShowAutoSearchTrialEarlyStopInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAutoSearchTrialEarlyStopInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAutoSearchTrialEarlyStopInvoker) Invoke() (*model.ShowAutoSearchTrialEarlyStopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAutoSearchTrialEarlyStopResponse), nil
	}
}

type ShowAutoSearchTrialsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAutoSearchTrialsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAutoSearchTrialsInvoker) Invoke() (*model.ShowAutoSearchTrialsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAutoSearchTrialsResponse), nil
	}
}

type ShowAutoSearchYamlTemplateContentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAutoSearchYamlTemplateContentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAutoSearchYamlTemplateContentInvoker) Invoke() (*model.ShowAutoSearchYamlTemplateContentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAutoSearchYamlTemplateContentResponse), nil
	}
}

type ShowAutoSearchYamlTemplatesInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAutoSearchYamlTemplatesInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAutoSearchYamlTemplatesInfoInvoker) Invoke() (*model.ShowAutoSearchYamlTemplatesInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAutoSearchYamlTemplatesInfoResponse), nil
	}
}

type ShowDynamicStorageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDynamicStorageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDynamicStorageInvoker) Invoke() (*model.ShowDynamicStorageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDynamicStorageResponse), nil
	}
}

type ShowImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowImageInvoker) Invoke() (*model.ShowImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowImageResponse), nil
	}
}

type ShowInferDeploymentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInferDeploymentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInferDeploymentInvoker) Invoke() (*model.ShowInferDeploymentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInferDeploymentResponse), nil
	}
}

type ShowInferDeploymentVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInferDeploymentVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInferDeploymentVersionInvoker) Invoke() (*model.ShowInferDeploymentVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInferDeploymentVersionResponse), nil
	}
}

type ShowInferServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInferServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInferServiceInvoker) Invoke() (*model.ShowInferServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInferServiceResponse), nil
	}
}

type ShowInferServiceClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInferServiceClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInferServiceClusterInvoker) Invoke() (*model.ShowInferServiceClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInferServiceClusterResponse), nil
	}
}

type ShowInferServiceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInferServiceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInferServiceTagsInvoker) Invoke() (*model.ShowInferServiceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInferServiceTagsResponse), nil
	}
}

type ShowNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNetworkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNetworkInvoker) Invoke() (*model.ShowNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNetworkResponse), nil
	}
}

type ShowNetworkAvailableIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNetworkAvailableIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNetworkAvailableIpInvoker) Invoke() (*model.ShowNetworkAvailableIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNetworkAvailableIpResponse), nil
	}
}

type ShowNodeConfigTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNodeConfigTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNodeConfigTemplateInvoker) Invoke() (*model.ShowNodeConfigTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNodeConfigTemplateResponse), nil
	}
}

type ShowNodePoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNodePoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNodePoolInvoker) Invoke() (*model.ShowNodePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNodePoolResponse), nil
	}
}

type ShowObsUrlOfTrainingJobLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowObsUrlOfTrainingJobLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowObsUrlOfTrainingJobLogsInvoker) Invoke() (*model.ShowObsUrlOfTrainingJobLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowObsUrlOfTrainingJobLogsResponse), nil
	}
}

type ShowOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowOrderInvoker) Invoke() (*model.ShowOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOrderResponse), nil
	}
}

type ShowOsConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOsConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowOsConfigInvoker) Invoke() (*model.ShowOsConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOsConfigResponse), nil
	}
}

type ShowOsQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOsQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowOsQuotaInvoker) Invoke() (*model.ShowOsQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOsQuotaResponse), nil
	}
}

type ShowPluginTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPluginTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPluginTemplateInvoker) Invoke() (*model.ShowPluginTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPluginTemplateResponse), nil
	}
}

type ShowPoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPoolInvoker) Invoke() (*model.ShowPoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPoolResponse), nil
	}
}

type ShowPoolMonitorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPoolMonitorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPoolMonitorInvoker) Invoke() (*model.ShowPoolMonitorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPoolMonitorResponse), nil
	}
}

type ShowPoolNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPoolNodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPoolNodeInvoker) Invoke() (*model.ShowPoolNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPoolNodeResponse), nil
	}
}

type ShowPoolNodeConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPoolNodeConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPoolNodeConfigInvoker) Invoke() (*model.ShowPoolNodeConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPoolNodeConfigResponse), nil
	}
}

type ShowPoolNodeConfigTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPoolNodeConfigTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPoolNodeConfigTemplateInvoker) Invoke() (*model.ShowPoolNodeConfigTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPoolNodeConfigTemplateResponse), nil
	}
}

type ShowPoolRuntimeMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPoolRuntimeMetricsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPoolRuntimeMetricsInvoker) Invoke() (*model.ShowPoolRuntimeMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPoolRuntimeMetricsResponse), nil
	}
}

type ShowPoolStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPoolStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPoolStatisticsInvoker) Invoke() (*model.ShowPoolStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPoolStatisticsResponse), nil
	}
}

type ShowPoolTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPoolTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPoolTagsInvoker) Invoke() (*model.ShowPoolTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPoolTagsResponse), nil
	}
}

type ShowSaveImageJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSaveImageJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSaveImageJobInvoker) Invoke() (*model.ShowSaveImageJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSaveImageJobResponse), nil
	}
}

type ShowSearchAlgorithmsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSearchAlgorithmsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSearchAlgorithmsInvoker) Invoke() (*model.ShowSearchAlgorithmsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSearchAlgorithmsResponse), nil
	}
}

type ShowTrainJobTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTrainJobTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTrainJobTagsInvoker) Invoke() (*model.ShowTrainJobTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTrainJobTagsResponse), nil
	}
}

type ShowTrainingExperimentDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTrainingExperimentDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTrainingExperimentDetailsInvoker) Invoke() (*model.ShowTrainingExperimentDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTrainingExperimentDetailsResponse), nil
	}
}

type ShowTrainingJobDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTrainingJobDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTrainingJobDetailsInvoker) Invoke() (*model.ShowTrainingJobDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTrainingJobDetailsResponse), nil
	}
}

type ShowTrainingJobEnginesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTrainingJobEnginesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTrainingJobEnginesInvoker) Invoke() (*model.ShowTrainingJobEnginesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTrainingJobEnginesResponse), nil
	}
}

type ShowTrainingJobFlavorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTrainingJobFlavorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTrainingJobFlavorsInvoker) Invoke() (*model.ShowTrainingJobFlavorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTrainingJobFlavorsResponse), nil
	}
}

type ShowTrainingJobLogsPreviewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTrainingJobLogsPreviewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTrainingJobLogsPreviewInvoker) Invoke() (*model.ShowTrainingJobLogsPreviewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTrainingJobLogsPreviewResponse), nil
	}
}

type ShowTrainingJobMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTrainingJobMetricsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTrainingJobMetricsInvoker) Invoke() (*model.ShowTrainingJobMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTrainingJobMetricsResponse), nil
	}
}

type ShowTrainingQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTrainingQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTrainingQuotasInvoker) Invoke() (*model.ShowTrainingQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTrainingQuotasResponse), nil
	}
}

type ShowWorkloadStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkloadStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWorkloadStatisticsInvoker) Invoke() (*model.ShowWorkloadStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkloadStatisticsResponse), nil
	}
}

type ShowWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWorkspaceInvoker) Invoke() (*model.ShowWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkspaceResponse), nil
	}
}

type ShowWorkspaceQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkspaceQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWorkspaceQuotasInvoker) Invoke() (*model.ShowWorkspaceQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkspaceQuotasResponse), nil
	}
}

type StartInferDeploymentInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartInferDeploymentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartInferDeploymentInvoker) Invoke() (*model.StartInferDeploymentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartInferDeploymentResponse), nil
	}
}

type StartInferServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartInferServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartInferServiceInvoker) Invoke() (*model.StartInferServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartInferServiceResponse), nil
	}
}

type StopInferDeploymentInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopInferDeploymentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopInferDeploymentInvoker) Invoke() (*model.StopInferDeploymentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopInferDeploymentResponse), nil
	}
}

type StopInferServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopInferServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopInferServiceInvoker) Invoke() (*model.StopInferServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopInferServiceResponse), nil
	}
}

type StopTrainingJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopTrainingJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopTrainingJobInvoker) Invoke() (*model.StopTrainingJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopTrainingJobResponse), nil
	}
}

type SwitchInferDeploymentVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchInferDeploymentVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchInferDeploymentVersionInvoker) Invoke() (*model.SwitchInferDeploymentVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchInferDeploymentVersionResponse), nil
	}
}

type SyncImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *SyncImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SyncImageInvoker) Invoke() (*model.SyncImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SyncImageResponse), nil
	}
}

type UnbindInferApiKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UnbindInferApiKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UnbindInferApiKeyInvoker) Invoke() (*model.UnbindInferApiKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnbindInferApiKeyResponse), nil
	}
}

type UpdateAuthModeInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAuthModeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAuthModeInvoker) Invoke() (*model.UpdateAuthModeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAuthModeResponse), nil
	}
}

type UpdateImageGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateImageGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateImageGroupInvoker) Invoke() (*model.UpdateImageGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateImageGroupResponse), nil
	}
}

type UpdateInferDeploymentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInferDeploymentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInferDeploymentInvoker) Invoke() (*model.UpdateInferDeploymentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInferDeploymentResponse), nil
	}
}

type UpdateInferDeploymentScaleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInferDeploymentScaleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInferDeploymentScaleInvoker) Invoke() (*model.UpdateInferDeploymentScaleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInferDeploymentScaleResponse), nil
	}
}

type UpdateInferIntranetConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInferIntranetConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInferIntranetConnectionInvoker) Invoke() (*model.UpdateInferIntranetConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInferIntranetConnectionResponse), nil
	}
}

type UpdateInferServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInferServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInferServiceInvoker) Invoke() (*model.UpdateInferServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInferServiceResponse), nil
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

type UpdateWorkspaceQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWorkspaceQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateWorkspaceQuotasInvoker) Invoke() (*model.UpdateWorkspaceQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWorkspaceQuotasResponse), nil
	}
}

type ValidateAuthorizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ValidateAuthorizationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ValidateAuthorizationInvoker) Invoke() (*model.ValidateAuthorizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ValidateAuthorizationResponse), nil
	}
}

type CreateInferDeploymentHpaInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInferDeploymentHpaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInferDeploymentHpaInvoker) Invoke() (*model.CreateInferDeploymentHpaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInferDeploymentHpaResponse), nil
	}
}

type DeleteInferDeploymentHpaInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInferDeploymentHpaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInferDeploymentHpaInvoker) Invoke() (*model.DeleteInferDeploymentHpaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInferDeploymentHpaResponse), nil
	}
}

type ListInferDeploymentHpaEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInferDeploymentHpaEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInferDeploymentHpaEventsInvoker) Invoke() (*model.ListInferDeploymentHpaEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInferDeploymentHpaEventsResponse), nil
	}
}

type ShowInferDeploymentHpaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInferDeploymentHpaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInferDeploymentHpaInvoker) Invoke() (*model.ShowInferDeploymentHpaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInferDeploymentHpaResponse), nil
	}
}

type UpdateInferDeploymentHpaInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInferDeploymentHpaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInferDeploymentHpaInvoker) Invoke() (*model.UpdateInferDeploymentHpaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInferDeploymentHpaResponse), nil
	}
}

type CreateInferHraInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInferHraInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInferHraInvoker) Invoke() (*model.CreateInferHraResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInferHraResponse), nil
	}
}

type ShowInferHraInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInferHraInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInferHraInvoker) Invoke() (*model.ShowInferHraResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInferHraResponse), nil
	}
}

type UpdateInferHraInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInferHraInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInferHraInvoker) Invoke() (*model.UpdateInferHraResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInferHraResponse), nil
	}
}

type AttachDevServerVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachDevServerVolumeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AttachDevServerVolumeInvoker) Invoke() (*model.AttachDevServerVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachDevServerVolumeResponse), nil
	}
}

type BatchDevServersActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDevServersActionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDevServersActionInvoker) Invoke() (*model.BatchDevServersActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDevServersActionResponse), nil
	}
}

type BindDevServerPublicIPInvoker struct {
	*invoker.BaseInvoker
}

func (i *BindDevServerPublicIPInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BindDevServerPublicIPInvoker) Invoke() (*model.BindDevServerPublicIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BindDevServerPublicIpResponse), nil
	}
}

type ChangeDevServerOSInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeDevServerOSInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeDevServerOSInvoker) Invoke() (*model.ChangeDevServerOsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeDevServerOsResponse), nil
	}
}

type ChangeHyperinstanceOSInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeHyperinstanceOSInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeHyperinstanceOSInvoker) Invoke() (*model.ChangeHyperinstanceOsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeHyperinstanceOsResponse), nil
	}
}

type CreateDevServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDevServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDevServerInvoker) Invoke() (*model.CreateDevServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDevServerResponse), nil
	}
}

type CreateDevServerJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDevServerJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDevServerJobInvoker) Invoke() (*model.CreateDevServerJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDevServerJobResponse), nil
	}
}

type CreateHyperClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHyperClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateHyperClusterInvoker) Invoke() (*model.CreateHyperClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHyperClusterResponse), nil
	}
}

type CreateHyperinstanceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHyperinstanceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateHyperinstanceTagsInvoker) Invoke() (*model.CreateHyperinstanceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHyperinstanceTagsResponse), nil
	}
}

type CreateRoceNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRoceNetworkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRoceNetworkInvoker) Invoke() (*model.CreateRoceNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRoceNetworkResponse), nil
	}
}

type DeleteDevServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDevServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDevServerInvoker) Invoke() (*model.DeleteDevServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDevServerResponse), nil
	}
}

type DeleteDevServerJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDevServerJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDevServerJobsInvoker) Invoke() (*model.DeleteDevServerJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDevServerJobsResponse), nil
	}
}

type DeleteHyperClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHyperClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteHyperClusterInvoker) Invoke() (*model.DeleteHyperClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHyperClusterResponse), nil
	}
}

type DeleteHyperinstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHyperinstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteHyperinstanceInvoker) Invoke() (*model.DeleteHyperinstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHyperinstanceResponse), nil
	}
}

type DeleteHyperinstanceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHyperinstanceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteHyperinstanceTagsInvoker) Invoke() (*model.DeleteHyperinstanceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHyperinstanceTagsResponse), nil
	}
}

type DetachDevServerVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetachDevServerVolumeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetachDevServerVolumeInvoker) Invoke() (*model.DetachDevServerVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetachDevServerVolumeResponse), nil
	}
}

type GetDevServerImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetDevServerImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetDevServerImageInvoker) Invoke() (*model.GetDevServerImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetDevServerImageResponse), nil
	}
}

type GetDevServerJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetDevServerJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetDevServerJobInvoker) Invoke() (*model.GetDevServerJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetDevServerJobResponse), nil
	}
}

type GetDevServerJobServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetDevServerJobServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetDevServerJobServiceInvoker) Invoke() (*model.GetDevServerJobServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetDevServerJobServiceResponse), nil
	}
}

type GetDevServerJobTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetDevServerJobTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetDevServerJobTemplateInvoker) Invoke() (*model.GetDevServerJobTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetDevServerJobTemplateResponse), nil
	}
}

type GetDevServerOperationInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetDevServerOperationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetDevServerOperationInvoker) Invoke() (*model.GetDevServerOperationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetDevServerOperationResponse), nil
	}
}

type GetHyperClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetHyperClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetHyperClusterInvoker) Invoke() (*model.GetHyperClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetHyperClusterResponse), nil
	}
}

type GetHyperinstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetHyperinstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetHyperinstanceInvoker) Invoke() (*model.GetHyperinstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetHyperinstanceResponse), nil
	}
}

type GetScaleEvaluationsDevServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetScaleEvaluationsDevServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetScaleEvaluationsDevServerInvoker) Invoke() (*model.GetScaleEvaluationsDevServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetScaleEvaluationsDevServerResponse), nil
	}
}

type GetTopologiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetTopologiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetTopologiesInvoker) Invoke() (*model.GetTopologiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetTopologiesResponse), nil
	}
}

type ListAllDevServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllDevServersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAllDevServersInvoker) Invoke() (*model.ListAllDevServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllDevServersResponse), nil
	}
}

type ListAllHyperinstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllHyperinstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAllHyperinstancesInvoker) Invoke() (*model.ListAllHyperinstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllHyperinstancesResponse), nil
	}
}

type ListDevServerFlavorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDevServerFlavorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDevServerFlavorsInvoker) Invoke() (*model.ListDevServerFlavorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDevServerFlavorsResponse), nil
	}
}

type ListDevServerImagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDevServerImagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDevServerImagesInvoker) Invoke() (*model.ListDevServerImagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDevServerImagesResponse), nil
	}
}

type ListDevServerJobTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDevServerJobTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDevServerJobTemplatesInvoker) Invoke() (*model.ListDevServerJobTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDevServerJobTemplatesResponse), nil
	}
}

type ListDevServerJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDevServerJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDevServerJobsInvoker) Invoke() (*model.ListDevServerJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDevServerJobsResponse), nil
	}
}

type ListDevServerPublicIPInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDevServerPublicIPInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDevServerPublicIPInvoker) Invoke() (*model.ListDevServerPublicIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDevServerPublicIpResponse), nil
	}
}

type ListDevServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDevServersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDevServersInvoker) Invoke() (*model.ListDevServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDevServersResponse), nil
	}
}

type ListHyperClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHyperClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHyperClusterInvoker) Invoke() (*model.ListHyperClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHyperClusterResponse), nil
	}
}

type ListHyperinstanceClustersCapacityInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHyperinstanceClustersCapacityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHyperinstanceClustersCapacityInvoker) Invoke() (*model.ListHyperinstanceClustersCapacityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHyperinstanceClustersCapacityResponse), nil
	}
}

type ListHyperinstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHyperinstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHyperinstancesInvoker) Invoke() (*model.ListHyperinstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHyperinstancesResponse), nil
	}
}

type QueryHyperinstanceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *QueryHyperinstanceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *QueryHyperinstanceTagsInvoker) Invoke() (*model.QueryHyperinstanceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.QueryHyperinstanceTagsResponse), nil
	}
}

type RebootDevServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *RebootDevServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RebootDevServerInvoker) Invoke() (*model.RebootDevServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RebootDevServerResponse), nil
	}
}

type ReinstallDevServerOSInvoker struct {
	*invoker.BaseInvoker
}

func (i *ReinstallDevServerOSInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ReinstallDevServerOSInvoker) Invoke() (*model.ReinstallDevServerOsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReinstallDevServerOsResponse), nil
	}
}

type ScaleDownHyperinstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ScaleDownHyperinstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ScaleDownHyperinstanceInvoker) Invoke() (*model.ScaleDownHyperinstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ScaleDownHyperinstanceResponse), nil
	}
}

type ScaleUpHyperinstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ScaleUpHyperinstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ScaleUpHyperinstanceInvoker) Invoke() (*model.ScaleUpHyperinstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ScaleUpHyperinstanceResponse), nil
	}
}

type ShowDevServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDevServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDevServerInvoker) Invoke() (*model.ShowDevServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDevServerResponse), nil
	}
}

type StartDevServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartDevServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartDevServerInvoker) Invoke() (*model.StartDevServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartDevServerResponse), nil
	}
}

type StartHyperinstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartHyperinstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartHyperinstanceInvoker) Invoke() (*model.StartHyperinstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartHyperinstanceResponse), nil
	}
}

type StopDevServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopDevServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopDevServerInvoker) Invoke() (*model.StopDevServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopDevServerResponse), nil
	}
}

type StopHyperinstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopHyperinstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopHyperinstanceInvoker) Invoke() (*model.StopHyperinstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopHyperinstanceResponse), nil
	}
}

type SyncDevServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *SyncDevServersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SyncDevServersInvoker) Invoke() (*model.SyncDevServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SyncDevServersResponse), nil
	}
}

type UpdateDevServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDevServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDevServerInvoker) Invoke() (*model.UpdateDevServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDevServerResponse), nil
	}
}

type CreateImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateImageInvoker) Invoke() (*model.CreateImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateImageResponse), nil
	}
}

type CreateNotebookInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNotebookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateNotebookInvoker) Invoke() (*model.CreateNotebookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNotebookResponse), nil
	}
}

type CreateNotebookTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNotebookTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateNotebookTagsInvoker) Invoke() (*model.CreateNotebookTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNotebookTagsResponse), nil
	}
}

type DeleteNotebookInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNotebookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteNotebookInvoker) Invoke() (*model.DeleteNotebookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNotebookResponse), nil
	}
}

type DeleteNotebookTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNotebookTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteNotebookTagsInvoker) Invoke() (*model.DeleteNotebookTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNotebookTagsResponse), nil
	}
}

type ListAllNotebooksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllNotebooksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAllNotebooksInvoker) Invoke() (*model.ListAllNotebooksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllNotebooksResponse), nil
	}
}

type ListAuthoringClustersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuthoringClustersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuthoringClustersInvoker) Invoke() (*model.ListAuthoringClustersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuthoringClustersResponse), nil
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

type ListFlavorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlavorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFlavorsInvoker) Invoke() (*model.ListFlavorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlavorsResponse), nil
	}
}

type ListNotebooksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNotebooksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNotebooksInvoker) Invoke() (*model.ListNotebooksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNotebooksResponse), nil
	}
}

type RenewLeaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *RenewLeaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RenewLeaseInvoker) Invoke() (*model.RenewLeaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RenewLeaseResponse), nil
	}
}

type ShowClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowClusterInvoker) Invoke() (*model.ShowClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClusterResponse), nil
	}
}

type ShowLeaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLeaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLeaseInvoker) Invoke() (*model.ShowLeaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLeaseResponse), nil
	}
}

type ShowNotebookInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNotebookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNotebookInvoker) Invoke() (*model.ShowNotebookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNotebookResponse), nil
	}
}

type ShowNotebookTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNotebookTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNotebookTagsInvoker) Invoke() (*model.ShowNotebookTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNotebookTagsResponse), nil
	}
}

type ShowSwitchableFlavorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSwitchableFlavorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSwitchableFlavorsInvoker) Invoke() (*model.ShowSwitchableFlavorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSwitchableFlavorsResponse), nil
	}
}

type StartNotebookInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartNotebookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartNotebookInvoker) Invoke() (*model.StartNotebookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartNotebookResponse), nil
	}
}

type StopNotebookInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopNotebookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopNotebookInvoker) Invoke() (*model.StopNotebookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopNotebookResponse), nil
	}
}

type UpdateNotebookInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNotebookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateNotebookInvoker) Invoke() (*model.UpdateNotebookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNotebookResponse), nil
	}
}

type CreateWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWorkflowInvoker) Invoke() (*model.CreateWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWorkflowResponse), nil
	}
}

type CreateWorkflowPurchasePoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWorkflowPurchasePoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWorkflowPurchasePoolInvoker) Invoke() (*model.CreateWorkflowPurchasePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWorkflowPurchasePoolResponse), nil
	}
}

type CreateWorkflowServiceAuthInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWorkflowServiceAuthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWorkflowServiceAuthInvoker) Invoke() (*model.CreateWorkflowServiceAuthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWorkflowServiceAuthResponse), nil
	}
}

type DeleteWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteWorkflowInvoker) Invoke() (*model.DeleteWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteWorkflowResponse), nil
	}
}

type ListWorkflowsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkflowsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkflowsInvoker) Invoke() (*model.ListWorkflowsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkflowsResponse), nil
	}
}

type ShowWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWorkflowInvoker) Invoke() (*model.ShowWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkflowResponse), nil
	}
}

type ShowWorkflowLabelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkflowLabelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWorkflowLabelsInvoker) Invoke() (*model.ShowWorkflowLabelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkflowLabelsResponse), nil
	}
}

type ShowWorkflowsOverviewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkflowsOverviewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWorkflowsOverviewInvoker) Invoke() (*model.ShowWorkflowsOverviewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkflowsOverviewResponse), nil
	}
}

type ShowWorkflowsTodolistInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkflowsTodolistInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWorkflowsTodolistInvoker) Invoke() (*model.ShowWorkflowsTodolistResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkflowsTodolistResponse), nil
	}
}

type UpdateWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateWorkflowInvoker) Invoke() (*model.UpdateWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWorkflowResponse), nil
	}
}

type CreateWorkflowExecutionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWorkflowExecutionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWorkflowExecutionInvoker) Invoke() (*model.CreateWorkflowExecutionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWorkflowExecutionResponse), nil
	}
}

type CreateWorkflowExecutionsActionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWorkflowExecutionsActionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWorkflowExecutionsActionsInvoker) Invoke() (*model.CreateWorkflowExecutionsActionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWorkflowExecutionsActionsResponse), nil
	}
}

type CreateWorkflowStepExecutionsActionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWorkflowStepExecutionsActionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWorkflowStepExecutionsActionsInvoker) Invoke() (*model.CreateWorkflowStepExecutionsActionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWorkflowStepExecutionsActionsResponse), nil
	}
}

type DeleteWorkflowExecutionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteWorkflowExecutionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteWorkflowExecutionInvoker) Invoke() (*model.DeleteWorkflowExecutionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteWorkflowExecutionResponse), nil
	}
}

type ListExecutionLabelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListExecutionLabelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListExecutionLabelsInvoker) Invoke() (*model.ListExecutionLabelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListExecutionLabelsResponse), nil
	}
}

type ListWorkflowExecutionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkflowExecutionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkflowExecutionsInvoker) Invoke() (*model.ListWorkflowExecutionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkflowExecutionsResponse), nil
	}
}

type ListWorkflowStepExecutionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkflowStepExecutionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkflowStepExecutionInvoker) Invoke() (*model.ListWorkflowStepExecutionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkflowStepExecutionResponse), nil
	}
}

type ShowWorkflowExecutionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkflowExecutionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWorkflowExecutionInvoker) Invoke() (*model.ShowWorkflowExecutionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkflowExecutionResponse), nil
	}
}

type ShowWorkflowStepExecutionMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkflowStepExecutionMetricsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWorkflowStepExecutionMetricsInvoker) Invoke() (*model.ShowWorkflowStepExecutionMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkflowStepExecutionMetricsResponse), nil
	}
}

type UpdateWorkflowExecutionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWorkflowExecutionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateWorkflowExecutionInvoker) Invoke() (*model.UpdateWorkflowExecutionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWorkflowExecutionResponse), nil
	}
}

type CreateWorkflowScheduleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWorkflowScheduleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWorkflowScheduleInvoker) Invoke() (*model.CreateWorkflowScheduleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWorkflowScheduleResponse), nil
	}
}

type DeleteWorkflowScheduleIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteWorkflowScheduleIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteWorkflowScheduleIdInvoker) Invoke() (*model.DeleteWorkflowScheduleIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteWorkflowScheduleIdResponse), nil
	}
}

type ShowWorkflowScheduleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkflowScheduleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWorkflowScheduleInvoker) Invoke() (*model.ShowWorkflowScheduleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkflowScheduleResponse), nil
	}
}

type ShowWorkflowScheduleListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkflowScheduleListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWorkflowScheduleListInvoker) Invoke() (*model.ShowWorkflowScheduleListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkflowScheduleListResponse), nil
	}
}

type UpdateWorkflowScheduleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWorkflowScheduleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateWorkflowScheduleInvoker) Invoke() (*model.UpdateWorkflowScheduleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWorkflowScheduleResponse), nil
	}
}

type CreateWorkflowSubscriptionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWorkflowSubscriptionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWorkflowSubscriptionsInvoker) Invoke() (*model.CreateWorkflowSubscriptionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWorkflowSubscriptionsResponse), nil
	}
}

type DeleteWorkflowSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteWorkflowSubscriptionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteWorkflowSubscriptionInvoker) Invoke() (*model.DeleteWorkflowSubscriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteWorkflowSubscriptionResponse), nil
	}
}

type ShowWorkflowSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkflowSubscriptionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWorkflowSubscriptionInvoker) Invoke() (*model.ShowWorkflowSubscriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkflowSubscriptionResponse), nil
	}
}

type UpdateWorkflowSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWorkflowSubscriptionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateWorkflowSubscriptionInvoker) Invoke() (*model.UpdateWorkflowSubscriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWorkflowSubscriptionResponse), nil
	}
}
