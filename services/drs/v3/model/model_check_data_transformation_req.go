package model

import (
	"encoding/json"

	"strings"
)

// 校验数据加工规则请求体
type CheckDataTransformationReq struct {
	// 任务id

	JobId *string `json:"job_id,omitempty"`
	// 对象信息

	ObjectInfo []DatabaseObjectVo `json:"object_info"`

	TransformationInfo *TransformationInfo `json:"transformation_info"`

	ConfigTransformation *ConfigTransformationVo `json:"config_transformation,omitempty"`
}

func (o CheckDataTransformationReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CheckDataTransformationReq struct{}"
	}

	return strings.Join([]string{"CheckDataTransformationReq", string(data)}, " ")
}
