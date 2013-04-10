%w(.bash_login .bash_logout .bashrc postinstall.sh .profile .sudo_as_admin_successful .zlogin VBoxGuestAdditions_4.1.*.iso).each do |f|
  file "#{ENV['HOME']}/#{f}" do
    action :delete
  end
end
