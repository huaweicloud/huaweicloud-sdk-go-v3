package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/mapds/v1/model"
)

type MapDSClient struct {
	HcClient *http_client.HcHttpClient
}

func NewMapDSClient(hcClient *http_client.HcHttpClient) *MapDSClient {
	return &MapDSClient{HcClient: hcClient}
}

func MapDSClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CreateCredential 创建凭证
//
// 该接口用于创建访问地图数据服务的凭证。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MapDSClient) CreateCredential(request *model.CreateCredentialRequest) (*model.CreateCredentialResponse, error) {
	requestDef := GenReqDefForCreateCredential()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCredentialResponse), nil
	}
}

// CreateCredentialInvoker 创建凭证
func (c *MapDSClient) CreateCredentialInvoker(request *model.CreateCredentialRequest) *CreateCredentialInvoker {
	requestDef := GenReqDefForCreateCredential()
	return &CreateCredentialInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSasToken 创建SAS Token
//
// 创建SAS token凭据，用于访问瓦片。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MapDSClient) CreateSasToken(request *model.CreateSasTokenRequest) (*model.CreateSasTokenResponse, error) {
	requestDef := GenReqDefForCreateSasToken()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSasTokenResponse), nil
	}
}

// CreateSasTokenInvoker 创建SAS Token
func (c *MapDSClient) CreateSasTokenInvoker(request *model.CreateSasTokenRequest) *CreateSasTokenInvoker {
	requestDef := GenReqDefForCreateSasToken()
	return &CreateSasTokenInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCedential 删除凭证
//
// 该接口用于删除访问地图数据服务的凭证。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MapDSClient) DeleteCedential(request *model.DeleteCedentialRequest) (*model.DeleteCedentialResponse, error) {
	requestDef := GenReqDefForDeleteCedential()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCedentialResponse), nil
	}
}

// DeleteCedentialInvoker 删除凭证
func (c *MapDSClient) DeleteCedentialInvoker(request *model.DeleteCedentialRequest) *DeleteCedentialInvoker {
	requestDef := GenReqDefForDeleteCedential()
	return &DeleteCedentialInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCredential 查询凭证
//
// 该接口用于查询访问地图数据服务的凭证。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MapDSClient) ShowCredential(request *model.ShowCredentialRequest) (*model.ShowCredentialResponse, error) {
	requestDef := GenReqDefForShowCredential()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCredentialResponse), nil
	}
}

// ShowCredentialInvoker 查询凭证
func (c *MapDSClient) ShowCredentialInvoker(request *model.ShowCredentialRequest) *ShowCredentialInvoker {
	requestDef := GenReqDefForShowCredential()
	return &ShowCredentialInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMapTile 获取地图瓦片
//
// 该接口用于获取地图瓦片。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MapDSClient) ShowMapTile(request *model.ShowMapTileRequest) (*model.ShowMapTileResponse, error) {
	requestDef := GenReqDefForShowMapTile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMapTileResponse), nil
	}
}

// ShowMapTileInvoker 获取地图瓦片
func (c *MapDSClient) ShowMapTileInvoker(request *model.ShowMapTileRequest) *ShowMapTileInvoker {
	requestDef := GenReqDefForShowMapTile()
	return &ShowMapTileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
