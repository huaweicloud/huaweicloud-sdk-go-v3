package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateV2AgencyResponse Response Object
type CreateV2AgencyResponse struct {

	// **参数解释**：委托名称。 **取值范围**：长度1~128个字符，只能由英文字母、数字及下划线组成。
	AgencyName *string `json:"agency_name,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateV2AgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateV2AgencyResponse struct{}"
	}

	return strings.Join([]string{"CreateV2AgencyResponse", string(data)}, " ")
}
