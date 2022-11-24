package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dms/v2/model"
)

type BatchCreateOrDeleteQueueTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateOrDeleteQueueTagInvoker) Invoke() (*model.BatchCreateOrDeleteQueueTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateOrDeleteQueueTagResponse), nil
	}
}

type ConfirmConsumptionMessagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ConfirmConsumptionMessagesInvoker) Invoke() (*model.ConfirmConsumptionMessagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ConfirmConsumptionMessagesResponse), nil
	}
}

type ConfirmDeadLettersMessagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ConfirmDeadLettersMessagesInvoker) Invoke() (*model.ConfirmDeadLettersMessagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ConfirmDeadLettersMessagesResponse), nil
	}
}

type ConsumeDeadlettersMessageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ConsumeDeadlettersMessageInvoker) Invoke() (*model.ConsumeDeadlettersMessageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ConsumeDeadlettersMessageResponse), nil
	}
}

type ConsumeMessagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ConsumeMessagesInvoker) Invoke() (*model.ConsumeMessagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ConsumeMessagesResponse), nil
	}
}

type CreateConsumerGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateConsumerGroupInvoker) Invoke() (*model.CreateConsumerGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConsumerGroupResponse), nil
	}
}

type CreateQueueInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateQueueInvoker) Invoke() (*model.CreateQueueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateQueueResponse), nil
	}
}

type DeleteQueueInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteQueueInvoker) Invoke() (*model.DeleteQueueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteQueueResponse), nil
	}
}

type DeleteSpecifiedConsumerGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSpecifiedConsumerGroupInvoker) Invoke() (*model.DeleteSpecifiedConsumerGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSpecifiedConsumerGroupResponse), nil
	}
}

type ListConsumerGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConsumerGroupsInvoker) Invoke() (*model.ListConsumerGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConsumerGroupsResponse), nil
	}
}

type ListQueuesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQueuesInvoker) Invoke() (*model.ListQueuesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQueuesResponse), nil
	}
}

type SendMessagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *SendMessagesInvoker) Invoke() (*model.SendMessagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendMessagesResponse), nil
	}
}

type ShowQueueInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQueueInvoker) Invoke() (*model.ShowQueueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQueueResponse), nil
	}
}

type ShowQueueProjectTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQueueProjectTagsInvoker) Invoke() (*model.ShowQueueProjectTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQueueProjectTagsResponse), nil
	}
}

type ShowQueueTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQueueTagsInvoker) Invoke() (*model.ShowQueueTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQueueTagsResponse), nil
	}
}

type ShowQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuotasInvoker) Invoke() (*model.ShowQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuotasResponse), nil
	}
}
