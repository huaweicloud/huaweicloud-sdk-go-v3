package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rocketmq/v2/model"
)

type BatchCreateOrDeleteRocketmqTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateOrDeleteRocketmqTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateOrDeleteRocketmqTagInvoker) Invoke() (*model.BatchCreateOrDeleteRocketmqTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateOrDeleteRocketmqTagResponse), nil
	}
}

type BatchDeleteDiagnosisRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteDiagnosisRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteDiagnosisRecordsInvoker) Invoke() (*model.BatchDeleteDiagnosisRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteDiagnosisRecordsResponse), nil
	}
}

type BatchDeleteDiagnosisRecordsForRocketMqInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteDiagnosisRecordsForRocketMqInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteDiagnosisRecordsForRocketMqInvoker) Invoke() (*model.BatchDeleteDiagnosisRecordsForRocketMqResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteDiagnosisRecordsForRocketMqResponse), nil
	}
}

type BatchDeleteInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteInstancesInvoker) Invoke() (*model.BatchDeleteInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteInstancesResponse), nil
	}
}

type BatchDeleteRocketMqMigrationTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteRocketMqMigrationTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteRocketMqMigrationTaskInvoker) Invoke() (*model.BatchDeleteRocketMqMigrationTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteRocketMqMigrationTaskResponse), nil
	}
}

type BatchUpdateConsumerGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateConsumerGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateConsumerGroupOrBatchDeleteConsumerGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateConsumerGroupOrBatchDeleteConsumerGroupInvoker) Invoke() (*model.CreateConsumerGroupOrBatchDeleteConsumerGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConsumerGroupOrBatchDeleteConsumerGroupResponse), nil
	}
}

type CreateDiagnosisTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDiagnosisTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDiagnosisTaskInvoker) Invoke() (*model.CreateDiagnosisTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDiagnosisTaskResponse), nil
	}
}

type CreateDiagnosisTaskForRocketMqInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDiagnosisTaskForRocketMqInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDiagnosisTaskForRocketMqInvoker) Invoke() (*model.CreateDiagnosisTaskForRocketMqResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDiagnosisTaskForRocketMqResponse), nil
	}
}

type CreateInstanceByEngineInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceByEngineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

type CreatePostPaidInstanceForRocketMqInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePostPaidInstanceForRocketMqInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePostPaidInstanceForRocketMqInvoker) Invoke() (*model.CreatePostPaidInstanceForRocketMqResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePostPaidInstanceForRocketMqResponse), nil
	}
}

type CreateRocketMqMigrationTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRocketMqMigrationTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

type DeleteConsumerGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteConsumerGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

type DeleteScheduledTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteScheduledTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteScheduledTaskInvoker) Invoke() (*model.DeleteScheduledTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteScheduledTaskResponse), nil
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

type EnableDnsInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableDnsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableDnsInvoker) Invoke() (*model.EnableDnsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableDnsResponse), nil
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

type ListBrokersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBrokersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBrokersInvoker) Invoke() (*model.ListBrokersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBrokersResponse), nil
	}
}

type ListConfigFeaturesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConfigFeaturesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListConfigFeaturesInvoker) Invoke() (*model.ListConfigFeaturesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConfigFeaturesResponse), nil
	}
}

type ListConsumeGroupAccessPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConsumeGroupAccessPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListConsumeGroupAccessPolicyInvoker) Invoke() (*model.ListConsumeGroupAccessPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConsumeGroupAccessPolicyResponse), nil
	}
}

type ListConsumeGroupAccessPolicyForRocketMqInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConsumeGroupAccessPolicyForRocketMqInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListConsumeGroupAccessPolicyForRocketMqInvoker) Invoke() (*model.ListConsumeGroupAccessPolicyForRocketMqResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConsumeGroupAccessPolicyForRocketMqResponse), nil
	}
}

type ListDiagnosisReportsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDiagnosisReportsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDiagnosisReportsInvoker) Invoke() (*model.ListDiagnosisReportsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDiagnosisReportsResponse), nil
	}
}

type ListDiagnosisReportsForRocketMqInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDiagnosisReportsForRocketMqInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDiagnosisReportsForRocketMqInvoker) Invoke() (*model.ListDiagnosisReportsForRocketMqResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDiagnosisReportsForRocketMqResponse), nil
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

type ListInstanceConsumerGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceConsumerGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListMessageTraceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMessageTraceInvoker) Invoke() (*model.ListMessageTraceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMessageTraceResponse), nil
	}
}

type ListMessageTraceForRocketMqInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMessageTraceForRocketMqInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMessageTraceForRocketMqInvoker) Invoke() (*model.ListMessageTraceForRocketMqResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMessageTraceForRocketMqResponse), nil
	}
}

type ListMessagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMessagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMessagesInvoker) Invoke() (*model.ListMessagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMessagesResponse), nil
	}
}

type ListMessagesForRocketMqInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMessagesForRocketMqInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMessagesForRocketMqInvoker) Invoke() (*model.ListMessagesForRocketMqResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMessagesForRocketMqResponse), nil
	}
}

type ListRocketMqMigrationTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRocketMqMigrationTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRocketMqMigrationTaskInvoker) Invoke() (*model.ListRocketMqMigrationTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRocketMqMigrationTaskResponse), nil
	}
}

type ListScheduledTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScheduledTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScheduledTasksInvoker) Invoke() (*model.ListScheduledTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScheduledTasksResponse), nil
	}
}

type ListTopicAccessPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTopicAccessPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

type ModifyInstanceSslConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyInstanceSslConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyInstanceSslConfigInvoker) Invoke() (*model.ModifyInstanceSslConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyInstanceSslConfigResponse), nil
	}
}

type ModifyRecyclePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyRecyclePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyRecyclePolicyInvoker) Invoke() (*model.ModifyRecyclePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyRecyclePolicyResponse), nil
	}
}

type ResetConsumeOffsetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetConsumeOffsetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetConsumeOffsetInvoker) Invoke() (*model.ResetConsumeOffsetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetConsumeOffsetResponse), nil
	}
}

type ResetConsumeOffsetForRocketMqInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetConsumeOffsetForRocketMqInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetConsumeOffsetForRocketMqInvoker) Invoke() (*model.ResetConsumeOffsetForRocketMqResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetConsumeOffsetForRocketMqResponse), nil
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

type ResizeInstanceForRocketMqInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeInstanceForRocketMqInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResizeInstanceForRocketMqInvoker) Invoke() (*model.ResizeInstanceForRocketMqResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeInstanceForRocketMqResponse), nil
	}
}

type RestoreRecycleInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreRecycleInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RestoreRecycleInstanceInvoker) Invoke() (*model.RestoreRecycleInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreRecycleInstanceResponse), nil
	}
}

type SendDlqMessageInvoker struct {
	*invoker.BaseInvoker
}

func (i *SendDlqMessageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SendDlqMessageInvoker) Invoke() (*model.SendDlqMessageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendDlqMessageResponse), nil
	}
}

type SendDlqMessageForRocketMqInvoker struct {
	*invoker.BaseInvoker
}

func (i *SendDlqMessageForRocketMqInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SendDlqMessageForRocketMqInvoker) Invoke() (*model.SendDlqMessageForRocketMqResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendDlqMessageForRocketMqResponse), nil
	}
}

type SendMessageInvoker struct {
	*invoker.BaseInvoker
}

func (i *SendMessageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SendMessageInvoker) Invoke() (*model.SendMessageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendMessageResponse), nil
	}
}

type SendMessageForRocketMqInvoker struct {
	*invoker.BaseInvoker
}

func (i *SendMessageForRocketMqInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SendMessageForRocketMqInvoker) Invoke() (*model.SendMessageForRocketMqResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendMessageForRocketMqResponse), nil
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

type ShowBackgroundTaskProgressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBackgroundTaskProgressInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBackgroundTaskProgressInvoker) Invoke() (*model.ShowBackgroundTaskProgressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBackgroundTaskProgressResponse), nil
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

type ShowConsumerConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConsumerConnectionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowConsumerConnectionsInvoker) Invoke() (*model.ShowConsumerConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConsumerConnectionsResponse), nil
	}
}

type ShowConsumerConnectionsForRocketMqInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConsumerConnectionsForRocketMqInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowConsumerConnectionsForRocketMqInvoker) Invoke() (*model.ShowConsumerConnectionsForRocketMqResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConsumerConnectionsForRocketMqResponse), nil
	}
}

type ShowConsumerListOrDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConsumerListOrDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowConsumerListOrDetailsInvoker) Invoke() (*model.ShowConsumerListOrDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConsumerListOrDetailsResponse), nil
	}
}

type ShowDiagnosisReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDiagnosisReportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDiagnosisReportInvoker) Invoke() (*model.ShowDiagnosisReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDiagnosisReportResponse), nil
	}
}

type ShowDiagnosisReportForRocketMqInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDiagnosisReportForRocketMqInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDiagnosisReportForRocketMqInvoker) Invoke() (*model.ShowDiagnosisReportForRocketMqResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDiagnosisReportForRocketMqResponse), nil
	}
}

type ShowDiagnosisStackInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDiagnosisStackInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDiagnosisStackInvoker) Invoke() (*model.ShowDiagnosisStackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDiagnosisStackResponse), nil
	}
}

type ShowDiagnosisStackForRocketMqInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDiagnosisStackForRocketMqInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDiagnosisStackForRocketMqInvoker) Invoke() (*model.ShowDiagnosisStackForRocketMqResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDiagnosisStackForRocketMqResponse), nil
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

type ShowEngineInstanceExtendProductInfoForRocketMqInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEngineInstanceExtendProductInfoForRocketMqInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEngineInstanceExtendProductInfoForRocketMqInvoker) Invoke() (*model.ShowEngineInstanceExtendProductInfoForRocketMqResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEngineInstanceExtendProductInfoForRocketMqResponse), nil
	}
}

type ShowGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

type ShowInstanceNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceNodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceNodesInvoker) Invoke() (*model.ShowInstanceNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceNodesResponse), nil
	}
}

type ShowQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowQuotasInvoker) Invoke() (*model.ShowQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuotasResponse), nil
	}
}

type ShowRecycleInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRecycleInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRecycleInstancesInvoker) Invoke() (*model.ShowRecycleInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRecycleInstancesResponse), nil
	}
}

type ShowRocketMqConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRocketMqConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRocketMqConfigsInvoker) Invoke() (*model.ShowRocketMqConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRocketMqConfigsResponse), nil
	}
}

type ShowRocketMqProductCoresInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRocketMqProductCoresInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRocketMqProductCoresInvoker) Invoke() (*model.ShowRocketMqProductCoresResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRocketMqProductCoresResponse), nil
	}
}

type ShowRocketMqScalePreCheckInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRocketMqScalePreCheckInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRocketMqScalePreCheckInfoInvoker) Invoke() (*model.ShowRocketMqScalePreCheckInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRocketMqScalePreCheckInfoResponse), nil
	}
}

type ShowRocketmqProjectTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRocketmqProjectTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowRocketmqTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowUserInvoker) Invoke() (*model.ShowUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUserResponse), nil
	}
}

type ShowVolumeExpandConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVolumeExpandConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowVolumeExpandConfigInvoker) Invoke() (*model.ShowVolumeExpandConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVolumeExpandConfigResponse), nil
	}
}

type UpdateConsumerGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateConsumerGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

type UpdateRocketMqConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRocketMqConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRocketMqConfigsInvoker) Invoke() (*model.UpdateRocketMqConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRocketMqConfigsResponse), nil
	}
}

type UpdateScheduledTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateScheduledTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateScheduledTaskInvoker) Invoke() (*model.UpdateScheduledTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateScheduledTaskResponse), nil
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

type UpdateVolumeExpansionConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVolumeExpansionConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateVolumeExpansionConfigInvoker) Invoke() (*model.UpdateVolumeExpansionConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVolumeExpansionConfigResponse), nil
	}
}

type ValidateConsumedMessageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ValidateConsumedMessageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ValidateConsumedMessageInvoker) Invoke() (*model.ValidateConsumedMessageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ValidateConsumedMessageResponse), nil
	}
}

type ValidateConsumedMessageForRocketMqInvoker struct {
	*invoker.BaseInvoker
}

func (i *ValidateConsumedMessageForRocketMqInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ValidateConsumedMessageForRocketMqInvoker) Invoke() (*model.ValidateConsumedMessageForRocketMqResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ValidateConsumedMessageForRocketMqResponse), nil
	}
}

type CreateTopicOrBatchDeleteTopicInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTopicOrBatchDeleteTopicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteTopicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListConsumerGroupOfTopicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListRocketInstanceTopicsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowOneTopicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowTopicStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateTopicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTopicInvoker) Invoke() (*model.UpdateTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTopicResponse), nil
	}
}
