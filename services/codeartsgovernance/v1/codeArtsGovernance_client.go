package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codeartsgovernance/v1/model"
)

type CodeArtsGovernanceClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCodeArtsGovernanceClient(hcClient *httpclient.HcHttpClient) *CodeArtsGovernanceClient {
	return &CodeArtsGovernanceClient{HcClient: hcClient}
}

func CodeArtsGovernanceClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateSecAppTask 创建移动应用安全任务并启动
//
// 创建移动应用安全任务并启动
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsGovernanceClient) CreateSecAppTask(request *model.CreateSecAppTaskRequest) (*model.CreateSecAppTaskResponse, error) {
	requestDef := GenReqDefForCreateSecAppTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecAppTaskResponse), nil
	}
}

// CreateSecAppTaskInvoker 创建移动应用安全任务并启动
func (c *CodeArtsGovernanceClient) CreateSecAppTaskInvoker(request *model.CreateSecAppTaskRequest) *CreateSecAppTaskInvoker {
	requestDef := GenReqDefForCreateSecAppTask()
	return &CreateSecAppTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSecAppTask 删除移动应用安全任务
//
// 删除移动应用安全任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsGovernanceClient) DeleteSecAppTask(request *model.DeleteSecAppTaskRequest) (*model.DeleteSecAppTaskResponse, error) {
	requestDef := GenReqDefForDeleteSecAppTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecAppTaskResponse), nil
	}
}

// DeleteSecAppTaskInvoker 删除移动应用安全任务
func (c *CodeArtsGovernanceClient) DeleteSecAppTaskInvoker(request *model.DeleteSecAppTaskRequest) *DeleteSecAppTaskInvoker {
	requestDef := GenReqDefForDeleteSecAppTask()
	return &DeleteSecAppTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHmReport 获取鸿蒙生态应用检查结果
//
// 获取鸿蒙生态应用检查结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsGovernanceClient) ShowHmReport(request *model.ShowHmReportRequest) (*model.ShowHmReportResponse, error) {
	requestDef := GenReqDefForShowHmReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHmReportResponse), nil
	}
}

// ShowHmReportInvoker 获取鸿蒙生态应用检查结果
func (c *CodeArtsGovernanceClient) ShowHmReportInvoker(request *model.ShowHmReportRequest) *ShowHmReportInvoker {
	requestDef := GenReqDefForShowHmReport()
	return &ShowHmReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSecAppTaskResult 获取移动应用安全任务结果
//
// 获取移动应用安全任务结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsGovernanceClient) ShowSecAppTaskResult(request *model.ShowSecAppTaskResultRequest) (*model.ShowSecAppTaskResultResponse, error) {
	requestDef := GenReqDefForShowSecAppTaskResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecAppTaskResultResponse), nil
	}
}

// ShowSecAppTaskResultInvoker 获取移动应用安全任务结果
func (c *CodeArtsGovernanceClient) ShowSecAppTaskResultInvoker(request *model.ShowSecAppTaskResultRequest) *ShowSecAppTaskResultInvoker {
	requestDef := GenReqDefForShowSecAppTaskResult()
	return &ShowSecAppTaskResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSecAppTaskStatus 查询移动应用安全任务状态
//
// 查询移动应用安全任务状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsGovernanceClient) ShowSecAppTaskStatus(request *model.ShowSecAppTaskStatusRequest) (*model.ShowSecAppTaskStatusResponse, error) {
	requestDef := GenReqDefForShowSecAppTaskStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecAppTaskStatusResponse), nil
	}
}

// ShowSecAppTaskStatusInvoker 查询移动应用安全任务状态
func (c *CodeArtsGovernanceClient) ShowSecAppTaskStatusInvoker(request *model.ShowSecAppTaskStatusRequest) *ShowSecAppTaskStatusInvoker {
	requestDef := GenReqDefForShowSecAppTaskStatus()
	return &ShowSecAppTaskStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
