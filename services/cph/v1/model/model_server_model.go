package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerModel 服务器规格。
type ServerModel struct {

	// 云手机服务器的规格名称，不超过64字节。
	ServerModelName *string `json:"server_model_name,omitempty"`

	// 云手机服务器的型号，如Hi1616。不超过32字节。
	ServerType *string `json:"server_type,omitempty"`

	// 云手机服务器的CPU核数。
	Cpu *int32 `json:"cpu,omitempty"`

	// 云手机服务器的内存大小，单位G。
	Memory *int32 `json:"memory,omitempty"`

	ExtendSpec *ServerModelExtendSpec `json:"extend_spec,omitempty"`

	// 产品类型。 - 0：云手机 - 1：云手游
	ProductType *int32 `json:"product_type,omitempty"`

	// 服务器磁盘的免费配额，单位G。
	FreeSize *int32 `json:"free_size,omitempty"`
}

func (o ServerModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerModel struct{}"
	}

	return strings.Join([]string{"ServerModel", string(data)}, " ")
}
