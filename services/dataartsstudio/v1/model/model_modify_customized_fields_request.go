package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyCustomizedFieldsRequest Request Object
type ModifyCustomizedFieldsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *CustomizedFieldsVoList `json:"body,omitempty"`
}

func (o ModifyCustomizedFieldsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyCustomizedFieldsRequest struct{}"
	}

	return strings.Join([]string{"ModifyCustomizedFieldsRequest", string(data)}, " ")
}
