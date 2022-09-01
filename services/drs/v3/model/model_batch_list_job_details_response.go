package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchListJobDetailsResponse struct {

	// 任务数量
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 任务详细信息
	Results        *[]QueryJobResp `json:"results,omitempty" xml:"results"`
	HttpStatusCode int             `json:"-"`
}

func (o BatchListJobDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListJobDetailsResponse struct{}"
	}

	return strings.Join([]string{"BatchListJobDetailsResponse", string(data)}, " ")
}
