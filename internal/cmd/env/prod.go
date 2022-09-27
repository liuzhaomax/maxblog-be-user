//go:build prod

package env

import "flag"

func LoadEnv() *string {
	config := flag.String("c", "env/prod.yaml", "配置文件")
	flag.Parse()
	return config
}
