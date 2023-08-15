package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObsContentReq OBS数据源配置内容
type ObsContentReq struct {

	// 桶名称
	BucketName string `json:"bucket_name"`

	// 租户的AK
	Ak string `json:"ak"`

	// 租户的SK
	Sk string `json:"sk"`
}

func (o ObsContentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsContentReq struct{}"
	}

	return strings.Join([]string{"ObsContentReq", string(data)}, " ")
}
