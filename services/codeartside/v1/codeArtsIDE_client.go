package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codeartside/v1/model"
)

type CodeArtsIDEClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCodeArtsIDEClient(hcClient *httpclient.HcHttpClient) *CodeArtsIDEClient {
	return &CodeArtsIDEClient{HcClient: hcClient}
}

func CodeArtsIDEClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// ShowLatestUpgradableRelease 查询升级版本
//
// 查询升级版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsIDEClient) ShowLatestUpgradableRelease(request *model.ShowLatestUpgradableReleaseRequest) (*model.ShowLatestUpgradableReleaseResponse, error) {
	requestDef := GenReqDefForShowLatestUpgradableRelease()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLatestUpgradableReleaseResponse), nil
	}
}

// ShowLatestUpgradableReleaseInvoker 查询升级版本
func (c *CodeArtsIDEClient) ShowLatestUpgradableReleaseInvoker(request *model.ShowLatestUpgradableReleaseRequest) *ShowLatestUpgradableReleaseInvoker {
	requestDef := GenReqDefForShowLatestUpgradableRelease()
	return &ShowLatestUpgradableReleaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ValidateWhitelistUser 是否白名单用户
//
// 是否白名单用户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeArtsIDEClient) ValidateWhitelistUser(request *model.ValidateWhitelistUserRequest) (*model.ValidateWhitelistUserResponse, error) {
	requestDef := GenReqDefForValidateWhitelistUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateWhitelistUserResponse), nil
	}
}

// ValidateWhitelistUserInvoker 是否白名单用户
func (c *CodeArtsIDEClient) ValidateWhitelistUserInvoker(request *model.ValidateWhitelistUserRequest) *ValidateWhitelistUserInvoker {
	requestDef := GenReqDefForValidateWhitelistUser()
	return &ValidateWhitelistUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
