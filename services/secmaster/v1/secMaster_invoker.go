package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/secmaster/v1/model"
)

type BatchCreateDataobjectRelationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateDataobjectRelationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateDataobjectRelationsInvoker) Invoke() (*model.BatchCreateDataobjectRelationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateDataobjectRelationsResponse), nil
	}
}

type BatchCreateDatapanelObjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateDatapanelObjectsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateDatapanelObjectsInvoker) Invoke() (*model.BatchCreateDatapanelObjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateDatapanelObjectsResponse), nil
	}
}

type BatchSearchMetricHitsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchSearchMetricHitsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchSearchMetricHitsInvoker) Invoke() (*model.BatchSearchMetricHitsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchSearchMetricHitsResponse), nil
	}
}

type BatchTagResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchTagResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchTagResourcesInvoker) Invoke() (*model.BatchTagResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchTagResourcesResponse), nil
	}
}

type BatchUntagResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUntagResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUntagResourcesInvoker) Invoke() (*model.BatchUntagResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUntagResourcesResponse), nil
	}
}

type BatchUpdateCatalogueInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateCatalogueInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdateCatalogueInvoker) Invoke() (*model.BatchUpdateCatalogueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateCatalogueResponse), nil
	}
}

type ChangeAlertInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeAlertInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeAlertInvoker) Invoke() (*model.ChangeAlertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeAlertResponse), nil
	}
}

type ChangeAlertsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeAlertsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeAlertsInvoker) Invoke() (*model.ChangeAlertsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeAlertsResponse), nil
	}
}

type ChangeIncidentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeIncidentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeIncidentInvoker) Invoke() (*model.ChangeIncidentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeIncidentResponse), nil
	}
}

type ChangeIncidentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeIncidentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeIncidentsInvoker) Invoke() (*model.ChangeIncidentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeIncidentsResponse), nil
	}
}

type ChangePlaybookInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangePlaybookInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangePlaybookInstanceInvoker) Invoke() (*model.ChangePlaybookInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangePlaybookInstanceResponse), nil
	}
}

type ChangeResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeResourceInvoker) Invoke() (*model.ChangeResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeResourceResponse), nil
	}
}

type CopyMappingInvoker struct {
	*invoker.BaseInvoker
}

func (i *CopyMappingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CopyMappingInvoker) Invoke() (*model.CopyMappingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CopyMappingResponse), nil
	}
}

type CopyPlaybookVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CopyPlaybookVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CopyPlaybookVersionInvoker) Invoke() (*model.CopyPlaybookVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CopyPlaybookVersionResponse), nil
	}
}

type CountResourceInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountResourceInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountResourceInstanceInvoker) Invoke() (*model.CountResourceInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountResourceInstanceResponse), nil
	}
}

type CreateAlertInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAlertInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAlertInvoker) Invoke() (*model.CreateAlertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAlertResponse), nil
	}
}

type CreateAlertRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAlertRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAlertRuleInvoker) Invoke() (*model.CreateAlertRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAlertRuleResponse), nil
	}
}

type CreateAlertRuleSimulationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAlertRuleSimulationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAlertRuleSimulationInvoker) Invoke() (*model.CreateAlertRuleSimulationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAlertRuleSimulationResponse), nil
	}
}

type CreateAopWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAopWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAopWorkflowInvoker) Invoke() (*model.CreateAopWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAopWorkflowResponse), nil
	}
}

type CreateAopWorkflowVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAopWorkflowVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAopWorkflowVersionInvoker) Invoke() (*model.CreateAopWorkflowVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAopWorkflowVersionResponse), nil
	}
}

type CreateAopWorkflowVersionApprovelInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAopWorkflowVersionApprovelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAopWorkflowVersionApprovelInvoker) Invoke() (*model.CreateAopWorkflowVersionApprovelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAopWorkflowVersionApprovelResponse), nil
	}
}

type CreateBatchOrderAlertsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBatchOrderAlertsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateBatchOrderAlertsInvoker) Invoke() (*model.CreateBatchOrderAlertsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBatchOrderAlertsResponse), nil
	}
}

type CreateCatalogueInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCatalogueInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCatalogueInvoker) Invoke() (*model.CreateCatalogueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCatalogueResponse), nil
	}
}

type CreateClassifierInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClassifierInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateClassifierInvoker) Invoke() (*model.CreateClassifierResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClassifierResponse), nil
	}
}

type CreateCollectConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCollectConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCollectConfigInvoker) Invoke() (*model.CreateCollectConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCollectConfigResponse), nil
	}
}

type CreateCollectorChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCollectorChannelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCollectorChannelInvoker) Invoke() (*model.CreateCollectorChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCollectorChannelResponse), nil
	}
}

type CreateCollectorChannelGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCollectorChannelGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCollectorChannelGroupInvoker) Invoke() (*model.CreateCollectorChannelGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCollectorChannelGroupResponse), nil
	}
}

type CreateCollectorChannelOperationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCollectorChannelOperationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCollectorChannelOperationInvoker) Invoke() (*model.CreateCollectorChannelOperationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCollectorChannelOperationResponse), nil
	}
}

type CreateCollectorConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCollectorConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCollectorConnectionInvoker) Invoke() (*model.CreateCollectorConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCollectorConnectionResponse), nil
	}
}

type CreateCollectorFilesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCollectorFilesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCollectorFilesInvoker) Invoke() (*model.CreateCollectorFilesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCollectorFilesResponse), nil
	}
}

type CreateCollectorParserInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCollectorParserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCollectorParserInvoker) Invoke() (*model.CreateCollectorParserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCollectorParserResponse), nil
	}
}

type CreateComponentTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateComponentTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateComponentTemplateInvoker) Invoke() (*model.CreateComponentTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateComponentTemplateResponse), nil
	}
}

type CreateConfigurationApplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateConfigurationApplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateConfigurationApplicationInvoker) Invoke() (*model.CreateConfigurationApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConfigurationApplicationResponse), nil
	}
}

type CreateConfigurationDictionariesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateConfigurationDictionariesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateConfigurationDictionariesInvoker) Invoke() (*model.CreateConfigurationDictionariesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConfigurationDictionariesResponse), nil
	}
}

type CreateConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateConnectionInvoker) Invoke() (*model.CreateConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConnectionResponse), nil
	}
}

type CreateDataclassTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDataclassTypeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDataclassTypeInvoker) Invoke() (*model.CreateDataclassTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDataclassTypeResponse), nil
	}
}

type CreateDataobjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDataobjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDataobjectInvoker) Invoke() (*model.CreateDataobjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDataobjectResponse), nil
	}
}

type CreateDataobjectRelationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDataobjectRelationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDataobjectRelationInvoker) Invoke() (*model.CreateDataobjectRelationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDataobjectRelationResponse), nil
	}
}

type CreateDataspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDataspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDataspaceInvoker) Invoke() (*model.CreateDataspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDataspaceResponse), nil
	}
}

type CreateIncidentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateIncidentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateIncidentInvoker) Invoke() (*model.CreateIncidentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateIncidentResponse), nil
	}
}

type CreateIndicatorInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateIndicatorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateIndicatorInvoker) Invoke() (*model.CreateIndicatorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateIndicatorResponse), nil
	}
}

type CreateLayoutInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLayoutInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateLayoutInvoker) Invoke() (*model.CreateLayoutResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLayoutResponse), nil
	}
}

type CreateLayoutWizardInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLayoutWizardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateLayoutWizardInvoker) Invoke() (*model.CreateLayoutWizardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLayoutWizardResponse), nil
	}
}

type CreateMapperInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMapperInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMapperInvoker) Invoke() (*model.CreateMapperResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMapperResponse), nil
	}
}

type CreateMetricInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMetricInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMetricInvoker) Invoke() (*model.CreateMetricResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMetricResponse), nil
	}
}

type CreateModuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateModuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateModuleInvoker) Invoke() (*model.CreateModuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateModuleResponse), nil
	}
}

type CreateNoteInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNoteInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateNoteInvoker) Invoke() (*model.CreateNoteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNoteResponse), nil
	}
}

type CreatePipeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePipeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePipeInvoker) Invoke() (*model.CreatePipeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePipeResponse), nil
	}
}

type CreatePipeConsumptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePipeConsumptionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePipeConsumptionInvoker) Invoke() (*model.CreatePipeConsumptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePipeConsumptionResponse), nil
	}
}

type CreatePlaybookInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePlaybookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePlaybookInvoker) Invoke() (*model.CreatePlaybookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePlaybookResponse), nil
	}
}

type CreatePlaybookActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePlaybookActionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePlaybookActionInvoker) Invoke() (*model.CreatePlaybookActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePlaybookActionResponse), nil
	}
}

type CreatePlaybookApproveInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePlaybookApproveInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePlaybookApproveInvoker) Invoke() (*model.CreatePlaybookApproveResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePlaybookApproveResponse), nil
	}
}

type CreatePlaybookRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePlaybookRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePlaybookRuleInvoker) Invoke() (*model.CreatePlaybookRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePlaybookRuleResponse), nil
	}
}

type CreatePlaybookVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePlaybookVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePlaybookVersionInvoker) Invoke() (*model.CreatePlaybookVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePlaybookVersionResponse), nil
	}
}

type CreatePreProcessRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePreProcessRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePreProcessRulesInvoker) Invoke() (*model.CreatePreProcessRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePreProcessRulesResponse), nil
	}
}

type CreateReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateReportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateReportInvoker) Invoke() (*model.CreateReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateReportResponse), nil
	}
}

type CreateResourceConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResourceConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateResourceConfigInvoker) Invoke() (*model.CreateResourceConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResourceConfigResponse), nil
	}
}

type CreateRetryPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRetryPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRetryPolicyInvoker) Invoke() (*model.CreateRetryPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRetryPolicyResponse), nil
	}
}

type CreateSearchAnalysisInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSearchAnalysisInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSearchAnalysisInvoker) Invoke() (*model.CreateSearchAnalysisResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSearchAnalysisResponse), nil
	}
}

type CreateSearchConditionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSearchConditionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSearchConditionInvoker) Invoke() (*model.CreateSearchConditionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSearchConditionResponse), nil
	}
}

type CreateServiceAgencyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateServiceAgencyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateServiceAgencyInvoker) Invoke() (*model.CreateServiceAgencyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateServiceAgencyResponse), nil
	}
}

type CreateShipperInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateShipperInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateShipperInvoker) Invoke() (*model.CreateShipperResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateShipperResponse), nil
	}
}

type CreateShipperDelegateAuthInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateShipperDelegateAuthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateShipperDelegateAuthInvoker) Invoke() (*model.CreateShipperDelegateAuthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateShipperDelegateAuthResponse), nil
	}
}

type CreateSubscriptionOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSubscriptionOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSubscriptionOrderInvoker) Invoke() (*model.CreateSubscriptionOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSubscriptionOrderResponse), nil
	}
}

type CreateWorkflowInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWorkflowInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWorkflowInstanceInvoker) Invoke() (*model.CreateWorkflowInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWorkflowInstanceResponse), nil
	}
}

type CreateWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWorkspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWorkspaceInvoker) Invoke() (*model.CreateWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWorkspaceResponse), nil
	}
}

type DeleteAlertInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAlertInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAlertInvoker) Invoke() (*model.DeleteAlertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAlertResponse), nil
	}
}

type DeleteAlertRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAlertRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAlertRuleInvoker) Invoke() (*model.DeleteAlertRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAlertRuleResponse), nil
	}
}

type DeleteAlertsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAlertsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAlertsInvoker) Invoke() (*model.DeleteAlertsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAlertsResponse), nil
	}
}

type DeleteAopWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAopWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAopWorkflowInvoker) Invoke() (*model.DeleteAopWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAopWorkflowResponse), nil
	}
}

type DeleteAopWorkflowVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAopWorkflowVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAopWorkflowVersionInvoker) Invoke() (*model.DeleteAopWorkflowVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAopWorkflowVersionResponse), nil
	}
}

type DeleteCatalogueInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCatalogueInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCatalogueInvoker) Invoke() (*model.DeleteCatalogueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCatalogueResponse), nil
	}
}

type DeleteClassifierInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteClassifierInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteClassifierInvoker) Invoke() (*model.DeleteClassifierResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteClassifierResponse), nil
	}
}

type DeleteCollectorChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCollectorChannelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCollectorChannelInvoker) Invoke() (*model.DeleteCollectorChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCollectorChannelResponse), nil
	}
}

type DeleteCollectorChannelGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCollectorChannelGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCollectorChannelGroupInvoker) Invoke() (*model.DeleteCollectorChannelGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCollectorChannelGroupResponse), nil
	}
}

type DeleteCollectorConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCollectorConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCollectorConnectionInvoker) Invoke() (*model.DeleteCollectorConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCollectorConnectionResponse), nil
	}
}

type DeleteCollectorParserInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCollectorParserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCollectorParserInvoker) Invoke() (*model.DeleteCollectorParserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCollectorParserResponse), nil
	}
}

type DeleteComponentTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteComponentTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteComponentTemplateInvoker) Invoke() (*model.DeleteComponentTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteComponentTemplateResponse), nil
	}
}

type DeleteConfigurationDictionariesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteConfigurationDictionariesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteConfigurationDictionariesInvoker) Invoke() (*model.DeleteConfigurationDictionariesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteConfigurationDictionariesResponse), nil
	}
}

type DeleteConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteConnectionInvoker) Invoke() (*model.DeleteConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteConnectionResponse), nil
	}
}

type DeleteDataclassTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDataclassTypeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDataclassTypeInvoker) Invoke() (*model.DeleteDataclassTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDataclassTypeResponse), nil
	}
}

type DeleteDataobjectRelationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDataobjectRelationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDataobjectRelationInvoker) Invoke() (*model.DeleteDataobjectRelationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDataobjectRelationResponse), nil
	}
}

type DeleteDataobjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDataobjectsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDataobjectsInvoker) Invoke() (*model.DeleteDataobjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDataobjectsResponse), nil
	}
}

type DeleteDataspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDataspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDataspaceInvoker) Invoke() (*model.DeleteDataspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDataspaceResponse), nil
	}
}

type DeleteIncidentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteIncidentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteIncidentInvoker) Invoke() (*model.DeleteIncidentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteIncidentResponse), nil
	}
}

type DeleteIncidentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteIncidentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteIncidentsInvoker) Invoke() (*model.DeleteIncidentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteIncidentsResponse), nil
	}
}

type DeleteIndicatorInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteIndicatorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteIndicatorInvoker) Invoke() (*model.DeleteIndicatorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteIndicatorResponse), nil
	}
}

type DeleteLayoutWizardInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLayoutWizardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteLayoutWizardInvoker) Invoke() (*model.DeleteLayoutWizardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLayoutWizardResponse), nil
	}
}

type DeleteLayoutsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLayoutsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteLayoutsInvoker) Invoke() (*model.DeleteLayoutsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLayoutsResponse), nil
	}
}

type DeleteLtsCloudLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLtsCloudLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteLtsCloudLogInvoker) Invoke() (*model.DeleteLtsCloudLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLtsCloudLogResponse), nil
	}
}

type DeleteMappingInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMappingInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMappingInfoInvoker) Invoke() (*model.DeleteMappingInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMappingInfoResponse), nil
	}
}

type DeleteMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMetricsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMetricsInvoker) Invoke() (*model.DeleteMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMetricsResponse), nil
	}
}

type DeleteModuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteModuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteModuleInvoker) Invoke() (*model.DeleteModuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteModuleResponse), nil
	}
}

type DeleteNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteNodeInvoker) Invoke() (*model.DeleteNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNodeResponse), nil
	}
}

type DeleteNotesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNotesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteNotesInvoker) Invoke() (*model.DeleteNotesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNotesResponse), nil
	}
}

type DeletePipeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePipeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePipeInvoker) Invoke() (*model.DeletePipeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePipeResponse), nil
	}
}

type DeletePipeConsumptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePipeConsumptionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePipeConsumptionInvoker) Invoke() (*model.DeletePipeConsumptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePipeConsumptionResponse), nil
	}
}

type DeletePlaybookInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePlaybookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePlaybookInvoker) Invoke() (*model.DeletePlaybookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePlaybookResponse), nil
	}
}

type DeletePlaybookActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePlaybookActionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePlaybookActionInvoker) Invoke() (*model.DeletePlaybookActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePlaybookActionResponse), nil
	}
}

type DeletePlaybookRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePlaybookRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePlaybookRuleInvoker) Invoke() (*model.DeletePlaybookRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePlaybookRuleResponse), nil
	}
}

type DeletePlaybookVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePlaybookVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePlaybookVersionInvoker) Invoke() (*model.DeletePlaybookVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePlaybookVersionResponse), nil
	}
}

type DeletePoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePoliciesInvoker) Invoke() (*model.DeletePoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePoliciesResponse), nil
	}
}

type DeleteReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteReportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteReportInvoker) Invoke() (*model.DeleteReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteReportResponse), nil
	}
}

type DeleteResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteResourceInvoker) Invoke() (*model.DeleteResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteResourceResponse), nil
	}
}

type DeleteSearchConditionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSearchConditionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSearchConditionInvoker) Invoke() (*model.DeleteSearchConditionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSearchConditionResponse), nil
	}
}

type DeleteShipperInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteShipperInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteShipperInvoker) Invoke() (*model.DeleteShipperResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteShipperResponse), nil
	}
}

type DeleteSingleMapperInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSingleMapperInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSingleMapperInvoker) Invoke() (*model.DeleteSingleMapperResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSingleMapperResponse), nil
	}
}

type DeleteSubscriptionOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSubscriptionOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSubscriptionOrderInvoker) Invoke() (*model.DeleteSubscriptionOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSubscriptionOrderResponse), nil
	}
}

type DeleteTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTagsInvoker) Invoke() (*model.DeleteTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTagsResponse), nil
	}
}

type DeleteWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteWorkspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteWorkspaceInvoker) Invoke() (*model.DeleteWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteWorkspaceResponse), nil
	}
}

type DisableAlertRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableAlertRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisableAlertRuleInvoker) Invoke() (*model.DisableAlertRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableAlertRuleResponse), nil
	}
}

type DownloadAlertTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadAlertTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadAlertTemplateInvoker) Invoke() (*model.DownloadAlertTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadAlertTemplateResponse), nil
	}
}

type DownloadAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadAttachmentInvoker) Invoke() (*model.DownloadAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadAttachmentResponse), nil
	}
}

type DownloadIncidentTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadIncidentTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadIncidentTemplateInvoker) Invoke() (*model.DownloadIncidentTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadIncidentTemplateResponse), nil
	}
}

type DownloadIndicatorTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadIndicatorTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadIndicatorTemplateInvoker) Invoke() (*model.DownloadIndicatorTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadIndicatorTemplateResponse), nil
	}
}

type DownloadResourceTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadResourceTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadResourceTemplateInvoker) Invoke() (*model.DownloadResourceTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadResourceTemplateResponse), nil
	}
}

type DownloadVulnerabilityTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadVulnerabilityTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadVulnerabilityTemplateInvoker) Invoke() (*model.DownloadVulnerabilityTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadVulnerabilityTemplateResponse), nil
	}
}

type EnableAlertRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableAlertRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableAlertRuleInvoker) Invoke() (*model.EnableAlertRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableAlertRuleResponse), nil
	}
}

type EnableDataclassTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableDataclassTypeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableDataclassTypeInvoker) Invoke() (*model.EnableDataclassTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableDataclassTypeResponse), nil
	}
}

type ExecuteLayoutInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteLayoutInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteLayoutInvoker) Invoke() (*model.ExecuteLayoutResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteLayoutResponse), nil
	}
}

type ExecuteReportActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteReportActionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteReportActionInvoker) Invoke() (*model.ExecuteReportActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteReportActionResponse), nil
	}
}

type ExportAlertsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportAlertsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportAlertsInvoker) Invoke() (*model.ExportAlertsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportAlertsResponse), nil
	}
}

type ExportAopworkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportAopworkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportAopworkflowInvoker) Invoke() (*model.ExportAopworkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportAopworkflowResponse), nil
	}
}

type ExportCollectorParserInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportCollectorParserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportCollectorParserInvoker) Invoke() (*model.ExportCollectorParserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportCollectorParserResponse), nil
	}
}

type ExportDataobjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportDataobjectsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportDataobjectsInvoker) Invoke() (*model.ExportDataobjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportDataobjectsResponse), nil
	}
}

type ExportIncidentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportIncidentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportIncidentsInvoker) Invoke() (*model.ExportIncidentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportIncidentsResponse), nil
	}
}

type ExportIndicatorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportIndicatorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportIndicatorsInvoker) Invoke() (*model.ExportIndicatorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportIndicatorsResponse), nil
	}
}

type ExportPlaybookInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportPlaybookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportPlaybookInvoker) Invoke() (*model.ExportPlaybookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportPlaybookResponse), nil
	}
}

type ExportResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportResourcesInvoker) Invoke() (*model.ExportResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportResourcesResponse), nil
	}
}

type ExportVulnerabilitiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportVulnerabilitiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportVulnerabilitiesInvoker) Invoke() (*model.ExportVulnerabilitiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportVulnerabilitiesResponse), nil
	}
}

type HandleShipperAuthorizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *HandleShipperAuthorizationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *HandleShipperAuthorizationInvoker) Invoke() (*model.HandleShipperAuthorizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.HandleShipperAuthorizationResponse), nil
	}
}

type ImportAlertsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportAlertsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportAlertsInvoker) Invoke() (*model.ImportAlertsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportAlertsResponse), nil
	}
}

type ImportAopworkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportAopworkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportAopworkflowInvoker) Invoke() (*model.ImportAopworkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportAopworkflowResponse), nil
	}
}

type ImportCollectorParserInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportCollectorParserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportCollectorParserInvoker) Invoke() (*model.ImportCollectorParserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportCollectorParserResponse), nil
	}
}

type ImportDataobjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportDataobjectsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportDataobjectsInvoker) Invoke() (*model.ImportDataobjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportDataobjectsResponse), nil
	}
}

type ImportIncidentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportIncidentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportIncidentsInvoker) Invoke() (*model.ImportIncidentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportIncidentsResponse), nil
	}
}

type ImportIndicatorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportIndicatorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportIndicatorsInvoker) Invoke() (*model.ImportIndicatorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportIndicatorsResponse), nil
	}
}

type ImportPlaybookInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportPlaybookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportPlaybookInvoker) Invoke() (*model.ImportPlaybookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportPlaybookResponse), nil
	}
}

type ImportResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportResourceInvoker) Invoke() (*model.ImportResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportResourceResponse), nil
	}
}

type ImportVulnerabilitiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportVulnerabilitiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportVulnerabilitiesInvoker) Invoke() (*model.ImportVulnerabilitiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportVulnerabilitiesResponse), nil
	}
}

type ListAlertRuleMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlertRuleMetricsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAlertRuleMetricsInvoker) Invoke() (*model.ListAlertRuleMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlertRuleMetricsResponse), nil
	}
}

type ListAlertRuleTemplateMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlertRuleTemplateMetricsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAlertRuleTemplateMetricsInvoker) Invoke() (*model.ListAlertRuleTemplateMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlertRuleTemplateMetricsResponse), nil
	}
}

type ListAlertRuleTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlertRuleTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAlertRuleTemplatesInvoker) Invoke() (*model.ListAlertRuleTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlertRuleTemplatesResponse), nil
	}
}

type ListAlertRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlertRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAlertRulesInvoker) Invoke() (*model.ListAlertRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlertRulesResponse), nil
	}
}

type ListAlertsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlertsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAlertsInvoker) Invoke() (*model.ListAlertsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlertsResponse), nil
	}
}

type ListAllSecondCatalogueInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllSecondCatalogueInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAllSecondCatalogueInvoker) Invoke() (*model.ListAllSecondCatalogueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllSecondCatalogueResponse), nil
	}
}

type ListAopWorkflowInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAopWorkflowInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAopWorkflowInstanceInvoker) Invoke() (*model.ListAopWorkflowInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAopWorkflowInstanceResponse), nil
	}
}

type ListAopWorkflowVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAopWorkflowVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAopWorkflowVersionsInvoker) Invoke() (*model.ListAopWorkflowVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAopWorkflowVersionsResponse), nil
	}
}

type ListCatalogueInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCatalogueInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCatalogueInvoker) Invoke() (*model.ListCatalogueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCatalogueResponse), nil
	}
}

type ListCloudLogAliasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCloudLogAliasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCloudLogAliasInvoker) Invoke() (*model.ListCloudLogAliasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCloudLogAliasResponse), nil
	}
}

type ListCloudLogManagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCloudLogManagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCloudLogManagesInvoker) Invoke() (*model.ListCloudLogManagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCloudLogManagesResponse), nil
	}
}

type ListCloudLogPlatformInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCloudLogPlatformInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCloudLogPlatformInvoker) Invoke() (*model.ListCloudLogPlatformResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCloudLogPlatformResponse), nil
	}
}

type ListCloudLogResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCloudLogResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCloudLogResourcesInvoker) Invoke() (*model.ListCloudLogResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCloudLogResourcesResponse), nil
	}
}

type ListCollectConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCollectConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCollectConfigInvoker) Invoke() (*model.ListCollectConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCollectConfigResponse), nil
	}
}

type ListCollectorChannelGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCollectorChannelGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCollectorChannelGroupInvoker) Invoke() (*model.ListCollectorChannelGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCollectorChannelGroupResponse), nil
	}
}

type ListCollectorChannelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCollectorChannelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCollectorChannelsInvoker) Invoke() (*model.ListCollectorChannelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCollectorChannelsResponse), nil
	}
}

type ListCollectorConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCollectorConnectionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCollectorConnectionsInvoker) Invoke() (*model.ListCollectorConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCollectorConnectionsResponse), nil
	}
}

type ListCollectorInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCollectorInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCollectorInstancesInvoker) Invoke() (*model.ListCollectorInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCollectorInstancesResponse), nil
	}
}

type ListCollectorModuleRestrictionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCollectorModuleRestrictionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCollectorModuleRestrictionsInvoker) Invoke() (*model.ListCollectorModuleRestrictionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCollectorModuleRestrictionsResponse), nil
	}
}

type ListCollectorModuleTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCollectorModuleTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCollectorModuleTemplateInvoker) Invoke() (*model.ListCollectorModuleTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCollectorModuleTemplateResponse), nil
	}
}

type ListCollectorNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCollectorNodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCollectorNodeInvoker) Invoke() (*model.ListCollectorNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCollectorNodeResponse), nil
	}
}

type ListCollectorParserTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCollectorParserTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCollectorParserTemplateInvoker) Invoke() (*model.ListCollectorParserTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCollectorParserTemplateResponse), nil
	}
}

type ListCollectorParsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCollectorParsersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCollectorParsersInvoker) Invoke() (*model.ListCollectorParsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCollectorParsersResponse), nil
	}
}

type ListComponentActionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListComponentActionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListComponentActionsInvoker) Invoke() (*model.ListComponentActionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListComponentActionsResponse), nil
	}
}

type ListComponentConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListComponentConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListComponentConfigurationInvoker) Invoke() (*model.ListComponentConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListComponentConfigurationResponse), nil
	}
}

type ListComponentTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListComponentTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListComponentTemplateInvoker) Invoke() (*model.ListComponentTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListComponentTemplateResponse), nil
	}
}

type ListComponentTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListComponentTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListComponentTemplatesInvoker) Invoke() (*model.ListComponentTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListComponentTemplatesResponse), nil
	}
}

type ListComponentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListComponentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListComponentsInvoker) Invoke() (*model.ListComponentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListComponentsResponse), nil
	}
}

type ListConfigurationDictionariesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConfigurationDictionariesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListConfigurationDictionariesInvoker) Invoke() (*model.ListConfigurationDictionariesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConfigurationDictionariesResponse), nil
	}
}

type ListConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConnectionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListConnectionsInvoker) Invoke() (*model.ListConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConnectionsResponse), nil
	}
}

type ListDataclassInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDataclassInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDataclassInvoker) Invoke() (*model.ListDataclassResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDataclassResponse), nil
	}
}

type ListDataclassFieldsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDataclassFieldsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDataclassFieldsInvoker) Invoke() (*model.ListDataclassFieldsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDataclassFieldsResponse), nil
	}
}

type ListDataobjectRelationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDataobjectRelationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDataobjectRelationsInvoker) Invoke() (*model.ListDataobjectRelationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDataobjectRelationsResponse), nil
	}
}

type ListDataobjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDataobjectsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDataobjectsInvoker) Invoke() (*model.ListDataobjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDataobjectsResponse), nil
	}
}

type ListDatapanelObjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatapanelObjectsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDatapanelObjectsInvoker) Invoke() (*model.ListDatapanelObjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatapanelObjectsResponse), nil
	}
}

type ListDataspacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDataspacesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDataspacesInvoker) Invoke() (*model.ListDataspacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDataspacesResponse), nil
	}
}

type ListHistoryComponentConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHistoryComponentConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHistoryComponentConfigurationInvoker) Invoke() (*model.ListHistoryComponentConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHistoryComponentConfigurationResponse), nil
	}
}

type ListIncidentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIncidentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIncidentsInvoker) Invoke() (*model.ListIncidentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIncidentsResponse), nil
	}
}

type ListIndicatorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIndicatorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIndicatorsInvoker) Invoke() (*model.ListIndicatorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIndicatorsResponse), nil
	}
}

type ListInstallationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstallationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstallationInvoker) Invoke() (*model.ListInstallationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstallationResponse), nil
	}
}

type ListIsapComponentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIsapComponentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIsapComponentsInvoker) Invoke() (*model.ListIsapComponentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIsapComponentsResponse), nil
	}
}

type ListLayoutInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLayoutInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLayoutInvoker) Invoke() (*model.ListLayoutResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLayoutResponse), nil
	}
}

type ListLayoutWizardsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLayoutWizardsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLayoutWizardsInvoker) Invoke() (*model.ListLayoutWizardsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLayoutWizardsResponse), nil
	}
}

type ListMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMetricsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMetricsInvoker) Invoke() (*model.ListMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMetricsResponse), nil
	}
}

type ListModulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListModulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListModulesInvoker) Invoke() (*model.ListModulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListModulesResponse), nil
	}
}

type ListNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNodesInvoker) Invoke() (*model.ListNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNodesResponse), nil
	}
}

type ListNotesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNotesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNotesInvoker) Invoke() (*model.ListNotesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNotesResponse), nil
	}
}

type ListPipesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPipesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPipesInvoker) Invoke() (*model.ListPipesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPipesResponse), nil
	}
}

type ListPlaybookActionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPlaybookActionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPlaybookActionsInvoker) Invoke() (*model.ListPlaybookActionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPlaybookActionsResponse), nil
	}
}

type ListPlaybookApprovesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPlaybookApprovesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPlaybookApprovesInvoker) Invoke() (*model.ListPlaybookApprovesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPlaybookApprovesResponse), nil
	}
}

type ListPlaybookAuditLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPlaybookAuditLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPlaybookAuditLogsInvoker) Invoke() (*model.ListPlaybookAuditLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPlaybookAuditLogsResponse), nil
	}
}

type ListPlaybookInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPlaybookInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPlaybookInstancesInvoker) Invoke() (*model.ListPlaybookInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPlaybookInstancesResponse), nil
	}
}

type ListPlaybookVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPlaybookVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPlaybookVersionsInvoker) Invoke() (*model.ListPlaybookVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPlaybookVersionsResponse), nil
	}
}

type ListPlaybooksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPlaybooksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPlaybooksInvoker) Invoke() (*model.ListPlaybooksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPlaybooksResponse), nil
	}
}

type ListProjectTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectTagInvoker) Invoke() (*model.ListProjectTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectTagResponse), nil
	}
}

type ListRecipientsStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRecipientsStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRecipientsStatusInvoker) Invoke() (*model.ListRecipientsStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecipientsStatusResponse), nil
	}
}

type ListRegionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRegionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRegionsInvoker) Invoke() (*model.ListRegionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRegionsResponse), nil
	}
}

type ListReportsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListReportsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListReportsInvoker) Invoke() (*model.ListReportsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListReportsResponse), nil
	}
}

type ListResourceInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceInstanceInvoker) Invoke() (*model.ListResourceInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceInstanceResponse), nil
	}
}

type ListResourceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceTagInvoker) Invoke() (*model.ListResourceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceTagResponse), nil
	}
}

type ListResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourcesInvoker) Invoke() (*model.ListResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourcesResponse), nil
	}
}

type ListRunningNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRunningNodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRunningNodesInvoker) Invoke() (*model.ListRunningNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRunningNodesResponse), nil
	}
}

type ListSearchConditionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSearchConditionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSearchConditionsInvoker) Invoke() (*model.ListSearchConditionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSearchConditionsResponse), nil
	}
}

type ListSearchHistogramsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSearchHistogramsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSearchHistogramsInvoker) Invoke() (*model.ListSearchHistogramsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSearchHistogramsResponse), nil
	}
}

type ListSearchLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSearchLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSearchLogsInvoker) Invoke() (*model.ListSearchLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSearchLogsResponse), nil
	}
}

type ListServiceAgencyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServiceAgencyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServiceAgencyInvoker) Invoke() (*model.ListServiceAgencyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServiceAgencyResponse), nil
	}
}

type ListShipperAuthorizationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListShipperAuthorizationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListShipperAuthorizationsInvoker) Invoke() (*model.ListShipperAuthorizationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListShipperAuthorizationsResponse), nil
	}
}

type ListShippersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListShippersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListShippersInvoker) Invoke() (*model.ListShippersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListShippersResponse), nil
	}
}

type ListSubscriptionGlobalOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubscriptionGlobalOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSubscriptionGlobalOrderInvoker) Invoke() (*model.ListSubscriptionGlobalOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubscriptionGlobalOrderResponse), nil
	}
}

type ListSubscriptionOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubscriptionOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSubscriptionOrderInvoker) Invoke() (*model.ListSubscriptionOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubscriptionOrderResponse), nil
	}
}

type ListSubscriptionProductInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubscriptionProductInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSubscriptionProductInvoker) Invoke() (*model.ListSubscriptionProductResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubscriptionProductResponse), nil
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

type ListVpcEndpointServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVpcEndpointServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVpcEndpointServiceInvoker) Invoke() (*model.ListVpcEndpointServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVpcEndpointServiceResponse), nil
	}
}

type ListVulnerabilitiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVulnerabilitiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVulnerabilitiesInvoker) Invoke() (*model.ListVulnerabilitiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVulnerabilitiesResponse), nil
	}
}

type ListWorkflowsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkflowsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkflowsInvoker) Invoke() (*model.ListWorkflowsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkflowsResponse), nil
	}
}

type ListWorkspacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkspacesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkspacesInvoker) Invoke() (*model.ListWorkspacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkspacesResponse), nil
	}
}

type PauseShipperInvoker struct {
	*invoker.BaseInvoker
}

func (i *PauseShipperInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PauseShipperInvoker) Invoke() (*model.PauseShipperResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PauseShipperResponse), nil
	}
}

type ResumeShipperInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResumeShipperInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResumeShipperInvoker) Invoke() (*model.ResumeShipperResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResumeShipperResponse), nil
	}
}

type RetryShipperInvoker struct {
	*invoker.BaseInvoker
}

func (i *RetryShipperInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RetryShipperInvoker) Invoke() (*model.RetryShipperResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetryShipperResponse), nil
	}
}

type RetryShipperAuthorizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *RetryShipperAuthorizationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RetryShipperAuthorizationInvoker) Invoke() (*model.RetryShipperAuthorizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetryShipperAuthorizationResponse), nil
	}
}

type SearchPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchPolicyInvoker) Invoke() (*model.SearchPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchPolicyResponse), nil
	}
}

type SearchPolicyRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchPolicyRecordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchPolicyRecordInvoker) Invoke() (*model.SearchPolicyRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchPolicyRecordResponse), nil
	}
}

type SearchPolicyRecordByPolicyIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchPolicyRecordByPolicyIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchPolicyRecordByPolicyIdInvoker) Invoke() (*model.SearchPolicyRecordByPolicyIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchPolicyRecordByPolicyIdResponse), nil
	}
}

type ShowAlertInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAlertInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAlertInvoker) Invoke() (*model.ShowAlertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAlertResponse), nil
	}
}

type ShowAlertRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAlertRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAlertRuleInvoker) Invoke() (*model.ShowAlertRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAlertRuleResponse), nil
	}
}

type ShowAlertRuleTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAlertRuleTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAlertRuleTemplateInvoker) Invoke() (*model.ShowAlertRuleTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAlertRuleTemplateResponse), nil
	}
}

type ShowAopWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAopWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAopWorkflowInvoker) Invoke() (*model.ShowAopWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAopWorkflowResponse), nil
	}
}

type ShowAopWorkflowInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAopWorkflowInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAopWorkflowInstanceInvoker) Invoke() (*model.ShowAopWorkflowInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAopWorkflowInstanceResponse), nil
	}
}

type ShowAopWorkflowVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAopWorkflowVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAopWorkflowVersionInvoker) Invoke() (*model.ShowAopWorkflowVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAopWorkflowVersionResponse), nil
	}
}

type ShowAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAttachmentInvoker) Invoke() (*model.ShowAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAttachmentResponse), nil
	}
}

type ShowClassifierInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClassifierInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowClassifierInfoInvoker) Invoke() (*model.ShowClassifierInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClassifierInfoResponse), nil
	}
}

type ShowCloudLogTenantWhitelistInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCloudLogTenantWhitelistInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCloudLogTenantWhitelistInvoker) Invoke() (*model.ShowCloudLogTenantWhitelistResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCloudLogTenantWhitelistResponse), nil
	}
}

type ShowCollectorChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCollectorChannelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCollectorChannelInvoker) Invoke() (*model.ShowCollectorChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCollectorChannelResponse), nil
	}
}

type ShowCollectorConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCollectorConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCollectorConnectionInvoker) Invoke() (*model.ShowCollectorConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCollectorConnectionResponse), nil
	}
}

type ShowCollectorParserInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCollectorParserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCollectorParserInvoker) Invoke() (*model.ShowCollectorParserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCollectorParserResponse), nil
	}
}

type ShowComponentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowComponentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowComponentInvoker) Invoke() (*model.ShowComponentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowComponentResponse), nil
	}
}

type ShowComponentActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowComponentActionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowComponentActionInvoker) Invoke() (*model.ShowComponentActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowComponentActionResponse), nil
	}
}

type ShowComponentTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowComponentTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowComponentTemplateInvoker) Invoke() (*model.ShowComponentTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowComponentTemplateResponse), nil
	}
}

type ShowConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowConnectionInvoker) Invoke() (*model.ShowConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConnectionResponse), nil
	}
}

type ShowDataClassInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDataClassInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDataClassInfoInvoker) Invoke() (*model.ShowDataClassInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDataClassInfoResponse), nil
	}
}

type ShowDataobjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDataobjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDataobjectInvoker) Invoke() (*model.ShowDataobjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDataobjectResponse), nil
	}
}

type ShowDatapanelObjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDatapanelObjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDatapanelObjectInvoker) Invoke() (*model.ShowDatapanelObjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDatapanelObjectResponse), nil
	}
}

type ShowDataspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDataspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDataspaceInvoker) Invoke() (*model.ShowDataspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDataspaceResponse), nil
	}
}

type ShowIncidentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIncidentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowIncidentInvoker) Invoke() (*model.ShowIncidentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIncidentResponse), nil
	}
}

type ShowIndicatorDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIndicatorDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowIndicatorDetailInvoker) Invoke() (*model.ShowIndicatorDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIndicatorDetailResponse), nil
	}
}

type ShowIsapComponentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIsapComponentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowIsapComponentInvoker) Invoke() (*model.ShowIsapComponentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIsapComponentResponse), nil
	}
}

type ShowLayoutInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLayoutInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLayoutInvoker) Invoke() (*model.ShowLayoutResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLayoutResponse), nil
	}
}

type ShowLayoutWizardInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLayoutWizardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLayoutWizardInvoker) Invoke() (*model.ShowLayoutWizardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLayoutWizardResponse), nil
	}
}

type ShowLayoutWizardByFieldInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLayoutWizardByFieldInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLayoutWizardByFieldInvoker) Invoke() (*model.ShowLayoutWizardByFieldResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLayoutWizardByFieldResponse), nil
	}
}

type ShowMapperDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMapperDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMapperDetailInvoker) Invoke() (*model.ShowMapperDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMapperDetailResponse), nil
	}
}

type ShowMapperListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMapperListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMapperListInvoker) Invoke() (*model.ShowMapperListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMapperListResponse), nil
	}
}

type ShowMappingFunctionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMappingFunctionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMappingFunctionInvoker) Invoke() (*model.ShowMappingFunctionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMappingFunctionResponse), nil
	}
}

type ShowMappingInfoListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMappingInfoListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMappingInfoListInvoker) Invoke() (*model.ShowMappingInfoListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMappingInfoListResponse), nil
	}
}

type ShowMetricMetaDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMetricMetaDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMetricMetaDataInvoker) Invoke() (*model.ShowMetricMetaDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMetricMetaDataResponse), nil
	}
}

type ShowModuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowModuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowModuleInvoker) Invoke() (*model.ShowModuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowModuleResponse), nil
	}
}

type ShowMoniterMetricStatsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMoniterMetricStatsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMoniterMetricStatsInvoker) Invoke() (*model.ShowMoniterMetricStatsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMoniterMetricStatsResponse), nil
	}
}

type ShowPipeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPipeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPipeInvoker) Invoke() (*model.ShowPipeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPipeResponse), nil
	}
}

type ShowPipeConsumptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPipeConsumptionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPipeConsumptionInvoker) Invoke() (*model.ShowPipeConsumptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPipeConsumptionResponse), nil
	}
}

type ShowPipeIndexInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPipeIndexInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPipeIndexInvoker) Invoke() (*model.ShowPipeIndexResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPipeIndexResponse), nil
	}
}

type ShowPlatformManagedInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPlatformManagedInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPlatformManagedInvoker) Invoke() (*model.ShowPlatformManagedResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPlatformManagedResponse), nil
	}
}

type ShowPlaybookInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPlaybookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPlaybookInvoker) Invoke() (*model.ShowPlaybookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPlaybookResponse), nil
	}
}

type ShowPlaybookInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPlaybookInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPlaybookInstanceInvoker) Invoke() (*model.ShowPlaybookInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPlaybookInstanceResponse), nil
	}
}

type ShowPlaybookMonitorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPlaybookMonitorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPlaybookMonitorsInvoker) Invoke() (*model.ShowPlaybookMonitorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPlaybookMonitorsResponse), nil
	}
}

type ShowPlaybookRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPlaybookRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPlaybookRuleInvoker) Invoke() (*model.ShowPlaybookRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPlaybookRuleResponse), nil
	}
}

type ShowPlaybookStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPlaybookStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPlaybookStatisticsInvoker) Invoke() (*model.ShowPlaybookStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPlaybookStatisticsResponse), nil
	}
}

type ShowPlaybookTopologyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPlaybookTopologyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPlaybookTopologyInvoker) Invoke() (*model.ShowPlaybookTopologyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPlaybookTopologyResponse), nil
	}
}

type ShowPlaybookVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPlaybookVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPlaybookVersionInvoker) Invoke() (*model.ShowPlaybookVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPlaybookVersionResponse), nil
	}
}

type ShowPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPolicyInvoker) Invoke() (*model.ShowPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPolicyResponse), nil
	}
}

type ShowPreProcessRulesListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPreProcessRulesListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPreProcessRulesListInvoker) Invoke() (*model.ShowPreProcessRulesListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPreProcessRulesListResponse), nil
	}
}

type ShowReportInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowReportInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowReportInfoInvoker) Invoke() (*model.ShowReportInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowReportInfoResponse), nil
	}
}

type ShowResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResourceInvoker) Invoke() (*model.ShowResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceResponse), nil
	}
}

type ShowSearchConditionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSearchConditionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSearchConditionInvoker) Invoke() (*model.ShowSearchConditionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSearchConditionResponse), nil
	}
}

type ShowShipperInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowShipperInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowShipperInvoker) Invoke() (*model.ShowShipperResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowShipperResponse), nil
	}
}

type ShowShipperDelegateAuthInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowShipperDelegateAuthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowShipperDelegateAuthInvoker) Invoke() (*model.ShowShipperDelegateAuthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowShipperDelegateAuthResponse), nil
	}
}

type ShowShipperParamInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowShipperParamInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowShipperParamInvoker) Invoke() (*model.ShowShipperParamResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowShipperParamResponse), nil
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

type ShowVulnerabilityInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVulnerabilityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowVulnerabilityInvoker) Invoke() (*model.ShowVulnerabilityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVulnerabilityResponse), nil
	}
}

type ShowWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWorkspaceInvoker) Invoke() (*model.ShowWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkspaceResponse), nil
	}
}

type UpdateAlertRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAlertRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAlertRuleInvoker) Invoke() (*model.UpdateAlertRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAlertRuleResponse), nil
	}
}

type UpdateAopWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAopWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAopWorkflowInvoker) Invoke() (*model.UpdateAopWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAopWorkflowResponse), nil
	}
}

type UpdateAopWorkflowVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAopWorkflowVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAopWorkflowVersionInvoker) Invoke() (*model.UpdateAopWorkflowVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAopWorkflowVersionResponse), nil
	}
}

type UpdateCatalogueInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCatalogueInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCatalogueInvoker) Invoke() (*model.UpdateCatalogueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCatalogueResponse), nil
	}
}

type UpdateClassifierInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateClassifierInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateClassifierInvoker) Invoke() (*model.UpdateClassifierResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateClassifierResponse), nil
	}
}

type UpdateCollectorChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCollectorChannelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCollectorChannelInvoker) Invoke() (*model.UpdateCollectorChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCollectorChannelResponse), nil
	}
}

type UpdateCollectorChannelGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCollectorChannelGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCollectorChannelGroupInvoker) Invoke() (*model.UpdateCollectorChannelGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCollectorChannelGroupResponse), nil
	}
}

type UpdateCollectorConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCollectorConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCollectorConnectionInvoker) Invoke() (*model.UpdateCollectorConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCollectorConnectionResponse), nil
	}
}

type UpdateComponentConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateComponentConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateComponentConfigurationInvoker) Invoke() (*model.UpdateComponentConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateComponentConfigurationResponse), nil
	}
}

type UpdateComponentTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateComponentTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateComponentTemplateInvoker) Invoke() (*model.UpdateComponentTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateComponentTemplateResponse), nil
	}
}

type UpdateConfigurationDictionariesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateConfigurationDictionariesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateConfigurationDictionariesInvoker) Invoke() (*model.UpdateConfigurationDictionariesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateConfigurationDictionariesResponse), nil
	}
}

type UpdateConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateConnectionInvoker) Invoke() (*model.UpdateConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateConnectionResponse), nil
	}
}

type UpdateDataobjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDataobjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDataobjectInvoker) Invoke() (*model.UpdateDataobjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDataobjectResponse), nil
	}
}

type UpdateDataspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDataspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDataspaceInvoker) Invoke() (*model.UpdateDataspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDataspaceResponse), nil
	}
}

type UpdateIndicatorInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateIndicatorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateIndicatorInvoker) Invoke() (*model.UpdateIndicatorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateIndicatorResponse), nil
	}
}

type UpdateLayoutInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLayoutInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateLayoutInvoker) Invoke() (*model.UpdateLayoutResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLayoutResponse), nil
	}
}

type UpdateLayoutWizardsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLayoutWizardsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateLayoutWizardsInvoker) Invoke() (*model.UpdateLayoutWizardsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLayoutWizardsResponse), nil
	}
}

type UpdateMapperInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMapperInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMapperInvoker) Invoke() (*model.UpdateMapperResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMapperResponse), nil
	}
}

type UpdateMappingInfoStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMappingInfoStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMappingInfoStatusInvoker) Invoke() (*model.UpdateMappingInfoStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMappingInfoStatusResponse), nil
	}
}

type UpdateMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMetricsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMetricsInvoker) Invoke() (*model.UpdateMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMetricsResponse), nil
	}
}

type UpdateModuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateModuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateModuleInvoker) Invoke() (*model.UpdateModuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateModuleResponse), nil
	}
}

type UpdateNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateNodeInvoker) Invoke() (*model.UpdateNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNodeResponse), nil
	}
}

type UpdatePipeInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePipeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePipeInvoker) Invoke() (*model.UpdatePipeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePipeResponse), nil
	}
}

type UpdatePipeIndexInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePipeIndexInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePipeIndexInvoker) Invoke() (*model.UpdatePipeIndexResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePipeIndexResponse), nil
	}
}

type UpdatePlaybookInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePlaybookInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePlaybookInvoker) Invoke() (*model.UpdatePlaybookResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePlaybookResponse), nil
	}
}

type UpdatePlaybookActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePlaybookActionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePlaybookActionInvoker) Invoke() (*model.UpdatePlaybookActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePlaybookActionResponse), nil
	}
}

type UpdatePlaybookRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePlaybookRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePlaybookRuleInvoker) Invoke() (*model.UpdatePlaybookRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePlaybookRuleResponse), nil
	}
}

type UpdatePlaybookVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePlaybookVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePlaybookVersionInvoker) Invoke() (*model.UpdatePlaybookVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePlaybookVersionResponse), nil
	}
}

type UpdateRecipientsEmailStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRecipientsEmailStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRecipientsEmailStatusInvoker) Invoke() (*model.UpdateRecipientsEmailStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRecipientsEmailStatusResponse), nil
	}
}

type UpdateReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateReportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateReportInvoker) Invoke() (*model.UpdateReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateReportResponse), nil
	}
}

type UpdateSearchConditionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSearchConditionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSearchConditionInvoker) Invoke() (*model.UpdateSearchConditionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSearchConditionResponse), nil
	}
}

type UpdateSubscriptionOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSubscriptionOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSubscriptionOrderInvoker) Invoke() (*model.UpdateSubscriptionOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSubscriptionOrderResponse), nil
	}
}

type UpdateTagValueInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTagValueInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTagValueInvoker) Invoke() (*model.UpdateTagValueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTagValueResponse), nil
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

type UpdateVpcEndpointServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVpcEndpointServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateVpcEndpointServiceInvoker) Invoke() (*model.UpdateVpcEndpointServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVpcEndpointServiceResponse), nil
	}
}

type UpdateWorkflowInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWorkflowInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateWorkflowInstanceInvoker) Invoke() (*model.UpdateWorkflowInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWorkflowInstanceResponse), nil
	}
}

type UpdateWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWorkspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateWorkspaceInvoker) Invoke() (*model.UpdateWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWorkspaceResponse), nil
	}
}

type UploadAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UploadAttachmentInvoker) Invoke() (*model.UploadAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadAttachmentResponse), nil
	}
}

type ValidateAopWorkflowVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ValidateAopWorkflowVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ValidateAopWorkflowVersionInvoker) Invoke() (*model.ValidateAopWorkflowVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ValidateAopWorkflowVersionResponse), nil
	}
}
