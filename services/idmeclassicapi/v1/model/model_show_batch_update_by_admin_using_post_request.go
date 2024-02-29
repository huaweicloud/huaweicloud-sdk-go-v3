package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchUpdateByAdminUsingPostRequest Request Object
type ShowBatchUpdateByAdminUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListVersionModel `json:"body,omitempty"`
}

func (o ShowBatchUpdateByAdminUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchUpdateByAdminUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ShowBatchUpdateByAdminUsingPostRequest", string(data)}, " ")
}
