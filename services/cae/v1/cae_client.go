package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cae/v1/model"
)

type CaeClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCaeClient(hcClient *http_client.HcHttpClient) *CaeClient {
	return &CaeClient{HcClient: hcClient}
}

func CaeClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CreateAgency 创建委托
//
// 本接口用于创建cae_trust委托
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) CreateAgency(request *model.CreateAgencyRequest) (*model.CreateAgencyResponse, error) {
	requestDef := GenReqDefForCreateAgency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAgencyResponse), nil
	}
}

// CreateAgencyInvoker 创建委托
func (c *CaeClient) CreateAgencyInvoker(request *model.CreateAgencyRequest) *CreateAgencyInvoker {
	requestDef := GenReqDefForCreateAgency()
	return &CreateAgencyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAgency 获取委托
//
// 本接口用于获取cae_trust委托，如果委托不存在则创建委托
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) ShowAgency(request *model.ShowAgencyRequest) (*model.ShowAgencyResponse, error) {
	requestDef := GenReqDefForShowAgency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAgencyResponse), nil
	}
}

// ShowAgencyInvoker 获取委托
func (c *CaeClient) ShowAgencyInvoker(request *model.ShowAgencyRequest) *ShowAgencyInvoker {
	requestDef := GenReqDefForShowAgency()
	return &ShowAgencyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateApplication 创建应用
//
// 本接口用于创建一个应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) CreateApplication(request *model.CreateApplicationRequest) (*model.CreateApplicationResponse, error) {
	requestDef := GenReqDefForCreateApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApplicationResponse), nil
	}
}

// CreateApplicationInvoker 创建应用
func (c *CaeClient) CreateApplicationInvoker(request *model.CreateApplicationRequest) *CreateApplicationInvoker {
	requestDef := GenReqDefForCreateApplication()
	return &CreateApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApplication 删除应用
//
// 本接口用于删除指定应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) DeleteApplication(request *model.DeleteApplicationRequest) (*model.DeleteApplicationResponse, error) {
	requestDef := GenReqDefForDeleteApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApplicationResponse), nil
	}
}

// DeleteApplicationInvoker 删除应用
func (c *CaeClient) DeleteApplicationInvoker(request *model.DeleteApplicationRequest) *DeleteApplicationInvoker {
	requestDef := GenReqDefForDeleteApplication()
	return &DeleteApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApplications 获取应用列表
//
// 本接口用于获取当前环境下的应用列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) ListApplications(request *model.ListApplicationsRequest) (*model.ListApplicationsResponse, error) {
	requestDef := GenReqDefForListApplications()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationsResponse), nil
	}
}

// ListApplicationsInvoker 获取应用列表
func (c *CaeClient) ListApplicationsInvoker(request *model.ListApplicationsRequest) *ListApplicationsInvoker {
	requestDef := GenReqDefForListApplications()
	return &ListApplicationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApplication 获取应用
//
// 本接口用于获取指定应用详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) ShowApplication(request *model.ShowApplicationRequest) (*model.ShowApplicationResponse, error) {
	requestDef := GenReqDefForShowApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApplicationResponse), nil
	}
}

// ShowApplicationInvoker 获取应用
func (c *CaeClient) ShowApplicationInvoker(request *model.ShowApplicationRequest) *ShowApplicationInvoker {
	requestDef := GenReqDefForShowApplication()
	return &ShowApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateComponent 创建组件
//
// 本接口用于创建一个组件，组件是CAE的最小部署单位，支持将用户的源码，部署包，镜像等资源部署到组件上
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) CreateComponent(request *model.CreateComponentRequest) (*model.CreateComponentResponse, error) {
	requestDef := GenReqDefForCreateComponent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateComponentResponse), nil
	}
}

// CreateComponentInvoker 创建组件
func (c *CaeClient) CreateComponentInvoker(request *model.CreateComponentRequest) *CreateComponentInvoker {
	requestDef := GenReqDefForCreateComponent()
	return &CreateComponentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteComponent 删除组件
//
// 本接口用于删除指定的组件，组件是CAE的最小部署单位，支持将用户的源码，部署包，镜像等资源部署到组件上
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) DeleteComponent(request *model.DeleteComponentRequest) (*model.DeleteComponentResponse, error) {
	requestDef := GenReqDefForDeleteComponent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteComponentResponse), nil
	}
}

// DeleteComponentInvoker 删除组件
func (c *CaeClient) DeleteComponentInvoker(request *model.DeleteComponentRequest) *DeleteComponentInvoker {
	requestDef := GenReqDefForDeleteComponent()
	return &DeleteComponentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteAction 操作组件
//
// 本接口用于对组件执行指定操作，如部署、升级、重启、停止、启动、伸缩、重试、配置、回滚
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) ExecuteAction(request *model.ExecuteActionRequest) (*model.ExecuteActionResponse, error) {
	requestDef := GenReqDefForExecuteAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteActionResponse), nil
	}
}

// ExecuteActionInvoker 操作组件
func (c *CaeClient) ExecuteActionInvoker(request *model.ExecuteActionRequest) *ExecuteActionInvoker {
	requestDef := GenReqDefForExecuteAction()
	return &ExecuteActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListComponentSnapshots 获取组件快照列表
//
// 本接口用于获取组件快照版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) ListComponentSnapshots(request *model.ListComponentSnapshotsRequest) (*model.ListComponentSnapshotsResponse, error) {
	requestDef := GenReqDefForListComponentSnapshots()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListComponentSnapshotsResponse), nil
	}
}

// ListComponentSnapshotsInvoker 获取组件快照列表
func (c *CaeClient) ListComponentSnapshotsInvoker(request *model.ListComponentSnapshotsRequest) *ListComponentSnapshotsInvoker {
	requestDef := GenReqDefForListComponentSnapshots()
	return &ListComponentSnapshotsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListComponents 获取组件列表
//
// 本接口用于获取组件列表，组件是CAE的最小部署单位，支持将用户的源码，部署包，镜像等资源部署到组件上
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) ListComponents(request *model.ListComponentsRequest) (*model.ListComponentsResponse, error) {
	requestDef := GenReqDefForListComponents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListComponentsResponse), nil
	}
}

// ListComponentsInvoker 获取组件列表
func (c *CaeClient) ListComponentsInvoker(request *model.ListComponentsRequest) *ListComponentsInvoker {
	requestDef := GenReqDefForListComponents()
	return &ListComponentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEvents 获取事件列表
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) ListEvents(request *model.ListEventsRequest) (*model.ListEventsResponse, error) {
	requestDef := GenReqDefForListEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEventsResponse), nil
	}
}

// ListEventsInvoker 获取事件列表
func (c *CaeClient) ListEventsInvoker(request *model.ListEventsRequest) *ListEventsInvoker {
	requestDef := GenReqDefForListEvents()
	return &ListEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstances 获取组件实例列表
//
// 本接口用于获取组件实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

// ListInstancesInvoker 获取组件实例列表
func (c *CaeClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowComponent 获取组件
//
// 本接口用于获取指定的组件，组件是CAE的最小部署单位，支持将用户的源码，部署包，镜像等资源部署到组件上
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) ShowComponent(request *model.ShowComponentRequest) (*model.ShowComponentResponse, error) {
	requestDef := GenReqDefForShowComponent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowComponentResponse), nil
	}
}

// ShowComponentInvoker 获取组件
func (c *CaeClient) ShowComponentInvoker(request *model.ShowComponentRequest) *ShowComponentInvoker {
	requestDef := GenReqDefForShowComponent()
	return &ShowComponentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateComponent 更新组件
//
// 本接口用于更新指定的组件，组件是CAE的最小部署单位，支持将用户的源码，部署包，镜像等资源部署到组件上
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) UpdateComponent(request *model.UpdateComponentRequest) (*model.UpdateComponentResponse, error) {
	requestDef := GenReqDefForUpdateComponent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateComponentResponse), nil
	}
}

// UpdateComponentInvoker 更新组件
func (c *CaeClient) UpdateComponentInvoker(request *model.UpdateComponentRequest) *UpdateComponentInvoker {
	requestDef := GenReqDefForUpdateComponent()
	return &UpdateComponentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateComponentConfiguration 创建组件配置
//
// 本接口用于创建组件配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) CreateComponentConfiguration(request *model.CreateComponentConfigurationRequest) (*model.CreateComponentConfigurationResponse, error) {
	requestDef := GenReqDefForCreateComponentConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateComponentConfigurationResponse), nil
	}
}

// CreateComponentConfigurationInvoker 创建组件配置
func (c *CaeClient) CreateComponentConfigurationInvoker(request *model.CreateComponentConfigurationRequest) *CreateComponentConfigurationInvoker {
	requestDef := GenReqDefForCreateComponentConfiguration()
	return &CreateComponentConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteComponentConfiguration 删除组件配置
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) DeleteComponentConfiguration(request *model.DeleteComponentConfigurationRequest) (*model.DeleteComponentConfigurationResponse, error) {
	requestDef := GenReqDefForDeleteComponentConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteComponentConfigurationResponse), nil
	}
}

// DeleteComponentConfigurationInvoker 删除组件配置
func (c *CaeClient) DeleteComponentConfigurationInvoker(request *model.DeleteComponentConfigurationRequest) *DeleteComponentConfigurationInvoker {
	requestDef := GenReqDefForDeleteComponentConfiguration()
	return &DeleteComponentConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConfigurations 获取组件配置列表
//
// 本接口用于获取组件配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) ListConfigurations(request *model.ListConfigurationsRequest) (*model.ListConfigurationsResponse, error) {
	requestDef := GenReqDefForListConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigurationsResponse), nil
	}
}

// ListConfigurationsInvoker 获取组件配置列表
func (c *CaeClient) ListConfigurationsInvoker(request *model.ListConfigurationsRequest) *ListConfigurationsInvoker {
	requestDef := GenReqDefForListConfigurations()
	return &ListConfigurationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEnvironment 创建环境
//
// 本接口用于创建一个环境，环境是CAE定义的一个资源维度，所有的用户组件都放在环境下
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) CreateEnvironment(request *model.CreateEnvironmentRequest) (*model.CreateEnvironmentResponse, error) {
	requestDef := GenReqDefForCreateEnvironment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEnvironmentResponse), nil
	}
}

// CreateEnvironmentInvoker 创建环境
func (c *CaeClient) CreateEnvironmentInvoker(request *model.CreateEnvironmentRequest) *CreateEnvironmentInvoker {
	requestDef := GenReqDefForCreateEnvironment()
	return &CreateEnvironmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEnvironment 删除环境
//
// 本接口用于删除环境，暂未开放。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) DeleteEnvironment(request *model.DeleteEnvironmentRequest) (*model.DeleteEnvironmentResponse, error) {
	requestDef := GenReqDefForDeleteEnvironment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEnvironmentResponse), nil
	}
}

// DeleteEnvironmentInvoker 删除环境
func (c *CaeClient) DeleteEnvironmentInvoker(request *model.DeleteEnvironmentRequest) *DeleteEnvironmentInvoker {
	requestDef := GenReqDefForDeleteEnvironment()
	return &DeleteEnvironmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnvironments 获取环境列表
//
// 本接口用于获取当前租户环境信息，环境是CAE定义的一个资源维度，所有的用户组件都放在环境下
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) ListEnvironments(request *model.ListEnvironmentsRequest) (*model.ListEnvironmentsResponse, error) {
	requestDef := GenReqDefForListEnvironments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnvironmentsResponse), nil
	}
}

// ListEnvironmentsInvoker 获取环境列表
func (c *CaeClient) ListEnvironmentsInvoker(request *model.ListEnvironmentsRequest) *ListEnvironmentsInvoker {
	requestDef := GenReqDefForListEnvironments()
	return &ListEnvironmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RetryJob 重试任务
//
// 本接口用于重试任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) RetryJob(request *model.RetryJobRequest) (*model.RetryJobResponse, error) {
	requestDef := GenReqDefForRetryJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RetryJobResponse), nil
	}
}

// RetryJobInvoker 重试任务
func (c *CaeClient) RetryJobInvoker(request *model.RetryJobRequest) *RetryJobInvoker {
	requestDef := GenReqDefForRetryJob()
	return &RetryJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJob 获取任务详情
//
// 本接口用于获取任务详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) ShowJob(request *model.ShowJobRequest) (*model.ShowJobResponse, error) {
	requestDef := GenReqDefForShowJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobResponse), nil
	}
}

// ShowJobInvoker 获取任务详情
func (c *CaeClient) ShowJobInvoker(request *model.ShowJobRequest) *ShowJobInvoker {
	requestDef := GenReqDefForShowJob()
	return &ShowJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVolume 创建卷
//
// 本接口用于创建卷
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) CreateVolume(request *model.CreateVolumeRequest) (*model.CreateVolumeResponse, error) {
	requestDef := GenReqDefForCreateVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVolumeResponse), nil
	}
}

// CreateVolumeInvoker 创建卷
func (c *CaeClient) CreateVolumeInvoker(request *model.CreateVolumeRequest) *CreateVolumeInvoker {
	requestDef := GenReqDefForCreateVolume()
	return &CreateVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVolume 删除卷
//
// 本接口用于创建卷
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) DeleteVolume(request *model.DeleteVolumeRequest) (*model.DeleteVolumeResponse, error) {
	requestDef := GenReqDefForDeleteVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVolumeResponse), nil
	}
}

// DeleteVolumeInvoker 删除卷
func (c *CaeClient) DeleteVolumeInvoker(request *model.DeleteVolumeRequest) *DeleteVolumeInvoker {
	requestDef := GenReqDefForDeleteVolume()
	return &DeleteVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVolumes 获取卷列表
//
// 本接口用于获取卷列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) ListVolumes(request *model.ListVolumesRequest) (*model.ListVolumesResponse, error) {
	requestDef := GenReqDefForListVolumes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVolumesResponse), nil
	}
}

// ListVolumesInvoker 获取卷列表
func (c *CaeClient) ListVolumesInvoker(request *model.ListVolumesRequest) *ListVolumesInvoker {
	requestDef := GenReqDefForListVolumes()
	return &ListVolumesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
