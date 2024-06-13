package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourcePoolsRequest Request Object
type ListResourcePoolsRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`
}

func (o ListResourcePoolsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcePoolsRequest struct{}"
	}

	return strings.Join([]string{"ListResourcePoolsRequest", string(data)}, " ")
}
