# hkdefault

hkdefault is the default plugin for the fast Heroku client [hk](https://github.com/kr/hk)
default will execute the heroku gem in PATH with all parameters given to hk if the given command or hk plugin is not found.

### Install

	$ mkdir -p /usr/local/lib/hk/plugin
	$ wget -qO- https://github.com/downloads/heroku/hkdefault/darwin-amd64-hkdefault-1.gz | zcat >/usr/local/lib/hk/plugin/default
	$ chmod +x /usr/local/lib/hk/plugin/default
