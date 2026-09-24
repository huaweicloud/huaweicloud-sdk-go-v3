package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/optverse/v1/model"
)

type OptVerseClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewOptVerseClient(hcClient *httpclient.HcHttpClient) *OptVerseClient {
	return &OptVerseClient{HcClient: hcClient}
}

func OptVerseClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// BatchDeleteEvolveTask 删除算法演化任务
//
// 删除算法演化任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) BatchDeleteEvolveTask(request *model.BatchDeleteEvolveTaskRequest) (*model.BatchDeleteEvolveTaskResponse, error) {
	requestDef := GenReqDefForBatchDeleteEvolveTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteEvolveTaskResponse), nil
	}
}

// BatchDeleteEvolveTaskInvoker 删除算法演化任务
func (c *OptVerseClient) BatchDeleteEvolveTaskInvoker(request *model.BatchDeleteEvolveTaskRequest) *BatchDeleteEvolveTaskInvoker {
	requestDef := GenReqDefForBatchDeleteEvolveTask()
	return &BatchDeleteEvolveTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAlgorithm 创建设计项目
//
// 创建设计项目
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) CreateAlgorithm(request *model.CreateAlgorithmRequest) (*model.CreateAlgorithmResponse, error) {
	requestDef := GenReqDefForCreateAlgorithm()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAlgorithmResponse), nil
	}
}

// CreateAlgorithmInvoker 创建设计项目
func (c *OptVerseClient) CreateAlgorithmInvoker(request *model.CreateAlgorithmRequest) *CreateAlgorithmInvoker {
	requestDef := GenReqDefForCreateAlgorithm()
	return &CreateAlgorithmInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEvolveTask 创建演化任务
//
// 创建演化任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) CreateEvolveTask(request *model.CreateEvolveTaskRequest) (*model.CreateEvolveTaskResponse, error) {
	requestDef := GenReqDefForCreateEvolveTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEvolveTaskResponse), nil
	}
}

// CreateEvolveTaskInvoker 创建演化任务
func (c *OptVerseClient) CreateEvolveTaskInvoker(request *model.CreateEvolveTaskRequest) *CreateEvolveTaskInvoker {
	requestDef := GenReqDefForCreateEvolveTask()
	return &CreateEvolveTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAlgorithm 删除算法设计项目
//
// 删除算法设计项目
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) DeleteAlgorithm(request *model.DeleteAlgorithmRequest) (*model.DeleteAlgorithmResponse, error) {
	requestDef := GenReqDefForDeleteAlgorithm()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAlgorithmResponse), nil
	}
}

// DeleteAlgorithmInvoker 删除算法设计项目
func (c *OptVerseClient) DeleteAlgorithmInvoker(request *model.DeleteAlgorithmRequest) *DeleteAlgorithmInvoker {
	requestDef := GenReqDefForDeleteAlgorithm()
	return &DeleteAlgorithmInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAlgorithmFile 删除算法设计项目中的文件
//
// 删除算法设计项目中的文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) DeleteAlgorithmFile(request *model.DeleteAlgorithmFileRequest) (*model.DeleteAlgorithmFileResponse, error) {
	requestDef := GenReqDefForDeleteAlgorithmFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAlgorithmFileResponse), nil
	}
}

// DeleteAlgorithmFileInvoker 删除算法设计项目中的文件
func (c *OptVerseClient) DeleteAlgorithmFileInvoker(request *model.DeleteAlgorithmFileRequest) *DeleteAlgorithmFileInvoker {
	requestDef := GenReqDefForDeleteAlgorithmFile()
	return &DeleteAlgorithmFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEvolveTask 删除算法演化任务
//
// 删除算法演化任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) DeleteEvolveTask(request *model.DeleteEvolveTaskRequest) (*model.DeleteEvolveTaskResponse, error) {
	requestDef := GenReqDefForDeleteEvolveTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEvolveTaskResponse), nil
	}
}

// DeleteEvolveTaskInvoker 删除算法演化任务
func (c *OptVerseClient) DeleteEvolveTaskInvoker(request *model.DeleteEvolveTaskRequest) *DeleteEvolveTaskInvoker {
	requestDef := GenReqDefForDeleteEvolveTask()
	return &DeleteEvolveTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportAlgorithmFile 上传算法文件（ZIP包）
//
// 上传算法文件（ZIP包）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) ImportAlgorithmFile(request *model.ImportAlgorithmFileRequest) (*model.ImportAlgorithmFileResponse, error) {
	requestDef := GenReqDefForImportAlgorithmFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportAlgorithmFileResponse), nil
	}
}

// ImportAlgorithmFileInvoker 上传算法文件（ZIP包）
func (c *OptVerseClient) ImportAlgorithmFileInvoker(request *model.ImportAlgorithmFileRequest) *ImportAlgorithmFileInvoker {
	requestDef := GenReqDefForImportAlgorithmFile()
	return &ImportAlgorithmFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlgorithms 批量查询设计项目列表
//
// 批量查询设计项目列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) ListAlgorithms(request *model.ListAlgorithmsRequest) (*model.ListAlgorithmsResponse, error) {
	requestDef := GenReqDefForListAlgorithms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlgorithmsResponse), nil
	}
}

// ListAlgorithmsInvoker 批量查询设计项目列表
func (c *OptVerseClient) ListAlgorithmsInvoker(request *model.ListAlgorithmsRequest) *ListAlgorithmsInvoker {
	requestDef := GenReqDefForListAlgorithms()
	return &ListAlgorithmsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDirectoryByAlgorithmId 获取某一算法设计项目文件目录
//
// 获取某一算法设计项目文件目录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) ListDirectoryByAlgorithmId(request *model.ListDirectoryByAlgorithmIdRequest) (*model.ListDirectoryByAlgorithmIdResponse, error) {
	requestDef := GenReqDefForListDirectoryByAlgorithmId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDirectoryByAlgorithmIdResponse), nil
	}
}

// ListDirectoryByAlgorithmIdInvoker 获取某一算法设计项目文件目录
func (c *OptVerseClient) ListDirectoryByAlgorithmIdInvoker(request *model.ListDirectoryByAlgorithmIdRequest) *ListDirectoryByAlgorithmIdInvoker {
	requestDef := GenReqDefForListDirectoryByAlgorithmId()
	return &ListDirectoryByAlgorithmIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDirectoryByResultCommitId 获取某一演化任务某次结果的commit的目录
//
// 获取某一演化任务某次结果的commit的目录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) ListDirectoryByResultCommitId(request *model.ListDirectoryByResultCommitIdRequest) (*model.ListDirectoryByResultCommitIdResponse, error) {
	requestDef := GenReqDefForListDirectoryByResultCommitId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDirectoryByResultCommitIdResponse), nil
	}
}

// ListDirectoryByResultCommitIdInvoker 获取某一演化任务某次结果的commit的目录
func (c *OptVerseClient) ListDirectoryByResultCommitIdInvoker(request *model.ListDirectoryByResultCommitIdRequest) *ListDirectoryByResultCommitIdInvoker {
	requestDef := GenReqDefForListDirectoryByResultCommitId()
	return &ListDirectoryByResultCommitIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEvolveTaskMetas 批量查询演化项目列表
//
// 批量查询演化项目列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) ListEvolveTaskMetas(request *model.ListEvolveTaskMetasRequest) (*model.ListEvolveTaskMetasResponse, error) {
	requestDef := GenReqDefForListEvolveTaskMetas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEvolveTaskMetasResponse), nil
	}
}

// ListEvolveTaskMetasInvoker 批量查询演化项目列表
func (c *OptVerseClient) ListEvolveTaskMetasInvoker(request *model.ListEvolveTaskMetasRequest) *ListEvolveTaskMetasInvoker {
	requestDef := GenReqDefForListEvolveTaskMetas()
	return &ListEvolveTaskMetasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEvolveTaskStats 查询演化任务状态统计
//
// 查询演化任务状态统计
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) ListEvolveTaskStats(request *model.ListEvolveTaskStatsRequest) (*model.ListEvolveTaskStatsResponse, error) {
	requestDef := GenReqDefForListEvolveTaskStats()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEvolveTaskStatsResponse), nil
	}
}

// ListEvolveTaskStatsInvoker 查询演化任务状态统计
func (c *OptVerseClient) ListEvolveTaskStatsInvoker(request *model.ListEvolveTaskStatsRequest) *ListEvolveTaskStatsInvoker {
	requestDef := GenReqDefForListEvolveTaskStats()
	return &ListEvolveTaskStatsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SaveAlgorithmFile 上传算法文件
//
// 上传算法文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) SaveAlgorithmFile(request *model.SaveAlgorithmFileRequest) (*model.SaveAlgorithmFileResponse, error) {
	requestDef := GenReqDefForSaveAlgorithmFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SaveAlgorithmFileResponse), nil
	}
}

// SaveAlgorithmFileInvoker 上传算法文件
func (c *OptVerseClient) SaveAlgorithmFileInvoker(request *model.SaveAlgorithmFileRequest) *SaveAlgorithmFileInvoker {
	requestDef := GenReqDefForSaveAlgorithmFile()
	return &SaveAlgorithmFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAlgorithm 获取某一算法信息详情
//
// 获取某一算法信息详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) ShowAlgorithm(request *model.ShowAlgorithmRequest) (*model.ShowAlgorithmResponse, error) {
	requestDef := GenReqDefForShowAlgorithm()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlgorithmResponse), nil
	}
}

// ShowAlgorithmInvoker 获取某一算法信息详情
func (c *OptVerseClient) ShowAlgorithmInvoker(request *model.ShowAlgorithmRequest) *ShowAlgorithmInvoker {
	requestDef := GenReqDefForShowAlgorithm()
	return &ShowAlgorithmInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAlgorithmFile 获取某一算法设计项目某一文件中的内容
//
// 获取某一算法设计项目某一文件中的内容
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) ShowAlgorithmFile(request *model.ShowAlgorithmFileRequest) (*model.ShowAlgorithmFileResponse, error) {
	requestDef := GenReqDefForShowAlgorithmFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlgorithmFileResponse), nil
	}
}

// ShowAlgorithmFileInvoker 获取某一算法设计项目某一文件中的内容
func (c *OptVerseClient) ShowAlgorithmFileInvoker(request *model.ShowAlgorithmFileRequest) *ShowAlgorithmFileInvoker {
	requestDef := GenReqDefForShowAlgorithmFile()
	return &ShowAlgorithmFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTaskDetails 获取某一演化任务详情
//
// 获取某一演化任务详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) ShowTaskDetails(request *model.ShowTaskDetailsRequest) (*model.ShowTaskDetailsResponse, error) {
	requestDef := GenReqDefForShowTaskDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskDetailsResponse), nil
	}
}

// ShowTaskDetailsInvoker 获取某一演化任务详情
func (c *OptVerseClient) ShowTaskDetailsInvoker(request *model.ShowTaskDetailsRequest) *ShowTaskDetailsInvoker {
	requestDef := GenReqDefForShowTaskDetails()
	return &ShowTaskDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTaskResultCommit 获取某一演化任务某次结果的commit文件
//
// 获取某一演化任务某次结果的commit文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) ShowTaskResultCommit(request *model.ShowTaskResultCommitRequest) (*model.ShowTaskResultCommitResponse, error) {
	requestDef := GenReqDefForShowTaskResultCommit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskResultCommitResponse), nil
	}
}

// ShowTaskResultCommitInvoker 获取某一演化任务某次结果的commit文件
func (c *OptVerseClient) ShowTaskResultCommitInvoker(request *model.ShowTaskResultCommitRequest) *ShowTaskResultCommitInvoker {
	requestDef := GenReqDefForShowTaskResultCommit()
	return &ShowTaskResultCommitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTaskResultList 获取某一演化任务运行详结果列表
//
// 获取某一演化任务运行详结果列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) ShowTaskResultList(request *model.ShowTaskResultListRequest) (*model.ShowTaskResultListResponse, error) {
	requestDef := GenReqDefForShowTaskResultList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskResultListResponse), nil
	}
}

// ShowTaskResultListInvoker 获取某一演化任务运行详结果列表
func (c *OptVerseClient) ShowTaskResultListInvoker(request *model.ShowTaskResultListRequest) *ShowTaskResultListInvoker {
	requestDef := GenReqDefForShowTaskResultList()
	return &ShowTaskResultListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTaskRunningDetails 获取某一演化任务运行详情
//
// 获取某一演化任务运行详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) ShowTaskRunningDetails(request *model.ShowTaskRunningDetailsRequest) (*model.ShowTaskRunningDetailsResponse, error) {
	requestDef := GenReqDefForShowTaskRunningDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskRunningDetailsResponse), nil
	}
}

// ShowTaskRunningDetailsInvoker 获取某一演化任务运行详情
func (c *OptVerseClient) ShowTaskRunningDetailsInvoker(request *model.ShowTaskRunningDetailsRequest) *ShowTaskRunningDetailsInvoker {
	requestDef := GenReqDefForShowTaskRunningDetails()
	return &ShowTaskRunningDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTaskRunningLog 获取某一演化任务运行日志
//
// 获取某一演化任务运行日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) ShowTaskRunningLog(request *model.ShowTaskRunningLogRequest) (*model.ShowTaskRunningLogResponse, error) {
	requestDef := GenReqDefForShowTaskRunningLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskRunningLogResponse), nil
	}
}

// ShowTaskRunningLogInvoker 获取某一演化任务运行日志
func (c *OptVerseClient) ShowTaskRunningLogInvoker(request *model.ShowTaskRunningLogRequest) *ShowTaskRunningLogInvoker {
	requestDef := GenReqDefForShowTaskRunningLog()
	return &ShowTaskRunningLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartEvolveTask 启动演化任务
//
// 启动演化任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) StartEvolveTask(request *model.StartEvolveTaskRequest) (*model.StartEvolveTaskResponse, error) {
	requestDef := GenReqDefForStartEvolveTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartEvolveTaskResponse), nil
	}
}

// StartEvolveTaskInvoker 启动演化任务
func (c *OptVerseClient) StartEvolveTaskInvoker(request *model.StartEvolveTaskRequest) *StartEvolveTaskInvoker {
	requestDef := GenReqDefForStartEvolveTask()
	return &StartEvolveTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopEvolveTask 停止演化任务
//
// 停止演化任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) StopEvolveTask(request *model.StopEvolveTaskRequest) (*model.StopEvolveTaskResponse, error) {
	requestDef := GenReqDefForStopEvolveTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopEvolveTaskResponse), nil
	}
}

// StopEvolveTaskInvoker 停止演化任务
func (c *OptVerseClient) StopEvolveTaskInvoker(request *model.StopEvolveTaskRequest) *StopEvolveTaskInvoker {
	requestDef := GenReqDefForStopEvolveTask()
	return &StopEvolveTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAlgorithm 更新算法设计项目信息
//
// 更新算法设计项目信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) UpdateAlgorithm(request *model.UpdateAlgorithmRequest) (*model.UpdateAlgorithmResponse, error) {
	requestDef := GenReqDefForUpdateAlgorithm()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAlgorithmResponse), nil
	}
}

// UpdateAlgorithmInvoker 更新算法设计项目信息
func (c *OptVerseClient) UpdateAlgorithmInvoker(request *model.UpdateAlgorithmRequest) *UpdateAlgorithmInvoker {
	requestDef := GenReqDefForUpdateAlgorithm()
	return &UpdateAlgorithmInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEvolveTask 更新算法演化任务信息
//
// 更新算法演化任务信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) UpdateEvolveTask(request *model.UpdateEvolveTaskRequest) (*model.UpdateEvolveTaskResponse, error) {
	requestDef := GenReqDefForUpdateEvolveTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEvolveTaskResponse), nil
	}
}

// UpdateEvolveTaskInvoker 更新算法演化任务信息
func (c *OptVerseClient) UpdateEvolveTaskInvoker(request *model.UpdateEvolveTaskRequest) *UpdateEvolveTaskInvoker {
	requestDef := GenReqDefForUpdateEvolveTask()
	return &UpdateEvolveTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteModelAsset 删除模型资产
//
// 删除模型资产。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) DeleteModelAsset(request *model.DeleteModelAssetRequest) (*model.DeleteModelAssetResponse, error) {
	requestDef := GenReqDefForDeleteModelAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteModelAssetResponse), nil
	}
}

// DeleteModelAssetInvoker 删除模型资产
func (c *OptVerseClient) DeleteModelAssetInvoker(request *model.DeleteModelAssetRequest) *DeleteModelAssetInvoker {
	requestDef := GenReqDefForDeleteModelAsset()
	return &DeleteModelAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListModelAssets 获取模型资产列表
//
// 获取模型资产列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) ListModelAssets(request *model.ListModelAssetsRequest) (*model.ListModelAssetsResponse, error) {
	requestDef := GenReqDefForListModelAssets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListModelAssetsResponse), nil
	}
}

// ListModelAssetsInvoker 获取模型资产列表
func (c *OptVerseClient) ListModelAssetsInvoker(request *model.ListModelAssetsRequest) *ListModelAssetsInvoker {
	requestDef := GenReqDefForListModelAssets()
	return &ListModelAssetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowModelAssetDetail 查询资产详情
//
// 查询资产详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) ShowModelAssetDetail(request *model.ShowModelAssetDetailRequest) (*model.ShowModelAssetDetailResponse, error) {
	requestDef := GenReqDefForShowModelAssetDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowModelAssetDetailResponse), nil
	}
}

// ShowModelAssetDetailInvoker 查询资产详情
func (c *OptVerseClient) ShowModelAssetDetailInvoker(request *model.ShowModelAssetDetailRequest) *ShowModelAssetDetailInvoker {
	requestDef := GenReqDefForShowModelAssetDetail()
	return &ShowModelAssetDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateModelAsset 编辑模型资产
//
// 编辑模型资产描述。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) UpdateModelAsset(request *model.UpdateModelAssetRequest) (*model.UpdateModelAssetResponse, error) {
	requestDef := GenReqDefForUpdateModelAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateModelAssetResponse), nil
	}
}

// UpdateModelAssetInvoker 编辑模型资产
func (c *OptVerseClient) UpdateModelAssetInvoker(request *model.UpdateModelAssetRequest) *UpdateModelAssetInvoker {
	requestDef := GenReqDefForUpdateModelAsset()
	return &UpdateModelAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelChat 取消对话
//
// 取消对话。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) CancelChat(request *model.CancelChatRequest) (*model.CancelChatResponse, error) {
	requestDef := GenReqDefForCancelChat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelChatResponse), nil
	}
}

// CancelChatInvoker 取消对话
func (c *OptVerseClient) CancelChatInvoker(request *model.CancelChatRequest) *CancelChatInvoker {
	requestDef := GenReqDefForCancelChat()
	return &CancelChatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateArtifacts 产物中心
//
// 创建产物中心产物。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) CreateArtifacts(request *model.CreateArtifactsRequest) (*model.CreateArtifactsResponse, error) {
	requestDef := GenReqDefForCreateArtifacts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateArtifactsResponse), nil
	}
}

// CreateArtifactsInvoker 产物中心
func (c *OptVerseClient) CreateArtifactsInvoker(request *model.CreateArtifactsRequest) *CreateArtifactsInvoker {
	requestDef := GenReqDefForCreateArtifacts()
	return &CreateArtifactsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteChat 删除对话
//
// 删除对话。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) DeleteChat(request *model.DeleteChatRequest) (*model.DeleteChatResponse, error) {
	requestDef := GenReqDefForDeleteChat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteChatResponse), nil
	}
}

// DeleteChatInvoker 删除对话
func (c *OptVerseClient) DeleteChatInvoker(request *model.DeleteChatRequest) *DeleteChatInvoker {
	requestDef := GenReqDefForDeleteChat()
	return &DeleteChatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadFile 下载文件
//
// 下载文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) DownloadFile(request *model.DownloadFileRequest) (*model.DownloadFileResponse, error) {
	requestDef := GenReqDefForDownloadFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadFileResponse), nil
	}
}

// DownloadFileInvoker 下载文件
func (c *OptVerseClient) DownloadFileInvoker(request *model.DownloadFileRequest) *DownloadFileInvoker {
	requestDef := GenReqDefForDownloadFile()
	return &DownloadFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListArtifacts 获取产物中心列表
//
// 获取产物中心列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) ListArtifacts(request *model.ListArtifactsRequest) (*model.ListArtifactsResponse, error) {
	requestDef := GenReqDefForListArtifacts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListArtifactsResponse), nil
	}
}

// ListArtifactsInvoker 获取产物中心列表
func (c *OptVerseClient) ListArtifactsInvoker(request *model.ListArtifactsRequest) *ListArtifactsInvoker {
	requestDef := GenReqDefForListArtifacts()
	return &ListArtifactsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListChat 获取对话列表
//
// 获取对话列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) ListChat(request *model.ListChatRequest) (*model.ListChatResponse, error) {
	requestDef := GenReqDefForListChat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListChatResponse), nil
	}
}

// ListChatInvoker 获取对话列表
func (c *OptVerseClient) ListChatInvoker(request *model.ListChatRequest) *ListChatInvoker {
	requestDef := GenReqDefForListChat()
	return &ListChatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PublishChat 发布助手
//
// 发布助手。
// 发布前会校验当前助手最后一个阶段的文档是否已确认，确认后才可发布。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) PublishChat(request *model.PublishChatRequest) (*model.PublishChatResponse, error) {
	requestDef := GenReqDefForPublishChat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PublishChatResponse), nil
	}
}

// PublishChatInvoker 发布助手
func (c *OptVerseClient) PublishChatInvoker(request *model.PublishChatRequest) *PublishChatInvoker {
	requestDef := GenReqDefForPublishChat()
	return &PublishChatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowChat 获取对话详情
//
// 获取对话详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) ShowChat(request *model.ShowChatRequest) (*model.ShowChatResponse, error) {
	requestDef := GenReqDefForShowChat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowChatResponse), nil
	}
}

// ShowChatInvoker 获取对话详情
func (c *OptVerseClient) ShowChatInvoker(request *model.ShowChatRequest) *ShowChatInvoker {
	requestDef := GenReqDefForShowChat()
	return &ShowChatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateChat 更新对话
//
// 更新对话。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) UpdateChat(request *model.UpdateChatRequest) (*model.UpdateChatResponse, error) {
	requestDef := GenReqDefForUpdateChat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateChatResponse), nil
	}
}

// UpdateChatInvoker 更新对话
func (c *OptVerseClient) UpdateChatInvoker(request *model.UpdateChatRequest) *UpdateChatInvoker {
	requestDef := GenReqDefForUpdateChat()
	return &UpdateChatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadFile 上传文件
//
// 上传文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) UploadFile(request *model.UploadFileRequest) (*model.UploadFileResponse, error) {
	requestDef := GenReqDefForUploadFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadFileResponse), nil
	}
}

// UploadFileInvoker 上传文件
func (c *OptVerseClient) UploadFileInvoker(request *model.UploadFileRequest) *UploadFileInvoker {
	requestDef := GenReqDefForUploadFile()
	return &UploadFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateModelService 创建模型服务
//
// 创建模型服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) CreateModelService(request *model.CreateModelServiceRequest) (*model.CreateModelServiceResponse, error) {
	requestDef := GenReqDefForCreateModelService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateModelServiceResponse), nil
	}
}

// CreateModelServiceInvoker 创建模型服务
func (c *OptVerseClient) CreateModelServiceInvoker(request *model.CreateModelServiceRequest) *CreateModelServiceInvoker {
	requestDef := GenReqDefForCreateModelService()
	return &CreateModelServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateModelServiceTask 调用模型服务创建任务
//
// 调用模型服务创建任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) CreateModelServiceTask(request *model.CreateModelServiceTaskRequest) (*model.CreateModelServiceTaskResponse, error) {
	requestDef := GenReqDefForCreateModelServiceTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateModelServiceTaskResponse), nil
	}
}

// CreateModelServiceTaskInvoker 调用模型服务创建任务
func (c *OptVerseClient) CreateModelServiceTaskInvoker(request *model.CreateModelServiceTaskRequest) *CreateModelServiceTaskInvoker {
	requestDef := GenReqDefForCreateModelServiceTask()
	return &CreateModelServiceTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteModelService 删除模型服务
//
// 删除模型服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) DeleteModelService(request *model.DeleteModelServiceRequest) (*model.DeleteModelServiceResponse, error) {
	requestDef := GenReqDefForDeleteModelService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteModelServiceResponse), nil
	}
}

// DeleteModelServiceInvoker 删除模型服务
func (c *OptVerseClient) DeleteModelServiceInvoker(request *model.DeleteModelServiceRequest) *DeleteModelServiceInvoker {
	requestDef := GenReqDefForDeleteModelService()
	return &DeleteModelServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListModelServiceTasks 获取模型服务任务列表
//
// 获取模型服务任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) ListModelServiceTasks(request *model.ListModelServiceTasksRequest) (*model.ListModelServiceTasksResponse, error) {
	requestDef := GenReqDefForListModelServiceTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListModelServiceTasksResponse), nil
	}
}

// ListModelServiceTasksInvoker 获取模型服务任务列表
func (c *OptVerseClient) ListModelServiceTasksInvoker(request *model.ListModelServiceTasksRequest) *ListModelServiceTasksInvoker {
	requestDef := GenReqDefForListModelServiceTasks()
	return &ListModelServiceTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowModelServiceDetail 获取模型服务详情
//
// 获取模型服务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) ShowModelServiceDetail(request *model.ShowModelServiceDetailRequest) (*model.ShowModelServiceDetailResponse, error) {
	requestDef := GenReqDefForShowModelServiceDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowModelServiceDetailResponse), nil
	}
}

// ShowModelServiceDetailInvoker 获取模型服务详情
func (c *OptVerseClient) ShowModelServiceDetailInvoker(request *model.ShowModelServiceDetailRequest) *ShowModelServiceDetailInvoker {
	requestDef := GenReqDefForShowModelServiceDetail()
	return &ShowModelServiceDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowModelServiceList 获取模型服务列表
//
// 获取模型服务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) ShowModelServiceList(request *model.ShowModelServiceListRequest) (*model.ShowModelServiceListResponse, error) {
	requestDef := GenReqDefForShowModelServiceList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowModelServiceListResponse), nil
	}
}

// ShowModelServiceListInvoker 获取模型服务列表
func (c *OptVerseClient) ShowModelServiceListInvoker(request *model.ShowModelServiceListRequest) *ShowModelServiceListInvoker {
	requestDef := GenReqDefForShowModelServiceList()
	return &ShowModelServiceListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowModelServiceTask 获取模型服务任务详情
//
// 获取模型服务任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) ShowModelServiceTask(request *model.ShowModelServiceTaskRequest) (*model.ShowModelServiceTaskResponse, error) {
	requestDef := GenReqDefForShowModelServiceTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowModelServiceTaskResponse), nil
	}
}

// ShowModelServiceTaskInvoker 获取模型服务任务详情
func (c *OptVerseClient) ShowModelServiceTaskInvoker(request *model.ShowModelServiceTaskRequest) *ShowModelServiceTaskInvoker {
	requestDef := GenReqDefForShowModelServiceTask()
	return &ShowModelServiceTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartModelService 启动模型服务
//
// 启动模型服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) StartModelService(request *model.StartModelServiceRequest) (*model.StartModelServiceResponse, error) {
	requestDef := GenReqDefForStartModelService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartModelServiceResponse), nil
	}
}

// StartModelServiceInvoker 启动模型服务
func (c *OptVerseClient) StartModelServiceInvoker(request *model.StartModelServiceRequest) *StartModelServiceInvoker {
	requestDef := GenReqDefForStartModelService()
	return &StartModelServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopModelService 停止模型服务
//
// 停止模型服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) StopModelService(request *model.StopModelServiceRequest) (*model.StopModelServiceResponse, error) {
	requestDef := GenReqDefForStopModelService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopModelServiceResponse), nil
	}
}

// StopModelServiceInvoker 停止模型服务
func (c *OptVerseClient) StopModelServiceInvoker(request *model.StopModelServiceRequest) *StopModelServiceInvoker {
	requestDef := GenReqDefForStopModelService()
	return &StopModelServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateModelService 编辑推理服务
//
// 编辑推理服务，仅支持修改服务名称和服务描述。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) UpdateModelService(request *model.UpdateModelServiceRequest) (*model.UpdateModelServiceResponse, error) {
	requestDef := GenReqDefForUpdateModelService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateModelServiceResponse), nil
	}
}

// UpdateModelServiceInvoker 编辑推理服务
func (c *OptVerseClient) UpdateModelServiceInvoker(request *model.UpdateModelServiceRequest) *UpdateModelServiceInvoker {
	requestDef := GenReqDefForUpdateModelService()
	return &UpdateModelServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadModelServiceTaskFile 上传任务依赖的输入文件
//
// 上传任务依赖的输入文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) UploadModelServiceTaskFile(request *model.UploadModelServiceTaskFileRequest) (*model.UploadModelServiceTaskFileResponse, error) {
	requestDef := GenReqDefForUploadModelServiceTaskFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadModelServiceTaskFileResponse), nil
	}
}

// UploadModelServiceTaskFileInvoker 上传任务依赖的输入文件
func (c *OptVerseClient) UploadModelServiceTaskFileInvoker(request *model.UploadModelServiceTaskFileRequest) *UploadModelServiceTaskFileInvoker {
	requestDef := GenReqDefForUploadModelServiceTaskFile()
	return &UploadModelServiceTaskFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PublishModel 发布模型
//
// 发布训练任务生成的模型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) PublishModel(request *model.PublishModelRequest) (*model.PublishModelResponse, error) {
	requestDef := GenReqDefForPublishModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PublishModelResponse), nil
	}
}

// PublishModelInvoker 发布模型
func (c *OptVerseClient) PublishModelInvoker(request *model.PublishModelRequest) *PublishModelInvoker {
	requestDef := GenReqDefForPublishModel()
	return &PublishModelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AuthorizePermission 授权
//
// 授予LLM4AD操作用户桶的权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) AuthorizePermission(request *model.AuthorizePermissionRequest) (*model.AuthorizePermissionResponse, error) {
	requestDef := GenReqDefForAuthorizePermission()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AuthorizePermissionResponse), nil
	}
}

// AuthorizePermissionInvoker 授权
func (c *OptVerseClient) AuthorizePermissionInvoker(request *model.AuthorizePermissionRequest) *AuthorizePermissionInvoker {
	requestDef := GenReqDefForAuthorizePermission()
	return &AuthorizePermissionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBuckets 获取Bucket清单
//
// 获取Bucket清单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) ListBuckets(request *model.ListBucketsRequest) (*model.ListBucketsResponse, error) {
	requestDef := GenReqDefForListBuckets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBucketsResponse), nil
	}
}

// ListBucketsInvoker 获取Bucket清单
func (c *OptVerseClient) ListBucketsInvoker(request *model.ListBucketsRequest) *ListBucketsInvoker {
	requestDef := GenReqDefForListBuckets()
	return &ListBucketsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListObject 获取Object清单
//
// 获取Object清单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) ListObject(request *model.ListObjectRequest) (*model.ListObjectResponse, error) {
	requestDef := GenReqDefForListObject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListObjectResponse), nil
	}
}

// ListObjectInvoker 获取Object清单
func (c *OptVerseClient) ListObjectInvoker(request *model.ListObjectRequest) *ListObjectInvoker {
	requestDef := GenReqDefForListObject()
	return &ListObjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPermission 检查桶的权限
//
// 检查桶的权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) ListPermission(request *model.ListPermissionRequest) (*model.ListPermissionResponse, error) {
	requestDef := GenReqDefForListPermission()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPermissionResponse), nil
	}
}

// ListPermissionInvoker 检查桶的权限
func (c *OptVerseClient) ListPermissionInvoker(request *model.ListPermissionRequest) *ListPermissionInvoker {
	requestDef := GenReqDefForListPermission()
	return &ListPermissionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RevokePermission 取消授权
//
// 取消LLM4AD对用户桶的权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *OptVerseClient) RevokePermission(request *model.RevokePermissionRequest) (*model.RevokePermissionResponse, error) {
	requestDef := GenReqDefForRevokePermission()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RevokePermissionResponse), nil
	}
}

// RevokePermissionInvoker 取消授权
func (c *OptVerseClient) RevokePermissionInvoker(request *model.RevokePermissionRequest) *RevokePermissionInvoker {
	requestDef := GenReqDefForRevokePermission()
	return &RevokePermissionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
