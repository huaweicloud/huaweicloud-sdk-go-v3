package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAreasResponse Response Object
type ListAreasResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 大区列表。
	Areas          []Area `json:"areas"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAreasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAreasResponse struct{}"
	}

	return strings.Join([]string{"ListAreasResponse", string(data)}, " ")
}
