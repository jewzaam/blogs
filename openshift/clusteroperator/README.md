# Reporting state of Operators

In OpenShift v4 the Operator is king.  This blog will walk you through configuring resources to report the status of your operator, providing a valuable tool for quickly triaging the state of a cluster that's consistent with the OpenShift way of doing things.

You'll learn how to use existing librarires to easily create and manage a ClusterOperator custom resource for your application

## What is an Operator?


## What is a ClusterOperator?


# How to Manage ClusterOperators

This blog walks through each discrete step to bring your ClusterOperator to life.  You will walk through each of the components used in the library-go package, how to configure them in your code, and how they work together.

In summary you will:

- Create a custom VersionGetter to report the version of your Operator
- ConfigClient
- ConfigInformer

Each section gives the snippet of what must be configured.  A final section brings it all together and shows how to create a StatusSyncer.

## Imports

```go
import (
	configv1 "github.com/openshift/api/config/v1"
	"github.com/openshift/library-go/pkg/operator/status"

	configclient "github.com/openshift/client-go/config/clientset/versioned"
	configinformers "github.com/openshift/client-go/config/informers/externalversions"
    operatorversionedclient "github.com/openshift/client-go/operator/clientset/versioned"
    operatorinformers "github.com/openshift/client-go/operator/informers/externalversions"
    
	operatorv1client "github.com/openshift/client-go/operator/clientset/versioned/typed/operator/v1"
	operatorv1informers "github.com/openshift/client-go/operator/informers/externalversions"
)
```

## Constants

For this example we're going to assume some constants are defined in a single struct.  These values may change between versions of the operator, but they're immutable within a single version.

```go
const (
    // operator's name
    operatorName        string = "my-operator-name"
    // operator's namespace
    operatorNamespace   string = "my-operator-namespace"
    // operator's version
    operatorVersion     string = "0.0.1"
    // resync frequency for ClusterOperator
    operatorResync      int    = 10 * time.Minute
    // clusteroperator's name
    clusterOperatorName string = "my-operator"
)
```

## Verson Management

Every operator should report a version.  The version reported on the ClusterOperator needs to reflect the current operator version, so must change over time.  The library-go package provides an interface called VersionGetter that enables you to easily manage this.

Create a VersionGetter object for your controller to use and set the current version:

```go
    versionGetter := status.NewVersionGetter()
    versionGetter.SetVersion("operator", operatorVersion)
```

## Config Client

```go
func RunOperator(ctx *controllercmd.ControllerContext) error {
    ...
    configClient, err := configclient.NewForConfig(ctx.KubeConfig)
    if err != nil {
        return err
    }
    ...
}
```

## Operator Config Client

```go
    operatorConfigClient, err := operatorversionedclient.NewForConfig(ctx.KubeConfig)
    if err != nil {
        return err
    }
```

## Informer

```go
    operatorConfigInformers := configinformers.NewSharedInformerFactoryWithOptions(
        configClient,
        operatorResync,
    )
```

## Operator Client

```go
type OperatorClient struct {
	Informers operatorv1informers.SharedInformerFactory
	Client    operatorv1client.ConsolesGetter
}
...
	operatorClient := &OperatorClient{
		Informers: operatorConfigInformers,
		Client:    operatorConfigClient.OperatorV1(),
	}
```

## Related Objects

Define an array of objects related to the Operator.  Used to...

TODO add what it's used for!

```go
    relatedObjects := []configv1.ObjectReference{        
        {Resource: "namespaces", Name: operatorNamespace},
    }
```

## ClusterOperator Status Controller

```go
func RunOperator(ctx *controllercmd.ControllerContext) error {
    ...
    clusterOperatorStatus := status.NewClusterOperatorStatusController(
        clusterOperatorName,
        relatedObjects,
        configClient.ConfigV1(),
        operatorConfigInformers.Config().V1().ClusterOperators(),
        operatorClient,
        versionGetter,
        ctx.EventRecorder,
    )
    ...
}
```

## Final Solution

how to view the outcome of the work

# Conclusion

what you did
why it's important
what's next