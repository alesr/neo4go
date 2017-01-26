Vagrant.configure("2") do |config|

  # Default VM spec
  config.vm.box = "bento/centos-7.2"
  config.ssh.forward_agent = true
  config.vm.provider "virtualbox" do |vb|
    vb.customize [
      "modifyvm", :id,
      "--memory", "2048",
      "--cpus", "1",
      "--ioapic", "on",
      "--pae", "on",
      "--hwvirtex", "on",
      "--vtxvpid", "on",
      "--vtxux", "on",
      "--nestedpaging", "off"
    ]
  end

  $controller = {
    "tag" => "controller",
    "IP" => "192.168.11.2"
  }

  # Morpheus yells: MACHINES!
  $nodes = {
    "restapi" => {
      "tag" => "neo4j",
      "IP" => "192.168.11.3",
      "synced_folder_disabled" => true
    }
  }

  $service_restart = "nmcli connection reload; systemctl restart network.service"

  # Boot them'all (nodes)
  $nodes.each do |key, value|
    tag = value["tag"]
    ip = value["IP"]
    syncf = value["synced_folder_disabled"]
    config.vm.define tag do |t|
      t.vm.hostname = tag
      t.vm.network "private_network", ip: ip
      t.vm.synced_folder '.', '/vagrant', disabled: syncf
      t.vm.provision "shell", inline: $service_restart
      t.vm.provision "shell", privileged: false, inline: <<-EOF
        echo "#{key} running on local server address http://#{ip}"
      EOF
    end
  end

  # controller up!
  config.vm.define "controller" do |controller|
    controller.vm.hostname = $controller["tag"]
    controller.vm.network "private_network", ip: $controller["IP"]
    controller.vm.provider :virtualbox do |vb|
      vb.customize ["modifyvm", :id, "--memory", "512"]
    end
    # Execute provision playbook through Ansible instance installed on your local
    # enviroment to set up Ansible on controller node
    controller.vm.provision "ansible" do |ansible|
      ansible.playbook = "provision/playbooks/provision.yml"
    end
    controller.vm.provision "shell", inline: $service_restart
    controller.vm.provision "shell", privileged: false, inline: <<-EOF
      echo "#{$controller["tag"]} running on local server address http://#{$controller["IP"]}"
    EOF
  end
end
