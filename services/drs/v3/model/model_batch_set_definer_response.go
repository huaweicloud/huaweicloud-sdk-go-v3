package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchSetDefinerResponse struct {
	// 总数

	Count *int32 `json:"count,omitempty"`
	// 批量修改任务返回列表

	Results        *[]ModifyJobResp `json:"results,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o BatchSetDefinerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetDefinerResponse struct{}"
	}

	return strings.Join([]string{"BatchSetDefinerResponse", string(data)}, " ")
}
