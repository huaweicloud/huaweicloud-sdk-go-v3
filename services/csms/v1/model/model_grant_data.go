package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GrantData 授权返回数据
type GrantData struct {

	// 授权id，授权给个人时存在
	Uuid string `json:"uuid"`

	// 资源id
	ResourceId string `json:"resourceId"`

	// 授权类型（SECRET，GROUP）
	Type string `json:"type"`

	// 授权目标用户id
	GranteeUser string `json:"granteeUser"`

	// 创建时间
	CreateTime *int64 `json:"createTime,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"updateTime,omitempty"`

	// 有效期
	ValidityTime *int64 `json:"validityTime,omitempty"`

	// 状态
	State *int32 `json:"state,omitempty"`

	// 签名
	Signature *string `json:"signature,omitempty"`
}

func (o GrantData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GrantData struct{}"
	}

	return strings.Join([]string{"GrantData", string(data)}, " ")
}
