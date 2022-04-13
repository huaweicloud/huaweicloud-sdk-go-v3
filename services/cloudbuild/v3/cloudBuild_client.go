package v3

import (
	http_client "code.byted.org/ti/huaweicloud-sdk-go-v3/core"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/services/cloudbuild/v3/model"
)

type CloudBuildClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCloudBuildClient(hcClient *http_client.HcHttpClient) *CloudBuildClient {
	return &CloudBuildClient{HcClient: hcClient}
}

func CloudBuildClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//下载指定租户下的KeyStore文件
func (c *CloudBuildClient) DownloadKeystore(request *model.DownloadKeystoreRequest) (*model.DownloadKeystoreResponse, error) {
	requestDef := GenReqDefForDownloadKeystore()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadKeystoreResponse), nil
	}
}

//执行构建任务,可传自定义参数。
func (c *CloudBuildClient) RunJob(request *model.RunJobRequest) (*model.RunJobResponse, error) {
	requestDef := GenReqDefForRunJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunJobResponse), nil
	}
}

//获取构建历史详情信息接口
func (c *CloudBuildClient) ShowHistoryDetails(request *model.ShowHistoryDetailsRequest) (*model.ShowHistoryDetailsResponse, error) {
	requestDef := GenReqDefForShowHistoryDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHistoryDetailsResponse), nil
	}
}

//查看项目下用户的构建任务列表
func (c *CloudBuildClient) ShowJobListByProjectId(request *model.ShowJobListByProjectIdRequest) (*model.ShowJobListByProjectIdResponse, error) {
	requestDef := GenReqDefForShowJobListByProjectId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobListByProjectIdResponse), nil
	}
}

//查看任务运行状态
func (c *CloudBuildClient) ShowJobStatus(request *model.ShowJobStatusRequest) (*model.ShowJobStatusResponse, error) {
	requestDef := GenReqDefForShowJobStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobStatusResponse), nil
	}
}

//查询指定代码仓库最近一次成功的构建历史
func (c *CloudBuildClient) ShowLastHistory(request *model.ShowLastHistoryRequest) (*model.ShowLastHistoryResponse, error) {
	requestDef := GenReqDefForShowLastHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLastHistoryResponse), nil
	}
}

//查看构建任务的构建历史列表
func (c *CloudBuildClient) ShowListHistory(request *model.ShowListHistoryRequest) (*model.ShowListHistoryResponse, error) {
	requestDef := GenReqDefForShowListHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowListHistoryResponse), nil
	}
}

//根据开始时间和结束时间查看构建任务的构建历史列表
func (c *CloudBuildClient) ShowListPeriodHistory(request *model.ShowListPeriodHistoryRequest) (*model.ShowListPeriodHistoryResponse, error) {
	requestDef := GenReqDefForShowListPeriodHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowListPeriodHistoryResponse), nil
	}
}
