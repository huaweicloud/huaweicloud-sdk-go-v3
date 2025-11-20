package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceByInstanceIdResponse Response Object
type ShowInstanceByInstanceIdResponse struct {

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 版本
	Version *string `json:"version,omitempty"`

	// 过期时间
	ExpireTime *int32 `json:"expire_time,omitempty"`

	// 创建时间
	CreateTime *int32 `json:"create_time,omitempty"`

	// 当前时间
	CurrentTime *int32 `json:"current_time,omitempty"`

	ProductSpecData *ProductSpecData `json:"product_spec_data,omitempty"`

	InstanceConfig *InstanceConfig `json:"instance_config,omitempty"`

	// 弹性业务带宽是否可更新
	ElasticServiceBwUpdateEnable *bool `json:"elastic_service_bw_update_enable,omitempty"`
	HttpStatusCode               int   `json:"-"`
}

func (o ShowInstanceByInstanceIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceByInstanceIdResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceByInstanceIdResponse", string(data)}, " ")
}
