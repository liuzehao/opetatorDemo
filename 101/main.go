package main

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func main() {

	//create config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}

	//create client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	podLW := cache.NewListWatchFromClient(clientset.CoreV1().RESTClient(), "pods", "default", fields.Everything())
	//list function
	list, err := podLW.List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	podList := list.(*v1.PodList)
	for _, pod := range podList.Items {
		fmt.Printf(pod.Name)
	}

	//	watch function
	watcher, err := podLW.Watch(metav1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	for {
		select {
		case v, ok := <-watcher.ResultChan():
			if ok {
				fmt.Println(v.Type, ":", v.Object.(*v1.Pod).Name)
			}
		}
	}
}
