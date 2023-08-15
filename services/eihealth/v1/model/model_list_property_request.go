package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPropertyRequest Request Object
type ListPropertyRequest struct {

	// 属性，支持LABEL/PUBLISHER/CATEGORY
	Property string `json:"property"`
}

func (o ListPropertyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPropertyRequest struct{}"
	}

	return strings.Join([]string{"ListPropertyRequest", string(data)}, " ")
}
