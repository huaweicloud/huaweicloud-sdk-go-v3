package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstancesNewResponse Response Object
type DeleteInstancesNewResponse struct {

	// 响应状态
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteInstancesNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstancesNewResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstancesNewResponse", string(data)}, " ")
}
