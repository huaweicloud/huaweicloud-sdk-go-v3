package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCompareJobResponse Response Object
type DeleteCompareJobResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteCompareJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCompareJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteCompareJobResponse", string(data)}, " ")
}
