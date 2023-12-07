package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGaussMySqlInstanceInfoUnifyStatusResponse Response Object
type ShowGaussMySqlInstanceInfoUnifyStatusResponse struct {
	Instance       *MysqlInstanceInfoDetailUnifyStatus `json:"instance,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ShowGaussMySqlInstanceInfoUnifyStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGaussMySqlInstanceInfoUnifyStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowGaussMySqlInstanceInfoUnifyStatusResponse", string(data)}, " ")
}
