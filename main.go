package main

import (
	"log"

	"github.com/go-kratos/kratos/v2"
	"github.com/tkeel-io/tkeel-template/cmd/project"
	"github.com/tkeel-io/tkeel-template/cmd/proto"
	"github.com/tkeel-io/tkeel-template/cmd/run"
	"github.com/tkeel-io/tkeel-template/cmd/upgrade"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "kratos",
	Short:   "Kratos: An elegant toolkit for Go microservices.",
	Long:    `Kratos: An elegant toolkit for Go microservices.`,
	Version: kratos.Release,
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(proto.CmdProto)
	rootCmd.AddCommand(upgrade.CmdUpgrade)
	rootCmd.AddCommand(run.CmdRun)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
