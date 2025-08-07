package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/clouddc/v1/model"
)

type CloudDCClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCloudDCClient(hcClient *httpclient.HcHttpClient) *CloudDCClient {
	return &CloudDCClient{HcClient: hcClient}
}

func CloudDCClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("basic.Credentials")
	return builder
}

// BatchCreateIrackTags 批量创建机柜标签
//
// 批量创建机柜标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) BatchCreateIrackTags(request *model.BatchCreateIrackTagsRequest) (*model.BatchCreateIrackTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateIrackTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateIrackTagsResponse), nil
	}
}

// BatchCreateIrackTagsInvoker 批量创建机柜标签
func (c *CloudDCClient) BatchCreateIrackTagsInvoker(request *model.BatchCreateIrackTagsRequest) *BatchCreateIrackTagsInvoker {
	requestDef := GenReqDefForBatchCreateIrackTags()
	return &BatchCreateIrackTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateTags 批量创建资源标签
//
// 批量创建资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) BatchCreateTags(request *model.BatchCreateTagsRequest) (*model.BatchCreateTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateTagsResponse), nil
	}
}

// BatchCreateTagsInvoker 批量创建资源标签
func (c *CloudDCClient) BatchCreateTagsInvoker(request *model.BatchCreateTagsRequest) *BatchCreateTagsInvoker {
	requestDef := GenReqDefForBatchCreateTags()
	return &BatchCreateTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteIrackTags 批量删除机柜标签
//
// 批量删除机柜标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) BatchDeleteIrackTags(request *model.BatchDeleteIrackTagsRequest) (*model.BatchDeleteIrackTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteIrackTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteIrackTagsResponse), nil
	}
}

// BatchDeleteIrackTagsInvoker 批量删除机柜标签
func (c *CloudDCClient) BatchDeleteIrackTagsInvoker(request *model.BatchDeleteIrackTagsRequest) *BatchDeleteIrackTagsInvoker {
	requestDef := GenReqDefForBatchDeleteIrackTags()
	return &BatchDeleteIrackTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteTags 批量删除资源标签
//
// 批量删除资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) BatchDeleteTags(request *model.BatchDeleteTagsRequest) (*model.BatchDeleteTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteTagsResponse), nil
	}
}

// BatchDeleteTagsInvoker 批量删除资源标签
func (c *CloudDCClient) BatchDeleteTagsInvoker(request *model.BatchDeleteTagsRequest) *BatchDeleteTagsInvoker {
	requestDef := GenReqDefForBatchDeleteTags()
	return &BatchDeleteTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeInstancePassword 批量修改实例密码
//
// 修改服务器管理账号（root用户或Administrator用户）密码
// 前提条件：Instance state为running
// 该接口为同步接口，全部成功或者全部失败
// 约束：
// 无符合安全要求的密码复杂度检查，非安全密码输入后，无错误提示。
// 服务器开机或重启后，新密码生效（调用**ChangeServerPowerState**接口重启）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) ChangeInstancePassword(request *model.ChangeInstancePasswordRequest) (*model.ChangeInstancePasswordResponse, error) {
	requestDef := GenReqDefForChangeInstancePassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeInstancePasswordResponse), nil
	}
}

// ChangeInstancePasswordInvoker 批量修改实例密码
func (c *CloudDCClient) ChangeInstancePasswordInvoker(request *model.ChangeInstancePasswordRequest) *ChangeInstancePasswordInvoker {
	requestDef := GenReqDefForChangeInstancePassword()
	return &ChangeInstancePasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeServerPowerState 批量修改物理服务器电源状态
//
// 修改物理服务器电源状态，此接口为异步接口，电源状态修改成功与否需要根据查询物理服务器信息获得
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) ChangeServerPowerState(request *model.ChangeServerPowerStateRequest) (*model.ChangeServerPowerStateResponse, error) {
	requestDef := GenReqDefForChangeServerPowerState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeServerPowerStateResponse), nil
	}
}

// ChangeServerPowerStateInvoker 批量修改物理服务器电源状态
func (c *CloudDCClient) ChangeServerPowerStateInvoker(request *model.ChangeServerPowerStateRequest) *ChangeServerPowerStateInvoker {
	requestDef := GenReqDefForChangeServerPowerState()
	return &ChangeServerPowerStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstance 创建实例
//
// 创建实例，支持指定IP等更多个性化参数创建实例。
//
// 调度策略支持：
// 1、指定服务器
// 2、基于空闲随机调度策略
//
// 支持VPC网络及AI参数面网络配置。
//
// 此接口为异步接口，实例的创建和启动不是立即完成的，通过接口 **ShowInstanceStatus** 查询实例状态为 **running** 代表实例创建成功。
// 接口约束：服务器**manage_state**为 **ready**
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) CreateInstance(request *model.CreateInstanceRequest) (*model.CreateInstanceResponse, error) {
	requestDef := GenReqDefForCreateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceResponse), nil
	}
}

// CreateInstanceInvoker 创建实例
func (c *CloudDCClient) CreateInstanceInvoker(request *model.CreateInstanceRequest) *CreateInstanceInvoker {
	requestDef := GenReqDefForCreateInstance()
	return &CreateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstance 删除实例
//
// 指定物理服务器删除实例，此接口为异步接口，可通过 **ShowInstanceStatus** 查询实例状态，实例状态从 **shutting-down** 变为 **terminated**，表明删除实例成功。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResponse), nil
	}
}

// DeleteInstanceInvoker 删除实例
func (c *CloudDCClient) DeleteInstanceInvoker(request *model.DeleteInstanceRequest) *DeleteInstanceInvoker {
	requestDef := GenReqDefForDeleteInstance()
	return &DeleteInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstances 批量删除实例
//
// 指定物理服务器批量删除实例，此接口为异步接口，可通过 **ShowInstanceStatus** 查询实例状态，实例状态从 **shutting-down** 变为 **terminated**，表明删除实例成功。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) DeleteInstances(request *model.DeleteInstancesRequest) (*model.DeleteInstancesResponse, error) {
	requestDef := GenReqDefForDeleteInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstancesResponse), nil
	}
}

// DeleteInstancesInvoker 批量删除实例
func (c *CloudDCClient) DeleteInstancesInvoker(request *model.DeleteInstancesRequest) *DeleteInstancesInvoker {
	requestDef := GenReqDefForDeleteInstances()
	return &DeleteInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadServerLogs 下载日志文件
//
// 下载服务器日志文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) DownloadServerLogs(request *model.DownloadServerLogsRequest) (*model.DownloadServerLogsResponse, error) {
	requestDef := GenReqDefForDownloadServerLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadServerLogsResponse), nil
	}
}

// DownloadServerLogsInvoker 下载日志文件
func (c *CloudDCClient) DownloadServerLogsInvoker(request *model.DownloadServerLogsRequest) *DownloadServerLogsInvoker {
	requestDef := GenReqDefForDownloadServerLogs()
	return &DownloadServerLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportServerLogs 导出服务器日志请求
//
// 创建服务器日志导出异步任务。根据ShowLogsExportStatus查询status，当status状态为completed时，调用DownloadServerLogs下载日志文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) ExportServerLogs(request *model.ExportServerLogsRequest) (*model.ExportServerLogsResponse, error) {
	requestDef := GenReqDefForExportServerLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportServerLogsResponse), nil
	}
}

// ExportServerLogsInvoker 导出服务器日志请求
func (c *CloudDCClient) ExportServerLogsInvoker(request *model.ExportServerLogsRequest) *ExportServerLogsInvoker {
	requestDef := GenReqDefForExportServerLogs()
	return &ExportServerLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarms 服务器告警列表
//
// 该 API 用于查询服务器告警列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) ListAlarms(request *model.ListAlarmsRequest) (*model.ListAlarmsResponse, error) {
	requestDef := GenReqDefForListAlarms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmsResponse), nil
	}
}

// ListAlarmsInvoker 服务器告警列表
func (c *CloudDCClient) ListAlarmsInvoker(request *model.ListAlarmsRequest) *ListAlarmsInvoker {
	requestDef := GenReqDefForListAlarms()
	return &ListAlarmsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEvents 服务器事件列表
//
// 该 API 用于查询服务器事件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) ListEvents(request *model.ListEventsRequest) (*model.ListEventsResponse, error) {
	requestDef := GenReqDefForListEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEventsResponse), nil
	}
}

// ListEventsInvoker 服务器事件列表
func (c *CloudDCClient) ListEventsInvoker(request *model.ListEventsRequest) *ListEventsInvoker {
	requestDef := GenReqDefForListEvents()
	return &ListEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIDcs 查询 IDC 列表
//
// 查询 IDC 列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) ListIDcs(request *model.ListIDcsRequest) (*model.ListIDcsResponse, error) {
	requestDef := GenReqDefForListIDcs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIDcsResponse), nil
	}
}

// ListIDcsInvoker 查询 IDC 列表
func (c *CloudDCClient) ListIDcsInvoker(request *model.ListIDcsRequest) *ListIDcsInvoker {
	requestDef := GenReqDefForListIDcs()
	return &ListIDcsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIRacks 查询 iRack 实例列表
//
// 用户下单后，用户上报iRack设备列表。该 API 可以查看 iRack 实例与关联imetal数量列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) ListIRacks(request *model.ListIRacksRequest) (*model.ListIRacksResponse, error) {
	requestDef := GenReqDefForListIRacks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIRacksResponse), nil
	}
}

// ListIRacksInvoker 查询 iRack 实例列表
func (c *CloudDCClient) ListIRacksInvoker(request *model.ListIRacksRequest) *ListIRacksInvoker {
	requestDef := GenReqDefForListIRacks()
	return &ListIRacksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstances 批量查询实例
//
// 批量查询实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

// ListInstancesInvoker 批量查询实例
func (c *CloudDCClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServers 批量查询物理服务器
//
// List imetal servers
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) ListServers(request *model.ListServersRequest) (*model.ListServersResponse, error) {
	requestDef := GenReqDefForListServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServersResponse), nil
	}
}

// ListServersInvoker 批量查询物理服务器
func (c *CloudDCClient) ListServersInvoker(request *model.ListServersRequest) *ListServersInvoker {
	requestDef := GenReqDefForListServers()
	return &ListServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyInstanceIp 修改实例ip
//
// 用户可选择指定的 iMetal 实例，修改ip。
//
// 接口约束：iMetal 实例处于就绪状态（已调测成功）或修改ip失败同时为下电状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) ModifyInstanceIp(request *model.ModifyInstanceIpRequest) (*model.ModifyInstanceIpResponse, error) {
	requestDef := GenReqDefForModifyInstanceIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyInstanceIpResponse), nil
	}
}

// ModifyInstanceIpInvoker 修改实例ip
func (c *CloudDCClient) ModifyInstanceIpInvoker(request *model.ModifyInstanceIpRequest) *ModifyInstanceIpInvoker {
	requestDef := GenReqDefForModifyInstanceIp()
	return &ModifyInstanceIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ReinstallOs 批量重新安装OS
//
// 指定新OS镜像重新安装OS，此接口为异步接口，通过 **ShowInstanceStatus** 查询实例状态，当状态变为 **pending** 表明正在重装中，状态为 **running** 表明完成重装。
// 前提条件：Instance state为running
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) ReinstallOs(request *model.ReinstallOsRequest) (*model.ReinstallOsResponse, error) {
	requestDef := GenReqDefForReinstallOs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReinstallOsResponse), nil
	}
}

// ReinstallOsInvoker 批量重新安装OS
func (c *CloudDCClient) ReinstallOsInvoker(request *model.ReinstallOsRequest) *ReinstallOsInvoker {
	requestDef := GenReqDefForReinstallOs()
	return &ReinstallOsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunInstances 批量创建实例
//
// 批量创建实例，其中调度策略支持：
// 1、指定服务器
// 2、基于空闲随机调度策略
//
// 支持VPC网络及AI参数面网络配置。
//
// 此接口为异步接口，实例的创建和启动不是立即完成的，通过接口 **ShowInstanceStatus** 查询实例状态为 **running** 代表实例创建成功。
// 接口约束：服务器**manage_state**为 **ready**
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) RunInstances(request *model.RunInstancesRequest) (*model.RunInstancesResponse, error) {
	requestDef := GenReqDefForRunInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunInstancesResponse), nil
	}
}

// RunInstancesInvoker 批量创建实例
func (c *CloudDCClient) RunInstancesInvoker(request *model.RunInstancesRequest) *RunInstancesInvoker {
	requestDef := GenReqDefForRunInstances()
	return &RunInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAlarmSummary 服务器告警概览
//
// 该 API 用于查询服务器告警概览
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) ShowAlarmSummary(request *model.ShowAlarmSummaryRequest) (*model.ShowAlarmSummaryResponse, error) {
	requestDef := GenReqDefForShowAlarmSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlarmSummaryResponse), nil
	}
}

// ShowAlarmSummaryInvoker 服务器告警概览
func (c *CloudDCClient) ShowAlarmSummaryInvoker(request *model.ShowAlarmSummaryRequest) *ShowAlarmSummaryInvoker {
	requestDef := GenReqDefForShowAlarmSummary()
	return &ShowAlarmSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAlarmTrend 服务器告警趋势
//
// 该 API 用于查询服务器概览、服务器开机状态和服务器运行状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) ShowAlarmTrend(request *model.ShowAlarmTrendRequest) (*model.ShowAlarmTrendResponse, error) {
	requestDef := GenReqDefForShowAlarmTrend()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlarmTrendResponse), nil
	}
}

// ShowAlarmTrendInvoker 服务器告警趋势
func (c *CloudDCClient) ShowAlarmTrendInvoker(request *model.ShowAlarmTrendRequest) *ShowAlarmTrendInvoker {
	requestDef := GenReqDefForShowAlarmTrend()
	return &ShowAlarmTrendInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEvent 查询事件定义
//
// 查询事件定义
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) ShowEvent(request *model.ShowEventRequest) (*model.ShowEventResponse, error) {
	requestDef := GenReqDefForShowEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEventResponse), nil
	}
}

// ShowEventInvoker 查询事件定义
func (c *CloudDCClient) ShowEventInvoker(request *model.ShowEventRequest) *ShowEventInvoker {
	requestDef := GenReqDefForShowEvent()
	return &ShowEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceStatus 查询实例状态
//
// 查询实例状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) ShowInstanceStatus(request *model.ShowInstanceStatusRequest) (*model.ShowInstanceStatusResponse, error) {
	requestDef := GenReqDefForShowInstanceStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceStatusResponse), nil
	}
}

// ShowInstanceStatusInvoker 查询实例状态
func (c *CloudDCClient) ShowInstanceStatusInvoker(request *model.ShowInstanceStatusRequest) *ShowInstanceStatusInvoker {
	requestDef := GenReqDefForShowInstanceStatus()
	return &ShowInstanceStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLogsExportStatus 查询日志导出状态
//
// 查询日志采集状态及进度。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) ShowLogsExportStatus(request *model.ShowLogsExportStatusRequest) (*model.ShowLogsExportStatusResponse, error) {
	requestDef := GenReqDefForShowLogsExportStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLogsExportStatusResponse), nil
	}
}

// ShowLogsExportStatusInvoker 查询日志导出状态
func (c *CloudDCClient) ShowLogsExportStatusInvoker(request *model.ShowLogsExportStatusRequest) *ShowLogsExportStatusInvoker {
	requestDef := GenReqDefForShowLogsExportStatus()
	return &ShowLogsExportStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRemoteConsole 获取console地址信息
//
// 获取console信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) ShowRemoteConsole(request *model.ShowRemoteConsoleRequest) (*model.ShowRemoteConsoleResponse, error) {
	requestDef := GenReqDefForShowRemoteConsole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRemoteConsoleResponse), nil
	}
}

// ShowRemoteConsoleInvoker 获取console地址信息
func (c *CloudDCClient) ShowRemoteConsoleInvoker(request *model.ShowRemoteConsoleRequest) *ShowRemoteConsoleInvoker {
	requestDef := GenReqDefForShowRemoteConsole()
	return &ShowRemoteConsoleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowServer 查询物理服务器信息
//
// Get imetal server by id
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) ShowServer(request *model.ShowServerRequest) (*model.ShowServerResponse, error) {
	requestDef := GenReqDefForShowServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerResponse), nil
	}
}

// ShowServerInvoker 查询物理服务器信息
func (c *CloudDCClient) ShowServerInvoker(request *model.ShowServerRequest) *ShowServerInvoker {
	requestDef := GenReqDefForShowServer()
	return &ShowServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowServerFirmwareAttributes 查询服务器固件详细信息
//
// 获取详细固件信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) ShowServerFirmwareAttributes(request *model.ShowServerFirmwareAttributesRequest) (*model.ShowServerFirmwareAttributesResponse, error) {
	requestDef := GenReqDefForShowServerFirmwareAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerFirmwareAttributesResponse), nil
	}
}

// ShowServerFirmwareAttributesInvoker 查询服务器固件详细信息
func (c *CloudDCClient) ShowServerFirmwareAttributesInvoker(request *model.ShowServerFirmwareAttributesRequest) *ShowServerFirmwareAttributesInvoker {
	requestDef := GenReqDefForShowServerFirmwareAttributes()
	return &ShowServerFirmwareAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowServerHardwareAttributes 查询服务器硬件详细信息
//
// 获取详细硬件信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) ShowServerHardwareAttributes(request *model.ShowServerHardwareAttributesRequest) (*model.ShowServerHardwareAttributesResponse, error) {
	requestDef := GenReqDefForShowServerHardwareAttributes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerHardwareAttributesResponse), nil
	}
}

// ShowServerHardwareAttributesInvoker 查询服务器硬件详细信息
func (c *CloudDCClient) ShowServerHardwareAttributesInvoker(request *model.ShowServerHardwareAttributesRequest) *ShowServerHardwareAttributesInvoker {
	requestDef := GenReqDefForShowServerHardwareAttributes()
	return &ShowServerHardwareAttributesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowServerStatus 服务器概览
//
// 该 API 用于查询服务器概览、服务器开机状态和服务器运行状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) ShowServerStatus(request *model.ShowServerStatusRequest) (*model.ShowServerStatusResponse, error) {
	requestDef := GenReqDefForShowServerStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerStatusResponse), nil
	}
}

// ShowServerStatusInvoker 服务器概览
func (c *CloudDCClient) ShowServerStatusInvoker(request *model.ShowServerStatusRequest) *ShowServerStatusInvoker {
	requestDef := GenReqDefForShowServerStatus()
	return &ShowServerStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateIDcs 修改 IDC 描述
//
// 修改 IDC 描述
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) UpdateIDcs(request *model.UpdateIDcsRequest) (*model.UpdateIDcsResponse, error) {
	requestDef := GenReqDefForUpdateIDcs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIDcsResponse), nil
	}
}

// UpdateIDcsInvoker 修改 IDC 描述
func (c *CloudDCClient) UpdateIDcsInvoker(request *model.UpdateIDcsRequest) *UpdateIDcsInvoker {
	requestDef := GenReqDefForUpdateIDcs()
	return &UpdateIDcsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateIRack 更新 iRack 实例
//
// 更新iRack信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudDCClient) UpdateIRack(request *model.UpdateIRackRequest) (*model.UpdateIRackResponse, error) {
	requestDef := GenReqDefForUpdateIRack()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIRackResponse), nil
	}
}

// UpdateIRackInvoker 更新 iRack 实例
func (c *CloudDCClient) UpdateIRackInvoker(request *model.UpdateIRackRequest) *UpdateIRackInvoker {
	requestDef := GenReqDefForUpdateIRack()
	return &UpdateIRackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
