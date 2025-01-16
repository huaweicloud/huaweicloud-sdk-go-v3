package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VersionMetadata 凭据版本被标记的状态。
type VersionMetadata struct {

	// 凭据的版本号标识符，凭据对象下唯一。
	Id *string `json:"id,omitempty"`

	// 凭据版本创建时间，时间戳，即从1970年1月1日至该时间的总秒数。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 凭据版本过期时间，时间戳，即从1970年1月1日至该时间的总秒数。默认为空，凭据订阅“版本过期”事件类型时，有效期判断所依据的值。
	ExpireTime *int64 `json:"expire_time,omitempty"`

	// 加密版本凭据值的KMS主密钥ID。
	KmsKeyId *string `json:"kms_key_id,omitempty"`

	// 凭据名称。
	SecretName *string `json:"secret_name,omitempty"`

	// 凭据版本被标记的状态列表。每个版本标签对于凭据对象下版本是唯一存在的，如果创建版本时，指定的是同一凭据对象下的一个已经标记在其他版本上的状态，该标签将自动从其他版本上删除，并附加到此版本上。  如果未指定version_stage的值，则凭据管理服务会自动移动临时标签SYSCURRENT到此新版本。
	VersionStages *[]string `json:"version_stages,omitempty"`
}

func (o VersionMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionMetadata struct{}"
	}

	return strings.Join([]string{"VersionMetadata", string(data)}, " ")
}
