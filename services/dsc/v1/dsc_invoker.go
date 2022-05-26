package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dsc/v1/model"
)

type BatchAddDataMaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAddDataMaskInvoker) Invoke() (*model.BatchAddDataMaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAddDataMaskResponse), nil
	}
}

type CreateDatabaseWaterMarkInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDatabaseWaterMarkInvoker) Invoke() (*model.CreateDatabaseWaterMarkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDatabaseWaterMarkResponse), nil
	}
}

type CreateDocWatermarkInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDocWatermarkInvoker) Invoke() (*model.CreateDocWatermarkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDocWatermarkResponse), nil
	}
}

type CreateDocWatermarkByAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDocWatermarkByAddressInvoker) Invoke() (*model.CreateDocWatermarkByAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDocWatermarkByAddressResponse), nil
	}
}

type CreateImageWatermarkInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateImageWatermarkInvoker) Invoke() (*model.CreateImageWatermarkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateImageWatermarkResponse), nil
	}
}

type CreateImageWatermarkByAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateImageWatermarkByAddressInvoker) Invoke() (*model.CreateImageWatermarkByAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateImageWatermarkByAddressResponse), nil
	}
}

type ShowDatabaseWaterMarkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDatabaseWaterMarkInvoker) Invoke() (*model.ShowDatabaseWaterMarkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDatabaseWaterMarkResponse), nil
	}
}

type ShowDocWatermarkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDocWatermarkInvoker) Invoke() (*model.ShowDocWatermarkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDocWatermarkResponse), nil
	}
}

type ShowDocWatermarkByAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDocWatermarkByAddressInvoker) Invoke() (*model.ShowDocWatermarkByAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDocWatermarkByAddressResponse), nil
	}
}

type ShowImageWatermarkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowImageWatermarkInvoker) Invoke() (*model.ShowImageWatermarkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowImageWatermarkResponse), nil
	}
}

type ShowImageWatermarkByAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowImageWatermarkByAddressInvoker) Invoke() (*model.ShowImageWatermarkByAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowImageWatermarkByAddressResponse), nil
	}
}

type ShowImageWatermarkWithImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowImageWatermarkWithImageInvoker) Invoke() (*model.ShowImageWatermarkWithImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowImageWatermarkWithImageResponse), nil
	}
}

type ShowImageWatermarkWithImageByAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowImageWatermarkWithImageByAddressInvoker) Invoke() (*model.ShowImageWatermarkWithImageByAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowImageWatermarkWithImageByAddressResponse), nil
	}
}

type ShowScanJobResultsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowScanJobResultsInvoker) Invoke() (*model.ShowScanJobResultsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowScanJobResultsResponse), nil
	}
}

type ShowScanJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowScanJobsInvoker) Invoke() (*model.ShowScanJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowScanJobsResponse), nil
	}
}

type ShowOpenApiCalledRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOpenApiCalledRecordsInvoker) Invoke() (*model.ShowOpenApiCalledRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOpenApiCalledRecordsResponse), nil
	}
}
