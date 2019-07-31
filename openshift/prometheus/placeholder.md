# Deploy Custom Prometheus Metrics and Alerts in OCP v4

What is prometheus..
Installed by CVO
why metrics
why alerts

## What is Prometheus?

Prometheus is an open-source systems monitoring and alerting toolkit.  It provides multi-dimentional time-series metrics and the ability to trigger alerts on those metrics.  It is at the core of OpenShift Container Platform v4.x's monitoring and alerting capability.

## What is CVO?

The Cluster Version Operator, or CVO for short, is the component in OCP v4 clusters that manages the installation and configuration of core components of OpenShift.  Once installation is complete, CVO is responsible for ensuring the cluster resources is manages remain in a state it expects.

This is important when we look at how to introduce new recording rules and alerts into Prometheus.  More on that later!

TODO: look for some language from product to base this on

## Overview



Application w/ /metrics in correct format
Exposed to Prometheus
Custom rules and alerts

## Application

namespace label

## Expose Metrics

Service

## Make available to prometheus

ServiceMonitor
Role
RoleBinding

## Custom Rules

PrometheusRules

## View it

Prometheus UI

# Conclusion

Easy to extend the cluster.
SRE exporters and operators.
Ship to any prometheus, external too.