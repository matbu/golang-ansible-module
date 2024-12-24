# Golang Ansible Module

### Golang Ansible Module Demo

This repository demonstrates how to use **Golang** as an Ansible module within your Ansible collections or roles. Follow the steps below to get started!

---

## Clone the Repository

```bash
git clone https://github.com/matbu/golang-ansible-module
cd golang-ansible-module
```

---

## Rebuild the Binary (Optional)

If needed, rebuild the `hello` binary by running:

```bash
go build -o plugins/modules/hello/hello plugins/modules/hello/hello_src.go
```

---

## Configure the Ansible Library Path

Ensure your Ansible library points to the current path:

```bash
export ANSIBLE_LIBRARY=$PWD/plugins/modules
```

---

## Run the Demo Playbook

Execute the demo playbook using:

```bash
ansible-playbook -i localhost playbooks/demo.yml
```

---

## Customize and Extend

Now you can adapt the code to create your own Golang-powered Ansible modules.
Enjoy building something awesome!