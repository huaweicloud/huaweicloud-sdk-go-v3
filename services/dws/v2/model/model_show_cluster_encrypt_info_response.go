package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterEncryptInfoResponse Response Object
type ShowClusterEncryptInfoResponse struct {

	// **参数解释**： 加密ID，仅DWS内部使用。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 集群ID。 **取值范围**： 不涉及。
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： KMS密钥ID。 **取值范围**： 不涉及。
	MasterKeyId *string `json:"master_key_id,omitempty"`

	// **参数解释**： KMS密钥名称。 **取值范围**： 不涉及。
	MasterKeyName *string `json:"master_key_name,omitempty"`

	// **参数解释**： 最后做密钥轮转的时间。 **取值范围**： 不涉及。
	LastRotateKeyTime *string `json:"last_rotate_key_time,omitempty"`

	// **参数解释**： 加密方式。 **取值范围**： 不涉及。
	CryptAlgorithm *string `json:"crypt_algorithm,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowClusterEncryptInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterEncryptInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterEncryptInfoResponse", string(data)}, " ")
}
