package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserPopupInfoRequest Request Object
type ListUserPopupInfoRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`
}

func (o ListUserPopupInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserPopupInfoRequest struct{}"
	}

	return strings.Join([]string{"ListUserPopupInfoRequest", string(data)}, " ")
}
