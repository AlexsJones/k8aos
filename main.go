package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/AlexsJones/k8oas/core"
	"github.com/abiosoft/ishell"
	"github.com/dimiro1/banner"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/clientcmd"
)

const b = `
{{ .AnsiColor.Red }} __   ___  _______     ______      __        ________
{{ .AnsiColor.Red }}|/"| /  ")/"  _  \\   /    " \    /""\      /"       )
{{ .AnsiColor.Red }}(: |/   /|:  _ /  :| // ____  \  /    \    (:   \___/
{{ .AnsiColor.Red }}|    __/  \___/___/ /  /    ) :)/' /\  \    \___  \
{{ .AnsiColor.Red }}(// _  \  //  /_ \\(: (____/ ////  __'  \    __/  \\
{{ .AnsiColor.Red }}|: | \  \|:  /_   :|\        //   /  \\  \  /" \   :)
{{ .AnsiColor.Red }}(__|  \__)\_______/  \"_____/(___/    \___)(_______/
{{ .AnsiColor.Default }}
`

func main() {
	banner.Init(os.Stdout, true, true, bytes.NewBufferString(b))

	var clientSet *kubernetes.Clientset

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
			fmt.Println("Connected to active cluster...")
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "inspect",
		Help: "inspect the current cluster containers",
		Func: func(c *ishell.Context) {

			p := core.NewProbe(clientSet)
			p.Inspect()
		}})

	shell.AddCmd(&ishell.Cmd{
		Name: "mischief",
		Help: "Destroy a pod in a random namespace (can specify with second argument)",
		Func: func(c *ishell.Context) {

		}})

	shell.Start()
}
