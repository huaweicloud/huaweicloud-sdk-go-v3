package v3

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/scm/v3/model"
)

type DeleteCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCertificateInvoker) Invoke() (*model.DeleteCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCertificateResponse), nil
	}
}

type ExportCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportCertificateInvoker) Invoke() (*model.ExportCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportCertificateResponse), nil
	}
}

type ImportCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportCertificateInvoker) Invoke() (*model.ImportCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportCertificateResponse), nil
	}
}

type ListCertificatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCertificatesInvoker) Invoke() (*model.ListCertificatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCertificatesResponse), nil
	}
}

type PushCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *PushCertificateInvoker) Invoke() (*model.PushCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PushCertificateResponse), nil
	}
}

type ShowCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCertificateInvoker) Invoke() (*model.ShowCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCertificateResponse), nil
	}
}
