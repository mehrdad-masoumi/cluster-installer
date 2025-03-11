package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
)

func ExecCommand(name string, args ...string) error {

	cmd := exec.Command(name, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf(Red+"error executing %s: %v"+Reset, name, err)
	}

	fmt.Printf(Green+"%s successfuly\n"+Red, name)
	return nil
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("please provide a command")
	}

	switch os.Args[1] {

	case "--install-requirements":

		err := ExecCommand("apt-get", "update", "-y")
		if err != nil {
			return
		}

		err = ExecCommand("sudo", "apt", "install", "-y", "python3", "python3-pip")
		if err != nil {
			return
		}

		err = ExecCommand("pip3", "install", "--upgrade", "pip")
		if err != nil {
			return
		}

		err = ExecCommand("pip3", "install", "ansible")
		if err != nil {
			return
		}

		err = ExecCommand("pip3", "install", "--upgrade", "pip")
		if err != nil {
			return
		}

		fmt.Println(Green + "success install requirements" + Reset)

	case "--kubernetes-requirements":

		err := ExecCommand("sudo", "git", "config", "--global", "--add", "safe.directory", "/home/ubuntu/kubespray")
		if err != nil {
			return
		}

		fmt.Println("clone kubespray...")
		if _, err := os.Stat("kubespray"); os.IsNotExist(err) {
			err = ExecCommand("sudo", "git", "clone", "https://github.com/kubernetes-sigs/kubespray.git")
			if err != nil {
				return
			}
		} else {
			fmt.Println(Yellow + "kubespray already exists, update directly" + Reset)
			err := ExecCommand("sudo", "git", "-C", "kubespray", "pull")
			if err != nil {
				return
			}
		}

		fmt.Println(Yellow + "install requirements..." + Reset)

		err = ExecCommand("pip3", "install", "-r", "kubespray/requirements.txt")
		if err != nil {
			return
		}

	case "--kubernetes-cluster":

		if len(os.Args) < 3 {
			log.Fatal("please provide a inventory file")
		}

		fmt.Println("ping to all servers is running...")

		err := ExecCommand("ansible", "kubernetes_servers", "-i", os.Args[2], "-m", "ping")
		if err != nil {
			log.Fatalf(Red+"connection is not successful: %v"+Reset, err)
		}

		err = os.Chdir("kubespray")
		if err != nil {
			log.Fatalf(Red+"kubespray: %v"+Reset, err)
		}

		err = ExecCommand("ansible-playbook", "-i", os.Args[2], "--become", "--become-user=root", "cluster.yml")
		if err != nil {
			return
		}

	case "--haproxy":

		if len(os.Args) < 3 {
			log.Fatal("please provide a inventory file")
		}

		if len(os.Args) < 4 {
			log.Fatal("please provide a playbook file")
		}

		err := ExecCommand("ansible-playbook", "-i", os.Args[2], "--become", "--become-user=root", os.Args[3])
		if err != nil {
			return
		}

	case "--docker":

		if len(os.Args) < 3 {
			log.Fatal("please provide a inventory file")
		}

		if len(os.Args) < 4 {
			log.Fatal("please provide a playbook file")
		}

		err := ExecCommand("ansible-playbook", "-i", os.Args[2], "--become", "--become-user=root", os.Args[3])
		if err != nil {
			return
		}

	case "--gitlab-kubernetes":

		if len(os.Args) < 3 {
			log.Fatal("please provide a inventory file")
		}

		if len(os.Args) < 4 {
			log.Fatal("please provide a playbook file")
		}

		err := ExecCommand("ansible-playbook", "-i", os.Args[2], "--become", "--become-user=root", os.Args[3])
		if err != nil {
			return
		}

	case "--gitlab":

		if len(os.Args) < 3 {
			log.Fatal("please provide a domain name")
		}

		err := ExecCommand("apt-get", "update", "-y")
		if err != nil {
			return
		}

		err = ExecCommand("sudo", "apt-get", "install", "-y", "curl", "openssh-server", "ca-certificates", "tzdata", "perl")
		if err != nil {
			return
		}

		err = ExecCommand("sudo", "apt-get", "install", "-y", "postfix")
		if err != nil {
			return
		}

		err = ExecCommand("curl", "https://packages.gitlab.com/install/repositories/gitlab/gitlab-ce/script.deb.sh", "|", "sudo", "bash")
		if err != nil {
			return
		}

		err = ExecCommand("sudo", "EXTERNAL_URL=", os.Args[2], "apt-get", "install", "gitlab-ce")
		if err != nil {
			return
		}

		fmt.Println(Green + "success install gitlab" + Reset)

	case "--help":
		fmt.Println("ansible-playbook  -i /mnt/d/ws/ansible/inventory.yml --become --become-user=root /mnt/d/ws/ansible/tasks/install_docker.yml")
		fmt.Println("ansible-playbook  -i /mnt/d/ws/ansible/inventory.yml --become --become-user=root /mnt/d/ws/ansible/tasks/install_gitlab.yml")
	default:
		log.Fatal("please provide a command")
	}

}
