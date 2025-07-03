package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyFlowSourcesRequestBody 修改入流信息请求体
type ModifyFlowSourcesRequestBody struct {
	Source *FlowSource `json:"source"`
}

func (o ModifyFlowSourcesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyFlowSourcesRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyFlowSourcesRequestBody", string(data)}, " ")
}
