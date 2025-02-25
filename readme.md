
# Kubernetes Cluster Setup atomates with Go

This project includes a Go script that automates the installation and configuration of the required packages to set up **Kubernetes** and **HAProxy**. The script uses **Ansible** to execute playbooks and allows you to easily deploy and configure a Kubernetes cluster and HAProxy.

## 📥 Installation
Install WSL on your device,
if you want use WSL please flow these steps:

**make sure your ssh key has been copied on WSL**


Clone the repository on your device:
```bash
git clone https://github.com/mehrdad-masoumi/cluster-installer.git
cd cluster-installer
```

Build the project on your device:
```bash
GOOS=linux GOARCH=amd64 go build -o cluster-installer main.go
```


open your terminal as administrator and run below command

```bash
wsl -d ubuntu
```

Move the binary to `/usr/local/bin` for global access:
```bash
sudo cp /mnt/.../cluster-installer /usr/local/bin
```

Make the binary executable:
```bash
chmod +x cluster-installer
```

## 🔧 Usage
Run the installer with different flags on WSL:

### Install requirements:
```bash
cluster-installer --install-requirements
```

### Install Kubernetes:
```bash
cluster-installer --cluster-kubernetes inventory.yml
```

Example inventory

``` code
[all]
master1 ansible_host=185.204.170.20 ip=192.168.2.254
worker1 ansible_host=185.204.170.215 ip=192.168.3.114
worker2 ansible_host=185.204.171.141 ip=192.168.3.70

[kube_control_plane]
master1

[etcd]
master1

[kube_node]
worker1
worker2

[k8s_cluster:children]
kube_control_plane
kube_node


[all:vars]
ansible_user=ubuntu
ansible_ssh_private_key_file=/home/ubuntu/.ssh/id_rsa
```
### HAProxy install:
```bash
cluster-installer --cluster-haproxy
```


