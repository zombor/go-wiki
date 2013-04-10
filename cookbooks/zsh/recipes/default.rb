package 'zsh' do
  version '4.3.17-1ubuntu1'
  action :install
end

directory "#{ENV['HOME']}/.bin" do
  owner ENV['USER']
  action :create
end

vcprompt_exec = "#{ENV['HOME']}/.bin/vcprompt"
bash 'install vcprompt' do
  user ENV['USER']
  code <<-EOH
  curl -sL https://github.com/djl/vcprompt/raw/master/bin/vcprompt > #{vcprompt_exec}
  chmod 755 #{vcprompt_exec}
  EOH
  not_if "test -e #{vcprompt_exec}"
end

bash 'clone zsh-syntax-highlighting repo' do
  user ENV['USER']
  code "git clone https://github.com/zsh-users/zsh-syntax-highlighting.git #{ENV['HOME']}/zsh-syntax-highlighting"
  not_if "test -d #{ENV['HOME']}/zsh-syntax-highlighting"
end

template "#{ENV['HOME']}/.zshrc" do
  owner ENV['USER']
  action :create
end

cookbook_file "#{ENV['HOME']}/.aliases" do
  owner ENV['USER']
  action :create_if_missing
end

bash 'make ZSH the default login shell' do
  code "sudo chsh -s `which zsh` #{ENV['USER']}"
end
