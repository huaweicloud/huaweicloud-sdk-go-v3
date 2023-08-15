package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Extend extend对象
type Extend struct {

	// 扩容后文件系统的新容量，以GB为单位。扩容步长大于等于100GB。  普通文件系统容量，取值范围500~32768。  带宽型文件系统，容量范围是10240~327680
	NewSize int32 `json:"new_size"`

	// 扩缩带宽后文件系统的新带宽，以 GB 为单位。仅支持 hpc cache 型文件系统
	NewBandwidth *int64 `json:"new_bandwidth,omitempty"`

	BssParam *BssInfoExtend `json:"bss_param,omitempty"`
}

func (o Extend) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Extend struct{}"
	}

	return strings.Join([]string{"Extend", string(data)}, " ")
}
