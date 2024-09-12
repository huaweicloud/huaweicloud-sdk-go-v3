package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kvs/v1/model"
)

type CreateTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTableInvoker) Invoke() (*model.CreateTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTableResponse), nil
	}
}

type DescribeTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *DescribeTableInvoker) Invoke() (*model.DescribeTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DescribeTableResponse), nil
	}
}

type ListStoreInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStoreInvoker) Invoke() (*model.ListStoreResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStoreResponse), nil
	}
}

type ListTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTableInvoker) Invoke() (*model.ListTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTableResponse), nil
	}
}

type BatchWriteKvInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchWriteKvInvoker) Invoke() (*model.BatchWriteKvResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchWriteKvResponse), nil
	}
}

type DeleteKvInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteKvInvoker) Invoke() (*model.DeleteKvResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteKvResponse), nil
	}
}

type GetKvInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetKvInvoker) Invoke() (*model.GetKvResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetKvResponse), nil
	}
}

type PutKvInvoker struct {
	*invoker.BaseInvoker
}

func (i *PutKvInvoker) Invoke() (*model.PutKvResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PutKvResponse), nil
	}
}

type ScanKvInvoker struct {
	*invoker.BaseInvoker
}

func (i *ScanKvInvoker) Invoke() (*model.ScanKvResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ScanKvResponse), nil
	}
}

type ScanSkeyKvInvoker struct {
	*invoker.BaseInvoker
}

func (i *ScanSkeyKvInvoker) Invoke() (*model.ScanSkeyKvResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ScanSkeyKvResponse), nil
	}
}

type UpdateKvInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateKvInvoker) Invoke() (*model.UpdateKvResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateKvResponse), nil
	}
}
