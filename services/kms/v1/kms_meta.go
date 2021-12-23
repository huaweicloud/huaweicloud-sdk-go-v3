package v1

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/def"

	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/services/kms/v1/model"
	"net/http"
)

func GenReqDefForBatchCreateKmsTags() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/{key_id}/tags/action").
		WithResponse(new(model.BatchCreateKmsTagsResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("KeyId").
		WithJsonTag("key_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCancelGrant() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/revoke-grant").
		WithResponse(new(model.CancelGrantResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCancelKeyDeletion() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/cancel-key-deletion").
		WithResponse(new(model.CancelKeyDeletionResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCancelSelfGrant() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/retire-grant").
		WithResponse(new(model.CancelSelfGrantResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateDatakey() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/create-datakey").
		WithResponse(new(model.CreateDatakeyResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateDatakeyWithoutPlaintext() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/create-datakey-without-plaintext").
		WithResponse(new(model.CreateDatakeyWithoutPlaintextResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateGrant() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/create-grant").
		WithResponse(new(model.CreateGrantResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateKey() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/create-key").
		WithResponse(new(model.CreateKeyResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateKmsTag() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/{key_id}/tags").
		WithResponse(new(model.CreateKmsTagResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("KeyId").
		WithJsonTag("key_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateParametersForImport() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/get-parameters-for-import").
		WithResponse(new(model.CreateParametersForImportResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateRandom() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/gen-random").
		WithResponse(new(model.CreateRandomResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateSecret() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/{project_id}/secrets").
		WithResponse(new(model.CreateSecretResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateSecretVersion() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/{project_id}/secrets/{secret_id}/versions").
		WithResponse(new(model.CreateSecretVersionResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("SecretId").
		WithJsonTag("secret_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDecryptData() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/decrypt-data").
		WithResponse(new(model.DecryptDataResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDecryptDatakey() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/decrypt-datakey").
		WithResponse(new(model.DecryptDatakeyResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteImportedKeyMaterial() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/delete-imported-key-material").
		WithResponse(new(model.DeleteImportedKeyMaterialResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteKey() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/schedule-key-deletion").
		WithResponse(new(model.DeleteKeyResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteSecret() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v1/{project_id}/secrets/{secret_id}").
		WithResponse(new(model.DeleteSecretResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("SecretId").
		WithJsonTag("secret_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteSecretForSchedule() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/{project_id}/secrets/{secret_id}/scheduled-deleted-tasks/create").
		WithResponse(new(model.DeleteSecretForScheduleResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("SecretId").
		WithJsonTag("secret_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteSecretStage() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v1/{project_id}/secrets/{secret_id}/stages/{stage_name}").
		WithResponse(new(model.DeleteSecretStageResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("SecretId").
		WithJsonTag("secret_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("StageName").
		WithJsonTag("stage_name").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteTag() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/{version_id}/{project_id}/kms/{key_id}/tags/{key}").
		WithResponse(new(model.DeleteTagResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("KeyId").
		WithJsonTag("key_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Key").
		WithJsonTag("key").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDisableKey() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/disable-key").
		WithResponse(new(model.DisableKeyResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDisableKeyRotation() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/disable-key-rotation").
		WithResponse(new(model.DisableKeyRotationResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForEnableKey() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/enable-key").
		WithResponse(new(model.EnableKeyResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForEnableKeyRotation() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/enable-key-rotation").
		WithResponse(new(model.EnableKeyRotationResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForEncryptData() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/encrypt-data").
		WithResponse(new(model.EncryptDataResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForEncryptDatakey() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/encrypt-datakey").
		WithResponse(new(model.EncryptDatakeyResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForImportKeyMaterial() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/import-key-material").
		WithResponse(new(model.ImportKeyMaterialResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListGrants() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/list-grants").
		WithResponse(new(model.ListGrantsResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListKeyDetail() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/describe-key").
		WithResponse(new(model.ListKeyDetailResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListKeys() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/list-keys").
		WithResponse(new(model.ListKeysResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListKmsByTags() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/{resource_instances}/action").
		WithResponse(new(model.ListKmsByTagsResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ResourceInstances").
		WithJsonTag("resource_instances").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListKmsTags() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/{version_id}/{project_id}/kms/tags").
		WithResponse(new(model.ListKmsTagsResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListRetirableGrants() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/list-retirable-grants").
		WithResponse(new(model.ListRetirableGrantsResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListSecretStage() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v1/{project_id}/secrets/{secret_id}/stages/{stage_name}").
		WithResponse(new(model.ListSecretStageResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("SecretId").
		WithJsonTag("secret_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("StageName").
		WithJsonTag("stage_name").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListSecretVersions() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v1/{project_id}/secrets/{secret_id}/versions").
		WithResponse(new(model.ListSecretVersionsResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("SecretId").
		WithJsonTag("secret_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListSecrets() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v1/{project_id}/secrets").
		WithResponse(new(model.ListSecretsResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Limit").
		WithJsonTag("limit").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Marker").
		WithJsonTag("marker").
		WithLocationType(def.Query))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForRestoreSecret() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/{project_id}/secrets/{secret_id}/scheduled-deleted-tasks/cancel").
		WithResponse(new(model.RestoreSecretResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("SecretId").
		WithJsonTag("secret_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowKeyRotationStatus() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/get-key-rotation-status").
		WithResponse(new(model.ShowKeyRotationStatusResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowKmsTags() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/{version_id}/{project_id}/kms/{key_id}/tags").
		WithResponse(new(model.ShowKmsTagsResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("KeyId").
		WithJsonTag("key_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowPublicKey() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/get-publickey").
		WithResponse(new(model.ShowPublicKeyResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowSecret() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v1/{project_id}/secrets/{secret_id}").
		WithResponse(new(model.ShowSecretResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("SecretId").
		WithJsonTag("secret_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowSecretVersion() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v1/{project_id}/secrets/{secret_id}/versions/{version_id}").
		WithResponse(new(model.ShowSecretVersionResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("SecretId").
		WithJsonTag("secret_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowUserInstances() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/{version_id}/{project_id}/kms/user-instances").
		WithResponse(new(model.ShowUserInstancesResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowUserQuotas() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/{version_id}/{project_id}/kms/user-quotas").
		WithResponse(new(model.ShowUserQuotasResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForSign() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/sign").
		WithResponse(new(model.SignResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateKeyAlias() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/update-key-alias").
		WithResponse(new(model.UpdateKeyAliasResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateKeyDescription() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/update-key-description").
		WithResponse(new(model.UpdateKeyDescriptionResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateKeyRotationInterval() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/update-key-rotation-interval").
		WithResponse(new(model.UpdateKeyRotationIntervalResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateSecret() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v1/{project_id}/secrets/{secret_id}").
		WithResponse(new(model.UpdateSecretResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("SecretId").
		WithJsonTag("secret_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateSecretStage() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v1/{project_id}/secrets/{secret_id}/stages/{stage_name}").
		WithResponse(new(model.UpdateSecretStageResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("SecretId").
		WithJsonTag("secret_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("StageName").
		WithJsonTag("stage_name").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForValidateSignature() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/{version_id}/{project_id}/kms/verify").
		WithResponse(new(model.ValidateSignatureResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowVersion() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/{version_id}").
		WithResponse(new(model.ShowVersionResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("VersionId").
		WithJsonTag("version_id").
		WithLocationType(def.Path))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowVersions() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/").
		WithResponse(new(model.ShowVersionsResponse)).
		WithContentType("application/json")

	requestDef := reqDefBuilder.Build()
	return requestDef
}
