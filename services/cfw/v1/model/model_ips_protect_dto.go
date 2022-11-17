package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// description
type IpsProtectDto struct {

	// 防护对象id
	ObjectId *string `json:"object_id,omitempty"`

	// ips防护模式，0：观察模式，1：严格模式，2：中等模式，3：宽松模式
	Mode *int32 `json:"mode,omitempty"`
}

func (o IpsProtectDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpsProtectDto struct{}"
	}

	return strings.Join([]string{"IpsProtectDto", string(data)}, " ")
}
