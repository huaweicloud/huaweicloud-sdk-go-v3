package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBaremetalServerMetadataReq This is a auto create Body Object
type UpdateBaremetalServerMetadataReq struct {

	// 用户自定义metadata键值对。  结构体允许为空，取值为空时不更新数据。
	Metadata map[string]string `json:"metadata"`
}

func (o UpdateBaremetalServerMetadataReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBaremetalServerMetadataReq struct{}"
	}

	return strings.Join([]string{"UpdateBaremetalServerMetadataReq", string(data)}, " ")
}
