package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUpdateByAdminUsingPostRequest Request Object
type ShowUpdateByAdminUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModel `json:"body,omitempty"`
}

func (o ShowUpdateByAdminUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUpdateByAdminUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowUpdateByAdminUsingPostRequest", string(data)}, " ")
}
