package v1

import (
	http_client "github.com/RandolphCYG/huaweicloud-sdk-go-v3/core"

	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/services/sfsturbo/v1/model"
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

//指定共享批量添加标签。  一个共享上最多有10个标签。 一个共享上的多个标签的key不允许重复。 此接口为幂等接口：如果要添加的key在共享上已存在，则覆盖更新标签。
func (c *SFSTurboClient) BatchAddSharedTags(request *model.BatchAddSharedTagsRequest) (*model.BatchAddSharedTagsResponse, error) {
	requestDef := GenReqDefForBatchAddSharedTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddSharedTagsResponse), nil
	}
}

//修改SFS Turbo文件系统绑定的安全组。修改安全组为异步任务，可以通过“查询单个文件系统”返回的子状态字段“sub_status”来判断是否修改安全组状态，子状态为“232”即为修改安全组成功。
func (c *SFSTurboClient) ChangeSecurityGroup(request *model.ChangeSecurityGroupRequest) (*model.ChangeSecurityGroupResponse, error) {
	requestDef := GenReqDefForChangeSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeSecurityGroupResponse), nil
	}
}

//创建文件系统。
func (c *SFSTurboClient) CreateShare(request *model.CreateShareRequest) (*model.CreateShareResponse, error) {
	requestDef := GenReqDefForCreateShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateShareResponse), nil
	}
}

//指定共享添加一个标签。 一个共享上最多有10个标签。 一个共享上的多个标签的key不允许重复。 此接口为幂等接口：如果要添加的key在共享上已存在，则覆盖更新标签。
func (c *SFSTurboClient) CreateSharedTag(request *model.CreateSharedTagRequest) (*model.CreateSharedTagResponse, error) {
	requestDef := GenReqDefForCreateSharedTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSharedTagResponse), nil
	}
}

//删除文件系统。
func (c *SFSTurboClient) DeleteShare(request *model.DeleteShareRequest) (*model.DeleteShareResponse, error) {
	requestDef := GenReqDefForDeleteShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteShareResponse), nil
	}
}

//指定共享删除一个标签。当共享中不存在指定要删除的key时，接口调用将会返回404错误。
func (c *SFSTurboClient) DeleteSharedTag(request *model.DeleteSharedTagRequest) (*model.DeleteSharedTagResponse, error) {
	requestDef := GenReqDefForDeleteSharedTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSharedTagResponse), nil
	}
}

//扩容文件系统。
func (c *SFSTurboClient) ExpandShare(request *model.ExpandShareRequest) (*model.ExpandShareResponse, error) {
	requestDef := GenReqDefForExpandShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandShareResponse), nil
	}
}

//查询租户所有共享的标签集合。
func (c *SFSTurboClient) ListSharedTags(request *model.ListSharedTagsRequest) (*model.ListSharedTagsResponse, error) {
	requestDef := GenReqDefForListSharedTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSharedTagsResponse), nil
	}
}

//获取文件系统列表
func (c *SFSTurboClient) ListShares(request *model.ListSharesRequest) (*model.ListSharesResponse, error) {
	requestDef := GenReqDefForListShares()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSharesResponse), nil
	}
}

//查询SFS Turbo文件系统详细信息。
func (c *SFSTurboClient) ShowShare(request *model.ShowShareRequest) (*model.ShowShareResponse, error) {
	requestDef := GenReqDefForShowShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowShareResponse), nil
	}
}

//查询指定共享的所有标签信息。
func (c *SFSTurboClient) ShowSharedTags(request *model.ShowSharedTagsRequest) (*model.ShowSharedTagsResponse, error) {
	requestDef := GenReqDefForShowSharedTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSharedTagsResponse), nil
	}
}
