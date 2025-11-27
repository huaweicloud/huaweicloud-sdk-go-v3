package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ConfigBody 配置参数body
type ConfigBody struct {

	// 配置类型 MIGRATE_EXCLUDE_DIR: 迁移时设定的不迁移目录 SYNC_EXCLUDE_DIR: 同步时设定的不同步目录 ONLY_SYNC_DIR: 同步时设定的同步目录 CONSISTENCY_DIR: 一致性校验的目录 CONSISTENCY_DIR_ILLEGAL: 一致性校验后非法目录 LINUX_BLOCK_COMPRESS_THREAD_NUM: linux块迁移压缩线程个数 MIGRATE_DST_IP: 迁移目的ip LINUX_BLOCK_CACHE_SIZE: linux块迁移缓存大小 LINUX_CPU_LIMIT: linux的cpu限制 LINUX_MEM_LIMIT: linux的内存限制 LINUX_IO_LIMIT: linux的IO限制 NUM_PROCESS_MIGRATE: 迁移进程数 NUM_PROCESS_SYNC: 同步进程数 CONSISTENCY_RECHECK: 一致性校验再检 CONSISTENCY_MODE: 一致性校验模式 DYNAMIC_PORT: 动态端口
	ConfigKey ConfigBodyConfigKey `json:"config_key"`

	// 具体配置参数字段，保存于数据库，最终在agent端进行解析
	ConfigValue string `json:"config_value"`

	// 描述配置状态的保留字段
	ConfigStatus *string `json:"config_status,omitempty"`
}

func (o ConfigBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigBody struct{}"
	}

	return strings.Join([]string{"ConfigBody", string(data)}, " ")
}

type ConfigBodyConfigKey struct {
	value string
}

type ConfigBodyConfigKeyEnum struct {
	MIGRATE_EXCLUDE_DIR             ConfigBodyConfigKey
	SYNC_EXCLUDE_DIR                ConfigBodyConfigKey
	ONLY_SYNC_DIR                   ConfigBodyConfigKey
	CONSISTENCY_DIR                 ConfigBodyConfigKey
	CONSISTENCY_DIR_ILLEGAL         ConfigBodyConfigKey
	LINUX_BLOCK_COMPRESS_THREAD_NUM ConfigBodyConfigKey
	MIGRATE_DST_IP                  ConfigBodyConfigKey
	LINUX_BLOCK_CACHE_SIZE          ConfigBodyConfigKey
	LINUX_CPU_LIMIT                 ConfigBodyConfigKey
	LINUX_MEM_LIMIT                 ConfigBodyConfigKey
	LINUX_IO_LIMIT                  ConfigBodyConfigKey
	NUM_PROCESS_MIGRATE             ConfigBodyConfigKey
	NUM_PROCESS_SYNC                ConfigBodyConfigKey
	CONSISTENCY_RECHECK             ConfigBodyConfigKey
	CONSISTENCY_MODE                ConfigBodyConfigKey
	DYNAMIC_PORT                    ConfigBodyConfigKey
}

func GetConfigBodyConfigKeyEnum() ConfigBodyConfigKeyEnum {
	return ConfigBodyConfigKeyEnum{
		MIGRATE_EXCLUDE_DIR: ConfigBodyConfigKey{
			value: "MIGRATE_EXCLUDE_DIR",
		},
		SYNC_EXCLUDE_DIR: ConfigBodyConfigKey{
			value: "SYNC_EXCLUDE_DIR",
		},
		ONLY_SYNC_DIR: ConfigBodyConfigKey{
			value: "ONLY_SYNC_DIR",
		},
		CONSISTENCY_DIR: ConfigBodyConfigKey{
			value: "CONSISTENCY_DIR",
		},
		CONSISTENCY_DIR_ILLEGAL: ConfigBodyConfigKey{
			value: "CONSISTENCY_DIR_ILLEGAL",
		},
		LINUX_BLOCK_COMPRESS_THREAD_NUM: ConfigBodyConfigKey{
			value: "LINUX_BLOCK_COMPRESS_THREAD_NUM",
		},
		MIGRATE_DST_IP: ConfigBodyConfigKey{
			value: "MIGRATE_DST_IP",
		},
		LINUX_BLOCK_CACHE_SIZE: ConfigBodyConfigKey{
			value: "LINUX_BLOCK_CACHE_SIZE",
		},
		LINUX_CPU_LIMIT: ConfigBodyConfigKey{
			value: "LINUX_CPU_LIMIT",
		},
		LINUX_MEM_LIMIT: ConfigBodyConfigKey{
			value: "LINUX_MEM_LIMIT",
		},
		LINUX_IO_LIMIT: ConfigBodyConfigKey{
			value: "LINUX_IO_LIMIT",
		},
		NUM_PROCESS_MIGRATE: ConfigBodyConfigKey{
			value: "NUM_PROCESS_MIGRATE",
		},
		NUM_PROCESS_SYNC: ConfigBodyConfigKey{
			value: "NUM_PROCESS_SYNC",
		},
		CONSISTENCY_RECHECK: ConfigBodyConfigKey{
			value: "CONSISTENCY_RECHECK",
		},
		CONSISTENCY_MODE: ConfigBodyConfigKey{
			value: "CONSISTENCY_MODE",
		},
		DYNAMIC_PORT: ConfigBodyConfigKey{
			value: "DYNAMIC_PORT",
		},
	}
}

func (c ConfigBodyConfigKey) Value() string {
	return c.value
}

func (c ConfigBodyConfigKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfigBodyConfigKey) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
