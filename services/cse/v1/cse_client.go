package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

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

// 创建微服务引擎专享版
//
// 创建微服务引擎专享版。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CseClient) CreateEngine(request *model.CreateEngineRequest) (*model.CreateEngineResponse, error) {
	requestDef := GenReqDefForCreateEngine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEngineResponse), nil
	}
}

// 删除微服务引擎专享版
//
// 删除微服务引擎专享版。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CseClient) DeleteEngine(request *model.DeleteEngineRequest) (*model.DeleteEngineResponse, error) {
	requestDef := GenReqDefForDeleteEngine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEngineResponse), nil
	}
}

// 导出kie配置
//
// 导出kie配置
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CseClient) DownloadKie(request *model.DownloadKieRequest) (*model.DownloadKieResponse, error) {
	requestDef := GenReqDefForDownloadKie()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadKieResponse), nil
	}
}

// 查询微服务引擎列表
//
// 查询微服务引擎列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CseClient) ListEngines(request *model.ListEnginesRequest) (*model.ListEnginesResponse, error) {
	requestDef := GenReqDefForListEngines()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnginesResponse), nil
	}
}

// 查询微服务引擎专享版的规格列表
//
// 查询微服务引擎专享版的规格列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CseClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsResponse), nil
	}
}

// 查询微服务引擎专享版详情
//
// 查询微服务引擎专享版详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CseClient) ShowEngine(request *model.ShowEngineRequest) (*model.ShowEngineResponse, error) {
	requestDef := GenReqDefForShowEngine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEngineResponse), nil
	}
}

// 查询微服务引擎任务详情
//
// 查询微服务引擎任务详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CseClient) ShowEngineJob(request *model.ShowEngineJobRequest) (*model.ShowEngineJobResponse, error) {
	requestDef := GenReqDefForShowEngineJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEngineJobResponse), nil
	}
}

// 导入kie配置
//
// 导入kie配置
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CseClient) UploadKie(request *model.UploadKieRequest) (*model.UploadKieResponse, error) {
	requestDef := GenReqDefForUploadKie()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadKieResponse), nil
	}
}
