package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunCheckDataResponse struct {

	// 检查数据完成返回success。
	Result *string `json:"result,omitempty"`

	Data           *CheckRestInfo `json:"data,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o RunCheckDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCheckDataResponse struct{}"
	}

	return strings.Join([]string{"RunCheckDataResponse", string(data)}, " ")
}
