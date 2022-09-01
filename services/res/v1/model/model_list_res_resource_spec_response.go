package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListResResourceSpecResponse struct {

	// 离线计算规格。
	OfflineSpec *[]string `json:"offline_spec,omitempty" xml:"offline_spec"`

	// 实时计算规格。
	NearlineSpec *[]string `json:"nearline_spec,omitempty" xml:"nearline_spec"`

	// 排序模型计算规格。
	DeepLearningSpec *[]string `json:"deep_learning_spec,omitempty" xml:"deep_learning_spec"`

	// 请求是否成功。
	IsSuccess *bool `json:"is_success,omitempty" xml:"is_success"`

	// 返回消息（请求成功时，不返回此字段）。
	Message *string `json:"message,omitempty" xml:"message"`

	// 错误码（请求成功时，不返回此字段）。
	ErrorCode      *string `json:"error_code,omitempty" xml:"error_code"`
	HttpStatusCode int     `json:"-"`
}

func (o ListResResourceSpecResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResResourceSpecResponse struct{}"
	}

	return strings.Join([]string{"ListResResourceSpecResponse", string(data)}, " ")
}
