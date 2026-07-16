package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InferDeploymentVersionItemResp 部署版本概要信息
type InferDeploymentVersionItemResp struct {

	// **参数解释：** 部署id。 **取值范围：** 不涉及。
	InferName *string `json:"infer_name,omitempty"`

	// **参数解释：** 部署版本 **取值范围：** 不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释：** 版本状态 **取值范围：** 不涉及。
	VersionStatus *string `json:"version_status,omitempty"`

	// **参数解释：** 版本描述 **取值范围：** 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 创建时间 **取值范围：** 不涉及。
	CreateAt *sdktime.SdkTime `json:"create_at,omitempty"`

	// **参数解释：** 更新时间 **取值范围：** 不涉及。
	UpdateAt *sdktime.SdkTime `json:"update_at,omitempty"`

	// **参数解释：** 部署类型。 **取值范围：** - SINGLE：单机单卡。 - MULTI：多机多卡。
	DeployType *string `json:"deploy_type,omitempty"`
}

func (o InferDeploymentVersionItemResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InferDeploymentVersionItemResp struct{}"
	}

	return strings.Join([]string{"InferDeploymentVersionItemResp", string(data)}, " ")
}
