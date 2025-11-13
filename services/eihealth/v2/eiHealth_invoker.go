package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eihealth/v2/model"
)

type ShowDrugJobStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDrugJobStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDrugJobStatisticsInvoker) Invoke() (*model.ShowDrugJobStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDrugJobStatisticsResponse), nil
	}
}

type CreateMessageFeedbackInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMessageFeedbackInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMessageFeedbackInvoker) Invoke() (*model.CreateMessageFeedbackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMessageFeedbackResponse), nil
	}
}

type CreateAssistantModelInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAssistantModelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAssistantModelInvoker) Invoke() (*model.CreateAssistantModelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAssistantModelResponse), nil
	}
}

type DeleteAssistantModelInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAssistantModelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAssistantModelInvoker) Invoke() (*model.DeleteAssistantModelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAssistantModelResponse), nil
	}
}

type ListAllAssistantModelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllAssistantModelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAllAssistantModelsInvoker) Invoke() (*model.ListAllAssistantModelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllAssistantModelsResponse), nil
	}
}

type ListAssistantModelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAssistantModelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAssistantModelsInvoker) Invoke() (*model.ListAssistantModelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAssistantModelsResponse), nil
	}
}

type UpdateAssistantModelInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAssistantModelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAssistantModelInvoker) Invoke() (*model.UpdateAssistantModelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAssistantModelResponse), nil
	}
}

type CreateModelVendorInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateModelVendorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateModelVendorInvoker) Invoke() (*model.CreateModelVendorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateModelVendorResponse), nil
	}
}

type DeleteModelVendorInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteModelVendorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteModelVendorInvoker) Invoke() (*model.DeleteModelVendorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteModelVendorResponse), nil
	}
}

type ListModelVendorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListModelVendorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListModelVendorsInvoker) Invoke() (*model.ListModelVendorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListModelVendorsResponse), nil
	}
}

type UpdateModelVendorInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateModelVendorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateModelVendorInvoker) Invoke() (*model.UpdateModelVendorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateModelVendorResponse), nil
	}
}
