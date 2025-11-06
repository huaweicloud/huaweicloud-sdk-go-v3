package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHttpsPasswordSettingResponse Response Object
type UpdateHttpsPasswordSettingResponse struct {

	// **参数解释：** 是否用https的认证方式 true,false。 **取值范围：** 字符串长度不少于1，不超过1000。
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateHttpsPasswordSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpsPasswordSettingResponse struct{}"
	}

	return strings.Join([]string{"UpdateHttpsPasswordSettingResponse", string(data)}, " ")
}
