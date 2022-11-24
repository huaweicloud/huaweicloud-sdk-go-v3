package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListResoleRulesRequest struct {

	// 待查询的resolverrule的域名。
	DomainName *string `json:"domain_name,omitempty"`

	// 每页返回的资源个数。 取值范围：0~500，取值一般为10，20，50。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询起始页码，起始值为0。 当前设置marker不为空时，以marker为分页起始标识。取值范围：0~2147483647。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListResoleRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResoleRulesRequest struct{}"
	}

	return strings.Join([]string{"ListResoleRulesRequest", string(data)}, " ")
}
