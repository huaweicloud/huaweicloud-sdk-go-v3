package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchRemoveOrgsFromChannelRequestBody struct {

	// 组织名称列表
	OrgNames []string `json:"org_names"`
}

func (o BatchRemoveOrgsFromChannelRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRemoveOrgsFromChannelRequestBody struct{}"
	}

	return strings.Join([]string{"BatchRemoveOrgsFromChannelRequestBody", string(data)}, " ")
}
