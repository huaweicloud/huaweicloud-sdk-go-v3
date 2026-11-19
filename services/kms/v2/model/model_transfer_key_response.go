package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TransferKeyResponse Response Object
type TransferKeyResponse struct {

	// **参数解释：** 转移后密钥的ID **取值范围：** 不涉及
	KeyId *string `json:"key_id,omitempty"`

	// **参数解释：** 密钥所属的目标密钥空间ID **取值范围：** 不涉及
	KeyspaceId     *string `json:"keyspace_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o TransferKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferKeyResponse struct{}"
	}

	return strings.Join([]string{"TransferKeyResponse", string(data)}, " ")
}
