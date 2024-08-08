package v5

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iotdm/v5/model"
)

type IoTDMClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewIoTDMClient(hcClient *httpclient.HcHttpClient) *IoTDMClient {
	return &IoTDMClient{HcClient: hcClient}
}

func IoTDMClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// BindInstanceTags 添加实例标签
//
// 添加实例标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTDMClient) BindInstanceTags(request *model.BindInstanceTagsRequest) (*model.BindInstanceTagsResponse, error) {
	requestDef := GenReqDefForBindInstanceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BindInstanceTagsResponse), nil
	}
}

// BindInstanceTagsInvoker 添加实例标签
func (c *IoTDMClient) BindInstanceTagsInvoker(request *model.BindInstanceTagsRequest) *BindInstanceTagsInvoker {
	requestDef := GenReqDefForBindInstanceTags()
	return &BindInstanceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeInstanceChargeMode 修改实例计费模式
//
// 修改设备接入实例的计费模式，支持将按需计费模式修改为包年/包月计费模式。
// 接口约束：当前实例的规格支持包年/包月计费模式时，才可以修改实例的计费模式为包年/包月。支持的实例规格请参见[[产品规格说明](https://support.huaweicloud.com/productdesc-iothub/iot_04_0014.html)](tag:hws)[[产品规格说明](https://support.huaweicloud.com/intl/zh-cn/productdesc-iothub/iot_04_0014.html)](tag:hws_hk)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTDMClient) ChangeInstanceChargeMode(request *model.ChangeInstanceChargeModeRequest) (*model.ChangeInstanceChargeModeResponse, error) {
	requestDef := GenReqDefForChangeInstanceChargeMode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeInstanceChargeModeResponse), nil
	}
}

// ChangeInstanceChargeModeInvoker 修改实例计费模式
func (c *IoTDMClient) ChangeInstanceChargeModeInvoker(request *model.ChangeInstanceChargeModeRequest) *ChangeInstanceChargeModeInvoker {
	requestDef := GenReqDefForChangeInstanceChargeMode()
	return &ChangeInstanceChargeModeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstance 创建设备接入实例
//
// 用户可以调用此接口创建一个设备接入实例。支持的实例规格请参见[[产品规格说明](https://support.huaweicloud.com/productdesc-iothub/iot_04_0014.html)](tag:hws)[[产品规格说明](https://support.huaweicloud.com/intl/zh-cn/productdesc-iothub/iot_04_0014.html)](tag:hws_hk)。
// [接口约束：
// - 请保证账户余额充足，此接口无法使用优惠券支付，在创建包年/包月实例时，若余额不足会创建一个待支付订单。
// - 若想使用优惠券，请将请求中的is_auto_pay字段设置为false，参考[\&quot;支付包年/包月产品订单\&quot;](https://support.huaweicloud.com/api-bpconsole/api_order_00016.html#section0)进行支付，或者在华为云官网页面使用优惠券进行支付。
// - 如果您需要退订包年/包月资源，请参考[\&quot;退订包年/包月资源\&quot;](https://support.huaweicloud.com/api-bpconsole/api_order_00019.html)。](tag:hws)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTDMClient) CreateInstance(request *model.CreateInstanceRequest) (*model.CreateInstanceResponse, error) {
	requestDef := GenReqDefForCreateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceResponse), nil
	}
}

// CreateInstanceInvoker 创建设备接入实例
func (c *IoTDMClient) CreateInstanceInvoker(request *model.CreateInstanceRequest) *CreateInstanceInvoker {
	requestDef := GenReqDefForCreateInstance()
	return &CreateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstance 删除实例
//
// 删除设备接入实例。约束：此接口仅支持删除按需计费的实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTDMClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResponse), nil
	}
}

// DeleteInstanceInvoker 删除实例
func (c *IoTDMClient) DeleteInstanceInvoker(request *model.DeleteInstanceRequest) *DeleteInstanceInvoker {
	requestDef := GenReqDefForDeleteInstance()
	return &DeleteInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstances 查询实例列表
//
// 用户可以调用此接口查询设备接入实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTDMClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

// ListInstancesInvoker 查询实例列表
func (c *IoTDMClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeInstance 修改实例规格信息
//
// 修改设备接入实例的规格。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTDMClient) ResizeInstance(request *model.ResizeInstanceRequest) (*model.ResizeInstanceResponse, error) {
	requestDef := GenReqDefForResizeInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeInstanceResponse), nil
	}
}

// ResizeInstanceInvoker 修改实例规格信息
func (c *IoTDMClient) ResizeInstanceInvoker(request *model.ResizeInstanceRequest) *ResizeInstanceInvoker {
	requestDef := GenReqDefForResizeInstance()
	return &ResizeInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstance 查询实例详情
//
// 查询设备接入实例详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTDMClient) ShowInstance(request *model.ShowInstanceRequest) (*model.ShowInstanceResponse, error) {
	requestDef := GenReqDefForShowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceResponse), nil
	}
}

// ShowInstanceInvoker 查询实例详情
func (c *IoTDMClient) ShowInstanceInvoker(request *model.ShowInstanceRequest) *ShowInstanceInvoker {
	requestDef := GenReqDefForShowInstance()
	return &ShowInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnbindInstanceTags 删除实例标签
//
// 删除实例标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTDMClient) UnbindInstanceTags(request *model.UnbindInstanceTagsRequest) (*model.UnbindInstanceTagsResponse, error) {
	requestDef := GenReqDefForUnbindInstanceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnbindInstanceTagsResponse), nil
	}
}

// UnbindInstanceTagsInvoker 删除实例标签
func (c *IoTDMClient) UnbindInstanceTagsInvoker(request *model.UnbindInstanceTagsRequest) *UnbindInstanceTagsInvoker {
	requestDef := GenReqDefForUnbindInstanceTags()
	return &UnbindInstanceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstance 修改实例信息
//
// 修改设备接入实例信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IoTDMClient) UpdateInstance(request *model.UpdateInstanceRequest) (*model.UpdateInstanceResponse, error) {
	requestDef := GenReqDefForUpdateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceResponse), nil
	}
}

// UpdateInstanceInvoker 修改实例信息
func (c *IoTDMClient) UpdateInstanceInvoker(request *model.UpdateInstanceRequest) *UpdateInstanceInvoker {
	requestDef := GenReqDefForUpdateInstance()
	return &UpdateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
