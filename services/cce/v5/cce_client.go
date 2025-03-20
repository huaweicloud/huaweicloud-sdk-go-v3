package v5

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cce/v5/model"
)

type CceClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCceClient(hcClient *httpclient.HcHttpClient) *CceClient {
	return &CceClient{HcClient: hcClient}
}

func CceClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateImageCache 创建镜像缓存
//
// 创建镜像缓存。创建过程会在指定集群中启动临时Pod进行镜像缓存构建，创建镜像缓存后，可在负载创建时通过使用镜像缓存功能大幅减少下载容器镜像的耗时，实现容器的快速启动。单租户创建镜像缓存数量上限为50。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) CreateImageCache(request *model.CreateImageCacheRequest) (*model.CreateImageCacheResponse, error) {
	requestDef := GenReqDefForCreateImageCache()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateImageCacheResponse), nil
	}
}

// CreateImageCacheInvoker 创建镜像缓存
func (c *CceClient) CreateImageCacheInvoker(request *model.CreateImageCacheRequest) *CreateImageCacheInvoker {
	requestDef := GenReqDefForCreateImageCache()
	return &CreateImageCacheInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteImageCache 删除镜像缓存
//
// 删除镜像缓存
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) DeleteImageCache(request *model.DeleteImageCacheRequest) (*model.DeleteImageCacheResponse, error) {
	requestDef := GenReqDefForDeleteImageCache()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteImageCacheResponse), nil
	}
}

// DeleteImageCacheInvoker 删除镜像缓存
func (c *CceClient) DeleteImageCacheInvoker(request *model.DeleteImageCacheRequest) *DeleteImageCacheInvoker {
	requestDef := GenReqDefForDeleteImageCache()
	return &DeleteImageCacheInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImageCaches 查询镜像缓存列表
//
// 查询镜像缓存列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) ListImageCaches(request *model.ListImageCachesRequest) (*model.ListImageCachesResponse, error) {
	requestDef := GenReqDefForListImageCaches()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImageCachesResponse), nil
	}
}

// ListImageCachesInvoker 查询镜像缓存列表
func (c *CceClient) ListImageCachesInvoker(request *model.ListImageCachesRequest) *ListImageCachesInvoker {
	requestDef := GenReqDefForListImageCaches()
	return &ListImageCachesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowImageCache 查询镜像缓存详情
//
// 查询镜像缓存详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CceClient) ShowImageCache(request *model.ShowImageCacheRequest) (*model.ShowImageCacheResponse, error) {
	requestDef := GenReqDefForShowImageCache()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowImageCacheResponse), nil
	}
}

// ShowImageCacheInvoker 查询镜像缓存详情
func (c *CceClient) ShowImageCacheInvoker(request *model.ShowImageCacheRequest) *ShowImageCacheInvoker {
	requestDef := GenReqDefForShowImageCache()
	return &ShowImageCacheInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
