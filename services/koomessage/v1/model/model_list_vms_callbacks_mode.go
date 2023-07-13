package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVmsCallbacksMode 响应数据。
type ListVmsCallbacksMode struct {

	// 回执接口列表。
	Callbacks *[]VmsCallBack `json:"callbacks,omitempty"`
}

func (o ListVmsCallbacksMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVmsCallbacksMode struct{}"
	}

	return strings.Join([]string{"ListVmsCallbacksMode", string(data)}, " ")
}
