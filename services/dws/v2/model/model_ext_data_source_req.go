package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtDataSourceReq **参数解释**： 数据源请求。 **取值范围**： 不涉及。
type ExtDataSourceReq struct {

	// **参数解释**： 外部数据源ID。当数据源为MRS时为必选字段。 **取值范围**： 不涉及。
	DataSourceId *string `json:"data_source_id,omitempty"`

	// **参数解释**： 数据源类型。 **取值范围**： - OBS: obs数据源。 - LAKE_FORMATION: lake_formation数据源。 - MRS: mrs数据源。
	Type string `json:"type"`

	// **参数解释**： 外部数据源名称。 **取值范围**： 仅可包含大小写字母、数字、下划线，1到64个字符。
	DataSourceName string `json:"data_source_name"`

	// **参数解释**： 用户名。数据源类型为OBS时，传对应OBS委托名称。 **取值范围**： 仅可包含大小写字母、数字、下划线，3到20个字符。
	UserName string `json:"user_name"`

	// **参数解释**： 密码。当数据源为MRS时为必选字段。 **取值范围**： 不涉及。
	UserPwd *string `json:"user_pwd,omitempty"`

	// **参数解释**： 描述。 **取值范围**： 除!<>'=&等特殊字符外的字符。长度256个字符以内。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 是否重启集群。 **取值范围**： 不涉及。
	Reboot *bool `json:"reboot,omitempty"`

	// **参数解释**： 连接的数据库。当数据源为OBS时为必选字段。 **取值范围**： 不涉及。
	ConnectInfo *string `json:"connect_info,omitempty"`
}

func (o ExtDataSourceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtDataSourceReq struct{}"
	}

	return strings.Join([]string{"ExtDataSourceReq", string(data)}, " ")
}
