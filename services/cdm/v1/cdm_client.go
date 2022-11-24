package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cdm/v1/model"
)

type CdmClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCdmClient(hcClient *http_client.HcHttpClient) *CdmClient {
	return &CdmClient{HcClient: hcClient}
}

func CdmClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CreateAndStartRandomClusterJob 随机集群创建作业并执行
//
// 随机集群创建作业并执行接口。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdmClient) CreateAndStartRandomClusterJob(request *model.CreateAndStartRandomClusterJobRequest) (*model.CreateAndStartRandomClusterJobResponse, error) {
	requestDef := GenReqDefForCreateAndStartRandomClusterJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAndStartRandomClusterJobResponse), nil
	}
}

// CreateAndStartRandomClusterJobInvoker 随机集群创建作业并执行
func (c *CdmClient) CreateAndStartRandomClusterJobInvoker(request *model.CreateAndStartRandomClusterJobRequest) *CreateAndStartRandomClusterJobInvoker {
	requestDef := GenReqDefForCreateAndStartRandomClusterJob()
	return &CreateAndStartRandomClusterJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCluster 创建集群
//
// 创建集群接口。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdmClient) CreateCluster(request *model.CreateClusterRequest) (*model.CreateClusterResponse, error) {
	requestDef := GenReqDefForCreateCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterResponse), nil
	}
}

// CreateClusterInvoker 创建集群
func (c *CdmClient) CreateClusterInvoker(request *model.CreateClusterRequest) *CreateClusterInvoker {
	requestDef := GenReqDefForCreateCluster()
	return &CreateClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateJob 指定集群创建作业
//
// 指定集群创建作业接口。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdmClient) CreateJob(request *model.CreateJobRequest) (*model.CreateJobResponse, error) {
	requestDef := GenReqDefForCreateJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateJobResponse), nil
	}
}

// CreateJobInvoker 指定集群创建作业
func (c *CdmClient) CreateJobInvoker(request *model.CreateJobRequest) *CreateJobInvoker {
	requestDef := GenReqDefForCreateJob()
	return &CreateJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLink 创建连接
//
// 创建连接接口。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdmClient) CreateLink(request *model.CreateLinkRequest) (*model.CreateLinkResponse, error) {
	requestDef := GenReqDefForCreateLink()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLinkResponse), nil
	}
}

// CreateLinkInvoker 创建连接
func (c *CdmClient) CreateLinkInvoker(request *model.CreateLinkRequest) *CreateLinkInvoker {
	requestDef := GenReqDefForCreateLink()
	return &CreateLinkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCluster 删除集群
//
// 删除集群接口。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdmClient) DeleteCluster(request *model.DeleteClusterRequest) (*model.DeleteClusterResponse, error) {
	requestDef := GenReqDefForDeleteCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClusterResponse), nil
	}
}

// DeleteClusterInvoker 删除集群
func (c *CdmClient) DeleteClusterInvoker(request *model.DeleteClusterRequest) *DeleteClusterInvoker {
	requestDef := GenReqDefForDeleteCluster()
	return &DeleteClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteJob 删除作业
//
// 删除作业接口。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdmClient) DeleteJob(request *model.DeleteJobRequest) (*model.DeleteJobResponse, error) {
	requestDef := GenReqDefForDeleteJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteJobResponse), nil
	}
}

// DeleteJobInvoker 删除作业
func (c *CdmClient) DeleteJobInvoker(request *model.DeleteJobRequest) *DeleteJobInvoker {
	requestDef := GenReqDefForDeleteJob()
	return &DeleteJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLink 删除连接
//
// 删除连接接口。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdmClient) DeleteLink(request *model.DeleteLinkRequest) (*model.DeleteLinkResponse, error) {
	requestDef := GenReqDefForDeleteLink()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLinkResponse), nil
	}
}

// DeleteLinkInvoker 删除连接
func (c *CdmClient) DeleteLinkInvoker(request *model.DeleteLinkRequest) *DeleteLinkInvoker {
	requestDef := GenReqDefForDeleteLink()
	return &DeleteLinkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusters 查询集群列表
//
// 查询集群列表接口。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdmClient) ListClusters(request *model.ListClustersRequest) (*model.ListClustersResponse, error) {
	requestDef := GenReqDefForListClusters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClustersResponse), nil
	}
}

// ListClustersInvoker 查询集群列表
func (c *CdmClient) ListClustersInvoker(request *model.ListClustersRequest) *ListClustersInvoker {
	requestDef := GenReqDefForListClusters()
	return &ListClustersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestartCluster 重启集群
//
// 重启集群接口。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdmClient) RestartCluster(request *model.RestartClusterRequest) (*model.RestartClusterResponse, error) {
	requestDef := GenReqDefForRestartCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartClusterResponse), nil
	}
}

// RestartClusterInvoker 重启集群
func (c *CdmClient) RestartClusterInvoker(request *model.RestartClusterRequest) *RestartClusterInvoker {
	requestDef := GenReqDefForRestartCluster()
	return &RestartClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowClusterDetail 查询集群详情
//
// 查询集群详情接口。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdmClient) ShowClusterDetail(request *model.ShowClusterDetailRequest) (*model.ShowClusterDetailResponse, error) {
	requestDef := GenReqDefForShowClusterDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterDetailResponse), nil
	}
}

// ShowClusterDetailInvoker 查询集群详情
func (c *CdmClient) ShowClusterDetailInvoker(request *model.ShowClusterDetailRequest) *ShowClusterDetailInvoker {
	requestDef := GenReqDefForShowClusterDetail()
	return &ShowClusterDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobStatus 查询作业状态
//
// 查询作业状态接口。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdmClient) ShowJobStatus(request *model.ShowJobStatusRequest) (*model.ShowJobStatusResponse, error) {
	requestDef := GenReqDefForShowJobStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobStatusResponse), nil
	}
}

// ShowJobStatusInvoker 查询作业状态
func (c *CdmClient) ShowJobStatusInvoker(request *model.ShowJobStatusRequest) *ShowJobStatusInvoker {
	requestDef := GenReqDefForShowJobStatus()
	return &ShowJobStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobs 查询作业
//
// 查询作业接口。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdmClient) ShowJobs(request *model.ShowJobsRequest) (*model.ShowJobsResponse, error) {
	requestDef := GenReqDefForShowJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobsResponse), nil
	}
}

// ShowJobsInvoker 查询作业
func (c *CdmClient) ShowJobsInvoker(request *model.ShowJobsRequest) *ShowJobsInvoker {
	requestDef := GenReqDefForShowJobs()
	return &ShowJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLink 查询连接
//
// 查询连接接口。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdmClient) ShowLink(request *model.ShowLinkRequest) (*model.ShowLinkResponse, error) {
	requestDef := GenReqDefForShowLink()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLinkResponse), nil
	}
}

// ShowLinkInvoker 查询连接
func (c *CdmClient) ShowLinkInvoker(request *model.ShowLinkRequest) *ShowLinkInvoker {
	requestDef := GenReqDefForShowLink()
	return &ShowLinkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSubmissions 查询作业执行历史
//
// 查询作业执行历史接口。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdmClient) ShowSubmissions(request *model.ShowSubmissionsRequest) (*model.ShowSubmissionsResponse, error) {
	requestDef := GenReqDefForShowSubmissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSubmissionsResponse), nil
	}
}

// ShowSubmissionsInvoker 查询作业执行历史
func (c *CdmClient) ShowSubmissionsInvoker(request *model.ShowSubmissionsRequest) *ShowSubmissionsInvoker {
	requestDef := GenReqDefForShowSubmissions()
	return &ShowSubmissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartCluster 启动集群
//
// 启动集群接口。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdmClient) StartCluster(request *model.StartClusterRequest) (*model.StartClusterResponse, error) {
	requestDef := GenReqDefForStartCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartClusterResponse), nil
	}
}

// StartClusterInvoker 启动集群
func (c *CdmClient) StartClusterInvoker(request *model.StartClusterRequest) *StartClusterInvoker {
	requestDef := GenReqDefForStartCluster()
	return &StartClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartJob 启动作业
//
// 启动作业接口。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdmClient) StartJob(request *model.StartJobRequest) (*model.StartJobResponse, error) {
	requestDef := GenReqDefForStartJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartJobResponse), nil
	}
}

// StartJobInvoker 启动作业
func (c *CdmClient) StartJobInvoker(request *model.StartJobRequest) *StartJobInvoker {
	requestDef := GenReqDefForStartJob()
	return &StartJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopCluster 停止集群
//
// 停止集群接口。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdmClient) StopCluster(request *model.StopClusterRequest) (*model.StopClusterResponse, error) {
	requestDef := GenReqDefForStopCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopClusterResponse), nil
	}
}

// StopClusterInvoker 停止集群
func (c *CdmClient) StopClusterInvoker(request *model.StopClusterRequest) *StopClusterInvoker {
	requestDef := GenReqDefForStopCluster()
	return &StopClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopJob 停止作业
//
// 停止作业接口。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdmClient) StopJob(request *model.StopJobRequest) (*model.StopJobResponse, error) {
	requestDef := GenReqDefForStopJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopJobResponse), nil
	}
}

// StopJobInvoker 停止作业
func (c *CdmClient) StopJobInvoker(request *model.StopJobRequest) *StopJobInvoker {
	requestDef := GenReqDefForStopJob()
	return &StopJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateJob 修改作业
//
// 修改作业接口。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdmClient) UpdateJob(request *model.UpdateJobRequest) (*model.UpdateJobResponse, error) {
	requestDef := GenReqDefForUpdateJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateJobResponse), nil
	}
}

// UpdateJobInvoker 修改作业
func (c *CdmClient) UpdateJobInvoker(request *model.UpdateJobRequest) *UpdateJobInvoker {
	requestDef := GenReqDefForUpdateJob()
	return &UpdateJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLink 修改连接
//
// 修改连接接口。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CdmClient) UpdateLink(request *model.UpdateLinkRequest) (*model.UpdateLinkResponse, error) {
	requestDef := GenReqDefForUpdateLink()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLinkResponse), nil
	}
}

// UpdateLinkInvoker 修改连接
func (c *CdmClient) UpdateLinkInvoker(request *model.UpdateLinkRequest) *UpdateLinkInvoker {
	requestDef := GenReqDefForUpdateLink()
	return &UpdateLinkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
