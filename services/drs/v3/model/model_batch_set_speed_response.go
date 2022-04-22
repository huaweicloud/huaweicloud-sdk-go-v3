package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchSetSpeedResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 批量修改任务返回列表
	Results        *[]ModifyJobResp `json:"results,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o BatchSetSpeedResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetSpeedResponse struct{}"
	}

	return strings.Join([]string{"BatchSetSpeedResponse", string(data)}, " ")
}
