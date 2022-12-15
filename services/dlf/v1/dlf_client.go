package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dlf/v1/model"
)

type DlfClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDlfClient(hcClient *http_client.HcHttpClient) *DlfClient {
	return &DlfClient{HcClient: hcClient}
}

func DlfClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CancelScript 停止脚本实例的执行
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) CancelScript(request *model.CancelScriptRequest) (*model.CancelScriptResponse, error) {
	requestDef := GenReqDefForCancelScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelScriptResponse), nil
	}
}

// CancelScriptInvoker 停止脚本实例的执行
func (c *DlfClient) CancelScriptInvoker(request *model.CancelScriptRequest) *CancelScriptInvoker {
	requestDef := GenReqDefForCancelScript()
	return &CancelScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConnection 创建连接
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) CreateConnection(request *model.CreateConnectionRequest) (*model.CreateConnectionResponse, error) {
	requestDef := GenReqDefForCreateConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConnectionResponse), nil
	}
}

// CreateConnectionInvoker 创建连接
func (c *DlfClient) CreateConnectionInvoker(request *model.CreateConnectionRequest) *CreateConnectionInvoker {
	requestDef := GenReqDefForCreateConnection()
	return &CreateConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateJob 创建作业
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) CreateJob(request *model.CreateJobRequest) (*model.CreateJobResponse, error) {
	requestDef := GenReqDefForCreateJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateJobResponse), nil
	}
}

// CreateJobInvoker 创建作业
func (c *DlfClient) CreateJobInvoker(request *model.CreateJobRequest) *CreateJobInvoker {
	requestDef := GenReqDefForCreateJob()
	return &CreateJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateResource 创建资源
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) CreateResource(request *model.CreateResourceRequest) (*model.CreateResourceResponse, error) {
	requestDef := GenReqDefForCreateResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResourceResponse), nil
	}
}

// CreateResourceInvoker 创建资源
func (c *DlfClient) CreateResourceInvoker(request *model.CreateResourceRequest) *CreateResourceInvoker {
	requestDef := GenReqDefForCreateResource()
	return &CreateResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateScript 创建脚本
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) CreateScript(request *model.CreateScriptRequest) (*model.CreateScriptResponse, error) {
	requestDef := GenReqDefForCreateScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateScriptResponse), nil
	}
}

// CreateScriptInvoker 创建脚本
func (c *DlfClient) CreateScriptInvoker(request *model.CreateScriptRequest) *CreateScriptInvoker {
	requestDef := GenReqDefForCreateScript()
	return &CreateScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteConnction 删除连接
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) DeleteConnction(request *model.DeleteConnctionRequest) (*model.DeleteConnctionResponse, error) {
	requestDef := GenReqDefForDeleteConnction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConnctionResponse), nil
	}
}

// DeleteConnctionInvoker 删除连接
func (c *DlfClient) DeleteConnctionInvoker(request *model.DeleteConnctionRequest) *DeleteConnctionInvoker {
	requestDef := GenReqDefForDeleteConnction()
	return &DeleteConnctionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteJob 删除作业
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) DeleteJob(request *model.DeleteJobRequest) (*model.DeleteJobResponse, error) {
	requestDef := GenReqDefForDeleteJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteJobResponse), nil
	}
}

// DeleteJobInvoker 删除作业
func (c *DlfClient) DeleteJobInvoker(request *model.DeleteJobRequest) *DeleteJobInvoker {
	requestDef := GenReqDefForDeleteJob()
	return &DeleteJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteResource 删除资源
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) DeleteResource(request *model.DeleteResourceRequest) (*model.DeleteResourceResponse, error) {
	requestDef := GenReqDefForDeleteResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResourceResponse), nil
	}
}

// DeleteResourceInvoker 删除资源
func (c *DlfClient) DeleteResourceInvoker(request *model.DeleteResourceRequest) *DeleteResourceInvoker {
	requestDef := GenReqDefForDeleteResource()
	return &DeleteResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteScript 删除脚本
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) DeleteScript(request *model.DeleteScriptRequest) (*model.DeleteScriptResponse, error) {
	requestDef := GenReqDefForDeleteScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteScriptResponse), nil
	}
}

// DeleteScriptInvoker 删除脚本
func (c *DlfClient) DeleteScriptInvoker(request *model.DeleteScriptRequest) *DeleteScriptInvoker {
	requestDef := GenReqDefForDeleteScript()
	return &DeleteScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteScript 执行脚本
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) ExecuteScript(request *model.ExecuteScriptRequest) (*model.ExecuteScriptResponse, error) {
	requestDef := GenReqDefForExecuteScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteScriptResponse), nil
	}
}

// ExecuteScriptInvoker 执行脚本
func (c *DlfClient) ExecuteScriptInvoker(request *model.ExecuteScriptRequest) *ExecuteScriptInvoker {
	requestDef := GenReqDefForExecuteScript()
	return &ExecuteScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportConnections 导出连接
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) ExportConnections(request *model.ExportConnectionsRequest) (*model.ExportConnectionsResponse, error) {
	requestDef := GenReqDefForExportConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportConnectionsResponse), nil
	}
}

// ExportConnectionsInvoker 导出连接
func (c *DlfClient) ExportConnectionsInvoker(request *model.ExportConnectionsRequest) *ExportConnectionsInvoker {
	requestDef := GenReqDefForExportConnections()
	return &ExportConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportJob 导出作业
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) ExportJob(request *model.ExportJobRequest) (*model.ExportJobResponse, error) {
	requestDef := GenReqDefForExportJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportJobResponse), nil
	}
}

// ExportJobInvoker 导出作业
func (c *DlfClient) ExportJobInvoker(request *model.ExportJobRequest) *ExportJobInvoker {
	requestDef := GenReqDefForExportJob()
	return &ExportJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportJobList 批量导出作业
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) ExportJobList(request *model.ExportJobListRequest) (*model.ExportJobListResponse, error) {
	requestDef := GenReqDefForExportJobList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportJobListResponse), nil
	}
}

// ExportJobListInvoker 批量导出作业
func (c *DlfClient) ExportJobListInvoker(request *model.ExportJobListRequest) *ExportJobListInvoker {
	requestDef := GenReqDefForExportJobList()
	return &ExportJobListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportConnections 导入连接
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) ImportConnections(request *model.ImportConnectionsRequest) (*model.ImportConnectionsResponse, error) {
	requestDef := GenReqDefForImportConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportConnectionsResponse), nil
	}
}

// ImportConnectionsInvoker 导入连接
func (c *DlfClient) ImportConnectionsInvoker(request *model.ImportConnectionsRequest) *ImportConnectionsInvoker {
	requestDef := GenReqDefForImportConnections()
	return &ImportConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportJob 导入作业
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) ImportJob(request *model.ImportJobRequest) (*model.ImportJobResponse, error) {
	requestDef := GenReqDefForImportJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportJobResponse), nil
	}
}

// ImportJobInvoker 导入作业
func (c *DlfClient) ImportJobInvoker(request *model.ImportJobRequest) *ImportJobInvoker {
	requestDef := GenReqDefForImportJob()
	return &ImportJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConnections 查询连接列表
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) ListConnections(request *model.ListConnectionsRequest) (*model.ListConnectionsResponse, error) {
	requestDef := GenReqDefForListConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConnectionsResponse), nil
	}
}

// ListConnectionsInvoker 查询连接列表
func (c *DlfClient) ListConnectionsInvoker(request *model.ListConnectionsRequest) *ListConnectionsInvoker {
	requestDef := GenReqDefForListConnections()
	return &ListConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJobInstances 查询作业实例列表
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) ListJobInstances(request *model.ListJobInstancesRequest) (*model.ListJobInstancesResponse, error) {
	requestDef := GenReqDefForListJobInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobInstancesResponse), nil
	}
}

// ListJobInstancesInvoker 查询作业实例列表
func (c *DlfClient) ListJobInstancesInvoker(request *model.ListJobInstancesRequest) *ListJobInstancesInvoker {
	requestDef := GenReqDefForListJobInstances()
	return &ListJobInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJobs 查询作业列表
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) ListJobs(request *model.ListJobsRequest) (*model.ListJobsResponse, error) {
	requestDef := GenReqDefForListJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobsResponse), nil
	}
}

// ListJobsInvoker 查询作业列表
func (c *DlfClient) ListJobsInvoker(request *model.ListJobsRequest) *ListJobsInvoker {
	requestDef := GenReqDefForListJobs()
	return &ListJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResources 查询资源列表
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) ListResources(request *model.ListResourcesRequest) (*model.ListResourcesResponse, error) {
	requestDef := GenReqDefForListResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourcesResponse), nil
	}
}

// ListResourcesInvoker 查询资源列表
func (c *DlfClient) ListResourcesInvoker(request *model.ListResourcesRequest) *ListResourcesInvoker {
	requestDef := GenReqDefForListResources()
	return &ListResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScriptResults 查询脚本实例执行结果
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) ListScriptResults(request *model.ListScriptResultsRequest) (*model.ListScriptResultsResponse, error) {
	requestDef := GenReqDefForListScriptResults()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScriptResultsResponse), nil
	}
}

// ListScriptResultsInvoker 查询脚本实例执行结果
func (c *DlfClient) ListScriptResultsInvoker(request *model.ListScriptResultsRequest) *ListScriptResultsInvoker {
	requestDef := GenReqDefForListScriptResults()
	return &ListScriptResultsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScripts 查询脚本列表
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) ListScripts(request *model.ListScriptsRequest) (*model.ListScriptsResponse, error) {
	requestDef := GenReqDefForListScripts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScriptsResponse), nil
	}
}

// ListScriptsInvoker 查询脚本列表
func (c *DlfClient) ListScriptsInvoker(request *model.ListScriptsRequest) *ListScriptsInvoker {
	requestDef := GenReqDefForListScripts()
	return &ListScriptsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSystemTasks 查询系统任务详情
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) ListSystemTasks(request *model.ListSystemTasksRequest) (*model.ListSystemTasksResponse, error) {
	requestDef := GenReqDefForListSystemTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSystemTasksResponse), nil
	}
}

// ListSystemTasksInvoker 查询系统任务详情
func (c *DlfClient) ListSystemTasksInvoker(request *model.ListSystemTasksRequest) *ListSystemTasksInvoker {
	requestDef := GenReqDefForListSystemTasks()
	return &ListSystemTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestoreJobInstance 重跑作业实例
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) RestoreJobInstance(request *model.RestoreJobInstanceRequest) (*model.RestoreJobInstanceResponse, error) {
	requestDef := GenReqDefForRestoreJobInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreJobInstanceResponse), nil
	}
}

// RestoreJobInstanceInvoker 重跑作业实例
func (c *DlfClient) RestoreJobInstanceInvoker(request *model.RestoreJobInstanceRequest) *RestoreJobInstanceInvoker {
	requestDef := GenReqDefForRestoreJobInstance()
	return &RestoreJobInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunOnce 单次执行作业
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) RunOnce(request *model.RunOnceRequest) (*model.RunOnceResponse, error) {
	requestDef := GenReqDefForRunOnce()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunOnceResponse), nil
	}
}

// RunOnceInvoker 单次执行作业
func (c *DlfClient) RunOnceInvoker(request *model.RunOnceRequest) *RunOnceInvoker {
	requestDef := GenReqDefForRunOnce()
	return &RunOnceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConnection 查询连接详情
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) ShowConnection(request *model.ShowConnectionRequest) (*model.ShowConnectionResponse, error) {
	requestDef := GenReqDefForShowConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConnectionResponse), nil
	}
}

// ShowConnectionInvoker 查询连接详情
func (c *DlfClient) ShowConnectionInvoker(request *model.ShowConnectionRequest) *ShowConnectionInvoker {
	requestDef := GenReqDefForShowConnection()
	return &ShowConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDirectoryTree 查询目录树
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) ShowDirectoryTree(request *model.ShowDirectoryTreeRequest) (*model.ShowDirectoryTreeResponse, error) {
	requestDef := GenReqDefForShowDirectoryTree()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDirectoryTreeResponse), nil
	}
}

// ShowDirectoryTreeInvoker 查询目录树
func (c *DlfClient) ShowDirectoryTreeInvoker(request *model.ShowDirectoryTreeRequest) *ShowDirectoryTreeInvoker {
	requestDef := GenReqDefForShowDirectoryTree()
	return &ShowDirectoryTreeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFileInfo 检查导入作业文件中的作业和脚本
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) ShowFileInfo(request *model.ShowFileInfoRequest) (*model.ShowFileInfoResponse, error) {
	requestDef := GenReqDefForShowFileInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFileInfoResponse), nil
	}
}

// ShowFileInfoInvoker 检查导入作业文件中的作业和脚本
func (c *DlfClient) ShowFileInfoInvoker(request *model.ShowFileInfoRequest) *ShowFileInfoInvoker {
	requestDef := GenReqDefForShowFileInfo()
	return &ShowFileInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJob 查询作业详情
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) ShowJob(request *model.ShowJobRequest) (*model.ShowJobResponse, error) {
	requestDef := GenReqDefForShowJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobResponse), nil
	}
}

// ShowJobInvoker 查询作业详情
func (c *DlfClient) ShowJobInvoker(request *model.ShowJobRequest) *ShowJobInvoker {
	requestDef := GenReqDefForShowJob()
	return &ShowJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobInstance 查询作业实例详情
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) ShowJobInstance(request *model.ShowJobInstanceRequest) (*model.ShowJobInstanceResponse, error) {
	requestDef := GenReqDefForShowJobInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobInstanceResponse), nil
	}
}

// ShowJobInstanceInvoker 查询作业实例详情
func (c *DlfClient) ShowJobInstanceInvoker(request *model.ShowJobInstanceRequest) *ShowJobInstanceInvoker {
	requestDef := GenReqDefForShowJobInstance()
	return &ShowJobInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobStatus 查询实时作业的运行状态
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) ShowJobStatus(request *model.ShowJobStatusRequest) (*model.ShowJobStatusResponse, error) {
	requestDef := GenReqDefForShowJobStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobStatusResponse), nil
	}
}

// ShowJobStatusInvoker 查询实时作业的运行状态
func (c *DlfClient) ShowJobStatusInvoker(request *model.ShowJobStatusRequest) *ShowJobStatusInvoker {
	requestDef := GenReqDefForShowJobStatus()
	return &ShowJobStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResource 查询资源详情
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) ShowResource(request *model.ShowResourceRequest) (*model.ShowResourceResponse, error) {
	requestDef := GenReqDefForShowResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceResponse), nil
	}
}

// ShowResourceInvoker 查询资源详情
func (c *DlfClient) ShowResourceInvoker(request *model.ShowResourceRequest) *ShowResourceInvoker {
	requestDef := GenReqDefForShowResource()
	return &ShowResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowScript 查询脚本信息
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) ShowScript(request *model.ShowScriptRequest) (*model.ShowScriptResponse, error) {
	requestDef := GenReqDefForShowScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowScriptResponse), nil
	}
}

// ShowScriptInvoker 查询脚本信息
func (c *DlfClient) ShowScriptInvoker(request *model.ShowScriptRequest) *ShowScriptInvoker {
	requestDef := GenReqDefForShowScript()
	return &ShowScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartJob 启动作业
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) StartJob(request *model.StartJobRequest) (*model.StartJobResponse, error) {
	requestDef := GenReqDefForStartJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartJobResponse), nil
	}
}

// StartJobInvoker 启动作业
func (c *DlfClient) StartJobInvoker(request *model.StartJobRequest) *StartJobInvoker {
	requestDef := GenReqDefForStartJob()
	return &StartJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopJob 停止作业
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) StopJob(request *model.StopJobRequest) (*model.StopJobResponse, error) {
	requestDef := GenReqDefForStopJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopJobResponse), nil
	}
}

// StopJobInvoker 停止作业
func (c *DlfClient) StopJobInvoker(request *model.StopJobRequest) *StopJobInvoker {
	requestDef := GenReqDefForStopJob()
	return &StopJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopJobInstance 停止作业实例
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) StopJobInstance(request *model.StopJobInstanceRequest) (*model.StopJobInstanceResponse, error) {
	requestDef := GenReqDefForStopJobInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopJobInstanceResponse), nil
	}
}

// StopJobInstanceInvoker 停止作业实例
func (c *DlfClient) StopJobInstanceInvoker(request *model.StopJobInstanceRequest) *StopJobInstanceInvoker {
	requestDef := GenReqDefForStopJobInstance()
	return &StopJobInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateConnection 修改连接
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) UpdateConnection(request *model.UpdateConnectionRequest) (*model.UpdateConnectionResponse, error) {
	requestDef := GenReqDefForUpdateConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateConnectionResponse), nil
	}
}

// UpdateConnectionInvoker 修改连接
func (c *DlfClient) UpdateConnectionInvoker(request *model.UpdateConnectionRequest) *UpdateConnectionInvoker {
	requestDef := GenReqDefForUpdateConnection()
	return &UpdateConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateJob 修改作业
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) UpdateJob(request *model.UpdateJobRequest) (*model.UpdateJobResponse, error) {
	requestDef := GenReqDefForUpdateJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateJobResponse), nil
	}
}

// UpdateJobInvoker 修改作业
func (c *DlfClient) UpdateJobInvoker(request *model.UpdateJobRequest) *UpdateJobInvoker {
	requestDef := GenReqDefForUpdateJob()
	return &UpdateJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateResource 修改资源
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) UpdateResource(request *model.UpdateResourceRequest) (*model.UpdateResourceResponse, error) {
	requestDef := GenReqDefForUpdateResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResourceResponse), nil
	}
}

// UpdateResourceInvoker 修改资源
func (c *DlfClient) UpdateResourceInvoker(request *model.UpdateResourceRequest) *UpdateResourceInvoker {
	requestDef := GenReqDefForUpdateResource()
	return &UpdateResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateScript 修改脚本内容
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DlfClient) UpdateScript(request *model.UpdateScriptRequest) (*model.UpdateScriptResponse, error) {
	requestDef := GenReqDefForUpdateScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateScriptResponse), nil
	}
}

// UpdateScriptInvoker 修改脚本内容
func (c *DlfClient) UpdateScriptInvoker(request *model.UpdateScriptRequest) *UpdateScriptInvoker {
	requestDef := GenReqDefForUpdateScript()
	return &UpdateScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
