package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchListStructDetailResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 批量查询灾备初始化对象详情返回列表
	Results        *[]QueryStructDetailResp `json:"results,omitempty" xml:"results"`
	HttpStatusCode int                      `json:"-"`
}

func (o BatchListStructDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListStructDetailResponse struct{}"
	}

	return strings.Join([]string{"BatchListStructDetailResponse", string(data)}, " ")
}
