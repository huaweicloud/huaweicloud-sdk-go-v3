package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cce/v3/model"
)

type AddNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddNodeInvoker) Invoke() (*model.AddNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddNodeResponse), nil
	}
}

type AwakeClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *AwakeClusterInvoker) Invoke() (*model.AwakeClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AwakeClusterResponse), nil
	}
}

type CreateAddonInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAddonInstanceInvoker) Invoke() (*model.CreateAddonInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAddonInstanceResponse), nil
	}
}

type CreateCloudPersistentVolumeClaimsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCloudPersistentVolumeClaimsInvoker) Invoke() (*model.CreateCloudPersistentVolumeClaimsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCloudPersistentVolumeClaimsResponse), nil
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

type CreateKubernetesClusterCertInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateKubernetesClusterCertInvoker) Invoke() (*model.CreateKubernetesClusterCertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateKubernetesClusterCertResponse), nil
	}
}

type CreateNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNodeInvoker) Invoke() (*model.CreateNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNodeResponse), nil
	}
}

type CreateNodePoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNodePoolInvoker) Invoke() (*model.CreateNodePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNodePoolResponse), nil
	}
}

type DeleteAddonInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAddonInstanceInvoker) Invoke() (*model.DeleteAddonInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAddonInstanceResponse), nil
	}
}

type DeleteCloudPersistentVolumeClaimsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCloudPersistentVolumeClaimsInvoker) Invoke() (*model.DeleteCloudPersistentVolumeClaimsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCloudPersistentVolumeClaimsResponse), nil
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

type DeleteNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNodeInvoker) Invoke() (*model.DeleteNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNodeResponse), nil
	}
}

type DeleteNodePoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNodePoolInvoker) Invoke() (*model.DeleteNodePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNodePoolResponse), nil
	}
}

type HibernateClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *HibernateClusterInvoker) Invoke() (*model.HibernateClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.HibernateClusterResponse), nil
	}
}

type ListAddonInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAddonInstancesInvoker) Invoke() (*model.ListAddonInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAddonInstancesResponse), nil
	}
}

type ListAddonTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAddonTemplatesInvoker) Invoke() (*model.ListAddonTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAddonTemplatesResponse), nil
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

type ListNodePoolsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNodePoolsInvoker) Invoke() (*model.ListNodePoolsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNodePoolsResponse), nil
	}
}

type ListNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNodesInvoker) Invoke() (*model.ListNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNodesResponse), nil
	}
}

type MigrateNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *MigrateNodeInvoker) Invoke() (*model.MigrateNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.MigrateNodeResponse), nil
	}
}

type RemoveNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *RemoveNodeInvoker) Invoke() (*model.RemoveNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RemoveNodeResponse), nil
	}
}

type ResetNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetNodeInvoker) Invoke() (*model.ResetNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetNodeResponse), nil
	}
}

type ShowAddonInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAddonInstanceInvoker) Invoke() (*model.ShowAddonInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAddonInstanceResponse), nil
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

type ShowJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobInvoker) Invoke() (*model.ShowJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobResponse), nil
	}
}

type ShowNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNodeInvoker) Invoke() (*model.ShowNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNodeResponse), nil
	}
}

type ShowNodePoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNodePoolInvoker) Invoke() (*model.ShowNodePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNodePoolResponse), nil
	}
}

type ShowQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuotasInvoker) Invoke() (*model.ShowQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuotasResponse), nil
	}
}

type UpdateAddonInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAddonInstanceInvoker) Invoke() (*model.UpdateAddonInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAddonInstanceResponse), nil
	}
}

type UpdateClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateClusterInvoker) Invoke() (*model.UpdateClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateClusterResponse), nil
	}
}

type UpdateNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNodeInvoker) Invoke() (*model.UpdateNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNodeResponse), nil
	}
}

type UpdateNodePoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNodePoolInvoker) Invoke() (*model.UpdateNodePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNodePoolResponse), nil
	}
}
