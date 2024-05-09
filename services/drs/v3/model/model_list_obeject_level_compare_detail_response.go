package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListObejectLevelCompareDetailResponse Response Object
type ListObejectLevelCompareDetailResponse struct {

	// 对比数量。
	Count *int32 `json:"count,omitempty"`

	// 对象级对比详情。
	CompareDetail  *[]ObjectsCompareDetailInfo `json:"compare_detail,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListObejectLevelCompareDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListObejectLevelCompareDetailResponse struct{}"
	}

	return strings.Join([]string{"ListObejectLevelCompareDetailResponse", string(data)}, " ")
}
