package main

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"

	"github.com/Sirupsen/logrus"
	metav1"k8s.io/apimachinery/pkg/apis/meta/v1"
	api_v1"k8s.io/api/core/v1"
	"github.com/skippbox/kubewatch/pkg/handlers"
	"k8s.io/client-go/tools/cache"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
)

type Controller struct {
	logger 		 *logrus.Entry
	clientset 	 kubernetes.Interface
	queue 		 workqueue.RateLimitingInterface
	informer     cache.SharedIndexInformer
	eventHandler handlers.Handler
}
func main()  {
	fmt.Println("Hello World!!!")

	c, err := clientcmd.BuildConfigFromFlags("", "/home/tigerworks/.kube/config")
	if err != nil {
		fmt.Println(err)
		return
	}
	kubeClient := kubernetes.NewForConfigOrDie(c)
	pods, err := kubeClient.CoreV1().Pods(metav1.NamespaceAll).List(metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pods)
}

func newControllerPod(client kubernetes.Interface, eventHandler handlers.Handler) *Controller {
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	informer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				return client.CoreV1().Pods(metav1.NamespaceAll).List(metav1.ListOptions{})
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return client.CoreV1().Pods(metav1.NamespaceAll).Watch(metav1.ListOptions{})
			},
		},
		&api_v1.Pod{},
		0,
		cache.Indexers{},
	)
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
		UpdateFunc: func(old, new interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(new)
			if err == nil {
				queue.Add(key)
			}
		},
		DeleteFunc: func(obj interface{}) {
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
	})

	return &Controller{
		logger:       logrus.WithField("pkg", "kubewatch-pod"),
		clientset:    client,
		informer:     informer,
		queue:        queue,
		eventHandler: eventHandler,
	}
}
