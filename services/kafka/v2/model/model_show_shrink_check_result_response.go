package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowShrinkCheckResultResponse Response Object
type ShowShrinkCheckResultResponse struct {

	// 缩容检查是否通过
	Success *bool `json:"success,omitempty"`

	// broker检查结果
	CheckDetail    *[]ShowShrinkCheckResponseBodyCheckDetail `json:"check_detail,omitempty"`
	HttpStatusCode int                                       `json:"-"`
}

func (o ShowShrinkCheckResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowShrinkCheckResultResponse struct{}"
	}

	return strings.Join([]string{"ShowShrinkCheckResultResponse", string(data)}, " ")
}
