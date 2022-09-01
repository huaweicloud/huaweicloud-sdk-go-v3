package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchChangeDataResponse struct {

	// 批量数据加工响应列表
	Results *[]DataTransformationResp `json:"results,omitempty" xml:"results"`

	// 总数
	Count          *int32 `json:"count,omitempty" xml:"count"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchChangeDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchChangeDataResponse struct{}"
	}

	return strings.Join([]string{"BatchChangeDataResponse", string(data)}, " ")
}
