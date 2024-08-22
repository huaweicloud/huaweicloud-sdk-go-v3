package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cae/v1/model"
)

type CaeClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCaeClient(hcClient *httpclient.HcHttpClient) *CaeClient {
	return &CaeClient{HcClient: hcClient}
}

func CaeClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateAgency 创建委托
//
// 创建委托。
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

// ListAgencies 获取委托列表
//
// 获取委托列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) ListAgencies(request *model.ListAgenciesRequest) (*model.ListAgenciesResponse, error) {
	requestDef := GenReqDefForListAgencies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAgenciesResponse), nil
	}
}

// ListAgenciesInvoker 获取委托列表
func (c *CaeClient) ListAgenciesInvoker(request *model.ListAgenciesRequest) *ListAgenciesInvoker {
	requestDef := GenReqDefForListAgencies()
	return &ListAgenciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateApplication 创建应用
//
// 创建应用。
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
// 删除应用。
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
// 获取应用列表。
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

// ShowApplication 获取应用详情
//
// 获取应用详情。
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

// ShowApplicationInvoker 获取应用详情
func (c *CaeClient) ShowApplicationInvoker(request *model.ShowApplicationRequest) *ShowApplicationInvoker {
	requestDef := GenReqDefForShowApplication()
	return &ShowApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCertificate 创建证书
//
// 创建证书。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) CreateCertificate(request *model.CreateCertificateRequest) (*model.CreateCertificateResponse, error) {
	requestDef := GenReqDefForCreateCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCertificateResponse), nil
	}
}

// CreateCertificateInvoker 创建证书
func (c *CaeClient) CreateCertificateInvoker(request *model.CreateCertificateRequest) *CreateCertificateInvoker {
	requestDef := GenReqDefForCreateCertificate()
	return &CreateCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCertificate 删除证书
//
// 删除证书。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) DeleteCertificate(request *model.DeleteCertificateRequest) (*model.DeleteCertificateResponse, error) {
	requestDef := GenReqDefForDeleteCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCertificateResponse), nil
	}
}

// DeleteCertificateInvoker 删除证书
func (c *CaeClient) DeleteCertificateInvoker(request *model.DeleteCertificateRequest) *DeleteCertificateInvoker {
	requestDef := GenReqDefForDeleteCertificate()
	return &DeleteCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCertificates 获取证书列表
//
// 获取证书列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) ListCertificates(request *model.ListCertificatesRequest) (*model.ListCertificatesResponse, error) {
	requestDef := GenReqDefForListCertificates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCertificatesResponse), nil
	}
}

// ListCertificatesInvoker 获取证书列表
func (c *CaeClient) ListCertificatesInvoker(request *model.ListCertificatesRequest) *ListCertificatesInvoker {
	requestDef := GenReqDefForListCertificates()
	return &ListCertificatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCertificate 修改证书
//
// 修改证书。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) UpdateCertificate(request *model.UpdateCertificateRequest) (*model.UpdateCertificateResponse, error) {
	requestDef := GenReqDefForUpdateCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCertificateResponse), nil
	}
}

// UpdateCertificateInvoker 修改证书
func (c *CaeClient) UpdateCertificateInvoker(request *model.UpdateCertificateRequest) *UpdateCertificateInvoker {
	requestDef := GenReqDefForUpdateCertificate()
	return &UpdateCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateComponent 创建组件
//
// 创建组件。
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

// CreateComponentWithConfiguration 创建、生效配置并部署组件
//
// 创建、生效配置并部署组件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) CreateComponentWithConfiguration(request *model.CreateComponentWithConfigurationRequest) (*model.CreateComponentWithConfigurationResponse, error) {
	requestDef := GenReqDefForCreateComponentWithConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateComponentWithConfigurationResponse), nil
	}
}

// CreateComponentWithConfigurationInvoker 创建、生效配置并部署组件
func (c *CaeClient) CreateComponentWithConfigurationInvoker(request *model.CreateComponentWithConfigurationRequest) *CreateComponentWithConfigurationInvoker {
	requestDef := GenReqDefForCreateComponentWithConfiguration()
	return &CreateComponentWithConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteComponent 删除组件
//
// 删除组件。
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
// 对组件执行指定操作，如部署、升级、重启、停止、启动、伸缩、配置、回滚。
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

// ListComponentInstances 获取组件实例列表
//
// 获取组件实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) ListComponentInstances(request *model.ListComponentInstancesRequest) (*model.ListComponentInstancesResponse, error) {
	requestDef := GenReqDefForListComponentInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListComponentInstancesResponse), nil
	}
}

// ListComponentInstancesInvoker 获取组件实例列表
func (c *CaeClient) ListComponentInstancesInvoker(request *model.ListComponentInstancesRequest) *ListComponentInstancesInvoker {
	requestDef := GenReqDefForListComponentInstances()
	return &ListComponentInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListComponentSnapshots 获取组件快照列表
//
// 获取组件快照列表。
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
// 获取组件列表。
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

// ShowComponent 获取组件详情
//
// 获取组件详情。
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

// ShowComponentInvoker 获取组件详情
func (c *CaeClient) ShowComponentInvoker(request *model.ShowComponentRequest) *ShowComponentInvoker {
	requestDef := GenReqDefForShowComponent()
	return &ShowComponentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateComponent 更新组件
//
// 更新组件。
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
// 创建组件配置。
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
// 删除组件配置。
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

// ListComponentConfigurations 获取组件配置列表
//
// 获取组件配置列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) ListComponentConfigurations(request *model.ListComponentConfigurationsRequest) (*model.ListComponentConfigurationsResponse, error) {
	requestDef := GenReqDefForListComponentConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListComponentConfigurationsResponse), nil
	}
}

// ListComponentConfigurationsInvoker 获取组件配置列表
func (c *CaeClient) ListComponentConfigurationsInvoker(request *model.ListComponentConfigurationsRequest) *ListComponentConfigurationsInvoker {
	requestDef := GenReqDefForListComponentConfigurations()
	return &ListComponentConfigurationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDomain 创建域名
//
// 创建域名。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) CreateDomain(request *model.CreateDomainRequest) (*model.CreateDomainResponse, error) {
	requestDef := GenReqDefForCreateDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDomainResponse), nil
	}
}

// CreateDomainInvoker 创建域名
func (c *CaeClient) CreateDomainInvoker(request *model.CreateDomainRequest) *CreateDomainInvoker {
	requestDef := GenReqDefForCreateDomain()
	return &CreateDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDomain 删除域名
//
// 删除域名。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) DeleteDomain(request *model.DeleteDomainRequest) (*model.DeleteDomainResponse, error) {
	requestDef := GenReqDefForDeleteDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDomainResponse), nil
	}
}

// DeleteDomainInvoker 删除域名
func (c *CaeClient) DeleteDomainInvoker(request *model.DeleteDomainRequest) *DeleteDomainInvoker {
	requestDef := GenReqDefForDeleteDomain()
	return &DeleteDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDomains 获取域名列表
//
// 获取域名列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) ListDomains(request *model.ListDomainsRequest) (*model.ListDomainsResponse, error) {
	requestDef := GenReqDefForListDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainsResponse), nil
	}
}

// ListDomainsInvoker 获取域名列表
func (c *CaeClient) ListDomainsInvoker(request *model.ListDomainsRequest) *ListDomainsInvoker {
	requestDef := GenReqDefForListDomains()
	return &ListDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEips 获取集群节点弹性公网IP列表
//
// 获取集群节点弹性公网IP列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) ListEips(request *model.ListEipsRequest) (*model.ListEipsResponse, error) {
	requestDef := GenReqDefForListEips()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEipsResponse), nil
	}
}

// ListEipsInvoker 获取集群节点弹性公网IP列表
func (c *CaeClient) ListEipsInvoker(request *model.ListEipsRequest) *ListEipsInvoker {
	requestDef := GenReqDefForListEips()
	return &ListEipsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEip 修改出入网带宽以及开闭状态
//
// 修改出入网带宽以及开闭状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) UpdateEip(request *model.UpdateEipRequest) (*model.UpdateEipResponse, error) {
	requestDef := GenReqDefForUpdateEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEipResponse), nil
	}
}

// UpdateEipInvoker 修改出入网带宽以及开闭状态
func (c *CaeClient) UpdateEipInvoker(request *model.UpdateEipRequest) *UpdateEipInvoker {
	requestDef := GenReqDefForUpdateEip()
	return &UpdateEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEnvironment 创建环境
//
// 创建环境。
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
// 删除环境。
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
// 获取环境列表。
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
// 重试任务。
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
// 获取任务详情。
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

// CreateMonitorSystem 创建监控系统配置
//
// 创建监控系统配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) CreateMonitorSystem(request *model.CreateMonitorSystemRequest) (*model.CreateMonitorSystemResponse, error) {
	requestDef := GenReqDefForCreateMonitorSystem()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMonitorSystemResponse), nil
	}
}

// CreateMonitorSystemInvoker 创建监控系统配置
func (c *CaeClient) CreateMonitorSystemInvoker(request *model.CreateMonitorSystemRequest) *CreateMonitorSystemInvoker {
	requestDef := GenReqDefForCreateMonitorSystem()
	return &CreateMonitorSystemInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMonitorSystem 获取监控系统配置
//
// 获取监控系统配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) ShowMonitorSystem(request *model.ShowMonitorSystemRequest) (*model.ShowMonitorSystemResponse, error) {
	requestDef := GenReqDefForShowMonitorSystem()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMonitorSystemResponse), nil
	}
}

// ShowMonitorSystemInvoker 获取监控系统配置
func (c *CaeClient) ShowMonitorSystemInvoker(request *model.ShowMonitorSystemRequest) *ShowMonitorSystemInvoker {
	requestDef := GenReqDefForShowMonitorSystem()
	return &ShowMonitorSystemInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMonitorSystem 更新监控系统配置
//
// 更新监控系统配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) UpdateMonitorSystem(request *model.UpdateMonitorSystemRequest) (*model.UpdateMonitorSystemResponse, error) {
	requestDef := GenReqDefForUpdateMonitorSystem()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMonitorSystemResponse), nil
	}
}

// UpdateMonitorSystemInvoker 更新监控系统配置
func (c *CaeClient) UpdateMonitorSystemInvoker(request *model.UpdateMonitorSystemRequest) *UpdateMonitorSystemInvoker {
	requestDef := GenReqDefForUpdateMonitorSystem()
	return &UpdateMonitorSystemInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateNoticeRule 创建事件通知规则。
//
// 创建事件通知规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) CreateNoticeRule(request *model.CreateNoticeRuleRequest) (*model.CreateNoticeRuleResponse, error) {
	requestDef := GenReqDefForCreateNoticeRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNoticeRuleResponse), nil
	}
}

// CreateNoticeRuleInvoker 创建事件通知规则。
func (c *CaeClient) CreateNoticeRuleInvoker(request *model.CreateNoticeRuleRequest) *CreateNoticeRuleInvoker {
	requestDef := GenReqDefForCreateNoticeRule()
	return &CreateNoticeRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteNoticeRule 删除事件通知规则。
//
// 删除事件通知规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) DeleteNoticeRule(request *model.DeleteNoticeRuleRequest) (*model.DeleteNoticeRuleResponse, error) {
	requestDef := GenReqDefForDeleteNoticeRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNoticeRuleResponse), nil
	}
}

// DeleteNoticeRuleInvoker 删除事件通知规则。
func (c *CaeClient) DeleteNoticeRuleInvoker(request *model.DeleteNoticeRuleRequest) *DeleteNoticeRuleInvoker {
	requestDef := GenReqDefForDeleteNoticeRule()
	return &DeleteNoticeRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNoticeRules 查询事件通知规则列表。
//
// 查询事件通知规则列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) ListNoticeRules(request *model.ListNoticeRulesRequest) (*model.ListNoticeRulesResponse, error) {
	requestDef := GenReqDefForListNoticeRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNoticeRulesResponse), nil
	}
}

// ListNoticeRulesInvoker 查询事件通知规则列表。
func (c *CaeClient) ListNoticeRulesInvoker(request *model.ListNoticeRulesRequest) *ListNoticeRulesInvoker {
	requestDef := GenReqDefForListNoticeRules()
	return &ListNoticeRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNoticeRule 查询事件通知规则。
//
// 查询事件通知规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) ShowNoticeRule(request *model.ShowNoticeRuleRequest) (*model.ShowNoticeRuleResponse, error) {
	requestDef := GenReqDefForShowNoticeRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNoticeRuleResponse), nil
	}
}

// ShowNoticeRuleInvoker 查询事件通知规则。
func (c *CaeClient) ShowNoticeRuleInvoker(request *model.ShowNoticeRuleRequest) *ShowNoticeRuleInvoker {
	requestDef := GenReqDefForShowNoticeRule()
	return &ShowNoticeRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNoticeRule 修改事件通知规则。
//
// 修改事件通知规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) UpdateNoticeRule(request *model.UpdateNoticeRuleRequest) (*model.UpdateNoticeRuleResponse, error) {
	requestDef := GenReqDefForUpdateNoticeRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNoticeRuleResponse), nil
	}
}

// UpdateNoticeRuleInvoker 修改事件通知规则。
func (c *CaeClient) UpdateNoticeRuleInvoker(request *model.UpdateNoticeRuleRequest) *UpdateNoticeRuleInvoker {
	requestDef := GenReqDefForUpdateNoticeRule()
	return &UpdateNoticeRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTimerRule 创建定时启停规则
//
// 创建定时启停规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) CreateTimerRule(request *model.CreateTimerRuleRequest) (*model.CreateTimerRuleResponse, error) {
	requestDef := GenReqDefForCreateTimerRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTimerRuleResponse), nil
	}
}

// CreateTimerRuleInvoker 创建定时启停规则
func (c *CaeClient) CreateTimerRuleInvoker(request *model.CreateTimerRuleRequest) *CreateTimerRuleInvoker {
	requestDef := GenReqDefForCreateTimerRule()
	return &CreateTimerRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTimerRule 删除定时启停规则
//
// 删除定时启停规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) DeleteTimerRule(request *model.DeleteTimerRuleRequest) (*model.DeleteTimerRuleResponse, error) {
	requestDef := GenReqDefForDeleteTimerRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTimerRuleResponse), nil
	}
}

// DeleteTimerRuleInvoker 删除定时启停规则
func (c *CaeClient) DeleteTimerRuleInvoker(request *model.DeleteTimerRuleRequest) *DeleteTimerRuleInvoker {
	requestDef := GenReqDefForDeleteTimerRule()
	return &DeleteTimerRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTimerRules 获取定时启停规则列表
//
// 获取定时启停规则列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) ListTimerRules(request *model.ListTimerRulesRequest) (*model.ListTimerRulesResponse, error) {
	requestDef := GenReqDefForListTimerRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTimerRulesResponse), nil
	}
}

// ListTimerRulesInvoker 获取定时启停规则列表
func (c *CaeClient) ListTimerRulesInvoker(request *model.ListTimerRulesRequest) *ListTimerRulesInvoker {
	requestDef := GenReqDefForListTimerRules()
	return &ListTimerRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowExecutionResult 获取上次定时启停规则的执行情况
//
// 获取上次定时启停规则的执行情况。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) ShowExecutionResult(request *model.ShowExecutionResultRequest) (*model.ShowExecutionResultResponse, error) {
	requestDef := GenReqDefForShowExecutionResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowExecutionResultResponse), nil
	}
}

// ShowExecutionResultInvoker 获取上次定时启停规则的执行情况
func (c *CaeClient) ShowExecutionResultInvoker(request *model.ShowExecutionResultRequest) *ShowExecutionResultInvoker {
	requestDef := GenReqDefForShowExecutionResult()
	return &ShowExecutionResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTimerRule 修改定时启停规则
//
// 修改定时启停规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) UpdateTimerRule(request *model.UpdateTimerRuleRequest) (*model.UpdateTimerRuleResponse, error) {
	requestDef := GenReqDefForUpdateTimerRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTimerRuleResponse), nil
	}
}

// UpdateTimerRuleInvoker 修改定时启停规则
func (c *CaeClient) UpdateTimerRuleInvoker(request *model.UpdateTimerRuleRequest) *UpdateTimerRuleInvoker {
	requestDef := GenReqDefForUpdateTimerRule()
	return &UpdateTimerRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVolume 授权云存储
//
// 授权云存储。
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

// CreateVolumeInvoker 授权云存储
func (c *CaeClient) CreateVolumeInvoker(request *model.CreateVolumeRequest) *CreateVolumeInvoker {
	requestDef := GenReqDefForCreateVolume()
	return &CreateVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVolume 解绑云存储
//
// 解绑云存储。
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

// DeleteVolumeInvoker 解绑云存储
func (c *CaeClient) DeleteVolumeInvoker(request *model.DeleteVolumeRequest) *DeleteVolumeInvoker {
	requestDef := GenReqDefForDeleteVolume()
	return &DeleteVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVolumes 获取云存储列表
//
// 获取云存储列表。
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

// ListVolumesInvoker 获取云存储列表
func (c *CaeClient) ListVolumesInvoker(request *model.ListVolumesRequest) *ListVolumesInvoker {
	requestDef := GenReqDefForListVolumes()
	return &ListVolumesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVpcEgress 创建CAE环境访问VPC配置
//
// 创建CAE环境访问VPC配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) CreateVpcEgress(request *model.CreateVpcEgressRequest) (*model.CreateVpcEgressResponse, error) {
	requestDef := GenReqDefForCreateVpcEgress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVpcEgressResponse), nil
	}
}

// CreateVpcEgressInvoker 创建CAE环境访问VPC配置
func (c *CaeClient) CreateVpcEgressInvoker(request *model.CreateVpcEgressRequest) *CreateVpcEgressInvoker {
	requestDef := GenReqDefForCreateVpcEgress()
	return &CreateVpcEgressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVpcEgress 删除CAE环境访问VPC配置
//
// 删除CAE环境访问VPC配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) DeleteVpcEgress(request *model.DeleteVpcEgressRequest) (*model.DeleteVpcEgressResponse, error) {
	requestDef := GenReqDefForDeleteVpcEgress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVpcEgressResponse), nil
	}
}

// DeleteVpcEgressInvoker 删除CAE环境访问VPC配置
func (c *CaeClient) DeleteVpcEgressInvoker(request *model.DeleteVpcEgressRequest) *DeleteVpcEgressInvoker {
	requestDef := GenReqDefForDeleteVpcEgress()
	return &DeleteVpcEgressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVpcEgress 获取CAE环境访问VPC配置
//
// 获取CAE环境访问VPC配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CaeClient) ListVpcEgress(request *model.ListVpcEgressRequest) (*model.ListVpcEgressResponse, error) {
	requestDef := GenReqDefForListVpcEgress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpcEgressResponse), nil
	}
}

// ListVpcEgressInvoker 获取CAE环境访问VPC配置
func (c *CaeClient) ListVpcEgressInvoker(request *model.ListVpcEgressRequest) *ListVpcEgressInvoker {
	requestDef := GenReqDefForListVpcEgress()
	return &ListVpcEgressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
