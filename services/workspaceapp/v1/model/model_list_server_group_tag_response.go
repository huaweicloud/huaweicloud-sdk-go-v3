package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServerGroupTagResponse Response Object
type ListServerGroupTagResponse struct {

	// 标签列表。
	Tags           *[]TmsTagValues `json:"tags,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListServerGroupTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerGroupTagResponse struct{}"
	}

	return strings.Join([]string{"ListServerGroupTagResponse", string(data)}, " ")
}
