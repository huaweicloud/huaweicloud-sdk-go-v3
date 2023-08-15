package v3

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/iotedge/v3/model"
)

type CreateAppInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowClusterInvoker) Invoke() (*model.ShowClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClusterResponse), nil
	}
}
