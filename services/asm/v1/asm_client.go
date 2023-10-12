package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/asm/v1/model"
)

type AsmClient struct {
	HcClient *http_client.HcHttpClient
}

func NewAsmClient(hcClient *http_client.HcHttpClient) *AsmClient {
	return &AsmClient{HcClient: hcClient}
}

func AsmClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder().WithCredentialsType("basic.Credentials")
	return builder
}

// CreateMesh 新建网格
//
// 该API用于创建一个网格。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsmClient) CreateMesh(request *model.CreateMeshRequest) (*model.CreateMeshResponse, error) {
	requestDef := GenReqDefForCreateMesh()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMeshResponse), nil
	}
}

// CreateMeshInvoker 新建网格
func (c *AsmClient) CreateMeshInvoker(request *model.CreateMeshRequest) *CreateMeshInvoker {
	requestDef := GenReqDefForCreateMesh()
	return &CreateMeshInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMesh 删除网格
//
// 该API用于删除一个指定的网格。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsmClient) DeleteMesh(request *model.DeleteMeshRequest) (*model.DeleteMeshResponse, error) {
	requestDef := GenReqDefForDeleteMesh()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMeshResponse), nil
	}
}

// DeleteMeshInvoker 删除网格
func (c *AsmClient) DeleteMeshInvoker(request *model.DeleteMeshRequest) *DeleteMeshInvoker {
	requestDef := GenReqDefForDeleteMesh()
	return &DeleteMeshInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMeshes 查询网格列表
//
// 该API用于获取用户所有网格的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsmClient) ListMeshes(request *model.ListMeshesRequest) (*model.ListMeshesResponse, error) {
	requestDef := GenReqDefForListMeshes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMeshesResponse), nil
	}
}

// ListMeshesInvoker 查询网格列表
func (c *AsmClient) ListMeshesInvoker(request *model.ListMeshesRequest) *ListMeshesInvoker {
	requestDef := GenReqDefForListMeshes()
	return &ListMeshesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMesh 查询网格
//
// 该API用于获取指定网格的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AsmClient) ShowMesh(request *model.ShowMeshRequest) (*model.ShowMeshResponse, error) {
	requestDef := GenReqDefForShowMesh()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMeshResponse), nil
	}
}

// ShowMeshInvoker 查询网格
func (c *AsmClient) ShowMeshInvoker(request *model.ShowMeshRequest) *ShowMeshInvoker {
	requestDef := GenReqDefForShowMesh()
	return &ShowMeshInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
