---
title: "Installing Inventory"
slug: inventory-installation
type: "docs"
toc: true

back: /docs/escape-installation/
backLabel: Installation
next: /docs/quickstart-configure-escape/
nextLabel: Configuring Escape
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
curl -O https://storage.googleapis.com/escape-releases-eu/escape-inventory/0.12.13/escape-inventory-v0.12.13-linux-amd64.tgz
tar -xvzf escape-v0.12.13-linux-amd64.tgz
sudo mv escape-inventory /usr/bin/escape-inventory
sudo mkdir -p /var/lib/escape/releases/ && chown -R `whoami` /var/lib/escape/
escape-inventory
```

## MacOS

```
curl -O https://storage.googleapis.com/escape-releases-eu/escape-inventory/0.12.13/escape-inventory-v0.12.13-darwin-amd64.tgz
tar -xvzf escape-v0.12.13-darwin-amd64.tgz
sudo mv escape-inventory /usr/bin/escape-inventory
sudo mkdir -p /var/lib/escape/releases/ && chown -R `whoami` /var/lib/escape/
escape-inventory
```

## Docker

We also provide a Docker image on the public [Docker
Hub](https://hub.docker.com/r/ankyra/escape-inventory/):

```
docker run -P -it ankyra/escape-inventory:v0.12.13
```

We can use a [Docker
Volume](https://docs.docker.com/engine/admin/volumes/volumes/) for the
`/var/lib/escape/` directory to make sure our data is persisted:

```
docker create -v /var/lib/escape/ --name inventory ankyra/escape-inventory:v0.12.13
docker run --volumes-from inventory -P -it ankyra/escape-inventory:v0.12.13
```