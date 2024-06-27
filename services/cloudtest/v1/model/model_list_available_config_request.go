package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailableConfigRequest Request Object
type ListAvailableConfigRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`
}

func (o ListAvailableConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableConfigRequest struct{}"
	}

	return strings.Join([]string{"ListAvailableConfigRequest", string(data)}, " ")
}
