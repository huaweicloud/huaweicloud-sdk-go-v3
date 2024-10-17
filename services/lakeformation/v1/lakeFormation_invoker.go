package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/lakeformation/v1/model"
)

type ApplyForAccessInvoker struct {
	*invoker.BaseInvoker
}

func (i *ApplyForAccessInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ApplyForAccessInvoker) Invoke() (*model.ApplyForAccessResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApplyForAccessResponse), nil
	}
}

type BatchAuthorizeInterfaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAuthorizeInterfaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchAuthorizeInterfaceInvoker) Invoke() (*model.BatchAuthorizeInterfaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAuthorizeInterfaceResponse), nil
	}
}

type BatchCancelAuthorizationInterfaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCancelAuthorizationInterfaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCancelAuthorizationInterfaceInvoker) Invoke() (*model.BatchCancelAuthorizationInterfaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCancelAuthorizationInterfaceResponse), nil
	}
}

type BatchCheckPermissionInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCheckPermissionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCheckPermissionInvoker) Invoke() (*model.BatchCheckPermissionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCheckPermissionResponse), nil
	}
}

type CreateAccessClientInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAccessClientInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAccessClientInvoker) Invoke() (*model.CreateAccessClientResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAccessClientResponse), nil
	}
}

type DeleteAccessClientInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAccessClientInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAccessClientInvoker) Invoke() (*model.DeleteAccessClientResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAccessClientResponse), nil
	}
}

type ListAccessClientInfosInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccessClientInfosInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAccessClientInfosInvoker) Invoke() (*model.ListAccessClientInfosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccessClientInfosResponse), nil
	}
}

type ListAccessInfosInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccessInfosInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAccessInfosInvoker) Invoke() (*model.ListAccessInfosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccessInfosResponse), nil
	}
}

type ListInterfacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInterfacesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInterfacesInvoker) Invoke() (*model.ListInterfacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInterfacesResponse), nil
	}
}

type ListPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPolicyInvoker) Invoke() (*model.ListPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPolicyResponse), nil
	}
}

type ShowAccessClientInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAccessClientInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAccessClientInvoker) Invoke() (*model.ShowAccessClientResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAccessClientResponse), nil
	}
}

type ShowSyncPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSyncPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSyncPolicyInvoker) Invoke() (*model.ShowSyncPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSyncPolicyResponse), nil
	}
}

type UpdateAccessClientInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAccessClientInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAccessClientInvoker) Invoke() (*model.UpdateAccessClientResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAccessClientResponse), nil
	}
}

type CreateAgencyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAgencyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAgencyInvoker) Invoke() (*model.CreateAgencyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAgencyResponse), nil
	}
}

type DeleteAgencyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAgencyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAgencyInvoker) Invoke() (*model.DeleteAgencyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAgencyResponse), nil
	}
}

type ShowAgencyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAgencyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAgencyInvoker) Invoke() (*model.ShowAgencyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAgencyResponse), nil
	}
}

type CreateCatalogInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCatalogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCatalogInvoker) Invoke() (*model.CreateCatalogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCatalogResponse), nil
	}
}

type DeleteCatalogInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCatalogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCatalogInvoker) Invoke() (*model.DeleteCatalogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCatalogResponse), nil
	}
}

type ListCatalogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCatalogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCatalogsInvoker) Invoke() (*model.ListCatalogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCatalogsResponse), nil
	}
}

type ShowCatalogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCatalogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCatalogInvoker) Invoke() (*model.ShowCatalogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCatalogResponse), nil
	}
}

type UpdateCatalogInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCatalogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCatalogInvoker) Invoke() (*model.UpdateCatalogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCatalogResponse), nil
	}
}

type ListConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListConfigsInvoker) Invoke() (*model.ListConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConfigsResponse), nil
	}
}

type ShowCredentialInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCredentialInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCredentialInvoker) Invoke() (*model.ShowCredentialResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCredentialResponse), nil
	}
}

type CreateDatabaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDatabaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDatabaseInvoker) Invoke() (*model.CreateDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDatabaseResponse), nil
	}
}

type DeleteDatabaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDatabaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDatabaseInvoker) Invoke() (*model.DeleteDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDatabaseResponse), nil
	}
}

type ListDatabaseNamesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatabaseNamesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDatabaseNamesInvoker) Invoke() (*model.ListDatabaseNamesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatabaseNamesResponse), nil
	}
}

type ListDatabasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatabasesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDatabasesInvoker) Invoke() (*model.ListDatabasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatabasesResponse), nil
	}
}

type ShowDatabaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDatabaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDatabaseInvoker) Invoke() (*model.ShowDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDatabaseResponse), nil
	}
}

type UpdateDatabaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDatabaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDatabaseInvoker) Invoke() (*model.UpdateDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDatabaseResponse), nil
	}
}

type CreateFunctionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFunctionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFunctionInvoker) Invoke() (*model.CreateFunctionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFunctionResponse), nil
	}
}

type DeleteFunctionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFunctionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteFunctionInvoker) Invoke() (*model.DeleteFunctionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFunctionResponse), nil
	}
}

type ListAllFunctionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllFunctionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAllFunctionsInvoker) Invoke() (*model.ListAllFunctionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllFunctionsResponse), nil
	}
}

type ListFunctionNamesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFunctionNamesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFunctionNamesInvoker) Invoke() (*model.ListFunctionNamesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFunctionNamesResponse), nil
	}
}

type ListFunctionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFunctionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFunctionsInvoker) Invoke() (*model.ListFunctionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFunctionsResponse), nil
	}
}

type ShowFunctionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFunctionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFunctionInvoker) Invoke() (*model.ShowFunctionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFunctionResponse), nil
	}
}

type UpdateFunctionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFunctionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateFunctionInvoker) Invoke() (*model.UpdateFunctionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFunctionResponse), nil
	}
}

type AuthorizeAccessServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *AuthorizeAccessServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AuthorizeAccessServiceInvoker) Invoke() (*model.AuthorizeAccessServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthorizeAccessServiceResponse), nil
	}
}

type CreateAgreementInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAgreementInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAgreementInvoker) Invoke() (*model.CreateAgreementResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAgreementResponse), nil
	}
}

type DeleteAgreementInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAgreementInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAgreementInvoker) Invoke() (*model.DeleteAgreementResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAgreementResponse), nil
	}
}

type ShowAccessServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAccessServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAccessServiceInvoker) Invoke() (*model.ShowAccessServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAccessServiceResponse), nil
	}
}

type ShowAgreementInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAgreementInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAgreementInvoker) Invoke() (*model.ShowAgreementResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAgreementResponse), nil
	}
}

type ShowAgreementRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAgreementRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAgreementRuleInvoker) Invoke() (*model.ShowAgreementRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAgreementRuleResponse), nil
	}
}

type CountMetaObjInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountMetaObjInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountMetaObjInvoker) Invoke() (*model.CountMetaObjResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountMetaObjResponse), nil
	}
}

type CreateLakeFormationInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLakeFormationInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateLakeFormationInstanceInvoker) Invoke() (*model.CreateLakeFormationInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLakeFormationInstanceResponse), nil
	}
}

type DeleteLakeFormationInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLakeFormationInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteLakeFormationInstanceInvoker) Invoke() (*model.DeleteLakeFormationInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLakeFormationInstanceResponse), nil
	}
}

type ListLakeFormationInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLakeFormationInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLakeFormationInstancesInvoker) Invoke() (*model.ListLakeFormationInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLakeFormationInstancesResponse), nil
	}
}

type MoveLakeFormationInstanceOutRecycleBinInvoker struct {
	*invoker.BaseInvoker
}

func (i *MoveLakeFormationInstanceOutRecycleBinInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *MoveLakeFormationInstanceOutRecycleBinInvoker) Invoke() (*model.MoveLakeFormationInstanceOutRecycleBinResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.MoveLakeFormationInstanceOutRecycleBinResponse), nil
	}
}

type ShowLakeFormationInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLakeFormationInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLakeFormationInstanceInvoker) Invoke() (*model.ShowLakeFormationInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLakeFormationInstanceResponse), nil
	}
}

type UpdateLakeFormationInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLakeFormationInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateLakeFormationInstanceInvoker) Invoke() (*model.UpdateLakeFormationInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLakeFormationInstanceResponse), nil
	}
}

type UpdateLakeFormationInstanceDefaultInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLakeFormationInstanceDefaultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateLakeFormationInstanceDefaultInvoker) Invoke() (*model.UpdateLakeFormationInstanceDefaultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLakeFormationInstanceDefaultResponse), nil
	}
}

type UpdateLakeFormationInstanceScaleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLakeFormationInstanceScaleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateLakeFormationInstanceScaleInvoker) Invoke() (*model.UpdateLakeFormationInstanceScaleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLakeFormationInstanceScaleResponse), nil
	}
}

type ListObsBucketsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListObsBucketsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListObsBucketsInvoker) Invoke() (*model.ListObsBucketsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListObsBucketsResponse), nil
	}
}

type ListObsObjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListObsObjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListObsObjectInvoker) Invoke() (*model.ListObsObjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListObsObjectResponse), nil
	}
}

type AddPartitionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddPartitionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddPartitionsInvoker) Invoke() (*model.AddPartitionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddPartitionsResponse), nil
	}
}

type BatchDeletePartitionInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeletePartitionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeletePartitionInvoker) Invoke() (*model.BatchDeletePartitionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeletePartitionResponse), nil
	}
}

type BatchDeletePartitionedStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeletePartitionedStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeletePartitionedStatisticsInvoker) Invoke() (*model.BatchDeletePartitionedStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeletePartitionedStatisticsResponse), nil
	}
}

type BatchListPartitionByValuesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchListPartitionByValuesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchListPartitionByValuesInvoker) Invoke() (*model.BatchListPartitionByValuesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchListPartitionByValuesResponse), nil
	}
}

type BatchUpdatePartitionInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdatePartitionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdatePartitionInvoker) Invoke() (*model.BatchUpdatePartitionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdatePartitionResponse), nil
	}
}

type ListPartitionNamesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPartitionNamesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPartitionNamesInvoker) Invoke() (*model.ListPartitionNamesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPartitionNamesResponse), nil
	}
}

type ListPartitionNamesWithoutLimitInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPartitionNamesWithoutLimitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPartitionNamesWithoutLimitInvoker) Invoke() (*model.ListPartitionNamesWithoutLimitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPartitionNamesWithoutLimitResponse), nil
	}
}

type ListPartitionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPartitionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPartitionsInvoker) Invoke() (*model.ListPartitionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPartitionsResponse), nil
	}
}

type BatchShowPartitionColumnStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchShowPartitionColumnStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchShowPartitionColumnStatisticsInvoker) Invoke() (*model.BatchShowPartitionColumnStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchShowPartitionColumnStatisticsResponse), nil
	}
}

type DeletePartitionColumnStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePartitionColumnStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePartitionColumnStatisticsInvoker) Invoke() (*model.DeletePartitionColumnStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePartitionColumnStatisticsResponse), nil
	}
}

type SetPartitionColumnStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetPartitionColumnStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetPartitionColumnStatisticsInvoker) Invoke() (*model.SetPartitionColumnStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetPartitionColumnStatisticsResponse), nil
	}
}

type ListQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQuotasInvoker) Invoke() (*model.ListQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotasResponse), nil
	}
}

type AssociatePrincipalsInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociatePrincipalsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociatePrincipalsInvoker) Invoke() (*model.AssociatePrincipalsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociatePrincipalsResponse), nil
	}
}

type CreateRoleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRoleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRoleInvoker) Invoke() (*model.CreateRoleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRoleResponse), nil
	}
}

type DeleteRoleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRoleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRoleInvoker) Invoke() (*model.DeleteRoleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRoleResponse), nil
	}
}

type ListPrincipalsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPrincipalsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPrincipalsInvoker) Invoke() (*model.ListPrincipalsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPrincipalsResponse), nil
	}
}

type ListRoleNamesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRoleNamesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRoleNamesInvoker) Invoke() (*model.ListRoleNamesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRoleNamesResponse), nil
	}
}

type ListRolesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRolesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRolesInvoker) Invoke() (*model.ListRolesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRolesResponse), nil
	}
}

type RevokePrincipalsInvoker struct {
	*invoker.BaseInvoker
}

func (i *RevokePrincipalsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RevokePrincipalsInvoker) Invoke() (*model.RevokePrincipalsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RevokePrincipalsResponse), nil
	}
}

type ShowRoleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRoleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRoleInvoker) Invoke() (*model.ShowRoleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRoleResponse), nil
	}
}

type UpdatePrincipalsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePrincipalsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePrincipalsInvoker) Invoke() (*model.UpdatePrincipalsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePrincipalsResponse), nil
	}
}

type UpdateRoleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRoleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRoleInvoker) Invoke() (*model.UpdateRoleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRoleResponse), nil
	}
}

type ListSpecsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSpecsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSpecsInvoker) Invoke() (*model.ListSpecsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSpecsResponse), nil
	}
}

type CreateTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTableInvoker) Invoke() (*model.CreateTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTableResponse), nil
	}
}

type DeleteAllTablesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAllTablesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAllTablesInvoker) Invoke() (*model.DeleteAllTablesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAllTablesResponse), nil
	}
}

type DeleteTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTableInvoker) Invoke() (*model.DeleteTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTableResponse), nil
	}
}

type ListTableMetaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTableMetaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTableMetaInvoker) Invoke() (*model.ListTableMetaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTableMetaResponse), nil
	}
}

type ListTableNamesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTableNamesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTableNamesInvoker) Invoke() (*model.ListTableNamesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTableNamesResponse), nil
	}
}

type ListTablesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTablesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTablesInvoker) Invoke() (*model.ListTablesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTablesResponse), nil
	}
}

type ListTablesByNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTablesByNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTablesByNameInvoker) Invoke() (*model.ListTablesByNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTablesByNameResponse), nil
	}
}

type ShowTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTableInvoker) Invoke() (*model.ShowTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTableResponse), nil
	}
}

type UpdateTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTableInvoker) Invoke() (*model.UpdateTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTableResponse), nil
	}
}

type DeleteTableColumnStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTableColumnStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTableColumnStatisticsInvoker) Invoke() (*model.DeleteTableColumnStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTableColumnStatisticsResponse), nil
	}
}

type ListTableColumnStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTableColumnStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTableColumnStatisticsInvoker) Invoke() (*model.ListTableColumnStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTableColumnStatisticsResponse), nil
	}
}

type SetTableColumnStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetTableColumnStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetTableColumnStatisticsInvoker) Invoke() (*model.SetTableColumnStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetTableColumnStatisticsResponse), nil
	}
}

type BatchCreateConstraintInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateConstraintInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateConstraintInvoker) Invoke() (*model.BatchCreateConstraintResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateConstraintResponse), nil
	}
}

type DeleteConstraintInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteConstraintInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteConstraintInvoker) Invoke() (*model.DeleteConstraintResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteConstraintResponse), nil
	}
}

type ListConstraintsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConstraintsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListConstraintsInvoker) Invoke() (*model.ListConstraintsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConstraintsResponse), nil
	}
}

type BatchUpdateLakeFormationInstanceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateLakeFormationInstanceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdateLakeFormationInstanceTagsInvoker) Invoke() (*model.BatchUpdateLakeFormationInstanceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateLakeFormationInstanceTagsResponse), nil
	}
}

type ListLakeFormationInstanceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLakeFormationInstanceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLakeFormationInstanceTagsInvoker) Invoke() (*model.ListLakeFormationInstanceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLakeFormationInstanceTagsResponse), nil
	}
}

type AssociateRolesInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateRolesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateRolesInvoker) Invoke() (*model.AssociateRolesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateRolesResponse), nil
	}
}

type ListUserRolesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserRolesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUserRolesInvoker) Invoke() (*model.ListUserRolesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserRolesResponse), nil
	}
}

type ListUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUsersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUsersInvoker) Invoke() (*model.ListUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUsersResponse), nil
	}
}

type RevokeRolesInvoker struct {
	*invoker.BaseInvoker
}

func (i *RevokeRolesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RevokeRolesInvoker) Invoke() (*model.RevokeRolesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RevokeRolesResponse), nil
	}
}

type UpdateRolesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRolesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRolesInvoker) Invoke() (*model.UpdateRolesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRolesResponse), nil
	}
}

type ListGroupsForDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupsForDomainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupsForDomainInvoker) Invoke() (*model.ListGroupsForDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupsForDomainResponse), nil
	}
}
