package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ugo/v1/model"
)

type UgoClient struct {
	HcClient *http_client.HcHttpClient
}

func NewUgoClient(hcClient *http_client.HcHttpClient) *UgoClient {
	return &UgoClient{HcClient: hcClient}
}

func UgoClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CheckPermission 目标库权限检查。
//
// 目标库权限检查。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UgoClient) CheckPermission(request *model.CheckPermissionRequest) (*model.CheckPermissionResponse, error) {
	requestDef := GenReqDefForCheckPermission()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckPermissionResponse), nil
	}
}

// CheckPermissionInvoker 目标库权限检查。
func (c *UgoClient) CheckPermissionInvoker(request *model.CheckPermissionRequest) *CheckPermissionInvoker {
	requestDef := GenReqDefForCheckPermission()
	return &CheckPermissionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CommitSyntaxConversion 提交语法转换
//
// 提交语法转换。只有migration_project_status为\&quot;READY\&quot;时才能调用该接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UgoClient) CommitSyntaxConversion(request *model.CommitSyntaxConversionRequest) (*model.CommitSyntaxConversionResponse, error) {
	requestDef := GenReqDefForCommitSyntaxConversion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CommitSyntaxConversionResponse), nil
	}
}

// CommitSyntaxConversionInvoker 提交语法转换
func (c *UgoClient) CommitSyntaxConversionInvoker(request *model.CommitSyntaxConversionRequest) *CommitSyntaxConversionInvoker {
	requestDef := GenReqDefForCommitSyntaxConversion()
	return &CommitSyntaxConversionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CommitVerification 提交验证。
//
// 提交验证。语法转换完成后，才能调用该接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UgoClient) CommitVerification(request *model.CommitVerificationRequest) (*model.CommitVerificationResponse, error) {
	requestDef := GenReqDefForCommitVerification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CommitVerificationResponse), nil
	}
}

// CommitVerificationInvoker 提交验证。
func (c *UgoClient) CommitVerificationInvoker(request *model.CommitVerificationRequest) *CommitVerificationInvoker {
	requestDef := GenReqDefForCommitVerification()
	return &CommitVerificationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ConfirmTargetDbType 评估项目确认目标数据库类型。
//
// 评估项目确认目标数据库类型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UgoClient) ConfirmTargetDbType(request *model.ConfirmTargetDbTypeRequest) (*model.ConfirmTargetDbTypeResponse, error) {
	requestDef := GenReqDefForConfirmTargetDbType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ConfirmTargetDbTypeResponse), nil
	}
}

// ConfirmTargetDbTypeInvoker 评估项目确认目标数据库类型。
func (c *UgoClient) ConfirmTargetDbTypeInvoker(request *model.ConfirmTargetDbTypeRequest) *ConfirmTargetDbTypeInvoker {
	requestDef := GenReqDefForConfirmTargetDbType()
	return &ConfirmTargetDbTypeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEvaluationProject 创建评估项目。
//
// 创建评估项目。评估项目分2个阶段：采集、评估。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UgoClient) CreateEvaluationProject(request *model.CreateEvaluationProjectRequest) (*model.CreateEvaluationProjectResponse, error) {
	requestDef := GenReqDefForCreateEvaluationProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEvaluationProjectResponse), nil
	}
}

// CreateEvaluationProjectInvoker 创建评估项目。
func (c *UgoClient) CreateEvaluationProjectInvoker(request *model.CreateEvaluationProjectRequest) *CreateEvaluationProjectInvoker {
	requestDef := GenReqDefForCreateEvaluationProject()
	return &CreateEvaluationProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMigrationProject 创建迁移项目。
//
// 创建迁移项目。创建迁移项目需要关联状态为“COMPLETED”的评估项目。迁移项目依次经历以下几个阶段：目标库权限检查、语法转换、验证、下载迁移失败的报告、删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UgoClient) CreateMigrationProject(request *model.CreateMigrationProjectRequest) (*model.CreateMigrationProjectResponse, error) {
	requestDef := GenReqDefForCreateMigrationProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMigrationProjectResponse), nil
	}
}

// CreateMigrationProjectInvoker 创建迁移项目。
func (c *UgoClient) CreateMigrationProjectInvoker(request *model.CreateMigrationProjectRequest) *CreateMigrationProjectInvoker {
	requestDef := GenReqDefForCreateMigrationProject()
	return &CreateMigrationProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEvaluationProject 删除评估项目。
//
// 删除评估项目。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UgoClient) DeleteEvaluationProject(request *model.DeleteEvaluationProjectRequest) (*model.DeleteEvaluationProjectResponse, error) {
	requestDef := GenReqDefForDeleteEvaluationProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEvaluationProjectResponse), nil
	}
}

// DeleteEvaluationProjectInvoker 删除评估项目。
func (c *UgoClient) DeleteEvaluationProjectInvoker(request *model.DeleteEvaluationProjectRequest) *DeleteEvaluationProjectInvoker {
	requestDef := GenReqDefForDeleteEvaluationProject()
	return &DeleteEvaluationProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMigrationProject 删除迁移项目
//
// 删除迁移项目。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UgoClient) DeleteMigrationProject(request *model.DeleteMigrationProjectRequest) (*model.DeleteMigrationProjectResponse, error) {
	requestDef := GenReqDefForDeleteMigrationProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMigrationProjectResponse), nil
	}
}

// DeleteMigrationProjectInvoker 删除迁移项目
func (c *UgoClient) DeleteMigrationProjectInvoker(request *model.DeleteMigrationProjectRequest) *DeleteMigrationProjectInvoker {
	requestDef := GenReqDefForDeleteMigrationProject()
	return &DeleteMigrationProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadFailureReport 下载迁移错误报告。
//
// 下载迁移错误报告。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UgoClient) DownloadFailureReport(request *model.DownloadFailureReportRequest) (*model.DownloadFailureReportResponse, error) {
	requestDef := GenReqDefForDownloadFailureReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadFailureReportResponse), nil
	}
}

// DownloadFailureReportInvoker 下载迁移错误报告。
func (c *UgoClient) DownloadFailureReportInvoker(request *model.DownloadFailureReportRequest) *DownloadFailureReportInvoker {
	requestDef := GenReqDefForDownloadFailureReport()
	return &DownloadFailureReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEvaluationProjects 查询评估项目列表。
//
// 查询评估项目列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UgoClient) ListEvaluationProjects(request *model.ListEvaluationProjectsRequest) (*model.ListEvaluationProjectsResponse, error) {
	requestDef := GenReqDefForListEvaluationProjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEvaluationProjectsResponse), nil
	}
}

// ListEvaluationProjectsInvoker 查询评估项目列表。
func (c *UgoClient) ListEvaluationProjectsInvoker(request *model.ListEvaluationProjectsRequest) *ListEvaluationProjectsInvoker {
	requestDef := GenReqDefForListEvaluationProjects()
	return &ListEvaluationProjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMigrationProjects 查询迁移项目列表。
//
// 查询迁移项目列表。创建迁移项目之后，调用该接口，根据项目名称，获取项目ID。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UgoClient) ListMigrationProjects(request *model.ListMigrationProjectsRequest) (*model.ListMigrationProjectsResponse, error) {
	requestDef := GenReqDefForListMigrationProjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMigrationProjectsResponse), nil
	}
}

// ListMigrationProjectsInvoker 查询迁移项目列表。
func (c *UgoClient) ListMigrationProjectsInvoker(request *model.ListMigrationProjectsRequest) *ListMigrationProjectsInvoker {
	requestDef := GenReqDefForListMigrationProjects()
	return &ListMigrationProjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPermissionCheckResult 查询权限检查结果。
//
// 查询权限检查结果。permission_check_status 为 \&quot;SUCCESS\&quot; 或者 \&quot;FAILED\&quot; 时，才能调用该接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UgoClient) ListPermissionCheckResult(request *model.ListPermissionCheckResultRequest) (*model.ListPermissionCheckResultResponse, error) {
	requestDef := GenReqDefForListPermissionCheckResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPermissionCheckResultResponse), nil
	}
}

// ListPermissionCheckResultInvoker 查询权限检查结果。
func (c *UgoClient) ListPermissionCheckResultInvoker(request *model.ListPermissionCheckResultRequest) *ListPermissionCheckResultInvoker {
	requestDef := GenReqDefForListPermissionCheckResult()
	return &ListPermissionCheckResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuotas 查询配额。
//
// 查询单租户的配额，包括评估项目配额、迁移项目配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UgoClient) ListQuotas(request *model.ListQuotasRequest) (*model.ListQuotasResponse, error) {
	requestDef := GenReqDefForListQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotasResponse), nil
	}
}

// ListQuotasInvoker 查询配额。
func (c *UgoClient) ListQuotasInvoker(request *model.ListQuotasRequest) *ListQuotasInvoker {
	requestDef := GenReqDefForListQuotas()
	return &ListQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSyntaxConversionProgress 查询语法转换的进度。
//
// 查询语法转换的进度。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UgoClient) ListSyntaxConversionProgress(request *model.ListSyntaxConversionProgressRequest) (*model.ListSyntaxConversionProgressResponse, error) {
	requestDef := GenReqDefForListSyntaxConversionProgress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSyntaxConversionProgressResponse), nil
	}
}

// ListSyntaxConversionProgressInvoker 查询语法转换的进度。
func (c *UgoClient) ListSyntaxConversionProgressInvoker(request *model.ListSyntaxConversionProgressRequest) *ListSyntaxConversionProgressInvoker {
	requestDef := GenReqDefForListSyntaxConversionProgress()
	return &ListSyntaxConversionProgressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVerificationProgress 查询验证进度。
//
// 查询验证进度。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UgoClient) ListVerificationProgress(request *model.ListVerificationProgressRequest) (*model.ListVerificationProgressResponse, error) {
	requestDef := GenReqDefForListVerificationProgress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVerificationProgressResponse), nil
	}
}

// ListVerificationProgressInvoker 查询验证进度。
func (c *UgoClient) ListVerificationProgressInvoker(request *model.ListVerificationProgressRequest) *ListVerificationProgressInvoker {
	requestDef := GenReqDefForListVerificationProgress()
	return &ListVerificationProgressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEvaluationProjectDetail 查询评估项目详情。
//
// 查询评估项目详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UgoClient) ShowEvaluationProjectDetail(request *model.ShowEvaluationProjectDetailRequest) (*model.ShowEvaluationProjectDetailResponse, error) {
	requestDef := GenReqDefForShowEvaluationProjectDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEvaluationProjectDetailResponse), nil
	}
}

// ShowEvaluationProjectDetailInvoker 查询评估项目详情。
func (c *UgoClient) ShowEvaluationProjectDetailInvoker(request *model.ShowEvaluationProjectDetailRequest) *ShowEvaluationProjectDetailInvoker {
	requestDef := GenReqDefForShowEvaluationProjectDetail()
	return &ShowEvaluationProjectDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEvaluationProjectStatus 查询评估项目状态。
//
// 查询评估项目状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UgoClient) ShowEvaluationProjectStatus(request *model.ShowEvaluationProjectStatusRequest) (*model.ShowEvaluationProjectStatusResponse, error) {
	requestDef := GenReqDefForShowEvaluationProjectStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEvaluationProjectStatusResponse), nil
	}
}

// ShowEvaluationProjectStatusInvoker 查询评估项目状态。
func (c *UgoClient) ShowEvaluationProjectStatusInvoker(request *model.ShowEvaluationProjectStatusRequest) *ShowEvaluationProjectStatusInvoker {
	requestDef := GenReqDefForShowEvaluationProjectStatus()
	return &ShowEvaluationProjectStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMigrationProjectDetail 查询迁移项目详情。
//
// 查询迁移项目详情。只有migration_project_status为\&quot;READY\&quot;时才能调用该接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UgoClient) ShowMigrationProjectDetail(request *model.ShowMigrationProjectDetailRequest) (*model.ShowMigrationProjectDetailResponse, error) {
	requestDef := GenReqDefForShowMigrationProjectDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMigrationProjectDetailResponse), nil
	}
}

// ShowMigrationProjectDetailInvoker 查询迁移项目详情。
func (c *UgoClient) ShowMigrationProjectDetailInvoker(request *model.ShowMigrationProjectDetailRequest) *ShowMigrationProjectDetailInvoker {
	requestDef := GenReqDefForShowMigrationProjectDetail()
	return &ShowMigrationProjectDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMigrationProjectStatus 查询迁移项目状态。
//
// 查询迁移项目状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UgoClient) ShowMigrationProjectStatus(request *model.ShowMigrationProjectStatusRequest) (*model.ShowMigrationProjectStatusResponse, error) {
	requestDef := GenReqDefForShowMigrationProjectStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMigrationProjectStatusResponse), nil
	}
}

// ShowMigrationProjectStatusInvoker 查询迁移项目状态。
func (c *UgoClient) ShowMigrationProjectStatusInvoker(request *model.ShowMigrationProjectStatusRequest) *ShowMigrationProjectStatusInvoker {
	requestDef := GenReqDefForShowMigrationProjectStatus()
	return &ShowMigrationProjectStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiVersions 查询API版本信息列表。
//
// 查询API版本信息列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UgoClient) ListApiVersions(request *model.ListApiVersionsRequest) (*model.ListApiVersionsResponse, error) {
	requestDef := GenReqDefForListApiVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionsResponse), nil
	}
}

// ListApiVersionsInvoker 查询API版本信息列表。
func (c *UgoClient) ListApiVersionsInvoker(request *model.ListApiVersionsRequest) *ListApiVersionsInvoker {
	requestDef := GenReqDefForListApiVersions()
	return &ListApiVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApiVersionInfo 查询指定版本号的API版本信息。
//
// 查询指定版本号的API版本信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UgoClient) ShowApiVersionInfo(request *model.ShowApiVersionInfoRequest) (*model.ShowApiVersionInfoResponse, error) {
	requestDef := GenReqDefForShowApiVersionInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApiVersionInfoResponse), nil
	}
}

// ShowApiVersionInfoInvoker 查询指定版本号的API版本信息。
func (c *UgoClient) ShowApiVersionInfoInvoker(request *model.ShowApiVersionInfoRequest) *ShowApiVersionInfoInvoker {
	requestDef := GenReqDefForShowApiVersionInfo()
	return &ShowApiVersionInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunSqlConversion SQL语句转换。
//
// SQL语句转换。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *UgoClient) RunSqlConversion(request *model.RunSqlConversionRequest) (*model.RunSqlConversionResponse, error) {
	requestDef := GenReqDefForRunSqlConversion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunSqlConversionResponse), nil
	}
}

// RunSqlConversionInvoker SQL语句转换。
func (c *UgoClient) RunSqlConversionInvoker(request *model.RunSqlConversionRequest) *RunSqlConversionInvoker {
	requestDef := GenReqDefForRunSqlConversion()
	return &RunSqlConversionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
