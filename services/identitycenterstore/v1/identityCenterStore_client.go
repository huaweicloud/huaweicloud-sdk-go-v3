package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/identitycenterstore/v1/model"
)

type IdentityCenterStoreClient struct {
	HcClient *http_client.HcHttpClient
}

func NewIdentityCenterStoreClient(hcClient *http_client.HcHttpClient) *IdentityCenterStoreClient {
	return &IdentityCenterStoreClient{HcClient: hcClient}
}

func IdentityCenterStoreClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// ListUsers 列出用户
//
// 查询指定IdentityStore下的IdentityCenter用户列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) ListUsers(request *model.ListUsersRequest) (*model.ListUsersResponse, error) {
	requestDef := GenReqDefForListUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsersResponse), nil
	}
}

// ListUsersInvoker 列出用户
func (c *IdentityCenterStoreClient) ListUsersInvoker(request *model.ListUsersRequest) *ListUsersInvoker {
	requestDef := GenReqDefForListUsers()
	return &ListUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
