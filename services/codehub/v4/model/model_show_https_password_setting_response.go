package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpsPasswordSettingResponse Response Object
type ShowHttpsPasswordSettingResponse struct {

	// **参数解释：** 是否用https的认证方式 true,false。 **取值范围：** 字符串长度不少于1，不超过1000。
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowHttpsPasswordSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpsPasswordSettingResponse struct{}"
	}

	return strings.Join([]string{"ShowHttpsPasswordSettingResponse", string(data)}, " ")
}
