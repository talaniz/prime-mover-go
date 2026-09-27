# frozen_string_literal: true

# Debian host for local Temporal workflow development.
#
# Useful overrides:
#   PRIME_MOVER_VM_IP=192.168.56.50 vagrant up
#   TEMPORAL_VERSION=1.25.2 TEMPORAL_UI_VERSION=2.31.2 vagrant provision
#
# Guest endpoints:
#   Temporal frontend: 192.168.56.50:7233
#   Temporal UI:       http://192.168.56.50:8233

VAGRANTFILE_API_VERSION = "2"

VM_BOX = ENV.fetch("PRIME_MOVER_VM_BOX", "generic/debian12")
VM_HOSTNAME = ENV.fetch("PRIME_MOVER_VM_HOSTNAME", "prime-mover-temporal")
VM_IP = ENV.fetch("PRIME_MOVER_VM_IP", "192.168.56.50")
VM_CPUS = ENV.fetch("PRIME_MOVER_VM_CPUS", "2").to_i
VM_MEMORY = ENV.fetch("PRIME_MOVER_VM_MEMORY", "4096").to_i
TEMPORAL_VERSION = ENV.fetch("TEMPORAL_VERSION", "1.25.2")
TEMPORAL_UI_VERSION = ENV.fetch("TEMPORAL_UI_VERSION", "2.31.2")
POSTGRES_VERSION = ENV.fetch("POSTGRES_VERSION", "16")

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = VM_BOX
  config.vm.hostname = VM_HOSTNAME

  config.vm.network "private_network", ip: VM_IP
  config.vm.network "forwarded_port", guest: 7233, host: 7233, auto_correct: true
  config.vm.network "forwarded_port", guest: 8233, host: 8233, auto_correct: true

  config.vm.provider "virtualbox" do |vb|
    vb.name = VM_HOSTNAME
    vb.cpus = VM_CPUS
    vb.memory = VM_MEMORY
  end

  config.vm.provision "ansible_local" do |ansible|
    ansible.install = true
    ansible.playbook = "provisioning/playbook.yml"
    ansible.extra_vars = {
      "temporal_version" => TEMPORAL_VERSION,
      "temporal_ui_version" => TEMPORAL_UI_VERSION,
      "postgres_version" => POSTGRES_VERSION
    }
  end

  config.vm.post_up_message = <<~MESSAGE
    Temporal VM is configured.

    Guest Temporal address: #{VM_IP}:7233
    Guest UI:               http://#{VM_IP}:8233

    Useful commands:
      vagrant ssh
      temporal-compose ps
      temporal-compose logs -f temporal
      docker run --rm --network host temporalio/admin-tools:#{TEMPORAL_VERSION} temporal operator cluster health
  MESSAGE
end
