# frozen_string_literal: true

# Debian host for local Temporal workflow development.
#
# Useful overrides:
#   PRIME_MOVER_VM_IP=192.168.56.50 vagrant up
#   PRIME_MOVER_VM_BOX=bento/debian-12 VAGRANT_DEFAULT_PROVIDER=vmware_desktop vagrant up
#   TEMPORAL_VERSION=1.25.2 TEMPORAL_UI_VERSION=2.31.2 vagrant provision
#
# Host endpoints:
#   Temporal frontend: 127.0.0.1:7233
#   Temporal UI:       http://127.0.0.1:8233

require "rbconfig"

VAGRANTFILE_API_VERSION = "2"

HOST_ARM = RbConfig::CONFIG["host_cpu"].match?(/arm64|aarch64/)
DEFAULT_PROVIDER = HOST_ARM ? "vmware_desktop" : "virtualbox"
DEFAULT_BOX = HOST_ARM ? "bento/debian-12" : "generic/debian12"
DEFAULT_BOX_ARCHITECTURE = HOST_ARM ? "arm64" : "amd64"

ENV["VAGRANT_DEFAULT_PROVIDER"] ||= DEFAULT_PROVIDER

VM_BOX = ENV.fetch("PRIME_MOVER_VM_BOX", DEFAULT_BOX)
VM_BOX_ARCHITECTURE = ENV.fetch("PRIME_MOVER_VM_BOX_ARCHITECTURE", DEFAULT_BOX_ARCHITECTURE)
VM_HOSTNAME = ENV.fetch("PRIME_MOVER_VM_HOSTNAME", "prime-mover-temporal")
VM_IP = ENV.fetch("PRIME_MOVER_VM_IP", "192.168.56.50")
VM_PRIVATE_NETWORK = ENV.fetch("PRIME_MOVER_VM_PRIVATE_NETWORK", HOST_ARM ? "false" : "true") == "true"
VM_CPUS = ENV.fetch("PRIME_MOVER_VM_CPUS", "2").to_i
VM_MEMORY = ENV.fetch("PRIME_MOVER_VM_MEMORY", "4096").to_i
TEMPORAL_VERSION = ENV.fetch("TEMPORAL_VERSION", "1.25.2")
TEMPORAL_UI_VERSION = ENV.fetch("TEMPORAL_UI_VERSION", "2.31.2")
POSTGRES_VERSION = ENV.fetch("POSTGRES_VERSION", "16")
DOCKER_DNS_SERVERS = ENV.fetch("PRIME_MOVER_DOCKER_DNS_SERVERS", "1.1.1.1,8.8.8.8").split(",").map(&:strip).reject(&:empty?)

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = VM_BOX
  config.vm.box_architecture = VM_BOX_ARCHITECTURE unless VM_BOX_ARCHITECTURE.empty?
  config.vm.hostname = VM_HOSTNAME

  config.vm.network "private_network", ip: VM_IP if VM_PRIVATE_NETWORK
  config.vm.network "forwarded_port", guest: 7233, host: 7233, auto_correct: true
  config.vm.network "forwarded_port", guest: 8233, host: 8233, auto_correct: true

  config.vm.provider "virtualbox" do |vb|
    vb.name = VM_HOSTNAME
    vb.cpus = VM_CPUS
    vb.memory = VM_MEMORY
  end

  config.vm.provider "vmware_desktop" do |vmware|
    vmware.vmx["displayName"] = VM_HOSTNAME
    vmware.vmx["numvcpus"] = VM_CPUS.to_s
    vmware.vmx["memsize"] = VM_MEMORY.to_s
    vmware.vmx["ethernet0.pcislotnumber"] = "160"
    vmware.vmx["ethernet1.pcislotnumber"] = "224"
  end

  config.vm.provision "shell",
                      name: "refresh package index for ansible_local",
                      inline: "apt-get clean && rm -rf /var/lib/apt/lists/* && apt-get update"

  config.vm.provision "ansible_local" do |ansible|
    ansible.install = true
    ansible.playbook = "provisioning/playbook.yml"
    ansible.extra_vars = {
      "temporal_version" => TEMPORAL_VERSION,
      "temporal_ui_version" => TEMPORAL_UI_VERSION,
      "postgres_version" => POSTGRES_VERSION,
      "docker_dns_servers" => DOCKER_DNS_SERVERS
    }
  end

  config.vm.post_up_message = <<~MESSAGE
    Temporal VM is configured.

    Host Temporal address:  127.0.0.1:7233
    Host UI:                http://127.0.0.1:8233
    Guest Temporal address: #{VM_PRIVATE_NETWORK ? "#{VM_IP}:7233" : "not configured"}
    Guest UI:               #{VM_PRIVATE_NETWORK ? "http://#{VM_IP}:8233" : "not configured"}

    Useful commands:
      vagrant ssh
      temporal-compose ps
      temporal-compose logs -f temporal
      docker run --rm --network host temporalio/admin-tools:#{TEMPORAL_VERSION} temporal operator cluster health
  MESSAGE
end
