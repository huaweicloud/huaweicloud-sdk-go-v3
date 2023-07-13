package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResResourceSpecResponse Response Object
type ListResResourceSpecResponse struct {

	// 离线计算规格。
	OfflineSpec *[]string `json:"offline_spec,omitempty"`

	// 实时计算规格。
	NearlineSpec *[]string `json:"nearline_spec,omitempty"`

	// 排序模型计算规格。
	DeepLearningSpec *[]string `json:"deep_learning_spec,omitempty"`

	// 请求是否成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 返回消息（请求成功时，不返回此字段）。
	Message *string `json:"message,omitempty"`

	// 错误码（请求成功时，不返回此字段）。
	ErrorCode      *string `json:"error_code,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListResResourceSpecResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResResourceSpecResponse struct{}"
	}

	return strings.Join([]string{"ListResResourceSpecResponse", string(data)}, " ")
}
