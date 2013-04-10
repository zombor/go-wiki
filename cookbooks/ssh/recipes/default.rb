cookbook_file "#{ENV['HOME']}/.ssh/config" do
  owner ENV['USER']
  action :create
end
