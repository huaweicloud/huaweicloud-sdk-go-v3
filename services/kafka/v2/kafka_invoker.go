package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kafka/v2/model"
)

type BatchCreateOrDeleteKafkaTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateOrDeleteKafkaTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateOrDeleteKafkaTagInvoker) Invoke() (*model.BatchCreateOrDeleteKafkaTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateOrDeleteKafkaTagResponse), nil
	}
}

type BatchDeleteGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteGroupInvoker) Invoke() (*model.BatchDeleteGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteGroupResponse), nil
	}
}

type BatchDeleteInstanceTopicInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteInstanceTopicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteInstanceTopicInvoker) Invoke() (*model.BatchDeleteInstanceTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteInstanceTopicResponse), nil
	}
}

type BatchDeleteInstanceUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteInstanceUsersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteInstanceUsersInvoker) Invoke() (*model.BatchDeleteInstanceUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteInstanceUsersResponse), nil
	}
}

type BatchDeleteMessageDiagnosisReportsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteMessageDiagnosisReportsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteMessageDiagnosisReportsInvoker) Invoke() (*model.BatchDeleteMessageDiagnosisReportsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteMessageDiagnosisReportsResponse), nil
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

type CloseKafkaManagerInvoker struct {
	*invoker.BaseInvoker
}

func (i *CloseKafkaManagerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CloseKafkaManagerInvoker) Invoke() (*model.CloseKafkaManagerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CloseKafkaManagerResponse), nil
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

type CreateInstanceTopicInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceTopicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstanceTopicInvoker) Invoke() (*model.CreateInstanceTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceTopicResponse), nil
	}
}

type CreateInstanceUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstanceUserInvoker) Invoke() (*model.CreateInstanceUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceUserResponse), nil
	}
}

type CreateKafkaConsumerGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateKafkaConsumerGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateKafkaConsumerGroupInvoker) Invoke() (*model.CreateKafkaConsumerGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateKafkaConsumerGroupResponse), nil
	}
}

type CreateKafkaReassignmentTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateKafkaReassignmentTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateKafkaReassignmentTaskInvoker) Invoke() (*model.CreateKafkaReassignmentTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateKafkaReassignmentTaskResponse), nil
	}
}

type CreateKafkaRebalanceLogTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateKafkaRebalanceLogTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateKafkaRebalanceLogTaskInvoker) Invoke() (*model.CreateKafkaRebalanceLogTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateKafkaRebalanceLogTaskResponse), nil
	}
}

type CreateKafkaTopicQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateKafkaTopicQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateKafkaTopicQuotaInvoker) Invoke() (*model.CreateKafkaTopicQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateKafkaTopicQuotaResponse), nil
	}
}

type CreateKafkaUserClientQuotaTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateKafkaUserClientQuotaTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateKafkaUserClientQuotaTaskInvoker) Invoke() (*model.CreateKafkaUserClientQuotaTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateKafkaUserClientQuotaTaskResponse), nil
	}
}

type CreateMessageDiagnosisTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMessageDiagnosisTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMessageDiagnosisTaskInvoker) Invoke() (*model.CreateMessageDiagnosisTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMessageDiagnosisTaskResponse), nil
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

type CreatePostPaidKafkaInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePostPaidKafkaInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePostPaidKafkaInstanceInvoker) Invoke() (*model.CreatePostPaidKafkaInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePostPaidKafkaInstanceResponse), nil
	}
}

type CreateReassignmentTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateReassignmentTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateReassignmentTaskInvoker) Invoke() (*model.CreateReassignmentTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateReassignmentTaskResponse), nil
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

type DeleteConsumerGroupOffsetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteConsumerGroupOffsetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteConsumerGroupOffsetsInvoker) Invoke() (*model.DeleteConsumerGroupOffsetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteConsumerGroupOffsetsResponse), nil
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

type DeleteKafkaMessageInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteKafkaMessageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteKafkaMessageInvoker) Invoke() (*model.DeleteKafkaMessageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteKafkaMessageResponse), nil
	}
}

type DeleteKafkaTopicQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteKafkaTopicQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteKafkaTopicQuotaInvoker) Invoke() (*model.DeleteKafkaTopicQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteKafkaTopicQuotaResponse), nil
	}
}

type DeleteKafkaUserClientQuotaTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteKafkaUserClientQuotaTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteKafkaUserClientQuotaTaskInvoker) Invoke() (*model.DeleteKafkaUserClientQuotaTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteKafkaUserClientQuotaTaskResponse), nil
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

type ListInstanceConsumerGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceConsumerGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceConsumerGroupInvoker) Invoke() (*model.ListInstanceConsumerGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceConsumerGroupResponse), nil
	}
}

type ListInstanceConsumerGroupMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceConsumerGroupMembersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceConsumerGroupMembersInvoker) Invoke() (*model.ListInstanceConsumerGroupMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceConsumerGroupMembersResponse), nil
	}
}

type ListInstanceConsumerGroupMessageOffsetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceConsumerGroupMessageOffsetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceConsumerGroupMessageOffsetInvoker) Invoke() (*model.ListInstanceConsumerGroupMessageOffsetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceConsumerGroupMessageOffsetResponse), nil
	}
}

type ListInstanceConsumerGroupTopicsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceConsumerGroupTopicsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceConsumerGroupTopicsInvoker) Invoke() (*model.ListInstanceConsumerGroupTopicsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceConsumerGroupTopicsResponse), nil
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

type ListInstanceTopicsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceTopicsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceTopicsInvoker) Invoke() (*model.ListInstanceTopicsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceTopicsResponse), nil
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

type ListMessageDiagnosisReportsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMessageDiagnosisReportsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMessageDiagnosisReportsInvoker) Invoke() (*model.ListMessageDiagnosisReportsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMessageDiagnosisReportsResponse), nil
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

type ListTopicPartitionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTopicPartitionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTopicPartitionsInvoker) Invoke() (*model.ListTopicPartitionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTopicPartitionsResponse), nil
	}
}

type ListTopicProducersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTopicProducersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTopicProducersInvoker) Invoke() (*model.ListTopicProducersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTopicProducersResponse), nil
	}
}

type ModifyInstanceConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyInstanceConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyInstanceConfigsInvoker) Invoke() (*model.ModifyInstanceConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyInstanceConfigsResponse), nil
	}
}

type ModifyKafkaTopicQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyKafkaTopicQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyKafkaTopicQuotaInvoker) Invoke() (*model.ModifyKafkaTopicQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyKafkaTopicQuotaResponse), nil
	}
}

type ResetManagerPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetManagerPasswordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetManagerPasswordInvoker) Invoke() (*model.ResetManagerPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetManagerPasswordResponse), nil
	}
}

type ResetMessageOffsetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetMessageOffsetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetMessageOffsetInvoker) Invoke() (*model.ResetMessageOffsetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetMessageOffsetResponse), nil
	}
}

type ResetMessageOffsetWithEngineInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetMessageOffsetWithEngineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetMessageOffsetWithEngineInvoker) Invoke() (*model.ResetMessageOffsetWithEngineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetMessageOffsetWithEngineResponse), nil
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

type ResetUserPasswrodInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetUserPasswrodInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetUserPasswrodInvoker) Invoke() (*model.ResetUserPasswrodResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetUserPasswrodResponse), nil
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

type ResizeKafkaInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeKafkaInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResizeKafkaInstanceInvoker) Invoke() (*model.ResizeKafkaInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeKafkaInstanceResponse), nil
	}
}

type RestartManagerInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestartManagerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RestartManagerInvoker) Invoke() (*model.RestartManagerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartManagerResponse), nil
	}
}

type SendKafkaMessageInvoker struct {
	*invoker.BaseInvoker
}

func (i *SendKafkaMessageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SendKafkaMessageInvoker) Invoke() (*model.SendKafkaMessageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendKafkaMessageResponse), nil
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

type ShowCoordinatorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCoordinatorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCoordinatorsInvoker) Invoke() (*model.ShowCoordinatorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCoordinatorsResponse), nil
	}
}

type ShowDiagnosisPreCheckInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDiagnosisPreCheckInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDiagnosisPreCheckInvoker) Invoke() (*model.ShowDiagnosisPreCheckResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDiagnosisPreCheckResponse), nil
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

type ShowGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGroupsInvoker) Invoke() (*model.ShowGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGroupsResponse), nil
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

type ShowInstanceConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceConfigsInvoker) Invoke() (*model.ShowInstanceConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceConfigsResponse), nil
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

type ShowInstanceMessagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceMessagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceMessagesInvoker) Invoke() (*model.ShowInstanceMessagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceMessagesResponse), nil
	}
}

type ShowInstanceTopicDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceTopicDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceTopicDetailInvoker) Invoke() (*model.ShowInstanceTopicDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceTopicDetailResponse), nil
	}
}

type ShowInstanceUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceUsersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceUsersInvoker) Invoke() (*model.ShowInstanceUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceUsersResponse), nil
	}
}

type ShowKafkaInstanceExtendProductInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowKafkaInstanceExtendProductInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowKafkaInstanceExtendProductInfoInvoker) Invoke() (*model.ShowKafkaInstanceExtendProductInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowKafkaInstanceExtendProductInfoResponse), nil
	}
}

type ShowKafkaProjectTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowKafkaProjectTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowKafkaProjectTagsInvoker) Invoke() (*model.ShowKafkaProjectTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowKafkaProjectTagsResponse), nil
	}
}

type ShowKafkaRebalanceLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowKafkaRebalanceLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowKafkaRebalanceLogInvoker) Invoke() (*model.ShowKafkaRebalanceLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowKafkaRebalanceLogResponse), nil
	}
}

type ShowKafkaTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowKafkaTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowKafkaTagsInvoker) Invoke() (*model.ShowKafkaTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowKafkaTagsResponse), nil
	}
}

type ShowKafkaTopicPartitionDiskusageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowKafkaTopicPartitionDiskusageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowKafkaTopicPartitionDiskusageInvoker) Invoke() (*model.ShowKafkaTopicPartitionDiskusageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowKafkaTopicPartitionDiskusageResponse), nil
	}
}

type ShowKafkaTopicQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowKafkaTopicQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowKafkaTopicQuotaInvoker) Invoke() (*model.ShowKafkaTopicQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowKafkaTopicQuotaResponse), nil
	}
}

type ShowKafkaUserClientQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowKafkaUserClientQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowKafkaUserClientQuotaInvoker) Invoke() (*model.ShowKafkaUserClientQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowKafkaUserClientQuotaResponse), nil
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

type ShowMessageDiagnosisReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMessageDiagnosisReportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMessageDiagnosisReportInvoker) Invoke() (*model.ShowMessageDiagnosisReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMessageDiagnosisReportResponse), nil
	}
}

type ShowMessagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMessagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMessagesInvoker) Invoke() (*model.ShowMessagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMessagesResponse), nil
	}
}

type ShowPartitionBeginningMessageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPartitionBeginningMessageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPartitionBeginningMessageInvoker) Invoke() (*model.ShowPartitionBeginningMessageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPartitionBeginningMessageResponse), nil
	}
}

type ShowPartitionEndMessageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPartitionEndMessageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPartitionEndMessageInvoker) Invoke() (*model.ShowPartitionEndMessageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPartitionEndMessageResponse), nil
	}
}

type ShowPartitionMessageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPartitionMessageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPartitionMessageInvoker) Invoke() (*model.ShowPartitionMessageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPartitionMessageResponse), nil
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

type ShowTopicAccessPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTopicAccessPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTopicAccessPolicyInvoker) Invoke() (*model.ShowTopicAccessPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTopicAccessPolicyResponse), nil
	}
}

type StopKafkaRebalanceLogTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopKafkaRebalanceLogTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopKafkaRebalanceLogTaskInvoker) Invoke() (*model.StopKafkaRebalanceLogTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopKafkaRebalanceLogTaskResponse), nil
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

type UpdateInstanceAutoCreateTopicInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceAutoCreateTopicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceAutoCreateTopicInvoker) Invoke() (*model.UpdateInstanceAutoCreateTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceAutoCreateTopicResponse), nil
	}
}

type UpdateInstanceConsumerGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceConsumerGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceConsumerGroupInvoker) Invoke() (*model.UpdateInstanceConsumerGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceConsumerGroupResponse), nil
	}
}

type UpdateInstanceCrossVpcIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceCrossVpcIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceCrossVpcIpInvoker) Invoke() (*model.UpdateInstanceCrossVpcIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceCrossVpcIpResponse), nil
	}
}

type UpdateInstanceGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceGroupInvoker) Invoke() (*model.UpdateInstanceGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceGroupResponse), nil
	}
}

type UpdateInstanceTopicInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceTopicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceTopicInvoker) Invoke() (*model.UpdateInstanceTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceTopicResponse), nil
	}
}

type UpdateInstanceUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceUserInvoker) Invoke() (*model.UpdateInstanceUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceUserResponse), nil
	}
}

type UpdateKafkaPortProtocolInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateKafkaPortProtocolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateKafkaPortProtocolInvoker) Invoke() (*model.UpdateKafkaPortProtocolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateKafkaPortProtocolResponse), nil
	}
}

type UpdateKafkaUserClientQuotaTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateKafkaUserClientQuotaTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateKafkaUserClientQuotaTaskInvoker) Invoke() (*model.UpdateKafkaUserClientQuotaTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateKafkaUserClientQuotaTaskResponse), nil
	}
}

type UpdateTopicAccessPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTopicAccessPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTopicAccessPolicyInvoker) Invoke() (*model.UpdateTopicAccessPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTopicAccessPolicyResponse), nil
	}
}

type UpdateTopicReplicaInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTopicReplicaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTopicReplicaInvoker) Invoke() (*model.UpdateTopicReplicaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTopicReplicaResponse), nil
	}
}

type CreateConnectorInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateConnectorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateConnectorInvoker) Invoke() (*model.CreateConnectorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConnectorResponse), nil
	}
}

type CreateConnectorTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateConnectorTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateConnectorTaskInvoker) Invoke() (*model.CreateConnectorTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConnectorTaskResponse), nil
	}
}

type DeleteConnectorInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteConnectorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteConnectorInvoker) Invoke() (*model.DeleteConnectorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteConnectorResponse), nil
	}
}

type DeleteConnectorTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteConnectorTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteConnectorTaskInvoker) Invoke() (*model.DeleteConnectorTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteConnectorTaskResponse), nil
	}
}

type ListConnectorTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConnectorTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListConnectorTasksInvoker) Invoke() (*model.ListConnectorTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConnectorTasksResponse), nil
	}
}

type PauseConnectorTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *PauseConnectorTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PauseConnectorTaskInvoker) Invoke() (*model.PauseConnectorTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PauseConnectorTaskResponse), nil
	}
}

type RestartConnectorTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestartConnectorTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RestartConnectorTaskInvoker) Invoke() (*model.RestartConnectorTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartConnectorTaskResponse), nil
	}
}

type RestartSmartConnectorTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestartSmartConnectorTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RestartSmartConnectorTaskInvoker) Invoke() (*model.RestartSmartConnectorTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartSmartConnectorTaskResponse), nil
	}
}

type ResumeConnectorTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResumeConnectorTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResumeConnectorTaskInvoker) Invoke() (*model.ResumeConnectorTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResumeConnectorTaskResponse), nil
	}
}

type ShowConnectorTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConnectorTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowConnectorTaskInvoker) Invoke() (*model.ShowConnectorTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConnectorTaskResponse), nil
	}
}
