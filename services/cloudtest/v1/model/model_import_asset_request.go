package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportAssetRequest Request Object
type ImportAssetRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	Body *ImportAssetRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ImportAssetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportAssetRequest struct{}"
	}

	return strings.Join([]string{"ImportAssetRequest", string(data)}, " ")
}
