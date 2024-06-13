package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cloudtest/v1/model"
)

type CloudtestClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCloudtestClient(hcClient *httpclient.HcHttpClient) *CloudtestClient {
	return &CloudtestClient{HcClient: hcClient}
}

func CloudtestClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// BatchAddRelationsByOneCase 添加需求/缺陷和多个用例关联关系
//
// 添加需求/缺陷和多个用例关联关系
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) BatchAddRelationsByOneCase(request *model.BatchAddRelationsByOneCaseRequest) (*model.BatchAddRelationsByOneCaseResponse, error) {
	requestDef := GenReqDefForBatchAddRelationsByOneCase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddRelationsByOneCaseResponse), nil
	}
}

// BatchAddRelationsByOneCaseInvoker 添加需求/缺陷和多个用例关联关系
func (c *CloudtestClient) BatchAddRelationsByOneCaseInvoker(request *model.BatchAddRelationsByOneCaseRequest) *BatchAddRelationsByOneCaseInvoker {
	requestDef := GenReqDefForBatchAddRelationsByOneCase()
	return &BatchAddRelationsByOneCaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteTestCase 批量删除自定义测试服务类型用例
//
// 批量删除自定义测试服务类型用例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) BatchDeleteTestCase(request *model.BatchDeleteTestCaseRequest) (*model.BatchDeleteTestCaseResponse, error) {
	requestDef := GenReqDefForBatchDeleteTestCase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteTestCaseResponse), nil
	}
}

// BatchDeleteTestCaseInvoker 批量删除自定义测试服务类型用例
func (c *CloudtestClient) BatchDeleteTestCaseInvoker(request *model.BatchDeleteTestCaseRequest) *BatchDeleteTestCaseInvoker {
	requestDef := GenReqDefForBatchDeleteTestCase()
	return &BatchDeleteTestCaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteTestReport 根据测试报告uri列表，删除测试报告
//
// 根据测试报告uri列表，删除测试报告
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) BatchDeleteTestReport(request *model.BatchDeleteTestReportRequest) (*model.BatchDeleteTestReportResponse, error) {
	requestDef := GenReqDefForBatchDeleteTestReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteTestReportResponse), nil
	}
}

// BatchDeleteTestReportInvoker 根据测试报告uri列表，删除测试报告
func (c *CloudtestClient) BatchDeleteTestReportInvoker(request *model.BatchDeleteTestReportRequest) *BatchDeleteTestReportInvoker {
	requestDef := GenReqDefForBatchDeleteTestReport()
	return &BatchDeleteTestReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckPermission 检查项目权限
//
// 检查项目权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) CheckPermission(request *model.CheckPermissionRequest) (*model.CheckPermissionResponse, error) {
	requestDef := GenReqDefForCheckPermission()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckPermissionResponse), nil
	}
}

// CheckPermissionInvoker 检查项目权限
func (c *CloudtestClient) CheckPermissionInvoker(request *model.CheckPermissionRequest) *CheckPermissionInvoker {
	requestDef := GenReqDefForCheckPermission()
	return &CheckPermissionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePlan 项目下创建计划
//
// 项目下创建计划
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) CreatePlan(request *model.CreatePlanRequest) (*model.CreatePlanResponse, error) {
	requestDef := GenReqDefForCreatePlan()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePlanResponse), nil
	}
}

// CreatePlanInvoker 项目下创建计划
func (c *CloudtestClient) CreatePlanInvoker(request *model.CreatePlanRequest) *CreatePlanInvoker {
	requestDef := GenReqDefForCreatePlan()
	return &CreatePlanInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRelationsByOneCase 添加一个用例和多个需求/缺陷关联关系
//
// 添加一个用例和多个需求/缺陷关联关系
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) CreateRelationsByOneCase(request *model.CreateRelationsByOneCaseRequest) (*model.CreateRelationsByOneCaseResponse, error) {
	requestDef := GenReqDefForCreateRelationsByOneCase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRelationsByOneCaseResponse), nil
	}
}

// CreateRelationsByOneCaseInvoker 添加一个用例和多个需求/缺陷关联关系
func (c *CloudtestClient) CreateRelationsByOneCaseInvoker(request *model.CreateRelationsByOneCaseRequest) *CreateRelationsByOneCaseInvoker {
	requestDef := GenReqDefForCreateRelationsByOneCase()
	return &CreateRelationsByOneCaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateResourceUri 生成资源URI
//
// 生成资源URI
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) CreateResourceUri(request *model.CreateResourceUriRequest) (*model.CreateResourceUriResponse, error) {
	requestDef := GenReqDefForCreateResourceUri()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResourceUriResponse), nil
	}
}

// CreateResourceUriInvoker 生成资源URI
func (c *CloudtestClient) CreateResourceUriInvoker(request *model.CreateResourceUriRequest) *CreateResourceUriInvoker {
	requestDef := GenReqDefForCreateResourceUri()
	return &CreateResourceUriInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateService 新测试类型服务注册
//
// 通过接口CreateService注册成为自定义服务。 注册完成后界面将会出现此自定义测试类型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) CreateService(request *model.CreateServiceRequest) (*model.CreateServiceResponse, error) {
	requestDef := GenReqDefForCreateService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateServiceResponse), nil
	}
}

// CreateServiceInvoker 新测试类型服务注册
func (c *CloudtestClient) CreateServiceInvoker(request *model.CreateServiceRequest) *CreateServiceInvoker {
	requestDef := GenReqDefForCreateService()
	return &CreateServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTestCase 创建自定义测试服务类型用例
//
// 创建自定义测试服务类型用例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) CreateTestCase(request *model.CreateTestCaseRequest) (*model.CreateTestCaseResponse, error) {
	requestDef := GenReqDefForCreateTestCase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTestCaseResponse), nil
	}
}

// CreateTestCaseInvoker 创建自定义测试服务类型用例
func (c *CloudtestClient) CreateTestCaseInvoker(request *model.CreateTestCaseRequest) *CreateTestCaseInvoker {
	requestDef := GenReqDefForCreateTestCase()
	return &CreateTestCaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTestCaseInPlan 计划中批量添加测试用例
//
// 计划中批量添加测试用例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) CreateTestCaseInPlan(request *model.CreateTestCaseInPlanRequest) (*model.CreateTestCaseInPlanResponse, error) {
	requestDef := GenReqDefForCreateTestCaseInPlan()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTestCaseInPlanResponse), nil
	}
}

// CreateTestCaseInPlanInvoker 计划中批量添加测试用例
func (c *CloudtestClient) CreateTestCaseInPlanInvoker(request *model.CreateTestCaseInPlanRequest) *CreateTestCaseInPlanInvoker {
	requestDef := GenReqDefForCreateTestCaseInPlan()
	return &CreateTestCaseInPlanInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRelationsByOneCase 删除一个用例和多个需求/缺陷关联关系
//
// 删除一个用例和多个需求/缺陷关联关系
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) DeleteRelationsByOneCase(request *model.DeleteRelationsByOneCaseRequest) (*model.DeleteRelationsByOneCaseResponse, error) {
	requestDef := GenReqDefForDeleteRelationsByOneCase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRelationsByOneCaseResponse), nil
	}
}

// DeleteRelationsByOneCaseInvoker 删除一个用例和多个需求/缺陷关联关系
func (c *CloudtestClient) DeleteRelationsByOneCaseInvoker(request *model.DeleteRelationsByOneCaseRequest) *DeleteRelationsByOneCaseInvoker {
	requestDef := GenReqDefForDeleteRelationsByOneCase()
	return &DeleteRelationsByOneCaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteService 删除已注册服务
//
// 删除已注册服务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) DeleteService(request *model.DeleteServiceRequest) (*model.DeleteServiceResponse, error) {
	requestDef := GenReqDefForDeleteService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServiceResponse), nil
	}
}

// DeleteServiceInvoker 删除已注册服务
func (c *CloudtestClient) DeleteServiceInvoker(request *model.DeleteServiceRequest) *DeleteServiceInvoker {
	requestDef := GenReqDefForDeleteService()
	return &DeleteServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllBranches 获取分支列表
//
// 获取分支列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListAllBranches(request *model.ListAllBranchesRequest) (*model.ListAllBranchesResponse, error) {
	requestDef := GenReqDefForListAllBranches()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllBranchesResponse), nil
	}
}

// ListAllBranchesInvoker 获取分支列表
func (c *CloudtestClient) ListAllBranchesInvoker(request *model.ListAllBranchesRequest) *ListAllBranchesInvoker {
	requestDef := GenReqDefForListAllBranches()
	return &ListAllBranchesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllIterators 查询项目下所有迭代计划
//
// 查询项目下所有迭代计划
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListAllIterators(request *model.ListAllIteratorsRequest) (*model.ListAllIteratorsResponse, error) {
	requestDef := GenReqDefForListAllIterators()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllIteratorsResponse), nil
	}
}

// ListAllIteratorsInvoker 查询项目下所有迭代计划
func (c *CloudtestClient) ListAllIteratorsInvoker(request *model.ListAllIteratorsRequest) *ListAllIteratorsInvoker {
	requestDef := GenReqDefForListAllIterators()
	return &ListAllIteratorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAttachments 查询附件列表
//
// 查询附件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListAttachments(request *model.ListAttachmentsRequest) (*model.ListAttachmentsResponse, error) {
	requestDef := GenReqDefForListAttachments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAttachmentsResponse), nil
	}
}

// ListAttachmentsInvoker 查询附件列表
func (c *CloudtestClient) ListAttachmentsInvoker(request *model.ListAttachmentsRequest) *ListAttachmentsInvoker {
	requestDef := GenReqDefForListAttachments()
	return &ListAttachmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBranches 获取分支列表
//
// 获取分支列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListBranches(request *model.ListBranchesRequest) (*model.ListBranchesResponse, error) {
	requestDef := GenReqDefForListBranches()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBranchesResponse), nil
	}
}

// ListBranchesInvoker 获取分支列表
func (c *CloudtestClient) ListBranchesInvoker(request *model.ListBranchesRequest) *ListBranchesInvoker {
	requestDef := GenReqDefForListBranches()
	return &ListBranchesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIssueTree 查询需求树
//
// 查询需求树
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListIssueTree(request *model.ListIssueTreeRequest) (*model.ListIssueTreeResponse, error) {
	requestDef := GenReqDefForListIssueTree()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIssueTreeResponse), nil
	}
}

// ListIssueTreeInvoker 查询需求树
func (c *CloudtestClient) ListIssueTreeInvoker(request *model.ListIssueTreeRequest) *ListIssueTreeInvoker {
	requestDef := GenReqDefForListIssueTree()
	return &ListIssueTreeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectFieldConfigs 查询项目字段配置
//
// 查询项目字段配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListProjectFieldConfigs(request *model.ListProjectFieldConfigsRequest) (*model.ListProjectFieldConfigsResponse, error) {
	requestDef := GenReqDefForListProjectFieldConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectFieldConfigsResponse), nil
	}
}

// ListProjectFieldConfigsInvoker 查询项目字段配置
func (c *CloudtestClient) ListProjectFieldConfigsInvoker(request *model.ListProjectFieldConfigsRequest) *ListProjectFieldConfigsInvoker {
	requestDef := GenReqDefForListProjectFieldConfigs()
	return &ListProjectFieldConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectTestCaseFields 获取项目测试用例自定义字段列表
//
// 获取项目测试用例自定义字段列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListProjectTestCaseFields(request *model.ListProjectTestCaseFieldsRequest) (*model.ListProjectTestCaseFieldsResponse, error) {
	requestDef := GenReqDefForListProjectTestCaseFields()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectTestCaseFieldsResponse), nil
	}
}

// ListProjectTestCaseFieldsInvoker 获取项目测试用例自定义字段列表
func (c *CloudtestClient) ListProjectTestCaseFieldsInvoker(request *model.ListProjectTestCaseFieldsRequest) *ListProjectTestCaseFieldsInvoker {
	requestDef := GenReqDefForListProjectTestCaseFields()
	return &ListProjectTestCaseFieldsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListReports 页面报表展示
//
// 页面报表展示
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListReports(request *model.ListReportsRequest) (*model.ListReportsResponse, error) {
	requestDef := GenReqDefForListReports()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListReportsResponse), nil
	}
}

// ListReportsInvoker 页面报表展示
func (c *CloudtestClient) ListReportsInvoker(request *model.ListReportsRequest) *ListReportsInvoker {
	requestDef := GenReqDefForListReports()
	return &ListReportsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourcePools 获取资源池列表
//
// 获取资源池列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListResourcePools(request *model.ListResourcePoolsRequest) (*model.ListResourcePoolsResponse, error) {
	requestDef := GenReqDefForListResourcePools()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourcePoolsResponse), nil
	}
}

// ListResourcePoolsInvoker 获取资源池列表
func (c *CloudtestClient) ListResourcePoolsInvoker(request *model.ListResourcePoolsRequest) *ListResourcePoolsInvoker {
	requestDef := GenReqDefForListResourcePools()
	return &ListResourcePoolsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTaskTestCases 查询用例关联的测试任务列表
//
// 查询用例关联的测试任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListTaskTestCases(request *model.ListTaskTestCasesRequest) (*model.ListTaskTestCasesResponse, error) {
	requestDef := GenReqDefForListTaskTestCases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTaskTestCasesResponse), nil
	}
}

// ListTaskTestCasesInvoker 查询用例关联的测试任务列表
func (c *CloudtestClient) ListTaskTestCasesInvoker(request *model.ListTaskTestCasesRequest) *ListTaskTestCasesInvoker {
	requestDef := GenReqDefForListTaskTestCases()
	return &ListTaskTestCasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTestCaseHistories 查询用例修改历史记录
//
// 查询用例修改历史记录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListTestCaseHistories(request *model.ListTestCaseHistoriesRequest) (*model.ListTestCaseHistoriesResponse, error) {
	requestDef := GenReqDefForListTestCaseHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTestCaseHistoriesResponse), nil
	}
}

// ListTestCaseHistoriesInvoker 查询用例修改历史记录
func (c *CloudtestClient) ListTestCaseHistoriesInvoker(request *model.ListTestCaseHistoriesRequest) *ListTestCaseHistoriesInvoker {
	requestDef := GenReqDefForListTestCaseHistories()
	return &ListTestCaseHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTestCases 查询用例列表
//
// 查询用例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListTestCases(request *model.ListTestCasesRequest) (*model.ListTestCasesResponse, error) {
	requestDef := GenReqDefForListTestCases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTestCasesResponse), nil
	}
}

// ListTestCasesInvoker 查询用例列表
func (c *CloudtestClient) ListTestCasesInvoker(request *model.ListTestCasesRequest) *ListTestCasesInvoker {
	requestDef := GenReqDefForListTestCases()
	return &ListTestCasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTestReportsByCondition 根据查询条件获取测试报告列表
//
// 根据查询条件获取测试报告列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListTestReportsByCondition(request *model.ListTestReportsByConditionRequest) (*model.ListTestReportsByConditionResponse, error) {
	requestDef := GenReqDefForListTestReportsByCondition()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTestReportsByConditionResponse), nil
	}
}

// ListTestReportsByConditionInvoker 根据查询条件获取测试报告列表
func (c *CloudtestClient) ListTestReportsByConditionInvoker(request *model.ListTestReportsByConditionRequest) *ListTestReportsByConditionInvoker {
	requestDef := GenReqDefForListTestReportsByCondition()
	return &ListTestReportsByConditionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTestTypes 获取测试类型列表
//
// 获取测试类型列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListTestTypes(request *model.ListTestTypesRequest) (*model.ListTestTypesResponse, error) {
	requestDef := GenReqDefForListTestTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTestTypesResponse), nil
	}
}

// ListTestTypesInvoker 获取测试类型列表
func (c *CloudtestClient) ListTestTypesInvoker(request *model.ListTestTypesRequest) *ListTestTypesInvoker {
	requestDef := GenReqDefForListTestTypes()
	return &ListTestTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTestcasesByProjectIssuesRelation 查询项目下关联了需求的用例列表
//
// 查询项目下关联了需求的用例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListTestcasesByProjectIssuesRelation(request *model.ListTestcasesByProjectIssuesRelationRequest) (*model.ListTestcasesByProjectIssuesRelationResponse, error) {
	requestDef := GenReqDefForListTestcasesByProjectIssuesRelation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTestcasesByProjectIssuesRelationResponse), nil
	}
}

// ListTestcasesByProjectIssuesRelationInvoker 查询项目下关联了需求的用例列表
func (c *CloudtestClient) ListTestcasesByProjectIssuesRelationInvoker(request *model.ListTestcasesByProjectIssuesRelationRequest) *ListTestcasesByProjectIssuesRelationInvoker {
	requestDef := GenReqDefForListTestcasesByProjectIssuesRelation()
	return &ListTestcasesByProjectIssuesRelationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUsageInfos 获取租户订单已用资源信息
//
// 获取租户订单已用资源信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListUsageInfos(request *model.ListUsageInfosRequest) (*model.ListUsageInfosResponse, error) {
	requestDef := GenReqDefForListUsageInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsageInfosResponse), nil
	}
}

// ListUsageInfosInvoker 获取租户订单已用资源信息
func (c *CloudtestClient) ListUsageInfosInvoker(request *model.ListUsageInfosRequest) *ListUsageInfosInvoker {
	requestDef := GenReqDefForListUsageInfos()
	return &ListUsageInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserPackageUsage ListUserPackageUsage
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListUserPackageUsage(request *model.ListUserPackageUsageRequest) (*model.ListUserPackageUsageResponse, error) {
	requestDef := GenReqDefForListUserPackageUsage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserPackageUsageResponse), nil
	}
}

// ListUserPackageUsageInvoker ListUserPackageUsage
func (c *CloudtestClient) ListUserPackageUsageInvoker(request *model.ListUserPackageUsageRequest) *ListUserPackageUsageInvoker {
	requestDef := GenReqDefForListUserPackageUsage()
	return &ListUserPackageUsageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserPopupInfo ListUserPopupInfo
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListUserPopupInfo(request *model.ListUserPopupInfoRequest) (*model.ListUserPopupInfoResponse, error) {
	requestDef := GenReqDefForListUserPopupInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserPopupInfoResponse), nil
	}
}

// ListUserPopupInfoInvoker ListUserPopupInfo
func (c *CloudtestClient) ListUserPopupInfoInvoker(request *model.ListUserPopupInfoRequest) *ListUserPopupInfoInvoker {
	requestDef := GenReqDefForListUserPopupInfo()
	return &ListUserPopupInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunTestCase 批量执行测试用例
//
// 批量执行测试用例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) RunTestCase(request *model.RunTestCaseRequest) (*model.RunTestCaseResponse, error) {
	requestDef := GenReqDefForRunTestCase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunTestCaseResponse), nil
	}
}

// RunTestCaseInvoker 批量执行测试用例
func (c *CloudtestClient) RunTestCaseInvoker(request *model.RunTestCaseRequest) *RunTestCaseInvoker {
	requestDef := GenReqDefForRunTestCase()
	return &RunTestCaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAllFeatureChildren 获取特性树V5
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowAllFeatureChildren(request *model.ShowAllFeatureChildrenRequest) (*model.ShowAllFeatureChildrenResponse, error) {
	requestDef := GenReqDefForShowAllFeatureChildren()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAllFeatureChildrenResponse), nil
	}
}

// ShowAllFeatureChildrenInvoker 获取特性树V5
func (c *CloudtestClient) ShowAllFeatureChildrenInvoker(request *model.ShowAllFeatureChildrenRequest) *ShowAllFeatureChildrenInvoker {
	requestDef := GenReqDefForShowAllFeatureChildren()
	return &ShowAllFeatureChildrenInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApiTestcaseHistories 获取用例历史执行数据
//
// 获取用例历史执行数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowApiTestcaseHistories(request *model.ShowApiTestcaseHistoriesRequest) (*model.ShowApiTestcaseHistoriesResponse, error) {
	requestDef := GenReqDefForShowApiTestcaseHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApiTestcaseHistoriesResponse), nil
	}
}

// ShowApiTestcaseHistoriesInvoker 获取用例历史执行数据
func (c *CloudtestClient) ShowApiTestcaseHistoriesInvoker(request *model.ShowApiTestcaseHistoriesRequest) *ShowApiTestcaseHistoriesInvoker {
	requestDef := GenReqDefForShowApiTestcaseHistories()
	return &ShowApiTestcaseHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBackgroundInfo 获取测试报告的模板设置
//
// 获取测试报告的模板设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowBackgroundInfo(request *model.ShowBackgroundInfoRequest) (*model.ShowBackgroundInfoResponse, error) {
	requestDef := GenReqDefForShowBackgroundInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackgroundInfoResponse), nil
	}
}

// ShowBackgroundInfoInvoker 获取测试报告的模板设置
func (c *CloudtestClient) ShowBackgroundInfoInvoker(request *model.ShowBackgroundInfoRequest) *ShowBackgroundInfoInvoker {
	requestDef := GenReqDefForShowBackgroundInfo()
	return &ShowBackgroundInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDisclaimerRecord 查询用户免责声明
//
// 查询用户免责声明
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowDisclaimerRecord(request *model.ShowDisclaimerRecordRequest) (*model.ShowDisclaimerRecordResponse, error) {
	requestDef := GenReqDefForShowDisclaimerRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDisclaimerRecordResponse), nil
	}
}

// ShowDisclaimerRecordInvoker 查询用户免责声明
func (c *CloudtestClient) ShowDisclaimerRecordInvoker(request *model.ShowDisclaimerRecordRequest) *ShowDisclaimerRecordInvoker {
	requestDef := GenReqDefForShowDisclaimerRecord()
	return &ShowDisclaimerRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDomainInfo 根据domainId获取加密的testbirdkey
//
// 根据domainId获取加密的testbirdkey
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowDomainInfo(request *model.ShowDomainInfoRequest) (*model.ShowDomainInfoResponse, error) {
	requestDef := GenReqDefForShowDomainInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainInfoResponse), nil
	}
}

// ShowDomainInfoInvoker 根据domainId获取加密的testbirdkey
func (c *CloudtestClient) ShowDomainInfoInvoker(request *model.ShowDomainInfoRequest) *ShowDomainInfoInvoker {
	requestDef := GenReqDefForShowDomainInfo()
	return &ShowDomainInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFeatureChildren 获取目录\\特性树
//
// 获取目录\\特性树
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowFeatureChildren(request *model.ShowFeatureChildrenRequest) (*model.ShowFeatureChildrenResponse, error) {
	requestDef := GenReqDefForShowFeatureChildren()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFeatureChildrenResponse), nil
	}
}

// ShowFeatureChildrenInvoker 获取目录\\特性树
func (c *CloudtestClient) ShowFeatureChildrenInvoker(request *model.ShowFeatureChildrenRequest) *ShowFeatureChildrenInvoker {
	requestDef := GenReqDefForShowFeatureChildren()
	return &ShowFeatureChildrenInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFreeDeclaration 查询限时免费用户免责声明记录
//
// 查询限时免费用户免责声明记录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowFreeDeclaration(request *model.ShowFreeDeclarationRequest) (*model.ShowFreeDeclarationResponse, error) {
	requestDef := GenReqDefForShowFreeDeclaration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFreeDeclarationResponse), nil
	}
}

// ShowFreeDeclarationInvoker 查询限时免费用户免责声明记录
func (c *CloudtestClient) ShowFreeDeclarationInvoker(request *model.ShowFreeDeclarationRequest) *ShowFreeDeclarationInvoker {
	requestDef := GenReqDefForShowFreeDeclaration()
	return &ShowFreeDeclarationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIssuesByPlanId 查询某个测试计划下的需求树
//
// 查询某个测试计划下的需求列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowIssuesByPlanId(request *model.ShowIssuesByPlanIdRequest) (*model.ShowIssuesByPlanIdResponse, error) {
	requestDef := GenReqDefForShowIssuesByPlanId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIssuesByPlanIdResponse), nil
	}
}

// ShowIssuesByPlanIdInvoker 查询某个测试计划下的需求树
func (c *CloudtestClient) ShowIssuesByPlanIdInvoker(request *model.ShowIssuesByPlanIdRequest) *ShowIssuesByPlanIdInvoker {
	requestDef := GenReqDefForShowIssuesByPlanId()
	return &ShowIssuesByPlanIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIteratorByDefect 查询缺陷相关联测试计划
//
// 查询缺陷相关联测试计划
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowIteratorByDefect(request *model.ShowIteratorByDefectRequest) (*model.ShowIteratorByDefectResponse, error) {
	requestDef := GenReqDefForShowIteratorByDefect()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIteratorByDefectResponse), nil
	}
}

// ShowIteratorByDefectInvoker 查询缺陷相关联测试计划
func (c *CloudtestClient) ShowIteratorByDefectInvoker(request *model.ShowIteratorByDefectRequest) *ShowIteratorByDefectInvoker {
	requestDef := GenReqDefForShowIteratorByDefect()
	return &ShowIteratorByDefectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMindmapByPage 根据条件分页获取脑图对象V3
//
// 根据条件分页获取脑图对象V3
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowMindmapByPage(request *model.ShowMindmapByPageRequest) (*model.ShowMindmapByPageResponse, error) {
	requestDef := GenReqDefForShowMindmapByPage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMindmapByPageResponse), nil
	}
}

// ShowMindmapByPageInvoker 根据条件分页获取脑图对象V3
func (c *CloudtestClient) ShowMindmapByPageInvoker(request *model.ShowMindmapByPageRequest) *ShowMindmapByPageInvoker {
	requestDef := GenReqDefForShowMindmapByPage()
	return &ShowMindmapByPageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMindmapCreatorName 获取脑图创建人V2
//
// 获取脑图创建人V2
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowMindmapCreatorName(request *model.ShowMindmapCreatorNameRequest) (*model.ShowMindmapCreatorNameResponse, error) {
	requestDef := GenReqDefForShowMindmapCreatorName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMindmapCreatorNameResponse), nil
	}
}

// ShowMindmapCreatorNameInvoker 获取脑图创建人V2
func (c *CloudtestClient) ShowMindmapCreatorNameInvoker(request *model.ShowMindmapCreatorNameRequest) *ShowMindmapCreatorNameInvoker {
	requestDef := GenReqDefForShowMindmapCreatorName()
	return &ShowMindmapCreatorNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlanJournals 查询某测试计划下的操作历史
//
// 查询某测试计划下的操作历史
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowPlanJournals(request *model.ShowPlanJournalsRequest) (*model.ShowPlanJournalsResponse, error) {
	requestDef := GenReqDefForShowPlanJournals()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlanJournalsResponse), nil
	}
}

// ShowPlanJournalsInvoker 查询某测试计划下的操作历史
func (c *CloudtestClient) ShowPlanJournalsInvoker(request *model.ShowPlanJournalsRequest) *ShowPlanJournalsInvoker {
	requestDef := GenReqDefForShowPlanJournals()
	return &ShowPlanJournalsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlanList 项目下查询测试计划列表v2
//
// 项目下查询测试计划列表v2
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowPlanList(request *model.ShowPlanListRequest) (*model.ShowPlanListResponse, error) {
	requestDef := GenReqDefForShowPlanList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlanListResponse), nil
	}
}

// ShowPlanListInvoker 项目下查询测试计划列表v2
func (c *CloudtestClient) ShowPlanListInvoker(request *model.ShowPlanListRequest) *ShowPlanListInvoker {
	requestDef := GenReqDefForShowPlanList()
	return &ShowPlanListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlans 项目下查询测试计划列表
//
// 项目下查询测试计划列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowPlans(request *model.ShowPlansRequest) (*model.ShowPlansResponse, error) {
	requestDef := GenReqDefForShowPlans()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlansResponse), nil
	}
}

// ShowPlansInvoker 项目下查询测试计划列表
func (c *CloudtestClient) ShowPlansInvoker(request *model.ShowPlansRequest) *ShowPlansInvoker {
	requestDef := GenReqDefForShowPlans()
	return &ShowPlansInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProgress 获取异步进度
//
// 获取异步进度
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowProgress(request *model.ShowProgressRequest) (*model.ShowProgressResponse, error) {
	requestDef := GenReqDefForShowProgress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProgressResponse), nil
	}
}

// ShowProgressInvoker 获取异步进度
func (c *CloudtestClient) ShowProgressInvoker(request *model.ShowProgressRequest) *ShowProgressInvoker {
	requestDef := GenReqDefForShowProgress()
	return &ShowProgressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectDataDashboard 查询质量报告看板统计信息
//
// 查询质量报告看板统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowProjectDataDashboard(request *model.ShowProjectDataDashboardRequest) (*model.ShowProjectDataDashboardResponse, error) {
	requestDef := GenReqDefForShowProjectDataDashboard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectDataDashboardResponse), nil
	}
}

// ShowProjectDataDashboardInvoker 查询质量报告看板统计信息
func (c *CloudtestClient) ShowProjectDataDashboardInvoker(request *model.ShowProjectDataDashboardRequest) *ShowProjectDataDashboardInvoker {
	requestDef := GenReqDefForShowProjectDataDashboard()
	return &ShowProjectDataDashboardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRegisterService 用户获取自己当前已经注册的服务
//
// 用户获取自己当前已经注册的服务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowRegisterService(request *model.ShowRegisterServiceRequest) (*model.ShowRegisterServiceResponse, error) {
	requestDef := GenReqDefForShowRegisterService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRegisterServiceResponse), nil
	}
}

// ShowRegisterServiceInvoker 用户获取自己当前已经注册的服务
func (c *CloudtestClient) ShowRegisterServiceInvoker(request *model.ShowRegisterServiceRequest) *ShowRegisterServiceInvoker {
	requestDef := GenReqDefForShowRegisterService()
	return &ShowRegisterServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowReport 实时计算单个自定义报表
//
// 实时计算单个自定义报表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowReport(request *model.ShowReportRequest) (*model.ShowReportResponse, error) {
	requestDef := GenReqDefForShowReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReportResponse), nil
	}
}

// ShowReportInvoker 实时计算单个自定义报表
func (c *CloudtestClient) ShowReportInvoker(request *model.ShowReportRequest) *ShowReportInvoker {
	requestDef := GenReqDefForShowReport()
	return &ShowReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRequirementsOverview 质量报告需求测试情况统计
//
// 质量报告需求测试情况统计
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowRequirementsOverview(request *model.ShowRequirementsOverviewRequest) (*model.ShowRequirementsOverviewResponse, error) {
	requestDef := GenReqDefForShowRequirementsOverview()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRequirementsOverviewResponse), nil
	}
}

// ShowRequirementsOverviewInvoker 质量报告需求测试情况统计
func (c *CloudtestClient) ShowRequirementsOverviewInvoker(request *model.ShowRequirementsOverviewRequest) *ShowRequirementsOverviewInvoker {
	requestDef := GenReqDefForShowRequirementsOverview()
	return &ShowRequirementsOverviewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSystemConfigs 根据入参动态查询系统配置中的信息
//
// 根据入参动态查询系统配置中的信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowSystemConfigs(request *model.ShowSystemConfigsRequest) (*model.ShowSystemConfigsResponse, error) {
	requestDef := GenReqDefForShowSystemConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSystemConfigsResponse), nil
	}
}

// ShowSystemConfigsInvoker 根据入参动态查询系统配置中的信息
func (c *CloudtestClient) ShowSystemConfigsInvoker(request *model.ShowSystemConfigsRequest) *ShowSystemConfigsInvoker {
	requestDef := GenReqDefForShowSystemConfigs()
	return &ShowSystemConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTestCase 查询用例详情
//
// 查询用例详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowTestCase(request *model.ShowTestCaseRequest) (*model.ShowTestCaseResponse, error) {
	requestDef := GenReqDefForShowTestCase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTestCaseResponse), nil
	}
}

// ShowTestCaseInvoker 查询用例详情
func (c *CloudtestClient) ShowTestCaseInvoker(request *model.ShowTestCaseRequest) *ShowTestCaseInvoker {
	requestDef := GenReqDefForShowTestCase()
	return &ShowTestCaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTestCaseAndDefectInfo 查询用户用例关联缺陷的统计信息
//
// 查询用户用例关联缺陷的统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowTestCaseAndDefectInfo(request *model.ShowTestCaseAndDefectInfoRequest) (*model.ShowTestCaseAndDefectInfoResponse, error) {
	requestDef := GenReqDefForShowTestCaseAndDefectInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTestCaseAndDefectInfoResponse), nil
	}
}

// ShowTestCaseAndDefectInfoInvoker 查询用户用例关联缺陷的统计信息
func (c *CloudtestClient) ShowTestCaseAndDefectInfoInvoker(request *model.ShowTestCaseAndDefectInfoRequest) *ShowTestCaseAndDefectInfoInvoker {
	requestDef := GenReqDefForShowTestCaseAndDefectInfo()
	return &ShowTestCaseAndDefectInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTestCaseDetail 获取测试用例详情
//
// 获取测试用例详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowTestCaseDetail(request *model.ShowTestCaseDetailRequest) (*model.ShowTestCaseDetailResponse, error) {
	requestDef := GenReqDefForShowTestCaseDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTestCaseDetailResponse), nil
	}
}

// ShowTestCaseDetailInvoker 获取测试用例详情
func (c *CloudtestClient) ShowTestCaseDetailInvoker(request *model.ShowTestCaseDetailRequest) *ShowTestCaseDetailInvoker {
	requestDef := GenReqDefForShowTestCaseDetail()
	return &ShowTestCaseDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTestCaseDetailV2 通过用例编号获取测试用例详情
//
// 通过用例编号获取测试用例详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowTestCaseDetailV2(request *model.ShowTestCaseDetailV2Request) (*model.ShowTestCaseDetailV2Response, error) {
	requestDef := GenReqDefForShowTestCaseDetailV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTestCaseDetailV2Response), nil
	}
}

// ShowTestCaseDetailV2Invoker 通过用例编号获取测试用例详情
func (c *CloudtestClient) ShowTestCaseDetailV2Invoker(request *model.ShowTestCaseDetailV2Request) *ShowTestCaseDetailV2Invoker {
	requestDef := GenReqDefForShowTestCaseDetailV2()
	return &ShowTestCaseDetailV2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUserAccessInfo 获取租户订单信息
//
// 获取租户订单信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowUserAccessInfo(request *model.ShowUserAccessInfoRequest) (*model.ShowUserAccessInfoResponse, error) {
	requestDef := GenReqDefForShowUserAccessInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserAccessInfoResponse), nil
	}
}

// ShowUserAccessInfoInvoker 获取租户订单信息
func (c *CloudtestClient) ShowUserAccessInfoInvoker(request *model.ShowUserAccessInfoRequest) *ShowUserAccessInfoInvoker {
	requestDef := GenReqDefForShowUserAccessInfo()
	return &ShowUserAccessInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUserExecuteTestCaseInfo 查询时段内用例的执行情况
//
// 查询时段内用例的执行情况
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowUserExecuteTestCaseInfo(request *model.ShowUserExecuteTestCaseInfoRequest) (*model.ShowUserExecuteTestCaseInfoResponse, error) {
	requestDef := GenReqDefForShowUserExecuteTestCaseInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserExecuteTestCaseInfoResponse), nil
	}
}

// ShowUserExecuteTestCaseInfoInvoker 查询时段内用例的执行情况
func (c *CloudtestClient) ShowUserExecuteTestCaseInfoInvoker(request *model.ShowUserExecuteTestCaseInfoRequest) *ShowUserExecuteTestCaseInfoInvoker {
	requestDef := GenReqDefForShowUserExecuteTestCaseInfo()
	return &ShowUserExecuteTestCaseInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateService 更新已注册服务
//
// 更新已注册服务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) UpdateService(request *model.UpdateServiceRequest) (*model.UpdateServiceResponse, error) {
	requestDef := GenReqDefForUpdateService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateServiceResponse), nil
	}
}

// UpdateServiceInvoker 更新已注册服务
func (c *CloudtestClient) UpdateServiceInvoker(request *model.UpdateServiceRequest) *UpdateServiceInvoker {
	requestDef := GenReqDefForUpdateService()
	return &UpdateServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTestCase 更新自定义测试服务类型用例
//
// 更新自定义测试服务类型用例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) UpdateTestCase(request *model.UpdateTestCaseRequest) (*model.UpdateTestCaseResponse, error) {
	requestDef := GenReqDefForUpdateTestCase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTestCaseResponse), nil
	}
}

// UpdateTestCaseInvoker 更新自定义测试服务类型用例
func (c *CloudtestClient) UpdateTestCaseInvoker(request *model.UpdateTestCaseRequest) *UpdateTestCaseInvoker {
	requestDef := GenReqDefForUpdateTestCase()
	return &UpdateTestCaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTestCaseResult 批量更新测试用例结果
//
// 批量更新测试用例结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) UpdateTestCaseResult(request *model.UpdateTestCaseResultRequest) (*model.UpdateTestCaseResultResponse, error) {
	requestDef := GenReqDefForUpdateTestCaseResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTestCaseResultResponse), nil
	}
}

// UpdateTestCaseResultInvoker 批量更新测试用例结果
func (c *CloudtestClient) UpdateTestCaseResultInvoker(request *model.UpdateTestCaseResultRequest) *UpdateTestCaseResultInvoker {
	requestDef := GenReqDefForUpdateTestCaseResult()
	return &UpdateTestCaseResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlertGroupsByCondition 查询告警组列表
//
// 查询告警组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListAlertGroupsByCondition(request *model.ListAlertGroupsByConditionRequest) (*model.ListAlertGroupsByConditionResponse, error) {
	requestDef := GenReqDefForListAlertGroupsByCondition()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlertGroupsByConditionResponse), nil
	}
}

// ListAlertGroupsByConditionInvoker 查询告警组列表
func (c *CloudtestClient) ListAlertGroupsByConditionInvoker(request *model.ListAlertGroupsByConditionRequest) *ListAlertGroupsByConditionInvoker {
	requestDef := GenReqDefForListAlertGroupsByCondition()
	return &ListAlertGroupsByConditionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlertTemplates 查询告警模板
//
// 查询告警模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListAlertTemplates(request *model.ListAlertTemplatesRequest) (*model.ListAlertTemplatesResponse, error) {
	requestDef := GenReqDefForListAlertTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlertTemplatesResponse), nil
	}
}

// ListAlertTemplatesInvoker 查询告警模板
func (c *CloudtestClient) ListAlertTemplatesInvoker(request *model.ListAlertTemplatesRequest) *ListAlertTemplatesInvoker {
	requestDef := GenReqDefForListAlertTemplates()
	return &ListAlertTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllConfigItemByType 查询任务告警信息
//
// 查询任务告警信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListAllConfigItemByType(request *model.ListAllConfigItemByTypeRequest) (*model.ListAllConfigItemByTypeResponse, error) {
	requestDef := GenReqDefForListAllConfigItemByType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllConfigItemByTypeResponse), nil
	}
}

// ListAllConfigItemByTypeInvoker 查询任务告警信息
func (c *CloudtestClient) ListAllConfigItemByTypeInvoker(request *model.ListAllConfigItemByTypeRequest) *ListAllConfigItemByTypeInvoker {
	requestDef := GenReqDefForListAllConfigItemByType()
	return &ListAllConfigItemByTypeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SaveTaskSetting 保存任务配置
//
// 保存任务配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) SaveTaskSetting(request *model.SaveTaskSettingRequest) (*model.SaveTaskSettingResponse, error) {
	requestDef := GenReqDefForSaveTaskSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SaveTaskSettingResponse), nil
	}
}

// SaveTaskSettingInvoker 保存任务配置
func (c *CloudtestClient) SaveTaskSettingInvoker(request *model.SaveTaskSettingRequest) *SaveTaskSettingInvoker {
	requestDef := GenReqDefForSaveTaskSetting()
	return &SaveTaskSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAllConfigValueByTypeAndKey 查询任务配置
//
// 查询任务配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowAllConfigValueByTypeAndKey(request *model.ShowAllConfigValueByTypeAndKeyRequest) (*model.ShowAllConfigValueByTypeAndKeyResponse, error) {
	requestDef := GenReqDefForShowAllConfigValueByTypeAndKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAllConfigValueByTypeAndKeyResponse), nil
	}
}

// ShowAllConfigValueByTypeAndKeyInvoker 查询任务配置
func (c *CloudtestClient) ShowAllConfigValueByTypeAndKeyInvoker(request *model.ShowAllConfigValueByTypeAndKeyRequest) *ShowAllConfigValueByTypeAndKeyInvoker {
	requestDef := GenReqDefForShowAllConfigValueByTypeAndKey()
	return &ShowAllConfigValueByTypeAndKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIfTaskNameRepeat 查询告警模板名称是否重复
//
// 查询告警模板名称是否重复
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowIfTaskNameRepeat(request *model.ShowIfTaskNameRepeatRequest) (*model.ShowIfTaskNameRepeatResponse, error) {
	requestDef := GenReqDefForShowIfTaskNameRepeat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIfTaskNameRepeatResponse), nil
	}
}

// ShowIfTaskNameRepeatInvoker 查询告警模板名称是否重复
func (c *CloudtestClient) ShowIfTaskNameRepeatInvoker(request *model.ShowIfTaskNameRepeatRequest) *ShowIfTaskNameRepeatInvoker {
	requestDef := GenReqDefForShowIfTaskNameRepeat()
	return &ShowIfTaskNameRepeatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIfUserNameRepeat 查询告警组用户名是否重复
//
// 查询告警组用户名是否重复
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowIfUserNameRepeat(request *model.ShowIfUserNameRepeatRequest) (*model.ShowIfUserNameRepeatResponse, error) {
	requestDef := GenReqDefForShowIfUserNameRepeat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIfUserNameRepeatResponse), nil
	}
}

// ShowIfUserNameRepeatInvoker 查询告警组用户名是否重复
func (c *CloudtestClient) ShowIfUserNameRepeatInvoker(request *model.ShowIfUserNameRepeatRequest) *ShowIfUserNameRepeatInvoker {
	requestDef := GenReqDefForShowIfUserNameRepeat()
	return &ShowIfUserNameRepeatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateApiTestSuiteByRepoFile 通过导入仓库中的文件生成接口测试套
//
// 通过导入仓库中的文件生成接口测试套
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) CreateApiTestSuiteByRepoFile(request *model.CreateApiTestSuiteByRepoFileRequest) (*model.CreateApiTestSuiteByRepoFileResponse, error) {
	requestDef := GenReqDefForCreateApiTestSuiteByRepoFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApiTestSuiteByRepoFileResponse), nil
	}
}

// CreateApiTestSuiteByRepoFileInvoker 通过导入仓库中的文件生成接口测试套
func (c *CloudtestClient) CreateApiTestSuiteByRepoFileInvoker(request *model.CreateApiTestSuiteByRepoFileRequest) *CreateApiTestSuiteByRepoFileInvoker {
	requestDef := GenReqDefForCreateApiTestSuiteByRepoFile()
	return &CreateApiTestSuiteByRepoFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnvironments 获取环境参数分组列表
//
// 获取环境参数分组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListEnvironments(request *model.ListEnvironmentsRequest) (*model.ListEnvironmentsResponse, error) {
	requestDef := GenReqDefForListEnvironments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnvironmentsResponse), nil
	}
}

// ListEnvironmentsInvoker 获取环境参数分组列表
func (c *CloudtestClient) ListEnvironmentsInvoker(request *model.ListEnvironmentsRequest) *ListEnvironmentsInvoker {
	requestDef := GenReqDefForListEnvironments()
	return &ListEnvironmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBasicAw 根据id获取单个basicAW信息
//
// 根据id获取单个basicAW信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListBasicAw(request *model.ListBasicAwRequest) (*model.ListBasicAwResponse, error) {
	requestDef := GenReqDefForListBasicAw()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBasicAwResponse), nil
	}
}

// ListBasicAwInvoker 根据id获取单个basicAW信息
func (c *CloudtestClient) ListBasicAwInvoker(request *model.ListBasicAwRequest) *ListBasicAwInvoker {
	requestDef := GenReqDefForListBasicAw()
	return &ListBasicAwInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPublicLibAndAws 获取工程关联的公共aw信息和公共aw所属公共aw库信息
//
// 获取工程关联的公共aw信息和公共aw所属公共aw库信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListPublicLibAndAws(request *model.ListPublicLibAndAwsRequest) (*model.ListPublicLibAndAwsResponse, error) {
	requestDef := GenReqDefForListPublicLibAndAws()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPublicLibAndAwsResponse), nil
	}
}

// ListPublicLibAndAwsInvoker 获取工程关联的公共aw信息和公共aw所属公共aw库信息
func (c *CloudtestClient) ListPublicLibAndAwsInvoker(request *model.ListPublicLibAndAwsRequest) *ListPublicLibAndAwsInvoker {
	requestDef := GenReqDefForListPublicLibAndAws()
	return &ListPublicLibAndAwsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserDnsMapping 查询用户DNS映射
//
// 查询用户DNS映射
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListUserDnsMapping(request *model.ListUserDnsMappingRequest) (*model.ListUserDnsMappingResponse, error) {
	requestDef := GenReqDefForListUserDnsMapping()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserDnsMappingResponse), nil
	}
}

// ListUserDnsMappingInvoker 查询用户DNS映射
func (c *CloudtestClient) ListUserDnsMappingInvoker(request *model.ListUserDnsMappingRequest) *ListUserDnsMappingInvoker {
	requestDef := GenReqDefForListUserDnsMapping()
	return &ListUserDnsMappingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVariables 查询全局变量参数列表V4
//
// 查询全局变量参数列表V4
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListVariables(request *model.ListVariablesRequest) (*model.ListVariablesResponse, error) {
	requestDef := GenReqDefForListVariables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVariablesResponse), nil
	}
}

// ListVariablesInvoker 查询全局变量参数列表V4
func (c *CloudtestClient) ListVariablesInvoker(request *model.ListVariablesRequest) *ListVariablesInvoker {
	requestDef := GenReqDefForListVariables()
	return &ListVariablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
