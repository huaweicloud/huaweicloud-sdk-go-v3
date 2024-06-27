package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/servicestage/v2/model"
)

type ServiceStageClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewServiceStageClient(hcClient *httpclient.HcHttpClient) *ServiceStageClient {
	return &ServiceStageClient{HcClient: hcClient}
}

func ServiceStageClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// ChangeApplication 修改应用信息
//
// 此API通过应用ID修改应用信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ChangeApplication(request *model.ChangeApplicationRequest) (*model.ChangeApplicationResponse, error) {
	requestDef := GenReqDefForChangeApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeApplicationResponse), nil
	}
}

// ChangeApplicationInvoker 修改应用信息
func (c *ServiceStageClient) ChangeApplicationInvoker(request *model.ChangeApplicationRequest) *ChangeApplicationInvoker {
	requestDef := GenReqDefForChangeApplication()
	return &ChangeApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeApplicationConfiguration 修改应用配置信息
//
// 通过此API修改应用配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ChangeApplicationConfiguration(request *model.ChangeApplicationConfigurationRequest) (*model.ChangeApplicationConfigurationResponse, error) {
	requestDef := GenReqDefForChangeApplicationConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeApplicationConfigurationResponse), nil
	}
}

// ChangeApplicationConfigurationInvoker 修改应用配置信息
func (c *ServiceStageClient) ChangeApplicationConfigurationInvoker(request *model.ChangeApplicationConfigurationRequest) *ChangeApplicationConfigurationInvoker {
	requestDef := GenReqDefForChangeApplicationConfiguration()
	return &ChangeApplicationConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeComponent 根据组件ID修改组件信息
//
// 此API通过组件ID修改组件信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ChangeComponent(request *model.ChangeComponentRequest) (*model.ChangeComponentResponse, error) {
	requestDef := GenReqDefForChangeComponent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeComponentResponse), nil
	}
}

// ChangeComponentInvoker 根据组件ID修改组件信息
func (c *ServiceStageClient) ChangeComponentInvoker(request *model.ChangeComponentRequest) *ChangeComponentInvoker {
	requestDef := GenReqDefForChangeComponent()
	return &ChangeComponentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeEnvironment 修改环境信息
//
// 此API通过环境ID修改环境信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ChangeEnvironment(request *model.ChangeEnvironmentRequest) (*model.ChangeEnvironmentResponse, error) {
	requestDef := GenReqDefForChangeEnvironment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeEnvironmentResponse), nil
	}
}

// ChangeEnvironmentInvoker 修改环境信息
func (c *ServiceStageClient) ChangeEnvironmentInvoker(request *model.ChangeEnvironmentRequest) *ChangeEnvironmentInvoker {
	requestDef := GenReqDefForChangeEnvironment()
	return &ChangeEnvironmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeInstance 修改应用组件实例
//
// 通过此API修改应用组件实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ChangeInstance(request *model.ChangeInstanceRequest) (*model.ChangeInstanceResponse, error) {
	requestDef := GenReqDefForChangeInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeInstanceResponse), nil
	}
}

// ChangeInstanceInvoker 修改应用组件实例
func (c *ServiceStageClient) ChangeInstanceInvoker(request *model.ChangeInstanceRequest) *ChangeInstanceInvoker {
	requestDef := GenReqDefForChangeInstance()
	return &ChangeInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeResourceInEnvironment 修改环境资源
//
// 此API用来修改环境资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ChangeResourceInEnvironment(request *model.ChangeResourceInEnvironmentRequest) (*model.ChangeResourceInEnvironmentResponse, error) {
	requestDef := GenReqDefForChangeResourceInEnvironment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeResourceInEnvironmentResponse), nil
	}
}

// ChangeResourceInEnvironmentInvoker 修改环境资源
func (c *ServiceStageClient) ChangeResourceInEnvironmentInvoker(request *model.ChangeResourceInEnvironmentRequest) *ChangeResourceInEnvironmentInvoker {
	requestDef := GenReqDefForChangeResourceInEnvironment()
	return &ChangeResourceInEnvironmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateApplication 创建应用
//
// 应用是一个功能相对完备的业务系统，由一个或多个特性相关的组件组成。
//
// 此API用来创建应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) CreateApplication(request *model.CreateApplicationRequest) (*model.CreateApplicationResponse, error) {
	requestDef := GenReqDefForCreateApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApplicationResponse), nil
	}
}

// CreateApplicationInvoker 创建应用
func (c *ServiceStageClient) CreateApplicationInvoker(request *model.CreateApplicationRequest) *CreateApplicationInvoker {
	requestDef := GenReqDefForCreateApplication()
	return &CreateApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCamInstance 创建、更新实例
//
// 创建、更新实例
func (c *ServiceStageClient) CreateCamInstance(request *model.CreateCamInstanceRequest) (*model.CreateCamInstanceResponse, error) {
	requestDef := GenReqDefForCreateCamInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCamInstanceResponse), nil
	}
}

// CreateCamInstanceInvoker 创建、更新实例
func (c *ServiceStageClient) CreateCamInstanceInvoker(request *model.CreateCamInstanceRequest) *CreateCamInstanceInvoker {
	requestDef := GenReqDefForCreateCamInstance()
	return &CreateCamInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateComponent 应用中创建组件
//
// 应用组件是组成应用的某个业务特性实现，以代码或者软件包为载体，可独立部署运行。
//
// 此API用来在应用中创建组件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) CreateComponent(request *model.CreateComponentRequest) (*model.CreateComponentResponse, error) {
	requestDef := GenReqDefForCreateComponent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateComponentResponse), nil
	}
}

// CreateComponentInvoker 应用中创建组件
func (c *ServiceStageClient) CreateComponentInvoker(request *model.CreateComponentRequest) *CreateComponentInvoker {
	requestDef := GenReqDefForCreateComponent()
	return &CreateComponentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEnvironment 创建环境
//
// 环境是用于应用部署和运行的计算、存储、网络等基础设施的集合。Servicestage把相同VPC下的CCE集群加上多个ELB、RDS、DCS实例组合为一个环境，如：开发环境，测试环境，预生产环境，生产环境。环境内网络互通，可以按环境维度来管理资源、部署服务，减少具体基础设施运维管理的复杂性。
//
// 此API用来创建环境。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) CreateEnvironment(request *model.CreateEnvironmentRequest) (*model.CreateEnvironmentResponse, error) {
	requestDef := GenReqDefForCreateEnvironment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEnvironmentResponse), nil
	}
}

// CreateEnvironmentInvoker 创建环境
func (c *ServiceStageClient) CreateEnvironmentInvoker(request *model.CreateEnvironmentRequest) *CreateEnvironmentInvoker {
	requestDef := GenReqDefForCreateEnvironment()
	return &CreateEnvironmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstance 创建组件实例
//
// 此API用来创建应用组件实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) CreateInstance(request *model.CreateInstanceRequest) (*model.CreateInstanceResponse, error) {
	requestDef := GenReqDefForCreateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceResponse), nil
	}
}

// CreateInstanceInvoker 创建组件实例
func (c *ServiceStageClient) CreateInstanceInvoker(request *model.CreateInstanceRequest) *CreateInstanceInvoker {
	requestDef := GenReqDefForCreateInstance()
	return &CreateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTemplate 创建模板
//
// 创建模板
func (c *ServiceStageClient) CreateTemplate(request *model.CreateTemplateRequest) (*model.CreateTemplateResponse, error) {
	requestDef := GenReqDefForCreateTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTemplateResponse), nil
	}
}

// CreateTemplateInvoker 创建模板
func (c *ServiceStageClient) CreateTemplateInvoker(request *model.CreateTemplateRequest) *CreateTemplateInvoker {
	requestDef := GenReqDefForCreateTemplate()
	return &CreateTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApplication 根据应用ID删除应用
//
// 此API通过应用ID删除应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) DeleteApplication(request *model.DeleteApplicationRequest) (*model.DeleteApplicationResponse, error) {
	requestDef := GenReqDefForDeleteApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApplicationResponse), nil
	}
}

// DeleteApplicationInvoker 根据应用ID删除应用
func (c *ServiceStageClient) DeleteApplicationInvoker(request *model.DeleteApplicationRequest) *DeleteApplicationInvoker {
	requestDef := GenReqDefForDeleteApplication()
	return &DeleteApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApplicationConfiguration 删除应用配置
//
// 通过此API删除应用配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) DeleteApplicationConfiguration(request *model.DeleteApplicationConfigurationRequest) (*model.DeleteApplicationConfigurationResponse, error) {
	requestDef := GenReqDefForDeleteApplicationConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApplicationConfigurationResponse), nil
	}
}

// DeleteApplicationConfigurationInvoker 删除应用配置
func (c *ServiceStageClient) DeleteApplicationConfigurationInvoker(request *model.DeleteApplicationConfigurationRequest) *DeleteApplicationConfigurationInvoker {
	requestDef := GenReqDefForDeleteApplicationConfiguration()
	return &DeleteApplicationConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteComponent 根据应用组件ID删除应用组件
//
// 此API通过应用组件ID删除应用组件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) DeleteComponent(request *model.DeleteComponentRequest) (*model.DeleteComponentResponse, error) {
	requestDef := GenReqDefForDeleteComponent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteComponentResponse), nil
	}
}

// DeleteComponentInvoker 根据应用组件ID删除应用组件
func (c *ServiceStageClient) DeleteComponentInvoker(request *model.DeleteComponentRequest) *DeleteComponentInvoker {
	requestDef := GenReqDefForDeleteComponent()
	return &DeleteComponentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEnvironment 根据环境ID删除环境
//
// 此API通过环境ID删除环境。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) DeleteEnvironment(request *model.DeleteEnvironmentRequest) (*model.DeleteEnvironmentResponse, error) {
	requestDef := GenReqDefForDeleteEnvironment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEnvironmentResponse), nil
	}
}

// DeleteEnvironmentInvoker 根据环境ID删除环境
func (c *ServiceStageClient) DeleteEnvironmentInvoker(request *model.DeleteEnvironmentRequest) *DeleteEnvironmentInvoker {
	requestDef := GenReqDefForDeleteEnvironment()
	return &DeleteEnvironmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstance 删除应用组件实例
//
// 通过此API删除应用组件实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResponse), nil
	}
}

// DeleteInstanceInvoker 删除应用组件实例
func (c *ServiceStageClient) DeleteInstanceInvoker(request *model.DeleteInstanceRequest) *DeleteInstanceInvoker {
	requestDef := GenReqDefForDeleteInstance()
	return &DeleteInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstanceById 删除实例
//
// 删除实例
func (c *ServiceStageClient) DeleteInstanceById(request *model.DeleteInstanceByIdRequest) (*model.DeleteInstanceByIdResponse, error) {
	requestDef := GenReqDefForDeleteInstanceById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceByIdResponse), nil
	}
}

// DeleteInstanceByIdInvoker 删除实例
func (c *ServiceStageClient) DeleteInstanceByIdInvoker(request *model.DeleteInstanceByIdRequest) *DeleteInstanceByIdInvoker {
	requestDef := GenReqDefForDeleteInstanceById()
	return &DeleteInstanceByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTemplate 删除模板
//
// 删除模板
func (c *ServiceStageClient) DeleteTemplate(request *model.DeleteTemplateRequest) (*model.DeleteTemplateResponse, error) {
	requestDef := GenReqDefForDeleteTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTemplateResponse), nil
	}
}

// DeleteTemplateInvoker 删除模板
func (c *ServiceStageClient) DeleteTemplateInvoker(request *model.DeleteTemplateRequest) *DeleteTemplateInvoker {
	requestDef := GenReqDefForDeleteTemplate()
	return &DeleteTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeployInstance 部署实例
//
// 部署实例
func (c *ServiceStageClient) DeployInstance(request *model.DeployInstanceRequest) (*model.DeployInstanceResponse, error) {
	requestDef := GenReqDefForDeployInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeployInstanceResponse), nil
	}
}

// DeployInstanceInvoker 部署实例
func (c *ServiceStageClient) DeployInstanceInvoker(request *model.DeployInstanceRequest) *DeployInstanceInvoker {
	requestDef := GenReqDefForDeployInstance()
	return &DeployInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApplications 获取所有应用
//
// 通过此API可以获取所有已经创建的应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ListApplications(request *model.ListApplicationsRequest) (*model.ListApplicationsResponse, error) {
	requestDef := GenReqDefForListApplications()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationsResponse), nil
	}
}

// ListApplicationsInvoker 获取所有应用
func (c *ServiceStageClient) ListApplicationsInvoker(request *model.ListApplicationsRequest) *ListApplicationsInvoker {
	requestDef := GenReqDefForListApplications()
	return &ListApplicationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListComponentOverviews 获取应用所有组件部署信息
//
// 通过此API获取应用下所有应用组件部署信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ListComponentOverviews(request *model.ListComponentOverviewsRequest) (*model.ListComponentOverviewsResponse, error) {
	requestDef := GenReqDefForListComponentOverviews()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListComponentOverviewsResponse), nil
	}
}

// ListComponentOverviewsInvoker 获取应用所有组件部署信息
func (c *ServiceStageClient) ListComponentOverviewsInvoker(request *model.ListComponentOverviewsRequest) *ListComponentOverviewsInvoker {
	requestDef := GenReqDefForListComponentOverviews()
	return &ListComponentOverviewsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListComponents 获取应用所有组件
//
// 通过此API获取应用下所有应用组件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ListComponents(request *model.ListComponentsRequest) (*model.ListComponentsResponse, error) {
	requestDef := GenReqDefForListComponents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListComponentsResponse), nil
	}
}

// ListComponentsInvoker 获取应用所有组件
func (c *ServiceStageClient) ListComponentsInvoker(request *model.ListComponentsRequest) *ListComponentsInvoker {
	requestDef := GenReqDefForListComponents()
	return &ListComponentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnvironments 获取所有环境
//
// 此API用来获取所有已经创建环境。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ListEnvironments(request *model.ListEnvironmentsRequest) (*model.ListEnvironmentsResponse, error) {
	requestDef := GenReqDefForListEnvironments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnvironmentsResponse), nil
	}
}

// ListEnvironmentsInvoker 获取所有环境
func (c *ServiceStageClient) ListEnvironmentsInvoker(request *model.ListEnvironmentsRequest) *ListEnvironmentsInvoker {
	requestDef := GenReqDefForListEnvironments()
	return &ListEnvironmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceSnapshots 获取组件实例快照
//
// 通过此API获取应用组件实例快照信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ListInstanceSnapshots(request *model.ListInstanceSnapshotsRequest) (*model.ListInstanceSnapshotsResponse, error) {
	requestDef := GenReqDefForListInstanceSnapshots()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceSnapshotsResponse), nil
	}
}

// ListInstanceSnapshotsInvoker 获取组件实例快照
func (c *ServiceStageClient) ListInstanceSnapshotsInvoker(request *model.ListInstanceSnapshotsRequest) *ListInstanceSnapshotsInvoker {
	requestDef := GenReqDefForListInstanceSnapshots()
	return &ListInstanceSnapshotsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstances 获取应用组件实例
//
// 通过此API获取组件下的所有组件实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

// ListInstancesInvoker 获取应用组件实例
func (c *ServiceStageClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApplicationConfiguration 获取应用配置
//
// 通过此API获取应用配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ShowApplicationConfiguration(request *model.ShowApplicationConfigurationRequest) (*model.ShowApplicationConfigurationResponse, error) {
	requestDef := GenReqDefForShowApplicationConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApplicationConfigurationResponse), nil
	}
}

// ShowApplicationConfigurationInvoker 获取应用配置
func (c *ServiceStageClient) ShowApplicationConfigurationInvoker(request *model.ShowApplicationConfigurationRequest) *ShowApplicationConfigurationInvoker {
	requestDef := GenReqDefForShowApplicationConfiguration()
	return &ShowApplicationConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApplicationDetail 根据应用ID获取应用详细信息
//
// 此API通过应用ID获取应用详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ShowApplicationDetail(request *model.ShowApplicationDetailRequest) (*model.ShowApplicationDetailResponse, error) {
	requestDef := GenReqDefForShowApplicationDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApplicationDetailResponse), nil
	}
}

// ShowApplicationDetailInvoker 根据应用ID获取应用详细信息
func (c *ServiceStageClient) ShowApplicationDetailInvoker(request *model.ShowApplicationDetailRequest) *ShowApplicationDetailInvoker {
	requestDef := GenReqDefForShowApplicationDetail()
	return &ShowApplicationDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowComponentDetail 根据组件ID获取应用组件信息
//
// 通过组件ID获取应用组件信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ShowComponentDetail(request *model.ShowComponentDetailRequest) (*model.ShowComponentDetailResponse, error) {
	requestDef := GenReqDefForShowComponentDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowComponentDetailResponse), nil
	}
}

// ShowComponentDetailInvoker 根据组件ID获取应用组件信息
func (c *ServiceStageClient) ShowComponentDetailInvoker(request *model.ShowComponentDetailRequest) *ShowComponentDetailInvoker {
	requestDef := GenReqDefForShowComponentDetail()
	return &ShowComponentDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEnvironmentDetail 根据环境ID获取环境详细信息
//
// 此API通过环境ID获取环境详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ShowEnvironmentDetail(request *model.ShowEnvironmentDetailRequest) (*model.ShowEnvironmentDetailResponse, error) {
	requestDef := GenReqDefForShowEnvironmentDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEnvironmentDetailResponse), nil
	}
}

// ShowEnvironmentDetailInvoker 根据环境ID获取环境详细信息
func (c *ServiceStageClient) ShowEnvironmentDetailInvoker(request *model.ShowEnvironmentDetailRequest) *ShowEnvironmentDetailInvoker {
	requestDef := GenReqDefForShowEnvironmentDetail()
	return &ShowEnvironmentDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceDetail 根据实例ID获取实例详细信息
//
// 此API通过实例ID获取实例详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ShowInstanceDetail(request *model.ShowInstanceDetailRequest) (*model.ShowInstanceDetailResponse, error) {
	requestDef := GenReqDefForShowInstanceDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceDetailResponse), nil
	}
}

// ShowInstanceDetailInvoker 根据实例ID获取实例详细信息
func (c *ServiceStageClient) ShowInstanceDetailInvoker(request *model.ShowInstanceDetailRequest) *ShowInstanceDetailInvoker {
	requestDef := GenReqDefForShowInstanceDetail()
	return &ShowInstanceDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobDetail 获取部署任务详细信息
//
// 通过此API获取部署任务详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ShowJobDetail(request *model.ShowJobDetailRequest) (*model.ShowJobDetailResponse, error) {
	requestDef := GenReqDefForShowJobDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobDetailResponse), nil
	}
}

// ShowJobDetailInvoker 获取部署任务详细信息
func (c *ServiceStageClient) ShowJobDetailInvoker(request *model.ShowJobDetailRequest) *ShowJobDetailInvoker {
	requestDef := GenReqDefForShowJobDetail()
	return &ShowJobDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceAction 对组件实例的操作
//
// 通过此API获取对组件实例的操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) UpdateInstanceAction(request *model.UpdateInstanceActionRequest) (*model.UpdateInstanceActionResponse, error) {
	requestDef := GenReqDefForUpdateInstanceAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceActionResponse), nil
	}
}

// UpdateInstanceActionInvoker 对组件实例的操作
func (c *ServiceStageClient) UpdateInstanceActionInvoker(request *model.UpdateInstanceActionRequest) *UpdateInstanceActionInvoker {
	requestDef := GenReqDefForUpdateInstanceAction()
	return &UpdateInstanceActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTemplate 更新模板
//
// 更新模板
func (c *ServiceStageClient) UpdateTemplate(request *model.UpdateTemplateRequest) (*model.UpdateTemplateResponse, error) {
	requestDef := GenReqDefForUpdateTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTemplateResponse), nil
	}
}

// UpdateTemplateInvoker 更新模板
func (c *ServiceStageClient) UpdateTemplateInvoker(request *model.UpdateTemplateRequest) *UpdateTemplateInvoker {
	requestDef := GenReqDefForUpdateTemplate()
	return &UpdateTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFile 创建仓库文件
//
// 在指定仓库项目下创建文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) CreateFile(request *model.CreateFileRequest) (*model.CreateFileResponse, error) {
	requestDef := GenReqDefForCreateFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFileResponse), nil
	}
}

// CreateFileInvoker 创建仓库文件
func (c *ServiceStageClient) CreateFileInvoker(request *model.CreateFileRequest) *CreateFileInvoker {
	requestDef := GenReqDefForCreateFile()
	return &CreateFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHook 创建项目hook
//
// 创建指定项目的hook。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) CreateHook(request *model.CreateHookRequest) (*model.CreateHookResponse, error) {
	requestDef := GenReqDefForCreateHook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHookResponse), nil
	}
}

// CreateHookInvoker 创建项目hook
func (c *ServiceStageClient) CreateHookInvoker(request *model.CreateHookRequest) *CreateHookInvoker {
	requestDef := GenReqDefForCreateHook()
	return &CreateHookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateOAuth 创建OAuth授权
//
// 创建指定Git仓库类型的OAuth授权。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) CreateOAuth(request *model.CreateOAuthRequest) (*model.CreateOAuthResponse, error) {
	requestDef := GenReqDefForCreateOAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOAuthResponse), nil
	}
}

// CreateOAuthInvoker 创建OAuth授权
func (c *ServiceStageClient) CreateOAuthInvoker(request *model.CreateOAuthRequest) *CreateOAuthInvoker {
	requestDef := GenReqDefForCreateOAuth()
	return &CreateOAuthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePasswordAuth 创建口令授权
//
// 创建指定Git仓库类型的口令授权。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) CreatePasswordAuth(request *model.CreatePasswordAuthRequest) (*model.CreatePasswordAuthResponse, error) {
	requestDef := GenReqDefForCreatePasswordAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePasswordAuthResponse), nil
	}
}

// CreatePasswordAuthInvoker 创建口令授权
func (c *ServiceStageClient) CreatePasswordAuthInvoker(request *model.CreatePasswordAuthRequest) *CreatePasswordAuthInvoker {
	requestDef := GenReqDefForCreatePasswordAuth()
	return &CreatePasswordAuthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePersonalAuth 创建私人令牌授权
//
// 创建指定Git仓库类型的私人令牌授权。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) CreatePersonalAuth(request *model.CreatePersonalAuthRequest) (*model.CreatePersonalAuthResponse, error) {
	requestDef := GenReqDefForCreatePersonalAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePersonalAuthResponse), nil
	}
}

// CreatePersonalAuthInvoker 创建私人令牌授权
func (c *ServiceStageClient) CreatePersonalAuthInvoker(request *model.CreatePersonalAuthRequest) *CreatePersonalAuthInvoker {
	requestDef := GenReqDefForCreatePersonalAuth()
	return &CreatePersonalAuthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateProject 创建软件仓库项目
//
// 创建指定组织下的软件仓库项目。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) CreateProject(request *model.CreateProjectRequest) (*model.CreateProjectResponse, error) {
	requestDef := GenReqDefForCreateProject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProjectResponse), nil
	}
}

// CreateProjectInvoker 创建软件仓库项目
func (c *ServiceStageClient) CreateProjectInvoker(request *model.CreateProjectRequest) *CreateProjectInvoker {
	requestDef := GenReqDefForCreateProject()
	return &CreateProjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTag 创建项目tag标签
//
// 创建指定项目的tag标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) CreateTag(request *model.CreateTagRequest) (*model.CreateTagResponse, error) {
	requestDef := GenReqDefForCreateTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTagResponse), nil
	}
}

// CreateTagInvoker 创建项目tag标签
func (c *ServiceStageClient) CreateTagInvoker(request *model.CreateTagRequest) *CreateTagInvoker {
	requestDef := GenReqDefForCreateTag()
	return &CreateTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAuthorize 删除仓库授权
//
// 通过名称删除仓库授权。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) DeleteAuthorize(request *model.DeleteAuthorizeRequest) (*model.DeleteAuthorizeResponse, error) {
	requestDef := GenReqDefForDeleteAuthorize()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAuthorizeResponse), nil
	}
}

// DeleteAuthorizeInvoker 删除仓库授权
func (c *ServiceStageClient) DeleteAuthorizeInvoker(request *model.DeleteAuthorizeRequest) *DeleteAuthorizeInvoker {
	requestDef := GenReqDefForDeleteAuthorize()
	return &DeleteAuthorizeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFile 删除仓库文件
//
// 删除指定项目仓库下的文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) DeleteFile(request *model.DeleteFileRequest) (*model.DeleteFileResponse, error) {
	requestDef := GenReqDefForDeleteFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFileResponse), nil
	}
}

// DeleteFileInvoker 删除仓库文件
func (c *ServiceStageClient) DeleteFileInvoker(request *model.DeleteFileRequest) *DeleteFileInvoker {
	requestDef := GenReqDefForDeleteFile()
	return &DeleteFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHook 删除项目hook
//
// 删除指定项目的hook。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) DeleteHook(request *model.DeleteHookRequest) (*model.DeleteHookResponse, error) {
	requestDef := GenReqDefForDeleteHook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHookResponse), nil
	}
}

// DeleteHookInvoker 删除项目hook
func (c *ServiceStageClient) DeleteHookInvoker(request *model.DeleteHookRequest) *DeleteHookInvoker {
	requestDef := GenReqDefForDeleteHook()
	return &DeleteHookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTag 删除项目tag标签
//
// 删除指定项目的tag标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) DeleteTag(request *model.DeleteTagRequest) (*model.DeleteTagResponse, error) {
	requestDef := GenReqDefForDeleteTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTagResponse), nil
	}
}

// DeleteTagInvoker 删除项目tag标签
func (c *ServiceStageClient) DeleteTagInvoker(request *model.DeleteTagRequest) *DeleteTagInvoker {
	requestDef := GenReqDefForDeleteTag()
	return &DeleteTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuthorizations 获取仓库授权列表
//
// 获取所有Git仓库授权信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ListAuthorizations(request *model.ListAuthorizationsRequest) (*model.ListAuthorizationsResponse, error) {
	requestDef := GenReqDefForListAuthorizations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuthorizationsResponse), nil
	}
}

// ListAuthorizationsInvoker 获取仓库授权列表
func (c *ServiceStageClient) ListAuthorizationsInvoker(request *model.ListAuthorizationsRequest) *ListAuthorizationsInvoker {
	requestDef := GenReqDefForListAuthorizations()
	return &ListAuthorizationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBranches 获取项目分支
//
// 获取指定项目的所有分支列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ListBranches(request *model.ListBranchesRequest) (*model.ListBranchesResponse, error) {
	requestDef := GenReqDefForListBranches()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBranchesResponse), nil
	}
}

// ListBranchesInvoker 获取项目分支
func (c *ServiceStageClient) ListBranchesInvoker(request *model.ListBranchesRequest) *ListBranchesInvoker {
	requestDef := GenReqDefForListBranches()
	return &ListBranchesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCommits 获取项目commit提交记录
//
// 获取指定项目的最近10次commit提交记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ListCommits(request *model.ListCommitsRequest) (*model.ListCommitsResponse, error) {
	requestDef := GenReqDefForListCommits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCommitsResponse), nil
	}
}

// ListCommitsInvoker 获取项目commit提交记录
func (c *ServiceStageClient) ListCommitsInvoker(request *model.ListCommitsRequest) *ListCommitsInvoker {
	requestDef := GenReqDefForListCommits()
	return &ListCommitsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHooks 获取项目hooks
//
// 获取指定项目的所有hooks
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ListHooks(request *model.ListHooksRequest) (*model.ListHooksResponse, error) {
	requestDef := GenReqDefForListHooks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHooksResponse), nil
	}
}

// ListHooksInvoker 获取项目hooks
func (c *ServiceStageClient) ListHooksInvoker(request *model.ListHooksRequest) *ListHooksInvoker {
	requestDef := GenReqDefForListHooks()
	return &ListHooksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNamespaces 获取仓库的namespaces
//
// 获取仓库的namespaces。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ListNamespaces(request *model.ListNamespacesRequest) (*model.ListNamespacesResponse, error) {
	requestDef := GenReqDefForListNamespaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNamespacesResponse), nil
	}
}

// ListNamespacesInvoker 获取仓库的namespaces
func (c *ServiceStageClient) ListNamespacesInvoker(request *model.ListNamespacesRequest) *ListNamespacesInvoker {
	requestDef := GenReqDefForListNamespaces()
	return &ListNamespacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjects 获取组织下所有项目
//
// 获取指定组织下的所有项目。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ListProjects(request *model.ListProjectsRequest) (*model.ListProjectsResponse, error) {
	requestDef := GenReqDefForListProjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectsResponse), nil
	}
}

// ListProjectsInvoker 获取组织下所有项目
func (c *ServiceStageClient) ListProjectsInvoker(request *model.ListProjectsRequest) *ListProjectsInvoker {
	requestDef := GenReqDefForListProjects()
	return &ListProjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTags 获取项目的所有tag标签
//
// 获取指定项目的所有tag标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ListTags(request *model.ListTagsRequest) (*model.ListTagsResponse, error) {
	requestDef := GenReqDefForListTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsResponse), nil
	}
}

// ListTagsInvoker 获取项目的所有tag标签
func (c *ServiceStageClient) ListTagsInvoker(request *model.ListTagsRequest) *ListTagsInvoker {
	requestDef := GenReqDefForListTags()
	return &ListTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTrees 获取仓库文件列表
//
// 获取指定项目仓库的文件列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ListTrees(request *model.ListTreesRequest) (*model.ListTreesResponse, error) {
	requestDef := GenReqDefForListTrees()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTreesResponse), nil
	}
}

// ListTreesInvoker 获取仓库文件列表
func (c *ServiceStageClient) ListTreesInvoker(request *model.ListTreesRequest) *ListTreesInvoker {
	requestDef := GenReqDefForListTrees()
	return &ListTreesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowContent 获取仓库文件内容
//
// 获取指定项目仓库下文件的内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ShowContent(request *model.ShowContentRequest) (*model.ShowContentResponse, error) {
	requestDef := GenReqDefForShowContent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowContentResponse), nil
	}
}

// ShowContentInvoker 获取仓库文件内容
func (c *ServiceStageClient) ShowContentInvoker(request *model.ShowContentRequest) *ShowContentInvoker {
	requestDef := GenReqDefForShowContent()
	return &ShowContentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectDetail 通过clone url 获取仓库信息
//
// 通过指定的clone url 获取仓库信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ShowProjectDetail(request *model.ShowProjectDetailRequest) (*model.ShowProjectDetailResponse, error) {
	requestDef := GenReqDefForShowProjectDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectDetailResponse), nil
	}
}

// ShowProjectDetailInvoker 通过clone url 获取仓库信息
func (c *ServiceStageClient) ShowProjectDetailInvoker(request *model.ShowProjectDetailRequest) *ShowProjectDetailInvoker {
	requestDef := GenReqDefForShowProjectDetail()
	return &ShowProjectDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRedirectUrl 获取授权重定向URL
//
// 获取指定Git仓库类型的授权重定向URL。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ShowRedirectUrl(request *model.ShowRedirectUrlRequest) (*model.ShowRedirectUrlResponse, error) {
	requestDef := GenReqDefForShowRedirectUrl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRedirectUrlResponse), nil
	}
}

// ShowRedirectUrlInvoker 获取授权重定向URL
func (c *ServiceStageClient) ShowRedirectUrlInvoker(request *model.ShowRedirectUrlRequest) *ShowRedirectUrlInvoker {
	requestDef := GenReqDefForShowRedirectUrl()
	return &ShowRedirectUrlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFile 更新仓库文件内容
//
// 更新指定项目仓库下的文件内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) UpdateFile(request *model.UpdateFileRequest) (*model.UpdateFileResponse, error) {
	requestDef := GenReqDefForUpdateFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFileResponse), nil
	}
}

// UpdateFileInvoker 更新仓库文件内容
func (c *ServiceStageClient) UpdateFileInvoker(request *model.UpdateFileRequest) *UpdateFileInvoker {
	requestDef := GenReqDefForUpdateFile()
	return &UpdateFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlavors 获取所有支持的应用资源规格
//
// 通过此API获取所用支持的应用资源规格。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsResponse), nil
	}
}

// ListFlavorsInvoker 获取所有支持的应用资源规格
func (c *ServiceStageClient) ListFlavorsInvoker(request *model.ListFlavorsRequest) *ListFlavorsInvoker {
	requestDef := GenReqDefForListFlavors()
	return &ListFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRuntimes 获取所有支持的应用组件运行时类型
//
// 此API用来获取所有支持应用组件运行时类型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ListRuntimes(request *model.ListRuntimesRequest) (*model.ListRuntimesResponse, error) {
	requestDef := GenReqDefForListRuntimes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRuntimesResponse), nil
	}
}

// ListRuntimesInvoker 获取所有支持的应用组件运行时类型
func (c *ServiceStageClient) ListRuntimesInvoker(request *model.ListRuntimesRequest) *ListRuntimesInvoker {
	requestDef := GenReqDefForListRuntimes()
	return &ListRuntimesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTemplates 获取所有支持的应用组件模板
//
// 此API用来获取所有内置应用组件模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ServiceStageClient) ListTemplates(request *model.ListTemplatesRequest) (*model.ListTemplatesResponse, error) {
	requestDef := GenReqDefForListTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplatesResponse), nil
	}
}

// ListTemplatesInvoker 获取所有支持的应用组件模板
func (c *ServiceStageClient) ListTemplatesInvoker(request *model.ListTemplatesRequest) *ListTemplatesInvoker {
	requestDef := GenReqDefForListTemplates()
	return &ListTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
