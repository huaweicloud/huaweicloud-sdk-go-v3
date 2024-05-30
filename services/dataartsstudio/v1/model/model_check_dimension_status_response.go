package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckDimensionStatusResponse Response Object
type CheckDimensionStatusResponse struct {
	Data           *CheckDimensionStatusResultData `json:"data,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o CheckDimensionStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDimensionStatusResponse struct{}"
	}

	return strings.Join([]string{"CheckDimensionStatusResponse", string(data)}, " ")
}
