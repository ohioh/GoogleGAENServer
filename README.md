Infomration by Tjark @all:

Google open-sourced its Exposure Notification reference server (https://github.com/google/exposure-notifications-server), h/t @richard.taylor. It’s responsible for the following functions:
• Accepting the temporary exposure keys of affected users from mobile devices.
• Validating the temporary exposure keys using the device attestation API.
• Storing the temporary exposure keys in a database.
• Periodically generating incremental files that will be downloaded by mobile devices to perform the key matching algorithm on the mobile device.
• Sending a public key to devices, and digitally signing the incremental files with a private key.
• Periodically deleting old temporary exposure keys. After 14 days, or configured time period, the exposure keys can no longer be matched to a device.
What this means for us; other observations:
• To speed up our time to market and use security/scaling best-practices, we can use it as is.
• Interoperability is handled using a federation model where servers sync the temporary exposure keys between each other periodically. (This means a US server might contain temporary exposure keys from an MX server if the former is set up that way).
• Deployment of services can be made on the Google Cloud Platform entirely.
• Verification of published temporary exposure keys is done once when the mobile device uploads them to the publish service. Alongside region, and SafetyNet / DeviceCheck verification, there are hints that the service can be modified to verify the verificationPayload received from the verification authority (PHA), which is also attached to the report.
• If verification was successful then the keys will be published and made available for apps to download. During server federation sync, servers will receive a VerificationAuthorityName that did the verification but not the verificationPayload.
I’ve been discussing with @Madhava and @Josh (California) how we can use it, and we’re interested to read your ideas.
23:28 Uhr
did you see this in the TCN slack















-------------------------------------------------------------------------------------------------------------------------------------
# Exposure Notification Reference Server

[https://www.google.com/covid19/exposurenotifications/](https://www.google.com/covid19/exposurenotifications/)

In our continued effort to help governments and health authorities during the COVID-19 pandemic, we have authored an open source reference implementation of an Exposure Notifications server. The server reference design implements the Exposure Notifications API and provides reference code for working with both Android and iOS apps that are built by public health authorities. The reference server source code is available on GitHub and can be deployed on any modern infrastructure or cloud provider selected by a public health authority. 

The reference server implementation accepts, validates, and stores temporary exposure keys from verified mobile devices. It also periodically generates and signs incremental files that will later be downloaded by clients to perform the on-device key matching algorithm. Our hope is by making this privacy-preserving server implementation available to health authorities we can enable their developers to use the open source code to get started quickly. 

This repository contains a reference implementation of an exposure notification
server for use in a mobile exposure notification system.

## Overview

The server is responsible for the following functions:

* Accepting the temporary exposure keys of affected users from mobile devices.

* Validating the temporary exposure keys using the device attestation API.

* Storing the temporary exposure keys in a database.

* Periodically generating incremental files that will be downloaded by mobile
  devices to perform the key matching algorithm on the mobile device.

* Sending a public key to devices, and digitally signing the incremental files with
  a private key.

* Periodically deleting old temporary exposure keys. After 14 days, or
  configured time period, the exposure keys can no longer be matched to a device.

## Tutorials and reference documentation

You can read tutorials on deploying and using the reference Exposure Notification
Server here:

* [Overview](docs/index.md)
* [Deployment Guide](docs/deploying.md)
* [Reference Documentation](https://pkg.go.dev/mod/github.com/google/exposure-notifications-server)
* [Server Functional Requirements](docs/server_functional_requirements.md)
* [Server Deployment Options](docs/server_deployment_options.md)

# Issues and Questions

You can open a
[GitHub Issue](https://github.com/google/exposure-notifications-server/issues/new).
Please be sure to include as much detail as you can to help aid in addressing
your concern. If you wish to reach out privately, you can send an e-mail
exposure-notifications-feedback@google.com.

## Contributing to this project

Contributions to this project are welcomed. For more information about
contributing to this project, see the [contribution guidelines](CONTRIBUTING.md).

