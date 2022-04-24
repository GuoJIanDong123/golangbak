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






