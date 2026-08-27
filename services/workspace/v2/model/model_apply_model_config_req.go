package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyModelConfigReq 批量关联模型分组与资源请求。每项指定一个模型分组与一个资源（桌面/桌面标签）的关联。
type ApplyModelConfigReq struct {

	// 关联项列表，每项包含模型分组ID、资源ID及资源类型。
	Items []ModelConfigItem `json:"items"`
}

func (o ApplyModelConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyModelConfigReq struct{}"
	}

	return strings.Join([]string{"ApplyModelConfigReq", string(data)}, " ")
}
