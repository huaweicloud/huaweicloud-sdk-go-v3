package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/koomessage/v1/model"
)

type AddCallBackInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddCallBackInvoker) Invoke() (*model.AddCallBackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddCallBackResponse), nil
	}
}

type ListAimCallbacksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAimCallbacksInvoker) Invoke() (*model.ListAimCallbacksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAimCallbacksResponse), nil
	}
}

type CheckMobileCapabilityInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckMobileCapabilityInvoker) Invoke() (*model.CheckMobileCapabilityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckMobileCapabilityResponse), nil
	}
}

type CreateResolveTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResolveTaskInvoker) Invoke() (*model.CreateResolveTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResolveTaskResponse), nil
	}
}

type ListAimResolveDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAimResolveDetailsInvoker) Invoke() (*model.ListAimResolveDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAimResolveDetailsResponse), nil
	}
}

type ListResolveTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResolveTasksInvoker) Invoke() (*model.ListResolveTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResolveTasksResponse), nil
	}
}

type CreateAimSendTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAimSendTaskInvoker) Invoke() (*model.CreateAimSendTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAimSendTaskResponse), nil
	}
}

type ListAimSendDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAimSendDetailsInvoker) Invoke() (*model.ListAimSendDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAimSendDetailsResponse), nil
	}
}

type ListAimSendReportsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAimSendReportsInvoker) Invoke() (*model.ListAimSendReportsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAimSendReportsResponse), nil
	}
}

type ListAimSendTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAimSendTasksInvoker) Invoke() (*model.ListAimSendTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAimSendTasksResponse), nil
	}
}

type CreateAimPersonalTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAimPersonalTemplateInvoker) Invoke() (*model.CreateAimPersonalTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAimPersonalTemplateResponse), nil
	}
}

type DeleteAimPersonalTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAimPersonalTemplateInvoker) Invoke() (*model.DeleteAimPersonalTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAimPersonalTemplateResponse), nil
	}
}

type DeleteTemplateMaterialInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTemplateMaterialInvoker) Invoke() (*model.DeleteTemplateMaterialResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTemplateMaterialResponse), nil
	}
}

type ListAimTemplateMaterialsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAimTemplateMaterialsInvoker) Invoke() (*model.ListAimTemplateMaterialsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAimTemplateMaterialsResponse), nil
	}
}

type ListAimTemplateReportsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAimTemplateReportsInvoker) Invoke() (*model.ListAimTemplateReportsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAimTemplateReportsResponse), nil
	}
}

type ListAimTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAimTemplatesInvoker) Invoke() (*model.ListAimTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAimTemplatesResponse), nil
	}
}

type UpdatePersonalTemplateStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePersonalTemplateStateInvoker) Invoke() (*model.UpdatePersonalTemplateStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePersonalTemplateStateResponse), nil
	}
}

type UploadAimTemplateMaterialInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadAimTemplateMaterialInvoker) Invoke() (*model.UploadAimTemplateMaterialResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadAimTemplateMaterialResponse), nil
	}
}

type ListMenusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMenusInvoker) Invoke() (*model.ListMenusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMenusResponse), nil
	}
}

type UpdateMenuInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMenuInvoker) Invoke() (*model.UpdateMenuResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMenuResponse), nil
	}
}

type DeletePortInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePortInfoInvoker) Invoke() (*model.DeletePortInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePortInfoResponse), nil
	}
}

type ListPortInfosInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPortInfosInvoker) Invoke() (*model.ListPortInfosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPortInfosResponse), nil
	}
}

type LockPortInvoker struct {
	*invoker.BaseInvoker
}

func (i *LockPortInvoker) Invoke() (*model.LockPortResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.LockPortResponse), nil
	}
}

type RegisterPortInvoker struct {
	*invoker.BaseInvoker
}

func (i *RegisterPortInvoker) Invoke() (*model.RegisterPortResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterPortResponse), nil
	}
}

type UnlockPortInvoker struct {
	*invoker.BaseInvoker
}

func (i *UnlockPortInvoker) Invoke() (*model.UnlockPortResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnlockPortResponse), nil
	}
}

type ListPortalInfosInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPortalInfosInvoker) Invoke() (*model.ListPortalInfosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPortalInfosResponse), nil
	}
}

type UpdatePortalInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePortalInfoInvoker) Invoke() (*model.UpdatePortalInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePortalInfoResponse), nil
	}
}

type FreezePubInvoker struct {
	*invoker.BaseInvoker
}

func (i *FreezePubInvoker) Invoke() (*model.FreezePubResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.FreezePubResponse), nil
	}
}

type ListPubInfosInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPubInfosInvoker) Invoke() (*model.ListPubInfosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPubInfosResponse), nil
	}
}

type UnfreezePubInvoker struct {
	*invoker.BaseInvoker
}

func (i *UnfreezePubInvoker) Invoke() (*model.UnfreezePubResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnfreezePubResponse), nil
	}
}

type UpdatePubInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePubInfoInvoker) Invoke() (*model.UpdatePubInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePubInfoResponse), nil
	}
}

type CreatePubInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePubInfoInvoker) Invoke() (*model.CreatePubInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePubInfoResponse), nil
	}
}

type PushMenuInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *PushMenuInfoInvoker) Invoke() (*model.PushMenuInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PushMenuInfoResponse), nil
	}
}

type PushPortalInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *PushPortalInfoInvoker) Invoke() (*model.PushPortalInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PushPortalInfoResponse), nil
	}
}

type UploadMediaInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadMediaInvoker) Invoke() (*model.UploadMediaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadMediaResponse), nil
	}
}

type AddVmsCallBackInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddVmsCallBackInvoker) Invoke() (*model.AddVmsCallBackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddVmsCallBackResponse), nil
	}
}

type CreateVmsSendTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVmsSendTaskInvoker) Invoke() (*model.CreateVmsSendTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVmsSendTaskResponse), nil
	}
}

type ListVmsCallbacksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVmsCallbacksInvoker) Invoke() (*model.ListVmsCallbacksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVmsCallbacksResponse), nil
	}
}

type ListVmsSendTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVmsSendTasksInvoker) Invoke() (*model.ListVmsSendTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVmsSendTasksResponse), nil
	}
}

type CreateVmsTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVmsTemplateInvoker) Invoke() (*model.CreateVmsTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVmsTemplateResponse), nil
	}
}

type ListVmsTemplateStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVmsTemplateStatusInvoker) Invoke() (*model.ListVmsTemplateStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVmsTemplateStatusResponse), nil
	}
}
