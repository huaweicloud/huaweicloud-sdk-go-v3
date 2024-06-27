package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBasicAwInfoListSupportsSearchResponse Response Object
type ListBasicAwInfoListSupportsSearchResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	Error *CommonResponseErrorPageResultBasicAwInfo `json:"error,omitempty"`

	// 失败原因
	Reason *string `json:"reason,omitempty"`

	Result *PageResultBasicAwInfo `json:"result,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListBasicAwInfoListSupportsSearchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBasicAwInfoListSupportsSearchResponse struct{}"
	}

	return strings.Join([]string{"ListBasicAwInfoListSupportsSearchResponse", string(data)}, " ")
}
