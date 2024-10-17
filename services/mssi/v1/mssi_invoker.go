package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/mssi/v1/model"
)

type CreateConnectionInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateConnectionInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateConnectionInfoInvoker) Invoke() (*model.CreateConnectionInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConnectionInfoResponse), nil
	}
}

type CreateCustomConnectorFromOpenapiInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCustomConnectorFromOpenapiInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCustomConnectorFromOpenapiInvoker) Invoke() (*model.CreateCustomConnectorFromOpenapiResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCustomConnectorFromOpenapiResponse), nil
	}
}

type CreateFlowInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFlowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFlowInvoker) Invoke() (*model.CreateFlowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFlowResponse), nil
	}
}

type CreateFlowTemplateFromFlowInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFlowTemplateFromFlowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFlowTemplateFromFlowInvoker) Invoke() (*model.CreateFlowTemplateFromFlowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFlowTemplateFromFlowResponse), nil
	}
}

type DeleteConnectionInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteConnectionInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteConnectionInfoInvoker) Invoke() (*model.DeleteConnectionInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteConnectionInfoResponse), nil
	}
}

type DeleteCustomConnectorInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCustomConnectorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCustomConnectorInvoker) Invoke() (*model.DeleteCustomConnectorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCustomConnectorResponse), nil
	}
}

type DeleteFlowInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFlowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteFlowInvoker) Invoke() (*model.DeleteFlowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFlowResponse), nil
	}
}

type SearchFlowByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchFlowByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchFlowByIdInvoker) Invoke() (*model.SearchFlowByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchFlowByIdResponse), nil
	}
}

type ShowAllConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAllConnectionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAllConnectionsInvoker) Invoke() (*model.ShowAllConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAllConnectionsResponse), nil
	}
}

type ShowAllFlowsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAllFlowsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAllFlowsInvoker) Invoke() (*model.ShowAllFlowsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAllFlowsResponse), nil
	}
}

type ShowConnectorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConnectorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowConnectorsInvoker) Invoke() (*model.ShowConnectorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConnectorsResponse), nil
	}
}

type ShowCustomConnectorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCustomConnectorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCustomConnectorInvoker) Invoke() (*model.ShowCustomConnectorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCustomConnectorResponse), nil
	}
}

type ShowCustomConnectorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCustomConnectorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCustomConnectorsInvoker) Invoke() (*model.ShowCustomConnectorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCustomConnectorsResponse), nil
	}
}

type ShowSingleConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSingleConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSingleConnectionInvoker) Invoke() (*model.ShowSingleConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSingleConnectionResponse), nil
	}
}

type UpdateConnectionInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateConnectionInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateConnectionInfoInvoker) Invoke() (*model.UpdateConnectionInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateConnectionInfoResponse), nil
	}
}

type UpdateFlowInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFlowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateFlowInvoker) Invoke() (*model.UpdateFlowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFlowResponse), nil
	}
}
