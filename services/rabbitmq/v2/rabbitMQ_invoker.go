package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rabbitmq/v2/model"
)

type BatchCreateOrDeleteRabbitMqTagInvoker struct {
	*invoker.BaseInvoker
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

func (i *CreatePostPaidInstanceInvoker) Invoke() (*model.CreatePostPaidInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePostPaidInstanceResponse), nil
	}
}

type DeleteBackgroundTaskInvoker struct {
	*invoker.BaseInvoker
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

func (i *DeleteInstanceInvoker) Invoke() (*model.DeleteInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceResponse), nil
	}
}

type ListAvailableZonesInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListProductsInvoker) Invoke() (*model.ListProductsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProductsResponse), nil
	}
}

type ResetPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetPasswordInvoker) Invoke() (*model.ResetPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetPasswordResponse), nil
	}
}

type ResizeInstanceInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowBackgroundTaskInvoker) Invoke() (*model.ShowBackgroundTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBackgroundTaskResponse), nil
	}
}

type ShowInstanceInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdatePluginsInvoker) Invoke() (*model.UpdatePluginsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePluginsResponse), nil
	}
}
