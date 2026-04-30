package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthObject 鉴权对象
type AuthObject struct {
	AppId *string `json:"app_id,omitempty"`
}

func (o AuthObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthObject struct{}"
	}

	return strings.Join([]string{"AuthObject", string(data)}, " ")
}
