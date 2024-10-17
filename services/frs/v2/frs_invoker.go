package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/frs/v2/model"
)

type AddFacesByBase64Invoker struct {
	*invoker.BaseInvoker
}

func (i *AddFacesByBase64Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddFacesByBase64Invoker) Invoke() (*model.AddFacesByBase64Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddFacesByBase64Response), nil
	}
}

type AddFacesByFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddFacesByFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddFacesByFileInvoker) Invoke() (*model.AddFacesByFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddFacesByFileResponse), nil
	}
}

type AddFacesByUrlInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddFacesByUrlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddFacesByUrlInvoker) Invoke() (*model.AddFacesByUrlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddFacesByUrlResponse), nil
	}
}

type BatchDeleteFacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteFacesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteFacesInvoker) Invoke() (*model.BatchDeleteFacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteFacesResponse), nil
	}
}

type CompareFaceByBase64Invoker struct {
	*invoker.BaseInvoker
}

func (i *CompareFaceByBase64Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CompareFaceByBase64Invoker) Invoke() (*model.CompareFaceByBase64Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CompareFaceByBase64Response), nil
	}
}

type CompareFaceByFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *CompareFaceByFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CompareFaceByFileInvoker) Invoke() (*model.CompareFaceByFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CompareFaceByFileResponse), nil
	}
}

type CompareFaceByUrlInvoker struct {
	*invoker.BaseInvoker
}

func (i *CompareFaceByUrlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CompareFaceByUrlInvoker) Invoke() (*model.CompareFaceByUrlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CompareFaceByUrlResponse), nil
	}
}

type CreateFaceSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFaceSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFaceSetInvoker) Invoke() (*model.CreateFaceSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFaceSetResponse), nil
	}
}

type DeleteFaceByExternalImageIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFaceByExternalImageIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteFaceByExternalImageIdInvoker) Invoke() (*model.DeleteFaceByExternalImageIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFaceByExternalImageIdResponse), nil
	}
}

type DeleteFaceByFaceIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFaceByFaceIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteFaceByFaceIdInvoker) Invoke() (*model.DeleteFaceByFaceIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFaceByFaceIdResponse), nil
	}
}

type DeleteFaceSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFaceSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteFaceSetInvoker) Invoke() (*model.DeleteFaceSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFaceSetResponse), nil
	}
}

type DetectFaceByBase64Invoker struct {
	*invoker.BaseInvoker
}

func (i *DetectFaceByBase64Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetectFaceByBase64Invoker) Invoke() (*model.DetectFaceByBase64Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetectFaceByBase64Response), nil
	}
}

type DetectFaceByBase64IntlInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetectFaceByBase64IntlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetectFaceByBase64IntlInvoker) Invoke() (*model.DetectFaceByBase64IntlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetectFaceByBase64IntlResponse), nil
	}
}

type DetectFaceByFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetectFaceByFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetectFaceByFileInvoker) Invoke() (*model.DetectFaceByFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetectFaceByFileResponse), nil
	}
}

type DetectFaceByFileIntlInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetectFaceByFileIntlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetectFaceByFileIntlInvoker) Invoke() (*model.DetectFaceByFileIntlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetectFaceByFileIntlResponse), nil
	}
}

type DetectFaceByUrlInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetectFaceByUrlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetectFaceByUrlInvoker) Invoke() (*model.DetectFaceByUrlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetectFaceByUrlResponse), nil
	}
}

type DetectFaceByUrlIntlInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetectFaceByUrlIntlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetectFaceByUrlIntlInvoker) Invoke() (*model.DetectFaceByUrlIntlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetectFaceByUrlIntlResponse), nil
	}
}

type DetectLiveByBase64Invoker struct {
	*invoker.BaseInvoker
}

func (i *DetectLiveByBase64Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetectLiveByBase64Invoker) Invoke() (*model.DetectLiveByBase64Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetectLiveByBase64Response), nil
	}
}

type DetectLiveByBase64IntlInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetectLiveByBase64IntlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetectLiveByBase64IntlInvoker) Invoke() (*model.DetectLiveByBase64IntlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetectLiveByBase64IntlResponse), nil
	}
}

type DetectLiveByFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetectLiveByFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetectLiveByFileInvoker) Invoke() (*model.DetectLiveByFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetectLiveByFileResponse), nil
	}
}

type DetectLiveByFileIntlInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetectLiveByFileIntlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetectLiveByFileIntlInvoker) Invoke() (*model.DetectLiveByFileIntlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetectLiveByFileIntlResponse), nil
	}
}

type DetectLiveByUrlInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetectLiveByUrlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetectLiveByUrlInvoker) Invoke() (*model.DetectLiveByUrlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetectLiveByUrlResponse), nil
	}
}

type DetectLiveByUrlIntlInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetectLiveByUrlIntlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetectLiveByUrlIntlInvoker) Invoke() (*model.DetectLiveByUrlIntlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetectLiveByUrlIntlResponse), nil
	}
}

type DetectLiveFaceByBase64Invoker struct {
	*invoker.BaseInvoker
}

func (i *DetectLiveFaceByBase64Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetectLiveFaceByBase64Invoker) Invoke() (*model.DetectLiveFaceByBase64Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetectLiveFaceByBase64Response), nil
	}
}

type DetectLiveFaceByFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetectLiveFaceByFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetectLiveFaceByFileInvoker) Invoke() (*model.DetectLiveFaceByFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetectLiveFaceByFileResponse), nil
	}
}

type DetectLiveFaceByUrlInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetectLiveFaceByUrlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetectLiveFaceByUrlInvoker) Invoke() (*model.DetectLiveFaceByUrlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetectLiveFaceByUrlResponse), nil
	}
}

type SearchFaceByBase64Invoker struct {
	*invoker.BaseInvoker
}

func (i *SearchFaceByBase64Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchFaceByBase64Invoker) Invoke() (*model.SearchFaceByBase64Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchFaceByBase64Response), nil
	}
}

type SearchFaceByFaceIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchFaceByFaceIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchFaceByFaceIdInvoker) Invoke() (*model.SearchFaceByFaceIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchFaceByFaceIdResponse), nil
	}
}

type SearchFaceByFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchFaceByFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchFaceByFileInvoker) Invoke() (*model.SearchFaceByFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchFaceByFileResponse), nil
	}
}

type SearchFaceByUrlInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchFaceByUrlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchFaceByUrlInvoker) Invoke() (*model.SearchFaceByUrlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchFaceByUrlResponse), nil
	}
}

type ShowAllFaceSetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAllFaceSetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAllFaceSetsInvoker) Invoke() (*model.ShowAllFaceSetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAllFaceSetsResponse), nil
	}
}

type ShowFaceSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFaceSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFaceSetInvoker) Invoke() (*model.ShowFaceSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFaceSetResponse), nil
	}
}

type ShowFacesByFaceIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFacesByFaceIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFacesByFaceIdInvoker) Invoke() (*model.ShowFacesByFaceIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFacesByFaceIdResponse), nil
	}
}

type ShowFacesByLimitInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFacesByLimitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFacesByLimitInvoker) Invoke() (*model.ShowFacesByLimitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFacesByLimitResponse), nil
	}
}

type UpdateFaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateFaceInvoker) Invoke() (*model.UpdateFaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFaceResponse), nil
	}
}
