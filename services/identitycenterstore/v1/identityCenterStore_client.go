package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/identitycenterstore/v1/model"
)

type IdentityCenterStoreClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewIdentityCenterStoreClient(hcClient *httpclient.HcHttpClient) *IdentityCenterStoreClient {
	return &IdentityCenterStoreClient{HcClient: hcClient}
}

func IdentityCenterStoreClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// ListUsers 列出用户
//
// 查询指定身份源下的IAM身份中心用户列表。
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
