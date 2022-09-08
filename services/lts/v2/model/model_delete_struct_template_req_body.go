package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 删除结构化配置参数
type DeleteStructTemplateReqBody struct {

	// 结构化规则ID
	Id string `json:"id"`
}

func (o DeleteStructTemplateReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStructTemplateReqBody struct{}"
	}

	return strings.Join([]string{"DeleteStructTemplateReqBody", string(data)}, " ")
}
