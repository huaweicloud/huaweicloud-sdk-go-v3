package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFreeDeclarationResponse Response Object
type ShowFreeDeclarationResponse struct {

	// 实际的数据类型：单个对象，集合 或 NULL
	Value          *bool `json:"value,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowFreeDeclarationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFreeDeclarationResponse struct{}"
	}

	return strings.Join([]string{"ShowFreeDeclarationResponse", string(data)}, " ")
}
