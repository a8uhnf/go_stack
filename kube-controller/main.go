package main

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	// "k8s.io/client-go/rest"
	// "k8s.io/client-go/"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func main()  {
	config, err := clientcmd.BuildConfigFromFlags("", "/home/hanifa/.kube/config")
	if err != nil {
		fmt.Println("Error Occured", err)
		return
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("Error Occured", err)
		return
	}
	fmt.Println("Hello Hanifa")
	pods, err := client.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error Occured", err)
		return
	}
	
	for key, value := range pods.Items {
		fmt.Println("pod key", key)
		fmt.Println("pod value", value.Name)
	}
}
