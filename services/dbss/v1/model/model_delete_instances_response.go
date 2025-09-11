package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstancesResponse Response Object
type DeleteInstancesResponse struct {

	// 操作结果  - success: 成功  - failed: 失败
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstancesResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstancesResponse", string(data)}, " ")
}
