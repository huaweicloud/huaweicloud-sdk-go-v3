package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/scm/v3/model"
)

type ApplyCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ApplyCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ApplyCertificateInvoker) Invoke() (*model.ApplyCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApplyCertificateResponse), nil
	}
}

type BatchPushCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchPushCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchPushCertificateInvoker) Invoke() (*model.BatchPushCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchPushCertificateResponse), nil
	}
}

type CancelCertificateRequestInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelCertificateRequestInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelCertificateRequestInvoker) Invoke() (*model.CancelCertificateRequestResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelCertificateRequestResponse), nil
	}
}

type DeleteCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCertificateInvoker) Invoke() (*model.DeleteCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCertificateResponse), nil
	}
}

type DeployCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeployCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeployCertificateInvoker) Invoke() (*model.DeployCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployCertificateResponse), nil
	}
}

type ExportCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ImportCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListCertificatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCertificatesInvoker) Invoke() (*model.ListCertificatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCertificatesResponse), nil
	}
}

type ListDeployedResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDeployedResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDeployedResourcesInvoker) Invoke() (*model.ListDeployedResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDeployedResourcesResponse), nil
	}
}

type PushCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *PushCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCertificateInvoker) Invoke() (*model.ShowCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCertificateResponse), nil
	}
}

type SubscribeCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *SubscribeCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SubscribeCertificateInvoker) Invoke() (*model.SubscribeCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SubscribeCertificateResponse), nil
	}
}

type UnsubscribeCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UnsubscribeCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UnsubscribeCertificateInvoker) Invoke() (*model.UnsubscribeCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnsubscribeCertificateResponse), nil
	}
}

type CreateCsrInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCsrInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCsrInvoker) Invoke() (*model.CreateCsrResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCsrResponse), nil
	}
}

type DeleteCsrInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCsrInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCsrInvoker) Invoke() (*model.DeleteCsrResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCsrResponse), nil
	}
}

type ListCsrInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCsrInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCsrInvoker) Invoke() (*model.ListCsrResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCsrResponse), nil
	}
}

type ShowCsrInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCsrInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCsrInvoker) Invoke() (*model.ShowCsrResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCsrResponse), nil
	}
}

type ShowCsrPrivateKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCsrPrivateKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCsrPrivateKeyInvoker) Invoke() (*model.ShowCsrPrivateKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCsrPrivateKeyResponse), nil
	}
}

type UpdateCsrInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCsrInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCsrInvoker) Invoke() (*model.UpdateCsrResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCsrResponse), nil
	}
}

type UploadCsrInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadCsrInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UploadCsrInvoker) Invoke() (*model.UploadCsrResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadCsrResponse), nil
	}
}
