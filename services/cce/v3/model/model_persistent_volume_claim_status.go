package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type PersistentVolumeClaimStatus struct {

	// 显示volume实际具有的访问模式。
	AccessModes *[]string `json:"accessModes,omitempty" xml:"accessModes"`

	// 底层卷的实际资源
	Capacity *string `json:"capacity,omitempty" xml:"capacity"`

	// PersistentVolumeClaim当前所处的状态
	Phase *string `json:"phase,omitempty" xml:"phase"`
}

func (o PersistentVolumeClaimStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistentVolumeClaimStatus struct{}"
	}

	return strings.Join([]string{"PersistentVolumeClaimStatus", string(data)}, " ")
}
