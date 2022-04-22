package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kms/v2/model"
)

type KmsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewKmsClient(hcClient *http_client.HcHttpClient) *KmsClient {
	return &KmsClient{HcClient: hcClient}
}

func KmsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// 批量添加删除密钥标签
//
// - 功能介绍：批量添加删除密钥标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) BatchCreateKmsTags(request *model.BatchCreateKmsTagsRequest) (*model.BatchCreateKmsTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateKmsTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateKmsTagsResponse), nil
	}
}

// 撤销授权
//
// - 功能介绍：撤销授权，授权用户撤销被授权用户操作密钥的权限。
// - 说明：
//    - 创建密钥的用户才能撤销该密钥授权。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) CancelGrant(request *model.CancelGrantRequest) (*model.CancelGrantResponse, error) {
	requestDef := GenReqDefForCancelGrant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelGrantResponse), nil
	}
}

// 取消计划删除密钥
//
// - 功能介绍：取消计划删除密钥。
// - 说明：密钥处于“计划删除”状态才能取消计划删除密钥。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) CancelKeyDeletion(request *model.CancelKeyDeletionRequest) (*model.CancelKeyDeletionResponse, error) {
	requestDef := GenReqDefForCancelKeyDeletion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelKeyDeletionResponse), nil
	}
}

// 退役授权
//
// - 功能介绍：退役授权，表示被授权用户不再具有授权密钥的操作权。
//   例如：用户A授权用户B可以操作密钥A/key，同时授权用户C可以撤销该授权，
//   那么用户A、B、C均可退役该授权，退役授权后，用户B不再可以使用A/key。
// - 须知：
//      可执行退役授权的主体包括：
//    - 创建授权的用户；
//    - 授权中retiring_principal指向的用户；
//    - 当授权的操作列表中包含retire-grant时，grantee_principal指向的用户。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) CancelSelfGrant(request *model.CancelSelfGrantRequest) (*model.CancelSelfGrantResponse, error) {
	requestDef := GenReqDefForCancelSelfGrant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelSelfGrantResponse), nil
	}
}

// 创建数据密钥
//
// - 功能介绍：创建数据密钥，返回结果包含明文和密文。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) CreateDatakey(request *model.CreateDatakeyRequest) (*model.CreateDatakeyResponse, error) {
	requestDef := GenReqDefForCreateDatakey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatakeyResponse), nil
	}
}

// 创建不含明文数据密钥
//
// - 功能介绍：创建数据密钥，返回结果只包含密文。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) CreateDatakeyWithoutPlaintext(request *model.CreateDatakeyWithoutPlaintextRequest) (*model.CreateDatakeyWithoutPlaintextResponse, error) {
	requestDef := GenReqDefForCreateDatakeyWithoutPlaintext()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatakeyWithoutPlaintextResponse), nil
	}
}

// 创建授权
//
// - 功能介绍：创建授权，被授权用户可以对授权密钥进行操作。
// - 说明：
//    - 服务默认主密钥（密钥别名后缀为“/default”）不可以授权。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) CreateGrant(request *model.CreateGrantRequest) (*model.CreateGrantResponse, error) {
	requestDef := GenReqDefForCreateGrant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGrantResponse), nil
	}
}

// 创建密钥
//
// 创建用户主密钥，用户主密钥可以为对称密钥或非对称密钥。
// - 对称密钥为长度为256位AES密钥或者128位SM4密钥，可用于小量数据的加密或者用于加密数据密钥。
// - 非对称密钥可以为RSA密钥对或者ECC密钥对（包含SM2密钥对），可用于加解密数据、数字签名及验签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) CreateKey(request *model.CreateKeyRequest) (*model.CreateKeyResponse, error) {
	requestDef := GenReqDefForCreateKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateKeyResponse), nil
	}
}

// 添加密钥标签
//
// - 功能介绍：添加密钥标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) CreateKmsTag(request *model.CreateKmsTagRequest) (*model.CreateKmsTagResponse, error) {
	requestDef := GenReqDefForCreateKmsTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateKmsTagResponse), nil
	}
}

// 获取密钥导入参数
//
// - 功能介绍：获取导入密钥的必要参数，包括密钥导入令牌和密钥加密公钥。
// - 说明：返回的公钥类型默认为RSA_2048。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) CreateParametersForImport(request *model.CreateParametersForImportRequest) (*model.CreateParametersForImportResponse, error) {
	requestDef := GenReqDefForCreateParametersForImport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateParametersForImportResponse), nil
	}
}

// 创建随机数
//
// - 功能介绍：
//   生成8~8192bit范围内的随机数。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) CreateRandom(request *model.CreateRandomRequest) (*model.CreateRandomResponse, error) {
	requestDef := GenReqDefForCreateRandom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRandomResponse), nil
	}
}

// 解密数据
//
// - 功能介绍：解密数据。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) DecryptData(request *model.DecryptDataRequest) (*model.DecryptDataResponse, error) {
	requestDef := GenReqDefForDecryptData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DecryptDataResponse), nil
	}
}

// 解密数据密钥
//
// - 功能介绍：解密数据密钥，用指定的主密钥解密数据密钥。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) DecryptDatakey(request *model.DecryptDatakeyRequest) (*model.DecryptDatakeyResponse, error) {
	requestDef := GenReqDefForDecryptDatakey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DecryptDatakeyResponse), nil
	}
}

// 删除密钥材料
//
// - 功能介绍：删除密钥材料信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) DeleteImportedKeyMaterial(request *model.DeleteImportedKeyMaterialRequest) (*model.DeleteImportedKeyMaterialResponse, error) {
	requestDef := GenReqDefForDeleteImportedKeyMaterial()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteImportedKeyMaterialResponse), nil
	}
}

// 计划删除密钥
//
// - 功能介绍：计划多少天后删除密钥，可设置7天～1096天内删除密钥。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) DeleteKey(request *model.DeleteKeyRequest) (*model.DeleteKeyResponse, error) {
	requestDef := GenReqDefForDeleteKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteKeyResponse), nil
	}
}

// 删除密钥标签
//
// - 功能介绍：删除密钥标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) DeleteTag(request *model.DeleteTagRequest) (*model.DeleteTagResponse, error) {
	requestDef := GenReqDefForDeleteTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTagResponse), nil
	}
}

// 禁用密钥
//
// - 功能介绍：禁用密钥，密钥禁用后不可以使用。
// - 说明：密钥为启用状态才能禁用密钥。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) DisableKey(request *model.DisableKeyRequest) (*model.DisableKeyResponse, error) {
	requestDef := GenReqDefForDisableKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableKeyResponse), nil
	}
}

// 关闭密钥轮换
//
// - 功能介绍：关闭用户主密钥轮换。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) DisableKeyRotation(request *model.DisableKeyRotationRequest) (*model.DisableKeyRotationResponse, error) {
	requestDef := GenReqDefForDisableKeyRotation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableKeyRotationResponse), nil
	}
}

// 启用密钥
//
// - 功能介绍：启用密钥，密钥启用后才可以使用。
// - 说明：密钥为禁用状态才能启用密钥。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) EnableKey(request *model.EnableKeyRequest) (*model.EnableKeyResponse, error) {
	requestDef := GenReqDefForEnableKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableKeyResponse), nil
	}
}

// 开启密钥轮换
//
// - 功能介绍：开启用户主密钥轮换。
// - 说明：
//   - 开启密钥轮换后，默认轮换间隔时间为365天。
//   - 默认主密钥及外部导入密钥不支持轮换操作。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) EnableKeyRotation(request *model.EnableKeyRotationRequest) (*model.EnableKeyRotationResponse, error) {
	requestDef := GenReqDefForEnableKeyRotation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableKeyRotationResponse), nil
	}
}

// 加密数据
//
// - 功能介绍：加密数据，用指定的用户主密钥加密数据。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) EncryptData(request *model.EncryptDataRequest) (*model.EncryptDataResponse, error) {
	requestDef := GenReqDefForEncryptData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EncryptDataResponse), nil
	}
}

// 加密数据密钥
//
// - 功能介绍：加密数据密钥，用指定的主密钥加密数据密钥。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) EncryptDatakey(request *model.EncryptDatakeyRequest) (*model.EncryptDatakeyResponse, error) {
	requestDef := GenReqDefForEncryptDatakey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EncryptDatakeyResponse), nil
	}
}

// 导入密钥材料
//
// - 功能介绍：导入密钥材料。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ImportKeyMaterial(request *model.ImportKeyMaterialRequest) (*model.ImportKeyMaterialResponse, error) {
	requestDef := GenReqDefForImportKeyMaterial()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportKeyMaterialResponse), nil
	}
}

// 查询授权列表
//
// - 功能介绍：查询密钥的授权列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ListGrants(request *model.ListGrantsRequest) (*model.ListGrantsResponse, error) {
	requestDef := GenReqDefForListGrants()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGrantsResponse), nil
	}
}

// 查询密钥信息
//
// - 功能介绍：查询密钥详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ListKeyDetail(request *model.ListKeyDetailRequest) (*model.ListKeyDetailResponse, error) {
	requestDef := GenReqDefForListKeyDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKeyDetailResponse), nil
	}
}

// 查询密钥列表
//
// - 功能介绍：查询用户所有密钥列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ListKeys(request *model.ListKeysRequest) (*model.ListKeysResponse, error) {
	requestDef := GenReqDefForListKeys()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKeysResponse), nil
	}
}

// 查询密钥实例
//
// - 功能介绍：查询密钥实例。通过标签过滤，查询指定用户主密钥的详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ListKmsByTags(request *model.ListKmsByTagsRequest) (*model.ListKmsByTagsResponse, error) {
	requestDef := GenReqDefForListKmsByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKmsByTagsResponse), nil
	}
}

// 查询项目标签
//
// - 功能介绍：查询用户在指定项目下的所有标签集合。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ListKmsTags(request *model.ListKmsTagsRequest) (*model.ListKmsTagsResponse, error) {
	requestDef := GenReqDefForListKmsTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKmsTagsResponse), nil
	}
}

// 查询可退役授权列表
//
// - 功能介绍：查询用户可以退役的授权列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ListRetirableGrants(request *model.ListRetirableGrantsRequest) (*model.ListRetirableGrantsResponse, error) {
	requestDef := GenReqDefForListRetirableGrants()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRetirableGrantsResponse), nil
	}
}

// 查询密钥轮换状态
//
// - 功能介绍：查询用户主密钥轮换状态。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ShowKeyRotationStatus(request *model.ShowKeyRotationStatusRequest) (*model.ShowKeyRotationStatusResponse, error) {
	requestDef := GenReqDefForShowKeyRotationStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKeyRotationStatusResponse), nil
	}
}

// 查询密钥标签
//
// - 功能介绍：查询密钥标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ShowKmsTags(request *model.ShowKmsTagsRequest) (*model.ShowKmsTagsResponse, error) {
	requestDef := GenReqDefForShowKmsTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKmsTagsResponse), nil
	}
}

// 查询公钥信息
//
// - 查询用户指定非对称密钥的公钥信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ShowPublicKey(request *model.ShowPublicKeyRequest) (*model.ShowPublicKeyResponse, error) {
	requestDef := GenReqDefForShowPublicKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPublicKeyResponse), nil
	}
}

// 查询实例数
//
// - 功能介绍：查询实例数，获取用户已经创建的用户主密钥数量。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ShowUserInstances(request *model.ShowUserInstancesRequest) (*model.ShowUserInstancesResponse, error) {
	requestDef := GenReqDefForShowUserInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserInstancesResponse), nil
	}
}

// 查询配额
//
// - 功能介绍：查询配额，查询用户可以创建的用户主密钥配额总数及当前使用量信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ShowUserQuotas(request *model.ShowUserQuotasRequest) (*model.ShowUserQuotasResponse, error) {
	requestDef := GenReqDefForShowUserQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserQuotasResponse), nil
	}
}

// 签名数据
//
// - 功能介绍：使用非对称密钥的私钥对消息或消息摘要进行数字签名。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) Sign(request *model.SignRequest) (*model.SignResponse, error) {
	requestDef := GenReqDefForSign()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SignResponse), nil
	}
}

// 修改密钥别名
//
// - 功能介绍：修改用户主密钥别名。
// - 说明：
//    - 服务默认主密钥（密钥别名后缀为“/default”）不可以修改。
//    - 密钥处于“计划删除”状态，密钥别名不可以修改。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) UpdateKeyAlias(request *model.UpdateKeyAliasRequest) (*model.UpdateKeyAliasResponse, error) {
	requestDef := GenReqDefForUpdateKeyAlias()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateKeyAliasResponse), nil
	}
}

// 修改密钥描述
//
// - 功能介绍：修改用户主密钥描述信息。
// - 说明：
//    - 服务默认主密钥（密钥别名后缀为“/default”）不可以修改。
//    - 密钥处于“计划删除”状态，密钥描述不可以修改。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) UpdateKeyDescription(request *model.UpdateKeyDescriptionRequest) (*model.UpdateKeyDescriptionResponse, error) {
	requestDef := GenReqDefForUpdateKeyDescription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateKeyDescriptionResponse), nil
	}
}

// 修改密钥轮换周期
//
// - 功能介绍：修改用户主密钥轮换周期。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) UpdateKeyRotationInterval(request *model.UpdateKeyRotationIntervalRequest) (*model.UpdateKeyRotationIntervalResponse, error) {
	requestDef := GenReqDefForUpdateKeyRotationInterval()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateKeyRotationIntervalResponse), nil
	}
}

// 验证签名
//
// - 功能介绍：使用非对称密钥的私钥对消息或消息摘要进行签名验证。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ValidateSignature(request *model.ValidateSignatureRequest) (*model.ValidateSignatureResponse, error) {
	requestDef := GenReqDefForValidateSignature()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateSignatureResponse), nil
	}
}

// 查询指定版本信息
//
// - 功能介绍：查指定API版本信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ShowVersion(request *model.ShowVersionRequest) (*model.ShowVersionResponse, error) {
	requestDef := GenReqDefForShowVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVersionResponse), nil
	}
}

// 查询版本信息列表
//
// - 功能介绍：查询API版本信息列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *KmsClient) ShowVersions(request *model.ShowVersionsRequest) (*model.ShowVersionsResponse, error) {
	requestDef := GenReqDefForShowVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVersionsResponse), nil
	}
}
