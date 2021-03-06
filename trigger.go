package main

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"time"
)

func trigger(namespace string, name string) error {

	//namespaceByte, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	//if err != nil {
	//	return err
	//}
	config, err := rest.InClusterConfig()
	if err != nil {
		return err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	_, err = clientset.AppsV1().Deployments(namespace).Patch(context.TODO(),
		name, types.StrategicMergePatchType, []byte(fmt.Sprintf("{\"spec\":{\"template\":{\"metadata\":{\"labels\":{\"time\":\"%d\"}}}}}", time.Now().Unix())), v1.PatchOptions{})
	if err != nil {
		return err
	}
	return nil
}
