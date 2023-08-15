package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSha256Request Request Object
type ShowSha256Request struct {

	// 关键字
	Key string `json:"key"`
}

func (o ShowSha256Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSha256Request struct{}"
	}

	return strings.Join([]string{"ShowSha256Request", string(data)}, " ")
}
