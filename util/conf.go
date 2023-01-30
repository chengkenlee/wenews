package util

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func usage() {
	fmt.Printf("\nUsage: %s [-c conf] [-h]\n\nOptions:\n", filepath.Base(os.Args[0]))
	flag.PrintDefaults()
	fmt.Println()
}

func init() {
	var (
		conf        string
		help        bool
		defaultConf string
	)

	execDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	defaultConf = fmt.Sprintf("%s/.%s.yaml", execDir, filepath.Base(os.Args[0]))

	flag.StringVar(&conf, "c", defaultConf, "conf file")
	flag.BoolVar(&help, "h", false, "show help information")

	flag.Parse()
	flag.Usage = usage

	if help {
		flag.Usage()
		os.Exit(-1)
	}

	paths, name := filepath.Split(conf)
	Config = viper.New()
	Config.SetConfigFile(fmt.Sprintf("%s%s", paths, name))
}

func InitParStr(keyStr string) string {
	if err := Config.ReadInConfig(); err != nil {
		Logger.Info(err.Error())
	}
	return Config.GetString(keyStr)
}

func Parm() {
}
