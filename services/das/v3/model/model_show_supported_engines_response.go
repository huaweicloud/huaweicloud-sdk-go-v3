package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSupportedEnginesResponse Response Object
type ShowSupportedEnginesResponse struct {

	// 所有EngineType
	AllEngineTypes *[]string `json:"all_engine_types,omitempty"`

	// 支持的EngineType
	SupportedEngineTypes *[]string `json:"supported_engine_types,omitempty"`

	// 支持的NetWorkType和EngineType
	SupportedNetWorkTypes *[]SupportNetWorkTypeResponse `json:"supported_net_work_types,omitempty"`

	// 支持的CloudDBA的NetWorkType和EngineType
	SupportedCloudDbaTypes *[]SupportNetWorkTypeResponse `json:"supported_cloud_dba_types,omitempty"`
	HttpStatusCode         int                           `json:"-"`
}

func (o ShowSupportedEnginesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSupportedEnginesResponse struct{}"
	}

	return strings.Join([]string{"ShowSupportedEnginesResponse", string(data)}, " ")
}
