package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSparkJobImageConfigResponse **参数解释**：镜像配置信息，包含镜像ID和版本信息。
type ShowSparkJobImageConfigResponse struct {

	// **参数解释**：镜像ID，唯一标识Spark作业使用的镜像。 **取值范围**：长度为1~128个字符。
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**：镜像版本，指定镜像的版本号。 **取值范围**：长度为1~64个字符。
	ImageVersionId *string `json:"image_version_id,omitempty"`

	// **参数解释**：Spark作业的版本信息(AIDalake.xxx(Sparkx.xx.x,JRExx.x))。 **取值范围**：长度为1~512个字符。
	ReleaseVersion *string `json:"release_version,omitempty"`

	// **参数解释**：Spark作业使用的镜像版本类型。 **取值范围**： - STABLE：稳定。 - RECOMMEND：推荐。 - BETA：测试。
	ImageVersionType *string `json:"image_version_type,omitempty"`
}

func (o ShowSparkJobImageConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkJobImageConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowSparkJobImageConfigResponse", string(data)}, " ")
}
