package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/functiongraph/v2/model"
	"net/http"
)

func GenReqDefForAsyncInvokeFunction() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v2/{project_id}/fgs/functions/{function_urn}/invocations-async").
		WithResponse(new(model.AsyncInvokeFunctionResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FunctionUrn").
		WithJsonTag("function_urn").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateFunction() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v2/{project_id}/fgs/functions").
		WithResponse(new(model.CreateFunctionResponse)).
		WithContentType("application/json")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateFunctionVersion() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v2/{project_id}/fgs/functions/{function_urn}/versions").
		WithResponse(new(model.CreateFunctionVersionResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FunctionUrn").
		WithJsonTag("function_urn").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateVersionAlias() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v2/{project_id}/fgs/functions/{function_urn}/aliases").
		WithResponse(new(model.CreateVersionAliasResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FunctionUrn").
		WithJsonTag("function_urn").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteFunction() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v2/{project_id}/fgs/functions/{function_urn}").
		WithResponse(new(model.DeleteFunctionResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FunctionUrn").
		WithJsonTag("function_urn").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteVersionAlias() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v2/{project_id}/fgs/functions/{function_urn}/aliases/{name}").
		WithResponse(new(model.DeleteVersionAliasResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FunctionUrn").
		WithJsonTag("function_urn").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Name").
		WithJsonTag("name").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForInvokeFunction() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v2/{project_id}/fgs/functions/{function_urn}/invocations").
		WithResponse(new(model.InvokeFunctionResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FunctionUrn").
		WithJsonTag("function_urn").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XCffLogType").
		WithJsonTag("X-Cff-Log-Type").
		WithLocationType(def.Header))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XCFFRequestVersion").
		WithJsonTag("X-CFF-Request-Version").
		WithLocationType(def.Header))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListFunctionStatistics() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/fgs/functions/{func_urn}/statistics/{period}").
		WithResponse(new(model.ListFunctionStatisticsResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FuncUrn").
		WithJsonTag("func_urn").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Period").
		WithJsonTag("period").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListFunctionVersions() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/fgs/functions/{function_urn}/versions").
		WithResponse(new(model.ListFunctionVersionsResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FunctionUrn").
		WithJsonTag("function_urn").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Marker").
		WithJsonTag("marker").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Maxitems").
		WithJsonTag("maxitems").
		WithLocationType(def.Query))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListFunctions() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/fgs/functions").
		WithResponse(new(model.ListFunctionsResponse))

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Marker").
		WithJsonTag("marker").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Maxitems").
		WithJsonTag("maxitems").
		WithLocationType(def.Query))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListQuotas() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/fgs/quotas").
		WithResponse(new(model.ListQuotasResponse))

	// request

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListStatistics() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/fgs/functions/statistics").
		WithResponse(new(model.ListStatisticsResponse))

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Filter").
		WithJsonTag("filter").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Period").
		WithJsonTag("period").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("MonthCode").
		WithJsonTag("month_code").
		WithLocationType(def.Query))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListVersionAliases() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/fgs/functions/{function_urn}/aliases").
		WithResponse(new(model.ListVersionAliasesResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FunctionUrn").
		WithJsonTag("function_urn").
		WithLocationType(def.Path))

	// response
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowFunctionCode() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/fgs/functions/{function_urn}/code").
		WithResponse(new(model.ShowFunctionCodeResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FunctionUrn").
		WithJsonTag("function_urn").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowFunctionConfig() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/fgs/functions/{function_urn}/config").
		WithResponse(new(model.ShowFunctionConfigResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FunctionUrn").
		WithJsonTag("function_urn").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowVersionAlias() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/fgs/functions/{function_urn}/aliases/{name}").
		WithResponse(new(model.ShowVersionAliasResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FunctionUrn").
		WithJsonTag("function_urn").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Name").
		WithJsonTag("name").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateFunctionCode() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v2/{project_id}/fgs/functions/{function_urn}/code").
		WithResponse(new(model.UpdateFunctionCodeResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FunctionUrn").
		WithJsonTag("function_urn").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateFunctionConfig() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v2/{project_id}/fgs/functions/{function_urn}/config").
		WithResponse(new(model.UpdateFunctionConfigResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FunctionUrn").
		WithJsonTag("function_urn").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateVersionAlias() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v2/{project_id}/fgs/functions/{function_urn}/aliases/{name}").
		WithResponse(new(model.UpdateVersionAliasResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FunctionUrn").
		WithJsonTag("function_urn").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Name").
		WithJsonTag("name").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForBatchDeleteFunctionTriggers() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v2/{project_id}/fgs/triggers/{function_urn}").
		WithResponse(new(model.BatchDeleteFunctionTriggersResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FunctionUrn").
		WithJsonTag("function_urn").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateFunctionTrigger() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v2/{project_id}/fgs/triggers/{function_urn}").
		WithResponse(new(model.CreateFunctionTriggerResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FunctionUrn").
		WithJsonTag("function_urn").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteFunctionTrigger() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v2/{project_id}/fgs/triggers/{function_urn}/{trigger_type_code}/{triggerId}").
		WithResponse(new(model.DeleteFunctionTriggerResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FunctionUrn").
		WithJsonTag("function_urn").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("TriggerTypeCode").
		WithJsonTag("trigger_type_code").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("TriggerId").
		WithJsonTag("triggerId").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListFunctionTriggers() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/fgs/triggers/{function_urn}").
		WithResponse(new(model.ListFunctionTriggersResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FunctionUrn").
		WithJsonTag("function_urn").
		WithLocationType(def.Path))

	// response
	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowFunctionTrigger() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/fgs/triggers/{function_urn}/{trigger_type_code}/{triggerId}").
		WithResponse(new(model.ShowFunctionTriggerResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FunctionUrn").
		WithJsonTag("function_urn").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("TriggerTypeCode").
		WithJsonTag("trigger_type_code").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("TriggerId").
		WithJsonTag("triggerId").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateTrigger() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v2/{project_id}/fgs/triggers/{function_urn}/{trigger_type_code}/{triggerId}").
		WithResponse(new(model.UpdateTriggerResponse)).
		WithContentType("application/json")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("FunctionUrn").
		WithJsonTag("function_urn").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("TriggerTypeCode").
		WithJsonTag("trigger_type_code").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("TriggerId").
		WithJsonTag("triggerId").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}
