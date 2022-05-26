package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kms/v2/model"
)

type BatchCreateKmsTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateKmsTagsInvoker) Invoke() (*model.BatchCreateKmsTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateKmsTagsResponse), nil
	}
}

type CancelGrantInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelGrantInvoker) Invoke() (*model.CancelGrantResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelGrantResponse), nil
	}
}

type CancelKeyDeletionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelKeyDeletionInvoker) Invoke() (*model.CancelKeyDeletionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelKeyDeletionResponse), nil
	}
}

type CancelSelfGrantInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelSelfGrantInvoker) Invoke() (*model.CancelSelfGrantResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelSelfGrantResponse), nil
	}
}

type CreateDatakeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDatakeyInvoker) Invoke() (*model.CreateDatakeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDatakeyResponse), nil
	}
}

type CreateDatakeyWithoutPlaintextInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDatakeyWithoutPlaintextInvoker) Invoke() (*model.CreateDatakeyWithoutPlaintextResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDatakeyWithoutPlaintextResponse), nil
	}
}

type CreateGrantInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGrantInvoker) Invoke() (*model.CreateGrantResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGrantResponse), nil
	}
}

type CreateKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateKeyInvoker) Invoke() (*model.CreateKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateKeyResponse), nil
	}
}

type CreateKmsTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateKmsTagInvoker) Invoke() (*model.CreateKmsTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateKmsTagResponse), nil
	}
}

type CreateParametersForImportInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateParametersForImportInvoker) Invoke() (*model.CreateParametersForImportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateParametersForImportResponse), nil
	}
}

type CreateRandomInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRandomInvoker) Invoke() (*model.CreateRandomResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRandomResponse), nil
	}
}

type DecryptDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *DecryptDataInvoker) Invoke() (*model.DecryptDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DecryptDataResponse), nil
	}
}

type DecryptDatakeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DecryptDatakeyInvoker) Invoke() (*model.DecryptDatakeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DecryptDatakeyResponse), nil
	}
}

type DeleteImportedKeyMaterialInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteImportedKeyMaterialInvoker) Invoke() (*model.DeleteImportedKeyMaterialResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteImportedKeyMaterialResponse), nil
	}
}

type DeleteKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteKeyInvoker) Invoke() (*model.DeleteKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteKeyResponse), nil
	}
}

type DeleteTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTagInvoker) Invoke() (*model.DeleteTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTagResponse), nil
	}
}

type DisableKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableKeyInvoker) Invoke() (*model.DisableKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableKeyResponse), nil
	}
}

type DisableKeyRotationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableKeyRotationInvoker) Invoke() (*model.DisableKeyRotationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableKeyRotationResponse), nil
	}
}

type EnableKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableKeyInvoker) Invoke() (*model.EnableKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableKeyResponse), nil
	}
}

type EnableKeyRotationInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableKeyRotationInvoker) Invoke() (*model.EnableKeyRotationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableKeyRotationResponse), nil
	}
}

type EncryptDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *EncryptDataInvoker) Invoke() (*model.EncryptDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EncryptDataResponse), nil
	}
}

type EncryptDatakeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *EncryptDatakeyInvoker) Invoke() (*model.EncryptDatakeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EncryptDatakeyResponse), nil
	}
}

type ImportKeyMaterialInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportKeyMaterialInvoker) Invoke() (*model.ImportKeyMaterialResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportKeyMaterialResponse), nil
	}
}

type ListGrantsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGrantsInvoker) Invoke() (*model.ListGrantsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGrantsResponse), nil
	}
}

type ListKeyDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListKeyDetailInvoker) Invoke() (*model.ListKeyDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListKeyDetailResponse), nil
	}
}

type ListKeysInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListKeysInvoker) Invoke() (*model.ListKeysResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListKeysResponse), nil
	}
}

type ListKmsByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListKmsByTagsInvoker) Invoke() (*model.ListKmsByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListKmsByTagsResponse), nil
	}
}

type ListKmsTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListKmsTagsInvoker) Invoke() (*model.ListKmsTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListKmsTagsResponse), nil
	}
}

type ListRetirableGrantsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRetirableGrantsInvoker) Invoke() (*model.ListRetirableGrantsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRetirableGrantsResponse), nil
	}
}

type ShowKeyRotationStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowKeyRotationStatusInvoker) Invoke() (*model.ShowKeyRotationStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowKeyRotationStatusResponse), nil
	}
}

type ShowKmsTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowKmsTagsInvoker) Invoke() (*model.ShowKmsTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowKmsTagsResponse), nil
	}
}

type ShowPublicKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPublicKeyInvoker) Invoke() (*model.ShowPublicKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPublicKeyResponse), nil
	}
}

type ShowUserInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUserInstancesInvoker) Invoke() (*model.ShowUserInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUserInstancesResponse), nil
	}
}

type ShowUserQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUserQuotasInvoker) Invoke() (*model.ShowUserQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUserQuotasResponse), nil
	}
}

type SignInvoker struct {
	*invoker.BaseInvoker
}

func (i *SignInvoker) Invoke() (*model.SignResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SignResponse), nil
	}
}

type UpdateKeyAliasInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateKeyAliasInvoker) Invoke() (*model.UpdateKeyAliasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateKeyAliasResponse), nil
	}
}

type UpdateKeyDescriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateKeyDescriptionInvoker) Invoke() (*model.UpdateKeyDescriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateKeyDescriptionResponse), nil
	}
}

type UpdateKeyRotationIntervalInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateKeyRotationIntervalInvoker) Invoke() (*model.UpdateKeyRotationIntervalResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateKeyRotationIntervalResponse), nil
	}
}

type ValidateSignatureInvoker struct {
	*invoker.BaseInvoker
}

func (i *ValidateSignatureInvoker) Invoke() (*model.ValidateSignatureResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ValidateSignatureResponse), nil
	}
}

type ShowVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVersionInvoker) Invoke() (*model.ShowVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVersionResponse), nil
	}
}

type ShowVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVersionsInvoker) Invoke() (*model.ShowVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVersionsResponse), nil
	}
}
