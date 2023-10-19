package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Description 实例描述。
type Description struct {

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`
}

func (o Description) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Description struct{}"
	}

	return strings.Join([]string{"Description", string(data)}, " ")
}
