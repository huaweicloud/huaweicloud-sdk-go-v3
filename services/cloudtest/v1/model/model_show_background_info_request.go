package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackgroundInfoRequest Request Object
type ShowBackgroundInfoRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`
}

func (o ShowBackgroundInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackgroundInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowBackgroundInfoRequest", string(data)}, " ")
}
