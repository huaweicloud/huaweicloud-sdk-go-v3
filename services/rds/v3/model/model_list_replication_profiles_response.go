package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListReplicationProfilesResponse Response Object
type ListReplicationProfilesResponse struct {

	// 配置文件列表。
	Profiles *[]ListReplicationProfilesResult `json:"profiles,omitempty"`

	// 配置文件总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListReplicationProfilesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReplicationProfilesResponse struct{}"
	}

	return strings.Join([]string{"ListReplicationProfilesResponse", string(data)}, " ")
}
