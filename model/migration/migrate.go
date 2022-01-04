package migration

import (
	"fmt"
	"github.com/DowneyL/august/model/migration/config"
)

func Migration(option Option) {
	config.Init(option.Po.GormtPath)
	if err := resetConfig(option); err != nil {
		panic(err)
	}

	if err := poMigrate(option); err != nil {
		fmt.Println(err)
		return
	}

	if err := repoMigrate(option); err != nil {
		fmt.Println(err)
	}
}

func resetConfig(option Option) error {
	config.SetConf("out_dir", option.Po.To)
	config.SetConf("table_prefix", option.TablePrefix)
	config.SetConf("table_names", option.TableNames)
	config.SetConf("db_info.host", option.Do.Host)
	config.SetConf("db_info.port", option.Do.Port)
	config.SetConf("db_info.username", option.Do.Username)
	config.SetConf("db_info.password", option.Do.Password)
	config.SetConf("db_info.database", option.Do.DBName)

	return config.Conf().WriteConfigAs(config.Filepath())
}
