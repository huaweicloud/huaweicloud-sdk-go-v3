package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 分页查询对象Query
type SlaveRegister struct {

	// cluster ID
	ClusterId string `json:"cluster_id" xml:"cluster_id"`

	// Slave名称
	SlaveName string `json:"slave_name" xml:"slave_name"`

	// Slave工作空间
	WorkDir string `json:"work_dir" xml:"work_dir"`

	// Slave label
	Label *string `json:"label,omitempty" xml:"label"`

	// agent版本
	Version *string `json:"version,omitempty" xml:"version"`

	// 是否重试
	Retry *bool `json:"retry,omitempty" xml:"retry"`

	// Slave ownerType
	OwnerType *string `json:"owner_type,omitempty" xml:"owner_type"`
}

func (o SlaveRegister) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlaveRegister struct{}"
	}

	return strings.Join([]string{"SlaveRegister", string(data)}, " ")
}
