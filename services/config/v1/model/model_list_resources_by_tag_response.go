package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourcesByTagResponse Response Object
type ListResourcesByTagResponse struct {
	Resources *[]ResourceResp `json:"resources,omitempty"`

	// 总记录数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListResourcesByTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesByTagResponse struct{}"
	}

	return strings.Join([]string{"ListResourcesByTagResponse", string(data)}, " ")
}
