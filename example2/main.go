package main

import (
	"context"
	"path/filepath"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	home := homedir.HomeDir()
	kubeconfig := filepath.Join(home, ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	coreClient := clientset.CoreV1()

	ctx := context.Background()
	pod := &v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "my-pod",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  "my-ctr",
					Image: "httpd:alpine",
				},
			},
		},
	}
	_, err = coreClient.Pods("default").Create(ctx, pod, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
}
