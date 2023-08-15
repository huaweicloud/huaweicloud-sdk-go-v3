package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EncryptData 加密数据信息
type EncryptData struct {

	// 加密数据ID
	Id string `json:"id"`

	// 加密数据名称
	Name string `json:"name"`

	// 加密数据描述
	Description string `json:"description"`

	// 加密数据项配置
	Config []EncryptDataItem `json:"config"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 铂金版实例ID，专业版实例为default
	IefInstanceId *string `json:"ief_instance_id,omitempty"`

	// 租户账户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 加密数据创建时间
	CreatedTime *int64 `json:"created_time,omitempty"`

	// 加密数据更新时间
	UpdatedTime *int64 `json:"updated_time,omitempty"`
}

func (o EncryptData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncryptData struct{}"
	}

	return strings.Join([]string{"EncryptData", string(data)}, " ")
}
