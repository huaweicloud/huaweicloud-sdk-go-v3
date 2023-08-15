package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClusterSettingResponse Response Object
type UpdateClusterSettingResponse struct {

	// 配置修改结果
	ModifyResult   *bool `json:"modify_result,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o UpdateClusterSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterSettingResponse struct{}"
	}

	return strings.Join([]string{"UpdateClusterSettingResponse", string(data)}, " ")
}
