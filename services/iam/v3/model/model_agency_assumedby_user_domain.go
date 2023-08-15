package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgencyAssumedbyUserDomain
type AgencyAssumedbyUserDomain struct {

	// 被委托方B的账号名称。
	Name string `json:"name"`

	// 被委托方B的账号ID。
	Id string `json:"id"`
}

func (o AgencyAssumedbyUserDomain) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyAssumedbyUserDomain struct{}"
	}

	return strings.Join([]string{"AgencyAssumedbyUserDomain", string(data)}, " ")
}
