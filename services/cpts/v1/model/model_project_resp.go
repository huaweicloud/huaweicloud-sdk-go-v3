package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProjectResp struct {

	// 状态码
	Code *string `json:"code,omitempty" xml:"code"`

	// 描述
	Message *string `json:"message,omitempty" xml:"message"`
}

func (o ProjectResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectResp struct{}"
	}

	return strings.Join([]string{"ProjectResp", string(data)}, " ")
}
