package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateMySqlCompatibilityRequestBody struct {

	// M兼容端口，可选范围为：0, 1024-39989。   - 如下端口不可设置： 2378，2379，2380，2400，4999，5000，5001，5100，5500，5999，6000，6001，6009，6010，6500，8015，8097，8098，8181，9090，9100，9180，9187，9200，12016，12017，20049，20050，21731，21732，32122，32123，32124，32125，32126，39001，[数据库端口, 数据库端口+10]。   - 取值为0，表示关闭M兼容端口。
	Port string `json:"port"`
}

func (o UpdateMySqlCompatibilityRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMySqlCompatibilityRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateMySqlCompatibilityRequestBody", string(data)}, " ")
}
