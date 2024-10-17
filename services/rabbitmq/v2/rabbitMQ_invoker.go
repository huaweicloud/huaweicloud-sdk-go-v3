package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rabbitmq/v2/model"
)

type BatchCreateOrDeleteRabbitMqTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateOrDeleteRabbitMqTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateOrDeleteRabbitMqTagInvoker) Invoke() (*model.BatchCreateOrDeleteRabbitMqTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateOrDeleteRabbitMqTagResponse), nil
	}
}

type BatchRestartOrDeleteInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchRestartOrDeleteInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchRestartOrDeleteInstancesInvoker) Invoke() (*model.BatchRestartOrDeleteInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchRestartOrDeleteInstancesResponse), nil
	}
}

type CreatePostPaidInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePostPaidInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePostPaidInstanceInvoker) Invoke() (*model.CreatePostPaidInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePostPaidInstanceResponse), nil
	}
}

type CreatePostPaidInstanceByEngineInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePostPaidInstanceByEngineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePostPaidInstanceByEngineInvoker) Invoke() (*model.CreatePostPaidInstanceByEngineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePostPaidInstanceByEngineResponse), nil
	}
}

type CreateUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateUserInvoker) Invoke() (*model.CreateUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateUserResponse), nil
	}
}

type DeleteBackgroundTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBackgroundTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteBackgroundTaskInvoker) Invoke() (*model.DeleteBackgroundTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBackgroundTaskResponse), nil
	}
}

type DeleteInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstanceInvoker) Invoke() (*model.DeleteInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceResponse), nil
	}
}

type DeleteUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteUserInvoker) Invoke() (*model.DeleteUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteUserResponse), nil
	}
}

type ListAvailableZonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAvailableZonesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAvailableZonesInvoker) Invoke() (*model.ListAvailableZonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAvailableZonesResponse), nil
	}
}

type ListBackgroundTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBackgroundTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBackgroundTasksInvoker) Invoke() (*model.ListBackgroundTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBackgroundTasksResponse), nil
	}
}

type ListEngineProductsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEngineProductsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEngineProductsInvoker) Invoke() (*model.ListEngineProductsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEngineProductsResponse), nil
	}
}

type ListInstancesDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstancesDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstancesDetailsInvoker) Invoke() (*model.ListInstancesDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstancesDetailsResponse), nil
	}
}

type ListPluginsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPluginsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPluginsInvoker) Invoke() (*model.ListPluginsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPluginsResponse), nil
	}
}

type ListProductsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProductsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProductsInvoker) Invoke() (*model.ListProductsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProductsResponse), nil
	}
}

type ListUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUserInvoker) Invoke() (*model.ListUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserResponse), nil
	}
}

type ResetPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetPasswordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetPasswordInvoker) Invoke() (*model.ResetPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetPasswordResponse), nil
	}
}

type ResizeEngineInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeEngineInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResizeEngineInstanceInvoker) Invoke() (*model.ResizeEngineInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeEngineInstanceResponse), nil
	}
}

type ResizeInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResizeInstanceInvoker) Invoke() (*model.ResizeInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeInstanceResponse), nil
	}
}

type ShowBackgroundTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBackgroundTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBackgroundTaskInvoker) Invoke() (*model.ShowBackgroundTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBackgroundTaskResponse), nil
	}
}

type ShowCesHierarchyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCesHierarchyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCesHierarchyInvoker) Invoke() (*model.ShowCesHierarchyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCesHierarchyResponse), nil
	}
}

type ShowEngineInstanceExtendProductInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEngineInstanceExtendProductInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEngineInstanceExtendProductInfoInvoker) Invoke() (*model.ShowEngineInstanceExtendProductInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEngineInstanceExtendProductInfoResponse), nil
	}
}

type ShowInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceInvoker) Invoke() (*model.ShowInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceResponse), nil
	}
}

type ShowInstanceExtendProductInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceExtendProductInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceExtendProductInfoInvoker) Invoke() (*model.ShowInstanceExtendProductInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceExtendProductInfoResponse), nil
	}
}

type ShowMaintainWindowsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMaintainWindowsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMaintainWindowsInvoker) Invoke() (*model.ShowMaintainWindowsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMaintainWindowsResponse), nil
	}
}

type ShowRabbitMqProjectTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRabbitMqProjectTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRabbitMqProjectTagsInvoker) Invoke() (*model.ShowRabbitMqProjectTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRabbitMqProjectTagsResponse), nil
	}
}

type ShowRabbitMqTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRabbitMqTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRabbitMqTagsInvoker) Invoke() (*model.ShowRabbitMqTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRabbitMqTagsResponse), nil
	}
}

type UpdateInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceInvoker) Invoke() (*model.UpdateInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceResponse), nil
	}
}

type UpdatePluginsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePluginsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePluginsInvoker) Invoke() (*model.UpdatePluginsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePluginsResponse), nil
	}
}

type UpdateUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateUserInvoker) Invoke() (*model.UpdateUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUserResponse), nil
	}
}

type CreateBindingInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBindingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateBindingInvoker) Invoke() (*model.CreateBindingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBindingResponse), nil
	}
}

type DeleteBindingInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBindingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteBindingInvoker) Invoke() (*model.DeleteBindingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBindingResponse), nil
	}
}

type ListBindingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBindingsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBindingsInvoker) Invoke() (*model.ListBindingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBindingsResponse), nil
	}
}

type BatchDeleteExchangesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteExchangesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteExchangesInvoker) Invoke() (*model.BatchDeleteExchangesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteExchangesResponse), nil
	}
}

type CreateExchangeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateExchangeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateExchangeInvoker) Invoke() (*model.CreateExchangeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateExchangeResponse), nil
	}
}

type ListExchangesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListExchangesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListExchangesInvoker) Invoke() (*model.ListExchangesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListExchangesResponse), nil
	}
}

type BatchDeleteQueuesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteQueuesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteQueuesInvoker) Invoke() (*model.BatchDeleteQueuesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteQueuesResponse), nil
	}
}

type CreateQueueInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateQueueInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateQueueInvoker) Invoke() (*model.CreateQueueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateQueueResponse), nil
	}
}

type DeleteQueueInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteQueueInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteQueueInfoInvoker) Invoke() (*model.DeleteQueueInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteQueueInfoResponse), nil
	}
}

type ListQueuesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQueuesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQueuesInvoker) Invoke() (*model.ListQueuesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQueuesResponse), nil
	}
}

type ShowQueueDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQueueDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowQueueDetailsInvoker) Invoke() (*model.ShowQueueDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQueueDetailsResponse), nil
	}
}

type BatchDeleteVhostsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteVhostsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteVhostsInvoker) Invoke() (*model.BatchDeleteVhostsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteVhostsResponse), nil
	}
}

type CreateVhostInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVhostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateVhostInvoker) Invoke() (*model.CreateVhostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVhostResponse), nil
	}
}

type ListVhostsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVhostsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVhostsInvoker) Invoke() (*model.ListVhostsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVhostsResponse), nil
	}
}
