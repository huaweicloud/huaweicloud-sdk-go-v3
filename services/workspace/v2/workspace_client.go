package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/workspace/v2/model"
)

type WorkspaceClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewWorkspaceClient(hcClient *httpclient.HcHttpClient) *WorkspaceClient {
	return &WorkspaceClient{HcClient: hcClient}
}

func WorkspaceClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// ListAccessAddressBackupConfig 获取云办公服务接入地址备份配置
//
// 该接口用于获取云办公服务接入地址备份配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListAccessAddressBackupConfig(request *model.ListAccessAddressBackupConfigRequest) (*model.ListAccessAddressBackupConfigResponse, error) {
	requestDef := GenReqDefForListAccessAddressBackupConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccessAddressBackupConfigResponse), nil
	}
}

// ListAccessAddressBackupConfigInvoker 获取云办公服务接入地址备份配置
func (c *WorkspaceClient) ListAccessAddressBackupConfigInvoker(request *model.ListAccessAddressBackupConfigRequest) *ListAccessAddressBackupConfigInvoker {
	requestDef := GenReqDefForListAccessAddressBackupConfig()
	return &ListAccessAddressBackupConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAccessAddressBackupConfig 修改云办公服务接入地址备份配置
//
// 该接口用于修改云办公服务接入地址备份配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateAccessAddressBackupConfig(request *model.UpdateAccessAddressBackupConfigRequest) (*model.UpdateAccessAddressBackupConfigResponse, error) {
	requestDef := GenReqDefForUpdateAccessAddressBackupConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAccessAddressBackupConfigResponse), nil
	}
}

// UpdateAccessAddressBackupConfigInvoker 修改云办公服务接入地址备份配置
func (c *WorkspaceClient) UpdateAccessAddressBackupConfigInvoker(request *model.UpdateAccessAddressBackupConfigRequest) *UpdateAccessAddressBackupConfigInvoker {
	requestDef := GenReqDefForUpdateAccessAddressBackupConfig()
	return &UpdateAccessAddressBackupConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteAccessPolicies 删除接入策略
//
// 该接口用于删除指定接入策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchDeleteAccessPolicies(request *model.BatchDeleteAccessPoliciesRequest) (*model.BatchDeleteAccessPoliciesResponse, error) {
	requestDef := GenReqDefForBatchDeleteAccessPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteAccessPoliciesResponse), nil
	}
}

// BatchDeleteAccessPoliciesInvoker 删除接入策略
func (c *WorkspaceClient) BatchDeleteAccessPoliciesInvoker(request *model.BatchDeleteAccessPoliciesRequest) *BatchDeleteAccessPoliciesInvoker {
	requestDef := GenReqDefForBatchDeleteAccessPolicies()
	return &BatchDeleteAccessPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAccessPolicy 创建接入策略
//
// 该接口用于创建接入策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreateAccessPolicy(request *model.CreateAccessPolicyRequest) (*model.CreateAccessPolicyResponse, error) {
	requestDef := GenReqDefForCreateAccessPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAccessPolicyResponse), nil
	}
}

// CreateAccessPolicyInvoker 创建接入策略
func (c *WorkspaceClient) CreateAccessPolicyInvoker(request *model.CreateAccessPolicyRequest) *CreateAccessPolicyInvoker {
	requestDef := GenReqDefForCreateAccessPolicy()
	return &CreateAccessPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccessPolicies 查询接入策略
//
// 该接口用于查询接入策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListAccessPolicies(request *model.ListAccessPoliciesRequest) (*model.ListAccessPoliciesResponse, error) {
	requestDef := GenReqDefForListAccessPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccessPoliciesResponse), nil
	}
}

// ListAccessPoliciesInvoker 查询接入策略
func (c *WorkspaceClient) ListAccessPoliciesInvoker(request *model.ListAccessPoliciesRequest) *ListAccessPoliciesInvoker {
	requestDef := GenReqDefForListAccessPolicies()
	return &ListAccessPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccessPolicyObjects 查询指定接入策略的应用对象
//
// 该接口用于查询指定接入策略的应用对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListAccessPolicyObjects(request *model.ListAccessPolicyObjectsRequest) (*model.ListAccessPolicyObjectsResponse, error) {
	requestDef := GenReqDefForListAccessPolicyObjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccessPolicyObjectsResponse), nil
	}
}

// ListAccessPolicyObjectsInvoker 查询指定接入策略的应用对象
func (c *WorkspaceClient) ListAccessPolicyObjectsInvoker(request *model.ListAccessPolicyObjectsRequest) *ListAccessPolicyObjectsInvoker {
	requestDef := GenReqDefForListAccessPolicyObjects()
	return &ListAccessPolicyObjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAccessPolicy 更新指定接入策略
//
// 该接口用于更新指定接入策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateAccessPolicy(request *model.UpdateAccessPolicyRequest) (*model.UpdateAccessPolicyResponse, error) {
	requestDef := GenReqDefForUpdateAccessPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAccessPolicyResponse), nil
	}
}

// UpdateAccessPolicyInvoker 更新指定接入策略
func (c *WorkspaceClient) UpdateAccessPolicyInvoker(request *model.UpdateAccessPolicyRequest) *UpdateAccessPolicyInvoker {
	requestDef := GenReqDefForUpdateAccessPolicy()
	return &UpdateAccessPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAccessPolicyObjects 更新指定接入策略的应用对象
//
// 该接口用于更新指定接入策略的应用对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateAccessPolicyObjects(request *model.UpdateAccessPolicyObjectsRequest) (*model.UpdateAccessPolicyObjectsResponse, error) {
	requestDef := GenReqDefForUpdateAccessPolicyObjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAccessPolicyObjectsResponse), nil
	}
}

// UpdateAccessPolicyObjectsInvoker 更新指定接入策略的应用对象
func (c *WorkspaceClient) UpdateAccessPolicyObjectsInvoker(request *model.UpdateAccessPolicyObjectsRequest) *UpdateAccessPolicyObjectsInvoker {
	requestDef := GenReqDefForUpdateAccessPolicyObjects()
	return &UpdateAccessPolicyObjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAgencies 开通委托功能
//
// 开通委托功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreateAgencies(request *model.CreateAgenciesRequest) (*model.CreateAgenciesResponse, error) {
	requestDef := GenReqDefForCreateAgencies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAgenciesResponse), nil
	}
}

// CreateAgenciesInvoker 开通委托功能
func (c *WorkspaceClient) CreateAgenciesInvoker(request *model.CreateAgenciesRequest) *CreateAgenciesInvoker {
	requestDef := GenReqDefForCreateAgencies()
	return &CreateAgenciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAgencies 查询委托功能
//
// 查询委托功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListAgencies(request *model.ListAgenciesRequest) (*model.ListAgenciesResponse, error) {
	requestDef := GenReqDefForListAgencies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAgenciesResponse), nil
	}
}

// ListAgenciesInvoker 查询委托功能
func (c *WorkspaceClient) ListAgenciesInvoker(request *model.ListAgenciesRequest) *ListAgenciesInvoker {
	requestDef := GenReqDefForListAgencies()
	return &ListAgenciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarmStatistics 查询告警统计
//
// 返回各级别告警数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListAlarmStatistics(request *model.ListAlarmStatisticsRequest) (*model.ListAlarmStatisticsResponse, error) {
	requestDef := GenReqDefForListAlarmStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmStatisticsResponse), nil
	}
}

// ListAlarmStatisticsInvoker 查询告警统计
func (c *WorkspaceClient) ListAlarmStatisticsInvoker(request *model.ListAlarmStatisticsRequest) *ListAlarmStatisticsInvoker {
	requestDef := GenReqDefForListAlarmStatistics()
	return &ListAlarmStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarms 查询告警列表
//
// 从ces查询告警列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListAlarms(request *model.ListAlarmsRequest) (*model.ListAlarmsResponse, error) {
	requestDef := GenReqDefForListAlarms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmsResponse), nil
	}
}

// ListAlarmsInvoker 查询告警列表
func (c *WorkspaceClient) ListAlarmsInvoker(request *model.ListAlarmsRequest) *ListAlarmsInvoker {
	requestDef := GenReqDefForListAlarms()
	return &ListAlarmsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteApps 批量删除应用
//
// 批量删除应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchDeleteApps(request *model.BatchDeleteAppsRequest) (*model.BatchDeleteAppsResponse, error) {
	requestDef := GenReqDefForBatchDeleteApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteAppsResponse), nil
	}
}

// BatchDeleteAppsInvoker 批量删除应用
func (c *WorkspaceClient) BatchDeleteAppsInvoker(request *model.BatchDeleteAppsRequest) *BatchDeleteAppsInvoker {
	requestDef := GenReqDefForBatchDeleteApps()
	return &BatchDeleteAppsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteJobs 批量删除job
//
// 批量删除job。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchDeleteJobs(request *model.BatchDeleteJobsRequest) (*model.BatchDeleteJobsResponse, error) {
	requestDef := GenReqDefForBatchDeleteJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteJobsResponse), nil
	}
}

// BatchDeleteJobsInvoker 批量删除job
func (c *WorkspaceClient) BatchDeleteJobsInvoker(request *model.BatchDeleteJobsRequest) *BatchDeleteJobsInvoker {
	requestDef := GenReqDefForBatchDeleteJobs()
	return &BatchDeleteJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDisableApps 批量设置不可见
//
// 批量设置不可见。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchDisableApps(request *model.BatchDisableAppsRequest) (*model.BatchDisableAppsResponse, error) {
	requestDef := GenReqDefForBatchDisableApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDisableAppsResponse), nil
	}
}

// BatchDisableAppsInvoker 批量设置不可见
func (c *WorkspaceClient) BatchDisableAppsInvoker(request *model.BatchDisableAppsRequest) *BatchDisableAppsInvoker {
	requestDef := GenReqDefForBatchDisableApps()
	return &BatchDisableAppsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchEnableApps 批量设置可见
//
// 批量设置可见。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchEnableApps(request *model.BatchEnableAppsRequest) (*model.BatchEnableAppsResponse, error) {
	requestDef := GenReqDefForBatchEnableApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchEnableAppsResponse), nil
	}
}

// BatchEnableAppsInvoker 批量设置可见
func (c *WorkspaceClient) BatchEnableAppsInvoker(request *model.BatchEnableAppsRequest) *BatchEnableAppsInvoker {
	requestDef := GenReqDefForBatchEnableApps()
	return &BatchEnableAppsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchInstallApps 批量自动安装安装应用
//
// 批量自动安装安装应用 (应用需支持静默安装或者解压安装)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchInstallApps(request *model.BatchInstallAppsRequest) (*model.BatchInstallAppsResponse, error) {
	requestDef := GenReqDefForBatchInstallApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchInstallAppsResponse), nil
	}
}

// BatchInstallAppsInvoker 批量自动安装安装应用
func (c *WorkspaceClient) BatchInstallAppsInvoker(request *model.BatchInstallAppsRequest) *BatchInstallAppsInvoker {
	requestDef := GenReqDefForBatchInstallApps()
	return &BatchInstallAppsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateAppAuthorizations 批量设置应用授权
//
// 批量设置应用授权。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchUpdateAppAuthorizations(request *model.BatchUpdateAppAuthorizationsRequest) (*model.BatchUpdateAppAuthorizationsResponse, error) {
	requestDef := GenReqDefForBatchUpdateAppAuthorizations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateAppAuthorizationsResponse), nil
	}
}

// BatchUpdateAppAuthorizationsInvoker 批量设置应用授权
func (c *WorkspaceClient) BatchUpdateAppAuthorizationsInvoker(request *model.BatchUpdateAppAuthorizationsRequest) *BatchUpdateAppAuthorizationsInvoker {
	requestDef := GenReqDefForBatchUpdateAppAuthorizations()
	return &BatchUpdateAppAuthorizationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAndAuthorizeBucket 添加并授权默认桶
//
// 添加并授权默认桶,桶不存在时会自动创建OBS桶。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreateAndAuthorizeBucket(request *model.CreateAndAuthorizeBucketRequest) (*model.CreateAndAuthorizeBucketResponse, error) {
	requestDef := GenReqDefForCreateAndAuthorizeBucket()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAndAuthorizeBucketResponse), nil
	}
}

// CreateAndAuthorizeBucketInvoker 添加并授权默认桶
func (c *WorkspaceClient) CreateAndAuthorizeBucketInvoker(request *model.CreateAndAuthorizeBucketRequest) *CreateAndAuthorizeBucketInvoker {
	requestDef := GenReqDefForCreateAndAuthorizeBucket()
	return &CreateAndAuthorizeBucketInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateBucketCredential 生成访问凭证信息
//
// 生成桶凭证信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreateBucketCredential(request *model.CreateBucketCredentialRequest) (*model.CreateBucketCredentialResponse, error) {
	requestDef := GenReqDefForCreateBucketCredential()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBucketCredentialResponse), nil
	}
}

// CreateBucketCredentialInvoker 生成访问凭证信息
func (c *WorkspaceClient) CreateBucketCredentialInvoker(request *model.CreateBucketCredentialRequest) *CreateBucketCredentialInvoker {
	requestDef := GenReqDefForCreateBucketCredential()
	return &CreateBucketCredentialInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApp 删除应用
//
// 删除应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) DeleteApp(request *model.DeleteAppRequest) (*model.DeleteAppResponse, error) {
	requestDef := GenReqDefForDeleteApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppResponse), nil
	}
}

// DeleteAppInvoker 删除应用
func (c *WorkspaceClient) DeleteAppInvoker(request *model.DeleteAppRequest) *DeleteAppInvoker {
	requestDef := GenReqDefForDeleteApp()
	return &DeleteAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// InstallApp 自动安装应用
//
// 自动安装应用(应用需支持静默安装或者解压安装)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) InstallApp(request *model.InstallAppRequest) (*model.InstallAppResponse, error) {
	requestDef := GenReqDefForInstallApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.InstallAppResponse), nil
	}
}

// InstallAppInvoker 自动安装应用
func (c *WorkspaceClient) InstallAppInvoker(request *model.InstallAppRequest) *InstallAppInvoker {
	requestDef := GenReqDefForInstallApp()
	return &InstallAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppAuthorizations 查询应用授权信息
//
// 查询应用授权信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListAppAuthorizations(request *model.ListAppAuthorizationsRequest) (*model.ListAppAuthorizationsResponse, error) {
	requestDef := GenReqDefForListAppAuthorizations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppAuthorizationsResponse), nil
	}
}

// ListAppAuthorizationsInvoker 查询应用授权信息
func (c *WorkspaceClient) ListAppAuthorizationsInvoker(request *model.ListAppAuthorizationsRequest) *ListAppAuthorizationsInvoker {
	requestDef := GenReqDefForListAppAuthorizations()
	return &ListAppAuthorizationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppCatalogs 查询应用分类信息
//
// 查询应用分类信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListAppCatalogs(request *model.ListAppCatalogsRequest) (*model.ListAppCatalogsResponse, error) {
	requestDef := GenReqDefForListAppCatalogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppCatalogsResponse), nil
	}
}

// ListAppCatalogsInvoker 查询应用分类信息
func (c *WorkspaceClient) ListAppCatalogsInvoker(request *model.ListAppCatalogsRequest) *ListAppCatalogsInvoker {
	requestDef := GenReqDefForListAppCatalogs()
	return &ListAppCatalogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApps 按照名称分页查询应用
//
// 按照名称分页查询应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListApps(request *model.ListAppsRequest) (*model.ListAppsResponse, error) {
	requestDef := GenReqDefForListApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppsResponse), nil
	}
}

// ListAppsInvoker 按照名称分页查询应用
func (c *WorkspaceClient) ListAppsInvoker(request *model.ListAppsRequest) *ListAppsInvoker {
	requestDef := GenReqDefForListApps()
	return &ListAppsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJobs 查询应用安装job信息
//
// 查询应用安装job信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListJobs(request *model.ListJobsRequest) (*model.ListJobsResponse, error) {
	requestDef := GenReqDefForListJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobsResponse), nil
	}
}

// ListJobsInvoker 查询应用安装job信息
func (c *WorkspaceClient) ListJobsInvoker(request *model.ListJobsRequest) *ListJobsInvoker {
	requestDef := GenReqDefForListJobs()
	return &ListJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RetryJobs 重试失败job
//
// 重试指定失败job(仅支持失败job重试，其他类型job重试会响应错误)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) RetryJobs(request *model.RetryJobsRequest) (*model.RetryJobsResponse, error) {
	requestDef := GenReqDefForRetryJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RetryJobsResponse), nil
	}
}

// RetryJobsInvoker 重试失败job
func (c *WorkspaceClient) RetryJobsInvoker(request *model.RetryJobsRequest) *RetryJobsInvoker {
	requestDef := GenReqDefForRetryJobs()
	return &RetryJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAppAuthorizations 设置应用授权
//
// 设置应用授权。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateAppAuthorizations(request *model.UpdateAppAuthorizationsRequest) (*model.UpdateAppAuthorizationsResponse, error) {
	requestDef := GenReqDefForUpdateAppAuthorizations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppAuthorizationsResponse), nil
	}
}

// UpdateAppAuthorizationsInvoker 设置应用授权
func (c *WorkspaceClient) UpdateAppAuthorizationsInvoker(request *model.UpdateAppAuthorizationsRequest) *UpdateAppAuthorizationsInvoker {
	requestDef := GenReqDefForUpdateAppAuthorizations()
	return &UpdateAppAuthorizationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUploadedApp 修改应用
//
// 修改应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateUploadedApp(request *model.UpdateUploadedAppRequest) (*model.UpdateUploadedAppResponse, error) {
	requestDef := GenReqDefForUpdateUploadedApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUploadedAppResponse), nil
	}
}

// UpdateUploadedAppInvoker 修改应用
func (c *WorkspaceClient) UpdateUploadedAppInvoker(request *model.UpdateUploadedAppRequest) *UpdateUploadedAppInvoker {
	requestDef := GenReqDefForUpdateUploadedApp()
	return &UpdateUploadedAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadApp 添加应用
//
// 添加应用应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UploadApp(request *model.UploadAppRequest) (*model.UploadAppResponse, error) {
	requestDef := GenReqDefForUploadApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadAppResponse), nil
	}
}

// UploadAppInvoker 添加应用
func (c *WorkspaceClient) UploadAppInvoker(request *model.UploadAppRequest) *UploadAppInvoker {
	requestDef := GenReqDefForUploadApp()
	return &UploadAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddRestrictedRule 增加管控规则
//
// 增加管控规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) AddRestrictedRule(request *model.AddRestrictedRuleRequest) (*model.AddRestrictedRuleResponse, error) {
	requestDef := GenReqDefForAddRestrictedRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddRestrictedRuleResponse), nil
	}
}

// AddRestrictedRuleInvoker 增加管控规则
func (c *WorkspaceClient) AddRestrictedRuleInvoker(request *model.AddRestrictedRuleRequest) *AddRestrictedRuleInvoker {
	requestDef := GenReqDefForAddRestrictedRule()
	return &AddRestrictedRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteAppRules 批量删除规则
//
// 批量删除规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchDeleteAppRules(request *model.BatchDeleteAppRulesRequest) (*model.BatchDeleteAppRulesResponse, error) {
	requestDef := GenReqDefForBatchDeleteAppRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteAppRulesResponse), nil
	}
}

// BatchDeleteAppRulesInvoker 批量删除规则
func (c *WorkspaceClient) BatchDeleteAppRulesInvoker(request *model.BatchDeleteAppRulesRequest) *BatchDeleteAppRulesInvoker {
	requestDef := GenReqDefForBatchDeleteAppRules()
	return &BatchDeleteAppRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAppRule 创建应用规则
//
// 创建应用规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreateAppRule(request *model.CreateAppRuleRequest) (*model.CreateAppRuleResponse, error) {
	requestDef := GenReqDefForCreateAppRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppRuleResponse), nil
	}
}

// CreateAppRuleInvoker 创建应用规则
func (c *WorkspaceClient) CreateAppRuleInvoker(request *model.CreateAppRuleRequest) *CreateAppRuleInvoker {
	requestDef := GenReqDefForCreateAppRule()
	return &CreateAppRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAppRule 删除应用规则
//
// 删除应用规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) DeleteAppRule(request *model.DeleteAppRuleRequest) (*model.DeleteAppRuleResponse, error) {
	requestDef := GenReqDefForDeleteAppRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppRuleResponse), nil
	}
}

// DeleteAppRuleInvoker 删除应用规则
func (c *WorkspaceClient) DeleteAppRuleInvoker(request *model.DeleteAppRuleRequest) *DeleteAppRuleInvoker {
	requestDef := GenReqDefForDeleteAppRule()
	return &DeleteAppRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRestrictedRule 批量删除管控规则列表
//
// 批量删除管控规则列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) DeleteRestrictedRule(request *model.DeleteRestrictedRuleRequest) (*model.DeleteRestrictedRuleResponse, error) {
	requestDef := GenReqDefForDeleteRestrictedRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRestrictedRuleResponse), nil
	}
}

// DeleteRestrictedRuleInvoker 批量删除管控规则列表
func (c *WorkspaceClient) DeleteRestrictedRuleInvoker(request *model.DeleteRestrictedRuleRequest) *DeleteRestrictedRuleInvoker {
	requestDef := GenReqDefForDeleteRestrictedRule()
	return &DeleteRestrictedRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableRuleRestriction 禁用规则管控
//
// 禁用规则管控。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) DisableRuleRestriction(request *model.DisableRuleRestrictionRequest) (*model.DisableRuleRestrictionResponse, error) {
	requestDef := GenReqDefForDisableRuleRestriction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableRuleRestrictionResponse), nil
	}
}

// DisableRuleRestrictionInvoker 禁用规则管控
func (c *WorkspaceClient) DisableRuleRestrictionInvoker(request *model.DisableRuleRestrictionRequest) *DisableRuleRestrictionInvoker {
	requestDef := GenReqDefForDisableRuleRestriction()
	return &DisableRuleRestrictionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableRuleRestriction 启用规则管控
//
// 启用规则管控。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) EnableRuleRestriction(request *model.EnableRuleRestrictionRequest) (*model.EnableRuleRestrictionResponse, error) {
	requestDef := GenReqDefForEnableRuleRestriction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableRuleRestrictionResponse), nil
	}
}

// EnableRuleRestrictionInvoker 启用规则管控
func (c *WorkspaceClient) EnableRuleRestrictionInvoker(request *model.EnableRuleRestrictionRequest) *EnableRuleRestrictionInvoker {
	requestDef := GenReqDefForEnableRuleRestriction()
	return &EnableRuleRestrictionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppRule 查询应用规则
//
// 查询应用规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListAppRule(request *model.ListAppRuleRequest) (*model.ListAppRuleResponse, error) {
	requestDef := GenReqDefForListAppRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppRuleResponse), nil
	}
}

// ListAppRuleInvoker 查询应用规则
func (c *WorkspaceClient) ListAppRuleInvoker(request *model.ListAppRuleRequest) *ListAppRuleInvoker {
	requestDef := GenReqDefForListAppRule()
	return &ListAppRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRestrictedRule 查询管控规则列表
//
// 查询管控规则列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListRestrictedRule(request *model.ListRestrictedRuleRequest) (*model.ListRestrictedRuleResponse, error) {
	requestDef := GenReqDefForListRestrictedRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRestrictedRuleResponse), nil
	}
}

// ListRestrictedRuleInvoker 查询管控规则列表
func (c *WorkspaceClient) ListRestrictedRuleInvoker(request *model.ListRestrictedRuleRequest) *ListRestrictedRuleInvoker {
	requestDef := GenReqDefForListRestrictedRule()
	return &ListRestrictedRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAppRule 修改应用规则
//
// 修改应用规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateAppRule(request *model.UpdateAppRuleRequest) (*model.UpdateAppRuleResponse, error) {
	requestDef := GenReqDefForUpdateAppRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppRuleResponse), nil
	}
}

// UpdateAppRuleInvoker 修改应用规则
func (c *WorkspaceClient) UpdateAppRuleInvoker(request *model.UpdateAppRuleRequest) *UpdateAppRuleInvoker {
	requestDef := GenReqDefForUpdateAppRule()
	return &UpdateAppRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAssistAuthConfig 查询辅助认证配置
//
// 查询辅助认证的配置信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowAssistAuthConfig(request *model.ShowAssistAuthConfigRequest) (*model.ShowAssistAuthConfigResponse, error) {
	requestDef := GenReqDefForShowAssistAuthConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAssistAuthConfigResponse), nil
	}
}

// ShowAssistAuthConfigInvoker 查询辅助认证配置
func (c *WorkspaceClient) ShowAssistAuthConfigInvoker(request *model.ShowAssistAuthConfigRequest) *ShowAssistAuthConfigInvoker {
	requestDef := GenReqDefForShowAssistAuthConfig()
	return &ShowAssistAuthConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuthConfig 查询认证登录方式
//
// 查询认证登录方式配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowAuthConfig(request *model.ShowAuthConfigRequest) (*model.ShowAuthConfigResponse, error) {
	requestDef := GenReqDefForShowAuthConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuthConfigResponse), nil
	}
}

// ShowAuthConfigInvoker 查询认证登录方式
func (c *WorkspaceClient) ShowAuthConfigInvoker(request *model.ShowAuthConfigRequest) *ShowAuthConfigInvoker {
	requestDef := GenReqDefForShowAuthConfig()
	return &ShowAuthConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAssistAuthMethodConfig 更新辅助认证策略配置
//
// 更新辅助认证策略配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateAssistAuthMethodConfig(request *model.UpdateAssistAuthMethodConfigRequest) (*model.UpdateAssistAuthMethodConfigResponse, error) {
	requestDef := GenReqDefForUpdateAssistAuthMethodConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAssistAuthMethodConfigResponse), nil
	}
}

// UpdateAssistAuthMethodConfigInvoker 更新辅助认证策略配置
func (c *WorkspaceClient) UpdateAssistAuthMethodConfigInvoker(request *model.UpdateAssistAuthMethodConfigRequest) *UpdateAssistAuthMethodConfigInvoker {
	requestDef := GenReqDefForUpdateAssistAuthMethodConfig()
	return &UpdateAssistAuthMethodConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAuthMethodConfig 更新认证策略配置
//
// 更新认证策略配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateAuthMethodConfig(request *model.UpdateAuthMethodConfigRequest) (*model.UpdateAuthMethodConfigResponse, error) {
	requestDef := GenReqDefForUpdateAuthMethodConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAuthMethodConfigResponse), nil
	}
}

// UpdateAuthMethodConfigInvoker 更新认证策略配置
func (c *WorkspaceClient) UpdateAuthMethodConfigInvoker(request *model.UpdateAuthMethodConfigRequest) *UpdateAuthMethodConfigInvoker {
	requestDef := GenReqDefForUpdateAuthMethodConfig()
	return &UpdateAuthMethodConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAvailabilityZones 查询可用分区列表
//
// 该接口用于查询云桌面支持的可用分区列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListAvailabilityZones(request *model.ListAvailabilityZonesRequest) (*model.ListAvailabilityZonesResponse, error) {
	requestDef := GenReqDefForListAvailabilityZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailabilityZonesResponse), nil
	}
}

// ListAvailabilityZonesInvoker 查询可用分区列表
func (c *WorkspaceClient) ListAvailabilityZonesInvoker(request *model.ListAvailabilityZonesRequest) *ListAvailabilityZonesInvoker {
	requestDef := GenReqDefForListAvailabilityZones()
	return &ListAvailabilityZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAzs 查询可用分区列表概要
//
// 该接口用于查询云桌面支持的可用分区列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListAzs(request *model.ListAzsRequest) (*model.ListAzsResponse, error) {
	requestDef := GenReqDefForListAzs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAzsResponse), nil
	}
}

// ListAzsInvoker 查询可用分区列表概要
func (c *WorkspaceClient) ListAzsInvoker(request *model.ListAzsRequest) *ListAzsInvoker {
	requestDef := GenReqDefForListAzs()
	return &ListAzsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAzDetails 查询可用分区详情
//
// 该接口用于查询云桌面支持的可用分区列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowAzDetails(request *model.ShowAzDetailsRequest) (*model.ShowAzDetailsResponse, error) {
	requestDef := GenReqDefForShowAzDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAzDetailsResponse), nil
	}
}

// ShowAzDetailsInvoker 查询可用分区详情
func (c *WorkspaceClient) ShowAzDetailsInvoker(request *model.ShowAzDetailsRequest) *ShowAzDetailsInvoker {
	requestDef := GenReqDefForShowAzDetails()
	return &ShowAzDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportUserLoginInfoNew 导出连接记录
//
// 该接口用于导出连接记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ExportUserLoginInfoNew(request *model.ExportUserLoginInfoNewRequest) (*model.ExportUserLoginInfoNewResponse, error) {
	requestDef := GenReqDefForExportUserLoginInfoNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportUserLoginInfoNewResponse), nil
	}
}

// ExportUserLoginInfoNewInvoker 导出连接记录
func (c *WorkspaceClient) ExportUserLoginInfoNewInvoker(request *model.ExportUserLoginInfoNewRequest) *ExportUserLoginInfoNewInvoker {
	requestDef := GenReqDefForExportUserLoginInfoNew()
	return &ExportUserLoginInfoNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListDesktopsStatus 查询桌面登录状态
//
// 该接口用于查询桌面登录状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListDesktopsStatus(request *model.ListDesktopsStatusRequest) (*model.ListDesktopsStatusResponse, error) {
	requestDef := GenReqDefForListDesktopsStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDesktopsStatusResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListDesktopsStatusInvoker 查询桌面登录状态
func (c *WorkspaceClient) ListDesktopsStatusInvoker(request *model.ListDesktopsStatusRequest) *ListDesktopsStatusInvoker {
	requestDef := GenReqDefForListDesktopsStatus()
	return &ListDesktopsStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHistoryOnlineInfo 查询登录人数
//
// 该接口用于查询登录人数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListHistoryOnlineInfo(request *model.ListHistoryOnlineInfoRequest) (*model.ListHistoryOnlineInfoResponse, error) {
	requestDef := GenReqDefForListHistoryOnlineInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHistoryOnlineInfoResponse), nil
	}
}

// ListHistoryOnlineInfoInvoker 查询登录人数
func (c *WorkspaceClient) ListHistoryOnlineInfoInvoker(request *model.ListHistoryOnlineInfoRequest) *ListHistoryOnlineInfoInvoker {
	requestDef := GenReqDefForListHistoryOnlineInfo()
	return &ListHistoryOnlineInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHistoryOnlineInfoNew 查询登录人数
//
// 该接口用于查询登录人数，注意查询参数不可全空。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListHistoryOnlineInfoNew(request *model.ListHistoryOnlineInfoNewRequest) (*model.ListHistoryOnlineInfoNewResponse, error) {
	requestDef := GenReqDefForListHistoryOnlineInfoNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHistoryOnlineInfoNewResponse), nil
	}
}

// ListHistoryOnlineInfoNewInvoker 查询登录人数
func (c *WorkspaceClient) ListHistoryOnlineInfoNewInvoker(request *model.ListHistoryOnlineInfoNewRequest) *ListHistoryOnlineInfoNewInvoker {
	requestDef := GenReqDefForListHistoryOnlineInfoNew()
	return &ListHistoryOnlineInfoNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstancesStatus 查询桌面登录状态
//
// 该接口用于查询桌面登录状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListInstancesStatus(request *model.ListInstancesStatusRequest) (*model.ListInstancesStatusResponse, error) {
	requestDef := GenReqDefForListInstancesStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesStatusResponse), nil
	}
}

// ListInstancesStatusInvoker 查询桌面登录状态
func (c *WorkspaceClient) ListInstancesStatusInvoker(request *model.ListInstancesStatusRequest) *ListInstancesStatusInvoker {
	requestDef := GenReqDefForListInstancesStatus()
	return &ListInstancesStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLoginRecordsNew 查询登录信息
//
// 该接口用于查询登录信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListLoginRecordsNew(request *model.ListLoginRecordsNewRequest) (*model.ListLoginRecordsNewResponse, error) {
	requestDef := GenReqDefForListLoginRecordsNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLoginRecordsNewResponse), nil
	}
}

// ListLoginRecordsNewInvoker 查询登录信息
func (c *WorkspaceClient) ListLoginRecordsNewInvoker(request *model.ListLoginRecordsNewRequest) *ListLoginRecordsNewInvoker {
	requestDef := GenReqDefForListLoginRecordsNew()
	return &ListLoginRecordsNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AttachInstances 分配用户
//
// 将桌面分配给用户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) AttachInstances(request *model.AttachInstancesRequest) (*model.AttachInstancesResponse, error) {
	requestDef := GenReqDefForAttachInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachInstancesResponse), nil
	}
}

// AttachInstancesInvoker 分配用户
func (c *WorkspaceClient) AttachInstancesInvoker(request *model.AttachInstancesRequest) *AttachInstancesInvoker {
	requestDef := GenReqDefForAttachInstances()
	return &AttachInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAssociateInstances 预批量分配用户
//
// 预批量将桌面分配给用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchAssociateInstances(request *model.BatchAssociateInstancesRequest) (*model.BatchAssociateInstancesResponse, error) {
	requestDef := GenReqDefForBatchAssociateInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAssociateInstancesResponse), nil
	}
}

// BatchAssociateInstancesInvoker 预批量分配用户
func (c *WorkspaceClient) BatchAssociateInstancesInvoker(request *model.BatchAssociateInstancesRequest) *BatchAssociateInstancesInvoker {
	requestDef := GenReqDefForBatchAssociateInstances()
	return &BatchAssociateInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAttachInstances 批量分配用户
//
// 批量分配桌面给用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchAttachInstances(request *model.BatchAttachInstancesRequest) (*model.BatchAttachInstancesResponse, error) {
	requestDef := GenReqDefForBatchAttachInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAttachInstancesResponse), nil
	}
}

// BatchAttachInstancesInvoker 批量分配用户
func (c *WorkspaceClient) BatchAttachInstancesInvoker(request *model.BatchAttachInstancesRequest) *BatchAttachInstancesInvoker {
	requestDef := GenReqDefForBatchAttachInstances()
	return &BatchAttachInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchChangeDesktopNetwork 批量切换桌面网络
//
// 批量切换桌面vpc、子网、ip、安全组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchChangeDesktopNetwork(request *model.BatchChangeDesktopNetworkRequest) (*model.BatchChangeDesktopNetworkResponse, error) {
	requestDef := GenReqDefForBatchChangeDesktopNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchChangeDesktopNetworkResponse), nil
	}
}

// BatchChangeDesktopNetworkInvoker 批量切换桌面网络
func (c *WorkspaceClient) BatchChangeDesktopNetworkInvoker(request *model.BatchChangeDesktopNetworkRequest) *BatchChangeDesktopNetworkInvoker {
	requestDef := GenReqDefForBatchChangeDesktopNetwork()
	return &BatchChangeDesktopNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteDesktops 批量删除桌面
//
// 批量删除桌面，删除后无法恢复。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchDeleteDesktops(request *model.BatchDeleteDesktopsRequest) (*model.BatchDeleteDesktopsResponse, error) {
	requestDef := GenReqDefForBatchDeleteDesktops()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteDesktopsResponse), nil
	}
}

// BatchDeleteDesktopsInvoker 批量删除桌面
func (c *WorkspaceClient) BatchDeleteDesktopsInvoker(request *model.BatchDeleteDesktopsRequest) *BatchDeleteDesktopsInvoker {
	requestDef := GenReqDefForBatchDeleteDesktops()
	return &BatchDeleteDesktopsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDetachInstances 批量解绑用户
//
// 批量将桌面和用户解绑
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchDetachInstances(request *model.BatchDetachInstancesRequest) (*model.BatchDetachInstancesResponse, error) {
	requestDef := GenReqDefForBatchDetachInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDetachInstancesResponse), nil
	}
}

// BatchDetachInstancesInvoker 批量解绑用户
func (c *WorkspaceClient) BatchDetachInstancesInvoker(request *model.BatchDetachInstancesRequest) *BatchDetachInstancesInvoker {
	requestDef := GenReqDefForBatchDetachInstances()
	return &BatchDetachInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchInstallAgent 安装ces-agent
//
// 批量为桌面安装agent
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchInstallAgent(request *model.BatchInstallAgentRequest) (*model.BatchInstallAgentResponse, error) {
	requestDef := GenReqDefForBatchInstallAgent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchInstallAgentResponse), nil
	}
}

// BatchInstallAgentInvoker 安装ces-agent
func (c *WorkspaceClient) BatchInstallAgentInvoker(request *model.BatchInstallAgentRequest) *BatchInstallAgentInvoker {
	requestDef := GenReqDefForBatchInstallAgent()
	return &BatchInstallAgentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchLogoffDesktops 批量注销桌面
//
// 批量注销桌面。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchLogoffDesktops(request *model.BatchLogoffDesktopsRequest) (*model.BatchLogoffDesktopsResponse, error) {
	requestDef := GenReqDefForBatchLogoffDesktops()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchLogoffDesktopsResponse), nil
	}
}

// BatchLogoffDesktopsInvoker 批量注销桌面
func (c *WorkspaceClient) BatchLogoffDesktopsInvoker(request *model.BatchLogoffDesktopsRequest) *BatchLogoffDesktopsInvoker {
	requestDef := GenReqDefForBatchLogoffDesktops()
	return &BatchLogoffDesktopsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchRebuildDesktopsSystemDisk 重建桌面
//
// 批量重建桌面系统盘。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchRebuildDesktopsSystemDisk(request *model.BatchRebuildDesktopsSystemDiskRequest) (*model.BatchRebuildDesktopsSystemDiskResponse, error) {
	requestDef := GenReqDefForBatchRebuildDesktopsSystemDisk()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRebuildDesktopsSystemDiskResponse), nil
	}
}

// BatchRebuildDesktopsSystemDiskInvoker 重建桌面
func (c *WorkspaceClient) BatchRebuildDesktopsSystemDiskInvoker(request *model.BatchRebuildDesktopsSystemDiskRequest) *BatchRebuildDesktopsSystemDiskInvoker {
	requestDef := GenReqDefForBatchRebuildDesktopsSystemDisk()
	return &BatchRebuildDesktopsSystemDiskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchRunDesktops 操作桌面
//
// 批量操作桌面，用于批量开机、关机和重启。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchRunDesktops(request *model.BatchRunDesktopsRequest) (*model.BatchRunDesktopsResponse, error) {
	requestDef := GenReqDefForBatchRunDesktops()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRunDesktopsResponse), nil
	}
}

// BatchRunDesktopsInvoker 操作桌面
func (c *WorkspaceClient) BatchRunDesktopsInvoker(request *model.BatchRunDesktopsRequest) *BatchRunDesktopsInvoker {
	requestDef := GenReqDefForBatchRunDesktops()
	return &BatchRunDesktopsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelRemoteAssistance 取消远程协助
//
// 取消远程协助。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CancelRemoteAssistance(request *model.CancelRemoteAssistanceRequest) (*model.CancelRemoteAssistanceResponse, error) {
	requestDef := GenReqDefForCancelRemoteAssistance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelRemoteAssistanceResponse), nil
	}
}

// CancelRemoteAssistanceInvoker 取消远程协助
func (c *WorkspaceClient) CancelRemoteAssistanceInvoker(request *model.CancelRemoteAssistanceRequest) *CancelRemoteAssistanceInvoker {
	requestDef := GenReqDefForCancelRemoteAssistance()
	return &CancelRemoteAssistanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeDesktopNetwork 切换桌面网络
//
// 切换桌面vpc、子网、ip、安全组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ChangeDesktopNetwork(request *model.ChangeDesktopNetworkRequest) (*model.ChangeDesktopNetworkResponse, error) {
	requestDef := GenReqDefForChangeDesktopNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeDesktopNetworkResponse), nil
	}
}

// ChangeDesktopNetworkInvoker 切换桌面网络
func (c *WorkspaceClient) ChangeDesktopNetworkInvoker(request *model.ChangeDesktopNetworkRequest) *ChangeDesktopNetworkInvoker {
	requestDef := GenReqDefForChangeDesktopNetwork()
	return &ChangeDesktopNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeDesktopToImage 桌面转镜像
//
// 桌面转镜像
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ChangeDesktopToImage(request *model.ChangeDesktopToImageRequest) (*model.ChangeDesktopToImageResponse, error) {
	requestDef := GenReqDefForChangeDesktopToImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeDesktopToImageResponse), nil
	}
}

// ChangeDesktopToImageInvoker 桌面转镜像
func (c *WorkspaceClient) ChangeDesktopToImageInvoker(request *model.ChangeDesktopToImageRequest) *ChangeDesktopToImageInvoker {
	requestDef := GenReqDefForChangeDesktopToImage()
	return &ChangeDesktopToImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeUserPrivilegeGroup 批量修改用户权限组
//
// 批量修改用户权限组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ChangeUserPrivilegeGroup(request *model.ChangeUserPrivilegeGroupRequest) (*model.ChangeUserPrivilegeGroupResponse, error) {
	requestDef := GenReqDefForChangeUserPrivilegeGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeUserPrivilegeGroupResponse), nil
	}
}

// ChangeUserPrivilegeGroupInvoker 批量修改用户权限组
func (c *WorkspaceClient) ChangeUserPrivilegeGroupInvoker(request *model.ChangeUserPrivilegeGroupRequest) *ChangeUserPrivilegeGroupInvoker {
	requestDef := GenReqDefForChangeUserPrivilegeGroup()
	return &ChangeUserPrivilegeGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDesktop 创建桌面
//
// 创建桌面，并将此桌面分配给用户，当桌面创建成功后用户可以登录使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreateDesktop(request *model.CreateDesktopRequest) (*model.CreateDesktopResponse, error) {
	requestDef := GenReqDefForCreateDesktop()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDesktopResponse), nil
	}
}

// CreateDesktopInvoker 创建桌面
func (c *WorkspaceClient) CreateDesktopInvoker(request *model.CreateDesktopRequest) *CreateDesktopInvoker {
	requestDef := GenReqDefForCreateDesktop()
	return &CreateDesktopInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRemoteAssistance 创建远程协助
//
// 创建远程协助。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreateRemoteAssistance(request *model.CreateRemoteAssistanceRequest) (*model.CreateRemoteAssistanceResponse, error) {
	requestDef := GenReqDefForCreateRemoteAssistance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRemoteAssistanceResponse), nil
	}
}

// CreateRemoteAssistanceInvoker 创建远程协助
func (c *WorkspaceClient) CreateRemoteAssistanceInvoker(request *model.CreateRemoteAssistanceRequest) *CreateRemoteAssistanceInvoker {
	requestDef := GenReqDefForCreateRemoteAssistance()
	return &CreateRemoteAssistanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDesktop 删除单个桌面
//
// 删除单个桌面，删除后无法恢复。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) DeleteDesktop(request *model.DeleteDesktopRequest) (*model.DeleteDesktopResponse, error) {
	requestDef := GenReqDefForDeleteDesktop()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDesktopResponse), nil
	}
}

// DeleteDesktopInvoker 删除单个桌面
func (c *WorkspaceClient) DeleteDesktopInvoker(request *model.DeleteDesktopRequest) *DeleteDesktopInvoker {
	requestDef := GenReqDefForDeleteDesktop()
	return &DeleteDesktopInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetachInstances 解绑用户
//
// 将桌面和用户解绑
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) DetachInstances(request *model.DetachInstancesRequest) (*model.DetachInstancesResponse, error) {
	requestDef := GenReqDefForDetachInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetachInstancesResponse), nil
	}
}

// DetachInstancesInvoker 解绑用户
func (c *WorkspaceClient) DetachInstancesInvoker(request *model.DetachInstancesRequest) *DetachInstancesInvoker {
	requestDef := GenReqDefForDetachInstances()
	return &DetachInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAgentsInstallCondition 查询桌面安装agent详情
//
// 展示桌面安装agent详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListAgentsInstallCondition(request *model.ListAgentsInstallConditionRequest) (*model.ListAgentsInstallConditionResponse, error) {
	requestDef := GenReqDefForListAgentsInstallCondition()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAgentsInstallConditionResponse), nil
	}
}

// ListAgentsInstallConditionInvoker 查询桌面安装agent详情
func (c *WorkspaceClient) ListAgentsInstallConditionInvoker(request *model.ListAgentsInstallConditionRequest) *ListAgentsInstallConditionInvoker {
	requestDef := GenReqDefForListAgentsInstallCondition()
	return &ListAgentsInstallConditionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDesktopActions 查询桌面开关机信息
//
// 获取桌面开关机信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListDesktopActions(request *model.ListDesktopActionsRequest) (*model.ListDesktopActionsResponse, error) {
	requestDef := GenReqDefForListDesktopActions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDesktopActionsResponse), nil
	}
}

// ListDesktopActionsInvoker 查询桌面开关机信息
func (c *WorkspaceClient) ListDesktopActionsInvoker(request *model.ListDesktopActionsRequest) *ListDesktopActionsInvoker {
	requestDef := GenReqDefForListDesktopActions()
	return &ListDesktopActionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDesktopDetachInfo 查询桌面解绑信息
//
// 查询桌面解绑信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListDesktopDetachInfo(request *model.ListDesktopDetachInfoRequest) (*model.ListDesktopDetachInfoResponse, error) {
	requestDef := GenReqDefForListDesktopDetachInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDesktopDetachInfoResponse), nil
	}
}

// ListDesktopDetachInfoInvoker 查询桌面解绑信息
func (c *WorkspaceClient) ListDesktopDetachInfoInvoker(request *model.ListDesktopDetachInfoRequest) *ListDesktopDetachInfoInvoker {
	requestDef := GenReqDefForListDesktopDetachInfo()
	return &ListDesktopDetachInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDesktops 查询桌面列表
//
// 该接口用于查询桌面虚拟机列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListDesktops(request *model.ListDesktopsRequest) (*model.ListDesktopsResponse, error) {
	requestDef := GenReqDefForListDesktops()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDesktopsResponse), nil
	}
}

// ListDesktopsInvoker 查询桌面列表
func (c *WorkspaceClient) ListDesktopsInvoker(request *model.ListDesktopsRequest) *ListDesktopsInvoker {
	requestDef := GenReqDefForListDesktops()
	return &ListDesktopsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDesktopsConnectStatus 查询桌面连接状态列表
//
// 查询桌面连接状态列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListDesktopsConnectStatus(request *model.ListDesktopsConnectStatusRequest) (*model.ListDesktopsConnectStatusResponse, error) {
	requestDef := GenReqDefForListDesktopsConnectStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDesktopsConnectStatusResponse), nil
	}
}

// ListDesktopsConnectStatusInvoker 查询桌面连接状态列表
func (c *WorkspaceClient) ListDesktopsConnectStatusInvoker(request *model.ListDesktopsConnectStatusRequest) *ListDesktopsConnectStatusInvoker {
	requestDef := GenReqDefForListDesktopsConnectStatus()
	return &ListDesktopsConnectStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDesktopsDetail 查询桌面详情列表
//
// 查询桌面详情信息列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListDesktopsDetail(request *model.ListDesktopsDetailRequest) (*model.ListDesktopsDetailResponse, error) {
	requestDef := GenReqDefForListDesktopsDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDesktopsDetailResponse), nil
	}
}

// ListDesktopsDetailInvoker 查询桌面详情列表
func (c *WorkspaceClient) ListDesktopsDetailInvoker(request *model.ListDesktopsDetailRequest) *ListDesktopsDetailInvoker {
	requestDef := GenReqDefForListDesktopsDetail()
	return &ListDesktopsDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RegisterDomain 重新加入AD域
//
// 该接口用于Windows桌面重新加入AD域，一般用于解决桌面脱域的情况使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) RegisterDomain(request *model.RegisterDomainRequest) (*model.RegisterDomainResponse, error) {
	requestDef := GenReqDefForRegisterDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterDomainResponse), nil
	}
}

// RegisterDomainInvoker 重新加入AD域
func (c *WorkspaceClient) RegisterDomainInvoker(request *model.RegisterDomainRequest) *RegisterDomainInvoker {
	requestDef := GenReqDefForRegisterDomain()
	return &RegisterDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeDesktop 变更规格
//
// 变更云桌面规格，只支持变更CPU和内存，不支持变更磁盘，不支持同规格互相变更。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ResizeDesktop(request *model.ResizeDesktopRequest) (*model.ResizeDesktopResponse, error) {
	requestDef := GenReqDefForResizeDesktop()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeDesktopResponse), nil
	}
}

// ResizeDesktopInvoker 变更规格
func (c *WorkspaceClient) ResizeDesktopInvoker(request *model.ResizeDesktopRequest) *ResizeDesktopInvoker {
	requestDef := GenReqDefForResizeDesktop()
	return &ResizeDesktopInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendNotifications 发送消息通知
//
// 用于管理员向桌面发送消息通知
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) SendNotifications(request *model.SendNotificationsRequest) (*model.SendNotificationsResponse, error) {
	requestDef := GenReqDefForSendNotifications()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendNotificationsResponse), nil
	}
}

// SendNotificationsInvoker 发送消息通知
func (c *WorkspaceClient) SendNotificationsInvoker(request *model.SendNotificationsRequest) *SendNotificationsInvoker {
	requestDef := GenReqDefForSendNotifications()
	return &SendNotificationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetMaintenanceMode 批量设置桌面维护模式
//
// 批量设置桌面管理员维护模式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) SetMaintenanceMode(request *model.SetMaintenanceModeRequest) (*model.SetMaintenanceModeResponse, error) {
	requestDef := GenReqDefForSetMaintenanceMode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetMaintenanceModeResponse), nil
	}
}

// SetMaintenanceModeInvoker 批量设置桌面维护模式
func (c *WorkspaceClient) SetMaintenanceModeInvoker(request *model.SetMaintenanceModeRequest) *SetMaintenanceModeInvoker {
	requestDef := GenReqDefForSetMaintenanceMode()
	return &SetMaintenanceModeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDesktopDetail 查询单个桌面详情
//
// 指定桌面Id查询详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowDesktopDetail(request *model.ShowDesktopDetailRequest) (*model.ShowDesktopDetailResponse, error) {
	requestDef := GenReqDefForShowDesktopDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDesktopDetailResponse), nil
	}
}

// ShowDesktopDetailInvoker 查询单个桌面详情
func (c *WorkspaceClient) ShowDesktopDetailInvoker(request *model.ShowDesktopDetailRequest) *ShowDesktopDetailInvoker {
	requestDef := GenReqDefForShowDesktopDetail()
	return &ShowDesktopDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDesktopMonitorData 查询桌面监控信息
//
// 该接口可获取某一计算机在一段时间段(范围：1小时到30天)的数据信息（例如CPU占用率、内存占用率、用户的在线时间段等），最长数据保存时间不能超过180天。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowDesktopMonitorData(request *model.ShowDesktopMonitorDataRequest) (*model.ShowDesktopMonitorDataResponse, error) {
	requestDef := GenReqDefForShowDesktopMonitorData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDesktopMonitorDataResponse), nil
	}
}

// ShowDesktopMonitorDataInvoker 查询桌面监控信息
func (c *WorkspaceClient) ShowDesktopMonitorDataInvoker(request *model.ShowDesktopMonitorDataRequest) *ShowDesktopMonitorDataInvoker {
	requestDef := GenReqDefForShowDesktopMonitorData()
	return &ShowDesktopMonitorDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDesktopNetwork 查询桌面网络
//
// 查询桌面vpc、子网、privateIp、EIP、安全组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowDesktopNetwork(request *model.ShowDesktopNetworkRequest) (*model.ShowDesktopNetworkResponse, error) {
	requestDef := GenReqDefForShowDesktopNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDesktopNetworkResponse), nil
	}
}

// ShowDesktopNetworkInvoker 查询桌面网络
func (c *WorkspaceClient) ShowDesktopNetworkInvoker(request *model.ShowDesktopNetworkRequest) *ShowDesktopNetworkInvoker {
	requestDef := GenReqDefForShowDesktopNetwork()
	return &ShowDesktopNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDesktopNetworks 批量查询桌面网络
//
// 查询桌面vpc、子网、privateIp、EIP、安全组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowDesktopNetworks(request *model.ShowDesktopNetworksRequest) (*model.ShowDesktopNetworksResponse, error) {
	requestDef := GenReqDefForShowDesktopNetworks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDesktopNetworksResponse), nil
	}
}

// ShowDesktopNetworksInvoker 批量查询桌面网络
func (c *WorkspaceClient) ShowDesktopNetworksInvoker(request *model.ShowDesktopNetworksRequest) *ShowDesktopNetworksInvoker {
	requestDef := GenReqDefForShowDesktopNetworks()
	return &ShowDesktopNetworksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDesktopRemoteAssistanceInfo 根据桌面id查询远程协助信息
//
// 根据桌面id查询远程协助信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowDesktopRemoteAssistanceInfo(request *model.ShowDesktopRemoteAssistanceInfoRequest) (*model.ShowDesktopRemoteAssistanceInfoResponse, error) {
	requestDef := GenReqDefForShowDesktopRemoteAssistanceInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDesktopRemoteAssistanceInfoResponse), nil
	}
}

// ShowDesktopRemoteAssistanceInfoInvoker 根据桌面id查询远程协助信息
func (c *WorkspaceClient) ShowDesktopRemoteAssistanceInfoInvoker(request *model.ShowDesktopRemoteAssistanceInfoRequest) *ShowDesktopRemoteAssistanceInfoInvoker {
	requestDef := GenReqDefForShowDesktopRemoteAssistanceInfo()
	return &ShowDesktopRemoteAssistanceInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRemoteConsoleAddress 远程登录控制台
//
// 用于直接获取远程登录控制台地址。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowRemoteConsoleAddress(request *model.ShowRemoteConsoleAddressRequest) (*model.ShowRemoteConsoleAddressResponse, error) {
	requestDef := GenReqDefForShowRemoteConsoleAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRemoteConsoleAddressResponse), nil
	}
}

// ShowRemoteConsoleAddressInvoker 远程登录控制台
func (c *WorkspaceClient) ShowRemoteConsoleAddressInvoker(request *model.ShowRemoteConsoleAddressRequest) *ShowRemoteConsoleAddressInvoker {
	requestDef := GenReqDefForShowRemoteConsoleAddress()
	return &ShowRemoteConsoleAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSysprepInfo 查询sysprep版本信息
//
// 查询sysprep版本信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowSysprepInfo(request *model.ShowSysprepInfoRequest) (*model.ShowSysprepInfoResponse, error) {
	requestDef := GenReqDefForShowSysprepInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSysprepInfoResponse), nil
	}
}

// ShowSysprepInfoInvoker 查询sysprep版本信息
func (c *WorkspaceClient) ShowSysprepInfoInvoker(request *model.ShowSysprepInfoRequest) *ShowSysprepInfoInvoker {
	requestDef := GenReqDefForShowSysprepInfo()
	return &ShowSysprepInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDesktop 修改桌面属性
//
// 修改桌面属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateDesktop(request *model.UpdateDesktopRequest) (*model.UpdateDesktopResponse, error) {
	requestDef := GenReqDefForUpdateDesktop()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDesktopResponse), nil
	}
}

// UpdateDesktopInvoker 修改桌面属性
func (c *WorkspaceClient) UpdateDesktopInvoker(request *model.UpdateDesktopRequest) *UpdateDesktopInvoker {
	requestDef := GenReqDefForUpdateDesktop()
	return &UpdateDesktopInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDesktopSids 更新桌面SID
//
// 该接口用于桌面sid和WindowsAD中的SID不同时，更新桌面SID使其与AD中的SID保持一致，一般用于解决桌面脱域的情况使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateDesktopSids(request *model.UpdateDesktopSidsRequest) (*model.UpdateDesktopSidsResponse, error) {
	requestDef := GenReqDefForUpdateDesktopSids()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDesktopSidsResponse), nil
	}
}

// UpdateDesktopSidsInvoker 更新桌面SID
func (c *WorkspaceClient) UpdateDesktopSidsInvoker(request *model.UpdateDesktopSidsRequest) *UpdateDesktopSidsInvoker {
	requestDef := GenReqDefForUpdateDesktopSids()
	return &UpdateDesktopSidsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDesktopUsername AD场景支持桌面更换关联用户名
//
// AD场景支持桌面更换关联用户名。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateDesktopUsername(request *model.UpdateDesktopUsernameRequest) (*model.UpdateDesktopUsernameResponse, error) {
	requestDef := GenReqDefForUpdateDesktopUsername()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDesktopUsernameResponse), nil
	}
}

// UpdateDesktopUsernameInvoker AD场景支持桌面更换关联用户名
func (c *WorkspaceClient) UpdateDesktopUsernameInvoker(request *model.UpdateDesktopUsernameRequest) *UpdateDesktopUsernameInvoker {
	requestDef := GenReqDefForUpdateDesktopUsername()
	return &UpdateDesktopUsernameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteDesktopNamePolicy 批量删除桌面名称策略
//
// 批量删除桌面名称策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchDeleteDesktopNamePolicy(request *model.BatchDeleteDesktopNamePolicyRequest) (*model.BatchDeleteDesktopNamePolicyResponse, error) {
	requestDef := GenReqDefForBatchDeleteDesktopNamePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteDesktopNamePolicyResponse), nil
	}
}

// BatchDeleteDesktopNamePolicyInvoker 批量删除桌面名称策略
func (c *WorkspaceClient) BatchDeleteDesktopNamePolicyInvoker(request *model.BatchDeleteDesktopNamePolicyRequest) *BatchDeleteDesktopNamePolicyInvoker {
	requestDef := GenReqDefForBatchDeleteDesktopNamePolicy()
	return &BatchDeleteDesktopNamePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDesktopNamePolicy 创建桌面名称策略
//
// 创建桌面名称策略，用于自动生成桌面名称，最多允许50个。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreateDesktopNamePolicy(request *model.CreateDesktopNamePolicyRequest) (*model.CreateDesktopNamePolicyResponse, error) {
	requestDef := GenReqDefForCreateDesktopNamePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDesktopNamePolicyResponse), nil
	}
}

// CreateDesktopNamePolicyInvoker 创建桌面名称策略
func (c *WorkspaceClient) CreateDesktopNamePolicyInvoker(request *model.CreateDesktopNamePolicyRequest) *CreateDesktopNamePolicyInvoker {
	requestDef := GenReqDefForCreateDesktopNamePolicy()
	return &CreateDesktopNamePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDesktopNamePolicy 获取桌面名称策略
//
// 获取桌面名称策略，用于自动生成桌面名称。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListDesktopNamePolicy(request *model.ListDesktopNamePolicyRequest) (*model.ListDesktopNamePolicyResponse, error) {
	requestDef := GenReqDefForListDesktopNamePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDesktopNamePolicyResponse), nil
	}
}

// ListDesktopNamePolicyInvoker 获取桌面名称策略
func (c *WorkspaceClient) ListDesktopNamePolicyInvoker(request *model.ListDesktopNamePolicyRequest) *ListDesktopNamePolicyInvoker {
	requestDef := GenReqDefForListDesktopNamePolicy()
	return &ListDesktopNamePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDesktopNamePolicy 更新桌面名称策略
//
// 更新桌面名称策略，用于自动生成桌面名称。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateDesktopNamePolicy(request *model.UpdateDesktopNamePolicyRequest) (*model.UpdateDesktopNamePolicyResponse, error) {
	requestDef := GenReqDefForUpdateDesktopNamePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDesktopNamePolicyResponse), nil
	}
}

// UpdateDesktopNamePolicyInvoker 更新桌面名称策略
func (c *WorkspaceClient) UpdateDesktopNamePolicyInvoker(request *model.UpdateDesktopNamePolicyRequest) *UpdateDesktopNamePolicyInvoker {
	requestDef := GenReqDefForUpdateDesktopNamePolicy()
	return &UpdateDesktopNamePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddDesktopPoolVolumes 桌面池批量添加磁盘
//
// 桌面池批量添加磁盘。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) AddDesktopPoolVolumes(request *model.AddDesktopPoolVolumesRequest) (*model.AddDesktopPoolVolumesResponse, error) {
	requestDef := GenReqDefForAddDesktopPoolVolumes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddDesktopPoolVolumesResponse), nil
	}
}

// AddDesktopPoolVolumesInvoker 桌面池批量添加磁盘
func (c *WorkspaceClient) AddDesktopPoolVolumesInvoker(request *model.AddDesktopPoolVolumesRequest) *AddDesktopPoolVolumesInvoker {
	requestDef := GenReqDefForAddDesktopPoolVolumes()
	return &AddDesktopPoolVolumesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDesktopPool 创建桌面池
//
// 创建桌面池，可将此桌面池分配给用户、用户组，用户登录时会绑定其中一个桌面。
// 注:需通过开通委托功能接口先对云服务进行授权才可以使用该功能
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreateDesktopPool(request *model.CreateDesktopPoolRequest) (*model.CreateDesktopPoolResponse, error) {
	requestDef := GenReqDefForCreateDesktopPool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDesktopPoolResponse), nil
	}
}

// CreateDesktopPoolInvoker 创建桌面池
func (c *WorkspaceClient) CreateDesktopPoolInvoker(request *model.CreateDesktopPoolRequest) *CreateDesktopPoolInvoker {
	requestDef := GenReqDefForCreateDesktopPool()
	return &CreateDesktopPoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDesktopPoolAuthorizedObjects 桌面池授权用户、用户组
//
// 该接口用于桌面池授权用户、用户组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreateDesktopPoolAuthorizedObjects(request *model.CreateDesktopPoolAuthorizedObjectsRequest) (*model.CreateDesktopPoolAuthorizedObjectsResponse, error) {
	requestDef := GenReqDefForCreateDesktopPoolAuthorizedObjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDesktopPoolAuthorizedObjectsResponse), nil
	}
}

// CreateDesktopPoolAuthorizedObjectsInvoker 桌面池授权用户、用户组
func (c *WorkspaceClient) CreateDesktopPoolAuthorizedObjectsInvoker(request *model.CreateDesktopPoolAuthorizedObjectsRequest) *CreateDesktopPoolAuthorizedObjectsInvoker {
	requestDef := GenReqDefForCreateDesktopPoolAuthorizedObjects()
	return &CreateDesktopPoolAuthorizedObjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDesktopPool 删除桌面池
//
// 当桌面池内无桌面时可删除桌面池，桌面池删除后无法恢复。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) DeleteDesktopPool(request *model.DeleteDesktopPoolRequest) (*model.DeleteDesktopPoolResponse, error) {
	requestDef := GenReqDefForDeleteDesktopPool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDesktopPoolResponse), nil
	}
}

// DeleteDesktopPoolInvoker 删除桌面池
func (c *WorkspaceClient) DeleteDesktopPoolInvoker(request *model.DeleteDesktopPoolRequest) *DeleteDesktopPoolInvoker {
	requestDef := GenReqDefForDeleteDesktopPool()
	return &DeleteDesktopPoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDesktopPoolVolumes 桌面池批量删除磁盘
//
// 桌面池批量删除磁盘。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) DeleteDesktopPoolVolumes(request *model.DeleteDesktopPoolVolumesRequest) (*model.DeleteDesktopPoolVolumesResponse, error) {
	requestDef := GenReqDefForDeleteDesktopPoolVolumes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDesktopPoolVolumesResponse), nil
	}
}

// DeleteDesktopPoolVolumesInvoker 桌面池批量删除磁盘
func (c *WorkspaceClient) DeleteDesktopPoolVolumesInvoker(request *model.DeleteDesktopPoolVolumesRequest) *DeleteDesktopPoolVolumesInvoker {
	requestDef := GenReqDefForDeleteDesktopPoolVolumes()
	return &DeleteDesktopPoolVolumesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteDesktopPoolAction 操作桌面池
//
// 操作桌面池，用于桌面池里面的桌面批量开机、关机、重启和休眠。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ExecuteDesktopPoolAction(request *model.ExecuteDesktopPoolActionRequest) (*model.ExecuteDesktopPoolActionResponse, error) {
	requestDef := GenReqDefForExecuteDesktopPoolAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteDesktopPoolActionResponse), nil
	}
}

// ExecuteDesktopPoolActionInvoker 操作桌面池
func (c *WorkspaceClient) ExecuteDesktopPoolActionInvoker(request *model.ExecuteDesktopPoolActionRequest) *ExecuteDesktopPoolActionInvoker {
	requestDef := GenReqDefForExecuteDesktopPoolAction()
	return &ExecuteDesktopPoolActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteDesktopPoolScript 桌面池批量执行脚本
//
// 桌面池批量执行脚本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ExecuteDesktopPoolScript(request *model.ExecuteDesktopPoolScriptRequest) (*model.ExecuteDesktopPoolScriptResponse, error) {
	requestDef := GenReqDefForExecuteDesktopPoolScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteDesktopPoolScriptResponse), nil
	}
}

// ExecuteDesktopPoolScriptInvoker 桌面池批量执行脚本
func (c *WorkspaceClient) ExecuteDesktopPoolScriptInvoker(request *model.ExecuteDesktopPoolScriptRequest) *ExecuteDesktopPoolScriptInvoker {
	requestDef := GenReqDefForExecuteDesktopPoolScript()
	return &ExecuteDesktopPoolScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExpandDesktopPool 扩容桌面池
//
// 扩容桌面池。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ExpandDesktopPool(request *model.ExpandDesktopPoolRequest) (*model.ExpandDesktopPoolResponse, error) {
	requestDef := GenReqDefForExpandDesktopPool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandDesktopPoolResponse), nil
	}
}

// ExpandDesktopPoolInvoker 扩容桌面池
func (c *WorkspaceClient) ExpandDesktopPoolInvoker(request *model.ExpandDesktopPoolRequest) *ExpandDesktopPoolInvoker {
	requestDef := GenReqDefForExpandDesktopPool()
	return &ExpandDesktopPoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExpandDesktopPoolVolumes 桌面池批量扩容磁盘
//
// 桌面池批量扩容磁盘。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ExpandDesktopPoolVolumes(request *model.ExpandDesktopPoolVolumesRequest) (*model.ExpandDesktopPoolVolumesResponse, error) {
	requestDef := GenReqDefForExpandDesktopPoolVolumes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandDesktopPoolVolumesResponse), nil
	}
}

// ExpandDesktopPoolVolumesInvoker 桌面池批量扩容磁盘
func (c *WorkspaceClient) ExpandDesktopPoolVolumesInvoker(request *model.ExpandDesktopPoolVolumesRequest) *ExpandDesktopPoolVolumesInvoker {
	requestDef := GenReqDefForExpandDesktopPoolVolumes()
	return &ExpandDesktopPoolVolumesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDesktopPoolAuthorizedObjects 查询桌面池授权的用户、用户组
//
// 该接口用于查询指定桌面池授权的用户、用户组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListDesktopPoolAuthorizedObjects(request *model.ListDesktopPoolAuthorizedObjectsRequest) (*model.ListDesktopPoolAuthorizedObjectsResponse, error) {
	requestDef := GenReqDefForListDesktopPoolAuthorizedObjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDesktopPoolAuthorizedObjectsResponse), nil
	}
}

// ListDesktopPoolAuthorizedObjectsInvoker 查询桌面池授权的用户、用户组
func (c *WorkspaceClient) ListDesktopPoolAuthorizedObjectsInvoker(request *model.ListDesktopPoolAuthorizedObjectsRequest) *ListDesktopPoolAuthorizedObjectsInvoker {
	requestDef := GenReqDefForListDesktopPoolAuthorizedObjects()
	return &ListDesktopPoolAuthorizedObjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDesktopPools 查询桌面池列表
//
// 该接口用于查询桌面池列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListDesktopPools(request *model.ListDesktopPoolsRequest) (*model.ListDesktopPoolsResponse, error) {
	requestDef := GenReqDefForListDesktopPools()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDesktopPoolsResponse), nil
	}
}

// ListDesktopPoolsInvoker 查询桌面池列表
func (c *WorkspaceClient) ListDesktopPoolsInvoker(request *model.ListDesktopPoolsRequest) *ListDesktopPoolsInvoker {
	requestDef := GenReqDefForListDesktopPools()
	return &ListDesktopPoolsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPoolDesktopsDetail 查询桌面池下的桌面信息
//
// 该接口用于查询桌面池下的桌面信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListPoolDesktopsDetail(request *model.ListPoolDesktopsDetailRequest) (*model.ListPoolDesktopsDetailResponse, error) {
	requestDef := GenReqDefForListPoolDesktopsDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPoolDesktopsDetailResponse), nil
	}
}

// ListPoolDesktopsDetailInvoker 查询桌面池下的桌面信息
func (c *WorkspaceClient) ListPoolDesktopsDetailInvoker(request *model.ListPoolDesktopsDetailRequest) *ListPoolDesktopsDetailInvoker {
	requestDef := GenReqDefForListPoolDesktopsDetail()
	return &ListPoolDesktopsDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RebuildDesktopPool 桌面池重建系统盘
//
// 桌面池重建系统盘。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) RebuildDesktopPool(request *model.RebuildDesktopPoolRequest) (*model.RebuildDesktopPoolResponse, error) {
	requestDef := GenReqDefForRebuildDesktopPool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RebuildDesktopPoolResponse), nil
	}
}

// RebuildDesktopPoolInvoker 桌面池重建系统盘
func (c *WorkspaceClient) RebuildDesktopPoolInvoker(request *model.RebuildDesktopPoolRequest) *RebuildDesktopPoolInvoker {
	requestDef := GenReqDefForRebuildDesktopPool()
	return &RebuildDesktopPoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeDesktopPool 桌面池变更规格
//
// 桌面池变更规格
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ResizeDesktopPool(request *model.ResizeDesktopPoolRequest) (*model.ResizeDesktopPoolResponse, error) {
	requestDef := GenReqDefForResizeDesktopPool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeDesktopPoolResponse), nil
	}
}

// ResizeDesktopPoolInvoker 桌面池变更规格
func (c *WorkspaceClient) ResizeDesktopPoolInvoker(request *model.ResizeDesktopPoolRequest) *ResizeDesktopPoolInvoker {
	requestDef := GenReqDefForResizeDesktopPool()
	return &ResizeDesktopPoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendDesktopPoolNotifications 发送消息通知
//
// 用于管理员向桌面发送消息通知
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) SendDesktopPoolNotifications(request *model.SendDesktopPoolNotificationsRequest) (*model.SendDesktopPoolNotificationsResponse, error) {
	requestDef := GenReqDefForSendDesktopPoolNotifications()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendDesktopPoolNotificationsResponse), nil
	}
}

// SendDesktopPoolNotificationsInvoker 发送消息通知
func (c *WorkspaceClient) SendDesktopPoolNotificationsInvoker(request *model.SendDesktopPoolNotificationsRequest) *SendDesktopPoolNotificationsInvoker {
	requestDef := GenReqDefForSendDesktopPoolNotifications()
	return &SendDesktopPoolNotificationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDesktopPoolDetail 查询桌面池详情
//
// 指定桌面池Id查询详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowDesktopPoolDetail(request *model.ShowDesktopPoolDetailRequest) (*model.ShowDesktopPoolDetailResponse, error) {
	requestDef := GenReqDefForShowDesktopPoolDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDesktopPoolDetailResponse), nil
	}
}

// ShowDesktopPoolDetailInvoker 查询桌面池详情
func (c *WorkspaceClient) ShowDesktopPoolDetailInvoker(request *model.ShowDesktopPoolDetailRequest) *ShowDesktopPoolDetailInvoker {
	requestDef := GenReqDefForShowDesktopPoolDetail()
	return &ShowDesktopPoolDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDesktopPoolsScriptExecTasks 查询桌面池的脚本执行任务列表
//
// 桌面池的脚本执行任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowDesktopPoolsScriptExecTasks(request *model.ShowDesktopPoolsScriptExecTasksRequest) (*model.ShowDesktopPoolsScriptExecTasksResponse, error) {
	requestDef := GenReqDefForShowDesktopPoolsScriptExecTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDesktopPoolsScriptExecTasksResponse), nil
	}
}

// ShowDesktopPoolsScriptExecTasksInvoker 查询桌面池的脚本执行任务列表
func (c *WorkspaceClient) ShowDesktopPoolsScriptExecTasksInvoker(request *model.ShowDesktopPoolsScriptExecTasksRequest) *ShowDesktopPoolsScriptExecTasksInvoker {
	requestDef := GenReqDefForShowDesktopPoolsScriptExecTasks()
	return &ShowDesktopPoolsScriptExecTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDesktopPool 修改桌面池属性
//
// 修改桌面池属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateDesktopPool(request *model.UpdateDesktopPoolRequest) (*model.UpdateDesktopPoolResponse, error) {
	requestDef := GenReqDefForUpdateDesktopPool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDesktopPoolResponse), nil
	}
}

// UpdateDesktopPoolInvoker 修改桌面池属性
func (c *WorkspaceClient) UpdateDesktopPoolInvoker(request *model.UpdateDesktopPoolRequest) *UpdateDesktopPoolInvoker {
	requestDef := GenReqDefForUpdateDesktopPool()
	return &UpdateDesktopPoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAddDesktopsTags 批量添加多个桌面标签
//
// 同时对多个桌面批量添加标签，如果创建的标签已经存在（key相同）则覆，最大支持100个桌面，每个桌面最大20个标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchAddDesktopsTags(request *model.BatchAddDesktopsTagsRequest) (*model.BatchAddDesktopsTagsResponse, error) {
	requestDef := GenReqDefForBatchAddDesktopsTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddDesktopsTagsResponse), nil
	}
}

// BatchAddDesktopsTagsInvoker 批量添加多个桌面标签
func (c *WorkspaceClient) BatchAddDesktopsTagsInvoker(request *model.BatchAddDesktopsTagsRequest) *BatchAddDesktopsTagsInvoker {
	requestDef := GenReqDefForBatchAddDesktopsTags()
	return &BatchAddDesktopsTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchChangeTags 批量添加删除标签
//
// 为指定桌面批量添加或删除标签。创建时，如果创建的标签已经存在（key相同），则覆盖。删除时，如果删除的标签不存在，默认处理成功
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchChangeTags(request *model.BatchChangeTagsRequest) (*model.BatchChangeTagsResponse, error) {
	requestDef := GenReqDefForBatchChangeTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchChangeTagsResponse), nil
	}
}

// BatchChangeTagsInvoker 批量添加删除标签
func (c *WorkspaceClient) BatchChangeTagsInvoker(request *model.BatchChangeTagsRequest) *BatchChangeTagsInvoker {
	requestDef := GenReqDefForBatchChangeTags()
	return &BatchChangeTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteDesktopsTags 批量删除多个桌面标签
//
// 同时对多个桌面批量添加标签，删除时，如果删除的标签不存在默认处理成功，最大支持100个桌面，每个桌面最大20个标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchDeleteDesktopsTags(request *model.BatchDeleteDesktopsTagsRequest) (*model.BatchDeleteDesktopsTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteDesktopsTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteDesktopsTagsResponse), nil
	}
}

// BatchDeleteDesktopsTagsInvoker 批量删除多个桌面标签
func (c *WorkspaceClient) BatchDeleteDesktopsTagsInvoker(request *model.BatchDeleteDesktopsTagsRequest) *BatchDeleteDesktopsTagsInvoker {
	requestDef := GenReqDefForBatchDeleteDesktopsTags()
	return &BatchDeleteDesktopsTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTag 创建桌面标签
//
// 该接口用于为桌面创建标签，一个桌面上最多有10个标签。创建时，如果创建的标签已经存在（key相同），则覆盖。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreateTag(request *model.CreateTagRequest) (*model.CreateTagResponse, error) {
	requestDef := GenReqDefForCreateTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTagResponse), nil
	}
}

// CreateTagInvoker 创建桌面标签
func (c *WorkspaceClient) CreateTagInvoker(request *model.CreateTagRequest) *CreateTagInvoker {
	requestDef := GenReqDefForCreateTag()
	return &CreateTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTag 删除桌面标签
//
// 该接口用于删除桌面标签。删除时，如果删除的标签不存在，默认处理成功。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) DeleteTag(request *model.DeleteTagRequest) (*model.DeleteTagResponse, error) {
	requestDef := GenReqDefForDeleteTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTagResponse), nil
	}
}

// DeleteTagInvoker 删除桌面标签
func (c *WorkspaceClient) DeleteTagInvoker(request *model.DeleteTagRequest) *DeleteTagInvoker {
	requestDef := GenReqDefForDeleteTag()
	return &DeleteTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDesktopByTags 使用标签过滤桌面
//
// 使用标签过滤桌面
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListDesktopByTags(request *model.ListDesktopByTagsRequest) (*model.ListDesktopByTagsResponse, error) {
	requestDef := GenReqDefForListDesktopByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDesktopByTagsResponse), nil
	}
}

// ListDesktopByTagsInvoker 使用标签过滤桌面
func (c *WorkspaceClient) ListDesktopByTagsInvoker(request *model.ListDesktopByTagsRequest) *ListDesktopByTagsInvoker {
	requestDef := GenReqDefForListDesktopByTags()
	return &ListDesktopByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectTags 查询项目标签
//
// 查询租户的所有标签集合
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListProjectTags(request *model.ListProjectTagsRequest) (*model.ListProjectTagsResponse, error) {
	requestDef := GenReqDefForListProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectTagsResponse), nil
	}
}

// ListProjectTagsInvoker 查询项目标签
func (c *WorkspaceClient) ListProjectTagsInvoker(request *model.ListProjectTagsRequest) *ListProjectTagsInvoker {
	requestDef := GenReqDefForListProjectTags()
	return &ListProjectTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTagByDesktopId 查询桌面标签
//
// 查询指定桌面的标签信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowTagByDesktopId(request *model.ShowTagByDesktopIdRequest) (*model.ShowTagByDesktopIdResponse, error) {
	requestDef := GenReqDefForShowTagByDesktopId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTagByDesktopIdResponse), nil
	}
}

// ShowTagByDesktopIdInvoker 查询桌面标签
func (c *WorkspaceClient) ShowTagByDesktopIdInvoker(request *model.ShowTagByDesktopIdRequest) *ShowTagByDesktopIdInvoker {
	requestDef := GenReqDefForShowTagByDesktopId()
	return &ShowTagByDesktopIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteUserGroups 批量删除用户组
//
// 该接口用于批量删除用户组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchDeleteUserGroups(request *model.BatchDeleteUserGroupsRequest) (*model.BatchDeleteUserGroupsResponse, error) {
	requestDef := GenReqDefForBatchDeleteUserGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteUserGroupsResponse), nil
	}
}

// BatchDeleteUserGroupsInvoker 批量删除用户组
func (c *WorkspaceClient) BatchDeleteUserGroupsInvoker(request *model.BatchDeleteUserGroupsRequest) *BatchDeleteUserGroupsInvoker {
	requestDef := GenReqDefForBatchDeleteUserGroups()
	return &BatchDeleteUserGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateUserGroup 创建用户组
//
// 该接口用于创建用户组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreateUserGroup(request *model.CreateUserGroupRequest) (*model.CreateUserGroupResponse, error) {
	requestDef := GenReqDefForCreateUserGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateUserGroupResponse), nil
	}
}

// CreateUserGroupInvoker 创建用户组
func (c *WorkspaceClient) CreateUserGroupInvoker(request *model.CreateUserGroupRequest) *CreateUserGroupInvoker {
	requestDef := GenReqDefForCreateUserGroup()
	return &CreateUserGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteUserGroup 删除用户组
//
// 删除用户组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) DeleteUserGroup(request *model.DeleteUserGroupRequest) (*model.DeleteUserGroupResponse, error) {
	requestDef := GenReqDefForDeleteUserGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteUserGroupResponse), nil
	}
}

// DeleteUserGroupInvoker 删除用户组
func (c *WorkspaceClient) DeleteUserGroupInvoker(request *model.DeleteUserGroupRequest) *DeleteUserGroupInvoker {
	requestDef := GenReqDefForDeleteUserGroup()
	return &DeleteUserGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserGroups 查询用户组列表
//
// 查询用户组列表，支持分页。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListUserGroups(request *model.ListUserGroupsRequest) (*model.ListUserGroupsResponse, error) {
	requestDef := GenReqDefForListUserGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserGroupsResponse), nil
	}
}

// ListUserGroupsInvoker 查询用户组列表
func (c *WorkspaceClient) ListUserGroupsInvoker(request *model.ListUserGroupsRequest) *ListUserGroupsInvoker {
	requestDef := GenReqDefForListUserGroups()
	return &ListUserGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUsersOfGroup 查询用户组中的用户
//
// 该接口用于查询用户组中的用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListUsersOfGroup(request *model.ListUsersOfGroupRequest) (*model.ListUsersOfGroupResponse, error) {
	requestDef := GenReqDefForListUsersOfGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsersOfGroupResponse), nil
	}
}

// ListUsersOfGroupInvoker 查询用户组中的用户
func (c *WorkspaceClient) ListUsersOfGroupInvoker(request *model.ListUsersOfGroupRequest) *ListUsersOfGroupInvoker {
	requestDef := GenReqDefForListUsersOfGroup()
	return &ListUsersOfGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunActionsOnGroup 操作用户组
//
// 操作用户组，如添加用户、删除用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) RunActionsOnGroup(request *model.RunActionsOnGroupRequest) (*model.RunActionsOnGroupResponse, error) {
	requestDef := GenReqDefForRunActionsOnGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunActionsOnGroupResponse), nil
	}
}

// RunActionsOnGroupInvoker 操作用户组
func (c *WorkspaceClient) RunActionsOnGroupInvoker(request *model.RunActionsOnGroupRequest) *RunActionsOnGroupInvoker {
	requestDef := GenReqDefForRunActionsOnGroup()
	return &RunActionsOnGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUserGroup 修改用户组信息
//
// 该接口用于修改用户组信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateUserGroup(request *model.UpdateUserGroupRequest) (*model.UpdateUserGroupResponse, error) {
	requestDef := GenReqDefForUpdateUserGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserGroupResponse), nil
	}
}

// UpdateUserGroupInvoker 修改用户组信息
func (c *WorkspaceClient) UpdateUserGroupInvoker(request *model.UpdateUserGroupRequest) *UpdateUserGroupInvoker {
	requestDef := GenReqDefForUpdateUserGroup()
	return &UpdateUserGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImages 查询产品镜像列表
//
// 该接口用于查询云桌面支持的产品镜像列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListImages(request *model.ListImagesRequest) (*model.ListImagesResponse, error) {
	requestDef := GenReqDefForListImages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImagesResponse), nil
	}
}

// ListImagesInvoker 查询产品镜像列表
func (c *WorkspaceClient) ListImagesInvoker(request *model.ListImagesRequest) *ListImagesInvoker {
	requestDef := GenReqDefForListImages()
	return &ListImagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMarketImages 获取云市场镜像列表
//
// 获取云市场镜像列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListMarketImages(request *model.ListMarketImagesRequest) (*model.ListMarketImagesResponse, error) {
	requestDef := GenReqDefForListMarketImages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMarketImagesResponse), nil
	}
}

// ListMarketImagesInvoker 获取云市场镜像列表
func (c *WorkspaceClient) ListMarketImagesInvoker(request *model.ListMarketImagesRequest) *ListMarketImagesInvoker {
	requestDef := GenReqDefForListMarketImages()
	return &ListMarketImagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EstimateAddResources 包周期桌面增配变更批量询价
//
// 包周期桌面增配变更批量询价
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) EstimateAddResources(request *model.EstimateAddResourcesRequest) (*model.EstimateAddResourcesResponse, error) {
	requestDef := GenReqDefForEstimateAddResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EstimateAddResourcesResponse), nil
	}
}

// EstimateAddResourcesInvoker 包周期桌面增配变更批量询价
func (c *WorkspaceClient) EstimateAddResourcesInvoker(request *model.EstimateAddResourcesRequest) *EstimateAddResourcesInvoker {
	requestDef := GenReqDefForEstimateAddResources()
	return &EstimateAddResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EstimateChangeImages 批量包周期桌面切换镜像询价
//
// 批量包周期桌面切换镜像(由不收费镜像变更至收费镜像)询价。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) EstimateChangeImages(request *model.EstimateChangeImagesRequest) (*model.EstimateChangeImagesResponse, error) {
	requestDef := GenReqDefForEstimateChangeImages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EstimateChangeImagesResponse), nil
	}
}

// EstimateChangeImagesInvoker 批量包周期桌面切换镜像询价
func (c *WorkspaceClient) EstimateChangeImagesInvoker(request *model.EstimateChangeImagesRequest) *EstimateChangeImagesInvoker {
	requestDef := GenReqDefForEstimateChangeImages()
	return &EstimateChangeImagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EstimateDesktopPoolAddVolume 包周期桌面池添加单个磁盘批量询价
//
// 包周期桌面池添加单个磁盘批量询价
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) EstimateDesktopPoolAddVolume(request *model.EstimateDesktopPoolAddVolumeRequest) (*model.EstimateDesktopPoolAddVolumeResponse, error) {
	requestDef := GenReqDefForEstimateDesktopPoolAddVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EstimateDesktopPoolAddVolumeResponse), nil
	}
}

// EstimateDesktopPoolAddVolumeInvoker 包周期桌面池添加单个磁盘批量询价
func (c *WorkspaceClient) EstimateDesktopPoolAddVolumeInvoker(request *model.EstimateDesktopPoolAddVolumeRequest) *EstimateDesktopPoolAddVolumeInvoker {
	requestDef := GenReqDefForEstimateDesktopPoolAddVolume()
	return &EstimateDesktopPoolAddVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EstimateDesktopPoolChangeImage 包周期桌面池切换镜像批量询价
//
// 包周期桌面池切换镜像(由不收费镜像变更至收费镜像)批量询价
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) EstimateDesktopPoolChangeImage(request *model.EstimateDesktopPoolChangeImageRequest) (*model.EstimateDesktopPoolChangeImageResponse, error) {
	requestDef := GenReqDefForEstimateDesktopPoolChangeImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EstimateDesktopPoolChangeImageResponse), nil
	}
}

// EstimateDesktopPoolChangeImageInvoker 包周期桌面池切换镜像批量询价
func (c *WorkspaceClient) EstimateDesktopPoolChangeImageInvoker(request *model.EstimateDesktopPoolChangeImageRequest) *EstimateDesktopPoolChangeImageInvoker {
	requestDef := GenReqDefForEstimateDesktopPoolChangeImage()
	return &EstimateDesktopPoolChangeImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EstimateDesktopPoolExtendVolume 包周期桌面池扩容磁盘批量询价
//
// 包周期桌面池扩容磁盘批量询价
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) EstimateDesktopPoolExtendVolume(request *model.EstimateDesktopPoolExtendVolumeRequest) (*model.EstimateDesktopPoolExtendVolumeResponse, error) {
	requestDef := GenReqDefForEstimateDesktopPoolExtendVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EstimateDesktopPoolExtendVolumeResponse), nil
	}
}

// EstimateDesktopPoolExtendVolumeInvoker 包周期桌面池扩容磁盘批量询价
func (c *WorkspaceClient) EstimateDesktopPoolExtendVolumeInvoker(request *model.EstimateDesktopPoolExtendVolumeRequest) *EstimateDesktopPoolExtendVolumeInvoker {
	requestDef := GenReqDefForEstimateDesktopPoolExtendVolume()
	return &EstimateDesktopPoolExtendVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EstimateDesktopPoolResize 包周期桌面池变更规格批量询价
//
// 包周期桌面池变更规格批量询价
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) EstimateDesktopPoolResize(request *model.EstimateDesktopPoolResizeRequest) (*model.EstimateDesktopPoolResizeResponse, error) {
	requestDef := GenReqDefForEstimateDesktopPoolResize()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EstimateDesktopPoolResizeResponse), nil
	}
}

// EstimateDesktopPoolResizeInvoker 包周期桌面池变更规格批量询价
func (c *WorkspaceClient) EstimateDesktopPoolResizeInvoker(request *model.EstimateDesktopPoolResizeRequest) *EstimateDesktopPoolResizeInvoker {
	requestDef := GenReqDefForEstimateDesktopPoolResize()
	return &EstimateDesktopPoolResizeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunActionsOnWorkspaceJob 重试任务
//
// 该接口用来对失败的任务进行重试，当前仅支持开户和销户的任务重试。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) RunActionsOnWorkspaceJob(request *model.RunActionsOnWorkspaceJobRequest) (*model.RunActionsOnWorkspaceJobResponse, error) {
	requestDef := GenReqDefForRunActionsOnWorkspaceJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunActionsOnWorkspaceJobResponse), nil
	}
}

// RunActionsOnWorkspaceJobInvoker 重试任务
func (c *WorkspaceClient) RunActionsOnWorkspaceJobInvoker(request *model.RunActionsOnWorkspaceJobRequest) *RunActionsOnWorkspaceJobInvoker {
	requestDef := GenReqDefForRunActionsOnWorkspaceJob()
	return &RunActionsOnWorkspaceJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteSubJobs 删除子任务
//
// 该接口用于删除子任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchDeleteSubJobs(request *model.BatchDeleteSubJobsRequest) (*model.BatchDeleteSubJobsResponse, error) {
	requestDef := GenReqDefForBatchDeleteSubJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteSubJobsResponse), nil
	}
}

// BatchDeleteSubJobsInvoker 删除子任务
func (c *WorkspaceClient) BatchDeleteSubJobsInvoker(request *model.BatchDeleteSubJobsRequest) *BatchDeleteSubJobsInvoker {
	requestDef := GenReqDefForBatchDeleteSubJobs()
	return &BatchDeleteSubJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListItaSubJobs 子任务查询
//
// 该接口用于查询异步任务执行情况，比如查询创建桌面的执行状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListItaSubJobs(request *model.ListItaSubJobsRequest) (*model.ListItaSubJobsResponse, error) {
	requestDef := GenReqDefForListItaSubJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListItaSubJobsResponse), nil
	}
}

// ListItaSubJobsInvoker 子任务查询
func (c *WorkspaceClient) ListItaSubJobsInvoker(request *model.ListItaSubJobsRequest) *ListItaSubJobsInvoker {
	requestDef := GenReqDefForListItaSubJobs()
	return &ListItaSubJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJob 查询任务详情
//
// 该接口用于查询异步任务的执行情况，比如查询创建桌面任务的执行状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowJob(request *model.ShowJobRequest) (*model.ShowJobResponse, error) {
	requestDef := GenReqDefForShowJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobResponse), nil
	}
}

// ShowJobInvoker 查询任务详情
func (c *WorkspaceClient) ShowJobInvoker(request *model.ShowJobRequest) *ShowJobInvoker {
	requestDef := GenReqDefForShowJob()
	return &ShowJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNatMappingConfigs 查询租户的NAT映射配置项
//
// 查询租户的NAT映射配置项。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListNatMappingConfigs(request *model.ListNatMappingConfigsRequest) (*model.ListNatMappingConfigsResponse, error) {
	requestDef := GenReqDefForListNatMappingConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNatMappingConfigsResponse), nil
	}
}

// ListNatMappingConfigsInvoker 查询租户的NAT映射配置项
func (c *WorkspaceClient) ListNatMappingConfigsInvoker(request *model.ListNatMappingConfigsRequest) *ListNatMappingConfigsInvoker {
	requestDef := GenReqDefForListNatMappingConfigs()
	return &ListNatMappingConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNatMappingConfigs 修改租户的NAT映射配置项
//
// 修改租户的NAT映射配置项。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateNatMappingConfigs(request *model.UpdateNatMappingConfigsRequest) (*model.UpdateNatMappingConfigsResponse, error) {
	requestDef := GenReqDefForUpdateNatMappingConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNatMappingConfigsResponse), nil
	}
}

// UpdateNatMappingConfigsInvoker 修改租户的NAT映射配置项
func (c *WorkspaceClient) UpdateNatMappingConfigsInvoker(request *model.UpdateNatMappingConfigsRequest) *UpdateNatMappingConfigsInvoker {
	requestDef := GenReqDefForUpdateNatMappingConfigs()
	return &UpdateNatMappingConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ApplyDesktopsInternet 开通桌面上网功能
//
// 开通桌面上网功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ApplyDesktopsInternet(request *model.ApplyDesktopsInternetRequest) (*model.ApplyDesktopsInternetResponse, error) {
	requestDef := GenReqDefForApplyDesktopsInternet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ApplyDesktopsInternetResponse), nil
	}
}

// ApplyDesktopsInternetInvoker 开通桌面上网功能
func (c *WorkspaceClient) ApplyDesktopsInternetInvoker(request *model.ApplyDesktopsInternetRequest) *ApplyDesktopsInternetInvoker {
	requestDef := GenReqDefForApplyDesktopsInternet()
	return &ApplyDesktopsInternetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ApplyInternet 开通上网功能
//
// 开通上网功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ApplyInternet(request *model.ApplyInternetRequest) (*model.ApplyInternetResponse, error) {
	requestDef := GenReqDefForApplyInternet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ApplyInternetResponse), nil
	}
}

// ApplyInternetInvoker 开通上网功能
func (c *WorkspaceClient) ApplyInternetInvoker(request *model.ApplyInternetRequest) *ApplyInternetInvoker {
	requestDef := GenReqDefForApplyInternet()
	return &ApplyInternetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ApplySubnetBandwidth 创建云办公带宽
//
// 创建按需云办公带宽。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ApplySubnetBandwidth(request *model.ApplySubnetBandwidthRequest) (*model.ApplySubnetBandwidthResponse, error) {
	requestDef := GenReqDefForApplySubnetBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ApplySubnetBandwidthResponse), nil
	}
}

// ApplySubnetBandwidthInvoker 创建云办公带宽
func (c *WorkspaceClient) ApplySubnetBandwidthInvoker(request *model.ApplySubnetBandwidthRequest) *ApplySubnetBandwidthInvoker {
	requestDef := GenReqDefForApplySubnetBandwidth()
	return &ApplySubnetBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateDesktopsEip 桌面绑定EIP
//
// 桌面绑定EIP。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) AssociateDesktopsEip(request *model.AssociateDesktopsEipRequest) (*model.AssociateDesktopsEipResponse, error) {
	requestDef := GenReqDefForAssociateDesktopsEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateDesktopsEipResponse), nil
	}
}

// AssociateDesktopsEipInvoker 桌面绑定EIP
func (c *WorkspaceClient) AssociateDesktopsEipInvoker(request *model.AssociateDesktopsEipRequest) *AssociateDesktopsEipInvoker {
	requestDef := GenReqDefForAssociateDesktopsEip()
	return &AssociateDesktopsEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDisassociateDesktopsEip 批量桌面解绑EIP
//
// 批量桌面解绑EIP。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchDisassociateDesktopsEip(request *model.BatchDisassociateDesktopsEipRequest) (*model.BatchDisassociateDesktopsEipResponse, error) {
	requestDef := GenReqDefForBatchDisassociateDesktopsEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDisassociateDesktopsEipResponse), nil
	}
}

// BatchDisassociateDesktopsEipInvoker 批量桌面解绑EIP
func (c *WorkspaceClient) BatchDisassociateDesktopsEipInvoker(request *model.BatchDisassociateDesktopsEipRequest) *BatchDisassociateDesktopsEipInvoker {
	requestDef := GenReqDefForBatchDisassociateDesktopsEip()
	return &BatchDisassociateDesktopsEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSubnetBandwidth 删除云办公带宽
//
// 删除云办公带宽。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) DeleteSubnetBandwidth(request *model.DeleteSubnetBandwidthRequest) (*model.DeleteSubnetBandwidthResponse, error) {
	requestDef := GenReqDefForDeleteSubnetBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSubnetBandwidthResponse), nil
	}
}

// DeleteSubnetBandwidthInvoker 删除云办公带宽
func (c *WorkspaceClient) DeleteSubnetBandwidthInvoker(request *model.DeleteSubnetBandwidthRequest) *DeleteSubnetBandwidthInvoker {
	requestDef := GenReqDefForDeleteSubnetBandwidth()
	return &DeleteSubnetBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDesktopsEips 查询已绑定桌面和未绑定的Eip
//
// 查询已绑定桌面和未绑定的Eip。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListDesktopsEips(request *model.ListDesktopsEipsRequest) (*model.ListDesktopsEipsResponse, error) {
	requestDef := GenReqDefForListDesktopsEips()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDesktopsEipsResponse), nil
	}
}

// ListDesktopsEipsInvoker 查询已绑定桌面和未绑定的Eip
func (c *WorkspaceClient) ListDesktopsEipsInvoker(request *model.ListDesktopsEipsRequest) *ListDesktopsEipsInvoker {
	requestDef := GenReqDefForListDesktopsEips()
	return &ListDesktopsEipsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInternet 查询上网功能
//
// 查询上网功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListInternet(request *model.ListInternetRequest) (*model.ListInternetResponse, error) {
	requestDef := GenReqDefForListInternet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInternetResponse), nil
	}
}

// ListInternetInvoker 查询上网功能
func (c *WorkspaceClient) ListInternetInvoker(request *model.ListInternetRequest) *ListInternetInvoker {
	requestDef := GenReqDefForListInternet()
	return &ListInternetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNatGateways 查询Nat网关列表
//
// 查询NAT网关列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListNatGateways(request *model.ListNatGatewaysRequest) (*model.ListNatGatewaysResponse, error) {
	requestDef := GenReqDefForListNatGateways()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNatGatewaysResponse), nil
	}
}

// ListNatGatewaysInvoker 查询Nat网关列表
func (c *WorkspaceClient) ListNatGatewaysInvoker(request *model.ListNatGatewaysRequest) *ListNatGatewaysInvoker {
	requestDef := GenReqDefForListNatGateways()
	return &ListNatGatewaysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPorts 查询端口列表
//
// 查询端口列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListPorts(request *model.ListPortsRequest) (*model.ListPortsResponse, error) {
	requestDef := GenReqDefForListPorts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPortsResponse), nil
	}
}

// ListPortsInvoker 查询端口列表
func (c *WorkspaceClient) ListPortsInvoker(request *model.ListPortsRequest) *ListPortsInvoker {
	requestDef := GenReqDefForListPorts()
	return &ListPortsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubnetBandwidths 查询云办公带宽列表
//
// 查询云办公带宽列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListSubnetBandwidths(request *model.ListSubnetBandwidthsRequest) (*model.ListSubnetBandwidthsResponse, error) {
	requestDef := GenReqDefForListSubnetBandwidths()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubnetBandwidthsResponse), nil
	}
}

// ListSubnetBandwidthsInvoker 查询云办公带宽列表
func (c *WorkspaceClient) ListSubnetBandwidthsInvoker(request *model.ListSubnetBandwidthsRequest) *ListSubnetBandwidthsInvoker {
	requestDef := GenReqDefForListSubnetBandwidths()
	return &ListSubnetBandwidthsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSubnetBandwidthControlList 查询云办公带宽的控制配置
//
// 查询云办公带宽的控制配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowSubnetBandwidthControlList(request *model.ShowSubnetBandwidthControlListRequest) (*model.ShowSubnetBandwidthControlListResponse, error) {
	requestDef := GenReqDefForShowSubnetBandwidthControlList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSubnetBandwidthControlListResponse), nil
	}
}

// ShowSubnetBandwidthControlListInvoker 查询云办公带宽的控制配置
func (c *WorkspaceClient) ShowSubnetBandwidthControlListInvoker(request *model.ShowSubnetBandwidthControlListRequest) *ShowSubnetBandwidthControlListInvoker {
	requestDef := GenReqDefForShowSubnetBandwidthControlList()
	return &ShowSubnetBandwidthControlListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUsingSubnets 查询正在被使用的子网id列表
//
// 根据子网id列表查询正在被桌面使用的子网id，并且返回subnetId列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowUsingSubnets(request *model.ShowUsingSubnetsRequest) (*model.ShowUsingSubnetsResponse, error) {
	requestDef := GenReqDefForShowUsingSubnets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUsingSubnetsResponse), nil
	}
}

// ShowUsingSubnetsInvoker 查询正在被使用的子网id列表
func (c *WorkspaceClient) ShowUsingSubnetsInvoker(request *model.ShowUsingSubnetsRequest) *ShowUsingSubnetsInvoker {
	requestDef := GenReqDefForShowUsingSubnets()
	return &ShowUsingSubnetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSubnetBandwidth 修改云办公带宽
//
// 修改云办公带宽
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateSubnetBandwidth(request *model.UpdateSubnetBandwidthRequest) (*model.UpdateSubnetBandwidthResponse, error) {
	requestDef := GenReqDefForUpdateSubnetBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSubnetBandwidthResponse), nil
	}
}

// UpdateSubnetBandwidthInvoker 修改云办公带宽
func (c *WorkspaceClient) UpdateSubnetBandwidthInvoker(request *model.UpdateSubnetBandwidthRequest) *UpdateSubnetBandwidthInvoker {
	requestDef := GenReqDefForUpdateSubnetBandwidth()
	return &UpdateSubnetBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSubnetBandwidthControlList 修改云办公带宽的控制配置
//
// 修改云办公带宽的控制配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateSubnetBandwidthControlList(request *model.UpdateSubnetBandwidthControlListRequest) (*model.UpdateSubnetBandwidthControlListResponse, error) {
	requestDef := GenReqDefForUpdateSubnetBandwidthControlList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSubnetBandwidthControlListResponse), nil
	}
}

// UpdateSubnetBandwidthControlListInvoker 修改云办公带宽的控制配置
func (c *WorkspaceClient) UpdateSubnetBandwidthControlListInvoker(request *model.UpdateSubnetBandwidthControlListRequest) *UpdateSubnetBandwidthControlListInvoker {
	requestDef := GenReqDefForUpdateSubnetBandwidthControlList()
	return &UpdateSubnetBandwidthControlListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateChangeOrder 创建变更订单
//
// 变更规格、扩容磁盘[、按需转包周期生成订单](tag:inner)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreateChangeOrder(request *model.CreateChangeOrderRequest) (*model.CreateChangeOrderResponse, error) {
	requestDef := GenReqDefForCreateChangeOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateChangeOrderResponse), nil
	}
}

// CreateChangeOrderInvoker 创建变更订单
func (c *WorkspaceClient) CreateChangeOrderInvoker(request *model.CreateChangeOrderRequest) *CreateChangeOrderInvoker {
	requestDef := GenReqDefForCreateChangeOrder()
	return &CreateChangeOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDesktopBatchOrder 包周期桌面批量变更下单
//
// 包周期桌面批量变更下单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreateDesktopBatchOrder(request *model.CreateDesktopBatchOrderRequest) (*model.CreateDesktopBatchOrderResponse, error) {
	requestDef := GenReqDefForCreateDesktopBatchOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDesktopBatchOrderResponse), nil
	}
}

// CreateDesktopBatchOrderInvoker 包周期桌面批量变更下单
func (c *WorkspaceClient) CreateDesktopBatchOrderInvoker(request *model.CreateDesktopBatchOrderRequest) *CreateDesktopBatchOrderInvoker {
	requestDef := GenReqDefForCreateDesktopBatchOrder()
	return &CreateDesktopBatchOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDesktopOrder 创建桌面订单
//
// 创建桌面订单。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreateDesktopOrder(request *model.CreateDesktopOrderRequest) (*model.CreateDesktopOrderResponse, error) {
	requestDef := GenReqDefForCreateDesktopOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDesktopOrderResponse), nil
	}
}

// CreateDesktopOrderInvoker 创建桌面订单
func (c *WorkspaceClient) CreateDesktopOrderInvoker(request *model.CreateDesktopOrderRequest) *CreateDesktopOrderInvoker {
	requestDef := GenReqDefForCreateDesktopOrder()
	return &CreateDesktopOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDesktopPoolChangeOrder 包周期桌面池批量变更下单
//
// 包周期桌面池批量变更下单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreateDesktopPoolChangeOrder(request *model.CreateDesktopPoolChangeOrderRequest) (*model.CreateDesktopPoolChangeOrderResponse, error) {
	requestDef := GenReqDefForCreateDesktopPoolChangeOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDesktopPoolChangeOrderResponse), nil
	}
}

// CreateDesktopPoolChangeOrderInvoker 包周期桌面池批量变更下单
func (c *WorkspaceClient) CreateDesktopPoolChangeOrderInvoker(request *model.CreateDesktopPoolChangeOrderRequest) *CreateDesktopPoolChangeOrderInvoker {
	requestDef := GenReqDefForCreateDesktopPoolChangeOrder()
	return &CreateDesktopPoolChangeOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateOrder 包周期下单
//
// 包周期资源（桌面、磁盘[、云办公主机](tag:ZQ)）下订单。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreateOrder(request *model.CreateOrderRequest) (*model.CreateOrderResponse, error) {
	requestDef := GenReqDefForCreateOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOrderResponse), nil
	}
}

// CreateOrderInvoker 包周期下单
func (c *WorkspaceClient) CreateOrderInvoker(request *model.CreateOrderRequest) *CreateOrderInvoker {
	requestDef := GenReqDefForCreateOrder()
	return &CreateOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSubnetBandwidthChangeOrder 包周期云办公带宽变更下单
//
// 包周期云办公带宽变更下单。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreateSubnetBandwidthChangeOrder(request *model.CreateSubnetBandwidthChangeOrderRequest) (*model.CreateSubnetBandwidthChangeOrderResponse, error) {
	requestDef := GenReqDefForCreateSubnetBandwidthChangeOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSubnetBandwidthChangeOrderResponse), nil
	}
}

// CreateSubnetBandwidthChangeOrderInvoker 包周期云办公带宽变更下单
func (c *WorkspaceClient) CreateSubnetBandwidthChangeOrderInvoker(request *model.CreateSubnetBandwidthChangeOrderRequest) *CreateSubnetBandwidthChangeOrderInvoker {
	requestDef := GenReqDefForCreateSubnetBandwidthChangeOrder()
	return &CreateSubnetBandwidthChangeOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddOu 新增OU信息
//
// 该接口用于创建OU。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) AddOu(request *model.AddOuRequest) (*model.AddOuResponse, error) {
	requestDef := GenReqDefForAddOu()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddOuResponse), nil
	}
}

// AddOuInvoker 新增OU信息
func (c *WorkspaceClient) AddOuInvoker(request *model.AddOuRequest) *AddOuInvoker {
	requestDef := GenReqDefForAddOu()
	return &AddOuInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteOu 删除OU信息
//
// 该接口用于删除OU信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) DeleteOu(request *model.DeleteOuRequest) (*model.DeleteOuResponse, error) {
	requestDef := GenReqDefForDeleteOu()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteOuResponse), nil
	}
}

// DeleteOuInvoker 删除OU信息
func (c *WorkspaceClient) DeleteOuInvoker(request *model.DeleteOuRequest) *DeleteOuInvoker {
	requestDef := GenReqDefForDeleteOu()
	return &DeleteOuInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAdOuUsers 查询OU下用户信息
//
// 查询OU下用户信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListAdOuUsers(request *model.ListAdOuUsersRequest) (*model.ListAdOuUsersResponse, error) {
	requestDef := GenReqDefForListAdOuUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAdOuUsersResponse), nil
	}
}

// ListAdOuUsersInvoker 查询OU下用户信息
func (c *WorkspaceClient) ListAdOuUsersInvoker(request *model.ListAdOuUsersRequest) *ListAdOuUsersInvoker {
	requestDef := GenReqDefForListAdOuUsers()
	return &ListAdOuUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAdOus 查询OU信息
//
// 查询OU信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListAdOus(request *model.ListAdOusRequest) (*model.ListAdOusResponse, error) {
	requestDef := GenReqDefForListAdOus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAdOusResponse), nil
	}
}

// ListAdOusInvoker 查询OU信息
func (c *WorkspaceClient) ListAdOusInvoker(request *model.ListAdOusRequest) *ListAdOusInvoker {
	requestDef := GenReqDefForListAdOus()
	return &ListAdOusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOuDetails 查询OU信息
//
// 查询OU信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListOuDetails(request *model.ListOuDetailsRequest) (*model.ListOuDetailsResponse, error) {
	requestDef := GenReqDefForListOuDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOuDetailsResponse), nil
	}
}

// ListOuDetailsInvoker 查询OU信息
func (c *WorkspaceClient) ListOuDetailsInvoker(request *model.ListOuDetailsRequest) *ListOuDetailsInvoker {
	requestDef := GenReqDefForListOuDetails()
	return &ListOuDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateOuInfo 更新OU信息
//
// 更新OU信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateOuInfo(request *model.UpdateOuInfoRequest) (*model.UpdateOuInfoResponse, error) {
	requestDef := GenReqDefForUpdateOuInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateOuInfoResponse), nil
	}
}

// UpdateOuInfoInvoker 更新OU信息
func (c *WorkspaceClient) UpdateOuInfoInvoker(request *model.UpdateOuInfoRequest) *UpdateOuInfoInvoker {
	requestDef := GenReqDefForUpdateOuInfo()
	return &UpdateOuInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateTargetOfPolicyGroup 修改策略组应用对象
//
// 批量增加、删除应用对象到指定策略组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchUpdateTargetOfPolicyGroup(request *model.BatchUpdateTargetOfPolicyGroupRequest) (*model.BatchUpdateTargetOfPolicyGroupResponse, error) {
	requestDef := GenReqDefForBatchUpdateTargetOfPolicyGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateTargetOfPolicyGroupResponse), nil
	}
}

// BatchUpdateTargetOfPolicyGroupInvoker 修改策略组应用对象
func (c *WorkspaceClient) BatchUpdateTargetOfPolicyGroupInvoker(request *model.BatchUpdateTargetOfPolicyGroupRequest) *BatchUpdateTargetOfPolicyGroupInvoker {
	requestDef := GenReqDefForBatchUpdateTargetOfPolicyGroup()
	return &BatchUpdateTargetOfPolicyGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePolicyGroup 新增策略组
//
// 新增策略组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreatePolicyGroup(request *model.CreatePolicyGroupRequest) (*model.CreatePolicyGroupResponse, error) {
	requestDef := GenReqDefForCreatePolicyGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePolicyGroupResponse), nil
	}
}

// CreatePolicyGroupInvoker 新增策略组
func (c *WorkspaceClient) CreatePolicyGroupInvoker(request *model.CreatePolicyGroupRequest) *CreatePolicyGroupInvoker {
	requestDef := GenReqDefForCreatePolicyGroup()
	return &CreatePolicyGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePolicyGroup 删除策略组
//
// 删除指定策略组，包含策略组对应的策略信息以及应用对象信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) DeletePolicyGroup(request *model.DeletePolicyGroupRequest) (*model.DeletePolicyGroupResponse, error) {
	requestDef := GenReqDefForDeletePolicyGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePolicyGroupResponse), nil
	}
}

// DeletePolicyGroupInvoker 删除策略组
func (c *WorkspaceClient) DeletePolicyGroupInvoker(request *model.DeletePolicyGroupRequest) *DeletePolicyGroupInvoker {
	requestDef := GenReqDefForDeletePolicyGroup()
	return &DeletePolicyGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOriginalPolicyInfo 查询初始策略项
//
// 查询初始策略项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListOriginalPolicyInfo(request *model.ListOriginalPolicyInfoRequest) (*model.ListOriginalPolicyInfoResponse, error) {
	requestDef := GenReqDefForListOriginalPolicyInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOriginalPolicyInfoResponse), nil
	}
}

// ListOriginalPolicyInfoInvoker 查询初始策略项
func (c *WorkspaceClient) ListOriginalPolicyInfoInvoker(request *model.ListOriginalPolicyInfoRequest) *ListOriginalPolicyInfoInvoker {
	requestDef := GenReqDefForListOriginalPolicyInfo()
	return &ListOriginalPolicyInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPoliciesOfPolicyGroup 查询策略组中的策略项
//
// 查询指定策略组内的策略项。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListPoliciesOfPolicyGroup(request *model.ListPoliciesOfPolicyGroupRequest) (*model.ListPoliciesOfPolicyGroupResponse, error) {
	requestDef := GenReqDefForListPoliciesOfPolicyGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPoliciesOfPolicyGroupResponse), nil
	}
}

// ListPoliciesOfPolicyGroupInvoker 查询策略组中的策略项
func (c *WorkspaceClient) ListPoliciesOfPolicyGroupInvoker(request *model.ListPoliciesOfPolicyGroupRequest) *ListPoliciesOfPolicyGroupInvoker {
	requestDef := GenReqDefForListPoliciesOfPolicyGroup()
	return &ListPoliciesOfPolicyGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicyDetailInfoById 查询策略组
//
// 根据策略组ID查询策略组详细信息，包含策略信息以及应用对象信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListPolicyDetailInfoById(request *model.ListPolicyDetailInfoByIdRequest) (*model.ListPolicyDetailInfoByIdResponse, error) {
	requestDef := GenReqDefForListPolicyDetailInfoById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyDetailInfoByIdResponse), nil
	}
}

// ListPolicyDetailInfoByIdInvoker 查询策略组
func (c *WorkspaceClient) ListPolicyDetailInfoByIdInvoker(request *model.ListPolicyDetailInfoByIdRequest) *ListPolicyDetailInfoByIdInvoker {
	requestDef := GenReqDefForListPolicyDetailInfoById()
	return &ListPolicyDetailInfoByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicyGroup 查询策略组列表
//
// 查询策略组概要信息列表，不包含策略信息以及应用对象信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListPolicyGroup(request *model.ListPolicyGroupRequest) (*model.ListPolicyGroupResponse, error) {
	requestDef := GenReqDefForListPolicyGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyGroupResponse), nil
	}
}

// ListPolicyGroupInvoker 查询策略组列表
func (c *WorkspaceClient) ListPolicyGroupInvoker(request *model.ListPolicyGroupRequest) *ListPolicyGroupInvoker {
	requestDef := GenReqDefForListPolicyGroup()
	return &ListPolicyGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicyGroupInfo 查询策略组详情列表
//
// 包含策略信息以及应用对象的信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListPolicyGroupInfo(request *model.ListPolicyGroupInfoRequest) (*model.ListPolicyGroupInfoResponse, error) {
	requestDef := GenReqDefForListPolicyGroupInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyGroupInfoResponse), nil
	}
}

// ListPolicyGroupInfoInvoker 查询策略组详情列表
func (c *WorkspaceClient) ListPolicyGroupInfoInvoker(request *model.ListPolicyGroupInfoRequest) *ListPolicyGroupInfoInvoker {
	requestDef := GenReqDefForListPolicyGroupInfo()
	return &ListPolicyGroupInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTargetOfPolicyGroup 查询策略组应用对象
//
// 查询指定策略组所应用的对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListTargetOfPolicyGroup(request *model.ListTargetOfPolicyGroupRequest) (*model.ListTargetOfPolicyGroupResponse, error) {
	requestDef := GenReqDefForListTargetOfPolicyGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTargetOfPolicyGroupResponse), nil
	}
}

// ListTargetOfPolicyGroupInvoker 查询策略组应用对象
func (c *WorkspaceClient) ListTargetOfPolicyGroupInvoker(request *model.ListTargetOfPolicyGroupRequest) *ListTargetOfPolicyGroupInvoker {
	requestDef := GenReqDefForListTargetOfPolicyGroup()
	return &ListTargetOfPolicyGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePoliciesOfPolicyGroup 修改策略组中的策略项
//
// 修改一个策略组的部分或者所有策略项。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdatePoliciesOfPolicyGroup(request *model.UpdatePoliciesOfPolicyGroupRequest) (*model.UpdatePoliciesOfPolicyGroupResponse, error) {
	requestDef := GenReqDefForUpdatePoliciesOfPolicyGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePoliciesOfPolicyGroupResponse), nil
	}
}

// UpdatePoliciesOfPolicyGroupInvoker 修改策略组中的策略项
func (c *WorkspaceClient) UpdatePoliciesOfPolicyGroupInvoker(request *model.UpdatePoliciesOfPolicyGroupRequest) *UpdatePoliciesOfPolicyGroupInvoker {
	requestDef := GenReqDefForUpdatePoliciesOfPolicyGroup()
	return &UpdatePoliciesOfPolicyGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePolicyGroup 修改策略组
//
// 修改指定策略组的信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdatePolicyGroup(request *model.UpdatePolicyGroupRequest) (*model.UpdatePolicyGroupResponse, error) {
	requestDef := GenReqDefForUpdatePolicyGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePolicyGroupResponse), nil
	}
}

// UpdatePolicyGroupInvoker 修改策略组
func (c *WorkspaceClient) UpdatePolicyGroupInvoker(request *model.UpdatePolicyGroupRequest) *UpdatePolicyGroupInvoker {
	requestDef := GenReqDefForUpdatePolicyGroup()
	return &UpdatePolicyGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHourPackagesType 查询可订购小时包类型
//
// 该接口用于查询可订购小时包类型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListHourPackagesType(request *model.ListHourPackagesTypeRequest) (*model.ListHourPackagesTypeResponse, error) {
	requestDef := GenReqDefForListHourPackagesType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHourPackagesTypeResponse), nil
	}
}

// ListHourPackagesTypeInvoker 查询可订购小时包类型
func (c *WorkspaceClient) ListHourPackagesTypeInvoker(request *model.ListHourPackagesTypeRequest) *ListHourPackagesTypeInvoker {
	requestDef := GenReqDefForListHourPackagesType()
	return &ListHourPackagesTypeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProducts 查询产品套餐列表
//
// 该接口用于查询云桌面支持的产品套餐列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListProducts(request *model.ListProductsRequest) (*model.ListProductsResponse, error) {
	requestDef := GenReqDefForListProducts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProductsResponse), nil
	}
}

// ListProductsInvoker 查询产品套餐列表
func (c *WorkspaceClient) ListProductsInvoker(request *model.ListProductsRequest) *ListProductsInvoker {
	requestDef := GenReqDefForListProducts()
	return &ListProductsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSharerProducts 查询协同套餐列表
//
// 该接口用于查询协同套餐列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListSharerProducts(request *model.ListSharerProductsRequest) (*model.ListSharerProductsResponse, error) {
	requestDef := GenReqDefForListSharerProducts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSharerProductsResponse), nil
	}
}

// ListSharerProductsInvoker 查询协同套餐列表
func (c *WorkspaceClient) ListSharerProductsInvoker(request *model.ListSharerProductsRequest) *ListSharerProductsInvoker {
	requestDef := GenReqDefForListSharerProducts()
	return &ListSharerProductsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTenantProfiles 查询租户功能状态
//
// 查询租户功能状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListTenantProfiles(request *model.ListTenantProfilesRequest) (*model.ListTenantProfilesResponse, error) {
	requestDef := GenReqDefForListTenantProfiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTenantProfilesResponse), nil
	}
}

// ListTenantProfilesInvoker 查询租户功能状态
func (c *WorkspaceClient) ListTenantProfilesInvoker(request *model.ListTenantProfilesRequest) *ListTenantProfilesInvoker {
	requestDef := GenReqDefForListTenantProfiles()
	return &ListTenantProfilesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTenantProfile 启禁用租户功能
//
// 启禁用租户功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateTenantProfile(request *model.UpdateTenantProfileRequest) (*model.UpdateTenantProfileResponse, error) {
	requestDef := GenReqDefForUpdateTenantProfile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTenantProfileResponse), nil
	}
}

// UpdateTenantProfileInvoker 启禁用租户功能
func (c *WorkspaceClient) UpdateTenantProfileInvoker(request *model.UpdateTenantProfileRequest) *UpdateTenantProfileInvoker {
	requestDef := GenReqDefForUpdateTenantProfile()
	return &UpdateTenantProfileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQuotaDetails 查询租户单个站点配额详情
//
// 查询租户单个站点配额详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowQuotaDetails(request *model.ShowQuotaDetailsRequest) (*model.ShowQuotaDetailsResponse, error) {
	requestDef := GenReqDefForShowQuotaDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotaDetailsResponse), nil
	}
}

// ShowQuotaDetailsInvoker 查询租户单个站点配额详情
func (c *WorkspaceClient) ShowQuotaDetailsInvoker(request *model.ShowQuotaDetailsRequest) *ShowQuotaDetailsInvoker {
	requestDef := GenReqDefForShowQuotaDetails()
	return &ShowQuotaDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQuotas 查询租户配额
//
// Console Framework和WKSDesktopManager调用该接口查询租户配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowQuotas(request *model.ShowQuotasRequest) (*model.ShowQuotasResponse, error) {
	requestDef := GenReqDefForShowQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotasResponse), nil
	}
}

// ShowQuotasInvoker 查询租户配额
func (c *WorkspaceClient) ShowQuotasInvoker(request *model.ShowQuotasRequest) *ShowQuotasInvoker {
	requestDef := GenReqDefForShowQuotas()
	return &ShowQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteScheduledTasks 批量删除定时任务
//
// 批量删除定时任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchDeleteScheduledTasks(request *model.BatchDeleteScheduledTasksRequest) (*model.BatchDeleteScheduledTasksResponse, error) {
	requestDef := GenReqDefForBatchDeleteScheduledTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteScheduledTasksResponse), nil
	}
}

// BatchDeleteScheduledTasksInvoker 批量删除定时任务
func (c *WorkspaceClient) BatchDeleteScheduledTasksInvoker(request *model.BatchDeleteScheduledTasksRequest) *BatchDeleteScheduledTasksInvoker {
	requestDef := GenReqDefForBatchDeleteScheduledTasks()
	return &BatchDeleteScheduledTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateScheduledTasks 创建定时任务
//
// 创建定时任务。
// 注:需通过开通委托功能接口先对云服务进行授权才可以使用该功能
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreateScheduledTasks(request *model.CreateScheduledTasksRequest) (*model.CreateScheduledTasksResponse, error) {
	requestDef := GenReqDefForCreateScheduledTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateScheduledTasksResponse), nil
	}
}

// CreateScheduledTasksInvoker 创建定时任务
func (c *WorkspaceClient) CreateScheduledTasksInvoker(request *model.CreateScheduledTasksRequest) *CreateScheduledTasksInvoker {
	requestDef := GenReqDefForCreateScheduledTasks()
	return &CreateScheduledTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteScheduledTasks 删除定时任务
//
// 删除定时任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) DeleteScheduledTasks(request *model.DeleteScheduledTasksRequest) (*model.DeleteScheduledTasksResponse, error) {
	requestDef := GenReqDefForDeleteScheduledTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteScheduledTasksResponse), nil
	}
}

// DeleteScheduledTasksInvoker 删除定时任务
func (c *WorkspaceClient) DeleteScheduledTasksInvoker(request *model.DeleteScheduledTasksRequest) *DeleteScheduledTasksInvoker {
	requestDef := GenReqDefForDeleteScheduledTasks()
	return &DeleteScheduledTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFutureExecutions 未来执行的具体时间列表
//
// 未来执行的具体时间列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListFutureExecutions(request *model.ListFutureExecutionsRequest) (*model.ListFutureExecutionsResponse, error) {
	requestDef := GenReqDefForListFutureExecutions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFutureExecutionsResponse), nil
	}
}

// ListFutureExecutionsInvoker 未来执行的具体时间列表
func (c *WorkspaceClient) ListFutureExecutionsInvoker(request *model.ListFutureExecutionsRequest) *ListFutureExecutionsInvoker {
	requestDef := GenReqDefForListFutureExecutions()
	return &ListFutureExecutionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScheduledTasks 查询定时任务列表
//
// 查询定时任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListScheduledTasks(request *model.ListScheduledTasksRequest) (*model.ListScheduledTasksResponse, error) {
	requestDef := GenReqDefForListScheduledTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScheduledTasksResponse), nil
	}
}

// ListScheduledTasksInvoker 查询定时任务列表
func (c *WorkspaceClient) ListScheduledTasksInvoker(request *model.ListScheduledTasksRequest) *ListScheduledTasksInvoker {
	requestDef := GenReqDefForListScheduledTasks()
	return &ListScheduledTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScheduledTasksRecords 查询定时任务执行记录
//
// 查询定时任务执行记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListScheduledTasksRecords(request *model.ListScheduledTasksRecordsRequest) (*model.ListScheduledTasksRecordsResponse, error) {
	requestDef := GenReqDefForListScheduledTasksRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScheduledTasksRecordsResponse), nil
	}
}

// ListScheduledTasksRecordsInvoker 查询定时任务执行记录
func (c *WorkspaceClient) ListScheduledTasksRecordsInvoker(request *model.ListScheduledTasksRecordsRequest) *ListScheduledTasksRecordsInvoker {
	requestDef := GenReqDefForListScheduledTasksRecords()
	return &ListScheduledTasksRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScheduledTasksRecordsDetails 查询定时任务执行记录详情
//
// 查询定时任务执行记录详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListScheduledTasksRecordsDetails(request *model.ListScheduledTasksRecordsDetailsRequest) (*model.ListScheduledTasksRecordsDetailsResponse, error) {
	requestDef := GenReqDefForListScheduledTasksRecordsDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScheduledTasksRecordsDetailsResponse), nil
	}
}

// ListScheduledTasksRecordsDetailsInvoker 查询定时任务执行记录详情
func (c *WorkspaceClient) ListScheduledTasksRecordsDetailsInvoker(request *model.ListScheduledTasksRecordsDetailsRequest) *ListScheduledTasksRecordsDetailsInvoker {
	requestDef := GenReqDefForListScheduledTasksRecordsDetails()
	return &ListScheduledTasksRecordsDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTimeZones 获取时区配置
//
// 获取时区配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListTimeZones(request *model.ListTimeZonesRequest) (*model.ListTimeZonesResponse, error) {
	requestDef := GenReqDefForListTimeZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTimeZonesResponse), nil
	}
}

// ListTimeZonesInvoker 获取时区配置
func (c *WorkspaceClient) ListTimeZonesInvoker(request *model.ListTimeZonesRequest) *ListTimeZonesInvoker {
	requestDef := GenReqDefForListTimeZones()
	return &ListTimeZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowScheduledTasks 查询定时任务详情
//
// 查询定时任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowScheduledTasks(request *model.ShowScheduledTasksRequest) (*model.ShowScheduledTasksResponse, error) {
	requestDef := GenReqDefForShowScheduledTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowScheduledTasksResponse), nil
	}
}

// ShowScheduledTasksInvoker 查询定时任务详情
func (c *WorkspaceClient) ShowScheduledTasksInvoker(request *model.ShowScheduledTasksRequest) *ShowScheduledTasksInvoker {
	requestDef := GenReqDefForShowScheduledTasks()
	return &ShowScheduledTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateScheduledTasks 修改定时任务
//
// 修改定时任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateScheduledTasks(request *model.UpdateScheduledTasksRequest) (*model.UpdateScheduledTasksResponse, error) {
	requestDef := GenReqDefForUpdateScheduledTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateScheduledTasksResponse), nil
	}
}

// UpdateScheduledTasksInvoker 修改定时任务
func (c *WorkspaceClient) UpdateScheduledTasksInvoker(request *model.UpdateScheduledTasksRequest) *UpdateScheduledTasksInvoker {
	requestDef := GenReqDefForUpdateScheduledTasks()
	return &UpdateScheduledTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteScreenRecords 批量删除录屏记录
//
// 批量删除录屏记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchDeleteScreenRecords(request *model.BatchDeleteScreenRecordsRequest) (*model.BatchDeleteScreenRecordsResponse, error) {
	requestDef := GenReqDefForBatchDeleteScreenRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteScreenRecordsResponse), nil
	}
}

// BatchDeleteScreenRecordsInvoker 批量删除录屏记录
func (c *WorkspaceClient) BatchDeleteScreenRecordsInvoker(request *model.BatchDeleteScreenRecordsRequest) *BatchDeleteScreenRecordsInvoker {
	requestDef := GenReqDefForBatchDeleteScreenRecords()
	return &BatchDeleteScreenRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDesktopOperations 查询桌面关键事件列表
//
// 查询桌面关键事件列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListDesktopOperations(request *model.ListDesktopOperationsRequest) (*model.ListDesktopOperationsResponse, error) {
	requestDef := GenReqDefForListDesktopOperations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDesktopOperationsResponse), nil
	}
}

// ListDesktopOperationsInvoker 查询桌面关键事件列表
func (c *WorkspaceClient) ListDesktopOperationsInvoker(request *model.ListDesktopOperationsRequest) *ListDesktopOperationsInvoker {
	requestDef := GenReqDefForListDesktopOperations()
	return &ListDesktopOperationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDownloadAddress 查询下载地址列表
//
// 查询下载地址列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListDownloadAddress(request *model.ListDownloadAddressRequest) (*model.ListDownloadAddressResponse, error) {
	requestDef := GenReqDefForListDownloadAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDownloadAddressResponse), nil
	}
}

// ListDownloadAddressInvoker 查询下载地址列表
func (c *WorkspaceClient) ListDownloadAddressInvoker(request *model.ListDownloadAddressRequest) *ListDownloadAddressInvoker {
	requestDef := GenReqDefForListDownloadAddress()
	return &ListDownloadAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScreenRecords 查询录屏记录列表
//
// 查询录屏记录列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListScreenRecords(request *model.ListScreenRecordsRequest) (*model.ListScreenRecordsResponse, error) {
	requestDef := GenReqDefForListScreenRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScreenRecordsResponse), nil
	}
}

// ListScreenRecordsInvoker 查询录屏记录列表
func (c *WorkspaceClient) ListScreenRecordsInvoker(request *model.ListScreenRecordsRequest) *ListScreenRecordsInvoker {
	requestDef := GenReqDefForListScreenRecords()
	return &ListScreenRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowScreenRecord 查询录屏详情
//
// 查询录屏详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowScreenRecord(request *model.ShowScreenRecordRequest) (*model.ShowScreenRecordResponse, error) {
	requestDef := GenReqDefForShowScreenRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowScreenRecordResponse), nil
	}
}

// ShowScreenRecordInvoker 查询录屏详情
func (c *WorkspaceClient) ShowScreenRecordInvoker(request *model.ShowScreenRecordRequest) *ShowScreenRecordInvoker {
	requestDef := GenReqDefForShowScreenRecord()
	return &ShowScreenRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateScript 新增脚本
//
// 新增脚本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreateScript(request *model.CreateScriptRequest) (*model.CreateScriptResponse, error) {
	requestDef := GenReqDefForCreateScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateScriptResponse), nil
	}
}

// CreateScriptInvoker 新增脚本
func (c *WorkspaceClient) CreateScriptInvoker(request *model.CreateScriptRequest) *CreateScriptInvoker {
	requestDef := GenReqDefForCreateScript()
	return &CreateScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteScript 删除脚本
//
// 删除脚本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) DeleteScript(request *model.DeleteScriptRequest) (*model.DeleteScriptResponse, error) {
	requestDef := GenReqDefForDeleteScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteScriptResponse), nil
	}
}

// DeleteScriptInvoker 删除脚本
func (c *WorkspaceClient) DeleteScriptInvoker(request *model.DeleteScriptRequest) *DeleteScriptInvoker {
	requestDef := GenReqDefForDeleteScript()
	return &DeleteScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteScriptOrCommand 批量执行脚本或命令
//
// 批量执行脚本或命令。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ExecuteScriptOrCommand(request *model.ExecuteScriptOrCommandRequest) (*model.ExecuteScriptOrCommandResponse, error) {
	requestDef := GenReqDefForExecuteScriptOrCommand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteScriptOrCommandResponse), nil
	}
}

// ExecuteScriptOrCommandInvoker 批量执行脚本或命令
func (c *WorkspaceClient) ExecuteScriptOrCommandInvoker(request *model.ExecuteScriptOrCommandRequest) *ExecuteScriptOrCommandInvoker {
	requestDef := GenReqDefForExecuteScriptOrCommand()
	return &ExecuteScriptOrCommandInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScriptRecords 查询脚本执行记录列表
//
// 查询脚本执行记录列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListScriptRecords(request *model.ListScriptRecordsRequest) (*model.ListScriptRecordsResponse, error) {
	requestDef := GenReqDefForListScriptRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScriptRecordsResponse), nil
	}
}

// ListScriptRecordsInvoker 查询脚本执行记录列表
func (c *WorkspaceClient) ListScriptRecordsInvoker(request *model.ListScriptRecordsRequest) *ListScriptRecordsInvoker {
	requestDef := GenReqDefForListScriptRecords()
	return &ListScriptRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScriptTasks 查询脚本任务列表
//
// 查询脚本任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListScriptTasks(request *model.ListScriptTasksRequest) (*model.ListScriptTasksResponse, error) {
	requestDef := GenReqDefForListScriptTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScriptTasksResponse), nil
	}
}

// ListScriptTasksInvoker 查询脚本任务列表
func (c *WorkspaceClient) ListScriptTasksInvoker(request *model.ListScriptTasksRequest) *ListScriptTasksInvoker {
	requestDef := GenReqDefForListScriptTasks()
	return &ListScriptTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScripts 查询脚本列表
//
// 查询脚本列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListScripts(request *model.ListScriptsRequest) (*model.ListScriptsResponse, error) {
	requestDef := GenReqDefForListScripts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScriptsResponse), nil
	}
}

// ListScriptsInvoker 查询脚本列表
func (c *WorkspaceClient) ListScriptsInvoker(request *model.ListScriptsRequest) *ListScriptsInvoker {
	requestDef := GenReqDefForListScripts()
	return &ListScriptsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RetryScriptExecution 重试脚本或执行脚本失败的任务
//
// 重试脚本或执行脚本失败的任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) RetryScriptExecution(request *model.RetryScriptExecutionRequest) (*model.RetryScriptExecutionResponse, error) {
	requestDef := GenReqDefForRetryScriptExecution()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RetryScriptExecutionResponse), nil
	}
}

// RetryScriptExecutionInvoker 重试脚本或执行脚本失败的任务
func (c *WorkspaceClient) RetryScriptExecutionInvoker(request *model.RetryScriptExecutionRequest) *RetryScriptExecutionInvoker {
	requestDef := GenReqDefForRetryScriptExecution()
	return &RetryScriptExecutionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowScriptDetail 查询脚本详情
//
// 查询脚本详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowScriptDetail(request *model.ShowScriptDetailRequest) (*model.ShowScriptDetailResponse, error) {
	requestDef := GenReqDefForShowScriptDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowScriptDetailResponse), nil
	}
}

// ShowScriptDetailInvoker 查询脚本详情
func (c *WorkspaceClient) ShowScriptDetailInvoker(request *model.ShowScriptDetailRequest) *ShowScriptDetailInvoker {
	requestDef := GenReqDefForShowScriptDetail()
	return &ShowScriptDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowScriptRecordDetail 查询脚本执行记录详情
//
// 查询脚本执行记录详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowScriptRecordDetail(request *model.ShowScriptRecordDetailRequest) (*model.ShowScriptRecordDetailResponse, error) {
	requestDef := GenReqDefForShowScriptRecordDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowScriptRecordDetailResponse), nil
	}
}

// ShowScriptRecordDetailInvoker 查询脚本执行记录详情
func (c *WorkspaceClient) ShowScriptRecordDetailInvoker(request *model.ShowScriptRecordDetailRequest) *ShowScriptRecordDetailInvoker {
	requestDef := GenReqDefForShowScriptRecordDetail()
	return &ShowScriptRecordDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopScriptExecution 停止脚本或者命令执行任务
//
// 停止脚本或者命令执行任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) StopScriptExecution(request *model.StopScriptExecutionRequest) (*model.StopScriptExecutionResponse, error) {
	requestDef := GenReqDefForStopScriptExecution()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopScriptExecutionResponse), nil
	}
}

// StopScriptExecutionInvoker 停止脚本或者命令执行任务
func (c *WorkspaceClient) StopScriptExecutionInvoker(request *model.StopScriptExecutionRequest) *StopScriptExecutionInvoker {
	requestDef := GenReqDefForStopScriptExecution()
	return &StopScriptExecutionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateScript 更新脚本
//
// 更新脚本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateScript(request *model.UpdateScriptRequest) (*model.UpdateScriptResponse, error) {
	requestDef := GenReqDefForUpdateScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateScriptResponse), nil
	}
}

// UpdateScriptInvoker 更新脚本
func (c *WorkspaceClient) UpdateScriptInvoker(request *model.UpdateScriptRequest) *UpdateScriptInvoker {
	requestDef := GenReqDefForUpdateScript()
	return &UpdateScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddDesktopSubResources 桌面购买附属资源
//
// 存量桌面购买附属资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) AddDesktopSubResources(request *model.AddDesktopSubResourcesRequest) (*model.AddDesktopSubResourcesResponse, error) {
	requestDef := GenReqDefForAddDesktopSubResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddDesktopSubResourcesResponse), nil
	}
}

// AddDesktopSubResourcesInvoker 桌面购买附属资源
func (c *WorkspaceClient) AddDesktopSubResourcesInvoker(request *model.AddDesktopSubResourcesRequest) *AddDesktopSubResourcesInvoker {
	requestDef := GenReqDefForAddDesktopSubResources()
	return &AddDesktopSubResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDesktopSubResources 桌面删除附属资源
//
// 桌面删除附属资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) DeleteDesktopSubResources(request *model.DeleteDesktopSubResourcesRequest) (*model.DeleteDesktopSubResourcesResponse, error) {
	requestDef := GenReqDefForDeleteDesktopSubResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDesktopSubResourcesResponse), nil
	}
}

// DeleteDesktopSubResourcesInvoker 桌面删除附属资源
func (c *WorkspaceClient) DeleteDesktopSubResourcesInvoker(request *model.DeleteDesktopSubResourcesRequest) *DeleteDesktopSubResourcesInvoker {
	requestDef := GenReqDefForDeleteDesktopSubResources()
	return &DeleteDesktopSubResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowShareSpaceConfig 查询协同桌面默认用户配置
//
// 查询协同桌面默认用户配置（当前功能公测中,需要使用请联系管理员申请使用）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowShareSpaceConfig(request *model.ShowShareSpaceConfigRequest) (*model.ShowShareSpaceConfigResponse, error) {
	requestDef := GenReqDefForShowShareSpaceConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowShareSpaceConfigResponse), nil
	}
}

// ShowShareSpaceConfigInvoker 查询协同桌面默认用户配置
func (c *WorkspaceClient) ShowShareSpaceConfigInvoker(request *model.ShowShareSpaceConfigRequest) *ShowShareSpaceConfigInvoker {
	requestDef := GenReqDefForShowShareSpaceConfig()
	return &ShowShareSpaceConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateShareSpaceConfig 设置协同桌面默认用户配置
//
// 设置协同桌面默认用户配置（当前功能公测中，需要使用请联系管理员申请使用）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateShareSpaceConfig(request *model.UpdateShareSpaceConfigRequest) (*model.UpdateShareSpaceConfigResponse, error) {
	requestDef := GenReqDefForUpdateShareSpaceConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateShareSpaceConfigResponse), nil
	}
}

// UpdateShareSpaceConfigInvoker 设置协同桌面默认用户配置
func (c *WorkspaceClient) UpdateShareSpaceConfigInvoker(request *model.UpdateShareSpaceConfigRequest) *UpdateShareSpaceConfigInvoker {
	requestDef := GenReqDefForUpdateShareSpaceConfig()
	return &UpdateShareSpaceConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddSite 新增站点
//
// 用于查询站点信息的接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) AddSite(request *model.AddSiteRequest) (*model.AddSiteResponse, error) {
	requestDef := GenReqDefForAddSite()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddSiteResponse), nil
	}
}

// AddSiteInvoker 新增站点
func (c *WorkspaceClient) AddSiteInvoker(request *model.AddSiteRequest) *AddSiteInvoker {
	requestDef := GenReqDefForAddSite()
	return &AddSiteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSite 删除站点
//
// 用于删除站点的接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) DeleteSite(request *model.DeleteSiteRequest) (*model.DeleteSiteResponse, error) {
	requestDef := GenReqDefForDeleteSite()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSiteResponse), nil
	}
}

// DeleteSiteInvoker 删除站点
func (c *WorkspaceClient) DeleteSiteInvoker(request *model.DeleteSiteRequest) *DeleteSiteInvoker {
	requestDef := GenReqDefForDeleteSite()
	return &DeleteSiteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSiteConfigs 查询站点信息
//
// 用于查询站点信息的接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListSiteConfigs(request *model.ListSiteConfigsRequest) (*model.ListSiteConfigsResponse, error) {
	requestDef := GenReqDefForListSiteConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSiteConfigsResponse), nil
	}
}

// ListSiteConfigsInvoker 查询站点信息
func (c *WorkspaceClient) ListSiteConfigsInvoker(request *model.ListSiteConfigsRequest) *ListSiteConfigsInvoker {
	requestDef := GenReqDefForListSiteConfigs()
	return &ListSiteConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWksEdgeSites 查询边缘小站列表
//
// 查询边缘小站列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListWksEdgeSites(request *model.ListWksEdgeSitesRequest) (*model.ListWksEdgeSitesResponse, error) {
	requestDef := GenReqDefForListWksEdgeSites()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWksEdgeSitesResponse), nil
	}
}

// ListWksEdgeSitesInvoker 查询边缘小站列表
func (c *WorkspaceClient) ListWksEdgeSitesInvoker(request *model.ListWksEdgeSitesRequest) *ListWksEdgeSitesInvoker {
	requestDef := GenReqDefForListWksEdgeSites()
	return &ListWksEdgeSitesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAccessMode 修改站点接入方式
//
// 用于修改站点接入方式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateAccessMode(request *model.UpdateAccessModeRequest) (*model.UpdateAccessModeResponse, error) {
	requestDef := GenReqDefForUpdateAccessMode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAccessModeResponse), nil
	}
}

// UpdateAccessModeInvoker 修改站点接入方式
func (c *WorkspaceClient) UpdateAccessModeInvoker(request *model.UpdateAccessModeRequest) *UpdateAccessModeInvoker {
	requestDef := GenReqDefForUpdateAccessMode()
	return &UpdateAccessModeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSubnetIds 修改站点业务子网
//
// 用于修改站点业务子网
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateSubnetIds(request *model.UpdateSubnetIdsRequest) (*model.UpdateSubnetIdsResponse, error) {
	requestDef := GenReqDefForUpdateSubnetIds()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSubnetIdsResponse), nil
	}
}

// UpdateSubnetIdsInvoker 修改站点业务子网
func (c *WorkspaceClient) UpdateSubnetIdsInvoker(request *model.UpdateSubnetIdsRequest) *UpdateSubnetIdsInvoker {
	requestDef := GenReqDefForUpdateSubnetIds()
	return &UpdateSubnetIdsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateDesktopSnapshot 批量创建快照
//
// 批量创建快照
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchCreateDesktopSnapshot(request *model.BatchCreateDesktopSnapshotRequest) (*model.BatchCreateDesktopSnapshotResponse, error) {
	requestDef := GenReqDefForBatchCreateDesktopSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateDesktopSnapshotResponse), nil
	}
}

// BatchCreateDesktopSnapshotInvoker 批量创建快照
func (c *WorkspaceClient) BatchCreateDesktopSnapshotInvoker(request *model.BatchCreateDesktopSnapshotRequest) *BatchCreateDesktopSnapshotInvoker {
	requestDef := GenReqDefForBatchCreateDesktopSnapshot()
	return &BatchCreateDesktopSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteDesktopSnapshot 批量删除快照
//
// 批量删除快照
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchDeleteDesktopSnapshot(request *model.BatchDeleteDesktopSnapshotRequest) (*model.BatchDeleteDesktopSnapshotResponse, error) {
	requestDef := GenReqDefForBatchDeleteDesktopSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteDesktopSnapshotResponse), nil
	}
}

// BatchDeleteDesktopSnapshotInvoker 批量删除快照
func (c *WorkspaceClient) BatchDeleteDesktopSnapshotInvoker(request *model.BatchDeleteDesktopSnapshotRequest) *BatchDeleteDesktopSnapshotInvoker {
	requestDef := GenReqDefForBatchDeleteDesktopSnapshot()
	return &BatchDeleteDesktopSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchRestoreDesktopSnapshot 批量恢复快照
//
// 批量恢快照
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchRestoreDesktopSnapshot(request *model.BatchRestoreDesktopSnapshotRequest) (*model.BatchRestoreDesktopSnapshotResponse, error) {
	requestDef := GenReqDefForBatchRestoreDesktopSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRestoreDesktopSnapshotResponse), nil
	}
}

// BatchRestoreDesktopSnapshotInvoker 批量恢复快照
func (c *WorkspaceClient) BatchRestoreDesktopSnapshotInvoker(request *model.BatchRestoreDesktopSnapshotRequest) *BatchRestoreDesktopSnapshotInvoker {
	requestDef := GenReqDefForBatchRestoreDesktopSnapshot()
	return &BatchRestoreDesktopSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDesktopSnapshot 查询快照列表
//
// 查询快照列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListDesktopSnapshot(request *model.ListDesktopSnapshotRequest) (*model.ListDesktopSnapshotResponse, error) {
	requestDef := GenReqDefForListDesktopSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDesktopSnapshotResponse), nil
	}
}

// ListDesktopSnapshotInvoker 查询快照列表
func (c *WorkspaceClient) ListDesktopSnapshotInvoker(request *model.ListDesktopSnapshotRequest) *ListDesktopSnapshotInvoker {
	requestDef := GenReqDefForListDesktopSnapshot()
	return &ListDesktopSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddMetricNotifyRule 新增通知规则
//
// 新增对应指标的通知规则;对应指标满足相应的规则条件时发送通知
// 同一指标的规则不允许重复;
// 统计指标名称，目前仅支持固定值：desktop_idle_duration
//   * &#x60;desktop_idle_duration&#x60; -  桌面空闲时长, 仅允许设置 &#39;&gt;&#x3D;&#39; 阈值
// 注：需先为云服务添加委托授权，否则无法正常发送通知到SMN
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) AddMetricNotifyRule(request *model.AddMetricNotifyRuleRequest) (*model.AddMetricNotifyRuleResponse, error) {
	requestDef := GenReqDefForAddMetricNotifyRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddMetricNotifyRuleResponse), nil
	}
}

// AddMetricNotifyRuleInvoker 新增通知规则
func (c *WorkspaceClient) AddMetricNotifyRuleInvoker(request *model.AddMetricNotifyRuleRequest) *AddMetricNotifyRuleInvoker {
	requestDef := GenReqDefForAddMetricNotifyRule()
	return &AddMetricNotifyRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMetricNotifyRule 删除通知规则
//
// 删除对应指标的通知规则;对应指标满足相应的规则条件时发送通知
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) DeleteMetricNotifyRule(request *model.DeleteMetricNotifyRuleRequest) (*model.DeleteMetricNotifyRuleResponse, error) {
	requestDef := GenReqDefForDeleteMetricNotifyRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMetricNotifyRuleResponse), nil
	}
}

// DeleteMetricNotifyRuleInvoker 删除通知规则
func (c *WorkspaceClient) DeleteMetricNotifyRuleInvoker(request *model.DeleteMetricNotifyRuleRequest) *DeleteMetricNotifyRuleInvoker {
	requestDef := GenReqDefForDeleteMetricNotifyRule()
	return &DeleteMetricNotifyRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppUserAccessData 查询云应用接入统计数据
//
// 查询云应用接入统计数据;
// 最多查询30天内的数据;
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListAppUserAccessData(request *model.ListAppUserAccessDataRequest) (*model.ListAppUserAccessDataResponse, error) {
	requestDef := GenReqDefForListAppUserAccessData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppUserAccessDataResponse), nil
	}
}

// ListAppUserAccessDataInvoker 查询云应用接入统计数据
func (c *WorkspaceClient) ListAppUserAccessDataInvoker(request *model.ListAppUserAccessDataRequest) *ListAppUserAccessDataInvoker {
	requestDef := GenReqDefForListAppUserAccessData()
	return &ListAppUserAccessDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDesktopUsageMetric 查询桌面使用情况统计数据
//
// 查询桌面使用统计信息;
// 云服务每天凌晨02:00进行聚合运算前一天00:00:00~23:59:59的使用时长,并将周期范围内的数据聚合到周期边界上
// 跨天的记录会按照统计周期进行计算
// 假设一天内桌面登录多次，09:00~12:00,13:00~21:00,22:00~01:00(次日):
// 则当天的累计使用时长数据会被汇聚到23:59:59这个点;总使用时长为 3hours(09:00~12:00)+8hours(13:00~21:00)+2hours(22:00~00:00)
// 仅能查询最近180天已进行汇聚计算的数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListDesktopUsageMetric(request *model.ListDesktopUsageMetricRequest) (*model.ListDesktopUsageMetricResponse, error) {
	requestDef := GenReqDefForListDesktopUsageMetric()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDesktopUsageMetricResponse), nil
	}
}

// ListDesktopUsageMetricInvoker 查询桌面使用情况统计数据
func (c *WorkspaceClient) ListDesktopUsageMetricInvoker(request *model.ListDesktopUsageMetricRequest) *ListDesktopUsageMetricInvoker {
	requestDef := GenReqDefForListDesktopUsageMetric()
	return &ListDesktopUsageMetricInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDesktopsStatistics 桌面统计
//
// 统计租户下的普通桌面、桌面池状态，默认仅统计总数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListDesktopsStatistics(request *model.ListDesktopsStatisticsRequest) (*model.ListDesktopsStatisticsResponse, error) {
	requestDef := GenReqDefForListDesktopsStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDesktopsStatisticsResponse), nil
	}
}

// ListDesktopsStatisticsInvoker 桌面统计
func (c *WorkspaceClient) ListDesktopsStatisticsInvoker(request *model.ListDesktopsStatisticsRequest) *ListDesktopsStatisticsInvoker {
	requestDef := GenReqDefForListDesktopsStatistics()
	return &ListDesktopsStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLoginState 登录状态统计
//
// 查询租户桌面登录状态统计。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListLoginState(request *model.ListLoginStateRequest) (*model.ListLoginStateResponse, error) {
	requestDef := GenReqDefForListLoginState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLoginStateResponse), nil
	}
}

// ListLoginStateInvoker 登录状态统计
func (c *WorkspaceClient) ListLoginStateInvoker(request *model.ListLoginStateRequest) *ListLoginStateInvoker {
	requestDef := GenReqDefForListLoginState()
	return &ListLoginStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMetricNotifyRecord 查询对应指标维度是否存在满足通知规则的记录
//
// 查询对应指标维度是否存在满足通知规则的记录;
// 查询结果仅表示满足相应指标维度下对应通知规则可产生的通知记录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListMetricNotifyRecord(request *model.ListMetricNotifyRecordRequest) (*model.ListMetricNotifyRecordResponse, error) {
	requestDef := GenReqDefForListMetricNotifyRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMetricNotifyRecordResponse), nil
	}
}

// ListMetricNotifyRecordInvoker 查询对应指标维度是否存在满足通知规则的记录
func (c *WorkspaceClient) ListMetricNotifyRecordInvoker(request *model.ListMetricNotifyRecordRequest) *ListMetricNotifyRecordInvoker {
	requestDef := GenReqDefForListMetricNotifyRecord()
	return &ListMetricNotifyRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMetricNotifyRule 查询通知规则
//
// 查询对应指标的通知规则;
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListMetricNotifyRule(request *model.ListMetricNotifyRuleRequest) (*model.ListMetricNotifyRuleResponse, error) {
	requestDef := GenReqDefForListMetricNotifyRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMetricNotifyRuleResponse), nil
	}
}

// ListMetricNotifyRuleInvoker 查询通知规则
func (c *WorkspaceClient) ListMetricNotifyRuleInvoker(request *model.ListMetricNotifyRuleRequest) *ListMetricNotifyRuleInvoker {
	requestDef := GenReqDefForListMetricNotifyRule()
	return &ListMetricNotifyRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMetrics 查询指标
//
// 查询指标
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListMetrics(request *model.ListMetricsRequest) (*model.ListMetricsResponse, error) {
	requestDef := GenReqDefForListMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMetricsResponse), nil
	}
}

// ListMetricsInvoker 查询指标
func (c *WorkspaceClient) ListMetricsInvoker(request *model.ListMetricsRequest) *ListMetricsInvoker {
	requestDef := GenReqDefForListMetrics()
	return &ListMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMetricsTrend 查询指标趋势
//
// 查询指标趋势
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListMetricsTrend(request *model.ListMetricsTrendRequest) (*model.ListMetricsTrendResponse, error) {
	requestDef := GenReqDefForListMetricsTrend()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMetricsTrendResponse), nil
	}
}

// ListMetricsTrendInvoker 查询指标趋势
func (c *WorkspaceClient) ListMetricsTrendInvoker(request *model.ListMetricsTrendRequest) *ListMetricsTrendInvoker {
	requestDef := GenReqDefForListMetricsTrend()
	return &ListMetricsTrendInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRunState 运行状态统计
//
// 租户桌面运行状态统计。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListRunState(request *model.ListRunStateRequest) (*model.ListRunStateResponse, error) {
	requestDef := GenReqDefForListRunState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRunStateResponse), nil
	}
}

// ListRunStateInvoker 运行状态统计
func (c *WorkspaceClient) ListRunStateInvoker(request *model.ListRunStateRequest) *ListRunStateInvoker {
	requestDef := GenReqDefForListRunState()
	return &ListRunStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUnusedDesktops 查询在指定时间段未使用的桌面
//
// 查询在指定时间段未使用的桌面。已废弃，不推荐使用。统计数据推荐使用[查询桌面使用情况统计数据接口](https://console.huaweicloud.com/apiexplorer/#/openapi/Workspace/doc?api&#x3D;ListDesktopUsageMetric)和[查询用户使用统计数据接口](https://console.huaweicloud.com/apiexplorer/#/openapi/Workspace/doc?api&#x3D;ListUserUsageMetric)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListUnusedDesktops(request *model.ListUnusedDesktopsRequest) (*model.ListUnusedDesktopsResponse, error) {
	requestDef := GenReqDefForListUnusedDesktops()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUnusedDesktopsResponse), nil
	}
}

// ListUnusedDesktopsInvoker 查询在指定时间段未使用的桌面
func (c *WorkspaceClient) ListUnusedDesktopsInvoker(request *model.ListUnusedDesktopsRequest) *ListUnusedDesktopsInvoker {
	requestDef := GenReqDefForListUnusedDesktops()
	return &ListUnusedDesktopsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUsedDesktopInfo 查询使用桌面的时长
//
// 查询使用桌面的时长。已废弃，不推荐使用。统计数据推荐使用[查询桌面使用情况统计数据接口](https://console.huaweicloud.com/apiexplorer/#/openapi/Workspace/doc?api&#x3D;ListDesktopUsageMetric)和[查询用户使用统计数据接口](https://console.huaweicloud.com/apiexplorer/#/openapi/Workspace/doc?api&#x3D;ListUserUsageMetric)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListUsedDesktopInfo(request *model.ListUsedDesktopInfoRequest) (*model.ListUsedDesktopInfoResponse, error) {
	requestDef := GenReqDefForListUsedDesktopInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsedDesktopInfoResponse), nil
	}
}

// ListUsedDesktopInfoInvoker 查询使用桌面的时长
func (c *WorkspaceClient) ListUsedDesktopInfoInvoker(request *model.ListUsedDesktopInfoRequest) *ListUsedDesktopInfoInvoker {
	requestDef := GenReqDefForListUsedDesktopInfo()
	return &ListUsedDesktopInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserUsageMetric 查询用户使用统计数据
//
// 查询用户使用统计信息;
// 最多查询30天内的数据;
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListUserUsageMetric(request *model.ListUserUsageMetricRequest) (*model.ListUserUsageMetricResponse, error) {
	requestDef := GenReqDefForListUserUsageMetric()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserUsageMetricResponse), nil
	}
}

// ListUserUsageMetricInvoker 查询用户使用统计数据
func (c *WorkspaceClient) ListUserUsageMetricInvoker(request *model.ListUserUsageMetricRequest) *ListUserUsageMetricInvoker {
	requestDef := GenReqDefForListUserUsageMetric()
	return &ListUserUsageMetricInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGrowthRate 查询指标环比值
//
// 查询指标环比值
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowGrowthRate(request *model.ShowGrowthRateRequest) (*model.ShowGrowthRateResponse, error) {
	requestDef := GenReqDefForShowGrowthRate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGrowthRateResponse), nil
	}
}

// ShowGrowthRateInvoker 查询指标环比值
func (c *WorkspaceClient) ShowGrowthRateInvoker(request *model.ShowGrowthRateRequest) *ShowGrowthRateInvoker {
	requestDef := GenReqDefForShowGrowthRate()
	return &ShowGrowthRateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUserAccessStages 查询接入云桌面或云应用各阶段时长数据
//
// 查询接入云桌面或云应用各阶段时长数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowUserAccessStages(request *model.ShowUserAccessStagesRequest) (*model.ShowUserAccessStagesResponse, error) {
	requestDef := GenReqDefForShowUserAccessStages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserAccessStagesResponse), nil
	}
}

// ShowUserAccessStagesInvoker 查询接入云桌面或云应用各阶段时长数据
func (c *WorkspaceClient) ShowUserAccessStagesInvoker(request *model.ShowUserAccessStagesRequest) *ShowUserAccessStagesInvoker {
	requestDef := GenReqDefForShowUserAccessStages()
	return &ShowUserAccessStagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMetricNotifyRule 更新通知规则
//
// 更新对应指标的通知规则;对应指标满足相应的规则条件时发送通知
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateMetricNotifyRule(request *model.UpdateMetricNotifyRuleRequest) (*model.UpdateMetricNotifyRuleResponse, error) {
	requestDef := GenReqDefForUpdateMetricNotifyRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMetricNotifyRuleResponse), nil
	}
}

// UpdateMetricNotifyRuleInvoker 更新通知规则
func (c *WorkspaceClient) UpdateMetricNotifyRuleInvoker(request *model.UpdateMetricNotifyRuleRequest) *UpdateMetricNotifyRuleInvoker {
	requestDef := GenReqDefForUpdateMetricNotifyRule()
	return &UpdateMetricNotifyRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAvailableIp 根据子网id查询该子网下可用的ip
//
// 根据子网id查询该子网下可用的ip。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowAvailableIp(request *model.ShowAvailableIpRequest) (*model.ShowAvailableIpResponse, error) {
	requestDef := GenReqDefForShowAvailableIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAvailableIpResponse), nil
	}
}

// ShowAvailableIpInvoker 根据子网id查询该子网下可用的ip
func (c *WorkspaceClient) ShowAvailableIpInvoker(request *model.ShowAvailableIpRequest) *ShowAvailableIpInvoker {
	requestDef := GenReqDefForShowAvailableIp()
	return &ShowAvailableIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTenantConfigs 查询租户个性配置列表
//
// 查询租户个性配置列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListTenantConfigs(request *model.ListTenantConfigsRequest) (*model.ListTenantConfigsResponse, error) {
	requestDef := GenReqDefForListTenantConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTenantConfigsResponse), nil
	}
}

// ListTenantConfigsInvoker 查询租户个性配置列表
func (c *WorkspaceClient) ListTenantConfigsInvoker(request *model.ListTenantConfigsRequest) *ListTenantConfigsInvoker {
	requestDef := GenReqDefForListTenantConfigs()
	return &ListTenantConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTenantConfig 修改租户个性配置
//
// 租户个性配置修改。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateTenantConfig(request *model.UpdateTenantConfigRequest) (*model.UpdateTenantConfigResponse, error) {
	requestDef := GenReqDefForUpdateTenantConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTenantConfigResponse), nil
	}
}

// UpdateTenantConfigInvoker 修改租户个性配置
func (c *WorkspaceClient) UpdateTenantConfigInvoker(request *model.UpdateTenantConfigRequest) *UpdateTenantConfigInvoker {
	requestDef := GenReqDefForUpdateTenantConfig()
	return &UpdateTenantConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTerminalsBindingDesktops 增加终端与桌面绑定配置
//
// 增加终端与桌面绑定配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreateTerminalsBindingDesktops(request *model.CreateTerminalsBindingDesktopsRequest) (*model.CreateTerminalsBindingDesktopsResponse, error) {
	requestDef := GenReqDefForCreateTerminalsBindingDesktops()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTerminalsBindingDesktopsResponse), nil
	}
}

// CreateTerminalsBindingDesktopsInvoker 增加终端与桌面绑定配置
func (c *WorkspaceClient) CreateTerminalsBindingDesktopsInvoker(request *model.CreateTerminalsBindingDesktopsRequest) *CreateTerminalsBindingDesktopsInvoker {
	requestDef := GenReqDefForCreateTerminalsBindingDesktops()
	return &CreateTerminalsBindingDesktopsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTerminalsBindingDesktops 删除终端与桌面绑定配置
//
// 删除终端与桌面绑定配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) DeleteTerminalsBindingDesktops(request *model.DeleteTerminalsBindingDesktopsRequest) (*model.DeleteTerminalsBindingDesktopsResponse, error) {
	requestDef := GenReqDefForDeleteTerminalsBindingDesktops()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTerminalsBindingDesktopsResponse), nil
	}
}

// DeleteTerminalsBindingDesktopsInvoker 删除终端与桌面绑定配置
func (c *WorkspaceClient) DeleteTerminalsBindingDesktopsInvoker(request *model.DeleteTerminalsBindingDesktopsRequest) *DeleteTerminalsBindingDesktopsInvoker {
	requestDef := GenReqDefForDeleteTerminalsBindingDesktops()
	return &DeleteTerminalsBindingDesktopsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTerminalsBindingDesktops 查询终端与桌面绑定配置列表
//
// 查询终端与桌面绑定配置列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListTerminalsBindingDesktops(request *model.ListTerminalsBindingDesktopsRequest) (*model.ListTerminalsBindingDesktopsResponse, error) {
	requestDef := GenReqDefForListTerminalsBindingDesktops()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTerminalsBindingDesktopsResponse), nil
	}
}

// ListTerminalsBindingDesktopsInvoker 查询终端与桌面绑定配置列表
func (c *WorkspaceClient) ListTerminalsBindingDesktopsInvoker(request *model.ListTerminalsBindingDesktopsRequest) *ListTerminalsBindingDesktopsInvoker {
	requestDef := GenReqDefForListTerminalsBindingDesktops()
	return &ListTerminalsBindingDesktopsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTerminalsBindingDesktopsConfig 查询终端与桌面绑定的开关配置信息
//
// 查询终端与桌面绑定的开关配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListTerminalsBindingDesktopsConfig(request *model.ListTerminalsBindingDesktopsConfigRequest) (*model.ListTerminalsBindingDesktopsConfigResponse, error) {
	requestDef := GenReqDefForListTerminalsBindingDesktopsConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTerminalsBindingDesktopsConfigResponse), nil
	}
}

// ListTerminalsBindingDesktopsConfigInvoker 查询终端与桌面绑定的开关配置信息
func (c *WorkspaceClient) ListTerminalsBindingDesktopsConfigInvoker(request *model.ListTerminalsBindingDesktopsConfigRequest) *ListTerminalsBindingDesktopsConfigInvoker {
	requestDef := GenReqDefForListTerminalsBindingDesktopsConfig()
	return &ListTerminalsBindingDesktopsConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTerminalsBindingDesktops 修改终端与桌面绑定配置
//
// 修改终端与桌面绑定配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateTerminalsBindingDesktops(request *model.UpdateTerminalsBindingDesktopsRequest) (*model.UpdateTerminalsBindingDesktopsResponse, error) {
	requestDef := GenReqDefForUpdateTerminalsBindingDesktops()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTerminalsBindingDesktopsResponse), nil
	}
}

// UpdateTerminalsBindingDesktopsInvoker 修改终端与桌面绑定配置
func (c *WorkspaceClient) UpdateTerminalsBindingDesktopsInvoker(request *model.UpdateTerminalsBindingDesktopsRequest) *UpdateTerminalsBindingDesktopsInvoker {
	requestDef := GenReqDefForUpdateTerminalsBindingDesktops()
	return &UpdateTerminalsBindingDesktopsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTerminalsBindingDesktopsConfig 设置终端与桌面绑定的开关配置
//
// 设置终端与桌面绑定的开关配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateTerminalsBindingDesktopsConfig(request *model.UpdateTerminalsBindingDesktopsConfigRequest) (*model.UpdateTerminalsBindingDesktopsConfigResponse, error) {
	requestDef := GenReqDefForUpdateTerminalsBindingDesktopsConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTerminalsBindingDesktopsConfigResponse), nil
	}
}

// UpdateTerminalsBindingDesktopsConfigInvoker 设置终端与桌面绑定的开关配置
func (c *WorkspaceClient) UpdateTerminalsBindingDesktopsConfigInvoker(request *model.UpdateTerminalsBindingDesktopsConfigRequest) *UpdateTerminalsBindingDesktopsConfigInvoker {
	requestDef := GenReqDefForUpdateTerminalsBindingDesktopsConfig()
	return &UpdateTerminalsBindingDesktopsConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateUsers 批量创建用户
//
// 该接口用于批量创建用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchCreateUsers(request *model.BatchCreateUsersRequest) (*model.BatchCreateUsersResponse, error) {
	requestDef := GenReqDefForBatchCreateUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateUsersResponse), nil
	}
}

// BatchCreateUsersInvoker 批量创建用户
func (c *WorkspaceClient) BatchCreateUsersInvoker(request *model.BatchCreateUsersRequest) *BatchCreateUsersInvoker {
	requestDef := GenReqDefForBatchCreateUsers()
	return &BatchCreateUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteOtpDevices 解绑OTP设备
//
// 该接口用于解绑用户的OTP设备。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchDeleteOtpDevices(request *model.BatchDeleteOtpDevicesRequest) (*model.BatchDeleteOtpDevicesResponse, error) {
	requestDef := GenReqDefForBatchDeleteOtpDevices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteOtpDevicesResponse), nil
	}
}

// BatchDeleteOtpDevicesInvoker 解绑OTP设备
func (c *WorkspaceClient) BatchDeleteOtpDevicesInvoker(request *model.BatchDeleteOtpDevicesRequest) *BatchDeleteOtpDevicesInvoker {
	requestDef := GenReqDefForBatchDeleteOtpDevices()
	return &BatchDeleteOtpDevicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteUser 批量删除用户
//
// 该接口用于批量删除桌面用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchDeleteUser(request *model.BatchDeleteUserRequest) (*model.BatchDeleteUserResponse, error) {
	requestDef := GenReqDefForBatchDeleteUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteUserResponse), nil
	}
}

// BatchDeleteUserInvoker 批量删除用户
func (c *WorkspaceClient) BatchDeleteUserInvoker(request *model.BatchDeleteUserRequest) *BatchDeleteUserInvoker {
	requestDef := GenReqDefForBatchDeleteUser()
	return &BatchDeleteUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeUserStatus 操作用户
//
// 该接口用于操作用户，包含三种操作：锁定、解锁和重置密码（重置密码建议使用/v2/{project_id}/users/{user_id}/random-password接口，在没有通知方式的情况下必须使用/v2/{project_id}/users/{user_id}/random-password接口）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ChangeUserStatus(request *model.ChangeUserStatusRequest) (*model.ChangeUserStatusResponse, error) {
	requestDef := GenReqDefForChangeUserStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeUserStatusResponse), nil
	}
}

// ChangeUserStatusInvoker 操作用户
func (c *WorkspaceClient) ChangeUserStatusInvoker(request *model.ChangeUserStatusRequest) *ChangeUserStatusInvoker {
	requestDef := GenReqDefForChangeUserStatus()
	return &ChangeUserStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDesktopUser 创建用户
//
// 该接口用于创建桌面用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreateDesktopUser(request *model.CreateDesktopUserRequest) (*model.CreateDesktopUserResponse, error) {
	requestDef := GenReqDefForCreateDesktopUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDesktopUserResponse), nil
	}
}

// CreateDesktopUserInvoker 创建用户
func (c *WorkspaceClient) CreateDesktopUserInvoker(request *model.CreateDesktopUserRequest) *CreateDesktopUserInvoker {
	requestDef := GenReqDefForCreateDesktopUser()
	return &CreateDesktopUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteUser 删除指定用户
//
// 删除指定的桌面用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) DeleteUser(request *model.DeleteUserRequest) (*model.DeleteUserResponse, error) {
	requestDef := GenReqDefForDeleteUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteUserResponse), nil
	}
}

// DeleteUserInvoker 删除指定用户
func (c *WorkspaceClient) DeleteUserInvoker(request *model.DeleteUserRequest) *DeleteUserInvoker {
	requestDef := GenReqDefForDeleteUser()
	return &DeleteUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOtpDevicesByUserId 查询OTP设备
//
// 该接口用于查询相应用户下面的OTP设备。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListOtpDevicesByUserId(request *model.ListOtpDevicesByUserIdRequest) (*model.ListOtpDevicesByUserIdResponse, error) {
	requestDef := GenReqDefForListOtpDevicesByUserId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOtpDevicesByUserIdResponse), nil
	}
}

// ListOtpDevicesByUserIdInvoker 查询OTP设备
func (c *WorkspaceClient) ListOtpDevicesByUserIdInvoker(request *model.ListOtpDevicesByUserIdRequest) *ListOtpDevicesByUserIdInvoker {
	requestDef := GenReqDefForListOtpDevicesByUserId()
	return &ListOtpDevicesByUserIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserDetail 查询用户详情信息
//
// 该接口用于查询指定的桌面用户详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListUserDetail(request *model.ListUserDetailRequest) (*model.ListUserDetailResponse, error) {
	requestDef := GenReqDefForListUserDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserDetailResponse), nil
	}
}

// ListUserDetailInvoker 查询用户详情信息
func (c *WorkspaceClient) ListUserDetailInvoker(request *model.ListUserDetailRequest) *ListUserDetailInvoker {
	requestDef := GenReqDefForListUserDetail()
	return &ListUserDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUsers 查询用户列表
//
// 该接口用于查询桌面用户列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListUsers(request *model.ListUsersRequest) (*model.ListUsersResponse, error) {
	requestDef := GenReqDefForListUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsersResponse), nil
	}
}

// ListUsersInvoker 查询用户列表
func (c *WorkspaceClient) ListUsersInvoker(request *model.ListUsersRequest) *ListUsersInvoker {
	requestDef := GenReqDefForListUsers()
	return &ListUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetRandomPassword 给用户重置随机密码
//
// 该接口用于给用户重置一个密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ResetRandomPassword(request *model.ResetRandomPasswordRequest) (*model.ResetRandomPasswordResponse, error) {
	requestDef := GenReqDefForResetRandomPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetRandomPasswordResponse), nil
	}
}

// ResetRandomPasswordInvoker 给用户重置随机密码
func (c *WorkspaceClient) ResetRandomPasswordInvoker(request *model.ResetRandomPasswordRequest) *ResetRandomPasswordInvoker {
	requestDef := GenReqDefForResetRandomPassword()
	return &ResetRandomPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendEmail 重新发送邮件
//
// 该接口用于重新发送邮件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) SendEmail(request *model.SendEmailRequest) (*model.SendEmailResponse, error) {
	requestDef := GenReqDefForSendEmail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendEmailResponse), nil
	}
}

// SendEmailInvoker 重新发送邮件
func (c *WorkspaceClient) SendEmailInvoker(request *model.SendEmailRequest) *SendEmailInvoker {
	requestDef := GenReqDefForSendEmail()
	return &SendEmailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUserInfo 修改用户信息
//
// 该接口用于修改桌面用户信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateUserInfo(request *model.UpdateUserInfoRequest) (*model.UpdateUserInfoResponse, error) {
	requestDef := GenReqDefForUpdateUserInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserInfoResponse), nil
	}
}

// UpdateUserInfoInvoker 修改用户信息
func (c *WorkspaceClient) UpdateUserInfoInvoker(request *model.UpdateUserInfoRequest) *UpdateUserInfoInvoker {
	requestDef := GenReqDefForUpdateUserInfo()
	return &UpdateUserInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserEvents 查询用户事件
//
// 查询用户事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListUserEvents(request *model.ListUserEventsRequest) (*model.ListUserEventsResponse, error) {
	requestDef := GenReqDefForListUserEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserEventsResponse), nil
	}
}

// ListUserEventsInvoker 查询用户事件
func (c *WorkspaceClient) ListUserEventsInvoker(request *model.ListUserEventsRequest) *ListUserEventsInvoker {
	requestDef := GenReqDefForListUserEvents()
	return &ListUserEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddDesktopVolumes 增加桌面磁盘
//
// 给单个桌面增加磁盘。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) AddDesktopVolumes(request *model.AddDesktopVolumesRequest) (*model.AddDesktopVolumesResponse, error) {
	requestDef := GenReqDefForAddDesktopVolumes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddDesktopVolumesResponse), nil
	}
}

// AddDesktopVolumesInvoker 增加桌面磁盘
func (c *WorkspaceClient) AddDesktopVolumesInvoker(request *model.AddDesktopVolumesRequest) *AddDesktopVolumesInvoker {
	requestDef := GenReqDefForAddDesktopVolumes()
	return &AddDesktopVolumesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddVolumes 增加桌面磁盘
//
// 增加桌面磁盘。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) AddVolumes(request *model.AddVolumesRequest) (*model.AddVolumesResponse, error) {
	requestDef := GenReqDefForAddVolumes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddVolumesResponse), nil
	}
}

// AddVolumesInvoker 增加桌面磁盘
func (c *WorkspaceClient) AddVolumesInvoker(request *model.AddVolumesRequest) *AddVolumesInvoker {
	requestDef := GenReqDefForAddVolumes()
	return &AddVolumesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDesktopVolumes 删除桌面数据盘
//
// 删除桌面数据盘，删除后无法恢复。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) DeleteDesktopVolumes(request *model.DeleteDesktopVolumesRequest) (*model.DeleteDesktopVolumesResponse, error) {
	requestDef := GenReqDefForDeleteDesktopVolumes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDesktopVolumesResponse), nil
	}
}

// DeleteDesktopVolumesInvoker 删除桌面数据盘
func (c *WorkspaceClient) DeleteDesktopVolumesInvoker(request *model.DeleteDesktopVolumesRequest) *DeleteDesktopVolumesInvoker {
	requestDef := GenReqDefForDeleteDesktopVolumes()
	return &DeleteDesktopVolumesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExpandDesktopVolume 扩容磁盘
//
// 扩容磁盘
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ExpandDesktopVolume(request *model.ExpandDesktopVolumeRequest) (*model.ExpandDesktopVolumeResponse, error) {
	requestDef := GenReqDefForExpandDesktopVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandDesktopVolumeResponse), nil
	}
}

// ExpandDesktopVolumeInvoker 扩容磁盘
func (c *WorkspaceClient) ExpandDesktopVolumeInvoker(request *model.ExpandDesktopVolumeRequest) *ExpandDesktopVolumeInvoker {
	requestDef := GenReqDefForExpandDesktopVolume()
	return &ExpandDesktopVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExpandVolumes 扩容桌面磁盘
//
// 扩容桌面磁盘。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ExpandVolumes(request *model.ExpandVolumesRequest) (*model.ExpandVolumesResponse, error) {
	requestDef := GenReqDefForExpandVolumes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandVolumesResponse), nil
	}
}

// ExpandVolumesInvoker 扩容桌面磁盘
func (c *WorkspaceClient) ExpandVolumesInvoker(request *model.ExpandVolumesRequest) *ExpandVolumesInvoker {
	requestDef := GenReqDefForExpandVolumes()
	return &ExpandVolumesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVolumeProductInfo 查询磁盘产品信息
//
// 查询磁盘产品信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListVolumeProductInfo(request *model.ListVolumeProductInfoRequest) (*model.ListVolumeProductInfoResponse, error) {
	requestDef := GenReqDefForListVolumeProductInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVolumeProductInfoResponse), nil
	}
}

// ListVolumeProductInfoInvoker 查询磁盘产品信息
func (c *WorkspaceClient) ListVolumeProductInfoInvoker(request *model.ListVolumeProductInfoRequest) *ListVolumeProductInfoInvoker {
	requestDef := GenReqDefForListVolumeProductInfo()
	return &ListVolumeProductInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ApplyWorkspace 开通云办公服务
//
// 该接口用于开通云办公服务。
//
// 作为异步接口，调用成功说明云办公服务后台收到了开通请求，但服务是否开通成功需要通过任务查询接口(GET /v2/{project_id}/workspace-sub-jobs)查询该任务的执行状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ApplyWorkspace(request *model.ApplyWorkspaceRequest) (*model.ApplyWorkspaceResponse, error) {
	requestDef := GenReqDefForApplyWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ApplyWorkspaceResponse), nil
	}
}

// ApplyWorkspaceInvoker 开通云办公服务
func (c *WorkspaceClient) ApplyWorkspaceInvoker(request *model.ApplyWorkspaceRequest) *ApplyWorkspaceInvoker {
	requestDef := GenReqDefForApplyWorkspace()
	return &ApplyWorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelWorkspace 注销云办公服务
//
// 该接口用于注销云办公服务。注销前请确保当前用户下的云桌面已经删除，注销后无法恢复。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CancelWorkspace(request *model.CancelWorkspaceRequest) (*model.CancelWorkspaceResponse, error) {
	requestDef := GenReqDefForCancelWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelWorkspaceResponse), nil
	}
}

// CancelWorkspaceInvoker 注销云办公服务
func (c *WorkspaceClient) CancelWorkspaceInvoker(request *model.CancelWorkspaceRequest) *CancelWorkspaceInvoker {
	requestDef := GenReqDefForCancelWorkspace()
	return &CancelWorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkspaces 查询云办公服务详情
//
// 该接口用于查询云办公服务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListWorkspaces(request *model.ListWorkspacesRequest) (*model.ListWorkspacesResponse, error) {
	requestDef := GenReqDefForListWorkspaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkspacesResponse), nil
	}
}

// ListWorkspacesInvoker 查询云办公服务详情
func (c *WorkspaceClient) ListWorkspacesInvoker(request *model.ListWorkspacesRequest) *ListWorkspacesInvoker {
	requestDef := GenReqDefForListWorkspaces()
	return &ListWorkspacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkspaceLock 查询云办公服务是否被锁定
//
// 查询云办公服务是否被锁定。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowWorkspaceLock(request *model.ShowWorkspaceLockRequest) (*model.ShowWorkspaceLockResponse, error) {
	requestDef := GenReqDefForShowWorkspaceLock()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkspaceLockResponse), nil
	}
}

// ShowWorkspaceLockInvoker 查询云办公服务是否被锁定
func (c *WorkspaceClient) ShowWorkspaceLockInvoker(request *model.ShowWorkspaceLockRequest) *ShowWorkspaceLockInvoker {
	requestDef := GenReqDefForShowWorkspaceLock()
	return &ShowWorkspaceLockInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnlockWorkspace 解除云办公服务锁定状态
//
// 该接口目前支持解除云办公服务锁定状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UnlockWorkspace(request *model.UnlockWorkspaceRequest) (*model.UnlockWorkspaceResponse, error) {
	requestDef := GenReqDefForUnlockWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnlockWorkspaceResponse), nil
	}
}

// UnlockWorkspaceInvoker 解除云办公服务锁定状态
func (c *WorkspaceClient) UnlockWorkspaceInvoker(request *model.UnlockWorkspaceRequest) *UnlockWorkspaceInvoker {
	requestDef := GenReqDefForUnlockWorkspace()
	return &UnlockWorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEnterpriseId 修改企业ID
//
// 修改租户的企业ID。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateEnterpriseId(request *model.UpdateEnterpriseIdRequest) (*model.UpdateEnterpriseIdResponse, error) {
	requestDef := GenReqDefForUpdateEnterpriseId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEnterpriseIdResponse), nil
	}
}

// UpdateEnterpriseIdInvoker 修改企业ID
func (c *WorkspaceClient) UpdateEnterpriseIdInvoker(request *model.UpdateEnterpriseIdRequest) *UpdateEnterpriseIdInvoker {
	requestDef := GenReqDefForUpdateEnterpriseId()
	return &UpdateEnterpriseIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWorkspace 修改云办公服务属性
//
// 该接口目前支持修改云办公服务属性。单次请求仅支持修改一种属性类型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateWorkspace(request *model.UpdateWorkspaceRequest) (*model.UpdateWorkspaceResponse, error) {
	requestDef := GenReqDefForUpdateWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWorkspaceResponse), nil
	}
}

// UpdateWorkspaceInvoker 修改云办公服务属性
func (c *WorkspaceClient) UpdateWorkspaceInvoker(request *model.UpdateWorkspaceRequest) *UpdateWorkspaceInvoker {
	requestDef := GenReqDefForUpdateWorkspace()
	return &UpdateWorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
