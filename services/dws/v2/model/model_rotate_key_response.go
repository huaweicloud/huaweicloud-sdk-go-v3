package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RotateKeyResponse Response Object
type RotateKeyResponse struct {

	// **参数解释**： KMS密钥名称 **取值范围**： 不涉及。
	MasterKeyName *string `json:"master_key_name,omitempty"`

	// **参数解释**： 集群ID。 **取值范围**： 不涉及。
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 轮转时间。 **取值范围**： 不涉及。
	RotateKeyTime  *string `json:"rotate_key_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RotateKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RotateKeyResponse struct{}"
	}

	return strings.Join([]string{"RotateKeyResponse", string(data)}, " ")
}
