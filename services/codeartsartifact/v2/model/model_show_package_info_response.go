package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPackageInfoResponse Response Object
type ShowPackageInfoResponse struct {

	// 结果状态
	Status *string `json:"status,omitempty"`

	// 请求id
	TraceId *string `json:"trace_id,omitempty"`

	// 请求返回结果，接口不同，返回不同
	Result         *interface{} `json:"result,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowPackageInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPackageInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowPackageInfoResponse", string(data)}, " ")
}
