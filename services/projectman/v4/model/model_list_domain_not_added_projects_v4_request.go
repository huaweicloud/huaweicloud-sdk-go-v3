package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListDomainNotAddedProjectsV4Request struct {

	// 分页索引，偏移量,offset是limit的整数倍，limit=10,offset=0,10,20...
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页显示的数量,每页最多显示100条
	Limit *int32 `json:"limit,omitempty" xml:"limit"`
}

func (o ListDomainNotAddedProjectsV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainNotAddedProjectsV4Request struct{}"
	}

	return strings.Join([]string{"ListDomainNotAddedProjectsV4Request", string(data)}, " ")
}
