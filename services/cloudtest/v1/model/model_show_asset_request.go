package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAssetRequest Request Object
type ShowAssetRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`
}

func (o ShowAssetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssetRequest struct{}"
	}

	return strings.Join([]string{"ShowAssetRequest", string(data)}, " ")
}
