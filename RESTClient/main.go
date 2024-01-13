package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// 获取系统家目录
func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}

	// for windows
	return os.Getenv("USERPROFILE")
}

func main() {
	var kubeConfig *string
	var err error
	var config *rest.Config

	// 获取 kubeconfig 文件路径
	if h := homeDir(); h != "" {
		kubeConfig = flag.String("kubeConfig", filepath.Join(h, ".kube", "config"), "use kubeconfig to access kube-apiserver")
	} else {
		kubeConfig = flag.String("kubeConfig", "", "use kubeconfig to access kube-apiserver")
	}
	flag.Parse()

	// 获取 kubeconfig
	config, err = clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		panic(err.Error())
	}

	// 使用 RESTClient 需要开发者自行设置资源 URL
	// pod 资源没有 group，在核心组，所以前缀是 api
	config.APIPath = "api"
	// 设置 corev1 groupVersion
	config.GroupVersion = &corev1.SchemeGroupVersion
	// 设置解析器，用于用于解析 scheme
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	// 初始化 RESTClient
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err.Error())
	}
	// 调用结果用 podList 解析
	result := &corev1.PodList{}
	// 获取 kube-system 命名空间的 pod
	namespace := "kube-system"
	// 链式调用 RESTClient 方法获取，并将结果解析到 corev1.PodList{}
	err = restClient.Get().Namespace(namespace).Resource("pods").Do(context.TODO()).Into(result)
	if err != nil {
		panic(err.Error())
	}

	// 打印结果
	for _, pod := range result.Items {
		fmt.Printf("namespace: %s, pod: %s\n", pod.Namespace, pod.Name)
	}
}
