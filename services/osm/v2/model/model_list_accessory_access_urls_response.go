package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccessoryAccessUrlsResponse Response Object
type ListAccessoryAccessUrlsResponse struct {

	// 附件链接
	AccessoryUrls  *[]AccessoryUrl `json:"accessory_urls,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListAccessoryAccessUrlsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessoryAccessUrlsResponse struct{}"
	}

	return strings.Join([]string{"ListAccessoryAccessUrlsResponse", string(data)}, " ")
}
