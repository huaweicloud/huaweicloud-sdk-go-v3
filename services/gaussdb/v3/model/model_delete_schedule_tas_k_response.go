package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScheduleTasKResponse Response Object
type DeleteScheduleTasKResponse struct {

	// 结果。
	Result         *bool `json:"result,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o DeleteScheduleTasKResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScheduleTasKResponse struct{}"
	}

	return strings.Join([]string{"DeleteScheduleTasKResponse", string(data)}, " ")
}
