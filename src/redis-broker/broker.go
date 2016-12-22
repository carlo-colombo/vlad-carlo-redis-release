package broker

import "github.com/pivotal-cf/brokerapi"
import "context"

func New() brokerapi.ServiceBroker {
	return &broker{}
}

type broker struct{}

// Services returns a []brokerapi.Service here, describing your service(s) and plan(s)
func (*broker) Services(context context.Context) []brokerapi.Service {
	freePlan := true
	redisService := brokerapi.Service{
		ID:          "f930e807-4573-41dc-8f2a-4de9869ea10b",
		Name:        "redis-service",
		Description: "A Redis Service given to you by Vlad & Carlo",
		Bindable:    true,
		Plans: []brokerapi.ServicePlan{
			brokerapi.ServicePlan{
				ID:          "699d339e-7b5b-415d-a766-5c06f30c0aa0",
				Name:        "standard",
				Description: "uhhhh, not another description",
				Free:        &freePlan,
			}},
	}
	return []brokerapi.Service{redisService}
}

// Provision a new instance here. If async is allowed, the broker can still
// chose to provision the instance synchronously.
func (*broker) Provision(context context.Context, instanceID string, details brokerapi.ProvisionDetails, asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error) {
	return brokerapi.ProvisionedServiceSpec{
		IsAsync:       false,
		DashboardURL:  "random-dashboard-url",
		OperationData: "random-operation-data",
	}, nil
}

// Deprovision a new instance here. If async is allowed, the broker can still
// chose to deprovision the instance synchronously, hence the first return value.
func (*broker) Deprovision(context context.Context, instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {
	return brokerapi.DeprovisionServiceSpec{}, nil
}

// Bind to instances here
// Return a binding which contains a credentials object that can be marshalled to JSON,
// and (optionally) a syslog drain URL.
func (*broker) Bind(context context.Context, instanceID, bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error) {
	return brokerapi.Binding{
		Credentials:     "random-credentials",
		SyslogDrainURL:  "random-syslog-url",
		RouteServiceURL: "random-route-service-url",
		VolumeMounts:    []brokerapi.VolumeMount{},
	}, nil
}

// Unbind from instances here
func (*broker) Unbind(context context.Context, instanceID, bindingID string, details brokerapi.UnbindDetails) error {
	return nil
}

// Update instance here
func (*broker) Update(context context.Context, instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error) {
	return brokerapi.UpdateServiceSpec{}, nil
}

// LastOperation
// If the broker provisions asynchronously, the Cloud Controller will poll this endpoint
// for the status of the provisioning operation.
// This also applies to deprovisioning (work in progress).
func (*broker) LastOperation(context context.Context, instanceID, operationData string) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{}, nil
}
