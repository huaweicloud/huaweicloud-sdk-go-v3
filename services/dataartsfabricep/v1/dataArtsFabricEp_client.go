package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dataartsfabricep/v1/model"
)

type DataArtsFabricEpClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewDataArtsFabricEpClient(hcClient *httpclient.HcHttpClient) *DataArtsFabricEpClient {
	return &DataArtsFabricEpClient{HcClient: hcClient}
}

func DataArtsFabricEpClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// DeleteServiceInstance 删除Service实例
//
// 删除Service实例，释放该实例的资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricEpClient) DeleteServiceInstance(request *model.DeleteServiceInstanceRequest) (*model.DeleteServiceInstanceResponse, error) {
	requestDef := GenReqDefForDeleteServiceInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServiceInstanceResponse), nil
	}
}

// DeleteServiceInstanceInvoker 删除Service实例
func (c *DataArtsFabricEpClient) DeleteServiceInstanceInvoker(request *model.DeleteServiceInstanceRequest) *DeleteServiceInstanceInvoker {
	requestDef := GenReqDefForDeleteServiceInstance()
	return &DeleteServiceInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeployServiceInstance 部署服务
//
// 部署一个Service实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricEpClient) DeployServiceInstance(request *model.DeployServiceInstanceRequest) (*model.DeployServiceInstanceResponse, error) {
	requestDef := GenReqDefForDeployServiceInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeployServiceInstanceResponse), nil
	}
}

// DeployServiceInstanceInvoker 部署服务
func (c *DataArtsFabricEpClient) DeployServiceInstanceInvoker(request *model.DeployServiceInstanceRequest) *DeployServiceInstanceInvoker {
	requestDef := GenReqDefForDeployServiceInstance()
	return &DeployServiceInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServicesInstances 列举已部署的Service实例
//
// 列举已部署的Service实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricEpClient) ListServicesInstances(request *model.ListServicesInstancesRequest) (*model.ListServicesInstancesResponse, error) {
	requestDef := GenReqDefForListServicesInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServicesInstancesResponse), nil
	}
}

// ListServicesInstancesInvoker 列举已部署的Service实例
func (c *DataArtsFabricEpClient) ListServicesInstancesInvoker(request *model.ListServicesInstancesRequest) *ListServicesInstancesInvoker {
	requestDef := GenReqDefForListServicesInstances()
	return &ListServicesInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowServiceInstanceDetail 查看部署的Service实例详情
//
// 查看部署后的Service实例的详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricEpClient) ShowServiceInstanceDetail(request *model.ShowServiceInstanceDetailRequest) (*model.ShowServiceInstanceDetailResponse, error) {
	requestDef := GenReqDefForShowServiceInstanceDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServiceInstanceDetailResponse), nil
	}
}

// ShowServiceInstanceDetailInvoker 查看部署的Service实例详情
func (c *DataArtsFabricEpClient) ShowServiceInstanceDetailInvoker(request *model.ShowServiceInstanceDetailRequest) *ShowServiceInstanceDetailInvoker {
	requestDef := GenReqDefForShowServiceInstanceDetail()
	return &ShowServiceInstanceDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateServiceInstance 更新Service实例
//
// 更新已部署的Service实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsFabricEpClient) UpdateServiceInstance(request *model.UpdateServiceInstanceRequest) (*model.UpdateServiceInstanceResponse, error) {
	requestDef := GenReqDefForUpdateServiceInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateServiceInstanceResponse), nil
	}
}

// UpdateServiceInstanceInvoker 更新Service实例
func (c *DataArtsFabricEpClient) UpdateServiceInstanceInvoker(request *model.UpdateServiceInstanceRequest) *UpdateServiceInstanceInvoker {
	requestDef := GenReqDefForUpdateServiceInstance()
	return &UpdateServiceInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
