package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListDownloadAddressRequestBody struct {

	// 录屏记录UUID列表。
	RecordIds *[]string `json:"record_ids,omitempty"`
}

func (o ListDownloadAddressRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDownloadAddressRequestBody struct{}"
	}

	return strings.Join([]string{"ListDownloadAddressRequestBody", string(data)}, " ")
}
