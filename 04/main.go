package main

import (
	"context"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	//RESTClient
	//config
	//config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	//if err != nil {
	//	panic(err)
	//}
	//config.GroupVersion = &v1.SchemeGroupVersion
	//config.NegotiatedSerializer = scheme.Codecs
	//config.APIPath = "/api"
	//
	//// client
	//restClient, err := rest.RESTClientFor(config)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// get data
	//pod := v1.Pod{}
	//err = restClient.Get().Namespace("default").Resource("pods").Name("nginx-deployment-8d545c96d-h9rhj").Do(context.TODO()).Into(&pod)
	//if err != nil {
	//	println(err)
	//} else {
	//	println(pod.Name)
	//}
	//// config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	coreV1 := clientset.CoreV1()
	var pods, _ = coreV1.Pods("default").List(context.TODO(), v1.ListOptions{})
	println(pods.ListMeta.Marshal())

}