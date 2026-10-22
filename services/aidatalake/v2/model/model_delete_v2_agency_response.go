package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteV2AgencyResponse Response Object
type DeleteV2AgencyResponse struct {

	// **参数解释**：执行请求是否成功。 **取值范围**：   - true：执行成功。   - false：执行失败。
	IsSuccess *bool `json:"is_success,omitempty"`

	// **参数解释**：系统提示信息。 **取值范围**：字符类型不限，长度为0~4096个字符。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteV2AgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteV2AgencyResponse struct{}"
	}

	return strings.Join([]string{"DeleteV2AgencyResponse", string(data)}, " ")
}
