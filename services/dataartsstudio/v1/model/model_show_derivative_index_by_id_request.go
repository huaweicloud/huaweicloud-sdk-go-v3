package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDerivativeIndexByIdRequest Request Object
type ShowDerivativeIndexByIdRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 实体id
	Id string `json:"id"`

	// 是否查询最新的
	Latest *bool `json:"latest,omitempty"`
}

func (o ShowDerivativeIndexByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDerivativeIndexByIdRequest struct{}"
	}

	return strings.Join([]string{"ShowDerivativeIndexByIdRequest", string(data)}, " ")
}
