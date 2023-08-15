package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyPortRequestBody struct {

	// 新端口号。端口有效范围为2100~9500，暂不支持8636、8637和8638。GaussDB(for Mongo)副本集4.0数据库实例端口有效范围为2100~9500，暂不支持8636、8637和8638。 GaussDB(for Cassandra)数据库实例端口有效范围为2100~9500，暂不支持7000，7001，7199，8636，8479，8484，8999，8018，2180，2887，3887，8079，8091，8092。 GaussDB(for Redis)数据库实例端口有效范围为1024~65535，暂不支持2180、2887、3887、6377、6378、6380、8018、8079、8091、8479、8484、8999、12017、12333、50069。
	Port int32 `json:"port"`
}

func (o ModifyPortRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyPortRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyPortRequestBody", string(data)}, " ")
}
