package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAssetRequest Request Object
type DeleteAssetRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 资产库ID
	Id string `json:"id"`
}

func (o DeleteAssetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAssetRequest struct{}"
	}

	return strings.Join([]string{"DeleteAssetRequest", string(data)}, " ")
}
