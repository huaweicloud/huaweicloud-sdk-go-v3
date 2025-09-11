package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eihealth/v1/model"
)

type EiHealthClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewEiHealthClient(hcClient *httpclient.HcHttpClient) *EiHealthClient {
	return &EiHealthClient{HcClient: hcClient}
}

func EiHealthClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AddDrugDatabaseFile 数据库追加文件
//
// 数据库追加文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) AddDrugDatabaseFile(request *model.AddDrugDatabaseFileRequest) (*model.AddDrugDatabaseFileResponse, error) {
	requestDef := GenReqDefForAddDrugDatabaseFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddDrugDatabaseFileResponse), nil
	}
}

// AddDrugDatabaseFileInvoker 数据库追加文件
func (c *EiHealthClient) AddDrugDatabaseFileInvoker(request *model.AddDrugDatabaseFileRequest) *AddDrugDatabaseFileInvoker {
	requestDef := GenReqDefForAddDrugDatabaseFile()
	return &AddDrugDatabaseFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCancelJob 批量取消作业
//
// 批量取消作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) BatchCancelJob(request *model.BatchCancelJobRequest) (*model.BatchCancelJobResponse, error) {
	requestDef := GenReqDefForBatchCancelJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCancelJobResponse), nil
	}
}

// BatchCancelJobInvoker 批量取消作业
func (c *EiHealthClient) BatchCancelJobInvoker(request *model.BatchCancelJobRequest) *BatchCancelJobInvoker {
	requestDef := GenReqDefForBatchCancelJob()
	return &BatchCancelJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteData 批量删除项目数据
//
// 批量删除项目数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) BatchDeleteData(request *model.BatchDeleteDataRequest) (*model.BatchDeleteDataResponse, error) {
	requestDef := GenReqDefForBatchDeleteData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteDataResponse), nil
	}
}

// BatchDeleteDataInvoker 批量删除项目数据
func (c *EiHealthClient) BatchDeleteDataInvoker(request *model.BatchDeleteDataRequest) *BatchDeleteDataInvoker {
	requestDef := GenReqDefForBatchDeleteData()
	return &BatchDeleteDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteJob 批量删除作业
//
// 批量删除作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) BatchDeleteJob(request *model.BatchDeleteJobRequest) (*model.BatchDeleteJobResponse, error) {
	requestDef := GenReqDefForBatchDeleteJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteJobResponse), nil
	}
}

// BatchDeleteJobInvoker 批量删除作业
func (c *EiHealthClient) BatchDeleteJobInvoker(request *model.BatchDeleteJobRequest) *BatchDeleteJobInvoker {
	requestDef := GenReqDefForBatchDeleteJob()
	return &BatchDeleteJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteLabel 批量删除标签
//
// 批量删除标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) BatchDeleteLabel(request *model.BatchDeleteLabelRequest) (*model.BatchDeleteLabelResponse, error) {
	requestDef := GenReqDefForBatchDeleteLabel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteLabelResponse), nil
	}
}

// BatchDeleteLabelInvoker 批量删除标签
func (c *EiHealthClient) BatchDeleteLabelInvoker(request *model.BatchDeleteLabelRequest) *BatchDeleteLabelInvoker {
	requestDef := GenReqDefForBatchDeleteLabel()
	return &BatchDeleteLabelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchImportApp 导入应用
//
// 批量导入应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) BatchImportApp(request *model.BatchImportAppRequest) (*model.BatchImportAppResponse, error) {
	requestDef := GenReqDefForBatchImportApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchImportAppResponse), nil
	}
}

// BatchImportAppInvoker 导入应用
func (c *EiHealthClient) BatchImportAppInvoker(request *model.BatchImportAppRequest) *BatchImportAppInvoker {
	requestDef := GenReqDefForBatchImportApp()
	return &BatchImportAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchRetryJob 批量重试作业
//
// 批量重试作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) BatchRetryJob(request *model.BatchRetryJobRequest) (*model.BatchRetryJobResponse, error) {
	requestDef := GenReqDefForBatchRetryJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRetryJobResponse), nil
	}
}

// BatchRetryJobInvoker 批量重试作业
func (c *EiHealthClient) BatchRetryJobInvoker(request *model.BatchRetryJobRequest) *BatchRetryJobInvoker {
	requestDef := GenReqDefForBatchRetryJob()
	return &BatchRetryJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelDataJob 取消数据作业
//
// 取消数据作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CancelDataJob(request *model.CancelDataJobRequest) (*model.CancelDataJobResponse, error) {
	requestDef := GenReqDefForCancelDataJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelDataJobResponse), nil
	}
}

// CancelDataJobInvoker 取消数据作业
func (c *EiHealthClient) CancelDataJobInvoker(request *model.CancelDataJobRequest) *CancelDataJobInvoker {
	requestDef := GenReqDefForCancelDataJob()
	return &CancelDataJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelDrugJob 取消药物作业
//
// 取消药物作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CancelDrugJob(request *model.CancelDrugJobRequest) (*model.CancelDrugJobResponse, error) {
	requestDef := GenReqDefForCancelDrugJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelDrugJobResponse), nil
	}
}

// CancelDrugJobInvoker 取消药物作业
func (c *EiHealthClient) CancelDrugJobInvoker(request *model.CancelDrugJobRequest) *CancelDrugJobInvoker {
	requestDef := GenReqDefForCancelDrugJob()
	return &CancelDrugJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelJob 取消或强制停止作业调度
//
// 取消或强制作业调度
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CancelJob(request *model.CancelJobRequest) (*model.CancelJobResponse, error) {
	requestDef := GenReqDefForCancelJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelJobResponse), nil
	}
}

// CancelJobInvoker 取消或强制停止作业调度
func (c *EiHealthClient) CancelJobInvoker(request *model.CancelJobRequest) *CancelJobInvoker {
	requestDef := GenReqDefForCancelJob()
	return &CancelJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangePassword 修改密码
//
// 修改密码
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ChangePassword(request *model.ChangePasswordRequest) (*model.ChangePasswordResponse, error) {
	requestDef := GenReqDefForChangePassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangePasswordResponse), nil
	}
}

// ChangePasswordInvoker 修改密码
func (c *EiHealthClient) ChangePasswordInvoker(request *model.ChangePasswordRequest) *ChangePasswordInvoker {
	requestDef := GenReqDefForChangePassword()
	return &ChangePasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckTokenVerification 校验token
//
// 校验token是否可访问当前环境
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CheckTokenVerification(request *model.CheckTokenVerificationRequest) (*model.CheckTokenVerificationResponse, error) {
	requestDef := GenReqDefForCheckTokenVerification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckTokenVerificationResponse), nil
	}
}

// CheckTokenVerificationInvoker 校验token
func (c *EiHealthClient) CheckTokenVerificationInvoker(request *model.CheckTokenVerificationRequest) *CheckTokenVerificationInvoker {
	requestDef := GenReqDefForCheckTokenVerification()
	return &CheckTokenVerificationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CopyData 复制项目数据
//
// 复制项目数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CopyData(request *model.CopyDataRequest) (*model.CopyDataResponse, error) {
	requestDef := GenReqDefForCopyData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyDataResponse), nil
	}
}

// CopyDataInvoker 复制项目数据
func (c *EiHealthClient) CopyDataInvoker(request *model.CopyDataRequest) *CopyDataInvoker {
	requestDef := GenReqDefForCopyData()
	return &CopyDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAdmetJob 创建分子属性预测作业
//
// 创建分子属性预测作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateAdmetJob(request *model.CreateAdmetJobRequest) (*model.CreateAdmetJobResponse, error) {
	requestDef := GenReqDefForCreateAdmetJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAdmetJobResponse), nil
	}
}

// CreateAdmetJobInvoker 创建分子属性预测作业
func (c *EiHealthClient) CreateAdmetJobInvoker(request *model.CreateAdmetJobRequest) *CreateAdmetJobInvoker {
	requestDef := GenReqDefForCreateAdmetJob()
	return &CreateAdmetJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateApp 创建应用
//
// 创建应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateApp(request *model.CreateAppRequest) (*model.CreateAppResponse, error) {
	requestDef := GenReqDefForCreateApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppResponse), nil
	}
}

// CreateAppInvoker 创建应用
func (c *EiHealthClient) CreateAppInvoker(request *model.CreateAppRequest) *CreateAppInvoker {
	requestDef := GenReqDefForCreateApp()
	return &CreateAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateClusterJob 创建分子聚类作业
//
// 创建分子聚类作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateClusterJob(request *model.CreateClusterJobRequest) (*model.CreateClusterJobResponse, error) {
	requestDef := GenReqDefForCreateClusterJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterJobResponse), nil
	}
}

// CreateClusterJobInvoker 创建分子聚类作业
func (c *EiHealthClient) CreateClusterJobInvoker(request *model.CreateClusterJobRequest) *CreateClusterJobInvoker {
	requestDef := GenReqDefForCreateClusterJob()
	return &CreateClusterJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateClusteringJob 创建聚类分析作业
//
// 创建聚类分析作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateClusteringJob(request *model.CreateClusteringJobRequest) (*model.CreateClusteringJobResponse, error) {
	requestDef := GenReqDefForCreateClusteringJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusteringJobResponse), nil
	}
}

// CreateClusteringJobInvoker 创建聚类分析作业
func (c *EiHealthClient) CreateClusteringJobInvoker(request *model.CreateClusteringJobRequest) *CreateClusteringJobInvoker {
	requestDef := GenReqDefForCreateClusteringJob()
	return &CreateClusteringJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCode 发送验证码
//
// 发送验证码
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateCode(request *model.CreateCodeRequest) (*model.CreateCodeResponse, error) {
	requestDef := GenReqDefForCreateCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCodeResponse), nil
	}
}

// CreateCodeInvoker 发送验证码
func (c *EiHealthClient) CreateCodeInvoker(request *model.CreateCodeRequest) *CreateCodeInvoker {
	requestDef := GenReqDefForCreateCode()
	return &CreateCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateComputingCluster 绑定计算集群
//
// 绑定计算集群。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateComputingCluster(request *model.CreateComputingClusterRequest) (*model.CreateComputingClusterResponse, error) {
	requestDef := GenReqDefForCreateComputingCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateComputingClusterResponse), nil
	}
}

// CreateComputingClusterInvoker 绑定计算集群
func (c *EiHealthClient) CreateComputingClusterInvoker(request *model.CreateComputingClusterRequest) *CreateComputingClusterInvoker {
	requestDef := GenReqDefForCreateComputingCluster()
	return &CreateComputingClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateData 创建文件夹
//
// 创建文件夹
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateData(request *model.CreateDataRequest) (*model.CreateDataResponse, error) {
	requestDef := GenReqDefForCreateData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDataResponse), nil
	}
}

// CreateDataInvoker 创建文件夹
func (c *EiHealthClient) CreateDataInvoker(request *model.CreateDataRequest) *CreateDataInvoker {
	requestDef := GenReqDefForCreateData()
	return &CreateDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDockingJob 创建分子对接作业
//
// 创建分子对接作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateDockingJob(request *model.CreateDockingJobRequest) (*model.CreateDockingJobResponse, error) {
	requestDef := GenReqDefForCreateDockingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDockingJobResponse), nil
	}
}

// CreateDockingJobInvoker 创建分子对接作业
func (c *EiHealthClient) CreateDockingJobInvoker(request *model.CreateDockingJobRequest) *CreateDockingJobInvoker {
	requestDef := GenReqDefForCreateDockingJob()
	return &CreateDockingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDrugDatabase 创建数据库
//
// 创建数据库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateDrugDatabase(request *model.CreateDrugDatabaseRequest) (*model.CreateDrugDatabaseResponse, error) {
	requestDef := GenReqDefForCreateDrugDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDrugDatabaseResponse), nil
	}
}

// CreateDrugDatabaseInvoker 创建数据库
func (c *EiHealthClient) CreateDrugDatabaseInvoker(request *model.CreateDrugDatabaseRequest) *CreateDrugDatabaseInvoker {
	requestDef := GenReqDefForCreateDrugDatabase()
	return &CreateDrugDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDrugModel 创建模型
//
// 创建模型
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateDrugModel(request *model.CreateDrugModelRequest) (*model.CreateDrugModelResponse, error) {
	requestDef := GenReqDefForCreateDrugModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDrugModelResponse), nil
	}
}

// CreateDrugModelInvoker 创建模型
func (c *EiHealthClient) CreateDrugModelInvoker(request *model.CreateDrugModelRequest) *CreateDrugModelInvoker {
	requestDef := GenReqDefForCreateDrugModel()
	return &CreateDrugModelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDrugModelResource 创建盘古药物分子大模型
//
// 创建盘古药物分子大模型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateDrugModelResource(request *model.CreateDrugModelResourceRequest) (*model.CreateDrugModelResourceResponse, error) {
	requestDef := GenReqDefForCreateDrugModelResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDrugModelResourceResponse), nil
	}
}

// CreateDrugModelResourceInvoker 创建盘古药物分子大模型
func (c *EiHealthClient) CreateDrugModelResourceInvoker(request *model.CreateDrugModelResourceRequest) *CreateDrugModelResourceInvoker {
	requestDef := GenReqDefForCreateDrugModelResource()
	return &CreateDrugModelResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFavorite 添加收藏
//
// 添加收藏。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateFavorite(request *model.CreateFavoriteRequest) (*model.CreateFavoriteResponse, error) {
	requestDef := GenReqDefForCreateFavorite()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFavoriteResponse), nil
	}
}

// CreateFavoriteInvoker 添加收藏
func (c *EiHealthClient) CreateFavoriteInvoker(request *model.CreateFavoriteRequest) *CreateFavoriteInvoker {
	requestDef := GenReqDefForCreateFavorite()
	return &CreateFavoriteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFepJob 创建自由能微扰作业
//
// 创建自由能微扰作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateFepJob(request *model.CreateFepJobRequest) (*model.CreateFepJobResponse, error) {
	requestDef := GenReqDefForCreateFepJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFepJobResponse), nil
	}
}

// CreateFepJobInvoker 创建自由能微扰作业
func (c *EiHealthClient) CreateFepJobInvoker(request *model.CreateFepJobRequest) *CreateFepJobInvoker {
	requestDef := GenReqDefForCreateFepJob()
	return &CreateFepJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGenJob 创建分子生成作业
//
// 创建分子生成作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateGenJob(request *model.CreateGenJobRequest) (*model.CreateGenJobResponse, error) {
	requestDef := GenReqDefForCreateGenJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGenJobResponse), nil
	}
}

// CreateGenJobInvoker 创建分子生成作业
func (c *EiHealthClient) CreateGenJobInvoker(request *model.CreateGenJobRequest) *CreateGenJobInvoker {
	requestDef := GenReqDefForCreateGenJob()
	return &CreateGenJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateImage 创建镜像
//
// 创建镜像
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateImage(request *model.CreateImageRequest) (*model.CreateImageResponse, error) {
	requestDef := GenReqDefForCreateImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateImageResponse), nil
	}
}

// CreateImageInvoker 创建镜像
func (c *EiHealthClient) CreateImageInvoker(request *model.CreateImageRequest) *CreateImageInvoker {
	requestDef := GenReqDefForCreateImage()
	return &CreateImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLabel 创建标签
//
// 创建标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateLabel(request *model.CreateLabelRequest) (*model.CreateLabelResponse, error) {
	requestDef := GenReqDefForCreateLabel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLabelResponse), nil
	}
}

// CreateLabelInvoker 创建标签
func (c *EiHealthClient) CreateLabelInvoker(request *model.CreateLabelRequest) *CreateLabelInvoker {
	requestDef := GenReqDefForCreateLabel()
	return &CreateLabelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMolBatchDownloadTask 创建分子或分子复合物批量下载任务
//
// 创建分子或分子复合物批量下载任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateMolBatchDownloadTask(request *model.CreateMolBatchDownloadTaskRequest) (*model.CreateMolBatchDownloadTaskResponse, error) {
	requestDef := GenReqDefForCreateMolBatchDownloadTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMolBatchDownloadTaskResponse), nil
	}
}

// CreateMolBatchDownloadTaskInvoker 创建分子或分子复合物批量下载任务
func (c *EiHealthClient) CreateMolBatchDownloadTaskInvoker(request *model.CreateMolBatchDownloadTaskRequest) *CreateMolBatchDownloadTaskInvoker {
	requestDef := GenReqDefForCreateMolBatchDownloadTask()
	return &CreateMolBatchDownloadTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMolDockingJob 单分子预对接
//
// 单分子预对接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateMolDockingJob(request *model.CreateMolDockingJobRequest) (*model.CreateMolDockingJobResponse, error) {
	requestDef := GenReqDefForCreateMolDockingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMolDockingJobResponse), nil
	}
}

// CreateMolDockingJobInvoker 单分子预对接
func (c *EiHealthClient) CreateMolDockingJobInvoker(request *model.CreateMolDockingJobRequest) *CreateMolDockingJobInvoker {
	requestDef := GenReqDefForCreateMolDockingJob()
	return &CreateMolDockingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateOptmJob 创建分子优化作业
//
// 创建分子优化作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateOptmJob(request *model.CreateOptmJobRequest) (*model.CreateOptmJobResponse, error) {
	requestDef := GenReqDefForCreateOptmJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOptmJobResponse), nil
	}
}

// CreateOptmJobInvoker 创建分子优化作业
func (c *EiHealthClient) CreateOptmJobInvoker(request *model.CreateOptmJobRequest) *CreateOptmJobInvoker {
	requestDef := GenReqDefForCreateOptmJob()
	return &CreateOptmJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePerformanceResource 购买性能加速资源
//
// 购买性能加速资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreatePerformanceResource(request *model.CreatePerformanceResourceRequest) (*model.CreatePerformanceResourceResponse, error) {
	requestDef := GenReqDefForCreatePerformanceResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePerformanceResourceResponse), nil
	}
}

// CreatePerformanceResourceInvoker 购买性能加速资源
func (c *EiHealthClient) CreatePerformanceResourceInvoker(request *model.CreatePerformanceResourceRequest) *CreatePerformanceResourceInvoker {
	requestDef := GenReqDefForCreatePerformanceResource()
	return &CreatePerformanceResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePocketDetectionJob 创建靶点口袋发现作业
//
// 创建靶点口袋发现作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreatePocketDetectionJob(request *model.CreatePocketDetectionJobRequest) (*model.CreatePocketDetectionJobResponse, error) {
	requestDef := GenReqDefForCreatePocketDetectionJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePocketDetectionJobResponse), nil
	}
}

// CreatePocketDetectionJobInvoker 创建靶点口袋发现作业
func (c *EiHealthClient) CreatePocketDetectionJobInvoker(request *model.CreatePocketDetectionJobRequest) *CreatePocketDetectionJobInvoker {
	requestDef := GenReqDefForCreatePocketDetectionJob()
	return &CreatePocketDetectionJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePocketMolDesignJob 创建靶点口袋分子设计作业
//
// 创建靶点口袋分子设计作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreatePocketMolDesignJob(request *model.CreatePocketMolDesignJobRequest) (*model.CreatePocketMolDesignJobResponse, error) {
	requestDef := GenReqDefForCreatePocketMolDesignJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePocketMolDesignJobResponse), nil
	}
}

// CreatePocketMolDesignJobInvoker 创建靶点口袋分子设计作业
func (c *EiHealthClient) CreatePocketMolDesignJobInvoker(request *model.CreatePocketMolDesignJobRequest) *CreatePocketMolDesignJobInvoker {
	requestDef := GenReqDefForCreatePocketMolDesignJob()
	return &CreatePocketMolDesignJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateProject 创建项目
//
// 创建项目
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateProject(request *model.CreateProjectRequest) (*model.CreateProjectResponse, error) {
	requestDef := GenReqDefForCreateProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProjectResponse), nil
	}
}

// CreateProjectInvoker 创建项目
func (c *EiHealthClient) CreateProjectInvoker(request *model.CreateProjectRequest) *CreateProjectInvoker {
	requestDef := GenReqDefForCreateProject()
	return &CreateProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSearchJob 创建分子搜索作业
//
// 创建分子搜索作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateSearchJob(request *model.CreateSearchJobRequest) (*model.CreateSearchJobResponse, error) {
	requestDef := GenReqDefForCreateSearchJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSearchJobResponse), nil
	}
}

// CreateSearchJobInvoker 创建分子搜索作业
func (c *EiHealthClient) CreateSearchJobInvoker(request *model.CreateSearchJobRequest) *CreateSearchJobInvoker {
	requestDef := GenReqDefForCreateSearchJob()
	return &CreateSearchJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTargetOptJob 创建靶点优化作业
//
// 创建靶点优化作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateTargetOptJob(request *model.CreateTargetOptJobRequest) (*model.CreateTargetOptJobResponse, error) {
	requestDef := GenReqDefForCreateTargetOptJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTargetOptJobResponse), nil
	}
}

// CreateTargetOptJobInvoker 创建靶点优化作业
func (c *EiHealthClient) CreateTargetOptJobInvoker(request *model.CreateTargetOptJobRequest) *CreateTargetOptJobInvoker {
	requestDef := GenReqDefForCreateTargetOptJob()
	return &CreateTargetOptJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateUser 创建用户
//
// 创建用户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateUser(request *model.CreateUserRequest) (*model.CreateUserResponse, error) {
	requestDef := GenReqDefForCreateUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateUserResponse), nil
	}
}

// CreateUserInvoker 创建用户
func (c *EiHealthClient) CreateUserInvoker(request *model.CreateUserRequest) *CreateUserInvoker {
	requestDef := GenReqDefForCreateUser()
	return &CreateUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkflow 创建流程
//
// 创建流程
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateWorkflow(request *model.CreateWorkflowRequest) (*model.CreateWorkflowResponse, error) {
	requestDef := GenReqDefForCreateWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkflowResponse), nil
	}
}

// CreateWorkflowInvoker 创建流程
func (c *EiHealthClient) CreateWorkflowInvoker(request *model.CreateWorkflowRequest) *CreateWorkflowInvoker {
	requestDef := GenReqDefForCreateWorkflow()
	return &CreateWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApp 删除应用
//
// 删除应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteApp(request *model.DeleteAppRequest) (*model.DeleteAppResponse, error) {
	requestDef := GenReqDefForDeleteApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppResponse), nil
	}
}

// DeleteAppInvoker 删除应用
func (c *EiHealthClient) DeleteAppInvoker(request *model.DeleteAppRequest) *DeleteAppInvoker {
	requestDef := GenReqDefForDeleteApp()
	return &DeleteAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteComputingCluster 解绑计算集群
//
// 解绑计算集群。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteComputingCluster(request *model.DeleteComputingClusterRequest) (*model.DeleteComputingClusterResponse, error) {
	requestDef := GenReqDefForDeleteComputingCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteComputingClusterResponse), nil
	}
}

// DeleteComputingClusterInvoker 解绑计算集群
func (c *EiHealthClient) DeleteComputingClusterInvoker(request *model.DeleteComputingClusterRequest) *DeleteComputingClusterInvoker {
	requestDef := GenReqDefForDeleteComputingCluster()
	return &DeleteComputingClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDataJob 删除数据作业
//
// 删除数据作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteDataJob(request *model.DeleteDataJobRequest) (*model.DeleteDataJobResponse, error) {
	requestDef := GenReqDefForDeleteDataJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDataJobResponse), nil
	}
}

// DeleteDataJobInvoker 删除数据作业
func (c *EiHealthClient) DeleteDataJobInvoker(request *model.DeleteDataJobRequest) *DeleteDataJobInvoker {
	requestDef := GenReqDefForDeleteDataJob()
	return &DeleteDataJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDrugDatabase 删除数据库
//
// 删除数据库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteDrugDatabase(request *model.DeleteDrugDatabaseRequest) (*model.DeleteDrugDatabaseResponse, error) {
	requestDef := GenReqDefForDeleteDrugDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDrugDatabaseResponse), nil
	}
}

// DeleteDrugDatabaseInvoker 删除数据库
func (c *EiHealthClient) DeleteDrugDatabaseInvoker(request *model.DeleteDrugDatabaseRequest) *DeleteDrugDatabaseInvoker {
	requestDef := GenReqDefForDeleteDrugDatabase()
	return &DeleteDrugDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDrugJob 删除药物作业
//
// 删除药物作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteDrugJob(request *model.DeleteDrugJobRequest) (*model.DeleteDrugJobResponse, error) {
	requestDef := GenReqDefForDeleteDrugJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDrugJobResponse), nil
	}
}

// DeleteDrugJobInvoker 删除药物作业
func (c *EiHealthClient) DeleteDrugJobInvoker(request *model.DeleteDrugJobRequest) *DeleteDrugJobInvoker {
	requestDef := GenReqDefForDeleteDrugJob()
	return &DeleteDrugJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDrugModel 删除模型
//
// 删除模型
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteDrugModel(request *model.DeleteDrugModelRequest) (*model.DeleteDrugModelResponse, error) {
	requestDef := GenReqDefForDeleteDrugModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDrugModelResponse), nil
	}
}

// DeleteDrugModelInvoker 删除模型
func (c *EiHealthClient) DeleteDrugModelInvoker(request *model.DeleteDrugModelRequest) *DeleteDrugModelInvoker {
	requestDef := GenReqDefForDeleteDrugModel()
	return &DeleteDrugModelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDrugModelResource 退订盘古药物分子大模型
//
// 退订盘古药物分子大模型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteDrugModelResource(request *model.DeleteDrugModelResourceRequest) (*model.DeleteDrugModelResourceResponse, error) {
	requestDef := GenReqDefForDeleteDrugModelResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDrugModelResourceResponse), nil
	}
}

// DeleteDrugModelResourceInvoker 退订盘古药物分子大模型
func (c *EiHealthClient) DeleteDrugModelResourceInvoker(request *model.DeleteDrugModelResourceRequest) *DeleteDrugModelResourceInvoker {
	requestDef := GenReqDefForDeleteDrugModelResource()
	return &DeleteDrugModelResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFavorite 取消收藏
//
// 取消收藏。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteFavorite(request *model.DeleteFavoriteRequest) (*model.DeleteFavoriteResponse, error) {
	requestDef := GenReqDefForDeleteFavorite()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFavoriteResponse), nil
	}
}

// DeleteFavoriteInvoker 取消收藏
func (c *EiHealthClient) DeleteFavoriteInvoker(request *model.DeleteFavoriteRequest) *DeleteFavoriteInvoker {
	requestDef := GenReqDefForDeleteFavorite()
	return &DeleteFavoriteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteImage 删除镜像仓库
//
// 删除镜像仓库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteImage(request *model.DeleteImageRequest) (*model.DeleteImageResponse, error) {
	requestDef := GenReqDefForDeleteImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteImageResponse), nil
	}
}

// DeleteImageInvoker 删除镜像仓库
func (c *EiHealthClient) DeleteImageInvoker(request *model.DeleteImageRequest) *DeleteImageInvoker {
	requestDef := GenReqDefForDeleteImage()
	return &DeleteImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteJob 删除作业
//
// 删除作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteJob(request *model.DeleteJobRequest) (*model.DeleteJobResponse, error) {
	requestDef := GenReqDefForDeleteJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteJobResponse), nil
	}
}

// DeleteJobInvoker 删除作业
func (c *EiHealthClient) DeleteJobInvoker(request *model.DeleteJobRequest) *DeleteJobInvoker {
	requestDef := GenReqDefForDeleteJob()
	return &DeleteJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLabel 删除标签
//
// 删除标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteLabel(request *model.DeleteLabelRequest) (*model.DeleteLabelResponse, error) {
	requestDef := GenReqDefForDeleteLabel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLabelResponse), nil
	}
}

// DeleteLabelInvoker 删除标签
func (c *EiHealthClient) DeleteLabelInvoker(request *model.DeleteLabelRequest) *DeleteLabelInvoker {
	requestDef := GenReqDefForDeleteLabel()
	return &DeleteLabelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMember 移除项目成员
//
// 移除项目成员
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteMember(request *model.DeleteMemberRequest) (*model.DeleteMemberResponse, error) {
	requestDef := GenReqDefForDeleteMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMemberResponse), nil
	}
}

// DeleteMemberInvoker 移除项目成员
func (c *EiHealthClient) DeleteMemberInvoker(request *model.DeleteMemberRequest) *DeleteMemberInvoker {
	requestDef := GenReqDefForDeleteMember()
	return &DeleteMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePerformanceResource 删除性能加速资源
//
// 删除性能加速资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeletePerformanceResource(request *model.DeletePerformanceResourceRequest) (*model.DeletePerformanceResourceResponse, error) {
	requestDef := GenReqDefForDeletePerformanceResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePerformanceResourceResponse), nil
	}
}

// DeletePerformanceResourceInvoker 删除性能加速资源
func (c *EiHealthClient) DeletePerformanceResourceInvoker(request *model.DeletePerformanceResourceRequest) *DeletePerformanceResourceInvoker {
	requestDef := GenReqDefForDeletePerformanceResource()
	return &DeletePerformanceResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteProject 删除项目
//
// 删除项目
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteProject(request *model.DeleteProjectRequest) (*model.DeleteProjectResponse, error) {
	requestDef := GenReqDefForDeleteProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProjectResponse), nil
	}
}

// DeleteProjectInvoker 删除项目
func (c *EiHealthClient) DeleteProjectInvoker(request *model.DeleteProjectRequest) *DeleteProjectInvoker {
	requestDef := GenReqDefForDeleteProject()
	return &DeleteProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTag 删除指定镜像tag
//
// 删除指定镜像tag
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteTag(request *model.DeleteTagRequest) (*model.DeleteTagResponse, error) {
	requestDef := GenReqDefForDeleteTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTagResponse), nil
	}
}

// DeleteTagInvoker 删除指定镜像tag
func (c *EiHealthClient) DeleteTagInvoker(request *model.DeleteTagRequest) *DeleteTagInvoker {
	requestDef := GenReqDefForDeleteTag()
	return &DeleteTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteUser 删除用户
//
// 删除用户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteUser(request *model.DeleteUserRequest) (*model.DeleteUserResponse, error) {
	requestDef := GenReqDefForDeleteUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteUserResponse), nil
	}
}

// DeleteUserInvoker 删除用户
func (c *EiHealthClient) DeleteUserInvoker(request *model.DeleteUserRequest) *DeleteUserInvoker {
	requestDef := GenReqDefForDeleteUser()
	return &DeleteUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWorkflow 删除流程
//
// 删除流程
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteWorkflow(request *model.DeleteWorkflowRequest) (*model.DeleteWorkflowResponse, error) {
	requestDef := GenReqDefForDeleteWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWorkflowResponse), nil
	}
}

// DeleteWorkflowInvoker 删除流程
func (c *EiHealthClient) DeleteWorkflowInvoker(request *model.DeleteWorkflowRequest) *DeleteWorkflowInvoker {
	requestDef := GenReqDefForDeleteWorkflow()
	return &DeleteWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteJob 启动作业
//
// 启动作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ExecuteJob(request *model.ExecuteJobRequest) (*model.ExecuteJobResponse, error) {
	requestDef := GenReqDefForExecuteJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteJobResponse), nil
	}
}

// ExecuteJobInvoker 启动作业
func (c *EiHealthClient) ExecuteJobInvoker(request *model.ExecuteJobRequest) *ExecuteJobInvoker {
	requestDef := GenReqDefForExecuteJob()
	return &ExecuteJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GenerateComplexCombine 将传入的蛋白和小分子拼接成复合物结构
//
// 将传入的蛋白和小分子拼接成复合物结构
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) GenerateComplexCombine(request *model.GenerateComplexCombineRequest) (*model.GenerateComplexCombineResponse, error) {
	requestDef := GenReqDefForGenerateComplexCombine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GenerateComplexCombineResponse), nil
	}
}

// GenerateComplexCombineInvoker 将传入的蛋白和小分子拼接成复合物结构
func (c *EiHealthClient) GenerateComplexCombineInvoker(request *model.GenerateComplexCombineRequest) *GenerateComplexCombineInvoker {
	requestDef := GenReqDefForGenerateComplexCombine()
	return &GenerateComplexCombineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GeneratePocketFile 根据center、size、padding参数生成可渲染的口袋文件内容
//
// 根据center、size、padding参数生成可渲染的口袋文件内容
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) GeneratePocketFile(request *model.GeneratePocketFileRequest) (*model.GeneratePocketFileResponse, error) {
	requestDef := GenReqDefForGeneratePocketFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GeneratePocketFileResponse), nil
	}
}

// GeneratePocketFileInvoker 根据center、size、padding参数生成可渲染的口袋文件内容
func (c *EiHealthClient) GeneratePocketFileInvoker(request *model.GeneratePocketFileRequest) *GeneratePocketFileInvoker {
	requestDef := GenReqDefForGeneratePocketFile()
	return &GeneratePocketFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GenerateSurfacePoints 根据表面离散点坐标集生成可渲染的文件内容
//
// 根据表面离散点坐标集生成可渲染的文件内容
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) GenerateSurfacePoints(request *model.GenerateSurfacePointsRequest) (*model.GenerateSurfacePointsResponse, error) {
	requestDef := GenReqDefForGenerateSurfacePoints()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GenerateSurfacePointsResponse), nil
	}
}

// GenerateSurfacePointsInvoker 根据表面离散点坐标集生成可渲染的文件内容
func (c *EiHealthClient) GenerateSurfacePointsInvoker(request *model.GenerateSurfacePointsRequest) *GenerateSurfacePointsInvoker {
	requestDef := GenReqDefForGenerateSurfacePoints()
	return &GenerateSurfacePointsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportData 导入项目数据
//
// 导入项目数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ImportData(request *model.ImportDataRequest) (*model.ImportDataResponse, error) {
	requestDef := GenReqDefForImportData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportDataResponse), nil
	}
}

// ImportDataInvoker 导入项目数据
func (c *EiHealthClient) ImportDataInvoker(request *model.ImportDataRequest) *ImportDataInvoker {
	requestDef := GenReqDefForImportData()
	return &ImportDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportImage 导入镜像
//
// 导入镜像
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ImportImage(request *model.ImportImageRequest) (*model.ImportImageResponse, error) {
	requestDef := GenReqDefForImportImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportImageResponse), nil
	}
}

// ImportImageInvoker 导入镜像
func (c *EiHealthClient) ImportImageInvoker(request *model.ImportImageRequest) *ImportImageInvoker {
	requestDef := GenReqDefForImportImage()
	return &ImportImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportNetworkData 导入网上数据
//
// 导入网上数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ImportNetworkData(request *model.ImportNetworkDataRequest) (*model.ImportNetworkDataResponse, error) {
	requestDef := GenReqDefForImportNetworkData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportNetworkDataResponse), nil
	}
}

// ImportNetworkDataInvoker 导入网上数据
func (c *EiHealthClient) ImportNetworkDataInvoker(request *model.ImportNetworkDataRequest) *ImportNetworkDataInvoker {
	requestDef := GenReqDefForImportNetworkData()
	return &ImportNetworkDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportUser 导入用户
//
// 导入用户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ImportUser(request *model.ImportUserRequest) (*model.ImportUserResponse, error) {
	requestDef := GenReqDefForImportUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportUserResponse), nil
	}
}

// ImportUserInvoker 导入用户
func (c *EiHealthClient) ImportUserInvoker(request *model.ImportUserRequest) *ImportUserInvoker {
	requestDef := GenReqDefForImportUser()
	return &ImportUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportWorkflow 导入流程
//
// 导入流程
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ImportWorkflow(request *model.ImportWorkflowRequest) (*model.ImportWorkflowResponse, error) {
	requestDef := GenReqDefForImportWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportWorkflowResponse), nil
	}
}

// ImportWorkflowInvoker 导入流程
func (c *EiHealthClient) ImportWorkflowInvoker(request *model.ImportWorkflowRequest) *ImportWorkflowInvoker {
	requestDef := GenReqDefForImportWorkflow()
	return &ImportWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// InitializePlatform 初始化平台
//
// 初始化平台。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) InitializePlatform(request *model.InitializePlatformRequest) (*model.InitializePlatformResponse, error) {
	requestDef := GenReqDefForInitializePlatform()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.InitializePlatformResponse), nil
	}
}

// InitializePlatformInvoker 初始化平台
func (c *EiHealthClient) InitializePlatformInvoker(request *model.InitializePlatformRequest) *InitializePlatformInvoker {
	requestDef := GenReqDefForInitializePlatform()
	return &InitializePlatformInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApp 获取应用列表
//
// 获取应用列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListApp(request *model.ListAppRequest) (*model.ListAppResponse, error) {
	requestDef := GenReqDefForListApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppResponse), nil
	}
}

// ListAppInvoker 获取应用列表
func (c *EiHealthClient) ListAppInvoker(request *model.ListAppRequest) *ListAppInvoker {
	requestDef := GenReqDefForListApp()
	return &ListAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAsset 获取资产列表
//
// 获取资产列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListAsset(request *model.ListAssetRequest) (*model.ListAssetResponse, error) {
	requestDef := GenReqDefForListAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAssetResponse), nil
	}
}

// ListAssetInvoker 获取资产列表
func (c *EiHealthClient) ListAssetInvoker(request *model.ListAssetRequest) *ListAssetInvoker {
	requestDef := GenReqDefForListAsset()
	return &ListAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBaseModel 获取基模型列表
//
// 获取基模型列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListBaseModel(request *model.ListBaseModelRequest) (*model.ListBaseModelResponse, error) {
	requestDef := GenReqDefForListBaseModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBaseModelResponse), nil
	}
}

// ListBaseModelInvoker 获取基模型列表
func (c *EiHealthClient) ListBaseModelInvoker(request *model.ListBaseModelRequest) *ListBaseModelInvoker {
	requestDef := GenReqDefForListBaseModel()
	return &ListBaseModelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBucket 获取桶列表
//
// 获取桶列表(包含当前项目桶和引用项目桶)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListBucket(request *model.ListBucketRequest) (*model.ListBucketResponse, error) {
	requestDef := GenReqDefForListBucket()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBucketResponse), nil
	}
}

// ListBucketInvoker 获取桶列表
func (c *EiHealthClient) ListBucketInvoker(request *model.ListBucketRequest) *ListBucketInvoker {
	requestDef := GenReqDefForListBucket()
	return &ListBucketInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCceCluster 获取CCE集群列表
//
// 获取CCE集群列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListCceCluster(request *model.ListCceClusterRequest) (*model.ListCceClusterResponse, error) {
	requestDef := GenReqDefForListCceCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCceClusterResponse), nil
	}
}

// ListCceClusterInvoker 获取CCE集群列表
func (c *EiHealthClient) ListCceClusterInvoker(request *model.ListCceClusterRequest) *ListCceClusterInvoker {
	requestDef := GenReqDefForListCceCluster()
	return &ListCceClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusterAllNodeLabel 获取节点标签集
//
// 获取节点标签集
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListClusterAllNodeLabel(request *model.ListClusterAllNodeLabelRequest) (*model.ListClusterAllNodeLabelResponse, error) {
	requestDef := GenReqDefForListClusterAllNodeLabel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterAllNodeLabelResponse), nil
	}
}

// ListClusterAllNodeLabelInvoker 获取节点标签集
func (c *EiHealthClient) ListClusterAllNodeLabelInvoker(request *model.ListClusterAllNodeLabelRequest) *ListClusterAllNodeLabelInvoker {
	requestDef := GenReqDefForListClusterAllNodeLabel()
	return &ListClusterAllNodeLabelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusterInstallStep 查询指定集群安装步骤列表
//
// 查询指定集群安装步骤列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListClusterInstallStep(request *model.ListClusterInstallStepRequest) (*model.ListClusterInstallStepResponse, error) {
	requestDef := GenReqDefForListClusterInstallStep()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterInstallStepResponse), nil
	}
}

// ListClusterInstallStepInvoker 查询指定集群安装步骤列表
func (c *EiHealthClient) ListClusterInstallStepInvoker(request *model.ListClusterInstallStepRequest) *ListClusterInstallStepInvoker {
	requestDef := GenReqDefForListClusterInstallStep()
	return &ListClusterInstallStepInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListComputingCluster 获取计算集群列表
//
// 获取计算集群列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListComputingCluster(request *model.ListComputingClusterRequest) (*model.ListComputingClusterResponse, error) {
	requestDef := GenReqDefForListComputingCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListComputingClusterResponse), nil
	}
}

// ListComputingClusterInvoker 获取计算集群列表
func (c *EiHealthClient) ListComputingClusterInvoker(request *model.ListComputingClusterRequest) *ListComputingClusterInvoker {
	requestDef := GenReqDefForListComputingCluster()
	return &ListComputingClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListData 查询数据列表
//
// 查询指定目录下的数据列表，如果不指定默认查询根目录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListData(request *model.ListDataRequest) (*model.ListDataResponse, error) {
	requestDef := GenReqDefForListData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataResponse), nil
	}
}

// ListDataInvoker 查询数据列表
func (c *EiHealthClient) ListDataInvoker(request *model.ListDataRequest) *ListDataInvoker {
	requestDef := GenReqDefForListData()
	return &ListDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDataJob 获取数据作业列表
//
// 获取数据作业列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListDataJob(request *model.ListDataJobRequest) (*model.ListDataJobResponse, error) {
	requestDef := GenReqDefForListDataJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataJobResponse), nil
	}
}

// ListDataJobInvoker 获取数据作业列表
func (c *EiHealthClient) ListDataJobInvoker(request *model.ListDataJobRequest) *ListDataJobInvoker {
	requestDef := GenReqDefForListDataJob()
	return &ListDataJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDrugDatabase 获取数据库列表
//
// 获取数据库列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListDrugDatabase(request *model.ListDrugDatabaseRequest) (*model.ListDrugDatabaseResponse, error) {
	requestDef := GenReqDefForListDrugDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDrugDatabaseResponse), nil
	}
}

// ListDrugDatabaseInvoker 获取数据库列表
func (c *EiHealthClient) ListDrugDatabaseInvoker(request *model.ListDrugDatabaseRequest) *ListDrugDatabaseInvoker {
	requestDef := GenReqDefForListDrugDatabase()
	return &ListDrugDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDrugJob 获取药物作业列表
//
// 获取药物作业列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListDrugJob(request *model.ListDrugJobRequest) (*model.ListDrugJobResponse, error) {
	requestDef := GenReqDefForListDrugJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDrugJobResponse), nil
	}
}

// ListDrugJobInvoker 获取药物作业列表
func (c *EiHealthClient) ListDrugJobInvoker(request *model.ListDrugJobRequest) *ListDrugJobInvoker {
	requestDef := GenReqDefForListDrugJob()
	return &ListDrugJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDrugModel 获取模型列表
//
// 获取模型列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListDrugModel(request *model.ListDrugModelRequest) (*model.ListDrugModelResponse, error) {
	requestDef := GenReqDefForListDrugModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDrugModelResponse), nil
	}
}

// ListDrugModelInvoker 获取模型列表
func (c *EiHealthClient) ListDrugModelInvoker(request *model.ListDrugModelRequest) *ListDrugModelInvoker {
	requestDef := GenReqDefForListDrugModel()
	return &ListDrugModelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDrugModelResource 查询盘古药物分子大模型
//
// 查询盘古药物分子大模型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListDrugModelResource(request *model.ListDrugModelResourceRequest) (*model.ListDrugModelResourceResponse, error) {
	requestDef := GenReqDefForListDrugModelResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDrugModelResourceResponse), nil
	}
}

// ListDrugModelResourceInvoker 查询盘古药物分子大模型
func (c *EiHealthClient) ListDrugModelResourceInvoker(request *model.ListDrugModelResourceRequest) *ListDrugModelResourceInvoker {
	requestDef := GenReqDefForListDrugModelResource()
	return &ListDrugModelResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFavorite 获取收藏夹列表
//
// 获取收藏夹列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListFavorite(request *model.ListFavoriteRequest) (*model.ListFavoriteResponse, error) {
	requestDef := GenReqDefForListFavorite()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFavoriteResponse), nil
	}
}

// ListFavoriteInvoker 获取收藏夹列表
func (c *EiHealthClient) ListFavoriteInvoker(request *model.ListFavoriteRequest) *ListFavoriteInvoker {
	requestDef := GenReqDefForListFavorite()
	return &ListFavoriteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGlobalWorkflowStatistic 统计全局流程、作业信息
//
// 统计全局流程、作业信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListGlobalWorkflowStatistic(request *model.ListGlobalWorkflowStatisticRequest) (*model.ListGlobalWorkflowStatisticResponse, error) {
	requestDef := GenReqDefForListGlobalWorkflowStatistic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGlobalWorkflowStatisticResponse), nil
	}
}

// ListGlobalWorkflowStatisticInvoker 统计全局流程、作业信息
func (c *EiHealthClient) ListGlobalWorkflowStatisticInvoker(request *model.ListGlobalWorkflowStatisticRequest) *ListGlobalWorkflowStatisticInvoker {
	requestDef := GenReqDefForListGlobalWorkflowStatistic()
	return &ListGlobalWorkflowStatisticInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIamGroupUsers 查询IAM用户组的用户列表
//
// 查询IAM用户组的用户列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListIamGroupUsers(request *model.ListIamGroupUsersRequest) (*model.ListIamGroupUsersResponse, error) {
	requestDef := GenReqDefForListIamGroupUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIamGroupUsersResponse), nil
	}
}

// ListIamGroupUsersInvoker 查询IAM用户组的用户列表
func (c *EiHealthClient) ListIamGroupUsersInvoker(request *model.ListIamGroupUsersRequest) *ListIamGroupUsersInvoker {
	requestDef := GenReqDefForListIamGroupUsers()
	return &ListIamGroupUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIamGroups 查询IAM用户组列表
//
// 查询IAM用户组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListIamGroups(request *model.ListIamGroupsRequest) (*model.ListIamGroupsResponse, error) {
	requestDef := GenReqDefForListIamGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIamGroupsResponse), nil
	}
}

// ListIamGroupsInvoker 查询IAM用户组列表
func (c *EiHealthClient) ListIamGroupsInvoker(request *model.ListIamGroupsRequest) *ListIamGroupsInvoker {
	requestDef := GenReqDefForListIamGroups()
	return &ListIamGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIamUsers 查询IAM用户列表
//
// 查询IAM用户列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListIamUsers(request *model.ListIamUsersRequest) (*model.ListIamUsersResponse, error) {
	requestDef := GenReqDefForListIamUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIamUsersResponse), nil
	}
}

// ListIamUsersInvoker 查询IAM用户列表
func (c *EiHealthClient) ListIamUsersInvoker(request *model.ListIamUsersRequest) *ListIamUsersInvoker {
	requestDef := GenReqDefForListIamUsers()
	return &ListIamUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImage 获取镜像列表
//
// 获取镜像列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListImage(request *model.ListImageRequest) (*model.ListImageResponse, error) {
	requestDef := GenReqDefForListImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImageResponse), nil
	}
}

// ListImageInvoker 获取镜像列表
func (c *EiHealthClient) ListImageInvoker(request *model.ListImageRequest) *ListImageInvoker {
	requestDef := GenReqDefForListImage()
	return &ListImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImageTag 获取指定镜像的tag列表
//
// 获取指定镜像的tag列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListImageTag(request *model.ListImageTagRequest) (*model.ListImageTagResponse, error) {
	requestDef := GenReqDefForListImageTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImageTagResponse), nil
	}
}

// ListImageTagInvoker 获取指定镜像的tag列表
func (c *EiHealthClient) ListImageTagInvoker(request *model.ListImageTagRequest) *ListImageTagInvoker {
	requestDef := GenReqDefForListImageTag()
	return &ListImageTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJob 获取作业列表
//
// 获取作业列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListJob(request *model.ListJobRequest) (*model.ListJobResponse, error) {
	requestDef := GenReqDefForListJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobResponse), nil
	}
}

// ListJobInvoker 获取作业列表
func (c *EiHealthClient) ListJobInvoker(request *model.ListJobRequest) *ListJobInvoker {
	requestDef := GenReqDefForListJob()
	return &ListJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLabel 获取标签列表
//
// 获取标签列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListLabel(request *model.ListLabelRequest) (*model.ListLabelResponse, error) {
	requestDef := GenReqDefForListLabel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLabelResponse), nil
	}
}

// ListLabelInvoker 获取标签列表
func (c *EiHealthClient) ListLabelInvoker(request *model.ListLabelRequest) *ListLabelInvoker {
	requestDef := GenReqDefForListLabel()
	return &ListLabelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMfa 获取可用的认证方法
//
// 获取可用的认证方法
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListMfa(request *model.ListMfaRequest) (*model.ListMfaResponse, error) {
	requestDef := GenReqDefForListMfa()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMfaResponse), nil
	}
}

// ListMfaInvoker 获取可用的认证方法
func (c *EiHealthClient) ListMfaInvoker(request *model.ListMfaRequest) *ListMfaInvoker {
	requestDef := GenReqDefForListMfa()
	return &ListMfaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPerformanceResourceStat 获取性能加速资源上统计信息
//
// 获取性能加速资源上统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListPerformanceResourceStat(request *model.ListPerformanceResourceStatRequest) (*model.ListPerformanceResourceStatResponse, error) {
	requestDef := GenReqDefForListPerformanceResourceStat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPerformanceResourceStatResponse), nil
	}
}

// ListPerformanceResourceStatInvoker 获取性能加速资源上统计信息
func (c *EiHealthClient) ListPerformanceResourceStatInvoker(request *model.ListPerformanceResourceStatRequest) *ListPerformanceResourceStatInvoker {
	requestDef := GenReqDefForListPerformanceResourceStat()
	return &ListPerformanceResourceStatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPerformanceResources 查询性能加速资源
//
// 查询性能加速资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListPerformanceResources(request *model.ListPerformanceResourcesRequest) (*model.ListPerformanceResourcesResponse, error) {
	requestDef := GenReqDefForListPerformanceResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPerformanceResourcesResponse), nil
	}
}

// ListPerformanceResourcesInvoker 查询性能加速资源
func (c *EiHealthClient) ListPerformanceResourcesInvoker(request *model.ListPerformanceResourcesRequest) *ListPerformanceResourcesInvoker {
	requestDef := GenReqDefForListPerformanceResources()
	return &ListPerformanceResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProject 获取项目列表
//
// 获取项目列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListProject(request *model.ListProjectRequest) (*model.ListProjectResponse, error) {
	requestDef := GenReqDefForListProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectResponse), nil
	}
}

// ListProjectInvoker 获取项目列表
func (c *EiHealthClient) ListProjectInvoker(request *model.ListProjectRequest) *ListProjectInvoker {
	requestDef := GenReqDefForListProject()
	return &ListProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectStatistics 获取当前用户所属空间资源统计信息
//
// 获取当前用户所属空间资源统计信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListProjectStatistics(request *model.ListProjectStatisticsRequest) (*model.ListProjectStatisticsResponse, error) {
	requestDef := GenReqDefForListProjectStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectStatisticsResponse), nil
	}
}

// ListProjectStatisticsInvoker 获取当前用户所属空间资源统计信息
func (c *EiHealthClient) ListProjectStatisticsInvoker(request *model.ListProjectStatisticsRequest) *ListProjectStatisticsInvoker {
	requestDef := GenReqDefForListProjectStatistics()
	return &ListProjectStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProperty 获取属性值列表
//
// 获取属性值列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListProperty(request *model.ListPropertyRequest) (*model.ListPropertyResponse, error) {
	requestDef := GenReqDefForListProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPropertyResponse), nil
	}
}

// ListPropertyInvoker 获取属性值列表
func (c *EiHealthClient) ListPropertyInvoker(request *model.ListPropertyRequest) *ListPropertyInvoker {
	requestDef := GenReqDefForListProperty()
	return &ListPropertyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuota 获取当前系统配额及资源使用情况
//
// 获取当前系统配额及资源使用情况
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListQuota(request *model.ListQuotaRequest) (*model.ListQuotaResponse, error) {
	requestDef := GenReqDefForListQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotaResponse), nil
	}
}

// ListQuotaInvoker 获取当前系统配额及资源使用情况
func (c *EiHealthClient) ListQuotaInvoker(request *model.ListQuotaRequest) *ListQuotaInvoker {
	requestDef := GenReqDefForListQuota()
	return &ListQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSfsTurbos 获取sfs-turbo资源列表
//
// 获取sfs-turbo资源列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListSfsTurbos(request *model.ListSfsTurbosRequest) (*model.ListSfsTurbosResponse, error) {
	requestDef := GenReqDefForListSfsTurbos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSfsTurbosResponse), nil
	}
}

// ListSfsTurbosInvoker 获取sfs-turbo资源列表
func (c *EiHealthClient) ListSfsTurbosInvoker(request *model.ListSfsTurbosRequest) *ListSfsTurbosInvoker {
	requestDef := GenReqDefForListSfsTurbos()
	return &ListSfsTurbosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUser 获取用户列表
//
// 获取用户列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListUser(request *model.ListUserRequest) (*model.ListUserResponse, error) {
	requestDef := GenReqDefForListUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserResponse), nil
	}
}

// ListUserInvoker 获取用户列表
func (c *EiHealthClient) ListUserInvoker(request *model.ListUserRequest) *ListUserInvoker {
	requestDef := GenReqDefForListUser()
	return &ListUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserApp 获取用户所属空间的应用列表
//
// 获取用户所属空间的应用列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListUserApp(request *model.ListUserAppRequest) (*model.ListUserAppResponse, error) {
	requestDef := GenReqDefForListUserApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserAppResponse), nil
	}
}

// ListUserAppInvoker 获取用户所属空间的应用列表
func (c *EiHealthClient) ListUserAppInvoker(request *model.ListUserAppRequest) *ListUserAppInvoker {
	requestDef := GenReqDefForListUserApp()
	return &ListUserAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserDrugJob 获取用户所属空间的药物作业列表
//
// 获取用户所属空间的药物作业列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListUserDrugJob(request *model.ListUserDrugJobRequest) (*model.ListUserDrugJobResponse, error) {
	requestDef := GenReqDefForListUserDrugJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserDrugJobResponse), nil
	}
}

// ListUserDrugJobInvoker 获取用户所属空间的药物作业列表
func (c *EiHealthClient) ListUserDrugJobInvoker(request *model.ListUserDrugJobRequest) *ListUserDrugJobInvoker {
	requestDef := GenReqDefForListUserDrugJob()
	return &ListUserDrugJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserImage 获取用户所属空间的镜像列表
//
// 获取用户所属空间的镜像列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListUserImage(request *model.ListUserImageRequest) (*model.ListUserImageResponse, error) {
	requestDef := GenReqDefForListUserImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserImageResponse), nil
	}
}

// ListUserImageInvoker 获取用户所属空间的镜像列表
func (c *EiHealthClient) ListUserImageInvoker(request *model.ListUserImageRequest) *ListUserImageInvoker {
	requestDef := GenReqDefForListUserImage()
	return &ListUserImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserJob 获取用户所属空间的作业列表
//
// 获取用户所属空间的作业列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListUserJob(request *model.ListUserJobRequest) (*model.ListUserJobResponse, error) {
	requestDef := GenReqDefForListUserJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserJobResponse), nil
	}
}

// ListUserJobInvoker 获取用户所属空间的作业列表
func (c *EiHealthClient) ListUserJobInvoker(request *model.ListUserJobRequest) *ListUserJobInvoker {
	requestDef := GenReqDefForListUserJob()
	return &ListUserJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserWorkflow 获取用户所属空间的流程列表
//
// 获取用户所属空间的流程列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListUserWorkflow(request *model.ListUserWorkflowRequest) (*model.ListUserWorkflowResponse, error) {
	requestDef := GenReqDefForListUserWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserWorkflowResponse), nil
	}
}

// ListUserWorkflowInvoker 获取用户所属空间的流程列表
func (c *EiHealthClient) ListUserWorkflowInvoker(request *model.ListUserWorkflowRequest) *ListUserWorkflowInvoker {
	requestDef := GenReqDefForListUserWorkflow()
	return &ListUserWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVendor 获取供应商列表
//
// 获取供应商列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListVendor(request *model.ListVendorRequest) (*model.ListVendorResponse, error) {
	requestDef := GenReqDefForListVendor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVendorResponse), nil
	}
}

// ListVendorInvoker 获取供应商列表
func (c *EiHealthClient) ListVendorInvoker(request *model.ListVendorRequest) *ListVendorInvoker {
	requestDef := GenReqDefForListVendor()
	return &ListVendorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkflow 获取流程列表
//
// 获取流程列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListWorkflow(request *model.ListWorkflowRequest) (*model.ListWorkflowResponse, error) {
	requestDef := GenReqDefForListWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkflowResponse), nil
	}
}

// ListWorkflowInvoker 获取流程列表
func (c *EiHealthClient) ListWorkflowInvoker(request *model.ListWorkflowRequest) *ListWorkflowInvoker {
	requestDef := GenReqDefForListWorkflow()
	return &ListWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// QuoteData 引用项目数据
//
// 引用项目数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) QuoteData(request *model.QuoteDataRequest) (*model.QuoteDataResponse, error) {
	requestDef := GenReqDefForQuoteData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.QuoteDataResponse), nil
	}
}

// QuoteDataInvoker 引用项目数据
func (c *EiHealthClient) QuoteDataInvoker(request *model.QuoteDataRequest) *QuoteDataInvoker {
	requestDef := GenReqDefForQuoteData()
	return &QuoteDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RetryDataJob 重试数据作业
//
// 重试数据作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) RetryDataJob(request *model.RetryDataJobRequest) (*model.RetryDataJobResponse, error) {
	requestDef := GenReqDefForRetryDataJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RetryDataJobResponse), nil
	}
}

// RetryDataJobInvoker 重试数据作业
func (c *EiHealthClient) RetryDataJobInvoker(request *model.RetryDataJobRequest) *RetryDataJobInvoker {
	requestDef := GenReqDefForRetryDataJob()
	return &RetryDataJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RetryJob 重试作业
//
// 重试作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) RetryJob(request *model.RetryJobRequest) (*model.RetryJobResponse, error) {
	requestDef := GenReqDefForRetryJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RetryJobResponse), nil
	}
}

// RetryJobInvoker 重试作业
func (c *EiHealthClient) RetryJobInvoker(request *model.RetryJobRequest) *RetryJobInvoker {
	requestDef := GenReqDefForRetryJob()
	return &RetryJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunFastaPreprocess 受体预处理（Fasta格式）
//
// 受体预处理（Fasta格式），用于前端计算预期扣费次数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) RunFastaPreprocess(request *model.RunFastaPreprocessRequest) (*model.RunFastaPreprocessResponse, error) {
	requestDef := GenReqDefForRunFastaPreprocess()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunFastaPreprocessResponse), nil
	}
}

// RunFastaPreprocessInvoker 受体预处理（Fasta格式）
func (c *EiHealthClient) RunFastaPreprocessInvoker(request *model.RunFastaPreprocessRequest) *RunFastaPreprocessInvoker {
	requestDef := GenReqDefForRunFastaPreprocess()
	return &RunFastaPreprocessInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunFormatConverter 单分子文件格式转换
//
// 单分子文件格式转换。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) RunFormatConverter(request *model.RunFormatConverterRequest) (*model.RunFormatConverterResponse, error) {
	requestDef := GenReqDefForRunFormatConverter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunFormatConverterResponse), nil
	}
}

// RunFormatConverterInvoker 单分子文件格式转换
func (c *EiHealthClient) RunFormatConverterInvoker(request *model.RunFormatConverterRequest) *RunFormatConverterInvoker {
	requestDef := GenReqDefForRunFormatConverter()
	return &RunFormatConverterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAdmetJob 查询分子属性预测作业详情
//
// 查询分子属性预测作业详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowAdmetJob(request *model.ShowAdmetJobRequest) (*model.ShowAdmetJobResponse, error) {
	requestDef := GenReqDefForShowAdmetJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAdmetJobResponse), nil
	}
}

// ShowAdmetJobInvoker 查询分子属性预测作业详情
func (c *EiHealthClient) ShowAdmetJobInvoker(request *model.ShowAdmetJobRequest) *ShowAdmetJobInvoker {
	requestDef := GenReqDefForShowAdmetJob()
	return &ShowAdmetJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAgency 获取业务委托
//
// 获取业务委托。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowAgency(request *model.ShowAgencyRequest) (*model.ShowAgencyResponse, error) {
	requestDef := GenReqDefForShowAgency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAgencyResponse), nil
	}
}

// ShowAgencyInvoker 获取业务委托
func (c *EiHealthClient) ShowAgencyInvoker(request *model.ShowAgencyRequest) *ShowAgencyInvoker {
	requestDef := GenReqDefForShowAgency()
	return &ShowAgencyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApp 获取应用详情
//
// 获取应用详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowApp(request *model.ShowAppRequest) (*model.ShowAppResponse, error) {
	requestDef := GenReqDefForShowApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppResponse), nil
	}
}

// ShowAppInvoker 获取应用详情
func (c *EiHealthClient) ShowAppInvoker(request *model.ShowAppRequest) *ShowAppInvoker {
	requestDef := GenReqDefForShowApp()
	return &ShowAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAsset 查询资产详情
//
// 查询资产详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowAsset(request *model.ShowAssetRequest) (*model.ShowAssetResponse, error) {
	requestDef := GenReqDefForShowAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAssetResponse), nil
	}
}

// ShowAssetInvoker 查询资产详情
func (c *EiHealthClient) ShowAssetInvoker(request *model.ShowAssetRequest) *ShowAssetInvoker {
	requestDef := GenReqDefForShowAsset()
	return &ShowAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAssetVersion 查询资产版本详情
//
// 查询资产版本详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowAssetVersion(request *model.ShowAssetVersionRequest) (*model.ShowAssetVersionResponse, error) {
	requestDef := GenReqDefForShowAssetVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAssetVersionResponse), nil
	}
}

// ShowAssetVersionInvoker 查询资产版本详情
func (c *EiHealthClient) ShowAssetVersionInvoker(request *model.ShowAssetVersionRequest) *ShowAssetVersionInvoker {
	requestDef := GenReqDefForShowAssetVersion()
	return &ShowAssetVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBucketStorage 获取桶存量信息
//
// 获取桶存量信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowBucketStorage(request *model.ShowBucketStorageRequest) (*model.ShowBucketStorageResponse, error) {
	requestDef := GenReqDefForShowBucketStorage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBucketStorageResponse), nil
	}
}

// ShowBucketStorageInvoker 获取桶存量信息
func (c *EiHealthClient) ShowBucketStorageInvoker(request *model.ShowBucketStorageRequest) *ShowBucketStorageInvoker {
	requestDef := GenReqDefForShowBucketStorage()
	return &ShowBucketStorageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowClusteringJob 查询聚类分析作业详情
//
// 查询聚类分析作业详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowClusteringJob(request *model.ShowClusteringJobRequest) (*model.ShowClusteringJobResponse, error) {
	requestDef := GenReqDefForShowClusteringJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusteringJobResponse), nil
	}
}

// ShowClusteringJobInvoker 查询聚类分析作业详情
func (c *EiHealthClient) ShowClusteringJobInvoker(request *model.ShowClusteringJobRequest) *ShowClusteringJobInvoker {
	requestDef := GenReqDefForShowClusteringJob()
	return &ShowClusteringJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowData 获取数据详情
//
// 获取指定数据对象的详细信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowData(request *model.ShowDataRequest) (*model.ShowDataResponse, error) {
	requestDef := GenReqDefForShowData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDataResponse), nil
	}
}

// ShowDataInvoker 获取数据详情
func (c *EiHealthClient) ShowDataInvoker(request *model.ShowDataRequest) *ShowDataInvoker {
	requestDef := GenReqDefForShowData()
	return &ShowDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDataJob 获取数据作业详细信息
//
// 获取数据作业详细信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowDataJob(request *model.ShowDataJobRequest) (*model.ShowDataJobResponse, error) {
	requestDef := GenReqDefForShowDataJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDataJobResponse), nil
	}
}

// ShowDataJobInvoker 获取数据作业详细信息
func (c *EiHealthClient) ShowDataJobInvoker(request *model.ShowDataJobRequest) *ShowDataJobInvoker {
	requestDef := GenReqDefForShowDataJob()
	return &ShowDataJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDockerLogin 获取docker login指令
//
// 获取docker login指令
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowDockerLogin(request *model.ShowDockerLoginRequest) (*model.ShowDockerLoginResponse, error) {
	requestDef := GenReqDefForShowDockerLogin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDockerLoginResponse), nil
	}
}

// ShowDockerLoginInvoker 获取docker login指令
func (c *EiHealthClient) ShowDockerLoginInvoker(request *model.ShowDockerLoginRequest) *ShowDockerLoginInvoker {
	requestDef := GenReqDefForShowDockerLogin()
	return &ShowDockerLoginInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDockingJob 查询分子对接作业详情
//
// 查询分子对接作业详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowDockingJob(request *model.ShowDockingJobRequest) (*model.ShowDockingJobResponse, error) {
	requestDef := GenReqDefForShowDockingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDockingJobResponse), nil
	}
}

// ShowDockingJobInvoker 查询分子对接作业详情
func (c *EiHealthClient) ShowDockingJobInvoker(request *model.ShowDockingJobRequest) *ShowDockingJobInvoker {
	requestDef := GenReqDefForShowDockingJob()
	return &ShowDockingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEnv 查询系统配置列表
//
// 获取系统配置列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowEnv(request *model.ShowEnvRequest) (*model.ShowEnvResponse, error) {
	requestDef := GenReqDefForShowEnv()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEnvResponse), nil
	}
}

// ShowEnvInvoker 查询系统配置列表
func (c *EiHealthClient) ShowEnvInvoker(request *model.ShowEnvRequest) *ShowEnvInvoker {
	requestDef := GenReqDefForShowEnv()
	return &ShowEnvInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFepJob 查询自由能微扰作业详情
//
// 查询自由能微扰作业详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowFepJob(request *model.ShowFepJobRequest) (*model.ShowFepJobResponse, error) {
	requestDef := GenReqDefForShowFepJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFepJobResponse), nil
	}
}

// ShowFepJobInvoker 查询自由能微扰作业详情
func (c *EiHealthClient) ShowFepJobInvoker(request *model.ShowFepJobRequest) *ShowFepJobInvoker {
	requestDef := GenReqDefForShowFepJob()
	return &ShowFepJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGenJob 查询分子生成作业详情
//
// 查询分子生成作业详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowGenJob(request *model.ShowGenJobRequest) (*model.ShowGenJobResponse, error) {
	requestDef := GenReqDefForShowGenJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGenJobResponse), nil
	}
}

// ShowGenJobInvoker 查询分子生成作业详情
func (c *EiHealthClient) ShowGenJobInvoker(request *model.ShowGenJobRequest) *ShowGenJobInvoker {
	requestDef := GenReqDefForShowGenJob()
	return &ShowGenJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJob 获取作业详情
//
// 获取作业详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowJob(request *model.ShowJobRequest) (*model.ShowJobResponse, error) {
	requestDef := GenReqDefForShowJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobResponse), nil
	}
}

// ShowJobInvoker 获取作业详情
func (c *EiHealthClient) ShowJobInvoker(request *model.ShowJobRequest) *ShowJobInvoker {
	requestDef := GenReqDefForShowJob()
	return &ShowJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobEvent 获取作业事件
//
// 获取作业事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowJobEvent(request *model.ShowJobEventRequest) (*model.ShowJobEventResponse, error) {
	requestDef := GenReqDefForShowJobEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobEventResponse), nil
	}
}

// ShowJobEventInvoker 获取作业事件
func (c *EiHealthClient) ShowJobEventInvoker(request *model.ShowJobEventRequest) *ShowJobEventInvoker {
	requestDef := GenReqDefForShowJobEvent()
	return &ShowJobEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobLog 获取作业日志
//
// 获取作业日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowJobLog(request *model.ShowJobLogRequest) (*model.ShowJobLogResponse, error) {
	requestDef := GenReqDefForShowJobLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobLogResponse), nil
	}
}

// ShowJobLogInvoker 获取作业日志
func (c *EiHealthClient) ShowJobLogInvoker(request *model.ShowJobLogRequest) *ShowJobLogInvoker {
	requestDef := GenReqDefForShowJobLog()
	return &ShowJobLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMolBatchDownloadTask 查询分子或分子复合物批量下载任务详情
//
// 查询分子或分子复合物批量下载任务详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowMolBatchDownloadTask(request *model.ShowMolBatchDownloadTaskRequest) (*model.ShowMolBatchDownloadTaskResponse, error) {
	requestDef := GenReqDefForShowMolBatchDownloadTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMolBatchDownloadTaskResponse), nil
	}
}

// ShowMolBatchDownloadTaskInvoker 查询分子或分子复合物批量下载任务详情
func (c *EiHealthClient) ShowMolBatchDownloadTaskInvoker(request *model.ShowMolBatchDownloadTaskRequest) *ShowMolBatchDownloadTaskInvoker {
	requestDef := GenReqDefForShowMolBatchDownloadTask()
	return &ShowMolBatchDownloadTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOptmJob 查询分子优化作业详情
//
// 查询分子优化作业详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowOptmJob(request *model.ShowOptmJobRequest) (*model.ShowOptmJobResponse, error) {
	requestDef := GenReqDefForShowOptmJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOptmJobResponse), nil
	}
}

// ShowOptmJobInvoker 查询分子优化作业详情
func (c *EiHealthClient) ShowOptmJobInvoker(request *model.ShowOptmJobRequest) *ShowOptmJobInvoker {
	requestDef := GenReqDefForShowOptmJob()
	return &ShowOptmJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPocketDetectionJob 查询靶点口袋发现作业详情
//
// 查询靶点口袋发现作业详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowPocketDetectionJob(request *model.ShowPocketDetectionJobRequest) (*model.ShowPocketDetectionJobResponse, error) {
	requestDef := GenReqDefForShowPocketDetectionJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPocketDetectionJobResponse), nil
	}
}

// ShowPocketDetectionJobInvoker 查询靶点口袋发现作业详情
func (c *EiHealthClient) ShowPocketDetectionJobInvoker(request *model.ShowPocketDetectionJobRequest) *ShowPocketDetectionJobInvoker {
	requestDef := GenReqDefForShowPocketDetectionJob()
	return &ShowPocketDetectionJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPocketMolDesignJob 查询靶点口袋分子设计作业详情
//
// 查询靶点口袋分子设计作业详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowPocketMolDesignJob(request *model.ShowPocketMolDesignJobRequest) (*model.ShowPocketMolDesignJobResponse, error) {
	requestDef := GenReqDefForShowPocketMolDesignJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPocketMolDesignJobResponse), nil
	}
}

// ShowPocketMolDesignJobInvoker 查询靶点口袋分子设计作业详情
func (c *EiHealthClient) ShowPocketMolDesignJobInvoker(request *model.ShowPocketMolDesignJobRequest) *ShowPocketMolDesignJobInvoker {
	requestDef := GenReqDefForShowPocketMolDesignJob()
	return &ShowPocketMolDesignJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProject 获取项目详情
//
// 获取项目详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowProject(request *model.ShowProjectRequest) (*model.ShowProjectResponse, error) {
	requestDef := GenReqDefForShowProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectResponse), nil
	}
}

// ShowProjectInvoker 获取项目详情
func (c *EiHealthClient) ShowProjectInvoker(request *model.ShowProjectRequest) *ShowProjectInvoker {
	requestDef := GenReqDefForShowProject()
	return &ShowProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSearchJob 查询分子搜索作业详情
//
// 查询分子搜索作业详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowSearchJob(request *model.ShowSearchJobRequest) (*model.ShowSearchJobResponse, error) {
	requestDef := GenReqDefForShowSearchJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSearchJobResponse), nil
	}
}

// ShowSearchJobInvoker 查询分子搜索作业详情
func (c *EiHealthClient) ShowSearchJobInvoker(request *model.ShowSearchJobRequest) *ShowSearchJobInvoker {
	requestDef := GenReqDefForShowSearchJob()
	return &ShowSearchJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTargetOptJob 查询靶点优化作业详情
//
// 查询靶点优化作业详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowTargetOptJob(request *model.ShowTargetOptJobRequest) (*model.ShowTargetOptJobResponse, error) {
	requestDef := GenReqDefForShowTargetOptJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTargetOptJobResponse), nil
	}
}

// ShowTargetOptJobInvoker 查询靶点优化作业详情
func (c *EiHealthClient) ShowTargetOptJobInvoker(request *model.ShowTargetOptJobRequest) *ShowTargetOptJobInvoker {
	requestDef := GenReqDefForShowTargetOptJob()
	return &ShowTargetOptJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTaskEvents 获取子任务启动事件
//
// 获取子任务启动事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowTaskEvents(request *model.ShowTaskEventsRequest) (*model.ShowTaskEventsResponse, error) {
	requestDef := GenReqDefForShowTaskEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskEventsResponse), nil
	}
}

// ShowTaskEventsInvoker 获取子任务启动事件
func (c *EiHealthClient) ShowTaskEventsInvoker(request *model.ShowTaskEventsRequest) *ShowTaskEventsInvoker {
	requestDef := GenReqDefForShowTaskEvents()
	return &ShowTaskEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTaskInstanceEvents 获取子任务中实例的事件
//
// 获取子任务中实例的事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowTaskInstanceEvents(request *model.ShowTaskInstanceEventsRequest) (*model.ShowTaskInstanceEventsResponse, error) {
	requestDef := GenReqDefForShowTaskInstanceEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskInstanceEventsResponse), nil
	}
}

// ShowTaskInstanceEventsInvoker 获取子任务中实例的事件
func (c *EiHealthClient) ShowTaskInstanceEventsInvoker(request *model.ShowTaskInstanceEventsRequest) *ShowTaskInstanceEventsInvoker {
	requestDef := GenReqDefForShowTaskInstanceEvents()
	return &ShowTaskInstanceEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTaskInstanceMetricData 获取子任务中实例的资源监控数据
//
// 获取子任务中实例的资源监控数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowTaskInstanceMetricData(request *model.ShowTaskInstanceMetricDataRequest) (*model.ShowTaskInstanceMetricDataResponse, error) {
	requestDef := GenReqDefForShowTaskInstanceMetricData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskInstanceMetricDataResponse), nil
	}
}

// ShowTaskInstanceMetricDataInvoker 获取子任务中实例的资源监控数据
func (c *EiHealthClient) ShowTaskInstanceMetricDataInvoker(request *model.ShowTaskInstanceMetricDataRequest) *ShowTaskInstanceMetricDataInvoker {
	requestDef := GenReqDefForShowTaskInstanceMetricData()
	return &ShowTaskInstanceMetricDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTaskInstancePod 获取子任务中实例的pod信息
//
// 获取子任务中实例的pod信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowTaskInstancePod(request *model.ShowTaskInstancePodRequest) (*model.ShowTaskInstancePodResponse, error) {
	requestDef := GenReqDefForShowTaskInstancePod()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskInstancePodResponse), nil
	}
}

// ShowTaskInstancePodInvoker 获取子任务中实例的pod信息
func (c *EiHealthClient) ShowTaskInstancePodInvoker(request *model.ShowTaskInstancePodRequest) *ShowTaskInstancePodInvoker {
	requestDef := GenReqDefForShowTaskInstancePod()
	return &ShowTaskInstancePodInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTaskInstances 获取子任务实例信息
//
// 获取子任务实例信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowTaskInstances(request *model.ShowTaskInstancesRequest) (*model.ShowTaskInstancesResponse, error) {
	requestDef := GenReqDefForShowTaskInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskInstancesResponse), nil
	}
}

// ShowTaskInstancesInvoker 获取子任务实例信息
func (c *EiHealthClient) ShowTaskInstancesInvoker(request *model.ShowTaskInstancesRequest) *ShowTaskInstancesInvoker {
	requestDef := GenReqDefForShowTaskInstances()
	return &ShowTaskInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUser 获取指定用户详情
//
// 获取指定用户详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowUser(request *model.ShowUserRequest) (*model.ShowUserResponse, error) {
	requestDef := GenReqDefForShowUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserResponse), nil
	}
}

// ShowUserInvoker 获取指定用户详情
func (c *EiHealthClient) ShowUserInvoker(request *model.ShowUserRequest) *ShowUserInvoker {
	requestDef := GenReqDefForShowUser()
	return &ShowUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUserSetting 查询用户设置
//
// 查询用户设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowUserSetting(request *model.ShowUserSettingRequest) (*model.ShowUserSettingResponse, error) {
	requestDef := GenReqDefForShowUserSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserSettingResponse), nil
	}
}

// ShowUserSettingInvoker 查询用户设置
func (c *EiHealthClient) ShowUserSettingInvoker(request *model.ShowUserSettingRequest) *ShowUserSettingInvoker {
	requestDef := GenReqDefForShowUserSetting()
	return &ShowUserSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVendor 获取供应商配置
//
// 获取供应商配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowVendor(request *model.ShowVendorRequest) (*model.ShowVendorResponse, error) {
	requestDef := GenReqDefForShowVendor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVendorResponse), nil
	}
}

// ShowVendorInvoker 获取供应商配置
func (c *EiHealthClient) ShowVendorInvoker(request *model.ShowVendorRequest) *ShowVendorInvoker {
	requestDef := GenReqDefForShowVendor()
	return &ShowVendorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkflow 获取流程详情
//
// 获取流程详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowWorkflow(request *model.ShowWorkflowRequest) (*model.ShowWorkflowResponse, error) {
	requestDef := GenReqDefForShowWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkflowResponse), nil
	}
}

// ShowWorkflowInvoker 获取流程详情
func (c *EiHealthClient) ShowWorkflowInvoker(request *model.ShowWorkflowRequest) *ShowWorkflowInvoker {
	requestDef := GenReqDefForShowWorkflow()
	return &ShowWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SubscribeApp 订阅应用
//
// 订阅应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) SubscribeApp(request *model.SubscribeAppRequest) (*model.SubscribeAppResponse, error) {
	requestDef := GenReqDefForSubscribeApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SubscribeAppResponse), nil
	}
}

// SubscribeAppInvoker 订阅应用
func (c *EiHealthClient) SubscribeAppInvoker(request *model.SubscribeAppRequest) *SubscribeAppInvoker {
	requestDef := GenReqDefForSubscribeApp()
	return &SubscribeAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SubscribeImage 订阅镜像
//
// 订阅镜像
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) SubscribeImage(request *model.SubscribeImageRequest) (*model.SubscribeImageResponse, error) {
	requestDef := GenReqDefForSubscribeImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SubscribeImageResponse), nil
	}
}

// SubscribeImageInvoker 订阅镜像
func (c *EiHealthClient) SubscribeImageInvoker(request *model.SubscribeImageRequest) *SubscribeImageInvoker {
	requestDef := GenReqDefForSubscribeImage()
	return &SubscribeImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SubscribeWorkflow 订阅流程
//
// 订阅流程
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) SubscribeWorkflow(request *model.SubscribeWorkflowRequest) (*model.SubscribeWorkflowResponse, error) {
	requestDef := GenReqDefForSubscribeWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SubscribeWorkflowResponse), nil
	}
}

// SubscribeWorkflowInvoker 订阅流程
func (c *EiHealthClient) SubscribeWorkflowInvoker(request *model.SubscribeWorkflowRequest) *SubscribeWorkflowInvoker {
	requestDef := GenReqDefForSubscribeWorkflow()
	return &SubscribeWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// TransferProject 转移项目
//
// 转移项目
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) TransferProject(request *model.TransferProjectRequest) (*model.TransferProjectResponse, error) {
	requestDef := GenReqDefForTransferProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.TransferProjectResponse), nil
	}
}

// TransferProjectInvoker 转移项目
func (c *EiHealthClient) TransferProjectInvoker(request *model.TransferProjectRequest) *TransferProjectInvoker {
	requestDef := GenReqDefForTransferProject()
	return &TransferProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAgency 更新业务委托
//
// 更新业务委托。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateAgency(request *model.UpdateAgencyRequest) (*model.UpdateAgencyResponse, error) {
	requestDef := GenReqDefForUpdateAgency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAgencyResponse), nil
	}
}

// UpdateAgencyInvoker 更新业务委托
func (c *EiHealthClient) UpdateAgencyInvoker(request *model.UpdateAgencyRequest) *UpdateAgencyInvoker {
	requestDef := GenReqDefForUpdateAgency()
	return &UpdateAgencyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateApp 更新应用
//
// 更新应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateApp(request *model.UpdateAppRequest) (*model.UpdateAppResponse, error) {
	requestDef := GenReqDefForUpdateApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppResponse), nil
	}
}

// UpdateAppInvoker 更新应用
func (c *EiHealthClient) UpdateAppInvoker(request *model.UpdateAppRequest) *UpdateAppInvoker {
	requestDef := GenReqDefForUpdateApp()
	return &UpdateAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDrugDatabase 更新药物数据库
//
// 更新药物数据库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateDrugDatabase(request *model.UpdateDrugDatabaseRequest) (*model.UpdateDrugDatabaseResponse, error) {
	requestDef := GenReqDefForUpdateDrugDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDrugDatabaseResponse), nil
	}
}

// UpdateDrugDatabaseInvoker 更新药物数据库
func (c *EiHealthClient) UpdateDrugDatabaseInvoker(request *model.UpdateDrugDatabaseRequest) *UpdateDrugDatabaseInvoker {
	requestDef := GenReqDefForUpdateDrugDatabase()
	return &UpdateDrugDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDrugJob 更新药物作业
//
// 更新药物作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateDrugJob(request *model.UpdateDrugJobRequest) (*model.UpdateDrugJobResponse, error) {
	requestDef := GenReqDefForUpdateDrugJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDrugJobResponse), nil
	}
}

// UpdateDrugJobInvoker 更新药物作业
func (c *EiHealthClient) UpdateDrugJobInvoker(request *model.UpdateDrugJobRequest) *UpdateDrugJobInvoker {
	requestDef := GenReqDefForUpdateDrugJob()
	return &UpdateDrugJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDrugModel 更新药物模型
//
// 更新药物模型
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateDrugModel(request *model.UpdateDrugModelRequest) (*model.UpdateDrugModelResponse, error) {
	requestDef := GenReqDefForUpdateDrugModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDrugModelResponse), nil
	}
}

// UpdateDrugModelInvoker 更新药物模型
func (c *EiHealthClient) UpdateDrugModelInvoker(request *model.UpdateDrugModelRequest) *UpdateDrugModelInvoker {
	requestDef := GenReqDefForUpdateDrugModel()
	return &UpdateDrugModelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateImage 更新镜像描述信息或者类型
//
// 更新镜像描述信息或者类型
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateImage(request *model.UpdateImageRequest) (*model.UpdateImageResponse, error) {
	requestDef := GenReqDefForUpdateImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateImageResponse), nil
	}
}

// UpdateImageInvoker 更新镜像描述信息或者类型
func (c *EiHealthClient) UpdateImageInvoker(request *model.UpdateImageRequest) *UpdateImageInvoker {
	requestDef := GenReqDefForUpdateImage()
	return &UpdateImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInitPassword 新用户重置密码
//
// 新用户重置密码
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateInitPassword(request *model.UpdateInitPasswordRequest) (*model.UpdateInitPasswordResponse, error) {
	requestDef := GenReqDefForUpdateInitPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInitPasswordResponse), nil
	}
}

// UpdateInitPasswordInvoker 新用户重置密码
func (c *EiHealthClient) UpdateInitPasswordInvoker(request *model.UpdateInitPasswordRequest) *UpdateInitPasswordInvoker {
	requestDef := GenReqDefForUpdateInitPassword()
	return &UpdateInitPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateJob 更新作业
//
// 更新作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateJob(request *model.UpdateJobRequest) (*model.UpdateJobResponse, error) {
	requestDef := GenReqDefForUpdateJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateJobResponse), nil
	}
}

// UpdateJobInvoker 更新作业
func (c *EiHealthClient) UpdateJobInvoker(request *model.UpdateJobRequest) *UpdateJobInvoker {
	requestDef := GenReqDefForUpdateJob()
	return &UpdateJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMember 更新或者添加项目成员角色
//
// 更新或者添加项目成员角色
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateMember(request *model.UpdateMemberRequest) (*model.UpdateMemberResponse, error) {
	requestDef := GenReqDefForUpdateMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMemberResponse), nil
	}
}

// UpdateMemberInvoker 更新或者添加项目成员角色
func (c *EiHealthClient) UpdateMemberInvoker(request *model.UpdateMemberRequest) *UpdateMemberInvoker {
	requestDef := GenReqDefForUpdateMember()
	return &UpdateMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePerformanceResource 更新性能加速资源配置
//
// 更新性能加速资源配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdatePerformanceResource(request *model.UpdatePerformanceResourceRequest) (*model.UpdatePerformanceResourceResponse, error) {
	requestDef := GenReqDefForUpdatePerformanceResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePerformanceResourceResponse), nil
	}
}

// UpdatePerformanceResourceInvoker 更新性能加速资源配置
func (c *EiHealthClient) UpdatePerformanceResourceInvoker(request *model.UpdatePerformanceResourceRequest) *UpdatePerformanceResourceInvoker {
	requestDef := GenReqDefForUpdatePerformanceResource()
	return &UpdatePerformanceResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProject 更新项目
//
// 更新项目
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateProject(request *model.UpdateProjectRequest) (*model.UpdateProjectResponse, error) {
	requestDef := GenReqDefForUpdateProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProjectResponse), nil
	}
}

// UpdateProjectInvoker 更新项目
func (c *EiHealthClient) UpdateProjectInvoker(request *model.UpdateProjectRequest) *UpdateProjectInvoker {
	requestDef := GenReqDefForUpdateProject()
	return &UpdateProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTopProject 置顶空间
//
// 置顶空间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateTopProject(request *model.UpdateTopProjectRequest) (*model.UpdateTopProjectResponse, error) {
	requestDef := GenReqDefForUpdateTopProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTopProjectResponse), nil
	}
}

// UpdateTopProjectInvoker 置顶空间
func (c *EiHealthClient) UpdateTopProjectInvoker(request *model.UpdateTopProjectRequest) *UpdateTopProjectInvoker {
	requestDef := GenReqDefForUpdateTopProject()
	return &UpdateTopProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUser 修改用户基本信息
//
// 修改用户基本信息（邮箱，手机）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateUser(request *model.UpdateUserRequest) (*model.UpdateUserResponse, error) {
	requestDef := GenReqDefForUpdateUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserResponse), nil
	}
}

// UpdateUserInvoker 修改用户基本信息
func (c *EiHealthClient) UpdateUserInvoker(request *model.UpdateUserRequest) *UpdateUserInvoker {
	requestDef := GenReqDefForUpdateUser()
	return &UpdateUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUserByDomain 最终租户修改子用户
//
// 最终租户修改子用户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateUserByDomain(request *model.UpdateUserByDomainRequest) (*model.UpdateUserByDomainResponse, error) {
	requestDef := GenReqDefForUpdateUserByDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserByDomainResponse), nil
	}
}

// UpdateUserByDomainInvoker 最终租户修改子用户
func (c *EiHealthClient) UpdateUserByDomainInvoker(request *model.UpdateUserByDomainRequest) *UpdateUserByDomainInvoker {
	requestDef := GenReqDefForUpdateUserByDomain()
	return &UpdateUserByDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUserRole 更新用户角色
//
// 更新用户角色
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateUserRole(request *model.UpdateUserRoleRequest) (*model.UpdateUserRoleResponse, error) {
	requestDef := GenReqDefForUpdateUserRole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserRoleResponse), nil
	}
}

// UpdateUserRoleInvoker 更新用户角色
func (c *EiHealthClient) UpdateUserRoleInvoker(request *model.UpdateUserRoleRequest) *UpdateUserRoleInvoker {
	requestDef := GenReqDefForUpdateUserRole()
	return &UpdateUserRoleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUserSetting 更新用户设置
//
// 更新用户设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateUserSetting(request *model.UpdateUserSettingRequest) (*model.UpdateUserSettingResponse, error) {
	requestDef := GenReqDefForUpdateUserSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserSettingResponse), nil
	}
}

// UpdateUserSettingInvoker 更新用户设置
func (c *EiHealthClient) UpdateUserSettingInvoker(request *model.UpdateUserSettingRequest) *UpdateUserSettingInvoker {
	requestDef := GenReqDefForUpdateUserSetting()
	return &UpdateUserSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVendor 设置供应商配置
//
// 设置供应商配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateVendor(request *model.UpdateVendorRequest) (*model.UpdateVendorResponse, error) {
	requestDef := GenReqDefForUpdateVendor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVendorResponse), nil
	}
}

// UpdateVendorInvoker 设置供应商配置
func (c *EiHealthClient) UpdateVendorInvoker(request *model.UpdateVendorRequest) *UpdateVendorInvoker {
	requestDef := GenReqDefForUpdateVendor()
	return &UpdateVendorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWorkflow 更新流程
//
// 更新流程
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateWorkflow(request *model.UpdateWorkflowRequest) (*model.UpdateWorkflowResponse, error) {
	requestDef := GenReqDefForUpdateWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWorkflowResponse), nil
	}
}

// UpdateWorkflowInvoker 更新流程
func (c *EiHealthClient) UpdateWorkflowInvoker(request *model.UpdateWorkflowRequest) *UpdateWorkflowInvoker {
	requestDef := GenReqDefForUpdateWorkflow()
	return &UpdateWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadData 上传数据文件
//
// 上传数据文件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UploadData(request *model.UploadDataRequest) (*model.UploadDataResponse, error) {
	requestDef := GenReqDefForUploadData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadDataResponse), nil
	}
}

// UploadDataInvoker 上传数据文件
func (c *EiHealthClient) UploadDataInvoker(request *model.UploadDataRequest) *UploadDataInvoker {
	requestDef := GenReqDefForUploadData()
	return &UploadDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ValidateCode 预验证
//
// 预验证
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ValidateCode(request *model.ValidateCodeRequest) (*model.ValidateCodeResponse, error) {
	requestDef := GenReqDefForValidateCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateCodeResponse), nil
	}
}

// ValidateCodeInvoker 预验证
func (c *EiHealthClient) ValidateCodeInvoker(request *model.ValidateCodeRequest) *ValidateCodeInvoker {
	requestDef := GenReqDefForValidateCode()
	return &ValidateCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAdmetProperties ADMET属性预测接口
//
// 计算小分子的物化性质，包括吸收(adsorption)、分布(distribution)、代谢(metabolism)、清除(excretion)与毒性(toxicity)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowAdmetProperties(request *model.ShowAdmetPropertiesRequest) (*model.ShowAdmetPropertiesResponse, error) {
	requestDef := GenReqDefForShowAdmetProperties()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAdmetPropertiesResponse), nil
	}
}

// ShowAdmetPropertiesInvoker ADMET属性预测接口
func (c *EiHealthClient) ShowAdmetPropertiesInvoker(request *model.ShowAdmetPropertiesRequest) *ShowAdmetPropertiesInvoker {
	requestDef := GenReqDefForShowAdmetProperties()
	return &ShowAdmetPropertiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCpiJob 创建CPI作业
//
// 创建CPI作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateCpiJob(request *model.CreateCpiJobRequest) (*model.CreateCpiJobResponse, error) {
	requestDef := GenReqDefForCreateCpiJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCpiJobResponse), nil
	}
}

// CreateCpiJobInvoker 创建CPI作业
func (c *EiHealthClient) CreateCpiJobInvoker(request *model.CreateCpiJobRequest) *CreateCpiJobInvoker {
	requestDef := GenReqDefForCreateCpiJob()
	return &CreateCpiJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCpiJob 查询CPI作业详情
//
// 查询CPI作业详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowCpiJob(request *model.ShowCpiJobRequest) (*model.ShowCpiJobResponse, error) {
	requestDef := GenReqDefForShowCpiJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCpiJobResponse), nil
	}
}

// ShowCpiJobInvoker 查询CPI作业详情
func (c *EiHealthClient) ShowCpiJobInvoker(request *model.ShowCpiJobRequest) *ShowCpiJobInvoker {
	requestDef := GenReqDefForShowCpiJob()
	return &ShowCpiJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCssCluster 绑定CSS集群
//
// 绑定CSS集群
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateCssCluster(request *model.CreateCssClusterRequest) (*model.CreateCssClusterResponse, error) {
	requestDef := GenReqDefForCreateCssCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCssClusterResponse), nil
	}
}

// CreateCssClusterInvoker 绑定CSS集群
func (c *EiHealthClient) CreateCssClusterInvoker(request *model.CreateCssClusterRequest) *CreateCssClusterInvoker {
	requestDef := GenReqDefForCreateCssCluster()
	return &CreateCssClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCssCluster CSS集群解绑
//
// CSS集群解绑
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteCssCluster(request *model.DeleteCssClusterRequest) (*model.DeleteCssClusterResponse, error) {
	requestDef := GenReqDefForDeleteCssCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCssClusterResponse), nil
	}
}

// DeleteCssClusterInvoker CSS集群解绑
func (c *EiHealthClient) DeleteCssClusterInvoker(request *model.DeleteCssClusterRequest) *DeleteCssClusterInvoker {
	requestDef := GenReqDefForDeleteCssCluster()
	return &DeleteCssClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCssCluster 获取CSS集群列表
//
// 获取CSS集群列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListCssCluster(request *model.ListCssClusterRequest) (*model.ListCssClusterResponse, error) {
	requestDef := GenReqDefForListCssCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCssClusterResponse), nil
	}
}

// ListCssClusterInvoker 获取CSS集群列表
func (c *EiHealthClient) ListCssClusterInvoker(request *model.ListCssClusterRequest) *ListCssClusterInvoker {
	requestDef := GenReqDefForListCssCluster()
	return &ListCssClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTermTenantCssCluster 获取最终租户CSS集群列表
//
// 获取最终租户CSS集群列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListTermTenantCssCluster(request *model.ListTermTenantCssClusterRequest) (*model.ListTermTenantCssClusterResponse, error) {
	requestDef := GenReqDefForListTermTenantCssCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTermTenantCssClusterResponse), nil
	}
}

// ListTermTenantCssClusterInvoker 获取最终租户CSS集群列表
func (c *EiHealthClient) ListTermTenantCssClusterInvoker(request *model.ListTermTenantCssClusterRequest) *ListTermTenantCssClusterInvoker {
	requestDef := GenReqDefForListTermTenantCssCluster()
	return &ListTermTenantCssClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ValidateCssConnection 测试CSS集群连接
//
// 测试CSS集群连接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ValidateCssConnection(request *model.ValidateCssConnectionRequest) (*model.ValidateCssConnectionResponse, error) {
	requestDef := GenReqDefForValidateCssConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateCssConnectionResponse), nil
	}
}

// ValidateCssConnectionInvoker 测试CSS集群连接
func (c *EiHealthClient) ValidateCssConnectionInvoker(request *model.ValidateCssConnectionRequest) *ValidateCssConnectionInvoker {
	requestDef := GenReqDefForValidateCssConnection()
	return &ValidateCssConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckDrugLigandDifference 计算配体间的3D结构差异
//
// 计算配体间的3D结构差异
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CheckDrugLigandDifference(request *model.CheckDrugLigandDifferenceRequest) (*model.CheckDrugLigandDifferenceResponse, error) {
	requestDef := GenReqDefForCheckDrugLigandDifference()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckDrugLigandDifferenceResponse), nil
	}
}

// CheckDrugLigandDifferenceInvoker 计算配体间的3D结构差异
func (c *EiHealthClient) CheckDrugLigandDifferenceInvoker(request *model.CheckDrugLigandDifferenceRequest) *CheckDrugLigandDifferenceInvoker {
	requestDef := GenReqDefForCheckDrugLigandDifference()
	return &CheckDrugLigandDifferenceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDrugLigandInteraction2dSvg 生成相互作用2D图
//
// 生成相互作用2D图，若不提供配体文件，则受体文件中必须包含配体；若提供配体文件，则受体中的配体（若有）则会被忽略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateDrugLigandInteraction2dSvg(request *model.CreateDrugLigandInteraction2dSvgRequest) (*model.CreateDrugLigandInteraction2dSvgResponse, error) {
	requestDef := GenReqDefForCreateDrugLigandInteraction2dSvg()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDrugLigandInteraction2dSvgResponse), nil
	}
}

// CreateDrugLigandInteraction2dSvgInvoker 生成相互作用2D图
func (c *EiHealthClient) CreateDrugLigandInteraction2dSvgInvoker(request *model.CreateDrugLigandInteraction2dSvgRequest) *CreateDrugLigandInteraction2dSvgInvoker {
	requestDef := GenReqDefForCreateDrugLigandInteraction2dSvg()
	return &CreateDrugLigandInteraction2dSvgInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDrugLigandPreviewTask 创建配体文件预览任务
//
// 创建配体文件预览任务，支持SMI、SDF、PDB、MOL2
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateDrugLigandPreviewTask(request *model.CreateDrugLigandPreviewTaskRequest) (*model.CreateDrugLigandPreviewTaskResponse, error) {
	requestDef := GenReqDefForCreateDrugLigandPreviewTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDrugLigandPreviewTaskResponse), nil
	}
}

// CreateDrugLigandPreviewTaskInvoker 创建配体文件预览任务
func (c *EiHealthClient) CreateDrugLigandPreviewTaskInvoker(request *model.CreateDrugLigandPreviewTaskRequest) *CreateDrugLigandPreviewTaskInvoker {
	requestDef := GenReqDefForCreateDrugLigandPreviewTask()
	return &CreateDrugLigandPreviewTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDrugLigandSdf 生成分子SDF三维结构
//
// 生成分子SDF三维结构
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateDrugLigandSdf(request *model.CreateDrugLigandSdfRequest) (*model.CreateDrugLigandSdfResponse, error) {
	requestDef := GenReqDefForCreateDrugLigandSdf()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDrugLigandSdfResponse), nil
	}
}

// CreateDrugLigandSdfInvoker 生成分子SDF三维结构
func (c *EiHealthClient) CreateDrugLigandSdfInvoker(request *model.CreateDrugLigandSdfRequest) *CreateDrugLigandSdfInvoker {
	requestDef := GenReqDefForCreateDrugLigandSdf()
	return &CreateDrugLigandSdfInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDrugLigandSimilarityGraphTask 创建配体相似性图计算任务
//
// 创建配体相似性图计算任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateDrugLigandSimilarityGraphTask(request *model.CreateDrugLigandSimilarityGraphTaskRequest) (*model.CreateDrugLigandSimilarityGraphTaskResponse, error) {
	requestDef := GenReqDefForCreateDrugLigandSimilarityGraphTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDrugLigandSimilarityGraphTaskResponse), nil
	}
}

// CreateDrugLigandSimilarityGraphTaskInvoker 创建配体相似性图计算任务
func (c *EiHealthClient) CreateDrugLigandSimilarityGraphTaskInvoker(request *model.CreateDrugLigandSimilarityGraphTaskRequest) *CreateDrugLigandSimilarityGraphTaskInvoker {
	requestDef := GenReqDefForCreateDrugLigandSimilarityGraphTask()
	return &CreateDrugLigandSimilarityGraphTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDrugLigandSvg 生成分子SVG图
//
// 生成分子SVG图
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateDrugLigandSvg(request *model.CreateDrugLigandSvgRequest) (*model.CreateDrugLigandSvgResponse, error) {
	requestDef := GenReqDefForCreateDrugLigandSvg()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDrugLigandSvgResponse), nil
	}
}

// CreateDrugLigandSvgInvoker 生成分子SVG图
func (c *EiHealthClient) CreateDrugLigandSvgInvoker(request *model.CreateDrugLigandSvgRequest) *CreateDrugLigandSvgInvoker {
	requestDef := GenReqDefForCreateDrugLigandSvg()
	return &CreateDrugLigandSvgInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDrugLigandPreviewTask 删除配体文件预览任务
//
// 删除配体文件预览任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteDrugLigandPreviewTask(request *model.DeleteDrugLigandPreviewTaskRequest) (*model.DeleteDrugLigandPreviewTaskResponse, error) {
	requestDef := GenReqDefForDeleteDrugLigandPreviewTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDrugLigandPreviewTaskResponse), nil
	}
}

// DeleteDrugLigandPreviewTaskInvoker 删除配体文件预览任务
func (c *EiHealthClient) DeleteDrugLigandPreviewTaskInvoker(request *model.DeleteDrugLigandPreviewTaskRequest) *DeleteDrugLigandPreviewTaskInvoker {
	requestDef := GenReqDefForDeleteDrugLigandPreviewTask()
	return &DeleteDrugLigandPreviewTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDrugLigandSimilarityGraphTask 删除配体相似性图计算任务
//
// 删除配体相似性图计算任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteDrugLigandSimilarityGraphTask(request *model.DeleteDrugLigandSimilarityGraphTaskRequest) (*model.DeleteDrugLigandSimilarityGraphTaskResponse, error) {
	requestDef := GenReqDefForDeleteDrugLigandSimilarityGraphTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDrugLigandSimilarityGraphTaskResponse), nil
	}
}

// DeleteDrugLigandSimilarityGraphTaskInvoker 删除配体相似性图计算任务
func (c *EiHealthClient) DeleteDrugLigandSimilarityGraphTaskInvoker(request *model.DeleteDrugLigandSimilarityGraphTaskRequest) *DeleteDrugLigandSimilarityGraphTaskInvoker {
	requestDef := GenReqDefForDeleteDrugLigandSimilarityGraphTask()
	return &DeleteDrugLigandSimilarityGraphTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ParseDrugReceptorInfo 受体信息解析
//
// 受体信息解析，如果有多个受体蛋白则只处理第一个，如果一个受体蛋白里结合了多个配体，则最多只处理前10个
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ParseDrugReceptorInfo(request *model.ParseDrugReceptorInfoRequest) (*model.ParseDrugReceptorInfoResponse, error) {
	requestDef := GenReqDefForParseDrugReceptorInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ParseDrugReceptorInfoResponse), nil
	}
}

// ParseDrugReceptorInfoInvoker 受体信息解析
func (c *EiHealthClient) ParseDrugReceptorInfoInvoker(request *model.ParseDrugReceptorInfoRequest) *ParseDrugReceptorInfoInvoker {
	requestDef := GenReqDefForParseDrugReceptorInfo()
	return &ParseDrugReceptorInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecognizeDrugReceptorPocket 受体口袋检测
//
// 检测受体口袋，检测类型基于配体，基于氨基酸残基，自动检测，自定义和全局对接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) RecognizeDrugReceptorPocket(request *model.RecognizeDrugReceptorPocketRequest) (*model.RecognizeDrugReceptorPocketResponse, error) {
	requestDef := GenReqDefForRecognizeDrugReceptorPocket()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecognizeDrugReceptorPocketResponse), nil
	}
}

// RecognizeDrugReceptorPocketInvoker 受体口袋检测
func (c *EiHealthClient) RecognizeDrugReceptorPocketInvoker(request *model.RecognizeDrugReceptorPocketRequest) *RecognizeDrugReceptorPocketInvoker {
	requestDef := GenReqDefForRecognizeDrugReceptorPocket()
	return &RecognizeDrugReceptorPocketInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunDrugLigandToSmilesConversion 配体格式转换为SMILES
//
// 配体格式转换为SMILES，若配体文件中存在多个分子，则只取第一个返回
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) RunDrugLigandToSmilesConversion(request *model.RunDrugLigandToSmilesConversionRequest) (*model.RunDrugLigandToSmilesConversionResponse, error) {
	requestDef := GenReqDefForRunDrugLigandToSmilesConversion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunDrugLigandToSmilesConversionResponse), nil
	}
}

// RunDrugLigandToSmilesConversionInvoker 配体格式转换为SMILES
func (c *EiHealthClient) RunDrugLigandToSmilesConversionInvoker(request *model.RunDrugLigandToSmilesConversionRequest) *RunDrugLigandToSmilesConversionInvoker {
	requestDef := GenReqDefForRunDrugLigandToSmilesConversion()
	return &RunDrugLigandToSmilesConversionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunDrugReceptorPreprocess 受体预处理
//
// 受体预处理，用于前端显示预处理后的受体
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) RunDrugReceptorPreprocess(request *model.RunDrugReceptorPreprocessRequest) (*model.RunDrugReceptorPreprocessResponse, error) {
	requestDef := GenReqDefForRunDrugReceptorPreprocess()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunDrugReceptorPreprocessResponse), nil
	}
}

// RunDrugReceptorPreprocessInvoker 受体预处理
func (c *EiHealthClient) RunDrugReceptorPreprocessInvoker(request *model.RunDrugReceptorPreprocessRequest) *RunDrugReceptorPreprocessInvoker {
	requestDef := GenReqDefForRunDrugReceptorPreprocess()
	return &RunDrugReceptorPreprocessInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDrugLigandPreviewTask 查询配体文件预览任务
//
// 查询配体文件预览任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowDrugLigandPreviewTask(request *model.ShowDrugLigandPreviewTaskRequest) (*model.ShowDrugLigandPreviewTaskResponse, error) {
	requestDef := GenReqDefForShowDrugLigandPreviewTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDrugLigandPreviewTaskResponse), nil
	}
}

// ShowDrugLigandPreviewTaskInvoker 查询配体文件预览任务
func (c *EiHealthClient) ShowDrugLigandPreviewTaskInvoker(request *model.ShowDrugLigandPreviewTaskRequest) *ShowDrugLigandPreviewTaskInvoker {
	requestDef := GenReqDefForShowDrugLigandPreviewTask()
	return &ShowDrugLigandPreviewTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDrugLigandSimilarityGraphTask 查询配体相似性图计算任务
//
// 查询配体相似性图计算任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowDrugLigandSimilarityGraphTask(request *model.ShowDrugLigandSimilarityGraphTaskRequest) (*model.ShowDrugLigandSimilarityGraphTaskResponse, error) {
	requestDef := GenReqDefForShowDrugLigandSimilarityGraphTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDrugLigandSimilarityGraphTaskResponse), nil
	}
}

// ShowDrugLigandSimilarityGraphTaskInvoker 查询配体相似性图计算任务
func (c *EiHealthClient) ShowDrugLigandSimilarityGraphTaskInvoker(request *model.ShowDrugLigandSimilarityGraphTaskRequest) *ShowDrugLigandSimilarityGraphTaskInvoker {
	requestDef := GenReqDefForShowDrugLigandSimilarityGraphTask()
	return &ShowDrugLigandSimilarityGraphTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadData 文件下载
//
// 文件下载
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DownloadData(request *model.DownloadDataRequest) (*model.DownloadDataResponse, error) {
	requestDef := GenReqDefForDownloadData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadDataResponse), nil
	}
}

// DownloadDataInvoker 文件下载
func (c *EiHealthClient) DownloadDataInvoker(request *model.DownloadDataRequest) *DownloadDataInvoker {
	requestDef := GenReqDefForDownloadData()
	return &DownloadDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOverview 获取医疗平台信息
//
// 获取医疗平台信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowOverview(request *model.ShowOverviewRequest) (*model.ShowOverviewResponse, error) {
	requestDef := GenReqDefForShowOverview()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOverviewResponse), nil
	}
}

// ShowOverviewInvoker 获取医疗平台信息
func (c *EiHealthClient) ShowOverviewInvoker(request *model.ShowOverviewRequest) *ShowOverviewInvoker {
	requestDef := GenReqDefForShowOverview()
	return &ShowOverviewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateNotebook 创建notebook
//
// 创建notebook
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) CreateNotebook(request *model.CreateNotebookRequest) (*model.CreateNotebookResponse, error) {
	requestDef := GenReqDefForCreateNotebook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNotebookResponse), nil
	}
}

// CreateNotebookInvoker 创建notebook
func (c *EiHealthClient) CreateNotebookInvoker(request *model.CreateNotebookRequest) *CreateNotebookInvoker {
	requestDef := GenReqDefForCreateNotebook()
	return &CreateNotebookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteNotebook 删除notebook
//
// 删除notebook
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DeleteNotebook(request *model.DeleteNotebookRequest) (*model.DeleteNotebookResponse, error) {
	requestDef := GenReqDefForDeleteNotebook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNotebookResponse), nil
	}
}

// DeleteNotebookInvoker 删除notebook
func (c *EiHealthClient) DeleteNotebookInvoker(request *model.DeleteNotebookRequest) *DeleteNotebookInvoker {
	requestDef := GenReqDefForDeleteNotebook()
	return &DeleteNotebookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNotebook 获取notebook列表
//
// 获取notebook列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListNotebook(request *model.ListNotebookRequest) (*model.ListNotebookResponse, error) {
	requestDef := GenReqDefForListNotebook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNotebookResponse), nil
	}
}

// ListNotebookInvoker 获取notebook列表
func (c *EiHealthClient) ListNotebookInvoker(request *model.ListNotebookRequest) *ListNotebookInvoker {
	requestDef := GenReqDefForListNotebook()
	return &ListNotebookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNotebookTool 获取notebook工作环境
//
// 获取notebook工作环境
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListNotebookTool(request *model.ListNotebookToolRequest) (*model.ListNotebookToolResponse, error) {
	requestDef := GenReqDefForListNotebookTool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNotebookToolResponse), nil
	}
}

// ListNotebookToolInvoker 获取notebook工作环境
func (c *EiHealthClient) ListNotebookToolInvoker(request *model.ListNotebookToolRequest) *ListNotebookToolInvoker {
	requestDef := GenReqDefForListNotebookTool()
	return &ListNotebookToolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserNotebook 获取用户所属空间的notebook列表
//
// 获取用户所属空间的notebook列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListUserNotebook(request *model.ListUserNotebookRequest) (*model.ListUserNotebookResponse, error) {
	requestDef := GenReqDefForListUserNotebook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserNotebookResponse), nil
	}
}

// ListUserNotebookInvoker 获取用户所属空间的notebook列表
func (c *EiHealthClient) ListUserNotebookInvoker(request *model.ListUserNotebookRequest) *ListUserNotebookInvoker {
	requestDef := GenReqDefForListUserNotebook()
	return &ListUserNotebookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNotebook 获取notebook详情
//
// 获取notebook详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowNotebook(request *model.ShowNotebookRequest) (*model.ShowNotebookResponse, error) {
	requestDef := GenReqDefForShowNotebook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNotebookResponse), nil
	}
}

// ShowNotebookInvoker 获取notebook详情
func (c *EiHealthClient) ShowNotebookInvoker(request *model.ShowNotebookRequest) *ShowNotebookInvoker {
	requestDef := GenReqDefForShowNotebook()
	return &ShowNotebookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNotebookToken 获取notebook鉴权信息
//
// 获取notebook鉴权信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowNotebookToken(request *model.ShowNotebookTokenRequest) (*model.ShowNotebookTokenResponse, error) {
	requestDef := GenReqDefForShowNotebookToken()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNotebookTokenResponse), nil
	}
}

// ShowNotebookTokenInvoker 获取notebook鉴权信息
func (c *EiHealthClient) ShowNotebookTokenInvoker(request *model.ShowNotebookTokenRequest) *ShowNotebookTokenInvoker {
	requestDef := GenReqDefForShowNotebookToken()
	return &ShowNotebookTokenInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopOrStartNotebook 启停notebook
//
// 启停notebook
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) StopOrStartNotebook(request *model.StopOrStartNotebookRequest) (*model.StopOrStartNotebookResponse, error) {
	requestDef := GenReqDefForStopOrStartNotebook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopOrStartNotebookResponse), nil
	}
}

// StopOrStartNotebookInvoker 启停notebook
func (c *EiHealthClient) StopOrStartNotebookInvoker(request *model.StopOrStartNotebookRequest) *StopOrStartNotebookInvoker {
	requestDef := GenReqDefForStopOrStartNotebook()
	return &StopOrStartNotebookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNotebook 更新notebook
//
// 更新notebook
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) UpdateNotebook(request *model.UpdateNotebookRequest) (*model.UpdateNotebookResponse, error) {
	requestDef := GenReqDefForUpdateNotebook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNotebookResponse), nil
	}
}

// UpdateNotebookInvoker 更新notebook
func (c *EiHealthClient) UpdateNotebookInvoker(request *model.UpdateNotebookRequest) *UpdateNotebookInvoker {
	requestDef := GenReqDefForUpdateNotebook()
	return &UpdateNotebookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadPublicData 文件下载
//
// 文件下载
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) DownloadPublicData(request *model.DownloadPublicDataRequest) (*model.DownloadPublicDataResponse, error) {
	requestDef := GenReqDefForDownloadPublicData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadPublicDataResponse), nil
	}
}

// DownloadPublicDataInvoker 文件下载
func (c *EiHealthClient) DownloadPublicDataInvoker(request *model.DownloadPublicDataRequest) *DownloadPublicDataInvoker {
	requestDef := GenReqDefForDownloadPublicData()
	return &DownloadPublicDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListObsBucket 获取用户OBS桶列表
//
// 获取用户OBS桶列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListObsBucket(request *model.ListObsBucketRequest) (*model.ListObsBucketResponse, error) {
	requestDef := GenReqDefForListObsBucket()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListObsBucketResponse), nil
	}
}

// ListObsBucketInvoker 获取用户OBS桶列表
func (c *EiHealthClient) ListObsBucketInvoker(request *model.ListObsBucketRequest) *ListObsBucketInvoker {
	requestDef := GenReqDefForListObsBucket()
	return &ListObsBucketInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListObsBucketObject 获取用户OBS桶内对象
//
// 获取用户OBS桶内对象
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ListObsBucketObject(request *model.ListObsBucketObjectRequest) (*model.ListObsBucketObjectResponse, error) {
	requestDef := GenReqDefForListObsBucketObject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListObsBucketObjectResponse), nil
	}
}

// ListObsBucketObjectInvoker 获取用户OBS桶内对象
func (c *EiHealthClient) ListObsBucketObjectInvoker(request *model.ListObsBucketObjectRequest) *ListObsBucketObjectInvoker {
	requestDef := GenReqDefForListObsBucketObject()
	return &ListObsBucketObjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
