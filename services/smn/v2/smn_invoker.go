package v2

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/smn/v2/model"
)

type AddSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddSubscriptionInvoker) Invoke() (*model.AddSubscriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddSubscriptionResponse), nil
	}
}

type BatchCreateOrDeleteResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateOrDeleteResourceTagsInvoker) Invoke() (*model.BatchCreateOrDeleteResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateOrDeleteResourceTagsResponse), nil
	}
}

type CancelSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelSubscriptionInvoker) Invoke() (*model.CancelSubscriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelSubscriptionResponse), nil
	}
}

type CreateLogtankInvoker struct {
	*invoker.BaseInvoker
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

func (i *CreateMessageTemplateInvoker) Invoke() (*model.CreateMessageTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMessageTemplateResponse), nil
	}
}

type CreateResourceTagInvoker struct {
	*invoker.BaseInvoker
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

func (i *CreateTopicInvoker) Invoke() (*model.CreateTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTopicResponse), nil
	}
}

type DeleteLogtankInvoker struct {
	*invoker.BaseInvoker
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

func (i *DeleteMessageTemplateInvoker) Invoke() (*model.DeleteMessageTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMessageTemplateResponse), nil
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

type DeleteTopicAttributeByNameInvoker struct {
	*invoker.BaseInvoker
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

func (i *DeleteTopicAttributesInvoker) Invoke() (*model.DeleteTopicAttributesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTopicAttributesResponse), nil
	}
}

type ListLogtankInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListProjectTagsInvoker) Invoke() (*model.ListProjectTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectTagsResponse), nil
	}
}

type ListResourceInstancesInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListTopicDetailsInvoker) Invoke() (*model.ListTopicDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTopicDetailsResponse), nil
	}
}

type ListTopicsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTopicsInvoker) Invoke() (*model.ListTopicsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTopicsResponse), nil
	}
}

type ListVersionInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListVersionsInvoker) Invoke() (*model.ListVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVersionsResponse), nil
	}
}

type PublishMessageInvoker struct {
	*invoker.BaseInvoker
}

func (i *PublishMessageInvoker) Invoke() (*model.PublishMessageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishMessageResponse), nil
	}
}

type UpdateLogtankInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdateMessageTemplateInvoker) Invoke() (*model.UpdateMessageTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMessageTemplateResponse), nil
	}
}

type UpdateSubscriptionInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdateApplicationEndpointInvoker) Invoke() (*model.UpdateApplicationEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateApplicationEndpointResponse), nil
	}
}
