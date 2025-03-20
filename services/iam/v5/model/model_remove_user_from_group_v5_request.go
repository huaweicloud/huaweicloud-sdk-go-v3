package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveUserFromGroupV5Request Request Object
type RemoveUserFromGroupV5Request struct {

	// 用户组ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	GroupId string `json:"group_id"`

	Body *RemoveUserFromGroupReqBody `json:"body,omitempty"`
}

func (o RemoveUserFromGroupV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveUserFromGroupV5Request struct{}"
	}

	return strings.Join([]string{"RemoveUserFromGroupV5Request", string(data)}, " ")
}
