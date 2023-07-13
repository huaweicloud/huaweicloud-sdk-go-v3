package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dwr/v3/model"
	"net/http"
)

func GenReqDefForAcceptServiceContract() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3/{project_id}/workflow-agreements").
		WithResponse(new(model.AcceptServiceContractResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Type").
		WithJsonTag("type").
		WithLocationType(def.Query))

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XRequestId").
		WithJsonTag("x-request-id").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("ContentLength").
		WithJsonTag("Content-Length").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForAsyncInvokeApiStartWorkflow() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3/{project_id}/workflows/{graph_name}/execute").
		WithResponse(new(model.AsyncInvokeApiStartWorkflowResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GraphName").
		WithJsonTag("graph_name").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XRequestId").
		WithJsonTag("x-request-id").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Connection").
		WithJsonTag("Connection").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("ContentLength").
		WithJsonTag("Content-Length").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Date").
		WithJsonTag("Date").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCheckWorkflowAuthentication() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/{project_id}/workflow-authorization").
		WithResponse(new(model.CheckWorkflowAuthenticationResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XRequestId").
		WithJsonTag("x-request-id").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Connection").
		WithJsonTag("Connection").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("ContentLength").
		WithJsonTag("Content-Length").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Date").
		WithJsonTag("Date").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateMyActionTemplate() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3/{project_id}/myactiontemplates/{template_name}").
		WithResponse(new(model.CreateMyActionTemplateResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("TemplateName").
		WithJsonTag("template_name").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XRequestId").
		WithJsonTag("X-request-id").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Connection").
		WithJsonTag("Connection").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("ContentLength").
		WithJsonTag("Content-Length").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Date").
		WithJsonTag("Date").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateWorkflowAuthentication() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3/{project_id}/workflow-authorization").
		WithResponse(new(model.CreateWorkflowAuthenticationResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XRequestId").
		WithJsonTag("x-request-id").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("ContentLength").
		WithJsonTag("Content-Length").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteMyActionTemplate() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v3/{project_id}/myactiontemplates/{template_name}").
		WithResponse(new(model.DeleteMyActionTemplateResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("TemplateName").
		WithJsonTag("template_name").
		WithLocationType(def.Path))

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XRequestId").
		WithJsonTag("x-request-id").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("ContentLength").
		WithJsonTag("Content-Length").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListMyActionTemplate() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/{project_id}/myactiontemplates").
		WithResponse(new(model.ListMyActionTemplateResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Prefix").
		WithJsonTag("prefix").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Status").
		WithJsonTag("status").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Category").
		WithJsonTag("category").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Offset").
		WithJsonTag("offset").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Limit").
		WithJsonTag("limit").
		WithLocationType(def.Query))

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XRequestId").
		WithJsonTag("x-request-id").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Connection").
		WithJsonTag("Connection").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("ContentLength").
		WithJsonTag("Content-Length").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Date").
		WithJsonTag("Date").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListSystemTemplates() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/{project_id}/actiontemplates").
		WithResponse(new(model.ListSystemTemplatesResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Prefix").
		WithJsonTag("prefix").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Category").
		WithJsonTag("category").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Offset").
		WithJsonTag("offset").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Limit").
		WithJsonTag("limit").
		WithLocationType(def.Query))

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XRequestId").
		WithJsonTag("x-request-id").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Connection").
		WithJsonTag("Connection").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("ContentLength").
		WithJsonTag("Content-Length").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Date").
		WithJsonTag("Date").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListWorkflowInstance() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/{project_id}/workflowexecutions").
		WithResponse(new(model.ListWorkflowInstanceResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Limit").
		WithJsonTag("limit").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GraphName").
		WithJsonTag("graph_name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("StartTime").
		WithJsonTag("start_time").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("EndTime").
		WithJsonTag("end_time").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Status").
		WithJsonTag("status").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Offset").
		WithJsonTag("offset").
		WithLocationType(def.Query))

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XRequestId").
		WithJsonTag("X-Request-Id").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("ContentLength").
		WithJsonTag("Content-Length").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Date").
		WithJsonTag("Date").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("ContentType").
		WithJsonTag("Content-Type").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForRestoreWorkflowExecution() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3/{project_id}/workflows/{graph_name}/workflowexecution/{execution_name}/retry").
		WithResponse(new(model.RestoreWorkflowExecutionResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ExecutionName").
		WithJsonTag("execution_name").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GraphName").
		WithJsonTag("graph_name").
		WithLocationType(def.Path))

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XRequestId").
		WithJsonTag("x-request-id").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Connection").
		WithJsonTag("Connection").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("ContentLength").
		WithJsonTag("Content-Length").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Date").
		WithJsonTag("Date").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowPublicActionList() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/{project_id}/publicactiontemplates").
		WithResponse(new(model.ShowPublicActionListResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Prefix").
		WithJsonTag("prefix").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Category").
		WithJsonTag("category").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Offset").
		WithJsonTag("offset").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Limit").
		WithJsonTag("limit").
		WithLocationType(def.Query))

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XRequestId").
		WithJsonTag("X-request-id").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Connection").
		WithJsonTag("Connection").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("ContentLength").
		WithJsonTag("Content-Length").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Date").
		WithJsonTag("Date").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowPublicTemplateInfo() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/{project_id}/publicactiontemplate/{template_name}").
		WithResponse(new(model.ShowPublicTemplateInfoResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("TemplateName").
		WithJsonTag("template_name").
		WithLocationType(def.Path))

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XRequestId").
		WithJsonTag("X-request-id").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Connection").
		WithJsonTag("Connection").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("ContentLength").
		WithJsonTag("Content-Length").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Date").
		WithJsonTag("Date").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowServiceContract() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/{project_id}/workflow-agreements").
		WithResponse(new(model.ShowServiceContractResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Type").
		WithJsonTag("type").
		WithLocationType(def.Query))

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XRequestId").
		WithJsonTag("x-request-id").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Connection").
		WithJsonTag("Connection").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("ContentLength").
		WithJsonTag("Content-Length").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Date").
		WithJsonTag("Date").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowSystemTemplateDetail() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/{project_id}/actiontemplate/{template_name}").
		WithResponse(new(model.ShowSystemTemplateDetailResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("TemplateName").
		WithJsonTag("template_name").
		WithLocationType(def.Path))

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XRequestId").
		WithJsonTag("x-request-id").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Connection").
		WithJsonTag("Connection").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("ContentLength").
		WithJsonTag("Content-Length").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Date").
		WithJsonTag("Date").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowThirdTemplateInfo() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/{project_id}/myactiontemplate/{template_name}").
		WithResponse(new(model.ShowThirdTemplateInfoResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("TemplateName").
		WithJsonTag("template_name").
		WithLocationType(def.Path))

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XRequestId").
		WithJsonTag("x-request-id").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Connection").
		WithJsonTag("Connection").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("ContentLength").
		WithJsonTag("Content-Length").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Date").
		WithJsonTag("Date").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowWorkflowInstance() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/{project_id}/workflows/{graph_name}/workflowexecution/{execution_name}").
		WithResponse(new(model.ShowWorkflowInstanceResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ExecutionName").
		WithJsonTag("execution_name").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GraphName").
		WithJsonTag("graph_name").
		WithLocationType(def.Path))

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XRequestId").
		WithJsonTag("x-request-id").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Connection").
		WithJsonTag("Connection").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("ContentLength").
		WithJsonTag("Content-Length").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Date").
		WithJsonTag("Date").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateMyActionTemplate() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v3/{project_id}/myactiontemplates/{template_name}").
		WithResponse(new(model.UpdateMyActionTemplateResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("TemplateName").
		WithJsonTag("template_name").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XRequestId").
		WithJsonTag("x-request-id").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Connection").
		WithJsonTag("Connection").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("ContentLength").
		WithJsonTag("Content-Length").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Date").
		WithJsonTag("Date").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateMyActionTemplateToDeprecated() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3/{project_id}/myactiontemplates/{template_name}/forbid").
		WithResponse(new(model.UpdateMyActionTemplateToDeprecatedResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("TemplateName").
		WithJsonTag("template_name").
		WithLocationType(def.Path))

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XRequestId").
		WithJsonTag("x-request-id").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("ContentLength").
		WithJsonTag("Content-Length").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateWorkflow() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3/{project_id}/workflows/{graph_name}").
		WithResponse(new(model.CreateWorkflowResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GraphName").
		WithJsonTag("graph_name").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XRequestId").
		WithJsonTag("x-request-id").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Connection").
		WithJsonTag("Connection").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("ContentLength").
		WithJsonTag("Content-Length").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Date").
		WithJsonTag("Date").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteWorkflow() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v3/{project_id}/workflows/{graph_name}").
		WithResponse(new(model.DeleteWorkflowResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GraphName").
		WithJsonTag("graph_name").
		WithLocationType(def.Path))

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XRequestId").
		WithJsonTag("x-request-id").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("ContentLength").
		WithJsonTag("Content-Length").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListWorkflows() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/{project_id}/workflows").
		WithResponse(new(model.ListWorkflowsResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Prefix").
		WithJsonTag("prefix").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Offset").
		WithJsonTag("offset").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Limit").
		WithJsonTag("limit").
		WithLocationType(def.Query))

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XRequestId").
		WithJsonTag("x-request-id").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Connection").
		WithJsonTag("Connection").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("ContentLength").
		WithJsonTag("Content-Length").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Date").
		WithJsonTag("Date").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowWorkflowInfo() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/{project_id}/workflows/{graph_name}").
		WithResponse(new(model.ShowWorkflowInfoResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GraphName").
		WithJsonTag("graph_name").
		WithLocationType(def.Path))

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XRequestId").
		WithJsonTag("x-request-id").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Connection").
		WithJsonTag("Connection").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("ContentLength").
		WithJsonTag("Content-Length").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Date").
		WithJsonTag("Date").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateWorkflow() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v3/{project_id}/workflows/{graph_name}").
		WithResponse(new(model.UpdateWorkflowResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GraphName").
		WithJsonTag("graph_name").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XRequestId").
		WithJsonTag("x-request-id").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Connection").
		WithJsonTag("Connection").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("ContentLength").
		WithJsonTag("Content-Length").
		WithKindName("string").
		WithLocationType(def.Header))
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Date").
		WithJsonTag("Date").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}
