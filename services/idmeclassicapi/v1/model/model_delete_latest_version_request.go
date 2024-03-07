package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLatestVersionRequest Request Object
type DeleteLatestVersionRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelVersionMasterModifierDto `json:"body,omitempty"`
}

func (o DeleteLatestVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLatestVersionRequest struct{}"
	}

	return strings.Join([]string{"DeleteLatestVersionRequest", string(data)}, " ")
}
