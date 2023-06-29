package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecretVersionsRequest Request Object
type ListSecretVersionsRequest struct {

	// 凭据名称。
	SecretName string `json:"secret_name"`

	// 分页参数，取值为上一页数据的最后一条记录的版本号。
	Marker *string `json:"marker,omitempty"`

	// 每页显示的条目数量。默认值50。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSecretVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecretVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListSecretVersionsRequest", string(data)}, " ")
}
