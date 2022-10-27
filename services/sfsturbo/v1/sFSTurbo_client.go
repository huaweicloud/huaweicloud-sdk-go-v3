package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/sfsturbo/v1/model"
)

type SFSTurboClient struct {
	HcClient *http_client.HcHttpClient
}

func NewSFSTurboClient(hcClient *http_client.HcHttpClient) *SFSTurboClient {
	return &SFSTurboClient{HcClient: hcClient}
}

func SFSTurboClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// BatchAddSharedTags 批量添加共享标签
//
// 指定共享批量添加标签。
//
// 一个共享上最多有10个标签。
// 一个共享上的多个标签的key不允许重复。
// 此接口为幂等接口：如果要添加的key在共享上已存在，则覆盖更新标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SFSTurboClient) BatchAddSharedTags(request *model.BatchAddSharedTagsRequest) (*model.BatchAddSharedTagsResponse, error) {
	requestDef := GenReqDefForBatchAddSharedTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddSharedTagsResponse), nil
	}
}

// BatchAddSharedTagsInvoker 批量添加共享标签
func (c *SFSTurboClient) BatchAddSharedTagsInvoker(request *model.BatchAddSharedTagsRequest) *BatchAddSharedTagsInvoker {
	requestDef := GenReqDefForBatchAddSharedTags()
	return &BatchAddSharedTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeSecurityGroup 修改文件系统绑定的安全组
//
// 修改SFS Turbo文件系统绑定的安全组。修改安全组为异步任务，可以通过“查询单个文件系统”返回的子状态字段“sub_status”来判断是否修改安全组状态，子状态为“232”即为修改安全组成功。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SFSTurboClient) ChangeSecurityGroup(request *model.ChangeSecurityGroupRequest) (*model.ChangeSecurityGroupResponse, error) {
	requestDef := GenReqDefForChangeSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeSecurityGroupResponse), nil
	}
}

// ChangeSecurityGroupInvoker 修改文件系统绑定的安全组
func (c *SFSTurboClient) ChangeSecurityGroupInvoker(request *model.ChangeSecurityGroupRequest) *ChangeSecurityGroupInvoker {
	requestDef := GenReqDefForChangeSecurityGroup()
	return &ChangeSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeShareName 修改文件系统名称
//
// 修改文件系统名称
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SFSTurboClient) ChangeShareName(request *model.ChangeShareNameRequest) (*model.ChangeShareNameResponse, error) {
	requestDef := GenReqDefForChangeShareName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeShareNameResponse), nil
	}
}

// ChangeShareNameInvoker 修改文件系统名称
func (c *SFSTurboClient) ChangeShareNameInvoker(request *model.ChangeShareNameRequest) *ChangeShareNameInvoker {
	requestDef := GenReqDefForChangeShareName()
	return &ChangeShareNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateShare 创建文件系统
//
// 创建文件系统。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SFSTurboClient) CreateShare(request *model.CreateShareRequest) (*model.CreateShareResponse, error) {
	requestDef := GenReqDefForCreateShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateShareResponse), nil
	}
}

// CreateShareInvoker 创建文件系统
func (c *SFSTurboClient) CreateShareInvoker(request *model.CreateShareRequest) *CreateShareInvoker {
	requestDef := GenReqDefForCreateShare()
	return &CreateShareInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSharedTag 创建共享标签
//
// 指定共享添加一个标签。
// 一个共享上最多有10个标签。
// 一个共享上的多个标签的key不允许重复。
// 此接口为幂等接口：如果要添加的key在共享上已存在，则覆盖更新标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SFSTurboClient) CreateSharedTag(request *model.CreateSharedTagRequest) (*model.CreateSharedTagResponse, error) {
	requestDef := GenReqDefForCreateSharedTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSharedTagResponse), nil
	}
}

// CreateSharedTagInvoker 创建共享标签
func (c *SFSTurboClient) CreateSharedTagInvoker(request *model.CreateSharedTagRequest) *CreateSharedTagInvoker {
	requestDef := GenReqDefForCreateSharedTag()
	return &CreateSharedTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteShare 删除文件系统
//
// 删除文件系统。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SFSTurboClient) DeleteShare(request *model.DeleteShareRequest) (*model.DeleteShareResponse, error) {
	requestDef := GenReqDefForDeleteShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteShareResponse), nil
	}
}

// DeleteShareInvoker 删除文件系统
func (c *SFSTurboClient) DeleteShareInvoker(request *model.DeleteShareRequest) *DeleteShareInvoker {
	requestDef := GenReqDefForDeleteShare()
	return &DeleteShareInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSharedTag 删除共享标签
//
// 指定共享删除一个标签。当共享中不存在指定要删除的key时，接口调用将会返回404错误。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SFSTurboClient) DeleteSharedTag(request *model.DeleteSharedTagRequest) (*model.DeleteSharedTagResponse, error) {
	requestDef := GenReqDefForDeleteSharedTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSharedTagResponse), nil
	}
}

// DeleteSharedTagInvoker 删除共享标签
func (c *SFSTurboClient) DeleteSharedTagInvoker(request *model.DeleteSharedTagRequest) *DeleteSharedTagInvoker {
	requestDef := GenReqDefForDeleteSharedTag()
	return &DeleteSharedTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExpandShare 扩容文件系统
//
// 扩容文件系统。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SFSTurboClient) ExpandShare(request *model.ExpandShareRequest) (*model.ExpandShareResponse, error) {
	requestDef := GenReqDefForExpandShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandShareResponse), nil
	}
}

// ExpandShareInvoker 扩容文件系统
func (c *SFSTurboClient) ExpandShareInvoker(request *model.ExpandShareRequest) *ExpandShareInvoker {
	requestDef := GenReqDefForExpandShare()
	return &ExpandShareInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSharedTags 查询租户所有共享的标签
//
// 查询租户所有共享的标签集合。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SFSTurboClient) ListSharedTags(request *model.ListSharedTagsRequest) (*model.ListSharedTagsResponse, error) {
	requestDef := GenReqDefForListSharedTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSharedTagsResponse), nil
	}
}

// ListSharedTagsInvoker 查询租户所有共享的标签
func (c *SFSTurboClient) ListSharedTagsInvoker(request *model.ListSharedTagsRequest) *ListSharedTagsInvoker {
	requestDef := GenReqDefForListSharedTags()
	return &ListSharedTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListShares 获取文件系统列表
//
// 获取文件系统列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SFSTurboClient) ListShares(request *model.ListSharesRequest) (*model.ListSharesResponse, error) {
	requestDef := GenReqDefForListShares()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSharesResponse), nil
	}
}

// ListSharesInvoker 获取文件系统列表
func (c *SFSTurboClient) ListSharesInvoker(request *model.ListSharesRequest) *ListSharesInvoker {
	requestDef := GenReqDefForListShares()
	return &ListSharesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowShare 查询文件系统详细信息
//
// 查询SFS Turbo文件系统详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SFSTurboClient) ShowShare(request *model.ShowShareRequest) (*model.ShowShareResponse, error) {
	requestDef := GenReqDefForShowShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowShareResponse), nil
	}
}

// ShowShareInvoker 查询文件系统详细信息
func (c *SFSTurboClient) ShowShareInvoker(request *model.ShowShareRequest) *ShowShareInvoker {
	requestDef := GenReqDefForShowShare()
	return &ShowShareInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSharedTags 查询共享标签
//
// 查询指定共享的所有标签信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *SFSTurboClient) ShowSharedTags(request *model.ShowSharedTagsRequest) (*model.ShowSharedTagsResponse, error) {
	requestDef := GenReqDefForShowSharedTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSharedTagsResponse), nil
	}
}

// ShowSharedTagsInvoker 查询共享标签
func (c *SFSTurboClient) ShowSharedTagsInvoker(request *model.ShowSharedTagsRequest) *ShowSharedTagsInvoker {
	requestDef := GenReqDefForShowSharedTags()
	return &ShowSharedTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
