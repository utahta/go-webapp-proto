package config

import "github.com/spf13/viper"

type Config struct {
	Database database
}

type database struct {
	Driver   string
	DB       string
	Host     string
	Port     int
	User     string
	Password string
}

var C *Config

func init() {
	C = new(Config)
}

// 設定ファイルを読み出し
func Load(env string, path string) error {
	viper.AddConfigPath(path)

	// 基本の設定を読み出し
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// 環境毎の設定を読み出し
	viper.SetConfigName(env)
	if err := viper.MergeInConfig(); err != nil {
		return err
	}

	// struct に変数がないときエラーを返すようにしとく
	v := viper.GetViper()
	if err := v.UnmarshalExact(C); err != nil {
		return err
	}
	return nil
}
