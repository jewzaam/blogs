package test

import (
    "time"

	configv1 "github.com/openshift/api/config/v1"
	"github.com/openshift/library-go/pkg/controller/controllercmd"
	"github.com/openshift/library-go/pkg/operator/status"

	configclient "github.com/openshift/client-go/config/clientset/versioned"
	configinformers "github.com/openshift/client-go/config/informers/externalversions"
    operatorversionedclient "github.com/openshift/client-go/operator/clientset/versioned"
    
	operatorv1client "github.com/openshift/client-go/operator/clientset/versioned/typed/operator/v1"
	operatorv1informers "github.com/openshift/client-go/operator/informers/externalversions"
)
const (
    // operator's name
    operatorName        string = "my-operator-name"
    // operator's namespace
    operatorNamespace   string = "my-operator-namespace"
    // operator's version
    operatorVersion     string = "0.0.1"
    // clusteroperator's name
    clusterOperatorName string = "my-operator"
)

type OperatorClient struct {
	Informers operatorv1informers.SharedInformerFactory
	Client    operatorv1client.ConsolesGetter
}

func RunOperator(ctx *controllercmd.ControllerContext) error {
    // resync frequency for ClusterOperator
    operatorResync := 10 * time.Minute

    versionGetter := status.NewVersionGetter()
    versionGetter.SetVersion("operator", operatorVersion)

    configClient, err := configclient.NewForConfig(ctx.KubeConfig)
    if err != nil {
        return err
    }

    operatorConfigClient, err := operatorversionedclient.NewForConfig(ctx.KubeConfig)
    if err != nil {
        return err
    }

    operatorConfigInformers := configinformers.NewSharedInformerFactoryWithOptions(
        configClient,
        operatorResync,
    )

	operatorClient := &OperatorClient{
		Informers: operatorConfigInformers,
		Client:    operatorConfigClient.OperatorV1(),
	}

    relatedObjects := []configv1.ObjectReference{        
        {Resource: "namespaces", Name: operatorNamespace},
    }

    clusterOperatorStatus := status.NewClusterOperatorStatusController(
        clusterOperatorName,
        relatedObjects,
        configClient.ConfigV1(),
        operatorConfigInformers.Config().V1().ClusterOperators(),
        operatorClient,
        versionGetter,
        ctx.EventRecorder,
    )
}

