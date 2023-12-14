package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dli/v1/model"
)

type AssociateQueueToElasticResourcePoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateQueueToElasticResourcePoolInvoker) Invoke() (*model.AssociateQueueToElasticResourcePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateQueueToElasticResourcePoolResponse), nil
	}
}

type AssociateQueueToEnhancedConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateQueueToEnhancedConnectionInvoker) Invoke() (*model.AssociateQueueToEnhancedConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateQueueToEnhancedConnectionResponse), nil
	}
}

type BatchDeleteQueuePlansInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *BatchDeleteQueuePlansInvoker) Invoke() (*model.BatchDeleteQueuePlansResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteQueuePlansResponse), nil
	}
}

type ChangeAuthorizationInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ChangeAuthorizationInvoker) Invoke() (*model.ChangeAuthorizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeAuthorizationResponse), nil
	}
}

type ChangeQueuePlanInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ChangeQueuePlanInvoker) Invoke() (*model.ChangeQueuePlanResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeQueuePlanResponse), nil
	}
}

type CreateAuthInfoInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *CreateAuthInfoInvoker) Invoke() (*model.CreateAuthInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAuthInfoResponse), nil
	}
}

type CreateConnectivityTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateConnectivityTaskInvoker) Invoke() (*model.CreateConnectivityTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConnectivityTaskResponse), nil
	}
}

type CreateDatasourceConnectionInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *CreateDatasourceConnectionInvoker) Invoke() (*model.CreateDatasourceConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDatasourceConnectionResponse), nil
	}
}

type CreateDliAgencyInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *CreateDliAgencyInvoker) Invoke() (*model.CreateDliAgencyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDliAgencyResponse), nil
	}
}

type CreateElasticResourcePoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateElasticResourcePoolInvoker) Invoke() (*model.CreateElasticResourcePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateElasticResourcePoolResponse), nil
	}
}

type CreateEnhancedConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEnhancedConnectionInvoker) Invoke() (*model.CreateEnhancedConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEnhancedConnectionResponse), nil
	}
}

type CreateEnhancedConnectionRoutesInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *CreateEnhancedConnectionRoutesInvoker) Invoke() (*model.CreateEnhancedConnectionRoutesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEnhancedConnectionRoutesResponse), nil
	}
}

type CreateGlobalVariableInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGlobalVariableInvoker) Invoke() (*model.CreateGlobalVariableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGlobalVariableResponse), nil
	}
}

type CreateJobAuthInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateJobAuthInfoInvoker) Invoke() (*model.CreateJobAuthInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateJobAuthInfoResponse), nil
	}
}

type CreateQueueInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateQueueInvoker) Invoke() (*model.CreateQueueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateQueueResponse), nil
	}
}

type CreateQueuePlanInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *CreateQueuePlanInvoker) Invoke() (*model.CreateQueuePlanResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateQueuePlanResponse), nil
	}
}

type CreateQueuePropertyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateQueuePropertyInvoker) Invoke() (*model.CreateQueuePropertyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateQueuePropertyResponse), nil
	}
}

type CreateRouteToEnhancedConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRouteToEnhancedConnectionInvoker) Invoke() (*model.CreateRouteToEnhancedConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRouteToEnhancedConnectionResponse), nil
	}
}

type DeleteAuthInfoInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DeleteAuthInfoInvoker) Invoke() (*model.DeleteAuthInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAuthInfoResponse), nil
	}
}

type DeleteDatasourceConnectionInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DeleteDatasourceConnectionInvoker) Invoke() (*model.DeleteDatasourceConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDatasourceConnectionResponse), nil
	}
}

type DeleteElasticResourcePoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteElasticResourcePoolInvoker) Invoke() (*model.DeleteElasticResourcePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteElasticResourcePoolResponse), nil
	}
}

type DeleteEnhancedConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEnhancedConnectionInvoker) Invoke() (*model.DeleteEnhancedConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEnhancedConnectionResponse), nil
	}
}

type DeleteEnhancedConnectionRoutesInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DeleteEnhancedConnectionRoutesInvoker) Invoke() (*model.DeleteEnhancedConnectionRoutesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEnhancedConnectionRoutesResponse), nil
	}
}

type DeleteGlobalVariableInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGlobalVariableInvoker) Invoke() (*model.DeleteGlobalVariableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGlobalVariableResponse), nil
	}
}

type DeleteJobAuthInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteJobAuthInfoInvoker) Invoke() (*model.DeleteJobAuthInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteJobAuthInfoResponse), nil
	}
}

type DeleteQueueInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteQueueInvoker) Invoke() (*model.DeleteQueueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteQueueResponse), nil
	}
}

type DeleteQueuePlanInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DeleteQueuePlanInvoker) Invoke() (*model.DeleteQueuePlanResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteQueuePlanResponse), nil
	}
}

type DeleteQueuePropertyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteQueuePropertyInvoker) Invoke() (*model.DeleteQueuePropertyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteQueuePropertyResponse), nil
	}
}

type DeleteResourceInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DeleteResourceInvoker) Invoke() (*model.DeleteResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteResourceResponse), nil
	}
}

type DeleteRouteFromEnhancedConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRouteFromEnhancedConnectionInvoker) Invoke() (*model.DeleteRouteFromEnhancedConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRouteFromEnhancedConnectionResponse), nil
	}
}

type DisassociateQueueFromEnhancedConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateQueueFromEnhancedConnectionInvoker) Invoke() (*model.DisassociateQueueFromEnhancedConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateQueueFromEnhancedConnectionResponse), nil
	}
}

type ListAuthInfoInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAuthInfoInvoker) Invoke() (*model.ListAuthInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuthInfoResponse), nil
	}
}

type ListAuthorizationPrivilegesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuthorizationPrivilegesInvoker) Invoke() (*model.ListAuthorizationPrivilegesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuthorizationPrivilegesResponse), nil
	}
}

type ListDatabaseUsersInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListDatabaseUsersInvoker) Invoke() (*model.ListDatabaseUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatabaseUsersResponse), nil
	}
}

type ListDatasourceConnectionsInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListDatasourceConnectionsInvoker) Invoke() (*model.ListDatasourceConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatasourceConnectionsResponse), nil
	}
}

type ListElasticResourcePoolQueuesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListElasticResourcePoolQueuesInvoker) Invoke() (*model.ListElasticResourcePoolQueuesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListElasticResourcePoolQueuesResponse), nil
	}
}

type ListElasticResourcePoolScaleRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListElasticResourcePoolScaleRecordsInvoker) Invoke() (*model.ListElasticResourcePoolScaleRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListElasticResourcePoolScaleRecordsResponse), nil
	}
}

type ListElasticResourcePoolsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListElasticResourcePoolsInvoker) Invoke() (*model.ListElasticResourcePoolsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListElasticResourcePoolsResponse), nil
	}
}

type ListEnhancedConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEnhancedConnectionsInvoker) Invoke() (*model.ListEnhancedConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnhancedConnectionsResponse), nil
	}
}

type ListGlobalVariablesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGlobalVariablesInvoker) Invoke() (*model.ListGlobalVariablesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGlobalVariablesResponse), nil
	}
}

type ListJobAuthInfosInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJobAuthInfosInvoker) Invoke() (*model.ListJobAuthInfosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJobAuthInfosResponse), nil
	}
}

type ListQueuePlansInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListQueuePlansInvoker) Invoke() (*model.ListQueuePlansResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQueuePlansResponse), nil
	}
}

type ListQueuePropertiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQueuePropertiesInvoker) Invoke() (*model.ListQueuePropertiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQueuePropertiesResponse), nil
	}
}

type ListQueueUsersInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListQueueUsersInvoker) Invoke() (*model.ListQueueUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQueueUsersResponse), nil
	}
}

type ListQueuesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQueuesInvoker) Invoke() (*model.ListQueuesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQueuesResponse), nil
	}
}

type ListResourcesInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListResourcesInvoker) Invoke() (*model.ListResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourcesResponse), nil
	}
}

type ListTablePrivilegesInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListTablePrivilegesInvoker) Invoke() (*model.ListTablePrivilegesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTablePrivilegesResponse), nil
	}
}

type ListTableUsersInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListTableUsersInvoker) Invoke() (*model.ListTableUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTableUsersResponse), nil
	}
}

type RegisterAuthorizedQueueInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *RegisterAuthorizedQueueInvoker) Invoke() (*model.RegisterAuthorizedQueueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterAuthorizedQueueResponse), nil
	}
}

type RunAuthorizationActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunAuthorizationActionInvoker) Invoke() (*model.RunAuthorizationActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunAuthorizationActionResponse), nil
	}
}

type RunQueueActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunQueueActionInvoker) Invoke() (*model.RunQueueActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunQueueActionResponse), nil
	}
}

type ShowConnectivityTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConnectivityTaskInvoker) Invoke() (*model.ShowConnectivityTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConnectivityTaskResponse), nil
	}
}

type ShowDatasourceConnectionInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShowDatasourceConnectionInvoker) Invoke() (*model.ShowDatasourceConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDatasourceConnectionResponse), nil
	}
}

type ShowDliAgencyInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShowDliAgencyInvoker) Invoke() (*model.ShowDliAgencyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDliAgencyResponse), nil
	}
}

type ShowEnhancedConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEnhancedConnectionInvoker) Invoke() (*model.ShowEnhancedConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEnhancedConnectionResponse), nil
	}
}

type ShowEnhancedConnectionPrivilegeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEnhancedConnectionPrivilegeInvoker) Invoke() (*model.ShowEnhancedConnectionPrivilegeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEnhancedConnectionPrivilegeResponse), nil
	}
}

type ShowQueueInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQueueInvoker) Invoke() (*model.ShowQueueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQueueResponse), nil
	}
}

type ShowQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuotaInvoker) Invoke() (*model.ShowQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuotaResponse), nil
	}
}

type ShowResourceInfoInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShowResourceInfoInvoker) Invoke() (*model.ShowResourceInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceInfoResponse), nil
	}
}

type UpdateAuthInfoInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *UpdateAuthInfoInvoker) Invoke() (*model.UpdateAuthInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAuthInfoResponse), nil
	}
}

type UpdateElasticResourcePoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateElasticResourcePoolInvoker) Invoke() (*model.UpdateElasticResourcePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateElasticResourcePoolResponse), nil
	}
}

type UpdateElasticResourcePoolQueueInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateElasticResourcePoolQueueInvoker) Invoke() (*model.UpdateElasticResourcePoolQueueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateElasticResourcePoolQueueResponse), nil
	}
}

type UpdateEnhancedConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEnhancedConnectionInvoker) Invoke() (*model.UpdateEnhancedConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEnhancedConnectionResponse), nil
	}
}

type UpdateGlobalVariableInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGlobalVariableInvoker) Invoke() (*model.UpdateGlobalVariableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGlobalVariableResponse), nil
	}
}

type UpdateGroupOrResourceOwnerInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *UpdateGroupOrResourceOwnerInvoker) Invoke() (*model.UpdateGroupOrResourceOwnerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGroupOrResourceOwnerResponse), nil
	}
}

type UpdateJobAuthInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateJobAuthInfoInvoker) Invoke() (*model.UpdateJobAuthInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateJobAuthInfoResponse), nil
	}
}

type UpdateQueueCidrInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *UpdateQueueCidrInvoker) Invoke() (*model.UpdateQueueCidrResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateQueueCidrResponse), nil
	}
}

type UpdateQueuePropertyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateQueuePropertyInvoker) Invoke() (*model.UpdateQueuePropertyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateQueuePropertyResponse), nil
	}
}

type UploadFilesInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *UploadFilesInvoker) Invoke() (*model.UploadFilesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadFilesResponse), nil
	}
}

type UploadJarsInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *UploadJarsInvoker) Invoke() (*model.UploadJarsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadJarsResponse), nil
	}
}

type UploadPythonFilesInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *UploadPythonFilesInvoker) Invoke() (*model.UploadPythonFilesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadPythonFilesResponse), nil
	}
}

type UploadResourcesInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *UploadResourcesInvoker) Invoke() (*model.UploadResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadResourcesResponse), nil
	}
}

type BatchDeleteFlinkJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteFlinkJobsInvoker) Invoke() (*model.BatchDeleteFlinkJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteFlinkJobsResponse), nil
	}
}

type BatchRunFlinkJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchRunFlinkJobsInvoker) Invoke() (*model.BatchRunFlinkJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchRunFlinkJobsResponse), nil
	}
}

type ChangeFlinkJobStatusReportInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ChangeFlinkJobStatusReportInvoker) Invoke() (*model.ChangeFlinkJobStatusReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeFlinkJobStatusReportResponse), nil
	}
}

type CreateFlinkJarJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFlinkJarJobInvoker) Invoke() (*model.CreateFlinkJarJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFlinkJarJobResponse), nil
	}
}

type CreateFlinkSqlJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFlinkSqlJobInvoker) Invoke() (*model.CreateFlinkSqlJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFlinkSqlJobResponse), nil
	}
}

type CreateFlinkSqlJobGraphInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFlinkSqlJobGraphInvoker) Invoke() (*model.CreateFlinkSqlJobGraphResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFlinkSqlJobGraphResponse), nil
	}
}

type CreateFlinkSqlJobTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFlinkSqlJobTemplateInvoker) Invoke() (*model.CreateFlinkSqlJobTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFlinkSqlJobTemplateResponse), nil
	}
}

type CreateIefMessageChannelInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *CreateIefMessageChannelInvoker) Invoke() (*model.CreateIefMessageChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateIefMessageChannelResponse), nil
	}
}

type CreateIefSystemEventsInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *CreateIefSystemEventsInvoker) Invoke() (*model.CreateIefSystemEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateIefSystemEventsResponse), nil
	}
}

type DeleteFlinkJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFlinkJobInvoker) Invoke() (*model.DeleteFlinkJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFlinkJobResponse), nil
	}
}

type DeleteFlinkSqlJobTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFlinkSqlJobTemplateInvoker) Invoke() (*model.DeleteFlinkSqlJobTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFlinkSqlJobTemplateResponse), nil
	}
}

type ExportFlinkJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportFlinkJobsInvoker) Invoke() (*model.ExportFlinkJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportFlinkJobsResponse), nil
	}
}

type ImportFlinkJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportFlinkJobsInvoker) Invoke() (*model.ImportFlinkJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportFlinkJobsResponse), nil
	}
}

type ListFlinkJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlinkJobsInvoker) Invoke() (*model.ListFlinkJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlinkJobsResponse), nil
	}
}

type ListFlinkSqlJobTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlinkSqlJobTemplatesInvoker) Invoke() (*model.ListFlinkSqlJobTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlinkSqlJobTemplatesResponse), nil
	}
}

type RegisterBucketInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *RegisterBucketInvoker) Invoke() (*model.RegisterBucketResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterBucketResponse), nil
	}
}

type RunIefJobActionCallBackInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *RunIefJobActionCallBackInvoker) Invoke() (*model.RunIefJobActionCallBackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunIefJobActionCallBackResponse), nil
	}
}

type ShowFlinkJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFlinkJobInvoker) Invoke() (*model.ShowFlinkJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFlinkJobResponse), nil
	}
}

type ShowFlinkJobExecutionGraphInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFlinkJobExecutionGraphInvoker) Invoke() (*model.ShowFlinkJobExecutionGraphResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFlinkJobExecutionGraphResponse), nil
	}
}

type ShowFlinkMetricInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShowFlinkMetricInvoker) Invoke() (*model.ShowFlinkMetricResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFlinkMetricResponse), nil
	}
}

type StopFlinkJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopFlinkJobsInvoker) Invoke() (*model.StopFlinkJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopFlinkJobsResponse), nil
	}
}

type UpdateFlinkJarJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFlinkJarJobInvoker) Invoke() (*model.UpdateFlinkJarJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFlinkJarJobResponse), nil
	}
}

type UpdateFlinkSqlJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFlinkSqlJobInvoker) Invoke() (*model.UpdateFlinkSqlJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFlinkSqlJobResponse), nil
	}
}

type UpdateFlinkSqlJobTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFlinkSqlJobTemplateInvoker) Invoke() (*model.UpdateFlinkSqlJobTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFlinkSqlJobTemplateResponse), nil
	}
}

type CancelSparkJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelSparkJobInvoker) Invoke() (*model.CancelSparkJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelSparkJobResponse), nil
	}
}

type CreateSparkJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSparkJobInvoker) Invoke() (*model.CreateSparkJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSparkJobResponse), nil
	}
}

type CreateSparkJobTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSparkJobTemplateInvoker) Invoke() (*model.CreateSparkJobTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSparkJobTemplateResponse), nil
	}
}

type ListSparkJobTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSparkJobTemplatesInvoker) Invoke() (*model.ListSparkJobTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSparkJobTemplatesResponse), nil
	}
}

type ListSparkJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSparkJobsInvoker) Invoke() (*model.ListSparkJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSparkJobsResponse), nil
	}
}

type ShowBatchLogInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShowBatchLogInvoker) Invoke() (*model.ShowBatchLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBatchLogResponse), nil
	}
}

type ShowSparkJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSparkJobInvoker) Invoke() (*model.ShowSparkJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSparkJobResponse), nil
	}
}

type ShowSparkJobStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSparkJobStatusInvoker) Invoke() (*model.ShowSparkJobStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSparkJobStatusResponse), nil
	}
}

type ShowSparkJobTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSparkJobTemplateInvoker) Invoke() (*model.ShowSparkJobTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSparkJobTemplateResponse), nil
	}
}

type UpdateSparkJobTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSparkJobTemplateInvoker) Invoke() (*model.UpdateSparkJobTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSparkJobTemplateResponse), nil
	}
}

type BatchDeleteSqlJobTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteSqlJobTemplatesInvoker) Invoke() (*model.BatchDeleteSqlJobTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteSqlJobTemplatesResponse), nil
	}
}

type CancelSqlJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelSqlJobInvoker) Invoke() (*model.CancelSqlJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelSqlJobResponse), nil
	}
}

type CheckSqlInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckSqlInvoker) Invoke() (*model.CheckSqlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckSqlResponse), nil
	}
}

type CreateDatabaseInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *CreateDatabaseInvoker) Invoke() (*model.CreateDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDatabaseResponse), nil
	}
}

type CreateSqlJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSqlJobInvoker) Invoke() (*model.CreateSqlJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSqlJobResponse), nil
	}
}

type CreateSqlJobTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSqlJobTemplateInvoker) Invoke() (*model.CreateSqlJobTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSqlJobTemplateResponse), nil
	}
}

type CreateTableInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *CreateTableInvoker) Invoke() (*model.CreateTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTableResponse), nil
	}
}

type DeleteDatabaseInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DeleteDatabaseInvoker) Invoke() (*model.DeleteDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDatabaseResponse), nil
	}
}

type DeleteTableInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DeleteTableInvoker) Invoke() (*model.DeleteTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTableResponse), nil
	}
}

type ExportSqlJobResultInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ExportSqlJobResultInvoker) Invoke() (*model.ExportSqlJobResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportSqlJobResultResponse), nil
	}
}

type ExportTableInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ExportTableInvoker) Invoke() (*model.ExportTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportTableResponse), nil
	}
}

type ImportTableInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ImportTableInvoker) Invoke() (*model.ImportTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportTableResponse), nil
	}
}

type ListAllTablesInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAllTablesInvoker) Invoke() (*model.ListAllTablesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllTablesResponse), nil
	}
}

type ListDatabasesInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListDatabasesInvoker) Invoke() (*model.ListDatabasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatabasesResponse), nil
	}
}

type ListSqlJobTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSqlJobTemplatesInvoker) Invoke() (*model.ListSqlJobTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSqlJobTemplatesResponse), nil
	}
}

type ListSqlJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSqlJobsInvoker) Invoke() (*model.ListSqlJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSqlJobsResponse), nil
	}
}

type PreviewSqlJobResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *PreviewSqlJobResultInvoker) Invoke() (*model.PreviewSqlJobResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PreviewSqlJobResultResponse), nil
	}
}

type ShowDescribeTableInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShowDescribeTableInvoker) Invoke() (*model.ShowDescribeTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDescribeTableResponse), nil
	}
}

type ShowPartitionsInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShowPartitionsInvoker) Invoke() (*model.ShowPartitionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPartitionsResponse), nil
	}
}

type ShowSqlJobDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSqlJobDetailInvoker) Invoke() (*model.ShowSqlJobDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlJobDetailResponse), nil
	}
}

type ShowSqlJobProgressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSqlJobProgressInvoker) Invoke() (*model.ShowSqlJobProgressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlJobProgressResponse), nil
	}
}

type ShowSqlJobStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSqlJobStatusInvoker) Invoke() (*model.ShowSqlJobStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlJobStatusResponse), nil
	}
}

type ShowSqlSampleTemplatesInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShowSqlSampleTemplatesInvoker) Invoke() (*model.ShowSqlSampleTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlSampleTemplatesResponse), nil
	}
}

type ShowTableContentInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShowTableContentInvoker) Invoke() (*model.ShowTableContentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTableContentResponse), nil
	}
}

type UpdateDatabaseOwnerInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *UpdateDatabaseOwnerInvoker) Invoke() (*model.UpdateDatabaseOwnerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDatabaseOwnerResponse), nil
	}
}

type UpdateSqlJobTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSqlJobTemplateInvoker) Invoke() (*model.UpdateSqlJobTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSqlJobTemplateResponse), nil
	}
}

type UpdateTableOwnerInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *UpdateTableOwnerInvoker) Invoke() (*model.UpdateTableOwnerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTableOwnerResponse), nil
	}
}
