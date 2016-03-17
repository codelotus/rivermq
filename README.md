[![Build Status](https://travis-ci.org/codelotus/rivermq.svg?branch=master)](https://travis-ci.org/codelotus/rivermq) [![codecov.io](https://codecov.io/github/codelotus/rivermq/coverage.svg?branch=master)](https://codecov.io/github/codelotus/rivermq?branch=master)

RiverMQ
========

WebHook based messaging system

RiverMQ will provide a WebHook based asynchronous messaging solution for distributed applications.

Clients will register subscriptions to a message "type" with a RiverMQ Server instance via a HTTP POST.  This subscription will include a "callback url" which the client expects to receive messages on.  When a separate client sends a message of a matching "type" to a RiverMQ server instance, that message will be sent, via a HTTP POST, to the "callback url" of all subscribed clients.  Based on the response code received by RiverMQ, message redelivery attempts will be retried for up to 30 mintes.  After 30 minutes RiverMQ will cease attempts to send the message to the client.

RiverMQ is inspired by [subpub](https://github.com/PearsonEducation/subpub)


Goals
-----

1. Provide a single executable with minimal configuration
1. Simple horizontal scaling
1. Automatic discovery and configuration of RiverMQ nodes via [Raft](https://raft.github.io/)
1. Registration with [Consul](https://www.consul.io/) and/or [Etcd](https://coreos.com/etcd/) for discovery by clients
1. Securred communication between RiverMQ nodes with [ZeroMQ](http://zeromq.org/)
1. Message storage and flowthrough visualization with [InfluxDB](https://influxdata.com/)
1. Message storage and flowthrough visualization wtih [MongoDB](https://www.mongodb.org/), [Angular.js](https://angularjs.org/), and some charting library.


Development
-----------
Clone the repo and inspect the Makefile for running tests and building.

Integration tests require an instance of [InfluxDB](https://influxdata.com/) running on the localhost and bound to default ports.  For convenience, a docker-compose file is included in the root of the project which will launch an [InfluxDB](https://influxdata.com/) and [Consul](https://www.consul.io/) container.  Simply run the following to launch both containers.
```bash
docker-compose -f docker-compose-dev.yml up
```

Test coverage result file concatination is done using [gover](https://github.com/modocache/gover).  This is a required dependency which must be installed using the following:
```bash
go get github.com/modocache/gover
```

Developed using [Atom](https://atom.io/) [configured for Go development](http://marcio.io/2015/07/supercharging-atom-editor-for-go-development).
