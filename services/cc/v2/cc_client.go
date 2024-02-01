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

// BatchCreateGcbResourceTags 批量添加账户全域互联带宽资源标签
//
// # TMS批量添加资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) BatchCreateGcbResourceTags(request *model.BatchCreateGcbResourceTagsRequest) (*model.BatchCreateGcbResourceTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateGcbResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateGcbResourceTagsResponse), nil
	}
}

// BatchCreateGcbResourceTagsInvoker 批量添加账户全域互联带宽资源标签
func (c *CcClient) BatchCreateGcbResourceTagsInvoker(request *model.BatchCreateGcbResourceTagsRequest) *BatchCreateGcbResourceTagsInvoker {
	requestDef := GenReqDefForBatchCreateGcbResourceTags()
	return &BatchCreateGcbResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteGcbResourceTags 批量删除账户全域互联带宽资源标签
//
// 批量删除账户全域互联带宽资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) BatchDeleteGcbResourceTags(request *model.BatchDeleteGcbResourceTagsRequest) (*model.BatchDeleteGcbResourceTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteGcbResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteGcbResourceTagsResponse), nil
	}
}

// BatchDeleteGcbResourceTagsInvoker 批量删除账户全域互联带宽资源标签
func (c *CcClient) BatchDeleteGcbResourceTagsInvoker(request *model.BatchDeleteGcbResourceTagsRequest) *BatchDeleteGcbResourceTagsInvoker {
	requestDef := GenReqDefForBatchDeleteGcbResourceTags()
	return &BatchDeleteGcbResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountGcbResourceByTag 查询账户全域互联带宽资源标签数量
//
// 查询账户全域互联带宽资源标签数量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) CountGcbResourceByTag(request *model.CountGcbResourceByTagRequest) (*model.CountGcbResourceByTagResponse, error) {
	requestDef := GenReqDefForCountGcbResourceByTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountGcbResourceByTagResponse), nil
	}
}

// CountGcbResourceByTagInvoker 查询账户全域互联带宽资源标签数量
func (c *CcClient) CountGcbResourceByTagInvoker(request *model.CountGcbResourceByTagRequest) *CountGcbResourceByTagInvoker {
	requestDef := GenReqDefForCountGcbResourceByTag()
	return &CountGcbResourceByTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGcbResourceTag 添加账户全域互联带宽资源标签
//
// 添加账户全域互联带宽资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) CreateGcbResourceTag(request *model.CreateGcbResourceTagRequest) (*model.CreateGcbResourceTagResponse, error) {
	requestDef := GenReqDefForCreateGcbResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGcbResourceTagResponse), nil
	}
}

// CreateGcbResourceTagInvoker 添加账户全域互联带宽资源标签
func (c *CcClient) CreateGcbResourceTagInvoker(request *model.CreateGcbResourceTagRequest) *CreateGcbResourceTagInvoker {
	requestDef := GenReqDefForCreateGcbResourceTag()
	return &CreateGcbResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGcbResourceTag 删除账户全域互联带宽资源标签
//
// 删除账户全域互联带宽资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) DeleteGcbResourceTag(request *model.DeleteGcbResourceTagRequest) (*model.DeleteGcbResourceTagResponse, error) {
	requestDef := GenReqDefForDeleteGcbResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGcbResourceTagResponse), nil
	}
}

// DeleteGcbResourceTagInvoker 删除账户全域互联带宽资源标签
func (c *CcClient) DeleteGcbResourceTagInvoker(request *model.DeleteGcbResourceTagRequest) *DeleteGcbResourceTagInvoker {
	requestDef := GenReqDefForDeleteGcbResourceTag()
	return &DeleteGcbResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGcbResourceByTag 查询账户全域互联带宽资源实例列表
//
// 查询账户全域互联带宽资源实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListGcbResourceByTag(request *model.ListGcbResourceByTagRequest) (*model.ListGcbResourceByTagResponse, error) {
	requestDef := GenReqDefForListGcbResourceByTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGcbResourceByTagResponse), nil
	}
}

// ListGcbResourceByTagInvoker 查询账户全域互联带宽资源实例列表
func (c *CcClient) ListGcbResourceByTagInvoker(request *model.ListGcbResourceByTagRequest) *ListGcbResourceByTagInvoker {
	requestDef := GenReqDefForListGcbResourceByTag()
	return &ListGcbResourceByTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGcbResourceTags 查询账户全域互联带宽资源的标签
//
// 查询账户全域互联带宽资源的标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListGcbResourceTags(request *model.ListGcbResourceTagsRequest) (*model.ListGcbResourceTagsResponse, error) {
	requestDef := GenReqDefForListGcbResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGcbResourceTagsResponse), nil
	}
}

// ListGcbResourceTagsInvoker 查询账户全域互联带宽资源的标签
func (c *CcClient) ListGcbResourceTagsInvoker(request *model.ListGcbResourceTagsRequest) *ListGcbResourceTagsInvoker {
	requestDef := GenReqDefForListGcbResourceTags()
	return &ListGcbResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGcbTenantTags 查询账户全域互联带宽所有资源标签
//
// 查询账户全域互联带宽所有资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListGcbTenantTags(request *model.ListGcbTenantTagsRequest) (*model.ListGcbTenantTagsResponse, error) {
	requestDef := GenReqDefForListGcbTenantTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGcbTenantTagsResponse), nil
	}
}

// ListGcbTenantTagsInvoker 查询账户全域互联带宽所有资源标签
func (c *CcClient) ListGcbTenantTagsInvoker(request *model.ListGcbTenantTagsRequest) *ListGcbTenantTagsInvoker {
	requestDef := GenReqDefForListGcbTenantTags()
	return &ListGcbTenantTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateDeleteTags 批量创建和删除资源标签
//
// 批量创建和删除标签
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
// 添加资源标签
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
// 删除资源标签
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
// 查询账户资源标签
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
// 查询资源实例
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
// 查询资源标签
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
