package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eip/v3/model"
)

type EipClient struct {
	HcClient *http_client.HcHttpClient
}

func NewEipClient(hcClient *http_client.HcHttpClient) *EipClient {
	return &EipClient{HcClient: hcClient}
}

func EipClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// 查询公共池列表
//
// 查询公共池列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EipClient) ListCommonPools(request *model.ListCommonPoolsRequest) (*model.ListCommonPoolsResponse, error) {
	requestDef := GenReqDefForListCommonPools()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCommonPoolsResponse), nil
	}
}

// 查询公共池分组列表
//
// 查询公共池分组列表，包含名称和位置信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EipClient) ListPublicBorderGroups(request *model.ListPublicBorderGroupsRequest) (*model.ListPublicBorderGroupsResponse, error) {
	requestDef := GenReqDefForListPublicBorderGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPublicBorderGroupsResponse), nil
	}
}

// 查询指定租户下的共享带宽类型列表
//
// 查询指定租户下的共享带宽类型列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EipClient) ListShareBandwidthTypes(request *model.ListShareBandwidthTypesRequest) (*model.ListShareBandwidthTypesResponse, error) {
	requestDef := GenReqDefForListShareBandwidthTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListShareBandwidthTypesResponse), nil
	}
}

// 绑定弹性公网IP
//
// 绑定弹性公网IP
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EipClient) AssociatePublicips(request *model.AssociatePublicipsRequest) (*model.AssociatePublicipsResponse, error) {
	requestDef := GenReqDefForAssociatePublicips()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociatePublicipsResponse), nil
	}
}

// 解绑弹性公网IP
//
// 解绑弹性公网IP
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EipClient) DisassociatePublicips(request *model.DisassociatePublicipsRequest) (*model.DisassociatePublicipsResponse, error) {
	requestDef := GenReqDefForDisassociatePublicips()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociatePublicipsResponse), nil
	}
}

// 查询公网IP池列表
//
// 全量查询公网IP池列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EipClient) ListPublicipPool(request *model.ListPublicipPoolRequest) (*model.ListPublicipPoolResponse, error) {
	requestDef := GenReqDefForListPublicipPool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPublicipPoolResponse), nil
	}
}

// 全量查询弹性公网IP列表
//
// 查询弹性公网IP列表信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EipClient) ListPublicips(request *model.ListPublicipsRequest) (*model.ListPublicipsResponse, error) {
	requestDef := GenReqDefForListPublicips()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPublicipsResponse), nil
	}
}

// 查询弹性公网IP详情
//
// 查询弹性公网IP详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EipClient) ShowPublicip(request *model.ShowPublicipRequest) (*model.ShowPublicipResponse, error) {
	requestDef := GenReqDefForShowPublicip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPublicipResponse), nil
	}
}

// 查询公网IP池详情
//
// 查询公网IP池详情
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *EipClient) ShowPublicipPool(request *model.ShowPublicipPoolRequest) (*model.ShowPublicipPoolResponse, error) {
	requestDef := GenReqDefForShowPublicipPool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPublicipPoolResponse), nil
	}
}
