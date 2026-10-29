package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVaultResourceInstancesRequest Request Object
type ShowVaultResourceInstancesRequest struct {

	// **参数解释：** 云备份企业项目ID，获取企业项目ID请参考[查询企业项目列表](https://support.huaweicloud.com/api-em/zh-cn_topic_0121230880.html)(tag:hws)[查询企业项目列表](https://support.huaweicloud.com/intl/zh-cn/api-em/zh-cn_topic_0121230880.html)(tag:hws_hk)[《企业管理API参考》“查询企业项目列表”章节](tag:ctc,dt,tm,sbc,hws_eu,fcs_vm,hcso_dt,srg) **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *VaultResourceInstancesReq `json:"body,omitempty"`
}

func (o ShowVaultResourceInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVaultResourceInstancesRequest struct{}"
	}

	return strings.Join([]string{"ShowVaultResourceInstancesRequest", string(data)}, " ")
}
