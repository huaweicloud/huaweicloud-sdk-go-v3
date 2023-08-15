package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestSetCohostBody 申请联席主持人请求。
type RestSetCohostBody struct {

	// - 0：撤销联席主持人 - 1：设置联席主持人
	IsCohost int32 `json:"isCohost"`
}

func (o RestSetCohostBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestSetCohostBody struct{}"
	}

	return strings.Join([]string{"RestSetCohostBody", string(data)}, " ")
}
