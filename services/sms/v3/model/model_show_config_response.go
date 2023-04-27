package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowConfigResponse struct {

	// mainRegion,obs_domain,disktype,process_and_it及以后增加的信息
	Config map[string]string `json:"config,omitempty"`

	// region数组
	Regions        *[]map[string]string `json:"regions,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowConfigResponse", string(data)}, " ")
}
