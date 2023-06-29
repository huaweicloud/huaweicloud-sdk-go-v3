package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEnvResponse Response Object
type ShowEnvResponse struct {

	// notebook是否使用专属资源池
	DevUserPool *bool `json:"dev_user_pool,omitempty"`

	// 是否集成开发环境
	HasDev *bool `json:"has_dev,omitempty"`

	// 是否部署药物虚拟筛选
	HasDrug *bool `json:"has_drug,omitempty"`

	// 是否显示加密按钮
	HasEncryptionButton *bool `json:"has_encryption_button,omitempty"`

	// 医疗智能体部署模式
	DeployMode *string `json:"deploy_mode,omitempty"`

	// 是否支持归档类型存储
	EnableColdArchive *bool `json:"enable_cold_archive,omitempty"`
	HttpStatusCode    int   `json:"-"`
}

func (o ShowEnvResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnvResponse struct{}"
	}

	return strings.Join([]string{"ShowEnvResponse", string(data)}, " ")
}
