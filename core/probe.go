package core

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//Probe ...
type Probe struct {
	ClientSet *kubernetes.Clientset
}

//NewProbe ...
func NewProbe(clientSet *kubernetes.Clientset) *Probe {

	return &Probe{clientSet}
}

//Inspect ...
func (p *Probe) Inspect() {
	pods, err := p.ClientSet.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, p := range pods.Items {
		fmt.Printf("%s --- %s\n", p.Namespace, p.Name)
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
}
