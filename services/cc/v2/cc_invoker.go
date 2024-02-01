package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cc/v2/model"
)

type BatchCreateGcbResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateGcbResourceTagsInvoker) Invoke() (*model.BatchCreateGcbResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateGcbResourceTagsResponse), nil
	}
}

type BatchDeleteGcbResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteGcbResourceTagsInvoker) Invoke() (*model.BatchDeleteGcbResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteGcbResourceTagsResponse), nil
	}
}

type CountGcbResourceByTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountGcbResourceByTagInvoker) Invoke() (*model.CountGcbResourceByTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountGcbResourceByTagResponse), nil
	}
}

type CreateGcbResourceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGcbResourceTagInvoker) Invoke() (*model.CreateGcbResourceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGcbResourceTagResponse), nil
	}
}

type DeleteGcbResourceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGcbResourceTagInvoker) Invoke() (*model.DeleteGcbResourceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGcbResourceTagResponse), nil
	}
}

type ListGcbResourceByTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGcbResourceByTagInvoker) Invoke() (*model.ListGcbResourceByTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGcbResourceByTagResponse), nil
	}
}

type ListGcbResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGcbResourceTagsInvoker) Invoke() (*model.ListGcbResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGcbResourceTagsResponse), nil
	}
}

type ListGcbTenantTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGcbTenantTagsInvoker) Invoke() (*model.ListGcbTenantTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGcbTenantTagsResponse), nil
	}
}

type BatchCreateDeleteTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateDeleteTagsInvoker) Invoke() (*model.BatchCreateDeleteTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateDeleteTagsResponse), nil
	}
}

type CreateTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTagInvoker) Invoke() (*model.CreateTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTagResponse), nil
	}
}

type DeleteTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTagInvoker) Invoke() (*model.DeleteTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTagResponse), nil
	}
}

type ListDomainTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDomainTagsInvoker) Invoke() (*model.ListDomainTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDomainTagsResponse), nil
	}
}

type ListResourceByFilterTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceByFilterTagInvoker) Invoke() (*model.ListResourceByFilterTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceByFilterTagResponse), nil
	}
}

type ListTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagsInvoker) Invoke() (*model.ListTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagsResponse), nil
	}
}
