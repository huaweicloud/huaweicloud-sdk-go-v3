package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/roma/v2/model"
)

type AddSubsetsToGatewayInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddSubsetsToGatewayInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddSubsetsToGatewayInvoker) Invoke() (*model.AddSubsetsToGatewayResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddSubsetsToGatewayResponse), nil
	}
}

type AssociateAppsForAppQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateAppsForAppQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateAppsForAppQuotaInvoker) Invoke() (*model.AssociateAppsForAppQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateAppsForAppQuotaResponse), nil
	}
}

type AssociateCertificateV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateCertificateV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateCertificateV2Invoker) Invoke() (*model.AssociateCertificateV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateCertificateV2Response), nil
	}
}

type AssociateDomainV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateDomainV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateDomainV2Invoker) Invoke() (*model.AssociateDomainV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateDomainV2Response), nil
	}
}

type AssociateSignatureKeyV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateSignatureKeyV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateSignatureKeyV2Invoker) Invoke() (*model.AssociateSignatureKeyV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateSignatureKeyV2Response), nil
	}
}

type AttachApiToPluginInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachApiToPluginInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AttachApiToPluginInvoker) Invoke() (*model.AttachApiToPluginResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachApiToPluginResponse), nil
	}
}

type AttachPluginToApiInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachPluginToApiInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AttachPluginToApiInvoker) Invoke() (*model.AttachPluginToApiResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachPluginToApiResponse), nil
	}
}

type BatchAddDeviceToGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAddDeviceToGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchAddDeviceToGroupInvoker) Invoke() (*model.BatchAddDeviceToGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAddDeviceToGroupResponse), nil
	}
}

type BatchDeleteDevicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteDevicesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteDevicesInvoker) Invoke() (*model.BatchDeleteDevicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteDevicesResponse), nil
	}
}

type BatchDeleteMqsInstanceTopicInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteMqsInstanceTopicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteMqsInstanceTopicInvoker) Invoke() (*model.BatchDeleteMqsInstanceTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteMqsInstanceTopicResponse), nil
	}
}

type BatchDeleteRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteRulesInvoker) Invoke() (*model.BatchDeleteRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteRulesResponse), nil
	}
}

type BatchFreezeDevicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchFreezeDevicesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchFreezeDevicesInvoker) Invoke() (*model.BatchFreezeDevicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchFreezeDevicesResponse), nil
	}
}

type BatchStartOrStopTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchStartOrStopTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchStartOrStopTasksInvoker) Invoke() (*model.BatchStartOrStopTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchStartOrStopTasksResponse), nil
	}
}

type CheckLivedataApisV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CheckLivedataApisV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckLivedataApisV2Invoker) Invoke() (*model.CheckLivedataApisV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckLivedataApisV2Response), nil
	}
}

type CountDevicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountDevicesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountDevicesInvoker) Invoke() (*model.CountDevicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountDevicesResponse), nil
	}
}

type CountTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountTasksInvoker) Invoke() (*model.CountTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountTasksResponse), nil
	}
}

type CreateAppCodeAutoV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAppCodeAutoV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAppCodeAutoV2Invoker) Invoke() (*model.CreateAppCodeAutoV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAppCodeAutoV2Response), nil
	}
}

type CreateAppCodeV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAppCodeV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAppCodeV2Invoker) Invoke() (*model.CreateAppCodeV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAppCodeV2Response), nil
	}
}

type CreateAppConfigV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAppConfigV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAppConfigV2Invoker) Invoke() (*model.CreateAppConfigV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAppConfigV2Response), nil
	}
}

type CreateAppQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAppQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAppQuotaInvoker) Invoke() (*model.CreateAppQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAppQuotaResponse), nil
	}
}

type CreateCommandInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCommandInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCommandInvoker) Invoke() (*model.CreateCommandResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCommandResponse), nil
	}
}

type CreateCommonTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCommonTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCommonTaskInvoker) Invoke() (*model.CreateCommonTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCommonTaskResponse), nil
	}
}

type CreateCustomAuthorizerV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCustomAuthorizerV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCustomAuthorizerV2Invoker) Invoke() (*model.CreateCustomAuthorizerV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCustomAuthorizerV2Response), nil
	}
}

type CreateDatasourceInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDatasourceInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDatasourceInfoInvoker) Invoke() (*model.CreateDatasourceInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDatasourceInfoResponse), nil
	}
}

type CreateDestinationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDestinationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDestinationInvoker) Invoke() (*model.CreateDestinationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDestinationResponse), nil
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

type CreateDeviceGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDeviceGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDeviceGroupInvoker) Invoke() (*model.CreateDeviceGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDeviceGroupResponse), nil
	}
}

type CreateDispatchesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDispatchesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDispatchesInvoker) Invoke() (*model.CreateDispatchesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDispatchesResponse), nil
	}
}

type CreateEnvironmentV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEnvironmentV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEnvironmentV2Invoker) Invoke() (*model.CreateEnvironmentV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEnvironmentV2Response), nil
	}
}

type CreateEnvironmentVariableV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEnvironmentVariableV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEnvironmentVariableV2Invoker) Invoke() (*model.CreateEnvironmentVariableV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEnvironmentVariableV2Response), nil
	}
}

type CreateFeatureV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFeatureV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFeatureV2Invoker) Invoke() (*model.CreateFeatureV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFeatureV2Response), nil
	}
}

type CreateLiveDataApiScriptV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLiveDataApiScriptV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateLiveDataApiScriptV2Invoker) Invoke() (*model.CreateLiveDataApiScriptV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLiveDataApiScriptV2Response), nil
	}
}

type CreateLiveDataApiV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLiveDataApiV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateLiveDataApiV2Invoker) Invoke() (*model.CreateLiveDataApiV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLiveDataApiV2Response), nil
	}
}

type CreateMqsInstanceTopicInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMqsInstanceTopicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMqsInstanceTopicInvoker) Invoke() (*model.CreateMqsInstanceTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMqsInstanceTopicResponse), nil
	}
}

type CreateMultiTaskMappingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMultiTaskMappingsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMultiTaskMappingsInvoker) Invoke() (*model.CreateMultiTaskMappingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMultiTaskMappingsResponse), nil
	}
}

type CreateMultiTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMultiTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMultiTasksInvoker) Invoke() (*model.CreateMultiTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMultiTasksResponse), nil
	}
}

type CreateNotificationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNotificationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateNotificationInvoker) Invoke() (*model.CreateNotificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNotificationResponse), nil
	}
}

type CreatePluginInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePluginInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePluginInvoker) Invoke() (*model.CreatePluginResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePluginResponse), nil
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

type CreateProductTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateProductTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateProductTemplateInvoker) Invoke() (*model.CreateProductTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateProductTemplateResponse), nil
	}
}

type CreateProductTopicInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateProductTopicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateProductTopicInvoker) Invoke() (*model.CreateProductTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateProductTopicResponse), nil
	}
}

type CreatePropertyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePropertyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePropertyInvoker) Invoke() (*model.CreatePropertyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePropertyResponse), nil
	}
}

type CreateRequestPropertyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRequestPropertyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRequestPropertyInvoker) Invoke() (*model.CreateRequestPropertyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRequestPropertyResponse), nil
	}
}

type CreateRequestThrottlingPolicyV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRequestThrottlingPolicyV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRequestThrottlingPolicyV2Invoker) Invoke() (*model.CreateRequestThrottlingPolicyV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRequestThrottlingPolicyV2Response), nil
	}
}

type CreateResponsePropertyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResponsePropertyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateResponsePropertyInvoker) Invoke() (*model.CreateResponsePropertyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResponsePropertyResponse), nil
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

type CreateSignatureKeyV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSignatureKeyV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSignatureKeyV2Invoker) Invoke() (*model.CreateSignatureKeyV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSignatureKeyV2Response), nil
	}
}

type CreateSourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSourceInvoker) Invoke() (*model.CreateSourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSourceResponse), nil
	}
}

type CreateSpecialThrottlingConfigurationV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSpecialThrottlingConfigurationV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSpecialThrottlingConfigurationV2Invoker) Invoke() (*model.CreateSpecialThrottlingConfigurationV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSpecialThrottlingConfigurationV2Response), nil
	}
}

type DebugLiveDataApiV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DebugLiveDataApiV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DebugLiveDataApiV2Invoker) Invoke() (*model.DebugLiveDataApiV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DebugLiveDataApiV2Response), nil
	}
}

type DebugRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DebugRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DebugRuleInvoker) Invoke() (*model.DebugRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DebugRuleResponse), nil
	}
}

type DeleteAppAclInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAppAclInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAppAclInvoker) Invoke() (*model.DeleteAppAclResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAppAclResponse), nil
	}
}

type DeleteAppCodeV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAppCodeV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAppCodeV2Invoker) Invoke() (*model.DeleteAppCodeV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAppCodeV2Response), nil
	}
}

type DeleteAppConfigV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAppConfigV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAppConfigV2Invoker) Invoke() (*model.DeleteAppConfigV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAppConfigV2Response), nil
	}
}

type DeleteAppQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAppQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAppQuotaInvoker) Invoke() (*model.DeleteAppQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAppQuotaResponse), nil
	}
}

type DeleteCommandInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCommandInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCommandInvoker) Invoke() (*model.DeleteCommandResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCommandResponse), nil
	}
}

type DeleteCustomAuthorizerV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCustomAuthorizerV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCustomAuthorizerV2Invoker) Invoke() (*model.DeleteCustomAuthorizerV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCustomAuthorizerV2Response), nil
	}
}

type DeleteDatasourceInfoByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDatasourceInfoByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDatasourceInfoByIdInvoker) Invoke() (*model.DeleteDatasourceInfoByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDatasourceInfoByIdResponse), nil
	}
}

type DeleteDestinationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDestinationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDestinationInvoker) Invoke() (*model.DeleteDestinationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDestinationResponse), nil
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

type DeleteDeviceFromGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDeviceFromGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDeviceFromGroupInvoker) Invoke() (*model.DeleteDeviceFromGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDeviceFromGroupResponse), nil
	}
}

type DeleteDeviceGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDeviceGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDeviceGroupInvoker) Invoke() (*model.DeleteDeviceGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDeviceGroupResponse), nil
	}
}

type DeleteEnvironmentV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEnvironmentV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEnvironmentV2Invoker) Invoke() (*model.DeleteEnvironmentV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEnvironmentV2Response), nil
	}
}

type DeleteEnvironmentVariableV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEnvironmentVariableV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEnvironmentVariableV2Invoker) Invoke() (*model.DeleteEnvironmentVariableV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEnvironmentVariableV2Response), nil
	}
}

type DeleteLiveDataApiV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLiveDataApiV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteLiveDataApiV2Invoker) Invoke() (*model.DeleteLiveDataApiV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLiveDataApiV2Response), nil
	}
}

type DeleteMqsInstanceTopicInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMqsInstanceTopicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMqsInstanceTopicInvoker) Invoke() (*model.DeleteMqsInstanceTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMqsInstanceTopicResponse), nil
	}
}

type DeleteMultiTaskMappingInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMultiTaskMappingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMultiTaskMappingInvoker) Invoke() (*model.DeleteMultiTaskMappingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMultiTaskMappingResponse), nil
	}
}

type DeleteNotificationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNotificationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteNotificationInvoker) Invoke() (*model.DeleteNotificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNotificationResponse), nil
	}
}

type DeletePluginInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePluginInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePluginInvoker) Invoke() (*model.DeletePluginResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePluginResponse), nil
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

type DeleteProductTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteProductTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteProductTemplateInvoker) Invoke() (*model.DeleteProductTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteProductTemplateResponse), nil
	}
}

type DeleteProductTopicInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteProductTopicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteProductTopicInvoker) Invoke() (*model.DeleteProductTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteProductTopicResponse), nil
	}
}

type DeletePropertyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePropertyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePropertyInvoker) Invoke() (*model.DeletePropertyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePropertyResponse), nil
	}
}

type DeleteRequestPropertyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRequestPropertyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRequestPropertyInvoker) Invoke() (*model.DeleteRequestPropertyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRequestPropertyResponse), nil
	}
}

type DeleteRequestThrottlingPolicyV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRequestThrottlingPolicyV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRequestThrottlingPolicyV2Invoker) Invoke() (*model.DeleteRequestThrottlingPolicyV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRequestThrottlingPolicyV2Response), nil
	}
}

type DeleteResponsePropertyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteResponsePropertyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteResponsePropertyInvoker) Invoke() (*model.DeleteResponsePropertyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteResponsePropertyResponse), nil
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

type DeleteSignatureKeyV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSignatureKeyV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSignatureKeyV2Invoker) Invoke() (*model.DeleteSignatureKeyV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSignatureKeyV2Response), nil
	}
}

type DeleteSourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSourceInvoker) Invoke() (*model.DeleteSourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSourceResponse), nil
	}
}

type DeleteSpecialThrottlingConfigurationV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSpecialThrottlingConfigurationV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSpecialThrottlingConfigurationV2Invoker) Invoke() (*model.DeleteSpecialThrottlingConfigurationV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSpecialThrottlingConfigurationV2Response), nil
	}
}

type DeleteTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTaskInvoker) Invoke() (*model.DeleteTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTaskResponse), nil
	}
}

type DetachApiFromPluginInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetachApiFromPluginInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetachApiFromPluginInvoker) Invoke() (*model.DetachApiFromPluginResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetachApiFromPluginResponse), nil
	}
}

type DetachPluginFromApiInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetachPluginFromApiInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetachPluginFromApiInvoker) Invoke() (*model.DetachPluginFromApiResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetachPluginFromApiResponse), nil
	}
}

type DisassociateAppQuotaWithAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateAppQuotaWithAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisassociateAppQuotaWithAppInvoker) Invoke() (*model.DisassociateAppQuotaWithAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateAppQuotaWithAppResponse), nil
	}
}

type DisassociateCertificateV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateCertificateV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisassociateCertificateV2Invoker) Invoke() (*model.DisassociateCertificateV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateCertificateV2Response), nil
	}
}

type DisassociateDomainV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateDomainV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisassociateDomainV2Invoker) Invoke() (*model.DisassociateDomainV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateDomainV2Response), nil
	}
}

type DisassociateSignatureKeyV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateSignatureKeyV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisassociateSignatureKeyV2Invoker) Invoke() (*model.DisassociateSignatureKeyV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateSignatureKeyV2Response), nil
	}
}

type DownloadProductsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadProductsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadProductsInvoker) Invoke() (*model.DownloadProductsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadProductsResponse), nil
	}
}

type ExportMqsInstanceTopicInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportMqsInstanceTopicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportMqsInstanceTopicInvoker) Invoke() (*model.ExportMqsInstanceTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportMqsInstanceTopicResponse), nil
	}
}

type ImportMqsInstanceTopicInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportMqsInstanceTopicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportMqsInstanceTopicInvoker) Invoke() (*model.ImportMqsInstanceTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportMqsInstanceTopicResponse), nil
	}
}

type InstallMultiTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *InstallMultiTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *InstallMultiTasksInvoker) Invoke() (*model.InstallMultiTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstallMultiTasksResponse), nil
	}
}

type ListApiAttachablePluginsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiAttachablePluginsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApiAttachablePluginsInvoker) Invoke() (*model.ListApiAttachablePluginsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApiAttachablePluginsResponse), nil
	}
}

type ListApiAttachedPluginsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiAttachedPluginsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApiAttachedPluginsInvoker) Invoke() (*model.ListApiAttachedPluginsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApiAttachedPluginsResponse), nil
	}
}

type ListApisBindedToSignatureKeyV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListApisBindedToSignatureKeyV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApisBindedToSignatureKeyV2Invoker) Invoke() (*model.ListApisBindedToSignatureKeyV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApisBindedToSignatureKeyV2Response), nil
	}
}

type ListApisNotBoundWithSignatureKeyV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListApisNotBoundWithSignatureKeyV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApisNotBoundWithSignatureKeyV2Invoker) Invoke() (*model.ListApisNotBoundWithSignatureKeyV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApisNotBoundWithSignatureKeyV2Response), nil
	}
}

type ListAppCodesV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppCodesV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppCodesV2Invoker) Invoke() (*model.ListAppCodesV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppCodesV2Response), nil
	}
}

type ListAppConfigsV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppConfigsV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppConfigsV2Invoker) Invoke() (*model.ListAppConfigsV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppConfigsV2Response), nil
	}
}

type ListAppQuotaBindableAppsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppQuotaBindableAppsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppQuotaBindableAppsInvoker) Invoke() (*model.ListAppQuotaBindableAppsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppQuotaBindableAppsResponse), nil
	}
}

type ListAppQuotaBoundAppsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppQuotaBoundAppsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppQuotaBoundAppsInvoker) Invoke() (*model.ListAppQuotaBoundAppsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppQuotaBoundAppsResponse), nil
	}
}

type ListAppQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppQuotasInvoker) Invoke() (*model.ListAppQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppQuotasResponse), nil
	}
}

type ListAppsV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppsV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppsV2Invoker) Invoke() (*model.ListAppsV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppsV2Response), nil
	}
}

type ListCommandsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCommandsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCommandsInvoker) Invoke() (*model.ListCommandsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCommandsResponse), nil
	}
}

type ListCustomAuthorizersV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListCustomAuthorizersV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCustomAuthorizersV2Invoker) Invoke() (*model.ListCustomAuthorizersV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCustomAuthorizersV2Response), nil
	}
}

type ListDatasourceColumnsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatasourceColumnsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDatasourceColumnsInvoker) Invoke() (*model.ListDatasourceColumnsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatasourceColumnsResponse), nil
	}
}

type ListDatasourceTablesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatasourceTablesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDatasourceTablesInvoker) Invoke() (*model.ListDatasourceTablesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatasourceTablesResponse), nil
	}
}

type ListDatasourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatasourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDatasourcesInvoker) Invoke() (*model.ListDatasourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatasourcesResponse), nil
	}
}

type ListDestinationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDestinationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDestinationsInvoker) Invoke() (*model.ListDestinationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDestinationsResponse), nil
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

type ListDevicesInProductInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDevicesInProductInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDevicesInProductInvoker) Invoke() (*model.ListDevicesInProductResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDevicesInProductResponse), nil
	}
}

type ListEnvironmentVariablesV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListEnvironmentVariablesV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEnvironmentVariablesV2Invoker) Invoke() (*model.ListEnvironmentVariablesV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnvironmentVariablesV2Response), nil
	}
}

type ListEnvironmentsV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListEnvironmentsV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEnvironmentsV2Invoker) Invoke() (*model.ListEnvironmentsV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnvironmentsV2Response), nil
	}
}

type ListFeaturesV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListFeaturesV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFeaturesV2Invoker) Invoke() (*model.ListFeaturesV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFeaturesV2Response), nil
	}
}

type ListLatelyApiStatisticsV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListLatelyApiStatisticsV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLatelyApiStatisticsV2Invoker) Invoke() (*model.ListLatelyApiStatisticsV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLatelyApiStatisticsV2Response), nil
	}
}

type ListLiveDataApiDeploymentHistoryV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListLiveDataApiDeploymentHistoryV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLiveDataApiDeploymentHistoryV2Invoker) Invoke() (*model.ListLiveDataApiDeploymentHistoryV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLiveDataApiDeploymentHistoryV2Response), nil
	}
}

type ListLiveDataApiTestHistoryV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListLiveDataApiTestHistoryV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLiveDataApiTestHistoryV2Invoker) Invoke() (*model.ListLiveDataApiTestHistoryV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLiveDataApiTestHistoryV2Response), nil
	}
}

type ListLiveDataApiV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListLiveDataApiV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLiveDataApiV2Invoker) Invoke() (*model.ListLiveDataApiV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLiveDataApiV2Response), nil
	}
}

type ListLiveDataDataSourcesV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListLiveDataDataSourcesV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLiveDataDataSourcesV2Invoker) Invoke() (*model.ListLiveDataDataSourcesV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLiveDataDataSourcesV2Response), nil
	}
}

type ListLiveDataQuotaV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListLiveDataQuotaV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLiveDataQuotaV2Invoker) Invoke() (*model.ListLiveDataQuotaV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLiveDataQuotaV2Response), nil
	}
}

type ListMonitorInfosInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMonitorInfosInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMonitorInfosInvoker) Invoke() (*model.ListMonitorInfosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMonitorInfosResponse), nil
	}
}

type ListMonitorLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMonitorLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMonitorLogInvoker) Invoke() (*model.ListMonitorLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMonitorLogResponse), nil
	}
}

type ListMqsInstanceTopicsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMqsInstanceTopicsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMqsInstanceTopicsInvoker) Invoke() (*model.ListMqsInstanceTopicsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMqsInstanceTopicsResponse), nil
	}
}

type ListNotificationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNotificationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNotificationInvoker) Invoke() (*model.ListNotificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNotificationResponse), nil
	}
}

type ListPluginAttachableApisInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPluginAttachableApisInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPluginAttachableApisInvoker) Invoke() (*model.ListPluginAttachableApisResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPluginAttachableApisResponse), nil
	}
}

type ListPluginAttachedApisInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPluginAttachedApisInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPluginAttachedApisInvoker) Invoke() (*model.ListPluginAttachedApisResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPluginAttachedApisResponse), nil
	}
}

type ListPluginsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPluginsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPluginsInvoker) Invoke() (*model.ListPluginsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPluginsResponse), nil
	}
}

type ListProductTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProductTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProductTemplatesInvoker) Invoke() (*model.ListProductTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProductTemplatesResponse), nil
	}
}

type ListProductTopicsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProductTopicsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProductTopicsInvoker) Invoke() (*model.ListProductTopicsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProductTopicsResponse), nil
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

type ListProjectCofigsV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectCofigsV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectCofigsV2Invoker) Invoke() (*model.ListProjectCofigsV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectCofigsV2Response), nil
	}
}

type ListPropertiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPropertiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPropertiesInvoker) Invoke() (*model.ListPropertiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPropertiesResponse), nil
	}
}

type ListRequestPropertiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRequestPropertiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRequestPropertiesInvoker) Invoke() (*model.ListRequestPropertiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRequestPropertiesResponse), nil
	}
}

type ListRequestThrottlingPolicyV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListRequestThrottlingPolicyV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRequestThrottlingPolicyV2Invoker) Invoke() (*model.ListRequestThrottlingPolicyV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRequestThrottlingPolicyV2Response), nil
	}
}

type ListResponsePropertiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResponsePropertiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResponsePropertiesInvoker) Invoke() (*model.ListResponsePropertiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResponsePropertiesResponse), nil
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

type ListShadowsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListShadowsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListShadowsInvoker) Invoke() (*model.ListShadowsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListShadowsResponse), nil
	}
}

type ListSignatureKeysBindedToApiV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListSignatureKeysBindedToApiV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSignatureKeysBindedToApiV2Invoker) Invoke() (*model.ListSignatureKeysBindedToApiV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSignatureKeysBindedToApiV2Response), nil
	}
}

type ListSignatureKeysV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListSignatureKeysV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSignatureKeysV2Invoker) Invoke() (*model.ListSignatureKeysV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSignatureKeysV2Response), nil
	}
}

type ListSourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSourcesInvoker) Invoke() (*model.ListSourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSourcesResponse), nil
	}
}

type ListSpecialThrottlingConfigurationsV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListSpecialThrottlingConfigurationsV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSpecialThrottlingConfigurationsV2Invoker) Invoke() (*model.ListSpecialThrottlingConfigurationsV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSpecialThrottlingConfigurationsV2Response), nil
	}
}

type ListStatisticsApiInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStatisticsApiInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStatisticsApiInvoker) Invoke() (*model.ListStatisticsApiResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStatisticsApiResponse), nil
	}
}

type ListSubsetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubsetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSubsetsInvoker) Invoke() (*model.ListSubsetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubsetsResponse), nil
	}
}

type ListTagsV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagsV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTagsV2Invoker) Invoke() (*model.ListTagsV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagsV2Response), nil
	}
}

type ListTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTasksInvoker) Invoke() (*model.ListTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTasksResponse), nil
	}
}

type ListTopicsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTopicsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTopicsInvoker) Invoke() (*model.ListTopicsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTopicsResponse), nil
	}
}

type PublishLiveDataApiV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *PublishLiveDataApiV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PublishLiveDataApiV2Invoker) Invoke() (*model.PublishLiveDataApiV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishLiveDataApiV2Response), nil
	}
}

type ResetAuthenticationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetAuthenticationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetAuthenticationInvoker) Invoke() (*model.ResetAuthenticationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetAuthenticationResponse), nil
	}
}

type ResetMessagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetMessagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetMessagesInvoker) Invoke() (*model.ResetMessagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetMessagesResponse), nil
	}
}

type ResetMultiTaskOffsetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetMultiTaskOffsetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetMultiTaskOffsetInvoker) Invoke() (*model.ResetMultiTaskOffsetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetMultiTaskOffsetResponse), nil
	}
}

type ResetProductAuthenticationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetProductAuthenticationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetProductAuthenticationInvoker) Invoke() (*model.ResetProductAuthenticationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetProductAuthenticationResponse), nil
	}
}

type RunTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunTaskInvoker) Invoke() (*model.RunTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunTaskResponse), nil
	}
}

type SendCommandInvoker struct {
	*invoker.BaseInvoker
}

func (i *SendCommandInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SendCommandInvoker) Invoke() (*model.SendCommandResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendCommandResponse), nil
	}
}

type ShowAppBoundAppQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppBoundAppQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAppBoundAppQuotaInvoker) Invoke() (*model.ShowAppBoundAppQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppBoundAppQuotaResponse), nil
	}
}

type ShowAppQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAppQuotaInvoker) Invoke() (*model.ShowAppQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppQuotaResponse), nil
	}
}

type ShowAuthenticationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuthenticationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuthenticationInvoker) Invoke() (*model.ShowAuthenticationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuthenticationResponse), nil
	}
}

type ShowCommandInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCommandInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCommandInvoker) Invoke() (*model.ShowCommandResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCommandResponse), nil
	}
}

type ShowDataourceDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDataourceDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDataourceDetailInvoker) Invoke() (*model.ShowDataourceDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDataourceDetailResponse), nil
	}
}

type ShowDetailsOfAppAclInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailsOfAppAclInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDetailsOfAppAclInvoker) Invoke() (*model.ShowDetailsOfAppAclResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailsOfAppAclResponse), nil
	}
}

type ShowDetailsOfAppCodeV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailsOfAppCodeV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDetailsOfAppCodeV2Invoker) Invoke() (*model.ShowDetailsOfAppCodeV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailsOfAppCodeV2Response), nil
	}
}

type ShowDetailsOfAppConfigV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailsOfAppConfigV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDetailsOfAppConfigV2Invoker) Invoke() (*model.ShowDetailsOfAppConfigV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailsOfAppConfigV2Response), nil
	}
}

type ShowDetailsOfAppV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailsOfAppV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDetailsOfAppV2Invoker) Invoke() (*model.ShowDetailsOfAppV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailsOfAppV2Response), nil
	}
}

type ShowDetailsOfCustomAuthorizersV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailsOfCustomAuthorizersV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDetailsOfCustomAuthorizersV2Invoker) Invoke() (*model.ShowDetailsOfCustomAuthorizersV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailsOfCustomAuthorizersV2Response), nil
	}
}

type ShowDetailsOfDomainNameCertificateV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailsOfDomainNameCertificateV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDetailsOfDomainNameCertificateV2Invoker) Invoke() (*model.ShowDetailsOfDomainNameCertificateV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailsOfDomainNameCertificateV2Response), nil
	}
}

type ShowDetailsOfEnvironmentVariableV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailsOfEnvironmentVariableV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDetailsOfEnvironmentVariableV2Invoker) Invoke() (*model.ShowDetailsOfEnvironmentVariableV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailsOfEnvironmentVariableV2Response), nil
	}
}

type ShowDetailsOfInstanceV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailsOfInstanceV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDetailsOfInstanceV2Invoker) Invoke() (*model.ShowDetailsOfInstanceV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailsOfInstanceV2Response), nil
	}
}

type ShowDetailsOfRequestThrottlingPolicyV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailsOfRequestThrottlingPolicyV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDetailsOfRequestThrottlingPolicyV2Invoker) Invoke() (*model.ShowDetailsOfRequestThrottlingPolicyV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailsOfRequestThrottlingPolicyV2Response), nil
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

type ShowDeviceGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeviceGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDeviceGroupInvoker) Invoke() (*model.ShowDeviceGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeviceGroupResponse), nil
	}
}

type ShowDeviceGroupTreeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeviceGroupTreeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDeviceGroupTreeInvoker) Invoke() (*model.ShowDeviceGroupTreeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeviceGroupTreeResponse), nil
	}
}

type ShowDevicesInGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDevicesInGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDevicesInGroupInvoker) Invoke() (*model.ShowDevicesInGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDevicesInGroupResponse), nil
	}
}

type ShowDispatchesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDispatchesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDispatchesInvoker) Invoke() (*model.ShowDispatchesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDispatchesResponse), nil
	}
}

type ShowLiveDataApiV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLiveDataApiV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLiveDataApiV2Invoker) Invoke() (*model.ShowLiveDataApiV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLiveDataApiV2Response), nil
	}
}

type ShowMqsInstanceMessagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMqsInstanceMessagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMqsInstanceMessagesInvoker) Invoke() (*model.ShowMqsInstanceMessagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMqsInstanceMessagesResponse), nil
	}
}

type ShowMqsInstanceTopicAccessPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMqsInstanceTopicAccessPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMqsInstanceTopicAccessPolicyInvoker) Invoke() (*model.ShowMqsInstanceTopicAccessPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMqsInstanceTopicAccessPolicyResponse), nil
	}
}

type ShowPluginInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPluginInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPluginInvoker) Invoke() (*model.ShowPluginResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPluginResponse), nil
	}
}

type ShowProductInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProductInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProductInvoker) Invoke() (*model.ShowProductResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProductResponse), nil
	}
}

type ShowProductAuthenticationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProductAuthenticationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProductAuthenticationInvoker) Invoke() (*model.ShowProductAuthenticationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProductAuthenticationResponse), nil
	}
}

type ShowProductTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProductTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProductTemplateInvoker) Invoke() (*model.ShowProductTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProductTemplateResponse), nil
	}
}

type ShowPropertyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPropertyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPropertyInvoker) Invoke() (*model.ShowPropertyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPropertyResponse), nil
	}
}

type ShowRequestPropertyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRequestPropertyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRequestPropertyInvoker) Invoke() (*model.ShowRequestPropertyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRequestPropertyResponse), nil
	}
}

type ShowResponsePropertyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResponsePropertyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResponsePropertyInvoker) Invoke() (*model.ShowResponsePropertyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResponsePropertyResponse), nil
	}
}

type ShowRestrictionOfInstanceV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRestrictionOfInstanceV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRestrictionOfInstanceV2Invoker) Invoke() (*model.ShowRestrictionOfInstanceV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRestrictionOfInstanceV2Response), nil
	}
}

type ShowRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRuleInvoker) Invoke() (*model.ShowRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRuleResponse), nil
	}
}

type ShowServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowServiceInvoker) Invoke() (*model.ShowServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServiceResponse), nil
	}
}

type ShowTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTaskInvoker) Invoke() (*model.ShowTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskResponse), nil
	}
}

type StartTestDatasourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartTestDatasourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartTestDatasourceInvoker) Invoke() (*model.StartTestDatasourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartTestDatasourceResponse), nil
	}
}

type StopTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopTaskInvoker) Invoke() (*model.StopTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopTaskResponse), nil
	}
}

type UnpublishLiveDataApiV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *UnpublishLiveDataApiV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UnpublishLiveDataApiV2Invoker) Invoke() (*model.UnpublishLiveDataApiV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnpublishLiveDataApiV2Response), nil
	}
}

type UpdateAppAclInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAppAclInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAppAclInvoker) Invoke() (*model.UpdateAppAclResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAppAclResponse), nil
	}
}

type UpdateAppConfigV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAppConfigV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAppConfigV2Invoker) Invoke() (*model.UpdateAppConfigV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAppConfigV2Response), nil
	}
}

type UpdateAppQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAppQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAppQuotaInvoker) Invoke() (*model.UpdateAppQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAppQuotaResponse), nil
	}
}

type UpdateCommandInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCommandInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCommandInvoker) Invoke() (*model.UpdateCommandResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCommandResponse), nil
	}
}

type UpdateCustomAuthorizerV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCustomAuthorizerV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCustomAuthorizerV2Invoker) Invoke() (*model.UpdateCustomAuthorizerV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCustomAuthorizerV2Response), nil
	}
}

type UpdateDatasourceInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDatasourceInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDatasourceInfoInvoker) Invoke() (*model.UpdateDatasourceInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDatasourceInfoResponse), nil
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

type UpdateDeviceGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDeviceGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDeviceGroupInvoker) Invoke() (*model.UpdateDeviceGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDeviceGroupResponse), nil
	}
}

type UpdateDispatchesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDispatchesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDispatchesInvoker) Invoke() (*model.UpdateDispatchesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDispatchesResponse), nil
	}
}

type UpdateDomainV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDomainV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDomainV2Invoker) Invoke() (*model.UpdateDomainV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDomainV2Response), nil
	}
}

type UpdateEnvironmentV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEnvironmentV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEnvironmentV2Invoker) Invoke() (*model.UpdateEnvironmentV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEnvironmentV2Response), nil
	}
}

type UpdateEnvironmentVariableV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEnvironmentVariableV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEnvironmentVariableV2Invoker) Invoke() (*model.UpdateEnvironmentVariableV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEnvironmentVariableV2Response), nil
	}
}

type UpdateLiveDataApiV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLiveDataApiV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateLiveDataApiV2Invoker) Invoke() (*model.UpdateLiveDataApiV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLiveDataApiV2Response), nil
	}
}

type UpdateMqsInstanceTopicInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMqsInstanceTopicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMqsInstanceTopicInvoker) Invoke() (*model.UpdateMqsInstanceTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMqsInstanceTopicResponse), nil
	}
}

type UpdateMultiTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMultiTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMultiTasksInvoker) Invoke() (*model.UpdateMultiTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMultiTasksResponse), nil
	}
}

type UpdateNotificationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNotificationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateNotificationInvoker) Invoke() (*model.UpdateNotificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNotificationResponse), nil
	}
}

type UpdatePluginInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePluginInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePluginInvoker) Invoke() (*model.UpdatePluginResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePluginResponse), nil
	}
}

type UpdateProductInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProductInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProductInvoker) Invoke() (*model.UpdateProductResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProductResponse), nil
	}
}

type UpdateProductTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProductTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProductTemplateInvoker) Invoke() (*model.UpdateProductTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProductTemplateResponse), nil
	}
}

type UpdateProductTopicInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProductTopicInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProductTopicInvoker) Invoke() (*model.UpdateProductTopicResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProductTopicResponse), nil
	}
}

type UpdatePropertyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePropertyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePropertyInvoker) Invoke() (*model.UpdatePropertyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePropertyResponse), nil
	}
}

type UpdateRequestPropertyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRequestPropertyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRequestPropertyInvoker) Invoke() (*model.UpdateRequestPropertyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRequestPropertyResponse), nil
	}
}

type UpdateRequestThrottlingPolicyV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRequestThrottlingPolicyV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRequestThrottlingPolicyV2Invoker) Invoke() (*model.UpdateRequestThrottlingPolicyV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRequestThrottlingPolicyV2Response), nil
	}
}

type UpdateResponsePropertyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateResponsePropertyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateResponsePropertyInvoker) Invoke() (*model.UpdateResponsePropertyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateResponsePropertyResponse), nil
	}
}

type UpdateRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRuleInvoker) Invoke() (*model.UpdateRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRuleResponse), nil
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

type UpdateSignatureKeyV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSignatureKeyV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSignatureKeyV2Invoker) Invoke() (*model.UpdateSignatureKeyV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSignatureKeyV2Response), nil
	}
}

type UpdateSpecialThrottlingConfigurationV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSpecialThrottlingConfigurationV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSpecialThrottlingConfigurationV2Invoker) Invoke() (*model.UpdateSpecialThrottlingConfigurationV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSpecialThrottlingConfigurationV2Response), nil
	}
}

type UpdateTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTaskInvoker) Invoke() (*model.UpdateTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTaskResponse), nil
	}
}

type UpdateTopicAccessPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTopicAccessPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTopicAccessPolicyInvoker) Invoke() (*model.UpdateTopicAccessPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTopicAccessPolicyResponse), nil
	}
}

type UploadProductInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadProductInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UploadProductInvoker) Invoke() (*model.UploadProductResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadProductResponse), nil
	}
}

type BatchDeleteAclV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteAclV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteAclV2Invoker) Invoke() (*model.BatchDeleteAclV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteAclV2Response), nil
	}
}

type CreateAclStrategyV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAclStrategyV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAclStrategyV2Invoker) Invoke() (*model.CreateAclStrategyV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAclStrategyV2Response), nil
	}
}

type DeleteAclV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAclV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAclV2Invoker) Invoke() (*model.DeleteAclV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAclV2Response), nil
	}
}

type ListAclStrategiesV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListAclStrategiesV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAclStrategiesV2Invoker) Invoke() (*model.ListAclStrategiesV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAclStrategiesV2Response), nil
	}
}

type ShowDetailsOfAclPolicyV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailsOfAclPolicyV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDetailsOfAclPolicyV2Invoker) Invoke() (*model.ShowDetailsOfAclPolicyV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailsOfAclPolicyV2Response), nil
	}
}

type UpdateAclStrategyV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAclStrategyV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAclStrategyV2Invoker) Invoke() (*model.UpdateAclStrategyV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAclStrategyV2Response), nil
	}
}

type AssociateRequestThrottlingPolicyV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateRequestThrottlingPolicyV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateRequestThrottlingPolicyV2Invoker) Invoke() (*model.AssociateRequestThrottlingPolicyV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateRequestThrottlingPolicyV2Response), nil
	}
}

type BatchDisassociateThrottlingPolicyV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDisassociateThrottlingPolicyV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDisassociateThrottlingPolicyV2Invoker) Invoke() (*model.BatchDisassociateThrottlingPolicyV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDisassociateThrottlingPolicyV2Response), nil
	}
}

type BatchPublishOrOfflineApiV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *BatchPublishOrOfflineApiV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchPublishOrOfflineApiV2Invoker) Invoke() (*model.BatchPublishOrOfflineApiV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchPublishOrOfflineApiV2Response), nil
	}
}

type ChangeApiVersionV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeApiVersionV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeApiVersionV2Invoker) Invoke() (*model.ChangeApiVersionV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeApiVersionV2Response), nil
	}
}

type CheckApiGroupsV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CheckApiGroupsV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckApiGroupsV2Invoker) Invoke() (*model.CheckApiGroupsV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckApiGroupsV2Response), nil
	}
}

type CheckApisV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CheckApisV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckApisV2Invoker) Invoke() (*model.CheckApisV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckApisV2Response), nil
	}
}

type CreateApiGroupV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateApiGroupV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateApiGroupV2Invoker) Invoke() (*model.CreateApiGroupV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateApiGroupV2Response), nil
	}
}

type CreateApiV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateApiV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateApiV2Invoker) Invoke() (*model.CreateApiV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateApiV2Response), nil
	}
}

type CreateOrDeletePublishRecordForApiV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOrDeletePublishRecordForApiV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateOrDeletePublishRecordForApiV2Invoker) Invoke() (*model.CreateOrDeletePublishRecordForApiV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOrDeletePublishRecordForApiV2Response), nil
	}
}

type DebugApiV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DebugApiV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DebugApiV2Invoker) Invoke() (*model.DebugApiV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DebugApiV2Response), nil
	}
}

type DeleteApiByVersionIdV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteApiByVersionIdV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteApiByVersionIdV2Invoker) Invoke() (*model.DeleteApiByVersionIdV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteApiByVersionIdV2Response), nil
	}
}

type DeleteApiGroupV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteApiGroupV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteApiGroupV2Invoker) Invoke() (*model.DeleteApiGroupV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteApiGroupV2Response), nil
	}
}

type DeleteApiV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteApiV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteApiV2Invoker) Invoke() (*model.DeleteApiV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteApiV2Response), nil
	}
}

type DisassociateRequestThrottlingPolicyV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateRequestThrottlingPolicyV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisassociateRequestThrottlingPolicyV2Invoker) Invoke() (*model.DisassociateRequestThrottlingPolicyV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateRequestThrottlingPolicyV2Response), nil
	}
}

type ListApiGroupsV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiGroupsV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApiGroupsV2Invoker) Invoke() (*model.ListApiGroupsV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApiGroupsV2Response), nil
	}
}

type ListApiRuntimeDefinitionV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiRuntimeDefinitionV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApiRuntimeDefinitionV2Invoker) Invoke() (*model.ListApiRuntimeDefinitionV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApiRuntimeDefinitionV2Response), nil
	}
}

type ListApiVersionDetailV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiVersionDetailV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApiVersionDetailV2Invoker) Invoke() (*model.ListApiVersionDetailV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApiVersionDetailV2Response), nil
	}
}

type ListApiVersionsV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiVersionsV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApiVersionsV2Invoker) Invoke() (*model.ListApiVersionsV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApiVersionsV2Response), nil
	}
}

type ListApisBindedToRequestThrottlingPolicyV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListApisBindedToRequestThrottlingPolicyV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApisBindedToRequestThrottlingPolicyV2Invoker) Invoke() (*model.ListApisBindedToRequestThrottlingPolicyV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApisBindedToRequestThrottlingPolicyV2Response), nil
	}
}

type ListApisUnbindedToRequestThrottlingPolicyV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListApisUnbindedToRequestThrottlingPolicyV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApisUnbindedToRequestThrottlingPolicyV2Invoker) Invoke() (*model.ListApisUnbindedToRequestThrottlingPolicyV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApisUnbindedToRequestThrottlingPolicyV2Response), nil
	}
}

type ListApisV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListApisV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApisV2Invoker) Invoke() (*model.ListApisV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApisV2Response), nil
	}
}

type ListRequestThrottlingPoliciesBindedToApiV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListRequestThrottlingPoliciesBindedToApiV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRequestThrottlingPoliciesBindedToApiV2Invoker) Invoke() (*model.ListRequestThrottlingPoliciesBindedToApiV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRequestThrottlingPoliciesBindedToApiV2Response), nil
	}
}

type ShowDetailsOfApiGroupV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailsOfApiGroupV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDetailsOfApiGroupV2Invoker) Invoke() (*model.ShowDetailsOfApiGroupV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailsOfApiGroupV2Response), nil
	}
}

type ShowDetailsOfApiV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailsOfApiV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDetailsOfApiV2Invoker) Invoke() (*model.ShowDetailsOfApiV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailsOfApiV2Response), nil
	}
}

type UpdateApiGroupV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateApiGroupV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateApiGroupV2Invoker) Invoke() (*model.UpdateApiGroupV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateApiGroupV2Response), nil
	}
}

type UpdateApiV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateApiV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateApiV2Invoker) Invoke() (*model.UpdateApiV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateApiV2Response), nil
	}
}

type BatchDeleteApiAclBindingV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteApiAclBindingV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteApiAclBindingV2Invoker) Invoke() (*model.BatchDeleteApiAclBindingV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteApiAclBindingV2Response), nil
	}
}

type CreateApiAclBindingV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateApiAclBindingV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateApiAclBindingV2Invoker) Invoke() (*model.CreateApiAclBindingV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateApiAclBindingV2Response), nil
	}
}

type DeleteApiAclBindingV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteApiAclBindingV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteApiAclBindingV2Invoker) Invoke() (*model.DeleteApiAclBindingV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteApiAclBindingV2Response), nil
	}
}

type ListAclPolicyBindedToApiV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListAclPolicyBindedToApiV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAclPolicyBindedToApiV2Invoker) Invoke() (*model.ListAclPolicyBindedToApiV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAclPolicyBindedToApiV2Response), nil
	}
}

type ListApisBindedToAclPolicyV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListApisBindedToAclPolicyV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApisBindedToAclPolicyV2Invoker) Invoke() (*model.ListApisBindedToAclPolicyV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApisBindedToAclPolicyV2Response), nil
	}
}

type ListApisUnbindedToAclPolicyV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListApisUnbindedToAclPolicyV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApisUnbindedToAclPolicyV2Invoker) Invoke() (*model.ListApisUnbindedToAclPolicyV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApisUnbindedToAclPolicyV2Response), nil
	}
}

type CancelingAuthorizationV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CancelingAuthorizationV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelingAuthorizationV2Invoker) Invoke() (*model.CancelingAuthorizationV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelingAuthorizationV2Response), nil
	}
}

type CreateAuthorizingAppsV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAuthorizingAppsV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAuthorizingAppsV2Invoker) Invoke() (*model.CreateAuthorizingAppsV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAuthorizingAppsV2Response), nil
	}
}

type ListApisBindedToAppV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListApisBindedToAppV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApisBindedToAppV2Invoker) Invoke() (*model.ListApisBindedToAppV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApisBindedToAppV2Response), nil
	}
}

type ListApisUnbindedToAppV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListApisUnbindedToAppV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApisUnbindedToAppV2Invoker) Invoke() (*model.ListApisUnbindedToAppV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApisUnbindedToAppV2Response), nil
	}
}

type ListAppsBindedToApiV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppsBindedToApiV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppsBindedToApiV2Invoker) Invoke() (*model.ListAppsBindedToApiV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppsBindedToApiV2Response), nil
	}
}

type ListDuplicateApisForAppV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListDuplicateApisForAppV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDuplicateApisForAppV2Invoker) Invoke() (*model.ListDuplicateApisForAppV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDuplicateApisForAppV2Response), nil
	}
}

type AddUserToAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddUserToAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddUserToAppInvoker) Invoke() (*model.AddUserToAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddUserToAppResponse), nil
	}
}

type CheckAuthUsersOfAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckAuthUsersOfAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckAuthUsersOfAppInvoker) Invoke() (*model.CheckAuthUsersOfAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckAuthUsersOfAppResponse), nil
	}
}

type CheckCanAuthUsersOfAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckCanAuthUsersOfAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckCanAuthUsersOfAppInvoker) Invoke() (*model.CheckCanAuthUsersOfAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckCanAuthUsersOfAppResponse), nil
	}
}

type CheckRomaAppDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckRomaAppDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckRomaAppDetailsInvoker) Invoke() (*model.CheckRomaAppDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckRomaAppDetailsResponse), nil
	}
}

type CheckRomaAppSecretInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckRomaAppSecretInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckRomaAppSecretInvoker) Invoke() (*model.CheckRomaAppSecretResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckRomaAppSecretResponse), nil
	}
}

type CreateRomaAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRomaAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRomaAppInvoker) Invoke() (*model.CreateRomaAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRomaAppResponse), nil
	}
}

type DeleteRomaAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRomaAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRomaAppInvoker) Invoke() (*model.DeleteRomaAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRomaAppResponse), nil
	}
}

type ListRomaAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRomaAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRomaAppInvoker) Invoke() (*model.ListRomaAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRomaAppResponse), nil
	}
}

type ResetRomaAppSecretInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetRomaAppSecretInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetRomaAppSecretInvoker) Invoke() (*model.ResetRomaAppSecretResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetRomaAppSecretResponse), nil
	}
}

type UpdateRomaAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRomaAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRomaAppInvoker) Invoke() (*model.UpdateRomaAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRomaAppResponse), nil
	}
}

type ValidateRomaAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *ValidateRomaAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ValidateRomaAppInvoker) Invoke() (*model.ValidateRomaAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ValidateRomaAppResponse), nil
	}
}

type CheckAssetJobStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckAssetJobStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckAssetJobStatusInvoker) Invoke() (*model.CheckAssetJobStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckAssetJobStatusResponse), nil
	}
}

type DeleteAssetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAssetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAssetInvoker) Invoke() (*model.DeleteAssetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAssetResponse), nil
	}
}

type DownloadAssetArchiveInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadAssetArchiveInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadAssetArchiveInvoker) Invoke() (*model.DownloadAssetArchiveResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadAssetArchiveResponse), nil
	}
}

type ExportAssetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportAssetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportAssetInvoker) Invoke() (*model.ExportAssetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportAssetResponse), nil
	}
}

type ImportAssetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportAssetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportAssetInvoker) Invoke() (*model.ImportAssetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportAssetResponse), nil
	}
}

type CheckDictionaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckDictionaryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckDictionaryInvoker) Invoke() (*model.CheckDictionaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckDictionaryResponse), nil
	}
}

type CreateDictionaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDictionaryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDictionaryInvoker) Invoke() (*model.CreateDictionaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDictionaryResponse), nil
	}
}

type DeleteDictionaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDictionaryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDictionaryInvoker) Invoke() (*model.DeleteDictionaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDictionaryResponse), nil
	}
}

type ListDictionaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDictionaryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDictionaryInvoker) Invoke() (*model.ListDictionaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDictionaryResponse), nil
	}
}

type UpdateDictionaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDictionaryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDictionaryInvoker) Invoke() (*model.UpdateDictionaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDictionaryResponse), nil
	}
}

type ValidateDictionaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ValidateDictionaryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ValidateDictionaryInvoker) Invoke() (*model.ValidateDictionaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ValidateDictionaryResponse), nil
	}
}

type CheckRomaInstanceListV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CheckRomaInstanceListV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckRomaInstanceListV2Invoker) Invoke() (*model.CheckRomaInstanceListV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckRomaInstanceListV2Response), nil
	}
}

type ListMqsInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMqsInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMqsInstanceInvoker) Invoke() (*model.ListMqsInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMqsInstanceResponse), nil
	}
}

type ShowMqsInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMqsInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMqsInstanceInvoker) Invoke() (*model.ShowMqsInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMqsInstanceResponse), nil
	}
}

type ExportApiDefinitionsV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ExportApiDefinitionsV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportApiDefinitionsV2Invoker) Invoke() (*model.ExportApiDefinitionsV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportApiDefinitionsV2Response), nil
	}
}

type ExportLiveDataApiDefinitionsV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ExportLiveDataApiDefinitionsV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportLiveDataApiDefinitionsV2Invoker) Invoke() (*model.ExportLiveDataApiDefinitionsV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportLiveDataApiDefinitionsV2Response), nil
	}
}

type ImportApiDefinitionsV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ImportApiDefinitionsV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportApiDefinitionsV2Invoker) Invoke() (*model.ImportApiDefinitionsV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportApiDefinitionsV2Response), nil
	}
}

type ImportLiveDataApiDefinitionsV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ImportLiveDataApiDefinitionsV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportLiveDataApiDefinitionsV2Invoker) Invoke() (*model.ImportLiveDataApiDefinitionsV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportLiveDataApiDefinitionsV2Response), nil
	}
}

type BatchAssociateCertsV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAssociateCertsV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchAssociateCertsV2Invoker) Invoke() (*model.BatchAssociateCertsV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAssociateCertsV2Response), nil
	}
}

type BatchAssociateDomainsV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAssociateDomainsV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchAssociateDomainsV2Invoker) Invoke() (*model.BatchAssociateDomainsV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAssociateDomainsV2Response), nil
	}
}

type BatchDisassociateCertsV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDisassociateCertsV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDisassociateCertsV2Invoker) Invoke() (*model.BatchDisassociateCertsV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDisassociateCertsV2Response), nil
	}
}

type BatchDisassociateDomainsV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDisassociateDomainsV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDisassociateDomainsV2Invoker) Invoke() (*model.BatchDisassociateDomainsV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDisassociateDomainsV2Response), nil
	}
}

type CreateCertificateV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCertificateV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCertificateV2Invoker) Invoke() (*model.CreateCertificateV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCertificateV2Response), nil
	}
}

type DeleteCertificateV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCertificateV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCertificateV2Invoker) Invoke() (*model.DeleteCertificateV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCertificateV2Response), nil
	}
}

type ListAttachedDomainsV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListAttachedDomainsV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAttachedDomainsV2Invoker) Invoke() (*model.ListAttachedDomainsV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAttachedDomainsV2Response), nil
	}
}

type ListCertificatesV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListCertificatesV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCertificatesV2Invoker) Invoke() (*model.ListCertificatesV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCertificatesV2Response), nil
	}
}

type ShowDetailsOfCertificateV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailsOfCertificateV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDetailsOfCertificateV2Invoker) Invoke() (*model.ShowDetailsOfCertificateV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailsOfCertificateV2Response), nil
	}
}

type UpdateCertificateV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCertificateV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCertificateV2Invoker) Invoke() (*model.UpdateCertificateV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCertificateV2Response), nil
	}
}

type AddingBackendInstancesV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *AddingBackendInstancesV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddingBackendInstancesV2Invoker) Invoke() (*model.AddingBackendInstancesV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddingBackendInstancesV2Response), nil
	}
}

type BatchDisableMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDisableMembersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDisableMembersInvoker) Invoke() (*model.BatchDisableMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDisableMembersResponse), nil
	}
}

type BatchEnableMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchEnableMembersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchEnableMembersInvoker) Invoke() (*model.BatchEnableMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchEnableMembersResponse), nil
	}
}

type CreateMemberGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMemberGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMemberGroupInvoker) Invoke() (*model.CreateMemberGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMemberGroupResponse), nil
	}
}

type CreateProjectVpcChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateProjectVpcChannelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateProjectVpcChannelInvoker) Invoke() (*model.CreateProjectVpcChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateProjectVpcChannelResponse), nil
	}
}

type CreateProjectVpcChannelSyncsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateProjectVpcChannelSyncsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateProjectVpcChannelSyncsInvoker) Invoke() (*model.CreateProjectVpcChannelSyncsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateProjectVpcChannelSyncsResponse), nil
	}
}

type CreateVpcChannelV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVpcChannelV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateVpcChannelV2Invoker) Invoke() (*model.CreateVpcChannelV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVpcChannelV2Response), nil
	}
}

type DeleteBackendInstanceV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBackendInstanceV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteBackendInstanceV2Invoker) Invoke() (*model.DeleteBackendInstanceV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBackendInstanceV2Response), nil
	}
}

type DeleteMemberGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMemberGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMemberGroupInvoker) Invoke() (*model.DeleteMemberGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMemberGroupResponse), nil
	}
}

type DeleteVpcChannelV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVpcChannelV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteVpcChannelV2Invoker) Invoke() (*model.DeleteVpcChannelV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVpcChannelV2Response), nil
	}
}

type ListBackendInstancesV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListBackendInstancesV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBackendInstancesV2Invoker) Invoke() (*model.ListBackendInstancesV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBackendInstancesV2Response), nil
	}
}

type ListMemberGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMemberGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMemberGroupsInvoker) Invoke() (*model.ListMemberGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMemberGroupsResponse), nil
	}
}

type ListProjectVpcChannelsV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectVpcChannelsV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectVpcChannelsV2Invoker) Invoke() (*model.ListProjectVpcChannelsV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectVpcChannelsV2Response), nil
	}
}

type ListVpcChannelsV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ListVpcChannelsV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVpcChannelsV2Invoker) Invoke() (*model.ListVpcChannelsV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVpcChannelsV2Response), nil
	}
}

type ShowDetailsOfMemberGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailsOfMemberGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDetailsOfMemberGroupInvoker) Invoke() (*model.ShowDetailsOfMemberGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailsOfMemberGroupResponse), nil
	}
}

type ShowDetailsOfVpcChannelV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDetailsOfVpcChannelV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDetailsOfVpcChannelV2Invoker) Invoke() (*model.ShowDetailsOfVpcChannelV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDetailsOfVpcChannelV2Response), nil
	}
}

type UpdateBackendInstancesV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateBackendInstancesV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateBackendInstancesV2Invoker) Invoke() (*model.UpdateBackendInstancesV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateBackendInstancesV2Response), nil
	}
}

type UpdateHealthCheckInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHealthCheckInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHealthCheckInvoker) Invoke() (*model.UpdateHealthCheckResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHealthCheckResponse), nil
	}
}

type UpdateMemberGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMemberGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMemberGroupInvoker) Invoke() (*model.UpdateMemberGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMemberGroupResponse), nil
	}
}

type UpdateProjectVpcChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProjectVpcChannelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProjectVpcChannelInvoker) Invoke() (*model.UpdateProjectVpcChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProjectVpcChannelResponse), nil
	}
}

type UpdateVpcChannelV2Invoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVpcChannelV2Invoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateVpcChannelV2Invoker) Invoke() (*model.UpdateVpcChannelV2Response, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVpcChannelV2Response), nil
	}
}
