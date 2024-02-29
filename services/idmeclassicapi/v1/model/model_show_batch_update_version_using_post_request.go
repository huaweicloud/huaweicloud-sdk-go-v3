package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchUpdateVersionUsingPostRequest Request Object
type ShowBatchUpdateVersionUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListVersionModelVersionUpdateDto `json:"body,omitempty"`
}

func (o ShowBatchUpdateVersionUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchUpdateVersionUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowBatchUpdateVersionUsingPostRequest", string(data)}, " ")
}
