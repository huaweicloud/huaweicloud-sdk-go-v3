package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportFactorRequest Request Object
type ExportFactorRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 资产库ID
	AssetId string `json:"asset_id"`

	Body *CommRequestAssetExportParam `json:"body,omitempty"`
}

func (o ExportFactorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportFactorRequest struct{}"
	}

	return strings.Join([]string{"ExportFactorRequest", string(data)}, " ")
}
