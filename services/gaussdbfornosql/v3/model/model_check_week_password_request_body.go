package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckWeekPasswordRequestBody struct {

	// 密码
	Password string `json:"password"`
}

func (o CheckWeekPasswordRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckWeekPasswordRequestBody struct{}"
	}

	return strings.Join([]string{"CheckWeekPasswordRequestBody", string(data)}, " ")
}
