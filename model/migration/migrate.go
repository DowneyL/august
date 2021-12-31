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
	config.Conf().Set("out_dir", option.Po.To)
	config.Conf().Set("table_prefix", option.TablePrefix)
	config.Conf().Set("table_names", option.TableNames)
	config.Conf().Set("db_info.host", option.Do.Host)
	config.Conf().Set("db_info.port", option.Do.Port)
	config.Conf().Set("db_info.username", option.Do.Username)
	config.Conf().Set("db_info.password", option.Do.Password)
	config.Conf().Set("db_info.database", option.Do.DBName)

	return config.Conf().WriteConfigAs(config.Filepath())
}
