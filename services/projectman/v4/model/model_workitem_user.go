package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkitemUser struct {

	// 用户32位uuid
	Id *string `json:"id,omitempty"`

	// 用户名
	Name *string `json:"name,omitempty"`

	// 昵称
	NickName *string `json:"nick_name,omitempty"`
}

func (o WorkitemUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkitemUser struct{}"
	}

	return strings.Join([]string{"WorkitemUser", string(data)}, " ")
}
