package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStandardByIdRequest Request Object
type ShowStandardByIdRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 实体id
	Id string `json:"id"`
}

func (o ShowStandardByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStandardByIdRequest struct{}"
	}

	return strings.Join([]string{"ShowStandardByIdRequest", string(data)}, " ")
}
