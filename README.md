# AnsibleRoleReadme

## Features

This go application generates Ansible Role README.md file to "roles" Directories.
Generated README.md include role name(equal folder name),variables ( and default value).

## Usage

1.Build go script.
2.Put generated files to ansible's root directory(playbook directory).

```
/
  - Playbook.yml
  - {Generated Binary} <-- here
  - roles
```

3.Run binary in playbook directory.

example readme.md

```
#mariadb

## abstruct

## variables

|variable name|default value|
|---|---|
|mariadb_port|3306|
```

