package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowInformationAboutDatabaseProxyResponse struct {
	Proxy          *Proxy          `json:"proxy,omitempty"`
	MasterInstance *MasterInstance `json:"master_instance,omitempty"`
	// 只读实例信息
	ReadonlyInstances *[]ReadonlyInstances `json:"readonly_instances,omitempty"`
	HttpStatusCode    int                  `json:"-"`
}

func (o ShowInformationAboutDatabaseProxyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowInformationAboutDatabaseProxyResponse struct{}"
	}

	return strings.Join([]string{"ShowInformationAboutDatabaseProxyResponse", string(data)}, " ")
}
