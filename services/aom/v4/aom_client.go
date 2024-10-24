package v4

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/aom/v4/model"
)

type AomClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewAomClient(hcClient *httpclient.HcHttpClient) *AomClient {
	return &AomClient{HcClient: hcClient}
}

func AomClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// BatchImportAgent 下发批量安装UniAgent任务
//
// 该接口用于下发批量安装UniAgent任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) BatchImportAgent(request *model.BatchImportAgentRequest) (*model.BatchImportAgentResponse, error) {
	requestDef := GenReqDefForBatchImportAgent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchImportAgentResponse), nil
	}
}

// BatchImportAgentInvoker 下发批量安装UniAgent任务
func (c *AomClient) BatchImportAgentInvoker(request *model.BatchImportAgentRequest) *BatchImportAgentInvoker {
	requestDef := GenReqDefForBatchImportAgent()
	return &BatchImportAgentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateAgent 下发批量升级UniAgent任务
//
// 该接口用于下发批量升级UniAgent任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) BatchUpdateAgent(request *model.BatchUpdateAgentRequest) (*model.BatchUpdateAgentResponse, error) {
	requestDef := GenReqDefForBatchUpdateAgent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateAgentResponse), nil
	}
}

// BatchUpdateAgentInvoker 下发批量升级UniAgent任务
func (c *AomClient) BatchUpdateAgentInvoker(request *model.BatchUpdateAgentRequest) *BatchUpdateAgentInvoker {
	requestDef := GenReqDefForBatchUpdateAgent()
	return &BatchUpdateAgentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAgentInfos 查询UniAgent主机列表信息
//
// 该接口用于查询执行过安装UniAgent任务的主机列表信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AomClient) ShowAgentInfos(request *model.ShowAgentInfosRequest) (*model.ShowAgentInfosResponse, error) {
	requestDef := GenReqDefForShowAgentInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAgentInfosResponse), nil
	}
}

// ShowAgentInfosInvoker 查询UniAgent主机列表信息
func (c *AomClient) ShowAgentInfosInvoker(request *model.ShowAgentInfosRequest) *ShowAgentInfosInvoker {
	requestDef := GenReqDefForShowAgentInfos()
	return &ShowAgentInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
