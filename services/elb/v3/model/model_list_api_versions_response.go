package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApiVersionsResponse Response Object
type ListApiVersionsResponse struct {

	// 可用API版本列表。
	Versions       *[]ApiVersionInfo `json:"versions,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListApiVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListApiVersionsResponse", string(data)}, " ")
}
