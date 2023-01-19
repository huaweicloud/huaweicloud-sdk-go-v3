package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ccm/v1/model"
)

type CreateCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCertificateInvoker) Invoke() (*model.CreateCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCertificateResponse), nil
	}
}

type CreateCertificateAuthorityObsAgencyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCertificateAuthorityObsAgencyInvoker) Invoke() (*model.CreateCertificateAuthorityObsAgencyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCertificateAuthorityObsAgencyResponse), nil
	}
}

type CreateCertificateByCsrInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCertificateByCsrInvoker) Invoke() (*model.CreateCertificateByCsrResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCertificateByCsrResponse), nil
	}
}

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

type DisableCertificateAuthorityCrlInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableCertificateAuthorityCrlInvoker) Invoke() (*model.DisableCertificateAuthorityCrlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableCertificateAuthorityCrlResponse), nil
	}
}

type EnableCertificateAuthorityCrlInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableCertificateAuthorityCrlInvoker) Invoke() (*model.EnableCertificateAuthorityCrlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableCertificateAuthorityCrlResponse), nil
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

type ListCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCertificateInvoker) Invoke() (*model.ListCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCertificateResponse), nil
	}
}

type ListCertificateAuthorityObsBucketInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCertificateAuthorityObsBucketInvoker) Invoke() (*model.ListCertificateAuthorityObsBucketResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCertificateAuthorityObsBucketResponse), nil
	}
}

type ParseCertificateSigningRequestInvoker struct {
	*invoker.BaseInvoker
}

func (i *ParseCertificateSigningRequestInvoker) Invoke() (*model.ParseCertificateSigningRequestResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ParseCertificateSigningRequestResponse), nil
	}
}

type RevokeCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *RevokeCertificateInvoker) Invoke() (*model.RevokeCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RevokeCertificateResponse), nil
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

type ShowCertificateAuthorityObsAgencyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCertificateAuthorityObsAgencyInvoker) Invoke() (*model.ShowCertificateAuthorityObsAgencyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCertificateAuthorityObsAgencyResponse), nil
	}
}

type ShowCertificateQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCertificateQuotaInvoker) Invoke() (*model.ShowCertificateQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCertificateQuotaResponse), nil
	}
}

type CreateCertificateAuthorityInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCertificateAuthorityInvoker) Invoke() (*model.CreateCertificateAuthorityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCertificateAuthorityResponse), nil
	}
}

type DeleteCertificateAuthorityInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCertificateAuthorityInvoker) Invoke() (*model.DeleteCertificateAuthorityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCertificateAuthorityResponse), nil
	}
}

type DisableCertificateAuthorityInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableCertificateAuthorityInvoker) Invoke() (*model.DisableCertificateAuthorityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableCertificateAuthorityResponse), nil
	}
}

type EnableCertificateAuthorityInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableCertificateAuthorityInvoker) Invoke() (*model.EnableCertificateAuthorityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableCertificateAuthorityResponse), nil
	}
}

type ExportCertificateAuthorityCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportCertificateAuthorityCertificateInvoker) Invoke() (*model.ExportCertificateAuthorityCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportCertificateAuthorityCertificateResponse), nil
	}
}

type ExportCertificateAuthorityCsrInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportCertificateAuthorityCsrInvoker) Invoke() (*model.ExportCertificateAuthorityCsrResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportCertificateAuthorityCsrResponse), nil
	}
}

type ImportCertificateAuthorityCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportCertificateAuthorityCertificateInvoker) Invoke() (*model.ImportCertificateAuthorityCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportCertificateAuthorityCertificateResponse), nil
	}
}

type IssueCertificateAuthorityCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *IssueCertificateAuthorityCertificateInvoker) Invoke() (*model.IssueCertificateAuthorityCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.IssueCertificateAuthorityCertificateResponse), nil
	}
}

type ListCertificateAuthorityInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCertificateAuthorityInvoker) Invoke() (*model.ListCertificateAuthorityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCertificateAuthorityResponse), nil
	}
}

type RestoreCertificateAuthorityInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreCertificateAuthorityInvoker) Invoke() (*model.RestoreCertificateAuthorityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreCertificateAuthorityResponse), nil
	}
}

type RevokeCertificateAuthorityInvoker struct {
	*invoker.BaseInvoker
}

func (i *RevokeCertificateAuthorityInvoker) Invoke() (*model.RevokeCertificateAuthorityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RevokeCertificateAuthorityResponse), nil
	}
}

type ShowCertificateAuthorityInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCertificateAuthorityInvoker) Invoke() (*model.ShowCertificateAuthorityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCertificateAuthorityResponse), nil
	}
}

type ShowCertificateAuthorityQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCertificateAuthorityQuotaInvoker) Invoke() (*model.ShowCertificateAuthorityQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCertificateAuthorityQuotaResponse), nil
	}
}
