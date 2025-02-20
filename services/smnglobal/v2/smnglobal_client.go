package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/smnglobal/v2/model"
)

type SmnglobalClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewSmnglobalClient(hcClient *httpclient.HcHttpClient) *SmnglobalClient {
	return &SmnglobalClient{HcClient: hcClient}
}

func SmnglobalClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// CreateSubscriptionUser 添加订阅用户
//
// 添加订阅用户。如果订阅用户的状态为未确认，则会向订阅用户发送一条确认订阅消息。订阅用户点击订阅链接确认订阅后，则订阅用户的状态变更为已确认，同时会向订阅用户发送一条取消订阅消息，便于订阅用户随时可以取消订阅。订阅用户点击取消订阅链接后，则订阅用户的状态变更为已取消，同时会向订阅用户发送一条重新订阅消息，便于订阅用户可以重新订阅。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnglobalClient) CreateSubscriptionUser(request *model.CreateSubscriptionUserRequest) (*model.CreateSubscriptionUserResponse, error) {
	requestDef := GenReqDefForCreateSubscriptionUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSubscriptionUserResponse), nil
	}
}

// CreateSubscriptionUserInvoker 添加订阅用户
func (c *SmnglobalClient) CreateSubscriptionUserInvoker(request *model.CreateSubscriptionUserRequest) *CreateSubscriptionUserInvoker {
	requestDef := GenReqDefForCreateSubscriptionUser()
	return &CreateSubscriptionUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSubscriptionUser 删除订阅用户
//
// 删除订阅用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnglobalClient) DeleteSubscriptionUser(request *model.DeleteSubscriptionUserRequest) (*model.DeleteSubscriptionUserResponse, error) {
	requestDef := GenReqDefForDeleteSubscriptionUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSubscriptionUserResponse), nil
	}
}

// DeleteSubscriptionUserInvoker 删除订阅用户
func (c *SmnglobalClient) DeleteSubscriptionUserInvoker(request *model.DeleteSubscriptionUserRequest) *DeleteSubscriptionUserInvoker {
	requestDef := GenReqDefForDeleteSubscriptionUser()
	return &DeleteSubscriptionUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubscriptionUser 查询订阅用户列表
//
// 查询订阅用户列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnglobalClient) ListSubscriptionUser(request *model.ListSubscriptionUserRequest) (*model.ListSubscriptionUserResponse, error) {
	requestDef := GenReqDefForListSubscriptionUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubscriptionUserResponse), nil
	}
}

// ListSubscriptionUserInvoker 查询订阅用户列表
func (c *SmnglobalClient) ListSubscriptionUserInvoker(request *model.ListSubscriptionUserRequest) *ListSubscriptionUserInvoker {
	requestDef := GenReqDefForListSubscriptionUser()
	return &ListSubscriptionUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSubscriptionUser 更新订阅用户
//
// 更新订阅用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SmnglobalClient) UpdateSubscriptionUser(request *model.UpdateSubscriptionUserRequest) (*model.UpdateSubscriptionUserResponse, error) {
	requestDef := GenReqDefForUpdateSubscriptionUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSubscriptionUserResponse), nil
	}
}

// UpdateSubscriptionUserInvoker 更新订阅用户
func (c *SmnglobalClient) UpdateSubscriptionUserInvoker(request *model.UpdateSubscriptionUserRequest) *UpdateSubscriptionUserInvoker {
	requestDef := GenReqDefForUpdateSubscriptionUser()
	return &UpdateSubscriptionUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
