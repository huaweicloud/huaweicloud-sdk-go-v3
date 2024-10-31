package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportFactorRequest Request Object
type ImportFactorRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 资产库ID
	AssetId string `json:"asset_id"`

	Body *ImportFactorRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ImportFactorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportFactorRequest struct{}"
	}

	return strings.Join([]string{"ImportFactorRequest", string(data)}, " ")
}
