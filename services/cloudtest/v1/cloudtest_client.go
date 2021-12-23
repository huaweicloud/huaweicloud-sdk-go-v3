package v1

import (
	http_client "github.com/RandolphCYG/huaweicloud-sdk-go-v3/core"

	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/services/cloudtest/v1/model"
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

//通过导入仓库中的文件生成接口测试套
func (c *CloudtestClient) CreateApiTestSuiteByRepoFile(request *model.CreateApiTestSuiteByRepoFileRequest) (*model.CreateApiTestSuiteByRepoFileResponse, error) {
	requestDef := GenReqDefForCreateApiTestSuiteByRepoFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApiTestSuiteByRepoFileResponse), nil
	}
}

//获取云测的环境参数分组列表
func (c *CloudtestClient) ListEnvironments(request *model.ListEnvironmentsRequest) (*model.ListEnvironmentsResponse, error) {
	requestDef := GenReqDefForListEnvironments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnvironmentsResponse), nil
	}
}

//项目下创建计划
func (c *CloudtestClient) CreatePlan(request *model.CreatePlanRequest) (*model.CreatePlanResponse, error) {
	requestDef := GenReqDefForCreatePlan()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePlanResponse), nil
	}
}

//计划中批量添加测试用例
func (c *CloudtestClient) CreateTestCaseInPlan(request *model.CreateTestCaseInPlanRequest) (*model.CreateTestCaseInPlanResponse, error) {
	requestDef := GenReqDefForCreateTestCaseInPlan()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTestCaseInPlanResponse), nil
	}
}

//项目下查询测试计划列表
func (c *CloudtestClient) ShowPlans(request *model.ShowPlansRequest) (*model.ShowPlansResponse, error) {
	requestDef := GenReqDefForShowPlans()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlansResponse), nil
	}
}

//新测试类型服务注册到云测
func (c *CloudtestClient) CreateService(request *model.CreateServiceRequest) (*model.CreateServiceResponse, error) {
	requestDef := GenReqDefForCreateService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateServiceResponse), nil
	}
}

//删除已注册服务
func (c *CloudtestClient) DeleteService(request *model.DeleteServiceRequest) (*model.DeleteServiceResponse, error) {
	requestDef := GenReqDefForDeleteService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServiceResponse), nil
	}
}

//用户获取自己当前已经注册的服务
func (c *CloudtestClient) ShowRegisterService(request *model.ShowRegisterServiceRequest) (*model.ShowRegisterServiceResponse, error) {
	requestDef := GenReqDefForShowRegisterService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRegisterServiceResponse), nil
	}
}

//更新已注册服务
func (c *CloudtestClient) UpdateService(request *model.UpdateServiceRequest) (*model.UpdateServiceResponse, error) {
	requestDef := GenReqDefForUpdateService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateServiceResponse), nil
	}
}

//批量删除测试用例
func (c *CloudtestClient) BatchDeleteTestCase(request *model.BatchDeleteTestCaseRequest) (*model.BatchDeleteTestCaseResponse, error) {
	requestDef := GenReqDefForBatchDeleteTestCase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteTestCaseResponse), nil
	}
}

//创建测试用例
func (c *CloudtestClient) CreateTestCase(request *model.CreateTestCaseRequest) (*model.CreateTestCaseResponse, error) {
	requestDef := GenReqDefForCreateTestCase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTestCaseResponse), nil
	}
}

//批量执行测试用例
func (c *CloudtestClient) RunTestCase(request *model.RunTestCaseRequest) (*model.RunTestCaseResponse, error) {
	requestDef := GenReqDefForRunTestCase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunTestCaseResponse), nil
	}
}

//获取测试用例详情
func (c *CloudtestClient) ShowTestCaseDetail(request *model.ShowTestCaseDetailRequest) (*model.ShowTestCaseDetailResponse, error) {
	requestDef := GenReqDefForShowTestCaseDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTestCaseDetailResponse), nil
	}
}

//通过用例编号或用例ID获取测试用例详情
func (c *CloudtestClient) ShowTestCaseDetailV2(request *model.ShowTestCaseDetailV2Request) (*model.ShowTestCaseDetailV2Response, error) {
	requestDef := GenReqDefForShowTestCaseDetailV2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTestCaseDetailV2Response), nil
	}
}

//更新测试用例接口
func (c *CloudtestClient) UpdateTestCase(request *model.UpdateTestCaseRequest) (*model.UpdateTestCaseResponse, error) {
	requestDef := GenReqDefForUpdateTestCase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTestCaseResponse), nil
	}
}

//批量更新测试用例结果
func (c *CloudtestClient) UpdateTestCaseResult(request *model.UpdateTestCaseResultRequest) (*model.UpdateTestCaseResultResponse, error) {
	requestDef := GenReqDefForUpdateTestCaseResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTestCaseResultResponse), nil
	}
}
