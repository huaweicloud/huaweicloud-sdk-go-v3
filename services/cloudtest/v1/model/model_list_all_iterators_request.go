package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllIteratorsRequest Request Object
type ListAllIteratorsRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`
}

func (o ListAllIteratorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllIteratorsRequest struct{}"
	}

	return strings.Join([]string{"ListAllIteratorsRequest", string(data)}, " ")
}
