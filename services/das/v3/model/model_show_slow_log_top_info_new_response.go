package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSlowLogTopInfoNewResponse Response Object
type ShowSlowLogTopInfoNewResponse struct {

	// Top用户列表
	TopUserList *[]SlowLogTopInfo `json:"top_user_list,omitempty"`

	// Top IP列表
	TopIpList *[]SlowLogTopInfo `json:"top_ip_list,omitempty"`

	// Top数据库列表
	TopDbList      *[]SlowLogTopInfo `json:"top_db_list,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowSlowLogTopInfoNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowLogTopInfoNewResponse struct{}"
	}

	return strings.Join([]string{"ShowSlowLogTopInfoNewResponse", string(data)}, " ")
}
