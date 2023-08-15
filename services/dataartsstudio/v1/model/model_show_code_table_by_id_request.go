package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCodeTableByIdRequest Request Object
type ShowCodeTableByIdRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 实体id
	Id string `json:"id"`
}

func (o ShowCodeTableByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCodeTableByIdRequest struct{}"
	}

	return strings.Join([]string{"ShowCodeTableByIdRequest", string(data)}, " ")
}
