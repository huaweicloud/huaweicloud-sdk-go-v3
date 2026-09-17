package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iotedge/v3/model"
)

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

type ShowAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAppInvoker) Invoke() (*model.ShowAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppResponse), nil
	}
}

type CreateAppInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAppInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAppInstanceInvoker) Invoke() (*model.CreateAppInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAppInstanceResponse), nil
	}
}

type DeleteAppInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAppInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAppInstanceInvoker) Invoke() (*model.DeleteAppInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAppInstanceResponse), nil
	}
}

type ListAppInstanceHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppInstanceHistoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppInstanceHistoryInvoker) Invoke() (*model.ListAppInstanceHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppInstanceHistoryResponse), nil
	}
}

type ListAppInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppInstancesInvoker) Invoke() (*model.ListAppInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppInstancesResponse), nil
	}
}

type UpdateAppInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAppInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAppInstanceInvoker) Invoke() (*model.UpdateAppInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAppInstanceResponse), nil
	}
}

type CreateAppVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAppVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAppVersionInvoker) Invoke() (*model.CreateAppVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAppVersionResponse), nil
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

type DownloadAppVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadAppVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadAppVersionInvoker) Invoke() (*model.DownloadAppVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadAppVersionResponse), nil
	}
}

type ListAppImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppImageInvoker) Invoke() (*model.ListAppImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppImageResponse), nil
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

type ShowAppVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAppVersionInvoker) Invoke() (*model.ShowAppVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppVersionResponse), nil
	}
}

type CreateClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateClusterInvoker) Invoke() (*model.CreateClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClusterResponse), nil
	}
}

type CreateClusterInstallCmdInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClusterInstallCmdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateClusterInstallCmdInvoker) Invoke() (*model.CreateClusterInstallCmdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClusterInstallCmdResponse), nil
	}
}

type DeleteClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteClusterInvoker) Invoke() (*model.DeleteClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteClusterResponse), nil
	}
}

type ListClustersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClustersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClustersInvoker) Invoke() (*model.ListClustersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClustersResponse), nil
	}
}

type ShowClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowClusterInvoker) Invoke() (*model.ShowClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClusterResponse), nil
	}
}

type ShowClusterResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClusterResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowClusterResourcesInvoker) Invoke() (*model.ShowClusterResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClusterResourcesResponse), nil
	}
}

type UpdateResourceBindingInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateResourceBindingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateResourceBindingInvoker) Invoke() (*model.UpdateResourceBindingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateResourceBindingResponse), nil
	}
}

type InvokeKubeApiInvoker struct {
	*invoker.BaseInvoker
}

func (i *InvokeKubeApiInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *InvokeKubeApiInvoker) Invoke() (*model.InvokeKubeApiResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.InvokeKubeApiResponse), nil
	}
}

type ListClusterNamespacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterNamespacesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClusterNamespacesInvoker) Invoke() (*model.ListClusterNamespacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterNamespacesResponse), nil
	}
}

type CreateClusterNodesInstallCmdInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClusterNodesInstallCmdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateClusterNodesInstallCmdInvoker) Invoke() (*model.CreateClusterNodesInstallCmdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClusterNodesInstallCmdResponse), nil
	}
}

type ListClusterNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterNodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClusterNodesInvoker) Invoke() (*model.ListClusterNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterNodesResponse), nil
	}
}

type ShowClusterNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClusterNodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowClusterNodeInvoker) Invoke() (*model.ShowClusterNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClusterNodeResponse), nil
	}
}
