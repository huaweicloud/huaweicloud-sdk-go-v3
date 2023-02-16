package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/imagesearch/v2/model"
)

type ImageSearchClient struct {
	HcClient *http_client.HcHttpClient
}

func NewImageSearchClient(hcClient *http_client.HcHttpClient) *ImageSearchClient {
	return &ImageSearchClient{HcClient: hcClient}
}

func ImageSearchClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// RunAddData 添加数据
//
// 添加数据到指定服务实例中。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageSearchClient) RunAddData(request *model.RunAddDataRequest) (*model.RunAddDataResponse, error) {
	requestDef := GenReqDefForRunAddData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunAddDataResponse), nil
	}
}

// RunAddDataInvoker 添加数据
func (c *ImageSearchClient) RunAddDataInvoker(request *model.RunAddDataRequest) *RunAddDataInvoker {
	requestDef := GenReqDefForRunAddData()
	return &RunAddDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunCheckData 检查数据
//
// 检查指定服务实例中的对应数据，支持指定ID检查和条件检查。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageSearchClient) RunCheckData(request *model.RunCheckDataRequest) (*model.RunCheckDataResponse, error) {
	requestDef := GenReqDefForRunCheckData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunCheckDataResponse), nil
	}
}

// RunCheckDataInvoker 检查数据
func (c *ImageSearchClient) RunCheckDataInvoker(request *model.RunCheckDataRequest) *RunCheckDataInvoker {
	requestDef := GenReqDefForRunCheckData()
	return &RunCheckDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunDeleteData 删除数据
//
// 删除指定服务实例中的对应数据，支持指定ID删除和条件删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageSearchClient) RunDeleteData(request *model.RunDeleteDataRequest) (*model.RunDeleteDataResponse, error) {
	requestDef := GenReqDefForRunDeleteData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunDeleteDataResponse), nil
	}
}

// RunDeleteDataInvoker 删除数据
func (c *ImageSearchClient) RunDeleteDataInvoker(request *model.RunDeleteDataRequest) *RunDeleteDataInvoker {
	requestDef := GenReqDefForRunDeleteData()
	return &RunDeleteDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunSearch 搜索
//
// 从指定服务实例中进行搜索。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageSearchClient) RunSearch(request *model.RunSearchRequest) (*model.RunSearchResponse, error) {
	requestDef := GenReqDefForRunSearch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunSearchResponse), nil
	}
}

// RunSearchInvoker 搜索
func (c *ImageSearchClient) RunSearchInvoker(request *model.RunSearchRequest) *RunSearchInvoker {
	requestDef := GenReqDefForRunSearch()
	return &RunSearchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunUpdateData 更新数据
//
// 更新指定服务实例中的对应数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ImageSearchClient) RunUpdateData(request *model.RunUpdateDataRequest) (*model.RunUpdateDataResponse, error) {
	requestDef := GenReqDefForRunUpdateData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunUpdateDataResponse), nil
	}
}

// RunUpdateDataInvoker 更新数据
func (c *ImageSearchClient) RunUpdateDataInvoker(request *model.RunUpdateDataRequest) *RunUpdateDataInvoker {
	requestDef := GenReqDefForRunUpdateData()
	return &RunUpdateDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
