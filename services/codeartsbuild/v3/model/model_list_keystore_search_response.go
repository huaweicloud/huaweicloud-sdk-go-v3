package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListKeystoreSearchResponse Response Object
type ListKeystoreSearchResponse struct {
	Result *ListKeystoreSearchResponseBodyResult `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListKeystoreSearchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeystoreSearchResponse struct{}"
	}

	return strings.Join([]string{"ListKeystoreSearchResponse", string(data)}, " ")
}
