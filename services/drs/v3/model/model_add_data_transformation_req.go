package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDataTransformationReq 数据加工规则
type AddDataTransformationReq struct {

	// 对象信息。
	ObjectInfo []CreateDataCompareDatabaseObject `json:"object_info"`

	TransformationInfo *CreateDataCompareTransformationInfo `json:"transformation_info"`
}

func (o AddDataTransformationReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDataTransformationReq struct{}"
	}

	return strings.Join([]string{"AddDataTransformationReq", string(data)}, " ")
}
