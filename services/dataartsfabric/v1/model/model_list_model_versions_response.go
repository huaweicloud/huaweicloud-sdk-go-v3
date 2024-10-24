package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListModelVersionsResponse Response Object
type ListModelVersionsResponse struct {

	// 符合条件的Version总数
	Total *int32 `json:"total,omitempty"`

	// 列表信息
	Versions *[]ModelVersionInfo `json:"versions,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListModelVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListModelVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListModelVersionsResponse", string(data)}, " ")
}
