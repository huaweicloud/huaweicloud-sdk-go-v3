package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/mrs/v2/model"
)

type MrsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewMrsClient(hcClient *http_client.HcHttpClient) *MrsClient {
	return &MrsClient{HcClient: hcClient}
}

func MrsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//在MRS集群中批量删除作业。
func (c *MrsClient) BatchDeleteJobs(request *model.BatchDeleteJobsRequest) (*model.BatchDeleteJobsResponse, error) {
	requestDef := GenReqDefForBatchDeleteJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteJobsResponse), nil
	}
}

//创建一个MRS集群。  使用接口前，您需要先获取下的资源信息。 - 通过VPC创建或查询VPC、子网 - 通过ECS创建或查询密钥对 - 通过[终端节点](https://support.huaweicloud.com/api-mrs/mrs_02_0003.html)获取区域信息 - 参考[MRS服务支持的组件](https://support.huaweicloud.com/api-mrs/mrs_02_9001.html)获取MRS版本及对应版本支持的组件信息
func (c *MrsClient) CreateCluster(request *model.CreateClusterRequest) (*model.CreateClusterResponse, error) {
	requestDef := GenReqDefForCreateCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterResponse), nil
	}
}

//在MRS集群中新增并提交一个作业。  需要先在集群详情页的“概览”页签，单击“IAM用户同步”右侧的“同步”进行IAM用户同步，然后再通过该接口提交作业。  如需使用OBS加密功能，请先参考“MRS用户指南 > 管理现有集群 > 作业管理 > 使用OBS加密数据运行作业”页面进行相关配置后，再调用API接口运行作业。  所有示例中涉及的OBS路径、样例文件及终端节点和AKSK，请提前准备并在提交请求时根据实际情况替换。
func (c *MrsClient) CreateExecuteJob(request *model.CreateExecuteJobRequest) (*model.CreateExecuteJobResponse, error) {
	requestDef := GenReqDefForCreateExecuteJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateExecuteJobResponse), nil
	}
}

//获取用户（组）与IAM委托之间的映射关系的详细信息。
func (c *MrsClient) ShowAgencyMapping(request *model.ShowAgencyMappingRequest) (*model.ShowAgencyMappingResponse, error) {
	requestDef := GenReqDefForShowAgencyMapping()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAgencyMappingResponse), nil
	}
}

//在MRS指定集群中查询作业列表信息。
func (c *MrsClient) ShowJobExeListNew(request *model.ShowJobExeListNewRequest) (*model.ShowJobExeListNewResponse, error) {
	requestDef := GenReqDefForShowJobExeListNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobExeListNewResponse), nil
	}
}

//在MRS集群中查询指定作业的详细信息。
func (c *MrsClient) ShowSingleJobExe(request *model.ShowSingleJobExeRequest) (*model.ShowSingleJobExeResponse, error) {
	requestDef := GenReqDefForShowSingleJobExe()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSingleJobExeResponse), nil
	}
}

//在MRS集群中查询SparkSql和SparkScript两种类型作业的SQL语句运行完成后返回的查询结果。
func (c *MrsClient) ShowSqlResultWithJob(request *model.ShowSqlResultWithJobRequest) (*model.ShowSqlResultWithJobResponse, error) {
	requestDef := GenReqDefForShowSqlResultWithJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlResultWithJobResponse), nil
	}
}

//在MRS集群中终止指定作业。
func (c *MrsClient) StopJob(request *model.StopJobRequest) (*model.StopJobResponse, error) {
	requestDef := GenReqDefForStopJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopJobResponse), nil
	}
}

//更新用户（组）与IAM委托之间的映射关系。
func (c *MrsClient) UpdateAgencyMapping(request *model.UpdateAgencyMappingRequest) (*model.UpdateAgencyMappingResponse, error) {
	requestDef := GenReqDefForUpdateAgencyMapping()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAgencyMappingResponse), nil
	}
}

//在MRS集群中获取指定目录文件列表。
func (c *MrsClient) ShowHdfsFileList(request *model.ShowHdfsFileListRequest) (*model.ShowHdfsFileListResponse, error) {
	requestDef := GenReqDefForShowHdfsFileList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHdfsFileListResponse), nil
	}
}

//在MRS集群中取消一条SQL的执行任务。
func (c *MrsClient) CancelSql(request *model.CancelSqlRequest) (*model.CancelSqlResponse, error) {
	requestDef := GenReqDefForCancelSql()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelSqlResponse), nil
	}
}

//在MRS集群中提交并执行一条SQL语句。
func (c *MrsClient) ExecuteSql(request *model.ExecuteSqlRequest) (*model.ExecuteSqlResponse, error) {
	requestDef := GenReqDefForExecuteSql()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteSqlResponse), nil
	}
}

//在MRS集群中查询一条SQL的执行结果。
func (c *MrsClient) ShowSqlResult(request *model.ShowSqlResultRequest) (*model.ShowSqlResultResponse, error) {
	requestDef := GenReqDefForShowSqlResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlResultResponse), nil
	}
}
