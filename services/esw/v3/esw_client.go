package v3

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/esw/v3/model"
)

type EswClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewEswClient(hcClient *httpclient.HcHttpClient) *EswClient {
	return &EswClient{HcClient: hcClient}
}

func EswClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// BindVport 将一个虚拟IP绑定到二层连接上
//
// 当您的二层连接创建成功后，您可以通过调用此接口将一个虚拟IP绑定到二层连接上。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EswClient) BindVport(request *model.BindVportRequest) (*model.BindVportResponse, error) {
	requestDef := GenReqDefForBindVport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BindVportResponse), nil
	}
}

// BindVportInvoker 将一个虚拟IP绑定到二层连接上
func (c *EswClient) BindVportInvoker(request *model.BindVportRequest) *BindVportInvoker {
	requestDef := GenReqDefForBindVport()
	return &BindVportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConnection 创建二层连接
//
// 当您的ESW实例创建成功后，您可以通过调用此接口在该实例上创建一个二层连接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EswClient) CreateConnection(request *model.CreateConnectionRequest) (*model.CreateConnectionResponse, error) {
	requestDef := GenReqDefForCreateConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConnectionResponse), nil
	}
}

// CreateConnectionInvoker 创建二层连接
func (c *EswClient) CreateConnectionInvoker(request *model.CreateConnectionRequest) *CreateConnectionInvoker {
	requestDef := GenReqDefForCreateConnection()
	return &CreateConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstance 创建ESW实例
//
// 您可通过调用本接口创建一个ESW实例。该接口是一个异步接口，当前创建ESW实例的请求下发成功后，会响应job_id以及实例ID等信息，需要通过[调用查询任务的执行状态查询job状态](ListResourceJobs.xml)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EswClient) CreateInstance(request *model.CreateInstanceRequest) (*model.CreateInstanceResponse, error) {
	requestDef := GenReqDefForCreateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceResponse), nil
	}
}

// CreateInstanceInvoker 创建ESW实例
func (c *EswClient) CreateInstanceInvoker(request *model.CreateInstanceRequest) *CreateInstanceInvoker {
	requestDef := GenReqDefForCreateInstance()
	return &CreateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteConnection 删除二层连接
//
// 当您已创建的二层连接不再使用时，您可以通过调用该接口删除二层连接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EswClient) DeleteConnection(request *model.DeleteConnectionRequest) (*model.DeleteConnectionResponse, error) {
	requestDef := GenReqDefForDeleteConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConnectionResponse), nil
	}
}

// DeleteConnectionInvoker 删除二层连接
func (c *EswClient) DeleteConnectionInvoker(request *model.DeleteConnectionRequest) *DeleteConnectionInvoker {
	requestDef := GenReqDefForDeleteConnection()
	return &DeleteConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstance 删除ESW实例
//
// 当您创建的ESW实例不再使用时，您可以通过调用该接口删除ESW实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EswClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResponse), nil
	}
}

// DeleteInstanceInvoker 删除ESW实例
func (c *EswClient) DeleteInstanceInvoker(request *model.DeleteInstanceRequest) *DeleteInstanceInvoker {
	requestDef := GenReqDefForDeleteInstance()
	return &DeleteInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAvailabilityZones 查询ESW实例可用区
//
// 当您在创建ESW实例时，需要通过调用本接口获取ESW实例主备节点可分布的可用区列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EswClient) ListAvailabilityZones(request *model.ListAvailabilityZonesRequest) (*model.ListAvailabilityZonesResponse, error) {
	requestDef := GenReqDefForListAvailabilityZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailabilityZonesResponse), nil
	}
}

// ListAvailabilityZonesInvoker 查询ESW实例可用区
func (c *EswClient) ListAvailabilityZonesInvoker(request *model.ListAvailabilityZonesRequest) *ListAvailabilityZonesInvoker {
	requestDef := GenReqDefForListAvailabilityZones()
	return &ListAvailabilityZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConnections 查询实例下的二层连接列表
//
// 当您的二层连接创建成功后，您可以通过调用此接口查询ESW实例下的二层连接列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EswClient) ListConnections(request *model.ListConnectionsRequest) (*model.ListConnectionsResponse, error) {
	requestDef := GenReqDefForListConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConnectionsResponse), nil
	}
}

// ListConnectionsInvoker 查询实例下的二层连接列表
func (c *EswClient) ListConnectionsInvoker(request *model.ListConnectionsRequest) *ListConnectionsInvoker {
	requestDef := GenReqDefForListConnections()
	return &ListConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConnectionsAllInstances 查询二层连接列表
//
// 当您的二层连接创建成功后，您可以通过调用此接口查询租户所有二层连接信息列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EswClient) ListConnectionsAllInstances(request *model.ListConnectionsAllInstancesRequest) (*model.ListConnectionsAllInstancesResponse, error) {
	requestDef := GenReqDefForListConnectionsAllInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConnectionsAllInstancesResponse), nil
	}
}

// ListConnectionsAllInstancesInvoker 查询二层连接列表
func (c *EswClient) ListConnectionsAllInstancesInvoker(request *model.ListConnectionsAllInstancesRequest) *ListConnectionsAllInstancesInvoker {
	requestDef := GenReqDefForListConnectionsAllInstances()
	return &ListConnectionsAllInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlavors 查询ESW实例规格列表
//
// 查询租户可选用的ESW实例规格
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EswClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsResponse), nil
	}
}

// ListFlavorsInvoker 查询ESW实例规格列表
func (c *EswClient) ListFlavorsInvoker(request *model.ListFlavorsRequest) *ListFlavorsInvoker {
	requestDef := GenReqDefForListFlavors()
	return &ListFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstances 查询ESW实例列表
//
// 当您的ESW实例创建成功后，您可以通过调用此接口查询所有ESW实例的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EswClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

// ListInstancesInvoker 查询ESW实例列表
func (c *EswClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuotas 查询ESW实例配额
//
// 查询租户的ESW实例配额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EswClient) ListQuotas(request *model.ListQuotasRequest) (*model.ListQuotasResponse, error) {
	requestDef := GenReqDefForListQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotasResponse), nil
	}
}

// ListQuotasInvoker 查询ESW实例配额
func (c *EswClient) ListQuotasInvoker(request *model.ListQuotasRequest) *ListQuotasInvoker {
	requestDef := GenReqDefForListQuotas()
	return &ListQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceJobs 查询任务的执行状态
//
// 查询租户指定资源相关的任务信息列表，COMPLETED表示任务已成功完成，RUNNING表示任务正在执行中，FAILED表示任务执行失败。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EswClient) ListResourceJobs(request *model.ListResourceJobsRequest) (*model.ListResourceJobsResponse, error) {
	requestDef := GenReqDefForListResourceJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceJobsResponse), nil
	}
}

// ListResourceJobsInvoker 查询任务的执行状态
func (c *EswClient) ListResourceJobsInvoker(request *model.ListResourceJobsRequest) *ListResourceJobsInvoker {
	requestDef := GenReqDefForListResourceJobs()
	return &ListResourceJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConnection 查询二层连接详情
//
// 当您的二层连接创建成功后，您可以通过调用此接口查询单二层连接的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EswClient) ShowConnection(request *model.ShowConnectionRequest) (*model.ShowConnectionResponse, error) {
	requestDef := GenReqDefForShowConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConnectionResponse), nil
	}
}

// ShowConnectionInvoker 查询二层连接详情
func (c *EswClient) ShowConnectionInvoker(request *model.ShowConnectionRequest) *ShowConnectionInvoker {
	requestDef := GenReqDefForShowConnection()
	return &ShowConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstance 查询ESW实例详情
//
// 当您的ESW实例创建成功后，您可以通过调用此接口查询单个ESW实例的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EswClient) ShowInstance(request *model.ShowInstanceRequest) (*model.ShowInstanceResponse, error) {
	requestDef := GenReqDefForShowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceResponse), nil
	}
}

// ShowInstanceInvoker 查询ESW实例详情
func (c *EswClient) ShowInstanceInvoker(request *model.ShowInstanceRequest) *ShowInstanceInvoker {
	requestDef := GenReqDefForShowInstance()
	return &ShowInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnbindVport 将一个虚拟IP从二层连接解绑
//
// 当您的二层连接已经绑定虚拟IP时，您可以通过调用此接口将虚拟IP与二层连接解绑。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EswClient) UnbindVport(request *model.UnbindVportRequest) (*model.UnbindVportResponse, error) {
	requestDef := GenReqDefForUnbindVport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnbindVportResponse), nil
	}
}

// UnbindVportInvoker 将一个虚拟IP从二层连接解绑
func (c *EswClient) UnbindVportInvoker(request *model.UnbindVportRequest) *UnbindVportInvoker {
	requestDef := GenReqDefForUnbindVport()
	return &UnbindVportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateConnection 更新二层连接
//
// 当您的二层连接创建成功后，您可以通过调用此接口更新一个二层连接的名称。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EswClient) UpdateConnection(request *model.UpdateConnectionRequest) (*model.UpdateConnectionResponse, error) {
	requestDef := GenReqDefForUpdateConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateConnectionResponse), nil
	}
}

// UpdateConnectionInvoker 更新二层连接
func (c *EswClient) UpdateConnectionInvoker(request *model.UpdateConnectionRequest) *UpdateConnectionInvoker {
	requestDef := GenReqDefForUpdateConnection()
	return &UpdateConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstance 更新ESW实例
//
// 当您的ESW实例创建成功后，您可以通过调用此接口更新一个ESW实例的名称或者描述信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EswClient) UpdateInstance(request *model.UpdateInstanceRequest) (*model.UpdateInstanceResponse, error) {
	requestDef := GenReqDefForUpdateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceResponse), nil
	}
}

// UpdateInstanceInvoker 更新ESW实例
func (c *EswClient) UpdateInstanceInvoker(request *model.UpdateInstanceRequest) *UpdateInstanceInvoker {
	requestDef := GenReqDefForUpdateInstance()
	return &UpdateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
