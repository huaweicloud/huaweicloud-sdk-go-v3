package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRemoteDbResponse Response Object
type ListRemoteDbResponse struct {

	// 数据库信息。
	Databases *[]SqlserverDatabaseForDetail `json:"databases,omitempty"`

	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListRemoteDbResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRemoteDbResponse struct{}"
	}

	return strings.Join([]string{"ListRemoteDbResponse", string(data)}, " ")
}
