vim_dotfiles_dir = "#{ENV['HOME']}/.vim_dotfiles"
vim_version = '2:7.3.429-2ubuntu2.1'

package 'vim-common' do
  #version vim_version
  action :install
end

package 'vim-runtime' do
  #version vim_version
  action :install
end

package 'vim-nox' do
  #version vim_version
  action :install
end

bash "clone vim_dotfiles repo" do
  user ENV['USER']
  code <<-EOH
  git clone git@github.com:sittercity/vim_files.git #{vim_dotfiles_dir}
  cd #{vim_dotfiles_dir}
  git submodule update --init
  EOH
  not_if "test -d #{vim_dotfiles_dir}"
end

bash "symlink vimrc" do
  user ENV['USER']
  code "ln -s #{vim_dotfiles_dir}/vimrc #{ENV['HOME']}/.vimrc"
  not_if "ls -l #{ENV['HOME']}/.vimrc"
end

bash "symlink vim dir" do
  user ENV['USER']
  code "ln -s #{vim_dotfiles_dir}/vim #{ENV['HOME']}/.vim"
  not_if "ls -l #{ENV['HOME']}/.vim"
end

#bash "build command-t" do
#  user ENV['USER']
#  code <<-EOH
#  cd #{vim_dotfiles_dir}/vim/bundle/command-t/ruby/command-t
#  #{ENV['HOME']}/.rvm/bin/rvm ruby18 do ruby extconf.rb
#  make clean
#  make
#  EOH
#end
