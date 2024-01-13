package main

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

type WatchDog struct {
	lw      *cache.ListWatch
	objType runtime.Object
	h       cache.ResourceEventHandler

	reflector *cache.Reflector
	fifo      *cache.DeltaFIFO
	store     cache.Store
}

func NewWatchDog(lw *cache.ListWatch, objType runtime.Object, h cache.ResourceEventHandler) *WatchDog {
	store := cache.NewStore(cache.MetaNamespaceKeyFunc)

	fifo := cache.NewDeltaFIFOWithOptions(cache.DeltaFIFOOptions{
		KeyFunction:  cache.MetaNamespaceKeyFunc,
		KnownObjects: store,
	})
	reflector := cache.NewReflector(lw, objType, fifo, 0)
	return &WatchDog{lw: lw, objType: objType,
		h: h, reflector: reflector, fifo: fifo, store: store}
}
func (wd *WatchDog) Run() {
	ch := make(chan struct{})
	go func() {
		wd.reflector.Run(ch)
	}()
	for {
		wd.fifo.Pop(func(obj interface{}) error {
			for _, delta := range obj.(cache.Deltas) {
				switch delta.Type {
				case cache.Sync, cache.Added:
					wd.store.Add(delta.Object)
					wd.h.OnAdd(delta.Object)
				case cache.Updated:
					if old, exists, err := wd.store.Get(delta.Object); err == nil && exists {
						wd.store.Update(delta.Object)
						wd.h.OnUpdate(old, delta.Object)
					}
				case cache.Deleted:
					wd.store.Delete(delta.Object)
					wd.h.OnDelete(delta.Object)
				}
			}
			return nil
		})
	}
	close(ch) // 关闭通道以通知子协程退出
}

type PodHandler struct{}

var _ cache.ResourceEventHandler = &PodHandler{}

func (p PodHandler) OnAdd(obj interface{}) {
	fmt.Println("OnAdd", obj.(*v1.Pod).Name)
}
func (p PodHandler) OnUpdate(old0bj, newObj interface{}) {
	fmt.Println("OnA Add:", newObj.(*v1.Pod).Name)
}
func (p PodHandler) OnDelete(obj interface{}) {
	fmt.Println("OnAdd:", obj.(*v1.Pod).Name)
}

// reflector
func main() {
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

	wd := NewWatchDog(podLW, &v1.Pod{}, &PodHandler{})
	wd.Run()
}
