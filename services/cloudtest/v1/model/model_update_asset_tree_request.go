package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAssetTreeRequest Request Object
type UpdateAssetTreeRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	Body *CommRequestUpdateAssetTreeParam `json:"body,omitempty"`
}

func (o UpdateAssetTreeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssetTreeRequest struct{}"
	}

	return strings.Join([]string{"UpdateAssetTreeRequest", string(data)}, " ")
}
