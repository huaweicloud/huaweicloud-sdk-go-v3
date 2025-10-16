package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTranscodeConcurrencyNumResponse Response Object
type ListTranscodeConcurrencyNumResponse struct {

	// 采样数据列表
	DataList *[]TranscodeConNumData `json:"data_list,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTranscodeConcurrencyNumResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTranscodeConcurrencyNumResponse struct{}"
	}

	return strings.Join([]string{"ListTranscodeConcurrencyNumResponse", string(data)}, " ")
}
