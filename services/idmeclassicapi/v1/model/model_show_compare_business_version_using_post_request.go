package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCompareBusinessVersionUsingPostRequest Request Object
type ShowCompareBusinessVersionUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelCompareVersionVo `json:"body,omitempty"`
}

func (o ShowCompareBusinessVersionUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCompareBusinessVersionUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowCompareBusinessVersionUsingPostRequest", string(data)}, " ")
}
