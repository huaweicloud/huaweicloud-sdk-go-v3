package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectFieldConfigsRequest Request Object
type ListProjectFieldConfigsRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`
}

func (o ListProjectFieldConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectFieldConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListProjectFieldConfigsRequest", string(data)}, " ")
}
