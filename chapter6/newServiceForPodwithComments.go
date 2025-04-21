// newServiceForPod creates a new Kubernetes Service resource for a given RoarApp custom resource.
// The Service is configured as a NodePort service, exposing a specific port for external access.
//
// Parameters:
//   - cr: A pointer to a RoarApp custom resource instance.
//
// Returns:
//   - A pointer to a corev1.Service object configured with the specified labels, ports, and namespace.
func newServiceForPod(cr *roarappv1alpha1.RoarApp) *corev1.Service {

    // Convert the next available port to a string for use in the service name.
    strPort := strconv.Itoa(nextPort)

    // Define labels to associate the service with the corresponding RoarApp instance.
    labels := map[string]string{
        "app": cr.Name,
    }

    // Create and return a new Service object.
    return &corev1.Service{
        ObjectMeta: metav1.ObjectMeta{
            // The service name includes the RoarApp name and the port to ensure uniqueness.
            Name:      cr.Name + "-service-" + strPort,
            Namespace: cr.Namespace,
        },
        Spec: corev1.ServiceSpec{
            // Use the labels to select the pods associated with this service.
            Selector: labels,
            Ports: []corev1.ServicePort{{
                Protocol:   corev1.ProtocolTCP,       // Use TCP as the protocol.
                Port:       8089,                    // The port exposed by the service.
                TargetPort: intstr.FromInt(8080),    // The target port on the pods.
                NodePort:   int32(nextPort),         // The NodePort for external access.
            }},
            Type: corev1.ServiceTypeNodePort, // Configure the service as a NodePort type.
        },
    }
}
