package v5

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cce/v5/model"
)

type CreateImageCacheInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateImageCacheInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateImageCacheInvoker) Invoke() (*model.CreateImageCacheResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateImageCacheResponse), nil
	}
}

type DeleteImageCacheInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteImageCacheInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteImageCacheInvoker) Invoke() (*model.DeleteImageCacheResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteImageCacheResponse), nil
	}
}

type ListImageCachesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageCachesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageCachesInvoker) Invoke() (*model.ListImageCachesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageCachesResponse), nil
	}
}

type ShowImageCacheInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowImageCacheInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowImageCacheInvoker) Invoke() (*model.ShowImageCacheResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowImageCacheResponse), nil
	}
}
