package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSessionByUserNameResponse Response Object
type ListSessionByUserNameResponse struct {

	// 会话信息列表
	SessionInfoList *[]AppSession `json:"session_info_list,omitempty"`
	HttpStatusCode  int           `json:"-"`
}

func (o ListSessionByUserNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSessionByUserNameResponse struct{}"
	}

	return strings.Join([]string{"ListSessionByUserNameResponse", string(data)}, " ")
}
