package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPlayDomainStreamInfoResponse Response Object
type ListPlayDomainStreamInfoResponse struct {

	// 采样时间。日期格式按照ISO8601表示法，并使用UTC时间。 格式为：YYYY-MM-DDThh:mm:ssZ。
	Time *string `json:"time,omitempty"`

	// 采样数据列表
	DataList *[]PlayDomainStreamInfo `json:"data_list,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPlayDomainStreamInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlayDomainStreamInfoResponse struct{}"
	}

	return strings.Join([]string{"ListPlayDomainStreamInfoResponse", string(data)}, " ")
}
