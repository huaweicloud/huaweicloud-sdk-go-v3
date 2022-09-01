package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Creator struct {

	// 指定工作项的创建者用户id
	UserId string `json:"user_id" xml:"user_id"`
}

func (o Creator) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Creator struct{}"
	}

	return strings.Join([]string{"Creator", string(data)}, " ")
}
