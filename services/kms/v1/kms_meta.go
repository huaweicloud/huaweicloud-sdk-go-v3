package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kms/v1/model"
	"net/http"
)

func GenReqDefForBatchCreateKmsTags(request *model.BatchCreateKmsTagsRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/{key_id}/tags/action").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("key_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForBatchCreateKmsTags() (*model.BatchCreateKmsTagsResponse, *def.HttpResponseDef) {
	resp := new(model.BatchCreateKmsTagsResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForCancelGrant(request *model.CancelGrantRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/revoke-grant").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForCancelGrant() (*model.CancelGrantResponse, *def.HttpResponseDef) {
	resp := new(model.CancelGrantResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForCancelKeyDeletion(request *model.CancelKeyDeletionRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/cancel-key-deletion").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForCancelKeyDeletion() (*model.CancelKeyDeletionResponse, *def.HttpResponseDef) {
	resp := new(model.CancelKeyDeletionResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForCancelSelfGrant(request *model.CancelSelfGrantRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/retire-grant").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForCancelSelfGrant() (*model.CancelSelfGrantResponse, *def.HttpResponseDef) {
	resp := new(model.CancelSelfGrantResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForCreateDatakey(request *model.CreateDatakeyRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/create-datakey").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForCreateDatakey() (*model.CreateDatakeyResponse, *def.HttpResponseDef) {
	resp := new(model.CreateDatakeyResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForCreateDatakeyWithoutPlaintext(request *model.CreateDatakeyWithoutPlaintextRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/create-datakey-without-plaintext").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForCreateDatakeyWithoutPlaintext() (*model.CreateDatakeyWithoutPlaintextResponse, *def.HttpResponseDef) {
	resp := new(model.CreateDatakeyWithoutPlaintextResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForCreateGrant(request *model.CreateGrantRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/create-grant").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForCreateGrant() (*model.CreateGrantResponse, *def.HttpResponseDef) {
	resp := new(model.CreateGrantResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForCreateKey(request *model.CreateKeyRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/create-key").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForCreateKey() (*model.CreateKeyResponse, *def.HttpResponseDef) {
	resp := new(model.CreateKeyResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForCreateKmsTag(request *model.CreateKmsTagRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/{key_id}/tags").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("key_id").
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

func GenRespForCreateKmsTag() (*model.CreateKmsTagResponse, *def.HttpResponseDef) {
	resp := new(model.CreateKmsTagResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForCreateParametersForImport(request *model.CreateParametersForImportRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/get-parameters-for-import").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForCreateParametersForImport() (*model.CreateParametersForImportResponse, *def.HttpResponseDef) {
	resp := new(model.CreateParametersForImportResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForCreateRandom(request *model.CreateRandomRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/gen-random").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForCreateRandom() (*model.CreateRandomResponse, *def.HttpResponseDef) {
	resp := new(model.CreateRandomResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForDecryptData(request *model.DecryptDataRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/decrypt-data").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForDecryptData() (*model.DecryptDataResponse, *def.HttpResponseDef) {
	resp := new(model.DecryptDataResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForDecryptDatakey(request *model.DecryptDatakeyRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/decrypt-datakey").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForDecryptDatakey() (*model.DecryptDatakeyResponse, *def.HttpResponseDef) {
	resp := new(model.DecryptDatakeyResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForDeleteImportedKeyMaterial(request *model.DeleteImportedKeyMaterialRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/delete-imported-key-material").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForDeleteImportedKeyMaterial() (*model.DeleteImportedKeyMaterialResponse, *def.HttpResponseDef) {
	resp := new(model.DeleteImportedKeyMaterialResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForDeleteKey(request *model.DeleteKeyRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/schedule-key-deletion").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForDeleteKey() (*model.DeleteKeyResponse, *def.HttpResponseDef) {
	resp := new(model.DeleteKeyResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForDeleteTag(request *model.DeleteTagRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/{version_id}/{project_id}/kms/{key_id}/tags/{key}")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("key_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("key").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForDeleteTag() (*model.DeleteTagResponse, *def.HttpResponseDef) {
	resp := new(model.DeleteTagResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForDisableKey(request *model.DisableKeyRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/disable-key").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForDisableKey() (*model.DisableKeyResponse, *def.HttpResponseDef) {
	resp := new(model.DisableKeyResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForDisableKeyRotation(request *model.DisableKeyRotationRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/disable-key-rotation").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForDisableKeyRotation() (*model.DisableKeyRotationResponse, *def.HttpResponseDef) {
	resp := new(model.DisableKeyRotationResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForEnableKey(request *model.EnableKeyRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/enable-key").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForEnableKey() (*model.EnableKeyResponse, *def.HttpResponseDef) {
	resp := new(model.EnableKeyResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForEnableKeyRotation(request *model.EnableKeyRotationRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/enable-key-rotation").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForEnableKeyRotation() (*model.EnableKeyRotationResponse, *def.HttpResponseDef) {
	resp := new(model.EnableKeyRotationResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForEncryptData(request *model.EncryptDataRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/encrypt-data").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForEncryptData() (*model.EncryptDataResponse, *def.HttpResponseDef) {
	resp := new(model.EncryptDataResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForEncryptDatakey(request *model.EncryptDatakeyRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/encrypt-datakey").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForEncryptDatakey() (*model.EncryptDatakeyResponse, *def.HttpResponseDef) {
	resp := new(model.EncryptDatakeyResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForImportKeyMaterial(request *model.ImportKeyMaterialRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/import-key-material").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForImportKeyMaterial() (*model.ImportKeyMaterialResponse, *def.HttpResponseDef) {
	resp := new(model.ImportKeyMaterialResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForListGrants(request *model.ListGrantsRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/list-grants").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForListGrants() (*model.ListGrantsResponse, *def.HttpResponseDef) {
	resp := new(model.ListGrantsResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForListKeyDetail(request *model.ListKeyDetailRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/describe-key").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForListKeyDetail() (*model.ListKeyDetailResponse, *def.HttpResponseDef) {
	resp := new(model.ListKeyDetailResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForListKeys(request *model.ListKeysRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/list-keys").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForListKeys() (*model.ListKeysResponse, *def.HttpResponseDef) {
	resp := new(model.ListKeysResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForListKmsByTags(request *model.ListKmsByTagsRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/{resource_instances}/action").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("resource_instances").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForListKmsByTags() (*model.ListKmsByTagsResponse, *def.HttpResponseDef) {
	resp := new(model.ListKmsByTagsResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForListKmsTags(request *model.ListKmsTagsRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/{version_id}/{project_id}/kms/tags")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForListKmsTags() (*model.ListKmsTagsResponse, *def.HttpResponseDef) {
	resp := new(model.ListKmsTagsResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForListRetirableGrants(request *model.ListRetirableGrantsRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/list-retirable-grants").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForListRetirableGrants() (*model.ListRetirableGrantsResponse, *def.HttpResponseDef) {
	resp := new(model.ListRetirableGrantsResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForShowKeyRotationStatus(request *model.ShowKeyRotationStatusRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/get-key-rotation-status").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForShowKeyRotationStatus() (*model.ShowKeyRotationStatusResponse, *def.HttpResponseDef) {
	resp := new(model.ShowKeyRotationStatusResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForShowKmsTags(request *model.ShowKmsTagsRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/{version_id}/{project_id}/kms/{key_id}/tags")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("key_id").
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

func GenRespForShowKmsTags() (*model.ShowKmsTagsResponse, *def.HttpResponseDef) {
	resp := new(model.ShowKmsTagsResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForShowUserInstances(request *model.ShowUserInstancesRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/{version_id}/{project_id}/kms/user-instances")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForShowUserInstances() (*model.ShowUserInstancesResponse, *def.HttpResponseDef) {
	resp := new(model.ShowUserInstancesResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForShowUserQuotas(request *model.ShowUserQuotasRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/{version_id}/{project_id}/kms/user-quotas")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForShowUserQuotas() (*model.ShowUserQuotasResponse, *def.HttpResponseDef) {
	resp := new(model.ShowUserQuotasResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForUpdateKeyAlias(request *model.UpdateKeyAliasRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/update-key-alias").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForUpdateKeyAlias() (*model.UpdateKeyAliasResponse, *def.HttpResponseDef) {
	resp := new(model.UpdateKeyAliasResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForUpdateKeyDescription(request *model.UpdateKeyDescriptionRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/update-key-description").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForUpdateKeyDescription() (*model.UpdateKeyDescriptionResponse, *def.HttpResponseDef) {
	resp := new(model.UpdateKeyDescriptionResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForUpdateKeyRotationInterval(request *model.UpdateKeyRotationIntervalRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/update-key-rotation-interval").
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForUpdateKeyRotationInterval() (*model.UpdateKeyRotationIntervalResponse, *def.HttpResponseDef) {
	resp := new(model.UpdateKeyRotationIntervalResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForShowVersion(request *model.ShowVersionRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/{version_id}")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("version_id").
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

func GenRespForShowVersion() (*model.ShowVersionResponse, *def.HttpResponseDef) {
	resp := new(model.ShowVersionResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}

func GenReqDefForShowVersions(request *model.ShowVersionsRequest) *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("domain_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenRespForShowVersions() (*model.ShowVersionsResponse, *def.HttpResponseDef) {
	resp := new(model.ShowVersionsResponse)
	respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
	responseDef := respDefBuilder.Build()
	return resp, responseDef
}
