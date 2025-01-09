package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDownloadAddressResponse Response Object
type ListDownloadAddressResponse struct {

	// 下载地址列表。
	Address        *[]DownloadAddressForList `json:"address,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListDownloadAddressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDownloadAddressResponse struct{}"
	}

	return strings.Join([]string{"ListDownloadAddressResponse", string(data)}, " ")
}
