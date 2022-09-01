package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type KeystoneGroupResult struct {

	// 用户组描述信息。
	Description string `json:"description" xml:"description"`

	// 用户组ID。
	Id string `json:"id" xml:"id"`

	// 用户组所属账号ID。
	DomainId string `json:"domain_id" xml:"domain_id"`

	// 用户组名称。
	Name string `json:"name" xml:"name"`

	Links *Links `json:"links" xml:"links"`

	// 用户组创建时间。
	CreateTime int64 `json:"create_time" xml:"create_time"`
}

func (o KeystoneGroupResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneGroupResult struct{}"
	}

	return strings.Join([]string{"KeystoneGroupResult", string(data)}, " ")
}
