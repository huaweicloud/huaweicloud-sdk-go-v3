package v1

import (
	http_client "code.byted.org/ti/huaweicloud-sdk-go-v3/core"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/services/cdm/v1/model"
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

//随机集群创建作业并执行接口。
func (c *CdmClient) CreateAndStartRandomClusterJob(request *model.CreateAndStartRandomClusterJobRequest) (*model.CreateAndStartRandomClusterJobResponse, error) {
	requestDef := GenReqDefForCreateAndStartRandomClusterJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAndStartRandomClusterJobResponse), nil
	}
}

//创建集群接口。
func (c *CdmClient) CreateCluster(request *model.CreateClusterRequest) (*model.CreateClusterResponse, error) {
	requestDef := GenReqDefForCreateCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterResponse), nil
	}
}

//指定集群创建作业接口。
func (c *CdmClient) CreateJob(request *model.CreateJobRequest) (*model.CreateJobResponse, error) {
	requestDef := GenReqDefForCreateJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateJobResponse), nil
	}
}

//创建连接接口。
func (c *CdmClient) CreateLink(request *model.CreateLinkRequest) (*model.CreateLinkResponse, error) {
	requestDef := GenReqDefForCreateLink()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLinkResponse), nil
	}
}

//删除集群接口。
func (c *CdmClient) DeleteCluster(request *model.DeleteClusterRequest) (*model.DeleteClusterResponse, error) {
	requestDef := GenReqDefForDeleteCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClusterResponse), nil
	}
}

//删除作业接口。
func (c *CdmClient) DeleteJob(request *model.DeleteJobRequest) (*model.DeleteJobResponse, error) {
	requestDef := GenReqDefForDeleteJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteJobResponse), nil
	}
}

//删除连接接口。
func (c *CdmClient) DeleteLink(request *model.DeleteLinkRequest) (*model.DeleteLinkResponse, error) {
	requestDef := GenReqDefForDeleteLink()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLinkResponse), nil
	}
}

//查询集群列表接口。
func (c *CdmClient) ListClusters(request *model.ListClustersRequest) (*model.ListClustersResponse, error) {
	requestDef := GenReqDefForListClusters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClustersResponse), nil
	}
}

//重启集群接口。
func (c *CdmClient) RestartCluster(request *model.RestartClusterRequest) (*model.RestartClusterResponse, error) {
	requestDef := GenReqDefForRestartCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartClusterResponse), nil
	}
}

//查询集群详情接口。
func (c *CdmClient) ShowClusterDetail(request *model.ShowClusterDetailRequest) (*model.ShowClusterDetailResponse, error) {
	requestDef := GenReqDefForShowClusterDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterDetailResponse), nil
	}
}

//查询作业状态接口。
func (c *CdmClient) ShowJobStatus(request *model.ShowJobStatusRequest) (*model.ShowJobStatusResponse, error) {
	requestDef := GenReqDefForShowJobStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobStatusResponse), nil
	}
}

//查询作业接口。
func (c *CdmClient) ShowJobs(request *model.ShowJobsRequest) (*model.ShowJobsResponse, error) {
	requestDef := GenReqDefForShowJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobsResponse), nil
	}
}

//查询连接接口。
func (c *CdmClient) ShowLink(request *model.ShowLinkRequest) (*model.ShowLinkResponse, error) {
	requestDef := GenReqDefForShowLink()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLinkResponse), nil
	}
}

//查询作业执行历史接口。
func (c *CdmClient) ShowSubmissions(request *model.ShowSubmissionsRequest) (*model.ShowSubmissionsResponse, error) {
	requestDef := GenReqDefForShowSubmissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSubmissionsResponse), nil
	}
}

//启动集群接口。
func (c *CdmClient) StartCluster(request *model.StartClusterRequest) (*model.StartClusterResponse, error) {
	requestDef := GenReqDefForStartCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartClusterResponse), nil
	}
}

//启动作业接口。
func (c *CdmClient) StartJob(request *model.StartJobRequest) (*model.StartJobResponse, error) {
	requestDef := GenReqDefForStartJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartJobResponse), nil
	}
}

//停止集群接口。
func (c *CdmClient) StopCluster(request *model.StopClusterRequest) (*model.StopClusterResponse, error) {
	requestDef := GenReqDefForStopCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopClusterResponse), nil
	}
}

//停止作业接口。
func (c *CdmClient) StopJob(request *model.StopJobRequest) (*model.StopJobResponse, error) {
	requestDef := GenReqDefForStopJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopJobResponse), nil
	}
}

//修改作业接口。
func (c *CdmClient) UpdateJob(request *model.UpdateJobRequest) (*model.UpdateJobResponse, error) {
	requestDef := GenReqDefForUpdateJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateJobResponse), nil
	}
}

//修改连接接口。
func (c *CdmClient) UpdateLink(request *model.UpdateLinkRequest) (*model.UpdateLinkResponse, error) {
	requestDef := GenReqDefForUpdateLink()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLinkResponse), nil
	}
}
