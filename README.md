# pushnix

A cli tool to push your NixOS configuration repository to a remote host using git and then run `nixos-rebuild` on the host using ssh.

## Setup

1. Make sure you have the following dependencies on the remote host.
  - `git-receive-pack` in path (included in `git` in nixpkgs).
  - A running SSH server and a remote user you can connect to.
2. On the remote NixOS host create a directory where you want the NixOS configuration to be.
3. Inside the directory setup the repository.
  - `git init`
  - `git config receive.denyCurrentBranch updateInstead` ([info](https://git-scm.com/docs/git-config#Documentation/git-config.txt-receivedenyCurrentBranch))
4. In your local NixOS configuration git repository run `git remote add <remote_name> <user>@<host>:<path>`.
  - `<remote_name>` - Whatever name you want this remote host to be called when interacting with it using `git` or `pushnix`.
  - `<user>` - The remote user you can SSH to.
  - `<host>` - DNS or IP of remote host you can SSH to.
  - `<path>` - Path to git repository on the remote host.

Now you've set it up so that you can run `git push <remote_name>` and the git repository on the remote side will be updated (as long as it is checked out on the same branch as you are pushing).

## Usage

Let's say you created a remote called `my_server` in your NixOS configuration git repository.

Now you can run `pushnix deploy my_server` to push the configuration and rebuild the NixOS configuration on the remote host.

Everything after `--` will be passed onto the `nixos-rebuild` command.

```
pushnix deploy my_server -- -I nixos-config=/home/user/config/machines/my_server/configuration.nix
```

## Why?

[NixOps](https://github.com/NixOS/nixops) and [morph](https://github.com/DBCDK/morph) are much more featureful options that didn't fit one my use cases.

They build the derivations locally, copy them over to the remote host and run the activation script there. This doesn't make much sense when being run on a laptop where everything is built locally instead on the server itself.

Also, I want one of my servers to occationally receive updated configuration and then just use `system.autoUpgrade` option to keep it updated. This requires the configuration to live on the remote host.

This also simplifies testing a configuration on a remote host, instead of pushing to a central repository, SSH into the host, pull from the central repository and finally run `nixos-rebuild`.
