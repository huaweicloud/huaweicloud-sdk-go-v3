package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGroupV5Request Request Object
type UpdateGroupV5Request struct {

	// 用户组ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	GroupId string `json:"group_id"`

	Body *UpdateGroupReqBody `json:"body,omitempty"`
}

func (o UpdateGroupV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGroupV5Request struct{}"
	}

	return strings.Join([]string{"UpdateGroupV5Request", string(data)}, " ")
}
