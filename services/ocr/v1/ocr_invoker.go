package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ocr/v1/model"
)

type RecognizeAcceptanceBillInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeAcceptanceBillInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeAcceptanceBillInvoker) Invoke() (*model.RecognizeAcceptanceBillResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeAcceptanceBillResponse), nil
	}
}

type RecognizeAutoClassificationInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeAutoClassificationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeAutoClassificationInvoker) Invoke() (*model.RecognizeAutoClassificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeAutoClassificationResponse), nil
	}
}

type RecognizeBankReceiptInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeBankReceiptInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeBankReceiptInvoker) Invoke() (*model.RecognizeBankReceiptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeBankReceiptResponse), nil
	}
}

type RecognizeBankcardInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeBankcardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeBankcardInvoker) Invoke() (*model.RecognizeBankcardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeBankcardResponse), nil
	}
}

type RecognizeBusinessCardInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeBusinessCardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeBusinessCardInvoker) Invoke() (*model.RecognizeBusinessCardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeBusinessCardResponse), nil
	}
}

type RecognizeBusinessLicenseInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeBusinessLicenseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeBusinessLicenseInvoker) Invoke() (*model.RecognizeBusinessLicenseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeBusinessLicenseResponse), nil
	}
}

type RecognizeCambodianIdCardInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeCambodianIdCardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeCambodianIdCardInvoker) Invoke() (*model.RecognizeCambodianIdCardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeCambodianIdCardResponse), nil
	}
}

type RecognizeChileIdCardInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeChileIdCardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeChileIdCardInvoker) Invoke() (*model.RecognizeChileIdCardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeChileIdCardResponse), nil
	}
}

type RecognizeColombiaIdCardInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeColombiaIdCardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeColombiaIdCardInvoker) Invoke() (*model.RecognizeColombiaIdCardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeColombiaIdCardResponse), nil
	}
}

type RecognizeDriverLicenseInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeDriverLicenseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeDriverLicenseInvoker) Invoke() (*model.RecognizeDriverLicenseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeDriverLicenseResponse), nil
	}
}

type RecognizeExitEntryPermitInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeExitEntryPermitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeExitEntryPermitInvoker) Invoke() (*model.RecognizeExitEntryPermitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeExitEntryPermitResponse), nil
	}
}

type RecognizeFinancialStatementInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeFinancialStatementInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeFinancialStatementInvoker) Invoke() (*model.RecognizeFinancialStatementResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeFinancialStatementResponse), nil
	}
}

type RecognizeFlightItineraryInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeFlightItineraryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeFlightItineraryInvoker) Invoke() (*model.RecognizeFlightItineraryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeFlightItineraryResponse), nil
	}
}

type RecognizeGeneralTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeGeneralTableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeGeneralTableInvoker) Invoke() (*model.RecognizeGeneralTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeGeneralTableResponse), nil
	}
}

type RecognizeGeneralTextInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeGeneralTextInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeGeneralTextInvoker) Invoke() (*model.RecognizeGeneralTextResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeGeneralTextResponse), nil
	}
}

type RecognizeHandwritingInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeHandwritingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeHandwritingInvoker) Invoke() (*model.RecognizeHandwritingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeHandwritingResponse), nil
	}
}

type RecognizeHealthCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeHealthCodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeHealthCodeInvoker) Invoke() (*model.RecognizeHealthCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeHealthCodeResponse), nil
	}
}

type RecognizeHkIdCardInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeHkIdCardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeHkIdCardInvoker) Invoke() (*model.RecognizeHkIdCardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeHkIdCardResponse), nil
	}
}

type RecognizeHouseholdRegisterInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeHouseholdRegisterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeHouseholdRegisterInvoker) Invoke() (*model.RecognizeHouseholdRegisterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeHouseholdRegisterResponse), nil
	}
}

type RecognizeIdCardInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeIdCardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeIdCardInvoker) Invoke() (*model.RecognizeIdCardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeIdCardResponse), nil
	}
}

type RecognizeIdDocumentInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeIdDocumentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeIdDocumentInvoker) Invoke() (*model.RecognizeIdDocumentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeIdDocumentResponse), nil
	}
}

type RecognizeInsurancePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeInsurancePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeInsurancePolicyInvoker) Invoke() (*model.RecognizeInsurancePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeInsurancePolicyResponse), nil
	}
}

type RecognizeInvoiceVerificationInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeInvoiceVerificationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeInvoiceVerificationInvoker) Invoke() (*model.RecognizeInvoiceVerificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeInvoiceVerificationResponse), nil
	}
}

type RecognizeLicensePlateInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeLicensePlateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeLicensePlateInvoker) Invoke() (*model.RecognizeLicensePlateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeLicensePlateResponse), nil
	}
}

type RecognizeMacaoIdCardInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeMacaoIdCardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeMacaoIdCardInvoker) Invoke() (*model.RecognizeMacaoIdCardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeMacaoIdCardResponse), nil
	}
}

type RecognizeMainlandTravelPermitInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeMainlandTravelPermitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeMainlandTravelPermitInvoker) Invoke() (*model.RecognizeMainlandTravelPermitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeMainlandTravelPermitResponse), nil
	}
}

type RecognizeMvsInvoiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeMvsInvoiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeMvsInvoiceInvoker) Invoke() (*model.RecognizeMvsInvoiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeMvsInvoiceResponse), nil
	}
}

type RecognizeMyanmarDriverLicenseInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeMyanmarDriverLicenseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeMyanmarDriverLicenseInvoker) Invoke() (*model.RecognizeMyanmarDriverLicenseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeMyanmarDriverLicenseResponse), nil
	}
}

type RecognizeMyanmarIdcardInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeMyanmarIdcardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeMyanmarIdcardInvoker) Invoke() (*model.RecognizeMyanmarIdcardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeMyanmarIdcardResponse), nil
	}
}

type RecognizePassportInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizePassportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizePassportInvoker) Invoke() (*model.RecognizePassportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizePassportResponse), nil
	}
}

type RecognizePcrTestRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizePcrTestRecordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizePcrTestRecordInvoker) Invoke() (*model.RecognizePcrTestRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizePcrTestRecordResponse), nil
	}
}

type RecognizePeruIdCardInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizePeruIdCardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizePeruIdCardInvoker) Invoke() (*model.RecognizePeruIdCardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizePeruIdCardResponse), nil
	}
}

type RecognizeQualificationCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeQualificationCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeQualificationCertificateInvoker) Invoke() (*model.RecognizeQualificationCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeQualificationCertificateResponse), nil
	}
}

type RecognizeQuotaInvoiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeQuotaInvoiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeQuotaInvoiceInvoker) Invoke() (*model.RecognizeQuotaInvoiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeQuotaInvoiceResponse), nil
	}
}

type RecognizeRealEstateCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeRealEstateCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeRealEstateCertificateInvoker) Invoke() (*model.RecognizeRealEstateCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeRealEstateCertificateResponse), nil
	}
}

type RecognizeSealInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeSealInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeSealInvoker) Invoke() (*model.RecognizeSealResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeSealResponse), nil
	}
}

type RecognizeSmartDocumentRecognizerInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeSmartDocumentRecognizerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeSmartDocumentRecognizerInvoker) Invoke() (*model.RecognizeSmartDocumentRecognizerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeSmartDocumentRecognizerResponse), nil
	}
}

type RecognizeTaxiInvoiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeTaxiInvoiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeTaxiInvoiceInvoker) Invoke() (*model.RecognizeTaxiInvoiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeTaxiInvoiceResponse), nil
	}
}

type RecognizeThailandIdcardInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeThailandIdcardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeThailandIdcardInvoker) Invoke() (*model.RecognizeThailandIdcardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeThailandIdcardResponse), nil
	}
}

type RecognizeThailandLicensePlateInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeThailandLicensePlateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeThailandLicensePlateInvoker) Invoke() (*model.RecognizeThailandLicensePlateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeThailandLicensePlateResponse), nil
	}
}

type RecognizeTollInvoiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeTollInvoiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeTollInvoiceInvoker) Invoke() (*model.RecognizeTollInvoiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeTollInvoiceResponse), nil
	}
}

type RecognizeTrainTicketInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeTrainTicketInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeTrainTicketInvoker) Invoke() (*model.RecognizeTrainTicketResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeTrainTicketResponse), nil
	}
}

type RecognizeTransportationLicenseInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeTransportationLicenseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeTransportationLicenseInvoker) Invoke() (*model.RecognizeTransportationLicenseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeTransportationLicenseResponse), nil
	}
}

type RecognizeVatInvoiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeVatInvoiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeVatInvoiceInvoker) Invoke() (*model.RecognizeVatInvoiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeVatInvoiceResponse), nil
	}
}

type RecognizeVehicleCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeVehicleCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeVehicleCertificateInvoker) Invoke() (*model.RecognizeVehicleCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeVehicleCertificateResponse), nil
	}
}

type RecognizeVehicleLicenseInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeVehicleLicenseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeVehicleLicenseInvoker) Invoke() (*model.RecognizeVehicleLicenseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeVehicleLicenseResponse), nil
	}
}

type RecognizeVietnamIdCardInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeVietnamIdCardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeVietnamIdCardInvoker) Invoke() (*model.RecognizeVietnamIdCardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeVietnamIdCardResponse), nil
	}
}

type RecognizeWaybillElectronicInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeWaybillElectronicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeWaybillElectronicInvoker) Invoke() (*model.RecognizeWaybillElectronicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeWaybillElectronicResponse), nil
	}
}

type RecognizeWebImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeWebImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeWebImageInvoker) Invoke() (*model.RecognizeWebImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeWebImageResponse), nil
	}
}

type RecognizeCustomTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeCustomTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeCustomTemplateInvoker) Invoke() (*model.RecognizeCustomTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeCustomTemplateResponse), nil
	}
}

type RecognizeVinInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeVinInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecognizeVinInvoker) Invoke() (*model.RecognizeVinResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeVinResponse), nil
	}
}
