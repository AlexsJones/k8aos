package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/AlexsJones/k8aos/core"
	cm "github.com/AlexsJones/k8aos/core/configuration"
	"github.com/abiosoft/ishell"
	"github.com/dimiro1/banner"
	"github.com/fatih/color"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/clientcmd"
)

const b = `
{{ .AnsiColor.Red }} ____    ____         ____           _____             ______
{{ .AnsiColor.Red }}|    |  |    |   ____|\   \     ____|\    \        ___|\     \
{{ .AnsiColor.Red }}|    |  |    |  /    /\    \   /     /\    \      |    |\     \
{{ .AnsiColor.Red }}|    | /    // |    |  |    | /     /  \    \     |    |/____/|
{{ .AnsiColor.Red }}|    |/ _ _//  |    |__|    ||     |    |    | ___|    \|   | |
{{ .AnsiColor.Red }}|    |\    \'  |    .--.    ||     |    |    ||    \    \___|/
{{ .AnsiColor.Red }}|    | \    \  |    |  |    ||\     \  /    /||    |\     \
{{ .AnsiColor.Red }}|____|  \____\ |____|  |____|| \_____\/____/ ||\ ___\|_____|
{{ .AnsiColor.Red }}|    |   |    ||    |  |    | \ |    ||    | /| |    |     |
{{ .AnsiColor.Red }}|____|   |____||____|  |____|  \|____||____|/  \|____|_____|
{{ .AnsiColor.Red }}  \(       )/    \(      )/       \(    )/        \(    )/
{{ .AnsiColor.Red }} 	 '       '      '      '         '    '          '    '
{{ .AnsiColor.Default }}
`

func main() {
	banner.Init(os.Stdout, true, true, bytes.NewBufferString(b))

	var clientSet *kubernetes.Clientset
	var m *core.Mischief
	var conf *cm.MischiefConfig
	shell := ishell.New()

	shell.AddCmd(&ishell.Cmd{
		Name: "connect",
		Help: "Provide an absolute path to config as second argument e.g. connect /tmp/config",
		Func: func(c *ishell.Context) {

			if len(c.Args) < 1 {
				fmt.Println("Requires a full path noting the kubeconfig location")
				return
			}
			config, err := clientcmd.BuildConfigFromFlags("", c.Args[0])
			if err != nil {
				panic(err.Error())
			}
			clientSet, err = kubernetes.NewForConfig(config)
			if err != nil {
				panic(err.Error())
			}
			color.Blue("Connected to active cluster...")
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "inspect",
		Help: "inspect the current cluster containers",
		Func: func(c *ishell.Context) {
			if clientSet == nil {
				fmt.Println("Please connect first")
				return
			}
			p := core.NewProbe(clientSet)
			p.Inspect()
		}})

	shell.AddCmd(&ishell.Cmd{
		Name: "mischief",
		Help: "Destroy a pod in a random namespace (can specify with second argument)",
		Func: func(c *ishell.Context) {
			if clientSet == nil {
				fmt.Println("Please connect first")
				return
			}
			conf = cm.NewDefaultConfiguration()
			reader := bufio.NewReader(os.Stdin)
			color.Red("What namespace do you want to start some chaos in?[default options is: default]")
			text, _ := reader.ReadString('\n')

			if len(text) > 1 {
				conf.TargetNamespace = strings.TrimSpace(text)
				color.Blue("Setting namespace context to %s", text)
			}
			reader = bufio.NewReader(os.Stdin)
			color.Red("How many attacks?[default options is: 1]")
			text, _ = reader.ReadString('\n')

			if len(text) > 1 {
				conf.AttackCount, _ = strconv.Atoi(strings.TrimSpace(text))
			}
			m = core.NewMischief(clientSet)
			m.Chaos(conf)
		}})

	shell.AddCmd(&ishell.Cmd{
		Name: "again",
		Help: "Run the last mischief command again",
		Func: func(c *ishell.Context) {
			if m == nil {
				fmt.Println("You need to run mischief at least once first")
				return
			}
			m.Chaos(conf)
		}})

	shell.Start()
}
