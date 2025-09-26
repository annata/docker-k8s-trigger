package main

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func trigger(namespace string, name string) error {
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

var tagRegexp = regexp.MustCompile(`^[a-z|A-Z0-9._]+$`)

func triggerVersion(namespace string, name string, containerName string, tag string) error {
	if !tagRegexp.MatchString(tag) {
		return fmt.Errorf("invalid tag format: %s", tag)
	}
	config, err := rest.InClusterConfig()
	if err != nil {
		return err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	deployment, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), name, v1.GetOptions{})
	if err != nil {
		return err
	}
	for i, container := range deployment.Spec.Template.Spec.Containers {
		if container.Name == containerName {
			parts := strings.SplitN(container.Image, ":", 2)
			baseImage := parts[0]
			deployment.Spec.Template.Spec.Containers[i].Image = baseImage + ":" + tag
			break
		}
	}
	deployment.Spec.Template.Labels["time"] = strconv.FormatInt(time.Now().Unix(), 10)
	_, err = clientset.AppsV1().Deployments(namespace).Update(context.TODO(), deployment, v1.UpdateOptions{})
	if err != nil {
		return err
	}
	return nil
}
