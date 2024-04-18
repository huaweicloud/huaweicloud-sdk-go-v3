package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OpenGaussUpgradeRequest GaussDB(for openGauss)实例版本升级接口传参参数。
type OpenGaussUpgradeRequest struct {

	// 实例升级类型，包括就地升级，灰度升级，热补丁升级三种，三种升级方式的异同，详见接口描述。  inplace 就地升级  grey 灰度升级  hotfix 热补丁升级。
	UpgradeType string `json:"upgrade_type"`

	// 实例升级操作，就地升级无需传值。灰度升级包括升级自动提交，升级待观察，提交升级，升级回退四种。热补丁升级包括升级自动提交，升级回退两种。详见接口描述。  upgradeAutoCommit 升级自动提交  upgrade 升级待观察  commit 提交升级  rollback 升级回退。
	UpgradeAction *string `json:"upgrade_action,omitempty"`

	// 实例升级目标版本，非必填。如果未传值则默认升级到当前实例的优选版本。仅热补丁升级方式下支持传入多个值，升级操作为升级自动提交，则实例版本从小到大批量升级，升级操作为升级回退，则实例版本从大到小批量回退。
	TargetVersion *string `json:"target_version,omitempty"`

	// 分布式实例灰度升级，滚动升级分片数。分布式实例灰度升级，升级待观察必填。该值不能大于实例未升级分片总数。
	UpgradeShardNum *int32 `json:"upgrade_shard_num,omitempty"`

	// 主备版实例灰度升级，滚动升级az值，可以支持多az一起升级，az之间以’,’分割。集中式实例灰度升级，升级待观察必填。该值不能填入不属于该实例的az值。
	UpgradeAz *string `json:"upgrade_az,omitempty"`
}

func (o OpenGaussUpgradeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussUpgradeRequest struct{}"
	}

	return strings.Join([]string{"OpenGaussUpgradeRequest", string(data)}, " ")
}
