package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDataFilteringResultResponse Response Object
type ShowDataFilteringResultResponse struct {

	// 数据过滤规则校验成功的数量
	SuccessCount *int64 `json:"success_count,omitempty"`

	// 数据过滤规则校验失败的数量
	FailedCount *int64 `json:"failed_count,omitempty"`

	// 库表过滤规则校验结果
	DbObjectFilteringResult *[]DbObjectFilteringResult `json:"db_object_filtering_result,omitempty"`
	HttpStatusCode          int                        `json:"-"`
}

func (o ShowDataFilteringResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataFilteringResultResponse struct{}"
	}

	return strings.Join([]string{"ShowDataFilteringResultResponse", string(data)}, " ")
}
