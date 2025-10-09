package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomerUpgradeDatabaseVersionReqNew struct {

	// 是否延迟至可维护时间段内升级。 取值范围： - true：延迟升级。表示实例将在设置的可维护时间段内升级。 - false：立即升级，默认该方式。
	IsDelayed *bool `json:"is_delayed,omitempty"`

	// 设置仅对RDS for MySQL数据库实例（主备）生效。主备实例升级过程中，备机升级成功后，会触发主备倒换继续升级主机，主机升级完成后，若主备可用区不同则触发第二次倒换，保证可用区不变。若主备可用区相同，该选项无效。 取值范围： - true：默认该方式。表示升级过程中会进行二次倒换保证主备实例可用区不变。 - false：升级过程中不进行第二次主备倒换，适合对主备所在可用区不敏感，对业务连续性敏感的客户。
	SecondSwitch *bool `json:"second_switch,omitempty"`
}

func (o CustomerUpgradeDatabaseVersionReqNew) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomerUpgradeDatabaseVersionReqNew struct{}"
	}

	return strings.Join([]string{"CustomerUpgradeDatabaseVersionReqNew", string(data)}, " ")
}
