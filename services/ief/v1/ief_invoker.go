package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ief/v1/model"
)

type BatchAddDeleteTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAddDeleteTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchAddDeleteTagsInvoker) Invoke() (*model.BatchAddDeleteTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAddDeleteTagsResponse), nil
	}
}

type CreateAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAppInvoker) Invoke() (*model.CreateAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAppResponse), nil
	}
}

type CreateAppVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAppVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAppVersionsInvoker) Invoke() (*model.CreateAppVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAppVersionsResponse), nil
	}
}

type CreateBatchJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBatchJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateBatchJobInvoker) Invoke() (*model.CreateBatchJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBatchJobResponse), nil
	}
}

type CreateConfigMapInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateConfigMapInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateConfigMapInvoker) Invoke() (*model.CreateConfigMapResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConfigMapResponse), nil
	}
}

type CreateDeploymentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDeploymentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDeploymentsInvoker) Invoke() (*model.CreateDeploymentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDeploymentsResponse), nil
	}
}

type CreateDeviceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDeviceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDeviceInvoker) Invoke() (*model.CreateDeviceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDeviceResponse), nil
	}
}

type CreateDeviceTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDeviceTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDeviceTemplateInvoker) Invoke() (*model.CreateDeviceTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDeviceTemplateResponse), nil
	}
}

type CreateEdgeGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEdgeGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEdgeGroupInvoker) Invoke() (*model.CreateEdgeGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEdgeGroupResponse), nil
	}
}

type CreateEdgeGroupCertInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEdgeGroupCertInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEdgeGroupCertInvoker) Invoke() (*model.CreateEdgeGroupCertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEdgeGroupCertResponse), nil
	}
}

type CreateEdgeNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEdgeNodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEdgeNodeInvoker) Invoke() (*model.CreateEdgeNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEdgeNodeResponse), nil
	}
}

type CreateEdgeNodeCertsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEdgeNodeCertsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEdgeNodeCertsInvoker) Invoke() (*model.CreateEdgeNodeCertsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEdgeNodeCertsResponse), nil
	}
}

type CreateEncryptdatasInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEncryptdatasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEncryptdatasInvoker) Invoke() (*model.CreateEncryptdatasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEncryptdatasResponse), nil
	}
}

type CreateEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEndpointInvoker) Invoke() (*model.CreateEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEndpointResponse), nil
	}
}

type CreateNodeEncryptdatasInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNodeEncryptdatasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateNodeEncryptdatasInvoker) Invoke() (*model.CreateNodeEncryptdatasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNodeEncryptdatasResponse), nil
	}
}

type CreateProductInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateProductInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateProductInvoker) Invoke() (*model.CreateProductResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateProductResponse), nil
	}
}

type CreateRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRuleInvoker) Invoke() (*model.CreateRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRuleResponse), nil
	}
}

type CreateSecretInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSecretInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSecretInvoker) Invoke() (*model.CreateSecretResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSecretResponse), nil
	}
}

type CreateServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateServiceInvoker) Invoke() (*model.CreateServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateServiceResponse), nil
	}
}

type CreateSystemEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSystemEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSystemEventInvoker) Invoke() (*model.CreateSystemEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSystemEventResponse), nil
	}
}

type CreateTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTagInvoker) Invoke() (*model.CreateTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTagResponse), nil
	}
}

type DeleteAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAppInvoker) Invoke() (*model.DeleteAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAppResponse), nil
	}
}

type DeleteAppVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAppVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAppVersionInvoker) Invoke() (*model.DeleteAppVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAppVersionResponse), nil
	}
}

type DeleteBatchJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBatchJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteBatchJobInvoker) Invoke() (*model.DeleteBatchJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBatchJobResponse), nil
	}
}

type DeleteConfigMapInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteConfigMapInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteConfigMapInvoker) Invoke() (*model.DeleteConfigMapResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteConfigMapResponse), nil
	}
}

type DeleteDeploymentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDeploymentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDeploymentInvoker) Invoke() (*model.DeleteDeploymentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDeploymentResponse), nil
	}
}

type DeleteDeviceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDeviceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDeviceInvoker) Invoke() (*model.DeleteDeviceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDeviceResponse), nil
	}
}

type DeleteDeviceTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDeviceTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDeviceTemplateInvoker) Invoke() (*model.DeleteDeviceTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDeviceTemplateResponse), nil
	}
}

type DeleteEdgeGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEdgeGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEdgeGroupInvoker) Invoke() (*model.DeleteEdgeGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEdgeGroupResponse), nil
	}
}

type DeleteEdgeGroupCertInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEdgeGroupCertInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEdgeGroupCertInvoker) Invoke() (*model.DeleteEdgeGroupCertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEdgeGroupCertResponse), nil
	}
}

type DeleteEdgeNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEdgeNodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEdgeNodeInvoker) Invoke() (*model.DeleteEdgeNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEdgeNodeResponse), nil
	}
}

type DeleteEdgeNodeCertsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEdgeNodeCertsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEdgeNodeCertsInvoker) Invoke() (*model.DeleteEdgeNodeCertsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEdgeNodeCertsResponse), nil
	}
}

type DeleteEncryptdatasInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEncryptdatasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEncryptdatasInvoker) Invoke() (*model.DeleteEncryptdatasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEncryptdatasResponse), nil
	}
}

type DeleteEndPointInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEndPointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEndPointInvoker) Invoke() (*model.DeleteEndPointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEndPointResponse), nil
	}
}

type DeleteNodeEncryptdatasInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNodeEncryptdatasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteNodeEncryptdatasInvoker) Invoke() (*model.DeleteNodeEncryptdatasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNodeEncryptdatasResponse), nil
	}
}

type DeleteProductInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteProductInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteProductInvoker) Invoke() (*model.DeleteProductResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteProductResponse), nil
	}
}

type DeleteResourceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteResourceTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteResourceTagInvoker) Invoke() (*model.DeleteResourceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteResourceTagResponse), nil
	}
}

type DeleteRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRuleInvoker) Invoke() (*model.DeleteRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRuleResponse), nil
	}
}

type DeleteSecretInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSecretInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSecretInvoker) Invoke() (*model.DeleteSecretResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSecretResponse), nil
	}
}

type DeleteServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteServiceInvoker) Invoke() (*model.DeleteServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteServiceResponse), nil
	}
}

type DeleteSystemEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSystemEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSystemEventInvoker) Invoke() (*model.DeleteSystemEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSystemEventResponse), nil
	}
}

type EnableDisableEdgeNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableDisableEdgeNodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableDisableEdgeNodesInvoker) Invoke() (*model.EnableDisableEdgeNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableDisableEdgeNodesResponse), nil
	}
}

type ListAppVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppVersionsInvoker) Invoke() (*model.ListAppVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppVersionsResponse), nil
	}
}

type ListAppsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppsInvoker) Invoke() (*model.ListAppsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppsResponse), nil
	}
}

type ListBatchJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBatchJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBatchJobInvoker) Invoke() (*model.ListBatchJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBatchJobResponse), nil
	}
}

type ListConfigMapsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConfigMapsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListConfigMapsInvoker) Invoke() (*model.ListConfigMapsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConfigMapsResponse), nil
	}
}

type ListDeploymentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDeploymentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDeploymentsInvoker) Invoke() (*model.ListDeploymentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDeploymentsResponse), nil
	}
}

type ListDeviceTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDeviceTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDeviceTemplatesInvoker) Invoke() (*model.ListDeviceTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDeviceTemplatesResponse), nil
	}
}

type ListDevicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDevicesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDevicesInvoker) Invoke() (*model.ListDevicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDevicesResponse), nil
	}
}

type ListEdgeGroupCertsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEdgeGroupCertsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEdgeGroupCertsInvoker) Invoke() (*model.ListEdgeGroupCertsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEdgeGroupCertsResponse), nil
	}
}

type ListEdgeGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEdgeGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEdgeGroupsInvoker) Invoke() (*model.ListEdgeGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEdgeGroupsResponse), nil
	}
}

type ListEdgeNodeCertsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEdgeNodeCertsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEdgeNodeCertsInvoker) Invoke() (*model.ListEdgeNodeCertsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEdgeNodeCertsResponse), nil
	}
}

type ListEdgeNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEdgeNodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEdgeNodesInvoker) Invoke() (*model.ListEdgeNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEdgeNodesResponse), nil
	}
}

type ListEncryptdataNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEncryptdataNodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEncryptdataNodesInvoker) Invoke() (*model.ListEncryptdataNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEncryptdataNodesResponse), nil
	}
}

type ListEncryptdatasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEncryptdatasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEncryptdatasInvoker) Invoke() (*model.ListEncryptdatasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEncryptdatasResponse), nil
	}
}

type ListEndpointsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEndpointsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEndpointsInvoker) Invoke() (*model.ListEndpointsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEndpointsResponse), nil
	}
}

type ListNodeEncryptdatasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNodeEncryptdatasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNodeEncryptdatasInvoker) Invoke() (*model.ListNodeEncryptdatasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNodeEncryptdatasResponse), nil
	}
}

type ListPodsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPodsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPodsInvoker) Invoke() (*model.ListPodsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPodsResponse), nil
	}
}

type ListProductsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProductsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProductsInvoker) Invoke() (*model.ListProductsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProductsResponse), nil
	}
}

type ListResourceByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceByTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceByTagsInvoker) Invoke() (*model.ListResourceByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceByTagsResponse), nil
	}
}

type ListRuleErrorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRuleErrorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRuleErrorsInvoker) Invoke() (*model.ListRuleErrorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRuleErrorsResponse), nil
	}
}

type ListRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRulesInvoker) Invoke() (*model.ListRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRulesResponse), nil
	}
}

type ListSecretsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecretsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecretsInvoker) Invoke() (*model.ListSecretsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecretsResponse), nil
	}
}

type ListServicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServicesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServicesInvoker) Invoke() (*model.ListServicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServicesResponse), nil
	}
}

type ListSystemEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSystemEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSystemEventsInvoker) Invoke() (*model.ListSystemEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSystemEventsResponse), nil
	}
}

type ListTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTagsInvoker) Invoke() (*model.ListTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagsResponse), nil
	}
}

type ListTagsByResourceTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagsByResourceTypeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTagsByResourceTypeInvoker) Invoke() (*model.ListTagsByResourceTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagsByResourceTypeResponse), nil
	}
}

type RestartDeploymentsPodInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestartDeploymentsPodInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RestartDeploymentsPodInvoker) Invoke() (*model.RestartDeploymentsPodResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartDeploymentsPodResponse), nil
	}
}

type RestoreBatchJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreBatchJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RestoreBatchJobInvoker) Invoke() (*model.RestoreBatchJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreBatchJobResponse), nil
	}
}

type RetryBatchJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RetryBatchJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RetryBatchJobInvoker) Invoke() (*model.RetryBatchJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetryBatchJobResponse), nil
	}
}

type ShowAppDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAppDetailInvoker) Invoke() (*model.ShowAppDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppDetailResponse), nil
	}
}

type ShowAppVersionDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppVersionDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAppVersionDetailInvoker) Invoke() (*model.ShowAppVersionDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppVersionDetailResponse), nil
	}
}

type ShowBatchJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBatchJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBatchJobInvoker) Invoke() (*model.ShowBatchJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBatchJobResponse), nil
	}
}

type ShowConfigMapInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConfigMapInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowConfigMapInvoker) Invoke() (*model.ShowConfigMapResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConfigMapResponse), nil
	}
}

type ShowDeploymentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeploymentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDeploymentInvoker) Invoke() (*model.ShowDeploymentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeploymentResponse), nil
	}
}

type ShowDeviceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeviceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDeviceInvoker) Invoke() (*model.ShowDeviceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeviceResponse), nil
	}
}

type ShowDeviceTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeviceTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDeviceTemplateInvoker) Invoke() (*model.ShowDeviceTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeviceTemplateResponse), nil
	}
}

type ShowDeviceTwinInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeviceTwinInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDeviceTwinInvoker) Invoke() (*model.ShowDeviceTwinResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeviceTwinResponse), nil
	}
}

type ShowEdgeGroupCertDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEdgeGroupCertDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEdgeGroupCertDetailInvoker) Invoke() (*model.ShowEdgeGroupCertDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEdgeGroupCertDetailResponse), nil
	}
}

type ShowEdgeGroupDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEdgeGroupDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEdgeGroupDetailInvoker) Invoke() (*model.ShowEdgeGroupDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEdgeGroupDetailResponse), nil
	}
}

type ShowEdgeNodeDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEdgeNodeDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEdgeNodeDetailInvoker) Invoke() (*model.ShowEdgeNodeDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEdgeNodeDetailResponse), nil
	}
}

type ShowEncryptdatasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEncryptdatasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEncryptdatasInvoker) Invoke() (*model.ShowEncryptdatasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEncryptdatasResponse), nil
	}
}

type ShowEndPointDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEndPointDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEndPointDetailInvoker) Invoke() (*model.ShowEndPointDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEndPointDetailResponse), nil
	}
}

type ShowProductDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProductDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProductDetailInvoker) Invoke() (*model.ShowProductDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProductDetailResponse), nil
	}
}

type ShowQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowQuotaInvoker) Invoke() (*model.ShowQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuotaResponse), nil
	}
}

type ShowRuleDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRuleDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRuleDetailInvoker) Invoke() (*model.ShowRuleDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRuleDetailResponse), nil
	}
}

type ShowSecretInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSecretInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSecretInvoker) Invoke() (*model.ShowSecretResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSecretResponse), nil
	}
}

type ShowServiceDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServiceDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowServiceDetailInvoker) Invoke() (*model.ShowServiceDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServiceDetailResponse), nil
	}
}

type ShowSystemEventDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSystemEventDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSystemEventDetailInvoker) Invoke() (*model.ShowSystemEventDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSystemEventDetailResponse), nil
	}
}

type StartRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartRuleInvoker) Invoke() (*model.StartRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartRuleResponse), nil
	}
}

type StartSystemEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartSystemEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartSystemEventInvoker) Invoke() (*model.StartSystemEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartSystemEventResponse), nil
	}
}

type StopBatchJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopBatchJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopBatchJobInvoker) Invoke() (*model.StopBatchJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopBatchJobResponse), nil
	}
}

type StopRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopRuleInvoker) Invoke() (*model.StopRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopRuleResponse), nil
	}
}

type StopSystemEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopSystemEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopSystemEventInvoker) Invoke() (*model.StopSystemEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopSystemEventResponse), nil
	}
}

type UpdateAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAppInvoker) Invoke() (*model.UpdateAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAppResponse), nil
	}
}

type UpdateAppVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAppVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAppVersionInvoker) Invoke() (*model.UpdateAppVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAppVersionResponse), nil
	}
}

type UpdateConfigMapInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateConfigMapInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateConfigMapInvoker) Invoke() (*model.UpdateConfigMapResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateConfigMapResponse), nil
	}
}

type UpdateDeploymentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDeploymentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDeploymentInvoker) Invoke() (*model.UpdateDeploymentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDeploymentResponse), nil
	}
}

type UpdateDeviceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDeviceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDeviceInvoker) Invoke() (*model.UpdateDeviceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDeviceResponse), nil
	}
}

type UpdateDeviceTemplateByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDeviceTemplateByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDeviceTemplateByIdInvoker) Invoke() (*model.UpdateDeviceTemplateByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDeviceTemplateByIdResponse), nil
	}
}

type UpdateDeviceTwinInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDeviceTwinInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDeviceTwinInvoker) Invoke() (*model.UpdateDeviceTwinResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDeviceTwinResponse), nil
	}
}

type UpdateEdgeGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEdgeGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEdgeGroupInvoker) Invoke() (*model.UpdateEdgeGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEdgeGroupResponse), nil
	}
}

type UpdateEdgeGroupNodeBindingInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEdgeGroupNodeBindingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEdgeGroupNodeBindingInvoker) Invoke() (*model.UpdateEdgeGroupNodeBindingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEdgeGroupNodeBindingResponse), nil
	}
}

type UpdateEdgeNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEdgeNodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEdgeNodeInvoker) Invoke() (*model.UpdateEdgeNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEdgeNodeResponse), nil
	}
}

type UpdateEdgeNodeDeviceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEdgeNodeDeviceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEdgeNodeDeviceInvoker) Invoke() (*model.UpdateEdgeNodeDeviceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEdgeNodeDeviceResponse), nil
	}
}

type UpdateEncryptdatasInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEncryptdatasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEncryptdatasInvoker) Invoke() (*model.UpdateEncryptdatasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEncryptdatasResponse), nil
	}
}

type UpdateNodeByDeviceIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNodeByDeviceIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateNodeByDeviceIdInvoker) Invoke() (*model.UpdateNodeByDeviceIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNodeByDeviceIdResponse), nil
	}
}

type UpdateSecretInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSecretInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSecretInvoker) Invoke() (*model.UpdateSecretResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSecretResponse), nil
	}
}

type UpdateServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateServiceInvoker) Invoke() (*model.UpdateServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateServiceResponse), nil
	}
}

type UpgradeEdgeNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpgradeEdgeNodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpgradeEdgeNodeInvoker) Invoke() (*model.UpgradeEdgeNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpgradeEdgeNodeResponse), nil
	}
}
