package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eihealth/v2/model"
)

type EiHealthClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewEiHealthClient(hcClient *httpclient.HcHttpClient) *EiHealthClient {
	return &EiHealthClient{HcClient: hcClient}
}

func EiHealthClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// ShowDrugJobStatistics 统计药物作业数量和调用次数
//
// 统计药物作业数量和调用次数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowDrugJobStatistics(request *model.ShowDrugJobStatisticsRequest) (*model.ShowDrugJobStatisticsResponse, error) {
	requestDef := GenReqDefForShowDrugJobStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDrugJobStatisticsResponse), nil
	}
}

// ShowDrugJobStatisticsInvoker 统计药物作业数量和调用次数
func (c *EiHealthClient) ShowDrugJobStatisticsInvoker(request *model.ShowDrugJobStatisticsRequest) *ShowDrugJobStatisticsInvoker {
	requestDef := GenReqDefForShowDrugJobStatistics()
	return &ShowDrugJobStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMessageFeedback 创建问答内容反馈
//
// 创建问答内容反馈。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateMessageFeedback(request *model.CreateMessageFeedbackRequest) (*model.CreateMessageFeedbackResponse, error) {
	requestDef := GenReqDefForCreateMessageFeedback()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMessageFeedbackResponse), nil
	}
}

// CreateMessageFeedbackInvoker 创建问答内容反馈
func (c *EiHealthClient) CreateMessageFeedbackInvoker(request *model.CreateMessageFeedbackRequest) *CreateMessageFeedbackInvoker {
	requestDef := GenReqDefForCreateMessageFeedback()
	return &CreateMessageFeedbackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAssistantModel 创建助手模型
//
// 创建助手模型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateAssistantModel(request *model.CreateAssistantModelRequest) (*model.CreateAssistantModelResponse, error) {
	requestDef := GenReqDefForCreateAssistantModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAssistantModelResponse), nil
	}
}

// CreateAssistantModelInvoker 创建助手模型
func (c *EiHealthClient) CreateAssistantModelInvoker(request *model.CreateAssistantModelRequest) *CreateAssistantModelInvoker {
	requestDef := GenReqDefForCreateAssistantModel()
	return &CreateAssistantModelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAssistantModel 删除助手模型
//
// 删除助手模型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteAssistantModel(request *model.DeleteAssistantModelRequest) (*model.DeleteAssistantModelResponse, error) {
	requestDef := GenReqDefForDeleteAssistantModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAssistantModelResponse), nil
	}
}

// DeleteAssistantModelInvoker 删除助手模型
func (c *EiHealthClient) DeleteAssistantModelInvoker(request *model.DeleteAssistantModelRequest) *DeleteAssistantModelInvoker {
	requestDef := GenReqDefForDeleteAssistantModel()
	return &DeleteAssistantModelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllAssistantModels 获取模型列表
//
// 获取模型列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListAllAssistantModels(request *model.ListAllAssistantModelsRequest) (*model.ListAllAssistantModelsResponse, error) {
	requestDef := GenReqDefForListAllAssistantModels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllAssistantModelsResponse), nil
	}
}

// ListAllAssistantModelsInvoker 获取模型列表
func (c *EiHealthClient) ListAllAssistantModelsInvoker(request *model.ListAllAssistantModelsRequest) *ListAllAssistantModelsInvoker {
	requestDef := GenReqDefForListAllAssistantModels()
	return &ListAllAssistantModelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAssistantModels 获取供应商模型列表
//
// 获取供应商模型列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListAssistantModels(request *model.ListAssistantModelsRequest) (*model.ListAssistantModelsResponse, error) {
	requestDef := GenReqDefForListAssistantModels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAssistantModelsResponse), nil
	}
}

// ListAssistantModelsInvoker 获取供应商模型列表
func (c *EiHealthClient) ListAssistantModelsInvoker(request *model.ListAssistantModelsRequest) *ListAssistantModelsInvoker {
	requestDef := GenReqDefForListAssistantModels()
	return &ListAssistantModelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAssistantModel 更新助手模型
//
// 更新助手模型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateAssistantModel(request *model.UpdateAssistantModelRequest) (*model.UpdateAssistantModelResponse, error) {
	requestDef := GenReqDefForUpdateAssistantModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAssistantModelResponse), nil
	}
}

// UpdateAssistantModelInvoker 更新助手模型
func (c *EiHealthClient) UpdateAssistantModelInvoker(request *model.UpdateAssistantModelRequest) *UpdateAssistantModelInvoker {
	requestDef := GenReqDefForUpdateAssistantModel()
	return &UpdateAssistantModelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateModelVendor 创建模型供应商
//
// 创建模型供应商。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateModelVendor(request *model.CreateModelVendorRequest) (*model.CreateModelVendorResponse, error) {
	requestDef := GenReqDefForCreateModelVendor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateModelVendorResponse), nil
	}
}

// CreateModelVendorInvoker 创建模型供应商
func (c *EiHealthClient) CreateModelVendorInvoker(request *model.CreateModelVendorRequest) *CreateModelVendorInvoker {
	requestDef := GenReqDefForCreateModelVendor()
	return &CreateModelVendorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteModelVendor 删除模型供应商
//
// 删除模型供应商。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteModelVendor(request *model.DeleteModelVendorRequest) (*model.DeleteModelVendorResponse, error) {
	requestDef := GenReqDefForDeleteModelVendor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteModelVendorResponse), nil
	}
}

// DeleteModelVendorInvoker 删除模型供应商
func (c *EiHealthClient) DeleteModelVendorInvoker(request *model.DeleteModelVendorRequest) *DeleteModelVendorInvoker {
	requestDef := GenReqDefForDeleteModelVendor()
	return &DeleteModelVendorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListModelVendors 获取模型供应商列表
//
// 获取模型供应商列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListModelVendors(request *model.ListModelVendorsRequest) (*model.ListModelVendorsResponse, error) {
	requestDef := GenReqDefForListModelVendors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListModelVendorsResponse), nil
	}
}

// ListModelVendorsInvoker 获取模型供应商列表
func (c *EiHealthClient) ListModelVendorsInvoker(request *model.ListModelVendorsRequest) *ListModelVendorsInvoker {
	requestDef := GenReqDefForListModelVendors()
	return &ListModelVendorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateModelVendor 更新模型供应商
//
// 更新模型供应商。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateModelVendor(request *model.UpdateModelVendorRequest) (*model.UpdateModelVendorResponse, error) {
	requestDef := GenReqDefForUpdateModelVendor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateModelVendorResponse), nil
	}
}

// UpdateModelVendorInvoker 更新模型供应商
func (c *EiHealthClient) UpdateModelVendorInvoker(request *model.UpdateModelVendorRequest) *UpdateModelVendorInvoker {
	requestDef := GenReqDefForUpdateModelVendor()
	return &UpdateModelVendorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
