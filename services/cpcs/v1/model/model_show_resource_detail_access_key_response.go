package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceDetailAccessKeyResponse Response Object
type ShowResourceDetailAccessKeyResponse struct {

	// 资源名称
	MetricName *string `json:"metric_name,omitempty"`

	// 访问密钥列表
	AkList *[]ShowResourceDetailAccessKeyResponseBodyAkList `json:"ak_list,omitempty"`

	// 访问密钥总数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowResourceDetailAccessKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceDetailAccessKeyResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceDetailAccessKeyResponse", string(data)}, " ")
}
