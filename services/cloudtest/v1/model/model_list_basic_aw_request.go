package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBasicAwRequest Request Object
type ListBasicAwRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// AW ID
	AwId string `json:"aw_id"`
}

func (o ListBasicAwRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBasicAwRequest struct{}"
	}

	return strings.Join([]string{"ListBasicAwRequest", string(data)}, " ")
}
