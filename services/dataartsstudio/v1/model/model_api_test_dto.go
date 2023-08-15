package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiTestDto struct {

	// 请求体
	Body *string `json:"body,omitempty"`

	Paras *ApiTestParas `json:"paras,omitempty"`
}

func (o ApiTestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiTestDto struct{}"
	}

	return strings.Join([]string{"ApiTestDto", string(data)}, " ")
}
