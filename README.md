<!-- Header block for project -->
<hr>

<div align="center">

![logo](https://user-images.githubusercontent.com/3129134/163247802-ad001e00-15a6-4d6e-9824-2106cb022dd7.png)

<h1 align="center">Unity CS Terraform Transformer</h1>
<!-- ☝️ Replace with your repo name ☝️ -->

</div>

<pre align="center">Transform and valitate Terraform scripts to run under Unity CS</pre>
<!-- ☝️ Replace with a single sentence describing the purpose of your repo / proj ☝️ -->

<!-- Header block for project -->

![Terraform](https://img.shields.io/badge/Terraform-Could%20be%20worse-brightgreen)
![Golang](https://img.shields.io/badge/Golang-hacked%20together-yellow) [![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](code_of_conduct.md)
<!-- ☝️ Add badges via: https://shields.io e.g. ![](https://img.shields.io/github/your_chosen_action/your_org/your_repo) ☝️ -->


The project is designed to allow us to parse Terraform scripts written by developers and ensure they will run inside Unity CS. This includes
adding networking information, ensuring mandatory tags are applied and more. Because the developers wont know the network topology of the AWS
account the project is being deployed into, they need to be structred in a way that allows us to add this information in at deploy time.

This project parses the scripts, looks up blocks and attributes and makes the required adjustments, adding, removing or appending blocks with 
the correct information.
<!-- ☝️ Replace with a more detailed description of your repository, including why it was made and whom its intended for.  ☝️ -->

[Docs/Wiki](https://github.com/unity-sds/unity-cs/wiki/Terraform-Transformer-Component)
[Issue Tracker](https://github.com/unity-sds/unity-cs-terraform-transformer/issues)
<!-- example links>
[Website](INSERT WEBSITE LINK HERE) | [Docs/Wiki](INSERT DOCS/WIKI SITE LINK HERE) | [Discussion Board](INSERT DISCUSSION BOARD LINK HERE) | [Issue Tracker](INSERT ISSUE TRACKER LINK HERE)
-->

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

This guide provides a quick way to get started with our project. Please see our [docs]([INSERT LINK TO DOCS SITE / WIKI HERE]) for a more comprehensive overview.

### Requirements

* A PC
* Some terraform scripts
* [The Unity CS Terraform Transformer Binary](https://github.com/orgs/unity-sds/packages?repo_name=unity-cs-terraform-transformer)

<!-- ☝️ Replace with a numbered list of your requirements, including hardware if applicable ☝️ -->

### Setup Instructions

1. Grab binary from Packages page
2. Open a terminal

<!-- ☝️ Replace with a numbered list of how to set up your software prior to running ☝️ -->

### Run Instructions

1. ```./terraform-transformer parser -t tom -p test --creator tom@spicule.co.uk --venue dev --servicearea cs --capability hysds --capversion 0.0.1 --release G1.0.0 --component python --securityplan 644 --exposed true --experimental false --userfacing true --critinfra 2 --project testproj
```

<!-- ☝️ Replace with a numbered list of your run instructions, including expected results ☝️ -->

### Usage Examples

* [INSERT LIST OF COMMON USAGE EXAMPLES HERE, WITH OPTIONAL SCREENSHOTS]

<!-- ☝️ Replace with a list of your usage examples, including screenshots if possible, and link to external documentation for details ☝️ -->

### Build Instructions (if applicable)

1. [INSERT STEP-BY-STEP BUILD INSTRUCTIONS HERE, WITH OPTIONAL SCREENSHOTS]

<!-- ☝️ Replace with a numbered list of your build instructions, including expected results / outputs with optional screenshots ☝️ -->

### Test Instructions (if applicable)

1. [INSERT STEP-BY-STEP TEST INSTRUCTIONS HERE, WITH OPTIONAL SCREENSHOTS]

<!-- ☝️ Replace with a numbered list of your test instructions, including expected results / outputs with optional screenshots ☝️ -->

## Changelog

See our [CHANGELOG.md](CHANGELOG.md) for a history of our changes.

See our [releases page]([INSERT LINK TO YOUR RELEASES PAGE]) for our key versioned releases.

<!-- ☝️ Replace with links to your changelog and releases page ☝️ -->

## Frequently Asked Questions (FAQ)

[INSERT LINK TO FAQ PAGE OR PROVIDE FAQ INLINE HERE]
<!-- example link to FAQ PAGE>
Questions about our project? Please see our: [FAQ]([INSERT LINK TO FAQ / DISCUSSION BOARD])
-->

<!-- example FAQ inline format>
1. Question 1
   - Answer to question 1
2. Question 2
   - Answer to question 2
-->

<!-- example FAQ inline with no questions yet>
No questions yet. Propose a question to be added here by reaching out to our contributors! See support section below.
-->

<!-- ☝️ Replace with a list of frequently asked questions from your project, or post a link to your FAQ on a discussion board ☝️ -->

## Contributing

Interested in contributing to our project? Please see our: [CONTRIBUTING.md](CONTRIBUTING.md)

## License

See our: [LICENSE](LICENSE)

## Support

[INSERT CONTACT INFORMATION OR PROFILE LINKS TO MAINTAINERS AMONG COMMITTER LIST]

<!-- example list of contacts>
Key points of contact are: [@github-user-1](link to github profile) [@github-user-2](link to github profile)
-->

<!-- ☝️ Replace with the key individuals who should be contacted for questions ☝️ -->




