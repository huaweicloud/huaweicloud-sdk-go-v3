package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchGetSecretValueResponse Response Object
type BatchGetSecretValueResponse struct {

	// **参数解释：**  返回的凭据版本总数 **取值范围：** 不涉及
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释：**  凭据版本值 **取值范围：** 不涉及
	Versions       *[]Version `json:"versions,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o BatchGetSecretValueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchGetSecretValueResponse struct{}"
	}

	return strings.Join([]string{"BatchGetSecretValueResponse", string(data)}, " ")
}
