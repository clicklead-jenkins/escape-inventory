---
title: "Configuring an Inventory"
slug: quickstart-configure-inventory 
type: "docs"
toc: true

back: /docs/installation/
backLabel: Installation
next: /docs/quickstart-building-a-package/
nextLabel: Building a Package
contributeLink: https://example.com/
---

Before we can truly kick off we need to configure an Escape Inventory. The
Escape Inventory is used by Escape to store, retrieve and answer questions
about packages.  By default Escape is configured to use the central Escape
Inventory hosted by Ankyra, but at the moment of writing this does not provide
write access to members of the public. (Watch this space)

Setting up your own Inventory is easy though. As with Escape we can either use 
the pre-built binaries or run a Docker container. 

<div class='docling'>
The Inventory can be configured to store packages on Google Cloud Storage
instead of disk and use a Postgres database instead of `ql`, but right now
we're sticking to the default values. For more information on how to configure
the Inventory see the <a href='/docs/escape-inventory/'>Escape Inventory
Usage</a> page.
</div>

## Linux

```
curl -O https://storage.googleapis.com/escape-releases-eu/escape-inventory/{{ version }}/escape-inventory-v{{version}}-linux-amd64.tgz
tar -xvzf escape-v{{version}}-linux-amd64.tgz
sudo mv escape-inventory /usr/bin/escape-inventory
sudo mkdir -p /var/lib/escape/releases/ && chown -R `whoami` /var/lib/escape/
escape-inventory
```

## MacOS

```
curl -O https://storage.googleapis.com/escape-releases-eu/escape-inventory/{{ version }}/escape-inventory-v{{version}}-darwin-amd64.tgz
tar -xvzf escape-v{{version}}-darwin-amd64.tgz
sudo mv escape-inventory /usr/bin/escape-inventory
sudo mkdir -p /var/lib/escape/releases/ && chown -R `whoami` /var/lib/escape/
escape-inventory
```

## Docker

We also provide a Docker image on the public [Docker
Hub](https://hub.docker.com/r/ankyra/escape-inventory/):

```
docker run -P -it ankyra/escape-inventory:v{{version}}
```

We can use a [Docker
Volume](https://docs.docker.com/engine/admin/volumes/volumes/) for the
`/var/lib/escape/` directory to make sure our data is persisted:

```
docker create -v /var/lib/escape/ --name inventory ankyra/escape-inventory:v{{version}}
docker run --volumes-from inventory -P -it ankyra/escape-inventory:v{{version}}
```

# Configuring Escape

We have our own instance of the Inventory running now, but we need to tell
Escape to use it, because our default configuration profile is set up to go the
central Ankyra repository:

```
escape config profile
```

We can "login" to our local instance and create and activate a new profile 
using the [escape login](/docs/escape_login/) command:

```
escape login --url http://localhost:7770/ --target-profile local
```

If everything is working correctly we should now be able to query our empty repository 
from the command line using:

```
escape inventory query --json
```

Beautiful, let's build something!
