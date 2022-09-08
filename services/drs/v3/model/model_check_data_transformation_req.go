package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 校验数据加工规则请求体
type CheckDataTransformationReq struct {

	// 任务id
	JobId *string `json:"job_id,omitempty"`

	// 对象信息，生成加工规则时需要填写。
	ObjectInfo *[]DatabaseObjectVo `json:"object_info,omitempty"`

	TransformationInfo *TransformationInfo `json:"transformation_info"`

	ConfigTransformation *ConfigTransformationVo `json:"config_transformation,omitempty"`
}

func (o CheckDataTransformationReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDataTransformationReq struct{}"
	}

	return strings.Join([]string{"CheckDataTransformationReq", string(data)}, " ")
}
