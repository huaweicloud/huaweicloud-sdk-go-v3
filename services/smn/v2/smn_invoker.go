package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/smn/v2/model"
)

type AddSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddSubscriptionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddSubscriptionInvoker) Invoke() (*model.AddSubscriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddSubscriptionResponse), nil
	}
}

type AddSubscriptionFromSubscriptionUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddSubscriptionFromSubscriptionUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddSubscriptionFromSubscriptionUserInvoker) Invoke() (*model.AddSubscriptionFromSubscriptionUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddSubscriptionFromSubscriptionUserResponse), nil
	}
}

type BatchCreateOrDeleteResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateOrDeleteResourceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateOrDeleteResourceTagsInvoker) Invoke() (*model.BatchCreateOrDeleteResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateOrDeleteResourceTagsResponse), nil
	}
}

type BatchCreateSubscriptionsFilterPolicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateSubscriptionsFilterPolicesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateSubscriptionsFilterPolicesInvoker) Invoke() (*model.BatchCreateSubscriptionsFilterPolicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateSubscriptionsFilterPolicesResponse), nil
	}
}

type BatchDeleteSubscriptionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteSubscriptionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteSubscriptionsInvoker) Invoke() (*model.BatchDeleteSubscriptionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteSubscriptionsResponse), nil
	}
}

type BatchDeleteSubscriptionsByTopicInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteSubscriptionsByTopicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteSubscriptionsByTopicInvoker) Invoke() (*model.BatchDeleteSubscriptionsByTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteSubscriptionsByTopicResponse), nil
	}
}

type BatchDeleteSubscriptionsFilterPolicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteSubscriptionsFilterPolicesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteSubscriptionsFilterPolicesInvoker) Invoke() (*model.BatchDeleteSubscriptionsFilterPolicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteSubscriptionsFilterPolicesResponse), nil
	}
}

type BatchUpdateSubscriptionsFilterPolicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateSubscriptionsFilterPolicesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdateSubscriptionsFilterPolicesInvoker) Invoke() (*model.BatchUpdateSubscriptionsFilterPolicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateSubscriptionsFilterPolicesResponse), nil
	}
}

type CancelSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelSubscriptionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelSubscriptionInvoker) Invoke() (*model.CancelSubscriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelSubscriptionResponse), nil
	}
}

type ConfirmSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ConfirmSubscriptionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ConfirmSubscriptionInvoker) Invoke() (*model.ConfirmSubscriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ConfirmSubscriptionResponse), nil
	}
}

type CreateKmsKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateKmsKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateKmsKeyInvoker) Invoke() (*model.CreateKmsKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateKmsKeyResponse), nil
	}
}

type CreateLogtankInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLogtankInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateLogtankInvoker) Invoke() (*model.CreateLogtankResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLogtankResponse), nil
	}
}

type CreateMessageTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMessageTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMessageTemplateInvoker) Invoke() (*model.CreateMessageTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMessageTemplateResponse), nil
	}
}

type CreateNotifyPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNotifyPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateNotifyPolicyInvoker) Invoke() (*model.CreateNotifyPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNotifyPolicyResponse), nil
	}
}

type CreateResourceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResourceTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateResourceTagInvoker) Invoke() (*model.CreateResourceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResourceTagResponse), nil
	}
}

type CreateTopicInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTopicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTopicInvoker) Invoke() (*model.CreateTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTopicResponse), nil
	}
}

type DeleteKmsKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteKmsKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteKmsKeyInvoker) Invoke() (*model.DeleteKmsKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteKmsKeyResponse), nil
	}
}

type DeleteLogtankInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLogtankInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteLogtankInvoker) Invoke() (*model.DeleteLogtankResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLogtankResponse), nil
	}
}

type DeleteMessageTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMessageTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMessageTemplateInvoker) Invoke() (*model.DeleteMessageTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMessageTemplateResponse), nil
	}
}

type DeleteNotifyPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNotifyPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteNotifyPolicyInvoker) Invoke() (*model.DeleteNotifyPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNotifyPolicyResponse), nil
	}
}

type DeleteResourceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteResourceTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteResourceTagInvoker) Invoke() (*model.DeleteResourceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteResourceTagResponse), nil
	}
}

type DeleteSubscriptionsByTopicInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSubscriptionsByTopicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSubscriptionsByTopicInvoker) Invoke() (*model.DeleteSubscriptionsByTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSubscriptionsByTopicResponse), nil
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

type DeleteTopicAttributeByNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTopicAttributeByNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTopicAttributeByNameInvoker) Invoke() (*model.DeleteTopicAttributeByNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTopicAttributeByNameResponse), nil
	}
}

type DeleteTopicAttributesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTopicAttributesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTopicAttributesInvoker) Invoke() (*model.DeleteTopicAttributesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTopicAttributesResponse), nil
	}
}

type DownloadHttpCertInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadHttpCertInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadHttpCertInvoker) Invoke() (*model.DownloadHttpCertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadHttpCertResponse), nil
	}
}

type DownloadHttpSignCertInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadHttpSignCertInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadHttpSignCertInvoker) Invoke() (*model.DownloadHttpSignCertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadHttpSignCertResponse), nil
	}
}

type ListCloudServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCloudServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCloudServiceInvoker) Invoke() (*model.ListCloudServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCloudServiceResponse), nil
	}
}

type ListCloudServicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCloudServicesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCloudServicesInvoker) Invoke() (*model.ListCloudServicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCloudServicesResponse), nil
	}
}

type ListLogtankInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLogtankInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLogtankInvoker) Invoke() (*model.ListLogtankResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLogtankResponse), nil
	}
}

type ListMessageTemplateDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMessageTemplateDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMessageTemplateDetailsInvoker) Invoke() (*model.ListMessageTemplateDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMessageTemplateDetailsResponse), nil
	}
}

type ListMessageTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMessageTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMessageTemplatesInvoker) Invoke() (*model.ListMessageTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMessageTemplatesResponse), nil
	}
}

type ListProjectTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectTagsInvoker) Invoke() (*model.ListProjectTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectTagsResponse), nil
	}
}

type ListProtocolsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProtocolsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProtocolsInvoker) Invoke() (*model.ListProtocolsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProtocolsResponse), nil
	}
}

type ListResourceInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceInstancesInvoker) Invoke() (*model.ListResourceInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceInstancesResponse), nil
	}
}

type ListResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceTagsInvoker) Invoke() (*model.ListResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceTagsResponse), nil
	}
}

type ListSubscriptionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubscriptionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSubscriptionsInvoker) Invoke() (*model.ListSubscriptionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubscriptionsResponse), nil
	}
}

type ListSubscriptionsByTopicInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubscriptionsByTopicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSubscriptionsByTopicInvoker) Invoke() (*model.ListSubscriptionsByTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubscriptionsByTopicResponse), nil
	}
}

type ListTopicAttributesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTopicAttributesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTopicAttributesInvoker) Invoke() (*model.ListTopicAttributesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTopicAttributesResponse), nil
	}
}

type ListTopicDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTopicDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTopicDetailsInvoker) Invoke() (*model.ListTopicDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTopicDetailsResponse), nil
	}
}

type ListTopicMessageStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTopicMessageStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTopicMessageStatisticsInvoker) Invoke() (*model.ListTopicMessageStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTopicMessageStatisticsResponse), nil
	}
}

type ListTopicsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTopicsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTopicsInvoker) Invoke() (*model.ListTopicsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTopicsResponse), nil
	}
}

type ListTopicsWithAssociatedResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTopicsWithAssociatedResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTopicsWithAssociatedResourcesInvoker) Invoke() (*model.ListTopicsWithAssociatedResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTopicsWithAssociatedResourcesResponse), nil
	}
}

type ListVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVersionInvoker) Invoke() (*model.ListVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVersionResponse), nil
	}
}

type ListVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVersionsInvoker) Invoke() (*model.ListVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVersionsResponse), nil
	}
}

type PublishHttpDetectInvoker struct {
	*invoker.BaseInvoker
}

func (i *PublishHttpDetectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PublishHttpDetectInvoker) Invoke() (*model.PublishHttpDetectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishHttpDetectResponse), nil
	}
}

type PublishMessageInvoker struct {
	*invoker.BaseInvoker
}

func (i *PublishMessageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PublishMessageInvoker) Invoke() (*model.PublishMessageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishMessageResponse), nil
	}
}

type ShowHttpDetectResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHttpDetectResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHttpDetectResultInvoker) Invoke() (*model.ShowHttpDetectResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHttpDetectResultResponse), nil
	}
}

type ShowKmsKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowKmsKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowKmsKeyInvoker) Invoke() (*model.ShowKmsKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowKmsKeyResponse), nil
	}
}

type ShowNotifyPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNotifyPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNotifyPolicyInvoker) Invoke() (*model.ShowNotifyPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNotifyPolicyResponse), nil
	}
}

type SubscribeTopicInvoker struct {
	*invoker.BaseInvoker
}

func (i *SubscribeTopicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SubscribeTopicInvoker) Invoke() (*model.SubscribeTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubscribeTopicResponse), nil
	}
}

type UnsubscribeSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UnsubscribeSubscriptionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UnsubscribeSubscriptionInvoker) Invoke() (*model.UnsubscribeSubscriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnsubscribeSubscriptionResponse), nil
	}
}

type UnsubscribeTopicInvoker struct {
	*invoker.BaseInvoker
}

func (i *UnsubscribeTopicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UnsubscribeTopicInvoker) Invoke() (*model.UnsubscribeTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnsubscribeTopicResponse), nil
	}
}

type UpdateKmsKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateKmsKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateKmsKeyInvoker) Invoke() (*model.UpdateKmsKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateKmsKeyResponse), nil
	}
}

type UpdateLogtankInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLogtankInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateLogtankInvoker) Invoke() (*model.UpdateLogtankResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLogtankResponse), nil
	}
}

type UpdateMessageTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMessageTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMessageTemplateInvoker) Invoke() (*model.UpdateMessageTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMessageTemplateResponse), nil
	}
}

type UpdateNotifyPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNotifyPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateNotifyPolicyInvoker) Invoke() (*model.UpdateNotifyPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNotifyPolicyResponse), nil
	}
}

type UpdateSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSubscriptionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSubscriptionInvoker) Invoke() (*model.UpdateSubscriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSubscriptionResponse), nil
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

type UpdateTopicAttributeInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTopicAttributeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTopicAttributeInvoker) Invoke() (*model.UpdateTopicAttributeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTopicAttributeResponse), nil
	}
}

type CreateApplicationInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *CreateApplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *CreateApplicationInvoker) Invoke() (*model.CreateApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateApplicationResponse), nil
	}
}

type DeleteApplicationInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DeleteApplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DeleteApplicationInvoker) Invoke() (*model.DeleteApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteApplicationResponse), nil
	}
}

type ListApplicationAttributesInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListApplicationAttributesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListApplicationAttributesInvoker) Invoke() (*model.ListApplicationAttributesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApplicationAttributesResponse), nil
	}
}

type ListApplicationsInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListApplicationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListApplicationsInvoker) Invoke() (*model.ListApplicationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApplicationsResponse), nil
	}
}

type PublishAppMessageInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *PublishAppMessageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *PublishAppMessageInvoker) Invoke() (*model.PublishAppMessageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishAppMessageResponse), nil
	}
}

type UpdateApplicationInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *UpdateApplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *UpdateApplicationInvoker) Invoke() (*model.UpdateApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateApplicationResponse), nil
	}
}

type CreateApplicationEndpointInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *CreateApplicationEndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *CreateApplicationEndpointInvoker) Invoke() (*model.CreateApplicationEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateApplicationEndpointResponse), nil
	}
}

type DeleteApplicationEndpointInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DeleteApplicationEndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DeleteApplicationEndpointInvoker) Invoke() (*model.DeleteApplicationEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteApplicationEndpointResponse), nil
	}
}

type ListApplicationEndpointAttributesInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListApplicationEndpointAttributesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListApplicationEndpointAttributesInvoker) Invoke() (*model.ListApplicationEndpointAttributesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApplicationEndpointAttributesResponse), nil
	}
}

type ListApplicationEndpointsInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListApplicationEndpointsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListApplicationEndpointsInvoker) Invoke() (*model.ListApplicationEndpointsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApplicationEndpointsResponse), nil
	}
}

type UpdateApplicationEndpointInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *UpdateApplicationEndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *UpdateApplicationEndpointInvoker) Invoke() (*model.UpdateApplicationEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateApplicationEndpointResponse), nil
	}
}
