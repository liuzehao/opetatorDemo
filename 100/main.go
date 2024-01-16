package main

import (
	"fmt"
	"k8s.io/client-go/tools/cache"
)

type pod struct {
	Name  string
	Value int
}

func newPod(name string, v int) pod {
	return pod{Name: name, Value: v}
}

func podKeyFunc(obj interface{}) (string, error) {
	return obj.(pod).Name, nil
}

// demo: DeltaFIFO queue
func main() {
	df := cache.NewDeltaFIFOWithOptions(cache.DeltaFIFOOptions{KeyFunction: podKeyFunc})

	pod1 := newPod("pod1", 1)
	pod2 := newPod("pod2", 2)
	pod3 := newPod("pod3", 3)

	df.Add(pod1)
	df.Add(pod2)
	df.Add(pod3)
	pod1.Value = 10
	df.Update(pod1)
	//fmt.Println(df.List())
	df.Delete(pod1)

	df.Pop(func(obj interface{}) error {
		//fmt.Printf("%T", obj)
		for _, delta := range obj.(cache.Deltas) {
			fmt.Println(delta.Type, ":", delta.Object.(pod).Name, delta.Object.(pod).Value)
		}
		return nil
	})
}
