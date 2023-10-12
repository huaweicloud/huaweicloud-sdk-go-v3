package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/lakeformation/v1/model"
)

type ApplyForAccessInvoker struct {
	*invoker.BaseInvoker
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

func (i *BatchCancelAuthorizationInterfaceInvoker) Invoke() (*model.BatchCancelAuthorizationInterfaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCancelAuthorizationInterfaceResponse), nil
	}
}

type ListAccessInfosInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListPolicyInvoker) Invoke() (*model.ListPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPolicyResponse), nil
	}
}

type ShowSyncPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSyncPolicyInvoker) Invoke() (*model.ShowSyncPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSyncPolicyResponse), nil
	}
}

type CreateCatalogInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdateCatalogInvoker) Invoke() (*model.UpdateCatalogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCatalogResponse), nil
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

type ListDatabaseNamesInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListAllFunctionsInvoker) Invoke() (*model.ListAllFunctionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllFunctionsResponse), nil
	}
}

type ListFunctionsInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdateFunctionInvoker) Invoke() (*model.UpdateFunctionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFunctionResponse), nil
	}
}

type CreateAgreementInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAgreementInvoker) Invoke() (*model.CreateAgreementResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAgreementResponse), nil
	}
}

type ShowAgreementInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdateLakeFormationInstanceInvoker) Invoke() (*model.UpdateLakeFormationInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLakeFormationInstanceResponse), nil
	}
}

type ListObsBucketsInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListPartitionNamesInvoker) Invoke() (*model.ListPartitionNamesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPartitionNamesResponse), nil
	}
}

type ListPartitionsInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListQuotasInvoker) Invoke() (*model.ListQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotasResponse), nil
	}
}

type CreateRoleInvoker struct {
	*invoker.BaseInvoker
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

func (i *DeleteRoleInvoker) Invoke() (*model.DeleteRoleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRoleResponse), nil
	}
}

type ListRoleNamesInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListRolesInvoker) Invoke() (*model.ListRolesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRolesResponse), nil
	}
}

type ShowRoleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRoleInvoker) Invoke() (*model.ShowRoleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRoleResponse), nil
	}
}

type UpdateRoleInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListLakeFormationInstanceTagsInvoker) Invoke() (*model.ListLakeFormationInstanceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLakeFormationInstanceTagsResponse), nil
	}
}

type ListGroupsForDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupsForDomainInvoker) Invoke() (*model.ListGroupsForDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupsForDomainResponse), nil
	}
}
