package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dcos/v1/model"
)

type ListOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListOrderInvoker) Invoke() (*model.ListOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOrderResponse), nil
	}
}

type SaveOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *SaveOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SaveOrderInvoker) Invoke() (*model.SaveOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SaveOrderResponse), nil
	}
}

type ShowOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowOrderInvoker) Invoke() (*model.ShowOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOrderResponse), nil
	}
}

type ShowOrderCatalogueInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOrderCatalogueInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowOrderCatalogueInvoker) Invoke() (*model.ShowOrderCatalogueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOrderCatalogueResponse), nil
	}
}

type ShowOrderInformationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOrderInformationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowOrderInformationInvoker) Invoke() (*model.ShowOrderInformationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOrderInformationResponse), nil
	}
}

type ShowPageAssetListResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPageAssetListResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPageAssetListResultInvoker) Invoke() (*model.ShowPageAssetListResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPageAssetListResultResponse), nil
	}
}

type UploadFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UploadFileInvoker) Invoke() (*model.UploadFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadFileResponse), nil
	}
}

type VerifyOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *VerifyOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *VerifyOrderInvoker) Invoke() (*model.VerifyOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.VerifyOrderResponse), nil
	}
}
