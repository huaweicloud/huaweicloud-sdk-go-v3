package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRecycleBinResponse Response Object
type ShowRecycleBinResponse struct {

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 回收站配置开关
	Switch *string `json:"switch,omitempty"`

	Policy         *RecycleBinPolicys `json:"policy,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowRecycleBinResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecycleBinResponse struct{}"
	}

	return strings.Join([]string{"ShowRecycleBinResponse", string(data)}, " ")
}
