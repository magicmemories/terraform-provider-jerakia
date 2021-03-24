policy :default do
  lookup :default do
    datasource :file, {
      :docroot    => "/var/lib/jerakia/data",
      :enable_caching => true,
      :searchpath => [
        "host/#{scope[:hostname]}",
        "env/#{scope[:env]}",
        "common",
      ],
    }
    filter :strsub, scope
  end
end
