Simbal
===

Simbal is the Simple Backup Launcher tool.
- Backup dirs(files) to specified dirs with tar command.
- Archive backup dirs(files) with gzip(tar -z option).
- Generational management of backup files.

### Usage

Simple tar archive.

```sh
$ simbal backup -s /tmp/test -d /tmp/backup
```

If you want backup more than 8days, set argument `-n` and backup days(default 7days backup).

```sh
$ simbal backup -s /tmp/test -d /tmp/backup -n 14
```
