package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DocumentQueryResponseRequestParamsData 创建作业时传的data参数
type DocumentQueryResponseRequestParamsData struct {

	// 创建作业时传的url参数
	Url string `json:"url"`

	// 创建作业时传的format参数
	Format string `json:"format"`

	// 创建作业时传的frame_interval参数
	FrameInterval *int32 `json:"frame_interval,omitempty"`
}

func (o DocumentQueryResponseRequestParamsData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DocumentQueryResponseRequestParamsData struct{}"
	}

	return strings.Join([]string{"DocumentQueryResponseRequestParamsData", string(data)}, " ")
}
