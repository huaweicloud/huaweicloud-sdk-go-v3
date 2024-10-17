package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ces/v3/model"
)

type ListAgentStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAgentStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAgentStatusInvoker) Invoke() (*model.ListAgentStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAgentStatusResponse), nil
	}
}

type BatchCreateAgentInvocationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateAgentInvocationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateAgentInvocationsInvoker) Invoke() (*model.BatchCreateAgentInvocationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateAgentInvocationsResponse), nil
	}
}

type ListAgentInvocationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAgentInvocationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAgentInvocationsInvoker) Invoke() (*model.ListAgentInvocationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAgentInvocationsResponse), nil
	}
}
