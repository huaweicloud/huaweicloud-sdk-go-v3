package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ies/v1/model"
)

type IesClient struct {
	HcClient *http_client.HcHttpClient
}

func NewIesClient(hcClient *http_client.HcHttpClient) *IesClient {
	return &IesClient{HcClient: hcClient}
}

func IesClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// 创建边缘小站
//
// 创建边缘小站。
// - 一个边缘小站关联一个华为云指定的区域。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IesClient) CreateEdgeSite(request *model.CreateEdgeSiteRequest) (*model.CreateEdgeSiteResponse, error) {
	requestDef := GenReqDefForCreateEdgeSite()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEdgeSiteResponse), nil
	}
}

// 删除边缘小站
//
// 删除指定的边缘小站。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IesClient) DeleteEdgeSite(request *model.DeleteEdgeSiteRequest) (*model.DeleteEdgeSiteResponse, error) {
	requestDef := GenReqDefForDeleteEdgeSite()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEdgeSiteResponse), nil
	}
}

// 查询边缘小站列表
//
// 查询边缘小站列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IesClient) ListEdgeSites(request *model.ListEdgeSitesRequest) (*model.ListEdgeSitesResponse, error) {
	requestDef := GenReqDefForListEdgeSites()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEdgeSitesResponse), nil
	}
}

// 查询边缘小站详情
//
// 查询边缘小站详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IesClient) ShowEdgeSite(request *model.ShowEdgeSiteRequest) (*model.ShowEdgeSiteResponse, error) {
	requestDef := GenReqDefForShowEdgeSite()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEdgeSiteResponse), nil
	}
}

// 更新边缘小站
//
// 更新边缘小站。
// - 允许更新边缘小站描述或场地信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IesClient) UpdateEdgeSite(request *model.UpdateEdgeSiteRequest) (*model.UpdateEdgeSiteResponse, error) {
	requestDef := GenReqDefForUpdateEdgeSite()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEdgeSiteResponse), nil
	}
}

// 查看站点容量信息
//
// 查看站点容量信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IesClient) ListEdgeSiteMetrics(request *model.ListEdgeSiteMetricsRequest) (*model.ListEdgeSiteMetricsResponse, error) {
	requestDef := GenReqDefForListEdgeSiteMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEdgeSiteMetricsResponse), nil
	}
}

// 查询配额
//
// 查询租户资源配额。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IesClient) ListQuotas(request *model.ListQuotasRequest) (*model.ListQuotasResponse, error) {
	requestDef := GenReqDefForListQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotasResponse), nil
	}
}

// 查询支持的区域列表
//
// 查询支持智能边缘小站接入的华为云区域（region）列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *IesClient) ListSupportedRegions(request *model.ListSupportedRegionsRequest) (*model.ListSupportedRegionsResponse, error) {
	requestDef := GenReqDefForListSupportedRegions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSupportedRegionsResponse), nil
	}
}
