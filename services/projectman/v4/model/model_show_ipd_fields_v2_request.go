package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIpdFieldsV2Request Request Object
type ShowIpdFieldsV2Request struct {

	// 项目32位ID，项目唯一标识。通过查询IPD项目列表获取，响应消息体中的id字段的值就是项目ID。
	ProjectId string `json:"project_id"`

	// 工作项类型ID，工作项类型的唯一标识。 不同项目模型下可选值不同： - IPD-系统设备类：10021（RR）、10065（SF）、10020（IR）、10022（SR）、10029（AR）、10027（Task）、10033（Bug） - IPD-独立软件类：10021（RR）、10065（SF）、10020（IR）、10023（US）、10027（Task）、10033（Bug） - IPD-自运营软件/云服务类：10001（Epic）、10028（FE）、10021（RR）、10023（US）、10027（Task）、10033（Bug）
	CategoryId string `json:"category_id"`

	// 层级字段ID。用于过滤层级类型的字段，当需要按层级结构筛选字段时传入。
	CategoryLayerId *string `json:"category_layer_id,omitempty"`

	// 目标项目32位ID，项目唯一标识。通过查询IPD项目列表获取，响应消息体中的id字段的值就是项目ID。当需要查询其他项目的字段配置时传入。
	TargetProjectId *string `json:"target_project_id,omitempty"`
}

func (o ShowIpdFieldsV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpdFieldsV2Request struct{}"
	}

	return strings.Join([]string{"ShowIpdFieldsV2Request", string(data)}, " ")
}
