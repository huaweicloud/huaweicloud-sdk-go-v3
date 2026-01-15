package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dns/v2/model"
)

type AssociateEndpointIpaddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateEndpointIpaddressInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateEndpointIpaddressInvoker) Invoke() (*model.AssociateEndpointIpaddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateEndpointIpaddressResponse), nil
	}
}

type AssociateResolverQueryLogConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateResolverQueryLogConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateResolverQueryLogConfigInvoker) Invoke() (*model.AssociateResolverQueryLogConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateResolverQueryLogConfigResponse), nil
	}
}

type AssociateResolverRuleRouterInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateResolverRuleRouterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateResolverRuleRouterInvoker) Invoke() (*model.AssociateResolverRuleRouterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateResolverRuleRouterResponse), nil
	}
}

type AssociateRouterInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateRouterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateRouterInvoker) Invoke() (*model.AssociateRouterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateRouterResponse), nil
	}
}

type BatchCreateCombinedPublicRecordsetsTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateCombinedPublicRecordsetsTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateCombinedPublicRecordsetsTaskInvoker) Invoke() (*model.BatchCreateCombinedPublicRecordsetsTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateCombinedPublicRecordsetsTaskResponse), nil
	}
}

type BatchCreatePublicRecordsetsTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreatePublicRecordsetsTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreatePublicRecordsetsTaskInvoker) Invoke() (*model.BatchCreatePublicRecordsetsTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreatePublicRecordsetsTaskResponse), nil
	}
}

type BatchCreatePublicZonesTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreatePublicZonesTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreatePublicZonesTaskInvoker) Invoke() (*model.BatchCreatePublicZonesTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreatePublicZonesTaskResponse), nil
	}
}

type BatchCreateTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateTagInvoker) Invoke() (*model.BatchCreateTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateTagResponse), nil
	}
}

type BatchDeletePtrRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeletePtrRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeletePtrRecordsInvoker) Invoke() (*model.BatchDeletePtrRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeletePtrRecordsResponse), nil
	}
}

type BatchDeletePublicRecordsetsTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeletePublicRecordsetsTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeletePublicRecordsetsTaskInvoker) Invoke() (*model.BatchDeletePublicRecordsetsTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeletePublicRecordsetsTaskResponse), nil
	}
}

type BatchDeleteRecordSetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteRecordSetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteRecordSetsInvoker) Invoke() (*model.BatchDeleteRecordSetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteRecordSetsResponse), nil
	}
}

type BatchDeleteZonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteZonesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteZonesInvoker) Invoke() (*model.BatchDeleteZonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteZonesResponse), nil
	}
}

type BatchSetRecordSetsStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchSetRecordSetsStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchSetRecordSetsStatusInvoker) Invoke() (*model.BatchSetRecordSetsStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchSetRecordSetsStatusResponse), nil
	}
}

type BatchSetZonesStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchSetZonesStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchSetZonesStatusInvoker) Invoke() (*model.BatchSetZonesStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchSetZonesStatusResponse), nil
	}
}

type BatchTransferPublicZonesTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchTransferPublicZonesTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchTransferPublicZonesTaskInvoker) Invoke() (*model.BatchTransferPublicZonesTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchTransferPublicZonesTaskResponse), nil
	}
}

type BatchUpdatePublicRecordsetsTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdatePublicRecordsetsTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdatePublicRecordsetsTaskInvoker) Invoke() (*model.BatchUpdatePublicRecordsetsTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdatePublicRecordsetsTaskResponse), nil
	}
}

type CreateAuthorizeTxtRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAuthorizeTxtRecordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAuthorizeTxtRecordInvoker) Invoke() (*model.CreateAuthorizeTxtRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAuthorizeTxtRecordResponse), nil
	}
}

type CreateAuthorizeTxtRecordVerificationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAuthorizeTxtRecordVerificationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAuthorizeTxtRecordVerificationInvoker) Invoke() (*model.CreateAuthorizeTxtRecordVerificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAuthorizeTxtRecordVerificationResponse), nil
	}
}

type CreateCustomLineInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCustomLineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCustomLineInvoker) Invoke() (*model.CreateCustomLineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCustomLineResponse), nil
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

type CreateLineGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLineGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateLineGroupInvoker) Invoke() (*model.CreateLineGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLineGroupResponse), nil
	}
}

type CreatePrivateZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePrivateZoneInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePrivateZoneInvoker) Invoke() (*model.CreatePrivateZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePrivateZoneResponse), nil
	}
}

type CreatePublicZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePublicZoneInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePublicZoneInvoker) Invoke() (*model.CreatePublicZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePublicZoneResponse), nil
	}
}

type CreateResolverQueryLogConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResolverQueryLogConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateResolverQueryLogConfigInvoker) Invoke() (*model.CreateResolverQueryLogConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResolverQueryLogConfigResponse), nil
	}
}

type CreateResolverRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResolverRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateResolverRuleInvoker) Invoke() (*model.CreateResolverRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResolverRuleResponse), nil
	}
}

type CreateRetrievalInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRetrievalInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRetrievalInvoker) Invoke() (*model.CreateRetrievalResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRetrievalResponse), nil
	}
}

type CreateRetrievalVerificationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRetrievalVerificationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRetrievalVerificationInvoker) Invoke() (*model.CreateRetrievalVerificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRetrievalVerificationResponse), nil
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

type DeleteCustomLineInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCustomLineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCustomLineInvoker) Invoke() (*model.DeleteCustomLineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCustomLineResponse), nil
	}
}

type DeleteEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEndpointInvoker) Invoke() (*model.DeleteEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEndpointResponse), nil
	}
}

type DeleteLineGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLineGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteLineGroupInvoker) Invoke() (*model.DeleteLineGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLineGroupResponse), nil
	}
}

type DeletePrivateZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePrivateZoneInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePrivateZoneInvoker) Invoke() (*model.DeletePrivateZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePrivateZoneResponse), nil
	}
}

type DeletePublicZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePublicZoneInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePublicZoneInvoker) Invoke() (*model.DeletePublicZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePublicZoneResponse), nil
	}
}

type DeleteResolverQueryLogConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteResolverQueryLogConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteResolverQueryLogConfigInvoker) Invoke() (*model.DeleteResolverQueryLogConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteResolverQueryLogConfigResponse), nil
	}
}

type DeleteResolverRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteResolverRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteResolverRuleInvoker) Invoke() (*model.DeleteResolverRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteResolverRuleResponse), nil
	}
}

type DeleteTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTagInvoker) Invoke() (*model.DeleteTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTagResponse), nil
	}
}

type DisassociateEndpointIpaddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateEndpointIpaddressInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisassociateEndpointIpaddressInvoker) Invoke() (*model.DisassociateEndpointIpaddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateEndpointIpaddressResponse), nil
	}
}

type DisassociateResolverQueryLogConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateResolverQueryLogConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisassociateResolverQueryLogConfigInvoker) Invoke() (*model.DisassociateResolverQueryLogConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateResolverQueryLogConfigResponse), nil
	}
}

type DisassociateResolverRuleRouterInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateResolverRuleRouterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisassociateResolverRuleRouterInvoker) Invoke() (*model.DisassociateResolverRuleRouterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateResolverRuleRouterResponse), nil
	}
}

type DisassociateRouterInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateRouterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisassociateRouterInvoker) Invoke() (*model.DisassociateRouterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateRouterResponse), nil
	}
}

type ListApiVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApiVersionsInvoker) Invoke() (*model.ListApiVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApiVersionsResponse), nil
	}
}

type ListBatchOperationTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBatchOperationTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBatchOperationTasksInvoker) Invoke() (*model.ListBatchOperationTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBatchOperationTasksResponse), nil
	}
}

type ListCustomLineInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCustomLineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCustomLineInvoker) Invoke() (*model.ListCustomLineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCustomLineResponse), nil
	}
}

type ListEndpointIpaddressesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEndpointIpaddressesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEndpointIpaddressesInvoker) Invoke() (*model.ListEndpointIpaddressesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEndpointIpaddressesResponse), nil
	}
}

type ListEndpointVpcsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEndpointVpcsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEndpointVpcsInvoker) Invoke() (*model.ListEndpointVpcsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEndpointVpcsResponse), nil
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

type ListInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstancesInvoker) Invoke() (*model.ListInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstancesResponse), nil
	}
}

type ListLineGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLineGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLineGroupsInvoker) Invoke() (*model.ListLineGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLineGroupsResponse), nil
	}
}

type ListNameServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNameServersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNameServersInvoker) Invoke() (*model.ListNameServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNameServersResponse), nil
	}
}

type ListPrivateZonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPrivateZonesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPrivateZonesInvoker) Invoke() (*model.ListPrivateZonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPrivateZonesResponse), nil
	}
}

type ListPublicZonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPublicZonesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPublicZonesInvoker) Invoke() (*model.ListPublicZonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPublicZonesResponse), nil
	}
}

type ListResolverQueryLogConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResolverQueryLogConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResolverQueryLogConfigsInvoker) Invoke() (*model.ListResolverQueryLogConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResolverQueryLogConfigsResponse), nil
	}
}

type ListResolverRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResolverRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResolverRulesInvoker) Invoke() (*model.ListResolverRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResolverRulesResponse), nil
	}
}

type ListTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTagInvoker) Invoke() (*model.ListTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagResponse), nil
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

type SetPrivateZoneProxyPatternInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetPrivateZoneProxyPatternInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetPrivateZoneProxyPatternInvoker) Invoke() (*model.SetPrivateZoneProxyPatternResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetPrivateZoneProxyPatternResponse), nil
	}
}

type ShowApiInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApiInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowApiInfoInvoker) Invoke() (*model.ShowApiInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApiInfoResponse), nil
	}
}

type ShowAuthorizeTxtRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuthorizeTxtRecordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuthorizeTxtRecordInvoker) Invoke() (*model.ShowAuthorizeTxtRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuthorizeTxtRecordResponse), nil
	}
}

type ShowBatchOperationTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBatchOperationTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBatchOperationTaskInvoker) Invoke() (*model.ShowBatchOperationTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBatchOperationTaskResponse), nil
	}
}

type ShowDomainDetectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainDetectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDomainDetectionInvoker) Invoke() (*model.ShowDomainDetectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainDetectionResponse), nil
	}
}

type ShowDomainQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDomainQuotaInvoker) Invoke() (*model.ShowDomainQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainQuotaResponse), nil
	}
}

type ShowEmailRecordSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEmailRecordSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEmailRecordSetInvoker) Invoke() (*model.ShowEmailRecordSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEmailRecordSetResponse), nil
	}
}

type ShowEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEndpointInvoker) Invoke() (*model.ShowEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEndpointResponse), nil
	}
}

type ShowLineGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLineGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLineGroupInvoker) Invoke() (*model.ShowLineGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLineGroupResponse), nil
	}
}

type ShowPrivateZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPrivateZoneInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPrivateZoneInvoker) Invoke() (*model.ShowPrivateZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPrivateZoneResponse), nil
	}
}

type ShowPrivateZoneNameServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPrivateZoneNameServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPrivateZoneNameServerInvoker) Invoke() (*model.ShowPrivateZoneNameServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPrivateZoneNameServerResponse), nil
	}
}

type ShowPublicZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPublicZoneInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPublicZoneInvoker) Invoke() (*model.ShowPublicZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPublicZoneResponse), nil
	}
}

type ShowPublicZoneNameServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPublicZoneNameServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPublicZoneNameServerInvoker) Invoke() (*model.ShowPublicZoneNameServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPublicZoneNameServerResponse), nil
	}
}

type ShowResolverQueryLogConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResolverQueryLogConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResolverQueryLogConfigInvoker) Invoke() (*model.ShowResolverQueryLogConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResolverQueryLogConfigResponse), nil
	}
}

type ShowResolverRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResolverRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResolverRuleInvoker) Invoke() (*model.ShowResolverRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResolverRuleResponse), nil
	}
}

type ShowResourceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResourceTagInvoker) Invoke() (*model.ShowResourceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceTagResponse), nil
	}
}

type ShowRetrievalInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRetrievalInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRetrievalInvoker) Invoke() (*model.ShowRetrievalResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRetrievalResponse), nil
	}
}

type ShowRetrievalVerificationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRetrievalVerificationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRetrievalVerificationInvoker) Invoke() (*model.ShowRetrievalVerificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRetrievalVerificationResponse), nil
	}
}

type ShowWebsiteRecordSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWebsiteRecordSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWebsiteRecordSetInvoker) Invoke() (*model.ShowWebsiteRecordSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWebsiteRecordSetResponse), nil
	}
}

type ShowZoneNameServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowZoneNameServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowZoneNameServerInvoker) Invoke() (*model.ShowZoneNameServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowZoneNameServerResponse), nil
	}
}

type UpdateCustomLineInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCustomLineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCustomLineInvoker) Invoke() (*model.UpdateCustomLineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCustomLineResponse), nil
	}
}

type UpdateEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEndpointInvoker) Invoke() (*model.UpdateEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEndpointResponse), nil
	}
}

type UpdateLineGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLineGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateLineGroupsInvoker) Invoke() (*model.UpdateLineGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLineGroupsResponse), nil
	}
}

type UpdatePrivateZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePrivateZoneInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePrivateZoneInvoker) Invoke() (*model.UpdatePrivateZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePrivateZoneResponse), nil
	}
}

type UpdatePrivateZoneStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePrivateZoneStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePrivateZoneStatusInvoker) Invoke() (*model.UpdatePrivateZoneStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePrivateZoneStatusResponse), nil
	}
}

type UpdatePublicZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePublicZoneInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePublicZoneInvoker) Invoke() (*model.UpdatePublicZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePublicZoneResponse), nil
	}
}

type UpdatePublicZoneStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePublicZoneStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePublicZoneStatusInvoker) Invoke() (*model.UpdatePublicZoneStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePublicZoneStatusResponse), nil
	}
}

type UpdateResolverRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateResolverRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateResolverRuleInvoker) Invoke() (*model.UpdateResolverRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateResolverRuleResponse), nil
	}
}

type DisableDnssecConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableDnssecConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisableDnssecConfigInvoker) Invoke() (*model.DisableDnssecConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableDnssecConfigResponse), nil
	}
}

type EnableDnssecConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableDnssecConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableDnssecConfigInvoker) Invoke() (*model.EnableDnssecConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableDnssecConfigResponse), nil
	}
}

type ShowDnssecConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDnssecConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDnssecConfigInvoker) Invoke() (*model.ShowDnssecConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDnssecConfigResponse), nil
	}
}

type CreateEipRecordSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEipRecordSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEipRecordSetInvoker) Invoke() (*model.CreateEipRecordSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEipRecordSetResponse), nil
	}
}

type CreateRecordSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRecordSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRecordSetInvoker) Invoke() (*model.CreateRecordSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRecordSetResponse), nil
	}
}

type DeleteRecordSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRecordSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRecordSetInvoker) Invoke() (*model.DeleteRecordSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRecordSetResponse), nil
	}
}

type ListPtrRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPtrRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPtrRecordsInvoker) Invoke() (*model.ListPtrRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPtrRecordsResponse), nil
	}
}

type ListRecordSetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRecordSetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRecordSetsInvoker) Invoke() (*model.ListRecordSetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecordSetsResponse), nil
	}
}

type ListRecordSetsByZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRecordSetsByZoneInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRecordSetsByZoneInvoker) Invoke() (*model.ListRecordSetsByZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecordSetsByZoneResponse), nil
	}
}

type RestorePtrRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestorePtrRecordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RestorePtrRecordInvoker) Invoke() (*model.RestorePtrRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestorePtrRecordResponse), nil
	}
}

type ShowPtrRecordSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPtrRecordSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPtrRecordSetInvoker) Invoke() (*model.ShowPtrRecordSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPtrRecordSetResponse), nil
	}
}

type ShowRecordSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRecordSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRecordSetInvoker) Invoke() (*model.ShowRecordSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRecordSetResponse), nil
	}
}

type UpdatePtrRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePtrRecordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePtrRecordInvoker) Invoke() (*model.UpdatePtrRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePtrRecordResponse), nil
	}
}

type UpdateRecordSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRecordSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRecordSetInvoker) Invoke() (*model.UpdateRecordSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRecordSetResponse), nil
	}
}

type BatchCreateRecordSetsTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateRecordSetsTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateRecordSetsTaskInvoker) Invoke() (*model.BatchCreateRecordSetsTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateRecordSetsTaskResponse), nil
	}
}

type BatchDeleteRecordSetWithLineInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteRecordSetWithLineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteRecordSetWithLineInvoker) Invoke() (*model.BatchDeleteRecordSetWithLineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteRecordSetWithLineResponse), nil
	}
}

type BatchUpdateRecordSetWithLineInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateRecordSetWithLineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdateRecordSetWithLineInvoker) Invoke() (*model.BatchUpdateRecordSetWithLineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateRecordSetWithLineResponse), nil
	}
}

type CreatePtrInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePtrInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePtrInvoker) Invoke() (*model.CreatePtrResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePtrResponse), nil
	}
}

type CreateRecordSetWithBatchLinesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRecordSetWithBatchLinesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRecordSetWithBatchLinesInvoker) Invoke() (*model.CreateRecordSetWithBatchLinesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRecordSetWithBatchLinesResponse), nil
	}
}

type CreateRecordSetWithLineInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRecordSetWithLineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRecordSetWithLineInvoker) Invoke() (*model.CreateRecordSetWithLineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRecordSetWithLineResponse), nil
	}
}

type DeleteBatchCreateRecordSetsTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBatchCreateRecordSetsTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteBatchCreateRecordSetsTaskInvoker) Invoke() (*model.DeleteBatchCreateRecordSetsTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBatchCreateRecordSetsTaskResponse), nil
	}
}

type DeletePtrInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePtrInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePtrInvoker) Invoke() (*model.DeletePtrResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePtrResponse), nil
	}
}

type DeleteRecordSetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRecordSetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRecordSetsInvoker) Invoke() (*model.DeleteRecordSetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRecordSetsResponse), nil
	}
}

type ListPtrsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPtrsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPtrsInvoker) Invoke() (*model.ListPtrsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPtrsResponse), nil
	}
}

type ListPublicZoneLinesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPublicZoneLinesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPublicZoneLinesInvoker) Invoke() (*model.ListPublicZoneLinesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPublicZoneLinesResponse), nil
	}
}

type ListRecordSetsWithLineInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRecordSetsWithLineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRecordSetsWithLineInvoker) Invoke() (*model.ListRecordSetsWithLineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecordSetsWithLineResponse), nil
	}
}

type ListSystemLinesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSystemLinesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSystemLinesInvoker) Invoke() (*model.ListSystemLinesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSystemLinesResponse), nil
	}
}

type SetRecordSetsStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetRecordSetsStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetRecordSetsStatusInvoker) Invoke() (*model.SetRecordSetsStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetRecordSetsStatusResponse), nil
	}
}

type ShowBatchCreateRecordSetsTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBatchCreateRecordSetsTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBatchCreateRecordSetsTaskInvoker) Invoke() (*model.ShowBatchCreateRecordSetsTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBatchCreateRecordSetsTaskResponse), nil
	}
}

type ShowPtrInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPtrInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPtrInvoker) Invoke() (*model.ShowPtrResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPtrResponse), nil
	}
}

type ShowRecordSetByZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRecordSetByZoneInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRecordSetByZoneInvoker) Invoke() (*model.ShowRecordSetByZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRecordSetByZoneResponse), nil
	}
}

type ShowRecordSetWithLineInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRecordSetWithLineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRecordSetWithLineInvoker) Invoke() (*model.ShowRecordSetWithLineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRecordSetWithLineResponse), nil
	}
}

type UpdatePtrInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePtrInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePtrInvoker) Invoke() (*model.UpdatePtrResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePtrResponse), nil
	}
}

type UpdateRecordSetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRecordSetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRecordSetsInvoker) Invoke() (*model.UpdateRecordSetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRecordSetsResponse), nil
	}
}
