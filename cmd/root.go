package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	path, _ = os.Executable()
	_, exec = filepath.Split(path)
)

var (
	url    string
	vulnId int
)
var rootCmd = &cobra.Command{
	Use:   exec,
	Short: exec,
	Long:  `一个简单的致远OA安全测试工具，目的是为了协助漏洞自查、修复工作。`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}