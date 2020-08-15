package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/apig/v2/model"
	"net/http"
)

func GenReqDefForAssociateCertificateV2(request *model.AssociateCertificateV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/api-groups/{group_id}/domains/{domain_id}/certificate").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("group_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithBodyJson(request.Body)

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForAssociateCertificateV2() (*model.AssociateCertificateV2Response, *def.HttpResponseDef) {
	resp := new(model.AssociateCertificateV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForAssociateDomainV2(request *model.AssociateDomainV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/api-groups/{group_id}/domains").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("group_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithBodyJson(request.Body)

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForAssociateDomainV2() (*model.AssociateDomainV2Response, *def.HttpResponseDef) {
	resp := new(model.AssociateDomainV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForAssociateSignatureKeyV2(request *model.AssociateSignatureKeyV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/sign-bindings").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithBodyJson(request.Body)

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForAssociateSignatureKeyV2() (*model.AssociateSignatureKeyV2Response, *def.HttpResponseDef) {
	resp := new(model.AssociateSignatureKeyV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForCreateEnvironmentV2(request *model.CreateEnvironmentV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/envs").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithBodyJson(request.Body)

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForCreateEnvironmentV2() (*model.CreateEnvironmentV2Response, *def.HttpResponseDef) {
	resp := new(model.CreateEnvironmentV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForCreateEnvironmentVariableV2(request *model.CreateEnvironmentVariableV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/env-variables").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithBodyJson(request.Body)

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForCreateEnvironmentVariableV2() (*model.CreateEnvironmentVariableV2Response, *def.HttpResponseDef) {
	resp := new(model.CreateEnvironmentVariableV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForCreateRequestThrottlingPolicyV2(request *model.CreateRequestThrottlingPolicyV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/throttles").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithBodyJson(request.Body)

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForCreateRequestThrottlingPolicyV2() (*model.CreateRequestThrottlingPolicyV2Response, *def.HttpResponseDef) {
	resp := new(model.CreateRequestThrottlingPolicyV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForCreateSignatureKeyV2(request *model.CreateSignatureKeyV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/signs").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithBodyJson(request.Body)

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForCreateSignatureKeyV2() (*model.CreateSignatureKeyV2Response, *def.HttpResponseDef) {
	resp := new(model.CreateSignatureKeyV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForCreateSpecialThrottlingConfigurationV2(request *model.CreateSpecialThrottlingConfigurationV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/throttles/{throttle_id}/throttle-specials").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("throttle_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithBodyJson(request.Body)

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForCreateSpecialThrottlingConfigurationV2() (*model.CreateSpecialThrottlingConfigurationV2Response, *def.HttpResponseDef) {
	resp := new(model.CreateSpecialThrottlingConfigurationV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForDeleteEnvironmentV2(request *model.DeleteEnvironmentV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/envs/{env_id}")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("env_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForDeleteEnvironmentV2() (*model.DeleteEnvironmentV2Response, *def.HttpResponseDef) {
	resp := new(model.DeleteEnvironmentV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForDeleteEnvironmentVariableV2(request *model.DeleteEnvironmentVariableV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/env-variables/{env_variable_id}")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("env_variable_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForDeleteEnvironmentVariableV2() (*model.DeleteEnvironmentVariableV2Response, *def.HttpResponseDef) {
	resp := new(model.DeleteEnvironmentVariableV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForDeleteRequestThrottlingPolicyV2(request *model.DeleteRequestThrottlingPolicyV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/throttles/{throttle_id}")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("throttle_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForDeleteRequestThrottlingPolicyV2() (*model.DeleteRequestThrottlingPolicyV2Response, *def.HttpResponseDef) {
	resp := new(model.DeleteRequestThrottlingPolicyV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForDeleteSignatureKeyV2(request *model.DeleteSignatureKeyV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/signs/{sign_id}")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("sign_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForDeleteSignatureKeyV2() (*model.DeleteSignatureKeyV2Response, *def.HttpResponseDef) {
	resp := new(model.DeleteSignatureKeyV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForDeleteSpecialThrottlingConfigurationV2(request *model.DeleteSpecialThrottlingConfigurationV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/throttles/{throttle_id}/throttle-specials/{strategy_id}")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("throttle_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("strategy_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForDeleteSpecialThrottlingConfigurationV2() (*model.DeleteSpecialThrottlingConfigurationV2Response, *def.HttpResponseDef) {
	resp := new(model.DeleteSpecialThrottlingConfigurationV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForDisassociateCertificateV2(request *model.DisassociateCertificateV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/api-groups/{group_id}/domains/{domain_id}/certificate/{certificate_id}")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("group_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("certificate_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForDisassociateCertificateV2() (*model.DisassociateCertificateV2Response, *def.HttpResponseDef) {
	resp := new(model.DisassociateCertificateV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForDisassociateDomainV2(request *model.DisassociateDomainV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/api-groups/{group_id}/domains/{domain_id}")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("group_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForDisassociateDomainV2() (*model.DisassociateDomainV2Response, *def.HttpResponseDef) {
	resp := new(model.DisassociateDomainV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForDisassociateSignatureKeyV2(request *model.DisassociateSignatureKeyV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/sign-bindings/{sign_bindings_id}")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("sign_bindings_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForDisassociateSignatureKeyV2() (*model.DisassociateSignatureKeyV2Response, *def.HttpResponseDef) {
	resp := new(model.DisassociateSignatureKeyV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForListApiGroupsQuantitiesV2(request *model.ListApiGroupsQuantitiesV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/resources/outline/groups")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForListApiGroupsQuantitiesV2() (*model.ListApiGroupsQuantitiesV2Response, *def.HttpResponseDef) {
	resp := new(model.ListApiGroupsQuantitiesV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForListApiQuantitiesV2(request *model.ListApiQuantitiesV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/resources/outline/apis")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForListApiQuantitiesV2() (*model.ListApiQuantitiesV2Response, *def.HttpResponseDef) {
	resp := new(model.ListApiQuantitiesV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForListApisBindedToSignatureKeyV2(request *model.ListApisBindedToSignatureKeyV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/sign-bindings/binded-apis")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("sign_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("env_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("api_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("api_name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("group_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("offset").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("limit").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForListApisBindedToSignatureKeyV2() (*model.ListApisBindedToSignatureKeyV2Response, *def.HttpResponseDef) {
	resp := new(model.ListApisBindedToSignatureKeyV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForListApisNotBoundWithSignatureKeyV2(request *model.ListApisNotBoundWithSignatureKeyV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/sign-bindings/unbinded-apis")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("sign_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("env_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("api_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("api_name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("group_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("offset").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("limit").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForListApisNotBoundWithSignatureKeyV2() (*model.ListApisNotBoundWithSignatureKeyV2Response, *def.HttpResponseDef) {
	resp := new(model.ListApisNotBoundWithSignatureKeyV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForListAppQuantitiesV2(request *model.ListAppQuantitiesV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/resources/outline/apps")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForListAppQuantitiesV2() (*model.ListAppQuantitiesV2Response, *def.HttpResponseDef) {
	resp := new(model.ListAppQuantitiesV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForListEnvironmentVariablesV2(request *model.ListEnvironmentVariablesV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/env-variables")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("group_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("env_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("variable_name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("offset").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("limit").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("precise_search").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForListEnvironmentVariablesV2() (*model.ListEnvironmentVariablesV2Response, *def.HttpResponseDef) {
	resp := new(model.ListEnvironmentVariablesV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForListEnvironmentsV2(request *model.ListEnvironmentsV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/envs")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("offset").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("limit").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForListEnvironmentsV2() (*model.ListEnvironmentsV2Response, *def.HttpResponseDef) {
	resp := new(model.ListEnvironmentsV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForListRequestThrottlingPolicyV2(request *model.ListRequestThrottlingPolicyV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/throttles")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("offset").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("limit").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("precise_search").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForListRequestThrottlingPolicyV2() (*model.ListRequestThrottlingPolicyV2Response, *def.HttpResponseDef) {
	resp := new(model.ListRequestThrottlingPolicyV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForListSignatureKeysBindedToApiV2(request *model.ListSignatureKeysBindedToApiV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/sign-bindings/binded-signs")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("api_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("sign_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("sign_name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("env_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("offset").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("limit").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForListSignatureKeysBindedToApiV2() (*model.ListSignatureKeysBindedToApiV2Response, *def.HttpResponseDef) {
	resp := new(model.ListSignatureKeysBindedToApiV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForListSignatureKeysV2(request *model.ListSignatureKeysV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/signs")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("offset").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("limit").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("precise_search").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForListSignatureKeysV2() (*model.ListSignatureKeysV2Response, *def.HttpResponseDef) {
	resp := new(model.ListSignatureKeysV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForListSpecialThrottlingConfigurationsV2(request *model.ListSpecialThrottlingConfigurationsV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/throttles/{throttle_id}/throttle-specials")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("throttle_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("object_type").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("app_name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("offset").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("limit").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForListSpecialThrottlingConfigurationsV2() (*model.ListSpecialThrottlingConfigurationsV2Response, *def.HttpResponseDef) {
	resp := new(model.ListSpecialThrottlingConfigurationsV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForShowDetailsOfDomainNameCertificateV2(request *model.ShowDetailsOfDomainNameCertificateV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/api-groups/{group_id}/domains/{domain_id}/certificate/{certificate_id}")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("group_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("certificate_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForShowDetailsOfDomainNameCertificateV2() (*model.ShowDetailsOfDomainNameCertificateV2Response, *def.HttpResponseDef) {
	resp := new(model.ShowDetailsOfDomainNameCertificateV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForShowDetailsOfEnvironmentVariableV2(request *model.ShowDetailsOfEnvironmentVariableV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/env-variables/{env_variable_id}")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("env_variable_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForShowDetailsOfEnvironmentVariableV2() (*model.ShowDetailsOfEnvironmentVariableV2Response, *def.HttpResponseDef) {
	resp := new(model.ShowDetailsOfEnvironmentVariableV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForShowDetailsOfRequestThrottlingPolicyV2(request *model.ShowDetailsOfRequestThrottlingPolicyV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/throttles/{throttle_id}")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("throttle_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForShowDetailsOfRequestThrottlingPolicyV2() (*model.ShowDetailsOfRequestThrottlingPolicyV2Response, *def.HttpResponseDef) {
	resp := new(model.ShowDetailsOfRequestThrottlingPolicyV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForUpdateEnvironmentV2(request *model.UpdateEnvironmentV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/envs/{env_id}").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("env_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithBodyJson(request.Body)

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForUpdateEnvironmentV2() (*model.UpdateEnvironmentV2Response, *def.HttpResponseDef) {
	resp := new(model.UpdateEnvironmentV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForUpdateRequestThrottlingPolicyV2(request *model.UpdateRequestThrottlingPolicyV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/throttles/{throttle_id}").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("throttle_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithBodyJson(request.Body)

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForUpdateRequestThrottlingPolicyV2() (*model.UpdateRequestThrottlingPolicyV2Response, *def.HttpResponseDef) {
	resp := new(model.UpdateRequestThrottlingPolicyV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForUpdateSignatureKeyV2(request *model.UpdateSignatureKeyV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/signs/{sign_id}").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("sign_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithBodyJson(request.Body)

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForUpdateSignatureKeyV2() (*model.UpdateSignatureKeyV2Response, *def.HttpResponseDef) {
	resp := new(model.UpdateSignatureKeyV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForUpdateSpecialThrottlingConfigurationV2(request *model.UpdateSpecialThrottlingConfigurationV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/throttles/{throttle_id}/throttle-specials/{strategy_id}").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("throttle_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("strategy_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithBodyJson(request.Body)

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForUpdateSpecialThrottlingConfigurationV2() (*model.UpdateSpecialThrottlingConfigurationV2Response, *def.HttpResponseDef) {
	resp := new(model.UpdateSpecialThrottlingConfigurationV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForAssociateRequestThrottlingPolicyV2(request *model.AssociateRequestThrottlingPolicyV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/throttle-bindings").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithBodyJson(request.Body)

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForAssociateRequestThrottlingPolicyV2() (*model.AssociateRequestThrottlingPolicyV2Response, *def.HttpResponseDef) {
	resp := new(model.AssociateRequestThrottlingPolicyV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForBatchDisassociateThrottlingPolicyV2(request *model.BatchDisassociateThrottlingPolicyV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/throttle-bindings").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("action").
		WithLocationType(def.Query))

	reqDefBuilder.WithBodyJson(request.Body)

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForBatchDisassociateThrottlingPolicyV2() (*model.BatchDisassociateThrottlingPolicyV2Response, *def.HttpResponseDef) {
	resp := new(model.BatchDisassociateThrottlingPolicyV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForCreateApiGroupV2(request *model.CreateApiGroupV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/api-groups").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithBodyJson(request.Body)

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForCreateApiGroupV2() (*model.CreateApiGroupV2Response, *def.HttpResponseDef) {
	resp := new(model.CreateApiGroupV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForCreateApiV2(request *model.CreateApiV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/apis").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithBodyJson(request.Body)

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForCreateApiV2() (*model.CreateApiV2Response, *def.HttpResponseDef) {
	resp := new(model.CreateApiV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForCreateOrDeletePublishRecordForApiV2(request *model.CreateOrDeletePublishRecordForApiV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/apis/action").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithBodyJson(request.Body)

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForCreateOrDeletePublishRecordForApiV2() (*model.CreateOrDeletePublishRecordForApiV2Response, *def.HttpResponseDef) {
	resp := new(model.CreateOrDeletePublishRecordForApiV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForDeleteApiGroupV2(request *model.DeleteApiGroupV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/api-groups/{group_id}")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("group_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForDeleteApiGroupV2() (*model.DeleteApiGroupV2Response, *def.HttpResponseDef) {
	resp := new(model.DeleteApiGroupV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForDeleteApiV2(request *model.DeleteApiV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/apis/{api_id}")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("api_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForDeleteApiV2() (*model.DeleteApiV2Response, *def.HttpResponseDef) {
	resp := new(model.DeleteApiV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForDisassociateRequestThrottlingPolicyV2(request *model.DisassociateRequestThrottlingPolicyV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/throttle-bindings/{throttle_binding_id}")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("throttle_binding_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForDisassociateRequestThrottlingPolicyV2() (*model.DisassociateRequestThrottlingPolicyV2Response, *def.HttpResponseDef) {
	resp := new(model.DisassociateRequestThrottlingPolicyV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForListApiGroupsV2(request *model.ListApiGroupsV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/api-groups")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("offset").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("limit").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("precise_search").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForListApiGroupsV2() (*model.ListApiGroupsV2Response, *def.HttpResponseDef) {
	resp := new(model.ListApiGroupsV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForListApisBindedToRequestThrottlingPolicyV2(request *model.ListApisBindedToRequestThrottlingPolicyV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/throttle-bindings/binded-apis")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("throttle_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("env_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("group_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("api_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("api_name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("offset").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("limit").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForListApisBindedToRequestThrottlingPolicyV2() (*model.ListApisBindedToRequestThrottlingPolicyV2Response, *def.HttpResponseDef) {
	resp := new(model.ListApisBindedToRequestThrottlingPolicyV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForListApisUnbindedToRequestThrottlingPolicyV2(request *model.ListApisUnbindedToRequestThrottlingPolicyV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/throttle-bindings/unbinded-apis")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("throttle_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("env_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("group_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("api_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("api_name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("offset").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("limit").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForListApisUnbindedToRequestThrottlingPolicyV2() (*model.ListApisUnbindedToRequestThrottlingPolicyV2Response, *def.HttpResponseDef) {
	resp := new(model.ListApisUnbindedToRequestThrottlingPolicyV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForListApisV2(request *model.ListApisV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/apis")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("group_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("req_protocol").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("req_method").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("req_uri").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("auth_type").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("env_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("type").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("offset").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("limit").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("precise_search").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForListApisV2() (*model.ListApisV2Response, *def.HttpResponseDef) {
	resp := new(model.ListApisV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForListRequestThrottlingPoliciesBindedToApiV2(request *model.ListRequestThrottlingPoliciesBindedToApiV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/throttle-bindings/binded-throttles")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("api_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("throttle_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("throttle_name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("env_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("offset").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("limit").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForListRequestThrottlingPoliciesBindedToApiV2() (*model.ListRequestThrottlingPoliciesBindedToApiV2Response, *def.HttpResponseDef) {
	resp := new(model.ListRequestThrottlingPoliciesBindedToApiV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForShowDetailsOfApiGroupV2(request *model.ShowDetailsOfApiGroupV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/api-groups/{group_id}")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("group_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForShowDetailsOfApiGroupV2() (*model.ShowDetailsOfApiGroupV2Response, *def.HttpResponseDef) {
	resp := new(model.ShowDetailsOfApiGroupV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForShowDetailsOfApiV2(request *model.ShowDetailsOfApiV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/apis/{api_id}")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("api_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForShowDetailsOfApiV2() (*model.ShowDetailsOfApiV2Response, *def.HttpResponseDef) {
	resp := new(model.ShowDetailsOfApiV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForUpdateApiGroupV2(request *model.UpdateApiGroupV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/api-groups/{group_id}").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("group_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithBodyJson(request.Body)

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForUpdateApiGroupV2() (*model.UpdateApiGroupV2Response, *def.HttpResponseDef) {
	resp := new(model.UpdateApiGroupV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForUpdateApiV2(request *model.UpdateApiV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/apis/{api_id}").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("api_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithBodyJson(request.Body)

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForUpdateApiV2() (*model.UpdateApiV2Response, *def.HttpResponseDef) {
	resp := new(model.UpdateApiV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForCancelingAuthorizationV2(request *model.CancelingAuthorizationV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/app-auths/{app_auth_id}")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("app_auth_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForCancelingAuthorizationV2() (*model.CancelingAuthorizationV2Response, *def.HttpResponseDef) {
	resp := new(model.CancelingAuthorizationV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForCheckAppV2(request *model.CheckAppV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/apps/validation/{app_id}")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("app_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForCheckAppV2() (*model.CheckAppV2Response, *def.HttpResponseDef) {
	resp := new(model.CheckAppV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForCreateAnAppV2(request *model.CreateAnAppV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/apps").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithBodyJson(request.Body)

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForCreateAnAppV2() (*model.CreateAnAppV2Response, *def.HttpResponseDef) {
	resp := new(model.CreateAnAppV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForCreateAuthorizingAppsV2(request *model.CreateAuthorizingAppsV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/app-auths").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithBodyJson(request.Body)

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForCreateAuthorizingAppsV2() (*model.CreateAuthorizingAppsV2Response, *def.HttpResponseDef) {
	resp := new(model.CreateAuthorizingAppsV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForDeleteAppV2(request *model.DeleteAppV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/apps/{app_id}")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("app_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForDeleteAppV2() (*model.DeleteAppV2Response, *def.HttpResponseDef) {
	resp := new(model.DeleteAppV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForListApisBindedToAppV2(request *model.ListApisBindedToAppV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/app-auths/binded-apis")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("app_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("api_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("api_name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("group_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("group_name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("env_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("offset").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("limit").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForListApisBindedToAppV2() (*model.ListApisBindedToAppV2Response, *def.HttpResponseDef) {
	resp := new(model.ListApisBindedToAppV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForListApisUnbindedToAppV2(request *model.ListApisUnbindedToAppV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/app-auths/unbinded-apis")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("app_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("env_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("group_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("api_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("api_name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("offset").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("limit").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForListApisUnbindedToAppV2() (*model.ListApisUnbindedToAppV2Response, *def.HttpResponseDef) {
	resp := new(model.ListApisUnbindedToAppV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForListAppsBindedToApiV2(request *model.ListAppsBindedToApiV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/app-auths/binded-apps")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("api_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("app_name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("app_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("env_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("offset").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("limit").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForListAppsBindedToApiV2() (*model.ListAppsBindedToApiV2Response, *def.HttpResponseDef) {
	resp := new(model.ListAppsBindedToApiV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForListAppsV2(request *model.ListAppsV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/apps")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("status").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("app_key").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("creator").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("offset").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("limit").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("precise_search").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForListAppsV2() (*model.ListAppsV2Response, *def.HttpResponseDef) {
	resp := new(model.ListAppsV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForResettingAppSecretV2(request *model.ResettingAppSecretV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/apps/secret/{app_id}").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("app_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithBodyJson(request.Body)

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForResettingAppSecretV2() (*model.ResettingAppSecretV2Response, *def.HttpResponseDef) {
	resp := new(model.ResettingAppSecretV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForShowDetailsOfAppV2(request *model.ShowDetailsOfAppV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/apps/{app_id}")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("app_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForShowDetailsOfAppV2() (*model.ShowDetailsOfAppV2Response, *def.HttpResponseDef) {
	resp := new(model.ShowDetailsOfAppV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForUpdateAppV2(request *model.UpdateAppV2Request) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v2/{project_id}/apigw/instances/{instance_id}/apps/{app_id}").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("instance_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("app_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithBodyJson(request.Body)

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForUpdateAppV2() (*model.UpdateAppV2Response, *def.HttpResponseDef) {
	resp := new(model.UpdateAppV2Response)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}
