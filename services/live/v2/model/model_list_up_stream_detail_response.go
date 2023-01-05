package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListUpStreamDetailResponse struct {

	// 推流域名
	PublishDomain *string `json:"publish_domain,omitempty"`

	// 应用名
	App *string `json:"app,omitempty"`

	// 流名
	Stream *string `json:"stream,omitempty"`

	// 推流质量数据
	Data *[]UpStreamDetail `json:"data,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListUpStreamDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUpStreamDetailResponse struct{}"
	}

	return strings.Join([]string{"ListUpStreamDetailResponse", string(data)}, " ")
}
