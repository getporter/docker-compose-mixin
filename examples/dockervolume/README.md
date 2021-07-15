This Porter bundle explores some options for getting data into/out of a Docker Compose application.

## Bundle components

- The `docker-compose.yml` consists of one single app container which uses a Docker volume for its data ("myvol")
- The `helpers.sh` script:
  - Creates the "myvol" Docker volume (which will be created using the host's Docker daemon).
  - Copies in both a file baked into the invocation image ("config.toml") and a file generated at runtime from a Porter parameter ("echo.txt").
  - It uses a helper docker container to perform this data volume setup.  This is necessary as all shared volume data must flow through Docker.  A benefit of using a [Docker Volume](https://docs.docker.com/storage/volumes/) here is it is agnostic towards the filesystem layout of the host (as opposed to [bind mounts](https://docs.docker.com/storage/bind-mounts/)), so it can be used on any host with a running Docker daemon, regardless of OS.
- The `porter.yaml` manifest ties this all together, e.g. in the `install` step which:
  1. Creates the volume and populates it with data via a helper container
  2. Brings up the Docker Compose app
  3. Prints app logs via a logreader container


## Sample flow

```
 $ porter install --allow-docker-host-access
installing dockervolume...
executing install action from dockervolume (installation: dockervolume)
Create Docker Volume
myvol
Docker Compose Up
Creating app_awk_1 ... done
Print app logs
appconfig = true
Hello from Porter!
execution completed successfully!

 $ porter upgrade --allow-docker-host-access --param echo_text="Howdy from Porter!"
upgrading dockervolume...
executing upgrade action from dockervolume (installation: dockervolume)
Update Docker Volume
local     myvol
Docker Compose Up
Starting app_awk_1 ... done
Print app logs
appconfig = true
Howdy from Porter!
execution completed successfully!

 $ porter uninstall --allow-docker-host-access
uninstalling dockervolume...
executing uninstall action from dockervolume (installation: dockervolume)
Docker Compose Down
Removing app_awk_1 ... done
Removing network app_default
Remove Docker Volume
myvol
execution completed successfully!
```