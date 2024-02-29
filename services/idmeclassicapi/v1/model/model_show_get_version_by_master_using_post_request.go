package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGetVersionByMasterUsingPostRequest Request Object
type ShowGetVersionByMasterUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelVersionMasterQueryDto `json:"body,omitempty"`
}

func (o ShowGetVersionByMasterUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGetVersionByMasterUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowGetVersionByMasterUsingPostRequest", string(data)}, " ")
}
