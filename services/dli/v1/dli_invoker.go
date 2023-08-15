package v1

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/dli/v1/model"
)

type CreateJobTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateJobTemplatesInvoker) Invoke() (*model.CreateJobTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateJobTemplatesResponse), nil
	}
}

type CreateSqlTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSqlTemplatesInvoker) Invoke() (*model.CreateSqlTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSqlTemplatesResponse), nil
	}
}

type DeleteSqlTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSqlTemplatesInvoker) Invoke() (*model.DeleteSqlTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSqlTemplatesResponse), nil
	}
}

type ListJobTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJobTemplatesInvoker) Invoke() (*model.ListJobTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJobTemplatesResponse), nil
	}
}

type ShowJobTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobTemplateInvoker) Invoke() (*model.ShowJobTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobTemplateResponse), nil
	}
}

type ShowSqlSampleTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSqlSampleTemplatesInvoker) Invoke() (*model.ShowSqlSampleTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlSampleTemplatesResponse), nil
	}
}

type ShowSqlTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSqlTemplatesInvoker) Invoke() (*model.ShowSqlTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSqlTemplatesResponse), nil
	}
}

type UpdateJobTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateJobTemplatesInvoker) Invoke() (*model.UpdateJobTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateJobTemplatesResponse), nil
	}
}

type UpdateSqlTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSqlTemplatesInvoker) Invoke() (*model.UpdateSqlTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSqlTemplatesResponse), nil
	}
}

type AssociateConnectionQueueInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateConnectionQueueInvoker) Invoke() (*model.AssociateConnectionQueueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateConnectionQueueResponse), nil
	}
}

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

type AuthorizeResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *AuthorizeResourceInvoker) Invoke() (*model.AuthorizeResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthorizeResourceResponse), nil
	}
}

type BatchDeleteQueuePlansInvoker struct {
	*invoker.BaseInvoker
}

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

func (i *ChangeQueuePlanInvoker) Invoke() (*model.ChangeQueuePlanResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeQueuePlanResponse), nil
	}
}

type CheckConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckConnectionInvoker) Invoke() (*model.CheckConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckConnectionResponse), nil
	}
}

type CreateAuthInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAuthInfoInvoker) Invoke() (*model.CreateAuthInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAuthInfoResponse), nil
	}
}

type CreateDatasourceConnectionInvoker struct {
	*invoker.BaseInvoker
}

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

func (i *CreateEnhancedConnectionRoutesInvoker) Invoke() (*model.CreateEnhancedConnectionRoutesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEnhancedConnectionRoutesResponse), nil
	}
}

type CreateGlobalValueInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGlobalValueInvoker) Invoke() (*model.CreateGlobalValueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGlobalValueResponse), nil
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

func (i *CreateQueuePlanInvoker) Invoke() (*model.CreateQueuePlanResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateQueuePlanResponse), nil
	}
}

type DeleteAuthInfoInvoker struct {
	*invoker.BaseInvoker
}

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

func (i *DeleteEnhancedConnectionRoutesInvoker) Invoke() (*model.DeleteEnhancedConnectionRoutesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEnhancedConnectionRoutesResponse), nil
	}
}

type DeleteGlobalValueInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGlobalValueInvoker) Invoke() (*model.DeleteGlobalValueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGlobalValueResponse), nil
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

func (i *DeleteQueuePlanInvoker) Invoke() (*model.DeleteQueuePlanResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteQueuePlanResponse), nil
	}
}

type DeleteResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteResourceInvoker) Invoke() (*model.DeleteResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteResourceResponse), nil
	}
}

type DisassociateConnectionQueueInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateConnectionQueueInvoker) Invoke() (*model.DisassociateConnectionQueueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateConnectionQueueResponse), nil
	}
}

type ListAuthInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuthInfoInvoker) Invoke() (*model.ListAuthInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuthInfoResponse), nil
	}
}

type ListDatabaseUsersInvoker struct {
	*invoker.BaseInvoker
}

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

type ListGlobalValuesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGlobalValuesInvoker) Invoke() (*model.ListGlobalValuesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGlobalValuesResponse), nil
	}
}

type ListQueuePlansInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQueuePlansInvoker) Invoke() (*model.ListQueuePlansResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQueuePlansResponse), nil
	}
}

type ListQueueUsersInvoker struct {
	*invoker.BaseInvoker
}

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

func (i *RegisterAuthorizedQueueInvoker) Invoke() (*model.RegisterAuthorizedQueueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterAuthorizedQueueResponse), nil
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

type ShowDatasourceConnectionInvoker struct {
	*invoker.BaseInvoker
}

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

type ShowEnhancedPrivilegeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEnhancedPrivilegeInvoker) Invoke() (*model.ShowEnhancedPrivilegeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEnhancedPrivilegeResponse), nil
	}
}

type ShowNodeConnectivityInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNodeConnectivityInvoker) Invoke() (*model.ShowNodeConnectivityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNodeConnectivityResponse), nil
	}
}

type ShowObjectUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowObjectUserInvoker) Invoke() (*model.ShowObjectUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowObjectUserResponse), nil
	}
}

type ShowQueueDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQueueDetailInvoker) Invoke() (*model.ShowQueueDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQueueDetailResponse), nil
	}
}

type ShowResourceInfoInvoker struct {
	*invoker.BaseInvoker
}

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

type UpdateElasticResourcePoolQueueInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateElasticResourcePoolQueueInfoInvoker) Invoke() (*model.UpdateElasticResourcePoolQueueInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateElasticResourcePoolQueueInfoResponse), nil
	}
}

type UpdateGlobalValueInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGlobalValueInvoker) Invoke() (*model.UpdateGlobalValueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGlobalValueResponse), nil
	}
}

type UpdateGroupOrResourceOwnerInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGroupOrResourceOwnerInvoker) Invoke() (*model.UpdateGroupOrResourceOwnerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGroupOrResourceOwnerResponse), nil
	}
}

type UpdateHostMassageInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHostMassageInvoker) Invoke() (*model.UpdateHostMassageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHostMassageResponse), nil
	}
}

type UpdateQueueCidrInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateQueueCidrInvoker) Invoke() (*model.UpdateQueueCidrResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateQueueCidrResponse), nil
	}
}

type UploadFilesInvoker struct {
	*invoker.BaseInvoker
}

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

func (i *ChangeFlinkJobStatusReportInvoker) Invoke() (*model.ChangeFlinkJobStatusReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeFlinkJobStatusReportResponse), nil
	}
}

type CreateFlinkJarInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFlinkJarInvoker) Invoke() (*model.CreateFlinkJarResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFlinkJarResponse), nil
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

type CreateFlinkTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFlinkTemplateInvoker) Invoke() (*model.CreateFlinkTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFlinkTemplateResponse), nil
	}
}

type CreateIefMessageChannelInvoker struct {
	*invoker.BaseInvoker
}

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

func (i *CreateIefSystemEventsInvoker) Invoke() (*model.CreateIefSystemEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateIefSystemEventsResponse), nil
	}
}

type CreateStreamGraphInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateStreamGraphInvoker) Invoke() (*model.CreateStreamGraphResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStreamGraphResponse), nil
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

type DeleteFlinkTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFlinkTemplateInvoker) Invoke() (*model.DeleteFlinkTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFlinkTemplateResponse), nil
	}
}

type ExportFlinkJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportFlinkJobInvoker) Invoke() (*model.ExportFlinkJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportFlinkJobResponse), nil
	}
}

type ImportFlinkJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportFlinkJobInvoker) Invoke() (*model.ImportFlinkJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportFlinkJobResponse), nil
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

type ListFlinkTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlinkTemplatesInvoker) Invoke() (*model.ListFlinkTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlinkTemplatesResponse), nil
	}
}

type RegisterBucketInvoker struct {
	*invoker.BaseInvoker
}

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

func (i *RunIefJobActionCallBackInvoker) Invoke() (*model.RunIefJobActionCallBackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunIefJobActionCallBackResponse), nil
	}
}

type ShowFlinkExecuteGraphInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFlinkExecuteGraphInvoker) Invoke() (*model.ShowFlinkExecuteGraphResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFlinkExecuteGraphResponse), nil
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

type ShowFlinkMetricInvoker struct {
	*invoker.BaseInvoker
}

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

type UpdateFlinkJarInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFlinkJarInvoker) Invoke() (*model.UpdateFlinkJarResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFlinkJarResponse), nil
	}
}

type UpdateFlinkSqlInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFlinkSqlInvoker) Invoke() (*model.UpdateFlinkSqlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFlinkSqlResponse), nil
	}
}

type UpdateFlinkTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFlinkTemplateInvoker) Invoke() (*model.UpdateFlinkTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFlinkTemplateResponse), nil
	}
}

type CancelBatchJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelBatchJobInvoker) Invoke() (*model.CancelBatchJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelBatchJobResponse), nil
	}
}

type CreateBatchJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBatchJobInvoker) Invoke() (*model.CreateBatchJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBatchJobResponse), nil
	}
}

type ListBatchesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBatchesInvoker) Invoke() (*model.ListBatchesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBatchesResponse), nil
	}
}

type ShowBatchInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBatchInfoInvoker) Invoke() (*model.ShowBatchInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBatchInfoResponse), nil
	}
}

type ShowBatchLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBatchLogInvoker) Invoke() (*model.ShowBatchLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBatchLogResponse), nil
	}
}

type ShowBatchStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBatchStateInvoker) Invoke() (*model.ShowBatchStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBatchStateResponse), nil
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

type CreateTableInvoker struct {
	*invoker.BaseInvoker
}

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

func (i *ListDatabasesInvoker) Invoke() (*model.ListDatabasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatabasesResponse), nil
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

type PreviewJobResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *PreviewJobResultInvoker) Invoke() (*model.PreviewJobResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PreviewJobResultResponse), nil
	}
}

type ShowDescribeTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDescribeTableInvoker) Invoke() (*model.ShowDescribeTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDescribeTableResponse), nil
	}
}

type ShowJobProgressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobProgressInvoker) Invoke() (*model.ShowJobProgressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobProgressResponse), nil
	}
}

type ShowPartitionsInvoker struct {
	*invoker.BaseInvoker
}

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

type ShowTableContentInvoker struct {
	*invoker.BaseInvoker
}

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

func (i *UpdateDatabaseOwnerInvoker) Invoke() (*model.UpdateDatabaseOwnerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDatabaseOwnerResponse), nil
	}
}

type UpdateTableOwnerInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTableOwnerInvoker) Invoke() (*model.UpdateTableOwnerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTableOwnerResponse), nil
	}
}
