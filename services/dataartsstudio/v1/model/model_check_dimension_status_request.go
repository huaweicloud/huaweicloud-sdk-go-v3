package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckDimensionStatusRequest Request Object
type CheckDimensionStatusRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`
}

func (o CheckDimensionStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDimensionStatusRequest struct{}"
	}

	return strings.Join([]string{"CheckDimensionStatusRequest", string(data)}, " ")
}
