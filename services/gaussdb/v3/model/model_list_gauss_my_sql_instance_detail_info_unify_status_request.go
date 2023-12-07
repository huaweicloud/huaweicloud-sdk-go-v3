package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGaussMySqlInstanceDetailInfoUnifyStatusRequest Request Object
type ListGaussMySqlInstanceDetailInfoUnifyStatusRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID，严格匹配UUID规则。最多同时输入20个实例ID，用英文逗号分隔。
	InstanceIds string `json:"instance_ids"`
}

func (o ListGaussMySqlInstanceDetailInfoUnifyStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGaussMySqlInstanceDetailInfoUnifyStatusRequest struct{}"
	}

	return strings.Join([]string{"ListGaussMySqlInstanceDetailInfoUnifyStatusRequest", string(data)}, " ")
}
