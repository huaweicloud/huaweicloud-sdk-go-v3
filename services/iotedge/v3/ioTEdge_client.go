package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iotedge/v3/model"
)

type IoTEdgeClient struct {
	HcClient *http_client.HcHttpClient
}

func NewIoTEdgeClient(hcClient *http_client.HcHttpClient) *IoTEdgeClient {
	return &IoTEdgeClient{HcClient: hcClient}
}

func IoTEdgeClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CreateApp 创建应用模板
//
// 应用服务器可调用此接口为创建批量处理任务，对多个设备进行批量操作。当前支持批量软固件升级、批量创建设备、批量删除设备、批量冻结、批量解冻、批量下发同步命令、批量下发异步命令。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) CreateApp(request *model.CreateAppRequest) (*model.CreateAppResponse, error) {
	requestDef := GenReqDefForCreateApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppResponse), nil
	}
}

// CreateAppInvoker 创建应用模板
func (c *IoTEdgeClient) CreateAppInvoker(request *model.CreateAppRequest) *CreateAppInvoker {
	requestDef := GenReqDefForCreateApp()
	return &CreateAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApp 删除应用模板
//
// 应用服务器可调用此接口删除应用模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) DeleteApp(request *model.DeleteAppRequest) (*model.DeleteAppResponse, error) {
	requestDef := GenReqDefForDeleteApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppResponse), nil
	}
}

// DeleteAppInvoker 删除应用模板
func (c *IoTEdgeClient) DeleteAppInvoker(request *model.DeleteAppRequest) *DeleteAppInvoker {
	requestDef := GenReqDefForDeleteApp()
	return &DeleteAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApps 查询应用模板列表
//
// 应用服务器可调用此接口查询应用模板列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ListApps(request *model.ListAppsRequest) (*model.ListAppsResponse, error) {
	requestDef := GenReqDefForListApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppsResponse), nil
	}
}

// ListAppsInvoker 查询应用模板列表
func (c *IoTEdgeClient) ListAppsInvoker(request *model.ListAppsRequest) *ListAppsInvoker {
	requestDef := GenReqDefForListApps()
	return &ListAppsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApp 查询应用模板详情
//
// 应用服务器可调用此接口查询物联网平台中指定批量任务的信息，包括任务内容、任务状态、任务完成情况统计以及子任务列表等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ShowApp(request *model.ShowAppRequest) (*model.ShowAppResponse, error) {
	requestDef := GenReqDefForShowApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppResponse), nil
	}
}

// ShowAppInvoker 查询应用模板详情
func (c *IoTEdgeClient) ShowAppInvoker(request *model.ShowAppRequest) *ShowAppInvoker {
	requestDef := GenReqDefForShowApp()
	return &ShowAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAppInstance 创建应用实例
//
// 应用服务器可调用此接口为创建应用实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) CreateAppInstance(request *model.CreateAppInstanceRequest) (*model.CreateAppInstanceResponse, error) {
	requestDef := GenReqDefForCreateAppInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppInstanceResponse), nil
	}
}

// CreateAppInstanceInvoker 创建应用实例
func (c *IoTEdgeClient) CreateAppInstanceInvoker(request *model.CreateAppInstanceRequest) *CreateAppInstanceInvoker {
	requestDef := GenReqDefForCreateAppInstance()
	return &CreateAppInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAppInstance 删除应用实例
//
// 应用服务器可调用此接口为删除应用实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) DeleteAppInstance(request *model.DeleteAppInstanceRequest) (*model.DeleteAppInstanceResponse, error) {
	requestDef := GenReqDefForDeleteAppInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppInstanceResponse), nil
	}
}

// DeleteAppInstanceInvoker 删除应用实例
func (c *IoTEdgeClient) DeleteAppInstanceInvoker(request *model.DeleteAppInstanceRequest) *DeleteAppInstanceInvoker {
	requestDef := GenReqDefForDeleteAppInstance()
	return &DeleteAppInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppInstanceHistory 查询应用实例的历史版本列表
//
// 应用服务器可调用此接口查询应用实例的历史版本列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ListAppInstanceHistory(request *model.ListAppInstanceHistoryRequest) (*model.ListAppInstanceHistoryResponse, error) {
	requestDef := GenReqDefForListAppInstanceHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppInstanceHistoryResponse), nil
	}
}

// ListAppInstanceHistoryInvoker 查询应用实例的历史版本列表
func (c *IoTEdgeClient) ListAppInstanceHistoryInvoker(request *model.ListAppInstanceHistoryRequest) *ListAppInstanceHistoryInvoker {
	requestDef := GenReqDefForListAppInstanceHistory()
	return &ListAppInstanceHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppInstances 查询应用实例列表
//
// 应用服务器可调用此接口查询应用实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ListAppInstances(request *model.ListAppInstancesRequest) (*model.ListAppInstancesResponse, error) {
	requestDef := GenReqDefForListAppInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppInstancesResponse), nil
	}
}

// ListAppInstancesInvoker 查询应用实例列表
func (c *IoTEdgeClient) ListAppInstancesInvoker(request *model.ListAppInstancesRequest) *ListAppInstancesInvoker {
	requestDef := GenReqDefForListAppInstances()
	return &ListAppInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAppInstance 更新应用实例
//
// 应用服务器可调用此接口为更新应用实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) UpdateAppInstance(request *model.UpdateAppInstanceRequest) (*model.UpdateAppInstanceResponse, error) {
	requestDef := GenReqDefForUpdateAppInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppInstanceResponse), nil
	}
}

// UpdateAppInstanceInvoker 更新应用实例
func (c *IoTEdgeClient) UpdateAppInstanceInvoker(request *model.UpdateAppInstanceRequest) *UpdateAppInstanceInvoker {
	requestDef := GenReqDefForUpdateAppInstance()
	return &UpdateAppInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAppVersion 创建应用版本
//
// 应用服务器可调用此接口为创建应用版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) CreateAppVersion(request *model.CreateAppVersionRequest) (*model.CreateAppVersionResponse, error) {
	requestDef := GenReqDefForCreateAppVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppVersionResponse), nil
	}
}

// CreateAppVersionInvoker 创建应用版本
func (c *IoTEdgeClient) CreateAppVersionInvoker(request *model.CreateAppVersionRequest) *CreateAppVersionInvoker {
	requestDef := GenReqDefForCreateAppVersion()
	return &CreateAppVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAppVersion 删除应用版本
//
// 应用服务器可调用此接口删除应用版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) DeleteAppVersion(request *model.DeleteAppVersionRequest) (*model.DeleteAppVersionResponse, error) {
	requestDef := GenReqDefForDeleteAppVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppVersionResponse), nil
	}
}

// DeleteAppVersionInvoker 删除应用版本
func (c *IoTEdgeClient) DeleteAppVersionInvoker(request *model.DeleteAppVersionRequest) *DeleteAppVersionInvoker {
	requestDef := GenReqDefForDeleteAppVersion()
	return &DeleteAppVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadAppVersion 下载应用版本Chart包
//
// 应用服务器可调用此接口下载应用版本Chart包。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) DownloadAppVersion(request *model.DownloadAppVersionRequest) (*model.DownloadAppVersionResponse, error) {
	requestDef := GenReqDefForDownloadAppVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadAppVersionResponse), nil
	}
}

// DownloadAppVersionInvoker 下载应用版本Chart包
func (c *IoTEdgeClient) DownloadAppVersionInvoker(request *model.DownloadAppVersionRequest) *DownloadAppVersionInvoker {
	requestDef := GenReqDefForDownloadAppVersion()
	return &DownloadAppVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppImage 查询应用版本包含的镜像列表
//
// 应用服务器可调用此接口查询应用版本包含的镜像列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ListAppImage(request *model.ListAppImageRequest) (*model.ListAppImageResponse, error) {
	requestDef := GenReqDefForListAppImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppImageResponse), nil
	}
}

// ListAppImageInvoker 查询应用版本包含的镜像列表
func (c *IoTEdgeClient) ListAppImageInvoker(request *model.ListAppImageRequest) *ListAppImageInvoker {
	requestDef := GenReqDefForListAppImage()
	return &ListAppImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppVersions 查询应用版本列表
//
// 应用服务器可调用此接口查询应用版本列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ListAppVersions(request *model.ListAppVersionsRequest) (*model.ListAppVersionsResponse, error) {
	requestDef := GenReqDefForListAppVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppVersionsResponse), nil
	}
}

// ListAppVersionsInvoker 查询应用版本列表
func (c *IoTEdgeClient) ListAppVersionsInvoker(request *model.ListAppVersionsRequest) *ListAppVersionsInvoker {
	requestDef := GenReqDefForListAppVersions()
	return &ListAppVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAppVersion 查询应用版本详情
//
// 应用服务器可调用此接口查询应用版本详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ShowAppVersion(request *model.ShowAppVersionRequest) (*model.ShowAppVersionResponse, error) {
	requestDef := GenReqDefForShowAppVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppVersionResponse), nil
	}
}

// ShowAppVersionInvoker 查询应用版本详情
func (c *IoTEdgeClient) ShowAppVersionInvoker(request *model.ShowAppVersionRequest) *ShowAppVersionInvoker {
	requestDef := GenReqDefForShowAppVersion()
	return &ShowAppVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCluster 创建边缘集群
//
// 应用服务器可调用此接口为创建边缘集群。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) CreateCluster(request *model.CreateClusterRequest) (*model.CreateClusterResponse, error) {
	requestDef := GenReqDefForCreateCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterResponse), nil
	}
}

// CreateClusterInvoker 创建边缘集群
func (c *IoTEdgeClient) CreateClusterInvoker(request *model.CreateClusterRequest) *CreateClusterInvoker {
	requestDef := GenReqDefForCreateCluster()
	return &CreateClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateClusterInstallCmd 生成边缘集群安装命令
//
// 应用服务器可调用此接口生成边缘集群安装命令。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) CreateClusterInstallCmd(request *model.CreateClusterInstallCmdRequest) (*model.CreateClusterInstallCmdResponse, error) {
	requestDef := GenReqDefForCreateClusterInstallCmd()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterInstallCmdResponse), nil
	}
}

// CreateClusterInstallCmdInvoker 生成边缘集群安装命令
func (c *IoTEdgeClient) CreateClusterInstallCmdInvoker(request *model.CreateClusterInstallCmdRequest) *CreateClusterInstallCmdInvoker {
	requestDef := GenReqDefForCreateClusterInstallCmd()
	return &CreateClusterInstallCmdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCluster 删除边缘集群
//
// 应用服务器可调用此接口删除边缘集群。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) DeleteCluster(request *model.DeleteClusterRequest) (*model.DeleteClusterResponse, error) {
	requestDef := GenReqDefForDeleteCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClusterResponse), nil
	}
}

// DeleteClusterInvoker 删除边缘集群
func (c *IoTEdgeClient) DeleteClusterInvoker(request *model.DeleteClusterRequest) *DeleteClusterInvoker {
	requestDef := GenReqDefForDeleteCluster()
	return &DeleteClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusters 查询边缘集群列表
//
// 应用服务器可调用此接口查询边缘集群列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ListClusters(request *model.ListClustersRequest) (*model.ListClustersResponse, error) {
	requestDef := GenReqDefForListClusters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClustersResponse), nil
	}
}

// ListClustersInvoker 查询边缘集群列表
func (c *IoTEdgeClient) ListClustersInvoker(request *model.ListClustersRequest) *ListClustersInvoker {
	requestDef := GenReqDefForListClusters()
	return &ListClustersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCluster 查询边缘集群详情
//
// 应用服务器可调用此接口查询边缘集群详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTEdgeClient) ShowCluster(request *model.ShowClusterRequest) (*model.ShowClusterResponse, error) {
	requestDef := GenReqDefForShowCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterResponse), nil
	}
}

// ShowClusterInvoker 查询边缘集群详情
func (c *IoTEdgeClient) ShowClusterInvoker(request *model.ShowClusterRequest) *ShowClusterInvoker {
	requestDef := GenReqDefForShowCluster()
	return &ShowClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
