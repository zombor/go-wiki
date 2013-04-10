shared_dir = '/vagrant'
user = 'vagrant'
home = "/home/#{user}"

ENV['HOME'] = home
ENV['USER'] = user
ENV['SHARED'] = shared_dir

ruby_block 'Give root access to the forwarded ssh agent' do
  block do
    agents = {}
    ppid = Process.ppid
    Dir.glob('/tmp/ssh*/agent*').each do |fn|
      agents[fn.match(/agent\.(\d+)$/)[1]] = fn
    end
    while ppid != '1'
      if (agent = agents[ppid])
        ENV['SSH_AUTH_SOCK'] = agent
        break
      end
      File.open("/proc/#{ppid}/status", "r") do |file|
        ppid = file.read().match(/PPid:\s+(\d+)/)[1]
      end
    end
  end
  action :create
end

bash 'apt-get update' do
  code 'sudo apt-get update'
  not_if 'which git'
end

package 'libyaml-dev' do
  action :install
end

package 'curl' do
  action :install
end
