package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/ajmaluk/byteclaw/cmd/byteclaw/internal"
	"github.com/ajmaluk/byteclaw/cmd/byteclaw/internal/agent"
	"github.com/ajmaluk/byteclaw/cmd/byteclaw/internal/auth"
	"github.com/ajmaluk/byteclaw/cmd/byteclaw/internal/cron"
	"github.com/ajmaluk/byteclaw/cmd/byteclaw/internal/gateway"
	"github.com/ajmaluk/byteclaw/cmd/byteclaw/internal/migrate"
	"github.com/ajmaluk/byteclaw/cmd/byteclaw/internal/onboard"
	"github.com/ajmaluk/byteclaw/cmd/byteclaw/internal/skills"
	"github.com/ajmaluk/byteclaw/cmd/byteclaw/internal/status"
	"github.com/ajmaluk/byteclaw/cmd/byteclaw/internal/version"
)

func NewByteclawCommand() *cobra.Command {
	short := fmt.Sprintf("%s byteclaw - Personal AI Assistant v%s\n\n", internal.Logo, internal.GetVersion())

	cmd := &cobra.Command{
		Use:     "byteclaw",
		Short:   short,
		Example: "byteclaw list",
	}

	cmd.AddCommand(
		onboard.NewOnboardCommand(),
		agent.NewAgentCommand(),
		auth.NewAuthCommand(),
		gateway.NewGatewayCommand(),
		status.NewStatusCommand(),
		cron.NewCronCommand(),
		migrate.NewMigrateCommand(),
		skills.NewSkillsCommand(),
		version.NewVersionCommand(),
	)

	return cmd
}

const (
	colorBlue = "\033[1;38;2;62;93;185m"
	colorRed  = "\033[1;38;2;213;70;70m"
	banner    = "\r\n" +
		colorBlue + "██████╗ ██╗   ██╗████████╗███████╗" + colorRed + " ██████╗██╗      █████╗ ██╗    ██╗\n" +
		colorBlue + "██╔══██╗╚██╗ ██╔╝╚══██╔══╝██╔════╝" + colorRed + "██╔════╝██║     ██╔══██╗██║    ██║\n" +
		colorBlue + "██████╔╝ ╚████╔╝    ██║   █████╗  " + colorRed + "██║     ██║     ███████║██║ █╗ ██║\n" +
		colorBlue + "██╔══██╗  ╚██╔╝     ██║   ██╔══╝  " + colorRed + "██║     ██║     ██╔══██║██║███╗██║\n" +
		colorBlue + "██████╔╝   ██║      ██║   ███████╗" + colorRed + "╚██████╗███████╗██║  ██║╚███╔███╔╝\n" +
		colorBlue + "╚═════╝    ╚═╝      ╚═╝   ╚══════╝" + colorRed + " ╚═════╝╚══════╝╚═╝  ╚═╝ ╚══╝╚══╝\n" +
		"\033[0m\r\n"
)

func main() {
	fmt.Printf("%s", banner)
	cmd := NewByteclawCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
