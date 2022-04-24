## 云原生总结
### go语言
* 控制结构
* 常用数据结构(常量、变量、数组、切片、map、结构体、指针)
* 函数(Main函数、Init函数、返回值、内置函数、回调函数、闭包、方法、接口)
* 错误处理
* 多线程
* 协程
* 线程加锁
* 线程调度
* 内存管理
* 包引用与依赖管理
* Makefile

### docker核心技术
* 虚拟机和容器性能对比
![image](https://user-images.githubusercontent.com/35921801/164977036-918f276a-af8e-4515-a526-3fd7dd7b3254.png)
* 容器操作
    - 启动
    - 查看容器，进入容器
* 容器主要特性
    - 容器主要特性(安全性、隔离性、便携性、可配额)
* 网络
    - ![image](https://user-images.githubusercontent.com/35921801/164977211-42285d89-9f66-4339-b884-de9eefe73702.png)

### kubernetes架构原则和对象设计
* 核心对象
    - Node
    - Namespace
    - Pod
    - Service
* 主要组件
![image](https://user-images.githubusercontent.com/35921801/164977463-a4cd95cf-bd65-41dd-8a14-0355ef3e368e.png)
   - Kubernetes的主节点(Master Node)
   - Kubernetes的工作节点(worker Node)
   - etcd 基于Raft开发的分布式key-value存储
   - APIServer 认证，授权，准入
   - Controller Manager 集群的大脑
   - Scheduler 调度
   - Kubelet
   - kube-Proxy

* 深入理解Kubernetes
   - kebernetes生态系统
![image](https://user-images.githubusercontent.com/35921801/164977735-a785e60c-7c74-42cc-ac49-83a2fc07546c.png)
 
### Kubernetes控制平面组件：etcd
* etcd
   - etcd:分布式key-value存储
   - 主要功能，用于服务发现、共享配置以及一致性保障
   - 服务注册与发现
   ![image](https://user-images.githubusercontent.com/35921801/164978177-ed1fa8de-4e77-444c-93cc-447f31e0bc70.png)
   - 消息发布与订阅
   ![image](https://user-images.githubusercontent.com/35921801/164978208-51fe288f-ec19-40fc-a233-9a9892960ecd.png)

* Raft协议
![image](https://user-images.githubusercontent.com/35921801/164978227-6b28a16a-8731-4fd7-b9f1-1bf97b013b65.png)
* 高可用集群管理
![image](https://user-images.githubusercontent.com/35921801/164978264-4283e16c-dab9-414c-9459-656517de7a02.png)
* etcd常见问题

### 深入理解kube-apiserver

* 认证
   - 认证插件
   ![image](https://user-images.githubusercontent.com/35921801/164978524-e44ab34f-6491-4c21-a52b-73b63573bb21.png)
   - 基于Webhook的认证服务集成
* 鉴权
   - 授权
* 准入
   - Mutating
   - Validating
   - Admission
* 限流
* API Server对象的发现

### Kubernetes控制平面组件
* 调度
   - kube-scheduler
   - 调度器
   - Predicates策略
   - Priorities策略
* Controller Manager
   - 控制器的工作流程
  ![image](https://user-images.githubusercontent.com/35921801/164978709-1475394d-3d4d-4a24-9b87-7eb8700046e3.png)
   - 通用Controller 
* kubelet
   - kubelet架构
   ![image](https://user-images.githubusercontent.com/35921801/164978751-050415f1-7afe-42ba-98ca-c0c0dace6bc4.png)

* CRI
![image](https://user-images.githubusercontent.com/35921801/164978770-00e17617-2eaf-4e19-9c62-be19c4e971eb.png)
* CNI
* CSI

### Kubernetes控制平面组件：生命周期管理和服务发现
* 深入理解Pod的生命周期
![image](https://user-images.githubusercontent.com/35921801/164978906-b5583e71-7eae-4810-a16f-d4385cf09127.png)
* 服务发现
   - 服务发布
   ![image](https://user-images.githubusercontent.com/35921801/164978935-967fa533-a65e-479f-899f-80d7d82e58f4.png)
* 微服务架构下的高可用挑战
* service对象
* kube-proxy
![image](https://user-images.githubusercontent.com/35921801/164978970-d0f07c7d-a8fb-4171-a259-d269d4bde301.png)

### 生产化集群的管理
* 计算节点相关
![image](https://user-images.githubusercontent.com/35921801/164979156-0e5aa408-aa78-4d1d-a960-1a27951f7372.png)
* 操作系统选择
* 节点资源管理
![image](https://user-images.githubusercontent.com/35921801/164979184-12af7892-8ad9-4230-88d4-0ab06f7e6596.png)
* 节点异常检测
* 常用节点问题排查手段
![image](https://user-images.githubusercontent.com/35921801/164979208-1c1c492e-1712-4e5b-9ab2-a45e115adf2a.png)
* 基于 extended resource 扩展节点资源
* 构建和管理高可用集群
* Cluster Autoscaler
* 集群管理实践案例分享
* 多租户集群管理

### Kubernetes的生产化运维

* 镜像仓库
* 镜像安全
* 基于Kubernetes的DevOps
* 基于GitHub Action的自动化流水线
* 基于 Jenkins的自动化流水线
* Tekton
* Argo CD
* 监控和日志

### 将应用迁移至Kubernetes平台

* 应用接入最佳实践
   - 应用容器化
   - 将应用迁移至Kubernetes
   - 有状态应用的复杂性讨论

* Spec管理神器 - Helm
* metrics - server
* 自动扩容缩容 - HPA
* 自动扩容缩容 - VPA
* 如何解决社区基础弹性能力不足的问题

### 基于Istio的高级流量管理

* 微服务架构的演变
* 微服务到服务网格还缺什么
* 深入理解数据平面 Envoy
* Istio流量管理
* 跟踪采样







