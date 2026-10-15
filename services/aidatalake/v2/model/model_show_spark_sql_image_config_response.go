package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSparkSqlImageConfigResponse **参数解释**：镜像配置信息，包含镜像ID和版本信息。
type ShowSparkSqlImageConfigResponse struct {

	// **参数解释**：SparkSql集群的版本信息(AIDalake.xxx(Sparkx.xx.x,JRExx.x))。 **取值范围**：长度为1~512个字符。
	ReleaseVersion *string `json:"release_version,omitempty"`

	// **参数解释**：SparkSql集群使用的镜像版本类型。 **取值范围**： - STABLE：稳定。 - RECOMMEND：推荐。 - BETA：测试。
	ImageVersionType *string `json:"image_version_type,omitempty"`
}

func (o ShowSparkSqlImageConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkSqlImageConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowSparkSqlImageConfigResponse", string(data)}, " ")
}
