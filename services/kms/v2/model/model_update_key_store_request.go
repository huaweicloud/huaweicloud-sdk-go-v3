package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKeyStoreRequest Request Object
type UpdateKeyStoreRequest struct {

	// **参数解释：** 密钥库ID **约束限制：** 满足正则表达式\"[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$\" **取值范围：** 不涉及 **默认取值：** 不涉及
	KeystoreId string `json:"keystore_id"`

	Body *UpdateKeyStoreRequestBody `json:"body,omitempty"`
}

func (o UpdateKeyStoreRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeyStoreRequest struct{}"
	}

	return strings.Join([]string{"UpdateKeyStoreRequest", string(data)}, " ")
}
