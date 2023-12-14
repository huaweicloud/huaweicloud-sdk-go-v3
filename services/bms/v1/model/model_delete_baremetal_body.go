package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBaremetalBody This is a auto create Body Object
type DeleteBaremetalBody struct {

	// 所需要删除的裸金属服务器列表。
	Servers []ServersList `json:"servers"`

	// 删除裸金属服务器时是否删除裸金属服务器绑定的弹性公网IP。如果选择不删除，则系统仅做解绑定操作，保留弹性公网IP资源。 取值为true或false，默认为false。  true：删除裸金属服务器时会同时删除绑定在裸金属服务器上的弹性公网IP。 false：删除裸金属服务器时，仅解绑定在裸金属服务器上的弹性公网IP，不删除弹性公网IP。
	DeletePublicip bool `json:"delete_publicip"`

	// 删除裸金属服务器时是否删除裸金属服务器对应的数据盘。如果选择不删除，则系统仅做解绑定操作，保留数据盘资源。 取值为false或true，默认为false。  true：删除裸金属服务器时会同时删除挂载在裸金属服务器上的数据盘。 false：删除裸金属服务器时，仅卸载裸金属服务器上挂载的数据盘，不删除该数据盘。
	DeleteVolume *bool `json:"delete_volume,omitempty"`
}

func (o DeleteBaremetalBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBaremetalBody struct{}"
	}

	return strings.Join([]string{"DeleteBaremetalBody", string(data)}, " ")
}
