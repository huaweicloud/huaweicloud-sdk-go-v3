package v4

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/aom/v4/model"
)

type BatchImportAgentInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchImportAgentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchImportAgentInvoker) Invoke() (*model.BatchImportAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchImportAgentResponse), nil
	}
}

type BatchUpdateAgentInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateAgentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdateAgentInvoker) Invoke() (*model.BatchUpdateAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateAgentResponse), nil
	}
}

type ShowAgentInfosInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAgentInfosInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAgentInfosInvoker) Invoke() (*model.ShowAgentInfosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAgentInfosResponse), nil
	}
}
