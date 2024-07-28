// newServiceForPod creates a new Kubernetes Service for a given RoarApp instance.
// This service is designed to expose the RoarApp's pods on a specific node port,
// allowing external access to the application.
//
// Parameters:
// - cr: A pointer to the RoarApp custom resource instance. This parameter is crucial
//       as it contains the metadata and specifications of the RoarApp instance, such
//       as its name and namespace, which are used to configure the service accordingly.
//
// Returns:
// - A pointer to a corev1.Service object. This object represents the newly created service
//   configured to expose the RoarApp instance. The service is not yet applied to the Kubernetes
//   cluster; it merely represents the desired state.
//
// The function performs the following steps:
// 1. Converts the global variable `nextPort` to a string. This port number is used to
//    uniquely identify the service and is part of the service's name to ensure uniqueness.
// 2. Prepares a set of labels that will be applied to the service.These labels are used
//    to select the pods that the service will expose. In this case, the label is based on
//    the RoarApp's name, ensuring that the service only selects pods belonging to the specific
//    RoarApp instance.
// 3. Constructs a corev1.Service object with the necessary metadata and specifications.
//    This includes setting the service's name (which includes the RoarApp's name and the `nextPort`),
//    namespace (same as the RoarApp's namespace), selector (to select the correct pods),
//    ports (defining the protocol, service port, target port on the pods, and the node port for
//    external access), and type (NodePort, to allow external access)
//
// Note: The `nextPort` variable must be managed externally to ensure it is unique and within
//       the valid range for node ports in the Kubernetes cluster. This function assumes `nextPort`
//       has been appropriately incremented and managed to avoid conflicts with other services.
func newServiceForPod(cr *roarappv1alpha1.RoarApp) *corev1.Service {

    strPort := strconv.Itoa(nextPort) // Convert the next available port number to a string.
    labels := map[string]string{
        "app": cr.Name, // Define labels to be applied to the service, using the app's name.
    }

    // Create and return a pointer to the Service object.
    return &corev1.Service{
        ObjectMeta: metav1.ObjectMeta{
            Name:      cr.Name + "-service-" + strPort, // Name of the service includes the app name and port.
            Namespace: cr.Namespace,                    // Service is created in the same namespace as the RoarApp.
        },
        Spec: corev1.ServiceSpec{
            Selector: labels, // Selector matches the labels of the pods to target.
            Ports: []corev1.ServicePort{{
                Protocol:   corev1.ProtocolTCP,          // The service supports TCP protocol.
                Port:       8089,                         // The service is exposed on port 8089.
                TargetPort: intstr.FromInt(8080),        // The target port on the pod is 8080.
                NodePort:   int32(nextPort),              // Expose the service outside the cluster using a node port.
            }},
            Type: corev1.ServiceTypeNodePort, // The service type is NodePort, making it accessible externally.
        },
    }
}
