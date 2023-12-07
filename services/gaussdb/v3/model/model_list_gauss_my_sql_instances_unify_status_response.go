package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGaussMySqlInstancesUnifyStatusResponse Response Object
type ListGaussMySqlInstancesUnifyStatusResponse struct {

	// 实例列表信息。
	Instances *[]MysqlInstanceListInfoUnifyStatus `json:"instances,omitempty"`

	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListGaussMySqlInstancesUnifyStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGaussMySqlInstancesUnifyStatusResponse struct{}"
	}

	return strings.Join([]string{"ListGaussMySqlInstancesUnifyStatusResponse", string(data)}, " ")
}
