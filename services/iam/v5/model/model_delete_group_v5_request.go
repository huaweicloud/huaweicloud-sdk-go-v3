package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGroupV5Request Request Object
type DeleteGroupV5Request struct {

	// 用户组ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	GroupId string `json:"group_id"`
}

func (o DeleteGroupV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGroupV5Request struct{}"
	}

	return strings.Join([]string{"DeleteGroupV5Request", string(data)}, " ")
}
