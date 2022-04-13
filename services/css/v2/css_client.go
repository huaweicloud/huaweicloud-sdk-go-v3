package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/css/v2/model"
)

type CssClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCssClient(hcClient *http_client.HcHttpClient) *CssClient {
	return &CssClient{HcClient: hcClient}
}

func CssClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//该接口用于创建集群V2。
func (c *CssClient) CreateCluster(request *model.CreateClusterRequest) (*model.CreateClusterResponse, error) {
	requestDef := GenReqDefForCreateCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterResponse), nil
	}
}

//该接口用于重启集群。
func (c *CssClient) RestartCluster(request *model.RestartClusterRequest) (*model.RestartClusterResponse, error) {
	requestDef := GenReqDefForRestartCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartClusterResponse), nil
	}
}

//该接口用于滚动重启。
func (c *CssClient) RollingRestart(request *model.RollingRestartRequest) (*model.RollingRestartResponse, error) {
	requestDef := GenReqDefForRollingRestart()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RollingRestartResponse), nil
	}
}

//该接口用于打开自动创建快照功能。
func (c *CssClient) StartAutoCreateSnapshots(request *model.StartAutoCreateSnapshotsRequest) (*model.StartAutoCreateSnapshotsResponse, error) {
	requestDef := GenReqDefForStartAutoCreateSnapshots()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartAutoCreateSnapshotsResponse), nil
	}
}

//该接口用于关闭自动创建快照功能。
func (c *CssClient) StopAutoCreateSnapshots(request *model.StopAutoCreateSnapshotsRequest) (*model.StopAutoCreateSnapshotsResponse, error) {
	requestDef := GenReqDefForStopAutoCreateSnapshots()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopAutoCreateSnapshotsResponse), nil
	}
}
