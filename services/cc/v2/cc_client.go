package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cc/v2/model"
)

type CcClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCcClient(hcClient *httpclient.HcHttpClient) *CcClient {
	return &CcClient{HcClient: hcClient}
}

func CcClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// BatchCreateDeleteTags 批量创建和删除资源标签
//
// 批量创建和删除标签。此API为历史API，请优先使用《 创建云连接实例标签》、《 创建带宽包标签》、《 删除云连接实例标签》、《 删除带宽包标签》。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) BatchCreateDeleteTags(request *model.BatchCreateDeleteTagsRequest) (*model.BatchCreateDeleteTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateDeleteTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateDeleteTagsResponse), nil
	}
}

// BatchCreateDeleteTagsInvoker 批量创建和删除资源标签
func (c *CcClient) BatchCreateDeleteTagsInvoker(request *model.BatchCreateDeleteTagsRequest) *BatchCreateDeleteTagsInvoker {
	requestDef := GenReqDefForBatchCreateDeleteTags()
	return &BatchCreateDeleteTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTag 添加资源标签
//
// 添加资源标签。此API为历史API，请优先使用《 创建云连接实例标签》、《 创建带宽包标签》。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) CreateTag(request *model.CreateTagRequest) (*model.CreateTagResponse, error) {
	requestDef := GenReqDefForCreateTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTagResponse), nil
	}
}

// CreateTagInvoker 添加资源标签
func (c *CcClient) CreateTagInvoker(request *model.CreateTagRequest) *CreateTagInvoker {
	requestDef := GenReqDefForCreateTag()
	return &CreateTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTag 删除资源标签
//
// 删除资源标签。此API为历史API，请优先使用《 删除云连接实例标签》或《 删除带宽包标签》。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) DeleteTag(request *model.DeleteTagRequest) (*model.DeleteTagResponse, error) {
	requestDef := GenReqDefForDeleteTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTagResponse), nil
	}
}

// DeleteTagInvoker 删除资源标签
func (c *CcClient) DeleteTagInvoker(request *model.DeleteTagRequest) *DeleteTagInvoker {
	requestDef := GenReqDefForDeleteTag()
	return &DeleteTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDomainTags 查询账户资源标签
//
// 查询账户资源标签。此API为历史API，请优先使用《查询云连接实例的标签信息》、《查询带宽包的标签信息》。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListDomainTags(request *model.ListDomainTagsRequest) (*model.ListDomainTagsResponse, error) {
	requestDef := GenReqDefForListDomainTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainTagsResponse), nil
	}
}

// ListDomainTagsInvoker 查询账户资源标签
func (c *CcClient) ListDomainTagsInvoker(request *model.ListDomainTagsRequest) *ListDomainTagsInvoker {
	requestDef := GenReqDefForListDomainTags()
	return &ListDomainTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceByFilterTag 查询资源实例
//
// 查询资源实例。此API为历史API，请优先使用《通过标签过滤云连接实例》、《通过标签过滤带宽包实例》。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListResourceByFilterTag(request *model.ListResourceByFilterTagRequest) (*model.ListResourceByFilterTagResponse, error) {
	requestDef := GenReqDefForListResourceByFilterTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceByFilterTagResponse), nil
	}
}

// ListResourceByFilterTagInvoker 查询资源实例
func (c *CcClient) ListResourceByFilterTagInvoker(request *model.ListResourceByFilterTagRequest) *ListResourceByFilterTagInvoker {
	requestDef := GenReqDefForListResourceByFilterTag()
	return &ListResourceByFilterTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTags 查询资源标签
//
// 查询资源标签。此API为历史API，请优先使用《查询云连接实例的标签信息》、《查询带宽包的标签信息》。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListTags(request *model.ListTagsRequest) (*model.ListTagsResponse, error) {
	requestDef := GenReqDefForListTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsResponse), nil
	}
}

// ListTagsInvoker 查询资源标签
func (c *CcClient) ListTagsInvoker(request *model.ListTagsRequest) *ListTagsInvoker {
	requestDef := GenReqDefForListTags()
	return &ListTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
