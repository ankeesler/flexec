# flexec

This tool makes it easier to `fly exec` tasks from anywhere on your machine.

### IMPORTANT NOTE

This only works on CredHub workstations currently :)

## To Use:

```bash
$ go install github.com/ankeesler/flexec
$ flexec lint
flexec: reading task path /Users/pivotal/workspace/credhub-ci/tasks/lint/task.yml
flexec: resolved input credhub-ci to path /Users/pivotal/workspace/credhub-ci
enter value for input credhub-src (default '/Users/pivotal/workspace/credhub-release/src/credhub'):
flexec: resolved input credhub-src to path /Users/pivotal/workspace/credhub-release/src/credhub
flexec: running command:
  HOME=/Users/pivotal
  fly
  --target
  credhub
  execute
  --config
  /Users/pivotal/workspace/credhub-ci/tasks/lint/task.yml
  --input=credhub-ci=/Users/pivotal/workspace/credhub-ci
  --input=credhub-src=/Users/pivotal/workspace/credhub-release/src/credhub

executing build 81827 at https://credhub.ci.cf-app.com/builds/81827
initializing
...
```
