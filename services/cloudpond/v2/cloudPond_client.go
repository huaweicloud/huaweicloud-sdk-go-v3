package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cloudpond/v2/model"
)

type CloudPondClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCloudPondClient(hcClient *httpclient.HcHttpClient) *CloudPondClient {
	return &CloudPondClient{HcClient: hcClient}
}

func CloudPondClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// ListNetworkDevices 查询网络设备列表
//
// 查询网络设备列表。
// [- 该接口支持企业项目细粒度权限的校验，具体细粒度请参见 ies:edgeSite:listNetworkDevices](tag:hws)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudPondClient) ListNetworkDevices(request *model.ListNetworkDevicesRequest) (*model.ListNetworkDevicesResponse, error) {
	requestDef := GenReqDefForListNetworkDevices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNetworkDevicesResponse), nil
	}
}

// ListNetworkDevicesInvoker 查询网络设备列表
func (c *CloudPondClient) ListNetworkDevicesInvoker(request *model.ListNetworkDevicesRequest) *ListNetworkDevicesInvoker {
	requestDef := GenReqDefForListNetworkDevices()
	return &ListNetworkDevicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNetworkDevice 查询网络设备详情
//
// 查询网络设备详情。
// [- 该接口支持企业项目细粒度权限的校验，具体细粒度请参见 ies:edgeSite:getNetworkDevice](tag:hws)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudPondClient) ShowNetworkDevice(request *model.ShowNetworkDeviceRequest) (*model.ShowNetworkDeviceResponse, error) {
	requestDef := GenReqDefForShowNetworkDevice()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNetworkDeviceResponse), nil
	}
}

// ShowNetworkDeviceInvoker 查询网络设备详情
func (c *CloudPondClient) ShowNetworkDeviceInvoker(request *model.ShowNetworkDeviceRequest) *ShowNetworkDeviceInvoker {
	requestDef := GenReqDefForShowNetworkDevice()
	return &ShowNetworkDeviceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNetworkDeviceOfferings 查询网络设备商品列表
//
// 查询网络设备商品列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudPondClient) ListNetworkDeviceOfferings(request *model.ListNetworkDeviceOfferingsRequest) (*model.ListNetworkDeviceOfferingsResponse, error) {
	requestDef := GenReqDefForListNetworkDeviceOfferings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNetworkDeviceOfferingsResponse), nil
	}
}

// ListNetworkDeviceOfferingsInvoker 查询网络设备商品列表
func (c *CloudPondClient) ListNetworkDeviceOfferingsInvoker(request *model.ListNetworkDeviceOfferingsRequest) *ListNetworkDeviceOfferingsInvoker {
	requestDef := GenReqDefForListNetworkDeviceOfferings()
	return &ListNetworkDeviceOfferingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServerOfferings 查询服务器商品列表
//
// 查询服务器销售商品列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudPondClient) ListServerOfferings(request *model.ListServerOfferingsRequest) (*model.ListServerOfferingsResponse, error) {
	requestDef := GenReqDefForListServerOfferings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServerOfferingsResponse), nil
	}
}

// ListServerOfferingsInvoker 查询服务器商品列表
func (c *CloudPondClient) ListServerOfferingsInvoker(request *model.ListServerOfferingsRequest) *ListServerOfferingsInvoker {
	requestDef := GenReqDefForListServerOfferings()
	return &ListServerOfferingsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSaleCycles 查询可购买的销售周期
//
// 查询可购买的销售周期
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudPondClient) ListSaleCycles(request *model.ListSaleCyclesRequest) (*model.ListSaleCyclesResponse, error) {
	requestDef := GenReqDefForListSaleCycles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSaleCyclesResponse), nil
	}
}

// ListSaleCyclesInvoker 查询可购买的销售周期
func (c *CloudPondClient) ListSaleCyclesInvoker(request *model.ListSaleCyclesRequest) *ListSaleCyclesInvoker {
	requestDef := GenReqDefForListSaleCycles()
	return &ListSaleCyclesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServers 查询服务器列表
//
// 查询服务器列表。
// [- 该接口支持企业项目细粒度权限的校验，具体细粒度请参见 ies:edgeSite:listServers](tag:hws)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudPondClient) ListServers(request *model.ListServersRequest) (*model.ListServersResponse, error) {
	requestDef := GenReqDefForListServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServersResponse), nil
	}
}

// ListServersInvoker 查询服务器列表
func (c *CloudPondClient) ListServersInvoker(request *model.ListServersRequest) *ListServersInvoker {
	requestDef := GenReqDefForListServers()
	return &ListServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowServer 查询服务器详情
//
// 查询服务器详情。
// [- 该接口支持企业项目细粒度权限的校验，具体细粒度请参见 ies:edgeSite:getServer](tag:hws)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudPondClient) ShowServer(request *model.ShowServerRequest) (*model.ShowServerResponse, error) {
	requestDef := GenReqDefForShowServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerResponse), nil
	}
}

// ShowServerInvoker 查询服务器详情
func (c *CloudPondClient) ShowServerInvoker(request *model.ShowServerRequest) *ShowServerInvoker {
	requestDef := GenReqDefForShowServer()
	return &ShowServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStorageGears 查询存储计费档位
//
// 该接口仅返回支持的存储计费档位，实际可购买的存储容量单独定义。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudPondClient) ListStorageGears(request *model.ListStorageGearsRequest) (*model.ListStorageGearsResponse, error) {
	requestDef := GenReqDefForListStorageGears()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStorageGearsResponse), nil
	}
}

// ListStorageGearsInvoker 查询存储计费档位
func (c *CloudPondClient) ListStorageGearsInvoker(request *model.ListStorageGearsRequest) *ListStorageGearsInvoker {
	requestDef := GenReqDefForListStorageGears()
	return &ListStorageGearsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStoragePools 查询存储池列表
//
// 查询存储池列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudPondClient) ListStoragePools(request *model.ListStoragePoolsRequest) (*model.ListStoragePoolsResponse, error) {
	requestDef := GenReqDefForListStoragePools()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStoragePoolsResponse), nil
	}
}

// ListStoragePoolsInvoker 查询存储池列表
func (c *CloudPondClient) ListStoragePoolsInvoker(request *model.ListStoragePoolsRequest) *ListStoragePoolsInvoker {
	requestDef := GenReqDefForListStoragePools()
	return &ListStoragePoolsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStoragePool 查询存储池详情
//
// 查询存储池详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudPondClient) ShowStoragePool(request *model.ShowStoragePoolRequest) (*model.ShowStoragePoolResponse, error) {
	requestDef := GenReqDefForShowStoragePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStoragePoolResponse), nil
	}
}

// ShowStoragePoolInvoker 查询存储池详情
func (c *CloudPondClient) ShowStoragePoolInvoker(request *model.ShowStoragePoolRequest) *ShowStoragePoolInvoker {
	requestDef := GenReqDefForShowStoragePool()
	return &ShowStoragePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStorageTypes 查询存储类型列表
//
// 查询支持的存储类型列表，包括步长等信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudPondClient) ListStorageTypes(request *model.ListStorageTypesRequest) (*model.ListStorageTypesResponse, error) {
	requestDef := GenReqDefForListStorageTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStorageTypesResponse), nil
	}
}

// ListStorageTypesInvoker 查询存储类型列表
func (c *CloudPondClient) ListStorageTypesInvoker(request *model.ListStorageTypesRequest) *ListStorageTypesInvoker {
	requestDef := GenReqDefForListStorageTypes()
	return &ListStorageTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
