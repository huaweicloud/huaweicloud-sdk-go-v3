package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CinderListVolumeTransfersRequest Request Object
type CinderListVolumeTransfersRequest struct {

	// 返回结果个数限制，取值为大 于0的整数
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，偏移量为一个大于0小 于云硬盘过户记录总个数的整 数，表示查询该偏移量后面的 所有的云硬盘过户记录
	Offset *int32 `json:"offset,omitempty"`
}

func (o CinderListVolumeTransfersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderListVolumeTransfersRequest struct{}"
	}

	return strings.Join([]string{"CinderListVolumeTransfersRequest", string(data)}, " ")
}
