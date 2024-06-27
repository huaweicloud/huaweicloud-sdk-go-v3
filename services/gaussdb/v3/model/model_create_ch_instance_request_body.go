package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateChInstanceRequestBody 创建ClickHouse实例请求体。
type CreateChInstanceRequestBody struct {

	// ClickHouse实例名称。 - 4位到64位之间 - 必须以字母开头，可以包含字母、数字、中划线或下划线 - 不能包含其他特殊字符
	Name string `json:"name"`

	Engine *ClickHouseEngineInfo `json:"engine"`

	Ha *CreateChInstanceRequestBodyHa `json:"ha"`

	// 节点规格ID，可通过“HTAP查询规格信息”获取。
	FlavorId string `json:"flavor_id"`

	// root账户密码。 - 8~32个字符 - 包含大写字母、小写字母、数字或特殊字符(~ ! @ # % ^ * - _ = + ? ,)中的三种
	DbRootPwd string `json:"db_root_pwd"`

	// 可用区类型。 取值范围： - single：单可用区 - double：多可用区
	AzMode string `json:"az_mode"`

	Volume *CreateChInstanceRequestBodyVolume `json:"volume"`

	// 可用区码。  当ha中mode为Ha时，需要填写多个可用区，用\",\" 分隔。例如：cn-southwest-244b,cn-southwest-244a
	AzCode string `json:"az_code"`

	// 时区。默认为所属GaussDB(for MySQL)实例时区。
	TimeZone *string `json:"time_zone,omitempty"`

	TagsInfo *CreateChInstanceRequestBodyTagsInfo `json:"tags_info"`

	PayInfo *CreateChInstanceRequestBodyPayInfo `json:"pay_info,omitempty"`
}

func (o CreateChInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateChInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateChInstanceRequestBody", string(data)}, " ")
}
