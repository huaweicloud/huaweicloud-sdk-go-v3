package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserGpgKeysRequest Request Object
type ListUserGpgKeysRequest struct {

	// **参数解释：** key的标题名称。 **取值范围：** 字符串长度不少于1，不超过2000。
	Query *string `json:"query,omitempty"`
}

func (o ListUserGpgKeysRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserGpgKeysRequest struct{}"
	}

	return strings.Join([]string{"ListUserGpgKeysRequest", string(data)}, " ")
}
