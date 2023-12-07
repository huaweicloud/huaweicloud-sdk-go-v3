package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGaussMySqlInstanceDetailInfoUnifyStatusResponse Response Object
type ListGaussMySqlInstanceDetailInfoUnifyStatusResponse struct {

	// 实例详情。
	Instances      *[]MysqlInstanceInfoDetailUnifyStatus `json:"instances,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o ListGaussMySqlInstanceDetailInfoUnifyStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGaussMySqlInstanceDetailInfoUnifyStatusResponse struct{}"
	}

	return strings.Join([]string{"ListGaussMySqlInstanceDetailInfoUnifyStatusResponse", string(data)}, " ")
}
