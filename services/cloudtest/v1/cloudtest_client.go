package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cloudtest/v1/model"
)

type CloudtestClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCloudtestClient(hcClient *http_client.HcHttpClient) *CloudtestClient {
	return &CloudtestClient{HcClient: hcClient}
}

func CloudtestClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
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
