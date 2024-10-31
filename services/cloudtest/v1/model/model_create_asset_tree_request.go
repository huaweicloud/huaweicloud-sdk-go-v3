package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAssetTreeRequest Request Object
type CreateAssetTreeRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 资产ID
	AssetId string `json:"asset_id"`

	// 父目录ID
	ParentId string `json:"parent_id"`

	Body *CommRequestAssetTree `json:"body,omitempty"`
}

func (o CreateAssetTreeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssetTreeRequest struct{}"
	}

	return strings.Join([]string{"CreateAssetTreeRequest", string(data)}, " ")
}
