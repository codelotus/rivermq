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


Development
-----------
Tests are written with [Ginkgo](http://onsi.github.io/ginkgo/) and [Gomega](http://onsi.github.io/gomega/).  To install these tools run the following:
```
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega
```
