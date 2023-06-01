package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rocketmq/v2/model"
)

type BatchCreateOrDeleteRocketmqTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateOrDeleteRocketmqTagInvoker) Invoke() (*model.BatchCreateOrDeleteRocketmqTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateOrDeleteRocketmqTagResponse), nil
	}
}

type BatchDeleteInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteInstancesInvoker) Invoke() (*model.BatchDeleteInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteInstancesResponse), nil
	}
}

type BatchUpdateConsumerGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateConsumerGroupInvoker) Invoke() (*model.BatchUpdateConsumerGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateConsumerGroupResponse), nil
	}
}

type CreateConsumerGroupOrBatchDeleteConsumerGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateConsumerGroupOrBatchDeleteConsumerGroupInvoker) Invoke() (*model.CreateConsumerGroupOrBatchDeleteConsumerGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConsumerGroupOrBatchDeleteConsumerGroupResponse), nil
	}
}

type CreateInstanceByEngineInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceByEngineInvoker) Invoke() (*model.CreateInstanceByEngineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceByEngineResponse), nil
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

type CreateRocketMqMigrationTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRocketMqMigrationTaskInvoker) Invoke() (*model.CreateRocketMqMigrationTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRocketMqMigrationTaskResponse), nil
	}
}

type CreateUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateUserInvoker) Invoke() (*model.CreateUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateUserResponse), nil
	}
}

type DeleteConsumerGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteConsumerGroupInvoker) Invoke() (*model.DeleteConsumerGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteConsumerGroupResponse), nil
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

type DeleteRocketMqMigrationTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRocketMqMigrationTaskInvoker) Invoke() (*model.DeleteRocketMqMigrationTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRocketMqMigrationTaskResponse), nil
	}
}

type DeleteUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteUserInvoker) Invoke() (*model.DeleteUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteUserResponse), nil
	}
}

type ExportDlqMessageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportDlqMessageInvoker) Invoke() (*model.ExportDlqMessageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportDlqMessageResponse), nil
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

type ListBrokersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBrokersInvoker) Invoke() (*model.ListBrokersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBrokersResponse), nil
	}
}

type ListConsumeGroupAccessPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConsumeGroupAccessPolicyInvoker) Invoke() (*model.ListConsumeGroupAccessPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConsumeGroupAccessPolicyResponse), nil
	}
}

type ListInstanceConsumerGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceConsumerGroupsInvoker) Invoke() (*model.ListInstanceConsumerGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceConsumerGroupsResponse), nil
	}
}

type ListInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstancesInvoker) Invoke() (*model.ListInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstancesResponse), nil
	}
}

type ListMessageTraceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMessageTraceInvoker) Invoke() (*model.ListMessageTraceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMessageTraceResponse), nil
	}
}

type ListMessagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMessagesInvoker) Invoke() (*model.ListMessagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMessagesResponse), nil
	}
}

type ListRocketMqMigrationTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRocketMqMigrationTaskInvoker) Invoke() (*model.ListRocketMqMigrationTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRocketMqMigrationTaskResponse), nil
	}
}

type ListTopicAccessPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTopicAccessPolicyInvoker) Invoke() (*model.ListTopicAccessPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTopicAccessPolicyResponse), nil
	}
}

type ListUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserInvoker) Invoke() (*model.ListUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserResponse), nil
	}
}

type ResetConsumeOffsetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetConsumeOffsetInvoker) Invoke() (*model.ResetConsumeOffsetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetConsumeOffsetResponse), nil
	}
}

type SendRocketMqDlqMessageInvoker struct {
	*invoker.BaseInvoker
}

func (i *SendRocketMqDlqMessageInvoker) Invoke() (*model.SendRocketMqDlqMessageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendRocketMqDlqMessageResponse), nil
	}
}

type ShowConsumerConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConsumerConnectionsInvoker) Invoke() (*model.ShowConsumerConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConsumerConnectionsResponse), nil
	}
}

type ShowConsumerListOrDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConsumerListOrDetailsInvoker) Invoke() (*model.ShowConsumerListOrDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConsumerListOrDetailsResponse), nil
	}
}

type ShowGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGroupInvoker) Invoke() (*model.ShowGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGroupResponse), nil
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

type ShowRocketmqProjectTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRocketmqProjectTagsInvoker) Invoke() (*model.ShowRocketmqProjectTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRocketmqProjectTagsResponse), nil
	}
}

type ShowRocketmqTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRocketmqTagsInvoker) Invoke() (*model.ShowRocketmqTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRocketmqTagsResponse), nil
	}
}

type ShowUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUserInvoker) Invoke() (*model.ShowUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUserResponse), nil
	}
}

type UpdateConsumerGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateConsumerGroupInvoker) Invoke() (*model.UpdateConsumerGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateConsumerGroupResponse), nil
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

type UpdateUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUserInvoker) Invoke() (*model.UpdateUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUserResponse), nil
	}
}

type ValidateRocketMqConsumedMessageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ValidateRocketMqConsumedMessageInvoker) Invoke() (*model.ValidateRocketMqConsumedMessageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ValidateRocketMqConsumedMessageResponse), nil
	}
}

type CreateTopicOrBatchDeleteTopicInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTopicOrBatchDeleteTopicInvoker) Invoke() (*model.CreateTopicOrBatchDeleteTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTopicOrBatchDeleteTopicResponse), nil
	}
}

type DeleteTopicInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTopicInvoker) Invoke() (*model.DeleteTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTopicResponse), nil
	}
}

type ListConsumerGroupOfTopicInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConsumerGroupOfTopicInvoker) Invoke() (*model.ListConsumerGroupOfTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConsumerGroupOfTopicResponse), nil
	}
}

type ListRocketInstanceTopicsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRocketInstanceTopicsInvoker) Invoke() (*model.ListRocketInstanceTopicsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRocketInstanceTopicsResponse), nil
	}
}

type ShowOneTopicInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOneTopicInvoker) Invoke() (*model.ShowOneTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOneTopicResponse), nil
	}
}

type ShowTopicStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTopicStatusInvoker) Invoke() (*model.ShowTopicStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTopicStatusResponse), nil
	}
}

type UpdateTopicInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTopicInvoker) Invoke() (*model.UpdateTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTopicResponse), nil
	}
}
