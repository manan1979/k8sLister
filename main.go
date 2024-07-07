 package main

 import (
"context"     
"flag"
"k8s.io/client-go/tools/clientcmd"
"fmt"
"k8s.io/client-go/kubernetes"
metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
"k8s.io/client-go/rest"
 )

func main() {
    
   kubeconfig :=  flag.String("kubeconfig", "/home/ubuntu/.kube/config", "location to your kubeconfig  file")
//  flag.Parse()

   config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
   if err !=nil {
       fmt.Printf("error %s building config from flag\n", err.Error())
   config ,err = rest.InClusterConfig()
     if err != nil {
       fmt.Printf("error %s getting in clusterconfig", err.Error())
     }
   }
 
   clientset, err :=  kubernetes.NewForConfig(config)
   if err != nil {
       fmt.Printf("error %s creating clientset\n" , err.Error())
   }

  ctx := context.Background()
  pods, err := clientset.CoreV1().Pods("default").List(context.Background(),metav1.ListOptions{})
  if err != nil {
       fmt.Printf("error %s, while listening all the pods from default namespace\n", err.Error()) 
  }

  fmt.Println("pods from default namesoaces")
  for  _, pod := range pods.Items {
       fmt.Printf("%s",pod.Name)
  }

  fmt.Printf("Deployments area")
  deployments, err := clientset.AppsV1().Deployments("default").List(ctx, metav1.ListOptions{})
  if err != nil {
      fmt.Printf("listing deplyments %s\n", err.Error())
  }

  for _, d:= range deployments.Items {
       fmt.Printf("%s",d.Name)
  }


}
