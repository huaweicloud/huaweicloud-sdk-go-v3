package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

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

// 批量删除自定义测试服务类型用例
//
// 批量删除自定义测试服务类型用例
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudtestClient) BatchDeleteTestCase(request *model.BatchDeleteTestCaseRequest) (*model.BatchDeleteTestCaseResponse, error) {
	requestDef := GenReqDefForBatchDeleteTestCase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteTestCaseResponse), nil
	}
}

// 项目下创建计划
//
// 项目下创建计划
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudtestClient) CreatePlan(request *model.CreatePlanRequest) (*model.CreatePlanResponse, error) {
	requestDef := GenReqDefForCreatePlan()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePlanResponse), nil
	}
}

// 新测试类型服务注册到云测
//
// 通过接口CreateService注册成为云测的自定义服务。 注册完成后云测界面将会出现此自定义测试类型。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudtestClient) CreateService(request *model.CreateServiceRequest) (*model.CreateServiceResponse, error) {
	requestDef := GenReqDefForCreateService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateServiceResponse), nil
	}
}

// 创建自定义测试服务类型用例
//
// 创建自定义测试服务类型用例
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudtestClient) CreateTestCase(request *model.CreateTestCaseRequest) (*model.CreateTestCaseResponse, error) {
	requestDef := GenReqDefForCreateTestCase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTestCaseResponse), nil
	}
}

// 计划中批量添加测试用例
//
// 计划中批量添加测试用例
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudtestClient) CreateTestCaseInPlan(request *model.CreateTestCaseInPlanRequest) (*model.CreateTestCaseInPlanResponse, error) {
	requestDef := GenReqDefForCreateTestCaseInPlan()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTestCaseInPlanResponse), nil
	}
}

// 删除已注册服务
//
// 删除已注册服务
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudtestClient) DeleteService(request *model.DeleteServiceRequest) (*model.DeleteServiceResponse, error) {
	requestDef := GenReqDefForDeleteService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServiceResponse), nil
	}
}

// 批量执行测试用例
//
// 批量执行测试用例
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudtestClient) RunTestCase(request *model.RunTestCaseRequest) (*model.RunTestCaseResponse, error) {
	requestDef := GenReqDefForRunTestCase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunTestCaseResponse), nil
	}
}

// 查询某个测试计划下的需求树
//
// 查询某个测试计划下的需求列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudtestClient) ShowIssuesByPlanId(request *model.ShowIssuesByPlanIdRequest) (*model.ShowIssuesByPlanIdResponse, error) {
	requestDef := GenReqDefForShowIssuesByPlanId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIssuesByPlanIdResponse), nil
	}
}

// 查询某测试计划下的操作历史
//
// 查询某测试计划下的操作历史
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudtestClient) ShowPlanJournals(request *model.ShowPlanJournalsRequest) (*model.ShowPlanJournalsResponse, error) {
	requestDef := GenReqDefForShowPlanJournals()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlanJournalsResponse), nil
	}
}

// 项目下查询测试计划列表v2
//
// 项目下查询测试计划列表v2
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudtestClient) ShowPlanList(request *model.ShowPlanListRequest) (*model.ShowPlanListResponse, error) {
	requestDef := GenReqDefForShowPlanList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlanListResponse), nil
	}
}

// 项目下查询测试计划列表
//
// 项目下查询测试计划列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudtestClient) ShowPlans(request *model.ShowPlansRequest) (*model.ShowPlansResponse, error) {
	requestDef := GenReqDefForShowPlans()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlansResponse), nil
	}
}

// 用户获取自己当前已经注册的服务
//
// 用户获取自己当前已经注册的服务
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudtestClient) ShowRegisterService(request *model.ShowRegisterServiceRequest) (*model.ShowRegisterServiceResponse, error) {
	requestDef := GenReqDefForShowRegisterService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRegisterServiceResponse), nil
	}
}

// 获取测试用例详情
//
// 获取测试用例详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudtestClient) ShowTestCaseDetail(request *model.ShowTestCaseDetailRequest) (*model.ShowTestCaseDetailResponse, error) {
	requestDef := GenReqDefForShowTestCaseDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTestCaseDetailResponse), nil
	}
}

// 通过用例编号获取测试用例详情
//
// 通过用例编号获取测试用例详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudtestClient) ShowTestCaseDetailV2(request *model.ShowTestCaseDetailV2Request) (*model.ShowTestCaseDetailV2Response, error) {
	requestDef := GenReqDefForShowTestCaseDetailV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTestCaseDetailV2Response), nil
	}
}

// 更新已注册服务
//
// 更新已注册服务
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudtestClient) UpdateService(request *model.UpdateServiceRequest) (*model.UpdateServiceResponse, error) {
	requestDef := GenReqDefForUpdateService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateServiceResponse), nil
	}
}

// 更新自定义测试服务类型用例
//
// 更新自定义测试服务类型用例
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudtestClient) UpdateTestCase(request *model.UpdateTestCaseRequest) (*model.UpdateTestCaseResponse, error) {
	requestDef := GenReqDefForUpdateTestCase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTestCaseResponse), nil
	}
}

// 批量更新测试用例结果
//
// 批量更新测试用例结果
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudtestClient) UpdateTestCaseResult(request *model.UpdateTestCaseResultRequest) (*model.UpdateTestCaseResultResponse, error) {
	requestDef := GenReqDefForUpdateTestCaseResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTestCaseResultResponse), nil
	}
}

// 通过导入仓库中的文件生成接口测试套
//
// 通过导入仓库中的文件生成接口测试套
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudtestClient) CreateApiTestSuiteByRepoFile(request *model.CreateApiTestSuiteByRepoFileRequest) (*model.CreateApiTestSuiteByRepoFileResponse, error) {
	requestDef := GenReqDefForCreateApiTestSuiteByRepoFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApiTestSuiteByRepoFileResponse), nil
	}
}

// 获取云测的环境参数分组列表
//
// 获取云测的环境参数分组列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CloudtestClient) ListEnvironments(request *model.ListEnvironmentsRequest) (*model.ListEnvironmentsResponse, error) {
	requestDef := GenReqDefForListEnvironments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnvironmentsResponse), nil
	}
}
