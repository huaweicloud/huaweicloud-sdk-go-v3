package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyHbaConfRequestBody struct {
	BeforeConf *BeforeHbaConfOption `json:"before_conf"`

	AfterConf *AfterHbaConfOption `json:"after_conf"`
}

func (o ModifyHbaConfRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyHbaConfRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyHbaConfRequestBody", string(data)}, " ")
}
