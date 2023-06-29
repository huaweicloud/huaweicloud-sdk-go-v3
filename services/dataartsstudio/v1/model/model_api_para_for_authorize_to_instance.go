package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiParaForAuthorizeToInstance struct {

	// 截止时间
	Time *string `json:"time,omitempty"`

	// app编号列表
	AppIds *[]string `json:"app_ids,omitempty"`
}

func (o ApiParaForAuthorizeToInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiParaForAuthorizeToInstance struct{}"
	}

	return strings.Join([]string{"ApiParaForAuthorizeToInstance", string(data)}, " ")
}
