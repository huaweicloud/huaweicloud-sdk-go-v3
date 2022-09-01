package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowInformationAboutDatabaseProxyResponse struct {
	Proxy *Proxy `json:"proxy,omitempty" xml:"proxy"`

	MasterInstance *MasterInstance `json:"master_instance,omitempty" xml:"master_instance"`

	// 只读实例信息。
	ReadonlyInstances *[]ReadonlyInstances `json:"readonly_instances,omitempty" xml:"readonly_instances"`
	HttpStatusCode    int                  `json:"-"`
}

func (o ShowInformationAboutDatabaseProxyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInformationAboutDatabaseProxyResponse struct{}"
	}

	return strings.Join([]string{"ShowInformationAboutDatabaseProxyResponse", string(data)}, " ")
}
