package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cse/v1/model"
)

type CseClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCseClient(hcClient *http_client.HcHttpClient) *CseClient {
	return &CseClient{HcClient: hcClient}
}

func CseClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CreateEngine 创建微服务引擎专享版
//
// 创建微服务引擎专享版。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) CreateEngine(request *model.CreateEngineRequest) (*model.CreateEngineResponse, error) {
	requestDef := GenReqDefForCreateEngine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEngineResponse), nil
	}
}

// CreateEngineInvoker 创建微服务引擎专享版
func (c *CseClient) CreateEngineInvoker(request *model.CreateEngineRequest) *CreateEngineInvoker {
	requestDef := GenReqDefForCreateEngine()
	return &CreateEngineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEngine 删除微服务引擎专享版
//
// 删除微服务引擎专享版。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) DeleteEngine(request *model.DeleteEngineRequest) (*model.DeleteEngineResponse, error) {
	requestDef := GenReqDefForDeleteEngine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEngineResponse), nil
	}
}

// DeleteEngineInvoker 删除微服务引擎专享版
func (c *CseClient) DeleteEngineInvoker(request *model.DeleteEngineRequest) *DeleteEngineInvoker {
	requestDef := GenReqDefForDeleteEngine()
	return &DeleteEngineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadKie 导出kie配置
//
// 导出kie配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) DownloadKie(request *model.DownloadKieRequest) (*model.DownloadKieResponse, error) {
	requestDef := GenReqDefForDownloadKie()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadKieResponse), nil
	}
}

// DownloadKieInvoker 导出kie配置
func (c *CseClient) DownloadKieInvoker(request *model.DownloadKieRequest) *DownloadKieInvoker {
	requestDef := GenReqDefForDownloadKie()
	return &DownloadKieInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEngines 查询微服务引擎列表
//
// 查询微服务引擎列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) ListEngines(request *model.ListEnginesRequest) (*model.ListEnginesResponse, error) {
	requestDef := GenReqDefForListEngines()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnginesResponse), nil
	}
}

// ListEnginesInvoker 查询微服务引擎列表
func (c *CseClient) ListEnginesInvoker(request *model.ListEnginesRequest) *ListEnginesInvoker {
	requestDef := GenReqDefForListEngines()
	return &ListEnginesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlavors 查询微服务引擎专享版的规格列表
//
// 查询微服务引擎专享版的规格列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsResponse), nil
	}
}

// ListFlavorsInvoker 查询微服务引擎专享版的规格列表
func (c *CseClient) ListFlavorsInvoker(request *model.ListFlavorsRequest) *ListFlavorsInvoker {
	requestDef := GenReqDefForListFlavors()
	return &ListFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RetryEngine 对微服务引擎专享版进行重试
//
// 对微服务引擎专享版进行重试
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) RetryEngine(request *model.RetryEngineRequest) (*model.RetryEngineResponse, error) {
	requestDef := GenReqDefForRetryEngine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RetryEngineResponse), nil
	}
}

// RetryEngineInvoker 对微服务引擎专享版进行重试
func (c *CseClient) RetryEngineInvoker(request *model.RetryEngineRequest) *RetryEngineInvoker {
	requestDef := GenReqDefForRetryEngine()
	return &RetryEngineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEngine 查询微服务引擎专享版详情
//
// 查询微服务引擎专享版详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) ShowEngine(request *model.ShowEngineRequest) (*model.ShowEngineResponse, error) {
	requestDef := GenReqDefForShowEngine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEngineResponse), nil
	}
}

// ShowEngineInvoker 查询微服务引擎专享版详情
func (c *CseClient) ShowEngineInvoker(request *model.ShowEngineRequest) *ShowEngineInvoker {
	requestDef := GenReqDefForShowEngine()
	return &ShowEngineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEngineJob 查询微服务引擎任务详情
//
// 查询微服务引擎任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) ShowEngineJob(request *model.ShowEngineJobRequest) (*model.ShowEngineJobResponse, error) {
	requestDef := GenReqDefForShowEngineJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEngineJobResponse), nil
	}
}

// ShowEngineJobInvoker 查询微服务引擎任务详情
func (c *CseClient) ShowEngineJobInvoker(request *model.ShowEngineJobRequest) *ShowEngineJobInvoker {
	requestDef := GenReqDefForShowEngineJob()
	return &ShowEngineJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpgradeEngine 升级微服务引擎专享版
//
// 升级微服务引擎专享版
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) UpgradeEngine(request *model.UpgradeEngineRequest) (*model.UpgradeEngineResponse, error) {
	requestDef := GenReqDefForUpgradeEngine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpgradeEngineResponse), nil
	}
}

// UpgradeEngineInvoker 升级微服务引擎专享版
func (c *CseClient) UpgradeEngineInvoker(request *model.UpgradeEngineRequest) *UpgradeEngineInvoker {
	requestDef := GenReqDefForUpgradeEngine()
	return &UpgradeEngineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadKie 导入kie配置
//
// 导入kie配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) UploadKie(request *model.UploadKieRequest) (*model.UploadKieResponse, error) {
	requestDef := GenReqDefForUploadKie()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadKieResponse), nil
	}
}

// UploadKieInvoker 导入kie配置
func (c *CseClient) UploadKieInvoker(request *model.UploadKieRequest) *UploadKieInvoker {
	requestDef := GenReqDefForUploadKie()
	return &UploadKieInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
