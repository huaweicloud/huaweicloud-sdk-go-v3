package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ccm/v1/model"
)

type BatchCreateCaTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateCaTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateCaTagsInvoker) Invoke() (*model.BatchCreateCaTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateCaTagsResponse), nil
	}
}

type BatchCreateCertTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateCertTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateCertTagsInvoker) Invoke() (*model.BatchCreateCertTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateCertTagsResponse), nil
	}
}

type BatchDeleteCaTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteCaTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteCaTagsInvoker) Invoke() (*model.BatchDeleteCaTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteCaTagsResponse), nil
	}
}

type BatchDeleteCertTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteCertTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteCertTagsInvoker) Invoke() (*model.BatchDeleteCertTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteCertTagsResponse), nil
	}
}

type CountCaResourceInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountCaResourceInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountCaResourceInstancesInvoker) Invoke() (*model.CountCaResourceInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountCaResourceInstancesResponse), nil
	}
}

type CountCertResourceInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountCertResourceInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountCertResourceInstancesInvoker) Invoke() (*model.CountCertResourceInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountCertResourceInstancesResponse), nil
	}
}

type CreateCaTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCaTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCaTagInvoker) Invoke() (*model.CreateCaTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCaTagResponse), nil
	}
}

type CreateCertTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCertTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCertTagInvoker) Invoke() (*model.CreateCertTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCertTagResponse), nil
	}
}

type CreateCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateCertificateAuthorityObsAgencyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCertificateAuthorityObsAgencyInvoker) Invoke() (*model.CreateCertificateAuthorityObsAgencyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCertificateAuthorityObsAgencyResponse), nil
	}
}

type CreateCertificateAuthorityOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCertificateAuthorityOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCertificateAuthorityOrderInvoker) Invoke() (*model.CreateCertificateAuthorityOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCertificateAuthorityOrderResponse), nil
	}
}

type CreateCertificateByCsrInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCertificateByCsrInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

type DisableCertificateAuthorityCrlInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableCertificateAuthorityCrlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *EnableCertificateAuthorityCrlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

type ListCaResourceInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCaResourceInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCaResourceInstancesInvoker) Invoke() (*model.ListCaResourceInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCaResourceInstancesResponse), nil
	}
}

type ListCaTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCaTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCaTagsInvoker) Invoke() (*model.ListCaTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCaTagsResponse), nil
	}
}

type ListCertResourceInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCertResourceInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCertResourceInstancesInvoker) Invoke() (*model.ListCertResourceInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCertResourceInstancesResponse), nil
	}
}

type ListCertTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCertTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCertTagsInvoker) Invoke() (*model.ListCertTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCertTagsResponse), nil
	}
}

type ListCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListCertificateAuthorityObsBucketInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCertificateAuthorityObsBucketInvoker) Invoke() (*model.ListCertificateAuthorityObsBucketResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCertificateAuthorityObsBucketResponse), nil
	}
}

type ListDomainCaTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDomainCaTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDomainCaTagsInvoker) Invoke() (*model.ListDomainCaTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDomainCaTagsResponse), nil
	}
}

type ListDomainCertTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDomainCertTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDomainCertTagsInvoker) Invoke() (*model.ListDomainCertTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDomainCertTagsResponse), nil
	}
}

type ParseCertificateSigningRequestInvoker struct {
	*invoker.BaseInvoker
}

func (i *ParseCertificateSigningRequestInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *RevokeCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

type ShowCertificateAuthorityObsAgencyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCertificateAuthorityObsAgencyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowCertificateQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateCertificateAuthorityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteCertificateAuthorityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DisableCertificateAuthorityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *EnableCertificateAuthorityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ExportCertificateAuthorityCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ExportCertificateAuthorityCsrInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ImportCertificateAuthorityCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *IssueCertificateAuthorityCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListCertificateAuthorityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *RestoreCertificateAuthorityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *RevokeCertificateAuthorityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowCertificateAuthorityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowCertificateAuthorityQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCertificateAuthorityQuotaInvoker) Invoke() (*model.ShowCertificateAuthorityQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCertificateAuthorityQuotaResponse), nil
	}
}
