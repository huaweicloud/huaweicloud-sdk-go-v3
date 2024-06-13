package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSystemConfigsRequest Request Object
type ShowSystemConfigsRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	Body *CommRequestSystemConfig `json:"body,omitempty"`
}

func (o ShowSystemConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSystemConfigsRequest struct{}"
	}

	return strings.Join([]string{"ShowSystemConfigsRequest", string(data)}, " ")
}
