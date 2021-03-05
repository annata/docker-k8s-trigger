# k8s集群一键触发重新部署
## 介绍
- 本项目用户在k8s集群中部署一个简单的http服务,get访问某个地址可以触发k8s集群中某个deploy进行重新部署
- 可以用于手动在浏览器上访问某个url触发重新部署或者是在脚本中curl触发
- 原理是通过api_server对某个deploy进行patch修改
## 用法
- 最简单的通过命令 `kubectl apply -f k8s-trigger.yml` 即可一键部署
- 通过名称 `kubectl get svc -n k8s-trigger` 即可查看部署的elb的url
- 通过 `http://{elb的url}/{secret}/{命名空间}/{deploy名称即可重新部署deploy}`
