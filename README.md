GOPROXY=direct go get -u github.com/unity-sds/unity-cs-manager@main

![logo](https://user-images.githubusercontent.com/3129134/163247802-ad001e00-15a6-4d6e-9824-2106cb022dd7.png)

# Unity CS Terraform Transformer

## Transform and valitate Terraform scripts to run under Unity CS

<!-- Header block for project -->

![Terraform](https://img.shields.io/badge/Terraform-Could%20be%20worse-brightgreen)
![Golang](https://img.shields.io/badge/Golang-hacked%20together-yellow)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](code_of_conduct.md)

The project is designed to allow us to parse Terraform scripts written by developers
and ensure they will run inside Unity CS. This includes adding networking information,
ensuring mandatory tags are applied and more. Because the developers wont know the
network topology of the AWS account the project is being deployed into, they need
to be structred in a way that allows us to add this information in at deploy time.

This project parses the scripts, looks up blocks and attributes and makes the
required adjustments, adding, removing or appending blocks with the correct information.

[Docs/Wiki](https://github.com/unity-sds/unity-cs/wiki/Terraform-Transformer-Component)
[Issue Tracker](https://github.com/unity-sds/unity-cs-terraform-transformer/issues)

## Features

* Support for mandatory Unity Tags
* Support for VPC and Subnet injection

<!-- ☝️ Replace with a bullet-point list of your features ☝️ -->

## Contents

* [Quick Start](#quick-start)
* [Changelog](#changelog)
* [FAQ](#frequently-asked-questions-faq)
* [Contributing Guide](#contributing)
* [License](#license)
* [Support](#support)

## Quick Start

### Requirements

* A PC
* Some terraform scripts
* [The Unity CS Terraform Transformer Binary](https://github.com/orgs/unity-sds/packages?repo_name=unity-cs-terraform-transformer)

### Setup Instructions

1. Grab binary from Packages page.
1. Open a terminal.

#### Compiling protobuf

protoc --proto_path=unity-management-console/protobuf --go_out=internal/marketplace/ --go_opt=paths=source_relative config.proto extensions.proto marketplace.proto

### Run Instructions

```shell
2. $> ./terraform-transformer parser -t tom -p test --creator tom@spicule.co.uk
--venue dev --servicearea cs --capability hysds --capversion 0.0.1
--release G1.0.0 --component python --securityplan 644 --exposed true
--experimental false --userfacing true --critinfra 2 --project testproj
```

## Changelog

See our [CHANGELOG.md](CHANGELOG.md) for a history of our changes.

<!-- ☝️ Replace with links to your changelog and releases page ☝️ -->

## Contributing

Interested in contributing to our project? Please see our: [CONTRIBUTING.md](CONTRIBUTING.md)

## License

See our: [LICENSE](LICENSE)

## Support

Key points of contact are:

[@galenatjpl](https://github.com/galenatjpl)
<!-- ☝️ Replace with the key individuals who should be contacted for questions ☝️ -->
