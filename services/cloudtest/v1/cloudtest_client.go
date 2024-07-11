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

// AddTestCaseComment 新增用例评论
//
// 新增用例评论
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) AddTestCaseComment(request *model.AddTestCaseCommentRequest) (*model.AddTestCaseCommentResponse, error) {
	requestDef := GenReqDefForAddTestCaseComment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddTestCaseCommentResponse), nil
	}
}

// AddTestCaseCommentInvoker 新增用例评论
func (c *CloudtestClient) AddTestCaseCommentInvoker(request *model.AddTestCaseCommentRequest) *AddTestCaseCommentInvoker {
	requestDef := GenReqDefForAddTestCaseComment()
	return &AddTestCaseCommentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// BatchAddResourcesForIterator 向迭代中添加资源
//
// 向迭代中添加资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) BatchAddResourcesForIterator(request *model.BatchAddResourcesForIteratorRequest) (*model.BatchAddResourcesForIteratorResponse, error) {
	requestDef := GenReqDefForBatchAddResourcesForIterator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddResourcesForIteratorResponse), nil
	}
}

// BatchAddResourcesForIteratorInvoker 向迭代中添加资源
func (c *CloudtestClient) BatchAddResourcesForIteratorInvoker(request *model.BatchAddResourcesForIteratorRequest) *BatchAddResourcesForIteratorInvoker {
	requestDef := GenReqDefForBatchAddResourcesForIterator()
	return &BatchAddResourcesForIteratorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// BatchDeleteTestCases 批量删除用例
//
// 批量删除用例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) BatchDeleteTestCases(request *model.BatchDeleteTestCasesRequest) (*model.BatchDeleteTestCasesResponse, error) {
	requestDef := GenReqDefForBatchDeleteTestCases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteTestCasesResponse), nil
	}
}

// BatchDeleteTestCasesInvoker 批量删除用例
func (c *CloudtestClient) BatchDeleteTestCasesInvoker(request *model.BatchDeleteTestCasesRequest) *BatchDeleteTestCasesInvoker {
	requestDef := GenReqDefForBatchDeleteTestCases()
	return &BatchDeleteTestCasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// BatchRemoveTestCasesFromIterator 从迭代中批量移除用例
//
// 从迭代中批量移除用例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) BatchRemoveTestCasesFromIterator(request *model.BatchRemoveTestCasesFromIteratorRequest) (*model.BatchRemoveTestCasesFromIteratorResponse, error) {
	requestDef := GenReqDefForBatchRemoveTestCasesFromIterator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRemoveTestCasesFromIteratorResponse), nil
	}
}

// BatchRemoveTestCasesFromIteratorInvoker 从迭代中批量移除用例
func (c *CloudtestClient) BatchRemoveTestCasesFromIteratorInvoker(request *model.BatchRemoveTestCasesFromIteratorRequest) *BatchRemoveTestCasesFromIteratorInvoker {
	requestDef := GenReqDefForBatchRemoveTestCasesFromIterator()
	return &BatchRemoveTestCasesFromIteratorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateVersionTestCases 批量更新用例属性
//
// 批量更新用例属性
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) BatchUpdateVersionTestCases(request *model.BatchUpdateVersionTestCasesRequest) (*model.BatchUpdateVersionTestCasesResponse, error) {
	requestDef := GenReqDefForBatchUpdateVersionTestCases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateVersionTestCasesResponse), nil
	}
}

// BatchUpdateVersionTestCasesInvoker 批量更新用例属性
func (c *CloudtestClient) BatchUpdateVersionTestCasesInvoker(request *model.BatchUpdateVersionTestCasesRequest) *BatchUpdateVersionTestCasesInvoker {
	requestDef := GenReqDefForBatchUpdateVersionTestCases()
	return &BatchUpdateVersionTestCasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CreateIterator 新增迭代
//
// 新增迭代
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) CreateIterator(request *model.CreateIteratorRequest) (*model.CreateIteratorResponse, error) {
	requestDef := GenReqDefForCreateIterator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIteratorResponse), nil
	}
}

// CreateIteratorInvoker 新增迭代
func (c *CloudtestClient) CreateIteratorInvoker(request *model.CreateIteratorRequest) *CreateIteratorInvoker {
	requestDef := GenReqDefForCreateIterator()
	return &CreateIteratorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CreateProjectBranch 新增分支
//
// 新增分支
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) CreateProjectBranch(request *model.CreateProjectBranchRequest) (*model.CreateProjectBranchResponse, error) {
	requestDef := GenReqDefForCreateProjectBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProjectBranchResponse), nil
	}
}

// CreateProjectBranchInvoker 新增分支
func (c *CloudtestClient) CreateProjectBranchInvoker(request *model.CreateProjectBranchRequest) *CreateProjectBranchInvoker {
	requestDef := GenReqDefForCreateProjectBranch()
	return &CreateProjectBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CreateReport 保存单个自定义报表
//
// 保存单个自定义报表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) CreateReport(request *model.CreateReportRequest) (*model.CreateReportResponse, error) {
	requestDef := GenReqDefForCreateReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateReportResponse), nil
	}
}

// CreateReportInvoker 保存单个自定义报表
func (c *CloudtestClient) CreateReportInvoker(request *model.CreateReportRequest) *CreateReportInvoker {
	requestDef := GenReqDefForCreateReport()
	return &CreateReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CreateUserDefinedUrlKeyWord 新增用户自定义URL关键字
//
// 新增用户自定义URL关键字
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) CreateUserDefinedUrlKeyWord(request *model.CreateUserDefinedUrlKeyWordRequest) (*model.CreateUserDefinedUrlKeyWordResponse, error) {
	requestDef := GenReqDefForCreateUserDefinedUrlKeyWord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateUserDefinedUrlKeyWordResponse), nil
	}
}

// CreateUserDefinedUrlKeyWordInvoker 新增用户自定义URL关键字
func (c *CloudtestClient) CreateUserDefinedUrlKeyWordInvoker(request *model.CreateUserDefinedUrlKeyWordRequest) *CreateUserDefinedUrlKeyWordInvoker {
	requestDef := GenReqDefForCreateUserDefinedUrlKeyWord()
	return &CreateUserDefinedUrlKeyWordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVersionTestCase 在分支或者测试计划下创建用例
//
// 在分支或者测试计划下创建用例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) CreateVersionTestCase(request *model.CreateVersionTestCaseRequest) (*model.CreateVersionTestCaseResponse, error) {
	requestDef := GenReqDefForCreateVersionTestCase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVersionTestCaseResponse), nil
	}
}

// CreateVersionTestCaseInvoker 在分支或者测试计划下创建用例
func (c *CloudtestClient) CreateVersionTestCaseInvoker(request *model.CreateVersionTestCaseRequest) *CreateVersionTestCaseInvoker {
	requestDef := GenReqDefForCreateVersionTestCase()
	return &CreateVersionTestCaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBasicAwById 融合版本删除单个basicAw
//
// 融合版本删除单个basicAw
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) DeleteBasicAwById(request *model.DeleteBasicAwByIdRequest) (*model.DeleteBasicAwByIdResponse, error) {
	requestDef := GenReqDefForDeleteBasicAwById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBasicAwByIdResponse), nil
	}
}

// DeleteBasicAwByIdInvoker 融合版本删除单个basicAw
func (c *CloudtestClient) DeleteBasicAwByIdInvoker(request *model.DeleteBasicAwByIdRequest) *DeleteBasicAwByIdInvoker {
	requestDef := GenReqDefForDeleteBasicAwById()
	return &DeleteBasicAwByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// DeleteTestCaseComment 删除用例评论
//
// 删除用例评论
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) DeleteTestCaseComment(request *model.DeleteTestCaseCommentRequest) (*model.DeleteTestCaseCommentResponse, error) {
	requestDef := GenReqDefForDeleteTestCaseComment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTestCaseCommentResponse), nil
	}
}

// DeleteTestCaseCommentInvoker 删除用例评论
func (c *CloudtestClient) DeleteTestCaseCommentInvoker(request *model.DeleteTestCaseCommentRequest) *DeleteTestCaseCommentInvoker {
	requestDef := GenReqDefForDeleteTestCaseComment()
	return &DeleteTestCaseCommentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarmStatisticsUsing 查询告警统计数据
//
// 查询告警统计数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListAlarmStatisticsUsing(request *model.ListAlarmStatisticsUsingRequest) (*model.ListAlarmStatisticsUsingResponse, error) {
	requestDef := GenReqDefForListAlarmStatisticsUsing()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmStatisticsUsingResponse), nil
	}
}

// ListAlarmStatisticsUsingInvoker 查询告警统计数据
func (c *CloudtestClient) ListAlarmStatisticsUsingInvoker(request *model.ListAlarmStatisticsUsingRequest) *ListAlarmStatisticsUsingInvoker {
	requestDef := GenReqDefForListAlarmStatisticsUsing()
	return &ListAlarmStatisticsUsingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListAllTestCases 查询用例列表
//
// 查询用例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListAllTestCases(request *model.ListAllTestCasesRequest) (*model.ListAllTestCasesResponse, error) {
	requestDef := GenReqDefForListAllTestCases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllTestCasesResponse), nil
	}
}

// ListAllTestCasesInvoker 查询用例列表
func (c *CloudtestClient) ListAllTestCasesInvoker(request *model.ListAllTestCasesRequest) *ListAllTestCasesInvoker {
	requestDef := GenReqDefForListAllTestCases()
	return &ListAllTestCasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListAvailableConfig 获取当前局点功能是否开启
//
// 获取当前局点功能是否开启
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListAvailableConfig(request *model.ListAvailableConfigRequest) (*model.ListAvailableConfigResponse, error) {
	requestDef := GenReqDefForListAvailableConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailableConfigResponse), nil
	}
}

// ListAvailableConfigInvoker 获取当前局点功能是否开启
func (c *CloudtestClient) ListAvailableConfigInvoker(request *model.ListAvailableConfigRequest) *ListAvailableConfigInvoker {
	requestDef := GenReqDefForListAvailableConfig()
	return &ListAvailableConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListBasicAwInfoListSupportsSearch 分页获取工程BasicAW详细信息列表（含目录）
//
// 分页获取工程BasicAW列表（含目录）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListBasicAwInfoListSupportsSearch(request *model.ListBasicAwInfoListSupportsSearchRequest) (*model.ListBasicAwInfoListSupportsSearchResponse, error) {
	requestDef := GenReqDefForListBasicAwInfoListSupportsSearch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBasicAwInfoListSupportsSearchResponse), nil
	}
}

// ListBasicAwInfoListSupportsSearchInvoker 分页获取工程BasicAW详细信息列表（含目录）
func (c *CloudtestClient) ListBasicAwInfoListSupportsSearchInvoker(request *model.ListBasicAwInfoListSupportsSearchRequest) *ListBasicAwInfoListSupportsSearchInvoker {
	requestDef := GenReqDefForListBasicAwInfoListSupportsSearch()
	return &ListBasicAwInfoListSupportsSearchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListCasesStatus 批量获取用例状态
//
// 批量获取用例状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListCasesStatus(request *model.ListCasesStatusRequest) (*model.ListCasesStatusResponse, error) {
	requestDef := GenReqDefForListCasesStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCasesStatusResponse), nil
	}
}

// ListCasesStatusInvoker 批量获取用例状态
func (c *CloudtestClient) ListCasesStatusInvoker(request *model.ListCasesStatusRequest) *ListCasesStatusInvoker {
	requestDef := GenReqDefForListCasesStatus()
	return &ListCasesStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDomainVisibleServices 查询当前租户可见的第三方服务列表
//
// 查询当前租户可见的第三方服务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListDomainVisibleServices(request *model.ListDomainVisibleServicesRequest) (*model.ListDomainVisibleServicesResponse, error) {
	requestDef := GenReqDefForListDomainVisibleServices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainVisibleServicesResponse), nil
	}
}

// ListDomainVisibleServicesInvoker 查询当前租户可见的第三方服务列表
func (c *CloudtestClient) ListDomainVisibleServicesInvoker(request *model.ListDomainVisibleServicesRequest) *ListDomainVisibleServicesInvoker {
	requestDef := GenReqDefForListDomainVisibleServices()
	return &ListDomainVisibleServicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListIteratorIssueTree 查询迭代关联的需求列表或树
//
// 查询迭代关联的需求列表或树
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListIteratorIssueTree(request *model.ListIteratorIssueTreeRequest) (*model.ListIteratorIssueTreeResponse, error) {
	requestDef := GenReqDefForListIteratorIssueTree()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIteratorIssueTreeResponse), nil
	}
}

// ListIteratorIssueTreeInvoker 查询迭代关联的需求列表或树
func (c *CloudtestClient) ListIteratorIssueTreeInvoker(request *model.ListIteratorIssueTreeRequest) *ListIteratorIssueTreeInvoker {
	requestDef := GenReqDefForListIteratorIssueTree()
	return &ListIteratorIssueTreeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIterators 查询迭代计划列表，包含统计信息
//
// 查询迭代计划列表，包含统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListIterators(request *model.ListIteratorsRequest) (*model.ListIteratorsResponse, error) {
	requestDef := GenReqDefForListIterators()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIteratorsResponse), nil
	}
}

// ListIteratorsInvoker 查询迭代计划列表，包含统计信息
func (c *CloudtestClient) ListIteratorsInvoker(request *model.ListIteratorsRequest) *ListIteratorsInvoker {
	requestDef := GenReqDefForListIterators()
	return &ListIteratorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLinesUsing 查询仪表盘折线图数据
//
// 查询仪表盘折线图数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListLinesUsing(request *model.ListLinesUsingRequest) (*model.ListLinesUsingResponse, error) {
	requestDef := GenReqDefForListLinesUsing()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLinesUsingResponse), nil
	}
}

// ListLinesUsingInvoker 查询仪表盘折线图数据
func (c *CloudtestClient) ListLinesUsingInvoker(request *model.ListLinesUsingRequest) *ListLinesUsingInvoker {
	requestDef := GenReqDefForListLinesUsing()
	return &ListLinesUsingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMsgInfosUsing 仪表盘根据测试服务ID获取MsgInfo详细信息
//
// 成功返回MsgInfo的详细信息列表，返回值为Model的List
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListMsgInfosUsing(request *model.ListMsgInfosUsingRequest) (*model.ListMsgInfosUsingResponse, error) {
	requestDef := GenReqDefForListMsgInfosUsing()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMsgInfosUsingResponse), nil
	}
}

// ListMsgInfosUsingInvoker 仪表盘根据测试服务ID获取MsgInfo详细信息
func (c *CloudtestClient) ListMsgInfosUsingInvoker(request *model.ListMsgInfosUsingRequest) *ListMsgInfosUsingInvoker {
	requestDef := GenReqDefForListMsgInfosUsing()
	return &ListMsgInfosUsingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListOwnTestCases 获取责任人是自己的测试用例
//
// 获取责任人是自己的测试用例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListOwnTestCases(request *model.ListOwnTestCasesRequest) (*model.ListOwnTestCasesResponse, error) {
	requestDef := GenReqDefForListOwnTestCases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOwnTestCasesResponse), nil
	}
}

// ListOwnTestCasesInvoker 获取责任人是自己的测试用例
func (c *CloudtestClient) ListOwnTestCasesInvoker(request *model.ListOwnTestCasesRequest) *ListOwnTestCasesInvoker {
	requestDef := GenReqDefForListOwnTestCases()
	return &ListOwnTestCasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListScattersUsing 查询仪表盘散点图数据
//
// 查询仪表盘散点图数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListScattersUsing(request *model.ListScattersUsingRequest) (*model.ListScattersUsingResponse, error) {
	requestDef := GenReqDefForListScattersUsing()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScattersUsingResponse), nil
	}
}

// ListScattersUsingInvoker 查询仪表盘散点图数据
func (c *CloudtestClient) ListScattersUsingInvoker(request *model.ListScattersUsingRequest) *ListScattersUsingInvoker {
	requestDef := GenReqDefForListScattersUsing()
	return &ListScattersUsingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubTaskCaseOverstockUsing 查询subTestCase阻塞任务数据
//
// 成功返回子任务用例数据积压信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListSubTaskCaseOverstockUsing(request *model.ListSubTaskCaseOverstockUsingRequest) (*model.ListSubTaskCaseOverstockUsingResponse, error) {
	requestDef := GenReqDefForListSubTaskCaseOverstockUsing()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubTaskCaseOverstockUsingResponse), nil
	}
}

// ListSubTaskCaseOverstockUsingInvoker 查询subTestCase阻塞任务数据
func (c *CloudtestClient) ListSubTaskCaseOverstockUsingInvoker(request *model.ListSubTaskCaseOverstockUsingRequest) *ListSubTaskCaseOverstockUsingInvoker {
	requestDef := GenReqDefForListSubTaskCaseOverstockUsing()
	return &ListSubTaskCaseOverstockUsingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTaskAssignCases 获取测试套关联用例详情
//
// 获取测试套关联用例详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListTaskAssignCases(request *model.ListTaskAssignCasesRequest) (*model.ListTaskAssignCasesResponse, error) {
	requestDef := GenReqDefForListTaskAssignCases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTaskAssignCasesResponse), nil
	}
}

// ListTaskAssignCasesInvoker 获取测试套关联用例详情
func (c *CloudtestClient) ListTaskAssignCasesInvoker(request *model.ListTaskAssignCasesRequest) *ListTaskAssignCasesInvoker {
	requestDef := GenReqDefForListTaskAssignCases()
	return &ListTaskAssignCasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListTasks 查询测试任务列表
//
// 查询测试任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListTasks(request *model.ListTasksRequest) (*model.ListTasksResponse, error) {
	requestDef := GenReqDefForListTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTasksResponse), nil
	}
}

// ListTasksInvoker 查询测试任务列表
func (c *CloudtestClient) ListTasksInvoker(request *model.ListTasksRequest) *ListTasksInvoker {
	requestDef := GenReqDefForListTasks()
	return &ListTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTestCaseComments 查询用例评论
//
// 查询用例评论
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListTestCaseComments(request *model.ListTestCaseCommentsRequest) (*model.ListTestCaseCommentsResponse, error) {
	requestDef := GenReqDefForListTestCaseComments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTestCaseCommentsResponse), nil
	}
}

// ListTestCaseCommentsInvoker 查询用例评论
func (c *CloudtestClient) ListTestCaseCommentsInvoker(request *model.ListTestCaseCommentsRequest) *ListTestCaseCommentsInvoker {
	requestDef := GenReqDefForListTestCaseComments()
	return &ListTestCaseCommentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListTestCaseScriptDetail 获取用例脚本详细信息
//
// 获取用例脚本详细信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListTestCaseScriptDetail(request *model.ListTestCaseScriptDetailRequest) (*model.ListTestCaseScriptDetailResponse, error) {
	requestDef := GenReqDefForListTestCaseScriptDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTestCaseScriptDetailResponse), nil
	}
}

// ListTestCaseScriptDetailInvoker 获取用例脚本详细信息
func (c *CloudtestClient) ListTestCaseScriptDetailInvoker(request *model.ListTestCaseScriptDetailRequest) *ListTestCaseScriptDetailInvoker {
	requestDef := GenReqDefForListTestCaseScriptDetail()
	return &ListTestCaseScriptDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListTestCasesByIssue 查询需求下的用例列表
//
// 查询需求下的用例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListTestCasesByIssue(request *model.ListTestCasesByIssueRequest) (*model.ListTestCasesByIssueResponse, error) {
	requestDef := GenReqDefForListTestCasesByIssue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTestCasesByIssueResponse), nil
	}
}

// ListTestCasesByIssueInvoker 查询需求下的用例列表
func (c *CloudtestClient) ListTestCasesByIssueInvoker(request *model.ListTestCasesByIssueRequest) *ListTestCasesByIssueInvoker {
	requestDef := GenReqDefForListTestCasesByIssue()
	return &ListTestCasesByIssueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListUsingGet 查询仪表盘用户的看板
//
// 查询仪表盘用户的看板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ListUsingGet(request *model.ListUsingGetRequest) (*model.ListUsingGetResponse, error) {
	requestDef := GenReqDefForListUsingGet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsingGetResponse), nil
	}
}

// ListUsingGetInvoker 查询仪表盘用户的看板
func (c *CloudtestClient) ListUsingGetInvoker(request *model.ListUsingGetRequest) *ListUsingGetInvoker {
	requestDef := GenReqDefForListUsingGet()
	return &ListUsingGetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// RemoveIssuesFromIterator 从迭代中移除需求
//
// 从迭代中移除需求
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) RemoveIssuesFromIterator(request *model.RemoveIssuesFromIteratorRequest) (*model.RemoveIssuesFromIteratorResponse, error) {
	requestDef := GenReqDefForRemoveIssuesFromIterator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveIssuesFromIteratorResponse), nil
	}
}

// RemoveIssuesFromIteratorInvoker 从迭代中移除需求
func (c *CloudtestClient) RemoveIssuesFromIteratorInvoker(request *model.RemoveIssuesFromIteratorRequest) *RemoveIssuesFromIteratorInvoker {
	requestDef := GenReqDefForRemoveIssuesFromIterator()
	return &RemoveIssuesFromIteratorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowAllFeatureChildren 获取特性树V5
//
// 获取目录\\特性树
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

// ShowAsset 获取资产列表
//
// 获取资产列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowAsset(request *model.ShowAssetRequest) (*model.ShowAssetResponse, error) {
	requestDef := GenReqDefForShowAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAssetResponse), nil
	}
}

// ShowAssetInvoker 获取资产列表
func (c *CloudtestClient) ShowAssetInvoker(request *model.ShowAssetRequest) *ShowAssetInvoker {
	requestDef := GenReqDefForShowAsset()
	return &ShowAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAssetTree 获取资产树列表
//
// 获取资产树列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowAssetTree(request *model.ShowAssetTreeRequest) (*model.ShowAssetTreeResponse, error) {
	requestDef := GenReqDefForShowAssetTree()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAssetTreeResponse), nil
	}
}

// ShowAssetTreeInvoker 获取资产树列表
func (c *CloudtestClient) ShowAssetTreeInvoker(request *model.ShowAssetTreeRequest) *ShowAssetTreeInvoker {
	requestDef := GenReqDefForShowAssetTree()
	return &ShowAssetTreeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowBranch 获取分支详情
//
// 获取分支详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowBranch(request *model.ShowBranchRequest) (*model.ShowBranchResponse, error) {
	requestDef := GenReqDefForShowBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBranchResponse), nil
	}
}

// ShowBranchInvoker 获取分支详情
func (c *CloudtestClient) ShowBranchInvoker(request *model.ShowBranchRequest) *ShowBranchInvoker {
	requestDef := GenReqDefForShowBranch()
	return &ShowBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConcurrencyPackageUsing 查询租户测试并发套餐状态
//
// 查询租户测试并发套餐状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowConcurrencyPackageUsing(request *model.ShowConcurrencyPackageUsingRequest) (*model.ShowConcurrencyPackageUsingResponse, error) {
	requestDef := GenReqDefForShowConcurrencyPackageUsing()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConcurrencyPackageUsingResponse), nil
	}
}

// ShowConcurrencyPackageUsingInvoker 查询租户测试并发套餐状态
func (c *CloudtestClient) ShowConcurrencyPackageUsingInvoker(request *model.ShowConcurrencyPackageUsingRequest) *ShowConcurrencyPackageUsingInvoker {
	requestDef := GenReqDefForShowConcurrencyPackageUsing()
	return &ShowConcurrencyPackageUsingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowEchoTestPackageUsing 查询租户在线拨测套餐状态
//
// 查询租户在线拨测套餐状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowEchoTestPackageUsing(request *model.ShowEchoTestPackageUsingRequest) (*model.ShowEchoTestPackageUsingResponse, error) {
	requestDef := GenReqDefForShowEchoTestPackageUsing()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEchoTestPackageUsingResponse), nil
	}
}

// ShowEchoTestPackageUsingInvoker 查询租户在线拨测套餐状态
func (c *CloudtestClient) ShowEchoTestPackageUsingInvoker(request *model.ShowEchoTestPackageUsingRequest) *ShowEchoTestPackageUsingInvoker {
	requestDef := GenReqDefForShowEchoTestPackageUsing()
	return &ShowEchoTestPackageUsingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFactorByAssetId 根据目录查询因子
//
// 根据目录查询因子
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowFactorByAssetId(request *model.ShowFactorByAssetIdRequest) (*model.ShowFactorByAssetIdResponse, error) {
	requestDef := GenReqDefForShowFactorByAssetId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFactorByAssetIdResponse), nil
	}
}

// ShowFactorByAssetIdInvoker 根据目录查询因子
func (c *CloudtestClient) ShowFactorByAssetIdInvoker(request *model.ShowFactorByAssetIdRequest) *ShowFactorByAssetIdInvoker {
	requestDef := GenReqDefForShowFactorByAssetId()
	return &ShowFactorByAssetIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFactorById 根据id获取因子
//
// 根据id获取因子
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowFactorById(request *model.ShowFactorByIdRequest) (*model.ShowFactorByIdResponse, error) {
	requestDef := GenReqDefForShowFactorById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFactorByIdResponse), nil
	}
}

// ShowFactorByIdInvoker 根据id获取因子
func (c *CloudtestClient) ShowFactorByIdInvoker(request *model.ShowFactorByIdRequest) *ShowFactorByIdInvoker {
	requestDef := GenReqDefForShowFactorById()
	return &ShowFactorByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowIteratorDetail 查询迭代计划详情，包含统计信息
//
// 查询迭代计划详情，包含统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowIteratorDetail(request *model.ShowIteratorDetailRequest) (*model.ShowIteratorDetailResponse, error) {
	requestDef := GenReqDefForShowIteratorDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIteratorDetailResponse), nil
	}
}

// ShowIteratorDetailInvoker 查询迭代计划详情，包含统计信息
func (c *CloudtestClient) ShowIteratorDetailInvoker(request *model.ShowIteratorDetailRequest) *ShowIteratorDetailInvoker {
	requestDef := GenReqDefForShowIteratorDetail()
	return &ShowIteratorDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMindMapById 根据id获取脑图对象
//
// 根据id获取脑图对象
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowMindMapById(request *model.ShowMindMapByIdRequest) (*model.ShowMindMapByIdResponse, error) {
	requestDef := GenReqDefForShowMindMapById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMindMapByIdResponse), nil
	}
}

// ShowMindMapByIdInvoker 根据id获取脑图对象
func (c *CloudtestClient) ShowMindMapByIdInvoker(request *model.ShowMindMapByIdRequest) *ShowMindMapByIdInvoker {
	requestDef := GenReqDefForShowMindMapById()
	return &ShowMindMapByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowOperationalDataCurrentMonthUsing 查询运行面板信息
//
// 成功返回运行面板信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowOperationalDataCurrentMonthUsing(request *model.ShowOperationalDataCurrentMonthUsingRequest) (*model.ShowOperationalDataCurrentMonthUsingResponse, error) {
	requestDef := GenReqDefForShowOperationalDataCurrentMonthUsing()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOperationalDataCurrentMonthUsingResponse), nil
	}
}

// ShowOperationalDataCurrentMonthUsingInvoker 查询运行面板信息
func (c *CloudtestClient) ShowOperationalDataCurrentMonthUsingInvoker(request *model.ShowOperationalDataCurrentMonthUsingRequest) *ShowOperationalDataCurrentMonthUsingInvoker {
	requestDef := GenReqDefForShowOperationalDataCurrentMonthUsing()
	return &ShowOperationalDataCurrentMonthUsingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowReviewByPage 根据条件分页获取评审对象V2
//
// 根据条件分页获取评审对象V2
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowReviewByPage(request *model.ShowReviewByPageRequest) (*model.ShowReviewByPageResponse, error) {
	requestDef := GenReqDefForShowReviewByPage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReviewByPageResponse), nil
	}
}

// ShowReviewByPageInvoker 根据条件分页获取评审对象V2
func (c *CloudtestClient) ShowReviewByPageInvoker(request *model.ShowReviewByPageRequest) *ShowReviewByPageInvoker {
	requestDef := GenReqDefForShowReviewByPage()
	return &ShowReviewByPageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSceneByPage 根据条件分页获取场景对象V2
//
// 根据条件分页获取场景对象V2
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowSceneByPage(request *model.ShowSceneByPageRequest) (*model.ShowSceneByPageResponse, error) {
	requestDef := GenReqDefForShowSceneByPage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSceneByPageResponse), nil
	}
}

// ShowSceneByPageInvoker 根据条件分页获取场景对象V2
func (c *CloudtestClient) ShowSceneByPageInvoker(request *model.ShowSceneByPageRequest) *ShowSceneByPageInvoker {
	requestDef := GenReqDefForShowSceneByPage()
	return &ShowSceneByPageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStatisticById 根据脑图id查询统计数目
//
// 根据脑图id查询统计数目
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowStatisticById(request *model.ShowStatisticByIdRequest) (*model.ShowStatisticByIdResponse, error) {
	requestDef := GenReqDefForShowStatisticById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStatisticByIdResponse), nil
	}
}

// ShowStatisticByIdInvoker 根据脑图id查询统计数目
func (c *CloudtestClient) ShowStatisticByIdInvoker(request *model.ShowStatisticByIdRequest) *ShowStatisticByIdInvoker {
	requestDef := GenReqDefForShowStatisticById()
	return &ShowStatisticByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowTemplateByPage 根据条件分页获取模板V3
//
// 根据条件分页获取模板V3
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowTemplateByPage(request *model.ShowTemplateByPageRequest) (*model.ShowTemplateByPageResponse, error) {
	requestDef := GenReqDefForShowTemplateByPage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTemplateByPageResponse), nil
	}
}

// ShowTemplateByPageInvoker 根据条件分页获取模板V3
func (c *CloudtestClient) ShowTemplateByPageInvoker(request *model.ShowTemplateByPageRequest) *ShowTemplateByPageInvoker {
	requestDef := GenReqDefForShowTemplateByPage()
	return &ShowTemplateByPageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowTestCaseReviews 根据用例查询评审记录
//
// 根据用例查询评审记录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowTestCaseReviews(request *model.ShowTestCaseReviewsRequest) (*model.ShowTestCaseReviewsResponse, error) {
	requestDef := GenReqDefForShowTestCaseReviews()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTestCaseReviewsResponse), nil
	}
}

// ShowTestCaseReviewsInvoker 根据用例查询评审记录
func (c *CloudtestClient) ShowTestCaseReviewsInvoker(request *model.ShowTestCaseReviewsRequest) *ShowTestCaseReviewsInvoker {
	requestDef := GenReqDefForShowTestCaseReviews()
	return &ShowTestCaseReviewsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTestCasesChangeStatistics 版本测试用例变更统计（只统计分支，不统计基线）
//
// 版本测试用例变更统计（只统计分支，不统计基线）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowTestCasesChangeStatistics(request *model.ShowTestCasesChangeStatisticsRequest) (*model.ShowTestCasesChangeStatisticsResponse, error) {
	requestDef := GenReqDefForShowTestCasesChangeStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTestCasesChangeStatisticsResponse), nil
	}
}

// ShowTestCasesChangeStatisticsInvoker 版本测试用例变更统计（只统计分支，不统计基线）
func (c *CloudtestClient) ShowTestCasesChangeStatisticsInvoker(request *model.ShowTestCasesChangeStatisticsRequest) *ShowTestCasesChangeStatisticsInvoker {
	requestDef := GenReqDefForShowTestCasesChangeStatistics()
	return &ShowTestCasesChangeStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTestcaseByPage 根据条件分页获取测试用例对象V2
//
// 根据条件分页获取测试用例对象V2
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowTestcaseByPage(request *model.ShowTestcaseByPageRequest) (*model.ShowTestcaseByPageResponse, error) {
	requestDef := GenReqDefForShowTestcaseByPage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTestcaseByPageResponse), nil
	}
}

// ShowTestcaseByPageInvoker 根据条件分页获取测试用例对象V2
func (c *CloudtestClient) ShowTestcaseByPageInvoker(request *model.ShowTestcaseByPageRequest) *ShowTestcaseByPageInvoker {
	requestDef := GenReqDefForShowTestcaseByPage()
	return &ShowTestcaseByPageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTestpointByPage 根据条件分页获取测试点对象V2
//
// 根据条件分页获取测试点对象V2
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) ShowTestpointByPage(request *model.ShowTestpointByPageRequest) (*model.ShowTestpointByPageResponse, error) {
	requestDef := GenReqDefForShowTestpointByPage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTestpointByPageResponse), nil
	}
}

// ShowTestpointByPageInvoker 根据条件分页获取测试点对象V2
func (c *CloudtestClient) ShowTestpointByPageInvoker(request *model.ShowTestpointByPageRequest) *ShowTestpointByPageInvoker {
	requestDef := GenReqDefForShowTestpointByPage()
	return &ShowTestpointByPageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// UpdateBasicAwById 修改关键字信息接口
//
// 修改关键字信息接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) UpdateBasicAwById(request *model.UpdateBasicAwByIdRequest) (*model.UpdateBasicAwByIdResponse, error) {
	requestDef := GenReqDefForUpdateBasicAwById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBasicAwByIdResponse), nil
	}
}

// UpdateBasicAwByIdInvoker 修改关键字信息接口
func (c *CloudtestClient) UpdateBasicAwByIdInvoker(request *model.UpdateBasicAwByIdRequest) *UpdateBasicAwByIdInvoker {
	requestDef := GenReqDefForUpdateBasicAwById()
	return &UpdateBasicAwByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateIterator 修改测试计划
//
// 修改测试计划
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) UpdateIterator(request *model.UpdateIteratorRequest) (*model.UpdateIteratorResponse, error) {
	requestDef := GenReqDefForUpdateIterator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIteratorResponse), nil
	}
}

// UpdateIteratorInvoker 修改测试计划
func (c *CloudtestClient) UpdateIteratorInvoker(request *model.UpdateIteratorRequest) *UpdateIteratorInvoker {
	requestDef := GenReqDefForUpdateIterator()
	return &UpdateIteratorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// UpdateTestCaseAndScript 更新tmss用例和用例脚本
//
// 更新tmss用例和用例脚本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) UpdateTestCaseAndScript(request *model.UpdateTestCaseAndScriptRequest) (*model.UpdateTestCaseAndScriptResponse, error) {
	requestDef := GenReqDefForUpdateTestCaseAndScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTestCaseAndScriptResponse), nil
	}
}

// UpdateTestCaseAndScriptInvoker 更新tmss用例和用例脚本
func (c *CloudtestClient) UpdateTestCaseAndScriptInvoker(request *model.UpdateTestCaseAndScriptRequest) *UpdateTestCaseAndScriptInvoker {
	requestDef := GenReqDefForUpdateTestCaseAndScript()
	return &UpdateTestCaseAndScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTestCaseComment 修改用例评论
//
// 修改用例评论
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) UpdateTestCaseComment(request *model.UpdateTestCaseCommentRequest) (*model.UpdateTestCaseCommentResponse, error) {
	requestDef := GenReqDefForUpdateTestCaseComment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTestCaseCommentResponse), nil
	}
}

// UpdateTestCaseCommentInvoker 修改用例评论
func (c *CloudtestClient) UpdateTestCaseCommentInvoker(request *model.UpdateTestCaseCommentRequest) *UpdateTestCaseCommentInvoker {
	requestDef := GenReqDefForUpdateTestCaseComment()
	return &UpdateTestCaseCommentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// UpdateUserDnsMapping 更新用户DNS映射
//
// 更新用户DNS映射，执行器自定义映射
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) UpdateUserDnsMapping(request *model.UpdateUserDnsMappingRequest) (*model.UpdateUserDnsMappingResponse, error) {
	requestDef := GenReqDefForUpdateUserDnsMapping()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserDnsMappingResponse), nil
	}
}

// UpdateUserDnsMappingInvoker 更新用户DNS映射
func (c *CloudtestClient) UpdateUserDnsMappingInvoker(request *model.UpdateUserDnsMappingRequest) *UpdateUserDnsMappingInvoker {
	requestDef := GenReqDefForUpdateUserDnsMapping()
	return &UpdateUserDnsMappingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVersionTestCase 在分支或者测试计划下修改用例
//
// 在分支或者测试计划下修改用例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudtestClient) UpdateVersionTestCase(request *model.UpdateVersionTestCaseRequest) (*model.UpdateVersionTestCaseResponse, error) {
	requestDef := GenReqDefForUpdateVersionTestCase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVersionTestCaseResponse), nil
	}
}

// UpdateVersionTestCaseInvoker 在分支或者测试计划下修改用例
func (c *CloudtestClient) UpdateVersionTestCaseInvoker(request *model.UpdateVersionTestCaseRequest) *UpdateVersionTestCaseInvoker {
	requestDef := GenReqDefForUpdateVersionTestCase()
	return &UpdateVersionTestCaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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
