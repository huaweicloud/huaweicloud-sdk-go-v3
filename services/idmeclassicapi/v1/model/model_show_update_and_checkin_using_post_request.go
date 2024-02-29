package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUpdateAndCheckinUsingPostRequest Request Object
type ShowUpdateAndCheckinUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelVersionUpdateAndCheckinDtoVersionModel `json:"body,omitempty"`
}

func (o ShowUpdateAndCheckinUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUpdateAndCheckinUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowUpdateAndCheckinUsingPostRequest", string(data)}, " ")
}
